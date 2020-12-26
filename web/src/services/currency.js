
import {store} from "../store/store";
import { Currency, Status} from "../store/actions";
import axios from "../client";

export default class CurrencyService {

  constructor() {
    this.dispatch = store.dispatch
  }

  fetch() {
    this.dispatch({type: Currency.BeginLoad})
    axios.get("/currency").then(resp => {
      if(resp.data == null) {
        this.dispatch({type: Currency.EndLoad, payload: []})
      }
      else {
        this.dispatch({type: Currency.EndLoad, payload: resp.data})
      }
    }).catch(c => {
      this.dispatch({type: Currency.Error })
      if(typeof(c.response) !== 'undefined') {
        this.dispatch({type: Status.Error, payload: "The product cannot be loaded: " + c.response.data})
      }
    })
  }
}
