package plugin

import (
	"fmt"
	"time"

	"github.com/zipstack/pct-plugin-framework/fwhelpers"
	"github.com/zipstack/pct-plugin-framework/schema"

	"github.com/zipstack/pct-provider-airbyte-local/api"
)

type connectionResource struct {
	Client *api.Client
}

type connectionResourceModel struct {
	Name          string           `pctsdk:"name"`
	SourceID      string           `pctsdk:"source_id"`
	DestinationID string           `pctsdk:"destination_id"`
	ConnectionID  string           `pctsdk:"connection_id"`
	Status        string           `pctsdk:"status"`
	ScheduleType  string           `pctsdk:"schedule_type"`
	ScheduleData  connScheduleData `pctsdk:"schedule_data"`
	// OperatorConfiguration connOperatorConfig `pctsdk:"operator_configuration"`
}

type connScheduleData struct {
	BasicSchedule connScheduleDataBasicSchedule `pctsdk:"basic_schedule,omitempty"`
	Cron          connScheduleDataCron          `pctsdk:"cron,omitempty"`
}

type connScheduleDataBasicSchedule struct {
	TimeUnit string `pctsdk:"time_unit"`
	Units    int64  `pctsdk:"units"`
}

type connScheduleDataCron struct {
	CronExpression string `pctsdk:"cron_expression"`
	CronTimeZone   string `pctsdk:"cron_time_zone"`
}

// type connOperatorConfig struct {
// 	OperatorType  string                          `pctsdk:"operator_type"`
// 	Normalization connOperatorConfigNormalization `pctsdk:"normalization"`
// 	Dbt           connOperatorConfigDbt           `pctsdk:"dbt"`
// 	Webhook       connOperatorConfigWebhook       `pctsdk:"webhook"`
// }

// type connOperatorConfigNormalization struct {
// 	Option string `cty:"option"`
// }

// type connOperatorConfigDbt struct {
// 	GitRepoUrl    string `pctsdk:"git_repo_url"`
// 	GitRepoBranch string `pctsdk:"git_repo_branch"`
// 	DockerImage   string `pctsdk:"docker_image"`
// 	DbtArguments  string `pctsdk:"dbt_arguments"`
// }

// type connOperatorConfigWebhook struct {
// 	WebhookConfigId string                            `pctsdk:"webhook_config_id"`
// 	WebhookType     string                            `pctsdk:"webhook_type"`
// 	DbtCloud        connOperatorConfigWebhookDbtCloud `pctsdk:"dbt_cloud"`
// }

// type connOperatorConfigWebhookDbtCloud struct {
// 	AccountId int64 `pctsdk:"account_id"`
// 	JobId     int64 `pctsdk:"job_id"`
// }

// Ensure the implementation satisfies the expected interfaces.
var (
	_ schema.ResourceService = &connectionResource{}
)

// Helper function to return a resource service instance.
func NewConnectionResource() schema.ResourceService {
	return &connectionResource{}
}

// Metadata returns the resource type name.
// It is always provider name + "_" + resource type name.
func (r *connectionResource) Metadata(req *schema.ServiceRequest) *schema.ServiceResponse {
	return &schema.ServiceResponse{
		TypeName: req.TypeName + "_connection",
	}
}

// Configure adds the provider configured client to the resource.
func (r *connectionResource) Configure(req *schema.ServiceRequest) *schema.ServiceResponse {
	if req.ResourceData == "" {
		return schema.ErrorResponse(fmt.Errorf("no data provided to configure resource"))
	}

	var creds map[string]string
	err := fwhelpers.Decode(req.ResourceData, &creds)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	client, err := api.NewClient(
		creds["host"], creds["username"], creds["password"],
	)
	if err != nil {
		return schema.ErrorResponse(fmt.Errorf("malformed data provided to configure resource"))
	}

	r.Client = client

	return &schema.ServiceResponse{}
}

