package width_test

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/gnkf/dump"
	"github.com/spiegel-im-spiegel/gnkf/width"
)

func ExampleTranslateString() {
	txt := "12345 コンバンハ、セカイ ６７８９０ ｺﾝﾊﾞﾝﾊ､ﾆｯﾎﾟﾝ"
	str, err := width.ConvertString("narrow", txt)
	if err != nil {
		return
	}
	dump.UnicodePoint(os.Stdout, bytes.NewBufferString(str))
	fmt.Println()
	str, err = width.ConvertString("widen", txt)
	if err != nil {
		return
	}
	dump.UnicodePoint(os.Stdout, bytes.NewBufferString(str))
	fmt.Println()
	str, err = width.ConvertString("fold", txt)
	if err != nil {
		return
	}
	dump.UnicodePoint(os.Stdout, bytes.NewBufferString(str))
	fmt.Println()
	//Output:
	//0x0031, 0x0032, 0x0033, 0x0034, 0x0035, 0x0020, 0xff7a, 0xff9d, 0xff8a, 0xff9e, 0xff9d, 0xff8a, 0xff64, 0xff7e, 0xff76, 0xff72, 0x0020, 0x0036, 0x0037, 0x0038, 0x0039, 0x0030, 0x0020, 0xff7a, 0xff9d, 0xff8a, 0xff9e, 0xff9d, 0xff8a, 0xff64, 0xff86, 0xff6f, 0xff8e, 0xff9f, 0xff9d
	//0xff11, 0xff12, 0xff13, 0xff14, 0xff15, 0x3000, 0x30b3, 0x30f3, 0x30d0, 0x30f3, 0x30cf, 0x3001, 0x30bb, 0x30ab, 0x30a4, 0x3000, 0xff16, 0xff17, 0xff18, 0xff19, 0xff10, 0x3000, 0x30b3, 0x30f3, 0x30d0, 0x30f3, 0x30cf, 0x3001, 0x30cb, 0x30c3, 0x30dd, 0x30f3
	//0x0031, 0x0032, 0x0033, 0x0034, 0x0035, 0x0020, 0x30b3, 0x30f3, 0x30d0, 0x30f3, 0x30cf, 0x3001, 0x30bb, 0x30ab, 0x30a4, 0x0020, 0x0036, 0x0037, 0x0038, 0x0039, 0x0030, 0x0020, 0x30b3, 0x30f3, 0x30d0, 0x30f3, 0x30cf, 0x3001, 0x30cb, 0x30c3, 0x30dd, 0x30f3
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
