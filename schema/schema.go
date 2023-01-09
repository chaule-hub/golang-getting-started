package schema

type ToDo struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type ToDoList []ToDo

type ToDoDO struct {
	Text string `json:"text"`
}
