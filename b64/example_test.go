package b64_test

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/goark/gnkf/b64"
)

func ExampleEncode() {
	input := strings.NewReader("Hello World\n")
	output := &bytes.Buffer{}
	if err := b64.Encode(false, false, input, output); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output.String())
	// Output:
	// SGVsbG8gV29ybGQK
}

func ExampleDecode() {
	input := strings.NewReader("SGVsbG8gV29ybGQK")
	output := &bytes.Buffer{}
	if err := b64.Decode(false, false, input, output); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(output.String())
	// Output:
	// Hello World
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
