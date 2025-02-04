// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_boolplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/boolplanmodifier"
	speakeasy_int64planmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/int64planmodifier"
	speakeasy_listplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/listplanmodifier"
	speakeasy_mapplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/mapplanmodifier"
	speakeasy_numberplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/numberplanmodifier"
	speakeasy_objectplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
	"math/big"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayUpstreamResource{}
var _ resource.ResourceWithImportState = &GatewayUpstreamResource{}

func NewGatewayUpstreamResource() resource.Resource {
	return &GatewayUpstreamResource{}
}

// GatewayUpstreamResource defines the resource implementation.
type GatewayUpstreamResource struct {
	client *sdk.Konnect
}

// GatewayUpstreamResourceModel describes the resource data model.
type GatewayUpstreamResourceModel struct {
	Algorithm              types.String          `tfsdk:"algorithm"`
	ClientCertificate      *tfTypes.ACLConsumer  `tfsdk:"client_certificate"`
	ControlPlaneID         types.String          `tfsdk:"control_plane_id"`
	CreatedAt              types.Int64           `tfsdk:"created_at"`
	HashFallback           types.String          `tfsdk:"hash_fallback"`
	HashFallbackHeader     types.String          `tfsdk:"hash_fallback_header"`
	HashFallbackQueryArg   types.String          `tfsdk:"hash_fallback_query_arg"`
	HashFallbackURICapture types.String          `tfsdk:"hash_fallback_uri_capture"`
	HashOn                 types.String          `tfsdk:"hash_on"`
	HashOnCookie           types.String          `tfsdk:"hash_on_cookie"`
	HashOnCookiePath       types.String          `tfsdk:"hash_on_cookie_path"`
	HashOnHeader           types.String          `tfsdk:"hash_on_header"`
	HashOnQueryArg         types.String          `tfsdk:"hash_on_query_arg"`
	HashOnURICapture       types.String          `tfsdk:"hash_on_uri_capture"`
	Healthchecks           *tfTypes.Healthchecks `tfsdk:"healthchecks"`
	HostHeader             types.String          `tfsdk:"host_header"`
	ID                     types.String          `tfsdk:"id"`
	Name                   types.String          `tfsdk:"name"`
	Slots                  types.Int64           `tfsdk:"slots"`
	Tags                   []types.String        `tfsdk:"tags"`
	UseSrvName             types.Bool            `tfsdk:"use_srv_name"`
}

func (r *GatewayUpstreamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_upstream"
}

