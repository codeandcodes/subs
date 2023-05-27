import React, { useEffect, useState } from 'react';
import { useSearchParams, useNavigate } from 'react-router-dom';
import { getOauthToken } from '../../api/oauth';
import { setUserWithToken } from '../../store/session';
import { useDispatch } from 'react-redux';

function Authorize() {
  const dispatch = useDispatch();
  const [searchParams, setSearchParams] = useSearchParams();
  const [authErrors, setAuthErrors] = useState(null);
  const [oauthTokens, setOauthTokens] = useState(null);
  const navigate = useNavigate();

  const code = searchParams.get('code');

  useEffect(() => {
    getOauthToken(code).then((response) => {
      if (response.errors) {
        setAuthErrors(response.errors);
        navigate('/');
      } else {
        setOauthTokens(response);
        dispatch(setUserWithToken(response.access_token));
        navigate('/feed');
      }
    });
  }, [])

  // show a loader or something here?
  return (
    <div>
      {/* {oauthTokens
        ? 
          <div>
            <h1>Authorized!</h1>
            <p>{oauthTokens.access_token}</p>
          </div>
        :
          <div>
            <h1>Errors</h1>
            <p>{authErrors[0].detail}</p>
          </div>
      } */}
    </div>

  );
}

export default Authorize;