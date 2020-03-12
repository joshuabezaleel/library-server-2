package book

// Book domain model.
type Book struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// NewBook creates a new instance of Book domain model.
func NewBook(title string, description string) *Book {
	return &Book{
		Title:       title,
		Description: description,
	}
}
