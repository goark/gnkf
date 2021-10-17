package kana

import (
	"bytes"
	"io"
	"strings"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/kkconv"
)

//Convert function converts kana character in text stream.
func Convert(f Form, writer io.Writer, txt io.Reader, foldFlag bool) error {
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(txt); err != nil {
		return errs.Wrap(err)
	}
	if _, err := strings.NewReader(ConvertString(f, buf.String(), foldFlag)).WriteTo(writer); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

//ConvertString function converts kana character in text string.
func ConvertString(f Form, txt string, foldFlag bool) string {
	switch f {
	case Hiragana:
		return kkconv.Hiragana(txt, foldFlag)
	case Katakana:
		return kkconv.Katakana(txt, foldFlag)
	case Chokuon:
		return kkconv.Chokuon(txt, foldFlag)
	default:
		return txt
	}
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
