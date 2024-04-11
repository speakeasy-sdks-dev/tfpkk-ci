// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PortalAuthResource{}
var _ resource.ResourceWithImportState = &PortalAuthResource{}

func NewPortalAuthResource() resource.Resource {
	return &PortalAuthResource{}
}

// PortalAuthResource defines the resource implementation.
type PortalAuthResource struct {
	client *sdk.Konnect
}

// PortalAuthResourceModel describes the resource data model.
type PortalAuthResourceModel struct {
	BasicAuthEnabled       types.Bool                   `tfsdk:"basic_auth_enabled"`
	KonnectMappingEnabled  types.Bool                   `tfsdk:"konnect_mapping_enabled"`
	OidcAuthEnabled        types.Bool                   `tfsdk:"oidc_auth_enabled"`
	OidcClaimMappings      *tfTypes.PortalClaimMappings `tfsdk:"oidc_claim_mappings"`
	OidcClientID           types.String                 `tfsdk:"oidc_client_id"`
	OidcClientSecret       types.String                 `tfsdk:"oidc_client_secret"`
	OidcConfig             *tfTypes.PortalOIDCConfig    `tfsdk:"oidc_config"`
	OidcIssuer             types.String                 `tfsdk:"oidc_issuer"`
	OidcScopes             []types.String               `tfsdk:"oidc_scopes"`
	OidcTeamMappingEnabled types.Bool                   `tfsdk:"oidc_team_mapping_enabled"`
	PortalID               types.String                 `tfsdk:"portal_id"`
}

func (r *PortalAuthResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal_auth"
}

func (r *PortalAuthResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PortalAuth Resource",

		Attributes: map[string]schema.Attribute{
			"basic_auth_enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The organization has basic auth enabled.`,
			},
			"konnect_mapping_enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether a Konnect Identity Admin assigns teams to a developer.`,
			},
			"oidc_auth_enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The organization has OIDC disabled.`,
			},
			"oidc_claim_mappings": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("email"),
						Description: `Default: "email"`,
					},
					"groups": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("groups"),
						Description: `Default: "groups"`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("name"),
						Description: `Default: "name"`,
					},
				},
				Description: `Mappings from a portal developer atribute to an Identity Provider claim.`,
			},
			"oidc_client_id": schema.StringAttribute{
				Optional: true,
			},
			"oidc_client_secret": schema.StringAttribute{
				Optional: true,
			},
			"oidc_config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"claim_mappings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"email": schema.StringAttribute{
								Computed: true,
							},
							"groups": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
						Description: `Mappings from a portal developer atribute to an Identity Provider claim.`,
					},
					"client_id": schema.StringAttribute{
						Computed: true,
					},
					"issuer": schema.StringAttribute{
						Computed: true,
					},
					"scopes": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
				Description: `Configuration properties for an OpenID Connect Identity Provider.`,
			},
			"oidc_issuer": schema.StringAttribute{
				Optional: true,
			},
			"oidc_scopes": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"oidc_team_mapping_enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether IdP groups determine the Konnect Portal teams a developer has.`,
			},
			"portal_id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the portal.`,
			},
		},
	}
}

func (r *PortalAuthResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PortalAuthResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PortalAuthResourceModel
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

	portalID := data.PortalID.ValueString()
	portalAuthenticationSettingsUpdateRequest := data.ToSharedPortalAuthenticationSettingsUpdateRequest()
	request := operations.UpdatePortalAuthenticationSettingsRequest{
		PortalID: portalID,
		PortalAuthenticationSettingsUpdateRequest: portalAuthenticationSettingsUpdateRequest,
	}
	res, err := r.client.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, request)
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
	if res.PortalAuthenticationSettingsResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPortalAuthenticationSettingsResponse(res.PortalAuthenticationSettingsResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	portalId1 := data.PortalID.ValueString()
	request1 := operations.GetPortalAuthenticationSettingsRequest{
		PortalID: portalId1,
	}
	res1, err := r.client.PortalAuthSettings.GetPortalAuthenticationSettings(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.PortalAuthenticationSettingsResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedPortalAuthenticationSettingsResponse(res1.PortalAuthenticationSettingsResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalAuthResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PortalAuthResourceModel
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

	portalID := data.PortalID.ValueString()
	request := operations.GetPortalAuthenticationSettingsRequest{
		PortalID: portalID,
	}
	res, err := r.client.PortalAuthSettings.GetPortalAuthenticationSettings(ctx, request)
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
	if res.PortalAuthenticationSettingsResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPortalAuthenticationSettingsResponse(res.PortalAuthenticationSettingsResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalAuthResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PortalAuthResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	portalID := data.PortalID.ValueString()
	portalAuthenticationSettingsUpdateRequest := data.ToSharedPortalAuthenticationSettingsUpdateRequest()
	request := operations.UpdatePortalAuthenticationSettingsRequest{
		PortalID: portalID,
		PortalAuthenticationSettingsUpdateRequest: portalAuthenticationSettingsUpdateRequest,
	}
	res, err := r.client.PortalAuthSettings.UpdatePortalAuthenticationSettings(ctx, request)
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
	if res.PortalAuthenticationSettingsResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPortalAuthenticationSettingsResponse(res.PortalAuthenticationSettingsResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	portalId1 := data.PortalID.ValueString()
	request1 := operations.GetPortalAuthenticationSettingsRequest{
		PortalID: portalId1,
	}
	res1, err := r.client.PortalAuthSettings.GetPortalAuthenticationSettings(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.PortalAuthenticationSettingsResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedPortalAuthenticationSettingsResponse(res1.PortalAuthenticationSettingsResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalAuthResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PortalAuthResourceModel
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

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *PortalAuthResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("portal_id"), req.ID)...)
}
