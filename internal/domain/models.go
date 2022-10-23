package domain

import (
	"time"

	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"
	"github.com/google/uuid"
	null "github.com/volatiletech/null/v8"
)

type UserStatus string

// Enum values for UserStatus
const (
	UserStatusCreator       UserStatus = "creator"
	UserStatusAdministrator UserStatus = "administrator"
	UserStatusMember        UserStatus = "member"
	UserStatusRestricted    UserStatus = "restricted"
	UserStatusLeft          UserStatus = "left"
	UserStatusKicked        UserStatus = "kicked"
)

// Tguser is an object representing the database table.
type Tguser struct {
	UserID     int64      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	TelegramID int64      `boil:"telegram_id" json:"telegram_id" toml:"telegram_id" yaml:"telegram_id"`
	Username   string     `boil:"username" json:"username" toml:"username" yaml:"username"`
	Status     UserStatus `boil:"status" json:"status" toml:"status" yaml:"status" swaggertype:"string"`
}

func TguserDomainToBoil(user *Tguser, userout *boilmodels.Tguser) {
	if userout != nil {
		userout.UserID = user.UserID
		userout.TelegramID = user.TelegramID
		userout.Username = user.Username
		userout.Status = boilmodels.UserStatus(user.Status)
	}
}

func TguserBoilToDomain(user *boilmodels.Tguser, userout *Tguser) {
	if userout != nil {
		userout.UserID = user.UserID
		userout.TelegramID = user.TelegramID
		userout.Username = user.Username
		userout.Status = UserStatus(user.Status)
	}
}

type AccountRole string

// Enum values for AccountRole
const (
	AccountRoleCreator AccountRole = "creator"
	AccountRoleAdmin   AccountRole = "admin"
	AccountRoleBot     AccountRole = "bot"
)

// Account is an object representing the database table.
type Account struct {
	AccountID int64       `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	Password  []byte      `boil:"password" json:"password" toml:"password" yaml:"password"`
	Email     string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	Role      AccountRole `boil:"role" json:"role" toml:"role" yaml:"role"`
	ChanName  null.String `boil:"chan_name" json:"chan_name,omitempty" toml:"chan_name" yaml:"chan_name,omitempty" swaggertype:"string"`
	UserID    null.Int64  `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty" swaggertype:"integer"`
}

func AccountDomainToBoil(creator *Account, creatorout *boilmodels.Account) {
	if creatorout != nil {
		creatorout.AccountID = creator.AccountID
		creatorout.Password = creator.Password
		creatorout.Role = boilmodels.AccountRole(creator.Role)
		creatorout.Email = creator.Email
		creatorout.ChanName = creator.ChanName
		creatorout.UserID = creator.UserID
	}
}

func AccountBoilToDomain(creator *boilmodels.Account, creatorout *Account) {
	if creatorout != nil {
		creatorout.AccountID = creator.AccountID
		creatorout.Password = creator.Password
		creatorout.Email = creator.Email
		creatorout.ChanName = creator.ChanName
		creatorout.Role = AccountRole(creator.Role)
		creatorout.UserID = creator.UserID
	}
}

// Sub is an object representing the database table.
type Sub struct {
	UserID      int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	AccountID   int64     `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	ActivatedAt time.Time `boil:"activated_at" json:"activated_at" toml:"activated_at" yaml:"activated_at"`
	ExpiresAt   time.Time `boil:"expires_at" json:"expires_at" toml:"expires_at" yaml:"expires_at"`
	Status      SubStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
	Price       null.Int  `boil:"price" json:"price,omitempty" toml:"price" yaml:"price,omitempty" swaggertype:"integer"`
}

func SubDomainToBoil(sub *Sub, subout *boilmodels.Sub) {
	if subout != nil {
		subout.UserID = sub.UserID
		subout.AccountID = sub.AccountID
		subout.ActivatedAt = sub.ActivatedAt
		subout.ExpiresAt = sub.ExpiresAt
		subout.Status = boilmodels.SubStatus(sub.Status)
		subout.Price = sub.Price
	}
}

func SubBoilToDomain(sub *boilmodels.Sub, subout *Sub) {
	if subout != nil {
		subout.UserID = sub.UserID
		subout.AccountID = sub.AccountID
		subout.ActivatedAt = sub.ActivatedAt
		subout.ExpiresAt = sub.ExpiresAt
		subout.Status = SubStatus(sub.Status)
		subout.Price = sub.Price
	}
}

type SubStatus string

// Enum values for SubStatus
const (
	SubStatusExpired   SubStatus = "expired"
	SubStatusActive    SubStatus = "active"
	SubStatusCancelled SubStatus = "cancelled"
)

// SubHistory is an object representing the database table.
type SubHistory struct {
	UserID      int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	AccountID   int64     `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	ActivatedAt time.Time `boil:"activated_at" json:"activated_at" toml:"activated_at" yaml:"activated_at"`
	ExpiresAt   time.Time `boil:"expires_at" json:"expires_at" toml:"expires_at" yaml:"expires_at"`
	Status      SubStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
	Price       null.Int  `boil:"price" json:"price,omitempty" toml:"price" yaml:"price,omitempty" swaggertype:"integer"`
	SubHistID   int64     `boil:"sub_hist_id" json:"sub_hist_id" toml:"sub_hist_id" yaml:"sub_hist_id" swaggertype:"integer"`
}

func SubHistoryDomainToBoil(sub *SubHistory, subout *boilmodels.SubHistory) {
	if subout != nil {
		subout.SubHistID = sub.SubHistID
		subout.UserID = sub.UserID
		subout.AccountID = sub.AccountID
		subout.ActivatedAt = sub.ActivatedAt
		subout.ExpiresAt = sub.ExpiresAt
		subout.Status = boilmodels.SubStatus(sub.Status)
		subout.Price = sub.Price
	}
}

func SubHistoryBoilToDomain(sub *boilmodels.SubHistory, subout *SubHistory) {
	if subout != nil {
		subout.SubHistID = sub.SubHistID
		subout.UserID = sub.UserID
		subout.AccountID = sub.AccountID
		subout.ActivatedAt = sub.ActivatedAt
		subout.ExpiresAt = sub.ExpiresAt
		subout.Status = SubStatus(sub.Status)
		subout.Price = sub.Price
	}
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Instance     Account   `json:"instance"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIP     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}
