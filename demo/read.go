package main

import (
	"flag"
	"github.com/kpmy/rng/loader"
	"os"
)

//const defaultName = "OpenDocument-v1.2-os-schema.rng"

//const defaultName = "relaxng.rng"

const defaultName = "testSuite.rng"

var inputName string

func init() {
	flag.StringVar(&inputName, "i", defaultName, "-i <name.rng>")
}

func main() {
	flag.Parse()
	if file, err := os.Open(inputName); err == nil {
		defer file.Close()
		root, _ := loader.Load(file)
		test(root)
	}
}
