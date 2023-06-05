import { useDispatch } from 'react-redux';
import { fetchSubscriptions } from '../../store/subscription';
import { setCurrentUser } from '../../store/session';
import { useEffect } from 'react';
import SetupSubscriptionModal from '../setupSubscriptionModal/setupSubscriptionModal';
import SubscriptionTable from './subscriptionTable';
import {
  Container,
  Box,
  Button,
  Typography
} from '@mui/material';

function Feed() {
  const dispatch = useDispatch();

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
    <Container maxWidth="sm">
      <Box display="flex" alignItems="center" flexDirection="column" sx={{ padding: "24px"}}>
        <Typography variant="h4" sx={{ fontWeight: "600", borderBottom: "5px dotted #519872", marginBottom: "12px"}}>current subs</Typography>
        <SetupSubscriptionModal />
      </Box>
      <SubscriptionTable />
    </Container>
  )
}

export default Feed;
