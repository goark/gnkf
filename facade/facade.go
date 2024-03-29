package facade

import (
	"fmt"
	"runtime"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

var (
	//Name is applicatin name
	Name = "gnkf"
	//Version is version for applicatin
	Version = "developer version"
)
var (
	debugFlag bool //debug flag
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "Network Kanji Filter by Golang",
		Long:  "Network Kanji Filter by Golang",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugPrint(ui, errs.Wrap(ecode.ErrNoCommand))
		},
	}
	rootCmd.SilenceUsage = true
	rootCmd.SetArgs(args)            //arguments of command-line
	rootCmd.SetIn(ui.Reader())       //Stdin
	rootCmd.SetOut(ui.ErrorWriter()) //Stdout -> Stderr
	rootCmd.SetErr(ui.ErrorWriter()) //Stderr
	rootCmd.AddCommand(
		newVersionCmd(ui),
		newGuessCmd(ui),
		newDumpCmd(ui),
		newEncCmd(ui),
		newNormCmd(ui),
		newNwlnCmd(ui),
		newWidthCmd(ui),
		newKanaCmd(ui),
		newBase64Cmd(ui),
		newRemoveBomCmd(ui),
		newCompletionCmd(ui),
		newhashCmd(ui),
		newBCryptCmd(ui),
	)

	//global options
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "for debug")

	return rootCmd
}

func debugPrint(ui *rwi.RWI, err error) error {
	if debugFlag && err != nil {
		fmt.Fprintf(ui.Writer(), "%+v\n", err)
	}
	return err
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
