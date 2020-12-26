import Catalog from "./pages/catalog";

const routes = [
  {
    name: 'global',
    actual: 'catalog',
    routes: [
      {path:"catalog", name: "catalog", component: Catalog},

    ]
  }

]

export default routes
