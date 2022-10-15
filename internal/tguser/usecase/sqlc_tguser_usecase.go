package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BON4/gosubs/internal/domain"
	sqlcmodels "github.com/BON4/gosubs/internal/domain/sqlc_postgres"
)

type tgUserUsecaseSqlc struct {
	db sqlcmodels.Store
}

func NewSqlcTguserUsecase(db *sql.DB) domain.TgUserUsecase {
	return &tgUserUsecaseSqlc{
		db: sqlcmodels.NewStore(db),
	}
}

func (u *tgUserUsecaseSqlc) GetByID(ctx context.Context, id int64) (*domain.Tguser, error) {

	user, err := u.db.FindTguserID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user does not exist")
		}
		return nil, err
	}
	return domain.TguserSqlcToDomain(user), nil
}

func (u *tgUserUsecaseSqlc) GetByTelegramID(ctx context.Context, id int64) (*domain.Tguser, error) {
	user, err := u.db.FindTguserTelegramID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user does not exist")
		}
		return nil, err
	}
	return domain.TguserSqlcToDomain(user), nil
}

// Create - will create new user.
func (u *tgUserUsecaseSqlc) Create(ctx context.Context, tguser *domain.Tguser) error {
	if ok, err := u.db.IsExistTguser(ctx, tguser.TelegramID); err != nil {
		return err
	} else if ok {
		return errors.New("already exist")
	}

	sqlcUser := domain.TguserDomainToSqlc(tguser)

	var err error
	// Insert
	sqlcUser, err = u.db.InsertTguser(ctx, sqlcmodels.InsertTguserParams{
		TelegramID: sqlcUser.TelegramID,
		Username:   sqlcUser.Username,
		Status:     sqlcUser.Status,
	})

	if err != nil {
		return err
	}

	tguser = domain.TguserSqlcToDomain(sqlcUser)

	return nil
}

// Delete - will delete user. Subscription will be deleted also.
func (u *tgUserUsecaseSqlc) Delete(ctx context.Context, id int64) error {
	return u.db.DeleteUserTx(ctx, sqlcmodels.DeleteUserTxParams{
		UserID: id,
	})
}

func (u *tgUserUsecaseSqlc) Update(ctx context.Context, tguser *domain.Tguser) error {
	_, err := u.db.FindTguserID(ctx, tguser.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user does not exist")
		}
		return err
	}

	sqlcUser := domain.TguserDomainToSqlc(tguser)

	sqlcUser, err = u.db.UpdateTguser(ctx, sqlcmodels.UpdateTguserParams{
		TelegramID: sql.NullInt64{
			Int64: sqlcUser.TelegramID,
			Valid: true,
		},
		Username: sql.NullString{
			String: sqlcUser.Username,
			Valid:  true,
		},
		Status: sqlcmodels.NullUserStatus{
			UserStatus: sqlcUser.Status,
			Valid:      true,
		},
		UserID: sqlcUser.UserID,
	})

	tguser = domain.TguserSqlcToDomain(sqlcUser)
	return err
}

func (u *tgUserUsecaseSqlc) List(ctx context.Context, cond domain.FindUserRequest) ([]*domain.Tguser, error) {
	susers, err := u.db.ListTguser(ctx, sqlcmodels.ListTguserParams{
		Offset: int32(cond.PageSettings.PageNumber),
		Limit:  int32(cond.PageSettings.PageSize),
	})

	if err != nil {
		return make([]*domain.Tguser, 0), err
	}

	domainUsers := make([]*domain.Tguser, len(susers))

	for i, user := range susers {
		domainUsers[i] = domain.TguserSqlcToDomain(user)
	}

	return domainUsers, nil
}
