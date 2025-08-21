// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
package enrichment

import "fmt"

const (
	// EnrichmentSource defines the source for enrichment data.
	EnrichmentSource = "external"
)

// EnrichCDR enriches a CDR record.
func EnrichCDR(record string) string {
	fmt.Println("Enriching CDR:", record)
	return record
}
