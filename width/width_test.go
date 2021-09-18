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
		{
			//"ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶヷヸヹヺヽヾヿㇰㇱㇲㇳㇴㇵㇶㇷㇸㇹㇺㇻㇼㇽㇾㇿｦｧｨｩｪｫｬｭｮｯｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜﾝ" + string([]rune{0x1b164, 0x1b165, 0x1b166, 0x1b167}))
			inp:      []byte(string([]rune{0x30a1, 0x30a2, 0x30a3, 0x30a4, 0x30a5, 0x30a6, 0x30a7, 0x30a8, 0x30a9, 0x30aa, 0x30ab, 0x30ac, 0x30ad, 0x30ae, 0x30af, 0x30b0, 0x30b1, 0x30b2, 0x30b3, 0x30b4, 0x30b5, 0x30b6, 0x30b7, 0x30b8, 0x30b9, 0x30ba, 0x30bb, 0x30bc, 0x30bd, 0x30be, 0x30bf, 0x30c0, 0x30c1, 0x30c2, 0x30c3, 0x30c4, 0x30c5, 0x30c6, 0x30c7, 0x30c8, 0x30c9, 0x30ca, 0x30cb, 0x30cc, 0x30cd, 0x30ce, 0x30cf, 0x30d0, 0x30d1, 0x30d2, 0x30d3, 0x30d4, 0x30d5, 0x30d6, 0x30d7, 0x30d8, 0x30d9, 0x30da, 0x30db, 0x30dc, 0x30dd, 0x30de, 0x30df, 0x30e0, 0x30e1, 0x30e2, 0x30e3, 0x30e4, 0x30e5, 0x30e6, 0x30e7, 0x30e8, 0x30e9, 0x30ea, 0x30eb, 0x30ec, 0x30ed, 0x30ee, 0x30ef, 0x30f0, 0x30f1, 0x30f2, 0x30f3, 0x30f4, 0x30f5, 0x30f6, 0x30f7, 0x30f8, 0x30f9, 0x30fa, 0x30fd, 0x30fe, 0x30ff, 0x31f0, 0x31f1, 0x31f2, 0x31f3, 0x31f4, 0x31f5, 0x31f6, 0x31f7, 0x31f8, 0x31f9, 0x31fa, 0x31fb, 0x31fc, 0x31fd, 0x31fe, 0x31ff, 0xff66, 0xff67, 0xff68, 0xff69, 0xff6a, 0xff6b, 0xff6c, 0xff6d, 0xff6e, 0xff6f, 0xff71, 0xff72, 0xff73, 0xff74, 0xff75, 0xff76, 0xff77, 0xff78, 0xff79, 0xff7a, 0xff7b, 0xff7c, 0xff7d, 0xff7e, 0xff7f, 0xff80, 0xff81, 0xff82, 0xff83, 0xff84, 0xff85, 0xff86, 0xff87, 0xff88, 0xff89, 0xff8a, 0xff8b, 0xff8c, 0xff8d, 0xff8e, 0xff8f, 0xff90, 0xff91, 0xff92, 0xff93, 0xff94, 0xff95, 0xff96, 0xff97, 0xff98, 0xff99, 0xff9a, 0xff9b, 0xff9c, 0xff9d, 0x1b164, 0x1b165, 0x1b166, 0x1b167})),
			out:      []byte("ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶヷヸヹヺヽヾヿㇰㇱㇲㇳㇴㇵㇶㇷㇸㇹㇺㇻㇼㇽㇾㇿヲァィゥェォャュョッアイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワン" + string([]rune{0x1b164, 0x1b165, 0x1b166, 0x1b167})),
			formName: "widen",
			err:      nil,
		},
		{
			//"ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶヷヸヹヺヽヾヿㇰㇱㇲㇳㇴㇵㇶㇷㇸㇹㇺㇻㇼㇽㇾㇿｦｧｨｩｪｫｬｭｮｯｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜﾝ" + string([]rune{0x1b164, 0x1b165, 0x1b166, 0x1b167}))
			inp:      []byte(string([]rune{0x30a1, 0x30a2, 0x30a3, 0x30a4, 0x30a5, 0x30a6, 0x30a7, 0x30a8, 0x30a9, 0x30aa, 0x30ab, 0x30ac, 0x30ad, 0x30ae, 0x30af, 0x30b0, 0x30b1, 0x30b2, 0x30b3, 0x30b4, 0x30b5, 0x30b6, 0x30b7, 0x30b8, 0x30b9, 0x30ba, 0x30bb, 0x30bc, 0x30bd, 0x30be, 0x30bf, 0x30c0, 0x30c1, 0x30c2, 0x30c3, 0x30c4, 0x30c5, 0x30c6, 0x30c7, 0x30c8, 0x30c9, 0x30ca, 0x30cb, 0x30cc, 0x30cd, 0x30ce, 0x30cf, 0x30d0, 0x30d1, 0x30d2, 0x30d3, 0x30d4, 0x30d5, 0x30d6, 0x30d7, 0x30d8, 0x30d9, 0x30da, 0x30db, 0x30dc, 0x30dd, 0x30de, 0x30df, 0x30e0, 0x30e1, 0x30e2, 0x30e3, 0x30e4, 0x30e5, 0x30e6, 0x30e7, 0x30e8, 0x30e9, 0x30ea, 0x30eb, 0x30ec, 0x30ed, 0x30ee, 0x30ef, 0x30f0, 0x30f1, 0x30f2, 0x30f3, 0x30f4, 0x30f5, 0x30f6, 0x30f7, 0x30f8, 0x30f9, 0x30fa, 0x30fd, 0x30fe, 0x30ff, 0x31f0, 0x31f1, 0x31f2, 0x31f3, 0x31f4, 0x31f5, 0x31f6, 0x31f7, 0x31f8, 0x31f9, 0x31fa, 0x31fb, 0x31fc, 0x31fd, 0x31fe, 0x31ff, 0xff66, 0xff67, 0xff68, 0xff69, 0xff6a, 0xff6b, 0xff6c, 0xff6d, 0xff6e, 0xff6f, 0xff71, 0xff72, 0xff73, 0xff74, 0xff75, 0xff76, 0xff77, 0xff78, 0xff79, 0xff7a, 0xff7b, 0xff7c, 0xff7d, 0xff7e, 0xff7f, 0xff80, 0xff81, 0xff82, 0xff83, 0xff84, 0xff85, 0xff86, 0xff87, 0xff88, 0xff89, 0xff8a, 0xff8b, 0xff8c, 0xff8d, 0xff8e, 0xff8f, 0xff90, 0xff91, 0xff92, 0xff93, 0xff94, 0xff95, 0xff96, 0xff97, 0xff98, 0xff99, 0xff9a, 0xff9b, 0xff9c, 0xff9d, 0x1b164, 0x1b165, 0x1b166, 0x1b167})),
			out:      []byte("ｧｱｨｲｩｳｪｴｫｵｶｶﾞｷｷﾞｸｸﾞｹｹﾞｺｺﾞｻｻﾞｼｼﾞｽｽﾞｾｾﾞｿｿﾞﾀﾀﾞﾁﾁﾞｯﾂﾂﾞﾃﾃﾞﾄﾄﾞﾅﾆﾇﾈﾉﾊﾊﾞﾊﾟﾋﾋﾞﾋﾟﾌﾌﾞﾌﾟﾍﾍﾞﾍﾟﾎﾎﾞﾎﾟﾏﾐﾑﾒﾓｬﾔｭﾕｮﾖﾗﾘﾙﾚﾛﾜﾜｲｴｦﾝｳﾞｶｹﾜﾞｲﾞｴﾞｦﾞヽヾヿｸｼｽﾄﾇﾊﾋﾌﾍﾎﾑﾗﾘﾙﾚﾛｦｧｨｩｪｫｬｭｮｯｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜﾝｲｴｦﾝ"),
			formName: "narrow",
			err:      nil,
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
