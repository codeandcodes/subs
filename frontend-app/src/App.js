import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Splash from './components/splash/splash';
import Authorize from './components/authorize/authorize';
import Feed from './components/feed/feed';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Splash />} />
        <Route path='/authorize' element={<Authorize />} />
        <Route path='/feed' element={<Feed />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
