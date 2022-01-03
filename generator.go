package genid

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type generator struct {
	cli    *redis.Client
	expSec int
	db     int
}

func (g *generator) Create(ctx context.Context, biz string) (string, error) {
	key := key(biz)
	// 检测本地时间与服务器时间是否存在偏差
	if err := g.check(ctx); err != nil {
		return "", err
	}
	res, err := g.sequences(ctx, key, 1, g.expSec)
	if err != nil {
		return "", err
	}
	if res == 0 {
		return "", errors.New("lua script run fail")
	}
	return code(key, res), nil
}

func (g *generator) CreateBatch(ctx context.Context, biz string, num int) ([]string, error) {
	key := key(biz)
	if err := g.check(ctx); err != nil {
		return nil, err
	}
	res, err := g.sequences(ctx, key, num, g.expSec)
	if err != nil {
		return nil, err
	}
	if res == 0 {
		return nil, errors.New("lua script run fail")
	}
	var codes = []string{}
	for i := res - num + 1; i <= res; i++ {
		sn := i
		sc := code(key, sn)
		codes = append(codes, sc)
	}
	return codes, nil
}

func (g *generator) sequences(ctx context.Context, key string, num, expire int) (int, error) {
	result, err := atomGetSequences.Run(
		ctx,
		g.cli,
		[]string{key},
		[]string{
			strconv.Itoa(num),
			strconv.Itoa(expire),
			strconv.Itoa(g.db),
		},
	).Int()
	if err != nil {
		return 0, err
	}
	return result, nil
}

// check 检测本地时间和服务器时间的误差值是否小于1s
func (g *generator) check(ctx context.Context) error {
	res, _ := g.cli.Do(ctx, "time").Result()
	cts := res.([]interface{})[0].(string)
	ct, err := strconv.ParseInt(cts, 10, 64)
	if err != nil {
		return nil
	}
	lt := time.Now().Unix()
	var pt int64
	if ct > lt {
		pt = ct - lt
	} else {
		pt = lt - ct
	}
	fmt.Println("pt:", pt)
	if pt > 1 {
		return errors.New("check redis date faile")
	}
	return nil
}

func key(biz string) string {
	return biz + time.Now().Local().Format("20060102150304")
	// return biz + time.Now().Local().Format("20060102")
}

func code(key string, sequence int) string {
	return key + fmt.Sprintf("%07d", sequence)
}
