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
		unicode.CaseRange{Lo: 0x1b150, Hi: 0x1b152, Delta: [unicode.MaxCase]rune{0x1b164 - 0x1b150, 0, 0}},
		unicode.CaseRange{Lo: 0x1b164, Hi: 0x1b166, Delta: [unicode.MaxCase]rune{0, 0x1b150 - 0x1b164, 0}},
	}
	replacekanaMap = map[string]string{
		string([]rune{'ヷ'}): string([]rune{'わ', 0x3099}),
		string([]rune{'ヸ'}): string([]rune{'ゐ', 0x3099}),
		string([]rune{'ヹ'}): string([]rune{'ゑ', 0x3099}),
		string([]rune{'ヺ'}): string([]rune{'を', 0x3099}),
	}
	replacekanaMap2 = map[string]string{
		string([]rune{'ㇰ'}):     string([]rune{'く'}),
		string([]rune{'ㇱ'}):     string([]rune{'し'}),
		string([]rune{'ㇲ'}):     string([]rune{'す'}),
		string([]rune{'ㇳ'}):     string([]rune{'と'}),
		string([]rune{'ㇴ'}):     string([]rune{'ぬ'}),
		string([]rune{'ㇵ'}):     string([]rune{'は'}),
		string([]rune{'ㇶ'}):     string([]rune{'ひ'}),
		string([]rune{'ㇷ'}):     string([]rune{'ふ'}),
		string([]rune{'ㇸ'}):     string([]rune{'へ'}),
		string([]rune{'ㇹ'}):     string([]rune{'ほ'}),
		string([]rune{'ㇺ'}):     string([]rune{'む'}),
		string([]rune{'ㇻ'}):     string([]rune{'ら'}),
		string([]rune{'ㇼ'}):     string([]rune{'り'}),
		string([]rune{'ㇽ'}):     string([]rune{'る'}),
		string([]rune{'ㇾ'}):     string([]rune{'れ'}),
		string([]rune{'ㇿ'}):     string([]rune{'ろ'}),
		string([]rune{0x1b167}): string([]rune{'ん'}),
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
		unicode.CaseRange{Lo: 'ゐ', Hi: 'ゐ', Delta: [unicode.MaxCase]rune{0, 0x1b150 - 'ゐ', 0}},
		unicode.CaseRange{Lo: 'ゑ', Hi: 'ゑ', Delta: [unicode.MaxCase]rune{0, 0x1b151 - 'ゑ', 0}},
		unicode.CaseRange{Lo: 'を', Hi: 'を', Delta: [unicode.MaxCase]rune{0, 0x1b152 - 'を', 0}},

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
		unicode.CaseRange{Lo: 0x1b150, Hi: 0x1b150, Delta: [unicode.MaxCase]rune{'ゐ' - 0x1b150, 0, 0}},
		unicode.CaseRange{Lo: 0x1b151, Hi: 0x1b151, Delta: [unicode.MaxCase]rune{'ゑ' - 0x1b151, 0, 0}},
		unicode.CaseRange{Lo: 0x1b152, Hi: 0x1b152, Delta: [unicode.MaxCase]rune{'を' - 0x1b152, 0, 0}},

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
		unicode.CaseRange{Lo: 'ク', Hi: 'ク', Delta: [unicode.MaxCase]rune{0, 'ク' - 'ㇰ', 0}},
		unicode.CaseRange{Lo: 'シ', Hi: 'シ', Delta: [unicode.MaxCase]rune{0, 'シ' - 'ㇱ', 0}},
		unicode.CaseRange{Lo: 'ス', Hi: 'ス', Delta: [unicode.MaxCase]rune{0, 'ス' - 'ㇲ', 0}},
		unicode.CaseRange{Lo: 'ト', Hi: 'ト', Delta: [unicode.MaxCase]rune{0, 'ト' - 'ㇳ', 0}},
		unicode.CaseRange{Lo: 'ヌ', Hi: 'ヌ', Delta: [unicode.MaxCase]rune{0, 'ヌ' - 'ㇴ', 0}},
		unicode.CaseRange{Lo: 'ハ', Hi: 'ハ', Delta: [unicode.MaxCase]rune{0, 'ハ' - 'ㇵ', 0}},
		unicode.CaseRange{Lo: 'ヒ', Hi: 'ヒ', Delta: [unicode.MaxCase]rune{0, 'ヒ' - 'ㇶ', 0}},
		unicode.CaseRange{Lo: 'フ', Hi: 'フ', Delta: [unicode.MaxCase]rune{0, 'フ' - 'ㇷ', 0}},
		unicode.CaseRange{Lo: 'ヘ', Hi: 'ヘ', Delta: [unicode.MaxCase]rune{0, 'ヘ' - 'ㇸ', 0}},
		unicode.CaseRange{Lo: 'ホ', Hi: 'ホ', Delta: [unicode.MaxCase]rune{0, 'ホ' - 'ㇹ', 0}},
		unicode.CaseRange{Lo: 'ム', Hi: 'ム', Delta: [unicode.MaxCase]rune{0, 'ム' - 'ㇺ', 0}},
		unicode.CaseRange{Lo: 'ラ', Hi: 'ラ', Delta: [unicode.MaxCase]rune{0, 'ラ' - 'ㇻ', 0}},
		unicode.CaseRange{Lo: 'リ', Hi: 'リ', Delta: [unicode.MaxCase]rune{0, 'リ' - 'ㇼ', 0}},
		unicode.CaseRange{Lo: 'ル', Hi: 'ル', Delta: [unicode.MaxCase]rune{0, 'ル' - 'ㇽ', 0}},
		unicode.CaseRange{Lo: 'レ', Hi: 'レ', Delta: [unicode.MaxCase]rune{0, 'レ' - 'ㇾ', 0}},
		unicode.CaseRange{Lo: 'ロ', Hi: 'ロ', Delta: [unicode.MaxCase]rune{0, 'ロ' - 'ㇿ', 0}},
		unicode.CaseRange{Lo: 'ヰ', Hi: 'ヰ', Delta: [unicode.MaxCase]rune{0, 'ヰ' - 0x1b164, 0}},
		unicode.CaseRange{Lo: 'ヱ', Hi: 'ヱ', Delta: [unicode.MaxCase]rune{0, 'ヱ' - 0x1b165, 0}},
		unicode.CaseRange{Lo: 'ヲ', Hi: 'ヲ', Delta: [unicode.MaxCase]rune{0, 'ヲ' - 0x1b166, 0}},
		unicode.CaseRange{Lo: 'ン', Hi: 'ン', Delta: [unicode.MaxCase]rune{0, 'ン' - 0x1b167, 0}},

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
		unicode.CaseRange{Lo: 'ㇰ', Hi: 'ㇰ', Delta: [unicode.MaxCase]rune{'ク' - 'ㇰ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇱ', Hi: 'ㇱ', Delta: [unicode.MaxCase]rune{'シ' - 'ㇱ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇲ', Hi: 'ㇲ', Delta: [unicode.MaxCase]rune{'ス' - 'ㇲ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇳ', Hi: 'ㇳ', Delta: [unicode.MaxCase]rune{'ト' - 'ㇳ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇴ', Hi: 'ㇴ', Delta: [unicode.MaxCase]rune{'ヌ' - 'ㇴ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇵ', Hi: 'ㇵ', Delta: [unicode.MaxCase]rune{'ハ' - 'ㇵ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇶ', Hi: 'ㇶ', Delta: [unicode.MaxCase]rune{'ヒ' - 'ㇶ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇷ', Hi: 'ㇷ', Delta: [unicode.MaxCase]rune{'フ' - 'ㇷ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇸ', Hi: 'ㇸ', Delta: [unicode.MaxCase]rune{'ヘ' - 'ㇸ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇹ', Hi: 'ㇹ', Delta: [unicode.MaxCase]rune{'ホ' - 'ㇹ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇺ', Hi: 'ㇺ', Delta: [unicode.MaxCase]rune{'ム' - 'ㇺ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇻ', Hi: 'ㇻ', Delta: [unicode.MaxCase]rune{'ラ' - 'ㇻ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇼ', Hi: 'ㇼ', Delta: [unicode.MaxCase]rune{'リ' - 'ㇼ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇽ', Hi: 'ㇽ', Delta: [unicode.MaxCase]rune{'ル' - 'ㇽ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇾ', Hi: 'ㇾ', Delta: [unicode.MaxCase]rune{'レ' - 'ㇾ', 0, 0}},
		unicode.CaseRange{Lo: 'ㇿ', Hi: 'ㇿ', Delta: [unicode.MaxCase]rune{'ロ' - 'ㇿ', 0, 0}},
		unicode.CaseRange{Lo: 0x1b164, Hi: 0x1b164, Delta: [unicode.MaxCase]rune{'ヰ' - 0x1b164, 0, 0}},
		unicode.CaseRange{Lo: 0x1b165, Hi: 0x1b165, Delta: [unicode.MaxCase]rune{'ヱ' - 0x1b165, 0, 0}},
		unicode.CaseRange{Lo: 0x1b166, Hi: 0x1b166, Delta: [unicode.MaxCase]rune{'ヲ' - 0x1b166, 0, 0}},
		unicode.CaseRange{Lo: 0x1b167, Hi: 0x1b167, Delta: [unicode.MaxCase]rune{'ン' - 0x1b167, 0, 0}},

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
	for k, v := range replacekanaMap2 {
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
