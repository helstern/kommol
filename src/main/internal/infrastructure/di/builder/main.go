package builder

import (
	"github.com/sarulabs/di/v2"
	"reflect"
)

func typeOf(r interface{}) func() string {
	runtimeType := reflect.TypeOf(r).String()
	return func() string { return runtimeType }
}

func TypeName(r interface{}) string {
	return reflect.TypeOf(r).String()
}

type add struct {
	getName func() string
}

func (t add) Add(builder *di.Builder, def di.Def) error {
	def.Name = t.getName()
	return builder.Add(def)
}

func withName(r interface{}) struct{ add } {
	t := struct{ add }{}
	t.getName = typeOf(r)

	return t
}
