package react

import (
	"context"
	"fmt"

	"github.com/snivilised/lorax/boost"
)

type Consumer[O any] struct {
	RoutineName boost.GoRoutineName
	quitter     boost.WaitGroupAn
	OutputsCh   <-chan boost.JobOutput[O]
	Count       int
}

func StartConsumer[O any](
	ctx context.Context,
	cancel context.CancelFunc,
	quitter boost.WaitGroupAn,
	outputsCh <-chan boost.JobOutput[O],
) *Consumer[O] {
	consumer := &Consumer[O]{
		RoutineName: boost.GoRoutineName("💠 consumer"),
		quitter:     quitter,
		OutputsCh:   outputsCh,
	}
	go consumer.run(ctx, cancel)

	return consumer
}

func (c *Consumer[O]) run(ctx context.Context, _ context.CancelFunc) {
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
