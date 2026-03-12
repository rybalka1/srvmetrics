// internal/storage/storage.go
package storage

import "sync"

type MemStorage struct {
	metrics map[string]any
	mu      sync.Mutex
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]any),
	}
}

func (s *MemStorage) GetMetric(name string) (any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.metrics[name]
	return value, ok
}

func (s *MemStorage) SetGauge(name string, value float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.metrics[name] = value
}

func (s *MemStorage) AddCounter(name string, value int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if existing, ok := s.metrics[name].(int64); ok {
		s.metrics[name] = existing + value
	} else {
		s.metrics[name] = value
	}
}
