package models

import "testing"


func TestCreateUser(t *testing.T) {
	want := "User created"
	if got := CreateUser(); got != want {
		t.Errorf("CreateUser() = %q, want %q", got, want)
	}
}