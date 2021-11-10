package config

type rateLimit struct {
	Limit int
	Ttl   int
}

type rateLimitDirect struct {
	Limit int
	Ttl   int
}

type mysql struct {
	Host     string
	Port     string
	Dbname   string
	Username string
	Pwd      string
	Charset  string
	Prefix   string
}

type server struct {
	Name string
}

var RateLimit = new(rateLimit)
var RateLimitDirect = new(rateLimitDirect)
var Mysql = new(mysql)
var Server = new(server)
