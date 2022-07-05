package item

import (
	"context"
	"github.com/m-zagornyak/rest-api.git/internal/domain/item"
)

type Service interface {
	GetItemUUID(ctx context.Context, uuid string) *item.TodoItem
	CreateItemDTO(ctx context.Context, dto *item.CreateItemDTO) *item.TodoItem
}
