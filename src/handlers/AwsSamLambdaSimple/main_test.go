package main

import (
	"context"
	"testing"
)

func TestHandler(t *testing.T) {
	// Arrange
	want := "Hello, World!"
	
	// Act
	got, err := handler(context.Background())

	//Assert
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if got != want {
		t.Errorf("handler = %q, want %q", got, want)
	}
}
func TestRun(t *testing.T) {
	lambdaStart = func(handler any) {
	}
	main()
}
