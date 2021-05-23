package ecode

import "errors"

var (
	ErrNullPointer          = errors.New("null reference instance")
	ErrNoCommand            = errors.New("no command")
	ErrNoData               = errors.New("no data")
	ErrCannotDetect         = errors.New("cannot detect character encoding")
	ErrInvalidUTF8Text      = errors.New("invalid UTF-8 text")
	ErrNotSuppotEncoding    = errors.New("not support IANA encoding name")
	ErrInvalidEncoding      = errors.New("text is invalid encoding")
	ErrInvalidNormForm      = errors.New("invalid Unicode normalization form")
	ErrInvalidNewlineForm   = errors.New("invalid newline form")
	ErrInvalidWidthForm     = errors.New("invalid width form")
	ErrInvalidKanaForm      = errors.New("invalid kana form")
	ErrInvalidHashAlg       = errors.New("not support hash algorithm")
	ErrImproperlyHashFormat = errors.New("improperly formatted hash string")
	ErrUnmatchHashString    = errors.New("hash value did NOT match")
	ErrInvalidChekerFormat  = errors.New("invalid checker format")
)

/* Copyright 2020-2021 Spiegel
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
