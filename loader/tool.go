package loader

import (
	"github.com/kpmy/ypk/halt"
	"io"
)

type Walker struct {
	root *Node
}

func (w *Walker) forEach(n *Node, do func(w *Walker, n *Node)) {
	for _, v := range n.Inner {
		do(w, &v)
	}
}

func traverseWrap() func(w *Walker, n *Node) {
	return func(w *Walker, n *Node) {
		w.traverse(n)
	}
}

func (w *Walker) traverse(n *Node) {
	switch n.XMLName.Local {
	case "grammar":
		if start := n.FindByXMLName("start"); start != nil {
			w.forEach(start, traverseWrap())
		}
	case "ref":
		if ref := w.root.FindByName(n.Name); ref != nil {
			w.forEach(ref, traverseWrap())
		} else {
			halt.As(100, "ref not found", n.Name)
		}
	default:
		halt.As(100, n.XMLName)
	}
}

func Load(input io.Reader) {
	//only xml rng format for now
	xrd := &XMLReader{rd: input}
	if data, err := xrd.Read(); err == nil {
		w := &Walker{root: data}
		w.traverse(data)
	}
}
