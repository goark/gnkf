package facade

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gnkf/rbom"
	"github.com/goark/gnkf/width"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

// newNormCmd returns cobra.Command instance for show sub-command
func newWidthCmd(ui *rwi.RWI) *cobra.Command {
	widthCmd := &cobra.Command{
		Use:     "width",
		Aliases: []string{"wdth", "w"},
		Short:   "Convert character width in the text",
		Long:    "Convert character width in the text (UTF-8 encoding only).",
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
			form, ferr := cmd.Flags().GetString("conversion-form")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --conversion-form option", errs.WithCause(ferr)))
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
					err = debugPrint(ui, errs.Wrap(ferr, errs.WithContext("file", inp)))
					return
				}
				defer func() {
					err = errs.Join(err, file.Close())
				}()
				r = file
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
				b, rerr := rbom.RemoveBom(r)
				if rerr != nil {
					err = debugPrint(ui, errs.Wrap(rerr, errs.WithContext("file", inp), errs.WithContext("output", out)))
					return
				}
				r = bytes.NewReader(b)
			}

			//Run command
			if werr := width.Convert(form, w, r); werr != nil {
				err = debugPrint(ui, errs.Wrap(werr, errs.WithContext("file", inp), errs.WithContext("output", out)))
				return
			}
			return
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
