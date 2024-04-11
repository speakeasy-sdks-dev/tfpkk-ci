// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateResponsetransformeradvancedPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// Description of the Plugin
	CreateResponseTransformerAdvancedPlugin shared.CreateResponseTransformerAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *UpdateResponsetransformeradvancedPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateResponsetransformeradvancedPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateResponsetransformeradvancedPluginRequest) GetCreateResponseTransformerAdvancedPlugin() shared.CreateResponseTransformerAdvancedPlugin {
	if o == nil {
		return shared.CreateResponseTransformerAdvancedPlugin{}
	}
	return o.CreateResponseTransformerAdvancedPlugin
}

type UpdateResponsetransformeradvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Plugin
	ResponseTransformerAdvancedPlugin *shared.ResponseTransformerAdvancedPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetResponseTransformerAdvancedPlugin() *shared.ResponseTransformerAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.ResponseTransformerAdvancedPlugin
}

func (o *UpdateResponsetransformeradvancedPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
