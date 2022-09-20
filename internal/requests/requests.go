package requests

import (
	"context"
	"fmt"
	"net/http"

	"bitbucket.org/velmie/wallet-integration/internal/requests_rpc"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) DebitAccount(ctx context.Context, id uint64, amount string, description string) (*requests_rpc.RequestResponse, error) {
	req := &requests_rpc.CreateDARequest{
		AccountId:   id,
		Amount:      amount,
		Description: description,
	}
	return s.processor().CreateDA(ctx, req)
}

func (s *Service) processor() requests_rpc.Requests {
	url := fmt.Sprintf("%s:%s", "localhost", "8181")
	return requests_rpc.NewRequestsProtobufClient(url, http.DefaultClient)
}
