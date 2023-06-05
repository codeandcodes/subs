import {
  Avatar,
  Box,
  Divider,
  Typography
} from '@mui/material';
import { Paid } from '@mui/icons-material';
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { getPayerInfo } from '../../store/subscription';
import { formatFrequency } from '../../utils/format-frequency';

// subscription: {
//   subscription_plan_data: {
//     name: "for you",
//     amount: 101,
//     subscription_frequency: {
//       cadence: "MONTHLY",
//       periods: 8
//     }
//   }
//   subscriptions: [
//     {
//       start_date: "2023-06-14",
//       customer_id: "123asdf"
//     }
//   ]
// }

function SubscriptionCard({ subscription }) {
  const dispatch = useDispatch();

  const name = subscription.subscription_plan_data.name;
  const amount = (subscription.subscription_plan_data.amount/100).toLocaleString('en-US', {
    style: 'currency',
    currency: 'USD',
  });
  const subscriptions = useSelector(state => state.subscriptions);
  const currentSubscription = subscriptions[subscription.id];

  // TODO: convert frequncy to readable string and calculate end date from periods
  const cadence = subscription.subscription_plan_data.subscription_frequency.cadence;
  const periods = subscription.subscription_plan_data.subscription_frequency.periods;

  const startDate = subscription.subscriptions[0].start_date;
  const frequency = formatFrequency({ cadence, periods, startDate });

  useEffect(() => {
    const subscriptionId = subscription.id;

    subscription.subscriptions.map((customerSubscription => {
      const customerId = customerSubscription.customer_id;
    
      dispatch(getPayerInfo({ customerId, subscriptionId }));
    }))
  }, [dispatch]);

  return (
    <>
      <Box display="flex" paddingY="12px">
        <Box display="flex" flexDirection="column" justifyContent="center">
          <Avatar sx={{ bgcolor: "#519872" }}>
            <Paid />
          </Avatar>
        </Box>
        <Box display="flex" flexDirection="column" flexGrow={1} sx={{ paddingLeft: "12px"}}>
          <Box display="flex" flexDirection="row" justifyContent="space-between">
            <Typography variant="h6" component="span" fontWeight="600">{name}</Typography>
            <Typography variant="h6" component="span" sx={{ color: "#519872", fontWeight: "600" }} >{amount}</Typography>
          </Box>
          <Typography>Start Date: {startDate}</Typography>
          <Typography>{frequency}</Typography>
          <Box>
            {subscription.payers && Object.values(subscription.payers).map((payer) => {
              return (
                <>
                  <Typography>Payer: {payer.email_address}</Typography>
                </>
              )
            })}
          </Box>
          <Divider sx={{ paddingTop: "12px" }} />
        </Box>
      </Box>
    </>
  )
};

export default SubscriptionCard;

