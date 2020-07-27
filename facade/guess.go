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
		Short:   "Guess character encoding of text",
		Long:    "Guess character encoding of text",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			path, err := cmd.Flags().GetString("src")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --src option"))
			}
			flagAll, err := cmd.Flags().GetBool("all")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --all option"))
			}

			//Run command
			r := ui.Reader()
			if len(path) > 0 {
				file, err := os.Open(path)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("path", path)))
				}
				defer file.Close()
				r = file
			}
			ss, err := guess.Encoding(r)
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("path", path)))
			}
			if len(ss) == 0 {
				return debugPrint(ui, errs.Wrap(ecode.ErrNoData, "", errs.WithContext("path", path)))
			}
			if flagAll {
				err = ui.Outputln(strings.Join(ss, "\n"))
			} else {
				err = ui.Outputln(ss[0])
			}
			return debugPrint(ui, errs.Wrap(err, ""))
		},
	}
	guessCmd.Flags().StringP("src", "s", "", "path of source text")
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
