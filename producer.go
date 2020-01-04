package queue

import (
	"github.com/go-redis/redis"
	"log"
)

type Producer interface {
	Produce(Message)
}

type RedisProducer struct {
	listKey string
	redis   redis.Cmdable
}

func NewRedisProducer(listKey string, redis redis.Cmdable) *RedisProducer {
	return &RedisProducer{listKey: listKey, redis: redis}
}

func (r *RedisProducer) Produce(m Message) {
	encoded, err := m.Marshal()
	if err != nil {
		log.Println("producer", err.Error())
	}

	r.redis.LPush(r.listKey, encoded)
}
