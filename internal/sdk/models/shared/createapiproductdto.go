// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

// CreateAPIProductDTO - The request schema to create an API product.
type CreateAPIProductDTO struct {
	// The name of the API product.
	Name string `json:"name"`
	// The description of the API product.
	Description *string `default:"null" json:"description"`
	// description: A maximum of 5 user-defined labels are allowed on this resource.
	// Keys must not start with kong, konnect, insomnia, mesh, kic or _, which are reserved for Kong.
	// Keys are case-sensitive.
	//
	Labels map[string]string `json:"labels,omitempty"`
	// The list of portal identifiers which this API product should be published to
	PortalIds []string `json:"portal_ids,omitempty"`
}

func (c CreateAPIProductDTO) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAPIProductDTO) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAPIProductDTO) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateAPIProductDTO) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateAPIProductDTO) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *CreateAPIProductDTO) GetPortalIds() []string {
	if o == nil {
		return nil
	}
	return o.PortalIds
}
