package facade

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/nrm"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newNormCmd returns cobra.Command instance for show sub-command
func newNormCmd(ui *rwi.RWI) *cobra.Command {
	normCmd := &cobra.Command{
		Use:     "norm",
		Aliases: []string{"normalize", "nrm"},
		Short:   "Unicode normalization",
		Long:    "Unicode normalization (UTF-8 text only)",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			inp, err := cmd.Flags().GetString("file")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --file option"))
			}
			out, err := cmd.Flags().GetString("output")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --output option"))
			}
			form, err := cmd.Flags().GetString("norm-form")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --src-encoding option"))
			}
			krFlag, err := cmd.Flags().GetBool("kangxi-radicals")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --kangxi-radicals option"))
			}

			//Input stream
			r := ui.Reader()
			if len(inp) > 0 {
				file, err := os.Open(inp)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("file", inp)))
				}
				defer file.Close()
				r = file
			}

			//Output stream
			w := ui.Writer()
			if len(out) > 0 {
				file, err := os.Create(out)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("output", out)))
				}
				defer file.Close()
				w = file
			}

			//Run command
			if err := nrm.Normalize(form, w, r, krFlag); err != nil {
				return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("file", inp), errs.WithContext("output", out)))
			}
			return nil
		},
	}
	normCmd.Flags().StringP("file", "f", "", "path of input text file")
	normCmd.Flags().StringP("output", "o", "", "path of output file")
	normCmd.Flags().StringP("norm-form", "n", "nfc", fmt.Sprintf("Unicode normalization form: [%s]", strings.Join(nrm.FormList(), "|")))
	normCmd.Flags().BoolP("kangxi-radicals", "k", false, "normalize kangxi radicals only (with nfkc or nfkd form)")

	return normCmd
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
