// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Ingestion      IngestionConfig
	Enrichment     EnrichmentConfig
	Transformation TransformationConfig
	Filtering      FilteringConfig
	Exporting      ExportingConfig
	Service        ServiceConfig
}

type IngestionConfig struct {
	// Kafka, SFTP, file settings
}
type EnrichmentConfig struct {
	// Cache, lookup sources, plugin settings
}
type TransformationConfig struct {
	// Mapping rules, masking, calculated fields
}
type FilteringConfig struct {
	// Event types, timestamp windows, tags
}
type ExportingConfig struct {
	// Output formats, destinations, retry policy
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
type ServiceConfig struct {
	RuleSet           string              // Name or version of transformation/mapping rules
	EnrichmentSources []string            // List of enrichment sources (e.g., "subscriber", "equipment", "location", "service_plan")
	OutputFormats     map[string]string   // Destination name → format (e.g., "analytics": "avro", "legacy": "csv")
	DryRun            bool                // Enable dry-run mode
	AuditMode         bool                // Enable audit mode
	ReloadOnChange    bool                // Enable config reload without restart
	DeadLetterQueue   string              // DLQ destination (e.g., Kafka topic or SFTP path)
	ExportDestinations []ExportDestination // List of export destinations
}

type ExportDestination struct {
	Name   string            // Destination name
	Type   string            // "kafka", "sftp", "http"
	Config map[string]string // Arbitrary config for destination (e.g., topic, path, endpoint)
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
