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
	task "cantlang.org/cant/lang/task"
	xml "cantlang.org/cant/lang/xml"
)

func Cant(inputPath string) {

	var reg = make(map[string]*task.TaskDefn)
	reg["cant"] = &task.TaskDefn_Cant
	reg["echo"] = &task.TaskDefn_Echo
	reg["property"] = &task.TaskDefn_Property
	reg["sleep"] = &task.TaskDefn_Sleep
	reg["target"] = &task.TaskDefn_Target

	var cant = xml.Load(inputPath)
	var cantTask = task.NewTask(&cant)

	var c = task.Context{}
	var eval = reg[cantTask.Name()].Eval
	eval(&cantTask, &c)

}
