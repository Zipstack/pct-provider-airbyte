package api

import (
	"encoding/json"
	"fmt"
)

type SourceFreshdeskID struct {
	SourceId string `json:"sourceId"`
}

type SourceFreshdesk struct {
	Name                    string                    `json:"name"`
	SourceId                string                    `json:"sourceId,omitempty"`
	SourceDefinitionId      string                    `json:"sourceDefinitionId,omitempty"`
	WorkspaceId             string                    `json:"workspaceId,omitempty"`
	ConnectionConfiguration SourceFreshdeskConnConfig `json:"connectionConfiguration"`
}

type SourceFreshdeskConnConfig struct {
	Domain            string `json:"domain"`
	StartDate         string `json:"start_date"`
	ApiKey            string `json:"api_key"`
	RequestsPerMinute int    `json:"requests_per_minute,omitempty"`
}

func (c *Client) CreateFreshdeskSource(payload SourceFreshdesk) (SourceFreshdesk, error) {
	// logger := fwhelpers.GetLogger()
	method := "POST"
	url := c.Host + "/api/v1/sources/create"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceFreshdesk{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceFreshdesk{}, err
	}
	source := SourceFreshdesk{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &source)
		return source, err
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return source, err
		} else {
			return source, fmt.Errorf(msg)
		}
	}
}

func (c *Client) ReadFreshdeskSource(sourceId string) (SourceFreshdesk, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/get"
	sId := SourceFreshdeskID{sourceId}
	body, err := json.Marshal(sId)
	if err != nil {
		return SourceFreshdesk{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceFreshdesk{}, err
	}

	source := SourceFreshdesk{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &source)
		return source, err
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return source, err
		} else {
			return source, fmt.Errorf(msg)
		}
	}
}

func (c *Client) UpdateFreshdeskSource(payload SourceFreshdesk) (SourceFreshdesk, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/update"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceFreshdesk{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceFreshdesk{}, err
	}

	source := SourceFreshdesk{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &source)
		return source, err
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return source, err
		} else {
			return source, fmt.Errorf(msg)
		}
	}
}

func (c *Client) DeleteFreshdeskSource(sourceId string) error {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/delete"
	sId := SourceFreshdeskID{sourceId}
	body, err := json.Marshal(sId)
	if err != nil {
		return err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return err
	}

	if statusCode >= 200 && statusCode <= 299 {
		return nil
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return err
		} else {
			return fmt.Errorf(msg)
		}
	}
}
