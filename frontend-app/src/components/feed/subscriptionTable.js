import React from 'react';
import { useSelector } from 'react-redux';
import { useState } from 'react';

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
    <table>
      <thead>
        <tr>
          <th>Plan ID</th>
          <th>Updated At</th>
          <th>Plan Name</th>
          <th>Amount</th>
          <th>Cadence</th>
          <th>Periods</th>
          <th>Is Ongoing</th>
          <th>Start Date</th>
        </tr>
      </thead>
      <tbody>
        {Object.values(subscriptions).filter(Boolean).map((subscriptionPlan) => (
            <>
                <tr key={subscriptionPlan.id} onClick={() => handleRowClick(subscriptionPlan.id)} style={{
                backgroundColor: subscriptionPlan.subscriptions.length > 0 ? '#e0f7fa' : 'white',
                cursor: 'pointer'
              }}>

                    <td>{subscriptionPlan.id}</td>
                    <td>{subscriptionPlan.updated_at}</td>
                    <td>{subscriptionPlan.subscription_plan_data.name}</td>
                    <td>{subscriptionPlan.subscription_plan_data.amount}</td>
                    <td>{subscriptionPlan.subscription_plan_data.subscription_frequency.cadence}</td>
                    <td>{subscriptionPlan.subscription_plan_data.subscription_frequency.periods}</td>
                    <td>{subscriptionPlan.subscription_plan_data.subscription_frequency.is_ongoing.toString()}</td>
                    <td>{subscriptionPlan.subscriptions[0]?.start_date || 'N/A'}</td>
                </tr>
                {openRows.includes(subscriptionPlan.id) && subscriptionPlan.subscriptions.map(sub => (
                    <tr key={sub.id}>
                    <td>{sub.id}</td>
                    <td>{sub.start_date}</td>
                    <td>{sub.status}</td>
                    {/* Add more details as required */}
                    </tr>
                ))}
          </>
        ))}
      </tbody>
    </table>
  );
}

export default SubscriptionTable;
