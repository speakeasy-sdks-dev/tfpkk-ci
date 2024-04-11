// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginOauth2ResourceModel) ToSharedCreateOauth2Plugin() *shared.CreateOauth2Plugin {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []shared.CreateOauth2PluginProtocols = nil
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateOauth2PluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = nil
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateOauth2PluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateOauth2PluginConsumer{
			ID: id,
		}
	}
	var route *shared.CreateOauth2PluginRoute
	if r.Route != nil {
		id1 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id1 = r.Route.ID.ValueString()
		} else {
			id1 = nil
		}
		route = &shared.CreateOauth2PluginRoute{
			ID: id1,
		}
	}
	var service *shared.CreateOauth2PluginService
	if r.Service != nil {
		id2 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id2 = r.Service.ID.ValueString()
		} else {
			id2 = nil
		}
		service = &shared.CreateOauth2PluginService{
			ID: id2,
		}
	}
	var scopes []string = nil
	for _, scopesItem := range r.Config.Scopes {
		scopes = append(scopes, scopesItem.ValueString())
	}
	mandatoryScope := new(bool)
	if !r.Config.MandatoryScope.IsUnknown() && !r.Config.MandatoryScope.IsNull() {
		*mandatoryScope = r.Config.MandatoryScope.ValueBool()
	} else {
		mandatoryScope = nil
	}
	provisionKey := new(string)
	if !r.Config.ProvisionKey.IsUnknown() && !r.Config.ProvisionKey.IsNull() {
		*provisionKey = r.Config.ProvisionKey.ValueString()
	} else {
		provisionKey = nil
	}
	tokenExpiration := new(float64)
	if !r.Config.TokenExpiration.IsUnknown() && !r.Config.TokenExpiration.IsNull() {
		*tokenExpiration, _ = r.Config.TokenExpiration.ValueBigFloat().Float64()
	} else {
		tokenExpiration = nil
	}
	enableAuthorizationCode := new(bool)
	if !r.Config.EnableAuthorizationCode.IsUnknown() && !r.Config.EnableAuthorizationCode.IsNull() {
		*enableAuthorizationCode = r.Config.EnableAuthorizationCode.ValueBool()
	} else {
		enableAuthorizationCode = nil
	}
	enableImplicitGrant := new(bool)
	if !r.Config.EnableImplicitGrant.IsUnknown() && !r.Config.EnableImplicitGrant.IsNull() {
		*enableImplicitGrant = r.Config.EnableImplicitGrant.ValueBool()
	} else {
		enableImplicitGrant = nil
	}
	enableClientCredentials := new(bool)
	if !r.Config.EnableClientCredentials.IsUnknown() && !r.Config.EnableClientCredentials.IsNull() {
		*enableClientCredentials = r.Config.EnableClientCredentials.ValueBool()
	} else {
		enableClientCredentials = nil
	}
	enablePasswordGrant := new(bool)
	if !r.Config.EnablePasswordGrant.IsUnknown() && !r.Config.EnablePasswordGrant.IsNull() {
		*enablePasswordGrant = r.Config.EnablePasswordGrant.ValueBool()
	} else {
		enablePasswordGrant = nil
	}
	hideCredentials := new(bool)
	if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
		*hideCredentials = r.Config.HideCredentials.ValueBool()
	} else {
		hideCredentials = nil
	}
	acceptHTTPIfAlreadyTerminated := new(bool)
	if !r.Config.AcceptHTTPIfAlreadyTerminated.IsUnknown() && !r.Config.AcceptHTTPIfAlreadyTerminated.IsNull() {
		*acceptHTTPIfAlreadyTerminated = r.Config.AcceptHTTPIfAlreadyTerminated.ValueBool()
	} else {
		acceptHTTPIfAlreadyTerminated = nil
	}
	anonymous := new(string)
	if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
		*anonymous = r.Config.Anonymous.ValueString()
	} else {
		anonymous = nil
	}
	globalCredentials := new(bool)
	if !r.Config.GlobalCredentials.IsUnknown() && !r.Config.GlobalCredentials.IsNull() {
		*globalCredentials = r.Config.GlobalCredentials.ValueBool()
	} else {
		globalCredentials = nil
	}
	authHeaderName := new(string)
	if !r.Config.AuthHeaderName.IsUnknown() && !r.Config.AuthHeaderName.IsNull() {
		*authHeaderName = r.Config.AuthHeaderName.ValueString()
	} else {
		authHeaderName = nil
	}
	refreshTokenTTL := new(float64)
	if !r.Config.RefreshTokenTTL.IsUnknown() && !r.Config.RefreshTokenTTL.IsNull() {
		*refreshTokenTTL, _ = r.Config.RefreshTokenTTL.ValueBigFloat().Float64()
	} else {
		refreshTokenTTL = nil
	}
	reuseRefreshToken := new(bool)
	if !r.Config.ReuseRefreshToken.IsUnknown() && !r.Config.ReuseRefreshToken.IsNull() {
		*reuseRefreshToken = r.Config.ReuseRefreshToken.ValueBool()
	} else {
		reuseRefreshToken = nil
	}
	persistentRefreshToken := new(bool)
	if !r.Config.PersistentRefreshToken.IsUnknown() && !r.Config.PersistentRefreshToken.IsNull() {
		*persistentRefreshToken = r.Config.PersistentRefreshToken.ValueBool()
	} else {
		persistentRefreshToken = nil
	}
	pkce := new(shared.Pkce)
	if !r.Config.Pkce.IsUnknown() && !r.Config.Pkce.IsNull() {
		*pkce = shared.Pkce(r.Config.Pkce.ValueString())
	} else {
		pkce = nil
	}
	config := shared.CreateOauth2PluginConfig{
		Scopes:                        scopes,
		MandatoryScope:                mandatoryScope,
		ProvisionKey:                  provisionKey,
		TokenExpiration:               tokenExpiration,
		EnableAuthorizationCode:       enableAuthorizationCode,
		EnableImplicitGrant:           enableImplicitGrant,
		EnableClientCredentials:       enableClientCredentials,
		EnablePasswordGrant:           enablePasswordGrant,
		HideCredentials:               hideCredentials,
		AcceptHTTPIfAlreadyTerminated: acceptHTTPIfAlreadyTerminated,
		Anonymous:                     anonymous,
		GlobalCredentials:             globalCredentials,
		AuthHeaderName:                authHeaderName,
		RefreshTokenTTL:               refreshTokenTTL,
		ReuseRefreshToken:             reuseRefreshToken,
		PersistentRefreshToken:        persistentRefreshToken,
		Pkce:                          pkce,
	}
	out := shared.CreateOauth2Plugin{
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

func (r *GatewayPluginOauth2ResourceModel) RefreshFromSharedOauth2Plugin(resp *shared.Oauth2Plugin) {
	if resp != nil {
		r.Config.AcceptHTTPIfAlreadyTerminated = types.BoolPointerValue(resp.Config.AcceptHTTPIfAlreadyTerminated)
		r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
		r.Config.AuthHeaderName = types.StringPointerValue(resp.Config.AuthHeaderName)
		r.Config.EnableAuthorizationCode = types.BoolPointerValue(resp.Config.EnableAuthorizationCode)
		r.Config.EnableClientCredentials = types.BoolPointerValue(resp.Config.EnableClientCredentials)
		r.Config.EnableImplicitGrant = types.BoolPointerValue(resp.Config.EnableImplicitGrant)
		r.Config.EnablePasswordGrant = types.BoolPointerValue(resp.Config.EnablePasswordGrant)
		r.Config.GlobalCredentials = types.BoolPointerValue(resp.Config.GlobalCredentials)
		r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
		r.Config.MandatoryScope = types.BoolPointerValue(resp.Config.MandatoryScope)
		r.Config.PersistentRefreshToken = types.BoolPointerValue(resp.Config.PersistentRefreshToken)
		if resp.Config.Pkce != nil {
			r.Config.Pkce = types.StringValue(string(*resp.Config.Pkce))
		} else {
			r.Config.Pkce = types.StringNull()
		}
		r.Config.ProvisionKey = types.StringPointerValue(resp.Config.ProvisionKey)
		if resp.Config.RefreshTokenTTL != nil {
			r.Config.RefreshTokenTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.RefreshTokenTTL)))
		} else {
			r.Config.RefreshTokenTTL = types.NumberNull()
		}
		r.Config.ReuseRefreshToken = types.BoolPointerValue(resp.Config.ReuseRefreshToken)
		r.Config.Scopes = []types.String{}
		for _, v := range resp.Config.Scopes {
			r.Config.Scopes = append(r.Config.Scopes, types.StringValue(v))
		}
		if resp.Config.TokenExpiration != nil {
			r.Config.TokenExpiration = types.NumberValue(big.NewFloat(float64(*resp.Config.TokenExpiration)))
		} else {
			r.Config.TokenExpiration = types.NumberNull()
		}
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
