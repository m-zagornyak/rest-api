package item

import "github.com/m-zagornyak/rest-api.git/internal/domain/item"

type itemStorage struct {
}

func (bs *itemStorage) GetOne(uuid string) *item.TodoItem {
	return nil
}

func (bs *itemStorage) Create(item *item.TodoItem) *item.TodoItem {
	return nil
}

func (bs *itemStorage) Delete(item *item.TodoItem) error {
	return nil
}
