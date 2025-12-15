// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Ingestion      IngestionConfig      `yaml:"ingestion"`
	Enrichment     EnrichmentConfig     `yaml:"enrichment"`
	Transformation TransformationConfig `yaml:"transformation"`
	Filtering      FilteringConfig      `yaml:"filtering"`
	Exporting      ExportingConfig      `yaml:"exporting"`
}

type IngestionConfig struct {
	Type       string `yaml:"type"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	BatchSize  int    `yaml:"batch_size"`
	BufferSize int    `yaml:"buffer_size"`
}

type EnrichmentConfig struct {
	Enabled       bool              `yaml:"enabled"`
	LookupSources []string          `yaml:"lookup_sources"`
	CacheSize     int               `yaml:"cache_size"`
	CacheTTL      int               `yaml:"cache_ttl"`
	CustomFields  map[string]string `yaml:"custom_fields"`
}

type TransformationConfig struct {
	Rules      []TransformationRule `yaml:"rules"`
	DateFormat string               `yaml:"date_format"`
	TimeZone   string               `yaml:"time_zone"`
}

type TransformationRule struct {
	Field     string `yaml:"field"`
	Operation string `yaml:"operation"`
	Value     string `yaml:"value"`
}

type FilteringConfig struct {
	Enabled    bool           `yaml:"enabled"`
	Conditions []FilterRule   `yaml:"conditions"`
	Action     string         `yaml:"action"`
}

type FilterRule struct {
	Field    string `yaml:"field"`
	Operator string `yaml:"operator"`
	Value    string `yaml:"value"`
}

type ExportingConfig struct {
	Type         string            `yaml:"type"`
	Destination  string            `yaml:"destination"`
	Format       string            `yaml:"format"`
	BatchSize    int               `yaml:"batch_size"`
	FlushTimeout int               `yaml:"flush_timeout"`
	Headers      map[string]string `yaml:"headers"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
