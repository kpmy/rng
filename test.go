package main

import (
	"fmt"
	"github.com/kpmy/ypk/halt"
	"reflect"
	"rng/mappers"
	"rng/schema"
)

var level int

func tab() (ret string) {
	for i := 0; i < level; i++ {
		ret = fmt.Sprint(ret, " ")
	}
	return
}

func verbose(_g schema.Guide) {
	level++
	switch g := _g.(type) {
	case schema.Start:
		fmt.Println(tab(), "root")
	case schema.Choice:
		fmt.Println(tab(), "has one of")
	case schema.Element:
		fmt.Println(tab(), "element", g.Name())
	case schema.Attribute:
		fmt.Println(tab(), "attribute", g.Name())
	case schema.Interleave:
		fmt.Println(tab(), "some of")
	case schema.ZeroOrMore:
		fmt.Println(tab(), "zero or more")
	case schema.OneOrMore:
		fmt.Println(tab(), "one or more")
	case schema.Optional:
		fmt.Println(tab(), "optional")
	case schema.Group:
		fmt.Println(tab(), "group")
	case schema.AnyName:
		fmt.Println(tab(), "any name")
	case schema.Except:
		fmt.Println(tab(), "except")
	case schema.NSName:
		fmt.Println(tab(), "ns name")
	case schema.Text:
		fmt.Println(tab(), "text")
	case schema.Data:
		fmt.Println(tab(), "data")
	case schema.Value:
		fmt.Println(tab(), "value")
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
	default:
		halt.As(100, reflect.TypeOf(g))
	}
	mappers.Map(mappers.Iterate(_g), verbose)
	level--
}

func test(start schema.Start) {
	verbose(start)
}
