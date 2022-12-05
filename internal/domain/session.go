package domain

import (
	"time"

	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID      `json:"id"`
	Instance     models.Account `json:"instance"`
	RefreshToken string         `json:"refresh_token"`
	UserAgent    string         `json:"user_agent"`
	ClientIP     string         `json:"client_ip"`
	IsBlocked    bool           `json:"is_blocked"`
	ExpiresAt    time.Time      `json:"expires_at"`
}
