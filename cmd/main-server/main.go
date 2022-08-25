package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/khamiruf/go-toilets/database"
	"github.com/khamiruf/go-toilets/handler"
	"github.com/khamiruf/go-toilets/repository"
	"github.com/khamiruf/go-toilets/service"
	"github.com/pkg/errors"
)

func main() {
	// run
	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "main exit due to %s \n", err)
		os.Exit(1)
	}
}

func run() error {
	// connect to db here
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
		return errors.Wrap(err, "connection to database failed")
	}
	defer db.Close()

	repo := repository.NewBookRepository(db)
	service := service.NewBookService(repo)
	bookHandler := handler.NewBookHandler(service)

	http.HandleFunc("/books", bookHandler.ServeHTTP)
	err = http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err.Error())
		os.Exit(1)
	}

	return nil
}
