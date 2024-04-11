// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type CreateRequestTransformerPluginProtocols string

const (
	CreateRequestTransformerPluginProtocolsGrpc           CreateRequestTransformerPluginProtocols = "grpc"
	CreateRequestTransformerPluginProtocolsGrpcs          CreateRequestTransformerPluginProtocols = "grpcs"
	CreateRequestTransformerPluginProtocolsHTTP           CreateRequestTransformerPluginProtocols = "http"
	CreateRequestTransformerPluginProtocolsHTTPS          CreateRequestTransformerPluginProtocols = "https"
	CreateRequestTransformerPluginProtocolsTCP            CreateRequestTransformerPluginProtocols = "tcp"
	CreateRequestTransformerPluginProtocolsTLS            CreateRequestTransformerPluginProtocols = "tls"
	CreateRequestTransformerPluginProtocolsTLSPassthrough CreateRequestTransformerPluginProtocols = "tls_passthrough"
	CreateRequestTransformerPluginProtocolsUDP            CreateRequestTransformerPluginProtocols = "udp"
	CreateRequestTransformerPluginProtocolsWs             CreateRequestTransformerPluginProtocols = "ws"
	CreateRequestTransformerPluginProtocolsWss            CreateRequestTransformerPluginProtocols = "wss"
)

func (e CreateRequestTransformerPluginProtocols) ToPointer() *CreateRequestTransformerPluginProtocols {
	return &e
}

func (e *CreateRequestTransformerPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateRequestTransformerPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRequestTransformerPluginProtocols: %v", v)
	}
}

// CreateRequestTransformerPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateRequestTransformerPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTransformerPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestTransformerPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateRequestTransformerPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTransformerPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestTransformerPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateRequestTransformerPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTransformerPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Remove struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *Remove) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Remove) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Remove) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type Rename struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *Rename) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Rename) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Rename) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type Replace struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
	URI         *string  `json:"uri,omitempty"`
}

func (o *Replace) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Replace) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Replace) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

func (o *Replace) GetURI() *string {
	if o == nil {
		return nil
	}
	return o.URI
}

type Add struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *Add) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Add) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Add) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type Append struct {
	Body        []string `json:"body,omitempty"`
	Headers     []string `json:"headers,omitempty"`
	Querystring []string `json:"querystring,omitempty"`
}

func (o *Append) GetBody() []string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Append) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Append) GetQuerystring() []string {
	if o == nil {
		return nil
	}
	return o.Querystring
}

type CreateRequestTransformerPluginConfig struct {
	// A string representing an HTTP method, such as GET, POST, PUT, or DELETE. The string must contain only uppercase letters.
	HTTPMethod *string  `json:"http_method,omitempty"`
	Remove     *Remove  `json:"remove,omitempty"`
	Rename     *Rename  `json:"rename,omitempty"`
	Replace    *Replace `json:"replace,omitempty"`
	Add        *Add     `json:"add,omitempty"`
	Append     *Append  `json:"append,omitempty"`
}

func (o *CreateRequestTransformerPluginConfig) GetHTTPMethod() *string {
	if o == nil {
		return nil
	}
	return o.HTTPMethod
}

func (o *CreateRequestTransformerPluginConfig) GetRemove() *Remove {
	if o == nil {
		return nil
	}
	return o.Remove
}

func (o *CreateRequestTransformerPluginConfig) GetRename() *Rename {
	if o == nil {
		return nil
	}
	return o.Rename
}

func (o *CreateRequestTransformerPluginConfig) GetReplace() *Replace {
	if o == nil {
		return nil
	}
	return o.Replace
}

func (o *CreateRequestTransformerPluginConfig) GetAdd() *Add {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *CreateRequestTransformerPluginConfig) GetAppend() *Append {
	if o == nil {
		return nil
	}
	return o.Append
}

// CreateRequestTransformerPlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CreateRequestTransformerPlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"request-transformer" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateRequestTransformerPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *CreateRequestTransformerPluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateRequestTransformerPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateRequestTransformerPluginService `json:"service,omitempty"`
	Config  CreateRequestTransformerPluginConfig   `json:"config"`
}

func (c CreateRequestTransformerPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateRequestTransformerPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateRequestTransformerPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateRequestTransformerPlugin) GetName() string {
	return "request-transformer"
}

func (o *CreateRequestTransformerPlugin) GetProtocols() []CreateRequestTransformerPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateRequestTransformerPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateRequestTransformerPlugin) GetConsumer() *CreateRequestTransformerPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateRequestTransformerPlugin) GetRoute() *CreateRequestTransformerPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateRequestTransformerPlugin) GetService() *CreateRequestTransformerPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CreateRequestTransformerPlugin) GetConfig() CreateRequestTransformerPluginConfig {
	if o == nil {
		return CreateRequestTransformerPluginConfig{}
	}
	return o.Config
}
