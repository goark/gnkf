package facade

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
	"github.com/spiegel-im-spiegel/gnkf/guess"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newGuessCmd returns cobra.Command instance for show sub-command
func newGuessCmd(ui *rwi.RWI) *cobra.Command {
	guessCmd := &cobra.Command{
		Use:     "guess",
		Aliases: []string{"g"},
		Short:   "Guess character encoding of the text",
		Long:    "Guess character encoding of the text",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			path, err := cmd.Flags().GetString("file")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --file option"))
			}
			flagAll, err := cmd.Flags().GetBool("all")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --all option"))
			}

			//Input stream
			r := ui.Reader()
			if len(path) > 0 {
				file, err := os.Open(path)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("file", path)))
				}
				defer file.Close()
				r = file
			}

			//Run command
			ss, err := guess.Encoding(r)
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("file", path)))
			}
			if len(ss) == 0 {
				return debugPrint(ui, errs.WrapWithCause(ecode.ErrNoData, nil, errs.WithContext("file", path)))
			}
			if flagAll {
				err = ui.Outputln(strings.Join(ss, "\n"))
			} else {
				err = ui.Outputln(ss[0])
			}
			return debugPrint(ui, errs.Wrap(err, ""))
		},
	}
	guessCmd.Flags().StringP("file", "f", "", "path of input text file")
	guessCmd.Flags().BoolP("all", "", false, "print all guesses")

	return guessCmd
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
