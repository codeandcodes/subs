# Onlysubs/Sugarbaby

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
  go run main.go --config=config.yaml --creds=credential.yaml --dev-mode=true
  ```

  [Important!] If you want to test out actual flows using square oauth, you must start the server with --dev-mode=false

  This will enforce authentication and you will need to call /loginUser first, then follow the steps to copy the cookies below *see API Documentation*

  If you are testing from the UI, you should use --dev-mode=false.

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

2. If not started in dev mode, you must authenticate
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

3. Try APIs out. Once authentication is enabled, you'll no longer be able to use the Swagger UI in step 4 below. Instead, you can generate the curl command from swagger UI and then paste in an parameter to pass in cookies to curl. E.g. (-b):
  ```
  curl -X 'GET' \
  -b 'onlysubs_session=XXXXXTM3MzU2NXxMd3dBTEVoRVRubFdjVlJuUjBGVFJVOHpSMlExWjBkWlJVUjRUWEF3ZVd0NVpHdFhkR3QzUmpSZmRYZFdka0U5fDFMt8bX6e6IYhyiT9hYyDquYUA2f3mu7VZ6j7dwpNgM' \
  'http://localhost:3000/v1/getCustomers' \
  -H 'accept: application/json' -v
  ```

4. [Optional] Swagger UI to browse APIs

- [MANDATORY] To use Swagger UI, you must ensure that dev-mode is set to true (this is the default)

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

If you see "Unknown ruby interpreter version (do not know how to handle): >=2.6.10." make sure you update ruby

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

If you get errors like
 ```
 {
  "error": "Missing session token",
  "code": 16,
  "message": "Missing session token",
  "details": []
}
```
you probably started the server with --dev-mode=false. To make API calls from the backend, follow the steps above in ### API Documentation 