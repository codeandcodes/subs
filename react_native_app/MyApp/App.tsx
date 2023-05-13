import React from 'react';
import HomeScreen from './src/components/HomeScreen';
import store from './src/store';
import { Provider } from 'react-redux';

const App = (): JSX.Element => {
  return (
    <Provider store={store}>
      <HomeScreen />
    </Provider>
  );
};

export default App;
