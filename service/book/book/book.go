package book

// Book domain model.
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// NewBook creates a new instance of Book domain model.
func NewBook(id int, title string, description string) *Book {
	return &Book{
		ID:          id,
		Title:       title,
		Description: description,
	}
}
