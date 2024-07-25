package domain

import (
	"errors"
	"testing"
)

func TestNewEmailNoAt(t *testing.T) {
	_, err := NewEmail("x")

	if err == nil {
		t.Fatalf("no error returned")
	}

  if !errors.Is(err, ErrInvalidEmail) {
    t.Fatalf("error is not invalid email: %s", err)
  }
}

func TestNewEmailWithAt(t *testing.T) {
	want := Email("x@x")
	res, err := NewEmail("x@x")
	if res != want {
		t.Fatalf("should have result as %s, has result: %s", want, res)
	}

	if err != nil {
		t.Fatalf("error returned: %s", err)
	}
}

func TestNewEmailEmpty(t *testing.T) {
	_, err := NewEmail("")

	if err == nil {
		t.Fatalf("no error returned")
	}

  if !errors.Is(err, ErrInvalidEmail) {
    t.Fatalf("error is not invalid email: %s", err)
  }
}
