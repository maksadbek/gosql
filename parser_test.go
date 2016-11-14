package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	s := `SELECT first_name, last_name, age FROM my_table`
	expectedStmt := &SelectStatement{
		Fields:    []string{"first_name", "last_name", "age"},
		TableName: "my_table",
	}

	r := bytes.NewReader([]byte(s))

	p := NewParser(r)

	stmt, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", stmt)

	if stmt.TableName != expectedStmt.TableName {
		t.Fatalf("want %s, got %s", expectedStmt, stmt.TableName)
	}

	if !reflect.DeepEqual(stmt.TableName, expectedStmt.TableName) {
		t.Fatalf("fuck, not equal")
	}
}
