package user

type Item struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        *bool  `json:"false"`
}

type ItemList struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type TodoList struct {
}