package b64

import (
	"encoding/base64"
	"io"

	"github.com/spiegel-im-spiegel/errs"
)

//Encode outputs base64 encoding string from raw data.
func Encode(forURL, noPadding bool, r io.Reader, w io.Writer) error {
	wc := base64.NewEncoder(encoder(forURL, noPadding), w)
	defer wc.Close()
	if _, err := io.Copy(wc, r); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

//Decode outputs raw data from base64 encoding string.
func Decode(forURL, noPadding bool, r io.Reader, w io.Writer) error {
	if _, err := io.Copy(w, base64.NewDecoder(encoder(forURL, noPadding), r)); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func encoder(forURL, noPadding bool) *base64.Encoding {
	var enc *base64.Encoding
	if forURL {
		enc = base64.URLEncoding
	} else {
		enc = base64.StdEncoding
	}
	if noPadding {
		enc = enc.WithPadding(base64.NoPadding)
	}
	return enc
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
