package usecase

import (
	"context"
	"database/sql"
	"errors"

	domain "github.com/BON4/gosubs/internal/domain"
	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type tgUserUsecaseBoil struct {
	db *sql.DB
}

func NewBoilTgUserUsecase(db *sql.DB) domain.TgUserUsecase {
	return &tgUserUsecaseBoil{
		db: db,
	}
}

func (u *tgUserUsecaseBoil) GetByID(ctx context.Context, id int64) (*domain.Tguser, error) {
	user, err := boilmodels.FindTguser(ctx, u.db, id)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, errors.New("user does not exist")
		// }
		return nil, err
	}

	domainUser := &domain.Tguser{}
	domain.TguserBoilToDomain(user, domainUser)
	return domainUser, nil
}

func (u *tgUserUsecaseBoil) GetByTelegramID(ctx context.Context, id int64) (*domain.Tguser, error) {
	user := &boilmodels.Tguser{}
	if err := boilmodels.Tgusers(qm.Where("telegram_di=?", id), qm.Limit(1)).Bind(ctx, u.db, user); err != nil {
		// if err != nil {
		// 	return nil, errors.New("user does not exist")
		// }
		return nil, err
	}

	domainUser := &domain.Tguser{}
	domain.TguserBoilToDomain(user, domainUser)
	return domainUser, nil
}

// Create - will create new user.
func (u *tgUserUsecaseBoil) Create(ctx context.Context, tguser *domain.Tguser) error {
	// IsExist. Check if user is already exist.
	if _, err := boilmodels.Tgusers(qm.Where("telegram_id=?", tguser.TelegramID)).One(ctx, u.db); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else {
		//TODO create custom errors
		return errors.New("already exist")
	}

	boilUser := &boilmodels.Tguser{}
	domain.TguserDomainToBoil(tguser, boilUser)

	// Insert
	if err := boilUser.Insert(ctx, u.db, boil.Infer()); err != nil {
		return err
	}

	domain.TguserBoilToDomain(boilUser, tguser)

	return nil
}

// Delete - will delete user. Subscription will be deleted also.
func (u *tgUserUsecaseBoil) Delete(ctx context.Context, id int64) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//Delete user subscription
	if _, err := boilmodels.Subs(qm.Where("user_id=?", id)).DeleteAll(ctx, tx); err != nil {
		//TODO Rollback can cause error
		tx.Rollback()
		return err
	}

	// Delete tguser
	if _, err := boilmodels.Tgusers(qm.Where("user_id=?", id)).DeleteAll(ctx, tx); err != nil {
		//TODO Rollback can cause error
		tx.Rollback()
		return err
	}

	//TODO Commit can cause error
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (u *tgUserUsecaseBoil) Update(ctx context.Context, tguser *domain.Tguser) error {
	user, err := boilmodels.FindTguser(ctx, u.db, tguser.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user does not exist")
		}
		return err
	}

	boilUser := &boilmodels.Tguser{}
	domain.TguserDomainToBoil(tguser, boilUser)

	user.Username = boilUser.Username
	user.TelegramID = boilUser.TelegramID
	user.Status = boilUser.Status

	_, err = user.Update(ctx, u.db, boil.Infer())

	domain.TguserBoilToDomain(user, tguser)
	return err
}

func (u *tgUserUsecaseBoil) List(ctx context.Context, cond domain.FindUserRequest) ([]*domain.Tguser, error) {
	busers, err := boilmodels.Tgusers(qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize))).All(ctx, u.db)
	if err != nil {
		return make([]*domain.Tguser, 0), err
	}

	domainUsers := make([]*domain.Tguser, len(busers))

	for i, user := range busers {
		domainUsers[i] = &domain.Tguser{}
		domain.TguserBoilToDomain(user, domainUsers[i])
	}

	return domainUsers, nil
}
