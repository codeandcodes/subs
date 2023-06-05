import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { login, logout } from '../../store/session';
import { useNavigate } from 'react-router-dom';
import {
  Button,
  Typography
} from '@mui/material';
import './splash.css';

function Splash() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const user = useSelector(state => state.session.user);

  const handleLogin = (e) => {
    e.preventDefault();
    return dispatch(login());
  };

  const scope = 'CUSTOMERS_READ+PAYMENTS_WRITE+SUBSCRIPTIONS_WRITE+ITEMS_READ+ORDERS_WRITE+INVOICES_WRITE+ITEMS_WRITE+ITEMS_READ+CUSTOMERS_WRITE+SUBSCRIPTIONS_READ+MERCHANT_PROFILE_READ+MERCHANT_PROFILE_WRITE';
  const token = 'pretendingthatthisissomekindoftoken123';
  const authorizeUrl = `https://connect.squareupsandbox.com/oauth2/authorize?client_id=${process.env.REACT_APP_SQUARE_APPLICATION_ID}&scope=${scope}&session=false&state=${token}`;

  useEffect(() => {
    if (user && user.hasSquareAccessToken) {
      navigate('/feed');
    } 
  }, [navigate, user]);
  
  return (
    <div>
      {!!user
        ?
          <div className="authorize">
            <img src={user && user.photoUrl} />
            <Typography variant="h5">Welcome {user && user.displayName}</Typography>
            <Typography variant="subtitle" sx={{ padding: "24px 0"}}>Authorize square in order to continue</Typography>
            <Button variant="contained" href={authorizeUrl}>Authorize</Button>
          </div>
        :
          <div className="background">
            <div className="text-background">
              <Typography variant="h6" sx={{ color: "black"}}>Create peer to peer subscriptions to make regular payments easier between family, friends, and gig workers</Typography>
            </div>
            <Button onClick={handleLogin} variant="contained" sx={{ width: "max-content"}}>Sign in With Facebook</Button>
          </div>
          }
      </div>
  );
}

export default Splash;
