package genid

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func Test_New(t *testing.T) {
	g, _ := New()
	var wg sync.WaitGroup
	var n int = 100
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			code, err := g.Create(context.Background(), "1000")
			if err != nil {
				t.Error("x", err)
			}
			fmt.Println(code)
		}()
	}
	wg.Wait()
}
