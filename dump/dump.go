package dump

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
)

// Octet output io.Writer hex-dump of byte stream.
func Octet(w io.Writer, r io.Reader) (err error) {
	sep := ""
	inp := bufio.NewReader(r)
	for {
		b, ierr := inp.ReadByte()
		if ierr != nil {
			if errs.Is(ierr, io.EOF) {
				break
			}
			err = errs.Wrap(ierr)
			return
		}
		_, err = fmt.Fprintf(w, "%s0x%02x", sep, b)
		if err != nil {
			err = errs.Wrap(err)
			return
		}
		sep = ", "
	}
	return
}

// OctetString output hex-dump string.
func OctetString(r io.Reader) string {
	buf := &bytes.Buffer{}
	if err := Octet(buf, r); err != nil {
		return ""
	}
	return buf.String()
}

// UnicodePoint output io.Writer hex-dump of Unicode code point (input text is UTF-8 only).
func UnicodePoint(w io.Writer, r io.Reader) (err error) {
	buf := &bytes.Buffer{}
	if _, err = buf.ReadFrom(r); err != nil {
		err = errs.Wrap(err)
		return
	}
	if !utf8.Valid(buf.Bytes()) {
		err = errs.Wrap(ecode.ErrInvalidUTF8Text)
		return
	}

	sep := ""
	for _, rn := range buf.String() {
		if (rn & 0x7fff0000) == 0 {
			_, err = fmt.Fprintf(w, "%s0x%04x", sep, rn)
			if err != nil {
				err = errs.Wrap(err)
				return
			}
		} else {
			_, err = fmt.Fprintf(w, "%s0x%08x", sep, rn)
			if err != nil {
				err = errs.Wrap(err)
				return
			}
		}
		sep = ", "
	}
	return
}

// UnicodePointString output hex-dump string of Unicode code point (input text is UTF-8 only).
func UnicodePointString(r io.Reader) string {
	buf := &bytes.Buffer{}
	if err := UnicodePoint(buf, r); err != nil {
		return ""
	}
	return buf.String()
}

/* Copyright 2020-2026 Spiegel
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
