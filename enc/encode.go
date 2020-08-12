package enc

import (
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

//Encode converts UTF-8 from other character encoding text.
func Encode(ianaName string, writer io.Writer, txt io.Reader) error {
	encoder, err := Encoding(ianaName)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("ianaName", ianaName))
	}
	return encode(encoder, writer, txt)
}

func encode(encoder encoding.Encoding, writer io.Writer, txt io.Reader) error {
	if encoder == unicode.UTF8 {
		return notConvert(writer, txt)
	}
	if _, err := io.Copy(encoder.NewEncoder().Writer(writer), txt); err != nil {
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
