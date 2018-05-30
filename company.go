package front

import (
	"net/http"
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

	var c *Company

	_, err = s.gateway.call(req, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
