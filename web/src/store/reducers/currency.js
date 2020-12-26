import {Currency} from "../actions";

const initial = {
  items: [],
  actual: {},
  loading: false,
}


export default function apps(state = initial, action) {
  switch (action.type) {
    case Currency.BeginLoad:
      return {...state, loading: true}
    case Currency.EndLoad:
      return {...state, loading: false, items: action.payload}
    case Currency.BeginSave:
      return {...state, loading: true}
    case Currency.EndSave:
      return {...state, loading: false}
    case Currency.Edit:
      return {...state, actual: action.payload}
    case Currency.BeginEdit:
      return {...state, loading: true}
    case Currency.EndEdit:
      return {...state, loading: false}
    case Currency.BeginGet:
      return {...state, loading: true}
    case Currency.EndGet:
      return {...state, loading: false, actual: action.payload}
    case Currency.Error:
      return {...state, loading: false}
    default:
      return state
  }
}
