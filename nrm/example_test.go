package nrm_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/goark/gnkf/dump"
	"github.com/goark/gnkf/nrm"
)

func ExampleNormalize() {
	buf := &bytes.Buffer{}
	if err := nrm.Normalize("nfkc", buf, strings.NewReader("ﾍﾟﾝｷﾞﾝ"), false); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if err := dump.UnicodePoint(os.Stdout, buf); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	//Output:
	//0x30da, 0x30f3, 0x30ae, 0x30f3
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
