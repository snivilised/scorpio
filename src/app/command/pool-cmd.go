package command

import (
	"fmt"
	"runtime"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	xi18n "github.com/snivilised/extendio/i18n"

	"github.com/snivilised/scorpio/src/app/domain"
	"github.com/snivilised/scorpio/src/i18n"
)

// CLIENT-TODO: rename this pool command to something required
// by the application. Pool command is meant to just serve as
// an aid in creating custom commands and intended to either be
// replaced or renamed.

const poolPsName = "pool-ps"

func buildPoolCommand(container *assistant.CobraContainer) *cobra.Command {
	// to test: arcadia pool -d ./some-existing-file -p "P?<date>" -t 30
	//
	poolCommand := &cobra.Command{
		Use:   "pool",
		Short: xi18n.Text(i18n.PoolCmdShortDescTemplData{}),
		Long:  xi18n.Text(i18n.PoolCmdLongDescTemplData{}),
		RunE: func(cmd *cobra.Command, args []string) error {
			var appErr error

			ps := container.MustGetParamSet(poolPsName).(domain.PoolParamSetPtr) //nolint:errcheck // is Must call

			if err := ps.Validate(); err == nil {
				native := ps.Native

				// optionally invoke cross field validation
				//
				if xv := ps.CrossValidate(func(ps *domain.PoolParameterSet) error {
					return nil
				}); xv == nil {
					options := []string{}
					cmd.Flags().Visit(func(f *pflag.Flag) {
						options = append(options, fmt.Sprintf("--%v=%v", f.Name, f.Value))
					})
					fmt.Printf("%v %v Running pool, with options: '%v', args: '%v'\n",
						AppEmoji, ApplicationName, options, args,
					)

					appErr = domain.EnterPool(native)
				} else {
					return xv
				}
			} else {
				return err
			}

			return appErr
		},
	}
	paramSet := assistant.NewParamSet[domain.PoolParameterSet](poolCommand)

	// TODO: define the helper text for the flag (also applies to other flag)
	//

	// --cpu
	//
	defaultCPU := runtime.NumCPU()

	const (
		minCPU = 1
		maxCPU = 16
	)

	paramSet.BindValidatedIntWithin(
		assistant.NewFlagInfo("cpu", "c", defaultCPU),
		&paramSet.Native.NumCPU,
		minCPU,
		maxCPU,
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

	// --jobq int (1-20)
	//
	const (
		minJobQueueSize = 1
		maxJobQueueSize = 20
	)

	paramSet.BindValidatedIntWithin(
		assistant.NewFlagInfo("jobq", "j", defaultCPU),
		&paramSet.Native.JobQueueSize,
		minJobQueueSize,
		maxJobQueueSize,
	)

	// --resq int (1-16)
	//
	const (
		minResultQueueSize = 1
		maxResultQueueSize = 20
	)

	paramSet.BindValidatedIntWithin(
		assistant.NewFlagInfo("resq", "r", defaultCPU),
		&paramSet.Native.ResultQueueSize,
		minResultQueueSize,
		maxResultQueueSize,
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

	container.MustRegisterRootedCommand(poolCommand)
	container.MustRegisterParamSet(poolPsName, paramSet)

	return poolCommand
}
