
export const getOauthToken = async (code) => {
  const body = JSON.stringify({
    client_id: process.env.REACT_APP_SQUARE_APPLICATION_ID,
    client_secret: process.env.REACT_APP_SQUARE_CLIENT_SECRET,
    code,
    grant_type: 'authorization_code'
  });

  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body
  }

  const response = await fetch('https://connect.squareupsandbox.com/oauth2/token', requestOptions)
    .then((res) => res.json());

  return response;
}

export const revokeOauthToken = async (accessToken) => {
  const body = JSON.stringify({
    access_token: accessToken,
    client_id: process.env.REACT_APP_SQUARE_APPLICATION_ID
  });

  const requestOptions = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Client ${process.env.REACT_APP_SQUARE_CLIENT_SECRET}`
    }
  }

  const response = await fetch('https://connect.squareupsandbox.com/oauth2/token', requestOptions)
  .then((res) => res.json());

  return response;
}
