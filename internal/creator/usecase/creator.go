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

type creatorUsecase struct {
	db *sql.DB
}

func NewCretorUsecase(db *sql.DB) domain.CreatorUsecase {
	return &creatorUsecase{
		db: db,
	}
}

func (c *creatorUsecase) GetByID(ctx context.Context, id int64) (*domain.Creator, error) {
	creator, err := boilmodels.FindCreator(ctx, c.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("creator does not exist")
		}
		return nil, err
	}
	return domain.CreatorBoilToDomain(creator), nil
}

func (c *creatorUsecase) GetByTelegramID(ctx context.Context, id int64) (*domain.Creator, error) {
	creator := &boilmodels.Creator{}
	if err := boilmodels.Creators(qm.Where("telegram_di=?", id), qm.Limit(1)).Bind(ctx, c.db, creator); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("creator does not exist")
		}
		return nil, err
	}

	return domain.CreatorBoilToDomain(creator), nil
}

func (c *creatorUsecase) Create(ctx context.Context, creator *domain.Creator) error {
	if _, err := boilmodels.Creators(qm.Where("telegram_id=?", creator.TelegramID)).One(ctx, c.db); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else {
		//TODO create custom errors
		return errors.New("already exist")
	}

	boilCreator := domain.CreatorDomainToBoil(creator)

	// Insert
	if err := boilCreator.Insert(ctx, c.db, boil.Infer()); err != nil {
		return err
	}

	creator = domain.CreatorBoilToDomain(boilCreator)

	return nil
}

// Delete - will delete creator. Subscriptions will be deleted also.
func (c *creatorUsecase) Delete(ctx context.Context, id int64) error {
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

func (c *creatorUsecase) Update(ctx context.Context, creator *domain.Creator) error {
	foundCreator, err := boilmodels.FindCreator(ctx, c.db, creator.CreatorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("creator does not exist")
		}
		return err
	}

	boilCreator := domain.CreatorDomainToBoil(creator)

	foundCreator.TelegramID = boilCreator.TelegramID
	foundCreator.Username = boilCreator.Username
	foundCreator.Password = boilCreator.Password
	foundCreator.Email = boilCreator.Email
	foundCreator.ChanName = boilCreator.ChanName

	_, err = foundCreator.Update(ctx, c.db, boil.Infer())

	creator = domain.CreatorBoilToDomain(foundCreator)

	return err
}

func (c *creatorUsecase) List(ctx context.Context, cond domain.FindCreatorRequest) ([]*domain.Creator, error) {
	bcreators, err := boilmodels.Creators(qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize))).All(ctx, c.db)
	if err != nil {
		return make([]*domain.Creator, 0), err
	}

	domainCreators := make([]*domain.Creator, len(bcreators))

	for i, creator := range bcreators {
		domainCreators[i] = domain.CreatorBoilToDomain(creator)
	}

	return domainCreators, nil
}
