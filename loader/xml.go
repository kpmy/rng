package loader

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

type XMLReader struct {
	rd io.Reader
}

type Pattern struct {
}

func (p *Pattern) String() string {
	return "pattern"
}

type Grammar struct {
	PatternList []Pattern `xml:"define"`
	Start       []Pattern `xml:"start"`
}

func (g *Grammar) String() (ret string) {
	ret = fmt.Sprintln("grammar defines ", len(g.PatternList), " patterns")
	for _, v := range g.PatternList {
		ret = fmt.Sprintln(ret, &v)
	}
	return
}

func (x *XMLReader) Read() {
	root := &Grammar{}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, x.rd)
	xml.Unmarshal(buf.Bytes(), root)
	fmt.Println(root)
}
