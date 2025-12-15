// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package enrichment

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/demo/cdr-microservice/config"
)

type Service struct {
	config config.EnrichmentConfig
	cache  *Cache
}

type Cache struct {
	mu      sync.RWMutex
	data    map[string]CacheEntry
	maxSize int
	ttl     time.Duration
}

type CacheEntry struct {
	Value     interface{}
	ExpiresAt time.Time
}

func NewService(cfg config.EnrichmentConfig) *Service {
	return &Service{
		config: cfg,
		cache: &Cache{
			data:    make(map[string]CacheEntry),
			maxSize: cfg.CacheSize,
			ttl:     time.Duration(cfg.CacheTTL) * time.Second,
		},
	}
}

func (s *Service) Enrich(ctx context.Context, data []byte) ([]byte, error) {
	if !s.config.Enabled {
		return data, nil
	}

	var record map[string]interface{}
	if err := json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	for field, lookupPath := range s.config.CustomFields {
		if value, ok := s.lookup(lookupPath); ok {
			record[field] = value
		}
	}

	return json.Marshal(record)
}

func (s *Service) lookup(path string) (interface{}, bool) {
	s.cache.mu.RLock()
	defer s.cache.mu.RUnlock()

	if entry, ok := s.cache.data[path]; ok {
		if time.Now().Before(entry.ExpiresAt) {
			return entry.Value, true
		}
	}
	return nil, false
}

func (s *Service) cacheSet(key string, value interface{}) {
	s.cache.mu.Lock()
	defer s.cache.mu.Unlock()

	if len(s.cache.data) >= s.cache.maxSize {
		for k := range s.cache.data {
			delete(s.cache.data, k)
			break
		}
	}

	s.cache.data[key] = CacheEntry{
		Value:     value,
		ExpiresAt: time.Now().Add(s.cache.ttl),
	}
}
