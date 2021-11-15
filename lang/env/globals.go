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

///

type Globals struct {
	TaskRegistry map[string]*TaskDefn
	Properties   map[string]*TaskInst
	Targets      map[string]*TaskInst
}

func NewGlobals(taskRegistry map[string]*TaskDefn) *Globals {
	var g = Globals{
		taskRegistry,
		make(map[string]*TaskInst),
		make(map[string]*TaskInst)}
	return &g
}

func (g *Globals) NewContext() *Context {
	var c = Context{
		g,
		nil,
		make(map[string]*TaskInst)}
	return &c
}

func (g *Globals) TaskDefn(name string) *TaskDefn {
	if g.TaskRegistry != nil {
		td := g.TaskRegistry[name]
		if td != nil {
			return td
		}
	}
	return nil
}
