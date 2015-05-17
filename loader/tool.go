package loader

import (
	"fmt"
	"github.com/kpmy/ypk/halt"
	"io"
)

type Walker struct {
	root  *Node
	cache map[string]*Node
}

func (w *Walker) Init() *Walker {
	w.cache = make(map[string]*Node)
	return w
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
	//structure elements
	case "grammar":
		if start := n.FindByXMLName("start"); start != nil {
			w.forEach(start, traverseWrap())
		}
	case "ref":
		if ref := w.root.FindByName(n.Name); ref != nil {
			if w.cache[n.Name] == nil {
				w.cache[n.Name] = ref
				w.forEach(ref, traverseWrap())
			}
		} else {
			halt.As(100, "ref not found", n.Name)
		}
	case "choice":
		w.forEach(n, func(w *Walker, n *Node) {
			fmt.Println(n.XMLName.Local+":"+n.Name, "|")
			w.traverse(n)
		})
	case "interleave":
		w.forEach(n, func(w *Walker, n *Node) {
			fmt.Println(n.XMLName.Local+":"+n.Name, "*")
			w.traverse(n)
		})
	case "optional":
		w.forEach(n, func(w *Walker, n *Node) {
			fmt.Println(n.XMLName.Local+":"+n.Name, "*")
			w.traverse(n)
		})
	case "zeroOrMore":
		w.forEach(n, func(w *Walker, n *Node) {
			fmt.Println(n.XMLName.Local+":"+n.Name, "?")
			w.traverse(n)
		})
	case "oneOrMore":
		w.forEach(n, func(w *Walker, n *Node) {
			fmt.Println(n.XMLName.Local+":"+n.Name, "??")
			w.traverse(n)
		})
	case "group":
		w.forEach(n, func(w *Walker, n *Node) {
			fmt.Println(n.XMLName.Local+":"+n.Name, ".")
			w.traverse(n)
		})
	case "list":
		w.forEach(n, func(w *Walker, n *Node) {
			fmt.Println(n.XMLName.Local+":"+n.Name, ",")
			w.traverse(n)
		})
	case "mixed":
		w.forEach(n, func(w *Walker, n *Node) {
			fmt.Println(n.XMLName.Local+":"+n.Name, " ")
			w.traverse(n)
		})

	//content elements
	case "element":
		fmt.Println("BEGIN", n.Name)
		w.forEach(n, traverseWrap())
	case "attribute":
		fmt.Print("attr ", n.Name)
		w.forEach(n, traverseWrap())
	case "data":
		fmt.Println("data type ", n.Type)
	case "text":
		fmt.Println("text")
	case "value":
		fmt.Println("value ", n.Data())
	case "name":
		fmt.Println("name", n.Data())
	//constraints elements
	case "anyName":
		fmt.Println("anyName")
		w.forEach(n, traverseWrap())
	case "except":
		fmt.Println("except")
		w.forEach(n, traverseWrap())
	case "nsName":
		fmt.Println("nsName", n.NS)
		w.forEach(n, traverseWrap())
	case "empty":
		fmt.Println("empty")
	case "description": //dc:descriprion do nothing
	default:
		halt.As(100, n.XMLName)
	}
}

func Load(input io.Reader) {
	//only xml rng format for now
	xrd := &XMLReader{rd: input}
	if data, err := xrd.Read(); err == nil {
		w := &Walker{root: data}
		w.Init()
		w.traverse(data)
	}
}
