package usecase_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	domain "github.com/BON4/gosubs/internal/domain"
	models "github.com/BON4/gosubs/internal/domain/boil_postgres"
	"github.com/sirupsen/logrus"

	account_usecase "github.com/BON4/gosubs/internal/account/usecase/boil"
	sub_usecase "github.com/BON4/gosubs/internal/subscription/usecase/boil"
	tguser_usecase "github.com/BON4/gosubs/internal/tguser/usecase/boil"
	"github.com/BON4/gosubs/internal/utis/tests"
	_ "github.com/lib/pq"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func TestSubCreate(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	usr := tests.RandomUserBoil()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	crt := tests.RrandomAccountBoil(null.Int64From(usr.UserID))

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubBoil(usr.UserID, crt.AccountID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db, logger.WithContext(ctx))

	if err := subUc.Create(ctx, sub); err != nil {
		t.Fatal(err)
	}

	_, err := models.FindSub(ctx, db, sub.UserID, sub.AccountID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubDelete(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)
	ctx := context.TODO()

	usr := tests.RandomUserBoil()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	crt := tests.RrandomAccountBoil(null.Int64From(usr.UserID))

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubBoil(usr.UserID, crt.AccountID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db, logger.WithContext(ctx))

	if err := subUc.Create(ctx, sub); err != nil {
		t.Fatal(err)
	}

	if err := subUc.Delete(ctx, sub.UserID, sub.AccountID); err != nil {
		t.Fatal(err)
	}

	_, err := models.FindSub(ctx, db, sub.UserID, sub.AccountID)
	if err != sql.ErrNoRows {
		t.Fatal(err)
	}

}

func TestSubUpdate(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)
	ctx := context.TODO()

	usr := tests.RandomUserBoil()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	crt := tests.RrandomAccountBoil(null.Int64From(usr.UserID))

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubBoil(usr.UserID, crt.AccountID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db, logger.WithContext(ctx))

	if err := subUc.Create(ctx, sub); err != nil {
		t.Fatal(err)
	}

	sub.ActivatedAt = time.Now()
	sub.ExpiresAt = time.Now().Add(time.Hour)
	sub.Price = null.IntFrom(0)
	sub.Status = models.SubStatusCancelled

	if err := subUc.Update(ctx, sub); err != nil {
		t.Fatal(err)
	}

	found, err := models.FindSub(ctx, db, sub.UserID, sub.AccountID)
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
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)
	ctx := context.TODO()

	usr := tests.RandomUserBoil()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	crt := tests.RrandomAccountBoil(null.Int64From(usr.UserID))

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubBoil(usr.UserID, crt.AccountID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db, logger.WithContext(ctx))

	if err := subUc.Create(ctx, sub); err != nil {
		t.Fatal(err)
	}

	id, err := subUc.Save(ctx, sub)
	if err != nil {
		t.Fatal(err)
	}

	hist, err := models.FindSubHistory(ctx, db, id)
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

func BenchmarkSubList(b *testing.B) {
	b.StopTimer()
	tests.DeleteAllBoil(db)

	ctx := context.TODO()

	uuc := tguser_usecase.NewBoilTgUserUsecase(db, logger.WithContext(ctx))
	cuc := account_usecase.NewBoilAccountUsecase(db, logger.WithContext(ctx))
	suc := sub_usecase.NewBoilSubscriptionUsecase(db, logger.WithContext(ctx))

	total := b.N

	users := make([]*models.Tguser, total)

	domainAccount := &models.Account{}

	subs := make([]*models.Sub, total)

	b.StartTimer()

	if err := cuc.Create(ctx, domainAccount); err != nil {
		panic(err)
	}

	for i := 0; i < total; i++ {
		users[i] = &models.Tguser{Status: models.UserStatusMember}
		if err := uuc.Create(ctx, users[i]); err != nil {
			b.Fatal(err)
			return
		}

		subs[i] = &models.Sub{
			Status:    models.SubStatusActive,
			AccountID: domainAccount.AccountID,
			UserID:    users[i].UserID,
		}
		if err := suc.Create(ctx, subs[i]); err != nil {
			b.Fatal(err)
			return
		}
	}

	from := int64(0)
	to := int64(1000)

	_, err := suc.List(ctx, domain.FindSubRequest{
		PriceRange: &struct {
			From *int64 `json:"from,omitempty"`
			To   *int64 `json:"to,omitempty"`
		}{
			From: &from,
			To:   &to,
		},
		PageSettings: &struct {
			PageSize   uint "json:\"page_size\""
			PageNumber uint "json:\"page_number\""
		}{
			PageSize:   10,
			PageNumber: 0,
		},
	})
	if err != nil {
		b.Fatal(err)
		return

	}

	b.StopTimer()
	tests.DeleteAllBoil(db)
	b.StartTimer()
}

func BenchmarkUpdate(b *testing.B) {
	b.StopTimer()
	tests.DeleteAllBoil(db)

	ctx := context.TODO()

	uuc := tguser_usecase.NewBoilTgUserUsecase(db, logger.WithContext(ctx))
	cuc := account_usecase.NewBoilAccountUsecase(db, logger.WithContext(ctx))
	suc := sub_usecase.NewBoilSubscriptionUsecase(db, logger.WithContext(ctx))

	total := b.N

	users := make([]*models.Tguser, total)

	domainAccount := tests.RrandomAccountBoil(null.NewInt64(0, false))
	subs := make([]*models.Sub, total)

	b.StartTimer()

	if err := cuc.Create(ctx, domainAccount); err != nil {
		b.Fatal(err)
		return

	}

	for i := 0; i < total; i++ {
		users[i] = tests.RandomUserBoil()
		if err := uuc.Create(ctx, users[i]); err != nil {
			b.Fatal(err)
			return

		}

		subs[i] = tests.RandomSubBoil(users[i].UserID, domainAccount.AccountID)
		if err := suc.Create(ctx, subs[i]); err != nil {
			b.Fatal(err)
			return

		}
	}

	for _, sub := range subs {
		temp := tests.RandomSubBoil(sub.UserID, sub.AccountID)
		sub.ActivatedAt = temp.ActivatedAt
		sub.ExpiresAt = temp.ExpiresAt
		sub.Status = models.SubStatus(temp.Status)
		sub.Price = temp.Price
		if err := suc.Update(ctx, sub); err != nil {
			b.Fatal(err)
		}
	}

	b.StopTimer()
	tests.DeleteAllBoil(db)
	b.StartTimer()

}
