package react

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/lorax/async"
)

const (
	Delay            = 750
	Interval         = 1000
	DefaultNoWorkers = 5
)

// PoolParameterSet
type PoolParameterSet struct {
	StopAfter     int
	NoWorkers     int
	BatchSize     int
	JobsChSize    int
	ResultsChSize int
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

type TestJobResult = string
type TestResultStream chan async.JobResult[TestJobResult]

var greeter = func(j async.Job[TestJobInput]) (async.JobResult[TestJobResult], error) {
	r := rand.Intn(Interval) + 1 //nolint:gosec // trivial
	delay := time.Millisecond * time.Duration(r)
	time.Sleep(delay)

	result := async.JobResult[TestJobResult]{
		Payload: fmt.Sprintf("			---> 🍉🍉🍉 [Seq: %v] Hello: '%v'",
			j.Input.SequenceNo(), j.Input.Recipient,
		),
	}

	return result, nil
}
