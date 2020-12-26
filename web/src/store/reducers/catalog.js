import {Catalog} from "../actions";

const initial = {
  items: [],
  actual: {},
  loading: false,
}


export default function apps (state = initial, action) {
  switch (action.type) {
    case Catalog.BeginLoad:
      return {...state, loading: true}
    case Catalog.EndLoad:
      return {...state, loading: false, items: action.payload}
    case Catalog.BeginSave:
      return {...state, loading: true}
    case Catalog.EndSave:
      return {...state, loading: false}
    case Catalog.Edit:
      return {...state, actual: action.payload}
    case Catalog.BeginEdit:
      return {...state, loading: true}
    case Catalog.EndEdit:
      return {...state, loading: false}
    case Catalog.BeginGet:
      return {...state, loading: true}
    case Catalog.EndGet:
      return {...state, loading: false, actual: action.payload}
    case Catalog.Error:
      return {...state, loading: false}

    default:
      return state
  }
}
