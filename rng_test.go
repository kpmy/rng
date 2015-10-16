package rng

import (
	"bytes"
	"fmt"
	"github.com/kpmy/rng/auto"
	"github.com/kpmy/rng/fn"
	"github.com/kpmy/rng/gen"
	"github.com/kpmy/rng/loader"
	"github.com/kpmy/rng/schema"
	"github.com/kpmy/ypk/halt"
	"os"
	"testing"
)

func print(_g interface{}, meta ...interface{}) interface{} {
	r, _ := fn.Deref(_g.(schema.Guide))
	if r != nil {
		for _, g := range r {
			fmt.Print(g, fmt.Sprint(meta...))
		}
	} else {
		fmt.Print(_g, fmt.Sprint(meta...))
	}
	return _g
}

func println(_g interface{}, meta ...interface{}) interface{} {
	fmt.Println(_g, fmt.Sprint(meta...))
	return _g
}

func printChildren(_g interface{}, meta ...interface{}) interface{} {
	fmt.Println(_g, "[]")
	g := _g.(schema.Guide)
	fn.Map(fn.Iterate(g), print, " ")
	fmt.Println()
	fmt.Println()
	return _g
}

func elementFilter(name string) fn.Bool {
	return func(g schema.Guide, _ ...interface{}) bool {
		e, ok := g.(schema.Element)
		return ok && e.Name() == name
	}
}

func childAndDeref(root schema.Guide) (ret []schema.Guide) {
	fn.Map(fn.Iterate(root), func(_g interface{}, meta ...interface{}) interface{} {
		r, _ := fn.Deref(_g.(schema.Guide))
		if r != nil {
			for _, g := range r {
				ret = append(ret, g)
			}
		} else {
			ret = append(ret, _g.(schema.Guide))
		}
		return _g
	})
	return
}

func TestSchema(t *testing.T) {
	if file, err := os.Open("relaxng.rng"); err == nil {
		defer file.Close()
		root, _ := loader.Load(file)
		el := fn.Select(root, elementFilter("element"))[0]
		fmt.Println("(", el.Parent(), ")", el)
		fn.Map(fn.Iterate(el), printChildren)
	}
}

func TestSchemaManual(t *testing.T) {
	if file, err := os.Open("relaxng.rng"); err == nil {
		defer file.Close()
		root, _ := loader.Load(file)
		el := fn.Select(root, elementFilter("element"))[0]
		fmt.Println(el)
		ce := childAndDeref(el)
		//fn.Map(ce, printChildren)
		fn.Map(fn.Iterate(ce[0]), printChildren)
		fn.Map(fn.Iterate(fn.Iterate(ce[0])[1]), printChildren)
		fn.Map(fn.Iterate(fn.Iterate(fn.Iterate(ce[0])[1])[0]), printChildren)
	}
}

func TestAuto(t *testing.T) {
	auto.Null()
}

//что-то сложновато
func TestSchemaToInstance(t *testing.T) {
	if file, err := os.Open("relaxng.rng"); err == nil {
		defer file.Close()
		root, _ := loader.Load(file)
		el := fn.Select(root, elementFilter("element"))[0]
		fmt.Println(el)
		buf := bytes.NewBuffer(nil)
		wr := gen.NewWriter(el.(schema.Element), buf)
		var traverse fn.Unary
		traverse = func(_g interface{}, meta ...interface{}) interface{} {
			//fmt.Println(string(buf.Bytes()))
			switch g := _g.(type) {
			case schema.Choice, schema.Interleave, schema.ZeroOrMore, schema.Optional, schema.OneOrMore:
				wr.Begin(g.(schema.Guide))
				fn.Map(fn.Iterate(g.(schema.Guide)), traverse, meta...)
				wr.End()
			case schema.Element:
			case schema.Attribute:
				//get attr typ
				wr.Attr(g.Name(), "string")
			case schema.Ref:
				gl, _ := fn.Deref(g)
				fn.Map(gl, traverse, meta...)
			default:
				halt.As(100, g)
			}
			return _g
		}
		fn.Map(fn.Iterate(el), traverse)
		wr.End()
		fmt.Println(string(buf.Bytes()))
	}
}
