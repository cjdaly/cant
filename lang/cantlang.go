package lang

import (
	"encoding/xml"
	"log"
	"os"
)

func Cant(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Println("err: ", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	dom(decoder)
}

// https://stackoverflow.com/questions/30256729/how-to-traverse-through-xml-data-in-golang

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

func walk(nodes []Node, f func(Node) bool) {
	for _, n := range nodes {
		if f(n) {
			walk(n.Nodes, f)
		}
	}
}

func dom(decoder *xml.Decoder) {
	var n Node
	err := decoder.Decode(&n)
	if err != nil {
		panic(err)
	}

	walk([]Node{n}, func(n Node) bool {
		log.Println(n.XMLName)
		for _, a := range n.Attrs {
			log.Println(a.Name, " = ", a.Value)
		}
		return true
	})
}
