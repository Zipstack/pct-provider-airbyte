package api

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type SourceFakerID struct {
// 	SourceId string `json:"sourceId"`
// }

// type SourceFaker struct {
// 	Name                    string                `json:"name"`
// 	SourceDefinitionId      string                `json:"sourceDefinitionId,omitempty"`
// 	SourceId                string                `json:"sourceId,omitempty"`
// 	WorkspaceId             string                `json:"workspaceId,omitempty"`
// 	ConnectionConfiguration SourceFakerConnConfig `json:"connectionConfiguration"`
// }

// type SourceFakerConnConfig struct {
// 	Seed            int64 `json:"seed"`
// 	Count           int64 `json:"count"`
// 	RecordsPerSync  int64 `json:"records_per_sync"`
// 	RecordsPerSlice int64 `json:"records_per_slice"`
// }

// func (c *Client) CreateSource(payload SourceFaker) (SourceFaker, error) {
// 	// logger := fwhelpers.GetLogger()

// 	method := "POST"
// 	url := c.Host + "/api/v1/sources/create"
// 	body, err := json.Marshal(payload)
// 	if err != nil {
// 		return SourceFaker{}, err
// 	}

// 	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
// 	if err != nil {
// 		return SourceFaker{}, err
// 	}

// 	source := SourceFaker{}
// 	if statusCode >= 200 && statusCode <= 299 {
// 		err = json.Unmarshal(b, &source)
// 		return source, err
// 	} else {
// 		msg, err := c.getAPIError(b)
// 		if err != nil {
// 			return source, err
// 		} else {
// 			return source, fmt.Errorf(msg)
// 		}
// 	}
// }

// func (c *Client) ReadSource(sourceId string) (SourceFaker, error) {
// 	// logger := fwhelpers.GetLogger()

// 	method := "POST"
// 	url := c.Host + "/api/v1/sources/get"
// 	sId := SourceFakerID{sourceId}
// 	body, err := json.Marshal(sId)
// 	if err != nil {
// 		return SourceFaker{}, err
// 	}

// 	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
// 	if err != nil {
// 		return SourceFaker{}, err
// 	}

// 	source := SourceFaker{}
// 	if statusCode >= 200 && statusCode <= 299 {
// 		err = json.Unmarshal(b, &source)
// 		return source, err
// 	} else {
// 		msg, err := c.getAPIError(b)
// 		if err != nil {
// 			return source, err
// 		} else {
// 			return source, fmt.Errorf(msg)
// 		}
// 	}
// }

// func (c *Client) UpdateSource(payload SourceFaker) (SourceFaker, error) {
// 	// logger := fwhelpers.GetLogger()

// 	method := "POST"
// 	url := c.Host + "/api/v1/sources/update"
// 	body, err := json.Marshal(payload)
// 	if err != nil {
// 		return SourceFaker{}, err
// 	}

// 	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
// 	if err != nil {
// 		return SourceFaker{}, err
// 	}

// 	source := SourceFaker{}
// 	if statusCode >= 200 && statusCode <= 299 {
// 		err = json.Unmarshal(b, &source)
// 		return source, err
// 	} else {
// 		msg, err := c.getAPIError(b)
// 		if err != nil {
// 			return source, err
// 		} else {
// 			return source, fmt.Errorf(msg)
// 		}
// 	}
// }

// func (c *Client) DeleteSource(sourceId string) error {
// 	// logger := fwhelpers.GetLogger()

// 	method := "POST"
// 	url := c.Host + "/api/v1/sources/delete"
// 	sId := SourceFakerID{sourceId}
// 	body, err := json.Marshal(sId)
// 	if err != nil {
// 		return err
// 	}

// 	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
// 	if err != nil {
// 		return err
// 	}

// 	if statusCode >= 200 && statusCode <= 299 {
// 		return nil
// 	} else {
// 		msg, err := c.getAPIError(b)
// 		if err != nil {
// 			return err
// 		} else {
// 			return fmt.Errorf(msg)
// 		}
// 	}
// }
