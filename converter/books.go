package converter

import (
	"github.com/khamiruf/go-toilets/model"
	"github.com/khamiruf/go-toilets/storage"
)

func StorageToModelBook(book *storage.Book) *model.Book {
	if book == nil {
		return nil
	}

	return &model.Book{
		ID:          book.ID,
		Author:      book.Author,
		Title:       book.Title,
		Description: book.Description,
		Rating:      book.Rating,
		CreatedAt:   SqlToTime(&book.CreatedAt),
		ModifiedAt:  SqlToTime(&book.ModifiedAt),
		Disabled:    book.Disabled,
		DisabledAt:  SqlToTime(&book.DisabledAt),
	}
}

func StorageToModelBooks(books []*storage.Book) *model.Books {
	if len(books) <= 0 {
		return nil
	}
	listOfBooks := &model.Books{}

	modelBooks := make([]*model.Book, len(books))
	for i, bk := range books {
		modelBooks[i] = StorageToModelBook(bk)
	}

	listOfBooks.Books = modelBooks

	return listOfBooks
}
