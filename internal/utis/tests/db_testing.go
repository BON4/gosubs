package tests

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"

	"github.com/volatiletech/null/v8"
)

func ConnectTestDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/tgram_subs_test?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DeleteAll(db *sql.DB) {
	if _, err := boilmodels.SubHistories().DeleteAll(context.TODO(), db); err != nil {
		panic(err)
	}

	if _, err := boilmodels.Subs().DeleteAll(context.TODO(), db); err != nil {
		panic(err)
	}

	if _, err := boilmodels.Tgusers().DeleteAll(context.TODO(), db); err != nil {
		panic(err)
	}

	if _, err := boilmodels.Creators().DeleteAll(context.TODO(), db); err != nil {
		panic(err)
	}
}

func RandomUser() *boilmodels.Tguser {
	id := int64(rand.Uint32())
	user1 := &boilmodels.Tguser{
		TelegramID: id,
		Username:   fmt.Sprintf("user_%d", id),
		Status:     boilmodels.UserStatusMember,
	}
	return user1
}

func RrandomCreator() *boilmodels.Creator {
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

func RandomSub(userID, creatorID int64) *boilmodels.Sub {
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
