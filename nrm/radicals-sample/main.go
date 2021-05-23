// +build run

package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/spiegel-im-spiegel/csvdata"
	"golang.org/x/text/unicode/norm"
)

//go:embed equivalent-unified-ideograph.csv
var kanjiList string

func readData() (map[rune]rune, error) {
	kanjiMap := map[rune]rune{}
	cr := csvdata.New(strings.NewReader(kanjiList), true).WithFieldsPerRecord(3)
	for {
		if err := cr.Next(); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		key, err := cr.ColumnInt64("radicals", 16)
		if err != nil {
			return nil, err
		}
		value, err := cr.ColumnInt64("normalize", 16)
		if err != nil {
			return nil, err
		}
		kanjiMap[rune(key)] = rune(value)
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
