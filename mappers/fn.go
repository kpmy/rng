package mappers

import (
	"github.com/kpmy/ypk/assert"
	"reflect"
	"rng/schema"
)

type Iterable interface {
	List() []schema.Guide
}

func Iterate(g schema.Guide) []schema.Guide {
	i, ok := g.(Iterable)
	assert.For(ok, 20, reflect.TypeOf(g))
	return i.List()
}

type Unary func(interface{}, ...interface{}) interface{}

func Map(i []schema.Guide, fn Unary, meta ...interface{}) (ret []interface{}) {
	for _, v := range i {
		ret = append(ret, fn(v, meta...))
	}
	return
}
