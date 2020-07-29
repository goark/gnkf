package enc

import (
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func Translate(toIanaName string, writer io.Writer, fromIanaName string, txt io.Reader) error {
	encoder, err := GetEncoding(toIanaName)
	if err != nil {
		return errs.WrapWithCause(err, nil, errs.WithContext("toIanaName", toIanaName))
	}
	decoder, err := GetEncoding(fromIanaName)
	if err != nil {
		return errs.WrapWithCause(err, nil, errs.WithContext("fromIanaName", fromIanaName))
	}
	if encoder == unicode.UTF8 {
		return decode(decoder, writer, txt)
	}
	if decoder == unicode.UTF8 {
		return encode(encoder, writer, txt)
	}
	return translate(encoder, decoder, writer, txt)
}

func translate(encoder, decoder encoding.Encoding, writer io.Writer, txt io.Reader) error {
	if encoder == decoder {
		return notTranslate(writer, txt)
	}
	if _, err := io.Copy(encoder.NewEncoder().Writer(writer), decoder.NewDecoder().Reader(txt)); err != nil {
		return errs.WrapWithCause(ecode.ErrInvalidEncoding, err)
	}
	return nil
}

func notTranslate(writer io.Writer, txt io.Reader) error {
	if _, err := io.Copy(writer, txt); err != nil {
		return errs.WrapWithCause(err, nil)
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
