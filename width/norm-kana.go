package width

import "strings"

var normKatakanaMap = map[string]string{
	string([]rune{'ガ'}): string([]rune{'カ', 0x3099}),
	string([]rune{'ギ'}): string([]rune{'キ', 0x3099}),
	string([]rune{'グ'}): string([]rune{'ク', 0x3099}),
	string([]rune{'ゲ'}): string([]rune{'ケ', 0x3099}),
	string([]rune{'ゴ'}): string([]rune{'コ', 0x3099}),
	string([]rune{'ザ'}): string([]rune{'サ', 0x3099}),
	string([]rune{'ジ'}): string([]rune{'シ', 0x3099}),
	string([]rune{'ズ'}): string([]rune{'ス', 0x3099}),
	string([]rune{'ゼ'}): string([]rune{'セ', 0x3099}),
	string([]rune{'ゾ'}): string([]rune{'ソ', 0x3099}),
	string([]rune{'ダ'}): string([]rune{'タ', 0x3099}),
	string([]rune{'ヂ'}): string([]rune{'チ', 0x3099}),
	string([]rune{'ヅ'}): string([]rune{'ツ', 0x3099}),
	string([]rune{'デ'}): string([]rune{'テ', 0x3099}),
	string([]rune{'ド'}): string([]rune{'ト', 0x3099}),
	string([]rune{'バ'}): string([]rune{'ハ', 0x3099}),
	string([]rune{'ビ'}): string([]rune{'ヒ', 0x3099}),
	string([]rune{'ブ'}): string([]rune{'フ', 0x3099}),
	string([]rune{'ベ'}): string([]rune{'ヘ', 0x3099}),
	string([]rune{'ボ'}): string([]rune{'ホ', 0x3099}),
	string([]rune{'パ'}): string([]rune{'ハ', 0x309a}),
	string([]rune{'ピ'}): string([]rune{'ヒ', 0x309a}),
	string([]rune{'プ'}): string([]rune{'フ', 0x309a}),
	string([]rune{'ペ'}): string([]rune{'ヘ', 0x309a}),
	string([]rune{'ポ'}): string([]rune{'ホ', 0x309a}),
	string([]rune{'ヴ'}): string([]rune{'ウ', 0x3099}),
	string([]rune{'ヷ'}): string([]rune{'ワ', 0x3099}),
	string([]rune{'ヸ'}): string([]rune{'ヰ', 0x3099}),
	string([]rune{'ヹ'}): string([]rune{'ヱ', 0x3099}),
	string([]rune{'ヺ'}): string([]rune{'ヲ', 0x3099}),
}

//NewReplaceerkanaNFC returns strings.Replacer instance for NFC translation of katakana
func NewReplaceerkanaNFC() *strings.Replacer {
	ss := []string{}
	for k, v := range normKatakanaMap {
		ss = append(ss, v, k)
	}
	return strings.NewReplacer(ss...)
}

//NewReplaceerkanaNFD returns strings.Replacer instance for NFD translation of katakana
func NewReplaceerkanaNFD() *strings.Replacer {
	ss := []string{}
	for k, v := range normKatakanaMap {
		ss = append(ss, k, v)
	}
	return strings.NewReplacer(ss...)
}

var normKatakanaMap2 = map[string]string{
	string([]rune{'ヮ'}):     string([]rune{'ﾜ'}),
	string([]rune{'ヰ'}):     string([]rune{'ｲ'}),
	string([]rune{'ヱ'}):     string([]rune{'ｴ'}),
	string([]rune{'ヵ'}):     string([]rune{'ｶ'}),
	string([]rune{'ヶ'}):     string([]rune{'ｹ'}),
	string([]rune{'ㇰ'}):     string([]rune{'ｸ'}),
	string([]rune{'ㇱ'}):     string([]rune{'ｼ'}),
	string([]rune{'ㇲ'}):     string([]rune{'ｽ'}),
	string([]rune{'ㇳ'}):     string([]rune{'ﾄ'}),
	string([]rune{'ㇴ'}):     string([]rune{'ﾇ'}),
	string([]rune{'ㇵ'}):     string([]rune{'ﾊ'}),
	string([]rune{'ㇶ'}):     string([]rune{'ﾋ'}),
	string([]rune{'ㇷ'}):     string([]rune{'ﾌ'}),
	string([]rune{'ㇸ'}):     string([]rune{'ﾍ'}),
	string([]rune{'ㇹ'}):     string([]rune{'ﾎ'}),
	string([]rune{'ㇺ'}):     string([]rune{'ﾑ'}),
	string([]rune{'ㇻ'}):     string([]rune{'ﾗ'}),
	string([]rune{'ㇼ'}):     string([]rune{'ﾘ'}),
	string([]rune{'ㇽ'}):     string([]rune{'ﾙ'}),
	string([]rune{'ㇾ'}):     string([]rune{'ﾚ'}),
	string([]rune{'ㇿ'}):     string([]rune{'ﾛ'}),
	string([]rune{0x1b164}): string([]rune{'ｲ'}),
	string([]rune{0x1b165}): string([]rune{'ｴ'}),
	string([]rune{0x1b166}): string([]rune{'ｦ'}),
	string([]rune{0x1b167}): string([]rune{'ﾝ'}),
}

//NewReplaceerHalfWidthkana returns strings.Replacer instance for complement half-width katakana
func NewReplaceerHalfWidthkana() *strings.Replacer {
	ss := []string{}
	for k, v := range normKatakanaMap2 {
		ss = append(ss, k, v)
	}
	return strings.NewReplacer(ss...)
}

/* Copyright 2020-2021 Spiegel
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
