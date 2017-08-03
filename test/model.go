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

//go:generate go run ../main/main.go User
type User struct {
	Id     string `neo4j:"index"`
	Name   string
	Email  string `neo4j:"index"`
	Age    int
	Active bool
	Happy  bool
}

//go:generate go run ../main/main.go Address
type Address struct {
	Id     string
	Street string
	City   string
	State  string
	zip    string
}

//go:generate go run ../main/main.go CreditCard
type CreditCard struct {
	Id             string
	Number         string
	ExpirationDate string
	SecurityCode   string
	NameOnCard     string
	Typ            string
}
