package facade

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/bcrypt"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newNormCmd returns cobra.Command instance for show sub-command
func newBCryptCmd(ui *rwi.RWI) *cobra.Command {
	bcryptCmd := &cobra.Command{
		Use:     "bcrypt [flags] string [string...]",
		Aliases: []string{"bc"},
		Short:   "Hash and compare by BCrypt",
		Long:    "Hash and compare by BCrypt.",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			cost, err := cmd.Flags().GetInt("cost")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --cost option", errs.WithCause(err)))
			}
			hashed, err := cmd.Flags().GetString("compare")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --compare option", errs.WithCause(err)))
			}

			if len(args) == 0 {
				return debugPrint(ui, errs.Wrap(ecode.ErrNoData))
			}
			if len(hashed) > 0 {
				_ = ui.OutputErrln(fmt.Sprintf("compare BCrypt hashed string '%s' to...", hashed))
			}

			//Run command
			var lastErr error
			for _, s := range args {
				if len(hashed) > 0 {
					if err := bcrypt.Compare(hashed, s); err != nil {
						_ = ui.OutputErrln(s, ":", err)
					} else {
						_ = ui.OutputErrln(s, ":", "match!")
					}
				} else {
					if h, err := bcrypt.Hash(s, cost); err != nil {
						lastErr = errs.Wrap(err, errs.WithContext("string", s), errs.WithContext("cost", cost))
						_ = ui.OutputErrln(err)
					} else {
						_ = ui.Outputln(h)
					}
				}
			}
			return debugPrint(ui, lastErr)
		},
	}
	bcryptCmd.Flags().IntP("cost", "c", bcrypt.DefaultCost, fmt.Sprintf("BCrypt cost (%d-%d)", bcrypt.MinCost, bcrypt.MaxCost))
	bcryptCmd.Flags().StringP("compare", "", "", "compare to BCrypt hashed string")

	return bcryptCmd
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
