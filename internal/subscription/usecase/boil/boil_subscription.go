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

type subscriptionUsecaseBoil struct {
	db *sql.DB
}

func NewBoilSubscriptionUsecase(db *sql.DB) domain.SubscriptionUsecase {
	return &subscriptionUsecaseBoil{
		db: db,
	}
}

// Create - creates subscribtion
func (s *subscriptionUsecaseBoil) Create(ctx context.Context, sub *domain.Sub) error {
	if _, err := boilmodels.FindSub(ctx, s.db, sub.UserID, sub.CreatorID); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else {
		return errors.New("subscription for this user is already exist")
	}

	boilSub := &boilmodels.Sub{}
	domain.SubDomainToBoil(sub, boilSub)

	if err := boilSub.Insert(ctx, s.db, boil.Infer()); err != nil {
		return err
	}

	domain.SubBoilToDomain(boilSub, sub)

	return nil
}

// Save - saves subscription to history table.
func (s *subscriptionUsecaseBoil) Save(ctx context.Context, sub *domain.Sub) (int64, error) {
	boilSub := &boilmodels.Sub{}

	domain.SubDomainToBoil(sub, boilSub)

	subhist := boilmodels.SubHistory{
		UserID:      boilSub.UserID,
		CreatorID:   boilSub.CreatorID,
		ActivatedAt: boilSub.ActivatedAt,
		ExpiresAt:   boilSub.ExpiresAt,
		Status:      boilSub.Status,
		Price:       boilSub.Price,
	}

	err := subhist.Insert(ctx, s.db, boil.Infer())
	return subhist.SubHistID, err
}

func (s *subscriptionUsecaseBoil) Update(ctx context.Context, sub *domain.Sub) error {
	foundSub, err := boilmodels.FindSub(ctx, s.db, sub.UserID, sub.CreatorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user does not exist")
		}
		return err
	}
	boilSub := &boilmodels.Sub{}

	domain.SubDomainToBoil(sub, boilSub)

	foundSub.ActivatedAt = boilSub.ActivatedAt
	foundSub.ExpiresAt = boilSub.ExpiresAt
	foundSub.Status = boilSub.Status
	foundSub.Price = boilSub.Price

	_, err = foundSub.Update(ctx, s.db, boil.Infer())
	return err
}

func (s *subscriptionUsecaseBoil) Delete(ctx context.Context, userID int64, creatorID int64) error {
	// Delete subscription
	_, err := boilmodels.Subs(qm.Where("user_id=? and creator_id=?", userID, creatorID)).DeleteAll(ctx, s.db)
	return err
}

func (s *subscriptionUsecaseBoil) List(ctx context.Context, cond domain.FindSubRequest) ([]*domain.Sub, error) {
	var conds []qm.QueryMod = make([]qm.QueryMod, 0, 1)
	if cond.Price != nil {
		if cond.Price.Eq != nil {
			conds = append(conds, qm.Where("price=?", *cond.Price.Eq))
		} else if cond.Price.Range != nil {
			if cond.Price.Range.From != nil {
				conds = append(conds, qm.Where("price>?", *cond.Price.Range.From))
			}

			if cond.Price.Range.To != nil {
				conds = append(conds, qm.Where("price<?", *cond.Price.Range.To))
			}
		}
	}

	if cond.Status != nil {
		conds = append(conds, qm.Where("status=?", cond.Status.Eq))
	}

	if cond.TgUserID != nil {
		conds = append(conds, qm.Where("user_id=?", cond.TgUserID.Eq))
	}

	if cond.CreatorID != nil {
		conds = append(conds, qm.Where("creator_id=?", cond.CreatorID.Eq))
	}

	conds = append(conds, qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize)))

	bsubs, err := boilmodels.Subs(conds...).All(ctx, s.db)

	if err != nil {
		return make([]*domain.Sub, 0), err
	}

	domainSubs := make([]*domain.Sub, len(bsubs))

	for i, sub := range bsubs {
		domainSubs[i] = &domain.Sub{}
		domain.SubBoilToDomain(sub, domainSubs[i])
	}

	return domainSubs, nil
}