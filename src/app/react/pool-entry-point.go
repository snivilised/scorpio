package react

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	StopAfter = 5
)

func EnterPool(ps *PoolParameterSet) error {
	ctx := context.Background()

	fmt.Println("---> 🎯 orpheus(alpha) ...")

	pipe := start[TestJobInput, TestJobResult](ps.ResultsChSize)
	sequence := 0

	fmt.Println("👾 WAIT-GROUP ADD(producer)")

	pipe.startProducer(ctx, func() TestJobInput {
		recipient := rand.Intn(len(audience)) //nolint:gosec // trivial
		sequence++
		return TestJobInput{
			sequenceNo: sequence,
			Recipient:  audience[recipient],
		}
	}, ps.JobsChSize)

	fmt.Println("👾 WAIT-GROUP ADD(worker-pool)")

	pipe.startPool(ctx, greeter, ps.NoWorkers)

	fmt.Println("👾 WAIT-GROUP ADD(consumer)")

	pipe.startConsumer(ctx)

	fmt.Println("👾 NOW AWAITING TERMINATION")

	pipe.stopProducerAfter(ctx, time.Second*time.Duration(ps.StopAfter))
	pipe.wg.Wait()

	fmt.Printf("<--- orpheus(alpha) finished Counts >>> (Producer: '%v', Consumer: '%v'). 🎯🎯🎯\n",
		pipe.producer.Count,
		pipe.consumer.Count,
	)

	return nil
}
