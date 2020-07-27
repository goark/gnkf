package facade

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/dump"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newDumpCmd returns cobra.Command instance for show sub-command
func newDumpCmd(ui *rwi.RWI) *cobra.Command {
	dumpCmd := &cobra.Command{
		Use:     "dump",
		Aliases: []string{"hexdump", "d", "hd"},
		Short:   "Hexadecimal view of octet data stream",
		Long:    "Hexadecimal view of octet data stream with C language array style.",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			path, err := cmd.Flags().GetString("src")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --src option"))
			}
			flagUnicode, err := cmd.Flags().GetBool("unicode")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --unicode option"))
			}

			//Input stream
			r := ui.Reader()
			if len(path) > 0 {
				file, err := os.Open(path)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("path", path)))
				}
				defer file.Close()
				r = file
			}

			//Run command
			if flagUnicode {
				return debugPrint(ui, errs.Wrap(dump.UnicodePoint(ui.Writer(), r), "", errs.WithContext("path", path)))
			}
			return debugPrint(ui, errs.Wrap(dump.Octet(ui.Writer(), r), "", errs.WithContext("path", path)))
		},
	}
	dumpCmd.Flags().StringP("src", "s", "", "path of source text")
	dumpCmd.Flags().BoolP("unicode", "u", false, "print by Unicode code point (UTF-8 only)")

	return dumpCmd
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
