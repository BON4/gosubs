package usecase_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	creator_usecase "github.com/BON4/gosubs/internal/creator/usecase"
	"github.com/BON4/gosubs/internal/domain"
	sqlcmodels "github.com/BON4/gosubs/internal/domain/sqlc_postgres"
	sub_usecase "github.com/BON4/gosubs/internal/subscription/usecase/sqlc"
	tguser_usecase "github.com/BON4/gosubs/internal/tguser/usecase/sqlc"
	"github.com/BON4/gosubs/internal/utis/tests"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
)

var db sqlcmodels.Store
var rawDb *sql.DB

func TestMain(m *testing.M) {
	dbraw, err := tests.ConnectTestDB()
	if err != nil {
		panic(err)
	}

	db = sqlcmodels.NewStore(dbraw)
	rawDb = dbraw

	m.Run()
}

func TestSubCreateSqlc(t *testing.T) {
	tests.DeleteAllSqlc(db)
	defer tests.DeleteAllSqlc(db)

	ctx := context.TODO()

	usr := tests.RandomUserSqlc()
	crt := tests.RrandomCreatorSqlc()

	usr, err := db.InsertTguser(ctx, sqlcmodels.InsertTguserParams{
		TelegramID: usr.TelegramID,
		Username:   usr.Username,
		Status:     usr.Status,
	})
	if err != nil {
		t.Fatal(err)
	}

	crt, err = db.InsertCreator(ctx, sqlcmodels.InsertCreatorParams{
		TelegramID: crt.TelegramID,
		Username:   crt.Username,
		Password:   crt.Password,
		Email:      crt.Email,
		ChanName:   crt.ChanName,
	})
	if err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubSqlc(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewSqlcSubscriptionUsecase(rawDb)

	domainSub := &domain.Sub{}
	domain.SubSqlcToDomain(sub, domainSub)

	if err := subUc.Create(ctx, domainSub); err != nil {
		t.Fatal(err)
	}

	_, err = db.FindSubID(ctx, sqlcmodels.FindSubIDParams{
		UserID:    sub.UserID,
		CreatorID: sub.CreatorID,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubDelete(t *testing.T) {
	tests.DeleteAllSqlc(db)
	defer tests.DeleteAllSqlc(db)

	ctx := context.TODO()

	usr := tests.RandomUserSqlc()
	crt := tests.RrandomCreatorSqlc()

	usr, err := db.InsertTguser(ctx, sqlcmodels.InsertTguserParams{
		TelegramID: usr.TelegramID,
		Username:   usr.Username,
		Status:     usr.Status,
	})
	if err != nil {
		t.Fatal(err)
	}

	crt, err = db.InsertCreator(ctx, sqlcmodels.InsertCreatorParams{
		TelegramID: crt.TelegramID,
		Username:   crt.Username,
		Password:   crt.Password,
		Email:      crt.Email,
		ChanName:   crt.ChanName,
	})
	if err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubSqlc(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewSqlcSubscriptionUsecase(rawDb)

	domainSub := &domain.Sub{}
	domain.SubSqlcToDomain(sub, domainSub)

	if err := subUc.Create(ctx, domainSub); err != nil {
		t.Fatal(err)
	}

	if err := subUc.Delete(ctx, sub.UserID, sub.CreatorID); err != nil {
		t.Fatal(err)
	}

	_, err = db.FindSubID(ctx, sqlcmodels.FindSubIDParams{
		UserID:    sub.UserID,
		CreatorID: sub.CreatorID,
	})

	if err != sql.ErrNoRows {
		t.Fatal(err)
	}

}

func TestSubUpdate(t *testing.T) {
	tests.DeleteAllSqlc(db)
	defer tests.DeleteAllSqlc(db)

	ctx := context.TODO()

	usr := tests.RandomUserSqlc()
	crt := tests.RrandomCreatorSqlc()

	usr, err := db.InsertTguser(ctx, sqlcmodels.InsertTguserParams{
		TelegramID: usr.TelegramID,
		Username:   usr.Username,
		Status:     usr.Status,
	})
	if err != nil {
		t.Fatal(err)
	}

	crt, err = db.InsertCreator(ctx, sqlcmodels.InsertCreatorParams{
		TelegramID: crt.TelegramID,
		Username:   crt.Username,
		Password:   crt.Password,
		Email:      crt.Email,
		ChanName:   crt.ChanName,
	})
	if err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubSqlc(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewSqlcSubscriptionUsecase(rawDb)

	domainSub := &domain.Sub{}
	domain.SubSqlcToDomain(sub, domainSub)

	if err := subUc.Create(ctx, domainSub); err != nil {
		t.Fatal(err)
	}

	sub, err = db.FindSubID(ctx, sqlcmodels.FindSubIDParams{
		UserID:    sub.UserID,
		CreatorID: sub.CreatorID,
	})

	sub.ActivatedAt = time.Now()
	sub.ExpiresAt = time.Now().Add(time.Hour)
	sub.Price = sql.NullInt32{Int32: 0, Valid: true}
	sub.Status = sqlcmodels.SubStatusCancelled

	domain.SubSqlcToDomain(sub, domainSub)

	if err := subUc.Update(ctx, domainSub); err != nil {
		t.Fatal(err)
	}

	found, err := db.FindSubID(ctx, sqlcmodels.FindSubIDParams{
		UserID:    sub.UserID,
		CreatorID: sub.CreatorID,
	})
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
	tests.DeleteAllSqlc(db)
	defer tests.DeleteAllSqlc(db)

	ctx := context.TODO()

	usr := tests.RandomUserSqlc()
	crt := tests.RrandomCreatorSqlc()

	usr, err := db.InsertTguser(ctx, sqlcmodels.InsertTguserParams{
		TelegramID: usr.TelegramID,
		Username:   usr.Username,
		Status:     usr.Status,
	})
	if err != nil {
		t.Fatal(err)
	}

	crt, err = db.InsertCreator(ctx, sqlcmodels.InsertCreatorParams{
		TelegramID: crt.TelegramID,
		Username:   crt.Username,
		Password:   crt.Password,
		Email:      crt.Email,
		ChanName:   crt.ChanName,
	})
	if err != nil {
		t.Fatal(err)
	}

	sub := tests.RandomSubSqlc(usr.UserID, crt.CreatorID)

	subUc := sub_usecase.NewSqlcSubscriptionUsecase(rawDb)

	domainSub := &domain.Sub{}
	domain.SubSqlcToDomain(sub, domainSub)

	if err := subUc.Create(ctx, domainSub); err != nil {
		t.Fatal(err)
	}

	domain.SubSqlcToDomain(sub, domainSub)

	id, err := subUc.Save(ctx, domainSub)
	if err != nil {
		t.Fatal(err)
	}

	hist, err := db.FindSubHistory(ctx, id)
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
	tests.DeleteAllSqlc(db)

	total := b.N

	ctx := context.TODO()

	uuc := tguser_usecase.NewSqlcTguserUsecase(rawDb)
	cuc := creator_usecase.NewSqlcCretorUsecase(rawDb)
	suc := sub_usecase.NewSqlcSubscriptionUsecase(rawDb)

	users := make([]*domain.Tguser, total)

	domainCreator := &domain.Creator{}
	domain.CreatorSqlcToDomain(tests.RrandomCreatorSqlc(), domainCreator)

	subs := make([]*domain.Sub, total)

	b.StartTimer()

	if err := cuc.Create(ctx, domainCreator); err != nil {
		panic(err)
	}

	for i := 0; i < total; i++ {
		users[i] = &domain.Tguser{}
		domain.TguserSqlcToDomain(tests.RandomUserSqlc(), users[i])

		if err := uuc.Create(ctx, users[i]); err != nil {
			panic(err)
		}

		subs[i] = &domain.Sub{}
		domain.SubSqlcToDomain(tests.RandomSubSqlc(users[i].UserID, domainCreator.CreatorID), subs[i])

		if err := suc.Create(ctx, subs[i]); err != nil {
			panic(err)
		}
	}

	from := 0
	to := 1000
	found, err := suc.List(ctx, domain.FindSubRequest{
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

	if len(found) == 0 {
		b.Fatal("list select fail")
	}

	b.StopTimer()
	tests.DeleteAllSqlc(db)
	b.StartTimer()
}

func BenchmarkUpdate(b *testing.B) {
	b.StopTimer()
	tests.DeleteAllSqlc(db)

	total := b.N

	ctx := context.TODO()

	uuc := tguser_usecase.NewSqlcTguserUsecase(rawDb)
	cuc := creator_usecase.NewSqlcCretorUsecase(rawDb)
	suc := sub_usecase.NewSqlcSubscriptionUsecase(rawDb)

	users := make([]*domain.Tguser, total)

	domainCreator := &domain.Creator{}
	domain.CreatorSqlcToDomain(tests.RrandomCreatorSqlc(), domainCreator)

	subs := make([]*domain.Sub, total)

	b.StartTimer()

	if err := cuc.Create(ctx, domainCreator); err != nil {
		panic(err)
	}

	for i := 0; i < total; i++ {
		users[i] = &domain.Tguser{}
		domain.TguserSqlcToDomain(tests.RandomUserSqlc(), users[i])

		if err := uuc.Create(ctx, users[i]); err != nil {
			panic(err)
		}

		subs[i] = &domain.Sub{}
		domain.SubSqlcToDomain(tests.RandomSubSqlc(users[i].UserID, domainCreator.CreatorID), subs[i])

		if err := suc.Create(ctx, subs[i]); err != nil {
			panic(err)
		}
	}

	for _, sub := range subs {
		temp := tests.RandomSubSqlc(sub.UserID, sub.CreatorID)
		sub.ActivatedAt = temp.ActivatedAt
		sub.ExpiresAt = temp.ExpiresAt
		sub.Status = domain.SubStatus(temp.Status)
		sub.Price = null.NewInt(int(temp.Price.Int32), temp.Price.Valid)

		if err := suc.Update(ctx, sub); err != nil {
			b.Fatal(err)
		}

	}

	b.StopTimer()
	tests.DeleteAllSqlc(db)
	b.StartTimer()
}
