package kana

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/gnkf/dump"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

func TestFormList(t *testing.T) {
	res := "hiragana|katakana|chokuon"
	str := strings.Join(FormList(), "|")
	if str != res {
		t.Errorf("FormList() = \"%+v\", want \"%+v\".", str, res)
	}
}

func TestTranslate(t *testing.T) {
	testCases := []struct {
		inp, out []byte
		formName string
		err      error
	}{
		{
			inp:      []byte("あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			out:      []byte("あいうえおわゐゑをんゔゕゖゝゞあいうえおわゐゑをんゔゕゖゝゞｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			formName: "hiragana",
			err:      nil,
		},
		{
			inp:      []byte("あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			out:      []byte("アイウエオワヰヱヲンヴヵヶヽヾアイウエオワヰヱヲンヴヵヶヽヾｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			formName: "katakana",
			err:      nil,
		},
		{
			inp:      []byte("あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			out:      []byte("あいうえおわゐゑをんゔかけゝゞアイウエオワヰヱヲンヴカケヽヾｱｲｳｴｵﾂﾔﾕﾖｱｲｳｴｵﾂﾔﾕﾖ"),
			formName: "chokuon",
			err:      nil,
		},
		{
			inp:      []byte("あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			out:      []byte{},
			formName: "foo",
			err:      ecode.ErrInvalidKanaForm,
		},
	}
	for _, tc := range testCases {
		f, err := FormOf(tc.formName)
		if !errors.Is(err, tc.err) {
			t.Errorf("FormOf(%v) error = \"%+v\", want \"%+v\".", tc.formName, err, tc.err)
		}
		if err == nil {
			buf := &bytes.Buffer{}
			if err := Convert(f, buf, bytes.NewReader(tc.inp), false); err != nil {
				if err != nil {
					t.Errorf("Convert() error = \"%+v\", want nil.", err)
				}
			} else if !bytes.Equal(buf.Bytes(), tc.out) {
				fmt.Println(buf.String())
				t.Errorf("Convert(%s) -> %s, want %s", tc.formName, dump.OctetString(bytes.NewReader(tc.inp)), dump.OctetString(buf))
			}
		}
	}
}

func TestTranslateFold(t *testing.T) {
	testCases := []struct {
		inp, out []byte
		formName string
		err      error
	}{
		{
			inp:      []byte("あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			out:      []byte("あいうえおわゐゑをんゔゕゖゝゞあいうえおわゐゑをんゔゕゖゝゞあいうえおつやゆよぁぃぅぇぉっゃゅょ"),
			formName: "hiragana",
			err:      nil,
		},
		{
			inp:      []byte("あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			out:      []byte("アイウエオワヰヱヲンヴヵヶヽヾアイウエオワヰヱヲンヴヵヶヽヾアイウエオツヤユヨァィゥェォッャュョ"),
			formName: "katakana",
			err:      nil,
		},
		{
			inp:      []byte("あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			out:      []byte("あいうえおわゐゑをんゔかけゝゞアイウエオワヰヱヲンヴカケヽヾアイウエオツヤユヨアイウエオツヤユヨ"),
			formName: "chokuon",
			err:      nil,
		},
		{
			inp:      []byte("あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾｱｲｳｴｵﾂﾔﾕﾖｧｨｩｪｫｯｬｭｮ"),
			out:      []byte{},
			formName: "foo",
			err:      ecode.ErrInvalidKanaForm,
		},
	}
	for _, tc := range testCases {
		f, err := FormOf(tc.formName)
		if !errors.Is(err, tc.err) {
			t.Errorf("FormOf(%v) error = \"%+v\", want \"%+v\".", tc.formName, err, tc.err)
		}
		if err == nil {
			buf := &bytes.Buffer{}
			if err := Convert(f, buf, bytes.NewReader(tc.inp), true); err != nil {
				if err != nil {
					t.Errorf("Convert() error = \"%+v\", want nil.", err)
				}
			} else if !bytes.Equal(buf.Bytes(), tc.out) {
				fmt.Println(buf.String())
				t.Errorf("Translate(%s) -> %s, want %s", tc.formName, dump.OctetString(bytes.NewReader(tc.inp)), dump.OctetString(buf))
			}
		}
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
