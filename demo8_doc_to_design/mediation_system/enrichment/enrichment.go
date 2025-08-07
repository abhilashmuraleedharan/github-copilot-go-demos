// [AI GENERATED] LLM: GitHub Copilot, Mode: Agent, Date: 2025-08-06
package enrichment

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// CDR represents a Call Data Record with common fields
type CDR struct {
	ID               string            `json:"id"`
	IMSI             string            `json:"imsi"`
	MSISDN           string            `json:"msisdn"`
	IMEI             string            `json:"imei"`
	CellID           string            `json:"cell_id"`
	Timestamp        time.Time         `json:"timestamp"`
	EventType        string            `json:"event_type"`
	ServiceID        string            `json:"service_id"`
	Duration         int64             `json:"duration"`
	DataVolume       int64             `json:"data_volume"`
	SubscriberType   string            `json:"subscriber_type,omitempty"`
	DeviceType       string            `json:"device_type,omitempty"`
	LocationRegion   string            `json:"location_region,omitempty"`
	Latitude         float64           `json:"latitude,omitempty"`
	Longitude        float64           `json:"longitude,omitempty"`
	ServicePlan      string            `json:"service_plan,omitempty"`
	AccountType      string            `json:"account_type,omitempty"`
	EnrichmentStatus map[string]string `json:"enrichment_status,omitempty"`
}

// Enricher defines the interface for enriching CDR data
type Enricher interface {
	Enrich(data []byte) ([]byte, error)
	GetName() string
}

// EnrichmentCache provides caching for enrichment lookups
type EnrichmentCache struct {
	data map[string]interface{}
	mu   sync.RWMutex
	ttl  time.Duration
}

// NewEnrichmentCache creates a new enrichment cache
func NewEnrichmentCache(ttl time.Duration) *EnrichmentCache {
	return &EnrichmentCache{
		data: make(map[string]interface{}),
		ttl:  ttl,
	}
}

