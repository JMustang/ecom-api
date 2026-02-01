package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
	// TODO: Repository
}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListProducts(ctx context.Context) error {
	return nil
}
