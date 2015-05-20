package gen

import (
	"container/list"
	"fmt"
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"io"
	"reflect"
	"rng/schema"
	"strconv"
)

type Stage struct {
	name    string
	retName string
	guide   schema.Guide
	stack   *list.List
	parent  *Stage
	content []interface{}
	count   int
}

type generator struct {
	name string
}

type constructor struct {
	bind string
	name string
	typ  string
}

type setter struct {
	bind    string
	field   string
	typ     string
	retName string
}

type field struct {
	name string
	typ  string
}

type Writer struct {
	wr   io.Writer
	root *Stage
	this *Stage
}

func (w *Writer) Write(_s ...string) *Writer {
	var s []interface{}
	for _, v := range _s {
		s = append(s, v)
	}
	w.wr.Write([]byte(fmt.Sprint(s...)))
	return w
}

func (w *Writer) Ln() *Writer {
	return w.Write("\n")
}

func (s *Stage) push(n *Stage) {
	n.name = s.name + strconv.Itoa(s.count)
	s.count++
	n.parent = s
	s.stack.PushFront(s)
}

func (s *Stage) pop() {
	if el := s.stack.Front(); el != nil {
		s.stack.Remove(el)
	}
}

func (s *Stage) top() (ret *Stage) {
	if el := s.stack.Front(); el != nil {
		ret = el.Value.(*Stage)
	}
	return
}

func (s *Stage) writeTo(w *Writer) {
	w.beginStruct(s.name)
	hasMap := false
	for _, _v := range s.content {
		switch v := _v.(type) {
		case *field:
			if v.name != "" {
				w.Write("\t", v.name, " ", v.typ, "\n")
			} else if !hasMap {
				w.Write("\t _attr map[string]interface{}\n")
			}
		}
	}
	w.Write("}").Ln()
	for _, _v := range s.content {
		switch v := _v.(type) {
		case *constructor:
			w.writeConstructor(v.bind, v.name, v.typ)
		case *setter:
			w.writeSetter(v.bind, v.field, v.typ, v.retName)
		case *field: //do nothing
		case *generator:
			w.writeGenerator(v.name)
		default:
			halt.As(100, reflect.TypeOf(v))
		}
	}
}

func (w *Writer) writeSetter(root, _field, typ, retName string) {
	method := "Set" + capitalize(_field)
	if _field != "" {
		w.Write("func (x *", root, ") ", method, " (val ", typ, ") *", retName, "{").Ln()
		w.Write("\t", "x.", _field, "= val").Ln()
	} else {
		w.Write("func (x *", root, ") ", method, " (name string, val ", typ, ") *", retName, "{").Ln()
		w.Write("\t", "x._attr[name] = val").Ln()
	}
	w.Write("return x._parent", "}").Ln()
}

func (w *Writer) beginStruct(name string) {
	w.Write("type " + name + " struct{").Ln()
}

func (w *Writer) writeConstructor(root, method, name string) {
	w.Write("func (x *" + root + ") " + capitalize(method) + "() *" + name).Write(`{
	return &` + name + `{_parent: x}
	}`).Ln()
}

func (w *Writer) writeGenerator(name string) {
	w.Write("func New() *" + name).Write(`{
	return &` + name + `{}
	}`).Ln()
}

func (w *Writer) Begin(g schema.Guide) {
	if w.root == nil {
		s := &Stage{guide: g, stack: list.New(), name: "Stage0"}
		w.root = s
		w.this = s
		s.content = append(s.content, &generator{name: "Stage0"})
	} else {
		assert.For(g != nil, 20)
		switch g.(type) {
		case schema.Choice, schema.Interleave, schema.ZeroOrMore:
			s := &Stage{guide: g, stack: list.New(), retName: w.this.name}
			w.this.push(s)
			s.content = append(s.content, &field{name: "_parent", typ: "*" + w.this.name})
			s.content = append(s.content, &constructor{bind: w.this.name, name: g.String(), typ: s.name})
			w.this = s
		default:
			halt.As(100, reflect.TypeOf(g))
		}
	}
}

func (w *Writer) End() {
	assert.For(w.this != nil, 20)
	if s := w.this.top(); s != nil {
		s.writeTo(w)
	}
	if w.this.stack.Len() == 0 {
		w.this.writeTo(w)
		w.this = w.this.parent
	}
	if w.this != nil {
		w.this.pop()
	}
}

func (w *Writer) Attr(_name string, t string, values ...string) {
	var f *field
	if _name != "" {
		f = &field{name: _name, typ: t}
	} else {
		f = &field{}
	}
	tmp := w.this.content
	w.this.content = nil
	w.this.content = append(w.this.content, f)
	w.this.content = append(w.this.content, tmp...)
	w.this.content = append(w.this.content, &setter{bind: w.this.name, field: _name, typ: t, retName: w.this.retName})
}

func NewWriter(root schema.Element, w io.Writer) (ret *Writer) {
	ret = &Writer{wr: w}
	ret.Write("package ").Write(root.Name()).Ln()
	ret.Begin(root)
	return
}
