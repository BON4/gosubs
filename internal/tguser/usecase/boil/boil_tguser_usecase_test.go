package usecase_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/BON4/gosubs/internal/domain"
	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"

	user_usecase "github.com/BON4/gosubs/internal/tguser/usecase/boil"
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

func TestTguserCreate(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	userUC := user_usecase.NewBoilTgUserUsecase(db)

	domainUser := &domain.Tguser{}
	domain.TguserBoilToDomain(usr, domainUser)

	if err := userUC.Create(ctx, domainUser); err != nil {
		t.Fatal(err)
	}

	_, err := boilmodels.FindTguser(ctx, db, domainUser.UserID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTguserDelete(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	userUC := user_usecase.NewBoilTgUserUsecase(db)

	domainUser := &domain.Tguser{}
	domain.TguserBoilToDomain(usr, domainUser)

	if err := userUC.Create(ctx, domainUser); err != nil {
		t.Fatal(err)
	}

	if err := userUC.Delete(ctx, domainUser.UserID); err != nil {
		t.Fatal(err)
	}

	_, err := boilmodels.FindTguser(ctx, db, domainUser.UserID)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestSubUpdate(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	userUC := user_usecase.NewBoilTgUserUsecase(db)

	domainUser := &domain.Tguser{}
	domain.TguserBoilToDomain(usr, domainUser)

	if err := userUC.Create(ctx, domainUser); err != nil {
		t.Fatal(err)
	}

	domainUser.Username = "updated"
	domainUser.Status = domain.UserStatus(boilmodels.UserStatusKicked)

	if err := userUC.Update(ctx, domainUser); err != nil {
		t.Fatal(err)
	}

	found, err := boilmodels.FindTguser(ctx, db, domainUser.UserID)
	if err != nil {
		t.Fatal(err)
	}

	if found.UserID != domainUser.UserID ||
		found.TelegramID != domainUser.TelegramID ||
		found.Username != domainUser.Username ||
		found.Status != boilmodels.UserStatus(domainUser.Status) {
		t.Logf("Found: %+v\n", found)
		t.Logf("Expected: %+v\n", domainUser)
		t.Fatal("entities dont match")
	}
}
