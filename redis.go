package corekit

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

var ErrEmptyRedisConfig = errors.New("empty redis config")

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func (config *RedisConfig) Connect() (client *redis.Client, err error) {
	if config == nil {
		return nil, ErrEmptyRedisConfig
	}

	client = redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})

	if _, err = client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
