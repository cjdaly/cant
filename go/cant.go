package main

import (
  "encoding/xml"
  "fmt"
  "io"
  "os"
)

func events(decoder *xml.Decoder) {
  for {
    tok, err := decoder.Token()
    if tok == nil || err == io.EOF {
      break
    }

    switch tokType := tok.(type) {
    case xml.StartElement:
      fmt.Println("Start: ", tokType.Name)
    case xml.EndElement:
      fmt.Println("End: ", tokType.Name)
    case xml.Attr:
      fmt.Println("Attr: ", tokType.Name)
    }
  }
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
    fmt.Println(n.XMLName)
    for _, a := range(n.Attrs) {
      fmt.Println(a.Name, " = ", a.Value)
    }
    return true
  })
}

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    fmt.Println("err: ", err)
  }
  defer file.Close()

  decoder := xml.NewDecoder(file)
  // events(decoder)
  dom(decoder)
}

