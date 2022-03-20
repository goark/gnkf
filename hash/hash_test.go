package hash

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"errors"
	"fmt"
	"strings"
	"syscall"

	"testing"

	"github.com/goark/gnkf/ecode"
)

func TestAlgorithmList(t *testing.T) {
	res := "MD5|SHA-1|SHA-224|SHA-256|SHA-384|SHA-512|SHA-512/224|SHA-512/256"
	str := AlgorithmList("|")
	if str != res {
		t.Errorf("AlgorithmList() = \"%+v\", want \"%+v\".", str, res)
	}
}

func TestAlgorithm(t *testing.T) {
	testCases := []struct {
		name string
		alg  crypto.Hash
		err  error
	}{
		{name: "", alg: crypto.Hash(0), err: ecode.ErrInvalidHashAlg},
		{name: "foo", alg: crypto.Hash(0), err: ecode.ErrInvalidHashAlg},
		{name: "md5", alg: crypto.MD5, err: nil},
		{name: "SHA-1", alg: crypto.SHA1, err: nil},
		{name: "SHA-224", alg: crypto.SHA224, err: nil},
		{name: "SHA-256", alg: crypto.SHA256, err: nil},
		{name: "SHA-384", alg: crypto.SHA384, err: nil},
		{name: "SHA-512", alg: crypto.SHA512, err: nil},
		{name: "SHA-512/224", alg: crypto.SHA512_224, err: nil},
		{name: "SHA-512/256", alg: crypto.SHA512_256, err: nil},
	}
	for _, tc := range testCases {
		alg, err := Algorithm(tc.name)
		if !errors.Is(err, tc.err) {
			t.Errorf("Algorithm(%v) error = \"%+v\", want \"%+v\".", tc.name, err, tc.err)
		}
		if alg != tc.alg {
			t.Errorf("Algorithm(%v) = \"%+v\", want \"%+v\".", tc.name, alg.String(), tc.alg.String())
		}
	}
}

func TestCheck(t *testing.T) {
	testCases := []struct {
		algName string
		hashStr string
		res     bool
		err     error
	}{
		{algName: "", hashStr: "", res: false, err: ecode.ErrInvalidHashAlg},
		{algName: "foo", hashStr: "", res: false, err: ecode.ErrInvalidHashAlg},
		{algName: "md5", hashStr: "aa", res: false, err: ecode.ErrImproperlyHashFormat},
		{algName: "md5", hashStr: "d41d8cd98f00b204e9800998ecf8427a", res: false, err: nil},
		{algName: "md5", hashStr: "d41d8cd98f00b204e9800998ecf8427e", res: true, err: nil},                                                                                                     //see https://en.wikipedia.org/wiki/MD5
		{algName: "SHA-1", hashStr: "da39a3ee5e6b4b0d3255bfef95601890afd80709", res: true, err: nil},                                                                                           //see https://en.wikipedia.org/wiki/SHA-1
		{algName: "SHA-224", hashStr: "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", res: true, err: nil},                                                                         //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-256", hashStr: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", res: true, err: nil},                                                                 //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-384", hashStr: "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", res: true, err: nil},                                 //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-512", hashStr: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", res: true, err: nil}, //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-512/224", hashStr: "6ed0dd02806fa89e25de060c19d3ac86cabb87d6a0ddd05c333b84f4", res: true, err: nil},                                                                     //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-512/256", hashStr: "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a", res: true, err: nil},                                                             //see https://en.wikipedia.org/wiki/SHA-2
	}
	for _, tc := range testCases {
		alg, err := Algorithm(tc.algName)
		if err != nil {
			if !errors.Is(err, tc.err) {
				t.Errorf("Algorithm(%v) error = \"%+v\", want \"%+v\".", tc.algName, err, tc.err)
			}
		} else {
			res, err := Check(alg, strings.NewReader(""), tc.hashStr)
			if res != tc.res {
				t.Errorf("Check(%v) \"%v\", want \"%v\".", tc.algName, res, tc.res)
			}
			if !errors.Is(err, tc.err) {
				t.Errorf("Value(%v) error = \"%+v\", want \"%+v\".", tc.algName, err, tc.err)
			}
		}
	}
}

func TestCheckFile(t *testing.T) {
	testCases := []struct {
		algName string
		path    string
		hashStr string
		res     bool
		err     error
	}{
		{algName: "md5", path: "not-exist.dat", hashStr: "d41d8cd98f00b204e9800998ecf8427e", res: false, err: syscall.ENOENT},
		{algName: "md5", path: "testdata/null.dat", hashStr: "d41d8cd98f00b204e9800998ecf8427e", res: true, err: nil}, //see https://en.wikipedia.org/wiki/MD5
	}
	for _, tc := range testCases {
		alg, err := Algorithm(tc.algName)
		if err != nil {
			if !errors.Is(err, tc.err) {
				t.Errorf("Algorithm(%v) error = \"%+v\", want \"%+v\".", tc.algName, err, tc.err)
			}
		} else {
			res, err := CheckFile(alg, tc.path, tc.hashStr)
			if res != tc.res {
				t.Errorf("Check(%v) \"%v\", want \"%v\".", tc.algName, res, tc.res)
			}
			if !errors.Is(err, tc.err) {
				t.Errorf("Value(%v) error = \"%+v\", want \"%+v\".", tc.algName, err, tc.err)
			}
		}
	}
}

func TestValueFromBytes(t *testing.T) {
	testCases := []struct {
		algName string
		hashStr string
		err     error
	}{
		{algName: "", hashStr: "", err: ecode.ErrInvalidHashAlg},
		{algName: "foo", hashStr: "", err: ecode.ErrInvalidHashAlg},
		{algName: "md5", hashStr: "d41d8cd98f00b204e9800998ecf8427e", err: nil},                                                                                                     //see https://en.wikipedia.org/wiki/MD5
		{algName: "SHA-1", hashStr: "da39a3ee5e6b4b0d3255bfef95601890afd80709", err: nil},                                                                                           //see https://en.wikipedia.org/wiki/SHA-1
		{algName: "SHA-224", hashStr: "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f", err: nil},                                                                         //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-256", hashStr: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", err: nil},                                                                 //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-384", hashStr: "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", err: nil},                                 //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-512", hashStr: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", err: nil}, //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-512/224", hashStr: "6ed0dd02806fa89e25de060c19d3ac86cabb87d6a0ddd05c333b84f4", err: nil},                                                                     //see https://en.wikipedia.org/wiki/SHA-2
		{algName: "SHA-512/256", hashStr: "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a", err: nil},                                                             //see https://en.wikipedia.org/wiki/SHA-2
	}
	for _, tc := range testCases {
		alg, err := Algorithm(tc.algName)
		if !errors.Is(err, tc.err) {
			t.Errorf("Algorithm(%v) error = \"%+v\", want \"%+v\".", tc.algName, err, tc.err)
		} else {

			if v, err := ValueFromBytes(alg, []byte("")); !errors.Is(err, tc.err) {
				t.Errorf("Value(%v) error = \"%+v\", want \"%+v\".", tc.algName, err, tc.err)
			} else {
				str := fmt.Sprintf("%x", v)
				if str != tc.hashStr {
					t.Errorf("Value(%v) \"%+v\", want \"%+v\".", tc.algName, str, tc.hashStr)
				}
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
