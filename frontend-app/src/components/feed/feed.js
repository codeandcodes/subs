import { useDispatch, useSelector } from 'react-redux';
import { fetchSubscriptions } from '../../store/subscription';
import { setCurrentUser } from '../../store/session';
import { useEffect } from 'react';
import SetupSubscriptionModal from '../setupSubscriptionModal/setupSubscriptionModal';

function Feed() {
  const dispatch = useDispatch();
  const userName = useSelector(state => state.session.user?.displayName);

  useEffect(() => {
    const loggedInUser = localStorage.getItem('user');
  
    if (loggedInUser) {
      dispatch(setCurrentUser(JSON.parse(loggedInUser)));
    }
  }, []);

  useEffect(() => {
    dispatch(fetchSubscriptions());
  }, [dispatch]);

  const scope = 'CUSTOMERS_READ+PAYMENTS_WRITE+SUBSCRIPTIONS_WRITE+ITEMS_READ+ORDERS_WRITE+INVOICES_WRITE+ITEMS_WRITE+ITEMS_READ+CUSTOMERS_WRITE+SUBSCRIPTIONS_READ+MERCHANT_PROFILE_READ+MERCHANT_PROFILE_WRITE';
  const token = 'pretendingthatthisissomekindoftoken123';
  const authorizeUrl = `https://connect.squareupsandbox.com/oauth2/authorize?client_id=${process.env.REACT_APP_SQUARE_APPLICATION_ID}&scope=${scope}&session=false&state=${token}`;

  return(
    <div>
      <h1>This is the Feed/homepage</h1>
      <p>{userName}</p>
      <a href={authorizeUrl}>authorize square</a>
      <SetupSubscriptionModal />
    </div>
  )

}

export default Feed;
