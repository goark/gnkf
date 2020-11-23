package enc

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

var (
	textUTF8 = []byte("こんにちは，世界！\n私の名前は Spiegel です。")
	textSJIS = []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x81, 0x43, 0x90, 0xa2, 0x8a, 0x45, 0x81, 0x49, 0x0a, 0x8e, 0x84, 0x82, 0xcc, 0x96, 0xbc, 0x91, 0x4f, 0x82, 0xcd, 0x20, 0x53, 0x70, 0x69, 0x65, 0x67, 0x65, 0x6c, 0x20, 0x82, 0xc5, 0x82, 0xb7, 0x81, 0x42}
	textEUC  = []byte{0xa4, 0xb3, 0xa4, 0xf3, 0xa4, 0xcb, 0xa4, 0xc1, 0xa4, 0xcf, 0xa1, 0xa4, 0xc0, 0xa4, 0xb3, 0xa6, 0xa1, 0xaa, 0x0a, 0xbb, 0xe4, 0xa4, 0xce, 0xcc, 0xbe, 0xc1, 0xb0, 0xa4, 0xcf, 0x20, 0x53, 0x70, 0x69, 0x65, 0x67, 0x65, 0x6c, 0x20, 0xa4, 0xc7, 0xa4, 0xb9, 0xa1, 0xa3}
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		inp, out []byte
		ianaName string
		err      error
	}{
		{inp: textUTF8, out: textSJIS, ianaName: "shift_jis", err: nil},
		{inp: textUTF8, out: textEUC, ianaName: "euc-jp", err: nil},
		{inp: textUTF8, out: textUTF8, ianaName: "utf-8", err: nil},
		{inp: textUTF8, out: []byte{}, ianaName: "foo", err: ecode.ErrNotSuppotEncoding},
		{inp: textUTF8, out: []byte{}, ianaName: "us-ascii", err: ecode.ErrInvalidEncoding},
		{inp: textEUC, out: []byte{}, ianaName: "shift_jis", err: ecode.ErrInvalidEncoding},
		{inp: textSJIS, out: []byte{}, ianaName: "euc-jp", err: ecode.ErrInvalidEncoding},
	}
	for _, tc := range testCases {
		buf := &bytes.Buffer{}
		if err := Encode(tc.ianaName, buf, bytes.NewReader(tc.inp)); err != nil {
			if !errs.Is(err, tc.err) {
				t.Errorf("Encode() error = \"%+v\", want \"%+v\".", err, tc.err)
			}
		} else if !bytes.Equal(buf.Bytes(), tc.out) {
			t.Errorf("Encode(%s) result wrong translation.", tc.ianaName)
		}
	}
}

func TestDecode(t *testing.T) {
	testCases := []struct {
		inp, out []byte
		ianaName string
		err      error
	}{
		{inp: textSJIS, out: textUTF8, ianaName: "shift_jis", err: nil},
		{inp: textEUC, out: textUTF8, ianaName: "euc-jp", err: nil},
		{inp: textUTF8, out: textUTF8, ianaName: "utf-8", err: nil},
		{inp: textUTF8, out: []byte{}, ianaName: "foo", err: ecode.ErrNotSuppotEncoding},
	}
	for _, tc := range testCases {
		buf := &bytes.Buffer{}
		if err := Decode(buf, tc.ianaName, bytes.NewReader(tc.inp)); err != nil {
			if !errs.Is(err, tc.err) {
				t.Errorf("Decode() error = \"%+v\", want \"%+v\".", err, tc.err)
			}
		} else if !bytes.Equal(buf.Bytes(), tc.out) {
			t.Errorf("Decode(%s) result wrong translation.", tc.ianaName)
		}
	}
}

func TestTranslate(t *testing.T) {
	testCases := []struct {
		inp, out []byte
		from, to string
		err      error
	}{
		{inp: textUTF8, out: textSJIS, from: "utf-8", to: "shift_jis", err: nil},
		{inp: textUTF8, out: textEUC, from: "utf-8", to: "euc-jp", err: nil},
		{inp: textSJIS, out: textUTF8, from: "shift_jis", to: "utf-8", err: nil},
		{inp: textEUC, out: textUTF8, from: "euc-jp", to: "utf-8", err: nil},
		{inp: textSJIS, out: textEUC, from: "shift_jis", to: "euc-jp", err: nil},
		{inp: textEUC, out: textSJIS, from: "euc-jp", to: "shift_jis", err: nil},
		{inp: textSJIS, out: textSJIS, from: "shift_jis", to: "shift_jis", err: nil},
		{inp: textUTF8, out: textUTF8, from: "utf-8", to: "utf-8", err: nil},
		{inp: textUTF8, out: textUTF8, from: "foo", to: "utf-8", err: ecode.ErrNotSuppotEncoding},
		{inp: textUTF8, out: textUTF8, from: "utf-8", to: "bar", err: ecode.ErrNotSuppotEncoding},
		{inp: textSJIS, out: textEUC, from: "euc-jp", to: "shift_jis", err: ecode.ErrInvalidEncoding},
	}
	for _, tc := range testCases {
		buf := &bytes.Buffer{}
		if err := Convert(tc.to, buf, tc.from, bytes.NewReader(tc.inp)); err != nil {
			if !errs.Is(err, tc.err) {
				t.Errorf("Encode() error = \"%+v\", want \"%+v\".", err, tc.err)
			}
		} else if !bytes.Equal(buf.Bytes(), tc.out) {
			t.Errorf("Encode(%s -> %s) result wrong translation.", tc.from, tc.to)
		}
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
