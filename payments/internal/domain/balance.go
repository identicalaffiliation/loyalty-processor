package domain

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	UserID    uuid.UUID
	Bonuses   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBalance(id uuid.UUID, bonuses int64) *Balance {
	return &Balance{
		UserID:  id,
		Bonuses: bonuses,
	}
}
