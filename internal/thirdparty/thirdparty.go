package thirdparty

import "context"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) OnTransferCallback(onTransfer func(ctx context.Context, id uint64, amount string, description string) error) error {
	// some third party vendor integration code here: subscription and so on
	return nil
}
