// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateJqPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Create a new JQ plugin
	CreateJQPlugin shared.CreateJQPlugin `request:"mediaType=application/json"`
}

func (o *CreateJqPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateJqPluginRequest) GetCreateJQPlugin() shared.CreateJQPlugin {
	if o == nil {
		return shared.CreateJQPlugin{}
	}
	return o.CreateJQPlugin
}

type CreateJqPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Plugin
	JQPlugin *shared.JQPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *CreateJqPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateJqPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateJqPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateJqPluginResponse) GetJQPlugin() *shared.JQPlugin {
	if o == nil {
		return nil
	}
	return o.JQPlugin
}

func (o *CreateJqPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}
