package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	Title       string
	Description *string
	Stock       int
	CreatedAt   time.Time
}

func NewProduct(title string, desc *string, stock int) *Product {
	return &Product{
		Title:       title,
		Description: desc,
		Stock:       stock,
	}
}
