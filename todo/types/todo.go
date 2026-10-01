package types

type Todo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Due         string `json:"due"`
	Priority    string `json:"priority"`
	Description string `json:"description"`
}
