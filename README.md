# Onlysubs

## Installation Steps

### Prerequisites
- install node https://nodejs.org/en/download
- install homebrew https://brew.sh/
- install react native cli https://reactnative.dev/docs/next/environment-setup?guide=native
- install xcode for iOS simulator https://apps.apple.com/us/app/xcode/id497799835?mt=12


## Environment Setup Steps

### Backend Setup

1. Clone the repository:

2. Navigate to the backend directory:
  ``` 
  cd backend
  ```


3. Install the required dependencies:
  ``` 
  npm install
  ```

4. Start the backend server:
  ``` 
  node server.js
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
