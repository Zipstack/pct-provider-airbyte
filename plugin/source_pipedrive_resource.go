package plugin

import (
	"fmt"
	"time"

	"github.com/zipstack/pct-plugin-framework/fwhelpers"
	"github.com/zipstack/pct-plugin-framework/schema"

	"github.com/zipstack/pct-provider-airbyte-local/api"
)

// Resource implementation.
type sourcePipedriveResource struct {
	Client *api.Client
}

type sourcePipedriveResourceModel struct {
	Name                    string                         `pctsdk:"name"`
	SourceId                string                         `pctsdk:"source_id"`
	SourceDefinitionId      string                         `pctsdk:"source_definition_id"`
	WorkspaceId             string                         `pctsdk:"workspace_id"`
	ConnectionConfiguration sourcePipedriveConnConfigModel `pctsdk:"connection_configuration"`
}

type sourcePipedriveConnConfigModel struct {
	ReplicationStartDate string                         `pctsdk:"replication_start_date"`
	Authorization        sourcePipedriveAuthConfigModel `pctsdk:"authorization"`
}

type sourcePipedriveAuthConfigModel struct {
	AuthType string `pctsdk:"auth_type"`
	ApiToken string `pctsdk:"api_token"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ schema.ResourceService = &sourcePipedriveResource{}
)

// Helper function to return a resource service instance.
func NewSourcePipedriveResource() schema.ResourceService {
	return &sourcePipedriveResource{}
}

// Metadata returns the resource type name.
// It is always provider name + "_" + resource type name.
func (r *sourcePipedriveResource) Metadata(req *schema.ServiceRequest) *schema.ServiceResponse {
	return &schema.ServiceResponse{
		TypeName: req.TypeName + "_source_pipedrive",
	}
}

// Configure adds the provider configured client to the resource.
func (r *sourcePipedriveResource) Configure(req *schema.ServiceRequest) *schema.ServiceResponse {
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
func (r *sourcePipedriveResource) Schema() *schema.ServiceResponse {
	s := &schema.Schema{
		Description: "Source pipedrive resource for Airbyte",
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
			"source_definition_id": &schema.StringAttribute{
				Description: "Definition ID",
				Required:    true,
			},
			"workspace_id": &schema.StringAttribute{
				Description: "Workspace ID",
				Required:    true,
			},
			"connection_configuration": &schema.MapAttribute{
				Description: "Connection configuration",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"replication_start_date": &schema.StringAttribute{
						Description: "Replication Start Date",
						Required:    true,
					},
					"authorization": &schema.MapAttribute{
						Description: "authorization",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"auth_type": &schema.StringAttribute{
								Description: "Auth Type",
								Required:    true,
							},
							"api_token": &schema.StringAttribute{
								Description: "API Token",
								Required:    true,
								Sensitive:   true,
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
func (r *sourcePipedriveResource) Create(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	// Retrieve values from plan
	var plan sourcePipedriveResourceModel
	err := fwhelpers.UnpackModel(req.PlanContents, &plan)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Generate API request body from plan
	body := api.SourcePipedrive{}
	body.Name = plan.Name
	body.SourceDefinitionId = plan.SourceDefinitionId
	body.WorkspaceId = plan.WorkspaceId

	body.ConnectionConfiguration = api.SourcePipedriveConnConfig{}
	body.ConnectionConfiguration.ReplicationStartDate = plan.ConnectionConfiguration.ReplicationStartDate
	body.ConnectionConfiguration.Authorization = api.SourcePipedriveAuthConfigModel{}
	body.ConnectionConfiguration.Authorization.ApiToken = plan.ConnectionConfiguration.Authorization.ApiToken
	body.ConnectionConfiguration.Authorization.AuthType = plan.ConnectionConfiguration.Authorization.AuthType

	// Create new source
	source, err := r.Client.CreatePipedriveSource(body)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Update resource state with response body
	state := sourcePipedriveResourceModel{}
	state.Name = source.Name
	state.SourceDefinitionId = source.SourceDefinitionId
	state.SourceId = source.SourceId
	state.WorkspaceId = source.WorkspaceId

	state.ConnectionConfiguration = sourcePipedriveConnConfigModel{}
	state.ConnectionConfiguration.ReplicationStartDate = source.ConnectionConfiguration.ReplicationStartDate
	state.ConnectionConfiguration.Authorization = sourcePipedriveAuthConfigModel{}
	state.ConnectionConfiguration.Authorization.ApiToken = source.ConnectionConfiguration.Authorization.ApiToken
	state.ConnectionConfiguration.Authorization.AuthType = source.ConnectionConfiguration.Authorization.AuthType

	// Set refreshed state
	stateEnc, err := fwhelpers.PackModel(nil, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{
		StateID:          state.SourceId,
		StateContents:    stateEnc,
		StateLastUpdated: time.Now().Format(time.RFC850),
	}
}

// Read resource information
func (r *sourcePipedriveResource) Read(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	var state sourcePipedriveResourceModel

	// Get current state
	err := fwhelpers.UnpackModel(req.StateContents, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	res := schema.ServiceResponse{}

	if req.StateID != "" {
		// Query using existing previous state.
		source, err := r.Client.ReadPipedriveSource(req.StateID)
		if err != nil {
			return schema.ErrorResponse(err)
		}

		// Update state with refreshed value
		state.Name = source.Name
		state.SourceDefinitionId = source.SourceDefinitionId
		state.SourceId = source.SourceId
		state.WorkspaceId = source.WorkspaceId

		state.ConnectionConfiguration = sourcePipedriveConnConfigModel{}
		state.ConnectionConfiguration.ReplicationStartDate = source.ConnectionConfiguration.ReplicationStartDate
		state.ConnectionConfiguration.Authorization = sourcePipedriveAuthConfigModel{}
		state.ConnectionConfiguration.Authorization.AuthType = source.ConnectionConfiguration.Authorization.AuthType
		state.ConnectionConfiguration.Authorization.ApiToken = source.ConnectionConfiguration.Authorization.ApiToken

		res.StateID = state.SourceId
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

func (r *sourcePipedriveResource) Update(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	// Retrieve values from plan
	var plan sourcePipedriveResourceModel
	err := fwhelpers.UnpackModel(req.PlanContents, &plan)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Generate API request body from plan
	body := api.SourcePipedrive{}
	body.Name = plan.Name
	body.SourceId = plan.SourceId

	body.ConnectionConfiguration = api.SourcePipedriveConnConfig{}
	body.ConnectionConfiguration.ReplicationStartDate = plan.ConnectionConfiguration.ReplicationStartDate
	body.ConnectionConfiguration.Authorization = api.SourcePipedriveAuthConfigModel{}
	body.ConnectionConfiguration.Authorization.AuthType = plan.ConnectionConfiguration.Authorization.AuthType
	body.ConnectionConfiguration.Authorization.ApiToken = plan.ConnectionConfiguration.Authorization.ApiToken

	// Update existing source
	_, err = r.Client.UpdatePipedriveSource(body)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Fetch updated items
	source, err := r.Client.ReadPipedriveSource(req.PlanID)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Update state with refreshed value
	state := sourcePipedriveResourceModel{}
	state.Name = source.Name
	state.SourceDefinitionId = source.SourceDefinitionId
	state.SourceId = source.SourceId
	state.WorkspaceId = source.WorkspaceId

	state.ConnectionConfiguration = sourcePipedriveConnConfigModel{}
	state.ConnectionConfiguration.ReplicationStartDate = source.ConnectionConfiguration.ReplicationStartDate
	state.ConnectionConfiguration.Authorization = sourcePipedriveAuthConfigModel{}
	state.ConnectionConfiguration.Authorization.AuthType = source.ConnectionConfiguration.Authorization.AuthType
	state.ConnectionConfiguration.Authorization.ApiToken = source.ConnectionConfiguration.Authorization.ApiToken

	// Set refreshed state
	stateEnc, err := fwhelpers.PackModel(nil, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{
		StateID:          state.SourceId,
		StateContents:    stateEnc,
		StateLastUpdated: time.Now().Format(time.RFC850),
	}
}

// Delete deletes the resource and removes the state on success.
func (r *sourcePipedriveResource) Delete(req *schema.ServiceRequest) *schema.ServiceResponse {
	// Delete existing source
	err := r.Client.DeletePipedriveSource(req.StateID)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{}
}
