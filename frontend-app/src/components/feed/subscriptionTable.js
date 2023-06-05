import React from 'react';
import { useSelector } from 'react-redux';
import { useState } from 'react';
import SubscriptionCard from './subscriptionCard';
import { Stack } from '@mui/material';

function SubscriptionTable() {

  const subscriptions = useSelector(state => state.subscriptions.subscriptions);
  const [openRows, setOpenRows] = useState([]);
  if (!subscriptions) {
    return null;
    }

    const handleRowClick = (subscriptionPlanId) => {
        // If row is already open, close it
        if (openRows.includes(subscriptionPlanId)) {
        setOpenRows(openRows.filter(id => id !== subscriptionPlanId));
        } else {
        setOpenRows([...openRows, subscriptionPlanId]);
        }
    }
  return (
    // <Stack spacing={2}>
    <div>
      {Object.values(subscriptions).filter(Boolean).map((subscriptionPlan) => {
          if (subscriptionPlan.subscriptions.length > 0) {
            return <SubscriptionCard subscription={subscriptionPlan} />;
          }
        })
      }
    </div>

    // </Stack>
  );
}

export default SubscriptionTable;
