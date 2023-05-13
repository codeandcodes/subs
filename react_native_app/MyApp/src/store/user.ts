import auth from '@react-native-firebase/auth';
import { LoginManager, AccessToken } from 'react-native-fbsdk-next';

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
  // Attempt login with permissions
  const result = await LoginManager.logInWithPermissions(['public_profile', 'email']);

  if (result.isCancelled) {
    throw 'User cancelled the login process';
  }

  // Once signed in, get the users AccesToken
  const data = await AccessToken.getCurrentAccessToken();

  if (!data) {
    throw 'Something went wrong obtaining access token';
  }

  // Create a Firebase credential with the AccessToken
  const facebookCredential = auth.FacebookAuthProvider.credential(data.accessToken);
  
  // Sign-in the user with the credential
  return auth().signInWithCredential(facebookCredential);
}

export const login = () => async (dispatch) => {
  const userCredential = await signInWithFacebook();

  const user = {
    displayName: userCredential.user.displayName,
    photoURL: userCredential.user.photoURL
  }
  dispatch(setUser(user));
}

export const logout = () => async (dispatch) => {
  LoginManager.logOut();

  dispatch(removeUser());
}

const userReducer = ( state = { user: null }, action) => {
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

export default userReducer;