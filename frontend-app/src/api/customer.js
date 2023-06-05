
/**
 * Get customer info given a customer id
 * @param {string} customerId 
 * @returns 
 */
export const getCustomer = async (customerId) => {
  const body = JSON.stringify({
    customerId,
    includeSubscription: true
  });

  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body
  };

  const response = await fetch('v1/getCustomer', requestOptions)
    .then((res) => res.json());

  return response;
}