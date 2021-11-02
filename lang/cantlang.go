/* Copyright (c) 2021 Chris J Daly (github user cjdaly)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lang

import (
	"encoding/xml"
	"errors"
	"log"
	"os"
	"strings"
)

func Cant(inputPath string) {
	file, f_err := os.Open(inputPath)
	defer file.Close()
	if f_err != nil {
		log.Fatal(f_err)
	}

	var cant Node
	decoder := xml.NewDecoder(file)
	d_err := decoder.Decode(&cant)
	if d_err != nil {
		log.Fatal(d_err)
	}

	if strings.ToLower(cant.XMLName.Local) != "cant" {
		log.Fatal(errors.New("Missing <cant> top-level element!"))
	}

	mainTarget := "main"
	for _, attr := range cant.Attrs {
		if attr.Name.Local == "main" {
			mainTarget = attr.Value
		}
	}
	log.Println("main target: " + mainTarget)

	// var g G
	for _, node := range cant.Nodes {
		if strings.ToLower(node.XMLName.Local) == "target" {

		} else if strings.ToLower(node.XMLName.Local) == "property" {

		}
	}

	// dom(decoder)
}

type G struct {
	Targets    map[string]Node
	Properties map[string]Node
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
		log.Println(n.XMLName.Local)
		for _, a := range n.Attrs {
			log.Println(a.Name.Local, " = ", a.Value)
		}
		return true
	})
}
