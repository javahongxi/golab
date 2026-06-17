package main

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

type Heartbeat struct {
	interval time.Duration
	count    int64
	stopCh   chan struct{}
}

func NewHeartbeat(interval time.Duration) *Heartbeat {
	return &Heartbeat{
		interval: interval,
		stopCh:   make(chan struct{}),
	}
}

func (h *Heartbeat) Start() {
	ticker := time.NewTicker(h.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			atomic.AddInt64(&h.count, 1)
			log.Printf("Heartbeat #%d at %v", h.count, time.Now())
		case <-h.stopCh:
			log.Println("Heartbeat stopped")
			return
		}
	}
}

func (h *Heartbeat) Stop() {
	close(h.stopCh)
}

func (h *Heartbeat) Count() int64 {
	return atomic.LoadInt64(&h.count)
}

type PeriodicCleaner struct {
	interval time.Duration
	callback func()
	stopCh   chan struct{}
}

func NewPeriodicCleaner(interval time.Duration, callback func()) *PeriodicCleaner {
	return &PeriodicCleaner{
		interval: interval,
		callback: callback,
		stopCh:   make(chan struct{}),
	}
}

func (p *PeriodicCleaner) Start() {
	ticker := time.NewTicker(p.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Running periodic cleanup...")
			p.callback()
		case <-p.stopCh:
			log.Println("PeriodicCleaner stopped")
			return
		}
	}
}

func (p *PeriodicCleaner) Stop() {
	close(p.stopCh)
}

func main() {
	fmt.Println("Starting timer examples...")

	heartbeat := NewHeartbeat(1 * time.Second)
	go heartbeat.Start()

	cache := make(map[string]string)
	cache["key1"] = "value1"
	cache["key2"] = "value2"

	cleaner := NewPeriodicCleaner(3*time.Second, func() {
		log.Printf("Cache size before cleanup: %d", len(cache))
		cache = make(map[string]string)
		log.Printf("Cache size after cleanup: %d", len(cache))
	})
	go cleaner.Start()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	<-ctx.Done()
	fmt.Println("\nContext cancelled, stopping all timers...")

	heartbeat.Stop()
	cleaner.Stop()

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Total heartbeats: %d\n", heartbeat.Count())
}
