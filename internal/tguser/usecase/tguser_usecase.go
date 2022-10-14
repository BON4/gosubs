package usecase

import (
	"context"
	"database/sql"
	"errors"

	models "github.com/BON4/gosubs/internal/domain/postgres"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type tgUserUsecase struct {
	db *sql.DB
}

func NewTgUserUsecase(db *sql.DB) models.TgUserUsecase {
	return &tgUserUsecase{
		db: db,
	}
}

func (u *tgUserUsecase) GetByID(ctx context.Context, id int64) (*models.Tguser, error) {
	user, err := models.FindTguser(ctx, u.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user does not exist")
		}
		return nil, err
	}
	return user, nil
}

func (u *tgUserUsecase) GetByTelegramID(ctx context.Context, id int64) (*models.Tguser, error) {
	user := &models.Tguser{}
	if err := models.Tgusers(qm.Where("telegram_di=?", id), qm.Limit(1)).Bind(ctx, u.db, user); err != nil {
		if err != nil {
			return nil, errors.New("user does not exist")
		}
		return nil, err
	}

	return user, nil
}

// Create - will create new user.
func (u *tgUserUsecase) Create(ctx context.Context, tguser *models.Tguser) error {
	// IsExist. Check if user is already exist.
	if _, err := models.Tgusers(qm.Where("telegram_id=?", tguser.TelegramID)).One(ctx, u.db); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else {
		//TODO create custom errors
		return errors.New("already exist")
	}

	// Insert
	if err := tguser.Insert(ctx, u.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// Delete - will delete user. Subscription will be deleted also.
func (u *tgUserUsecase) Delete(ctx context.Context, id int64) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//Delete user subscription
	if _, err := models.Subs(qm.Where("user_id=?", id)).DeleteAll(ctx, tx); err != nil {
		//TODO Rollback can cause error
		tx.Rollback()
		return err
	}

	// Delete tguser
	if _, err := models.Tgusers(qm.Where("user_id=?", id)).DeleteAll(ctx, tx); err != nil {
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

func (u *tgUserUsecase) Update(ctx context.Context, tguser *models.Tguser) error {
	user, err := models.FindTguser(ctx, u.db, tguser.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user does not exist")
		}
		return err
	}

	user.Status = tguser.Status
	user.Username = tguser.Username
	user.TelegramID = tguser.TelegramID

	_, err = user.Update(ctx, u.db, boil.Infer())
	return err
}

func (u *tgUserUsecase) List(ctx context.Context, cond models.FindUserRequest) ([]*models.Tguser, error) {
	return models.Tgusers(qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize))).All(ctx, u.db)
}
