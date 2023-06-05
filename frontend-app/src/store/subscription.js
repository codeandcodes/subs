import { getCustomer } from '../api/customer';
import { getSubscriptions, setupSubscription } from '../api/subscription';
import { produce } from 'immer';

const GET_SUBSCRIPTIONS = 'getSubscriptions';
const SET_PAYER_INFO = 'setPayerInfo';
const ADD_SUBSCRIPTION = 'addSubscription';

const setSubscriptions = (subscriptions) => {
  return {
    type: GET_SUBSCRIPTIONS,
    subscriptions
  }
};

const addSubscription = (subscription) => {
  return {
    type: ADD_SUBSCRIPTION,
    subscription
  }
};

const setPayerInfo = (user) => {
  return {
    type: SET_PAYER_INFO,
    user
  }
};

export const fetchSubscriptions = () => async (dispatch) => {
  const response = await getSubscriptions();

  dispatch(setSubscriptions(response.subscriptions));
};

export const getPayerInfo = ({ customerId, subscriptionId }) => async (dispatch) => {
  const response = await getCustomer(customerId);
  const payerInfo = {
    subscriptionId,
    user: response.user
  };

  dispatch(setPayerInfo(payerInfo));

  return payerInfo;
}

export const addNewSubscription = (body) => async (dispatch) => {
  const response = await setupSubscription(body);

  const subscription = response.catalog_creation_result.subscription_plan;
  subscription.payers = {};
  subscription.subscriptions.push(Object.values(response.subscription_creation_results)[0].subscription);

  dispatch(addSubscription(subscription));

  return response;
}

const subscriptionsReducer = (state = { subscriptions: {} }, action) => {
  return produce(state, draft => {
    switch (action.type) {
      case GET_SUBSCRIPTIONS:
        Object.values(action.subscriptions).forEach(subscription => {
          draft[subscription.id] = subscription;
          draft[subscription.id].payers = {};

        });
        break;
      case SET_PAYER_INFO:
        draft[action.user.subscriptionId].payers[action.user.user.id] = action.user.user;
        break;
      case ADD_SUBSCRIPTION:
        draft[action.subscription.id] = action.subscription;
      default:
        break;
      }
  });
};

export default subscriptionsReducer;
