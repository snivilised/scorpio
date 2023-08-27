package react

import (
	"context"
	"time"

	"github.com/snivilised/lorax/async"
)

type pipeline[I, O any] struct {
	adder     async.AssistedAdder
	quitter   async.AssistedQuitter
	waiter    async.AssistedWaiter
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

	wgex := async.NewAnnotatedWaitGroup("🍂 scorpio")
	pipe := &pipeline[I, O]{
		adder:     wgex,
		quitter:   wgex,
		waiter:    wgex,
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
		p.adder,
		p.quitter,
		jobChSize,
		provider,
		Delay,
	)

	p.adder.Add(1, p.producer.RoutineName)
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
			Quitter:   p.quitter,
		})

	go p.pool.Start(ctx, p.outputsCh)

	p.adder.Add(1, p.pool.RoutineName)
}

func (p *pipeline[I, O]) consume(ctx context.Context) {
	p.consumer = StartConsumer(ctx,
		p.quitter,
		p.outputsCh,
	)

	p.adder.Add(1, p.consumer.RoutineName)
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
