package width

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/dump"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

func TestFormList(t *testing.T) {
	res := "fold|narrow|widen"
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
			inp:      []byte("ペンギン 12345 ヸヹヺ"),
			out:      []byte("ﾍﾟﾝｷﾞﾝ 12345 ｲﾞｴﾞｦﾞ"),
			formName: "narrow",
			err:      nil,
		},
		{
			inp:      []byte("ﾍﾟﾝｷﾞﾝ 12345 ヸヹｦﾞ"),
			out:      []byte("ペンギン　１２３４５　ヸヹヺ"),
			formName: "widen",
			err:      nil,
		},
		{
			inp:      []byte("ﾍﾟﾝｷﾞﾝ　１２３４５ ヸヹｦﾞ"),
			out:      []byte("ペンギン 12345 ヸヹヺ"),
			formName: "fold",
			err:      nil,
		},
		{
			inp:      []byte("ペンギン 12345"),
			out:      []byte{},
			formName: "foo",
			err:      ecode.ErrInvalidWidthForm,
		},
	}
	for _, tc := range testCases {
		buf := &bytes.Buffer{}
		if err := Convert(tc.formName, buf, bytes.NewReader(tc.inp)); err != nil {
			if !errs.Is(err, tc.err) {
				t.Errorf("Translate() error = \"%+v\", want \"%+v\".", err, tc.err)
			}
		} else if !bytes.Equal(buf.Bytes(), tc.out) {
			fmt.Println(buf.String())
			t.Errorf("Translate(%s) -> %s, want %s", tc.formName, dump.OctetString(bytes.NewReader(tc.inp)), dump.OctetString(buf))
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
