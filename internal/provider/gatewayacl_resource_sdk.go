// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayACLResourceModel) ToSharedCreateACLWithoutParents() *shared.CreateACLWithoutParents {
	group := new(string)
	if !r.Group.IsUnknown() && !r.Group.IsNull() {
		*group = r.Group.ValueString()
	} else {
		group = nil
	}
	var tags []string = nil
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.CreateACLWithoutParents{
		Group: group,
		Tags:  tags,
	}
	return &out
}

func (r *GatewayACLResourceModel) RefreshFromSharedACL(resp *shared.ACL) {
	if resp != nil {
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Group = types.StringPointerValue(resp.Group)
		r.ID = types.StringPointerValue(resp.ID)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}
