package persistence

import (
	"github.com/joshuabezaleel/library-server-2/service/book/book"
)

var allBooks = []*book.Book{
	&book.Book{ID: 1, Title: "Medic Without Shame", Author: "Richard Gizmo Seventh"},
	&book.Book{ID: 2, Title: "Blacksmith Of The Forest", Author: "Mark R. Steel"},
	&book.Book{ID: 3, Title: "Men With Determination", Author: "Maxwell Southwell"},
	&book.Book{ID: 4, Title: "Snakes Of My Imagination", Author: "M. R. Smitherd"},
	&book.Book{ID: 5, Title: "Soldiers And Witches", Author: "Gizmo Richards"},
	&book.Book{ID: 6, Title: "Lions And Men", Author: "Richard Gizmo Seventh"},
	&book.Book{ID: 7, Title: "Restoration Of Wood", Author: "Marissa Smirnova"},
	&book.Book{ID: 8, Title: "Curse Of The River", Author: "Maree Markson"},
	&book.Book{ID: 9, Title: "Write About Technology", Author: "Smith Markblood"},
	&book.Book{ID: 10, Title: "Duke Of The North", Author: "Mathias R. Seventh"},
	&book.Book{ID: 11, Title: "Tree Of Darkness", Author: "Ocean Roxie Northern"},
	&book.Book{ID: 12, Title: "Humans Of The Ancestors", Author: "Wendell Billerbeck"},
	&book.Book{ID: 13, Title: "Fish Of The Gods", Author: "Mathias R. Seventh"},
	&book.Book{ID: 14, Title: "Doctors And Owls", Author: "Wearmouth O. Northern"},
	&book.Book{ID: 15, Title: "Hunters And Assassins", Author: "Octavia Birszwilks"},
	&book.Book{ID: 16, Title: "Surprise With Honor", Author: "Richard Gizmo Seventh"},
	&book.Book{ID: 17, Title: "Choice Of Yesterday", Author: "Bishop Wenblood"},
	&book.Book{ID: 18, Title: "Sounds In The East", Author: "Wes O. Armedrobber"},
	&book.Book{ID: 19, Title: "Crying In The Stars", Author: "M. R. Smitherd"},
	&book.Book{ID: 20, Title: "King With Sins", Author: "Joshua Bezaleel Abednego"},
}

var booksTopics = map[int][]int{
	1:  []int{1},
	2:  []int{1},
	3:  []int{1, 2},
	4:  []int{1, 4},
	5:  []int{2},
	6:  []int{2, 3},
	7:  []int{2},
	8:  []int{2},
	9:  []int{2},
	10: []int{2},
	11: []int{2, 4},
	12: []int{3, 1},
	13: []int{3},
	14: []int{3},
	15: []int{3, 4},
	16: []int{4},
	17: []int{4},
	18: []int{4},
	19: []int{4},
	20: []int{4},
}

var allTopics = map[int]string{
	1: "Fantasy",
	2: "Crime",
	3: "Psychology",
	4: "Law",
}

type bookRepository struct{}

// NewBookRepository returns initialized implementations of the repository for
// Book domain model.
func NewBookRepository() book.Repository {
	return &bookRepository{}
}

func (repo *bookRepository) GetAll() ([]*book.Book, error) {
	// var oneBook *book.Book
	// var allBooks []*book.Book

	// rows, err := repo.DB.Query("SELECT * FROM books")
	// if err != nil {
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	err = rows.Scan(oneBook)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	allBooks = append(allBooks, oneBook)
	// }

	// return allBooks, nil
	return allBooks, nil
}

func (repo *bookRepository) GetBookTopicIDs(bookID int) ([]int, error) {
	return booksTopics[bookID], nil
}

func (repo *bookRepository) GetTopicsByID(topicIDs []int) ([]string, error) {
	var topics []string

	for _, topicID := range topicIDs {
		topics = append(topics, allTopics[topicID])
	}

	return topics, nil
}
