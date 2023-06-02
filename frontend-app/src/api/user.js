/**
 * Saves the square access token with the currently logged in user
 * @param {string} token - square access token
 * @returns 
 */
export const addSquareAccessToken = async (token) => {
  const body = JSON.stringify({
    squareAccessToken: token,
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

/**
 * Login the user and set the user's session cookie
 * @param {string} osUserId - the os user id returned from registerUser
 */
export const loginUser = async (osUserId) => {
  // login with username and password is currently not implemented, so we are just passing in the same string for now
  const body = JSON.stringify({
    os_user_id: osUserId,
    username: "osUser",
    password: "abc123"
  });

  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body
  };

  console.log(requestOptions);

  const response = await fetch('/loginUser', requestOptions)
    .then((res) => {
      return res;
    });

    return response;
}

/**
 * Get the user if it exists
 * @param {string} emailAddress - email address returned from FB user credential
 */
export const getUser = async (emailAddress) => {
  const response = await fetch(`/v1/getUser?emailAddress=${emailAddress}`)
    .then(res => res.json());

  return response;
}
