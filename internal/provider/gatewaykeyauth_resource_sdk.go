// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayKeyAuthResourceModel) ToSharedCreateKeyAuthWithoutParents() *shared.CreateKeyAuthWithoutParents {
	key := new(string)
	if !r.Key.IsUnknown() && !r.Key.IsNull() {
		*key = r.Key.ValueString()
	} else {
		key = nil
	}
	var tags []string = nil
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.CreateKeyAuthWithoutParents{
		Key:  key,
		Tags: tags,
	}
	return &out
}

func (r *GatewayKeyAuthResourceModel) RefreshFromSharedKeyAuth(resp *shared.KeyAuth) {
	if resp != nil {
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.ID = types.StringPointerValue(resp.ID)
		r.Key = types.StringPointerValue(resp.Key)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}