// Schema defines the schema for the resource.
func (r *connectionResource) Schema() *schema.ServiceResponse {
	s := &schema.Schema{
		Description: "Connection resource for Airbyte",
		Attributes: map[string]schema.Attribute{
			"name": &schema.StringAttribute{
				Description: "Name",
				Required:    true,
			},
			"source_id": &schema.StringAttribute{
				Description: "Source ID",
				Required:    true,
			},
			"destination_id": &schema.StringAttribute{
				Description: "Destination ID",
				Required:    true,
			},
			"connection_id": &schema.StringAttribute{
				Description: "Connection ID",
				Computed:    true,
			},
			"status": &schema.StringAttribute{
				Description: "Status",
				Required:    true,
			},
			"schedule_type": &schema.StringAttribute{
				Description: "Schedule type",
				Required:    true,
			},
			"schedule_data": &schema.MapAttribute{
				Description:  "Schedule data",
				Required:     true,
				ExactlyOneOf: []string{"basic_schedule", "cron"},
				Attributes: map[string]schema.Attribute{
					"basic_schedule": &schema.MapAttribute{
						Description: "Basic schedule",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"time_unit": &schema.StringAttribute{
								Description: "Time unit",
								Required:    true,
							},
							"units": &schema.IntAttribute{
								Description: "Units",
								Required:    true,
							},
						},
					},
					"cron": &schema.MapAttribute{
						Description: "Cron",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"cron_time_zone": &schema.StringAttribute{
								Description: "Cron time zone",
								Required:    true,
							},
							"cron_expression": &schema.StringAttribute{
								Description: "Cron expression",
								Required:    true,
							},
						},
					},
				},
			},
			// "operator_configuration": &schema.MapAttribute{
			// 	Description: "Operator configuration",
			// 	Required:    false,
			// 	Attributes: map[string]schema.Attribute{
			// 		"operator_type": &schema.StringAttribute{
			// 			Description: "Operator type",
			// 			Required:    false,
			// 		},
			// 		"normalization": &schema.MapAttribute{
			// 			Description: "Basic schedule",
			// 			Required:    false,
			// 			Attributes: map[string]schema.Attribute{
			// 				"option": &schema.StringAttribute{
			// 					Description: "Option",
			// 					Required:    false,
			// 				},
			// 			},
			// 		},
			// 		"dbt": &schema.MapAttribute{
			// 			Description: "DBT",
			// 			Required:    false,
			// 			Attributes: map[string]schema.Attribute{
			// 				"dbt_arguments": &schema.StringAttribute{
			// 					Description: "DBT arguments",
			// 					Required:    false,
			// 				},
			// 				"docker_image": &schema.StringAttribute{
			// 					Description: "Docker image",
			// 					Required:    false,
			// 				},
			// 				"git_repo_branch": &schema.StringAttribute{
			// 					Description: "Git repo branch",
			// 					Required:    false,
			// 				},
			// 				"git_repo_url": &schema.StringAttribute{
			// 					Description: "Git repo url",
			// 					Required:    false,
			// 				},
			// 			},
			// 		},
			// 		"webhook": &schema.MapAttribute{
			// 			Description: "Webhook",
			// 			Required:    false,
			// 			Attributes: map[string]schema.Attribute{
			// 				"webhook_config_id": &schema.StringAttribute{
			// 					Description: "Webhook config ID",
			// 					Required:    false,
			// 				},
			// 				"webhook_type": &schema.StringAttribute{
			// 					Description: "Webhook type",
			// 					Required:    false,
			// 				},
			// 				"dbt_cloud": &schema.MapAttribute{
			// 					Description: "DBT cloud",
			// 					Required:    false,
			// 					Attributes: map[string]schema.Attribute{
			// 						"account_id": &schema.IntAttribute{
			// 							Description: "Account ID",
			// 							Required:    false,
			// 						},
			// 						"job_id": &schema.IntAttribute{
			// 							Description: "Job ID",
			// 							Required:    false,
			// 						},
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
		},
	}

	sEnc, err := fwhelpers.Encode(s)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{
		SchemaContents: sEnc,
	}
}

