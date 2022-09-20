package account

import (
	"context"
	"fmt"
	"net/http"

	"bitbucket.org/velmie/wallet-integration/internal/account_rpc"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAccountByID(ctx context.Context, id uint64) (*account_rpc.AccountResponse, error) {
	req := &account_rpc.AccountByIdReq{Id: id}
	return s.processor().GetAccountById(ctx, req)
}

func (s *Service) processor() account_rpc.Accounts {
	url := fmt.Sprintf("%s:%s", "localhost", "8181")
	return account_rpc.NewAccountsProtobufClient(url, http.DefaultClient)
}
