package guess

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
)

var (
	textUTF8  = []byte("こんにちは，世界！\n私の名前は Spiegel です。")
	textSJIS  = []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x81, 0x43, 0x90, 0xa2, 0x8a, 0x45, 0x81, 0x49, 0x0a, 0x8e, 0x84, 0x82, 0xcc, 0x96, 0xbc, 0x91, 0x4f, 0x82, 0xcd, 0x20, 0x53, 0x70, 0x69, 0x65, 0x67, 0x65, 0x6c, 0x20, 0x82, 0xc5, 0x82, 0xb7, 0x81, 0x42}
	textEUC   = []byte{0xa4, 0xb3, 0xa4, 0xf3, 0xa4, 0xcb, 0xa4, 0xc1, 0xa4, 0xcf, 0xa1, 0xa4, 0xc0, 0xa4, 0xb3, 0xa6, 0xa1, 0xaa, 0x0a, 0xbb, 0xe4, 0xa4, 0xce, 0xcc, 0xbe, 0xc1, 0xb0, 0xa4, 0xcf, 0x20, 0x53, 0x70, 0x69, 0x65, 0x67, 0x65, 0x6c, 0x20, 0xa4, 0xc7, 0xa4, 0xb9, 0xa1, 0xa3}
	testCases = []struct {
		text []byte
		res  string
		err  error
	}{
		{text: textUTF8, res: "UTF-8,windows-1252,windows-1253,Shift_JIS,windows-1255", err: nil},
		{text: textSJIS, res: "Shift_JIS,windows-1252,Big5,GB-18030,KOI8-R", err: nil},
		{text: textEUC, res: "EUC-JP,Big5,GB-18030,ISO-8859-7,EUC-KR,Shift_JIS,ISO-8859-1", err: nil},
		{text: []byte{0xff}, res: "", err: ecode.ErrCannotDetect},
		{text: nil, res: "UTF-8", err: nil},
	}
)

func TestEncodingBytes(t *testing.T) {
	for _, tc := range testCases {
		res, err := EncodingBytes(tc.text)
		if !errs.Is(err, tc.err) {
			t.Errorf("EncodingBytes() error = \"%+v\", want \"%+v\".", err, tc.err)
		}
		str := strings.Join(res, ",")
		if str != tc.res {
			t.Errorf("EncodingBytes() = \"%v\", want \"%v\".", str, tc.res)
		}
	}
}

func TestEncoding(t *testing.T) {
	for _, tc := range testCases {
		res, err := Encoding(bytes.NewReader(tc.text))
		if !errs.Is(err, tc.err) {
			t.Errorf("Encoding() error = \"%+v\", want \"%+v\".", err, tc.err)
		}
		str := strings.Join(res, ",")
		if str != tc.res {
			t.Errorf("Encoding() = \"%v\", want \"%v\".", str, tc.res)
		}
	}
}

func TestEncodingNil(t *testing.T) {
	_, err := Encoding(io.Reader(nil))
	if !errs.Is(err, ecode.ErrNullPointer) {
		t.Errorf("Encoding() error = \"%+v\", want \"%+v\".", err, ecode.ErrNullPointer)
	}
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
