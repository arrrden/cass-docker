package models

import (
	"time"
)

// Post describes the shape of the post model
type Post struct {
	// Id is a unique reference for the Post
	Id        string     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Text      string     `json:"text,omitempty"`
	User      *User      `json:"user,omitempty"`
}
