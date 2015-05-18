package std

import (
	"github.com/kpmy/ypk/assert"
	"rng/schema"
)

type start struct {
	common
}

func (s *start) list() []schema.Guide {
	return s.guides
}

type choice struct {
	common
}

type element struct {
	common
}

type attribute struct {
	common
}

type zeroOrMore struct {
	common
}

type oneOrMore struct {
	common
}

type interleave struct {
	common
}

type mixed struct {
	common
}

type group struct {
	common
}

type list struct {
	common
}

type except struct {
	common
}

type optional struct {
	common
}

type common struct {
	parent schema.Guide
	guides []schema.Guide
}

func (c *common) Parent(p ...schema.Guide) schema.Guide {
	if len(p) == 1 {
		c.parent = p[0]
	}
	return c.parent
}

func (c *common) Add(g schema.Guide) {
	assert.For(g != nil, 20)
	c.guides = append(c.guides, g)
	g.Parent(c)
}

func Start() schema.Start {
	return &start{}
}

func Choice() schema.Guide {
	var ret schema.Choice
	ret = &choice{}
	return ret
}

func Interleave() schema.Guide {
	var ret schema.Interleave
	ret = &interleave{}
	return ret
}

func Element() schema.Guide {
	var ret schema.Element
	ret = &element{}
	return ret
}

func Attribute() schema.Guide {
	var ret schema.Attribute
	ret = &attribute{}
	return ret
}

func ZeroOrMore() schema.Guide {
	var ret schema.ZeroOrMore
	ret = &zeroOrMore{}
	return ret
}

func OneOrMore() schema.Guide {
	var ret schema.OneOrMore
	ret = &oneOrMore{}
	return ret
}

func Mixed() schema.Guide {
	var ret schema.Mixed
	ret = &mixed{}
	return ret
}

func Except() schema.Guide {
	var ret schema.Except
	ret = &except{}
	return ret
}

func List() schema.Guide {
	var ret schema.List
	ret = &list{}
	return ret
}

func Group() schema.Guide {
	var ret schema.Group
	ret = &group{}
	return ret
}

func Optional() schema.Guide {
	var ret schema.Optional
	ret = &optional{}
	return ret
}
