package builder

import "fmt"

const (
	defaultDb       = 0
	defaultPassword = ""
	defaultRegion   = "hb"
)

type RedisClient struct {
	addr     string
	port     int
	password string
	db       int
	region   string
}

type RedisClientOption func(*RedisClient)

func WithDb(db int) RedisClientOption {
	return func(client *RedisClient) {
		client.db = db
	}
}

func WithPassword(password string) RedisClientOption {
	return func(client *RedisClient) {
		client.password = password
	}
}

func WithRegion(region string) RedisClientOption {
	return func(client *RedisClient) {
		client.region = region
	}
}

func NewRedisClient(addr string, port int, opts ...RedisClientOption) (*RedisClient, error) {
	if addr == "" {
		return nil, fmt.Errorf("addr can not be empty")
	}
	if port <= 0 {
		return nil, fmt.Errorf("port can not be less than 0")
	}

	client := &RedisClient{
		addr:     addr,
		port:     port,
		db:       defaultDb,
		password: defaultPassword,
		region:   defaultRegion,
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.region == "" {
		return nil, fmt.Errorf("region can not be empty")
	}

	return client, nil
}
