package service

import "testing"

func TestGenerateSessionId(t *testing.T) {
	got := len(GenerateSessionId())
	want := 64
	if got != want {
		t.Errorf("GenerateSessionId() = %d; want %d", got, want)
	}
}

func TestHashPassword(t *testing.T) {
	password := "secretpassword123"
	_, err := HashPassword(password)
	if err != nil {
		t.Errorf("TestHashPassword() = %s", err.Error())
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "secretpassword123"
	hash, _ := HashPassword(password)
	got := VerifyPassword(password, hash)
	want := true
	if got != want {
		t.Errorf("TestVerifyPassword() = %t; want %t", got, want)
	}
}
