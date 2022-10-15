package usecase

import (
	"context"
	"database/sql"
	"errors"

	domain "github.com/BON4/gosubs/internal/domain"

	sqlcmodels "github.com/BON4/gosubs/internal/domain/sqlc_postgres"
)

type subscriptionUsecaseSqlc struct {
	db sqlcmodels.Store
}

func NewSqlcSubscriptionUsecase(db *sql.DB) domain.SubscriptionUsecase {
	return &subscriptionUsecaseSqlc{
		db: sqlcmodels.NewStore(db),
	}
}

// Create - creates subscribtio2n
func (s *subscriptionUsecaseSqlc) Create(ctx context.Context, sub *domain.Sub) error {
	if ok, err := s.db.IsExistSub(ctx, sqlcmodels.IsExistSubParams{
		UserID:    sub.UserID,
		CreatorID: sub.CreatorID,
	}); err != nil {
		return err
	} else if ok {
		return errors.New("sub already exist")
	}

	sqlcSub := domain.SubDomainToSqlc(sub)

	var err error
	// Insert
	sqlcSub, err = s.db.InsertSub(ctx, sqlcmodels.InsertSubParams{
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      sqlcSub.Status,
		Price:       sqlcSub.Price,
	})

	if err != nil {
		return err
	}

	sub = domain.SubSqlcToDomain(sqlcSub)

	return nil
}

// Save - saves subscription to history table.
func (s *subscriptionUsecaseSqlc) Save(ctx context.Context, sub *domain.Sub) (int64, error) {
	sqlcSub := domain.SubDomainToSqlc(sub)

	subhist, err := s.db.SaveSub(ctx, sqlcmodels.SaveSubParams{
		UserID:      sub.UserID,
		CreatorID:   sub.CreatorID,
		ActivatedAt: sub.ActivatedAt,
		ExpiresAt:   sub.ExpiresAt,
		Status:      sqlcSub.Status,
		Price:       sqlcSub.Price,
	})

	return subhist.SubHistID, err
}

func (s *subscriptionUsecaseSqlc) Update(ctx context.Context, sub *domain.Sub) error {
	_, err := s.db.FindSubID(ctx, sqlcmodels.FindSubIDParams{
		UserID:    sub.UserID,
		CreatorID: sub.CreatorID,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("sub does not exist")
		}
		return err
	}

	sqlcSub := domain.SubDomainToSqlc(sub)

	sqlcSub, err = s.db.UpdateSub(ctx, sqlcmodels.UpdateSubParams{
		UserID:    sqlcSub.UserID,
		CreatorID: sqlcSub.CreatorID,
		ActivatedAt: sql.NullTime{
			Time:  sqlcSub.ActivatedAt,
			Valid: true,
		},
		ExpiresAt: sql.NullTime{
			Time:  sqlcSub.ExpiresAt,
			Valid: true,
		},
		Status: sqlcmodels.NullSubStatus{
			SubStatus: sqlcSub.Status,
			Valid:     true,
		},
		Price: sqlcSub.Price,
	})

	sub = domain.SubSqlcToDomain(sqlcSub)
	return err
}

func (s *subscriptionUsecaseSqlc) Delete(ctx context.Context, userID int64, creatorID int64) error {
	// Delete subscription
	return s.db.DeleteSub(ctx, sqlcmodels.DeleteSubParams{
		UserID:    userID,
		CreatorID: creatorID,
	})
}

func (s *subscriptionUsecaseSqlc) List(ctx context.Context, cond domain.FindSubRequest) ([]*domain.Sub, error) {
	lstParams := sqlcmodels.ListSubParams{}

	if cond.CreatorID != nil {
		lstParams.IsCreatorIDEq = true
		lstParams.CreatorIDEq = cond.CreatorID.Eq
	}

	if cond.TgUserID != nil {
		lstParams.IsUserIDEq = true
		lstParams.UserIDEq = cond.TgUserID.Eq
	}

	if cond.Status != nil {
		lstParams.IsStatusEq = true
		lstParams.StatusEq = sqlcmodels.SubStatus(cond.Status.Eq)
	}

	if cond.Price.Eq != nil {
		lstParams.IsPriceEq = true
		lstParams.PriceEq = sql.NullInt32{
			Int32: int32(*cond.Price.Eq),
			Valid: true,
		}
	} else if cond.Price.Range != nil {
		if cond.Price.Range.From != nil {
			lstParams.IsPriceFrom = true
			lstParams.PriceFrom = sql.NullInt32{
				Int32: int32(*cond.Price.Range.From),
				Valid: true,
			}
		}

		if cond.Price.Range.To != nil {
			lstParams.IsPriceTo = true
			lstParams.PriceTo = sql.NullInt32{
				Int32: int32(*cond.Price.Range.To),
				Valid: true,
			}
		}

	}

	sqlcSubs, err := s.db.ListSub(ctx, lstParams)
	if err != nil {
		return make([]*domain.Sub, 0), err
	}

	domainSubs := make([]*domain.Sub, len(sqlcSubs))

	for i, sub := range sqlcSubs {
		domainSubs[i] = domain.SubSqlcToDomain(sub)
	}

	return domainSubs, nil
}
