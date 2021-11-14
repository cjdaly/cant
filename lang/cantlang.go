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
	env "cantlang.org/cant/lang/env"
	tasks "cantlang.org/cant/lang/tasks"
	xml "cantlang.org/cant/lang/xml"
)

func Cant(inputPath string) {
	var cantNode = xml.Load(inputPath)
	var cantTask = env.NewTask(&cantNode)
	var context = env.NewRootContext(newTaskRegistry())
	env.Eval(&cantTask, &context)
}

func newTaskRegistry() map[string]*env.TaskDefn {
	var reg = make(map[string]*env.TaskDefn)

	reg["cant"] = &tasks.TaskDefn_Cant
	reg["echo"] = &tasks.TaskDefn_Echo
	reg["property"] = &tasks.TaskDefn_Property
	reg["sleep"] = &tasks.TaskDefn_Sleep
	reg["target"] = &tasks.TaskDefn_Target

	return reg
}
