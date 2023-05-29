export const addSquareAccessToken = async (token, osUserId) => {
  const body = JSON.stringify({
    squareAccessToken: token,
    osUserId
  });

  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body
  };

  const response = await fetch('v1/addSquareAccessToken', requestOptions)
  .then((res) => res.json());

  return response;
};

export const registerUser = async (user) => {
  const body = JSON.stringify(user);

  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body
  };

  const response = await fetch('v1/registerUser', requestOptions)
    .then((res) => res.json());

    return response;
};
