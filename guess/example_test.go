package guess_test

import (
	"fmt"
	"strings"

	"github.com/goark/gnkf/guess"
)

func ExampleEncoding() {
	elist, err := guess.Encoding(strings.NewReader("こんにちは，世界！\n私の名前は Spiegel です。"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Join(elist, ","))
	//Output:
	//UTF-8,windows-1252,windows-1253,Shift_JIS,windows-1255
}

func ExampleEncodingBytes() {
	elist, err := guess.EncodingBytes([]byte("こんにちは，世界！\n私の名前は Spiegel です。"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Join(elist, ","))
	//Output:
	//UTF-8,windows-1252,windows-1253,Shift_JIS,windows-1255
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
