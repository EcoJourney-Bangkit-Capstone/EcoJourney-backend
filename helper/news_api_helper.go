package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Article struct {
	ID           string   `json:"id"`
	DatePublished string  `json:"date_published"`
	Publisher    string   `json:"publisher"`
	Title        string   `json:"title"`
	Topic        []string `json:"topic"`
	ImgURL       string   `json:"img_url"`
	Content      string   `json:"content"`
}

type NewsAPIResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func FetchArticlesByKeyword(keyword string) ([]Article, error) {
	apiKey := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", keyword, apiKey)

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

	return newsResponse.Articles, nil
}
