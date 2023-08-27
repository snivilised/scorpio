package react

import (
	"context"
	"fmt"

	"github.com/snivilised/lorax/async"
)

type Consumer[O any] struct {
	RoutineName async.GoRoutineName
	quitter     async.AssistedQuitter
	OutputsCh   <-chan async.JobOutput[O]
	Count       int
}

func StartConsumer[O any](
	ctx context.Context,
	quitter async.AssistedQuitter,
	outputsCh <-chan async.JobOutput[O],
) *Consumer[O] {
	consumer := &Consumer[O]{
		RoutineName: async.GoRoutineName("💠 consumer"),
		quitter:     quitter,
		OutputsCh:   outputsCh,
	}
	go consumer.run(ctx)

	return consumer
}

func (c *Consumer[O]) run(ctx context.Context) {
	defer func() {
		c.quitter.Done(c.RoutineName)
		fmt.Printf("<<<< consumer.run - finished (QUIT). 💠💠💠 \n")
	}()
	fmt.Printf("<<<< 💠 consumer.run ...\n")

	for running := true; running; {
		select {
		case <-ctx.Done():
			running = false

			fmt.Println("<<<< 💠 consumer.run - done received 💔💔💔")

		case result, ok := <-c.OutputsCh:
			if ok {
				c.Count++
				fmt.Printf("<<<< 💠 consumer.run - new result arrived(#%v): '%+v' \n",
					c.Count, result.Payload,
				)
			} else {
				running = false
				fmt.Printf("<<<< 💠 consumer.run - no more results available (running: %+v)\n", running)
			}
		}
	}
}
