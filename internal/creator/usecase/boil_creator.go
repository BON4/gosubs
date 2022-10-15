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

type creatorUsecaseBoil struct {
	db *sql.DB
}

func NewBoilCretorUsecase(db *sql.DB) domain.CreatorUsecase {
	return &creatorUsecaseBoil{
		db: db,
	}
}

func (c *creatorUsecaseBoil) GetByID(ctx context.Context, id int64) (*domain.Creator, error) {
	creator, err := boilmodels.FindCreator(ctx, c.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("creator does not exist")
		}
		return nil, err
	}
	domainCreator := &domain.Creator{}
	domain.CreatorBoilToDomain(creator, domainCreator)
	return domainCreator, nil
}

func (c *creatorUsecaseBoil) GetByTelegramID(ctx context.Context, id int64) (*domain.Creator, error) {
	creator := &boilmodels.Creator{}
	if err := boilmodels.Creators(qm.Where("telegram_di=?", id), qm.Limit(1)).Bind(ctx, c.db, creator); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("creator does not exist")
		}
		return nil, err
	}
	domainCreator := &domain.Creator{}
	domain.CreatorBoilToDomain(creator, domainCreator)
	return domainCreator, nil
}

func (c *creatorUsecaseBoil) Create(ctx context.Context, creator *domain.Creator) error {
	if _, err := boilmodels.Creators(qm.Where("telegram_id=?", creator.TelegramID)).One(ctx, c.db); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else {
		//TODO create custom errors
		return errors.New("already exist")
	}

	boilCreator := &boilmodels.Creator{}
	domain.CreatorDomainToBoil(creator, boilCreator)

	// Insert
	if err := boilCreator.Insert(ctx, c.db, boil.Infer()); err != nil {
		return err
	}

	domain.CreatorBoilToDomain(boilCreator, creator)

	return nil
}

// Delete - will delete creator. Subscriptions will be deleted also.
func (c *creatorUsecaseBoil) Delete(ctx context.Context, id int64) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//Delete user subscription
	if _, err := boilmodels.Subs(qm.Where("creator_id=?", id)).DeleteAll(ctx, tx); err != nil {
		//TODO Rollback can cause error
		tx.Rollback()
		return err
	}

	// Delete tguser
	if _, err := boilmodels.Tgusers(qm.Where("creator_id=?", id)).DeleteAll(ctx, tx); err != nil {
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

func (c *creatorUsecaseBoil) Update(ctx context.Context, creator *domain.Creator) error {
	foundCreator, err := boilmodels.FindCreator(ctx, c.db, creator.CreatorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("creator does not exist")
		}
		return err
	}

	boilCreator := &boilmodels.Creator{}
	domain.CreatorDomainToBoil(creator, boilCreator)

	foundCreator.TelegramID = boilCreator.TelegramID
	foundCreator.Username = boilCreator.Username
	foundCreator.Password = boilCreator.Password
	foundCreator.Email = boilCreator.Email
	foundCreator.ChanName = boilCreator.ChanName

	_, err = foundCreator.Update(ctx, c.db, boil.Infer())

	domain.CreatorBoilToDomain(boilCreator, creator)

	return err
}

func (c *creatorUsecaseBoil) List(ctx context.Context, cond domain.FindCreatorRequest) ([]*domain.Creator, error) {
	bcreators, err := boilmodels.Creators(qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize))).All(ctx, c.db)
	if err != nil {
		return make([]*domain.Creator, 0), err
	}

	domainCreators := make([]*domain.Creator, len(bcreators))

	for i, creator := range bcreators {
		domainCreators[i] = &domain.Creator{}
		domain.CreatorBoilToDomain(creator, domainCreators[i])
	}

	return domainCreators, nil
}
