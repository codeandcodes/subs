import './header.css';
import { 
  AppBar,
  Avatar,
  Box,
  Button,
  Toolbar,
  Typography
} from '@mui/material';
import { useSelector } from 'react-redux';
import { useState } from 'react';

function Header() {
  const user = useSelector(state => state.session.user);
  const [anchorEl, setAnchorEl] = useState(null);

  // const handleChange = (event) => {
  //   setAuth(event.target.checked);
  // };

  const handleMenu = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLogout = () => {
    console.log("logout");
  };

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            onlySubs.
          </Typography>
          {user &&
            <Box display="flex" alignItems="center" aria-controls="menu-appbar">
              <Avatar alt={user.displayName} src={user.photoUrl} />
              <Typography sx={{ paddingLeft: '8px' }}>{user.displayName}</Typography>
              <Button onClick={handleLogout} style={{ color: 'white' }}>LOGOUT</Button>
            </Box>
          }
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default Header;