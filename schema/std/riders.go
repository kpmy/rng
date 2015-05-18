package std

import (
	"github.com/kpmy/ypk/halt"
	"rng/schema"
)

type rideable interface {
	list() []schema.Guide
}

type rider struct {
	base schema.Guide
	semi []interface{}
}

func (r *rider) Map(fn func(interface{}) interface{}) schema.Rider {
	var s []interface{}
	for _, v := range r.semi {
		s = append(s, fn(v))
	}
	return &rider{base: r.base, semi: s}
}

func NewRider(_g schema.Guide) (ret schema.Rider) {
	if g, ok := _g.(rideable); ok {
		rd := &rider{base: _g}
		for _, v := range g.list() {
			rd.semi = append(rd.semi, v)
		}
		ret = rd
	} else {
		halt.As(100)
	}
	return
}
