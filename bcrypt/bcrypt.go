package bcrypt

import (
	"github.com/goark/errs"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = bcrypt.MinCost
	MaxCost     int = bcrypt.MaxCost
	DefaultCost int = bcrypt.DefaultCost
)

//Hash function returns hashed string by BCrypt algorithm.
func Hash(s string, cost int) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), cost)
	if err != nil {
		return "", errs.Wrap(err)
	}
	return string(b), nil
}

//Compare function compares a bcrypt hashed string with its possible plaintext equivalent. Returns nil on success, or an error on failure.
func Compare(h, s string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(h), []byte(s)); err != nil {
		return errs.Wrap(err)
	}
	return nil
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
