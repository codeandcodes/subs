import { useDispatch, useSelector } from 'react-redux';
import { fetchSubscriptions } from '../../store/subscription';
import { setCurrentUser } from '../../store/session';
import { useEffect } from 'react';
import SetupSubscriptionModal from '../setupSubscriptionModal/setupSubscriptionModal';
import SubscriptionTable from './subscriptionTable';
import Header from '../header/header';
import {
  Box,
  Button,
  Typography
} from '@mui/material';

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
      <Header />
      <Box display="flex" alignItems="center" flexDirection="column" sx={{ padding: "24px"}}>
        <Typography variant="h3" sx={{ paddingBottom: "12px"}}>my subs</Typography>
        <SetupSubscriptionModal />
      </Box>
      <SubscriptionTable />
    </div>
  )

}

export default Feed;
