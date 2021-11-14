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

package task

import (
	xml "cantlang.org/cant/lang/xml"
)

type Task interface {
	Name() string
	Attr(name string) string
}

type TaskInst struct {
	node *xml.Node
}

func NewTask(node *xml.Node) TaskInst {
	return TaskInst{node}
}

func (t TaskInst) Name() string {
	return t.node.XMLName.Local
}

func (t TaskInst) Attr(name string) string {
	for _, attr := range t.node.Attrs {
		if attr.Name.Local == name {
			return attr.Value
		}
	}
	return ""
}

///

type TaskDefn struct {
	Name string
	Eval func(*TaskInst, *Context)
}

///

type Context struct {
	Parent *Context
	Locals map[string]string
}
