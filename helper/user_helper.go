package helper

import (
	"context"
	"fmt"
	"time"
	"path/filepath"

	"cloud.google.com/go/firestore"
	"Ecojourney-backend/config"
	"google.golang.org/api/iterator"
)

// GenerateUniqueFileName generates a unique filename based on timestamp.
func GenerateUniqueFileName(filename string) string {
	ext := filepath.Ext(filename)
	ts := time.Now().UnixNano() / int64(time.Millisecond)
	return fmt.Sprintf("%d%s", ts, ext)
}

func UpdateUserProfilePicture(userID, imageURL string) error {
	ctx := context.Background()
	usersCollection := config.FirestoreClient.Collection("users")

	userRef := usersCollection.Doc(userID)
	updateData := map[string]interface{}{
		"photo_url": imageURL,
		"updated_at": time.Now(),
	}

	_, err := userRef.Set(ctx, updateData, firestore.MergeAll)
	if err != nil {
		return fmt.Errorf("failed to update profile picture URL: %v", err)
	}

	return nil
}

func GetUserImageURL(uid string) (string, error) {
    ctx := context.Background()
    firestoreClient := config.FirestoreClient

    // Mendapatkan referensi koleksi 'users'
    usersCollection := firestoreClient.Collection("users")

    // Membuat query untuk mencari dokumen berdasarkan UID
    query := usersCollection.Where("uid", "==", uid).Limit(1)

    // Menjalankan query dan mendapatkan iterator untuk hasilnya
    iter := query.Documents(ctx)
    defer iter.Stop()

    // Mengambil dokumen pertama (diasumsikan hanya ada satu)
    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return "", fmt.Errorf("failed to iterate documents: %v", err)
        }

        // Mengambil nilai imageURL dari dokumen
        imageURL, ok := doc.Data()["imageURL"].(string)
        if !ok {
            return "", fmt.Errorf("imageURL not found or not a string")
        }

        return imageURL, nil
    }

    return "", fmt.Errorf("user with UID %s not found", uid)
}