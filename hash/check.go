package hash

import (
	"crypto"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
)

// Check function returns true if computed hash value is match.
func Check(alg crypto.Hash, r io.Reader, hashStr string) (bool, error) {
	v, err := Value(alg, r)
	if err != nil {
		return false, errs.Wrap(ecode.ErrInvalidHashAlg, errs.WithContext("algorithm", AlgoString(alg)), errs.WithContext("hash", hashStr))
	}
	str := fmt.Sprintf("%x", v)
	if len(str) != len(hashStr) {
		return false, errs.Wrap(ecode.ErrImproperlyHashFormat, errs.WithContext("algorithm", AlgoString(alg)), errs.WithContext("hash", hashStr))
	}
	if !strings.EqualFold(str, hashStr) {
		return false, nil
	}
	return true, nil
}

// Check function returns true if computed hash value is match.
func CheckFile(alg crypto.Hash, path string, hashStr string) (res bool, err error) {
	file, ferr := os.Open(filepath.Clean(path))
	if ferr != nil {
		err = errs.Wrap(ferr, errs.WithContext("algorithm", AlgoString(alg)), errs.WithContext("path", path), errs.WithContext("hash", hashStr))
		return
	}
	defer func() {
		err = errs.Join(err, file.Close())
	}()
	res, err = Check(alg, file, hashStr)
	return
}

/* Copyright 2021-2026 Spiegel
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
