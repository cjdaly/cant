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
	"errors"

	env "cantlang.org/cant/lang/env"
	out "cantlang.org/cant/output"
)

var TaskDefn_Cant = env.NewTaskDefn("cant", eval_Cant)

func eval_Cant(t *env.TaskInst, c *env.Context) {
	mainTargetName := "main"
	for _, attr := range t.Node.Attrs {
		if attr.Name.Local == "main" {
			mainTargetName = attr.Value
		}
	}
	out.Logln("main target: " + mainTargetName)

	for _, node := range t.Node.Nodes {
		switch node.XMLName.Local {
		case "property":
			c.AddProperty(env.NewTask(node))
		case "target":
			c.AddTarget(env.NewTask(node))
		}
	}

	mainTarget := c.GetTarget(mainTargetName)
	if mainTarget == nil {
		out.LogFatal(errors.New("main target not found: " + mainTargetName))
	} else {
		env.Eval(mainTarget, c)
	}
}
