const { Client, Environment, ApiError } = require("square");

// Function to initialize the Square client
function createSquareClient(accessToken, environment) {
  return new Client({
    environment: environment === 'production' ? Environment.Production : Environment.Sandbox,
    accessToken: accessToken,
  });
}

// Function to fetch catalog items from Square API
async function getCatalogItems(squareClient) {
  const catalogApi = squareClient.catalogApi;
  const { result: catalogResponse } = await catalogApi.listCatalog();
  return catalogResponse?.objects ?? [];
}

// Function to create a subscription using Square API
async function createSubscription(squareClient, idempotencyKey, locationId, planId, customerId) {
  const subscriptionsApi = squareClient.subscriptionsApi;
  const requestBody = {
    idempotencyKey: idempotencyKey,
    locationId: locationId,
    planId: planId,
    customerId: customerId,
  };
  const { result: subscriptionResponse } = await subscriptionsApi.createSubscription(requestBody);
  return subscriptionResponse;
}

async function getLocations(squareClient) {
  const { locationsApi } = squareClient;

  try {
    let listLocationsResponse = await locationsApi.listLocations();

    let locations = listLocationsResponse.result.locations;

    locations.forEach(function (location) {
      console.log(
        location.id + ": " +
          location.name +", " +
          location.address.addressLine1 + ", " +
          location.address.locality
      );
    });
  } catch (error) {
    if (error instanceof ApiError) {
      error.result.errors.forEach(function (e) {
        console.log(e.category);
        console.log(e.code);
        console.log(e.detail);
      });
    } else {
      console.log("Unexpected error occurred: ", error);
    }
  }
}

// Export the functions
module.exports = { createSquareClient, getCatalogItems, createSubscription, getLocations };
