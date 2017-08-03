package main

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

var driver bolt.DriverPool

func init() {
	var err error
	driver, err = bolt.NewDriverPool("", 10)
	if err != nil {
		panic(err)
	}
}

//go:generate go run ../neo4jGenerator/main.go User
type User struct {
	Id     string `json:"id" neo4j:"index"`
	Name   string `json:"name"`
	Email  string `json:"email" neo4j:"index"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
	Happy  bool   `json:"happy"`
}

//go:generate go run ../neo4jGenerator/main.go Address
type Address struct {
	Id     string
	Street string
	City   string
	State  string
	zip    string
}

//go:generate go run ../neo4jGenerator/main.go CreditCard
type CreditCard struct {
	Id             string
	Number         string
	ExpirationDate string
	SecurityCode   string
	NameOnCard     string
	Typ            string
}
