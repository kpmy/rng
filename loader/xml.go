package loader

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type XMLReader struct {
	rd io.Reader
}

type Node struct {
	XMLName  xml.Name
	Name     string `xml:"name,attr"`
	NS       string `xml:"ns,attr"`
	Type     string `xml:"type,attr"`
	Combine  string `xml:"combine,attr"`
	DataType string `xml:"datatypeLibrary,attr"`
	CharData string `xml:",chardata"`
	Href     string `xml:"href,attr"`
	Inner    []Node `xml:",any"`
}

func (p *Node) Data() string {
	return strings.Trim(p.CharData, " \n\r\t")
}

func (p *Node) String() (ret string) {
	ret = fmt.Sprint("$", p.XMLName.Local, " ")
	if p.Name != "" {
		ret = fmt.Sprint(ret, "name: ", p.Name, " ")
	}
	if p.NS != "" {
		ret = fmt.Sprint(ret, "ns: ", p.NS, " ")
	}
	if p.Type != "" {
		ret = fmt.Sprint(ret, "type: ", p.Type, " ")
	}
	if p.Data() != "" {
		ret = fmt.Sprint(ret, "data: ", p.Data(), " ")
	}
	if p.Href != "" {
		ret = fmt.Sprint(ret, "href: ", p.Href, " ")
	}
	if len(p.Inner) > 0 {
		ret = fmt.Sprintln(ret)
	}
	for _, v := range p.Inner {
		ret = fmt.Sprint(ret, "\n", &v)
	}
	return
}

func (n *Node) FindByXMLName(name string) (ret *Node) {
	for i := 0; i < len(n.Inner) && ret == nil; i++ {
		if n.Inner[i].XMLName.Local == name {
			ret = &n.Inner[i]
		}
	}
	return
}

func (n *Node) FindByName(name string) (ret *Node) {
	for i := 0; i < len(n.Inner) && ret == nil; i++ {
		if n.Inner[i].Name == name {
			ret = &n.Inner[i]
		}
	}
	return
}

func (x *XMLReader) Read() (root *Node, err error) {
	root = &Node{}
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, x.rd); err == nil {
		err = xml.NewDecoder(buf).Decode(root)
		//fmt.Println(root)
	}
	return
}
