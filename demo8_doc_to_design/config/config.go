// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
package config

const (
	// DefaultConfigPath defines the default path to the config file.
	DefaultConfigPath = "config.yaml"
)

// Config represents the microservice configuration including rule selection, enrichment sources, and output formats.
type Config struct {
	RuleSelection     []string `json:"rule_selection" yaml:"rule_selection"`
	EnrichmentSources []string `json:"enrichment_sources" yaml:"enrichment_sources"`
	OutputFormats     []string `json:"output_formats" yaml:"output_formats"`
}

// LoadConfig loads the configuration.
func LoadConfig(path string) error {
	// Placeholder for config loading logic.
	return nil
}
