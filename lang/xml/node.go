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

package xml

import (
	"encoding/xml"
	"errors"
	"os"

	out "cantlang.org/cant/output"
)

// https://stackoverflow.com/questions/30256729/how-to-traverse-through-xml-data-in-golang

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

func Load(inputPath string) Node {
	file, f_err := os.Open(inputPath)
	defer file.Close()
	if f_err != nil {
		out.LogFatal(f_err)
	}

	var cant Node
	decoder := xml.NewDecoder(file)
	d_err := decoder.Decode(&cant)
	if d_err != nil {
		out.LogFatal(d_err)
	}

	if cant.XMLName.Local != "cant" {
		out.LogFatal(errors.New("Missing <cant> top-level element!"))
	}

	return cant
}
