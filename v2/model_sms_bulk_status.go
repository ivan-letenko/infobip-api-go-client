/**
 * Infobip Client API Libraries OpenAPI Specification
 *
 * OpenAPI specification containing public endpoints supported in client API libraries.
 *
 * Contact: support@infobip.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit the class manually.
 */

package infobip

import (
	"encoding/json"
	"fmt"
)

// SmsBulkStatus the model 'SmsBulkStatus'
type SmsBulkStatus string

// List of SmsBulkStatus
const (
	SMSBULKSTATUS_PENDING    SmsBulkStatus = "PENDING"
	SMSBULKSTATUS_PAUSED     SmsBulkStatus = "PAUSED"
	SMSBULKSTATUS_PROCESSING SmsBulkStatus = "PROCESSING"
	SMSBULKSTATUS_CANCELED   SmsBulkStatus = "CANCELED"
	SMSBULKSTATUS_FINISHED   SmsBulkStatus = "FINISHED"
	SMSBULKSTATUS_FAILED     SmsBulkStatus = "FAILED"
)

func (v *SmsBulkStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SmsBulkStatus(value)
	for _, existing := range []SmsBulkStatus{"PENDING", "PAUSED", "PROCESSING", "CANCELED", "FINISHED", "FAILED"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SmsBulkStatus", value)
}

// Ptr returns reference to SmsBulkStatus value
func (v SmsBulkStatus) Ptr() *SmsBulkStatus {
	return &v
}

type NullableSmsBulkStatus struct {
	value *SmsBulkStatus
	isSet bool
}

func (v NullableSmsBulkStatus) Get() *SmsBulkStatus {
	return v.value
}

func (v *NullableSmsBulkStatus) Set(val *SmsBulkStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsBulkStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsBulkStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsBulkStatus(val *SmsBulkStatus) *NullableSmsBulkStatus {
	return &NullableSmsBulkStatus{value: val, isSet: true}
}

func (v NullableSmsBulkStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsBulkStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
