package model

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}
