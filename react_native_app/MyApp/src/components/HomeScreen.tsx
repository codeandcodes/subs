import React, { useState } from 'react';
import { View, StyleSheet, Text, TouchableOpacity } from 'react-native';
import { useAppDispatch, useAppSelector } from '../../hooks';
import { login, logout } from '../store/user';


const HomeScreen = () => {
  const dispatch = useAppDispatch();
  const user = useAppSelector(state => state.user);
  const [clickedLogin, setClickedLogin] = useState(false);
  const [clickedLogout, setClickedLogout] = useState(false);

  console.log(user);
  const handleLogin = () => {
    return dispatch(login()).then(() => setClickedLogin(true));
  }

  const handleLogout = () => {
    return dispatch(logout()).then(() => setClickedLogout(true));
  }

  // TODO: remove nested user object?
  return (
    <View style={styles.container}>
      {user.user !== null
        ?
          <View>
            <Text>{user.user.displayName}</Text>
            <TouchableOpacity style={styles.button} onPress={handleLogout}>
              <Text style={styles.buttonText}>Logout</Text>
            </TouchableOpacity>
          </View>
        :
          <View>
            <Text>Square Up</Text>
            <Text>Set up or pay automatic payments</Text>
            <TouchableOpacity style={styles.button} onPress={handleLogin}>
              <Text style={styles.buttonText}>Login with Facebook</Text>
            </TouchableOpacity>
          </View>
      }
    </View>
  )
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  button: {
    backgroundColor: '#007AFF',
    borderRadius: 4,
    padding: 10,
    alignItems: 'center',
    margin: 10,
  },
  buttonText: {
    color: '#FFFFFF',
    fontSize: 16,
  },
});


export default HomeScreen;
