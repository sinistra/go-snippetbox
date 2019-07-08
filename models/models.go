package models

import (
	"time"
)

// Define a Snippet type to hold the information about an individual snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// For convenience we also define a Snippets type, which is a slice for holding multiple Snippet objects.
type Snippets []*Snippet
