package facade

import (
	"crypto"
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/goark/errs"
	"github.com/goark/gnkf/ecode"
	"github.com/goark/gnkf/hash"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

// newhashCmd returns cobra.Command instance for show sub-command
func newhashCmd(ui *rwi.RWI) *cobra.Command {
	hashCmd := &cobra.Command{
		Use:     "hash [flags] [file]",
		Aliases: []string{"h"},
		Short:   "Print or check hash value",
		Long:    "Print or check hash value.\n  Support algorithm:\n  " + hash.AlgorithmList(", "),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			//Options
			s, ferr := cmd.Flags().GetString("algorithm")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --algorithm option", errs.WithCause(ferr)))
				return
			}
			alg, herr := hash.Algorithm(s)
			if herr != nil {
				err = debugPrint(ui, errs.Wrap(herr, errs.WithContext("algorithm", s)))
				return
			}
			checkerFlag, ferr := cmd.Flags().GetBool("check")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --check option", errs.WithCause(ferr)))
				return
			}
			ignoreMissingFlag, ferr := cmd.Flags().GetBool("ignore-missing")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --ignore-missing option", errs.WithCause(ferr)))
				return
			}
			quietFlag, ferr := cmd.Flags().GetBool("quiet")
			if ferr != nil {
				err = debugPrint(ui, errs.New("Error in --quiet option", errs.WithCause(ferr)))
				return
			}

			//Input stream
			inp := "-"
			r := ui.Reader()
			if len(args) > 0 && args[0] != inp {
				inp = args[0]
				file, ferr := os.Open(inp)
				if ferr != nil {
					err = debugPrint(ui, errs.Wrap(ferr, errs.WithContext("file", inp)))
					return
				}
				defer func() {
					err = errs.Join(err, file.Close())
				}()
				r = file
			}

			//Run command
			var lastError error
			if checkerFlag {
				checkers, herr := hash.NewCheckers(r, alg)
				if herr != nil {
					err = debugPrint(ui, errs.Wrap(errs.Join(lastError, herr), errs.WithContext("algorithm", alg.String()), errs.WithContext("file", inp)))
					return
				}
				lastError = hashChecks(checkers, ui, ignoreMissingFlag, quietFlag)
				if hashValidCount(checkers) == 0 {
					lastError = errs.New(fmt.Sprintf("%s: no file was verified", inp), errs.WithContext("algorithm", alg.String()), errs.WithContext("file", inp))
				}
			} else {
				res, herr := newHashValue(alg, r, inp)
				if herr != nil {
					err = debugPrint(ui, errs.Wrap(errs.Join(lastError, herr), errs.WithContext("algorithm", alg.String()), errs.WithContext("file", inp)))
					return
				}
				lastError = ui.Outputln(res.String())
			}
			err = debugPrint(ui, errs.Wrap(lastError, errs.WithContext("algorithm", alg.String()), errs.WithContext("file", inp)))
			return
		},
	}
	hashCmd.Flags().StringP("algorithm", "a", "SHA-256", "hash algorithm")
	hashCmd.Flags().BoolP("check", "c", false, "don't fail or report status for missing files")
	hashCmd.Flags().BoolP("ignore-missing", "", false, "don't fail or report status for missing files (with check option)")
	hashCmd.Flags().BoolP("quiet", "", false, "don't print OK for each successfully verified file (with check option)")

	return hashCmd
}

type hashValue struct {
	alg   crypto.Hash
	path  string
	value []byte
}

func newHashValue(alg crypto.Hash, r io.Reader, path string) (*hashValue, error) {
	value, err := hash.Value(alg, r)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("algorithm", alg.String()))
	}
	return &hashValue{alg: alg, path: path, value: value}, nil
}

func (hv *hashValue) hashString() string {
	if hv == nil {
		return ""
	}
	return fmt.Sprintf("%x", hv.value)
}

func (hv *hashValue) String() string {
	if hv == nil {
		return ""
	}
	return fmt.Sprintf("%v  %s", hv.hashString(), hv.path)
}

type warn struct {
	count int
	err   error
}

func (w warn) Error() string {
	if w.count > 1 {
		return fmt.Sprintf("WARNING in %d items: %v", w.count, w.err)
	}
	return fmt.Sprintf("Warning in %d item: %v", w.count, w.err)
}

func appendHashError(wlist []*warn, err error) []*warn {
	if len(wlist) == 0 {
		return append(wlist, &warn{count: 1, err: err})
	}
	for i := 0; i < len(wlist); i++ {
		if errs.Is(err, wlist[i].err) {
			wlist[i].count++
			return wlist
		}
	}
	return append(wlist, &warn{count: 1, err: err})
}

func hashChecks(checkers []hash.Checker, ui *rwi.RWI, ignoreMissingFlag, quietFlag bool) error {
	wlist := []*warn{}
	var lastError error
	for _, chkr := range checkers {
		err := chkr.Check()
		if err != nil {
			switch true {
			case errs.Is(err, syscall.ENOENT):
				wlist = appendHashError(wlist, syscall.ENOENT)
				if !ignoreMissingFlag {
					lastError = ui.OutputErrln(err)
				}
			case errs.Is(err, ecode.ErrUnmatchHashString):
				wlist = appendHashError(wlist, ecode.ErrUnmatchHashString)
				lastError = ui.Outputln(fmt.Sprintf("%s: FAILED", chkr.Path()))
			default:
				wlist = appendHashError(wlist, err)
				lastError = ui.OutputErrln(err)
			}
		} else if !quietFlag {
			lastError = ui.Outputln(fmt.Sprintf("%s: OK", chkr.Path()))
		}
		if lastError != nil {
			return lastError
		}
	}
	for _, w := range wlist {
		if err := ui.Outputln(w); err != nil {
			return err
		}
	}
	return nil
}

func hashValidCount(checkers []hash.Checker) int {
	count := 0
	for _, chk := range checkers {
		if chk.Err() == nil {
			count++
		}
	}
	return count
}

/* Copyright 2021-2026 Spiegel
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
