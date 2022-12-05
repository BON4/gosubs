package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domain "github.com/BON4/gosubs/internal/domain"
	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	myerrors "github.com/BON4/gosubs/internal/errors"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	ErrNoAccounts = errors.New("No accounts has been found")
)

type accUsecaseBoil struct {
	db     *sql.DB
	logger *logrus.Entry
}

func NewBoilAccountUsecase(db *sql.DB, logger *logrus.Entry) domain.AccountUsecase {
	return &accUsecaseBoil{
		logger: logger,
		db:     db,
	}
}

func (c *accUsecaseBoil) GetByEmail(ctx context.Context, email string) (*models.Account, error) {
	acc := &models.Account{}
	err := models.Accounts(qm.Where("email=?", email), qm.Limit(1)).Bind(ctx, c.db, acc)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Account with email: %s, does not exists. Detail: %w", email, err)
		}
		return nil, err
	}

	return acc, nil
}

func (c *accUsecaseBoil) GetUser(ctx context.Context, id int64) (*models.Tguser, error) {
	acc, err := models.FindAccount(ctx, c.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Account with user_id: %d, does not exists. Detail: %w", id, err)
		}
		return nil, err
	}

	user, err := models.FindTguser(ctx, c.db, acc.UserID.Int64)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (c *accUsecaseBoil) GetByID(ctx context.Context, id int64) (*models.Account, error) {
	acc, err := models.FindAccount(ctx, c.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Account with id: %d, does not exists. Detail: %w", id, err)
		}
		return nil, err
	}
	return acc, nil
}

func (c *accUsecaseBoil) GetByTelegramID(ctx context.Context, id int64) (*models.Account, error) {
	acc := &models.Account{}
	if err := models.Accounts(qm.Where("telegram_di=?", id), qm.Limit(1)).Bind(ctx, c.db, acc); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Account with telegram_id: %d, does not exists. Detail: %w", id, err)
		}
		return nil, err
	}
	return acc, nil
}

func (c *accUsecaseBoil) Create(ctx context.Context, acc *models.Account) error {
	if acc.UserID.Valid {
		_, err := models.Accounts(qm.Where("user_id=?", acc.UserID)).One(ctx, c.db)

		//Have an error
		if err != nil {
			if err != sql.ErrNoRows {
				//Essantial error then return
				return err
			}

		} else {
			//Have no error means DB already have this object, we dont want this
			return fmt.Errorf("Account with user_id: %d, already exists. Detail: %w", acc.UserID.Int64, myerrors.ErrAlreadyExists)
		}
	}

	// Insert
	if err := acc.Insert(ctx, c.db, boil.Infer()); err != nil {
		return err
	}

	return nil

}

// Delete - will delete acc. Subscriptions will be deleted also.
func (c *accUsecaseBoil) Delete(ctx context.Context, id int64) error {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{
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
				c.logger.Error("Rollback failed:", err)
			}
		}
	}()

	//Delete user subscription
	_, err = models.Subs(qm.Where("account_id=?", id)).DeleteAll(ctx, tx)
	if err != nil {
		return err
	}

	// Delete Account
	_, err = models.Accounts(qm.Where("account_id=?", id)).DeleteAll(ctx, tx)
	if err != nil {
		return err
	}

	return nil
}

func (c *accUsecaseBoil) Update(ctx context.Context, acc *models.Account) error {
	foundAccount, err := models.FindAccount(ctx, c.db, acc.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Account with id: %d, does not exists. Detail: %w", acc.AccountID, err)
		}
		return err
	}

	foundAccount.Password = acc.Password
	foundAccount.Email = acc.Email
	foundAccount.ChanName = acc.ChanName
	foundAccount.AccountID = acc.AccountID
	foundAccount.UserID = acc.UserID
	foundAccount.Role = acc.Role

	_, err = foundAccount.Update(ctx, c.db, boil.Infer())
	return err
}

func (c *accUsecaseBoil) List(ctx context.Context, cond domain.FindAccountRequest) ([]*models.Account, error) {
	var conds []qm.QueryMod = make([]qm.QueryMod, 0, 1)

	if cond.Role != nil {
		if cond.Role.Eq != nil {
			conds = append(conds, qm.Where("role=?", *cond.Role.Eq))
		} else if cond.Role.Like != nil {
			conds = append(conds, qm.Where("role like ?%", *cond.Role.Like))
		}
	}

	conds = append(conds, qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize)))

	baccs, err := models.Accounts(conds...).All(ctx, c.db)
	if err != nil {
		return make([]*models.Account, 0), err
	}

	return baccs, nil
}
