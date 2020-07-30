package nrm

import (
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"golang.org/x/text/unicode/norm"
)

//Normalize function normalize Unicode text
func Normalize(formName string, writer io.Writer, txt io.Reader, krFlag bool) error {
	f, err := FormOf(formName)
	if err != nil {
		return errs.WrapWithCause(err, nil, errs.WithContext("formName", formName))
	}
	if (f == norm.NFKC || f == norm.NFKD) && krFlag {
		return NormKangxiRadicals(writer, txt)
	}

	if _, err := io.Copy(writer, f.Reader(txt)); err != nil {
		return errs.WrapWithCause(err, nil, errs.WithContext("formName", formName))
	}
	return nil
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