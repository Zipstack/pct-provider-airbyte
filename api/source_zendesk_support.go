package api

import (
	"encoding/json"
	"fmt"
)

type SourceZendeskSupportID struct {
	SourceId string `json:"sourceId"`
}

type SourceZendeskSupport struct {
	Name                    string                         `json:"name"`
	SourceId                string                         `json:"sourceId,omitempty"`
	SourceDefinitionId      string                         `json:"sourceDefinitionId,omitempty"`
	WorkspaceId             string                         `json:"workspaceId,omitempty"`
	ConnectionConfiguration SourceZendeskSupportConnConfig `json:"connectionConfiguration"`
}

type SourceZendeskSupportConnConfig struct {
	StartDate        string                              `json:"start_date"`
	Subdomain        string                              `json:"subdomain,omitempty"`
	IgnorePagination bool                                `json:"ignore_pagination"`
	Credentials      SourceZendeskSupportCredConfigModel `json:"credentials"`
}

type SourceZendeskSupportCredConfigModel struct {
	Credentials string `json:"credentials"`
	ApiToken    string `json:"api_token,omitempty"`
	Email       string `json:"email,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

func (c *Client) CreateZendeskSupportSource(payload SourceZendeskSupport) (SourceZendeskSupport, error) {
	// logger := fwhelpers.GetLogger()
	method := "POST"
	url := c.Host + "/api/v1/sources/create"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceZendeskSupport{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceZendeskSupport{}, err
	}
	source := SourceZendeskSupport{}
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

func (c *Client) ReadZendeskSupportSource(sourceId string) (SourceZendeskSupport, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/get"
	sId := SourceZendeskSupportID{sourceId}
	body, err := json.Marshal(sId)
	if err != nil {
		return SourceZendeskSupport{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceZendeskSupport{}, err
	}

	source := SourceZendeskSupport{}
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

func (c *Client) UpdateZendeskSupportSource(payload SourceZendeskSupport) (SourceZendeskSupport, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/update"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceZendeskSupport{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceZendeskSupport{}, err
	}

	source := SourceZendeskSupport{}
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

func (c *Client) DeleteZendeskSupportSource(sourceId string) error {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/delete"
	sId := SourceZendeskSupportID{sourceId}
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
