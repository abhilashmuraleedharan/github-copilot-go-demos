// [AI GENERATED] LLM: GitHub Copilot, Mode: Agent, Date: 2025-08-06
package enrichment

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCDRSerialization(t *testing.T) {
	// Test CDR JSON marshaling and unmarshaling
	originalCDR := CDR{
		ID:         "test_001",
		IMSI:       "001234567890123",
		MSISDN:     "+1234567890",
		IMEI:       "351234567890123",
		CellID:     "NYC001",
		Timestamp:  time.Now(),
		EventType:  "data_session",
		ServiceID:  "internet",
		Duration:   3600,
		DataVolume: 1024000,
	}

	// Marshal to JSON
	data, err := json.Marshal(originalCDR)
	if err != nil {
		t.Fatalf("Failed to marshal CDR: %v", err)
	}

	// Unmarshal back to CDR
	var reconstructedCDR CDR
	if err := json.Unmarshal(data, &reconstructedCDR); err != nil {
		t.Fatalf("Failed to unmarshal CDR: %v", err)
	}

	// Verify key fields
	if originalCDR.ID != reconstructedCDR.ID {
		t.Errorf("Expected ID %s, got %s", originalCDR.ID, reconstructedCDR.ID)
	}
	if originalCDR.IMSI != reconstructedCDR.IMSI {
		t.Errorf("Expected IMSI %s, got %s", originalCDR.IMSI, reconstructedCDR.IMSI)
	}
}

func TestEnrichmentCache(t *testing.T) {
	cache := NewEnrichmentCache(5 * time.Minute)

	// Test setting and getting values
	testKey := "test_key"
	testValue := "test_value"

	cache.Set(testKey, testValue)

	value, exists := cache.Get(testKey)
	if !exists {
		t.Error("Expected value to exist in cache")
	}

	if strValue, ok := value.(string); !ok || strValue != testValue {
		t.Errorf("Expected %s, got %v", testValue, value)
	}

	// Test non-existent key
	_, exists = cache.Get("non_existent_key")
	if exists {
		t.Error("Expected non-existent key to return false")
	}
}

func TestSubscriberEnricher(t *testing.T) {
	cache := NewEnrichmentCache(5 * time.Minute)
	enricher := NewSubscriberEnricher(cache)

	// Test CDR with premium IMSI pattern
	premiumCDR := CDR{
		ID:   "test_001",
		IMSI: "001234567890123",
	}

	data, err := json.Marshal(premiumCDR)
	if err != nil {
		t.Fatalf("Failed to marshal test CDR: %v", err)
	}

	enrichedData, err := enricher.Enrich(data)
	if err != nil {
		t.Fatalf("Enrichment failed: %v", err)
	}

	var enrichedCDR CDR
	if err := json.Unmarshal(enrichedData, &enrichedCDR); err != nil {
		t.Fatalf("Failed to unmarshal enriched CDR: %v", err)
	}

	// Verify premium subscriber enrichment
	expectedType := "premium"
	expectedPlan := "unlimited_5g"
	expectedAccount := "postpaid"

	if enrichedCDR.SubscriberType != expectedType {
		t.Errorf("Expected subscriber type %s, got %s", expectedType, enrichedCDR.SubscriberType)
	}
	if enrichedCDR.ServicePlan != expectedPlan {
		t.Errorf("Expected service plan %s, got %s", expectedPlan, enrichedCDR.ServicePlan)
	}
	if enrichedCDR.AccountType != expectedAccount {
		t.Errorf("Expected account type %s, got %s", expectedAccount, enrichedCDR.AccountType)
	}

	// Test caching by enriching the same CDR again
	enrichedData2, err := enricher.Enrich(data)
	if err != nil {
		t.Fatalf("Second enrichment failed: %v", err)
	}

	// Results should be the same
	var enrichedCDR2 CDR
	if err := json.Unmarshal(enrichedData2, &enrichedCDR2); err != nil {
		t.Fatalf("Failed to unmarshal second enriched CDR: %v", err)
	}

	if enrichedCDR2.SubscriberType != expectedType {
		t.Errorf("Cached enrichment failed: expected %s, got %s", expectedType, enrichedCDR2.SubscriberType)
	}
}

