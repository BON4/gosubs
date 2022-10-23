package usecase_test

import (
	"bytes"
	"context"
	"database/sql"
	"testing"

	"github.com/BON4/gosubs/internal/domain"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	account_usecase "github.com/BON4/gosubs/internal/account/usecase/boil"
	"github.com/BON4/gosubs/internal/utis/tests"
	_ "github.com/lib/pq"
)

var db *sql.DB

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
	creatorUC := account_usecase.NewBoilAccountUsecase(db)

	domainAccount := &domain.Account{}
	domain.AccountBoilToDomain(crt, domainAccount)

	if err := creatorUC.Create(ctx, domainAccount); err != nil {
		t.Fatal(err)
	}

	_, err := creatorUC.GetByID(ctx, domainAccount.AccountID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccountcDelete(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	crt := tests.RrandomAccountBoil(null.NewInt64(0, false))
	creatorUC := account_usecase.NewBoilAccountUsecase(db)

	domainAccount := &domain.Account{}
	domain.AccountBoilToDomain(crt, domainAccount)

	if err := creatorUC.Create(ctx, domainAccount); err != nil {
		t.Fatal(err)
	}

	if err := creatorUC.Delete(ctx, domainAccount.AccountID); err != nil {
		t.Fatal(err)
	}

	_, err := creatorUC.GetByID(ctx, domainAccount.AccountID)
	if err != sql.ErrNoRows {
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
	creatorUC := account_usecase.NewBoilAccountUsecase(db)

	domainAccount := &domain.Account{}
	domain.AccountBoilToDomain(crt, domainAccount)

	if err := creatorUC.Create(ctx, domainAccount); err != nil {
		t.Fatal(err)
	}

	domainAccount.ChanName = null.StringFrom("updated")
	domainAccount.Email = "updated"
	domainAccount.Password = []byte("UPDATED")
	domainAccount.Role = domain.AccountRoleCreator
	domainAccount.UserID = null.Int64From(usr.UserID)

	if err := creatorUC.Update(ctx, domainAccount); err != nil {
		t.Fatal(err)
	}

	found, err := creatorUC.GetByID(ctx, domainAccount.AccountID)
	if err != nil {
		t.Fatal(err)
	}

	if found.AccountID != domainAccount.AccountID ||
		found.ChanName != domainAccount.ChanName ||
		found.Role != domainAccount.Role ||
		found.UserID != domainAccount.UserID ||
		bytes.Compare(found.Password, domainAccount.Password) != 0 ||
		found.Email != domainAccount.Email {

		t.Logf("Found: %+v\n", found)
		t.Logf("Expected: %+v\n", domainAccount)
		t.Fatal("entities dont match")
	}
}
