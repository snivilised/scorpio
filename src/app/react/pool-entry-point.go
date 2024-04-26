package react

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/snivilised/lorax/boost"
)

const (
	StopAfter   = 5
	RoutineName = boost.GoRoutineName("👾 test-main")
)

func EnterPool(ps *PoolParameterSet) error {
	fmt.Println("---> 🎯 orpheus(alpha) ...")

	ctx, cancel := context.WithCancel(context.Background())
	cancellations := []context.CancelFunc{cancel}

	pipe := start[TestJobInput, TestJobOutput](ps.OutputsChSize)
	sequence := 0

	fmt.Println("👾 WAIT-GROUP ADD(producer)")

	pipe.produce(ctx, func() TestJobInput {
		recipient := rand.Intn(len(audience)) //nolint:gosec // trivial
		sequence++

		return TestJobInput{
			sequenceNo: sequence,
			Recipient:  audience[recipient],
		}
	}, ps.JobsChSize)

	fmt.Println("👾 WAIT-GROUP ADD(worker-pool)")

	pipe.process(ctx, cancel, greeter, ps.NoWorkers)

	fmt.Println("👾 WAIT-GROUP ADD(consumer)")

	pipe.consume(ctx, cancel)

	fmt.Println("👾 NOW AWAITING TERMINATION")

	if ps.DoCancel {
		pipe.cancel.After(ctx, time.Second*time.Duration(ps.After), cancellations...)
	} else {
		pipe.stop.After(ctx, time.Second*time.Duration(ps.After))
	}

	pipe.wgan.Wait(RoutineName)

	fmt.Printf("<--- orpheus(alpha) finished Counts >>> (Producer: '%v', Consumer: '%v'). 🎯🎯🎯\n",
		pipe.producer.Count,
		pipe.consumer.Count,
	)

	return nil
}
