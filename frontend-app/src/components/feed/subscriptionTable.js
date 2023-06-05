import React from 'react';
import { useSelector } from 'react-redux';
import SubscriptionCard from './subscriptionCard';

function SubscriptionTable() {
  const subscriptions = useSelector(state => state.subscriptions);

  if (!subscriptions) {
    return null;
  }

  return (
    <div>
      {Object.values(subscriptions).filter(Boolean).map((subscriptionPlan) => {
          if (subscriptionPlan.subscriptions && subscriptionPlan.subscriptions.length > 0) {
            return <SubscriptionCard subscription={subscriptionPlan} />;
          }
        })
      }
    </div>
  );
}

export default SubscriptionTable;
