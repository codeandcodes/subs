import { signInWithPopup, FacebookAuthProvider, getAuth, signOut } from 'firebase/auth';
import { authentication } from '../firebase-config';
import { addSquareAccessToken, registerUser } from '../api/user';

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

  const user = {
    emailAddress: userCredential.user.email,
    displayName: userCredential.user.displayName,
    photoUrl: userCredential.user.photoURL,
    fbUserId: userCredential.user.uid
  }

  const registeredUser = await registerUser(user);

  user.osUserId = registeredUser.os_user_id;

  localStorage.setItem('user', JSON.stringify(user));

  dispatch(setUser(user));
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

export const setUserWithToken = (token, osUserId) => async (dispatch) => {
  const loggedInUser = JSON.parse(localStorage.getItem('user'));

  loggedInUser.squareAccessToken = token;
  loggedInUser.osUserId = osUserId;

  localStorage.setItem('user', JSON.stringify(loggedInUser));

  const addTokenResponse = await addSquareAccessToken(token, osUserId);
  // handle response

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
