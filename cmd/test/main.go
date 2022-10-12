package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/BON4/gosubs/internal/models"
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
	if _, err := models.Subs().DeleteAll(context.TODO(), boil.GetContextDB()); err != nil {
		panic(err)
	}

	if _, err := models.Tgusers().DeleteAll(context.TODO(), boil.GetContextDB()); err != nil {
		panic(err)
	}

	if _, err := models.Creators().DeleteAll(context.TODO(), boil.GetContextDB()); err != nil {
		panic(err)
	}
}

func main() {
	// Open handle to database like normal
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	boil.SetDB(db)

	deleteAll()

	user1 := models.Tguser{
		TelegramID: 12345,
		Username:   "test1_username",
		Status:     models.UserStatusCreator,
	}

	if err := user1.InsertG(context.TODO(), boil.Infer()); err != nil {
		panic(err)
	}

	user2 := models.Tguser{
		TelegramID: 12346,
		Username:   "test2_username",
		Status:     models.UserStatusMember,
	}

	if err := user2.InsertG(context.TODO(), boil.Infer()); err != nil {
		panic(err)
	}

	users, err := models.Tgusers().All(context.TODO(), boil.GetContextDB())
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("User: %+v\n", user)
	}

	creator := models.Creator{
		TelegramID: user1.TelegramID,
		Username:   user1.Username,
		Password:   null.BytesFrom([]byte("password")),
		Email:      null.StringFrom("email@mail.com"),
		ChanName:   null.StringFrom("test_channel"),
	}

	if err := creator.InsertG(context.TODO(), boil.Infer()); err != nil {
		panic(err)
	}

	creators, err := models.Creators().All(context.TODO(), boil.GetContextDB())
	if err != nil {
		panic(err)
	}

	for _, creator := range creators {
		fmt.Printf("Creator: %+v\n", creator)
	}

	sub := models.Sub{
		UserID:      user2.UserID,
		CreatorID:   creator.CreatorID,
		ActivatedAt: time.Now(),
		ExpiresAt:   time.Now().Add(time.Hour),
		Status:      models.SubStatusActive,
		Price:       null.StringFrom("5.99"),
	}

	if err := sub.InsertG(context.TODO(), boil.Infer()); err != nil {
		panic(err)
	}

	subs, err := models.Subs().All(context.TODO(), boil.GetContextDB())
	if err != nil {
		panic(err)
	}

	for _, sub := range subs {
		fmt.Printf("Sub: %+v\n", sub)
	}

	// deleteAll()
}
