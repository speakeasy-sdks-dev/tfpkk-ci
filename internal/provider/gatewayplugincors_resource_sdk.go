// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginCORSResourceModel) ToSharedCreateCORSPlugin() *shared.CreateCORSPlugin {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []shared.CreateCORSPluginProtocols = nil
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateCORSPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = nil
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateCORSPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateCORSPluginConsumer{
			ID: id,
		}
	}
	var route *shared.CreateCORSPluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CreateCORSPluginRoute{
			ID: id1,
		}
	}
	var service *shared.CreateCORSPluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CreateCORSPluginService{
			ID: id2,
		}
	}
	var origins []string = nil
	for _, originsItem := range r.Config.Origins {
		origins = append(origins, originsItem.ValueString())
	}
	var headers []string = nil
	for _, headersItem := range r.Config.Headers {
		headers = append(headers, headersItem.ValueString())
	}
	var exposedHeaders []string = nil
	for _, exposedHeadersItem := range r.Config.ExposedHeaders {
		exposedHeaders = append(exposedHeaders, exposedHeadersItem.ValueString())
	}
	var methods []shared.Methods = nil
	for _, methodsItem := range r.Config.Methods {
		methods = append(methods, shared.Methods(methodsItem.ValueString()))
	}
	maxAge := new(float64)
	if !r.Config.MaxAge.IsUnknown() && !r.Config.MaxAge.IsNull() {
		*maxAge, _ = r.Config.MaxAge.ValueBigFloat().Float64()
	} else {
		maxAge = nil
	}
	credentials := new(bool)
	if !r.Config.Credentials.IsUnknown() && !r.Config.Credentials.IsNull() {
		*credentials = r.Config.Credentials.ValueBool()
	} else {
		credentials = nil
	}
	privateNetwork := new(bool)
	if !r.Config.PrivateNetwork.IsUnknown() && !r.Config.PrivateNetwork.IsNull() {
		*privateNetwork = r.Config.PrivateNetwork.ValueBool()
	} else {
		privateNetwork = nil
	}
	preflightContinue := new(bool)
	if !r.Config.PreflightContinue.IsUnknown() && !r.Config.PreflightContinue.IsNull() {
		*preflightContinue = r.Config.PreflightContinue.ValueBool()
	} else {
		preflightContinue = nil
	}
	config := shared.CreateCORSPluginConfig{
		Origins:           origins,
		Headers:           headers,
		ExposedHeaders:    exposedHeaders,
		Methods:           methods,
		MaxAge:            maxAge,
		Credentials:       credentials,
		PrivateNetwork:    privateNetwork,
		PreflightContinue: preflightContinue,
	}
	out := shared.CreateCORSPlugin{
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

func (r *GatewayPluginCORSResourceModel) RefreshFromSharedCORSPlugin(resp *shared.CORSPlugin) {
	if resp != nil {
		r.Config.Credentials = types.BoolPointerValue(resp.Config.Credentials)
		r.Config.ExposedHeaders = []types.String{}
		for _, v := range resp.Config.ExposedHeaders {
			r.Config.ExposedHeaders = append(r.Config.ExposedHeaders, types.StringValue(v))
		}
		r.Config.Headers = []types.String{}
		for _, v := range resp.Config.Headers {
			r.Config.Headers = append(r.Config.Headers, types.StringValue(v))
		}
		if resp.Config.MaxAge != nil {
			r.Config.MaxAge = types.NumberValue(big.NewFloat(float64(*resp.Config.MaxAge)))
		} else {
			r.Config.MaxAge = types.NumberNull()
		}
		r.Config.Methods = []types.String{}
		for _, v := range resp.Config.Methods {
			r.Config.Methods = append(r.Config.Methods, types.StringValue(string(v)))
		}
		r.Config.Origins = []types.String{}
		for _, v := range resp.Config.Origins {
			r.Config.Origins = append(r.Config.Origins, types.StringValue(v))
		}
		r.Config.PreflightContinue = types.BoolPointerValue(resp.Config.PreflightContinue)
		r.Config.PrivateNetwork = types.BoolPointerValue(resp.Config.PrivateNetwork)
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
