package hash

import (
	"bufio"
	"crypto"
	"io"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
)

type Checker interface {
	Path() string
	Err() error
	Check() error
}

//NewCheckers returns list of Checker instances from io.Reader.
func NewCheckers(r io.Reader, alg crypto.Hash) ([]Checker, error) {
	scanner := bufio.NewScanner(r)
	chks := []Checker{}
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) < 2 {
			return chks, errs.Wrap(ecode.ErrInvalidChekerFormat)
		}
		if line[len(line)-1] != "-" {
			chks = append(chks, newChecker(alg, line[len(line)-1], line[0]))
		}
	}
	if err := scanner.Err(); err != nil {
		return chks, errs.Wrap(err)
	}
	return chks, nil
}

// checker is hash checker class.
type checker struct {
	alg     crypto.Hash
	path    string
	hashStr string
	err     error
}

func newChecker(alg crypto.Hash, path string, hashStr string) Checker {
	return &checker{alg: alg, path: path, hashStr: hashStr, err: nil}
}

//Path method returns path element in checker.
func (c *checker) Path() string { return c.path }

//Err method returns error element in checker.
func (c *checker) Err() error { return c.err }

//Check method checks hash code with checker info.
func (c *checker) Check() error {
	if c == nil {
		return nil
	}
	if ok, err := CheckFile(c.alg, c.path, c.hashStr); err != nil {
		c.err = errs.Wrap(err, errs.WithContext("alg", c.alg.String()), errs.WithContext("path", c.path), errs.WithContext("hashStr", c.hashStr))
	} else if !ok {
		c.err = errs.Wrap(ecode.ErrUnmatchHashString, errs.WithContext("alg", c.alg.String()), errs.WithContext("path", c.path), errs.WithContext("hashStr", c.hashStr))
	} else {
		c.err = nil
	}
	return c.err
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
