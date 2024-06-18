package models

type ArticleRequest struct {
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	AuthorID string   `json:"authorId" binding:"required"`
	Topic    []string `json:"topic"`
}
