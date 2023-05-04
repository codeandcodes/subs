import React, { useEffect, useState } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { fetchBackendData } from './api';

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
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});

export default App;
