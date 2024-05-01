package services

import (
	"fmt"
	"testWallet/internal/models"
)


type PostgresRepository struct{
	Users []models.User
}

func NewPostgresRepository() (models.Cruder, error){

	return &PostgresRepository{}, nil
}

func (p PostgresRepository)Create(u models.User) error{
	fmt.Println("Create user!")
	return nil
}

func (p PostgresRepository)Delete(id string) error{
	fmt.Println("Create user!")
	return nil
}

func (p PostgresRepository)Update(id string, u models.User) error{
	fmt.Println("Create user!")
	return nil
}

func (p PostgresRepository)GetById(id string) error{
	fmt.Println("Create user!")
	return nil
}