package specs

import (
  "github.com/mlambda-net/net/pkg/spec"
)


func ById(id string) spec.Expression {
	return spec.NewEval("id", id, "=")
}

func ByName(name string)  spec.Expression {
  return spec.NewEval("name","%" + name + "%", "LIKE " )
}

func ByCode ( code string) spec.Expression  {
  return spec.NewEval("code", code, "=" )
}

func ByLookupName(name string) spec.Expression  {
  return spec.NewEval("name", name, "=")
}

func ByLookupParent(value string) spec.Expression  {
  return spec.NewEval("parent", value, "=")
}
