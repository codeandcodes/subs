import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { login, logout, setCurrentUser } from '../../store/session';

function Splash() {
  const dispatch = useDispatch();
  const user = useSelector(state => state.session.user);
  const [clickedLogin, setClickedLogin] = useState(false);
  const [clickedLogout, setClickedLogout] = useState(false);

  const handleLogin = (e) => {
    e.preventDefault();
    return dispatch(login()).then(() => setClickedLogin(true));
  };

  const handleLogout = (e) => {
    e.preventDefault();
    return dispatch(logout()).then(() => setClickedLogout(true));
  }

  const token = 'pretendingthatthisissomekindoftoken123';
  const authorizeUrl = `https://connect.squareupsandbox.com/oauth2/authorize?client_id=${process.env.REACT_APP_SQUARE_APPLICATION_ID}&scope=CUSTOMERS_WRITE+CUSTOMERS_READ&session=false&state=${token}`;

  return (
    <>
      {!!user
        ?
          <div>
            <h1>{user.displayName}</h1>
            <img src={user.photoUrl} />
            <a href={authorizeUrl}>authorize square</a>
            <button onClick={handleLogout}>Logout</button>
          </div>
        :
          <div>
            <h1>Pay Your Friends</h1>
            <button onClick={handleLogin}>Sign in With Facebook</button>
          </div>
      }
    </>

  );
}

export default Splash;