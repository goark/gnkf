package facade

import (
	"bytes"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/enc"
	"github.com/spiegel-im-spiegel/gnkf/guess"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var descriptionEnc = `Convert character encoding of the text.
 Using MIME and IANA name as the character encoding name.
 Refer: http://www.iana.org/assignments/character-sets/character-sets.xhtml`

//newEncCmd returns cobra.Command instance for show sub-command
func newEncCmd(ui *rwi.RWI) *cobra.Command {
	encCmd := &cobra.Command{
		Use:     "enc",
		Aliases: []string{"encoding", "e"},
		Short:   "Convert character encoding of the text",
		Long:    descriptionEnc,
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			inp, err := cmd.Flags().GetString("file")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --file option", errs.WithCause(err)))
			}
			out, err := cmd.Flags().GetString("output")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --output option", errs.WithCause(err)))
			}
			from, err := cmd.Flags().GetString("src-encoding")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --src-encoding option", errs.WithCause(err)))
			}
			to, err := cmd.Flags().GetString("dst-encoding")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --dst-encoding option", errs.WithCause(err)))
			}
			flagGuess, err := cmd.Flags().GetBool("guess")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --guess option", errs.WithCause(err)))
			}

			//Input stream
			r := ui.Reader()
			if len(inp) > 0 {
				file, err := os.Open(inp)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, errs.WithContext("file", inp)))
				}
				defer file.Close()
				r = file
			}
			if flagGuess {
				dup := &bytes.Buffer{}
				ss, err := guess.Encoding(io.TeeReader(r, dup))
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, errs.WithContext("file", inp)))
				}
				if len(ss) > 0 {
					from = ss[0]
				}
				r = dup
			}

			//Output stream
			w := ui.Writer()
			if len(out) > 0 {
				file, err := os.Create(out)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, errs.WithContext("output", out)))
				}
				defer file.Close()
				w = file
			}

			//Run command
			if err := enc.Convert(to, w, from, r); err != nil {
				return debugPrint(ui, errs.Wrap(err, errs.WithContext("file", inp), errs.WithContext("output", out)))
			}
			return nil
		},
	}
	encCmd.Flags().StringP("file", "f", "", "path of input text file")
	encCmd.Flags().StringP("output", "o", "", "path of output file")
	encCmd.Flags().StringP("src-encoding", "s", "utf-8", "character encoding name of source text")
	encCmd.Flags().StringP("dst-encoding", "d", "utf-8", "character encoding name of output text")
	encCmd.Flags().BoolP("guess", "g", false, "guess character encoding of source text")

	return encCmd
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
