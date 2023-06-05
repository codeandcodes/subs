import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Splash from './components/splash/splash';
import Authorize from './components/authorize/authorize';
import Feed from './components/feed/feed';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { setCurrentUser } from './store/session';
import Header from './components/header/header';
function App() {
  const dispatch = useDispatch();

  useEffect(() => {
    const loggedInUser = localStorage.getItem('user');
  
    if (loggedInUser) {
      dispatch(setCurrentUser(JSON.parse(loggedInUser)));
    }
  }, []);

  return (
    <BrowserRouter>
      <Header />
      <Routes>
        <Route path='/' element={<Splash />} />
        <Route path='/authorize' element={<Authorize />} />
        <Route path='/feed' element={<Feed />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
