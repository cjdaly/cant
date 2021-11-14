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

func Eval(t *TaskInst, c *Context) {
	td := c.TaskDefn(t.Name())
	td.Eval(t, c)
}

///

type Context struct {
	TaskRegistry map[string]*TaskDefn
	Parent       *Context
	Locals       map[string]*TaskInst
}

func NewRootContext(taskRegistry map[string]*TaskDefn) Context {
	var rc = Context{
		taskRegistry,
		nil,
		make(map[string]*TaskInst)}
	return rc
}

func (c Context) NewSubContext() Context {
	var sc = Context{
		nil,
		&c,
		make(map[string]*TaskInst)}
	return sc
}

func (c Context) IsRoot() bool {
	return c.Parent == nil
}

func (c Context) TaskDefn(name string) *TaskDefn {
	if c.TaskRegistry != nil {
		td := c.TaskRegistry[name]
		if td != nil {
			return td
		}
	}
	if c.IsRoot() {
		return nil
	} else {
		return c.Parent.TaskDefn(name)
	}
}
