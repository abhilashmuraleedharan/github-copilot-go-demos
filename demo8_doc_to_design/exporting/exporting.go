// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
package exporting

import (
	"demo8_doc_to_design/config"
	"demo8_doc_to_design/ingestion"
)

type Exporter interface {
	Export(cdr ingestion.CDR) error
}

func NewExporter(cfg config.ExportingConfig) Exporter {
	return &DefaultExporter{}
}

type DefaultExporter struct{}

func (e *DefaultExporter) Export(cdr ingestion.CDR) error {
	// TODO: implement export logic
	return nil
}
