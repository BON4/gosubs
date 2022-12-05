package usecase

import (
	"context"
	"database/sql"
	"fmt"

	domain "github.com/BON4/gosubs/internal/domain"
	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	myerrors "github.com/BON4/gosubs/internal/errors"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type tgUserUsecaseBoil struct {
	db     *sql.DB
	logger *logrus.Entry
}

func NewBoilTgUserUsecase(db *sql.DB, logger *logrus.Entry) domain.TgUserUsecase {
	return &tgUserUsecaseBoil{
		db:     db,
		logger: logger,
	}
}

func (u *tgUserUsecaseBoil) GetByID(ctx context.Context, id int64) (*models.Tguser, error) {
	user, err := models.FindTguser(ctx, u.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with id: %d, does not exists. Detail: %w", id, err)
		}

		return nil, err
	}

	return user, nil
}

func (u *tgUserUsecaseBoil) GetByTelegramID(ctx context.Context, id int64) (*models.Tguser, error) {
	user := &models.Tguser{}
	if err := models.Tgusers(qm.Where("telegram_di=?", id), qm.Limit(1)).Bind(ctx, u.db, user); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with telegram_id: %d, does not exists. Detail: %w", id, err)
		}
		return nil, err
	}

	return user, nil
}

// Create - will create new user.
func (u *tgUserUsecaseBoil) Create(ctx context.Context, tguser *models.Tguser) error {
	// IsExist. Check if user is already exist.
	_, err := models.Tgusers(qm.Where("telegram_id=?", tguser.TelegramID)).One(ctx, u.db)

	//Have an error
	if err != nil {
		if err != sql.ErrNoRows {
			//Essantial error then return
			return err
		}
	} else {
		//Have no error means DB already have acount with this user_id, we dont want this
		return fmt.Errorf("User with telegram_id: %d, already exists. Detail: %w", tguser.TelegramID, myerrors.ErrAlreadyExists)
	}

	// Insert
	if err := tguser.Insert(ctx, u.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// Delete - will delete user. Subscription will be deleted also.
func (u *tgUserUsecaseBoil) Delete(ctx context.Context, id int64) error {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = tx.Commit()
		} else {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				u.logger.Error("Rollback failed:", err)
			}
		}
	}()

	//Delete user subscription
	if _, err := models.Subs(qm.Where("user_id=?", id)).DeleteAll(ctx, tx); err != nil {
		return err
	}

	// Delete tguser
	if _, err := models.Tgusers(qm.Where("user_id=?", id)).DeleteAll(ctx, tx); err != nil {
		return err
	}

	return nil
}

func (u *tgUserUsecaseBoil) Update(ctx context.Context, tguser *models.Tguser) error {
	user, err := models.FindTguser(ctx, u.db, tguser.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("User with id: %d, does not exists. Detail: %w", tguser.UserID, err)
		}
		return err
	}

	user.Username = tguser.Username
	user.TelegramID = tguser.TelegramID
	user.Status = tguser.Status

	_, err = user.Update(ctx, u.db, boil.Infer())

	return err
}

func (u *tgUserUsecaseBoil) List(ctx context.Context, cond domain.FindUserRequest) ([]*models.Tguser, error) {
	var conds []qm.QueryMod = make([]qm.QueryMod, 0, 1)

	if cond.Status != nil {
		if cond.Status.Eq != nil {
			conds = append(conds, qm.Where("status=?", *cond.Status.Eq))
		} else if cond.Status.Like != nil {
			conds = append(conds, qm.Where("status like ?%", *cond.Status.Like))
		}
	}

	if cond.Username != nil {
		if cond.Username.Eq != nil {
			conds = append(conds, qm.Where("username=?", *cond.Username.Eq))
		} else if cond.Username.Like != nil {
			conds = append(conds, qm.Where("username like ?%", *cond.Username.Like))
		}
	}

	conds = append(conds, qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize)))

	busers, err := models.Tgusers(conds...).All(ctx, u.db)
	if err != nil {
		return make([]*models.Tguser, 0), err
	}

	return busers, nil
}
