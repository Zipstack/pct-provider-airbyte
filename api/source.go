package api

import (
	"encoding/json"
	"fmt"
)

type SourceFakerID struct {
	SourceId string `json:"sourceId"`
}

type SourceFaker struct {
	Name                    string                `cty:"name"`
	SourceDefinitionId      string                `cty:"sourceDefinitionId"`
	SourceId                string                `cty:"sourceId"`
	WorkspaceId             string                `cty:"workspaceId"`
	ConnectionConfiguration SourceFakerConnConfig `cty:"connectionConfiguration"`
}

type SourceFakerConnConfig struct {
	Seed            int64 `cty:"seed"`
	Count           int64 `cty:"count"`
	RecordsPerSync  int64 `cty:"records_per_sync"`
	RecordsPerSlice int64 `cty:"records_per_slice"`
}

func (c *Client) CreateSource(payload SourceFaker) (SourceFaker, error) {
	method := "POST"
	url := c.Host + "/api/v1/sources/create"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceFaker{}, err
	}

	b, statusCode, status, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceFaker{}, err
	}

	source := SourceFaker{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &source)
		return source, err
	} else {
		return source, fmt.Errorf(status)
	}
}

func (c *Client) ReadSource(sourceId string) (SourceFaker, error) {
	method := "POST"
	url := c.Host + "/api/v1/sources/get"
	sId := SourceFakerID{sourceId}
	body, err := json.Marshal(sId)
	if err != nil {
		return SourceFaker{}, err
	}

	b, statusCode, status, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceFaker{}, err
	}

	source := SourceFaker{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &source)
		return source, err
	} else {
		return source, fmt.Errorf(status)
	}
}

func (c *Client) UpdateSource() {}

func (c *Client) DeleteSource() {}
