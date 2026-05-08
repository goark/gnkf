package facade

import (
	"os"
	"path/filepath"

	"github.com/goark/errs"
	"github.com/goark/gnkf/b64"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

// newNormCmd returns cobra.Command instance for show sub-command
func newBase64Cmd(ui *rwi.RWI) *cobra.Command {
	base64Cmd := &cobra.Command{
		Use:     "base64 [flags] [file]",
		Aliases: []string{"b64"},
		Short:   "Encode/Decode BASE64",
		Long:    "Encode/Decode BASE64.",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			//Options
			out, ferr := cmd.Flags().GetString("output")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --output option", errs.WithCause(ferr)))
				return
			}
			decodeFlag, ferr := cmd.Flags().GetBool("decode")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --decode option", errs.WithCause(ferr)))
				return
			}
			noPadding, ferr := cmd.Flags().GetBool("no-padding")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --no-padding option", errs.WithCause(ferr)))
				return
			}
			forURL, ferr := cmd.Flags().GetBool("for-url")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --for-url option", errs.WithCause(ferr)))
				return
			}

			//Input stream
			r := ui.Reader()
			if len(args) > 0 {
				file, ferr := os.Open(filepath.Clean(args[0]))
				if ferr != nil {
					err = debugPrint(ui, errs.Wrap(ferr, errs.WithContext("file", args[0])))
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
			if decodeFlag {
				err = b64.Decode(forURL, noPadding, r, w)
			} else {
				err = b64.Encode(forURL, noPadding, r, w)
			}
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, errs.WithContext("output", out)))
			}
			return nil
		},
	}
	base64Cmd.Flags().StringP("output", "o", "", "path of output file")
	_ = base64Cmd.MarkFlagFilename("output")
	base64Cmd.Flags().BoolP("decode", "d", false, "decode BASE64 string")
	base64Cmd.Flags().BoolP("no-padding", "p", false, "no padding")
	base64Cmd.Flags().BoolP("for-url", "u", false, "encoding/decoding defined in RFC 4648")

	return base64Cmd
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
