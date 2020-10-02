package kana

import (
	"sort"
	"strings"
	"unicode"
)

var (
	kanaCase = unicode.SpecialCase{
		unicode.CaseRange{Lo: 'ぁ', Hi: 'ゖ', Delta: [unicode.MaxCase]rune{'ァ' - 'ぁ', 0, 0}},
		unicode.CaseRange{Lo: 'ゝ', Hi: 'ゞ', Delta: [unicode.MaxCase]rune{'ヽ' - 'ゝ', 0, 0}},
		unicode.CaseRange{Lo: 'ァ', Hi: 'ヶ', Delta: [unicode.MaxCase]rune{0, 'ぁ' - 'ァ', 0}},
		unicode.CaseRange{Lo: 'ヽ', Hi: 'ヾ', Delta: [unicode.MaxCase]rune{0, 'ゝ' - 'ヽ', 0}},
	}
	replacekanaMap = map[string]string{
		string([]rune{'ヷ'}): string([]rune{'わ', 0x3099}),
		string([]rune{'ヸ'}): string([]rune{'ゐ', 0x3099}),
		string([]rune{'ヹ'}): string([]rune{'ゑ', 0x3099}),
		string([]rune{'ヺ'}): string([]rune{'を', 0x3099}),
	}
	kanaCase2 = unicode.SpecialCase{
		unicode.CaseRange{Lo: 'あ', Hi: 'あ', Delta: [unicode.MaxCase]rune{0, 'ぁ' - 'あ', 0}},
		unicode.CaseRange{Lo: 'い', Hi: 'い', Delta: [unicode.MaxCase]rune{0, 'ぃ' - 'い', 0}},
		unicode.CaseRange{Lo: 'う', Hi: 'う', Delta: [unicode.MaxCase]rune{0, 'ぅ' - 'う', 0}},
		unicode.CaseRange{Lo: 'え', Hi: 'え', Delta: [unicode.MaxCase]rune{0, 'ぇ' - 'え', 0}},
		unicode.CaseRange{Lo: 'お', Hi: 'お', Delta: [unicode.MaxCase]rune{0, 'ぉ' - 'お', 0}},
		unicode.CaseRange{Lo: 'か', Hi: 'か', Delta: [unicode.MaxCase]rune{0, 'ゕ' - 'か', 0}},
		unicode.CaseRange{Lo: 'け', Hi: 'け', Delta: [unicode.MaxCase]rune{0, 'ゖ' - 'け', 0}},
		unicode.CaseRange{Lo: 'つ', Hi: 'つ', Delta: [unicode.MaxCase]rune{0, 'っ' - 'つ', 0}},
		unicode.CaseRange{Lo: 'や', Hi: 'や', Delta: [unicode.MaxCase]rune{0, 'ゃ' - 'や', 0}},
		unicode.CaseRange{Lo: 'ゆ', Hi: 'ゆ', Delta: [unicode.MaxCase]rune{0, 'ゅ' - 'ゆ', 0}},
		unicode.CaseRange{Lo: 'よ', Hi: 'よ', Delta: [unicode.MaxCase]rune{0, 'ょ' - 'よ', 0}},
		unicode.CaseRange{Lo: 'わ', Hi: 'わ', Delta: [unicode.MaxCase]rune{0, 'ゎ' - 'わ', 0}},

		unicode.CaseRange{Lo: 'ぁ', Hi: 'ぁ', Delta: [unicode.MaxCase]rune{'あ' - 'ぁ', 0, 0}},
		unicode.CaseRange{Lo: 'ぃ', Hi: 'ぃ', Delta: [unicode.MaxCase]rune{'い' - 'ぃ', 0, 0}},
		unicode.CaseRange{Lo: 'ぅ', Hi: 'ぅ', Delta: [unicode.MaxCase]rune{'う' - 'ぅ', 0, 0}},
		unicode.CaseRange{Lo: 'ぇ', Hi: 'ぇ', Delta: [unicode.MaxCase]rune{'え' - 'ぇ', 0, 0}},
		unicode.CaseRange{Lo: 'ぉ', Hi: 'ぉ', Delta: [unicode.MaxCase]rune{'お' - 'ぉ', 0, 0}},
		unicode.CaseRange{Lo: 'ゕ', Hi: 'ゕ', Delta: [unicode.MaxCase]rune{'か' - 'ゕ', 0, 0}},
		unicode.CaseRange{Lo: 'ゖ', Hi: 'ゖ', Delta: [unicode.MaxCase]rune{'け' - 'ゖ', 0, 0}},
		unicode.CaseRange{Lo: 'っ', Hi: 'っ', Delta: [unicode.MaxCase]rune{'つ' - 'っ', 0, 0}},
		unicode.CaseRange{Lo: 'ゃ', Hi: 'ゃ', Delta: [unicode.MaxCase]rune{'や' - 'ゃ', 0, 0}},
		unicode.CaseRange{Lo: 'ゅ', Hi: 'ゅ', Delta: [unicode.MaxCase]rune{'ゆ' - 'ゅ', 0, 0}},
		unicode.CaseRange{Lo: 'ょ', Hi: 'ょ', Delta: [unicode.MaxCase]rune{'よ' - 'ょ', 0, 0}},
		unicode.CaseRange{Lo: 'ゎ', Hi: 'ゎ', Delta: [unicode.MaxCase]rune{'わ' - 'ゎ', 0, 0}},

		unicode.CaseRange{Lo: 'ア', Hi: 'ア', Delta: [unicode.MaxCase]rune{0, 'ァ' - 'ア', 0}},
		unicode.CaseRange{Lo: 'イ', Hi: 'イ', Delta: [unicode.MaxCase]rune{0, 'ィ' - 'イ', 0}},
		unicode.CaseRange{Lo: 'ウ', Hi: 'ウ', Delta: [unicode.MaxCase]rune{0, 'ゥ' - 'ウ', 0}},
		unicode.CaseRange{Lo: 'エ', Hi: 'エ', Delta: [unicode.MaxCase]rune{0, 'ェ' - 'エ', 0}},
		unicode.CaseRange{Lo: 'オ', Hi: 'オ', Delta: [unicode.MaxCase]rune{0, 'ォ' - 'オ', 0}},
		unicode.CaseRange{Lo: 'カ', Hi: 'カ', Delta: [unicode.MaxCase]rune{0, 'ヵ' - 'カ', 0}},
		unicode.CaseRange{Lo: 'ケ', Hi: 'ケ', Delta: [unicode.MaxCase]rune{0, 'ヶ' - 'ケ', 0}},
		unicode.CaseRange{Lo: 'ツ', Hi: 'ツ', Delta: [unicode.MaxCase]rune{0, 'ッ' - 'ツ', 0}},
		unicode.CaseRange{Lo: 'ヤ', Hi: 'ヤ', Delta: [unicode.MaxCase]rune{0, 'ャ' - 'ヤ', 0}},
		unicode.CaseRange{Lo: 'ユ', Hi: 'ユ', Delta: [unicode.MaxCase]rune{0, 'ュ' - 'ユ', 0}},
		unicode.CaseRange{Lo: 'ヨ', Hi: 'ヨ', Delta: [unicode.MaxCase]rune{0, 'ョ' - 'ヨ', 0}},
		unicode.CaseRange{Lo: 'ワ', Hi: 'ワ', Delta: [unicode.MaxCase]rune{0, 'ヮ' - 'ワ', 0}},

		unicode.CaseRange{Lo: 'ァ', Hi: 'ァ', Delta: [unicode.MaxCase]rune{'ア' - 'ァ', 0, 0}},
		unicode.CaseRange{Lo: 'ィ', Hi: 'ィ', Delta: [unicode.MaxCase]rune{'イ' - 'ィ', 0, 0}},
		unicode.CaseRange{Lo: 'ゥ', Hi: 'ゥ', Delta: [unicode.MaxCase]rune{'ウ' - 'ゥ', 0, 0}},
		unicode.CaseRange{Lo: 'ェ', Hi: 'ェ', Delta: [unicode.MaxCase]rune{'エ' - 'ェ', 0, 0}},
		unicode.CaseRange{Lo: 'ォ', Hi: 'ォ', Delta: [unicode.MaxCase]rune{'オ' - 'ォ', 0, 0}},
		unicode.CaseRange{Lo: 'ヵ', Hi: 'ヵ', Delta: [unicode.MaxCase]rune{'カ' - 'ヵ', 0, 0}},
		unicode.CaseRange{Lo: 'ヶ', Hi: 'ヶ', Delta: [unicode.MaxCase]rune{'ケ' - 'ヶ', 0, 0}},
		unicode.CaseRange{Lo: 'ッ', Hi: 'ッ', Delta: [unicode.MaxCase]rune{'ツ' - 'ッ', 0, 0}},
		unicode.CaseRange{Lo: 'ャ', Hi: 'ャ', Delta: [unicode.MaxCase]rune{'ヤ' - 'ャ', 0, 0}},
		unicode.CaseRange{Lo: 'ュ', Hi: 'ュ', Delta: [unicode.MaxCase]rune{'ユ' - 'ュ', 0, 0}},
		unicode.CaseRange{Lo: 'ョ', Hi: 'ョ', Delta: [unicode.MaxCase]rune{'ヨ' - 'ョ', 0, 0}},
		unicode.CaseRange{Lo: 'ヮ', Hi: 'ヮ', Delta: [unicode.MaxCase]rune{'ワ' - 'ヮ', 0, 0}},

		unicode.CaseRange{Lo: 'ｱ', Hi: 'ｱ', Delta: [unicode.MaxCase]rune{0, 'ｧ' - 'ｱ', 0}},
		unicode.CaseRange{Lo: 'ｲ', Hi: 'ｲ', Delta: [unicode.MaxCase]rune{0, 'ｨ' - 'ｲ', 0}},
		unicode.CaseRange{Lo: 'ｳ', Hi: 'ｳ', Delta: [unicode.MaxCase]rune{0, 'ｩ' - 'ｳ', 0}},
		unicode.CaseRange{Lo: 'ｴ', Hi: 'ｴ', Delta: [unicode.MaxCase]rune{0, 'ｪ' - 'ｴ', 0}},
		unicode.CaseRange{Lo: 'ｵ', Hi: 'ｵ', Delta: [unicode.MaxCase]rune{0, 'ｫ' - 'ｵ', 0}},
		unicode.CaseRange{Lo: 'ﾂ', Hi: 'ﾂ', Delta: [unicode.MaxCase]rune{0, 'ｯ' - 'ﾂ', 0}},
		unicode.CaseRange{Lo: 'ﾔ', Hi: 'ﾔ', Delta: [unicode.MaxCase]rune{0, 'ｬ' - 'ﾔ', 0}},
		unicode.CaseRange{Lo: 'ﾕ', Hi: 'ﾕ', Delta: [unicode.MaxCase]rune{0, 'ｭ' - 'ﾕ', 0}},
		unicode.CaseRange{Lo: 'ﾖ', Hi: 'ﾖ', Delta: [unicode.MaxCase]rune{0, 'ｮ' - 'ﾖ', 0}},

		unicode.CaseRange{Lo: 'ｧ', Hi: 'ｧ', Delta: [unicode.MaxCase]rune{'ｱ' - 'ｧ', 0, 0}},
		unicode.CaseRange{Lo: 'ｨ', Hi: 'ｨ', Delta: [unicode.MaxCase]rune{'ｲ' - 'ｨ', 0, 0}},
		unicode.CaseRange{Lo: 'ｩ', Hi: 'ｩ', Delta: [unicode.MaxCase]rune{'ｳ' - 'ｩ', 0, 0}},
		unicode.CaseRange{Lo: 'ｪ', Hi: 'ｪ', Delta: [unicode.MaxCase]rune{'ｴ' - 'ｪ', 0, 0}},
		unicode.CaseRange{Lo: 'ｫ', Hi: 'ｫ', Delta: [unicode.MaxCase]rune{'ｵ' - 'ｫ', 0, 0}},
		unicode.CaseRange{Lo: 'ｯ', Hi: 'ｯ', Delta: [unicode.MaxCase]rune{'ﾂ' - 'ｯ', 0, 0}},
		unicode.CaseRange{Lo: 'ｬ', Hi: 'ｬ', Delta: [unicode.MaxCase]rune{'ﾔ' - 'ｬ', 0, 0}},
		unicode.CaseRange{Lo: 'ｭ', Hi: 'ｭ', Delta: [unicode.MaxCase]rune{'ﾕ' - 'ｭ', 0, 0}},
		unicode.CaseRange{Lo: 'ｮ', Hi: 'ｮ', Delta: [unicode.MaxCase]rune{'ﾖ' - 'ｮ', 0, 0}}}
)

func init() {
	sort.Slice(kanaCase2, func(i, j int) bool {
		return kanaCase2[i].Lo < kanaCase2[j].Lo
	})
}

//ReplaceHiragana replaces hiragana from katrakana (full-width kana kcharacter only).
func ReplaceHiragana(txt string) string {
	ss := []string{}
	for k, v := range replacekanaMap {
		ss = append(ss, k, v)
	}
	return strings.ToLowerSpecial(kanaCase, strings.NewReplacer(ss...).Replace(txt))
}

//ReplaceKatakana replaces katakana from hiragana (full-width kana kcharacter only).
func ReplaceKatakana(txt string) string {
	ss := []string{}
	for k, v := range replacekanaMap {
		ss = append(ss, v, k)
	}
	return strings.ToUpperSpecial(kanaCase, strings.NewReplacer(ss...).Replace(txt))
}

//ReplaceChokuon replaces chokuon (upper kana case).
func ReplaceChokuon(txt string) string {
	return strings.ToUpperSpecial(kanaCase2, txt)
}

//ReplaceYouon replaces youon (lower kana case).
func ReplaceYouon(txt string) string {
	return strings.ToLowerSpecial(kanaCase2, txt)
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
