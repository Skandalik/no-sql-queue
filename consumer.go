package queue

import (
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"log"
)

type Consumer interface {
	Consume() []Message
}

type RedisConsumer struct {
	list      string
	batchSize int
	redis     redis.Cmdable
}

func NewRedisConsumer(list string, batchSize int, redis redis.Cmdable) *RedisConsumer {
	return &RedisConsumer{list: list, batchSize: batchSize, redis: redis}
}

func (r *RedisConsumer) Consume() []Message {
	messages := make([]Message, 0)

	for len(messages) < r.batchSize {
		encoded := r.redis.RPop(r.list).Val()
		if encoded == "" {
			log.Println("no message found")
			break
		}

		message := &Message{}

		err := jsoniter.Unmarshal([]byte(encoded), message)
		if err != nil {
			log.Println("consumer", err.Error())
		}

		messages = append(messages, *message)
	}

	return messages
}
