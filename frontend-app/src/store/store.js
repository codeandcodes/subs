import { configureStore } from '@reduxjs/toolkit';
import sessionReducer from './session';
import subscriptionsReducer from './subscription';
import logger from 'redux-logger'

const preloadedState = {};

const store = configureStore({
  reducer: {
    session: sessionReducer,
    subscriptions: subscriptionsReducer
  }, preloadedState,
   middleware: (getDefaultMiddleware) => 
   getDefaultMiddleware().concat(logger) });

export default store;
