package genid

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	exp  int    = 2 // 默认两秒过期时间
	addr string = "127.0.0.1:16379"
	db   int    = 15
)

// Generator id generate interface
type Generator interface {
	// Creaet create one id code
	Create(ctx context.Context, biz string) (string, error)
	// CreateBatch batch create biz code
	CreateBatch(ctx context.Context, biz string, num int) ([]string, error)
}

func New(opts ...Option) (Generator, error) {
	o := option{
		exp: exp,
		cli: redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   db,
		}),
		db: db,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return &generator{
		cli:    o.cli,
		expSec: o.exp, // 默认两秒过期时间
		db:     o.db,
	}, nil
}
