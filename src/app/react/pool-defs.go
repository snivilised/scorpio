package react

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/lorax/boost"
)

const (
	Delay            = 750
	Interval         = 1000
	DefaultNoWorkers = 5
)

// PoolParameterSet
type PoolParameterSet struct {
	After         int
	DoCancel      bool
	NoWorkers     int
	BatchSize     int
	JobsChSize    int
	OutputsChSize int
	Delay         int
}

type PoolParamSetPtr = *assistant.ParamSet[PoolParameterSet]

var audience = []string{
	"👻 caspar",
	"🧙 gandalf",
	"😺 garfield",
	"👺 gobby",
	"👿 nick",
	"👹 ogre",
	"👽 paul",
	"🦄 pegasus",
	"💩 poo",
	"🤖 rusty",
	"💀 skeletor",
	"🐉 smaug",
	"🧛‍♀️ vampire",
	"👾 xenomorph",
}

type TestJobInput struct {
	sequenceNo int // allocated by observer
	Recipient  string
}

func (i TestJobInput) SequenceNo() int {
	return i.sequenceNo
}

type TestJobOutput = string
type TestResultStream chan boost.JobOutput[TestJobOutput]

var greeter = func(j boost.Job[TestJobInput]) (boost.JobOutput[TestJobOutput], error) {
	r := rand.Intn(Interval) + 1 //nolint:gosec // trivial
	delay := time.Millisecond * time.Duration(r)
	time.Sleep(delay)

	result := boost.JobOutput[TestJobOutput]{
		Payload: fmt.Sprintf("			---> 🍉🍉🍉 [Seq: %v] Hello: '%v'",
			j.Input.SequenceNo(), j.Input.Recipient,
		),
	}

	return result, nil
}

// TerminatorFunc brings the work pool processing to an end, eg
// by stopping or cancellation after the requested amount of time.
type TerminatorFunc[I, O any] func(ctx context.Context, delay time.Duration, funcs ...context.CancelFunc)

func (f TerminatorFunc[I, O]) After(ctx context.Context, delay time.Duration, funcs ...context.CancelFunc) {
	f(ctx, delay, funcs...)
}
