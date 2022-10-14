package usecase

import (
	"context"
	"database/sql"
	"errors"

	models "github.com/BON4/gosubs/internal/domain/postgres"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type subscriptionUsecase struct {
	db *sql.DB
}

func NewSubscriptionUsecase(db *sql.DB) models.SubscriptionUsecase {
	return &subscriptionUsecase{
		db: db,
	}
}

// Create - creates subscribtion
func (s *subscriptionUsecase) Create(ctx context.Context, sub *models.Sub) error {
	if _, err := models.FindSub(ctx, s.db, sub.UserID, sub.CreatorID); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else {
		return errors.New("subscription for this user is already exist")
	}

	if err := sub.Insert(ctx, s.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// Save - saves subscription to history table.
func (s *subscriptionUsecase) Save(ctx context.Context, sub *models.Sub) error {
	subhist := models.SubHistory{
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      sub.Status,
		Price:       sub.Price,
	}

	return subhist.Insert(ctx, s.db, boil.Infer())
}

func (s *subscriptionUsecase) Update(ctx context.Context, sub *models.Sub) error {
	foundSub, err := models.FindSub(ctx, s.db, sub.UserID, sub.CreatorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user does not exist")
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

func (s *subscriptionUsecase) Delete(ctx context.Context, userID int64, creatorID int64) error {
	// Delete subscription
	_, err := models.Subs(qm.Where("user_id=? and creator_id=?", userID, creatorID)).DeleteAll(ctx, s.db)
	return err
}

func (s *subscriptionUsecase) List(ctx context.Context, cond models.FindSubRequest) ([]*models.Sub, error) {
	var conds []qm.QueryMod = make([]qm.QueryMod, 0, 1)
	if cond.Price != nil {
		if cond.Price.Eq != nil {
			conds = append(conds, qm.Where("price=?", *cond.Price.Eq))
		} else if cond.Price.Range != nil {
			conds = append(conds, qm.Where("price>?", *cond.Price.Range.From))
			conds = append(conds, qm.Where("price<?", *cond.Price.Range.To))
		}
	}

	if cond.Status != nil {
		conds = append(conds, qm.Where("status=?", cond.Status.Eq))
	}

	if cond.CreatorID != nil {
		conds = append(conds, qm.Where("status=?", cond.CreatorID.Eq))
	}

	conds = append(conds, qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize)))

	return models.Subs(conds...).All(ctx, s.db)
}
