package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/codeandcodes/subs/protos"
	square "github.com/square/square-connect-go-sdk/swagger"
)

type SquareLocationService struct {
	Client *square.APIClient
}

const FIXED_ONLYSUBS_LOCATION_NAME = "onlysubs payments dev"

// Takes a SubscriptionSetupRequest
// For each payer, will find or create a customer record in Square API

/*
	"location": {
	    "id": "3Z4V4WHQK64X9",
	    "name": "Midtown",
	    "address": {
	      "address_line_1": "1234 Peachtree St. NE",
	      "locality": "Atlanta",
	      "administrative_district_level_1": "GA",
	      "postal_code": "30309"
	    },
	    "timezone": "America/New_York",
	    "capabilities": [
	      "CREDIT_CARD_PROCESSING"
	    ],
	    "status": "ACTIVE",
	    "created_at": "2022-02-19T17:58:25Z",
	    "merchant_id": "3MYCJG5GVYQ8Q",
	    "country": "US",
	    "language_code": "en-US",
	    "currency": "USD",
	    "type": "PHYSICAL",
	    "description": "Midtown Atlanta store",
	    "coordinates": {
	      "latitude": 33.7889,
	      "longitude": -84.3841
	    },
	    "business_name": "Jet Fuel Coffee",
	    "mcc": "7299"
	  }
*/
func (s *SquareLocationService) CreateLocation(ctx context.Context, in *pb.SubscriptionSetupRequest,
	out *pb.SubscriptionSetupResponse) error {
	log.Printf("Calling CreateLocation as %v", ctx.Value("UserId"))

	loc, httpResponse, err := s.GetLocations(ctx)
	var message string

	if err != nil {
		log.Printf("Error when retrieving location for user %v: %v", ctx.Value("UserId"), err)
		return err
	}

	if loc != nil {
		log.Printf("Location already exists. Returning.")
		message = fmt.Sprintf("Location %v already exists. Returning existing location.", loc.Id)
	} else {
		locationStatus := square.ACTIVE_LocationStatus
		country := square.US_Country

		slr := square.CreateLocationRequest{
			Location: &square.Location{
				Name:        FIXED_ONLYSUBS_LOCATION_NAME,
				Status:      &locationStatus,
				Description: "location created for accepting payments by onlysubs",
				//TODO: enable users to pass in an address in the future
				Address: &square.Address{
					AddressLine1:                 "1234 Peachtree St. NE",
					Locality:                     "Atlanta",
					AdministrativeDistrictLevel1: "GA",
					PostalCode:                   "30309",
				},
				Timezone: "America/New_York",
				Country:  &country,
			},
		}

		locResponse, httpResp, err := s.Client.LocationsApi.CreateLocation(ctx, slr)

		if err != nil {
			log.Printf("Error when creating location for user %v: %v", ctx.Value("UserId"), err)
			return err
		}

		loc = locResponse.Location
		message = "Location created successfully."
		httpResponse = httpResp
	}

	out.LocationCreationResult = &pb.LocationCreationResult{
		HttpResponse: &pb.HttpResponse{
			Message:    message,
			Status:     fmt.Sprintf("%v", httpResponse.Status),
			StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
			Error:      "",
		},
		Location: &pb.Location{
			LocationId:  loc.Id,
			Name:        loc.Name,
			CountryCode: string(*loc.Country),
		},
	}

	return nil
}

// Should try to return only a single location for now
func (s *SquareLocationService) GetLocations(ctx context.Context) (*square.Location, *http.Response, error) {
	log.Printf("Calling GetLocations as %v", ctx.Value("UserId"))

	locResponse, httpResponse, err := s.Client.LocationsApi.ListLocations(ctx)

	if err != nil {
		return nil, httpResponse, err
	}

	for _, loc := range locResponse.Locations {
		if loc.Name == FIXED_ONLYSUBS_LOCATION_NAME {
			return &loc, httpResponse, nil
		}
	}
	return nil, httpResponse, nil
}
