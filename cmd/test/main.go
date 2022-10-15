package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	creator_usecase "github.com/BON4/gosubs/internal/creator/usecase"
	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"
	sub_usecase "github.com/BON4/gosubs/internal/subscription/usecase"
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

	uuc := tguser_usecase.NewTgUserUsecase(db)
	cuc := creator_usecase.NewCretorUsecase(db)
	suc := sub_usecase.NewBoilSubscriptionUsecase(db)

	users := make([]*boilmodels.Tguser, 10)
	creator := randomCreator()
	subs := make([]*boilmodels.Sub, 10)

	if err := cuc.Create(ctx, creator); err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		users[i] = randomUser()
		if err := uuc.Create(ctx, users[i]); err != nil {
			panic(err)
		}

		subs[i] = randomSub(users[i].UserID, creator.CreatorID)

		if err := suc.Create(ctx, subs[i]); err != nil {
			panic(err)
		}
	}

	from := 100
	to := 500

	found, err := suc.List(ctx, boilmodels.FindSubRequest{
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
}

func main1() {
	// Open handle to database like normal
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	boil.SetDB(db)

	deleteAll()

	user1 := boilmodels.Tguser{
		TelegramID: 12345,
		Username:   "test1_username",
		Status:     boilmodels.UserStatusCreator,
	}

	user2 := boilmodels.Tguser{
		TelegramID: 12346,
		Username:   "test2_username",
		Status:     boilmodels.UserStatusMember,
	}

	if err := user2.InsertG(context.TODO(), boil.Infer()); err != nil {
		panic(err)
	}

	users, err := boilmodels.Tgusers().All(context.TODO(), boil.GetContextDB())
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("User: %+v\n", user)
	}

	creator := boilmodels.Creator{
		TelegramID: user1.TelegramID,
		Username:   user1.Username,
		Password:   null.BytesFrom([]byte("password")),
		Email:      null.StringFrom("email@mail.com"),
		ChanName:   null.StringFrom("test_channel"),
	}

	if err := creator.InsertG(context.TODO(), boil.Infer()); err != nil {
		panic(err)
	}

	creators, err := boilmodels.Creators().All(context.TODO(), boil.GetContextDB())
	if err != nil {
		panic(err)
	}

	for _, creator := range creators {
		fmt.Printf("Creator: %+v\n", creator)
	}

	sub := boilmodels.Sub{
		UserID:      user2.UserID,
		CreatorID:   creator.CreatorID,
		ActivatedAt: time.Now(),
		ExpiresAt:   time.Now().Add(time.Hour),
		Status:      boilmodels.SubStatusActive,
		Price:       null.IntFrom(599),
	}

	if err := sub.InsertG(context.TODO(), boil.Infer()); err != nil {
		panic(err)
	}

	subs, err := boilmodels.Subs().All(context.TODO(), boil.GetContextDB())
	if err != nil {
		panic(err)
	}

	for _, sub := range subs {
		fmt.Printf("Sub: %+v\n", sub)
	}

	// deleteAll()
}
