import { useDispatch, useSelector } from 'react-redux';
// import { getSubscriptions } from '../../api/subscription';
import { setCurrentUser } from '../../store/session';
import { useEffect } from 'react';

function Feed() {
  const dispatch = useDispatch();
  const userName = useSelector(state => state.session.user?.displayName);
  // const subscriptions = useSelector(state => state.subscriptions);

  useEffect(() => {
    const loggedInUser = localStorage.getItem('user');
  
    if (loggedInUser) {
      dispatch(setCurrentUser(JSON.parse(loggedInUser)));
    }
  }, []);

  // useEffect(() => {
  //   dispatch(getSubscriptions());
  // }, [dispatch]);

  const handleClick = () => {
    // getSubscriptions().then(res => {
    //   console.log(res);
    // })
  }

  // console.log(subscriptions);

  return(
    <div>
      <h1>This is the Feed/homepage</h1>
      <p>{userName}</p>
      <button onClick={handleClick}>Get Subscriptions</button>
    </div>
  )

}

export default Feed;
