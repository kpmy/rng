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

func Map(i []schema.Guide, fn func(schema.Guide)) {
	for _, v := range i {
		fn(v)
	}
}
