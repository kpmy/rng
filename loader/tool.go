package loader

import (
	"encoding/xml"
	"fmt"
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"io"
	"rng/schema"
	"rng/schema/std"
)

const urn_rng_1 = "http://relaxng.org/ns/structure/1.0"

var Constructors map[string]func() schema.Guide

func init() {
	Constructors = make(map[string]func() schema.Guide)
	Constructors["choice"] = std.Choice
	Constructors["element"] = std.Element
	Constructors["interleave"] = std.Interleave
	Constructors["attribute"] = std.Attribute
	Constructors["zeroOrMore"] = std.ZeroOrMore
	Constructors["oneOrMore"] = std.OneOrMore
	Constructors["group"] = std.Group
	Constructors["list"] = std.List
	Constructors["except"] = std.Except
	Constructors["mixed"] = std.Mixed
	Constructors["optional"] = std.Optional
	Constructors["anyName"] = std.AnyName
	Constructors["nsName"] = std.NSName
	Constructors["data"] = std.Data
	Constructors["text"] = std.Text
	Constructors["name"] = std.Name
	Constructors["empty"] = std.Empty
	Constructors["value"] = std.Value
	Constructors["param"] = std.Param
	Constructors["ref"] = std.Ref
	Constructors["externalRef"] = std.ExternalRef
}

func Construct(name xml.Name) (ret schema.Guide) {
	assert.For(name.Space == urn_rng_1, 20, name)
	fn := Constructors[name.Local]
	assert.For(fn != nil, 40, name)
	ret = fn()
	assert.For(ret != nil, 60, name)
	return
}

type Cached struct {
	node *Node
	root schema.Guide
}

type Walker struct {
	root  *Node
	cache map[string]*Cached
	start schema.Start
	pos   schema.Guide
}

func (w *Walker) Init() *Walker {
	w.cache = make(map[string]*Cached)
	w.start = std.Start()
	w.pos = w.start
	return w
}

func (w *Walker) GrowDown(g schema.Guide) {
	w.pos.Add(g)
	g.Parent(w.pos)
	w.pos = g
}

func (w *Walker) Grow(g schema.Guide) {
	w.pos.Add(g)
}

func (w *Walker) Up() {
	w.pos = w.pos.Parent()
}

func (w *Walker) forEach(n *Node, do func(w *Walker, n *Node)) {
	for _, v := range n.Inner {
		do(w, &v)
	}
}

func traverseWrap() func(w *Walker, n *Node) {
	return func(w *Walker, n *Node) {
		w.traverse(n)
	}
}

func (w *Walker) traverse(n *Node) {
	var (
		this    schema.Guide
		skip    *bool
		skipped = func() {
			s := true
			skip = &s
		}
		important = func() {
			s := false
			skip = &s
		}
	)

	switch n.XMLName.Local {
	//structure elements
	case "grammar":
		skipped()
		if start := n.FindByXMLName("start"); start != nil {
			w.forEach(start, traverseWrap())
		}
	case "ref":
		skipped()
		if ref := w.root.FindByName(n.Name); ref != nil {
			if cached := w.cache[n.Name]; cached == nil {
				this = Construct(n.XMLName)
				{
					std.NameAttr(this, n.Name)
					//std.RefAttr(this, cached.root)
				}
				w.cache[n.Name] = &Cached{node: ref, root: this}
				w.GrowDown(this)
				w.forEach(ref, traverseWrap())
				w.Up()
			} else {
				w.Grow(cached.root)
			}
		} else {
			halt.As(100, "ref not found", n.Name)
		}
	//content elements
	case "element", "attribute", "data", "text", "value", "name", "param":
		fallthrough
	//constraint elements
	case "choice", "interleave", "optional", "zeroOrMore", "oneOrMore", "group", "list", "mixed", "except", "anyName", "nsName", "empty", "externalRef":
		important()
		this = Construct(n.XMLName)
		{
			std.NameAttr(this, n.Name)
			std.CharDataAttr(this, n.Data())
			std.TypeAttr(this, n.Type)
			std.NSAttr(this, n.NS)
			std.DataTypeAttr(this, n.DataType)
			std.CombineAttr(this, n.Combine)
			std.HrefAttr(this, n.Href)
		}
		w.GrowDown(this)
		w.forEach(n, traverseWrap())
		w.Up()
	//skipped elements
	case "description": //dc:descriprion do nothing
	default:
		halt.As(100, n.XMLName)
	}
	if skip != nil {
		assert.For(*skip || this != nil, 60, "no result for ", n.XMLName)
	} else if this == nil {
		fmt.Println("unhandled", n.XMLName)
	}
}

func Load(input io.Reader) (ret schema.Start, err error) {
	//only xml rng format for now
	xrd := &XMLReader{rd: input}
	if data, err := xrd.Read(); err == nil {
		w := &Walker{root: data}
		w.Init()
		w.traverse(data)
		ret = w.start
	}
	return
}
