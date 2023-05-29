export const getSubscriptions = async () => {
  const response = await fetch('v1/getSubscriptions')
    .then((res) => res.json());

  return response;
}
