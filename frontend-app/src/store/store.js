import { configureStore } from '@reduxjs/toolkit';
import sessionReducer from './session';

const preloadedState = {};
let logger;

if (process.env.NODE_ENV !== 'production') {
  logger = require('redux-logger').default;
}

const store = configureStore({
  reducer: {
    session: sessionReducer
  }, preloadedState, middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(logger) });

export default store;
