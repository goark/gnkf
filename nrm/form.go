package nrm

import (
	"strings"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
	"golang.org/x/text/unicode/norm"
)

var formNamesMap = map[string]norm.Form{
	"nfc":  norm.NFC,
	"nfd":  norm.NFD,
	"nfkc": norm.NFKC,
	"nfkd": norm.NFKD,
}

func formName(f norm.Form) string {
	for key, value := range formNamesMap {
		if value == f {
			return key
		}
	}
	return ""
}

//FormList returns list of Unicode normalization form
func FormList() []string {
	return []string{
		formName(norm.NFC),
		formName(norm.NFD),
		formName(norm.NFKC),
		formName(norm.NFKD),
	}
}

//FormOf returns Unicode normalization form type from name string
func FormOf(name string) (norm.Form, error) {
	if f, ok := formNamesMap[strings.ToLower(name)]; ok {
		return f, nil
	}
	return norm.Form(0), errs.Wrap(ecode.ErrInvalidNormForm, errs.WithContext("name", name))
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
