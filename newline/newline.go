package newline

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/errs"
)

//Convert function convert newline in the text stream.
func Convert(formName string, writer io.Writer, txt io.Reader) error {
	f, err := FormOf(formName)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("formName", formName))
	}

	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(txt); err != nil {
		return errs.Wrap(err)
	}
	if _, err := NewReplacer(f).WriteString(writer, buf.String()); err != nil {
		return errs.Wrap(err)
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
