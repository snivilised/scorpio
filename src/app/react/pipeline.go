package react

import (
	"context"
	"sync"
	"time"

	"github.com/snivilised/lorax/async"
)

type pipeline[I, O any] struct {
	wg        sync.WaitGroup
	sequence  int
	outputsCh chan async.JobOutput[O]
	provider  ProviderFn[I]
	producer  *Producer[I, O]
	pool      *async.WorkerPool[I, O]
	consumer  *Consumer[O]
	cancel    TerminatorFunc[I, O]
	stop      TerminatorFunc[I, O]
}

func start[I, O any](outputsChSize int) *pipeline[I, O] {
	outputsCh := make(chan async.JobOutput[O], outputsChSize)

	pipe := &pipeline[I, O]{
		outputsCh: outputsCh,
	}

	return pipe
}

func (p *pipeline[I, O]) produce(
	ctx context.Context,
	provider ProviderFn[I],
	jobChSize int,
) {
	p.cancel = func(ctx context.Context, delay time.Duration, cancellations ...context.CancelFunc) {
		go CancelProducerAfter[I, O](
			delay,
			cancellations...,
		)
	}
	p.stop = func(ctx context.Context, delay time.Duration, _ ...context.CancelFunc) {
		go StopProducerAfter(
			ctx,
			p.producer,
			delay,
		)
	}

	p.producer = StartProducer[I, O](
		ctx,
		&p.wg,
		jobChSize,
		provider,
		Delay,
	)

	p.wg.Add(1)
}

func (p *pipeline[I, O]) process(
	ctx context.Context,
	executive async.ExecutiveFunc[I, O],
	noWorkers int,
) {
	p.pool = async.NewWorkerPool[I, O](
		&async.NewWorkerPoolParams[I, O]{
			NoWorkers: noWorkers,
			Exec:      executive,
			JobsCh:    p.producer.JobsCh,
			CancelCh:  make(async.CancelStream),
			Quit:      &p.wg,
		})

	go p.pool.Start(ctx, p.outputsCh)

	p.wg.Add(1)
}

func (p *pipeline[I, O]) consume(ctx context.Context) {
	p.consumer = StartConsumer(ctx,
		&p.wg,
		p.outputsCh,
	)

	p.wg.Add(1)
}

func (p *pipeline[I, O]) stopProducerAfter(
	ctx context.Context,
	after time.Duration,
) {
	go StopProducerAfter(
		ctx,
		p.producer,
		after,
	)
}
