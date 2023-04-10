package plugin

import (
	"fmt"
	"time"

	"github.com/zipstack/pct-plugin-framework/fwhelpers"
	"github.com/zipstack/pct-plugin-framework/schema"

	"github.com/zipstack/pct-provider-airbyte/api"
)

// Resource implementation.
type sourceZendeskSupportResource struct {
	Client *api.Client
}

type sourceZendeskSupportResourceModel struct {
	Name                    string                              `cty:"name"`
	SourceId                string                              `cty:"source_id"`
	SourceDefinitionId      string                              `cty:"source_definition_id"`
	WorkspaceId             string                              `cty:"workspace_id"`
	ConnectionConfiguration sourceZendeskSupportConnConfigModel `cty:"connection_configuration"`
}

type sourceZendeskSupportConnConfigModel struct {
	StartDate        string                        `cty:"start_date"`
	Subdomain        string                        `cty:"subdomain"`
	IgnorePagination bool                          `cty:"ignore_pagination"`
	Credentials      ZendeskSupportCredConfigModel `cty:"credentials"`
}

type ZendeskSupportCredConfigModel struct {
	Credentials string `cty:"credentials"`
	ApiToken    string `cty:"api_token"`
	Email       string `cty:"email"`
	AccessToken string `cty:"access_token"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ schema.ResourceService = &sourceZendeskSupportResource{}
)

// Helper function to return a resource service instance.
func NewsourceZendeskSupportResource() schema.ResourceService {
	return &sourceZendeskSupportResource{}
}

// Metadata returns the resource type name.
// It is always provider name + "_" + resource type name.
func (r *sourceZendeskSupportResource) Metadata(req *schema.ServiceRequest) *schema.ServiceResponse {
	return &schema.ServiceResponse{
		TypeName: req.TypeName + "_source_zendesk_support",
	}
}

// Configure adds the provider configured client to the resource.
func (r *sourceZendeskSupportResource) Configure(req *schema.ServiceRequest) *schema.ServiceResponse {
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
func (r *sourceZendeskSupportResource) Schema() *schema.ServiceResponse {
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
				//Sensitive:   true,
				Attributes: map[string]schema.Attribute{
					"start_date": &schema.StringAttribute{
						Description: "Start Date",
						Required:    true,
					},
					"subdomain": &schema.StringAttribute{
						Description: "subdomain",
						Required:    true,
					},
					"ignore_pagination": &schema.BoolAttribute{
						Description: "ignore_pagination",
						Required:    false,
					},
					"credentials": &schema.MapAttribute{
						Description: "credentials",
						Required:    true,

						Attributes: map[string]schema.Attribute{
							"credentials": &schema.StringAttribute{
								Description: "credentials",
								Required:    true,
							},
							"api_token": &schema.StringAttribute{
								Description: "API Token",
								Required:    false,
								Sensitive:   true,
							},
							"email": &schema.StringAttribute{
								Description: "Email",
								Required:    false,
							},
							"access_token": &schema.StringAttribute{
								Description: "Access Token",
								Required:    false,
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
func (r *sourceZendeskSupportResource) Create(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	// Retrieve values from plan
	var plan sourceZendeskSupportResourceModel
	err := fwhelpers.UnpackModel(req.PlanContents, &plan)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Generate API request body from plan
	body := api.SourceZendeskSupport{}
	body.Name = plan.Name
	body.SourceDefinitionId = plan.SourceDefinitionId
	body.WorkspaceId = plan.WorkspaceId

	body.ConnectionConfiguration = api.SourceZendeskSupportConnConfig{}
	body.ConnectionConfiguration.StartDate = plan.ConnectionConfiguration.StartDate
	body.ConnectionConfiguration.IgnorePagination = plan.ConnectionConfiguration.IgnorePagination
	body.ConnectionConfiguration.Subdomain = plan.ConnectionConfiguration.Subdomain
	body.ConnectionConfiguration.Credentials = api.ZendeskSupportCredConfigModel{}
	body.ConnectionConfiguration.Credentials.Credentials = plan.ConnectionConfiguration.Credentials.Credentials
	body.ConnectionConfiguration.Credentials.ApiToken = plan.ConnectionConfiguration.Credentials.ApiToken
	body.ConnectionConfiguration.Credentials.Email = plan.ConnectionConfiguration.Credentials.Email
	body.ConnectionConfiguration.Credentials.AccessToken = plan.ConnectionConfiguration.Credentials.AccessToken
	// Create new source
	source, err := r.Client.CreateZendeskSupportSource(body)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Update resource state with response body
	state := sourceZendeskSupportResourceModel{}
	state.Name = source.Name
	state.SourceDefinitionId = source.SourceDefinitionId
	state.SourceId = source.SourceId
	state.WorkspaceId = source.WorkspaceId

	state.ConnectionConfiguration = sourceZendeskSupportConnConfigModel{}
	state.ConnectionConfiguration.StartDate = source.ConnectionConfiguration.StartDate
	state.ConnectionConfiguration.IgnorePagination = source.ConnectionConfiguration.IgnorePagination
	state.ConnectionConfiguration.Subdomain = source.ConnectionConfiguration.Subdomain
	state.ConnectionConfiguration.Credentials = ZendeskSupportCredConfigModel{}
	state.ConnectionConfiguration.Credentials.Credentials = source.ConnectionConfiguration.Credentials.Credentials
	state.ConnectionConfiguration.Credentials.ApiToken = source.ConnectionConfiguration.Credentials.ApiToken
	state.ConnectionConfiguration.Credentials.Email = source.ConnectionConfiguration.Credentials.Email
	state.ConnectionConfiguration.Credentials.AccessToken = source.ConnectionConfiguration.Credentials.AccessToken

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
func (r *sourceZendeskSupportResource) Read(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	var state sourceZendeskSupportResourceModel

	// Get current state
	err := fwhelpers.UnpackModel(req.StateContents, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	res := schema.ServiceResponse{}

	if req.StateID != "" {
		// Query using existing previous state.
		source, err := r.Client.ReadZendeskSupportSource(req.StateID)
		if err != nil {
			return schema.ErrorResponse(err)
		}

		// Update state with refreshed value
		state.Name = source.Name
		state.SourceDefinitionId = source.SourceDefinitionId
		state.SourceId = source.SourceId
		state.WorkspaceId = source.WorkspaceId

		state.ConnectionConfiguration = sourceZendeskSupportConnConfigModel{}
		state.ConnectionConfiguration.StartDate = source.ConnectionConfiguration.StartDate
		state.ConnectionConfiguration.IgnorePagination = source.ConnectionConfiguration.IgnorePagination
		state.ConnectionConfiguration.Subdomain = source.ConnectionConfiguration.Subdomain
		state.ConnectionConfiguration.Credentials = ZendeskSupportCredConfigModel{}
		state.ConnectionConfiguration.Credentials.Credentials = source.ConnectionConfiguration.Credentials.Credentials
		state.ConnectionConfiguration.Credentials.ApiToken = source.ConnectionConfiguration.Credentials.ApiToken
		state.ConnectionConfiguration.Credentials.Email = source.ConnectionConfiguration.Credentials.Email
		state.ConnectionConfiguration.Credentials.AccessToken = source.ConnectionConfiguration.Credentials.AccessToken

		res.StateID = state.SourceId
	} else {
		// No previous state exists.
		res.StateID = ""
	}

	// Set refreshed state
	stateEnc, err := fwhelpers.PackModel(nil, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}
	res.StateContents = stateEnc

	return &res
}

func (r *sourceZendeskSupportResource) Update(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	// Retrieve values from plan
	var plan sourceZendeskSupportResourceModel
	err := fwhelpers.UnpackModel(req.PlanContents, &plan)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Generate API request body from plan
	body := api.SourceZendeskSupport{}
	body.Name = plan.Name
	body.SourceId = plan.SourceId

	body.ConnectionConfiguration = api.SourceZendeskSupportConnConfig{}
	body.ConnectionConfiguration.StartDate = plan.ConnectionConfiguration.StartDate
	body.ConnectionConfiguration.IgnorePagination = plan.ConnectionConfiguration.IgnorePagination
	body.ConnectionConfiguration.Subdomain = plan.ConnectionConfiguration.Subdomain
	body.ConnectionConfiguration.Credentials = api.ZendeskSupportCredConfigModel{}
	body.ConnectionConfiguration.Credentials.Credentials = plan.ConnectionConfiguration.Credentials.Credentials
	body.ConnectionConfiguration.Credentials.ApiToken = plan.ConnectionConfiguration.Credentials.ApiToken
	body.ConnectionConfiguration.Credentials.Email = plan.ConnectionConfiguration.Credentials.Email
	body.ConnectionConfiguration.Credentials.AccessToken = plan.ConnectionConfiguration.Credentials.AccessToken
	// Update existing source
	_, err = r.Client.UpdateZendeskSupportSource(body)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Fetch updated items
	source, err := r.Client.ReadZendeskSupportSource(req.PlanID)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Update state with refreshed value
	state := sourceZendeskSupportResourceModel{}
	state.Name = source.Name
	state.SourceDefinitionId = source.SourceDefinitionId
	state.SourceId = source.SourceId
	state.WorkspaceId = source.WorkspaceId

	state.ConnectionConfiguration = sourceZendeskSupportConnConfigModel{}
	state.ConnectionConfiguration.StartDate = source.ConnectionConfiguration.StartDate
	state.ConnectionConfiguration.IgnorePagination = source.ConnectionConfiguration.IgnorePagination
	state.ConnectionConfiguration.Subdomain = source.ConnectionConfiguration.Subdomain
	state.ConnectionConfiguration.Credentials = ZendeskSupportCredConfigModel{}
	state.ConnectionConfiguration.Credentials.Credentials = source.ConnectionConfiguration.Credentials.Credentials
	state.ConnectionConfiguration.Credentials.ApiToken = source.ConnectionConfiguration.Credentials.ApiToken
	state.ConnectionConfiguration.Credentials.Email = source.ConnectionConfiguration.Credentials.Email
	state.ConnectionConfiguration.Credentials.AccessToken = source.ConnectionConfiguration.Credentials.AccessToken

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
func (r *sourceZendeskSupportResource) Delete(req *schema.ServiceRequest) *schema.ServiceResponse {
	// Delete existing source
	err := r.Client.DeleteZendeskSupportSource(req.StateID)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{}
}
