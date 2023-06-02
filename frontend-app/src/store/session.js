import { signInWithPopup, FacebookAuthProvider, getAuth, signOut } from 'firebase/auth';
import { authentication } from '../firebase-config';
import { addSquareAccessToken, registerUser, loginUser, getUser } from '../api/user';

const SET_USER = 'setUser';
const REMOVE_USER = 'removeUser';

const setUser = (user) => {
  return {
    type: SET_USER,
    payload: user
  }
}
const removeUser = () => {
  return {
    type: REMOVE_USER
  }
}

const signInWithFacebook = async () => {
  const provider = new FacebookAuthProvider();

  return signInWithPopup(authentication, provider)
    .then((result) => result)
    .catch((err) => console.log(err.message)); 
}

export const login = () => async (dispatch) => {
  const userCredential = await signInWithFacebook();

  const userState = {
    emailAddress: userCredential.user.email,
    displayName: userCredential.user.displayName,
    photoUrl: userCredential.user.photoURL,
    fbUserId: userCredential.user.uid
  };

  // clunky for now to pass in user to registerUser method
  const user = {
    emailAddress: userCredential.user.email,
    displayName: userCredential.user.displayName,
    photoUrl: userCredential.user.photoURL,
    fbUserId: userCredential.user.uid
  }

  // check to see if user is already registered; TODO: check if square access token is also saved
  const osUserId = await getUser(userCredential.user.email)
    .then((res) => {
      userState.hasSquareAccessToken = res.has_square_access_token;
      userState.osUserId = res.os_user_id;
      return res.os_user_id;
    });

  if (osUserId) {
    await loginUser(user.osUserId);
  } else {
    const registeredUser = await registerUser(user);

    user.osUserId = registeredUser.os_user_id;

    await loginUser(user.osUserId);
  }

  // temp: use local storage to persist user info for redirect from square oauth
  localStorage.setItem('user', JSON.stringify(user));

  dispatch(setUser(userState));
}

export const logout = () => async (dispatch) => {
  const auth = getAuth();

  signOut(auth).then(() => {
    localStorage.removeItem('user');
    
    dispatch(removeUser());
  }).catch((error) => {
    console.log(error.message);
  });

}

export const setUserWithToken = (token) => async (dispatch) => {
  // ** don't need this from local storage anymore?
  const loggedInUser = JSON.parse(localStorage.getItem('user'));

  // loggedInUser.squareAccessToken = token;
  // loggedInUser.osUserId = osUserId;

  // localStorage.setItem('user', JSON.stringify(loggedInUser));

  const addTokenResponse = await addSquareAccessToken(token);
  // handle response
  console.log(addTokenResponse);
  dispatch(setUser(loggedInUser));
}

export const setCurrentUser = (user) => (dispatch) => {  
  dispatch(setUser(user));
}

const sessionReducer = ( state = { user: null }, action) => {
  let newState;

  switch (action.type) {
    case SET_USER:
      newState = Object.assign({}, state);
      newState.user = action.payload
      return newState;
    case REMOVE_USER:
      newState = Object.assign({}, state);
      newState.user = null;
      return newState;
    default:
      return state;
  }
}

export default sessionReducer;
