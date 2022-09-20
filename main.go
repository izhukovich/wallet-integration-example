package main

import (
	"context"
	"errors"

	"bitbucket.org/otom/pkg/mod/github.com/shopspring/decimal@v1.3.1"

	"bitbucket.org/velmie/wallet-integration/internal/account"
	"bitbucket.org/velmie/wallet-integration/internal/requests"
	"bitbucket.org/velmie/wallet-integration/internal/thirdparty"
)

func main() {
	accountService := account.NewService()
	requestService := requests.NewService()
	thirdpartyService := thirdparty.NewService()

	if err := thirdpartyService.OnTransferCallback(
		func(ctx context.Context, id uint64, amount string, description string) error {
			acc, err := accountService.GetAccountByID(ctx, id)
			if err != nil {
				return err
			}

			amountDec, _ := decimal.NewFromString(amount)
			availableDec, _ := decimal.NewFromString(acc.AvailableAmount)

			if availableDec.Sub(amountDec).LessThan(decimal.Zero) {
				return errors.New("not enough funds")
			}

			_, err = requestService.DebitAccount(ctx, id, amount, description)
			return err
		}); err != nil {
		panic(err)
	}
}
