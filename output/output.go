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

package output

import (
	"fmt"
	"log"
	"strings"
)

func init() {
	log.SetPrefix("CANT: ")
}

func Logln(message string) {
	log.Println(message)
}

func LogFatal(err error) {
	log.Fatal(err)
}

func Println(message string) {
	if !testing {
		fmt.Println(message)
	} else {
		testOutput.WriteString(message + "\n")
	}
}

var testing = false
var testOutput strings.Builder

func Testing() {
	testing = true
	testOutput.Reset()
}

func TestOutput() string {
	return string(testOutput.String())
}
