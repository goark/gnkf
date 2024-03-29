package facade

import (
	"fmt"
	"os"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gnkf/newline"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

//newNormCmd returns cobra.Command instance for show sub-command
func newNwlnCmd(ui *rwi.RWI) *cobra.Command {
	nwlnCmd := &cobra.Command{
		Use:     "newline",
		Aliases: []string{"nwln", "nl"},
		Short:   "Convert newline form in the text",
		Long:    "Convert newline form in the text.",
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
			form, err := cmd.Flags().GetString("newline-form")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --newline-form option", errs.WithCause(err)))
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
			if err := newline.Convert(form, w, r); err != nil {
				return debugPrint(ui, errs.Wrap(err, errs.WithContext("file", inp), errs.WithContext("output", out)))
			}
			return nil
		},
	}
	nwlnCmd.Flags().StringP("file", "f", "", "path of input text file")
	_ = nwlnCmd.MarkFlagFilename("file")
	nwlnCmd.Flags().StringP("output", "o", "", "path of output file")
	_ = nwlnCmd.MarkFlagFilename("output")
	nwlnCmd.Flags().StringP("newline-form", "n", "lf", fmt.Sprintf("newline form: [%s]", strings.Join(newline.FormList(), "|")))
	_ = nwlnCmd.RegisterFlagCompletionFunc("newline-form", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return newline.FormList(), cobra.ShellCompDirectiveNoFileComp
	})

	return nwlnCmd
}

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
