import React, { useEffect, useState } from 'react';
import { Button, View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import { fetchBackendData, fetchLocation } from './api';
import auth from '@react-native-firebase/auth';
import { LoginManager, AccessToken } from 'react-native-fbsdk-next';

const App = (): JSX.Element => {
  const [backendData, setBackendData] = useState('');

  const signInWithFacebook = async () => {
    // Attempt login with permissions
    const result = await LoginManager.logInWithPermissions(['public_profile', 'email']);

    if (result.isCancelled) {
      throw 'User cancelled the login process';
    }

    // Once signed in, get the users AccesToken
    const data = await AccessToken.getCurrentAccessToken();

    if (!data) {
      throw 'Something went wrong obtaining access token';
    }

    // Create a Firebase credential with the AccessToken
    const facebookCredential = auth.FacebookAuthProvider.credential(data.accessToken);

    // Sign-in the user with the credential
    return auth().signInWithCredential(facebookCredential);
  }

  useEffect(() => {
    const fetchData = async (): Promise<void> => {
      const data: string | undefined = await fetchBackendData();

      if (data) setBackendData(data);
    };

    fetchData();
  }, []);

  return (
    <View style={styles.container}>
      <Text>Welcome to My React Native App!</Text>
      <Text>Data from the backend: {backendData}</Text>
      <TouchableOpacity style={styles.button} onPress={fetchLocation}>
            <Text style={styles.buttonText}>Get Location</Text>
      </TouchableOpacity>
      <Button title="Sign in with Facebook" onPress={signInWithFacebook} />
    </View>
  );
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

export default App;
