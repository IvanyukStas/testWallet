/*
Package models is

*/
package models


//User is 
type User struct{
	ID 		string `json:"id,omitempty"`
	Name	string `json:"name" validate:"required"`	
	Wallet		   `json:"wallet,omitempty"`
}


//Wallet is wallet of user
type Wallet struct{
	ID string	   `json:"id,omitempty"`
	Balance string `json:"balance,omitempty"`
}


//Cruder is CRUD interface
type Cruder interface{
	Create(User) error
	Urdate(string, User) error
	Delete(string) error
	GetById(string) error
}