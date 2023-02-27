package api

import (
	"encoding/json"
	"fmt"

	"github.com/zipstack/pct-plugin-framework/fwhelpers"
)

type SourceFakerID struct {
	SourceId string `json:"sourceId"`
}

type SourceFaker struct {
	Name                    string                `json:"name"`
	SourceDefinitionId      string                `json:"sourceDefinitionId"`
	SourceId                string                `json:"sourceId,omitempty"`
	WorkspaceId             string                `json:"workspaceId"`
	ConnectionConfiguration SourceFakerConnConfig `json:"connectionConfiguration"`
}

type SourceFakerConnConfig struct {
	Seed            int64 `json:"seed"`
	Count           int64 `json:"count"`
	RecordsPerSync  int64 `json:"records_per_sync"`
	RecordsPerSlice int64 `json:"records_per_slice"`
}

func (c *Client) CreateSource(payload SourceFaker) (SourceFaker, error) {
	logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/create"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceFaker{}, err
	}
	logger.Printf("create payload: %s", string(body))

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
