import { signInWithPopup, FacebookAuthProvider, getAuth, signOut, onAuthStateChanged } from 'firebase/auth';
import { authentication } from '../firebase-config';

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
    displayName: userCredential.user.displayName,
    photoURL: userCredential.user.photoURL,
    squareAccessToken: null
  }

  dispatch(setUser(user));
}

export const logout = async (dispatch) => {
  const auth = getAuth();

  signOut(auth).then(() => {
    dispatch(removeUser());
  }).catch((error) => {
    console.log(error.message);
  });
}

export const setUserWithToken = (token) => (dispatch) => {
  const auth = getAuth();
  const currentUser = auth.currentUser;

  const user = {
    displayName: currentUser.displayName,
    photoURL: currentUser.photoURL,
    squareAccessToken: token
  }

  dispatch(setUser(user));
}

export const setCurrentUser = (user) => dispatch => {
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
