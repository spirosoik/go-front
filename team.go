package front

import (
	"fmt"
	"net/http"
)

//TeamService Http client
type TeamService service

//Team is the group of teammates
type Team struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links Links  `json:"_links"`
}

//TeamList the list with available teams in the company
type TeamList struct {
	Results []Team `json:"_results"`
	Links   Links  `json:"_links"`
}

//List all available teams
func (s *TeamService) List() (*TeamList, error) {
	path := "teams/"

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var t *TeamList

	_, err = s.gateway.call(req, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

//Get by ID the requested team
func (s *TeamService) Get(id string) (*Team, error) {
	path := fmt.Sprintf("teams/%s", id)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var t *Team

	_, err = s.gateway.call(req, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
