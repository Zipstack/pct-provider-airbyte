package api

import (
	"encoding/json"
	"fmt"
)

type SourceHubspotID struct {
	SourceId string `json:"sourceId"`
}

type SourceHubspot struct {
	Name                    string                  `json:"name"`
	SourceId                string                  `json:"sourceId,omitempty"`
	SourceDefinitionId      string                  `json:"sourceDefinitionId,omitempty"`
	WorkspaceId             string                  `json:"workspaceId,omitempty"`
	ConnectionConfiguration SourceHubspotConnConfig `json:"connectionConfiguration"`
}

type SourceHubspotConnConfig struct {
	StartDate   string                 `json:"start_date"`
	Credentials HubspotCredConfigModel `json:"credentials"`
}

type HubspotCredConfigModel struct {
	CredentialsTitle string `json:"credentials_title"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	AccessToken      string `json:"access_token,omitempty"`
	ClientSecret     string `json:"client_secret,omitempty"`
	ClientId         string `json:"client_id,omitempty"`
}

func (c *Client) CreateHubspotSource(payload SourceHubspot) (SourceHubspot, error) {
	// logger := fwhelpers.GetLogger()

	fmt.Printf("coming here %#v\n", payload)

	method := "POST"
	url := c.Host + "/api/v1/sources/create"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceHubspot{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceHubspot{}, err
	}
	fmt.Printf("statusCode %#v\n", statusCode)
	source := SourceHubspot{}
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

func (c *Client) ReadHubspotSource(sourceId string) (SourceHubspot, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/get"
	sId := SourcePipedriveID{sourceId}
	body, err := json.Marshal(sId)
	if err != nil {
		return SourceHubspot{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceHubspot{}, err
	}

	source := SourceHubspot{}
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

func (c *Client) UpdateHubspotSource(payload SourceHubspot) (SourceHubspot, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/update"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceHubspot{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceHubspot{}, err
	}

	source := SourceHubspot{}
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

func (c *Client) DeleteHubspotSource(sourceId string) error {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/delete"
	sId := SourceHubspotID{sourceId}
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
