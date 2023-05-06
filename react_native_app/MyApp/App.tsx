import React, { useEffect, useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import { fetchBackendData, fetchLocation } from './api';

const App = (): JSX.Element => {
  const [backendData, setBackendData] = useState('');

  useEffect(() => {
    const fetchData = async (): Promise<void> => {
      const data: string = await fetchBackendData();
      setBackendData(data);
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
