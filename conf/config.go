package conf

import "os"

var (
	RedisServer = os.Getenv("REDIS_URL")
	RedisPort   = os.Getenv("REDIS_PORT")

	MailgunAPIKey      = os.Getenv("MAILGUN_APIKEY")
	MailgunDomain      = os.Getenv("MAILGUN_DOMAIN")
	MailgunFromAddress = os.Getenv("MAILGUN_FROM_ADDRESS")
)
