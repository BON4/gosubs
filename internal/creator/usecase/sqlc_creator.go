package usecase

import (
	"context"
	"database/sql"
	"errors"

	sqlcmodels "github.com/BON4/gosubs/internal/domain/sqlc_postgres"

	domain "github.com/BON4/gosubs/internal/domain"
)

type creatorUsecaseSqlc struct {
	db sqlcmodels.Store
}

func NewSqlcCretorUsecase(db *sql.DB) domain.CreatorUsecase {
	return &creatorUsecaseSqlc{
		db: sqlcmodels.NewStore(db),
	}
}

func (c *creatorUsecaseSqlc) GetByID(ctx context.Context, id int64) (*domain.Creator, error) {
	creator, err := c.db.FindCreatorID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("creator does not exist")
		}
		return nil, err
	}
	return domain.CreatorSqlcToDomain(creator), nil
}

func (c *creatorUsecaseSqlc) GetByTelegramID(ctx context.Context, id int64) (*domain.Creator, error) {
	creator, err := c.db.FindCreatorTelegramID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("creator does not exist")
		}
		return nil, err
	}
	return domain.CreatorSqlcToDomain(creator), nil
}

func (c *creatorUsecaseSqlc) Create(ctx context.Context, creator *domain.Creator) error {
	if ok, err := c.db.IsExistTguser(ctx, creator.TelegramID); err != nil {
		return err
	} else if ok {
		return errors.New("already exist")
	}

	sqlcCreator := domain.CreatorDomainToSqlc(creator)

	var err error
	// Insert
	sqlcCreator, err = c.db.InsertCreator(ctx, sqlcmodels.InsertCreatorParams{
		TelegramID: sqlcCreator.TelegramID,
		Username:   sqlcCreator.Username,
		Email:      sqlcCreator.Email,
		Password:   sqlcCreator.Password,
		ChanName:   sqlcCreator.ChanName,
	})

	if err != nil {
		return err
	}

	creator = domain.CreatorSqlcToDomain(sqlcCreator)

	return nil
}

// Delete - will delete creator. Subscriptions will be deleted also.
func (c *creatorUsecaseSqlc) Delete(ctx context.Context, id int64) error {
	return c.db.DeleteCreatorTx(ctx, sqlcmodels.DeleteCreatorTxParams{
		CreatorID: id,
	})
}

func (c *creatorUsecaseSqlc) Update(ctx context.Context, creator *domain.Creator) error {
	_, err := c.db.FindCreatorID(ctx, creator.CreatorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user does not exist")
		}
		return err
	}

	sqlcCreator := domain.CreatorDomainToSqlc(creator)

	sqlcCreator, err = c.db.UpdateCreator(ctx, sqlcmodels.UpdateCreatorParams{
		TelegramID: sql.NullInt64{
			Int64: sqlcCreator.TelegramID,
			Valid: true,
		},
		Username: sql.NullString{
			String: sqlcCreator.Username,
			Valid:  true,
		},
		Password: sqlcCreator.Password,
		Email:    sqlcCreator.Email,
		ChanName: sqlcCreator.ChanName,

		CreatorID: sqlcCreator.CreatorID,
	})

	creator = domain.CreatorSqlcToDomain(sqlcCreator)
	return err
}

func (c *creatorUsecaseSqlc) List(ctx context.Context, cond domain.FindCreatorRequest) ([]*domain.Creator, error) {
	susers, err := c.db.ListCreator(ctx, sqlcmodels.ListCreatorParams{
		Offset: int32(cond.PageSettings.PageNumber),
		Limit:  int32(cond.PageSettings.PageSize),
	})

	if err != nil {
		return make([]*domain.Creator, 0), err
	}

	domainCretors := make([]*domain.Creator, len(susers))

	for i, user := range susers {
		domainCretors[i] = domain.CreatorSqlcToDomain(user)
	}

	return domainCretors, nil
}
