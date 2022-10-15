package domain

import (
	"database/sql"
	"time"

	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"
	sqlcmodels "github.com/BON4/gosubs/internal/domain/sqlc_postgres"
	"github.com/volatiletech/null/v8"
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

type Tguser struct {
	UserID     int64      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	TelegramID int64      `boil:"telegram_id" json:"telegram_id" toml:"telegram_id" yaml:"telegram_id"`
	Username   string     `boil:"username" json:"username" toml:"username" yaml:"username"`
	Status     UserStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
}

func TguserDomainToBoil(user *Tguser) *boilmodels.Tguser {
	return &boilmodels.Tguser{
		UserID:     user.UserID,
		TelegramID: user.TelegramID,
		Username:   user.Username,
		Status:     boilmodels.UserStatus(user.Status),
	}
}

func TguserDomainToSqlc(user *Tguser) *sqlcmodels.Tguser {
	return &boilmodels.Tguser{
		UserID:     user.UserID,
		TelegramID: user.TelegramID,
		Username:   user.Username,
		Status:     boilmodels.UserStatus(user.Status),
	}
}

func TguserBoilToDomain(user *boilmodels.Tguser) *Tguser {
	return &Tguser{
		UserID:     user.UserID,
		TelegramID: user.TelegramID,
		Username:   user.Username,
		Status:     UserStatus(user.Status),
	}
}

func TguserSqlcToDomain(user *sqlcmodels.Tguser) *Tguser {
	return &Tguser{
		UserID:     user.UserID,
		TelegramID: user.TelegramID,
		Username:   user.Username,
		Status:     UserStatus(user.Status),
	}
}

type Creator struct {
	CreatorID  int64       `boil:"creator_id" json:"creator_id" toml:"creator_id" yaml:"creator_id"`
	TelegramID int64       `boil:"telegram_id" json:"telegram_id" toml:"telegram_id" yaml:"telegram_id"`
	Username   string      `boil:"username" json:"username" toml:"username" yaml:"username"`
	Password   null.Bytes  `boil:"password" json:"password,omitempty" toml:"password" yaml:"password,omitempty"`
	Email      null.String `boil:"email" json:"email,omitempty" toml:"email" yaml:"email,omitempty"`
	ChanName   null.String `boil:"chan_name" json:"chan_name,omitempty" toml:"chan_name" yaml:"chan_name,omitempty"`
}

func CreatorDomainToBoil(creator *Creator) *boilmodels.Creator {
	return &boilmodels.Creator{
		CreatorID:  creator.CreatorID,
		TelegramID: creator.TelegramID,
		Username:   creator.Username,
		Password:   creator.Password,
		Email:      creator.Email,
		ChanName:   creator.ChanName,
	}
}

func CreatorBoilToDomain(creator *boilmodels.Creator) *Creator {
	return &Creator{
		CreatorID:  creator.CreatorID,
		TelegramID: creator.TelegramID,
		Username:   creator.Username,
		Password:   creator.Password,
		Email:      creator.Email,
		ChanName:   creator.ChanName,
	}
}

func CreatorDomainToSqlc(creator *Creator) *sqlcmodels.Creator {
	return &boilmodels.Creator{
		CreatorID:  creator.CreatorID,
		TelegramID: creator.TelegramID,
		Username:   creator.Username,
		Password:   creator.Password,
		Email:      creator.Email,
		ChanName:   creator.ChanName,
	}
}

func CreatorSqlcToDomain(creator *sqlcmodels.Creator) *Creator {
	return &Creator{
		CreatorID:  creator.CreatorID,
		TelegramID: creator.TelegramID,
		Username:   creator.Username,
		Password:   null.NewBytes(creator.Password, len(creator.Password) < 0),
		Email:      null.NewString(creator.Email.String, creator.Email.Valid),
		ChanName:   null.NewString(creator.ChanName.String, creator.ChanName.Valid),
	}
}

type Sub struct {
	UserID      int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatorID   int64     `boil:"creator_id" json:"creator_id" toml:"creator_id" yaml:"creator_id"`
	ActivatedAt time.Time `boil:"activated_at" json:"activated_at" toml:"activated_at" yaml:"activated_at"`
	ExpiresAt   time.Time `boil:"expires_at" json:"expires_at" toml:"expires_at" yaml:"expires_at"`
	Status      SubStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
	Price       null.Int  `boil:"price" json:"price,omitempty" toml:"price" yaml:"price,omitempty"`
}

func SubDomainToBoil(sub *Sub) *boilmodels.Sub {
	return &boilmodels.Sub{
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      boilmodels.SubStatus(sub.Status),
		Price:       sub.Price,
	}
}

func SubBoilToDomain(sub *boilmodels.Sub) *Sub {
	return &Sub{
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      SubStatus(sub.Status),
		Price:       sub.Price,
	}
}

func SubDomainToSqlc(sub *Sub) *sqlcmodels.Sub {
	return &sqlcmodels.Sub{
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      sqlcmodels.SubStatus(sub.Status),
		Price: sql.NullInt32{
			Int32: int32(sub.Price.Int),
			Valid: sub.Price.Valid,
		},
	}
}

func SubSqlcToDomain(sub *sqlcmodels.Sub) *Sub {
	return &Sub{
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      SubStatus(sub.Status),
		Price:       null.NewInt(int(sub.Price.Int32), sub.Price.Valid),
	}
}

type SubStatus string

// Enum values for SubStatus
const (
	SubStatusExpired   SubStatus = "expired"
	SubStatusActive    SubStatus = "active"
	SubStatusCancelled SubStatus = "cancelled"
)

type SubHistory struct {
	UserID      int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatorID   int64     `boil:"creator_id" json:"creator_id" toml:"creator_id" yaml:"creator_id"`
	ActivatedAt time.Time `boil:"activated_at" json:"activated_at" toml:"activated_at" yaml:"activated_at"`
	ExpiresAt   time.Time `boil:"expires_at" json:"expires_at" toml:"expires_at" yaml:"expires_at"`
	Status      SubStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
	Price       null.Int  `boil:"price" json:"price,omitempty" toml:"price" yaml:"price,omitempty"`
	SubHistID   int64     `boil:"sub_hist_id" json:"sub_hist_id" toml:"sub_hist_id" yaml:"sub_hist_id"`
}

func SubHistoryDomainToBoil(sub *SubHistory) *boilmodels.SubHistory {
	return &boilmodels.SubHistory{
		SubHistID:   sub.SubHistID,
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      boilmodels.SubStatus(sub.Status),
		Price:       sub.Price,
	}
}

func SubHistoryBoilToDomain(sub *boilmodels.SubHistory) *SubHistory {
	return &SubHistory{
		SubHistID:   sub.SubHistID,
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      SubStatus(sub.Status),
		Price:       sub.Price,
	}
}

func SubHistoryDomainToSqlc(sub *SubHistory) *sqlcmodels.SubHistory {
	return &sqlcmodels.SubHistory{
		SubHistID:   sub.SubHistID,
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      sqlcmodels.SubStatus(sub.Status),
		Price: sql.NullInt32{
			Int32: int32(sub.Price.Int),
			Valid: sub.Price.Valid,
		},
	}
}

func SubHistorySqlcToDomain(sub *sqlcmodels.SubHistory) *SubHistory {
	return &SubHistory{
		SubHistID:   sub.SubHistID,
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      SubStatus(sub.Status),
		Price:       null.NewInt(int(sub.Price.Int32), sub.Price.Valid),
	}
}
