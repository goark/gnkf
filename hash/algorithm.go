package hash

import (
	"crypto"
	"strings"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

var algOrder = []crypto.Hash{
	crypto.MD5,        //require "crypto/md5" package
	crypto.SHA1,       //require "crypto/sha1" package
	crypto.SHA224,     //require "crypto/sha256" package
	crypto.SHA256,     //require "crypto/sha256" package
	crypto.SHA384,     //require "crypto/sha512" package
	crypto.SHA512,     //require "crypto/sha512" package
	crypto.SHA512_224, //require "crypto/sha512" package
	crypto.SHA512_256, //require "crypto/sha512" package
}

//AlgorithmList returns string of hash functions list.
func AlgorithmList(sep string) string {
	ss := []string{}
	for _, alg := range algOrder {
		if s := AlgoString(alg); len(s) > 0 {
			ss = append(ss, s)
		}
	}
	return strings.Join(ss, sep)
}

//Algorithm returns crypto.Hash from string.
func Algorithm(s string) (crypto.Hash, error) {
	if len(s) == 0 {
		return crypto.Hash(0), errs.Wrap(ecode.ErrInvalidHashAlg, errs.WithContext("algorithm", s))
	}
	for _, alg := range algOrder {
		if strings.EqualFold(AlgoString(alg), s) {
			return alg, nil
		}
	}
	return crypto.Hash(0), errs.Wrap(ecode.ErrInvalidHashAlg, errs.WithContext("algorithm", s))
}

//AlgoString returns string of hash algorithm.
func AlgoString(alg crypto.Hash) string {
	if alg.Available() {
		return alg.String()
	}
	return ""
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
