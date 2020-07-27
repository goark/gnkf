package dump

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

//Octet output io.Writer hex-dump of byte stream
func Octet(w io.Writer, r io.Reader) error {
	sep := ""
	inp := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	for {
		b, err := inp.ReadByte()
		if err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return errs.Wrap(err, "")
		}
		fmt.Fprintf(out, "%s0x%02x", sep, b)
		sep = ", "
	}
	fmt.Fprintln(out)
	err := out.Flush()
	if err != nil {
		return errs.Wrap(err, "")
	}
	return nil
}

//UnicodePoint output io.Writer hex-dump of Unicode code point (input text is UTF-8 only)
func UnicodePoint(w io.Writer, r io.Reader) error {
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(r); err != nil {
		return errs.Wrap(err, "")
	}
	if !utf8.Valid(buf.Bytes()) {
		return errs.Wrap(ecode.ErrInvalidUTF8Text, "")
	}

	sep := ""
	out := bufio.NewWriter(w)
	for _, rn := range buf.String() {
		if (rn & 0x7fff0000) == 0 {
			fmt.Fprintf(out, "%s0x%04x", sep, rn)
		} else {
			fmt.Fprintf(out, "%s0x%08x", sep, rn)
		}
		sep = ", "
	}
	fmt.Fprintln(out)
	err := out.Flush()
	if err != nil {
		return errs.Wrap(err, "")
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
