package hash

import (
	"crypto"
	"io"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
)

//Value returns hash value string from io.Reader
func Value(alg crypto.Hash, r io.Reader) ([]byte, error) {
	if !alg.Available() {
		return nil, errs.Wrap(ecode.ErrInvalidHashAlg, errs.WithContext("algorithm", AlgoString(alg)))
	}
	h := alg.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("algorithm", AlgoString(alg)))
	}
	return h.Sum(nil), nil
}

//ValueFromBytes returns hash value string from []byte
func ValueFromBytes(alg crypto.Hash, b []byte) ([]byte, error) {
	if !alg.Available() {
		return nil, errs.Wrap(ecode.ErrInvalidHashAlg, errs.WithContext("algorithm", AlgoString(alg)))
	}
	return alg.New().Sum(b), nil
}

/* Copyright 2021 Spiegel
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
