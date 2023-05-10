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

### Backend Setup

1. Clone the repository:
  ```
  ```

2. Navigate to the backend directory:
  ``` 
  cd backend
  ```

3. Install the required dependencies:
  ``` 
  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
  go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
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
  cd backend/grpcserver
  go run main.go
  ```

5. Start the http gateway server (runs on port 3000):
  ``` 
  cd backend/httpserver
  go run main.go
  ```

6. To update dependencies
  ```
  # in root
  go mod tidy
  ```

### Example curl


### Proto Setup

1. Compile protos for backend (in src root dir). Ensure that googleapis references where you cloned the repo.
  ```
  protoc --proto_path=./protos --proto_path=$HOME/workspace/googleapis \
    --go_out=protos --go_opt=paths=source_relative \
    --go-grpc_out=protos --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=protos --grpc-gateway_opt=paths=source_relative \
    protos/*.proto
  ```

### React Native Setup

1. Open a new terminal window.

2. Navigate to the app directory:
  ``` 
  cd react_native_app
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

  ``` 
  cd MyApp
  bundle install
  cd ios && bundle exec pod install
  ```
