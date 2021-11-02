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
		log.Println(n.XMLName.Local)
		for _, a := range n.Attrs {
			log.Println(a.Name.Local, " = ", a.Value)
		}
		return true
	})
}
