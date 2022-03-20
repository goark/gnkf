package rbom

import (
	"bytes"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
)

var bom = []byte{0xef, 0xbb, 0xbf}

//RemoveBom removes BOM character in UTF-8 stream
func RemoveBom(r io.Reader) ([]byte, error) {
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, errs.Wrap(err)
	}
	return RemoveBomBytes(buf.Bytes())
}

//RemoveBomBytes removes BOM character in UTF-8 byte string
func RemoveBomBytes(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return []byte{}, nil
	}
	if !utf8.Valid(b) {
		return nil, errs.Wrap(ecode.ErrInvalidUTF8Text)
	}
	return bytes.ReplaceAll(b, bom, []byte{}), nil
}

//RemoveBomString removes BOM character in UTF-8 string
func RemoveBomString(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ReplaceAll(s, string(bom), "")
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
