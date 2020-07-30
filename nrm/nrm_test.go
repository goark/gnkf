package nrm

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/dump"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

func TestFormList(t *testing.T) {
	res := "nfc|nfd|nfkc|nfkd"
	str := strings.Join(FormList(), "|")
	if str != res {
		t.Errorf("FormList() = \"%+v\", want \"%+v\".", str, res)
	}
}

func TestNormalize(t *testing.T) {
	testCases := []struct {
		inp, out []byte
		formName string
		err      error
	}{
		{
			inp:      []byte{0xe3, 0x83, 0x98, 0xe3, 0x82, 0x9a, 0xe3, 0x83, 0xb3, 0xe3, 0x82, 0xad, 0xe3, 0x82, 0x99, 0xe3, 0x83, 0xb3},
			out:      []byte("ペンギン"),
			formName: "nfc",
			err:      nil,
		},
		{
			inp:      []byte("ペンギン"),
			out:      []byte{0xe3, 0x83, 0x98, 0xe3, 0x82, 0x9a, 0xe3, 0x83, 0xb3, 0xe3, 0x82, 0xad, 0xe3, 0x82, 0x99, 0xe3, 0x83, 0xb3},
			formName: "nfd",
			err:      nil,
		},
		{
			inp:      []byte("ﾍﾟﾝｷﾞﾝ"),
			out:      []byte("ペンギン"),
			formName: "nfkc",
			err:      nil,
		},
		{
			inp:      []byte("ﾍﾟﾝｷﾞﾝ"),
			out:      []byte{0xe3, 0x83, 0x98, 0xe3, 0x82, 0x9a, 0xe3, 0x83, 0xb3, 0xe3, 0x82, 0xad, 0xe3, 0x82, 0x99, 0xe3, 0x83, 0xb3},
			formName: "nfkd",
			err:      nil,
		},
		{
			inp:      []byte("ペンギン"),
			out:      []byte{},
			formName: "foo",
			err:      ecode.ErrInvalidNormForm,
		},
	}
	for _, tc := range testCases {
		buf := &bytes.Buffer{}
		if err := Normalize(tc.formName, buf, bytes.NewReader(tc.inp), false); err != nil {
			if !errs.Is(err, tc.err) {
				t.Errorf("Normalize() error = \"%+v\", want \"%+v\".", err, tc.err)
			}
		} else if !bytes.Equal(buf.Bytes(), tc.out) {
			t.Errorf("Normalize(%s) result wrong translation: ", tc.formName)
			dump.Octet(os.Stdout, buf)
		}
	}
}

func TestNormKangxiRadicals(t *testing.T) {
	testCases := []struct {
		inp, out []byte
		formName string
		err      error
	}{
		{
			inp:      []byte("㈱埼⽟"), //U+3231, U+57FC, U+2F5F
			out:      []byte("㈱埼⽟"), //U+3231, U+57FC, U+2F5F (not translate)
			formName: "nfc",
			err:      nil,
		},
		{
			inp:      []byte("㈱埼⽟"), //U+3231, U+57FC, U+2F5F
			out:      []byte("㈱埼⽟"), //U+3231, U+57FC, U+2F5F (not translate)
			formName: "nfd",
			err:      nil,
		},
		{
			inp:      []byte("㈱埼⽟"), //U+3231, U+57FC, U+2F5F
			out:      []byte("㈱埼玉"), //U+3231, U+57FC, U+7389
			formName: "nfkc",
			err:      nil,
		},
		{
			inp:      []byte("㈱埼⽟"), //U+3231, U+57FC, U+2F5F
			out:      []byte("㈱埼玉"), //U+3231, U+57FC, U+7389
			formName: "nfkd",
			err:      nil,
		},
		{
			inp:      []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x81, 0x43, 0x90, 0xa2, 0x8a, 0x45, 0x81, 0x49}, //"こんにちは，世界！" by Shift_JIS encoding
			out:      []byte{},
			formName: "nfkc",
			err:      ecode.ErrInvalidUTF8Text,
		},
	}
	for _, tc := range testCases {
		buf := &bytes.Buffer{}
		if err := Normalize(tc.formName, buf, bytes.NewReader(tc.inp), true); err != nil {
			if !errs.Is(err, tc.err) {
				t.Errorf("NormKangxiRadicals() error = \"%+v\", want \"%+v\".", err, tc.err)
			}
		} else if !bytes.Equal(buf.Bytes(), tc.out) {
			t.Error("NormKangxiRadicals() result wrong translation: ")
			dump.Octet(os.Stdout, buf)
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
