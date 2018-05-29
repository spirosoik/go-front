package front

import (
	"fmt"
	"net/http"

	"github.com/spirosoik/go-front/front/serializer"
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

	resp, err := s.gateway.client.Do(req)
	if err != nil {
		return nil, err
	}
	var c *TeamList
	deserialize := serializer.Decode(&c)

	if err := deserialize(resp); err != nil {
		return c, err
	}

	return c, nil
}

//Get by ID the requested team
func (s *TeamService) Get(id string) (*Team, error) {
	path := fmt.Sprintf("teams/%s", id)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.gateway.client.Do(req)
	if err != nil {
		return nil, err
	}
	var c *Team
	deserialize := serializer.Decode(&c)

	if err := deserialize(resp); err != nil {
		return c, err
	}

	return c, nil
}
