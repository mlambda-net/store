import { combineReducers} from "redux";
import auth from "./auth";
import status from "./status";
import catalog from "./catalog";
import currency from "./currency"
import brand from './brand'
import category from "./category";

export default combineReducers({
  auth, status,catalog, currency, brand, category
})
