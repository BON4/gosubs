// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type SubStatus string

const (
	SubStatusExpired   SubStatus = "expired"
	SubStatusActive    SubStatus = "active"
	SubStatusCancelled SubStatus = "cancelled"
	SubStatusInactive  SubStatus = "inactive"
)

func (e *SubStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SubStatus(s)
	case string:
		*e = SubStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for SubStatus: %T", src)
	}
	return nil
}

type NullSubStatus struct {
	SubStatus SubStatus
	Valid     bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullSubStatus) Scan(value interface{}) error {
	if value == nil {
		ns.SubStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.SubStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullSubStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.SubStatus, nil
}

type UserStatus string

const (
	UserStatusCreator       UserStatus = "creator"
	UserStatusAdministrator UserStatus = "administrator"
	UserStatusMember        UserStatus = "member"
	UserStatusRestricted    UserStatus = "restricted"
	UserStatusLeft          UserStatus = "left"
	UserStatusKicked        UserStatus = "kicked"
)

func (e *UserStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserStatus(s)
	case string:
		*e = UserStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UserStatus: %T", src)
	}
	return nil
}

type NullUserStatus struct {
	UserStatus UserStatus
	Valid      bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UserStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UserStatus, nil
}

type Creator struct {
	CreatorID  int64          `db:"creator_id"`
	TelegramID int64          `db:"telegram_id"`
	Username   string         `db:"username"`
	Password   []byte         `db:"password"`
	Email      sql.NullString `db:"email"`
	ChanName   sql.NullString `db:"chan_name"`
}

type Sub struct {
	UserID      int64         `db:"user_id"`
	CreatorID   int64         `db:"creator_id"`
	ActivatedAt time.Time     `db:"activated_at"`
	ExpiresAt   time.Time     `db:"expires_at"`
	Status      SubStatus     `db:"status"`
	Price       sql.NullInt32 `db:"price"`
}

type SubHistory struct {
	UserID      int64         `db:"user_id"`
	CreatorID   int64         `db:"creator_id"`
	ActivatedAt time.Time     `db:"activated_at"`
	ExpiresAt   time.Time     `db:"expires_at"`
	Status      SubStatus     `db:"status"`
	Price       sql.NullInt32 `db:"price"`
	SubHistID   int64         `db:"sub_hist_id"`
}

type Tguser struct {
	UserID     int64      `db:"user_id"`
	TelegramID int64      `db:"telegram_id"`
	Username   string     `db:"username"`
	Status     UserStatus `db:"status"`
}