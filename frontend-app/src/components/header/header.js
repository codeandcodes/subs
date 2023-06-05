import './header.css';
import { 
  AppBar,
  Avatar,
  Box,
  Button,
  Toolbar,
  Typography
} from '@mui/material';
import { useSelector, useDispatch } from 'react-redux';
import { useEffect, useState } from 'react';
import { logout } from '../../store/session';
import { useNavigate } from 'react-router-dom';

function Header() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const user = useSelector(state => state.session.user);

  const handleLogout = () => {
    return dispatch(logout());
  };

  useEffect(() => {
    if (!user) {
      navigate('/');
    }
  }, [navigate, user])

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h5" component="div" sx={{ flexGrow: 1, fontWeight: "600" }}>
            onlySubs.
          </Typography>
          {user &&
            <Box display="flex" alignItems="center" aria-controls="menu-appbar">
              <Avatar alt={user.displayName} src={user.photoUrl} />
              <Typography sx={{ paddingLeft: '8px' }}>{user.displayName}</Typography>
              <Button onClick={handleLogout} style={{ color: 'white' }}>
                <Typography variant="button">LOGOUT</Typography>
              </Button>
            </Box>
          }
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default Header;