// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginFileLogResourceModel) ToSharedCreateFileLogPlugin() *shared.CreateFileLogPlugin {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []shared.CreateFileLogPluginProtocols = nil
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateFileLogPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = nil
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateFileLogPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateFileLogPluginConsumer{
			ID: id,
		}
	}
	var route *shared.CreateFileLogPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CreateFileLogPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CreateFileLogPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CreateFileLogPluginService{
			ID: id2,
		}
	}
	path := new(string)
	if !r.Config.Path.IsUnknown() && !r.Config.Path.IsNull() {
		*path = r.Config.Path.ValueString()
	} else {
		path = nil
	}
	reopen := new(bool)
	if !r.Config.Reopen.IsUnknown() && !r.Config.Reopen.IsNull() {
		*reopen = r.Config.Reopen.ValueBool()
	} else {
		reopen = nil
	}
	customFieldsByLua := make(map[string]interface{})
	for customFieldsByLuaKey, customFieldsByLuaValue := range r.Config.CustomFieldsByLua {
		var customFieldsByLuaInst interface{}
		_ = json.Unmarshal([]byte(customFieldsByLuaValue.ValueString()), &customFieldsByLuaInst)
		customFieldsByLua[customFieldsByLuaKey] = customFieldsByLuaInst
	}
	config := shared.CreateFileLogPluginConfig{
		Path:              path,
		Reopen:            reopen,
		CustomFieldsByLua: customFieldsByLua,
	}
	out := shared.CreateFileLogPlugin{
		Enabled:   enabled,
		Protocols: protocols,
		Tags:      tags,
		Consumer:  consumer,
		Route:     route,
		Service:   service,
		Config:    config,
	}
	return &out
}

func (r *GatewayPluginFileLogResourceModel) RefreshFromSharedFileLogPlugin(resp *shared.FileLogPlugin) {
	if resp != nil {
		if len(resp.Config.CustomFieldsByLua) > 0 {
			r.Config.CustomFieldsByLua = make(map[string]types.String)
			for key, value := range resp.Config.CustomFieldsByLua {
				result, _ := json.Marshal(value)
				r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
			}
		}
		r.Config.Path = types.StringPointerValue(resp.Config.Path)
		r.Config.Reopen = types.BoolPointerValue(resp.Config.Reopen)
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
	}
}
