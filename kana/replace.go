package kana

import (
	"sort"
	"strings"
	"unicode"
)

var (
	kanaCase = unicode.SpecialCase{
		unicode.CaseRange{'ぁ', 'ゖ', [unicode.MaxCase]rune{'ァ' - 'ぁ', 0, 0}},
		unicode.CaseRange{'ゝ', 'ゞ', [unicode.MaxCase]rune{'ヽ' - 'ゝ', 0, 0}},
		unicode.CaseRange{'ァ', 'ヶ', [unicode.MaxCase]rune{0, 'ぁ' - 'ァ', 0}},
		unicode.CaseRange{'ヽ', 'ヾ', [unicode.MaxCase]rune{0, 'ゝ' - 'ヽ', 0}},
	}
	replacekanaMap = map[string]string{
		string([]rune{'ヷ'}): string([]rune{'わ', 0x3099}),
		string([]rune{'ヸ'}): string([]rune{'ゐ', 0x3099}),
		string([]rune{'ヹ'}): string([]rune{'ゑ', 0x3099}),
		string([]rune{'ヺ'}): string([]rune{'を', 0x3099}),
	}
	kanaCase2 = unicode.SpecialCase{
		unicode.CaseRange{'あ', 'あ', [unicode.MaxCase]rune{0, 'ぁ' - 'あ', 0}},
		unicode.CaseRange{'い', 'い', [unicode.MaxCase]rune{0, 'ぃ' - 'い', 0}},
		unicode.CaseRange{'う', 'う', [unicode.MaxCase]rune{0, 'ぅ' - 'う', 0}},
		unicode.CaseRange{'え', 'え', [unicode.MaxCase]rune{0, 'ぇ' - 'え', 0}},
		unicode.CaseRange{'お', 'お', [unicode.MaxCase]rune{0, 'ぉ' - 'お', 0}},
		unicode.CaseRange{'か', 'か', [unicode.MaxCase]rune{0, 'ゕ' - 'か', 0}},
		unicode.CaseRange{'け', 'け', [unicode.MaxCase]rune{0, 'ゖ' - 'け', 0}},
		unicode.CaseRange{'つ', 'つ', [unicode.MaxCase]rune{0, 'っ' - 'つ', 0}},
		unicode.CaseRange{'や', 'や', [unicode.MaxCase]rune{0, 'ゃ' - 'や', 0}},
		unicode.CaseRange{'ゆ', 'ゆ', [unicode.MaxCase]rune{0, 'ゅ' - 'ゆ', 0}},
		unicode.CaseRange{'よ', 'よ', [unicode.MaxCase]rune{0, 'ょ' - 'よ', 0}},
		unicode.CaseRange{'わ', 'わ', [unicode.MaxCase]rune{0, 'ゎ' - 'わ', 0}},

		unicode.CaseRange{'ぁ', 'ぁ', [unicode.MaxCase]rune{'あ' - 'ぁ', 0, 0}},
		unicode.CaseRange{'ぃ', 'ぃ', [unicode.MaxCase]rune{'い' - 'ぃ', 0, 0}},
		unicode.CaseRange{'ぅ', 'ぅ', [unicode.MaxCase]rune{'う' - 'ぅ', 0, 0}},
		unicode.CaseRange{'ぇ', 'ぇ', [unicode.MaxCase]rune{'え' - 'ぇ', 0, 0}},
		unicode.CaseRange{'ぉ', 'ぉ', [unicode.MaxCase]rune{'お' - 'ぉ', 0, 0}},
		unicode.CaseRange{'ゕ', 'ゕ', [unicode.MaxCase]rune{'か' - 'ゕ', 0, 0}},
		unicode.CaseRange{'ゖ', 'ゖ', [unicode.MaxCase]rune{'け' - 'ゖ', 0, 0}},
		unicode.CaseRange{'っ', 'っ', [unicode.MaxCase]rune{'つ' - 'っ', 0, 0}},
		unicode.CaseRange{'ゃ', 'ゃ', [unicode.MaxCase]rune{'や' - 'ゃ', 0, 0}},
		unicode.CaseRange{'ゅ', 'ゅ', [unicode.MaxCase]rune{'ゆ' - 'ゅ', 0, 0}},
		unicode.CaseRange{'ょ', 'ょ', [unicode.MaxCase]rune{'よ' - 'ょ', 0, 0}},
		unicode.CaseRange{'ゎ', 'ゎ', [unicode.MaxCase]rune{'わ' - 'ゎ', 0, 0}},

		unicode.CaseRange{'ア', 'ア', [unicode.MaxCase]rune{0, 'ァ' - 'ア', 0}},
		unicode.CaseRange{'イ', 'イ', [unicode.MaxCase]rune{0, 'ィ' - 'イ', 0}},
		unicode.CaseRange{'ウ', 'ウ', [unicode.MaxCase]rune{0, 'ゥ' - 'ウ', 0}},
		unicode.CaseRange{'エ', 'エ', [unicode.MaxCase]rune{0, 'ェ' - 'エ', 0}},
		unicode.CaseRange{'オ', 'オ', [unicode.MaxCase]rune{0, 'ォ' - 'オ', 0}},
		unicode.CaseRange{'カ', 'カ', [unicode.MaxCase]rune{0, 'ヵ' - 'カ', 0}},
		unicode.CaseRange{'ケ', 'ケ', [unicode.MaxCase]rune{0, 'ヶ' - 'ケ', 0}},
		unicode.CaseRange{'ツ', 'ツ', [unicode.MaxCase]rune{0, 'ッ' - 'ツ', 0}},
		unicode.CaseRange{'ヤ', 'ヤ', [unicode.MaxCase]rune{0, 'ャ' - 'ヤ', 0}},
		unicode.CaseRange{'ユ', 'ユ', [unicode.MaxCase]rune{0, 'ュ' - 'ユ', 0}},
		unicode.CaseRange{'ヨ', 'ヨ', [unicode.MaxCase]rune{0, 'ョ' - 'ヨ', 0}},
		unicode.CaseRange{'ワ', 'ワ', [unicode.MaxCase]rune{0, 'ヮ' - 'ワ', 0}},

		unicode.CaseRange{'ァ', 'ァ', [unicode.MaxCase]rune{'ア' - 'ァ', 0, 0}},
		unicode.CaseRange{'ィ', 'ィ', [unicode.MaxCase]rune{'イ' - 'ィ', 0, 0}},
		unicode.CaseRange{'ゥ', 'ゥ', [unicode.MaxCase]rune{'ウ' - 'ゥ', 0, 0}},
		unicode.CaseRange{'ェ', 'ェ', [unicode.MaxCase]rune{'エ' - 'ェ', 0, 0}},
		unicode.CaseRange{'ォ', 'ォ', [unicode.MaxCase]rune{'オ' - 'ォ', 0, 0}},
		unicode.CaseRange{'ヵ', 'ヵ', [unicode.MaxCase]rune{'カ' - 'ヵ', 0, 0}},
		unicode.CaseRange{'ヶ', 'ヶ', [unicode.MaxCase]rune{'ケ' - 'ヶ', 0, 0}},
		unicode.CaseRange{'ッ', 'ッ', [unicode.MaxCase]rune{'ツ' - 'ッ', 0, 0}},
		unicode.CaseRange{'ャ', 'ャ', [unicode.MaxCase]rune{'ヤ' - 'ャ', 0, 0}},
		unicode.CaseRange{'ュ', 'ュ', [unicode.MaxCase]rune{'ユ' - 'ュ', 0, 0}},
		unicode.CaseRange{'ョ', 'ョ', [unicode.MaxCase]rune{'ヨ' - 'ョ', 0, 0}},
		unicode.CaseRange{'ヮ', 'ヮ', [unicode.MaxCase]rune{'ワ' - 'ヮ', 0, 0}},

		unicode.CaseRange{'ｱ', 'ｱ', [unicode.MaxCase]rune{0, 'ｧ' - 'ｱ', 0}},
		unicode.CaseRange{'ｲ', 'ｲ', [unicode.MaxCase]rune{0, 'ｨ' - 'ｲ', 0}},
		unicode.CaseRange{'ｳ', 'ｳ', [unicode.MaxCase]rune{0, 'ｩ' - 'ｳ', 0}},
		unicode.CaseRange{'ｴ', 'ｴ', [unicode.MaxCase]rune{0, 'ｪ' - 'ｴ', 0}},
		unicode.CaseRange{'ｵ', 'ｵ', [unicode.MaxCase]rune{0, 'ｫ' - 'ｵ', 0}},
		unicode.CaseRange{'ﾂ', 'ﾂ', [unicode.MaxCase]rune{0, 'ｯ' - 'ﾂ', 0}},
		unicode.CaseRange{'ﾔ', 'ﾔ', [unicode.MaxCase]rune{0, 'ｬ' - 'ﾔ', 0}},
		unicode.CaseRange{'ﾕ', 'ﾕ', [unicode.MaxCase]rune{0, 'ｭ' - 'ﾕ', 0}},
		unicode.CaseRange{'ﾖ', 'ﾖ', [unicode.MaxCase]rune{0, 'ｮ' - 'ﾖ', 0}},

		unicode.CaseRange{'ｧ', 'ｧ', [unicode.MaxCase]rune{'ｱ' - 'ｧ', 0, 0}},
		unicode.CaseRange{'ｨ', 'ｨ', [unicode.MaxCase]rune{'ｲ' - 'ｨ', 0, 0}},
		unicode.CaseRange{'ｩ', 'ｩ', [unicode.MaxCase]rune{'ｳ' - 'ｩ', 0, 0}},
		unicode.CaseRange{'ｪ', 'ｪ', [unicode.MaxCase]rune{'ｴ' - 'ｪ', 0, 0}},
		unicode.CaseRange{'ｫ', 'ｫ', [unicode.MaxCase]rune{'ｵ' - 'ｫ', 0, 0}},
		unicode.CaseRange{'ｯ', 'ｯ', [unicode.MaxCase]rune{'ﾂ' - 'ｯ', 0, 0}},
		unicode.CaseRange{'ｬ', 'ｬ', [unicode.MaxCase]rune{'ﾔ' - 'ｬ', 0, 0}},
		unicode.CaseRange{'ｭ', 'ｭ', [unicode.MaxCase]rune{'ﾕ' - 'ｭ', 0, 0}},
		unicode.CaseRange{'ｮ', 'ｮ', [unicode.MaxCase]rune{'ﾖ' - 'ｮ', 0, 0}}}
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
