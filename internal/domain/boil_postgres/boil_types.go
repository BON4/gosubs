// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

type AccountRole string

// Enum values for AccountRole
const (
	AccountRoleCreator AccountRole = "creator"
	AccountRoleAdmin   AccountRole = "admin"
	AccountRoleBot     AccountRole = "bot"
)

func AllAccountRole() []AccountRole {
	return []AccountRole{
		AccountRoleCreator,
		AccountRoleAdmin,
		AccountRoleBot,
	}
}

func (e AccountRole) IsValid() error {
	switch e {
	case AccountRoleCreator, AccountRoleAdmin, AccountRoleBot:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e AccountRole) String() string {
	return string(e)
}

type SubStatus string

// Enum values for SubStatus
const (
	SubStatusExpired   SubStatus = "expired"
	SubStatusActive    SubStatus = "active"
	SubStatusCancelled SubStatus = "cancelled"
	SubStatusInactive  SubStatus = "inactive"
)

func AllSubStatus() []SubStatus {
	return []SubStatus{
		SubStatusExpired,
		SubStatusActive,
		SubStatusCancelled,
		SubStatusInactive,
	}
}

func (e SubStatus) IsValid() error {
	switch e {
	case SubStatusExpired, SubStatusActive, SubStatusCancelled, SubStatusInactive:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e SubStatus) String() string {
	return string(e)
}

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

func AllUserStatus() []UserStatus {
	return []UserStatus{
		UserStatusCreator,
		UserStatusAdministrator,
		UserStatusMember,
		UserStatusRestricted,
		UserStatusLeft,
		UserStatusKicked,
	}
}

func (e UserStatus) IsValid() error {
	switch e {
	case UserStatusCreator, UserStatusAdministrator, UserStatusMember, UserStatusRestricted, UserStatusLeft, UserStatusKicked:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e UserStatus) String() string {
	return string(e)
}
