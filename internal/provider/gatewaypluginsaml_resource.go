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
var _ resource.Resource = &GatewayPluginSamlResource{}
var _ resource.ResourceWithImportState = &GatewayPluginSamlResource{}

func NewGatewayPluginSamlResource() resource.Resource {
	return &GatewayPluginSamlResource{}
}

// GatewayPluginSamlResource defines the resource implementation.
type GatewayPluginSamlResource struct {
	client *sdk.Konnect
}

// GatewayPluginSamlResourceModel describes the resource data model.
type GatewayPluginSamlResourceModel struct {
	Config         tfTypes.CreateSamlPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer           `tfsdk:"consumer"`
	ControlPlaneID types.String                   `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                    `tfsdk:"created_at"`
	Enabled        types.Bool                     `tfsdk:"enabled"`
	ID             types.String                   `tfsdk:"id"`
	Protocols      []types.String                 `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer           `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer           `tfsdk:"service"`
	Tags           []types.String                 `tfsdk:"tags"`
}

func (r *GatewayPluginSamlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_saml"
}

func (r *GatewayPluginSamlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginSaml Resource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"anonymous": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `An optional string (consumer UUID or username) value to use as an “anonymous” consumer. If not set, a Kong Consumer must exist for the SAML IdP user credentials, mapping the username format to the Kong Consumer username.`,
					},
					"assertion_consumer_path": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).`,
					},
					"idp_certificate": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The public certificate provided by the IdP. This is used to validate responses from the IdP.  Only include the contents of the certificate. Do not include the header (` + "`" + `BEGIN CERTIFICATE` + "`" + `) and footer (` + "`" + `END CERTIFICATE` + "`" + `) lines.`,
					},
					"idp_sso_url": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"issuer": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The unique identifier of the IdP application. Formatted as a URL containing information about the IdP so the SP can validate that the SAML assertions it receives are issued from the correct IdP.`,
					},
					"nameid_format": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("EmailAddress"),
						Description: `The requested ` + "`" + `NameId` + "`" + ` format. Options available are: - ` + "`" + `Unspecified` + "`" + ` - ` + "`" + `EmailAddress` + "`" + ` - ` + "`" + `Persistent` + "`" + ` - ` + "`" + `Transient` + "`" + `. must be one of ["Unspecified", "EmailAddress", "Persistent", "Transient"]; Default: "EmailAddress"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Unspecified",
								"EmailAddress",
								"Persistent",
								"Transient",
							),
						},
					},
					"request_digest_algorithm": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("SHA256"),
						Description: `The digest algorithm for Authn requests: - ` + "`" + `SHA256` + "`" + ` - ` + "`" + `SHA1` + "`" + `. must be one of ["SHA256", "SHA1"]; Default: "SHA256"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"SHA256",
								"SHA1",
							),
						},
					},
					"request_signature_algorithm": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("SHA256"),
						Description: `The signature algorithm for signing Authn requests. Options available are: - ` + "`" + `SHA256` + "`" + ` - ` + "`" + `SHA384` + "`" + ` - ` + "`" + `SHA512` + "`" + `. must be one of ["SHA256", "SHA384", "SHA512"]; Default: "SHA256"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"SHA256",
								"SHA384",
								"SHA512",
							),
						},
					},
					"request_signing_certificate": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The certificate for signing requests.`,
					},
					"request_signing_key": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The private key for signing requests.  If this parameter is set, requests sent to the IdP are signed.  The ` + "`" + `request_signing_certificate` + "`" + ` parameter must be set as well.`,
					},
					"response_digest_algorithm": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("SHA256"),
						Description: `The algorithm for verifying digest in SAML responses: - ` + "`" + `SHA256` + "`" + ` - ` + "`" + `SHA1` + "`" + `. must be one of ["SHA256", "SHA1"]; Default: "SHA256"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"SHA256",
								"SHA1",
							),
						},
					},
					"response_encryption_key": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The private encryption key required to decrypt encrypted assertions.`,
					},
					"response_signature_algorithm": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("SHA256"),
						Description: `The algorithm for validating signatures in SAML responses. Options available are: - ` + "`" + `SHA256` + "`" + ` - ` + "`" + `SHA384` + "`" + ` - ` + "`" + `SHA512` + "`" + `. must be one of ["SHA256", "SHA384", "SHA512"]; Default: "SHA256"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"SHA256",
								"SHA384",
								"SHA512",
							),
						},
					},
					"session_absolute_timeout": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(86400)),
						Description: `The session cookie absolute timeout in seconds. Specifies how long the session can be used until it is no longer valid. Default: 86400`,
					},
					"session_audience": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("default"),
						Description: `The session audience, for example "my-application". Default: "default"`,
					},
					"session_cookie_domain": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The session cookie domain flag.`,
					},
					"session_cookie_http_only": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(true),
						Description: `Forbids JavaScript from accessing the cookie, for example, through the ` + "`" + `Document.cookie` + "`" + ` property. Default: true`,
					},
					"session_cookie_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("session"),
						Description: `The session cookie name. Default: "session"`,
					},
					"session_cookie_path": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("/"),
						Description: `A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes). Default: "/"`,
					},
					"session_cookie_same_site": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("Lax"),
						Description: `Controls whether a cookie is sent with cross-origin requests, providing some protection against cross-site request forgery attacks. must be one of ["Strict", "Lax", "None", "Default"]; Default: "Lax"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Strict",
								"Lax",
								"None",
								"Default",
							),
						},
					},
					"session_cookie_secure": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The cookie is only sent to the server when a request is made with the https:scheme (except on localhost), and therefore is more resistant to man-in-the-middle attacks.`,
					},
					"session_enforce_same_subject": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `When set to ` + "`" + `true` + "`" + `, audiences are forced to share the same subject. Default: false`,
					},
					"session_hash_storage_key": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `When set to ` + "`" + `true` + "`" + `, the storage key (session ID) is hashed for extra security. Hashing the storage key means it is impossible to decrypt data from the storage without a cookie. Default: false`,
					},
					"session_hash_subject": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `When set to ` + "`" + `true` + "`" + `, the value of subject is hashed before being stored. Only applies when ` + "`" + `session_store_metadata` + "`" + ` is enabled. Default: false`,
					},
					"session_idling_timeout": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(900)),
						Description: `The session cookie idle time in seconds. Default: 900`,
					},
					"session_memcached_host": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("127.0.0.1"),
						Description: `The memcached host. Default: "127.0.0.1"`,
					},
					"session_memcached_port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(11211),
						Description: `An integer representing a port number between 0 and 65535, inclusive. Default: 11211`,
					},
					"session_memcached_prefix": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The memcached session key prefix.`,
					},
					"session_memcached_socket": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The memcached unix socket path.`,
					},
					"session_redis_cluster_max_redirections": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The Redis cluster maximum redirects.`,
					},
					"session_redis_cluster_nodes": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Default:     stringdefault.StaticString("127.0.0.1"),
									Description: `A string representing a host name, such as example.com. Default: "127.0.0.1"`,
								},
								"port": schema.Int64Attribute{
									Computed:    true,
									Optional:    true,
									Default:     int64default.StaticInt64(6379),
									Description: `An integer representing a port number between 0 and 65535, inclusive. Default: 6379`,
								},
							},
						},
						Description: `The Redis cluster node host. Takes an array of host records, with either ` + "`" + `ip` + "`" + ` or ` + "`" + `host` + "`" + `, and ` + "`" + `port` + "`" + ` values.`,
					},
					"session_redis_connect_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The Redis connection timeout in milliseconds.`,
					},
					"session_redis_host": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("127.0.0.1"),
						Description: `The Redis host IP. Default: "127.0.0.1"`,
					},
					"session_redis_password": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Password to use for Redis connection when the ` + "`" + `redis` + "`" + ` session storage is defined. If undefined, no auth commands are sent to Redis. This value is pulled from`,
					},
					"session_redis_port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(6379),
						Description: `An integer representing a port number between 0 and 65535, inclusive. Default: 6379`,
					},
					"session_redis_prefix": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The Redis session key prefix.`,
					},
					"session_redis_read_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The Redis read timeout in milliseconds.`,
					},
					"session_redis_send_timeout": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `The Redis send timeout in milliseconds.`,
					},
					"session_redis_server_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The SNI used for connecting to the Redis server.`,
					},
					"session_redis_socket": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The Redis unix socket path.`,
					},
					"session_redis_ssl": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Use SSL/TLS for the Redis connection. Default: false`,
					},
					"session_redis_ssl_verify": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Verify the Redis server certificate. Default: false`,
					},
					"session_redis_username": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Redis username if the ` + "`" + `redis` + "`" + ` session storage is defined and ACL authentication is desired.If undefined, ACL authentication will not be performed.  This requires Redis v6.0.0+. The username **cannot** be set to ` + "`" + `default` + "`" + `.`,
					},
					"session_remember": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Enables or disables persistent sessions. Default: false`,
					},
					"session_remember_absolute_timeout": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(2592000)),
						Description: `Persistent session absolute timeout in seconds. Default: 2592000`,
					},
					"session_remember_cookie_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("remember"),
						Description: `Persistent session cookie name. Default: "remember"`,
					},
					"session_remember_rolling_timeout": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(604800)),
						Description: `Persistent session rolling timeout in seconds. Default: 604800`,
					},
					"session_request_headers": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"session_response_headers": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"session_rolling_timeout": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(3600)),
						Description: `The session cookie absolute timeout in seconds. Specifies how long the session can be used until it is no longer valid. Default: 3600`,
					},
					"session_secret": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The session secret. This must be a random string of 32 characters from the base64 alphabet (letters, numbers, ` + "`" + `/` + "`" + `, ` + "`" + `_` + "`" + ` and ` + "`" + `+` + "`" + `). It is used as the secret key for encrypting session data as well as state information that is sent to the IdP in the authentication exchange.`,
					},
					"session_storage": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("cookie"),
						Description: `The session storage for session data: - ` + "`" + `cookie` + "`" + `: stores session data with the session cookie. The session cannot be invalidated or revoked without changing the session secret, but is stateless, and doesn't require a database. - ` + "`" + `memcached` + "`" + `: stores session data in memcached - ` + "`" + `redis` + "`" + `: stores session data in Redis. must be one of ["cookie", "memcache", "memcached", "redis"]; Default: "cookie"`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"cookie",
								"memcache",
								"memcached",
								"redis",
							),
						},
					},
					"session_store_metadata": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Configures whether or not session metadata should be stored. This includes information about the active sessions for the ` + "`" + `specific_audience` + "`" + ` belonging to a specific subject. Default: false`,
					},
					"validate_assertion_signature": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(true),
						Description: `Enable signature validation for SAML responses. Default: true`,
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

func (r *GatewayPluginSamlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginSamlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginSamlResourceModel
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
	createSamlPlugin := *data.ToSharedCreateSamlPlugin()
	request := operations.CreateSamlPluginRequest{
		ControlPlaneID:   controlPlaneID,
		CreateSamlPlugin: createSamlPlugin,
	}
	res, err := r.client.Plugins.CreateSamlPlugin(ctx, request)
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
	if res.SamlPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSamlPlugin(res.SamlPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginSamlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginSamlResourceModel
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
	request := operations.GetSamlPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       pluginID,
	}
	res, err := r.client.Plugins.GetSamlPlugin(ctx, request)
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
	if res.SamlPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSamlPlugin(res.SamlPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginSamlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginSamlResourceModel
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
	createSamlPlugin := *data.ToSharedCreateSamlPlugin()
	request := operations.UpdateSamlPluginRequest{
		ControlPlaneID:   controlPlaneID,
		PluginID:         pluginID,
		CreateSamlPlugin: createSamlPlugin,
	}
	res, err := r.client.Plugins.UpdateSamlPlugin(ctx, request)
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
	if res.SamlPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSamlPlugin(res.SamlPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginSamlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginSamlResourceModel
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
	request := operations.DeleteSamlPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       pluginID,
	}
	res, err := r.client.Plugins.DeleteSamlPlugin(ctx, request)
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

func (r *GatewayPluginSamlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_plugin_saml. Reason: composite imports strings not supported.")
}
