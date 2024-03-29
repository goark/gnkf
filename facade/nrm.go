package facade

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gnkf/nrm"
	"github.com/goark/gnkf/rbom"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

//newNormCmd returns cobra.Command instance for show sub-command
func newNormCmd(ui *rwi.RWI) *cobra.Command {
	normCmd := &cobra.Command{
		Use:     "norm",
		Aliases: []string{"normalize", "nrm", "nm"},
		Short:   "Unicode normalization of the text",
		Long:    "Unicode normalization of the text (UTF-8 encoding only).",
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
			form, err := cmd.Flags().GetString("norm-form")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --norm-form option", errs.WithCause(err)))
			}
			krFlag, err := cmd.Flags().GetBool("kangxi-radicals")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --kangxi-radicals option", errs.WithCause(err)))
			}
			rbFlag, err := cmd.Flags().GetBool("remove-bom")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --remove-bom option", errs.WithCause(err)))
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

			//Remove BOM
			if rbFlag {
				b, err := rbom.RemoveBom(r)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, errs.WithContext("file", inp), errs.WithContext("output", out)))
				}
				r = bytes.NewReader(b)
			}

			//Run command
			if err := nrm.Normalize(form, w, r, krFlag); err != nil {
				return debugPrint(ui, errs.Wrap(err, errs.WithContext("file", inp), errs.WithContext("output", out)))
			}
			return nil
		},
	}
	normCmd.Flags().StringP("file", "f", "", "path of input text file")
	_ = normCmd.MarkFlagFilename("file")
	normCmd.Flags().StringP("output", "o", "", "path of output file")
	_ = normCmd.MarkFlagFilename("output")
	normCmd.Flags().StringP("norm-form", "n", "nfc", fmt.Sprintf("Unicode normalization form: [%s]", strings.Join(nrm.FormList(), "|")))
	_ = normCmd.RegisterFlagCompletionFunc("norm-form", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nrm.FormList(), cobra.ShellCompDirectiveDefault
	})
	normCmd.Flags().BoolP("kangxi-radicals", "k", false, "normalize kangxi radicals only (with nfkc or nfkd form)")
	normCmd.Flags().BoolP("remove-bom", "b", false, "remove BOM character")

	return normCmd
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
