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
var _ resource.Resource = &GatewayPluginRateLimitingAdvancedResource{}
var _ resource.ResourceWithImportState = &GatewayPluginRateLimitingAdvancedResource{}

func NewGatewayPluginRateLimitingAdvancedResource() resource.Resource {
	return &GatewayPluginRateLimitingAdvancedResource{}
}

// GatewayPluginRateLimitingAdvancedResource defines the resource implementation.
type GatewayPluginRateLimitingAdvancedResource struct {
	client *sdk.Konnect
}

// GatewayPluginRateLimitingAdvancedResourceModel describes the resource data model.
type GatewayPluginRateLimitingAdvancedResourceModel struct {
	Config         tfTypes.CreateRateLimitingAdvancedPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                           `tfsdk:"consumer"`
	ControlPlaneID types.String                                   `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                                    `tfsdk:"created_at"`
	Enabled        types.Bool                                     `tfsdk:"enabled"`
	ID             types.String                                   `tfsdk:"id"`
	Protocols      []types.String                                 `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                           `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                           `tfsdk:"service"`
	Tags           []types.String                                 `tfsdk:"tags"`
}

func (r *GatewayPluginRateLimitingAdvancedResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_rate_limiting_advanced"
}

func (r *GatewayPluginRateLimitingAdvancedResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginRateLimitingAdvanced Resource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"consumer_groups": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `List of consumer groups allowed to override the rate limiting settings for the given Route or Service. Required if ` + "`" + `enforce_consumer_groups` + "`" + ` is set to ` + "`" + `true` + "`" + `.`,
					},
					"dictionary_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("kong_rate_limiting_counters"),
						Description: `The shared dictionary where counters are stored. When the plugin is configured to synchronize counter data externally (that is ` + "`" + `config.strategy` + "`" + ` is ` + "`" + `cluster` + "`" + ` or ` + "`" + `redis` + "`" + ` and ` + "`" + `config.sync_rate` + "`" + ` isn't ` + "`" + `-1` + "`" + `), this dictionary serves as a buffer to populate counters in the data store on each synchronization cycle. Default: "kong_rate_limiting_counters"`,
					},
					"disable_penalty": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `If set to ` + "`" + `true` + "`" + `, this doesn't count denied requests (status = ` + "`" + `429` + "`" + `). If set to ` + "`" + `false` + "`" + `, all requests, including denied ones, are counted. This parameter only affects the ` + "`" + `sliding` + "`" + ` window_type. Default: false`,
					},
					"enforce_consumer_groups": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Determines if consumer groups are allowed to override the rate limiting settings for the given Route or Service. Flipping ` + "`" + `enforce_consumer_groups` + "`" + ` from ` + "`" + `true` + "`" + ` to ` + "`" + `false` + "`" + ` disables the group override, but does not clear the list of consumer groups. You can then flip ` + "`" + `enforce_consumer_groups` + "`" + ` to ` + "`" + `true` + "`" + ` to re-enforce the groups. Default: false`,
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
					"header_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing an HTTP header name.`,
					},
					"hide_client_headers": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Optionally hide informative response headers that would otherwise provide information about the current status of limits and counters. Default: false`,
					},
					"identifier": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("consumer"),
						Description: `The type of identifier used to generate the rate limit key. Defines the scope used to increment the rate limiting counters. Can be ` + "`" + `ip` + "`" + `, ` + "`" + `credential` + "`" + `, ` + "`" + `consumer` + "`" + `, ` + "`" + `service` + "`" + `, ` + "`" + `header` + "`" + `, ` + "`" + `path` + "`" + ` or ` + "`" + `consumer-group` + "`" + `. must be one of ["ip", "credential", "consumer", "service", "header", "path", "consumer-group"]; Default: "consumer"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"ip",
								"credential",
								"consumer",
								"service",
								"header",
								"path",
								"consumer-group",
							),
						},
					},
					"limit": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.NumberType,
						Description: `One or more requests-per-window limits to apply. There must be a matching number of window limits and sizes specified.`,
					},
					"namespace": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The rate limiting library namespace to use for this plugin instance. Counter data and sync configuration is isolated in each namespace. NOTE: For the plugin instances sharing the same namespace, all the configurations that are required for synchronizing counters, e.g. ` + "`" + `strategy` + "`" + `, ` + "`" + `redis` + "`" + `, ` + "`" + `sync_rate` + "`" + `, ` + "`" + `window_size` + "`" + `, ` + "`" + `dictionary_name` + "`" + `, need to be the same.`,
					},
					"path": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).`,
					},
					"redis": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"cluster_addresses": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `Cluster addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this value implies using Redis Cluster. Each string element must be a hostname. The minimum length of the array is 1 element.`,
							},
							"connect_timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
							},
							"database": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(0),
								Description: `Database to use for the Redis connection when using the ` + "`" + `redis` + "`" + ` strategy. Default: 0`,
							},
							"host": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `A string representing a host name, such as example.com.`,
							},
							"keepalive_backlog": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return ` + "`" + `nil` + "`" + `. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than ` + "`" + `keepalive_pool_size` + "`" + `. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than ` + "`" + `keepalive_pool_size` + "`" + `.`,
							},
							"keepalive_pool_size": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(256),
								Description: `The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither ` + "`" + `keepalive_pool_size` + "`" + ` nor ` + "`" + `keepalive_backlog` + "`" + ` is specified, no pool is created. If ` + "`" + `keepalive_pool_size` + "`" + ` isn't specified but ` + "`" + `keepalive_backlog` + "`" + ` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low. Default: 256`,
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.`,
							},
							"port": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a port number between 0 and 65535, inclusive.`,
							},
							"read_timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
							},
							"send_timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
							},
							"sentinel_addresses": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `Sentinel addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this value implies using Redis Sentinel. Each string element must be a hostname. The minimum length of the array is 1 element.`,
							},
							"sentinel_master": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.`,
							},
							"sentinel_password": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.`,
							},
							"sentinel_role": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel role to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this value implies using Redis Sentinel. must be one of ["master", "slave", "any"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"master",
										"slave",
										"any",
									),
								},
							},
							"sentinel_username": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.`,
							},
							"server_name": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `A string representing an SNI (server name indication) value for TLS.`,
							},
							"ssl": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
								Description: `If set to true, uses SSL to connect to Redis. Default: false`,
							},
							"ssl_verify": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
								Description: `If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure ` + "`" + `lua_ssl_trusted_certificate` + "`" + ` in ` + "`" + `kong.conf` + "`" + ` to specify the CA (or server) certificate used by your Redis server. You may also need to configure ` + "`" + `lua_ssl_verify_depth` + "`" + ` accordingly. Default: false`,
							},
							"timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(2000),
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2. Default: 2000`,
							},
							"username": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to ` + "`" + `default` + "`" + `.`,
							},
						},
					},
					"retry_after_jitter_max": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(0)),
						Description: `The upper bound of a jitter (random delay) in seconds to be added to the ` + "`" + `Retry-After` + "`" + ` header of denied requests (status = ` + "`" + `429` + "`" + `) in order to prevent all the clients from coming back at the same time. The lower bound of the jitter is ` + "`" + `0` + "`" + `; in this case, the ` + "`" + `Retry-After` + "`" + ` header is equal to the ` + "`" + `RateLimit-Reset` + "`" + ` header. Default: 0`,
					},
					"strategy": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("local"),
						Description: `The rate-limiting strategy to use for retrieving and incrementing the limits. Available values are: ` + "`" + `local` + "`" + ` and ` + "`" + `cluster` + "`" + `. must be one of ["cluster", "redis", "local"]; Default: "local"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"cluster",
								"redis",
								"local",
							),
						},
					},
					"sync_rate": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `How often to sync counter data to the central data store. A value of 0 results in synchronous behavior; a value of -1 ignores sync behavior entirely and only stores counters in node memory. A value greater than 0 will sync the counters in the specified number of seconds. The minimum allowed interval is 0.02 seconds (20ms).`,
					},
					"window_size": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.NumberType,
						Description: `One or more window sizes to apply a limit to (defined in seconds). There must be a matching number of window limits and sizes specified.`,
					},
					"window_type": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("sliding"),
						Description: `Sets the time window type to either ` + "`" + `sliding` + "`" + ` (default) or ` + "`" + `fixed` + "`" + `. Sliding windows apply the rate limiting logic while taking into account previous hit rates (from the window that immediately precedes the current) using a dynamic weight. Fixed windows consist of buckets that are statically assigned to a definitive time range, each request is mapped to only one fixed window based on its timestamp and will affect only that window's counters. must be one of ["fixed", "sliding"]; Default: "sliding"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"fixed",
								"sliding",
							),
						},
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

func (r *GatewayPluginRateLimitingAdvancedResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginRateLimitingAdvancedResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginRateLimitingAdvancedResourceModel
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
	createRateLimitingAdvancedPlugin := *data.ToSharedCreateRateLimitingAdvancedPlugin()
	request := operations.CreateRatelimitingadvancedPluginRequest{
		ControlPlaneID:                   controlPlaneID,
		CreateRateLimitingAdvancedPlugin: createRateLimitingAdvancedPlugin,
	}
	res, err := r.client.Plugins.CreateRatelimitingadvancedPlugin(ctx, request)
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
	if res.RateLimitingAdvancedPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRateLimitingAdvancedPlugin(res.RateLimitingAdvancedPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginRateLimitingAdvancedResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginRateLimitingAdvancedResourceModel
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
	request := operations.GetRatelimitingadvancedPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       pluginID,
	}
	res, err := r.client.Plugins.GetRatelimitingadvancedPlugin(ctx, request)
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
	if res.RateLimitingAdvancedPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRateLimitingAdvancedPlugin(res.RateLimitingAdvancedPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginRateLimitingAdvancedResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginRateLimitingAdvancedResourceModel
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
	createRateLimitingAdvancedPlugin := *data.ToSharedCreateRateLimitingAdvancedPlugin()
	request := operations.UpdateRatelimitingadvancedPluginRequest{
		ControlPlaneID:                   controlPlaneID,
		PluginID:                         pluginID,
		CreateRateLimitingAdvancedPlugin: createRateLimitingAdvancedPlugin,
	}
	res, err := r.client.Plugins.UpdateRatelimitingadvancedPlugin(ctx, request)
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
	if res.RateLimitingAdvancedPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRateLimitingAdvancedPlugin(res.RateLimitingAdvancedPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginRateLimitingAdvancedResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginRateLimitingAdvancedResourceModel
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
	request := operations.DeleteRatelimitingadvancedPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       pluginID,
	}
	res, err := r.client.Plugins.DeleteRatelimitingadvancedPlugin(ctx, request)
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

func (r *GatewayPluginRateLimitingAdvancedResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_plugin_rate_limiting_advanced. Reason: composite imports strings not supported.")
}
