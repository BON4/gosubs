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

type accUsecaseBoil struct {
	db *sql.DB
}

func NewBoilAccountUsecase(db *sql.DB) domain.AccountUsecase {
	return &accUsecaseBoil{
		db: db,
	}
}

func (c *accUsecaseBoil) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	acc := &boilmodels.Account{}
	err := boilmodels.Accounts(qm.Where("email=?", email), qm.Limit(1)).Bind(ctx, c.db, acc)
	if err != nil {
		return nil, err
	}

	domainAcc := &domain.Account{}
	domain.AccountBoilToDomain(acc, domainAcc)
	return domainAcc, nil
}

func (c *accUsecaseBoil) GetUser(ctx context.Context, id int64) (*domain.Tguser, error) {
	acc, err := boilmodels.FindAccount(ctx, c.db, id)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, errors.New("acc does not exist")
		// }
		return nil, err
	}

	user, err := boilmodels.FindTguser(ctx, c.db, acc.UserID.Int64)
	if err != nil {
		return nil, err
	}

	domainUser := &domain.Tguser{}
	domain.TguserBoilToDomain(user, domainUser)

	return domainUser, err
}

func (c *accUsecaseBoil) GetByID(ctx context.Context, id int64) (*domain.Account, error) {
	acc, err := boilmodels.FindAccount(ctx, c.db, id)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, errors.New("acc does not exist")
		// }
		return nil, err
	}
	domainAccount := &domain.Account{}
	domain.AccountBoilToDomain(acc, domainAccount)
	return domainAccount, nil
}

func (c *accUsecaseBoil) GetByTelegramID(ctx context.Context, id int64) (*domain.Account, error) {
	acc := &boilmodels.Account{}
	if err := boilmodels.Accounts(qm.Where("telegram_di=?", id), qm.Limit(1)).Bind(ctx, c.db, acc); err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, errors.New("acc does not exist")
		// }
		return nil, err
	}
	domainAccount := &domain.Account{}
	domain.AccountBoilToDomain(acc, domainAccount)
	return domainAccount, nil
}

func (c *accUsecaseBoil) Create(ctx context.Context, acc *domain.Account) error {
	if acc.UserID.Valid {
		if _, err := boilmodels.Accounts(qm.Where("user_id=?", acc.UserID)).One(ctx, c.db); err != nil {
			if err != sql.ErrNoRows {
				return err
			}
		} else {
			//TODO create custom errors
			return errors.New("account with this user_id already exist")
		}
	}

	boilAccount := &boilmodels.Account{}
	domain.AccountDomainToBoil(acc, boilAccount)

	// Insert
	if err := boilAccount.Insert(ctx, c.db, boil.Infer()); err != nil {
		return err
	}

	domain.AccountBoilToDomain(boilAccount, acc)

	return nil
}

// Delete - will delete acc. Subscriptions will be deleted also.
func (c *accUsecaseBoil) Delete(ctx context.Context, id int64) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//Delete user subscription
	if _, err := boilmodels.Subs(qm.Where("account_id=?", id)).DeleteAll(ctx, tx); err != nil {
		//TODO Rollback can cause error
		tx.Rollback()
		return err
	}

	// Delete Account
	if _, err := boilmodels.Accounts(qm.Where("account_id=?", id)).DeleteAll(ctx, tx); err != nil {
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

func (c *accUsecaseBoil) Update(ctx context.Context, acc *domain.Account) error {
	foundAccount, err := boilmodels.FindAccount(ctx, c.db, acc.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("acc does not exist")
		}
		return err
	}

	boilAccount := &boilmodels.Account{}
	domain.AccountDomainToBoil(acc, boilAccount)

	foundAccount.Password = boilAccount.Password
	foundAccount.Email = boilAccount.Email
	foundAccount.ChanName = boilAccount.ChanName
	foundAccount.AccountID = boilAccount.AccountID
	foundAccount.UserID = boilAccount.UserID
	foundAccount.Role = boilAccount.Role

	_, err = foundAccount.Update(ctx, c.db, boil.Infer())

	domain.AccountBoilToDomain(boilAccount, acc)

	return err
}

func (c *accUsecaseBoil) List(ctx context.Context, cond domain.FindAccountRequest) ([]*domain.Account, error) {
	baccs, err := boilmodels.Accounts(qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize))).All(ctx, c.db)
	if err != nil {
		return make([]*domain.Account, 0), err
	}

	domainAccounts := make([]*domain.Account, len(baccs))

	for i, acc := range baccs {
		domainAccounts[i] = &domain.Account{}
		domain.AccountBoilToDomain(acc, domainAccounts[i])
	}

	return domainAccounts, nil
}
