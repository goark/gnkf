package enc

import (
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
)

//GetEncoding returns encoding.Encoding instance from IANA name
func GetEncoding(ianaName string) (encoding.Encoding, error) {
	e, err := ianaindex.IANA.Encoding(ianaName)
	if err != nil {
		e, err = ianaindex.MIME.Encoding(ianaName)
		if err != nil {
			return nil, errs.WrapWithCause(ecode.ErrNotSuppotEncoding, err, errs.WithContext("ianaName", ianaName))
		}
	}
	if e == nil {
		return nil, errs.WrapWithCause(ecode.ErrNotSuppotEncoding, err, errs.WithContext("ianaName", ianaName))
	}
	return e, nil
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
