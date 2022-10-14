package usecase

import (
	"context"
	"database/sql"
	"errors"

	models "github.com/BON4/gosubs/internal/domain/postgres"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type creatorUsecase struct {
	db *sql.DB
}

func NewCretorUsecase(db *sql.DB) models.CreatorUsecase {
	return &creatorUsecase{
		db: db,
	}
}

func (c *creatorUsecase) GetByID(ctx context.Context, id int64) (*models.Creator, error) {
	creator, err := models.FindCreator(ctx, c.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("creator does not exist")
		}
		return nil, err
	}
	return creator, nil
}

func (c *creatorUsecase) GetByTelegramID(ctx context.Context, id int64) (*models.Creator, error) {
	creator := &models.Creator{}
	if err := models.Creators(qm.Where("telegram_di=?", id), qm.Limit(1)).Bind(ctx, c.db, creator); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("creator does not exist")
		}
		return nil, err
	}

	return creator, nil
}

func (c *creatorUsecase) Create(ctx context.Context, creator *models.Creator) error {
	if _, err := models.Creators(qm.Where("telegram_id=?", creator.TelegramID)).One(ctx, c.db); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else {
		//TODO create custom errors
		return errors.New("already exist")
	}

	// Insert
	if err := creator.Insert(ctx, c.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// Delete - will delete creator. Subscriptions will be deleted also.
func (c *creatorUsecase) Delete(ctx context.Context, id int64) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//Delete user subscription
	if _, err := models.Subs(qm.Where("creator_id=?", id)).DeleteAll(ctx, tx); err != nil {
		//TODO Rollback can cause error
		tx.Rollback()
		return err
	}

	// Delete tguser
	if _, err := models.Tgusers(qm.Where("creator_id=?", id)).DeleteAll(ctx, tx); err != nil {
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

func (c *creatorUsecase) Update(ctx context.Context, creator *models.Creator) error {
	foundCreator, err := models.FindCreator(ctx, c.db, creator.CreatorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("creator does not exist")
		}
		return err
	}

	foundCreator.TelegramID = creator.TelegramID
	foundCreator.Username = creator.Username
	foundCreator.Password = creator.Password
	foundCreator.Email = creator.Email
	foundCreator.ChanName = creator.ChanName

	_, err = foundCreator.Update(ctx, c.db, boil.Infer())
	return err
}

func (c *creatorUsecase) List(ctx context.Context, cond models.FindCreatorRequest) ([]*models.Creator, error) {
	return models.Creators(qm.Offset(int(cond.PageSettings.PageNumber)), qm.Limit(int(cond.PageSettings.PageSize))).All(ctx, c.db)
}
