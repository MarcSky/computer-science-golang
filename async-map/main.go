package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	taskSize = 5000
	maxRand  = 1000000
)

type ConcurrentMap struct {
	write      chan MapValue
	request    chan int
	readAnswer chan int
	data       map[int]int
	stop       chan struct{}
}

type MapValue struct {
	Key   int
	Value int
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		write:      make(chan MapValue),
		request:    make(chan int),
		readAnswer: make(chan int),
		data:       make(map[int]int),
		stop:       make(chan struct{}),
	}
}

func (c *ConcurrentMap) Start() {
	for {
		select {
		case v := <-c.write:
			c.data[v.Key] = v.Value
		case key := <-c.request:
			c.readAnswer <- c.data[key]
		case <-c.stop:
			return
		}
	}
}

func (c *ConcurrentMap) Close() {
	c.stop <- struct{}{}
	close(c.write)
	close(c.request)
	close(c.readAnswer)
	close(c.stop)
}

func (c *ConcurrentMap) Set(w MapValue) {
	c.write <- w
}

func (c *ConcurrentMap) Read(key int) int {
	c.request <- key
	for v := range c.readAnswer {
		return v
	}
	return 0
}

func Task(v int, m *ConcurrentMap, wg *sync.WaitGroup) {
	key := rand.Intn(maxRand)
	value := rand.Intn(maxRand)
	m.Set(MapValue{
		Key:   key,
		Value: value,
	})
	res := m.Read(key)
	if res != value {
		fmt.Println("sad value")
	} else {
		fmt.Println(v, "success")
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().Unix())

	m := NewConcurrentMap()
	go m.Start()

	wg := &sync.WaitGroup{}
	wg.Add(taskSize)
	for i := 0; i < taskSize; i++ {
		go Task(i, m, wg)
	}
	wg.Wait()

	m.Close()
}
