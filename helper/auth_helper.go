package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ConvertCustomTokenToIDToken(email string, password string) (string, error) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=" + os.Getenv("FIREBASE_API_KEY")
	payload := strings.NewReader(fmt.Sprintf(`{"email":"%s","password":"%s","returnSecureToken":true}`, email, password))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New(response["error"].(map[string]interface{})["message"].(string))
	}

	return response["idToken"].(string), nil
}
