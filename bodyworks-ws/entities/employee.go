package entities

type Employee struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Lastname string `json:"lastname"`
	Care string `json:"care"`
}
