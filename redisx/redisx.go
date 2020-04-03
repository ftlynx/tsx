package redisx

import "github.com/go-redis/redis"

func InitRedis(addr string, passwd string, db int) (*redis.Client, error) {
	if addr == "" {
		addr = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})

	_, err := client.Ping().Result()
	return client, err
}
