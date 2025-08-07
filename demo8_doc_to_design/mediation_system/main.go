// [AI GENERATED] LLM: GitHub Copilot, Mode: Agent, Date: 2025-08-06
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/abhilashmuraleedharan/copilotdemos/config"
	"github.com/abhilashmuraleedharan/copilotdemos/enrichment"
)

// CDRProcessor manages the CDR processing pipeline
type CDRProcessor struct {
	enricher enrichment.Enricher
	config   *config.Config
}

// NewCDRProcessor creates a new CDR processor with configured enrichers
func NewCDRProcessor(cfg *config.Config) *CDRProcessor {
	// Create enrichment cache with 5-minute TTL
	cache := enrichment.NewEnrichmentCache(5 * time.Minute)

	// Initialize individual enrichers
	var enrichers []enrichment.Enricher

	for _, source := range cfg.EnrichmentSources {
		switch source {
		case "subscriber":
			enrichers = append(enrichers, enrichment.NewSubscriberEnricher(cache))
		case "device":
			enrichers = append(enrichers, enrichment.NewDeviceEnricher(cache))
		case "geo":
			enrichers = append(enrichers, enrichment.NewGeoEnricher(cache))
		default:
			log.Printf("Unknown enrichment source: %s", source)
		}
	}

	// Create composite enricher
	compositeEnricher := enrichment.NewCompositeEnricher(enrichers, 5*time.Minute)

	return &CDRProcessor{
		enricher: compositeEnricher,
		config:   cfg,
	}
}

// ProcessCDR processes a single CDR through the enrichment pipeline
func (cp *CDRProcessor) ProcessCDR(cdrData []byte) ([]byte, error) {
	log.Printf("Processing CDR with enrichment sources: %v", cp.config.EnrichmentSources)

	// Enrich the CDR
	enrichedData, err := cp.enricher.Enrich(cdrData)
	if err != nil {
		return nil, fmt.Errorf("enrichment failed: %w", err)
	}

	log.Printf("CDR enrichment completed successfully")
	return enrichedData, nil
}

// ProcessBatch processes multiple CDRs
func (cp *CDRProcessor) ProcessBatch(cdrBatch [][]byte) ([][]byte, error) {
	var results [][]byte
	var errors []error

	for i, cdrData := range cdrBatch {
		enrichedData, err := cp.ProcessCDR(cdrData)
		if err != nil {
			log.Printf("Failed to process CDR %d: %v", i, err)
			errors = append(errors, err)
			continue
		}
		results = append(results, enrichedData)
	}

	if len(errors) > 0 {
		return results, fmt.Errorf("failed to process %d CDRs", len(errors))
	}

	return results, nil
}

func main() {
	fmt.Println("CDR Transformation Microservice starting...")

	// Initialize configuration
	cfg := &config.Config{
		RuleSet:           "default_v1",
		EnrichmentSources: []string{"subscriber", "device", "geo"},
		OutputFormats:     []string{"json", "csv"},
	}

	// Create CDR processor with enrichment pipeline
	processor := NewCDRProcessor(cfg)

	// Demo: Process sample CDRs
	sampleCDRs := createSampleCDRs()

	fmt.Printf("Processing %d sample CDRs...\n", len(sampleCDRs))

	for i, cdrData := range sampleCDRs {
		fmt.Printf("\n--- Processing CDR %d ---\n", i+1)

		// Show original CDR
		fmt.Println("Original CDR:")
		printFormattedJSON(cdrData)

		// Process CDR
		enrichedData, err := processor.ProcessCDR(cdrData)
		if err != nil {
			log.Printf("Failed to process CDR %d: %v", i+1, err)
			continue
		}

		// Show enriched CDR
		fmt.Println("\nEnriched CDR:")
		printFormattedJSON(enrichedData)
	}

	fmt.Println("\nCDR Transformation Microservice demo completed.")
}

// createSampleCDRs creates sample CDR data for demonstration
func createSampleCDRs() [][]byte {
	sampleCDRs := []enrichment.CDR{
		{
			ID:         "cdr_001",
			IMSI:       "001234567890123",
			MSISDN:     "+1234567890",
			IMEI:       "351234567890123",
			CellID:     "NYC001",
			Timestamp:  time.Now(),
			EventType:  "data_session",
			ServiceID:  "internet",
			Duration:   3600,
			DataVolume: 1024000,
		},
		{
			ID:         "cdr_002",
			IMSI:       "002345678901234",
			MSISDN:     "+2345678901",
			IMEI:       "862345678901234",
			CellID:     "LAX002",
			Timestamp:  time.Now(),
			EventType:  "voice_call",
			ServiceID:  "voice",
			Duration:   300,
			DataVolume: 0,
		},
		{
			ID:         "cdr_003",
			IMSI:       "003456789012345",
			MSISDN:     "+3456789012",
			IMEI:       "354567890123456",
			CellID:     "CHI003",
			Timestamp:  time.Now(),
			EventType:  "sms",
			ServiceID:  "messaging",
			Duration:   0,
			DataVolume: 160,
		},
	}

	var cdrDataList [][]byte
	for _, cdr := range sampleCDRs {
		data, _ := json.Marshal(cdr)
		cdrDataList = append(cdrDataList, data)
	}

	return cdrDataList
}

// printFormattedJSON prints JSON data in a formatted way
func printFormattedJSON(data []byte) {
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Printf("Raw data: %s\n", string(data))
		return
	}

	formatted, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Printf("Raw data: %s\n", string(data))
		return
	}

	fmt.Println(string(formatted))
}
