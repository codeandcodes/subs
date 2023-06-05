import {
  Avatar,
  Box,
  Divider,
  Typography
} from '@mui/material';
import { Paid } from '@mui/icons-material';

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
  const name = subscription.subscription_plan_data.name;
  const amount = `\$${subscription.subscription_plan_data.amount}`;
  
  // TODO: convert frequncy to readable string and calculate end date from periods
  const frequency = subscription.subscription_plan_data.subscription_frequency.cadence;
  const periods = subscription.subscription_plan_data.subscription_frequency.periods;

  // iterate through subscriptions for payers
  const startDate = subscription.subscriptions[0].start_date;

  return (
    <>
      <Box display="flex" paddingY="12px">
        <Box display="flex" flexDirection="column" justifyContent="center">
          <Avatar sx={{ bgcolor: "green" }}>
            <Paid />
          </Avatar>
        </Box>
        <Box display="flex" flexDirection="column" flexGrow={1} sx={{ paddingLeft: "12px"}}>
          <Box display="flex" flexDirection="row" justifyContent="space-between">
            <Typography variant="h6" component="span" fontWeight="600">{name}</Typography>
            <Typography variant="h6" component="span">{amount}</Typography>
          </Box>
          <Typography>Start Date: {startDate}</Typography>
          <Typography>{frequency}</Typography>
          <Typography>{periods}</Typography>
          <Divider />
        </Box>
      </Box>
    </>
  )
};

export default SubscriptionCard;

