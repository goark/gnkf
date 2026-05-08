package facade

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
	"github.com/goark/gnkf/guess"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

// newGuessCmd returns cobra.Command instance for show sub-command
func newGuessCmd(ui *rwi.RWI) *cobra.Command {
	guessCmd := &cobra.Command{
		Use:     "guess",
		Aliases: []string{"g"},
		Short:   "Guess character encoding of the text",
		Long:    "Guess character encoding of the text",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			//Options
			path, ferr := cmd.Flags().GetString("file")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --file option", errs.WithCause(ferr)))
				return
			}
			flagAll, ferr := cmd.Flags().GetBool("all")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --all option", errs.WithCause(ferr)))
				return
			}

			//Input stream
			r := ui.Reader()
			if len(path) > 0 {
				file, ferr := os.Open(filepath.Clean(path))
				if ferr != nil {
					err = debugPrint(ui, errs.Wrap(ferr, errs.WithContext("file", path)))
					return
				}
				defer func() {
					err = errs.Join(err, file.Close())
				}()
				r = file
			}

			//Run command
			ss, gerr := guess.Encoding(r)
			if gerr != nil {
				err = debugPrint(ui, errs.Wrap(gerr, errs.WithContext("file", path)))
				return
			}
			if len(ss) == 0 {
				err = debugPrint(ui, errs.Wrap(ecode.ErrNoData, errs.WithContext("file", path)))
				return
			}
			if flagAll {
				err = ui.Outputln(strings.Join(ss, "\n"))
			} else {
				err = ui.Outputln(ss[0])
			}
			err = debugPrint(ui, errs.Wrap(err))
			return
		},
	}
	guessCmd.Flags().StringP("file", "f", "", "path of input text file")
	_ = guessCmd.MarkFlagFilename("file")
	guessCmd.Flags().BoolP("all", "", false, "print all guesses")

	return guessCmd
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
