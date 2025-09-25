// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
package enrichment

import (
	"demo8_doc_to_design/config"
	"demo8_doc_to_design/ingestion"
)

type Enricher interface {
	Enrich(cdr ingestion.CDR) ingestion.CDR
}

func NewEnricher(cfg config.EnrichmentConfig) Enricher {
	return &DefaultEnricher{}
}

type DefaultEnricher struct{}

func (e *DefaultEnricher) Enrich(cdr ingestion.CDR) ingestion.CDR {
	// TODO: implement enrichment logic
	return cdr
}
