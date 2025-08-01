package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type CleanupFunc func(ctx context.Context) error

type ShutdownManager struct {
	m        sync.Mutex
	cleanups []CleanupFunc
}

func NewShutdownManager() *ShutdownManager {
	return &ShutdownManager{
		cleanups: make([]CleanupFunc, 0),
	}
}

func (s *ShutdownManager) Register(fn CleanupFunc) {
	s.m.Lock()
	defer s.m.Unlock()
	s.cleanups = append(s.cleanups, fn)
}

func (s *ShutdownManager) WaitAndShutdown(serverStop func(ctx context.Context) error) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("\n🛑 Shutdown signal received...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if serverStop != nil {
		if err := serverStop(ctx); err != nil {
			log.Printf("❌ Server shutdown error: %v", err)
		}
	}

	s.m.Lock()
	defer s.m.Unlock()

	for _, fn := range s.cleanups {
		if err := fn(ctx); err != nil {
			log.Printf("⚠️ Cleanup failed: %v", err)
		}
	}

	fmt.Println("✅ Shutdown complete")
}
