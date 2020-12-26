import {store} from "../store/store";
import {Category, Status} from "../store/actions";
import axios from "../client";


export default class CategoryService {
  constructor() {
    this.dispatch = store.dispatch
  }

  fetch() {
    this.dispatch({type: Category.BeginLoad})
    axios.get("/category/en").then(resp => {
      if(resp.data == null) {
        this.dispatch({type: Category.EndLoad, payload: []})
      }
      else {
        this.dispatch({type: Category.EndLoad, payload: resp.data})
      }
    }).catch(c => {
      this.dispatch({type: Category.Error })
      if(typeof(c.response) !== 'undefined') {
        this.dispatch({type: Status.Error, payload: "The product cannot be loaded: " + c.response.data})
      }
    })
  }
}
