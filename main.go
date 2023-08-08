package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"urlgo/config"
	"urlgo/controllers"
	"urlgo/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//setup mysql connection
	db, err := sql.Open("mysql", config.GetDBConnectionString())

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id INT AUTO_INCREMENT PRIMARY KEY,
			original VARCHAR(255) NOT NULL,
			shortCode VARCHAR(255) NOT NULL
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	urlService := services.NewUrlService(db)
	urlController := controllers.NewUrlController(urlService)

	http.HandleFunc("/create", urlController.CreateUrl)
	http.HandleFunc("/get", urlController.GetUrlByShortCode)

	port := 1515
	fmt.Printf("server is running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
