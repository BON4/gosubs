package usecase_test

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"testing"

	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	"github.com/sirupsen/logrus"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	account_usecase "github.com/BON4/gosubs/internal/account/usecase/boil"
	"github.com/BON4/gosubs/internal/utis/tests"
	_ "github.com/lib/pq"
)

var db *sql.DB
var logger = logrus.New()

func TestMain(m *testing.M) {
	var err error
	db, err = tests.ConnectTestDB()

	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestAccountCreate(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	crt := tests.RrandomAccountBoil(null.NewInt64(0, false))
	creatorUC := account_usecase.NewBoilAccountUsecase(db, logger.WithContext(ctx))

	if err := creatorUC.Create(ctx, crt); err != nil {
		t.Fatal(err)
	}

	_, err := creatorUC.GetByID(ctx, crt.AccountID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccountcDelete(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	crt := tests.RrandomAccountBoil(null.NewInt64(0, false))
	creatorUC := account_usecase.NewBoilAccountUsecase(db, logger.WithContext(ctx))

	if err := creatorUC.Create(ctx, crt); err != nil {
		t.Fatal(err)
	}

	if err := creatorUC.Delete(ctx, crt.AccountID); err != nil {
		t.Fatal(err)
	}

	_, err := creatorUC.GetByID(ctx, crt.AccountID)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatal(err)
	}
}

func TestAccountUpdate(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	crt := tests.RrandomAccountBoil(null.Int64From(usr.UserID))
	creatorUC := account_usecase.NewBoilAccountUsecase(db, logger.WithContext(ctx))

	if err := creatorUC.Create(ctx, crt); err != nil {
		t.Fatal(err)
	}

	crt.ChanName = null.StringFrom("updated")
	crt.Email = "updated"
	crt.Password = []byte("UPDATED")
	crt.Role = models.AccountRoleCreator
	crt.UserID = null.Int64From(usr.UserID)

	if err := creatorUC.Update(ctx, crt); err != nil {
		t.Fatal(err)
	}

	found, err := creatorUC.GetByID(ctx, crt.AccountID)
	if err != nil {
		t.Fatal(err)
	}

	if found.AccountID != crt.AccountID ||
		found.ChanName != crt.ChanName ||
		found.Role != crt.Role ||
		found.UserID != crt.UserID ||
		bytes.Compare(found.Password, crt.Password) != 0 ||
		found.Email != crt.Email {

		t.Logf("Found: %+v\n", found)
		t.Logf("Expected: %+v\n", crt)
		t.Fatal("entities dont match")
	}
}
