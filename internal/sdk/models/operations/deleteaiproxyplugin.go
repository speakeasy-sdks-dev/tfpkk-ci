// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type DeleteAiproxyPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
}

func (o *DeleteAiproxyPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *DeleteAiproxyPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

type DeleteAiproxyPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *DeleteAiproxyPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteAiproxyPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAiproxyPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteAiproxyPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
