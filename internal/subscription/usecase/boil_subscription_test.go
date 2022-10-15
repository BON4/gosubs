package usecase_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/BON4/gosubs/internal/domain"
	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"

	sub_usecase "github.com/BON4/gosubs/internal/subscription/usecase"
	"github.com/BON4/gosubs/internal/utis/tests"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func TestSubCreate(t *testing.T) {
	tests.DeleteAll(db)
	defer tests.DeleteAll(db)

	ctx := context.TODO()

	usr := tests.RandomUser()
	crt := tests.RrandomCreator()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSub(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db)

	if err := subUc.Create(ctx, domain.SubBoilToDomain(sub)); err != nil {
		t.Fatal(err)
	}

	_, err := boilmodels.FindSub(ctx, db, sub.UserID, sub.CreatorID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubDelete(t *testing.T) {
	tests.DeleteAll(db)
	defer tests.DeleteAll(db)
	ctx := context.TODO()

	usr := tests.RandomUser()
	crt := tests.RrandomCreator()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSub(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db)

	if err := subUc.Create(ctx, domain.SubBoilToDomain(sub)); err != nil {
		t.Fatal(err)
	}

	if err := subUc.Delete(ctx, sub.UserID, sub.CreatorID); err != nil {
		t.Fatal(err)
	}

	_, err := boilmodels.FindSub(ctx, db, sub.UserID, sub.CreatorID)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}

}

func TestSubUpdate(t *testing.T) {
	tests.DeleteAll(db)
	defer tests.DeleteAll(db)
	ctx := context.TODO()

	usr := tests.RandomUser()
	crt := tests.RrandomCreator()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSub(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db)

	if err := subUc.Create(ctx, domain.SubBoilToDomain(sub)); err != nil {
		t.Fatal(err)
	}

	sub.ActivatedAt = time.Now()
	sub.ExpiresAt = time.Now().Add(time.Hour)
	sub.Price = null.IntFrom(0)
	sub.Status = boilmodels.SubStatusCancelled

	if err := subUc.Update(ctx, domain.SubBoilToDomain(sub)); err != nil {
		t.Fatal(err)
	}

	found, err := boilmodels.FindSub(ctx, db, sub.UserID, sub.CreatorID)
	if err != nil {
		t.Fatal(err)
	}

	if found.ActivatedAt.Unix() != sub.ActivatedAt.Unix() ||
		found.ExpiresAt.Unix() != sub.ExpiresAt.Unix() ||
		found.Price != sub.Price ||
		found.Status != sub.Status {
		t.Logf("Found: %+v\n", found)
		t.Logf("Expected: %+v\n", sub)
		t.Fatal("entities dont match")
	}
}

func TestSubSave(t *testing.T) {
	tests.DeleteAll(db)
	defer tests.DeleteAll(db)
	ctx := context.TODO()

	usr := tests.RandomUser()
	crt := tests.RrandomCreator()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSub(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db)

	if err := subUc.Create(ctx, domain.SubBoilToDomain(sub)); err != nil {
		t.Fatal(err)
	}

	id, err := subUc.Save(ctx, domain.SubBoilToDomain(sub))
	if err != nil {
		t.Fatal(err)
	}

	hist, err := boilmodels.FindSubHistory(ctx, db, id)
	if err != nil {
		t.Fatal(err)
	}

	if hist.ActivatedAt.Unix() != sub.ActivatedAt.Unix() ||
		hist.ExpiresAt.Unix() != sub.ExpiresAt.Unix() ||
		hist.Price != sub.Price ||
		hist.Status != sub.Status {
		t.Logf("Found: %+v\n", hist)
		t.Logf("Expected: %+v\n", sub)
		t.Fatal("entities dont match")
	}
}
