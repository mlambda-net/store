
import axios from 'axios';
import {store} from "./store/store";
import {Auth} from "./store/actions";



axios.defaults.baseURL = process.env.REACT_APP_API_URL;

axios.interceptors.request.use((config) =>{
  const {auth} = store.getState()


  if(!auth.data.user.expired) {
    config.headers.Authorization = "bearer " + auth.data.user.access_token
  } else {
    auth.data.renew()
    store.dispatch({type: Auth.SetAuth, payload: auth.manager.getUser()})
  }
  return config;
})

axios.interceptors.response.use((response) => {
  return response
})

export default axios
