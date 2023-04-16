package api

import (
	"encoding/json"
	"fmt"
)

type ConnectionResourceID struct {
	ConnecctionID string `json:"connectionId"`
}

type ConnectionResource struct {
	Name          string           `json:"name"`
	SourceID      string           `json:"sourceId,omitempty"`
	DestinationID string           `json:"destinationId,omitempty"`
	ConnectionID  string           `json:"connectionId,omitempty"`
	SyncCatalog   interface{}      `json:"syncCatalog"`
	Status        string           `json:"status"`
	ScheduleType  string           `json:"scheduleType"`
	ScheduleData  ConnScheduleData `json:"scheduleData"`
	//OperatorConfiguration connOperatorConfig `json:"operator_configuration"`
}
type ConnScheduleData struct {
	BasicSchedule ConnScheduleDataBasicSchedule `json:"basicSchedule,omitempty"`
	Cron          ConnScheduleDataCron          `json:"cron,omitempty"`
}

type ConnScheduleDataBasicSchedule struct {
	TimeUnit string `json:"timeUnit,omitempty"`
	Units    int64  `json:"units,omitempty"`
}

type ConnScheduleDataCron struct {
	CronExpression string `json:"cronExpression,omitempty"`
	CronTimeZone   string `json:"cronTimeZone,omitempty"`
}

type DiscoverSourceSchemaCatalog struct {
	SourceID     string `json:"sourceId"`
	DisableCache bool   `json:"disable_cache"`
}

func (c *Client) CreateConnectionResource(payload ConnectionResource) (ConnectionResource, error) {
	// logger := fwhelpers.GetLogger()

	discover_payload := DiscoverSourceSchemaCatalog{
		SourceID:     payload.SourceID,
		DisableCache: true,
	}
	catalog := map[string]interface{}{}

	method := "POST"
	url := c.Host + "/api/v1/sources/discover_schema"
	body, err := json.Marshal(discover_payload)
	if err != nil {
		return ConnectionResource{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return ConnectionResource{}, err
	}

	err = json.Unmarshal(b, &catalog)
	if err != nil {
		return ConnectionResource{}, err
	}
	jobInfo := catalog["jobInfo"].(map[string]interface{})
	if !jobInfo["succeeded"].(bool) {
		return ConnectionResource{}, fmt.Errorf("failed to get source schema catalog")
	}
	if statusCode < 200 || statusCode > 299 {
		msg, err := c.getAPIError(b)
		if err != nil {
			return ConnectionResource{}, err
		} else {
			return ConnectionResource{}, fmt.Errorf(msg)
		}
	}

	payload.SyncCatalog = catalog["catalog"]

	method = "POST"
	url = c.Host + "/api/v1/connections/create"
	body, err = json.Marshal(payload)
	if err != nil {
		return ConnectionResource{}, err
	}

	b, statusCode, _, _, err = c.doRequest(method, url, body, nil)
	if err != nil {
		return ConnectionResource{}, err
	}

	connection := ConnectionResource{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &connection)
		return connection, err
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return connection, err
		} else {
			return connection, fmt.Errorf(msg)
		}
	}
}

func (c *Client) ReadConnectionResource(connectionId string) (ConnectionResource, error) {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/connections/get"
	sId := ConnectionResourceID{connectionId}
	body, err := json.Marshal(sId)
	if err != nil {
		return ConnectionResource{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return ConnectionResource{}, err
	}

	connection := ConnectionResource{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &connection)
		return connection, err
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return connection, err
		} else {
			return connection, fmt.Errorf(msg)
		}
	}
}

func (c *Client) UpdateConnectionResource(payload ConnectionResource) (ConnectionResource, error) {
	// logger := fwhelpers.GetLogger()

	discover_payload := DiscoverSourceSchemaCatalog{
		SourceID:     payload.SourceID,
		DisableCache: true,
	}
	catalog := map[string]interface{}{}

	method := "POST"
	url := c.Host + "/api/v1/sources/discover_schema"
	body, err := json.Marshal(discover_payload)
	if err != nil {
		return ConnectionResource{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return ConnectionResource{}, err
	}

	err = json.Unmarshal(b, &catalog)
	if err != nil {
		return ConnectionResource{}, err
	}
	jobInfo := catalog["jobInfo"].(map[string]interface{})
	if !jobInfo["succeeded"].(bool) {
		return ConnectionResource{}, fmt.Errorf("failed to get source schema catalog")
	}
	if statusCode < 200 || statusCode > 299 {
		msg, err := c.getAPIError(b)
		if err != nil {
			return ConnectionResource{}, err
		} else {
			return ConnectionResource{}, fmt.Errorf(msg)
		}
	}

	// payload.SyncCatalog = catalog["catalog"]

	method = "POST"
	url = c.Host + "/api/v1/connections/update"
	body, err = json.Marshal(payload)
	if err != nil {
		return ConnectionResource{}, err
	}

	b, statusCode, _, _, err = c.doRequest(method, url, body, nil)
	if err != nil {
		return ConnectionResource{}, err
	}

	connection := ConnectionResource{}
	if statusCode >= 200 && statusCode <= 299 {
		err = json.Unmarshal(b, &connection)
		return connection, err
	} else {
		msg, err := c.getAPIError(b)
		if err != nil {
			return ConnectionResource{}, err
		} else {
			return connection, fmt.Errorf(msg)
		}
	}
}

func (c *Client) DeleteConnectionResource(connectionId string) error {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/connections/delete"
	sId := ConnectionResourceID{connectionId}
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
