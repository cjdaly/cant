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

package tasks

import (
	env "cantlang.org/cant/lang/env"
	out "cantlang.org/cant/output"
)

var TaskDefn_Cant = env.NewTaskDefn("cant", eval_Cant)

func eval_Cant(t *env.TaskInst, c *env.Context) {
	mainTarget := "main"
	for _, attr := range t.Node.Attrs {
		if attr.Name.Local == "main" {
			mainTarget = attr.Value
		}
	}
	out.Logln("main target: " + mainTarget)

	for _, node := range t.Node.Nodes {
		out.Logln("  stuff : " + node.XMLName.Local)
	}
}
