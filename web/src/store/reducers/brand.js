import {Brand} from "../actions";

const initial = {
  items: [],
  actual: {},
  loading: false,
}


export default function brand(state = initial, action) {
  switch (action.type) {
    case Brand.BeginLoad:
      return {...state, loading: true}
    case Brand.EndLoad:
      return {...state, loading: false, items: action.payload}
    case Brand.BeginSave:
      return {...state, loading: true}
    case Brand.EndSave:
      return {...state, loading: false}
    case Brand.Edit:
      return {...state, actual: action.payload}
    case Brand.BeginEdit:
      return {...state, loading: true}
    case Brand.EndEdit:
      return {...state, loading: false}
    case Brand.BeginGet:
      return {...state, loading: true}
    case Brand.EndGet:
      return {...state, loading: false, actual: action.payload}
    case Brand.Error:
      return {...state, loading: false}

    default:
      return state
  }
}
