package models

import (
	"time"
)

// User describes the shape of the user model
type User struct {
	// Id is a unique reference for the User
	Id            string     `json:"id,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	Name          string     `json:"name,omitempty"`
	Email         string     `json:"email,omitempty"`
	EmailVerified *time.Time `json:"emailVerified,omitempty"`
	Posts         []*Post    `json:"posts,omitempty"`
}
