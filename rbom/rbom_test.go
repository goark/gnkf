package rbom_test

import (
	"bytes"
	"testing"

	"github.com/goark/gnkf/rbom"
)

func TestRemoveBomNil(t *testing.T) {
	testCases := []struct {
		inp  []byte
		outp []byte
	}{
		{inp: nil, outp: []byte{}},
	}
	for _, tc := range testCases {
		if b, err := rbom.RemoveBomBytes(tc.inp); err != nil {
			t.Errorf("RemoveBom() = \"%v\", want nil.", err)
		} else if !bytes.Equal(b, tc.outp) {
			t.Errorf("RemoveBom() = %v, want %v.", b, tc.outp)
		}
	}
}

func TestRemoveBom(t *testing.T) {
	testCases := []struct {
		inp  []byte
		outp []byte
	}{
		{inp: nil, outp: []byte{}},
		{inp: []byte("Hello"), outp: []byte("Hello")},
		{inp: []byte{0xef, 0xbb, 0xbf, 0x48, 0x65, 0x6c, 0x6c, 0xef, 0xbb, 0xbf, 0x6f}, outp: []byte("Hello")},
	}
	for _, tc := range testCases {
		if b, err := rbom.RemoveBom(bytes.NewReader(tc.inp)); err != nil {
			t.Errorf("RemoveBom() = \"%v\", want nil.", err)
		} else if !bytes.Equal(b, tc.outp) {
			t.Errorf("RemoveBom() = %v, want %v.", b, tc.outp)
		}
	}
}

func TestRemoveBomString(t *testing.T) {
	testCases := []struct {
		inp  string
		outp string
	}{
		{inp: "", outp: ""},
		{inp: "Hello", outp: "Hello"},
		{inp: string([]byte{0xef, 0xbb, 0xbf, 0x48, 0x65, 0x6c, 0x6c, 0x6f}), outp: "Hello"},
	}
	for _, tc := range testCases {
		s := rbom.RemoveBomString(tc.inp)
		if s != tc.outp {
			t.Errorf("RemoveBom() = %v, want %v.", s, tc.outp)
		}
	}
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
