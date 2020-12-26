import axios from "../client";
import {Status, Catalog} from "../store/actions";

export default class CatalogService {

  constructor(dispatch) {
    this.dispatch = dispatch
  }

  fetch(filter) {
    this.dispatch({type: Catalog.BeginLoad})
    axios.get("/product").then(resp => {
      if(resp.data == null) {
        this.dispatch({type: Catalog.EndLoad, payload: []})
      }
      else {
        this.dispatch({type: Catalog.EndLoad, payload: resp.data})
      }
    }).catch(c => {
      this.dispatch({type: Catalog.Error })
      if(typeof(c.response) !== 'undefined') {
        this.dispatch({type: Status.Error, payload: "The product cannot be loaded: " + c.response.data})
      }
    })
  }

  get(id) {
    this.dispatch({type: Catalog.BeginGet})
    axios.get("/product/" + id).then(resp => {
      this.dispatch({type: Catalog.EndGet, payload: resp.data})
    }).catch(c => {
      this.dispatch({type: Catalog.EndGet})
      this.dispatch({type: Status.Error, payload: "The product cannot be loaded: " + c.response.data})
    })
  }

  save(data) {
    return new Promise((resolve, reject) => {
      this.dispatch({type: Catalog.BeginSave})
      axios.post("/product", data)
        .then(r => {
          this.dispatch({type: Catalog.EndSave})
          this.dispatch({type: Status.Info, payload: "Saving the product " + data.name})
          resolve(r)
        })
        .catch(c => {
          this.dispatch({type: Catalog.EndSave})
          this.dispatch({type: Status.Error, payload: "The product cannot be added: " + c.response.data})
          reject(c)
        })
    })
  }


  update(data) {
    return new Promise((resolve, reject) => {
      this.dispatch({type: Catalog.BeginEdit})
      axios.put("/product", data)
        .then(r => {
          this.dispatch({type: Catalog.EndEdit})
          this.dispatch({type: Status.Info, payload: "Updating the product " + data.name})
          resolve(r)
        })
        .catch(c => {
          this.dispatch({type: Catalog.EndEdit})
          this.dispatch({type: Status.Error, payload: "The product " + data.name + " cannot be edited: " + c.response.data})
          reject(c)
        })
    })
  }


}
