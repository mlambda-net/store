
import {store} from "../store/store";
import axios from "../client";

export default class ImageService {

  constructor() {
    this.dispatch = store.dispatch
  }

  save(data) {
    return new Promise((resolve, reject) => {
      axios.post("/image", data, {
        headers: {
          'Content-Type': "multipart/form-data"
        }
      }).then(resp => {
        resolve(resp.data.url)
      }).catch(c => {
        reject(c)
      })
    })
  }

}
