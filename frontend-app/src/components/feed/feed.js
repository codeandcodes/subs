import { useSelector } from 'react-redux';

function Feed() {
  const accessToken = useSelector(state => state.session.user.squareAccessToken);

  return(
    <div>
      <h1>This is the Feed/homepage</h1>
      <p>Access token for square: {accessToken}</p>
    </div>
  )

}

export default Feed;
