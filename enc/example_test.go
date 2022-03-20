package enc_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/goark/gnkf/dump"
	"github.com/goark/gnkf/enc"
)

func ExampleConvert() {
	buf := &bytes.Buffer{}
	if err := enc.Convert("Shift_JIS", buf, "UTF-8", strings.NewReader("こんにちは，世界！\n私の名前は Spiegel です。")); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := dump.Octet(os.Stdout, buf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//Output:
	//0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x81, 0x43, 0x90, 0xa2, 0x8a, 0x45, 0x81, 0x49, 0x0a, 0x8e, 0x84, 0x82, 0xcc, 0x96, 0xbc, 0x91, 0x4f, 0x82, 0xcd, 0x20, 0x53, 0x70, 0x69, 0x65, 0x67, 0x65, 0x6c, 0x20, 0x82, 0xc5, 0x82, 0xb7, 0x81, 0x42
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
