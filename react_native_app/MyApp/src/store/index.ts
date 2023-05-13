import { configureStore } from '@reduxjs/toolkit';
import userReducer from './user';

const preloadedState = {};

let logger;

if (process.env.NODE_ENV !== 'production') {
  logger = require('redux-logger').default;
}

const store = configureStore({
  reducer: {
    user: userReducer,
  }, preloadedState, middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(logger) });

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch

export default store;