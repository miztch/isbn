package main

import (
	"testing"
)

func TestGetIsbn10(t *testing.T) {
	got, err := GetIsbn10("9781492077213")

	if err != nil {
		t.Error(err)
	}

	want, err := "1492077216", nil

	if got != want {
		t.Errorf("want: %s, got: %s", got, want)
	}
}

func TestGetIsbn13(t *testing.T) {
	got, err := GetIsbn13("1492077216")

	if err != nil {
		t.Error(err)
	}

	want, err := "9781492077213", nil

	if got != want {
		t.Errorf("want: %s, got: %s", got, want)
	}
}
