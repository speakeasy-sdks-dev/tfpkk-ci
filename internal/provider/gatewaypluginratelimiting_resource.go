// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"math/big"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginRateLimitingResource{}
var _ resource.ResourceWithImportState = &GatewayPluginRateLimitingResource{}

func NewGatewayPluginRateLimitingResource() resource.Resource {
	return &GatewayPluginRateLimitingResource{}
}

// GatewayPluginRateLimitingResource defines the resource implementation.
type GatewayPluginRateLimitingResource struct {
	client *sdk.Konnect
}

// GatewayPluginRateLimitingResourceModel describes the resource data model.
type GatewayPluginRateLimitingResourceModel struct {
	Config         tfTypes.CreateRateLimitingPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                   `tfsdk:"consumer"`
	ControlPlaneID types.String                           `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                            `tfsdk:"created_at"`
	Enabled        types.Bool                             `tfsdk:"enabled"`
	ID             types.String                           `tfsdk:"id"`
	Protocols      []types.String                         `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                   `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                   `tfsdk:"service"`
	Tags           []types.String                         `tfsdk:"tags"`
}

func (r *GatewayPluginRateLimitingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_rate_limiting"
}

func (r *GatewayPluginRateLimitingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginRateLimiting Resource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"day": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The number of HTTP requests that can be made per day.`,
					},
					"error_code": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(429)),
						Description: `Set a custom error code to return when the rate limit is exceeded. Default: 429`,
					},
					"error_message": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("API rate limit exceeded"),
						Description: `Set a custom error message to return when the rate limit is exceeded. Default: "API rate limit exceeded"`,
					},
					"fault_tolerant": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(true),
						Description: `A boolean value that determines if the requests should be proxied even if Kong has troubles connecting a third-party data store. If ` + "`" + `true` + "`" + `, requests will be proxied anyway, effectively disabling the rate-limiting function until the data store is working again. If ` + "`" + `false` + "`" + `, then the clients will see ` + "`" + `500` + "`" + ` errors. Default: true`,
					},
					"header_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing an HTTP header name.`,
					},
					"hide_client_headers": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Optionally hide informative response headers. Default: false`,
					},
					"hour": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The number of HTTP requests that can be made per hour.`,
					},
					"limit_by": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("consumer"),
						Description: `The entity that is used when aggregating the limits. must be one of ["consumer", "credential", "ip", "service", "header", "path", "consumer-group"]; Default: "consumer"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"consumer",
								"credential",
								"ip",
								"service",
								"header",
								"path",
								"consumer-group",
							),
						},
					},
					"minute": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The number of HTTP requests that can be made per minute.`,
					},
					"month": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The number of HTTP requests that can be made per month.`,
					},
					"path": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).`,
					},
					"policy": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("local"),
						Description: `The rate-limiting policies to use for retrieving and incrementing the limits. must be one of ["local", "cluster", "redis"]; Default: "local"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"local",
								"cluster",
								"redis",
							),
						},
					},
					"redis_database": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(0),
						Description: `When using the ` + "`" + `redis` + "`" + ` policy, this property specifies the Redis database to use. Default: 0`,
					},
					"redis_host": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a host name, such as example.com.`,
					},
					"redis_password": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `When using the ` + "`" + `redis` + "`" + ` policy, this property specifies the password to connect to the Redis server.`,
					},
					"redis_port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(6379),
						Description: `An integer representing a port number between 0 and 65535, inclusive. Default: 6379`,
					},
					"redis_server_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing an SNI (server name indication) value for TLS.`,
					},
					"redis_ssl": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `When using the ` + "`" + `redis` + "`" + ` policy, this property specifies if SSL is used to connect to the Redis server. Default: false`,
					},
					"redis_ssl_verify": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `When using the ` + "`" + `redis` + "`" + ` policy with ` + "`" + `redis_ssl` + "`" + ` set to ` + "`" + `true` + "`" + `, this property specifies it server SSL certificate is validated. Note that you need to configure the lua_ssl_trusted_certificate to specify the CA (or server) certificate used by your Redis server. You may also need to configure lua_ssl_verify_depth accordingly. Default: false`,
					},
					"redis_timeout": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(2000)),
						Description: `When using the ` + "`" + `redis` + "`" + ` policy, this property specifies the timeout in milliseconds of any command submitted to the Redis server. Default: 2000`,
					},
					"redis_username": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `When using the ` + "`" + `redis` + "`" + ` policy, this property specifies the username to connect to the Redis server when ACL authentication is desired.`,
					},
					"second": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The number of HTTP requests that can be made per second.`,
					},
					"sync_rate": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(-1)),
						Description: `How often to sync counter data to the central data store. A value of -1 results in synchronous behavior. Default: -1`,
					},
					"year": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The number of HTTP requests that can be made per year.`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
				Description: `Whether the plugin is applied. Default: true`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the Plugin to lookup`,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
		},
	}
}

func (r *GatewayPluginRateLimitingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GatewayPluginRateLimitingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginRateLimitingResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	createRateLimitingPlugin := *data.ToSharedCreateRateLimitingPlugin()
	request := operations.CreateRatelimitingPluginRequest{
		ControlPlaneID:           controlPlaneID,
		CreateRateLimitingPlugin: createRateLimitingPlugin,
	}
	res, err := r.client.Plugins.CreateRatelimitingPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.RateLimitingPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRateLimitingPlugin(res.RateLimitingPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginRateLimitingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginRateLimitingResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	pluginID := data.ID.ValueString()
	request := operations.GetRatelimitingPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       pluginID,
	}
	res, err := r.client.Plugins.GetRatelimitingPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.RateLimitingPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRateLimitingPlugin(res.RateLimitingPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginRateLimitingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginRateLimitingResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	pluginID := data.ID.ValueString()
	createRateLimitingPlugin := *data.ToSharedCreateRateLimitingPlugin()
	request := operations.UpdateRatelimitingPluginRequest{
		ControlPlaneID:           controlPlaneID,
		PluginID:                 pluginID,
		CreateRateLimitingPlugin: createRateLimitingPlugin,
	}
	res, err := r.client.Plugins.UpdateRatelimitingPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.RateLimitingPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRateLimitingPlugin(res.RateLimitingPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginRateLimitingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginRateLimitingResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	controlPlaneID := data.ControlPlaneID.ValueString()
	pluginID := data.ID.ValueString()
	request := operations.DeleteRatelimitingPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       pluginID,
	}
	res, err := r.client.Plugins.DeleteRatelimitingPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *GatewayPluginRateLimitingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_plugin_rate_limiting. Reason: composite imports strings not supported.")
}
