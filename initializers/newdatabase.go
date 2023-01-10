package initializers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// var NewDB *sql.DB

// func main() {
// 	var err error
// 	psqlInfo := os.Getenv("DB_URL")

// 	NewDB, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer NewDB.Close()

// 	err = NewDB.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected!")
// }

func main() {
	// Create a connection config

	// Connect to the database
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	fmt.Println("Connected to the database.")

}
