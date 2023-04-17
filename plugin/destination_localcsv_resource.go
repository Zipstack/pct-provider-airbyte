package plugin

import (
	"fmt"
	"time"

	"github.com/zipstack/pct-plugin-framework/fwhelpers"
	"github.com/zipstack/pct-plugin-framework/schema"

	"github.com/zipstack/pct-provider-airbyte-local/api"
)

// Resource implementation.
type destinationLocalCSVResource struct {
	Client *api.Client
}

type destinationLocalCSVResourceModel struct {
	Name                    string                             `pctsdk:"name"`
	DestinationId           string                             `pctsdk:"destination_id"`
	DestinationDefinitionId string                             `pctsdk:"destination_definition_id"`
	WorkspaceId             string                             `pctsdk:"workspace_id"`
	ConnectionConfiguration destinationLocalCSVConnConfigModel `pctsdk:"connection_configuration"`
}
type destinationLocalCSVConnConfigModel struct {
	DestinationPath string                          `pctsdk:"destination_path"`
	DelimiterType   destinationDelimiterConfigModel `pctsdk:"delimiter_type"`
}

type destinationDelimiterConfigModel struct {
	Delimiter string `pctsdk:"delimiter"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ schema.ResourceService = &destinationLocalCSVResource{}
)

// Helper function to return a resource service instance.
func NewDestinationLocalCSVResource() schema.ResourceService {
	return &destinationLocalCSVResource{}
}

// Metadata returns the resource type name.
// It is always provider name + "_" + resource type name.
func (r *destinationLocalCSVResource) Metadata(req *schema.ServiceRequest) *schema.ServiceResponse {
	return &schema.ServiceResponse{
		TypeName: req.TypeName + "_destination_localcsv",
	}
}

// Configure adds the provider configured client to the resource.
func (r *destinationLocalCSVResource) Configure(req *schema.ServiceRequest) *schema.ServiceResponse {
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
func (r *destinationLocalCSVResource) Schema() *schema.ServiceResponse {
	s := &schema.Schema{
		Description: "Destination local CSV resource for Airbyte",
		Attributes: map[string]schema.Attribute{
			"name": &schema.StringAttribute{
				Description: "Name",
				Required:    true,
			},
			"destination_id": &schema.StringAttribute{
				Description: "Destination ID",
				Required:    false,
				Computed:    true,
			},
			"destination_definition_id": &schema.StringAttribute{
				Description: "Destination Definition ID",
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
					"destination_path": &schema.StringAttribute{
						Description: "Destination Path",
						Required:    true,
					},
					"delimiter_type": &schema.MapAttribute{
						Description: "Delimiter Type",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"delimiter": &schema.StringAttribute{
								Description: "Delimiter",
								Required:    true,
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
func (r *destinationLocalCSVResource) Create(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	// Retrieve values from plan
	var plan destinationLocalCSVResourceModel
	err := fwhelpers.UnpackModel(req.PlanContents, &plan)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Generate API request body from plan
	body := api.DestinationLocalCSV{}
	body.Name = plan.Name
	body.DestinationDefinitionId = plan.DestinationDefinitionId
	body.WorkspaceId = plan.WorkspaceId

	body.ConnectionConfiguration = api.DestinationLocalCSVConnConfigModel{}
	body.ConnectionConfiguration.DestinationPath = plan.ConnectionConfiguration.DestinationPath
	body.ConnectionConfiguration.DelimiterType = api.DestinationDelimiterConfigModel{}
	body.ConnectionConfiguration.DelimiterType.Delimiter = plan.ConnectionConfiguration.DelimiterType.Delimiter

	// Create new source
	destination, err := r.Client.CreateLocalCSVDestination(body)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Update resource state with response body
	state := destinationLocalCSVResourceModel{}
	state.Name = destination.Name
	state.DestinationDefinitionId = destination.DestinationDefinitionId
	state.DestinationId = destination.DestinationId
	state.WorkspaceId = destination.WorkspaceId

	state.ConnectionConfiguration = destinationLocalCSVConnConfigModel{}
	state.ConnectionConfiguration.DestinationPath = destination.ConnectionConfiguration.DestinationPath
	state.ConnectionConfiguration.DelimiterType = destinationDelimiterConfigModel{}
	state.ConnectionConfiguration.DelimiterType.Delimiter = destination.ConnectionConfiguration.DelimiterType.Delimiter

	// Set refreshed state
	stateEnc, err := fwhelpers.PackModel(nil, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}
	return &schema.ServiceResponse{
		StateID:          state.DestinationId,
		StateContents:    stateEnc,
		StateLastUpdated: time.Now().Format(time.RFC850),
	}
}

// Read resource information
func (r *destinationLocalCSVResource) Read(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	var state destinationLocalCSVResourceModel

	// Get current state
	err := fwhelpers.UnpackModel(req.StateContents, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	res := schema.ServiceResponse{}

	if req.StateID != "" {
		// Query using existing previous state.
		destination, err := r.Client.ReadLocalCSVDestination(req.StateID)
		if err != nil {
			return schema.ErrorResponse(err)
		}

		// Update state with refreshed value
		state.Name = destination.Name
		state.DestinationDefinitionId = destination.DestinationDefinitionId
		state.DestinationId = destination.DestinationId
		state.WorkspaceId = destination.WorkspaceId

		state.ConnectionConfiguration = destinationLocalCSVConnConfigModel{}
		state.ConnectionConfiguration.DestinationPath = destination.ConnectionConfiguration.DestinationPath
		state.ConnectionConfiguration.DelimiterType = destinationDelimiterConfigModel{}
		state.ConnectionConfiguration.DelimiterType.Delimiter = destination.ConnectionConfiguration.DelimiterType.Delimiter

		res.StateID = state.DestinationId
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

func (r *destinationLocalCSVResource) Update(req *schema.ServiceRequest) *schema.ServiceResponse {
	// logger := fwhelpers.GetLogger()

	// Retrieve values from plan
	var plan destinationLocalCSVResourceModel
	err := fwhelpers.UnpackModel(req.PlanContents, &plan)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Generate API request body from plan
	body := api.DestinationLocalCSV{}
	body.Name = plan.Name
	body.DestinationId = plan.DestinationId

	body.ConnectionConfiguration = api.DestinationLocalCSVConnConfigModel{}
	body.ConnectionConfiguration.DestinationPath = plan.ConnectionConfiguration.DestinationPath
	body.ConnectionConfiguration.DelimiterType = api.DestinationDelimiterConfigModel{}
	body.ConnectionConfiguration.DelimiterType.Delimiter = plan.ConnectionConfiguration.DelimiterType.Delimiter

	// Update existing source
	_, err = r.Client.UpdateLocalCSVDestination(body)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Fetch updated items
	destination, err := r.Client.ReadLocalCSVDestination(req.PlanID)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	// Update state with refreshed value
	state := destinationLocalCSVResourceModel{}
	state.Name = destination.Name
	state.DestinationDefinitionId = destination.DestinationDefinitionId
	state.DestinationId = destination.DestinationId
	state.WorkspaceId = destination.WorkspaceId

	state.ConnectionConfiguration = destinationLocalCSVConnConfigModel{}
	state.ConnectionConfiguration.DestinationPath = destination.ConnectionConfiguration.DestinationPath
	state.ConnectionConfiguration.DelimiterType = destinationDelimiterConfigModel{}
	state.ConnectionConfiguration.DelimiterType.Delimiter = destination.ConnectionConfiguration.DelimiterType.Delimiter

	// Set refreshed state
	stateEnc, err := fwhelpers.PackModel(nil, &state)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{
		StateID:          state.DestinationId,
		StateContents:    stateEnc,
		StateLastUpdated: time.Now().Format(time.RFC850),
	}
}

// Delete deletes the resource and removes the state on success.
func (r *destinationLocalCSVResource) Delete(req *schema.ServiceRequest) *schema.ServiceResponse {
	// Delete existing source
	err := r.Client.DeleteLocalCSVDestination(req.StateID)
	if err != nil {
		return schema.ErrorResponse(err)
	}

	return &schema.ServiceResponse{}
}
