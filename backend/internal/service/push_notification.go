package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	BIP_project "github.com/nekitalek/bip_project/backend"
	"github.com/nekitalek/bip_project/backend/internal/repository"
)

type PushNotificationService struct {
	repo repository.PushNotification
}

func NewPushNotificationService(repo repository.PushNotification) *PushNotificationService {
	return &PushNotificationService{repo: repo}
}

func (s *PushNotificationService) CreatePushNotification(user_id int, token string) error {
	_, err := s.repo.CreatePushNotification(user_id, token)
	return err
}
func (s *PushNotificationService) DeletePushNotification(user_id int, token string) error {
	return s.repo.DeletePushNotification(user_id, token)
}

func (s *PushNotificationService) SendPushNotification(event_id int) error {
	//получить tokens из бд
	tokens, err := s.repo.GetPushNotification(event_id, BIP_project.Confirmed)
	// отправить http запрос к firebase
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	message := map[string]interface{}{
		"notification": map[string]string{
			"title": "new player join",
			"body":  "new player join",
		},
		"registration_ids": tokens,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		return err
	}
	// log.Println(string(bytesRepresentation))

	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(bytesRepresentation))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("AAAA8QK0YKY:APA91bE30uRqIgS1OzNsz6wEEcJRaPl6yw7hfhFWsOhFFe-vb8q3_deqokyXi7Zx0p1y5Ijc2dVNUb1XpM94P5a42ndCwr9kZYfEvCSmsZkrkGks3mlsFmUA6z7h2M-W-Y3Rtj2-N2aZ"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// log.Printf("client: status code: %d\n", resp.StatusCode)
	// log.Println(resp.Body)
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	// log.Println(result)
	return nil
}
