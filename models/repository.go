package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
}

type mysqlRepository struct {
	DB *sql.DB
}

func NewMySQLRepository(conn *sql.DB) Repository {
	return &mysqlRepository{DB: conn}
}

type testRepository struct {
	DB *sql.DB
}

func NewTestRepository(conn *sql.DB) Repository {
	return &testRepository{DB: nil}
}
