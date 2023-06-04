export const getSubscriptions = async () => {
  const response = await fetch('v1/getSubscriptions')
    .then((res) => res.json());

  return response;
}

export const setupSubscription = async ({
  name,
  amount,
  description,
  payee,
  payer,
  frequency
}) => {
  const body = JSON.stringify({
    name,
    description,
    amount,
    payee,
    payer,
    subscriptionFrequency: frequency
  });

  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body
  };

  const response = await fetch('/v1/setupSubscription', requestOptions)
    .then((res) => res.json());

  return response;
}
