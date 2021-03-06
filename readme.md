# genid

[![Go Refrence Card](https://pkg.go.dev/badge/github.com/dogslee/geni)](https://pkg.go.dev/github.com/dogslee/genid)
[![Go Passing Card](https://img.shields.io/badge/go-passing-success)](https://github.com/dogslee/genid/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dogslee/genid)](https://goreportcard.com/report/github.com/dogslee/genid)
[![Go Codecov Card](https://codecov.io/gh/dogslee/genid/main/graph/badge.svg)](https://codecov.io/gh/dogslee/genid/)
[![Go License Card](./img/license-mit.svg)](./LICENSE)

一个简单的24位 分布式自增的ID生成器

## 业务码生成规则如下

业务ID(4位)+时间戳(14位)+自增序列(8位)

如下图:
![alt ](./img/img.png)

其中inc id  是通过redis 中的原子自增实现的，理论上1s最多可以生成 99999999 个不同的ID

当然这依赖于redis的读写能力，超高并发场景下可将redis做集群化

## 快速开始

```go
package main

import (
	"context"
	"fmt"

	"github.com/dogslee/genid"
	"github.com/go-redis/redis/v8"
)

func main() {
	cli := redis.NewClient(
		&redis.Options{
			Addr: "127.0.0.1:16379",
		},
	)
	g, err := genid.New(
		genid.DB(15),
		genid.Cli(cli),
	)
	if err != nil {
		panic(err)
	}
	code, _ := g.Create(context.Background(), "1000")
	fmt.Println(code)
}
```