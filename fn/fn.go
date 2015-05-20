package fn

import (
	"github.com/kpmy/ypk/assert"
	"reflect"
	"rng/schema"
)

type iterable interface {
	List() []schema.Guide
}

func Iterate(g schema.Guide) []schema.Guide {
	i, ok := g.(iterable)
	assert.For(ok, 20, reflect.TypeOf(g))
	return i.List()
}

type Unary func(interface{}, ...interface{}) interface{}

type Bool func(schema.Guide, ...interface{}) bool

func Map(i []schema.Guide, fn Unary, meta ...interface{}) (ret []interface{}) {
	for _, v := range i {
		res := fn(v, meta...)
		assert.For(res != nil, 30)
		ret = append(ret, res)
	}
	return
}

func Filter(i []schema.Guide, fn Bool, meta ...interface{}) (ret []schema.Guide) {
	for _, v := range i {
		if fn(v, meta...) {
			ret = append(ret, v)
		}
	}
	return
}

func Traverse(root schema.Guide, fn Unary, meta ...interface{}) (ret []interface{}) {
	ref := make(map[string]schema.Guide)
	var foo Unary
	bar := func(g schema.Guide) {
		ret = append(ret, fn(g, meta...))
		Map(Iterate(g), foo, meta...)
	}
	foo = func(this interface{}, meta ...interface{}) interface{} {
		switch t := this.(type) {
		case schema.Ref:
			if ref[t.Name()] == nil {
				ref[t.Name()] = t
				bar(t)
			} else {
				ret = append(ret, fn(t, meta...))
			}
		default:
			bar(t.(schema.Guide))
		}
		return this
	}
	bar(root)
	return
}

func Select(root schema.Guide, fn Bool, meta ...interface{}) (ret []schema.Guide) {
	ref := make(map[string]schema.Guide)
	var foo Unary
	bar := func(g schema.Guide) {
		if fn(g, meta...) {
			ret = append(ret, g)
		}
		Map(Iterate(g), foo, meta...)
	}
	foo = func(this interface{}, meta ...interface{}) interface{} {
		switch t := this.(type) {
		case schema.Ref:
			if ref[t.Name()] == nil {
				ref[t.Name()] = t
				bar(t)
			} else {
				if fn(t, meta...) {
					ret = append(ret, t)
				}
			}
		default:
			bar(t.(schema.Guide))
		}
		return this
	}
	bar(root)
	return
}

func Deref(_r schema.Guide) (ret []schema.Guide, cycle bool) {
	ref := make(map[string]schema.Guide)
	if r, ok := _r.(schema.Ref); ok {
		Map(Iterate(r), func(_i interface{}, _ ...interface{}) interface{} {
			switch i := _i.(type) {
			case schema.Ref:
				if ref[i.Name()] == nil {
					ref[i.Name()] = i
					tmp, cycled := Deref(i)
					cycle = cycled
					ret = append(ret, tmp...)
				} else {
					ret = append(ret, i)
					cycle = true
				}
			default:
				ret = append(ret, _i.(schema.Guide))
			}
			return _i
		})
	}
	return
}
