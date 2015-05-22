package main

import (
	"flag"
	"fmt"
	"os"
	"rng/fn"
	"rng/loader"
	"rng/schema"
)

var inputName string

const defaultName = "text:p"
const schemaName = "OpenDocument-v1.2-os-schema.rng"

func init() {
	flag.StringVar(&inputName, "i", defaultName, "-i ns:name")
}

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

func main() {
	fmt.Println(inputName)
	if file, err := os.Open(schemaName); err == nil {
		defer file.Close()
		root, _ := loader.Load(file)
		el := fn.Select(root, elementFilter(inputName))[0]
		fmt.Println("(", el.Parent(), ")", el)
		fn.Map(fn.Iterate(el), printChildren)
	}
}
