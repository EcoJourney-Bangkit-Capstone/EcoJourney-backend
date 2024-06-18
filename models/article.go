package models

import "time"

type ArticleRequest struct {
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	AuthorID string   `json:"authorId" binding:"required"`
	Topic    []string `json:"topic"`
}

type ArticleResponse struct {
	ID             string    `json:"id"`
	DatePublished  time.Time `json:"date_published"`
	Publisher      string    `json:"publisher"`
	Title          string    `json:"title"`
	Topic          []string  `json:"topic"`
	ImgURL         string    `json:"img_url"`
	Content        string    `json:"content"`
}