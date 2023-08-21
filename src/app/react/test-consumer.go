package react

import (
	"context"
	"fmt"
	"sync"

	"github.com/snivilised/lorax/async"
)

type Consumer[O any] struct {
	quit      *sync.WaitGroup
	OutputsCh <-chan async.JobOutput[O]
	Count     int
}

func StartConsumer[O any](
	ctx context.Context,
	wg *sync.WaitGroup,
	outputsCh <-chan async.JobOutput[O],
) *Consumer[O] {
	consumer := &Consumer[O]{
		quit:      wg,
		OutputsCh: outputsCh,
	}
	go consumer.run(ctx)

	return consumer
}

func (c *Consumer[O]) run(ctx context.Context) {
	defer func() {
		c.quit.Done()
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
