package conf

import "os"

var (
	RedisServer = os.Getenv("REDIS_URL")
	RedisPort   = os.Getenv("REDIS_PORT")
)
