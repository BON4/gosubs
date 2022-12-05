package tokengen

import (
	"time"

	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
)

// Maker is an interface for managing tokens
type Generator interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(account *models.Account, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
