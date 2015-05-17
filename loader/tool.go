package loader

import (
	"io"
)

func Load(input io.Reader) {
	//only xml rng format for now
	xrd := &XMLReader{rd: input}
	xrd.Read()
}
