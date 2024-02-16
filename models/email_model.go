package models

import (
	"regexp"
)
type Email struct {
	Email string `json:"email" validate:"required"`
}

func (e *Email) IsEmailValid() bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return emailRegex.MatchString(e.Email)
}