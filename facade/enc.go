package facade

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/goark/errs"
	"github.com/goark/gnkf/enc"
	"github.com/goark/gnkf/guess"
	"github.com/goark/gnkf/rbom"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/unicode"
)

var descriptionEnc = `Convert character encoding of the text.
 Using MIME and IANA name as the character encoding name.
 Refer: http://www.iana.org/assignments/character-sets/character-sets.xhtml`

// newEncCmd returns cobra.Command instance for show sub-command
func newEncCmd(ui *rwi.RWI) *cobra.Command {
	encCmd := &cobra.Command{
		Use:     "enc",
		Aliases: []string{"encoding", "e"},
		Short:   "Convert character encoding of the text",
		Long:    descriptionEnc,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			//Options
			inp, ferr := cmd.Flags().GetString("file")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --file option", errs.WithCause(ferr)))
				return
			}
			out, ferr := cmd.Flags().GetString("output")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --output option", errs.WithCause(ferr)))
				return
			}
			from, ferr := cmd.Flags().GetString("src-encoding")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --src-encoding option", errs.WithCause(ferr)))
				return
			}
			to, ferr := cmd.Flags().GetString("dst-encoding")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --dst-encoding option", errs.WithCause(ferr)))
				return
			}
			flagGuess, ferr := cmd.Flags().GetBool("guess")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --guess option", errs.WithCause(ferr)))
				return
			}
			rbFlag, ferr := cmd.Flags().GetBool("remove-bom")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --remove-bom option", errs.WithCause(ferr)))
				return
			}

			//Input stream
			r := ui.Reader()
			if len(inp) > 0 {
				file, ferr := os.Open(filepath.Clean(inp))
				if ferr != nil {
					return debugPrint(ui, errs.Wrap(ferr, errs.WithContext("file", inp)))
				}
				defer func() {
					err = errs.Join(err, file.Close())
				}()
				r = file
			}
			if flagGuess {
				dup := &bytes.Buffer{}
				ss, eerr := guess.Encoding(io.TeeReader(r, dup))
				if eerr != nil {
					err = debugPrint(ui, errs.Wrap(eerr, errs.WithContext("file", inp)))
					return
				}
				if len(ss) > 0 {
					from = ss[0]
				}
				r = dup
			}

			//Output stream
			w := ui.Writer()
			if len(out) > 0 {
				file, ferr := os.Create(filepath.Clean(out))
				if ferr != nil {
					err = debugPrint(ui, errs.Wrap(ferr, errs.WithContext("output", out)))
					return
				}
				defer func() {
					err = errs.Join(err, file.Close())
				}()
				w = file
			}

			//Remove BOM
			if rbFlag {
				e, eerr := enc.Encoding(from)
				if eerr != nil {
					err = debugPrint(ui, errs.Wrap(eerr, errs.WithContext("file", inp), errs.WithContext("output", out)))
					return
				}
				if e == unicode.UTF8 {
					b, rerr := rbom.RemoveBom(r)
					if rerr != nil {
						err = debugPrint(ui, errs.Wrap(rerr, errs.WithContext("file", inp), errs.WithContext("output", out)))
						return
					}
					r = bytes.NewReader(b)
				}
			}

			//Run command
			if eerr := enc.Convert(to, w, from, r); eerr != nil {
				err = debugPrint(ui, errs.Wrap(eerr, errs.WithContext("file", inp), errs.WithContext("output", out)))
				return
			}
			return
		},
	}
	encCmd.Flags().StringP("file", "f", "", "path of input text file")
	_ = encCmd.MarkFlagFilename("file")
	encCmd.Flags().StringP("output", "o", "", "path of output file")
	_ = encCmd.MarkFlagFilename("output")
	encCmd.Flags().StringP("src-encoding", "s", "utf-8", "character encoding name of source text")
	encCmd.Flags().StringP("dst-encoding", "d", "utf-8", "character encoding name of output text")
	encCmd.Flags().BoolP("guess", "g", false, "guess character encoding of source text")
	encCmd.Flags().BoolP("remove-bom", "b", false, "remove BOM character in source text (UTF-8 only)")

	return encCmd
}

/* Copyright 2020-2026 Spiegel
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
