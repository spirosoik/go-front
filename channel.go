package front

import (
	"fmt"
	"net/http"
)

//ChannelService Http client
type ChannelService service

//ChannelType the source type of a channel
type ChannelType string

//Constants for channel types
const (
	SMTP     = ChannelType("smtp")
	Imap     = ChannelType("imap")
	Twilio   = ChannelType("twilio")
	Facebook = ChannelType("facebook")
	Smooch   = ChannelType("smooch")
	Intercom = ChannelType("intercom")
	Truly    = ChannelType("truly")
	//Only this supported right now by Front
	Custom = ChannelType("custom")
)

//Channel a DTO for API
type Channel struct {
	ID        string `json:"id"`
	Address   string `json:"address"`
	Type      string `json:"type"`
	SendAs    string `json:"send_as"`
	IsPrivate bool   `json:"is_private"`
	Links     Links  `json:"_links"`
}

//ChannelList which includes a list of Channel DTOs
type ChannelList struct {
	PageLinks Pagination `json:"_pagination"`
	Links     Links      `json:"_links"`
	Results   []Channel  `json:"_results"`
}

//ChannelCreatePayload to create a new channel
type ChannelCreatePayload struct {
	Type     ChannelType `json:"type"`
	Settings Settings    `json:"settings"`
}

//List channels of the api
func (s *ChannelService) List() (*ChannelList, error) {
	path := "channels/"

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}
	var c *ChannelList

	_, err = s.gateway.call(req, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

//Get by ID the channel details
func (s *ChannelService) Get(id string) (*Channel, error) {
	path := fmt.Sprintf("channels/%s", id)

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}

	var c *Channel
	_, err = s.gateway.call(req, &c)

	if err != nil {
		return nil, err
	}
	return c, nil
}

//Create a channel
func (s *ChannelService) Create(inboxID string, p *ChannelCreatePayload) error {
	path := fmt.Sprintf("inboxes/%s/channels", inboxID)

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
