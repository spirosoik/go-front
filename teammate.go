package front

import (
	"fmt"
	"net/http"

	"github.com/spirosoik/go-front/front/serializer"
)

//TeammateService Http client
type TeammateService service

//Teammate item
type Teammate struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	IsAdmin     bool   `json:"is_admin"`
	IsAvailable bool   `json:"is_available"`
	Links       Links  `json:"_links"`
}

//TeammateList the list with available teammates in the company
type TeammateList struct {
	Results []Teammate `json:"_results"`
	Links   Links      `json:"_links"`
}

//TeammateUpdatePayload to update the related teammate options
type TeammateUpdatePayload struct {
	Username    string `json:"username,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	IsAdmin     *bool  `json:"is_admin,omitempty"`
	IsAvailable *bool  `json:"is_available,omitempty"`
}

//List all available teammates
func (s *TeammateService) List() (*TeammateList, error) {
	path := "teammates/"

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.gateway.client.Do(req)
	if err != nil {
		return nil, err
	}
	var c *TeammateList
	deserialize := serializer.Decode(&c)

	if err := deserialize(resp); err != nil {
		return c, err
	}

	return c, nil
}

//Get by ID the requested teammate
func (s *TeammateService) Get(id string) (*Teammate, error) {
	path := fmt.Sprintf("teammates/%s", id)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.gateway.client.Do(req)
	if err != nil {
		return nil, err
	}
	var c *Teammate
	deserialize := serializer.Decode(&c)

	if err := deserialize(resp); err != nil {
		return c, err
	}

	return c, nil
}

//Update the teammate with the requested ID
func (s *TeammateService) Update(id string, p *TeammateUpdatePayload) error {
	path := fmt.Sprintf("teammates/%s", id)

	req, err := s.gateway.newRequest(http.MethodPatch, path, p)
	if err != nil {
		return err
	}

	_, err = s.gateway.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

//Conversations of requested ID of the teammate
func (s *TeammateService) Conversations(id string) (*ConversationList, error) {
	path := fmt.Sprintf("teammates/%s", id)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.gateway.client.Do(req)
	if err != nil {
		return nil, err
	}
	var c *ConversationList
	deserialize := serializer.Decode(&c)

	if err := deserialize(resp); err != nil {
		return c, err
	}

	return c, nil
}

//Inboxes of requested ID of the teammate
func (s *TeammateService) Inboxes(id string) (*InboxList, error) {
	path := fmt.Sprintf("teammates/%s", id)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.gateway.client.Do(req)
	if err != nil {
		return nil, err
	}
	var c *InboxList
	deserialize := serializer.Decode(&c)

	if err := deserialize(resp); err != nil {
		return c, err
	}

	return c, nil
}
