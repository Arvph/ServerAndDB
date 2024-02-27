package models

// User module defenition
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"Title"`
	Author  string `json:"Author"`
	Content string `json:"content"`
}
