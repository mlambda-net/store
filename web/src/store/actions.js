

const  Auth = {
  SetAuth: "SET_AUTH",
  Profile: "SET_PROFILE",
  Logout: "LOGOUT"
}


const Status = {
  Error: "STATUS_ERROR",
  Info : "STATUS_INFO",
  Warning: "STATUS_WARNING",
  Close: "STATUS_CLOSE",
}

const Catalog = {
  BeginSave: "CATALOG_BEGIN_SAVE",
  EndSave: "CATALOG_END_SAVE",
  BeginLoad: "CATALOG_BEGIN_LOAD",
  EndLoad: "CATALOG_END_LOAD",
  Edit: "CATALOG_EDIT",
  BeginEdit: "CATALOG_BEGIN_EDIT",
  EndEdit: "CATALOG_EDIT_END",
  BeginGet: "CATALOG_BEGIN_GET",
  EndGet: "CATALOG_END_GET",
  Error: "CATALOG_END_ERROR"
}

const Currency = {
  BeginSave: "CURRENCY_BEGIN_SAVE",
  EndSave: "CURRENCY_END_SAVE",
  BeginLoad: "CURRENCY_BEGIN_LOAD",
  EndLoad: "CURRENCY_END_LOAD",
  Edit: "CURRENCY_EDIT",
  BeginEdit: "CURRENCY_BEGIN_EDIT",
  EndEdit: "CURRENCY_EDIT_END",
  BeginGet: "CURRENCY_BEGIN_GET",
  EndGet: "CURRENCY_END_GET",
  Error: "CURRENCY_END_ERROR"
}

const Brand = {
  BeginSave: "BRAND_BEGIN_SAVE",
  EndSave: "BRAND_END_SAVE",
  BeginLoad: "BRAND_BEGIN_LOAD",
  EndLoad: "BRAND_END_LOAD",
  Edit: "BRAND_EDIT",
  BeginEdit: "BRAND_BEGIN_EDIT",
  EndEdit: "BRAND_EDIT_END",
  BeginGet: "BRAND_BEGIN_GET",
  EndGet: "BRAND_END_GET",
  Error: "BRAND_END_ERROR"
}

const Category = {
  BeginSave: "CATEGORY_BEGIN_SAVE",
  EndSave: "CATEGORY_END_SAVE",
  BeginLoad: "CATEGORY_BEGIN_LOAD",
  EndLoad: "CATEGORY_END_LOAD",
  Edit: "CATEGORY_EDIT",
  BeginEdit: "CATEGORY_BEGIN_EDIT",
  EndEdit: "CATEGORY_EDIT_END",
  BeginGet: "CATEGORY_BEGIN_GET",
  EndGet: "CATEGORY_END_GET",
  Error: "CATEGORY_END_ERROR"
}


export {
  Auth, Status, Catalog, Currency, Brand, Category
}
