package react

import (
	"context"
	"sync"
	"time"

	"github.com/snivilised/lorax/async"
)

type pipeline[I, R any] struct {
	wg        sync.WaitGroup
	sequence  int
	resultsCh chan async.JobResult[R]
	provider  ProviderFn[I]
	producer  *Producer[I, R]
	pool      *async.WorkerPool[I, R]
	consumer  *Consumer[R]
}

func start[I, R any](resultsChSize int) *pipeline[I, R] {
	resultsCh := make(chan async.JobResult[R], resultsChSize)

	pipe := &pipeline[I, R]{
		resultsCh: resultsCh,
	}

	return pipe
}

func (p *pipeline[I, R]) startProducer(
	ctx context.Context,
	provider ProviderFn[I],
	jobChSize int,
) {
	p.producer = StartProducer[I, R](
		ctx,
		&p.wg,
		jobChSize,
		provider,
		Delay,
	)

	p.wg.Add(1)
}

func (p *pipeline[I, R]) startPool(
	ctx context.Context,
	executive async.ExecutiveFunc[I, R],
	noWorkers int,
) {
	p.pool = async.NewWorkerPool[I, R](
		&async.NewWorkerPoolParams[I, R]{
			NoWorkers: noWorkers,
			Exec:      executive,
			JobsCh:    p.producer.JobsCh,
			CancelCh:  make(async.CancelStream),
			Quit:      &p.wg,
		})

	go p.pool.Start(ctx, p.resultsCh)

	p.wg.Add(1)
}

func (p *pipeline[I, R]) startConsumer(ctx context.Context) {
	p.consumer = StartConsumer(ctx,
		&p.wg,
		p.resultsCh,
	)

	p.wg.Add(1)
}

func (p *pipeline[I, R]) stopProducerAfter(
	ctx context.Context,
	after time.Duration,
) {
	go StopProducerAfter(
		ctx,
		p.producer,
		after,
	)
}
