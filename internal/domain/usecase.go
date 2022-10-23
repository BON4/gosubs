package domain

import (
	"context"
)

type FindSubRequest struct {
	// Username *struct {
	// 	Like string `json:"LIKE"`
	// 	eq   string `json:"EQ"`
	// } `json:"username"`
	// Email *struct {
	// 	Like string `json:"LIKE"`
	// 	eq   string `json:"EQ"`
	// } `json:"email"`
	// Role *struct{
	// 	eq   string `json:"EQ"`
	// } `json:"role"`
	// ID *struct {
	// 	eq int `json:"EQ"`
	// } `json:"id"`

	Price *struct {
		Eq    *int `json:"eq,omitempty"`
		Range *struct {
			From *int `json:"from,omitempty"`
			To   *int `json:"to,omitempty"`
		} `json:"range,omitempty"`
	} `json:"price,omitempty"`

	Status *struct {
		Eq string `json:"eq"`
	} `json:"status,omitempty"`

	AccountID *struct {
		Eq int64 `json:"eq"`
	} `json:"account_id,omitempty"`

	TgUserID *struct {
		Eq int64 `json:"eq"`
	} `json:"tguser_id,omitempty"`

	PageSettings *struct {
		PageSize   uint `json:"page_size"`
		PageNumber uint `json:"page_number"`
	} `json:"page_settings"`
}

type FindUserRequest struct {
	// Username *struct {
	// 	Like string `json:"LIKE"`
	// 	Eq   string `json:"EQ"`
	// } `json:"username"`
	// Email *struct {
	// 	Like string `json:"LIKE"`
	// 	Eq   string `json:"EQ"`
	// } `json:"email"`
	PageSettings *struct {
		PageSize   uint `json:"page_size"`
		PageNumber uint `json:"page_number"`
	} `json:"page_settings"`
}

type FindAccountRequest struct {
	// Username *struct {
	// 	Like string `json:"LIKE"`
	// 	Eq   string `json:"EQ"`
	// } `json:"username"`
	// Email *struct {
	// 	Like string `json:"LIKE"`
	// 	Eq   string `json:"EQ"`
	// } `json:"email"`
	PageSettings *struct {
		PageSize   uint `json:"page_size"`
		PageNumber uint `json:"page_number"`
	} `json:"page_settings"`
}

type TgUserUsecase interface {
	GetByID(ctx context.Context, id int64) (*Tguser, error)

	GetByTelegramID(ctx context.Context, id int64) (*Tguser, error)

	//Create - will create new user.
	Create(ctx context.Context, tguser *Tguser) error

	// Delete - will delete user. Subscription will be deleted also.
	Delete(ctx context.Context, id int64) error

	Update(ctx context.Context, tguser *Tguser) error

	List(ctx context.Context, cond FindUserRequest) ([]*Tguser, error)
}

type SubscriptionUsecase interface {
	GetByID(ctx context.Context, userID int64, creatorID int64) (*Sub, error)

	// Create - creates subscribtion
	Create(ctx context.Context, sub *Sub) error

	// Save - saves subscription to history table.
	Save(ctx context.Context, sub *Sub) (int64, error)

	Update(ctx context.Context, sub *Sub) error

	Delete(ctx context.Context, userID int64, creatorID int64) error

	List(ctx context.Context, cond FindSubRequest) ([]*Sub, error)
}

type AccountUsecase interface {
	GetByEmail(ctx context.Context, email string) (*Account, error)

	GetByID(ctx context.Context, id int64) (*Account, error)

	GetByTelegramID(ctx context.Context, id int64) (*Account, error)

	GetUser(ctx context.Context, id int64) (*Tguser, error)

	Create(ctx context.Context, creator *Account) error

	// Delete - will delete creator. Subscriptions will be deleted also.
	Delete(ctx context.Context, id int64) error

	Update(ctx context.Context, tguser *Account) error

	List(ctx context.Context, cond FindAccountRequest) ([]*Account, error)
}
