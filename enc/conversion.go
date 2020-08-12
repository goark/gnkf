package enc

import (
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

//Convert function converts character encoding text stream.
func Convert(toIanaName string, writer io.Writer, fromIanaName string, txt io.Reader) error {
	encoder, err := Encoding(toIanaName)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("toIanaName", toIanaName))
	}
	decoder, err := Encoding(fromIanaName)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("fromIanaName", fromIanaName))
	}
	if encoder == unicode.UTF8 {
		return decode(decoder, writer, txt)
	}
	if decoder == unicode.UTF8 {
		return encode(encoder, writer, txt)
	}
	return convert(encoder, decoder, writer, txt)
}

func convert(encoder, decoder encoding.Encoding, writer io.Writer, txt io.Reader) error {
	if encoder == decoder {
		return notConvert(writer, txt)
	}
	if _, err := io.Copy(encoder.NewEncoder().Writer(writer), decoder.NewDecoder().Reader(txt)); err != nil {
		return errs.Wrap(ecode.ErrInvalidEncoding, errs.WithCause(err))
	}
	return nil
}

func notConvert(writer io.Writer, txt io.Reader) error {
	if _, err := io.Copy(writer, txt); err != nil {
		return errs.Wrap(err)
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
