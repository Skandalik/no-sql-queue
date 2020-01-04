package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/msales/pkg/v3/clix"
	queue "no-sql-queue"
)

func createRedis(ctx *clix.Context) (redis.Cmdable, error) {
	options := &redis.UniversalOptions{
		Addrs:         ctx.StringSlice(flagRedisDSNs),
		RouteRandomly: true,
	}

	client := redis.NewUniversalClient(options)
	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return client, nil
}

func createProducer(ctx *clix.Context) (queue.Producer, error) {
	r, err := createRedis(ctx)
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return queue.NewRedisProducer(ctx.String(flagListKey), r), nil
}

func createConsumer(ctx *clix.Context) (queue.Consumer, error) {
	r, err := createRedis(ctx)
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return queue.NewRedisConsumer(ctx.String(flagListKey), ctx.Int(flagBatchSize), r), nil
}