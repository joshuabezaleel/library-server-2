package book

import (
	"errors"
)

// Errors definition.
var (
	ErrGetAll          = errors.New("Failed retrieving all Books")
	ErrGetBookTopicIDs = errors.New("Failed retrieving Book's topic IDs")
	ErrGetTopicsByID   = errors.New("Failed retrieving topics")
)

// Service provides basic operations on Book domain model.
type Service interface {
	GetAll(topicIDs []int) ([]*Book, error)
	GetBookTopicIDs(bookID int) ([]int, error)
	GetTopicsByID(topicIDs []int) ([]string, error)
}

type service struct {
	bookRepository Repository
}

// NewBookService creates an instance of the service for the Book domain model
// with all of the necessary dependencies.
func NewBookService(bookRepository Repository) Service {
	return &service{
		bookRepository: bookRepository,
	}
}

func (s *service) GetAll(topicIDs []int) ([]*Book, error) {
	books, err := s.bookRepository.GetAll(topicIDs)
	if err != nil {
		return nil, ErrGetAll
	}

	for _, book := range books {
		bookTopicIDs, err := s.GetBookTopicIDs(book.ID)
		if err != nil {
			return nil, ErrGetBookTopicIDs
		}

		bookTopics, err := s.GetTopicsByID(bookTopicIDs)
		if err != nil {
			return nil, ErrGetTopicsByID
		}

		book.Topics = bookTopics
	}

	return books, nil
}

func (s *service) GetBookTopicIDs(bookID int) ([]int, error) {
	bookTopicIDs, err := s.bookRepository.GetBookTopicIDs(bookID)
	if err != nil {
		return nil, ErrGetBookTopicIDs
	}

	return bookTopicIDs, nil
}

func (s *service) GetTopicsByID(topicIDs []int) ([]string, error) {
	topics, err := s.bookRepository.GetTopicsByID(topicIDs)
	if err != nil {
		return nil, ErrGetTopicsByID
	}

	return topics, nil
}
