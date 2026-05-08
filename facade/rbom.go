package facade

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/goark/errs"
	"github.com/goark/gnkf/rbom"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

// newNormCmd returns cobra.Command instance for show sub-command
func newRemoveBomCmd(ui *rwi.RWI) *cobra.Command {
	rbomCmd := &cobra.Command{
		Use:     "remove-bom",
		Aliases: []string{"rbom", "rb"},
		Short:   "Remove BOM character in UTF-8 string",
		Long:    "Remove BOM character in UTF-8 string.",
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

			//Run command
			b, rerr := rbom.RemoveBom(r)
			if rerr != nil {
				err = debugPrint(ui, errs.Wrap(rerr, errs.WithContext("file", inp), errs.WithContext("output", out)))
				return
			}
			if _, cerr := io.Copy(w, bytes.NewReader(b)); cerr != nil {
				err = debugPrint(ui, errs.Wrap(cerr, errs.WithContext("file", inp), errs.WithContext("output", out)))
				return
			}
			return
		},
	}
	rbomCmd.Flags().StringP("file", "f", "", "path of input text file")
	_ = rbomCmd.MarkFlagFilename("file")
	rbomCmd.Flags().StringP("output", "o", "", "path of output file")
	_ = rbomCmd.MarkFlagFilename("output")

	return rbomCmd
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