func (r *GatewayUpstreamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayUpstream Resource",

		Attributes: map[string]schema.Attribute{
			"algorithm": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Default:     stringdefault.StaticString("round-robin"),
				Description: `Which load balancing algorithm to use. Requires replacement if changed. ; must be one of ["consistent-hashing", "least-connections", "round-robin"]; Default: "round-robin"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"consistent-hashing",
						"least-connections",
						"round-robin",
					),
				},
			},
			"client_certificate": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Optional:    true,
						Description: `Requires replacement if changed. `,
					},
				},
				Description: `If set, the certificate to be used as client certificate while TLS handshaking to the upstream server. Requires replacement if changed. `,
			},
			"control_plane_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed. `,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"hash_fallback": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: `What to use as hashing input if the primary ` + "`" + `hash_on` + "`" + ` does not return a hash (eg. header is missing, or no Consumer identified). Not available if ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `cookie` + "`" + `. Requires replacement if changed. ; must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query_arg", "uri_capture"]; Default: "none"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"none",
						"consumer",
						"ip",
						"header",
						"cookie",
						"path",
						"query_arg",
						"uri_capture",
					),
				},
			},
			"hash_fallback_header": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `The header name to take the value from as hash input. Only required when ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `header` + "`" + `. Requires replacement if changed. `,
			},
			"hash_fallback_query_arg": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `The name of the query string argument to take the value from as hash input. Only required when ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `query_arg` + "`" + `. Requires replacement if changed. `,
			},
			"hash_fallback_uri_capture": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `The name of the route URI capture to take the value from as hash input. Only required when ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `uri_capture` + "`" + `. Requires replacement if changed. `,
			},
			"hash_on": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: `What to use as hashing input. Using ` + "`" + `none` + "`" + ` results in a weighted-round-robin scheme with no hashing. Requires replacement if changed. ; must be one of ["none", "consumer", "ip", "header", "cookie", "path", "query_arg", "uri_capture"]; Default: "none"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"none",
						"consumer",
						"ip",
						"header",
						"cookie",
						"path",
						"query_arg",
						"uri_capture",
					),
				},
			},
			"hash_on_cookie": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `The cookie name to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` or ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `cookie` + "`" + `. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response. Requires replacement if changed. `,
			},
			"hash_on_cookie_path": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Default:     stringdefault.StaticString("/"),
				Description: `The cookie path to set in the response headers. Only required when ` + "`" + `hash_on` + "`" + ` or ` + "`" + `hash_fallback` + "`" + ` is set to ` + "`" + `cookie` + "`" + `. Requires replacement if changed. ; Default: "/"`,
			},
			"hash_on_header": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `The header name to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `header` + "`" + `. Requires replacement if changed. `,
			},
			"hash_on_query_arg": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `The name of the query string argument to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `query_arg` + "`" + `. Requires replacement if changed. `,
			},
			"hash_on_uri_capture": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `The name of the route URI capture to take the value from as hash input. Only required when ` + "`" + `hash_on` + "`" + ` is set to ` + "`" + `uri_capture` + "`" + `. Requires replacement if changed. `,
			},
			"healthchecks": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"concurrency": schema.Int64Attribute{
								Computed: true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.RequiresReplaceIfConfigured(),
									speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Default:     int64default.StaticInt64(10),
								Description: `Requires replacement if changed. ; Default: 10`,
							},
							"headers": schema.MapAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Map{
									mapplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_mapplanmodifier.SuppressDiff(speakeasy_mapplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								ElementType: types.StringType,
								Description: `Requires replacement if changed. `,
								Validators: []validator.Map{
									mapvalidator.ValueStringsAre(validators.IsValidJSON()),
								},
							},
							"healthy": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"http_statuses": schema.ListAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										ElementType: types.Int64Type,
										Description: `Requires replacement if changed. `,
									},
									"interval": schema.NumberAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.Number{
											numberplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_numberplanmodifier.SuppressDiff(speakeasy_numberplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     numberdefault.StaticBigFloat(big.NewFloat(0)),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
									"successes": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplaceIfConfigured(),
											speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     int64default.StaticInt64(0),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
								},
								Description: `Requires replacement if changed. `,
							},
							"http_path": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Default:     stringdefault.StaticString("/"),
								Description: `Requires replacement if changed. ; Default: "/"`,
							},
							"https_sni": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Requires replacement if changed. `,
							},
							"https_verify_certificate": schema.BoolAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Default:     booldefault.StaticBool(true),
								Description: `Requires replacement if changed. ; Default: true`,
							},
							"timeout": schema.NumberAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Number{
									numberplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_numberplanmodifier.SuppressDiff(speakeasy_numberplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Default:     numberdefault.StaticBigFloat(big.NewFloat(1)),
								Description: `Requires replacement if changed. ; Default: 1`,
							},
							"type": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Default:     stringdefault.StaticString("http"),
								Description: `Requires replacement if changed. ; must be one of ["tcp", "http", "https", "grpc", "grpcs"]; Default: "http"`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"tcp",
										"http",
										"https",
										"grpc",
										"grpcs",
									),
								},
							},
							"unhealthy": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"http_failures": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplaceIfConfigured(),
											speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     int64default.StaticInt64(0),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
									"http_statuses": schema.ListAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										ElementType: types.Int64Type,
										Description: `Requires replacement if changed. `,
									},
									"interval": schema.NumberAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.Number{
											numberplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_numberplanmodifier.SuppressDiff(speakeasy_numberplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     numberdefault.StaticBigFloat(big.NewFloat(0)),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
									"tcp_failures": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplaceIfConfigured(),
											speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     int64default.StaticInt64(0),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
									"timeouts": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplaceIfConfigured(),
											speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     int64default.StaticInt64(0),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
								},
								Description: `Requires replacement if changed. `,
							},
						},
						Description: `Requires replacement if changed. `,
					},
					"passive": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"healthy": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"http_statuses": schema.ListAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										ElementType: types.Int64Type,
										Description: `Requires replacement if changed. `,
									},
									"successes": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplaceIfConfigured(),
											speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     int64default.StaticInt64(0),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
								},
								Description: `Requires replacement if changed. `,
							},
							"type": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Default:     stringdefault.StaticString("http"),
								Description: `Requires replacement if changed. ; must be one of ["tcp", "http", "https", "grpc", "grpcs"]; Default: "http"`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"tcp",
										"http",
										"https",
										"grpc",
										"grpcs",
									),
								},
							},
							"unhealthy": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"http_failures": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplaceIfConfigured(),
											speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     int64default.StaticInt64(0),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
									"http_statuses": schema.ListAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
										},
										Optional:    true,
										ElementType: types.Int64Type,
										Description: `Requires replacement if changed. `,
									},
									"tcp_failures": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplaceIfConfigured(),
											speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     int64default.StaticInt64(0),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
									"timeouts": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplaceIfConfigured(),
											speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
										},
										Optional:    true,
										Default:     int64default.StaticInt64(0),
										Description: `Requires replacement if changed. ; Default: 0`,
									},
								},
								Description: `Requires replacement if changed. `,
							},
						},
						Description: `Requires replacement if changed. `,
					},
					"threshold": schema.NumberAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Number{
							numberplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_numberplanmodifier.SuppressDiff(speakeasy_numberplanmodifier.ExplicitSuppress),
						},
						Optional:    true,
						Default:     numberdefault.StaticBigFloat(big.NewFloat(0)),
						Description: `Requires replacement if changed. ; Default: 0`,
					},
				},
				Description: `Requires replacement if changed. `,
			},
			"host_header": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `The hostname to be used as ` + "`" + `Host` + "`" + ` header when proxying requests through Kong. Requires replacement if changed. `,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `ID of the Upstream to lookup`,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `This is a hostname, which must be equal to the ` + "`" + `host` + "`" + ` of a Service. Requires replacement if changed. `,
			},
			"slots": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
					speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Default:     int64default.StaticInt64(10000),
				Description: `The number of slots in the load balancer algorithm. If ` + "`" + `algorithm` + "`" + ` is set to ` + "`" + `round-robin` + "`" + `, this setting determines the maximum number of slots. If ` + "`" + `algorithm` + "`" + ` is set to ` + "`" + `consistent-hashing` + "`" + `, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range ` + "`" + `10` + "`" + `-` + "`" + `65536` + "`" + `. Requires replacement if changed. ; Default: 10000`,
			},
			"tags": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Upstream for grouping and filtering. Requires replacement if changed. `,
			},
			"use_srv_name": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Default:     booldefault.StaticBool(false),
				Description: `If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream ` + "`" + `Host` + "`" + `. Requires replacement if changed. ; Default: false`,
			},
		},
	}
}

func (r *GatewayUpstreamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayUpstreamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayUpstreamResourceModel
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
	createUpstream := *data.ToSharedCreateUpstream()
	request := operations.CreateUpstreamRequest{
		ControlPlaneID: controlPlaneID,
		CreateUpstream: createUpstream,
	}
	res, err := r.client.Upstreams.CreateUpstream(ctx, request)
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
	if res.Upstream == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUpstream(res.Upstream)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayUpstreamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayUpstreamResourceModel
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
	upstreamID := data.ID.ValueString()
	request := operations.GetUpstreamRequest{
		ControlPlaneID: controlPlaneID,
		UpstreamID:     upstreamID,
	}
	res, err := r.client.Upstreams.GetUpstream(ctx, request)
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
	if res.Upstream == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUpstream(res.Upstream)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayUpstreamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayUpstreamResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayUpstreamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayUpstreamResourceModel
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
	upstreamID := data.ID.ValueString()
	request := operations.DeleteUpstreamRequest{
		ControlPlaneID: controlPlaneID,
		UpstreamID:     upstreamID,
	}
	res, err := r.client.Upstreams.DeleteUpstream(ctx, request)
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

func (r *GatewayUpstreamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_upstream. Reason: composite imports strings not supported.")
}
