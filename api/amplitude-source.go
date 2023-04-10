package api

import (
	"encoding/json"
	"fmt"
)

type SourceAmplitudeID struct {
	SourceId string `json:"sourceId"`
}

type SourceAmplitude struct {
	Name                    string                    `json:"name"`
	SourceId                string                    `json:"sourceId,omitempty"`
	SourceDefinitionId      string                    `json:"sourceDefinitionId,omitempty"`
	WorkspaceId             string                    `json:"workspaceId,omitempty"`
	ConnectionConfiguration SourceAmplitudeConnConfig `json:"connectionConfiguration"`
}

type SourceAmplitudeConnConfig struct {
	StartDate  string `json:"start_date"`
	DataRegion string `json:"data_region,omitempty"`
	SecretKey  string `json:"secret_key"`
	ApiKey     string `json:"api_key"`
}

func (c *Client) CreateAmplitudeSource(payload SourceAmplitude) (SourceAmplitude, error) {
	// logger := fwhelpers.GetLogger()

	fmt.Printf("coming here %#v\n", payload)

	method := "POST"
	url := c.Host + "/api/v1/sources/create"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceAmplitude{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceAmplitude{}, err
	}
	fmt.Printf("statusCode %#v\n", statusCode)
	source := SourceAmplitude{}
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

func (c *Client) ReadAmplitudeSource(sourceId string) (SourceAmplitude, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/get"
	sId := SourcePipedriveID{sourceId}
	body, err := json.Marshal(sId)
	if err != nil {
		return SourceAmplitude{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceAmplitude{}, err
	}

	source := SourceAmplitude{}
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

func (c *Client) UpdateAmplitudeSource(payload SourceAmplitude) (SourceAmplitude, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/update"
	body, err := json.Marshal(payload)
	if err != nil {
		return SourceAmplitude{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return SourceAmplitude{}, err
	}

	source := SourceAmplitude{}
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

func (c *Client) DeleteAmplitudeSource(sourceId string) error {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/sources/delete"
	sId := SourceAmplitudeID{sourceId}
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
