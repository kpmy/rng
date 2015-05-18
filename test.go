package main

import (
	"fmt"
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"reflect"
	"rng/schema"
	"rng/schema/std"
)

func test(start schema.Start) {
	std.NewRider(start).Map(func(_g interface{}) (ret interface{}) {
		switch g := _g.(type) {
		case schema.Choice:
			ret = func() {

			}
			fmt.Println("choice")
		default:
			halt.As(100, reflect.TypeOf(g))
		}
		assert.For(ret != nil, 20)
		return
	})
}
