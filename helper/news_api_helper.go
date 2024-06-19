package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Article struct {
	ID           string   `json:"id"`
	DatePublished string  `json:"publishedAt"`
	Publisher    string   `json:"author"`
	Title        string   `json:"title"`
	Topic        []string `json:"topic"`
	ImgURL       string   `json:"urlToImage"`
	Content      string   `json:"content"`
}

type NewsAPIResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func FetchArticlesByKeyword(keyword string) ([]Article, error) {
	apiKey := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?q=%s&category=health&apiKey=%s", keyword, apiKey)
	// https://newsapi.org/v2/top-headlines?q=%s&category=health&environtment=en&apiKey=

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch articles: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %s", resp.Status)
	}

	var newsResponse NewsAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&newsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Limit to first 10 articles
	articles := newsResponse.Articles
	if len(articles) > 10 {
		articles = articles[:10]
	}

	// Assign IDs from 1 to 10
	for i := 0; i < len(articles); i++ {
		articles[i].ID = fmt.Sprintf("%d", i+1)
	}

	return articles, nil
}