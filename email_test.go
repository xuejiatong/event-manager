package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	t.Parallel()
	if !isValidEmail("kyle.jensen@yale.edu") {
		t.Errorf("Expected true, got false for kyle.jensen@yale.edu")
	}
	if isValidEmail("blahblahblah") {
		t.Errorf("Expected false, got true for blahblahblah")
	}
}

func TestIsYaleEmail(t *testing.T) {
	t.Parallel()
	if !isYaleEmail("kyle.jensen@yale.edu") {
		t.Errorf("Expected true, got false for kyle.jensen@yale.edu")
	}
	if isYaleEmail("kyle.jensen@harvard.edu") {
		t.Errorf("Expected false, got true for kyle.jensen@harvard.edu")
	}
}

func TestCanRSVP(t *testing.T) {
	t.Parallel()
	if !canRSVP("kyle.jensen@yale.edu") {
		t.Errorf("Expected true, got false for kyle.jensen@yale.edu")
	}
	if canRSVP("kyle.jensen@harvard.edu") {
		t.Errorf("Expected false, got true for kyle.jensen@harvard.edu")
	}
	if canRSVP("blahblahblah") {
		t.Errorf("Expected false, got true for blahblahblah")
	}
}

func TestConfirmationCodeForEmail(t *testing.T) {
	t.Parallel()
	email := "ehofeh@yale.edu"
	expected := "752950a"
	actual := confirmationCodeForEmail(email)
	if actual != expected {
		t.Errorf("Expected %s, got %s for %s", expected, actual, email)
	}
}
