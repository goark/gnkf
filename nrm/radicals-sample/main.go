// +build run

package main

import (
	_ "embed"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/text/unicode/norm"
)

//go:embed equivalent-unified-ideograph.csv
var kanjiList string

var (
	ErrNoData        = errors.New("no data")
	ErrInvalidRecord = errors.New("invalid record")
)

//Reader is class of CSV reader
type Reader struct {
	reader        *csv.Reader
	cols          int
	headerFlag    bool
	headerStrings []string
}

//New function creates a new Reader instance.
func New(r io.Reader, cols int, headerFlag bool) *Reader {
	cr := csv.NewReader(r)
	cr.Comma = ','
	cr.LazyQuotes = true       // a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	cr.TrimLeadingSpace = true // leading
	return &Reader{reader: cr, cols: cols, headerFlag: headerFlag}
}

//readRecord method returns a new record.
func (r *Reader) readRecord() ([]string, error) {
	elms, err := r.reader.Read()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, err
		}
		return nil, ErrInvalidRecord
	}
	if len(elms) < r.cols {
		return nil, ErrInvalidRecord
	}
	return elms, nil
}

//Header method returns header strings.
func (r *Reader) Header() ([]string, error) {
	var err error
	if r.headerFlag {
		r.headerFlag = false
		r.headerStrings, err = r.readRecord()
	}
	return r.headerStrings, err
}

//Next method returns a next record.
func (r *Reader) Next() ([]string, error) {
	if r.headerFlag {
		if _, err := r.Header(); err != nil {
			return nil, err
		}
	}
	elms, err := r.readRecord()
	return elms, err
}

func readData() (map[rune]rune, error) {
	kanjiMap := map[rune]rune{}
	cr := New(strings.NewReader(kanjiList), 3, true)
	for {
		elms, err := cr.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		key, err := strconv.ParseUint(strings.TrimSpace(elms[0]), 16, 32)
		if err != nil {
			return nil, err
		}
		value, err := strconv.ParseUint(strings.TrimSpace(elms[1]), 16, 32)
		if err != nil {
			return nil, err
		}
		if len(elms) > 2 {
			kanjiMap[rune(key)] = rune(value)
		}
	}
	return kanjiMap, nil
}

func main() {
	kanjiMap, err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("var kangxiRadicals = unicode.SpecialCase{")
	for kr := rune(0x2e80); kr <= 0x2fdf; kr++ {
		ki, ok := kanjiMap[kr]
		if !ok {
			kis := []rune(norm.NFKC.String(string([]rune{kr})))
			ki = kis[0]
		}
		if kr != ki {
			fmt.Printf("\tunicode.CaseRange{Lo: %#[1]x, Hi: %#[1]x, Delta: [unicode.MaxCase]rune{%#[2]x - %#[1]x, 0, 0}}, // %#[1]U -> %#[2]U\n", kr, ki)
		}
	}
	fmt.Println("}")
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
