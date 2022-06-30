/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CveGetCveListResponseBody struct for CveGetCveListResponseBody
type CveGetCveListResponseBody struct {
	// Cves list
	Cves []CveListResponseBody `json:"cves,omitempty"`
	Paging *PagingResponseBody `json:"paging,omitempty"`
}

// NewCveGetCveListResponseBody instantiates a new CveGetCveListResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCveGetCveListResponseBody() *CveGetCveListResponseBody {
	this := CveGetCveListResponseBody{}
	return &this
}

// NewCveGetCveListResponseBodyWithDefaults instantiates a new CveGetCveListResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCveGetCveListResponseBodyWithDefaults() *CveGetCveListResponseBody {
	this := CveGetCveListResponseBody{}
	return &this
}

// GetCves returns the Cves field value if set, zero value otherwise.
func (o *CveGetCveListResponseBody) GetCves() []CveListResponseBody {
	if o == nil || o.Cves == nil {
		var ret []CveListResponseBody
		return ret
	}
	return o.Cves
}

// GetCvesOk returns a tuple with the Cves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CveGetCveListResponseBody) GetCvesOk() ([]CveListResponseBody, bool) {
	if o == nil || o.Cves == nil {
		return nil, false
	}
	return o.Cves, true
}

// HasCves returns a boolean if a field has been set.
func (o *CveGetCveListResponseBody) HasCves() bool {
	if o != nil && o.Cves != nil {
		return true
	}

	return false
}

// SetCves gets a reference to the given []CveListResponseBody and assigns it to the Cves field.
func (o *CveGetCveListResponseBody) SetCves(v []CveListResponseBody) {
	o.Cves = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CveGetCveListResponseBody) GetPaging() PagingResponseBody {
	if o == nil || o.Paging == nil {
		var ret PagingResponseBody
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CveGetCveListResponseBody) GetPagingOk() (*PagingResponseBody, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CveGetCveListResponseBody) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given PagingResponseBody and assigns it to the Paging field.
func (o *CveGetCveListResponseBody) SetPaging(v PagingResponseBody) {
	o.Paging = &v
}

func (o CveGetCveListResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cves != nil {
		toSerialize["cves"] = o.Cves
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableCveGetCveListResponseBody struct {
	value *CveGetCveListResponseBody
	isSet bool
}

func (v NullableCveGetCveListResponseBody) Get() *CveGetCveListResponseBody {
	return v.value
}

func (v *NullableCveGetCveListResponseBody) Set(val *CveGetCveListResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCveGetCveListResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCveGetCveListResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCveGetCveListResponseBody(val *CveGetCveListResponseBody) *NullableCveGetCveListResponseBody {
	return &NullableCveGetCveListResponseBody{value: val, isSet: true}
}

func (v NullableCveGetCveListResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCveGetCveListResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


