package command

import (
	"fmt"
	"runtime"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	xi18n "github.com/snivilised/extendio/i18n"

	"github.com/snivilised/scorpio/src/app/react"
	"github.com/snivilised/scorpio/src/i18n"
)

const poolPsName = "pool-ps"

func buildPoolCommand(container *assistant.CobraContainer) *cobra.Command {
	poolCommand := &cobra.Command{
		Use:   "pool",
		Short: xi18n.Text(i18n.PoolCmdShortDescTemplData{}),
		Long:  xi18n.Text(i18n.PoolCmdLongDescTemplData{}),
		RunE: func(cmd *cobra.Command, args []string) error {
			var appErr error

			ps := container.MustGetParamSet(poolPsName).(react.PoolParamSetPtr) //nolint:errcheck // is Must call

			if err := ps.Validate(); err == nil {
				native := ps.Native

				// optionally invoke cross field validation
				//
				if xv := ps.CrossValidate(func(ps *react.PoolParameterSet) error {
					return nil
				}); xv == nil {
					options := []string{}
					cmd.Flags().Visit(func(f *pflag.Flag) {
						options = append(options, fmt.Sprintf("--%v=%v", f.Name, f.Value))
					})
					fmt.Printf("%v %v Running pool, with options: '%v', args: '%v'\n",
						AppEmoji, ApplicationName, options, args,
					)

					appErr = react.EnterPool(native)
				} else {
					return xv
				}
			} else {
				return err
			}

			return appErr
		},
	}
	paramSet := assistant.NewParamSet[react.PoolParameterSet](poolCommand)

	// TODO: define the helper text for the flag (also applies to other flag)
	//

	// --after
	//
	defaultStopAfter := 1

	const (
		minStopAfter = 1
		maxStopAfter = 30
	)

	paramSet.BindValidatedIntWithin(
		assistant.NewFlagInfo("after", "a", defaultStopAfter),
		&paramSet.Native.After,
		minStopAfter,
		maxStopAfter,
	)

	// --cancel
	//
	paramSet.BindBool(
		assistant.NewFlagInfo("cancel", "c", false),
		&paramSet.Native.DoCancel,
	)

	// --now
	//
	defaultNoWorkers := runtime.NumCPU()

	const (
		minNoWorkers = 1
		maxNoWorkers = 16
	)

	paramSet.BindValidatedIntWithin(
		assistant.NewFlagInfo("now", "n", defaultNoWorkers),
		&paramSet.Native.NoWorkers,
		minNoWorkers,
		maxNoWorkers,
	)

	// --jobq int (1-20)
	//
	const (
		minJobQueueSize = 1
		maxJobQueueSize = 20
	)

	paramSet.BindValidatedIntWithin(
		assistant.NewFlagInfo("jobq", "j", defaultNoWorkers),
		&paramSet.Native.JobsChSize,
		minJobQueueSize,
		maxJobQueueSize,
	)

	// --resq int (1-16)
	//
	const (
		minResultsChSize = 1
		maxResultsChSize = 20
	)

	paramSet.BindValidatedIntWithin(
		assistant.NewFlagInfo("resq", "r", defaultNoWorkers),
		&paramSet.Native.OutputsChSize,
		minResultsChSize,
		maxResultsChSize,
	)

	// -- delay (1-20)
	//
	const (
		defaultDelay = 10
		minDelay     = 1
		maxDelay     = 20
	)

	paramSet.BindValidatedIntWithin(
		assistant.NewFlagInfo("delay", "d", defaultDelay),
		&paramSet.Native.Delay,
		minDelay,
		maxDelay,
	)

	// --batch int (1-50)
	//
	const (
		defaultBatch = 13
		minBatch     = 1
		maxBatch     = 16
	)

	paramSet.BindValidatedIntWithin(
		assistant.NewFlagInfo("batch", "b", defaultBatch),
		&paramSet.Native.BatchSize,
		minBatch,
		maxBatch,
	)

	container.MustRegisterRootedCommand(poolCommand)
	container.MustRegisterParamSet(poolPsName, paramSet)

	return poolCommand
}
