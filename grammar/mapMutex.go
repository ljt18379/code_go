package main

import (
	"math/rand"
	"sync"
	"time"
)

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}

// 不加锁
func ConcurrentWrite() {
	m := make(map[int]int)
	f := func() {
		k, v := rand.Int()%50, rand.Int()%1000
		m[k] = v
	}
	for i := 0; i < 10; i++ {
		go f()
	}
	time.Sleep(time.Second * 5)
}

func ConcurrentWrite3() {
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		k, v := rand.Int()%50, rand.Int()%1000
		m[k] = v
	}
	time.Sleep(time.Second * 5)
}

// 加锁
func ConcurrentWrite2() {
	var mutex sync.Mutex
	m := make(map[int]int)

	f := func() {
		k, v := rand.Int()%50, rand.Int()%1000
		mutex.Lock()
		m[k] = v
		mutex.Unlock()
	}
	for i := 0; i < 10; i++ {
		go f()
	}
	time.Sleep(time.Second * 5)
}

func Maptest() {
	//var mutex sync.Mutex
	m := make(map[string]int)
	f := func() {
		v := rand.Int() % 50
		//mutex.Lock()
		m["module"] = v
		m["user"] = v
		//mutex.Unlock()
	}
	for i := 0; i < 10; i++ {
		go f()
	}
	time.Sleep(time.Second * 5)
}
