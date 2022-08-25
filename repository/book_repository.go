package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/khamiruf/go-toilets/model"
	"github.com/khamiruf/go-toilets/storage"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		DB: db,
	}
}

func (r *BookRepository) GetBooks() ([]*storage.Book, error) {
	rows, err := r.DB.Query("SELECT * FROM book WHERE disabled = false;")
	if err != nil {
		return nil, err
	}

	books := make([]*storage.Book, 0)
	for rows.Next() {
		book := storage.Book{}
		if err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Description, &book.Rating, &book.CreatedAt, &book.ModifiedAt, &book.Disabled, &book.DisabledAt); err != nil {
			log.Fatal(err)
		}
		books = append(books, &book)
		// fmt.Println(books)
	}

	return books, nil
}

func (r *BookRepository) CreateBook(book *model.Book) (int64, error) {
	res, err := r.DB.Exec("INSERT INTO book (author, title, description, rating) VALUES ($1, $2, $3, $4);",
		book.Author,
		book.Title,
		book.Description,
		book.Rating,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *BookRepository) UpdateBook(book *model.Book, id int) (int64, error) {
	// call updateStatement
	var arg []interface{}
	qry, args := updateStatement(id, arg, book)

	res, err := r.DB.Exec(qry, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func updateStatement(id int, argArray []interface{}, updateData *model.Book) (string, []interface{}) {
	var queryString string
	i := 2 // update values
	argArray = append(argArray, id)
	if len(updateData.Author) > 0 {
		queryString += fmt.Sprintf(`author=$%d,`, i)
		argArray = append(argArray, updateData.Author)
		i++
	}

	if len(updateData.Title) > 0 {
		queryString += fmt.Sprintf(`title=$%d,`, i)
		argArray = append(argArray, updateData.Title)
		i++
	}

	if len(updateData.Description) > 0 {
		queryString += fmt.Sprintf(`description=$%d,`, i)
		argArray = append(argArray, updateData.Description)
		i++
	}

	if updateData.Rating != nil {
		queryString += fmt.Sprintf(`rating=$%d`, i)
		argArray = append(argArray, updateData.Rating)
		i++
	}

	qry := fmt.Sprintf("UPDATE book SET %s WHERE id=$1", strings.TrimSuffix(queryString, ","))
	fmt.Println(qry)
	fmt.Printf("%v+ \n", argArray)

	return qry, argArray
}

func (r *BookRepository) DeleteBook(id int64) (int64, error) {
	// implement soft delete
	sqlStatement := `UPDATE book
		SET disabled = true, disabled_at = now()
		WHERE id = $1;`
	res, err := r.DB.Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
