package facade

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gnkf/rbom"
	"github.com/goark/gnkf/width"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

//newNormCmd returns cobra.Command instance for show sub-command
func newWidthCmd(ui *rwi.RWI) *cobra.Command {
	widthCmd := &cobra.Command{
		Use:     "width",
		Aliases: []string{"wdth", "w"},
		Short:   "Convert character width in the text",
		Long:    "Convert character width in the text (UTF-8 encoding only).",
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
			form, err := cmd.Flags().GetString("conversion-form")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --conversion-form option", errs.WithCause(err)))
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
			if err := width.Convert(form, w, r); err != nil {
				return debugPrint(ui, errs.Wrap(err, errs.WithContext("file", inp), errs.WithContext("output", out)))
			}
			return nil
		},
	}
	widthCmd.Flags().StringP("file", "f", "", "path of input text file")
	_ = widthCmd.MarkFlagFilename("file")
	widthCmd.Flags().StringP("output", "o", "", "path of output file")
	_ = widthCmd.MarkFlagFilename("output")
	widthCmd.Flags().StringP("conversion-form", "c", "fold", fmt.Sprintf("conversion form: [%s]", strings.Join(width.FormList(), "|")))
	_ = widthCmd.RegisterFlagCompletionFunc("conversion-form", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return width.FormList(), cobra.ShellCompDirectiveDefault
	})
	widthCmd.Flags().BoolP("remove-bom", "b", false, "remove BOM character")

	return widthCmd
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
