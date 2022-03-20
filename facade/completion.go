package facade

import (
	"fmt"
	"os"
	"strings"

	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

var (
	longDescription = `To load completions:

Bash:

$ source <(%[1]v completion bash)

# To load completions for each session, execute once:
Linux:
$ %[1]v completion bash > /etc/bash_completion.d/%[1]v
MacOS:
$ %[1]v completion bash > /usr/local/etc/bash_completion.d/%[1]v

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ %[1]v completion zsh > "${fpath[1]}/_%[1]v"

# You will need to start a new shell for this setup to take effect.

Fish:

$ %[1]v completion fish | source

# To load completions for each session, execute once:
$ %[1]v completion fish > ~/.config/fish/completions/%[1]v.fish

Powershell:

PS> %[1]v completion powershell | Out-String | Invoke-Expression

# To load completions for every new session, run:
PS> %[1]v completion powershell > %[1]v.ps1
# and source this file from your powershell profile.
`

	shells = []string{"bash", "zsh", "fish", "powershell"}
)

//newCompletionCmd returns cobra.Command instance for show sub-command
func newCompletionCmd(ui *rwi.RWI) *cobra.Command {
	completionCmd := &cobra.Command{
		Use:       "completion [" + strings.Join(shells, "|") + "]",
		Aliases:   []string{"compl", "cmp"},
		Short:     "Generate completion script",
		Long:      fmt.Sprintf(longDescription, Name),
		ValidArgs: shells,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return debugPrint(ui, cmd.Root().GenBashCompletion(ui.Writer()))
			} else if len(args) == 1 {
				switch {
				case strings.EqualFold(args[0], "bash"):
					return debugPrint(ui, cmd.Root().GenBashCompletion(ui.Writer()))
				case strings.EqualFold(args[0], "zsh"):
					return debugPrint(ui, cmd.Root().GenZshCompletion(ui.Writer()))
				case strings.EqualFold(args[0], "fish"):
					return debugPrint(ui, cmd.Root().GenFishCompletion(ui.Writer(), true))
				case strings.EqualFold(args[0], "powershell"):
					return debugPrint(ui, cmd.Root().GenPowerShellCompletion(ui.Writer()))
				}
			}
			return debugPrint(ui, os.ErrInvalid)
		},
	}
	return completionCmd
}

/* Copyright 2021 Spiegel
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
