package tests

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	boilmodels "github.com/BON4/gosubs/internal/domain/boil_postgres"
	null "github.com/volatiletech/null/v8"
)

func ConnectTestDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/tgram_subs_test?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DeleteAllBoil(db *sql.DB) {
	if _, err := boilmodels.SubHistories().DeleteAll(context.TODO(), db); err != nil {
		panic(err)
	}

	if _, err := boilmodels.Subs().DeleteAll(context.TODO(), db); err != nil {
		panic(err)
	}
	if _, err := boilmodels.Accounts().DeleteAll(context.TODO(), db); err != nil {
		panic(err)
	}

	if _, err := boilmodels.Tgusers().DeleteAll(context.TODO(), db); err != nil {
		panic(err)
	}

}

func RandomUserBoil() *boilmodels.Tguser {
	id := int64(rand.Uint32())
	user1 := &boilmodels.Tguser{
		TelegramID: id,
		Username:   fmt.Sprintf("user_%d", id),
		Status:     boilmodels.UserStatusMember,
	}
	return user1
}

func RrandomAccountBoil(userID null.Int64) *boilmodels.Account {
	id := int64(rand.Uint32())
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("creator_pass%d", id)))

	user1 := &boilmodels.Account{
		UserID:   userID,
		Email:    fmt.Sprintf("creator_email%d", id),
		Password: h.Sum(nil),
		ChanName: null.StringFrom(fmt.Sprintf("creator_%d", id)),
		Role:     boilmodels.AccountRoleAdmin,
	}
	return user1
}

func RandomSubBoil(userID, creatorID int64) *boilmodels.Sub {
	t := time.Hour * time.Duration(rand.Intn(10000))
	sub := &boilmodels.Sub{
		UserID:      userID,
		AccountID:   creatorID,
		ActivatedAt: time.Now().Add(-t),
		ExpiresAt:   time.Now().Add(t),
		Status:      boilmodels.SubStatusActive,
		Price:       null.IntFrom(rand.Intn(1000)),
	}
	return sub
}
