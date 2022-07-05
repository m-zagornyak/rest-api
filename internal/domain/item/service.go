package item

import (
	"context"
	"github.com/m-zagornyak/rest-api.git/internal/adapters/api/item"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) item.Service {
	return &service{storage: storage}
}

func (s *service) GetItemUUID(ctx context.Context, uuid string) *TodoItem {
	return s.storage.GetOne(uuid)
}

func (s *service) CreateItemDTO(ctx context.Context, dto *CreateItemDTO) *TodoItem {
	return nil
}
