package client

import (
	"github.com/matheuspolitano/thetvdb/resource"
)

func (c *Client) ListCompanies() ([]resource.Companie, error) {
	req, err := c.NewRequest("GET", "/companies?page=0", nil)
	if err != nil {
		return nil, err
	}

	var responseCompanie resource.ResponseCompanie
	if err = c.Do(req, &responseCompanie); err != nil {
		return nil, err
	}
	return responseCompanie.Data, nil

}
