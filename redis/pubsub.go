package main

import (
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type Mgr struct {
	redis  *redis.Client
	pubsub *redis.PubSub
	wg     sync.WaitGroup
}

func NewMgr(redis *redis.Client) *Mgr {
	return &Mgr{
		redis: redis,
	}
}

func (m *Mgr) Start() {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		m.pubsub = m.redis.Subscribe("pubsub")
		for msg := range m.pubsub.Channel() {
			log.Printf("Received message: %s", msg.Payload)
		}
		log.Printf("finished reading channel")
	}()
}

func (m *Mgr) Stop() {
	_ = m.pubsub.Close()
	m.wg.Wait()
	log.Printf("graceful shutdown complete")
}

func main() {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	mgr := NewMgr(cli)
	mgr.Start()
	time.Sleep(time.Second * 5)
	mgr.Stop()
}
