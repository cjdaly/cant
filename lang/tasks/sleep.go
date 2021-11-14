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
	"strconv"
	"time"

	env "cantlang.org/cant/lang/env"
	out "cantlang.org/cant/output"
)

var TaskDefn_Sleep = env.NewTaskDefn("sleep", eval_Sleep)

func eval_Sleep(t *env.TaskInst, c *env.Context) {
	var hours, minutes, seconds, milliseconds int
	for _, attr := range t.Node.Attrs {
		switch attr.Name.Local {
		case "hours":
			hours, _ = strconv.Atoi(attr.Value)
		case "minutes":
			minutes, _ = strconv.Atoi(attr.Value)
		case "seconds":
			seconds, _ = strconv.Atoi(attr.Value)
		case "milliseconds":
			milliseconds, _ = strconv.Atoi(attr.Value)
		}

		if attr.Name.Local == "main" {
		}
	}

	var sleepTime = (hours * 60 * 60 * 1000) +
		(minutes * 60 * 1000) +
		(seconds * 1000) +
		milliseconds

	out.Logln("Sleep for " + strconv.Itoa(sleepTime) + "milliseconds")
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
}
