package service

import (
	"context"

	"github.com/khamiruf/go-toilets/converter"
	"github.com/khamiruf/go-toilets/model"
	"github.com/khamiruf/go-toilets/repository"
)

type BookRepository interface {
	GetBooks(context.Context) ([]*model.Book, error)
	AddBook(context.Context, model.Book) (int64, error)
	UpdateBook(context.Context, model.Book, int64) (int64, error)
	DeleteBook(context.Context, int64) (int64, error)
}

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) GetBooks() (*model.Books, error) {
	books, err := s.repo.GetBooks()
	if err != nil {
		return nil, err
	}

	return converter.StorageToModelBooks(books), nil
}

func (s *BookService) AddBook(book *model.Book) (int64, error) {
	rowsAffected, err := s.repo.CreateBook(book)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (s *BookService) UpdateBook(book *model.Book, id int) (int64, error) {
	rowsAffected, err := s.repo.UpdateBook(book, id)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (s *BookService) DeleteBook(id int64) (int64, error) {
	rowsAffected, err := s.repo.DeleteBook(id)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
