// [AI GENERATED] LLM: GitHub Copilot, Mode: Agent, Date: 2025-08-06
package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/abhilashmuraleedharan/copilotdemos/config"
	"github.com/abhilashmuraleedharan/copilotdemos/enrichment"
)

func TestCDRProcessorCreation(t *testing.T) {
	cfg := &config.Config{
		RuleSet:           "test_v1",
		EnrichmentSources: []string{"subscriber", "device", "geo"},
		OutputFormats:     []string{"json"},
	}

	processor := NewCDRProcessor(cfg)
	if processor == nil {
		t.Fatal("Expected CDRProcessor to be created")
	}

	if processor.config != cfg {
		t.Error("Expected processor to use provided config")
	}

	if processor.enricher == nil {
		t.Error("Expected processor to have enricher configured")
	}
}

func TestCDRProcessorWithDifferentEnrichmentSources(t *testing.T) {
	tests := []struct {
		name    string
		sources []string
	}{
		{"subscriber only", []string{"subscriber"}},
		{"device only", []string{"device"}},
		{"geo only", []string{"geo"}},
		{"subscriber and device", []string{"subscriber", "device"}},
		{"all sources", []string{"subscriber", "device", "geo"}},
		{"empty sources", []string{}},
		{"unknown source", []string{"unknown"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config.Config{
				RuleSet:           "test_v1",
				EnrichmentSources: tt.sources,
				OutputFormats:     []string{"json"},
			}

			processor := NewCDRProcessor(cfg)
			if processor == nil {
				t.Fatal("Expected CDRProcessor to be created")
			}

			// Test processing a sample CDR
			testCDR := enrichment.CDR{
				ID:         "test_001",
				IMSI:       "001234567890123",
				IMEI:       "351234567890123",
				CellID:     "NYC001",
				Timestamp:  time.Now(),
				EventType:  "data_session",
				ServiceID:  "internet",
				Duration:   3600,
				DataVolume: 1024000,
			}

			data, err := json.Marshal(testCDR)
			if err != nil {
				t.Fatalf("Failed to marshal test CDR: %v", err)
			}

			_, err = processor.ProcessCDR(data)
			if err != nil {
				t.Errorf("Failed to process CDR with sources %v: %v", tt.sources, err)
			}
		})
	}
}

func TestProcessSingleCDR(t *testing.T) {
	cfg := &config.Config{
		RuleSet:           "test_v1",
		EnrichmentSources: []string{"subscriber", "device", "geo"},
		OutputFormats:     []string{"json"},
	}

	processor := NewCDRProcessor(cfg)

	// Create test CDR
	testCDR := enrichment.CDR{
		ID:         "test_001",
		IMSI:       "001234567890123", // Premium pattern
		IMEI:       "351234567890123", // Smartphone pattern
		CellID:     "NYC001",          // NYC pattern
		Timestamp:  time.Now(),
		EventType:  "data_session",
		ServiceID:  "internet",
		Duration:   3600,
		DataVolume: 1024000,
	}

	data, err := json.Marshal(testCDR)
	if err != nil {
		t.Fatalf("Failed to marshal test CDR: %v", err)
	}

	enrichedData, err := processor.ProcessCDR(data)
	if err != nil {
		t.Fatalf("Failed to process CDR: %v", err)
	}

	// Verify enriched data
	var enrichedCDR enrichment.CDR
	if err := json.Unmarshal(enrichedData, &enrichedCDR); err != nil {
		t.Fatalf("Failed to unmarshal enriched CDR: %v", err)
	}

	// Verify enrichments were applied
	if enrichedCDR.SubscriberType != "premium" {
		t.Errorf("Expected subscriber type premium, got %s", enrichedCDR.SubscriberType)
	}
	if enrichedCDR.DeviceType != "smartphone" {
		t.Errorf("Expected device type smartphone, got %s", enrichedCDR.DeviceType)
	}
	if enrichedCDR.LocationRegion != "New York" {
		t.Errorf("Expected region New York, got %s", enrichedCDR.LocationRegion)
	}

	// Verify enrichment status
	if enrichedCDR.EnrichmentStatus == nil {
		t.Error("Expected enrichment status to be populated")
	}
}

func TestProcessBatch(t *testing.T) {
	cfg := &config.Config{
		RuleSet:           "test_v1",
		EnrichmentSources: []string{"subscriber", "device", "geo"},
		OutputFormats:     []string{"json"},
	}

	processor := NewCDRProcessor(cfg)

	// Create test CDR batch
	testCDRs := []enrichment.CDR{
		{
			ID:         "batch_001",
			IMSI:       "001234567890123",
			IMEI:       "351234567890123",
			CellID:     "NYC001",
			Timestamp:  time.Now(),
			EventType:  "data_session",
			ServiceID:  "internet",
			Duration:   3600,
			DataVolume: 1024000,
		},
		{
			ID:         "batch_002",
			IMSI:       "002345678901234",
			IMEI:       "862345678901234",
			CellID:     "LAX002",
			Timestamp:  time.Now(),
			EventType:  "voice_call",
			ServiceID:  "voice",
			Duration:   300,
			DataVolume: 0,
		},
	}

	var cdrBatch [][]byte
	for _, cdr := range testCDRs {
		data, err := json.Marshal(cdr)
		if err != nil {
			t.Fatalf("Failed to marshal test CDR: %v", err)
		}
		cdrBatch = append(cdrBatch, data)
	}

	results, err := processor.ProcessBatch(cdrBatch)
	if err != nil {
		t.Fatalf("Failed to process CDR batch: %v", err)
	}

	if len(results) != len(cdrBatch) {
		t.Errorf("Expected %d results, got %d", len(cdrBatch), len(results))
	}

	// Verify each result
	for i, result := range results {
		var enrichedCDR enrichment.CDR
		if err := json.Unmarshal(result, &enrichedCDR); err != nil {
			t.Errorf("Failed to unmarshal result %d: %v", i, err)
			continue
		}

		// Verify basic enrichment was applied
		if enrichedCDR.EnrichmentStatus == nil {
			t.Errorf("Expected enrichment status for CDR %d", i)
		}
	}
}

