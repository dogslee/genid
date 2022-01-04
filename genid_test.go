package genid

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func Test_New(t *testing.T) {

	var err error

	_, err = New()
	if err != nil {
		t.Error(err)
	}

	cli := redis.NewClient(&redis.Options{Addr: "127.1.1.1:6379", DialTimeout: time.Second})
	_, err = New(
		DB(10),
		Cli(cli),
		Exp(3),
	)
	if err == nil {
		fmt.Println(err)
	}
}

func Test_Create(t *testing.T) {
	g, _ := New()
	code, err := g.Create(context.TODO(), "1001")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("genid code:", code)
}

func Test_CreateBatch(t *testing.T) {
	var n int = 5
	g, _ := New()
	codes, err := g.CreateBatch(context.TODO(), "1001", n)
	if err != nil {
		t.Error(err)
	}
	if len(codes) != n {
		t.Errorf("CreateBatch code error")
	}
	fmt.Println("genid codes:", codes)
}

func Benchmark_Create(b *testing.B) {
	var wg sync.WaitGroup
	g, _ := New()
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			code, err := g.Create(context.Background(), "1000")
			if err != nil {
				b.Error("x", err)
			}
			fmt.Println("bench create code:", code)
		}()
	}
}

func Benchmark_CreateBatch(b *testing.B) {
	var wg sync.WaitGroup
	g, _ := New()
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			codes, err := g.CreateBatch(context.Background(), "2000", 5)
			if err != nil {
				b.Error("x", err)
			}
			fmt.Println("bench create batch codes:", codes)
		}()
	}
}
