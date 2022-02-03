package service

import "conference/model"

func (s *Service) CreateConference(conference model.CreateConferenceReq) (*model.Conference, error) {
	return s.conferenceRepo.SaveConference(model.Conference{
		Title:       conference.Title,
		Description: conference.Description,
		StartDate:   conference.StartDate,
		EndDate:     conference.EndDate,
	})
}
