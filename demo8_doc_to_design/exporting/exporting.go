// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2024-06-13
package exporting

import "fmt"

const (
	// ExportDestination defines the default export destination.
	ExportDestination = "output.csv"
)

// ExportCDR exports a CDR record.
func ExportCDR(record string) error {
	fmt.Println("Exporting CDR:", record)
	return nil
}
