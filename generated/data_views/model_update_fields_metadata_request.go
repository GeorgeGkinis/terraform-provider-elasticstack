/*
Data views

OpenAPI schema for data view endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_views

import (
	"encoding/json"
)

// checks if the UpdateFieldsMetadataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFieldsMetadataRequest{}

// UpdateFieldsMetadataRequest struct for UpdateFieldsMetadataRequest
type UpdateFieldsMetadataRequest struct {
	// The field object.
	Fields map[string]interface{} `json:"fields"`
}

// NewUpdateFieldsMetadataRequest instantiates a new UpdateFieldsMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFieldsMetadataRequest(fields map[string]interface{}) *UpdateFieldsMetadataRequest {
	this := UpdateFieldsMetadataRequest{}
	this.Fields = fields
	return &this
}

// NewUpdateFieldsMetadataRequestWithDefaults instantiates a new UpdateFieldsMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFieldsMetadataRequestWithDefaults() *UpdateFieldsMetadataRequest {
	this := UpdateFieldsMetadataRequest{}
	return &this
}

// GetFields returns the Fields field value
func (o *UpdateFieldsMetadataRequest) GetFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *UpdateFieldsMetadataRequest) GetFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *UpdateFieldsMetadataRequest) SetFields(v map[string]interface{}) {
	o.Fields = v
}

func (o UpdateFieldsMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFieldsMetadataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

type NullableUpdateFieldsMetadataRequest struct {
	value *UpdateFieldsMetadataRequest
	isSet bool
}

func (v NullableUpdateFieldsMetadataRequest) Get() *UpdateFieldsMetadataRequest {
	return v.value
}

func (v *NullableUpdateFieldsMetadataRequest) Set(val *UpdateFieldsMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFieldsMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFieldsMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFieldsMetadataRequest(val *UpdateFieldsMetadataRequest) *NullableUpdateFieldsMetadataRequest {
	return &NullableUpdateFieldsMetadataRequest{value: val, isSet: true}
}

func (v NullableUpdateFieldsMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFieldsMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}