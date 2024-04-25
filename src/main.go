package main

import (
	"log"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/gavril-s/erpico-test-assignemnt/repository"
	"github.com/gavril-s/erpico-test-assignemnt/routes"
	"github.com/gavril-s/erpico-test-assignemnt/service"
)

func main() {
	dsn := "host=postgres user=user password=password dbname=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer closeDb(db)

	repo := repository.NewGormRepository(db)
	svc := &service.Service{
		Repo: repo,
	}

	r := routes.SetupRoutes(svc)
	port := ":5000"
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func closeDb(db *gorm.DB) {
	DB, err := db.DB()
	if err == nil {
		DB.Close()
	}
}
