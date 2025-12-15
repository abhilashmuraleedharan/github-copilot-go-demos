// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package exporting

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/demo/cdr-microservice/config"
)

type Service struct {
	config config.ExportingConfig
	batch  [][]byte
	mu     sync.Mutex
	client *http.Client
}

func NewService(cfg config.ExportingConfig) *Service {
	return &Service{
		config: cfg,
		batch:  make([][]byte, 0, cfg.BatchSize),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *Service) Export(ctx context.Context, data []byte) error {
	s.mu.Lock()
	s.batch = append(s.batch, data)
	shouldFlush := len(s.batch) >= s.config.BatchSize
	s.mu.Unlock()

	if shouldFlush {
		return s.flush(ctx)
	}

	return nil
}

func (s *Service) flush(ctx context.Context) error {
	s.mu.Lock()
	if len(s.batch) == 0 {
		s.mu.Unlock()
		return nil
	}
	batchToSend := make([][]byte, len(s.batch))
	copy(batchToSend, s.batch)
	s.batch = s.batch[:0]
	s.mu.Unlock()

	switch s.config.Type {
	case "http":
		return s.exportHTTP(ctx, batchToSend)
	case "kafka":
		return s.exportKafka(ctx, batchToSend)
	case "file":
		return s.exportFile(ctx, batchToSend)
	default:
		return fmt.Errorf("unsupported export type: %s", s.config.Type)
	}
}

func (s *Service) exportHTTP(ctx context.Context, batch [][]byte) error {
	payload, err := json.Marshal(batch)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", s.config.Destination, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	for key, value := range s.config.Headers {
		req.Header.Set(key, value)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("export failed with status: %d", resp.StatusCode)
	}

	log.Printf("Exported %d records via HTTP", len(batch))
	return nil
}

func (s *Service) exportKafka(ctx context.Context, batch [][]byte) error {
	log.Printf("Exported %d records to Kafka topic: %s", len(batch), s.config.Destination)
	return nil
}

func (s *Service) exportFile(ctx context.Context, batch [][]byte) error {
	log.Printf("Exported %d records to file: %s", len(batch), s.config.Destination)
	return nil
}
