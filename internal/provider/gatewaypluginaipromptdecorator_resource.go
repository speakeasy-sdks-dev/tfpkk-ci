// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginAIPromptDecoratorResource{}
var _ resource.ResourceWithImportState = &GatewayPluginAIPromptDecoratorResource{}

func NewGatewayPluginAIPromptDecoratorResource() resource.Resource {
	return &GatewayPluginAIPromptDecoratorResource{}
}

// GatewayPluginAIPromptDecoratorResource defines the resource implementation.
type GatewayPluginAIPromptDecoratorResource struct {
	client *sdk.Konnect
}

// GatewayPluginAIPromptDecoratorResourceModel describes the resource data model.
type GatewayPluginAIPromptDecoratorResourceModel struct {
	Config         tfTypes.CreateAIPromptDecoratorPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                        `tfsdk:"consumer"`
	ControlPlaneID types.String                                `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                                 `tfsdk:"created_at"`
	Enabled        types.Bool                                  `tfsdk:"enabled"`
	ID             types.String                                `tfsdk:"id"`
	Protocols      []types.String                              `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                        `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                        `tfsdk:"service"`
	Tags           []types.String                              `tfsdk:"tags"`
}

func (r *GatewayPluginAIPromptDecoratorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_ai_prompt_decorator"
}

func (r *GatewayPluginAIPromptDecoratorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginAIPromptDecorator Resource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"prompts": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"append": schema.ListNestedAttribute{
								Computed: true,
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"content": schema.StringAttribute{
											Computed: true,
											Optional: true,
										},
										"role": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Default:     stringdefault.StaticString("system"),
											Description: `must be one of ["system", "assistant", "user"]; Default: "system"`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"system",
													"assistant",
													"user",
												),
											},
										},
									},
								},
								Description: `Insert chat messages at the end of the chat message array. This array preserves exact order when adding messages.`,
							},
							"prepend": schema.ListNestedAttribute{
								Computed: true,
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"content": schema.StringAttribute{
											Computed: true,
											Optional: true,
										},
										"role": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Default:     stringdefault.StaticString("system"),
											Description: `must be one of ["system", "assistant", "user"]; Default: "system"`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"system",
													"assistant",
													"user",
												),
											},
										},
									},
								},
								Description: `Insert chat messages at the beginning of the chat message array. This array preserves exact order when adding messages.`,
							},
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

func (r *GatewayPluginAIPromptDecoratorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginAIPromptDecoratorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginAIPromptDecoratorResourceModel
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
	createAIPromptDecoratorPlugin := *data.ToSharedCreateAIPromptDecoratorPlugin()
	request := operations.CreateAipromptdecoratorPluginRequest{
		ControlPlaneID:                controlPlaneID,
		CreateAIPromptDecoratorPlugin: createAIPromptDecoratorPlugin,
	}
	res, err := r.client.Plugins.CreateAipromptdecoratorPlugin(ctx, request)
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
	if res.AIPromptDecoratorPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAIPromptDecoratorPlugin(res.AIPromptDecoratorPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginAIPromptDecoratorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginAIPromptDecoratorResourceModel
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
	request := operations.GetAipromptdecoratorPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       pluginID,
	}
	res, err := r.client.Plugins.GetAipromptdecoratorPlugin(ctx, request)
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
	if res.AIPromptDecoratorPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAIPromptDecoratorPlugin(res.AIPromptDecoratorPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginAIPromptDecoratorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginAIPromptDecoratorResourceModel
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
	createAIPromptDecoratorPlugin := *data.ToSharedCreateAIPromptDecoratorPlugin()
	request := operations.UpdateAipromptdecoratorPluginRequest{
		ControlPlaneID:                controlPlaneID,
		PluginID:                      pluginID,
		CreateAIPromptDecoratorPlugin: createAIPromptDecoratorPlugin,
	}
	res, err := r.client.Plugins.UpdateAipromptdecoratorPlugin(ctx, request)
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
	if res.AIPromptDecoratorPlugin == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedAIPromptDecoratorPlugin(res.AIPromptDecoratorPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginAIPromptDecoratorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginAIPromptDecoratorResourceModel
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
	request := operations.DeleteAipromptdecoratorPluginRequest{
		ControlPlaneID: controlPlaneID,
		PluginID:       pluginID,
	}
	res, err := r.client.Plugins.DeleteAipromptdecoratorPlugin(ctx, request)
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

func (r *GatewayPluginAIPromptDecoratorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_plugin_ai_prompt_decorator. Reason: composite imports strings not supported.")
}
