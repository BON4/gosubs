package usecase_test

import (
	"context"
	"database/sql"
	"testing"

	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	"github.com/sirupsen/logrus"

	//myerrors "github.com/BON4/gosubs/internal/errors"

	user_usecase "github.com/BON4/gosubs/internal/tguser/usecase/boil"
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

func TestTguserCreate(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	userUC := user_usecase.NewBoilTgUserUsecase(db, logger.WithContext(ctx))

	if err := userUC.Create(ctx, usr); err != nil {
		t.Fatal(err)
	}

	_, err := models.FindTguser(ctx, db, usr.UserID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTguserDelete(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	userUC := user_usecase.NewBoilTgUserUsecase(db, logger.WithContext(ctx))

	if err := userUC.Create(ctx, usr); err != nil {
		t.Fatal(err)
	}

	if err := userUC.Delete(ctx, usr.UserID); err != nil {
		t.Fatal(err)
	}

	_, err := models.FindTguser(ctx, db, usr.UserID)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func TestTguserUpdate(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	userUC := user_usecase.NewBoilTgUserUsecase(db, logger.WithContext(ctx))

	if err := userUC.Create(ctx, usr); err != nil {
		t.Fatal(err)
	}

	usr.Username = "updated"
	usr.Status = models.UserStatusKicked

	if err := userUC.Update(ctx, usr); err != nil {
		t.Fatal(err)
	}

	found, err := models.FindTguser(ctx, db, usr.UserID)
	if err != nil {
		t.Fatal(err)
	}

	if found.UserID != usr.UserID ||
		found.TelegramID != usr.TelegramID ||
		found.Username != usr.Username ||
		found.Status != models.UserStatus(usr.Status) {
		t.Logf("Found: %+v\n", found)
		t.Logf("Expected: %+v\n", usr)
		t.Fatal("entities dont match")
	}
}
