package models

import (
	"time"

	"github.com/guregu/null"
)

type Bank struct {
	ID        uint
	Name      string `json:"name"`
	IBAN      string `json:"iban"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}
