package plugin

import (
	"fmt"

	"github.com/zipstack/pct-plugin-framework/fwhelpers"
	"github.com/zipstack/pct-plugin-framework/schema"

	"github.com/zipstack/pct-provider-airbyte/api"
)

// Resource implementation.
type connectionResource struct {
	Client *api.Client
}

type connectionResourceModel struct {
	Name                  string             `json:"name"`
	SourceID              string             `json:"sourceId"`
	DestinationID         string             `json:"destinationId"`
	ConnectionID          string             `json:"connectionId"`
	NamespaceDefinition   string             `json:"namespace_definition"`
	NamespaceFormat       string             `json:"namespace_format"`
	Status                string             `json:"status"`
	Prefix                string             `json:"prefix"`
	ScheduleType          string             `json:"schedule_type"`
	ScheduleData          connScheduleData   `json:"schedule_data"`
	OperatorConfiguration connOperatorConfig `json:"operator_configuration"`
}

type connScheduleData struct {
	BasicSchedule connScheduleDataBasicSchedule `json:"basic_schedule"`
	Cron          connScheduleDataCron          `json:"cron"`
}

type connScheduleDataBasicSchedule struct {
	TimeUnit string `json:"time_unit"`
	Units    int64  `json:"units"`
}

type connScheduleDataCron struct {
	CronExpression string `json:"cron_expression"`
	CronTimeZone   string `json:"cron_time_zone"`
}

type connOperatorConfig struct {
	OperatorType  string                          `json:"operator_type"`
	Normalization connOperatorConfigNormalization `json:"normalization"`
	Dbt           connOperatorConfigDbt           `json:"dbt"`
	Webhook       connOperatorConfigWebhook       `json:"webhook"`
}

type connOperatorConfigNormalization struct {
	Option string `json:"option"`
}

type connOperatorConfigDbt struct {
	GitRepoUrl    string `json:"git_repo_url"`
	GitRepoBranch string `json:"git_repo_branch"`
	DockerImage   string `json:"docker_image"`
	DbtArguments  string `json:"dbt_arguments"`
}

type connOperatorConfigWebhook struct {
	WebhookConfigId string                            `json:"webhook_config_id"`
	WebhookType     string                            `json:"webhook_type"`
	DbtCloud        connOperatorConfigWebhookDbtCloud `json:"dbt_cloud"`
}

