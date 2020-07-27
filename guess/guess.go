package guess

import (
	"bytes"
	"io"
	"sort"

	"github.com/saintfish/chardet"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

//Encoding detects guesses of character encoding name from byte stream
func Encoding(txt io.Reader) ([]string, error) {
	if txt == nil {
		return nil, errs.WrapWithCause(ecode.ErrNullPointer, nil)
	}
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(txt); err != nil {
		return nil, errs.Wrap(err, "")
	}
	return EncodingBytes(buf.Bytes())
}

//EncodingBytes detects guesses of character encoding name from byte array
func EncodingBytes(b []byte) ([]string, error) {
	all, err := chardet.NewTextDetector().DetectAll(b)
	if err != nil {
		return nil, errs.WrapWithCause(ecode.ErrCannotDetect, err)
	}
	sort.SliceStable(all, func(i, j int) bool {
		if all[i].Confidence != all[j].Confidence {
			return all[i].Confidence > all[j].Confidence
		}
		return all[i].Charset < all[j].Charset
	})
	ss := []string{}
	for _, r := range all {
		ss = append(ss, r.Charset)
	}
	return ss, nil
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
