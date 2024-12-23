package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// Define the connection string
	connStr := os.Getenv("DB_URL")
	var err error
	// Open a connection to the database
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to databse", err)
	}

	/*
	   // Example query
	   rows, err := db.Query("SELECT id, name FROM your_table")

	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   defer rows.Close()

	   // Iterate through the results

	   	for rows.Next() {
	   		var id int
	   		var name string
	   		if err := rows.Scan(&id, &name); err != nil {
	   			log.Fatal(err)
	   		}
	   		fmt.Printf("ID: %d, Name: %s\n", id, name)
	   	}

	   // Check for errors from iterating over rows

	   	if err := rows.Err(); err != nil {
	   		log.Fatal(err)
	   	}
	*/
}
