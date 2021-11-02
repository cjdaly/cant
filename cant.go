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

package main

import (
	"errors"
	"log"
	"os"

	lang "cantlang.org/cant/lang"
)

func main() {
	log.SetPrefix("CANT: ")

	if len(os.Args) < 2 {
		log.Fatal(errors.New("Missing input file argument!"))
	}

	arg1 := os.Args[1]
	log.Println("input:", arg1)

	lang.Cant(arg1)
}
