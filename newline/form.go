package newline

import (
	"strings"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

//Form is type of newline form
type Form int

const (
	LF   Form = iota //newline is '\n' only
	CR               //newline is '\r' only
	CRLF             //newline is '\r'+'\n'
)

var (
	formNamesMap = map[string]Form{
		"lf":   LF,
		"cr":   CR,
		"crlf": CRLF,
	}
	newlineCodeMap = map[Form]string{
		LF:   "\n",
		CR:   "\r",
		CRLF: "\r\n",
	}
)

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
		formName(LF),
		formName(CR),
		formName(CRLF),
	}
}

//FormOf returns newline form name string
func FormOf(name string) (Form, error) {
	if f, ok := formNamesMap[strings.ToLower(name)]; ok {
		return f, nil
	}
	return Form(0), errs.Wrap(ecode.ErrInvalidNewlineForm, errs.WithContext("name", name))
}

//Code returns newline code string
func (f Form) Code() string {
	if c, ok := newlineCodeMap[f]; ok {
		return c
	}
	return ""
}

//NewReplacer returns strings.Replacer instance for translating newline
func NewReplacer(frm Form) *strings.Replacer {
	return strings.NewReplacer(
		CRLF.Code(), frm.Code(),
		LF.Code(), frm.Code(),
		CR.Code(), frm.Code(),
	)
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
