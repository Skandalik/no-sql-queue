package main

import (
	"fmt"
	"github.com/msales/pkg/v3/clix"
	"gopkg.in/urfave/cli.v1"
	"log"
	"math/rand"
	"time"
)

func runConsumer(c *cli.Context) error {
	ctx, err := clix.NewContext(c)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	consumer, err := createConsumer(ctx)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	log.Println("Consuming.")

	for len(clix.WaitForSignals()) == 0 {
		msgs := consumer.Consume()

		toSleep := rand.Int63n(1000)
		log.Println("Consumed", len(msgs), "messages. Sleeping for", toSleep, "milliseconds")
		time.Sleep(time.Duration(toSleep) * time.Millisecond)
	}

	log.Println("Finished.")

	return nil
}