func TestDeviceEnricher(t *testing.T) {
	cache := NewEnrichmentCache(5 * time.Minute)
	enricher := NewDeviceEnricher(cache)

	// Test CDR with smartphone IMEI pattern
	smartphoneCDR := CDR{
		ID:   "test_001",
		IMEI: "351234567890123",
	}

	data, err := json.Marshal(smartphoneCDR)
	if err != nil {
		t.Fatalf("Failed to marshal test CDR: %v", err)
	}

	enrichedData, err := enricher.Enrich(data)
	if err != nil {
		t.Fatalf("Enrichment failed: %v", err)
	}

	var enrichedCDR CDR
	if err := json.Unmarshal(enrichedData, &enrichedCDR); err != nil {
		t.Fatalf("Failed to unmarshal enriched CDR: %v", err)
	}

	expectedDeviceType := "smartphone"
	if enrichedCDR.DeviceType != expectedDeviceType {
		t.Errorf("Expected device type %s, got %s", expectedDeviceType, enrichedCDR.DeviceType)
	}

	// Test IoT device IMEI pattern
	iotCDR := CDR{
		ID:   "test_002",
		IMEI: "861234567890123",
	}

	iotData, _ := json.Marshal(iotCDR)
	enrichedIoTData, err := enricher.Enrich(iotData)
	if err != nil {
		t.Fatalf("IoT enrichment failed: %v", err)
	}

	var enrichedIoTCDR CDR
	if err := json.Unmarshal(enrichedIoTData, &enrichedIoTCDR); err != nil {
		t.Fatalf("Failed to unmarshal enriched IoT CDR: %v", err)
	}

	expectedIoTType := "iot_device"
	if enrichedIoTCDR.DeviceType != expectedIoTType {
		t.Errorf("Expected IoT device type %s, got %s", expectedIoTType, enrichedIoTCDR.DeviceType)
	}
}

func TestGeoEnricher(t *testing.T) {
	cache := NewEnrichmentCache(5 * time.Minute)
	enricher := NewGeoEnricher(cache)

	// Test CDR with NYC cell ID pattern
	nycCDR := CDR{
		ID:     "test_001",
		CellID: "NYC001",
	}

	data, err := json.Marshal(nycCDR)
	if err != nil {
		t.Fatalf("Failed to marshal test CDR: %v", err)
	}

	enrichedData, err := enricher.Enrich(data)
	if err != nil {
		t.Fatalf("Enrichment failed: %v", err)
	}

	var enrichedCDR CDR
	if err := json.Unmarshal(enrichedData, &enrichedCDR); err != nil {
		t.Fatalf("Failed to unmarshal enriched CDR: %v", err)
	}

	expectedRegion := "New York"
	expectedLatitude := 40.7128
	expectedLongitude := -74.0060

	if enrichedCDR.LocationRegion != expectedRegion {
		t.Errorf("Expected region %s, got %s", expectedRegion, enrichedCDR.LocationRegion)
	}
	if enrichedCDR.Latitude != expectedLatitude {
		t.Errorf("Expected latitude %f, got %f", expectedLatitude, enrichedCDR.Latitude)
	}
	if enrichedCDR.Longitude != expectedLongitude {
		t.Errorf("Expected longitude %f, got %f", expectedLongitude, enrichedCDR.Longitude)
	}
}

