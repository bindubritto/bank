package api

import (
	db "github.com/bindubritto/bank/db/sqlc"
	"github.com/bindubritto/bank/utils"
)

func randomAccount(owner string) db.Account {
	return db.Account{
		ID:       utils.RandomInt(1, 1000),
		Owner:    owner,
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}
