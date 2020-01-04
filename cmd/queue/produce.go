package main

import (
	"fmt"
	lorem "github.com/drhodes/golorem"
	"github.com/msales/pkg/v3/clix"
	"gopkg.in/urfave/cli.v1"
	"log"
	"math/rand"
	queue "no-sql-queue"
	"time"
)

func runProducer(c *cli.Context) error {
	ctx, err := clix.NewContext(c)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	producer, err := createProducer(ctx)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	log.Println("Producing.")
	for len(clix.WaitForSignals()) == 0 {
		msg := queue.Message{
			Text: lorem.Word(0, rand.Intn(200)),
		}
		producer.Produce(msg)

		toSleep := rand.Int63n(200)
		log.Println("Produced message. Sleeping for", toSleep, "milliseconds")
		time.Sleep(time.Duration(toSleep) * time.Millisecond)
	}
	log.Println("Finished.")

	return nil
}