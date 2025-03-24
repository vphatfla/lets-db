package parser_test

import (
	"testing"
	"github.com/vphatfla/lets-db/db/parser"
)

func TestParser(t *testing.T) {
	selectQ := "Select abc"
	deleteQ := "deleTE abc"
	insertQ := "inserT abc value 123"

	// test select query parser
	if _, key, _, err := parser.ParseQuery(selectQ); err != nil {
		t.Errorf("Parsing select query error -> %v \n", err)
	} else {
		t.Logf("For query: %s, parser returned key: %s \n", selectQ, key)
	}

	// test delete query parser
	if _, key, _, err := parser.ParseQuery(deleteQ); err != nil {
		t.Errorf("Parsing delete query error -> %v \n", err)
	} else {
		t.Logf("For query: %s, parser returned key: %s \n", deleteQ, key)
	}

	// test insert query parser
	if _, key, value, err := parser.ParseQuery(insertQ); err != nil {
		t.Errorf("Parsing insert query error -> %v \n", err)
	} else if value != "123" {
		t.Errorf("Intert query, wrong value, given queyr %s, parser return value %s \n", insertQ, value)
	} else {
		t.Logf("For query: %s, parser returned key %s and value %s \n", insertQ, key, value)
	}
}


