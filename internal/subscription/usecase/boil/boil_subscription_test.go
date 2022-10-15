package usecase_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/BON4/gosubs/internal/domain"
	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"

	creator_usecase "github.com/BON4/gosubs/internal/creator/usecase"
	sub_usecase "github.com/BON4/gosubs/internal/subscription/usecase/boil"
	tguser_usecase "github.com/BON4/gosubs/internal/tguser/usecase/boil"
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
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)

	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	crt := tests.RrandomCreatorBoil()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubBoil(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db)

	domainSub := &domain.Sub{}
	domain.SubBoilToDomain(sub, domainSub)

	if err := subUc.Create(ctx, domainSub); err != nil {
		t.Fatal(err)
	}

	_, err := boilmodels.FindSub(ctx, db, sub.UserID, sub.CreatorID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubDelete(t *testing.T) {
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)
	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	crt := tests.RrandomCreatorBoil()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubBoil(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db)

	domainSub := &domain.Sub{}
	domain.SubBoilToDomain(sub, domainSub)

	if err := subUc.Create(ctx, domainSub); err != nil {
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
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)
	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	crt := tests.RrandomCreatorBoil()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubBoil(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db)

	domainSub := &domain.Sub{}
	domain.SubBoilToDomain(sub, domainSub)

	if err := subUc.Create(ctx, domainSub); err != nil {
		t.Fatal(err)
	}

	sub.ActivatedAt = time.Now()
	sub.ExpiresAt = time.Now().Add(time.Hour)
	sub.Price = null.IntFrom(0)
	sub.Status = boilmodels.SubStatusCancelled

	domain.SubBoilToDomain(sub, domainSub)

	if err := subUc.Update(ctx, domainSub); err != nil {
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
	tests.DeleteAllBoil(db)
	defer tests.DeleteAllBoil(db)
	ctx := context.TODO()

	usr := tests.RandomUserBoil()
	crt := tests.RrandomCreatorBoil()

	if err := usr.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err := crt.Insert(ctx, db, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubBoil(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewBoilSubscriptionUsecase(db)

	domainSub := &domain.Sub{}
	domain.SubBoilToDomain(sub, domainSub)

	if err := subUc.Create(ctx, domainSub); err != nil {
		t.Fatal(err)
	}

	domain.SubBoilToDomain(sub, domainSub)

	id, err := subUc.Save(ctx, domainSub)
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

func BenchmarkSubList(b *testing.B) {
	b.StopTimer()
	tests.DeleteAllBoil(db)

	ctx := context.TODO()

	uuc := tguser_usecase.NewBoilTgUserUsecase(db)
	cuc := creator_usecase.NewBoilCretorUsecase(db)
	suc := sub_usecase.NewBoilSubscriptionUsecase(db)

	total := b.N

	users := make([]*domain.Tguser, total)

	domainCreator := &domain.Creator{}
	domain.CreatorBoilToDomain(tests.RrandomCreatorBoil(), domainCreator)

	subs := make([]*domain.Sub, total)

	b.StartTimer()

	if err := cuc.Create(ctx, domainCreator); err != nil {
		panic(err)
	}

	for i := 0; i < total; i++ {
		users[i] = &domain.Tguser{}
		domain.TguserBoilToDomain(tests.RandomUserBoil(), users[i])

		if err := uuc.Create(ctx, users[i]); err != nil {
			panic(err)
		}
		subs[i] = &domain.Sub{}
		domain.SubBoilToDomain(tests.RandomSubBoil(users[i].UserID, domainCreator.CreatorID), subs[i])

		if err := suc.Create(ctx, subs[i]); err != nil {
			panic(err)
		}
	}

	from := 0
	to := 1000

	_, err := suc.List(ctx, domain.FindSubRequest{
		Price: &struct {
			Eq    *int `json:"eq,omitempty"`
			Range *struct {
				From *int `json:"from,omitempty"`
				To   *int `json:"to,omitempty"`
			} `json:"range,omitempty"`
		}{
			Range: &struct {
				From *int `json:"from,omitempty"`
				To   *int `json:"to,omitempty"`
			}{
				From: &from,
				To:   &to,
			},
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
		panic(err)
	}

	b.StopTimer()
	tests.DeleteAllBoil(db)
	b.StartTimer()
}

func BenchmarkCreate(b *testing.B) {
	b.StopTimer()
	tests.DeleteAllBoil(db)

	ctx := context.TODO()

	uuc := tguser_usecase.NewBoilTgUserUsecase(db)
	cuc := creator_usecase.NewBoilCretorUsecase(db)
	suc := sub_usecase.NewBoilSubscriptionUsecase(db)

	total := b.N

	users := make([]*domain.Tguser, total)

	domainCreator := &domain.Creator{}
	domain.CreatorBoilToDomain(tests.RrandomCreatorBoil(), domainCreator)

	subs := make([]*domain.Sub, total)

	b.StartTimer()

	if err := cuc.Create(ctx, domainCreator); err != nil {
		panic(err)
	}

	for i := 0; i < total; i++ {
		users[i] = &domain.Tguser{}
		domain.TguserBoilToDomain(tests.RandomUserBoil(), users[i])

		if err := uuc.Create(ctx, users[i]); err != nil {
			panic(err)
		}
		subs[i] = &domain.Sub{}
		domain.SubBoilToDomain(tests.RandomSubBoil(users[i].UserID, domainCreator.CreatorID), subs[i])

		if err := suc.Create(ctx, subs[i]); err != nil {
			panic(err)
		}
	}

	for _, sub := range subs {
		temp := tests.RandomSubBoil(sub.UserID, sub.CreatorID)
		sub.ActivatedAt = temp.ActivatedAt
		sub.ExpiresAt = temp.ExpiresAt
		sub.Status = domain.SubStatus(temp.Status)
		sub.Price = temp.Price
		if err := suc.Update(ctx, sub); err != nil {
			b.Fatal(err)
		}
	}

	b.StopTimer()
	tests.DeleteAllBoil(db)
	b.StartTimer()

}
