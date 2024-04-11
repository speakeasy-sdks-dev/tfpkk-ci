// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UnsupportedMediaTypeError - standard error
type UnsupportedMediaTypeError struct {
	Status   interface{} `json:"status"`
	Title    interface{} `json:"title"`
	Type     interface{} `json:"type,omitempty"`
	Instance interface{} `json:"instance"`
	Detail   interface{} `json:"detail"`
}

func (o *UnsupportedMediaTypeError) GetStatus() interface{} {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnsupportedMediaTypeError) GetTitle() interface{} {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UnsupportedMediaTypeError) GetType() interface{} {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UnsupportedMediaTypeError) GetInstance() interface{} {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *UnsupportedMediaTypeError) GetDetail() interface{} {
	if o == nil {
		return nil
	}
	return o.Detail
}
