package main

import "testing"

func TestFileType(t *testing.T) {
	ft := TemplateFunctionFileType("application/pdf")

	if ft != "pdf" {
		t.Fatalf("file type not detected for pdf, found: %s", ft)
	}
}

func TestTruncate(t *testing.T) {
	s := TemplateFunctionTruncate("The quick brown fox jumps over the lazy dog", 12)

	if s != "The quick br" {
		t.Fatalf("word not truncated, found: %s", s)
	}
}

func TestContact(t *testing.T) {
	s := TemplateFunctionContact("Example <example@example.com>")

	if s != "Example" {
		t.Fatalf("file type not detected for pdf, found: %s", s)
	}
}

func TestEmail(t *testing.T) {
	s := TemplateFunctionEmail("Example <example@example.com>")

	if s != "example@example.com" {
		t.Fatalf("file type not detected for pdf, found: %s", s)
	}
}

func TestGravatar(t *testing.T) {
	s := TemplateFunctionGravatar("Example <example@example.com>")
	expected := "https://www.gravatar.com/avatar/6864e11f40e6ddf6dbdc41d6888d3c26.jpg?rating=g&size=80"

	if s != expected {
		t.Fatalf("file type not detected for pdf, found: %s", s)
	}
}
