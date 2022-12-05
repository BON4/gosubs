package usecase

import (
	"context"
	"database/sql"
	"fmt"

	domain "github.com/BON4/gosubs/internal/domain"
	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	"github.com/sirupsen/logrus"

	myerrors "github.com/BON4/gosubs/internal/errors"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type subscriptionUsecaseBoil struct {
	db     *sql.DB
	logger *logrus.Entry
}

func NewBoilSubscriptionUsecase(db *sql.DB, logger *logrus.Entry) domain.SubscriptionUsecase {
	return &subscriptionUsecaseBoil{
		db:     db,
		logger: logger,
	}
}

func (s *subscriptionUsecaseBoil) GetByID(ctx context.Context, userID int64, creatorID int64) (*models.Sub, error) {
	found, err := models.FindSub(ctx, s.db, userID, creatorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Subscription with user_id: %d, does not exists. Detail: %w", userID, err)
		}
		return nil, err
	}

	return found, nil
}

// Create - creates subscribtion
func (s *subscriptionUsecaseBoil) Create(ctx context.Context, sub *models.Sub) error {
	_, err := models.FindSub(ctx, s.db, sub.UserID, sub.AccountID)

	//Have an error
	if err != nil {
		if err != sql.ErrNoRows {
			//Essantial error then return
			return err
		}

	} else {
		//Have no error means DB already have this object, we dont want this
		return fmt.Errorf("Subscription with user_id and acoount_id: %d - %d, already exists. Detail: %w", sub.UserID, sub.AccountID, myerrors.ErrAlreadyExists)
	}

	if err := sub.Insert(ctx, s.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// Save - saves subscription to history table.
func (s *subscriptionUsecaseBoil) Save(ctx context.Context, sub *models.Sub) (int64, error) {
	subhist := models.SubHistory{
		UserID:      sub.UserID,
		AccountID:   sub.AccountID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      sub.Status,
		Price:       sub.Price,
	}

	err := subhist.Insert(ctx, s.db, boil.Infer())
	return subhist.SubHistID, err
}

func (s *subscriptionUsecaseBoil) Update(ctx context.Context, sub *models.Sub) error {
	foundSub, err := models.FindSub(ctx, s.db, sub.UserID, sub.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Subscription with user_id and acoount_id: %d - %d, does not exists. Detail: %w", sub.UserID, sub.AccountID, err)
		}
		return err
	}

	foundSub.ActivatedAt = sub.ActivatedAt
	foundSub.ExpiresAt = sub.ExpiresAt
	foundSub.Status = sub.Status
	foundSub.Price = sub.Price

	_, err = foundSub.Update(ctx, s.db, boil.Infer())
	return err
}

func (s *subscriptionUsecaseBoil) Delete(ctx context.Context, userID int64, creatorID int64) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
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
				s.logger.Error("Rollback failed:", err)
			}
		}
	}()

	// Delete subscription
	_, err = models.Subs(qm.Where("user_id=? and account_id=?", userID, creatorID)).DeleteAll(ctx, tx)

	return err
}

func (s *subscriptionUsecaseBoil) List(ctx context.Context, cond domain.FindSubRequest) ([]*models.Sub, error) {
	var conds []qm.QueryMod = make([]qm.QueryMod, 0, 1)
	if cond.Price != nil {
		conds = append(conds, qm.Where("price=?", cond.Price.Eq))
	} else if cond.PriceRange != nil {
		if cond.PriceRange.From != nil {
			conds = append(conds, qm.Where("price>?", *cond.PriceRange.From))
		}

		if cond.PriceRange.To != nil {
			conds = append(conds, qm.Where("price<?", *cond.PriceRange.To))
		}
	}

	if cond.Status != nil {
		if cond.Status.Eq != nil {
			conds = append(conds, qm.Where("status=?", cond.Status.Eq))
		} else if cond.Status.Like != nil {
			conds = append(conds, qm.Where("status like ?%", cond.Status.Like))
		}
	}

	if cond.TgUserID != nil {
		conds = append(conds, qm.Where("user_id=?", cond.TgUserID.Eq))
	}

	if cond.AccountID != nil {
		conds = append(conds, qm.Where("account_id=?", cond.AccountID.Eq))
	}

	conds = append(conds, qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize)))

	bsubs, err := models.Subs(conds...).All(ctx, s.db)

	if err != nil {
		return make([]*models.Sub, 0), err
	}

	return bsubs, nil
}
