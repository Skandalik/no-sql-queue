package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/msales/pkg/v3/clix"

	"gopkg.in/urfave/cli.v1"
	"os"
)

const (
	flagRedisDSNs = "redis-dsns"
	flagBatchSize = "batch-size"
	flagListKey   = "list-key"
)

var commonFlags = clix.Flags{
	cli.StringSliceFlag{
		Name:   flagRedisDSNs,
		Usage:  "Redis DSNs to connect to.",
		EnvVar: "REDIS_DSNS",
	},
	cli.StringFlag{
		Name:   flagListKey,
		Usage:  "Key to save/read from.",
		EnvVar: "LIST_KEY",
	},
}

var commands = []cli.Command{
	{
		Name:   "producer",
		Usage:  "Run producer",
		Flags:  commonFlags,
		Action: runProducer,
	},
	{
		Name:  "consumer",
		Usage: "Run consumer",
		Flags: clix.Flags.Merge(
			commonFlags,
			clix.Flags{
				cli.StringFlag{
					Name:   flagBatchSize,
					Usage:  "Batch size for consumer.",
					EnvVar: "BATCH_SIZE",
				},
			},
		),
		Action: runConsumer,
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "Queue app"
	app.Commands = commands
	app.Version = "beta"

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
