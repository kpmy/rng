package std

import (
	"github.com/kpmy/ypk/assert"
	"rng/schema"
)

type start struct {
	common
}

type choice struct {
	common
}

type element struct {
	common
	named
}

type attribute struct {
	common
	named
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

type anyName struct {
	common
}

type nsName struct {
	common
}

type empty struct {
	common
}

type name struct {
	common
}

type data struct {
	common
}

type value struct {
	common
}

type text struct {
	common
}

type param struct {
	common
}

type common struct {
	schema.Guide
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

func (c *common) List() []schema.Guide {
	return c.guides
}

type named struct {
	name string
}

func (n *named) Name(val ...string) string {
	if len(val) == 1 {
		n.name = val[0]
	}
	return n.name
}

func (x *start) This() schema.Start           { return x }
func (x *choice) This() schema.Choice         { return x }
func (x *interleave) This() schema.Interleave { return x }
func (x *mixed) This() schema.Mixed           { return x }
func (x *element) This() schema.Element       { return x }
func (x *attribute) This() schema.Attribute   { return x }
func (x *group) This() schema.Group           { return x }
func (x *list) This() schema.List             { return x }
func (x *optional) This() schema.Optional     { return x }
func (x *oneOrMore) This() schema.OneOrMore   { return x }
func (x *zeroOrMore) This() schema.ZeroOrMore { return x }
func (x *except) This() schema.Except         { return x }
func (x *anyName) This() schema.AnyName       { return x }
func (x *nsName) This() schema.NSName         { return x }
func (x *name) This() schema.Name             { return x }
func (x *empty) This() schema.Empty           { return x }
func (x *value) This() schema.Value           { return x }
func (x *data) This() schema.Data             { return x }
func (x *text) This() schema.Text             { return x }
func (x *param) This() schema.Param           { return x }

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

func AnyName() schema.Guide {
	var ret schema.AnyName
	ret = &anyName{}
	return ret
}

func NSName() schema.Guide {
	var ret schema.NSName
	ret = &nsName{}
	return ret
}

func Name() schema.Guide {
	var ret schema.Name
	ret = &name{}
	return ret
}

func Data() schema.Guide {
	var ret schema.Data
	ret = &data{}
	return ret
}

func Value() schema.Guide {
	var ret schema.Value
	ret = &value{}
	return ret
}

func Empty() schema.Guide {
	var ret schema.Empty
	ret = &empty{}
	return ret
}

func Text() schema.Guide {
	var ret schema.Text
	ret = &text{}
	return ret
}

func Param() schema.Guide {
	var ret schema.Param
	ret = &param{}
	return ret
}

func NameAttr(g schema.Guide, n ...string) (ret string) {
	if x, ok := g.(schema.Named); ok {
		ret = x.Name(n...)
	}
	return
}

func TypeAttr(g schema.Guide, t ...string) (ret string) {
	if x, ok := g.(schema.Typed); ok {
		ret = x.Type(t...)
	}
	return
}

func NSAttr(g schema.Guide, ns ...string) (ret string) {
	if x, ok := g.(schema.NSed); ok {
		ret = x.NS(ns...)
	}
	return
}

func CharDataAttr(g schema.Guide, d ...string) (ret string) {
	if x, ok := g.(schema.Contented); ok {
		ret = x.Data(d...)
	}
	return
}
func DataTypeAttr(g schema.Guide, dt ...string) (ret string) {
	if x, ok := g.(schema.DataTyped); ok {
		ret = x.DataType(dt...)
	}
	return
}