// Get retrieves a value from the cache
func (c *EnrichmentCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

// Set stores a value in the cache
func (c *EnrichmentCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// CompositeEnricher manages multiple enrichers
type CompositeEnricher struct {
	enrichers []Enricher
	cache     *EnrichmentCache
}

// NewCompositeEnricher creates a new composite enricher
func NewCompositeEnricher(enrichers []Enricher, cacheTTL time.Duration) *CompositeEnricher {
	return &CompositeEnricher{
		enrichers: enrichers,
		cache:     NewEnrichmentCache(cacheTTL),
	}
}

// Enrich applies all enrichers to the CDR data
func (ce *CompositeEnricher) Enrich(data []byte) ([]byte, error) {
	var cdr CDR
	if err := json.Unmarshal(data, &cdr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CDR: %w", err)
	}

	// Initialize enrichment status tracking
	if cdr.EnrichmentStatus == nil {
		cdr.EnrichmentStatus = make(map[string]string)
	}

	// Apply each enricher
	for _, enricher := range ce.enrichers {
		enrichedData, err := enricher.Enrich(data)
		if err != nil {
			log.Printf("Enricher %s failed: %v", enricher.GetName(), err)
			cdr.EnrichmentStatus[enricher.GetName()] = "failed"
			continue
		}

		// Update CDR with enriched data
		if err := json.Unmarshal(enrichedData, &cdr); err != nil {
			log.Printf("Failed to unmarshal enriched data from %s: %v", enricher.GetName(), err)
			cdr.EnrichmentStatus[enricher.GetName()] = "failed"
			continue
		}

		cdr.EnrichmentStatus[enricher.GetName()] = "success"
		data = enrichedData
	}

	return json.Marshal(cdr)
}

// GetName returns the name of the composite enricher
func (ce *CompositeEnricher) GetName() string {
	return "CompositeEnricher"
}

// SubscriberEnricher enriches CDRs with subscriber information
type SubscriberEnricher struct {
	cache *EnrichmentCache
}

// NewSubscriberEnricher creates a new subscriber enricher
func NewSubscriberEnricher(cache *EnrichmentCache) *SubscriberEnricher {
	return &SubscriberEnricher{cache: cache}
}

// Enrich adds subscriber information to the CDR
func (se *SubscriberEnricher) Enrich(data []byte) ([]byte, error) {
	var cdr CDR
	if err := json.Unmarshal(data, &cdr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CDR: %w", err)
	}

	// Check cache first
	cacheKey := fmt.Sprintf("subscriber_%s", cdr.IMSI)
	if cached, exists := se.cache.Get(cacheKey); exists {
		if subscriberInfo, ok := cached.(map[string]string); ok {
			cdr.SubscriberType = subscriberInfo["type"]
			cdr.ServicePlan = subscriberInfo["plan"]
			cdr.AccountType = subscriberInfo["account_type"]
		}
	} else {
		// Simulate subscriber lookup
		subscriberInfo := se.lookupSubscriber(cdr.IMSI)
		cdr.SubscriberType = subscriberInfo["type"]
		cdr.ServicePlan = subscriberInfo["plan"]
		cdr.AccountType = subscriberInfo["account_type"]

		// Cache the result
		se.cache.Set(cacheKey, subscriberInfo)
	}

	return json.Marshal(cdr)
}

// GetName returns the enricher name
func (se *SubscriberEnricher) GetName() string {
	return "SubscriberEnricher"
}

// lookupSubscriber simulates a subscriber database lookup
func (se *SubscriberEnricher) lookupSubscriber(imsi string) map[string]string {
	// Simulate different subscriber types based on IMSI pattern
	if strings.HasPrefix(imsi, "001") {
		return map[string]string{
			"type":         "premium",
			"plan":         "unlimited_5g",
			"account_type": "postpaid",
		}
	} else if strings.HasPrefix(imsi, "002") {
		return map[string]string{
			"type":         "standard",
			"plan":         "basic_4g",
			"account_type": "prepaid",
		}
	}
	return map[string]string{
		"type":         "basic",
		"plan":         "voice_only",
		"account_type": "prepaid",
	}
}

// DeviceEnricher enriches CDRs with device information
type DeviceEnricher struct {
	cache *EnrichmentCache
}

// NewDeviceEnricher creates a new device enricher
func NewDeviceEnricher(cache *EnrichmentCache) *DeviceEnricher {
	return &DeviceEnricher{cache: cache}
}

// Enrich adds device information to the CDR
func (de *DeviceEnricher) Enrich(data []byte) ([]byte, error) {
	var cdr CDR
	if err := json.Unmarshal(data, &cdr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CDR: %w", err)
	}

	// Check cache first
	cacheKey := fmt.Sprintf("device_%s", cdr.IMEI)
	if cached, exists := de.cache.Get(cacheKey); exists {
		if deviceType, ok := cached.(string); ok {
			cdr.DeviceType = deviceType
		}
	} else {
		// Simulate device lookup
		deviceType := de.lookupDevice(cdr.IMEI)
		cdr.DeviceType = deviceType

		// Cache the result
		de.cache.Set(cacheKey, deviceType)
	}

	return json.Marshal(cdr)
}

// GetName returns the enricher name
func (de *DeviceEnricher) GetName() string {
	return "DeviceEnricher"
}

// lookupDevice simulates a device database lookup
func (de *DeviceEnricher) lookupDevice(imei string) string {
	// Simulate device type detection based on IMEI pattern
	if strings.HasPrefix(imei, "35") {
		return "smartphone"
	} else if strings.HasPrefix(imei, "86") {
		return "iot_device"
	}
	return "feature_phone"
}

// GeoEnricher enriches CDRs with geographic information
type GeoEnricher struct {
	cache *EnrichmentCache
}

// NewGeoEnricher creates a new geo enricher
func NewGeoEnricher(cache *EnrichmentCache) *GeoEnricher {
	return &GeoEnricher{cache: cache}
}

// Enrich adds geographic information to the CDR
func (ge *GeoEnricher) Enrich(data []byte) ([]byte, error) {
	var cdr CDR
	if err := json.Unmarshal(data, &cdr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CDR: %w", err)
	}

	// Check cache first
	cacheKey := fmt.Sprintf("geo_%s", cdr.CellID)
	if cached, exists := ge.cache.Get(cacheKey); exists {
		if geoInfo, ok := cached.(map[string]interface{}); ok {
			cdr.LocationRegion = geoInfo["region"].(string)
			cdr.Latitude = geoInfo["latitude"].(float64)
			cdr.Longitude = geoInfo["longitude"].(float64)
		}
	} else {
		// Simulate geo lookup
		geoInfo := ge.lookupGeoLocation(cdr.CellID)
		cdr.LocationRegion = geoInfo["region"].(string)
		cdr.Latitude = geoInfo["latitude"].(float64)
		cdr.Longitude = geoInfo["longitude"].(float64)

		// Cache the result
		ge.cache.Set(cacheKey, geoInfo)
	}

	return json.Marshal(cdr)
}

// GetName returns the enricher name
func (ge *GeoEnricher) GetName() string {
	return "GeoEnricher"
}

// lookupGeoLocation simulates a geo location lookup
func (ge *GeoEnricher) lookupGeoLocation(cellID string) map[string]interface{} {
	// Simulate different regions based on cell ID pattern
	if strings.HasPrefix(cellID, "NYC") {
		return map[string]interface{}{
			"region":    "New York",
			"latitude":  40.7128,
			"longitude": -74.0060,
		}
	} else if strings.HasPrefix(cellID, "LAX") {
		return map[string]interface{}{
			"region":    "Los Angeles",
			"latitude":  34.0522,
			"longitude": -118.2437,
		}
	}
	return map[string]interface{}{
		"region":    "Unknown",
		"latitude":  0.0,
		"longitude": 0.0,
	}
}
