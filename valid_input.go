package main

import (
	"crypto/sha256"
	"fmt"
	"net/mail"
	"strings"
)

// isValidEmail - returns whether an email string is valid
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// isYaleEmail - returns whether an input string has suffix @yale.edu
func isYaleEmail(email string) bool {
	return strings.HasSuffix(strings.ToLower(email), "@yale.edu")
}

// canRSVP - returns whether an input email can RSVP
func canRSVP(email string) bool {
	return isValidEmail(email) && isYaleEmail(email)
}

// confirmationCodeForEmail - returns confirmation code of an input email
func confirmationCodeForEmail(email string) string {
	hash := sha256.Sum256([]byte(email))
	return fmt.Sprintf("%x", hash)[:7]
}
