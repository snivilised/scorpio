package react

import (
	"context"
	"time"

	"github.com/snivilised/lorax/boost"
)

type pipeline[I, O any] struct {
	wgan      boost.WaitGroupAn
	sequence  int
	outputsCh chan boost.JobOutput[O]
	provider  ProviderFn[I]
	producer  *Producer[I, O]
	pool      *boost.WorkerPool[I, O]
	consumer  *Consumer[O]
	cancel    TerminatorFunc[I, O]
	stop      TerminatorFunc[I, O]
}

func start[I, O any](outputsChSize int) *pipeline[I, O] {
	outputsCh := make(chan boost.JobOutput[O], outputsChSize)

	wgan := boost.NewAnnotatedWaitGroup("🍂 scorpio")
	pipe := &pipeline[I, O]{
		wgan:      wgan,
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
		p.wgan,
		jobChSize,
		provider,
		Delay,
	)

	p.wgan.Add(1, p.producer.RoutineName)
}

func (p *pipeline[I, O]) process(
	ctx context.Context,
	cancel context.CancelFunc,
	executive boost.ExecutiveFunc[I, O],
	noWorkers int,
) {
	var outputTimeout = time.Second * 2

	p.pool = boost.NewWorkerPool[I, O](
		&boost.NewWorkerPoolParams[I, O]{
			NoWorkers:       noWorkers,
			OutputChTimeout: outputTimeout,
			Exec:            executive,
			JobsCh:          p.producer.JobsCh,
			CancelCh:        make(boost.CancelStream),
			WaitAQ:          p.wgan,
		})

	go p.pool.Start(ctx, cancel, p.outputsCh)

	p.wgan.Add(1, p.pool.RoutineName)
}

func (p *pipeline[I, O]) consume(ctx context.Context, cancel context.CancelFunc) {
	p.consumer = StartConsumer(
		ctx,
		cancel,
		p.wgan,
		p.outputsCh,
	)

	p.wgan.Add(1, p.consumer.RoutineName)
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
