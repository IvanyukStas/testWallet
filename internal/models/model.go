/*
Package models is
*/
package models

import "errors"

var ErrorUserExists = errors.New("User alreagy exists in database")

//User is 
type User struct{
	ID 		string `json:"id,omitempty"`
	Name	string `json:"name" validate:"required"`
	Time	string		
	Wallet		   `json:"wallet,omitempty"`
}


//Wallet is wallet of user
type Wallet struct{
	IDWallet string	   `json:"id,omitempty"`
	Balance string `json:"balance,omitempty"`
}


//Cruder is CRUD interface
type Cruder interface{
	Create(string) error
	Update(string, string)  error
	Delete(string) error
	GetByName(*User) (*User, error)
}