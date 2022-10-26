package domain

import (
	"context"
	"strconv"
	"strings"
)

type FindSubRequest struct {
	PriceRange *struct {
		From *int64 `json:"from,omitempty"`
		To   *int64 `json:"to,omitempty"`
	} `json:"price_range,omitempty"`

	Price *struct {
		Eq int64 `json:"eq"`
	} `json:"price,omitempty"`

	Status *struct {
		Eq   *string `json:"eq,omitempty"`
		Like *string `json:"like,omitempty"`
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

func ParseFindSubRequest(mapData map[string][]string) (FindSubRequest, error) {
	var req FindSubRequest = FindSubRequest{}

	req.Status = &struct {
		Eq   *string `json:"eq,omitempty"`
		Like *string `json:"like,omitempty"`
	}{}

	status_eq, ok := mapData["status_eq"]
	if ok && len(status_eq) > 0 {
		req.Status.Eq = &status_eq[0]
	}

	status_like, ok := mapData["status_like"]
	if ok && len(status_like) > 0 {
		req.Status.Like = &status_like[0]
	}

	acc_id, ok := mapData["account_id"]
	if ok && len(acc_id) > 0 {
		parsed_acc_id, err := strconv.ParseInt(acc_id[0], 10, 64)
		if err != nil {
			return FindSubRequest{}, err
		}

		req.AccountID = &struct {
			Eq int64 `json:"eq"`
		}{
			Eq: parsed_acc_id,
		}
	}

	usr_id, ok := mapData["user_id"]
	if ok && len(usr_id) > 0 {
		parsed_usr_id, err := strconv.ParseInt(usr_id[0], 10, 64)
		if err != nil {
			return FindSubRequest{}, err
		}

		req.TgUserID = &struct {
			Eq int64 `json:"eq"`
		}{
			Eq: parsed_usr_id,
		}
	}

	var parsed_page_number uint64
	var parsed_page_size uint64

	page_number, ok := mapData["page_number"]
	if ok && len(page_number) > 0 {
		var err error
		parsed_page_number, err = strconv.ParseUint(page_number[0], 10, 64)
		if err != nil {
			return FindSubRequest{}, err
		}
	}

	page_size, ok := mapData["page_size"]
	if ok && len(page_size) > 0 {
		var err error
		parsed_page_size, err = strconv.ParseUint(page_size[0], 10, 64)
		if err != nil {
			return FindSubRequest{}, err
		}
	}

	req.PageSettings = &struct {
		PageSize   uint `json:"page_size"`
		PageNumber uint `json:"page_number"`
	}{
		PageSize:   uint(parsed_page_size),
		PageNumber: uint(parsed_page_number),
	}

	var parsed_price_from int64
	var parsed_price_to int64
	price_range, ok := mapData["price_range"]

	if ok && len(page_number) > 0 {
		price_range = strings.Split(price_range[0], ",")

		req.PriceRange = &struct {
			From *int64 `json:"from,omitempty"`
			To   *int64 `json:"to,omitempty"`
		}{}

		price_from := price_range[0]

		var err error
		parsed_price_from, err = strconv.ParseInt(price_from, 10, 64)
		if err != nil {
			return FindSubRequest{}, err
		}
		req.PriceRange.From = &parsed_price_from

		price_to := price_range[len(price_range)-1]
		parsed_price_to, err = strconv.ParseInt(price_to, 10, 64)
		if err != nil {
			return FindSubRequest{}, err
		}
		req.PriceRange.To = &parsed_price_to

		if parsed_price_from == parsed_price_to {
			req.PriceRange = nil
			req.Price = &struct {
				Eq int64 `json:"eq"`
			}{
				Eq: parsed_price_from,
			}
		}
	}
	return req, nil
}

type FindUserRequest struct {
	Username *struct {
		Eq   *string `json:"eq,omitempty"`
		Like *string `json:"like,omitempty"`
	} `json:"username,omitempty"`

	Status *struct {
		Eq   *string `json:"eq,omitempty"`
		Like *string `json:"like,omitempty"`
	} `json:"status,omitempty"`

	PageSettings *struct {
		PageSize   uint `json:"page_size"`
		PageNumber uint `json:"page_number"`
	} `json:"page_settings"`
}

func ParseFindUserRequest(mapData map[string][]string) (FindUserRequest, error) {
	var req FindUserRequest = FindUserRequest{}

	req.Status = &struct {
		Eq   *string `json:"eq,omitempty"`
		Like *string `json:"like,omitempty"`
	}{}

	status_eq, ok := mapData["status_eq"]
	if ok && len(status_eq) > 0 {
		req.Status.Eq = &status_eq[0]
	}

	status_like, ok := mapData["status_like"]
	if ok && len(status_like) > 0 {
		req.Status.Like = &status_like[0]
	}

	req.Username = &struct {
		Eq   *string `json:"eq,omitempty"`
		Like *string `json:"like,omitempty"`
	}{}

	username_eq, ok := mapData["username_eq"]
	if ok && len(status_eq) > 0 {
		req.Username.Eq = &username_eq[0]
	}

	username_like, ok := mapData["username_like"]
	if ok && len(status_like) > 0 {
		req.Username.Like = &username_like[0]
	}

	var parsed_page_number uint64
	var parsed_page_size uint64

	page_number, ok := mapData["page_number"]
	if ok && len(page_number) > 0 {
		var err error
		parsed_page_number, err = strconv.ParseUint(page_number[0], 10, 64)
		if err != nil {
			return FindUserRequest{}, err
		}
	}

	page_size, ok := mapData["page_size"]
	if ok && len(page_size) > 0 {
		var err error
		parsed_page_size, err = strconv.ParseUint(page_size[0], 10, 64)
		if err != nil {
			return FindUserRequest{}, err
		}
	}

	req.PageSettings = &struct {
		PageSize   uint `json:"page_size"`
		PageNumber uint `json:"page_number"`
	}{
		PageSize:   uint(parsed_page_size),
		PageNumber: uint(parsed_page_number),
	}

	return req, nil
}

type FindAccountRequest struct {
	Role *struct {
		Eq   *string `json:"eq,omitempty"`
		Like *string `json:"like,omitempty"`
	} `json:"role,omitempty"`

	PageSettings *struct {
		PageSize   uint `json:"page_size"`
		PageNumber uint `json:"page_number"`
	} `json:"page_settings"`
}

func ParseFindAccountRequest(mapData map[string][]string) (FindAccountRequest, error) {
	var req FindAccountRequest = FindAccountRequest{}

	req.Role = &struct {
		Eq   *string `json:"eq,omitempty"`
		Like *string `json:"like,omitempty"`
	}{}

	role_eq, ok := mapData["role_eq"]
	if ok && len(role_eq) > 0 {
		req.Role.Eq = &role_eq[0]
	}

	role_like, ok := mapData["role_like"]
	if ok && len(role_like) > 0 {
		req.Role.Like = &role_like[0]
	}

	var parsed_page_number uint64
	var parsed_page_size uint64

	page_number, ok := mapData["page_number"]
	if ok && len(page_number) > 0 {
		var err error
		parsed_page_number, err = strconv.ParseUint(page_number[0], 10, 64)
		if err != nil {
			return FindAccountRequest{}, err
		}
	}

	page_size, ok := mapData["page_size"]
	if ok && len(page_size) > 0 {
		var err error
		parsed_page_size, err = strconv.ParseUint(page_size[0], 10, 64)
		if err != nil {
			return FindAccountRequest{}, err
		}
	}

	req.PageSettings = &struct {
		PageSize   uint `json:"page_size"`
		PageNumber uint `json:"page_number"`
	}{
		PageSize:   uint(parsed_page_size),
		PageNumber: uint(parsed_page_number),
	}

	return req, nil
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
