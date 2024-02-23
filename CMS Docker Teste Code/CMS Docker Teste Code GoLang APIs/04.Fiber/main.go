package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ChrisMarSilva/cms.golang.teste.api.fiber/internal/db"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.fiber
// go get -u github.com/gofiber/fiber/v3
// go install github.com/cosmtrek/air@latest
// go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
// go get -u github.com/jackc/pgx/v5
// go get -u github.com/jackc/pgx/v5/pgxpool
// go get -u github.com/go-sql-driver/mysql
// go mod tidy

// go run main.go
// air

// sqlc generate
// sqlc verify

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7004"
	}

	log.Printf("Server running at port: " + port)

	app := fiber.New()
	app.Get("/", handlerDefault)
	app.Get("/banco", handlerBancoPostgreSQL)
	// app.Get("/banco", handlerBancoMySQL)
	// app.Get("/banco", handlerBancoSQLite)
	log.Fatal(app.Listen(":" + port))
}

func handlerDefault(c fiber.Ctx) error {
	return c.SendString("Golang api 04.Fiber OK")
}

func handlerBancoPostgreSQL(c fiber.Ctx) error {
	ctx := context.Background()

	uri := "host=localhost port=5432 dbname=rinha user=admin password=123 sslmode=disable"

	//conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListAuthors failed: %v\n", err)
		return err
	}
	log.Println("authors")
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{Name: "Brian Kernighan", Bio: pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true}})
	if err != nil {
		fmt.Fprintf(os.Stderr, "CreateAuthor failed: %v\n", err)
		return err
	}
	log.Println("insertedAuthor", insertedAuthor)

	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
		return err
	}
	log.Println("Name", fetchedAuthor.Name)

	var greeting string
	err = conn.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	log.Println("greeting", greeting)

	dbpool, err := pgxpool.New(ctx, uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return err
	}
	defer dbpool.Close()

	var greeting2 string
	err = dbpool.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	log.Println("greeting2", greeting2)

	return c.SendString("OK")
}

// func handlerBancoMySQL(c fiber.Ctx) error {
// 	ctx := context.Background()
//
// 	db, err := sql.Open("mysql", "user:password@/dbname?parseTime=true")
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close(ctx)
//
// 	queries := tutorial.New(db)
//
// 	authors, err := queries.ListAuthors(ctx)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "ListAuthors failed: %v\n", err)
// 		return err
// 	}
// 	log.Println("authors")
//
// 	// create an author
// 	result, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{Name: "Brian Kernighan", Bio: sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true}})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "CreateAuthor failed: %v\n", err)
// 		return err
// 	}
// 	insertedAuthorID, err := result.LastInsertId()
// 	if err != nil {
// 		return err
// 	}
// 	log.Println("insertedAuthor", insertedAuthor)
//
// 	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
// 		return err
// 	}
// 	log.Println("Name", fetchedAuthor.Name)
//
// 	return c.SendString("OK")
// }

// func handlerBancoSQLite(c fiber.Ctx) error {
// 	ctx := context.Background()
//
// 	db, err := sql.Open("sqlite3", ":memory:")
// 	if err != nil {
// 		return err
// 	}
//
// 	//go:embed schema.sql
// 	var ddl string
//
// 	// create tables
// 	if _, err := db.ExecContext(ctx, ddl); err != nil {
// 		return err
// 	}
//
// 	queries := tutorial.New(db)
//
// 	authors, err := queries.ListAuthors(ctx)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "ListAuthors failed: %v\n", err)
// 		return err
// 	}
// 	log.Println("authors")
//
// 	// create an author
// 	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{Name: "Brian Kernighan", Bio: sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true}})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "CreateAuthor failed: %v\n", err)
// 		return err
// 	}
// 	log.Println("insertedAuthor", insertedAuthor)
//
// 	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "GetAuthor failed: %v\n", err)
// 		return err
// 	}
// 	log.Println("Name", fetchedAuthor.Name)
//
// 	return c.SendString("OK")
// }
