package myDatabase

import "testing"

func TestInitializeDB(t *testing.T) {
	err := InitializeDB()
	if err != nil {
		t.Fatalf("TestInitializeDB error: %v", err)
	}
}
