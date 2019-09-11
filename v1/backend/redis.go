package backend

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

func NewRedisQueue(client *redis.Client, opts ...*QueueOption) *RedisQueue {
	// TODO: add default option
	if len(opts) == 0 {
		panic(fmt.Errorf("options should not empy"))
	}
	opt := opts[0]

	return &RedisQueue{
		option: opt,
		client: client,
		name:   opt.Name,
	}
}

// RedisQueue single name redis storage(using go-redis client)
type RedisQueue struct {
	client *redis.Client
	option *QueueOption

	// name of redis name
	name string
}

// Push job payload into queue
func (rq *RedisQueue) Push(_ context.Context, payload []byte) error {
	return rq.client.LPush(rq.name, payload).Err()
}

// PopOne job from queue
func (rq *RedisQueue) Pop(_ context.Context) ([]byte, error) {
	return rq.client.RPop(rq.name).Bytes()
}