// Create a new resource
func (r *connectionResource) Create(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	var plan connectionResourceModel
	err := fwhelpers.UnpackModel(req.PlanContents, &plan)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	body := api.ConnectionResource{}
	body.Name = plan.Name
	body.SourceID = plan.SourceID
	body.DestinationID = plan.DestinationID

	body.ScheduleType = plan.ScheduleType
	body.ScheduleData = api.ConnScheduleData{}

	if plan.ScheduleType == "basic" {
		body.ScheduleData.BasicSchedule = api.ConnScheduleDataBasicSchedule{}
		body.ScheduleData.BasicSchedule.TimeUnit = plan.ScheduleData.BasicSchedule.TimeUnit
		body.ScheduleData.BasicSchedule.Units = plan.ScheduleData.BasicSchedule.Units
	} else if plan.ScheduleType == "cron" {
		body.ScheduleData.Cron = api.ConnScheduleDataCron{}
		body.ScheduleData.Cron.CronExpression = plan.ScheduleData.Cron.CronExpression
		body.ScheduleData.Cron.CronTimeZone = plan.ScheduleData.Cron.CronTimeZone
	}

	body.Status = plan.Status

	connection, err := r.Client.CreateConnectionResource(body)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Map response body to schema and populate Computed attribute values
	state := connectionResourceModel{}

	state.Name = connection.Name
	state.ConnectionID = connection.ConnectionID
	state.SourceID = connection.SourceID
	state.DestinationID = connection.DestinationID

	state.ScheduleType = connection.ScheduleType
	state.ScheduleData = connScheduleData{}

	if connection.ScheduleType == "basic" {
		state.ScheduleData.BasicSchedule = connScheduleDataBasicSchedule{}
		state.ScheduleData.BasicSchedule.TimeUnit = connection.ScheduleData.BasicSchedule.TimeUnit
		state.ScheduleData.BasicSchedule.Units = connection.ScheduleData.BasicSchedule.Units
	} else if connection.ScheduleType == "cron" {
		state.ScheduleData.Cron = connScheduleDataCron{}
		state.ScheduleData.Cron.CronExpression = connection.ScheduleData.Cron.CronExpression
		state.ScheduleData.Cron.CronTimeZone = connection.ScheduleData.Cron.CronTimeZone
	}

	state.Status = connection.Status

	stateEnc, err := fwhelpers.PackModel(nil, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{
		StateID:          state.ConnectionID,
		StateContents:    stateEnc,
		StateLastUpdated: time.Now().Format(time.RFC850),
	}
}

// Read resource information
func (r *connectionResource) Read(req *schema.ServiceRequest) *schema.ServiceResponse {
	var state connectionResourceModel

	// Get current state
	err := fwhelpers.UnpackModel(req.StateContents, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	res := schema.ServiceResponse{}

	if req.StateID != "" {
		// Query using existing previous state.
		connection, err := r.Client.ReadConnectionResource(req.StateID)
		if err != nil {
			return schema.ErrorResponse(err)
		}

		state = connectionResourceModel{}

		// Update state with refreshed value
		state.Name = connection.Name
		state.ConnectionID = connection.ConnectionID
		state.SourceID = connection.SourceID
		state.DestinationID = connection.DestinationID

		state.ScheduleType = connection.ScheduleType
		state.ScheduleData = connScheduleData{}

		if connection.ScheduleType == "basic" {
			state.ScheduleData.BasicSchedule = connScheduleDataBasicSchedule{}
			state.ScheduleData.BasicSchedule.TimeUnit = connection.ScheduleData.BasicSchedule.TimeUnit
			state.ScheduleData.BasicSchedule.Units = connection.ScheduleData.BasicSchedule.Units
		} else if connection.ScheduleType == "cron" {
			state.ScheduleData.Cron = connScheduleDataCron{}
			state.ScheduleData.Cron.CronExpression = connection.ScheduleData.Cron.CronExpression
			state.ScheduleData.Cron.CronTimeZone = connection.ScheduleData.Cron.CronTimeZone
		}

		state.Status = connection.Status

		res.StateID = connection.ConnectionID
	} else {
		// No previous state exists.
		res.StateID = ""
		res.StateLastUpdated = ""
	}

	// Set refreshed state
	stateEnc, err := fwhelpers.PackModel(nil, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}
	res.StateContents = stateEnc

	return &res
}

func (r *connectionResource) Update(req *schema.ServiceRequest) *schema.ServiceResponse {
	var plan connectionResourceModel
	err := fwhelpers.UnpackModel(req.PlanContents, &plan)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Generate API request body from plan
	body := api.ConnectionResource{}

	body.Name = plan.Name
	body.ConnectionID = plan.ConnectionID
	body.SourceID = plan.SourceID
	body.DestinationID = plan.DestinationID

	body.ScheduleType = plan.ScheduleType
	body.ScheduleData = api.ConnScheduleData{}

	if plan.ScheduleType == "basic" {
		body.ScheduleData.BasicSchedule = api.ConnScheduleDataBasicSchedule{}
		body.ScheduleData.BasicSchedule.TimeUnit = plan.ScheduleData.BasicSchedule.TimeUnit
		body.ScheduleData.BasicSchedule.Units = plan.ScheduleData.BasicSchedule.Units
	} else if plan.ScheduleType == "cron" {
		body.ScheduleData.Cron = api.ConnScheduleDataCron{}
		body.ScheduleData.Cron.CronExpression = plan.ScheduleData.Cron.CronExpression
		body.ScheduleData.Cron.CronTimeZone = plan.ScheduleData.Cron.CronTimeZone
	}

	body.Status = plan.Status

	// Update existing source
	_, err = r.Client.UpdateConnectionResource(body)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Fetch updated items
	connection, err := r.Client.ReadConnectionResource(req.PlanID)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Update state with refreshed value
	state := connectionResourceModel{}

	state.Name = connection.Name
	state.ConnectionID = connection.ConnectionID
	state.SourceID = connection.SourceID
	state.DestinationID = connection.DestinationID

	state.ScheduleType = connection.ScheduleType
	state.ScheduleData = connScheduleData{}

	if connection.ScheduleType == "basic" {
		state.ScheduleData.BasicSchedule = connScheduleDataBasicSchedule{}
		state.ScheduleData.BasicSchedule.TimeUnit = connection.ScheduleData.BasicSchedule.TimeUnit
		state.ScheduleData.BasicSchedule.Units = connection.ScheduleData.BasicSchedule.Units
	} else if connection.ScheduleType == "cron" {
		state.ScheduleData.Cron = connScheduleDataCron{}
		state.ScheduleData.Cron.CronExpression = connection.ScheduleData.Cron.CronExpression
		state.ScheduleData.Cron.CronTimeZone = connection.ScheduleData.Cron.CronTimeZone
	}

	state.Status = connection.Status

	// Set refreshed state
	stateEnc, err := fwhelpers.PackModel(nil, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{
		StateID:          state.ConnectionID,
		StateContents:    stateEnc,
		StateLastUpdated: time.Now().Format(time.RFC850),
	}
}

// Delete deletes the resource and removes the state on success.
func (r *connectionResource) Delete(req *schema.ServiceRequest) *schema.ServiceResponse {
	// Delete existing source
	err := r.Client.DeleteConnectionResource(req.StateID)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{}
}
