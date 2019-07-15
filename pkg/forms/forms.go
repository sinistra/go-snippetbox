package forms

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0]))")

// Declare a struct to hold the form values (and also a map to hold any validation failure messages).
type NewSnippet struct {
	Title    string
	Content  string
	Expires  string
	Failures map[string]string
}

// Implement an Valid() method which carries out validation checks on the form
// fields and returns true if there are no failures.
func (f *NewSnippet) Valid() bool {
	f.Failures = make(map[string]string)
	// Check that the Title field is not blank and is not more than 100 characters long.
	// If it fails either of those checks, add a message to the f.Failures
	// map using the field name as the key.
	if strings.TrimSpace(f.Title) == "" {
		f.Failures["Title"] = "Title is required"
	} else if utf8.RuneCountInString(f.Title) > 100 {
		f.Failures["Title"] = "Title cannot be longer than 100 characters"
	}
	// Validate the Content and Expires fields aren't blank in a similar way.
	if strings.TrimSpace(f.Content) == "" {
		f.Failures["Content"] = "Content is required"
	}
	// Check that the Expires field isn't blank and is one of a fixed list.
	// Using a lookup on a map keyed with the permitted options and values of true is a
	// neat trick which saves you looping over the permitted values.
	permitted := map[string]bool{
		"3600":     true,
		"86400":    true,
		"31536000": true,
	}

	if strings.TrimSpace(f.Expires) == "" {
		f.Failures["Expires"] = "Expiry time is required"
	} else if !permitted[f.Expires] {
		f.Failures["Expires"] = "Expiry time must be 3600, 86400 or 31536000 seconds"
	}
	// If there are no failure messages, return true.
	return len(f.Failures) == 0
}

type SignupUser struct {
	Name     string
	Email    string
	Password string
	Failures map[string]string
}

func (f *SignupUser) Valid() bool {
	f.Failures = make(map[string]string)
	if strings.TrimSpace(f.Name) == "" {
		f.Failures["Name"] = "Name is required"
	}
	if strings.TrimSpace(f.Email) == "" {
		f.Failures["Email"] = "Email is required"
	} else if len(f.Email) > 254 || !rxEmail.MatchString(f.Email) {
		f.Failures["Email"] = "Email is not a valid address"
	}
	if utf8.RuneCountInString(f.Password) < 8 {
		f.Failures["Password"] = "Password cannot be shorter than 8 characters"
	}
	return len(f.Failures) == 0
}

type LoginUser struct {
	Email    string
	Password string
	Failures map[string]string
}

func (f *LoginUser) Valid() bool {
	f.Failures = make(map[string]string)
	if strings.TrimSpace(f.Email) == "" {
		f.Failures["Email"] = "Email is required"
	}
	if strings.TrimSpace(f.Password) == "" {
		f.Failures["Password"] = "Password is required"
	}
	return len(f.Failures) == 0
}