func TestCompositeEnricher(t *testing.T) {
	cache := NewEnrichmentCache(5 * time.Minute)

	// Create individual enrichers
	subscriberEnricher := NewSubscriberEnricher(cache)
	deviceEnricher := NewDeviceEnricher(cache)
	geoEnricher := NewGeoEnricher(cache)

	enrichers := []Enricher{subscriberEnricher, deviceEnricher, geoEnricher}
	compositeEnricher := NewCompositeEnricher(enrichers, 5*time.Minute)

	// Test CDR that should be enriched by all enrichers
	testCDR := CDR{
		ID:     "test_001",
		IMSI:   "001234567890123", // Premium pattern
		IMEI:   "351234567890123", // Smartphone pattern
		CellID: "NYC001",          // NYC pattern
	}

	data, err := json.Marshal(testCDR)
	if err != nil {
		t.Fatalf("Failed to marshal test CDR: %v", err)
	}

	enrichedData, err := compositeEnricher.Enrich(data)
	if err != nil {
		t.Fatalf("Composite enrichment failed: %v", err)
	}

	var enrichedCDR CDR
	if err := json.Unmarshal(enrichedData, &enrichedCDR); err != nil {
		t.Fatalf("Failed to unmarshal enriched CDR: %v", err)
	}

	// Verify all enrichments were applied
	if enrichedCDR.SubscriberType != "premium" {
		t.Errorf("Expected subscriber type premium, got %s", enrichedCDR.SubscriberType)
	}
	if enrichedCDR.DeviceType != "smartphone" {
		t.Errorf("Expected device type smartphone, got %s", enrichedCDR.DeviceType)
	}
	if enrichedCDR.LocationRegion != "New York" {
		t.Errorf("Expected region New York, got %s", enrichedCDR.LocationRegion)
	}

	// Verify enrichment status tracking
	if enrichedCDR.EnrichmentStatus == nil {
		t.Error("Expected enrichment status to be populated")
	} else {
		expectedEnrichers := []string{"SubscriberEnricher", "DeviceEnricher", "GeoEnricher"}
		for _, enricherName := range expectedEnrichers {
			if status, exists := enrichedCDR.EnrichmentStatus[enricherName]; !exists || status != "success" {
				t.Errorf("Expected enricher %s to have success status, got %s", enricherName, status)
			}
		}
	}
}

func TestEnricherNames(t *testing.T) {
	cache := NewEnrichmentCache(5 * time.Minute)

	// Test individual enricher names
	subscriberEnricher := NewSubscriberEnricher(cache)
	if subscriberEnricher.GetName() != "SubscriberEnricher" {
		t.Errorf("Expected SubscriberEnricher name, got %s", subscriberEnricher.GetName())
	}

	deviceEnricher := NewDeviceEnricher(cache)
	if deviceEnricher.GetName() != "DeviceEnricher" {
		t.Errorf("Expected DeviceEnricher name, got %s", deviceEnricher.GetName())
	}

	geoEnricher := NewGeoEnricher(cache)
	if geoEnricher.GetName() != "GeoEnricher" {
		t.Errorf("Expected GeoEnricher name, got %s", geoEnricher.GetName())
	}

	// Test composite enricher name
	enrichers := []Enricher{subscriberEnricher, deviceEnricher, geoEnricher}
	compositeEnricher := NewCompositeEnricher(enrichers, 5*time.Minute)
	if compositeEnricher.GetName() != "CompositeEnricher" {
		t.Errorf("Expected CompositeEnricher name, got %s", compositeEnricher.GetName())
	}
}

func TestInvalidJSONHandling(t *testing.T) {
	cache := NewEnrichmentCache(5 * time.Minute)
	enricher := NewSubscriberEnricher(cache)

	// Test with invalid JSON
	invalidJSON := []byte(`{"invalid": json}`)

	_, err := enricher.Enrich(invalidJSON)
	if err == nil {
		t.Error("Expected error when processing invalid JSON")
	}
}

func BenchmarkSubscriberEnrichment(b *testing.B) {
	cache := NewEnrichmentCache(5 * time.Minute)
	enricher := NewSubscriberEnricher(cache)

	testCDR := CDR{
		ID:   "bench_001",
		IMSI: "001234567890123",
	}

	data, _ := json.Marshal(testCDR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := enricher.Enrich(data)
		if err != nil {
			b.Fatalf("Enrichment failed: %v", err)
		}
	}
}

func BenchmarkCompositeEnrichment(b *testing.B) {
	cache := NewEnrichmentCache(5 * time.Minute)

	subscriberEnricher := NewSubscriberEnricher(cache)
	deviceEnricher := NewDeviceEnricher(cache)
	geoEnricher := NewGeoEnricher(cache)

	enrichers := []Enricher{subscriberEnricher, deviceEnricher, geoEnricher}
	compositeEnricher := NewCompositeEnricher(enrichers, 5*time.Minute)

	testCDR := CDR{
		ID:     "bench_001",
		IMSI:   "001234567890123",
		IMEI:   "351234567890123",
		CellID: "NYC001",
	}

	data, _ := json.Marshal(testCDR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := compositeEnricher.Enrich(data)
		if err != nil {
			b.Fatalf("Composite enrichment failed: %v", err)
		}
	}
}
