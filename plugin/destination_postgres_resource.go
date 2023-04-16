package plugin

// import (
// 	"fmt"

// 	"github.com/zipstack/pct-plugin-framework/fwhelpers"
// 	"github.com/zipstack/pct-plugin-framework/schema"

// 	"github.com/zipstack/pct-provider-airbyte/api"
// )

// // Resource implementation.
// type destinationPostgresResource struct {
// 	Client *api.Client
// }

// type destinationPostgresResourceModel struct {
// 	Name                    string                 `pctsdk:"name"`
// 	DestinationDefinitionId string                 `pctsdk:"destinationDefinitionId"`
// 	DestinationId           string                 `pctsdk:"destinationId"`
// 	WorkspaceId             string                 `pctsdk:"workspaceId"`
// 	ConnectionConfiguration destPostgresConnConfig `pctsdk:"connectionConfiguration"`
// }

// type destPostgresConnConfig struct {
// 	Host          string                             `pctsdk:"host"`
// 	Port          int64                              `pctsdk:"port"`
// 	Schema        string                             `pctsdk:"schema"`
// 	Database      string                             `pctsdk:"database"`
// 	Username      string                             `pctsdk:"username"`
// 	Password      string                             `pctsdk:"password"`
// 	Ssl           bool                               `pctsdk:"ssl"`
// 	SslMode       destPostgresConnConfigSslMode      `pctsdk:"ssl_mode"`
// 	TunnelMethod  destPostgresConnConfigTunnelMethod `pctsdk:"tunnel_method"`
// 	JdbcUrlParams string                             `pctsdk:"jdbc_url_params"`
// }

// type destPostgresConnConfigSslMode struct {
// 	Mode              string `pctsdk:"mode"`
// 	ClientKey         string `pctsdk:"client_key"`
// 	CaCertificate     string `pctsdk:"ca_certificate"`
// 	ClientCertificate string `pctsdk:"client_certificate"`
// 	ClientKeyPassword string `pctsdk:"client_key_password"`
// }

// type destPostgresConnConfigTunnelMethod struct {
// 	TunnelMethod        string `pctsdk:"tunnel_method"`
// 	TunnelHost          string `pctsdk:"tunnel_host"`
// 	TunnelPort          int64  `pctsdk:"tunnel_port"`
// 	TunnelUser          string `pctsdk:"tunnel_user"`
// 	TunnelUser_password string `pctsdk:"tunnel_user_password"`
// }

// // Ensure the implementation satisfies the expected interfaces.
// var (
// 	_ schema.ResourceService = &destinationPostgresResource{}
// )

// // Helper function to return a resource service instance.
// func NewDestinationPostgresResource() schema.ResourceService {
// 	return &destinationPostgresResource{}
// }

// // Metadata returns the resource type name.
// // It is always provider name + "_" + resource type name.
// func (r *destinationPostgresResource) Metadata(req *schema.ServiceRequest) *schema.ServiceResponse {
// 	return &schema.ServiceResponse{
// 		TypeName: req.TypeName + "_destination_postgres",
// 	}
// }

// // Configure adds the provider configured client to the resource.
// func (r *destinationPostgresResource) Configure(req *schema.ServiceRequest) *schema.ServiceResponse {
// 	if req.ResourceData == "" {
// 		return schema.ErrorResponse(fmt.Errorf("no data provided to configure resource"))
// 	}

// 	var client *api.Client
// 	err := fwhelpers.Decode(req.ResourceData, &client)
// 	if err != nil {
// 		return schema.ErrorResponse(err)
// 	}

// 	r.Client = client

// 	return &schema.ServiceResponse{}
// }

