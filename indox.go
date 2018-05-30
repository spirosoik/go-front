package front

import (
	"fmt"
	"net/http"
)

//InboxService Http client
type InboxService service

//Inbox DTO for API
type Inbox struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
}

//InboxList the list DTO for Inbox items
type InboxList struct {
	Links   Links   `json:"_links"`
	Results []Inbox `json:"_results"`
}

//InboxCreatePayload DTO to create an Inbox
type InboxCreatePayload struct {
	Name        string   `json:"name"`
	TeamMateIds []string `json:"teammate_ids"`
}

//List inboxes of the api
func (s *InboxService) List() (*InboxList, error) {
	path := "inboxes/"

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var i *InboxList

	_, err = s.gateway.call(req, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

//Get by ID the inbox details
func (s *InboxService) Get(id string) (*Inbox, error) {
	path := fmt.Sprintf("inboxes/%s", id)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}
	var i *Inbox

	_, err = s.gateway.call(req, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

//Create an inbox of the api
func (s *InboxService) Create(p *InboxCreatePayload) error {
	path := "inboxes/"

	req, err := s.gateway.newRequest(http.MethodPost, path, p)
	if err != nil {
		return err
	}

	_, err = s.gateway.call(req, nil)
	if err != nil {
		return err
	}
	return nil
}

//Channels which are assigned to the provided Inbox id
func (s *InboxService) Channels(inboxID string) (*ChannelList, error) {
	path := fmt.Sprintf("inboxes/%s/channels", inboxID)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var chnl *ChannelList

	_, err = s.gateway.call(req, &chnl)
	if err != nil {
		return nil, err
	}
	return chnl, nil
}

//Conversations which are assigned to the provided Inbox id
func (s *InboxService) Conversations(inboxID string) (*ConversationList, error) {
	path := fmt.Sprintf("inboxes/%s/conversations", inboxID)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var c *ConversationList

	_, err = s.gateway.call(req, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

//Teammates which are assigned to the provided Inbox id
func (s *InboxService) Teammates(inboxID string) (*TeammateList, error) {
	path := fmt.Sprintf("inboxes/%s/teammates", inboxID)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var t *TeammateList

	_, err = s.gateway.call(req, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
