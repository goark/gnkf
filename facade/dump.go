package facade

import (
	"os"
	"path/filepath"

	"github.com/goark/errs"
	"github.com/goark/gnkf/dump"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

// newDumpCmd returns cobra.Command instance for show sub-command
func newDumpCmd(ui *rwi.RWI) *cobra.Command {
	dumpCmd := &cobra.Command{
		Use:     "dump",
		Aliases: []string{"hexdump", "d", "hd"},
		Short:   "Hexadecimal view of octet data stream",
		Long:    "Hexadecimal view of octet data stream with C language array style.",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			//Options
			path, ferr := cmd.Flags().GetString("file")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --file option", errs.WithCause(ferr)))
				return
			}
			flagUnicode, ferr := cmd.Flags().GetBool("unicode")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --unicode option", errs.WithCause(ferr)))
				return
			}

			//Input stream
			r := ui.Reader()
			if len(path) > 0 {
				file, ferr := os.Open(filepath.Clean(path))
				if ferr != nil {
					err = debugPrint(ui, errs.Wrap(ferr, errs.WithContext("file", path)))
					return
				}
				defer func() {
					err = errs.Join(err, file.Close())
				}()
				r = file
			}

			//Run command
			if flagUnicode {
				err = dump.UnicodePoint(ui.Writer(), r)
			} else {
				err = dump.Octet(ui.Writer(), r)
			}
			if err != nil {
				err = debugPrint(ui, errs.Wrap(err, errs.WithContext("file", path)))
				return
			}
			err = debugPrint(ui, errs.Wrap(ui.Outputln(), errs.WithContext("file", path)))
			return
		},
	}
	dumpCmd.Flags().StringP("file", "f", "", "path of input text file")
	_ = dumpCmd.MarkFlagFilename("file")
	dumpCmd.Flags().BoolP("unicode", "u", false, "print by Unicode code point (UTF-8 only)")

	return dumpCmd
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
