// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
package main

import (
	"log"
	"os"

	"demo8_doc_to_design/config"
	"demo8_doc_to_design/ingestion"
	"demo8_doc_to_design/enrichment"
	"demo8_doc_to_design/transformation"
	"demo8_doc_to_design/filtering"
	"demo8_doc_to_design/exporting"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
		os.Exit(1)
	}

	// Initialize pipeline components
	ingestor := ingestion.NewIngestor(cfg.Ingestion)
	enricher := enrichment.NewEnricher(cfg.Enrichment)
	transformer := transformation.NewTransformer(cfg.Transformation)
	filter := filtering.NewFilter(cfg.Filtering)
	exporter := exporting.NewExporter(cfg.Exporting)

	// Start pipeline (simplified)
	for cdr := range ingestor.Ingest() {
		enriched := enricher.Enrich(cdr)
		transformed := transformer.Transform(enriched)
		if filter.Filter(transformed) {
			exporter.Export(transformed)
		}
	}
}
