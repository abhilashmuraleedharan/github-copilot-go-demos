// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
package ingestion

import "demo8_doc_to_design/config"

type CDR struct {
	// Raw CDR fields
}

type Ingestor interface {
	Ingest() <-chan CDR
}

func NewIngestor(cfg config.IngestionConfig) Ingestor {
	return &DefaultIngestor{}
}

type DefaultIngestor struct{}

func (i *DefaultIngestor) Ingest() <-chan CDR {
	out := make(chan CDR)
	go func() {
		// TODO: implement Kafka/SFTP/file ingestion
		close(out)
	}()
	return out
}
