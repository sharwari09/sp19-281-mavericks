package main

type User struct {
	Id        string   `json:"id"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"zipcode,omitempty`
	Email     string   `json:"email"`
	Password  string   `json:"password,omitempty"`
}

type Address struct {
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
	Street  string `json:"street,omitempty"`
}
