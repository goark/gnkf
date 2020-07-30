package ecode

import "fmt"

//ECode is error codes for books-data
type ECode int

const (
	ErrNullPointer ECode = iota + 1
	ErrNoCommand
	ErrNoData
	ErrCannotDetect
	ErrInvalidUTF8Text
	ErrNotSuppotEncoding
	ErrInvalidEncoding
	ErrInvalidNormForm
	ErrInvalidNewlineForm
	ErrInvalidWidthForm
)

var errMessages = map[ECode]string{
	ErrNullPointer:        "Null reference instance",
	ErrNoCommand:          "No command",
	ErrNoData:             "No data",
	ErrCannotDetect:       "Cannot detect character encoding",
	ErrInvalidUTF8Text:    "Invalid UTF-8 text",
	ErrNotSuppotEncoding:  "Not Support IANA encoding name",
	ErrInvalidEncoding:    "Text is invalid encoding",
	ErrInvalidNormForm:    "Invalid Unicode normalization form",
	ErrInvalidNewlineForm: "Invalid newline form",
	ErrInvalidWidthForm:   "Invalid width form",
}

func (e ECode) Error() string {
	if s, ok := errMessages[e]; ok {
		return s
	}
	return fmt.Sprintf("unknown error (%d)", int(e))
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
