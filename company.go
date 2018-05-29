package front

import (
	"net/http"

	"github.com/spirosoik/go-front/front/serializer"
)

//CompanyService Http client
type CompanyService service

//Company DTO response
type Company struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links Links  `json:"_links"`
}

//Me endpoint for service
func (s *CompanyService) Me() (*Company, error) {
	path := "me/"

	req, err := s.gateway.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.gateway.client.Do(req)
	if err != nil {
		return nil, err
	}
	var c *Company
	deserialize := serializer.Decode(&c)

	if err := deserialize(resp); err != nil {
		return c, err
	}

	return c, nil
}
