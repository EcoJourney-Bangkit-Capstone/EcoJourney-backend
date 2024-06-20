package helper

import (
	"context"
	"time"

	"google.golang.org/api/iterator"
	"Ecojourney-backend/config"
)

type WasteHistory struct {
	HistoryID  string    `json:"historyId"`
	TotalType  int       `json:"total_type"`
	Type       []string  `json:"type"`
	ImageURL   string    `json:"imageURL"`
	Timestamp  time.Time `json:"timestamp"`
}

type History struct {
	UserID     string    `firestore:"userId"`
	ImageURL   string    `firestore:"imageURL"`
	WasteTypes []string  `firestore:"wasteTypes"`
	Timestamp  time.Time `firestore:"timestamp"`
}

func SaveHistory(uid, imageURL string, wasteTypes []string) (string, error) {
	ctx := context.Background()
	history := History{
		UserID:     uid,
		ImageURL:   imageURL,
		WasteTypes: wasteTypes,
		Timestamp:  time.Now(),
	}

	doc, _, err := config.FirestoreClient.Collection("wasteHistory").Add(ctx, history)
	if err != nil {
		return "", err
	}
	return doc.ID, nil
}

func GetWasteHistory(uid string) ([]WasteHistory, error) {
	ctx := context.Background()
	var histories []WasteHistory

	iter := config.FirestoreClient.Collection("wasteHistory").Where("userId", "==", uid).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var history History
		doc.DataTo(&history)

		histories = append(histories, WasteHistory{
			HistoryID: doc.Ref.ID,
			TotalType: len(history.WasteTypes),
			Type:      history.WasteTypes,
			ImageURL:  history.ImageURL,
			Timestamp: history.Timestamp,
		})
	}

	return histories, nil
}
