# Onlysubs/Sugarbaby

## Try it out

[OnlySubs Dev](https://onlysubs-dev.com)

**Instructions**

1. Click 'Login with Facebook'. 
2. Hit 'Authorize'. If you see this text: "To start the OAuth flow for a sandbox account, first launch the seller test account from the Developer Dashboard.", you'll then need to enable Square OAuth to create a subscription. Open up your square sandbox account in order to do this in square developer dashboard.
3. You will then see a list of current subscriptions. You can also try it out hit "Setup New Subscription".

## Installation Steps

### Prerequisites
- install homebrew https://brew.sh/
- install go https://formulae.brew.sh/formula/go
- install react native cli https://reactnative.dev/docs/next/environment-setup?guide=native
- install xcode for iOS simulator https://apps.apple.com/us/app/xcode/id497799835?mt=12
- install protobuf (note must be 3.20.x)
  - there's a bug with protoc 21.x and up which no longer includes the js compiler. There's a workaround: 
  ```
  # on mac
  brew install protobuf@3
  echo 'export PATH="/opt/homebrew/opt/protobuf@3/bin:$PATH"' >> ~/.zshrc
  protoc --version
  # should be libprotoc 3.20.3
  ```

## Environment Setup Steps

1. Clone the repository
  ```
  git clone git@github.com:codeandcodes/subs.git 
  ```

### Backend Setup

Prerequisites:
- note: instructions all relative from where you cloned the repo. E.g. ~/workspace/subs

2. Navigate to the backend directory:
  ``` 
  cd backend
  ```

3. Install the required dependencies:
  ``` 
  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

4. Add to $PATH var in .zshrc or .bash_profile
  ```
  export GO_PATH=~/go
  export PATH=$PATH:/$GO_PATH/bin
  ```

3. Add googleapis repository for generating grpc-gateway code
  ```
  # somewhere e.g. ~/workspace
  git clone https://github.com/googleapis/googleapis.git
  ```

4. Start the backend gRPC server:
  ``` 
  # edit config.yaml
  # add access_token to config.yaml, then
  cd backend/grpcserver
  go run main.go --config=config.yaml
  ```

5. Start a new terminal window/tab. Start the http gateway server (runs on port 3000):
  ``` 
  cd backend/httpserver
  go run main.go
  ```

6. To update dependencies
  ```
  # in root
  go mod tidy
  ```

### API Documentation

1. Check if backend server is responding
  ```
  curl -X GET http://localhost:3000/
  # output: Hello, this is the root route!%
  ```

2. Note: you must authenticate to onlysubs
  ```
  curl -X 'POST' \
  'http://localhost:3000/loginUser' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "user_id": "nOToaHZTFXDKn1OVN1Sj",
    "username": "abcde",
    "password": "12345"
  }' -v
  ```

  What this does is ask your browser to set a cookie that has the encoded session ID. This session ID is used to identify the caller on the GRPC side.
  You will get a response like this:
  ```
  < HTTP/1.1 200 OK
  < Set-Cookie: onlysubs_session=XXXXTM3MzU2NXxMd3dBTEVoRVRubFdjVlJuUjBGVFJVOHpSMlExWjBkWlJVUjRUWEF3ZVd0NVpHdFhkR3QzUmpSZmRYZFdka0U5fDFMt8bX6e6IYhyiT9hYyDquYUA2f3mu7VZ6j7dwpNgM; Path=/; HttpOnly; Secure; SameSite=Strict
  ```

  You then must copy and paste this cookie into all subsequent curl commands. If you do not, you will get an unauthorized error back from the server while making requests. See below. 

3. Enable square oauth to store the access token. 

You must go through the square oauth flow to get the square access token. This needs to then be called with /v1/addSquareAccessToken, otherwise you won't be able to call any square APIs.
  ```
  curl -X 'POST' \
    'http://localhost:3000/v1/addSquareAccessToken' \
    -b 'onlysubs_session=XXXXXTM3MzU2NXxMd3dBTEVoRVRubFdjVlJuUjBGVFJVOHpSMlExWjBkWlJVUjRUWEF3ZVd0NVpHdFhkR3QzUmpSZmRYZFdka0U5fDFMt8bX6e6IYhyiT9hYyDquYUA2f3mu7VZ6j7dwpNgM' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d '{
    "squareAccessToken": "string"
  }'
  ```

4. Try APIs out via curl. You can generate the curl command from swagger UI (see below) and then paste in an parameter to pass in cookies to curl. E.g. (-b):
  ```
  curl -X 'GET' \
  -b 'onlysubs_session=XXXXXTM3MzU2NXxMd3dBTEVoRVRubFdjVlJuUjBGVFJVOHpSMlExWjBkWlJVUjRUWEF3ZVd0NVpHdFhkR3QzUmpSZmRYZFdka0U5fDFMt8bX6e6IYhyiT9hYyDquYUA2f3mu7VZ6j7dwpNgM' \
  'http://localhost:3000/v1/getCustomers' \
  -H 'accept: application/json' -v
  ```

5. [Optional] Swagger UI to browse APIs

- [MANDATORY] Note, you can't use the **"try-it-out"** function in swagger UI because swagger UI does not support cookies. The suggested method is clicking try-it-out, then modifying the payload, copying into a curl command and adding the cookie manually.

- clone swagger UI repo
  ```
  git clone https://github.com/swagger-api/swagger-ui.git
  cd swagger-ui
  npm run dev
  # Wait a bit
  # Open http://localhost:3200/
  In searchbox, type in
  http://localhost:3000/static/protos/api.swagger.json
  ```

- [Optional] You can also modify swagger-ui/dev-helpers/dev-helper-initializer.js to automatically load this whenever your run $ npm run dev.
  - modify the line in dev-helper-initializer.js url to "http://localhost:3000/static/protos/api.swagger.json"

### Proto Setup

0. [Optional] Use script
  ```
  ./regenProtos.sh # this just calls the steps below
  ```

1. Compile protos for backend (in src root dir). Ensure that googleapis references where you cloned the repo.
  ```
  # from repo root /
  protoc --proto_path=./protos --proto_path=$HOME/workspace/googleapis \
    --go_out=protos --go_opt=paths=source_relative \
    --go-grpc_out=protos --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=protos --grpc-gateway_opt=paths=source_relative \
    protos/*.proto
  ```

2. [Optional] Generate OpenAPIv2 files
  ```
  # from repo root /
  protoc -I . --proto_path=./protos --proto_path=$HOME/workspace/googleapis \
    --openapiv2_out ./backend/httpserver/static/ \
    --openapiv2_opt logtostderr=true \
    protos/api.proto
  ```

### Frontend Setup
1. Add `.env` to `/frontend-app` and `firebase-config.js` to `frontend-app/src`

2. in `/frontend-app` run `npm install` (first time only)

3. To serve: `npm start`. Note that we are serving on `localhost:4000` (port is defined in .env file) since the http server is `localhost:3000`

### React Native Setup

1. Open a new terminal window.

2. Navigate to the app directory:
  ``` 
  cd react_native_app/MyApp
  ```

3. Install the required dependencies:
  ```
  npm install
  ```

4. Start the React Native development server:
- For Android:
  ```
  npx react-native run-android
  ```
- For iOS:
  ```
  npx react-native run-ios
  ```

## Troubleshooting

1. If you see "Unknown ruby interpreter version (do not know how to handle): >=2.6.10." make sure you update ruby

set ruby version with rbenv
  ``` 
  rbenv install 2.7.6
  rbenv global 2.7.6
  ```

  If you have errors running iOS app using npx, try manually installing dependencies
  ``` 
  cd MyApp
  bundle install
  cd ios && bundle exec pod install
  ```

2. Missing session tokens. 

If you get errors like
 ```
 {
  "error": "Missing session token",
  "code": 16,
  "message": "Missing session token",
  "details": []
}
```
you probably haven't authenticated and/or copied the cookie into your curl command.
To make API calls from the backend, follow the steps above in ### API Documentation 

3. Missing square access token.

If you see ```User %v has no associated square access token. Cannot call square services.``` 
then you probably haven't associated your square access token. Go through square oauth to get an access token, then store it via the /v1/addSquareAccessToken or directly in the db.

Productionization

1) npm run build - to build
2) mv subs/frontend-app/build to subs/backend/httpserver (need to serve out of go httpserver)
3) redirect onlysubs-dev.com to IP address per ec2
4) generate certs from certbot 
5) ensure that build/.well-known can be served
6) enable TLS on https.go (copy of main.go) which uses log.Fatal(http.ListenAndServeTLS(":443", "path/to/cert.pem", "path/to/key.pem", mux))
7) sudo go run https.go because port <1024 (e.g. https 443) traffic requires root
8) add domain to firebase https://onlysubs-dev.com/
9) Open up IP routes table in ec2 instance:
  262  sudo iptables -A PREROUTING -t nat -i eth0 -p tcp --dport 80 -j REDIRECT --to-port 8080
  263  sudo iptables -A INPUT -p tcp -m tcp --sport 80 -j ACCEPT
  264  sudo iptables -A OUTPUT -p tcp -m tcp --dport 80 -j ACCEPT
