import {store} from "../store/store";
import {Brand, Status} from "../store/actions";
import axios from "../client";


export default class BrandService {
  constructor() {
    this.dispatch = store.dispatch
  }

  fetch() {
    this.dispatch({type: Brand.BeginLoad})
    axios.get("/brand").then(resp => {
      if(resp.data == null) {
        this.dispatch({type: Brand.EndLoad, payload: []})
      }
      else {
        this.dispatch({type: Brand.EndLoad, payload: resp.data})
      }
    }).catch(c => {
      this.dispatch({type: Brand.Error })
      if(typeof(c.response) !== 'undefined') {
        this.dispatch({type: Status.Error, payload: "The product cannot be loaded: " + c.response.data})
      }
    })
  }
}
