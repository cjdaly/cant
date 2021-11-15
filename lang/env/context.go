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

package env

import (
	"errors"

	out "cantlang.org/cant/output"
)

type Context struct {
	G      *Globals
	Parent *Context
	Locals map[string]*TaskInst
}

func (c *Context) NewContext() *Context {
	var sc = Context{
		c.G,
		c,
		make(map[string]*TaskInst)}
	return &sc
}

func (c *Context) IsRoot() bool {
	return c.Parent == nil
}

func (c *Context) AddProperty(t *TaskInst) {
	if t.Name() != "property" {
		out.LogFatal(errors.New("AddProperty with <" + t.Name() + ">"))
	}

	n := t.Attr("name")
	if n == "" {
		out.Logln("Ignoring nameless property...")
	} else if _, found := c.G.Properties[n]; found {
		out.Logln("Ignoring duplicate property: " + n)
	} else {
		c.G.Properties[n] = t
		out.Logln("Added property: " + n)
	}
}

func (c *Context) AddTarget(t *TaskInst) {
	if t.Name() != "target" {
		out.LogFatal(errors.New("AddTarget with <" + t.Name() + ">"))
	}

	n := t.Attr("name")
	if n == "" {
		out.Logln("Ignoring nameless target...")
	} else if _, found := c.G.Targets[n]; found {
		out.Logln("Ignoring duplicate target: " + n)
	} else {
		c.G.Targets[n] = t
		out.Logln("Added target: " + n)
	}
}

func (c *Context) GetTarget(name string) *TaskInst {
	return c.G.Targets[name]
}
