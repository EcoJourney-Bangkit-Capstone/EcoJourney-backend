package helper

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"Ecojourney-backend/config"
	"fmt"
)

type WasteHistory struct {
	HistoryID       string    `json:"historyId"`
	UserID          string    `json:"userID"`
	RecognizedWaste string    `json:"recognizedWaste"`
	Timestamp       time.Time `json:"timestamp"`
	Article         Article   `json:"article"`
}

func SaveHistory(userID, imageURL, wasteType string, articles []Article) (string, error) {
	ctx := context.Background()
	historyCollection := config.FirestoreClient.Collection("wasteHistory")

	history := WasteHistory{
		UserID:          userID,
		RecognizedWaste: wasteType,
		Timestamp:       time.Now(),
		Article:         articles[0], // Assume the first article is relevant
	}

	doc, _, err := historyCollection.Add(ctx, history)
	if err != nil {
		return "", fmt.Errorf("failed to save history: %v", err)
	}

	return doc.ID, nil
}


func GetWasteHistory(userID string) ([]WasteHistory, error) {
	ctx := context.Background()
	historyCollection := config.FirestoreClient.Collection("wasteHistory")

	query := historyCollection.Where("userID", "==", userID).OrderBy("Timestamp", firestore.Desc)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve history: %v", err)
	}

	var history []WasteHistory
	for _, doc := range docs {
		var entry WasteHistory
		doc.DataTo(&entry)
		entry.HistoryID = doc.Ref.ID
		history = append(history, entry)
	}

	return history, nil
}
