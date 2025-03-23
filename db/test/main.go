package main

import (
	"log"

	"github.com/vphatfla/lets-db/db/parser"
)

func main() {
	selectQ := "Select abc"
	deleteQ := "deleTE abc"
	insertQ := "inserT abc value 123"

	// test select query parser
	if _, key, _, err := parser.ParseQuery(selectQ); err != nil {
		log.Printf("Parsing select query error -> %v \n", err)
	} else {
		log.Printf("For query: %s, parser returned key: %s \n", selectQ, key)
	}

	// test delete query parser
	if _, key, _, err := parser.ParseQuery(deleteQ); err != nil {
		log.Printf("Parsing delete query error -> %v \n", err)
	} else {
		log.Printf("For query: %s, parser returned key: %s \n", deleteQ, key)
	}

	// test insert query parser
	if _, key, value, err := parser.ParseQuery(insertQ); err != nil {
		log.Printf("Parsing insert query error -> %v \n", err)
	} else {
		log.Printf("For query: %s, parser returned key %s and value %s \n", insertQ, key, value)
	}

}
