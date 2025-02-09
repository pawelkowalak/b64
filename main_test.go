package main

import (
	"testing"
)

func TestReadArgsWithEncode(t *testing.T) {
	text, decode, err := readArgs([]string{"cmd", "hello"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if text != "hello" {
		t.Fatalf("expected text to be 'hello', got %v", text)
	}
	if decode {
		t.Fatalf("expected decode to be false, got %v", decode)
	}
}

func TestReadArgsWithDecode(t *testing.T) {
	text, decode, err := readArgs([]string{"cmd", "-d", "aGVsbG8="})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if text != "aGVsbG8=" {
		t.Fatalf("expected text to be 'aGVsbG8=', got %v", text)
	}
	if !decode {
		t.Fatalf("expected decode to be true, got %v", decode)
	}
}

func TestReadArgsWithInvalidArgs(t *testing.T) {
	_, _, err := readArgs([]string{"cmd"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "invalid input arguments" {
		t.Fatalf("expected 'invalid input arguments', got %v", err)
	}
}
