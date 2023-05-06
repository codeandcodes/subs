const API_URL = 'http://localhost:3000';

export const fetchBackendData = async (): Promise<string | undefined> => {
  try {
    const response = await fetch(`${API_URL}`);
    const data = await response.text();
    return data;
  } catch (error) {
    console.log(error);
  }
};

export const fetchLocation = async (): Promise<string | undefined> => {
  try {
    const response = await fetch(`${API_URL}/location`);
    const data = await response.text();
    return data;
  } catch (error) {
    console.log(error);
  }
};
