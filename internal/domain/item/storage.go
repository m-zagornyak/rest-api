package item

type Storage interface {
	GetOne(uuid string) *TodoItem
	Create(item *TodoItem) *TodoItem
	Delete(item *TodoItem) error
}