func TestProcessBatchWithErrors(t *testing.T) {
	cfg := &config.Config{
		RuleSet:           "test_v1",
		EnrichmentSources: []string{"subscriber"},
		OutputFormats:     []string{"json"},
	}

	processor := NewCDRProcessor(cfg)

	// Create batch with valid and invalid CDRs
	cdrBatch := [][]byte{
		[]byte(`{"id":"valid_001","imsi":"001234567890123"}`),
		[]byte(`{"invalid": json}`), // Invalid JSON
		[]byte(`{"id":"valid_002","imsi":"002345678901234"}`),
	}

	results, err := processor.ProcessBatch(cdrBatch)

	// Should return partial results and an error
	if err == nil {
		t.Error("Expected error when processing batch with invalid CDRs")
	}

	// Should have 2 valid results (skipping the invalid one)
	if len(results) != 2 {
		t.Errorf("Expected 2 valid results, got %d", len(results))
	}
}

func TestCreateSampleCDRs(t *testing.T) {
	sampleCDRs := createSampleCDRs()

	if len(sampleCDRs) == 0 {
		t.Error("Expected sample CDRs to be created")
	}

	expectedCount := 3
	if len(sampleCDRs) != expectedCount {
		t.Errorf("Expected %d sample CDRs, got %d", expectedCount, len(sampleCDRs))
	}

	// Verify each sample CDR can be unmarshaled
	for i, cdrData := range sampleCDRs {
		var cdr enrichment.CDR
		if err := json.Unmarshal(cdrData, &cdr); err != nil {
			t.Errorf("Failed to unmarshal sample CDR %d: %v", i, err)
		}

		// Verify required fields are present
		if cdr.ID == "" {
			t.Errorf("Sample CDR %d missing ID", i)
		}
		if cdr.IMSI == "" {
			t.Errorf("Sample CDR %d missing IMSI", i)
		}
		if cdr.EventType == "" {
			t.Errorf("Sample CDR %d missing EventType", i)
		}
	}
}

func TestPrintFormattedJSON(t *testing.T) {
	// Test with valid JSON
	testData := map[string]interface{}{
		"id":         "test_001",
		"imsi":       "001234567890123",
		"event_type": "data_session",
	}

	validJSON, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("Failed to create test JSON: %v", err)
	}

	// This should not panic
	printFormattedJSON(validJSON)

	// Test with invalid JSON
	invalidJSON := []byte(`{"invalid": json}`)

	// This should also not panic and fall back to raw output
	printFormattedJSON(invalidJSON)
}

func BenchmarkCDRProcessing(b *testing.B) {
	cfg := &config.Config{
		RuleSet:           "benchmark_v1",
		EnrichmentSources: []string{"subscriber", "device", "geo"},
		OutputFormats:     []string{"json"},
	}

	processor := NewCDRProcessor(cfg)

	testCDR := enrichment.CDR{
		ID:         "bench_001",
		IMSI:       "001234567890123",
		IMEI:       "351234567890123",
		CellID:     "NYC001",
		Timestamp:  time.Now(),
		EventType:  "data_session",
		ServiceID:  "internet",
		Duration:   3600,
		DataVolume: 1024000,
	}

	data, err := json.Marshal(testCDR)
	if err != nil {
		b.Fatalf("Failed to marshal test CDR: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := processor.ProcessCDR(data)
		if err != nil {
			b.Fatalf("Failed to process CDR: %v", err)
		}
	}
}

func BenchmarkBatchProcessing(b *testing.B) {
	cfg := &config.Config{
		RuleSet:           "benchmark_v1",
		EnrichmentSources: []string{"subscriber", "device", "geo"},
		OutputFormats:     []string{"json"},
	}

	processor := NewCDRProcessor(cfg)

	// Create a batch of 100 CDRs
	var cdrBatch [][]byte
	for i := 0; i < 100; i++ {
		testCDR := enrichment.CDR{
			ID:         fmt.Sprintf("bench_%03d", i),
			IMSI:       "001234567890123",
			IMEI:       "351234567890123",
			CellID:     "NYC001",
			Timestamp:  time.Now(),
			EventType:  "data_session",
			ServiceID:  "internet",
			Duration:   3600,
			DataVolume: 1024000,
		}

		data, err := json.Marshal(testCDR)
		if err != nil {
			b.Fatalf("Failed to marshal test CDR: %v", err)
		}
		cdrBatch = append(cdrBatch, data)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := processor.ProcessBatch(cdrBatch)
		if err != nil {
			b.Fatalf("Failed to process CDR batch: %v", err)
		}
	}
}
