package genid

import "github.com/go-redis/redis/v8"

// Option operator the genid
type Option func(o *option)

type option struct {
	exp int           // 过期时间 默认两秒
	cli *redis.Client // 默认连接本地redis:16379
	db  int           // 默认的redis db 默认15
}

// Exp 更改默认redis过期时间
func Exp(e int) Option {
	return func(o *option) {
		o.exp = e
	}
}

// Cli 更改默认redis 连接
func Cli(c *redis.Client) Option {
	return func(o *option) {
		o.cli = c
	}
}

// DB 更改默认redis db
func DB(d int) Option {
	return func(o *option) {
		o.db = d
	}
}
