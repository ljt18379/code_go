package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg2 = sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go moniter(ctx, i)
	}
	time.Sleep(time.Millisecond * 5)
	cancel()
	wg2.Wait()
	//time.Sleep(time.Second*1)
	fmt.Println("监控结束，退出主程序")

}

func moniter(ctx context.Context, i int) {
	defer wg2.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("进程%v收到取消信号！！！", i)
			return
		default:
			fmt.Printf("进程%v正在持续监控！！！", i)
			time.Sleep(time.Millisecond * 1)
		}
	}
}
