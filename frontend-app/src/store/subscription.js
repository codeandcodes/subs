import { getSubscriptions } from '../api/subscription';

const GET_SUBSCRIPTIONS = 'getSubscriptions';

const setSubscriptions = (subscriptions) => {
  return {
    type: GET_SUBSCRIPTIONS,
    subscriptions
  }
};

export const fetchSubscriptions = () => async (dispatch) => {
  const response = await getSubscriptions();

  dispatch(setSubscriptions(response.subscriptions));
};

const subscriptionsReducer = ( state = { subscriptions: null }, action ) => {
  let newState;

  switch (action.type) {
    case GET_SUBSCRIPTIONS:
      newState = Object.assign({}, state);
      newState.subscriptions = action.subscriptions;
      return newState;
    default:
      return state;
  }
};

export default subscriptionsReducer;
