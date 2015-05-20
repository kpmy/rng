package rng

import (
	"fmt"
	"os"
	"rng/loader"
	"rng/mappers"
	"rng/schema"
	"testing"
)

func print(_g interface{}, meta ...interface{}) interface{} {
	r := mappers.Deref(_g.(schema.Guide))
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
	fmt.Println(_g)
	g := _g.(schema.Guide)
	mappers.Map(mappers.Iterate(g), print, " ")
	fmt.Println()
	return _g
}

func elementFilter(name string) mappers.Bool {
	return func(g schema.Guide, _ ...interface{}) bool {
		e, ok := g.(schema.Element)
		return ok && e.Name() == name
	}
}

func TestSchema(t *testing.T) {
	if file, err := os.Open("relaxng.rng"); err == nil {
		defer file.Close()
		root, _ := loader.Load(file)
		el := mappers.Select(root, elementFilter("element"))[0]
		fmt.Println("(", el.Parent(), ")", el)
		mappers.Map(mappers.Iterate(el), printChildren)
	}
}

func TestSchemaToInstance(t *testing.T) {
	if file, err := os.Open("relaxng.rng"); err == nil {
		defer file.Close()
		root, _ := loader.Load(file)
		_ = mappers.Select(root, elementFilter("element"))[0]

	}
}
