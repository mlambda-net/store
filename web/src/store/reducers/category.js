import {Category} from "../actions";

const initial = {
  items: [],
  actual: {},
  loading: false,
}


export default function category(state = initial, action) {
  switch (action.type) {
    case Category.BeginLoad:
      return {...state, loading: true}
    case Category.EndLoad:
      return {...state, loading: false, items: action.payload}
    case Category.BeginSave:
      return {...state, loading: true}
    case Category.EndSave:
      return {...state, loading: false}
    case Category.Edit:
      return {...state, actual: action.payload}
    case Category.BeginEdit:
      return {...state, loading: true}
    case Category.EndEdit:
      return {...state, loading: false}
    case Category.BeginGet:
      return {...state, loading: true}
    case Category.EndGet:
      return {...state, loading: false, actual: action.payload}
    case Category.Error:
      return {...state, loading: false}

    default:
      return state
  }
}
