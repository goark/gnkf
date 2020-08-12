package width

import (
	"strings"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
	wdth "golang.org/x/text/width"
)

var formNamesMap = map[string]wdth.Transformer{
	"fold":   wdth.Fold,
	"narrow": wdth.Narrow,
	"widen":  wdth.Widen,
}

func formName(f wdth.Transformer) string {
	for key, value := range formNamesMap {
		if value == f {
			return key
		}
	}
	return ""
}

//FormList returns list of width form
func FormList() []string {
	return []string{
		formName(wdth.Fold),
		formName(wdth.Narrow),
		formName(wdth.Widen),
	}
}

//FormOf returns Unicode normalization form type from name string
func FormOf(name string) (wdth.Transformer, error) {
	if f, ok := formNamesMap[strings.ToLower(name)]; ok {
		return f, nil
	}
	return wdth.Fold, errs.Wrap(ecode.ErrInvalidWidthForm, errs.WithContext("name", name))
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
