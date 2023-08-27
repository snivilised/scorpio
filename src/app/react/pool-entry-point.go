package react

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/snivilised/lorax/async"
)

const (
	StopAfter   = 5
	RoutineName = async.GoRoutineName("👾 test-main")
)

func EnterPool(ps *PoolParameterSet) error {
	ctx := context.Background()

	fmt.Println("---> 🎯 orpheus(alpha) ...")

	ctxCancel, cancel := context.WithCancel(ctx)
	cancellations := []context.CancelFunc{cancel}

	pipe := start[TestJobInput, TestJobOutput](ps.OutputsChSize)
	sequence := 0

	fmt.Println("👾 WAIT-GROUP ADD(producer)")

	pipe.produce(ctxCancel, func() TestJobInput {
		recipient := rand.Intn(len(audience)) //nolint:gosec // trivial
		sequence++
		return TestJobInput{
			sequenceNo: sequence,
			Recipient:  audience[recipient],
		}
	}, ps.JobsChSize)

	fmt.Println("👾 WAIT-GROUP ADD(worker-pool)")

	pipe.process(ctxCancel, greeter, ps.NoWorkers)

	fmt.Println("👾 WAIT-GROUP ADD(consumer)")

	pipe.consume(ctxCancel)

	fmt.Println("👾 NOW AWAITING TERMINATION")

	if ps.DoCancel {
		pipe.cancel.After(ctxCancel, time.Second*time.Duration(ps.After), cancellations...)
	} else {
		pipe.stop.After(ctxCancel, time.Second*time.Duration(ps.After))
	}

	pipe.waiter.Wait(RoutineName)

	fmt.Printf("<--- orpheus(alpha) finished Counts >>> (Producer: '%v', Consumer: '%v'). 🎯🎯🎯\n",
		pipe.producer.Count,
		pipe.consumer.Count,
	)

	return nil
}
