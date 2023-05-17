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
  go run main.go --config=config.yaml --creds=credential.yaml
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

2. [Optional] Swagger UI to browse APIs
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
