package std

import (
	"fmt"
	"github.com/kpmy/rng/schema"
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/fn"
	"strconv"
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
	nsed
}

type empty struct {
	common
}

type name struct {
	common
}

type data struct {
	common
	typed
}

type value struct {
	common
	contented
}

type text struct {
	common
}

type param struct {
	common
}

type ref struct {
	named
	common
}

type extRef struct {
	common
	refed
}

type Identified interface {
	Id(...string) string
}

type common struct {
	parent schema.Guide
	guides []schema.Guide
	id     string
}

func (c *common) Id(i ...string) string {
	if len(i) == 1 {
		c.id = i[0]
	}
	return c.id
}

func (c *common) Parent(p ...schema.Guide) schema.Guide {
	if len(p) == 1 {
		c.parent = p[0]
	}
	return c.parent
}

func (c *common) Add(g schema.Guide) {
	assert.For(g != nil, 20)
	if c.Id() == "" {
		c.Id("0")
	}
	c.guides = append(c.guides, g)
	id := c.Id() + "." + strconv.Itoa(len(c.guides)-1)
	g.(Identified).Id(id)
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

type typed struct {
	typ string
}

func (a *typed) Type(val ...string) string {
	if len(val) == 1 {
		a.typ = val[0]
	}
	return a.typ
}

type contented struct {
	content string
}

func (a *contented) Data(val ...string) string {
	if len(val) == 1 {
		a.content = val[0]
	}
	return a.content
}

type nsed struct {
	ns string
}

func (a *nsed) NS(val ...string) string {
	if len(val) == 1 {
		a.ns = val[0]
	}
	return a.ns
}

type refed struct {
	r string
}

func (a *refed) Href(val ...string) string {
	if len(val) == 1 {
		a.r = val[0]
	}
	return a.r
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
func (x *ref) This() schema.Ref               { return x }
func (x *extRef) This() schema.ExternalRef    { return x }

func (x *start) String() string      { return "grammar" }
func (x *choice) String() string     { return "choice" }
func (x *interleave) String() string { return "interleave" }
func (x *mixed) String() string      { return "mixed" }
func (x *element) String() (ret string) {
	return fmt.Sprint("element", " ", fn.MaybeString("'", x.Name(), "'"))
}
func (x *attribute) String() string {
	return fmt.Sprint("attribute", " ", fn.MaybeString("'", x.Name(), "'"))
}
func (x *group) String() string      { return "group" }
func (x *list) String() string       { return "list" }
func (x *optional) String() string   { return "optional" }
func (x *oneOrMore) String() string  { return "oneOrMore" }
func (x *zeroOrMore) String() string { return "zeroOrMore" }
func (x *except) String() string     { return "except" }
func (x *anyName) String() string    { return "anyName" }
func (x *nsName) String() string     { return "nsName" }
func (x *name) String() string       { return "name" }
func (x *empty) String() string      { return "empty" }
func (x *value) String() string      { return "value" }
func (x *data) String() string {
	return fmt.Sprint("data", " ", fn.MaybeString("type: ", x.Type(), " "))
}
func (x *text) String() string  { return "text" }
func (x *param) String() string { return "param" }
func (x *ref) String() string {
	return fmt.Sprint("ref", " ", fn.MaybeString("'", x.Name(), "'"))
}
func (x *extRef) String() string { return "externalRef" }

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

func Ref() schema.Guide {
	var ret schema.Ref
	ret = &ref{}
	return ret
}

func ExternalRef() schema.Guide {
	var ret schema.ExternalRef
	ret = &extRef{}
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

func CombineAttr(g schema.Guide, d ...string) (ret string) {
	if x, ok := g.(schema.Combined); ok {
		ret = x.Combine(d...)
	}
	return
}

func DataTypeAttr(g schema.Guide, dt ...string) (ret string) {
	if x, ok := g.(schema.DataTyped); ok {
		ret = x.DataType(dt...)
	}
	return
}

func HrefAttr(g schema.Guide, r ...string) (ret string) {
	if x, ok := g.(schema.Referred); ok {
		ret = x.Href(r...)
	}
	return
}
