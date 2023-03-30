package api

import (
	"encoding/json"
	"fmt"

	"github.com/zipstack/pct-plugin-framework/fwhelpers"
)

type ConnectionResourceID struct {
	ConnecctionID string `json:"connectionId"`
}

type ConnectionResource struct {
	Name          string `json:"name"`
	SourceID      string `json:"sourceId"`
	DestinationID string `json:"destinationId"`
	ConnecctionID string `json:"connectionId,omitempty"`
	// NamespaceDefinition          string           `json:"namespaceDefinition"`
	// NamespaceFormat              string           `json:"namespaceFormat"`
	// Status                       string           `json:"status,omitempty"`
	// Prefix                       string           `json:"prefix,omitempty"`
	ScheduleType string `json:"scheduleType"`
	// NonBreakingChangesPreference string           `json:"nonBreakingChangesPreference"`
	ScheduleData    ConnScheduleData `json:"scheduleData"`
	SourceCatalogId string           `json:"sourceCatalogId"`
	//OperatorConfiguration connOperatorConfig `json:"operator_configuration"`
}
type ConnScheduleData struct {
	BasicSchedule ConnScheduleDataBasicSchedule `json:"basicSchedule,omitempty"`
	Cron          ConnScheduleDataCron          `json:"cron,omitempty"`
}
type ConnScheduleDataBasicSchedule struct {
	TimeUnit string `json:"timeUnit"`
	Units    int64  `json:"units"`
}

type ConnScheduleDataCron struct {
	CronExpression string `json:"cronExpression"`
	CronTimeZone   string `json:"cronTimeZone"`
}

func (c *Client) CreateConnectionResource(payload ConnectionResource) (ConnectionResource, error) {
	logger := fwhelpers.GetLogger()
	logger.Printf("CreateConnectionResource call %#v\n", payload)
	method := "POST"
	url := c.Host + "/api/v1/web_backend/connections/create"
	body, err := json.Marshal(payload)
	if err != nil {
		return ConnectionResource{}, err
	}

	b, statusCode, _, _, err := c.doRequest(method, url, body, nil)
	if err != nil {
		return ConnectionResource{}, err
	}

	connection := ConnectionResource{}

	if statusCode >= 200 && statusCode <= 299 {
		logger.Printf(" CreateConnectionResource  resposne %s", string(b))
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

	fmt.Printf("ConnectionResource %s\n", connectionId)

	method := "POST"
	url := c.Host + "/api/v1/web_backend/connections/get"
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

	method := "POST"
	url := c.Host + "/api/v1/web_backend/connections/update"
	body, err := json.Marshal(payload)
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
			return ConnectionResource{}, err
		} else {
			return connection, fmt.Errorf(msg)
		}
	}
}

func (c *Client) DeleteConnectionResource(connectionId string) error {
	// logger := fwhelpers.GetLogger()

	method := "POST"
	url := c.Host + "/api/v1/web_backend/connections/delete"
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
