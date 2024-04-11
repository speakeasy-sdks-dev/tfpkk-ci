// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type CreateBasicAuthPluginProtocols string

const (
	CreateBasicAuthPluginProtocolsGrpc           CreateBasicAuthPluginProtocols = "grpc"
	CreateBasicAuthPluginProtocolsGrpcs          CreateBasicAuthPluginProtocols = "grpcs"
	CreateBasicAuthPluginProtocolsHTTP           CreateBasicAuthPluginProtocols = "http"
	CreateBasicAuthPluginProtocolsHTTPS          CreateBasicAuthPluginProtocols = "https"
	CreateBasicAuthPluginProtocolsTCP            CreateBasicAuthPluginProtocols = "tcp"
	CreateBasicAuthPluginProtocolsTLS            CreateBasicAuthPluginProtocols = "tls"
	CreateBasicAuthPluginProtocolsTLSPassthrough CreateBasicAuthPluginProtocols = "tls_passthrough"
	CreateBasicAuthPluginProtocolsUDP            CreateBasicAuthPluginProtocols = "udp"
	CreateBasicAuthPluginProtocolsWs             CreateBasicAuthPluginProtocols = "ws"
	CreateBasicAuthPluginProtocolsWss            CreateBasicAuthPluginProtocols = "wss"
)

func (e CreateBasicAuthPluginProtocols) ToPointer() *CreateBasicAuthPluginProtocols {
	return &e
}

func (e *CreateBasicAuthPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = CreateBasicAuthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateBasicAuthPluginProtocols: %v", v)
	}
}

// CreateBasicAuthPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateBasicAuthPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateBasicAuthPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateBasicAuthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateBasicAuthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateBasicAuthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateBasicAuthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateBasicAuthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateBasicAuthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateBasicAuthPluginConfig struct {
	// An optional string (Consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request will fail with an authentication failure `4xx`. Please note that this value must refer to the Consumer `id` or `username` attribute, and **not** its `custom_id`.
	Anonymous *string `json:"anonymous,omitempty"`
	// An optional boolean value telling the plugin to show or hide the credential from the upstream service. If `true`, the plugin will strip the credential from the request (i.e. the `Authorization` header) before proxying it.
	HideCredentials *bool `default:"false" json:"hide_credentials"`
}

func (c CreateBasicAuthPluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateBasicAuthPluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateBasicAuthPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *CreateBasicAuthPluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

// CreateBasicAuthPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CreateBasicAuthPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"basic-auth" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateBasicAuthPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *CreateBasicAuthPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateBasicAuthPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateBasicAuthPluginService `json:"service,omitempty"`
	Config  CreateBasicAuthPluginConfig   `json:"config"`
}

func (c CreateBasicAuthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateBasicAuthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateBasicAuthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateBasicAuthPlugin) GetName() string {
	return "basic-auth"
}

func (o *CreateBasicAuthPlugin) GetProtocols() []CreateBasicAuthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateBasicAuthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateBasicAuthPlugin) GetConsumer() *CreateBasicAuthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateBasicAuthPlugin) GetRoute() *CreateBasicAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateBasicAuthPlugin) GetService() *CreateBasicAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CreateBasicAuthPlugin) GetConfig() CreateBasicAuthPluginConfig {
	if o == nil {
		return CreateBasicAuthPluginConfig{}
	}
	return o.Config
}
