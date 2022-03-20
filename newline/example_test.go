package newline_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/goark/gnkf/dump"
	"github.com/goark/gnkf/newline"
)

var text = `こんにちは
世界！`

func ExampleConvert() {
	buf := &bytes.Buffer{}
	if err := newline.Convert("crlf", buf, strings.NewReader(text)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := dump.UnicodePoint(os.Stdout, buf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//Output:
	//0x3053, 0x3093, 0x306b, 0x3061, 0x306f, 0x000d, 0x000a, 0x4e16, 0x754c, 0xff01
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
