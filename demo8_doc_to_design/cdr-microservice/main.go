// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/demo/cdr-microservice/config"
	"github.com/demo/cdr-microservice/enrichment"
	"github.com/demo/cdr-microservice/exporting"
	"github.com/demo/cdr-microservice/filtering"
	"github.com/demo/cdr-microservice/ingestion"
	"github.com/demo/cdr-microservice/transformation"
)

const (
	shutdownTimeout = 30 * time.Second
)

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	enrichmentSvc := enrichment.NewService(cfg.Enrichment)
	transformationSvc := transformation.NewService(cfg.Transformation)
	filteringSvc := filtering.NewService(cfg.Filtering)
	exportingSvc := exporting.NewService(cfg.Exporting)

	pipeline := &Pipeline{
		enrichment:     enrichmentSvc,
		transformation: transformationSvc,
		filtering:      filteringSvc,
		exporting:      exportingSvc,
	}

	ingestionSvc := ingestion.NewService(cfg.Ingestion, pipeline.Process)

	if err := ingestionSvc.Start(ctx); err != nil {
		log.Fatalf("Failed to start ingestion service: %v", err)
	}

	log.Println("CDR Transformation Microservice started successfully")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down gracefully...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownCancel()

	if err := ingestionSvc.Stop(shutdownCtx); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}

type Pipeline struct {
	enrichment     *enrichment.Service
	transformation *transformation.Service
	filtering      *filtering.Service
	exporting      *exporting.Service
}

func (p *Pipeline) Process(ctx context.Context, data []byte) error {
	enriched, err := p.enrichment.Enrich(ctx, data)
	if err != nil {
		return err
	}

	transformed, err := p.transformation.Transform(ctx, enriched)
	if err != nil {
		return err
	}

	if !p.filtering.ShouldProcess(ctx, transformed) {
		return nil
	}

	return p.exporting.Export(ctx, transformed)
}
