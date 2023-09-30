package model

type Book struct {
	BookID string `json:"id"`     // json tag
	Title  string `json:"title"`  // json tag
	Author string `json:"author"` // json tag
	Desc   string `json:"desc"`   // json tag
}
