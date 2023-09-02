package service

import (
	BIP_project "github.com/nekitalek/bip_project/backend"
	"github.com/nekitalek/bip_project/backend/internal/repository"
)

type EventItemService struct {
	repo repository.EventItem
}

func NewEventItemService(repo repository.EventItem) *EventItemService {
	return &EventItemService{repo: repo}
}

func (s *EventItemService) CreateEvent(user_id int, item *BIP_project.Event_items) (int, error) {
	item.Admin = user_id
	return s.repo.CreateEvent(item)
}

func (s *EventItemService) GetEvents(input *BIP_project.Event_items_input) ([]BIP_project.Event_items, error) {
	return s.repo.GetEvents(input)
}

func (s *EventItemService) UpdateEvent(userId, itemId int, input *BIP_project.Event_items_input) error {
	return s.repo.UpdateEvent(userId, itemId, input)
}

func (s *EventItemService) DeleteEvent(userId, itemId int) error {
	return s.repo.DeleteEvent(userId, itemId)
}
