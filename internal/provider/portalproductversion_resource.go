// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PortalProductVersionResource{}
var _ resource.ResourceWithImportState = &PortalProductVersionResource{}

func NewPortalProductVersionResource() resource.Resource {
	return &PortalProductVersionResource{}
}

// PortalProductVersionResource defines the resource implementation.
type PortalProductVersionResource struct {
	client *sdk.Konnect
}

// PortalProductVersionResourceModel describes the resource data model.
type PortalProductVersionResourceModel struct {
	ApplicationRegistrationEnabled types.Bool             `tfsdk:"application_registration_enabled"`
	AuthStrategies                 []tfTypes.AuthStrategy `tfsdk:"auth_strategies"`
	AuthStrategyIds                []types.String         `tfsdk:"auth_strategy_ids"`
	AutoApproveRegistration        types.Bool             `tfsdk:"auto_approve_registration"`
	CreatedAt                      types.String           `tfsdk:"created_at"`
	Deprecated                     types.Bool             `tfsdk:"deprecated"`
	ID                             types.String           `tfsdk:"id"`
	PortalID                       types.String           `tfsdk:"portal_id"`
	ProductVersionID               types.String           `tfsdk:"product_version_id"`
	PublishStatus                  types.String           `tfsdk:"publish_status"`
	UpdatedAt                      types.String           `tfsdk:"updated_at"`
}

func (r *PortalProductVersionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal_product_version"
}

func (r *PortalProductVersionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PortalProductVersion Resource",

		Attributes: map[string]schema.Attribute{
			"application_registration_enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the application registration on this portal for the api product version is enabled`,
			},
			"auth_strategies": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"client_credentials": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"auth_methods": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"credential_type": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["client_credentials", "self_managed_client_credentials"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"client_credentials",
											"self_managed_client_credentials",
										),
									},
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `The Application Auth Strategy ID.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
							Description: `Client Credential Auth strategy that the application uses.`,
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.Expressions{
									path.MatchRelative().AtParent().AtName("key_auth"),
								}...),
							},
						},
						"key_auth": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"credential_type": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["key_auth"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"key_auth",
										),
									},
								},
								"id": schema.StringAttribute{
									Computed:    true,
									Description: `The Application Auth Strategy ID.`,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
							},
							Description: `KeyAuth Auth strategy that the application uses.`,
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.Expressions{
									path.MatchRelative().AtParent().AtName("client_credentials"),
								}...),
							},
						},
					},
					Validators: []validator.Object{
						validators.ExactlyOneChild(),
					},
				},
				Description: `A list of authentication strategies`,
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
			},
			"auth_strategy_ids": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of authentication strategy IDs`,
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
			},
			"auto_approve_registration": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the application registration auto approval on this portal for the api product version is enabled`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"deprecated": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the api product version on the portal is deprecated`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Contains a unique identifier used by the API for this resource.`,
			},
			"portal_id": schema.StringAttribute{
				Required:    true,
				Description: `portal identifier`,
			},
			"product_version_id": schema.StringAttribute{
				Required:    true,
				Description: `API product version identifier`,
			},
			"publish_status": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Publication status of the API product version on the portal. must be one of ["published", "unpublished"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"published",
						"unpublished",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity update date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *PortalProductVersionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PortalProductVersionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PortalProductVersionResourceModel
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

	productVersionID := data.ProductVersionID.ValueString()
	portalID := data.PortalID.ValueString()
	updatePortalProductVersionPayload := *data.ToSharedUpdatePortalProductVersionPayload()
	request := operations.UpdatePortalProductVersionRequest{
		ProductVersionID:                  productVersionID,
		PortalID:                          portalID,
		UpdatePortalProductVersionPayload: updatePortalProductVersionPayload,
	}
	res, err := r.client.PortalProductVersions.UpdatePortalProductVersion(ctx, request)
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
	if res.PortalProductVersion == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPortalProductVersion(res.PortalProductVersion)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	productVersionId1 := data.ProductVersionID.ValueString()
	portalId1 := data.PortalID.ValueString()
	request1 := operations.GetPortalProductVersionRequest{
		ProductVersionID: productVersionId1,
		PortalID:         portalId1,
	}
	res1, err := r.client.PortalProductVersions.GetPortalProductVersion(ctx, request1)
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
	if res1.PortalProductVersion == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedPortalProductVersion(res1.PortalProductVersion)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalProductVersionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PortalProductVersionResourceModel
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

	productVersionID := data.ProductVersionID.ValueString()
	portalID := data.PortalID.ValueString()
	request := operations.GetPortalProductVersionRequest{
		ProductVersionID: productVersionID,
		PortalID:         portalID,
	}
	res, err := r.client.PortalProductVersions.GetPortalProductVersion(ctx, request)
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
	if res.PortalProductVersion == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPortalProductVersion(res.PortalProductVersion)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalProductVersionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PortalProductVersionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	productVersionID := data.ProductVersionID.ValueString()
	portalID := data.PortalID.ValueString()
	updatePortalProductVersionPayload := *data.ToSharedUpdatePortalProductVersionPayload()
	request := operations.UpdatePortalProductVersionRequest{
		ProductVersionID:                  productVersionID,
		PortalID:                          portalID,
		UpdatePortalProductVersionPayload: updatePortalProductVersionPayload,
	}
	res, err := r.client.PortalProductVersions.UpdatePortalProductVersion(ctx, request)
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
	if res.PortalProductVersion == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPortalProductVersion(res.PortalProductVersion)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	productVersionId1 := data.ProductVersionID.ValueString()
	portalId1 := data.PortalID.ValueString()
	request1 := operations.GetPortalProductVersionRequest{
		ProductVersionID: productVersionId1,
		PortalID:         portalId1,
	}
	res1, err := r.client.PortalProductVersions.GetPortalProductVersion(ctx, request1)
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
	if res1.PortalProductVersion == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedPortalProductVersion(res1.PortalProductVersion)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalProductVersionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PortalProductVersionResourceModel
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

	productVersionID := data.ProductVersionID.ValueString()
	portalID := data.PortalID.ValueString()
	request := operations.DeletePortalProductVersionRequest{
		ProductVersionID: productVersionID,
		PortalID:         portalID,
	}
	res, err := r.client.PortalProductVersions.DeletePortalProductVersion(ctx, request)
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

func (r *PortalProductVersionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource portal_product_version. Reason: composite imports strings not supported.")
}
