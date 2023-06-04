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

  return(
    <div>
      <h1>This is the Feed/homepage</h1>
      <p>{userName}</p>
      <SetupSubscriptionModal />
    </div>
  )

}

export default Feed;
