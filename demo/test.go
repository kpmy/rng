package main

import (
	"fmt"
	"github.com/kpmy/rng/fn"
	"github.com/kpmy/rng/schema"
	"github.com/kpmy/rng/schema/std"
	"github.com/kpmy/ypk/halt"
	"reflect"
)

var level int
var passed map[string]schema.Guide = make(map[string]schema.Guide)

func tab() (ret string) {
	for i := 0; i < level; i++ {
		ret = fmt.Sprint(ret, " ")
	}
	return
}

func verbose(_g interface{}, meta ...interface{}) (ret interface{}) {
	level++
	fmt.Print(meta...)
	delim := " "
	switch g := _g.(type) {
	case schema.Start:
		fmt.Println(tab(), "$start", g)
	case schema.Choice:
		fmt.Println(tab(), "choice")
	case schema.Element:
		fmt.Println(tab(), "$element", g.Name())
	case schema.Attribute:
		fmt.Println(tab(), "$attribute", g.Name())
	case schema.Interleave:
		fmt.Println(tab(), "interleave")
	case schema.ZeroOrMore:
		fmt.Println(tab(), "zero-or-more")
	case schema.OneOrMore:
		fmt.Println(tab(), "one-or-more")
	case schema.Optional:
		fmt.Println(tab(), "optional")
	case schema.Group:
		fmt.Println(tab(), "group")
	case schema.AnyName:
		fmt.Println(tab(), "any-name")
	case schema.Except:
		fmt.Println(tab(), "except")
	case schema.NSName:
		fmt.Println(tab(), "ns-name", g.NS())
	case schema.Text:
		fmt.Println(tab(), "text")
	case schema.Data:
		fmt.Print(tab(), "data")
		if g.Type() != "" {
			fmt.Print(" ", "type ", g.Type())
		}
		fmt.Println()
	case schema.Value:
		fmt.Println(tab(), "value", g.Data())
	case schema.Name:
		fmt.Println(tab(), "name")
	case schema.Empty:
		fmt.Println(tab(), "empty")
	case schema.List:
		fmt.Println(tab(), "list")
	case schema.Mixed:
		fmt.Println(tab(), "mixed")
	case schema.Param:
		fmt.Println(tab(), "param")
	case schema.Ref:
		fmt.Println(tab(), "ref", g.Name())
	case schema.ExternalRef:
		fmt.Println(tab(), "externalRef", g.Href())
	default:
		halt.As(100, reflect.TypeOf(g))
	}
	if id := _g.(std.Identified).Id(); passed[id] == nil {
		passed[id] = _g.(schema.Guide)
		fn.Map(fn.Iterate(_g.(schema.Guide)), verbose, delim)
	}
	level--
	return _g
}

func print(_g interface{}, meta ...interface{}) interface{} {
	fmt.Println(_g)
	return _g
}

func elementFilter(name string) fn.Bool {
	return func(g schema.Guide, _ ...interface{}) bool {
		e, ok := g.(schema.Element)
		return ok && e.Name() == name
	}
}

func test(start schema.Start) {
	verbose(start)
	fmt.Println("---")
	fn.Traverse(start, print)
	fmt.Println("---")
}
