package enc

import (
	"io"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

//Decode converts from UTF-8 encodeing text.
func Decode(writer io.Writer, ianaName string, txt io.Reader) error {
	decoder, err := Encoding(ianaName)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("ianaName", ianaName))
	}
	return decode(decoder, writer, txt)
}

func decode(decoder encoding.Encoding, writer io.Writer, txt io.Reader) error {
	if decoder == unicode.UTF8 {
		return notConvert(writer, txt)
	}
	if _, err := io.Copy(writer, decoder.NewDecoder().Reader(txt)); err != nil {
		return errs.Wrap(ecode.ErrInvalidEncoding, errs.WithCause(err))
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