// // Schema defines the schema for the resource.
// func (r *destinationPostgresResource) Schema() *schema.ServiceResponse {
// 	s := &schema.Schema{
// 		Description: "Destination Postgres resource for Airbyte",
// 		Attributes: map[string]schema.Attribute{
// 			"name": &schema.StringAttribute{
// 				Description: "Name",
// 				Required:    true,
// 			},
// 			"destination_definition_id": &schema.StringAttribute{
// 				Description: "Destination Definition ID",
// 				Required:    true,
// 			},
// 			"destination_id": &schema.StringAttribute{
// 				Description: "Destination ID",
// 				Required:    false,
// 				Computed:    true,
// 			},
// 			"workspace_id": &schema.StringAttribute{
// 				Description: "Workspace ID",
// 				Required:    true,
// 			},
// 			"connection_configuration": &schema.MapAttribute{
// 				Description: "Connection Configuration",
// 				Required:    true,
// 				Attributes: map[string]schema.Attribute{
// 					"host": &schema.StringAttribute{
// 						Description: "Host",
// 						Required:    true,
// 					},
// 					"port": &schema.IntAttribute{
// 						Description: "Port",
// 						Required:    true,
// 					},
// 					"schema": &schema.StringAttribute{
// 						Description: "Schema",
// 						Required:    true,
// 					},
// 					"database": &schema.StringAttribute{
// 						Description: "Database",
// 						Required:    true,
// 					},
// 					"username": &schema.StringAttribute{
// 						Description: "Username",
// 						Required:    true,
// 					},
// 					"password": &schema.StringAttribute{
// 						Description: "Password",
// 						Sensitive:   true,
// 						Required:    true,
// 					},
// 					"ssl": &schema.BoolAttribute{
// 						Description: "SSL",
// 						Required:    false,
// 					},
// 					"ssl_mode": &schema.MapAttribute{
// 						Description: "SSL mode",
// 						Required:    true,
// 						Attributes: map[string]schema.Attribute{
// 							"mode": &schema.StringAttribute{
// 								Description: "Mode",
// 								Required:    true,
// 							},
// 							"client_key": &schema.StringAttribute{
// 								Description: "Client Key",
// 								Required:    true,
// 							},
// 							"ca_certificate": &schema.StringAttribute{
// 								Description: "CA certificate",
// 								Required:    true,
// 							},
// 							"client_certificate": &schema.StringAttribute{
// 								Description: "Client certificate",
// 								Required:    true,
// 							},
// 							"client_key_password": &schema.StringAttribute{
// 								Description: "Client key password",
// 								Required:    true,
// 								Sensitive:   true,
// 							},
// 						},
// 					},
// 					"tunnel_method": &schema.MapAttribute{
// 						Description: "Tunnel method",
// 						Required:    true,
// 						Attributes: map[string]schema.Attribute{
// 							"tunnel_method": &schema.StringAttribute{
// 								Description: "Tunnel method",
// 								Required:    true,
// 							},
// 							"tunnel_host": &schema.StringAttribute{
// 								Description: "Tunnel host",
// 								Required:    true,
// 							},
// 							"tunnel_port": &schema.IntAttribute{
// 								Description: "Tunnel port",
// 								Required:    false,
// 							},
// 							"tunnel_user": &schema.StringAttribute{
// 								Description: "Tunnel user",
// 								Required:    true,
// 							},
// 							"tunnel_user_password": &schema.StringAttribute{
// 								Description: "Tunnel user password",
// 								Required:    true,
// 								Sensitive:   true,
// 							},
// 						},
// 					},
// 					"jdbc_url_params": &schema.StringAttribute{
// 						Description: "JDBC URL params",
// 						Required:    false,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	sEnc, err := fwhelpers.Encode(s)
// 	if err != nil {
// 		return schema.ErrorResponse(err)
// 	}

// 	return &schema.ServiceResponse{
// 		SchemaContents: sEnc,
// 	}
// }

// // Create a new resource
// func (r *destinationPostgresResource) Create(req *schema.ServiceRequest) *schema.ServiceResponse {
// 	return &schema.ServiceResponse{}

// 	// logger := fwhelpers.GetLogger()

// 	// var plan sourceResourceModel
// 	// err := fwhelpers.UnpackModel(req.PlanContents, &plan)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// // Generate API request body from plan
// 	// // plan??

// 	// // Create new order
// 	// source, err := r.Client.CreateSource(plan)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// // Map response body to schema and populate Computed attribute values
// 	// state := make(map[string]string)
// 	// // plan.ID = types.StringValue(strconv.Itoa(order.ID))
// 	// // plan.Items[int]model
// 	// // plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

// 	// stateEnc, err := fwhelpers.Encode(state)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// return &schema.ServiceResponse{
// 	// 	StateContents: stateEnc,
// 	// }
// }

// // Read resource information
// func (r *destinationPostgresResource) Read(req *schema.ServiceRequest) *schema.ServiceResponse {
// 	return &schema.ServiceResponse{}

// 	// // Get current state
// 	// var state sourceResourceModel
// 	// err := fwhelpers.UnpackModel(req.StateContents, &state)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// // Get refreshed source value
// 	// source, err := r.Client.GetSource(state.ID)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// // Overwrite items with refreshed state
// 	// state.Items = []sourceResourceModel{}
// 	// // fill state.Items with source

// 	// // Set refreshed state
// 	// stateEnc, err := fwhelpers.Encode(state)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// return &schema.ServiceResponse{
// 	// 	StateContents: stateEnc,
// 	// }
// }

// func (r *destinationPostgresResource) Update(req *schema.ServiceRequest) *schema.ServiceResponse {
// 	return &schema.ServiceResponse{}

// 	// // Retrieve values from plan
// 	// var plan sourceResourceModel
// 	// err := fwhelpers.UnpackModel(req.StateContents, &plan)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// // Generate API request body from plan
// 	// // plan??

// 	// // Update existing source
// 	// source, err := r.Client.UpdateSource(plan)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// // Fetch updated items from GetSource as UpdateSource items are not
// 	// // populated.
// 	// source, err := r.Client.GetSource(state.ID)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// // Update resource state with updated items and timestamp
// 	// state.Items = []sourceResourceModel{}
// 	// // fill state.Items with source
// 	// // plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

// 	// // Set refreshed state
// 	// stateEnc, err := fwhelpers.Encode(state)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// return &schema.ServiceResponse{
// 	// 	StateContents: stateEnc,
// 	// }
// }

// // Delete deletes the resource and removes the state on success.
// func (r *destinationPostgresResource) Delete(req *schema.ServiceRequest) *schema.ServiceResponse {
// 	return &schema.ServiceResponse{}

// 	// // Retrieve values from state
// 	// var state sourceResourceModel
// 	// err := fwhelpers.UnpackModel(req.StateContents, &state)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// // Delete existing source
// 	// source, err := r.Client.DeleteSource(state.ID)
// 	// if err != nil {
// 	// 	return schema.ErrorResponse(err)
// 	// }

// 	// return &schema.ServiceResponse{}
// }
