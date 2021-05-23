package hash

import (
	"bytes"
	"crypto"
	"errors"
	"syscall"
	"testing"

	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

const (
	checkerFile0 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85a`
	checkerFile1 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85a  not-exist.dat`
	checkerFile2 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855  testdata/null.dat`
	checkerFile3 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85a  testdata/null.dat`
	checkerFile4 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85aa  testdata/null.dat`
)

func TestCheckerFile(t *testing.T) {
	testCases := []struct {
		alg crypto.Hash
		inp string
		err error
	}{
		{alg: crypto.SHA256, inp: checkerFile0, err: ecode.ErrInvalidChekerFormat},
		{alg: crypto.SHA256, inp: checkerFile1, err: syscall.ENOENT},
		{alg: crypto.SHA256, inp: checkerFile2, err: nil},
		{alg: crypto.SHA256, inp: checkerFile3, err: ecode.ErrUnmatchHashString},
		{alg: crypto.SHA256, inp: checkerFile4, err: ecode.ErrImproperlyHashFormat},
	}
	for _, tc := range testCases {
		checkers, err := NewCheckers(bytes.NewReader([]byte(tc.inp)), tc.alg)
		if err != nil {
			if !errors.Is(err, tc.err) {
				t.Errorf("NewCheckers() error = \"%+v\", want \"%+v\".", err, tc.err)
			}
		} else if len(checkers) < 1 {
			t.Errorf("count ofNewCheckers() are %d.", len(checkers))
		} else {
			err := checkers[0].Check()
			if !errors.Is(err, tc.err) {
				t.Errorf("Checkers.Check() error = \"%+v\", want \"%+v\".", err, tc.err)
			}
		}
	}
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
