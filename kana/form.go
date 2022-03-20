package kana

import (
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
)

//Form is type of newline form
type Form int

const (
	Hiragana Form = iota //Hiragana form
	Katakana             //Katakana form
	Chokuon              //Chokuon (Upper kana) form
)

var formNamesMap = map[string]Form{
	"hiragana": Hiragana,
	"katakana": Katakana,
	"chokuon":  Chokuon,
}

func (f Form) String() string {
	return formName(f)
}

func formName(f Form) string {
	for key, value := range formNamesMap {
		if value == f {
			return key
		}
	}
	return ""
}

//FormList returns list of newline form
func FormList() []string {
	return []string{
		formName(Hiragana),
		formName(Katakana),
		formName(Chokuon),
	}
}

//FormOf returns newline form name string
func FormOf(name string) (Form, error) {
	if f, ok := formNamesMap[strings.ToLower(name)]; ok {
		return f, nil
	}
	return Form(0), errs.Wrap(ecode.ErrInvalidKanaForm, errs.WithContext("name", name))
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
