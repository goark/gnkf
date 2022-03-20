package facade

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gnkf/kana"
	"github.com/goark/gnkf/newline"
	"github.com/goark/gnkf/rbom"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

var descriptionKana = `Convert kana characters in the text.
 UTF-8 encoding only.
 "hiragana" and "katakana" forms are valid only for full-width kana character.`

//newNormCmd returns cobra.Command instance for show sub-command
func newKanaCmd(ui *rwi.RWI) *cobra.Command {
	kanaCmd := &cobra.Command{
		Use:     "kana",
		Aliases: []string{"k"},
		Short:   "Convert kana characters in the text",
		Long:    descriptionKana,
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
			formName, err := cmd.Flags().GetString("conversion-form")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --conversion-form option", errs.WithCause(err)))
			}
			form, err := kana.FormOf(formName)
			if err != nil {
				return debugPrint(ui, err)
			}
			foldFlag, err := cmd.Flags().GetBool("fold")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --fold option", errs.WithCause(err)))
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
			if err := kana.Convert(form, w, r, foldFlag); err != nil {
				return debugPrint(ui, errs.Wrap(err, errs.WithContext("file", inp), errs.WithContext("output", out)))
			}
			return nil
		},
	}
	kanaCmd.Flags().StringP("file", "f", "", "path of input text file")
	_ = kanaCmd.MarkFlagFilename("file")
	kanaCmd.Flags().StringP("output", "o", "", "path of output file")
	_ = kanaCmd.MarkFlagFilename("output")
	kanaCmd.Flags().StringP("conversion-form", "c", "katakana", fmt.Sprintf("conversion form: [%s]", strings.Join(kana.FormList(), "|")))
	_ = kanaCmd.RegisterFlagCompletionFunc("conversion-form", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return newline.FormList(), cobra.ShellCompDirectiveNoFileComp
	})
	kanaCmd.Flags().BoolP("fold", "", false, "convert character width by fold form")
	kanaCmd.Flags().BoolP("remove-bom", "b", false, "remove BOM character")

	return kanaCmd
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
