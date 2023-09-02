package service

import (
	BIP_project "github.com/nekitalek/bip_project/backend"
	"github.com/nekitalek/bip_project/backend/internal/repository"
)

type InvitationService struct {
	repo repository.Invitation
}

func NewInvitationService(repo repository.Invitation) *InvitationService {
	return &InvitationService{repo: repo}
}

func (s *InvitationService) CreateInvitation(user_id int, input *BIP_project.Event_invitations_input) (int, error) {
	return s.repo.CreateInvitation(user_id, input)
}
func (s *InvitationService) GetInvitation(user_id int) ([]BIP_project.Event_invitations, error) {
	return s.repo.GetInvitation(user_id)
}
func (s *InvitationService) UpdateInvitation(user_id, event_invitations_id int, input *BIP_project.Event_invitations_input) error {
	return s.repo.UpdateInvitation(user_id, event_invitations_id, input)
}
func (s *InvitationService) DeleteInvitation(user_id int, input *BIP_project.Event_invitations_input) error {
	return s.repo.DeleteInvitation(user_id, input)
}
