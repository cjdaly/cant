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
	"io/ioutil"
	"testing"

	out "cantlang.org/cant/output"
)

var tests = []string{
	"test_hello",
	"test_nondefault_main",
}

func TestBasic(t *testing.T) {
	for _, tn := range tests {
		out.Testing()

		testPath := "testdata/basic/" + tn + ".xml"
		goldenPath := "testdata/basic/" + tn + ".golden"

		out.Logln("testPath: " + testPath)

		Cant(testPath)

		testOut := out.TestOutput()
		out.Logln("testOut: " + testOut)

		testGoldenBytes, err := ioutil.ReadFile(goldenPath)
		if err != nil {
			t.Fatalf("failed reading golden file: %s", err)
		}

		testGolden := string(testGoldenBytes)
		out.Logln("testGolden: " + testGolden)

		if testOut != testGolden {
			t.Errorf("test: " + tn + ", output does not match golden file")
		}

	}
}
