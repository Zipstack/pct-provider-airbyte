package api

import (
	"encoding/json"
	"fmt"
)

type DestinationLocalCSVID struct {
	DestinationId string `json:"destinationId"`
}

type DestinationLocalCSV struct {
	Name                    string                             `json:"name"`
	DestinationId           string                             `json:"destinationId,omitempty"`
	DestinationDefinitionId string                             `json:"destinationDefinitionId,omitempty"`
	WorkspaceId             string                             `json:"workspaceId,omitempty"`
	ConnectionConfiguration DestinationLocalCSVConnConfigModel `json:"connectionConfiguration"`
}

type DestinationLocalCSVConnConfigModel struct {
	DestinationPath string                          `json:"destination_path"`
	DelimiterType   DestinationDelimiterConfigModel `json:"delimiter_type"`
}

type DestinationDelimiterConfigModel struct {
	Delimiter string `json:"delimiter"`
}

func (c *Client) CreateLocalCSVDestination(payload DestinationLocalCSV) (DestinationLocalCSV, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/destinations/create"
	body, err := json.Marshal(payload)
	if err != nil {
		return DestinationLocalCSV{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return DestinationLocalCSV{}, err
	}

	destination := DestinationLocalCSV{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &destination)
		return destination, err
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return destination, err
		} else {
			return destination, fmt.Errorf(msg)
		}
	}
}

func (c *Client) ReadLocalCSVDestination(sourceId string) (DestinationLocalCSV, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/destinations/get"
	sId := DestinationLocalCSVID{sourceId}
	body, err := json.Marshal(sId)
	if err != nil {
		return DestinationLocalCSV{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return DestinationLocalCSV{}, err
	}

	destination := DestinationLocalCSV{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &destination)
		return destination, err
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return destination, err
		} else {
			return destination, fmt.Errorf(msg)
		}
	}
}

func (c *Client) UpdateLocalCSVDestination(payload DestinationLocalCSV) (DestinationLocalCSV, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/destinations/update"
	body, err := json.Marshal(payload)
	if err != nil {
		return DestinationLocalCSV{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return DestinationLocalCSV{}, err
	}

	destination := DestinationLocalCSV{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &destination)
		return destination, err
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return destination, err
		} else {
			return destination, fmt.Errorf(msg)
		}
	}
}

func (c *Client) DeleteLocalCSVDestination(destinationId string) error {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/destinations/delete"
	sId := DestinationLocalCSVID{destinationId}
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