type connOperatorConfigWebhookDbtCloud struct {
	AccountId int64 `json:"account_id"`
	JobId     int64 `json:"job_id"`
}

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

	var client *api.Client
	err := fwhelpers.Decode(req.ResourceData, &client)
	if err != nil {
		return schema.ErrorResponse(err)
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
				Required:    false,
				Computed:    true,
			},
			"destination_id": &schema.StringAttribute{
				Description: "Destination ID",
				Required:    false,
				Computed:    true,
			},
			"connection_id": &schema.StringAttribute{
				Description: "Connection ID",
				Required:    false,
				Computed:    true,
			},
			"namespace_definition": &schema.StringAttribute{
				Description: "Namespace definition",
				Required:    false,
			},
			"namespace_format": &schema.StringAttribute{
				Description: "Namespace format",
				Required:    false,
			},
			"status": &schema.StringAttribute{
				Description: "Status",
				Required:    true,
			},
			"prefix": &schema.StringAttribute{
				Description: "Prefix",
				Required:    false,
			},
			"schedule_type": &schema.StringAttribute{
				Description: "Schedule type",
				Required:    true,
			},
			"schedule_data": &schema.MapAttribute{
				Description: "Schedule data",
				Required:    true,
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
			"operator_configuration": &schema.MapAttribute{
				Description: "Operator configuration",
				Required:    false,
				Attributes: map[string]schema.Attribute{
					"operator_type": &schema.StringAttribute{
						Description: "Operator type",
						Required:    false,
					},
					"normalization": &schema.MapAttribute{
						Description: "Basic schedule",
						Required:    false,
						Attributes: map[string]schema.Attribute{
							"option": &schema.StringAttribute{
								Description: "Option",
								Required:    false,
							},
						},
					},
					"dbt": &schema.MapAttribute{
						Description: "DBT",
						Required:    false,
						Attributes: map[string]schema.Attribute{
							"dbt_arguments": &schema.StringAttribute{
								Description: "DBT arguments",
								Required:    false,
							},
							"docker_image": &schema.StringAttribute{
								Description: "Docker image",
								Required:    false,
							},
							"git_repo_branch": &schema.StringAttribute{
								Description: "Git repo branch",
								Required:    false,
							},
							"git_repo_url": &schema.StringAttribute{
								Description: "Git repo url",
								Required:    false,
							},
						},
					},
					"webhook": &schema.MapAttribute{
						Description: "Webhook",
						Required:    false,
						Attributes: map[string]schema.Attribute{
							"webhook_config_id": &schema.StringAttribute{
								Description: "Webhook config ID",
								Required:    false,
							},
							"webhook_type": &schema.StringAttribute{
								Description: "Webhook type",
								Required:    false,
							},
							"dbt_cloud": &schema.MapAttribute{
								Description: "DBT cloud",
								Required:    false,
								Attributes: map[string]schema.Attribute{
									"account_id": &schema.IntAttribute{
										Description: "Account ID",
										Required:    false,
									},
									"job_id": &schema.IntAttribute{
										Description: "Job ID",
										Required:    false,
									},
								},
							},
						},
					},
				},
			},
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
	return &schema.ServiceResponse{}

	// logger := fwhelpers.GetLogger()

	// var plan sourceResourceModel
	// err := fwhelpers.UnpackModel(req.PlanContents, &plan)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// // Generate API request body from plan
	// // plan??

	// // Create new order
	// source, err := r.Client.CreateSource(plan)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// // Map response body to schema and populate Computed attribute values
	// state := make(map[string]string)
	// // plan.ID = types.StringValue(strconv.Itoa(order.ID))
	// // plan.Items[int]model
	// // plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// stateEnc, err := fwhelpers.Encode(state)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// return &schema.ServiceResponse{
	// 	StateContents: stateEnc,
	// }
}

// Read resource information
func (r *connectionResource) Read(req *schema.ServiceRequest) *schema.ServiceResponse {
	return &schema.ServiceResponse{}

	// // Get current state
	// var state sourceResourceModel
	// err := fwhelpers.UnpackModel(req.StateContents, &state)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// // Get refreshed source value
	// source, err := r.Client.GetSource(state.ID)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// // Overwrite items with refreshed state
	// state.Items = []sourceResourceModel{}
	// // fill state.Items with source

	// // Set refreshed state
	// stateEnc, err := fwhelpers.Encode(state)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// return &schema.ServiceResponse{
	// 	StateContents: stateEnc,
	// }
}

func (r *connectionResource) Update(req *schema.ServiceRequest) *schema.ServiceResponse {
	return &schema.ServiceResponse{}

	// // Retrieve values from plan
	// var plan sourceResourceModel
	// err := fwhelpers.UnpackModel(req.StateContents, &plan)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// // Generate API request body from plan
	// // plan??

	// // Update existing source
	// source, err := r.Client.UpdateSource(plan)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// // Fetch updated items from GetSource as UpdateSource items are not
	// // populated.
	// source, err := r.Client.GetSource(state.ID)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// // Update resource state with updated items and timestamp
	// state.Items = []sourceResourceModel{}
	// // fill state.Items with source
	// // plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// // Set refreshed state
	// stateEnc, err := fwhelpers.Encode(state)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// return &schema.ServiceResponse{
	// 	StateContents: stateEnc,
	// }
}

// Delete deletes the resource and removes the state on success.
func (r *connectionResource) Delete(req *schema.ServiceRequest) *schema.ServiceResponse {
	return &schema.ServiceResponse{}

	// // Retrieve values from state
	// var state sourceResourceModel
	// err := fwhelpers.UnpackModel(req.StateContents, &state)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// // Delete existing source
	// source, err := r.Client.DeleteSource(state.ID)
	// if err != nil {
	// 	return schema.ErrorResponse(err)
	// }

	// return &schema.ServiceResponse{}
}
