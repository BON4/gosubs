package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	creator_usecase "github.com/BON4/gosubs/internal/creator/usecase"
	"github.com/BON4/gosubs/internal/domain"
	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"
	sub_usecase "github.com/BON4/gosubs/internal/subscription/usecase/boil"
	tguser_usecase "github.com/BON4/gosubs/internal/tguser/usecase"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/tgram_subs?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func deleteAll() {
	if _, err := boilmodels.SubHistories().DeleteAll(context.TODO(), boil.GetContextDB()); err != nil {
		panic(err)
	}

	if _, err := boilmodels.Subs().DeleteAll(context.TODO(), boil.GetContextDB()); err != nil {
		panic(err)
	}

	if _, err := boilmodels.Tgusers().DeleteAll(context.TODO(), boil.GetContextDB()); err != nil {
		panic(err)
	}

	if _, err := boilmodels.Creators().DeleteAll(context.TODO(), boil.GetContextDB()); err != nil {
		panic(err)
	}
}

func randomUser() *boilmodels.Tguser {
	id := int64(rand.Uint32())
	user1 := &boilmodels.Tguser{
		TelegramID: id,
		Username:   fmt.Sprintf("user_%d", id),
		Status:     boilmodels.UserStatusMember,
	}
	return user1
}

func randomCreator() *boilmodels.Creator {
	id := int64(rand.Uint32())
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("creator_pass%d", id)))

	user1 := &boilmodels.Creator{
		TelegramID: id,
		Username:   fmt.Sprintf("creator_%d", id),
		Email:      null.StringFrom(fmt.Sprintf("creator_email%d", id)),
		Password:   null.BytesFrom(h.Sum(nil)),
		ChanName:   null.StringFrom(fmt.Sprintf("creator_%d", id)),
	}
	return user1
}

func randomSub(userID, creatorID int64) *boilmodels.Sub {
	t := time.Hour * time.Duration(rand.Intn(10000))
	sub := &boilmodels.Sub{
		UserID:      userID,
		CreatorID:   creatorID,
		ActivatedAt: time.Now().Add(-t),
		ExpiresAt:   time.Now().Add(t),
		Status:      boilmodels.SubStatusActive,
		Price:       null.IntFrom(rand.Intn(1000)),
	}
	return sub
}

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	boil.SetDB(db)

	deleteAll()

	ctx := context.TODO()

	uuc := tguser_usecase.NewBoilTgUserUsecase(db)
	cuc := creator_usecase.NewBoilCretorUsecase(db)
	suc := sub_usecase.NewBoilSubscriptionUsecase(db)

	users := make([]*domain.Tguser, 10)

	domainCreator := &domain.Creator{}
	domain.CreatorBoilToDomain(randomCreator(), domainCreator)

	subs := make([]*domain.Sub, 10)

	if err := cuc.Create(ctx, domainCreator); err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		users[i] = &domain.Tguser{}
		domain.TguserBoilToDomain(randomUser(), users[i])

		if err := uuc.Create(ctx, users[i]); err != nil {
			panic(err)
		}
		subs[i] = &domain.Sub{}
		domain.SubBoilToDomain(randomSub(users[i].UserID, domainCreator.CreatorID), subs[i])

		if err := suc.Create(ctx, subs[i]); err != nil {
			panic(err)
		}
	}

	from := 100
	to := 500

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

	for _, s := range found {
		fmt.Printf("%+v\n", s)
	}

	deleteAll()
}
