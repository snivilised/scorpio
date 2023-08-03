package domain

import "github.com/snivilised/cobrass/src/assistant"

// PoolParameterSet
type PoolParameterSet struct {
	NumCPU          int
	BatchSize       int
	JobQueueSize    int
	ResultQueueSize int
	Delay           int
}

type PoolParamSetPtr = *assistant.ParamSet[PoolParameterSet]
