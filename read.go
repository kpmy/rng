package main

import (
	"flag"
	"os"
	"rng/loader"
)

//const defaultName = "OpenDocument-v1.2-os-schema.rng"

const defaultName = "relaxng.rng"

var inputName string

func init() {
	flag.StringVar(&inputName, "i", "", "-i <name.rng>")
}

func main() {
	flag.Parse()
	if inputName == "" {
		inputName = defaultName
	}
	if file, err := os.Open(inputName); err == nil {
		defer file.Close()
		root, _ := loader.Load(file)
		test(root)
	}
}
