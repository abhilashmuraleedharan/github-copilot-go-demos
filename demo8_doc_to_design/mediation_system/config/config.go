// [AI GENERATED] LLM: GitHub Copilot, Mode: Edit, Date: 2024-06-09
package config

// Config represents the microservice configuration.
type Config struct {
	RuleSet           string   // Name of the transformation rule set to use
	EnrichmentSources []string // List of enrichment sources (e.g., geo, customer DB)
	OutputFormats     []string // Supported output formats (e.g., CSV, JSON)
}
