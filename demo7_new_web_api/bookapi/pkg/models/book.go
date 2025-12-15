// [AI GENERATED] LLM: GitHub Copilot, Mode: Inline Suggestion, Date: 2023-10-05
package models

// Book represents the data model for a book in the library.
type Book struct {
    ID     string `json:"id"`     // Unique identifier for the book
    Title  string `json:"title"`  // Title of the book
    Author string `json:"author"` // Author of the book
    Year   int    `json:"year"`   // Year of publication
}