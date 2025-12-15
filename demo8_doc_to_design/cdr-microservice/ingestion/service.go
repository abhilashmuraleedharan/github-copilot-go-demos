// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package ingestion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/demo/cdr-microservice/config"
	"github.com/gorilla/mux"
)

type ProcessFunc func(context.Context, []byte) error

type Service struct {
	config  config.IngestionConfig
	process ProcessFunc
	server  *http.Server
	buffer  chan []byte
}

func NewService(cfg config.IngestionConfig, process ProcessFunc) *Service {
	return &Service{
		config:  cfg,
		process: process,
		buffer:  make(chan []byte, cfg.BufferSize),
	}
}

func (s *Service) Start(ctx context.Context) error {
	router := mux.NewRouter()
	router.HandleFunc("/cdr", s.handleCDR).Methods("POST")
	router.HandleFunc("/health", s.handleHealth).Methods("GET")

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	s.server = &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go s.processBuffer(ctx)

	go func() {
		log.Printf("Starting ingestion service on %s", addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ingestion service error: %v", err)
		}
	}()

	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	close(s.buffer)
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

func (s *Service) handleCDR(w http.ResponseWriter, r *http.Request) {
	var data json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	select {
	case s.buffer <- data:
		w.WriteHeader(http.StatusAccepted)
	default:
		http.Error(w, "Buffer full", http.StatusServiceUnavailable)
	}
}

func (s *Service) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func (s *Service) processBuffer(ctx context.Context) {
	batch := make([][]byte, 0, s.config.BatchSize)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case data, ok := <-s.buffer:
			if !ok {
				s.processBatch(ctx, batch)
				return
			}
			batch = append(batch, data)
			if len(batch) >= s.config.BatchSize {
				s.processBatch(ctx, batch)
				batch = batch[:0]
			}
		case <-ticker.C:
			if len(batch) > 0 {
				s.processBatch(ctx, batch)
				batch = batch[:0]
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *Service) processBatch(ctx context.Context, batch [][]byte) {
	for _, data := range batch {
		if err := s.process(ctx, data); err != nil {
			log.Printf("Error processing CDR: %v", err)
		}
	}
}
