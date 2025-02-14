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
)

// SmsPreview struct for SmsPreview
type SmsPreview struct {
	// Number of remaining characters in the last SMS part.
	CharactersRemaining *int32 `json:"charactersRemaining,omitempty"`
	// Configuration that, when sent with the original text, results in this preview.
	Configuration *SmsLanguageConfiguration `json:"configuration,omitempty"`
	// Number of SMS message parts required to deliver the message.
	MessageCount *int32 `json:"messageCount,omitempty"`
	// Preview of the text as it should appear on the recipient’s device.
	TextPreview *string `json:"textPreview,omitempty"`
}

// NewSmsPreview instantiates a new SmsPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsPreview() *SmsPreview {
	this := SmsPreview{}
	return &this
}

// NewSmsPreviewWithDefaults instantiates a new SmsPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsPreviewWithDefaults() *SmsPreview {
	this := SmsPreview{}
	return &this
}

// GetCharactersRemaining returns the CharactersRemaining field value if set, zero value otherwise.
func (o *SmsPreview) GetCharactersRemaining() int32 {
	if o == nil || o.CharactersRemaining == nil {
		var ret int32
		return ret
	}
	return *o.CharactersRemaining
}

// GetCharactersRemainingOk returns a tuple with the CharactersRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsPreview) GetCharactersRemainingOk() (*int32, bool) {
	if o == nil || o.CharactersRemaining == nil {
		return nil, false
	}
	return o.CharactersRemaining, true
}

// HasCharactersRemaining returns a boolean if a field has been set.
func (o *SmsPreview) HasCharactersRemaining() bool {
	if o != nil && o.CharactersRemaining != nil {
		return true
	}

	return false
}

// SetCharactersRemaining gets a reference to the given int32 and assigns it to the CharactersRemaining field.
func (o *SmsPreview) SetCharactersRemaining(v int32) {
	o.CharactersRemaining = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *SmsPreview) GetConfiguration() SmsLanguageConfiguration {
	if o == nil || o.Configuration == nil {
		var ret SmsLanguageConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsPreview) GetConfigurationOk() (*SmsLanguageConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *SmsPreview) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given SmsLanguageConfiguration and assigns it to the Configuration field.
func (o *SmsPreview) SetConfiguration(v SmsLanguageConfiguration) {
	o.Configuration = &v
}

// GetMessageCount returns the MessageCount field value if set, zero value otherwise.
func (o *SmsPreview) GetMessageCount() int32 {
	if o == nil || o.MessageCount == nil {
		var ret int32
		return ret
	}
	return *o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsPreview) GetMessageCountOk() (*int32, bool) {
	if o == nil || o.MessageCount == nil {
		return nil, false
	}
	return o.MessageCount, true
}

// HasMessageCount returns a boolean if a field has been set.
func (o *SmsPreview) HasMessageCount() bool {
	if o != nil && o.MessageCount != nil {
		return true
	}

	return false
}

// SetMessageCount gets a reference to the given int32 and assigns it to the MessageCount field.
func (o *SmsPreview) SetMessageCount(v int32) {
	o.MessageCount = &v
}

// GetTextPreview returns the TextPreview field value if set, zero value otherwise.
func (o *SmsPreview) GetTextPreview() string {
	if o == nil || o.TextPreview == nil {
		var ret string
		return ret
	}
	return *o.TextPreview
}

// GetTextPreviewOk returns a tuple with the TextPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsPreview) GetTextPreviewOk() (*string, bool) {
	if o == nil || o.TextPreview == nil {
		return nil, false
	}
	return o.TextPreview, true
}

// HasTextPreview returns a boolean if a field has been set.
func (o *SmsPreview) HasTextPreview() bool {
	if o != nil && o.TextPreview != nil {
		return true
	}

	return false
}

// SetTextPreview gets a reference to the given string and assigns it to the TextPreview field.
func (o *SmsPreview) SetTextPreview(v string) {
	o.TextPreview = &v
}

func (o SmsPreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CharactersRemaining != nil {
		toSerialize["charactersRemaining"] = o.CharactersRemaining
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.MessageCount != nil {
		toSerialize["messageCount"] = o.MessageCount
	}
	if o.TextPreview != nil {
		toSerialize["textPreview"] = o.TextPreview
	}
	return json.Marshal(toSerialize)
}

type NullableSmsPreview struct {
	value *SmsPreview
	isSet bool
}

func (v NullableSmsPreview) Get() *SmsPreview {
	return v.value
}

func (v *NullableSmsPreview) Set(val *SmsPreview) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsPreview) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsPreview(val *SmsPreview) *NullableSmsPreview {
	return &NullableSmsPreview{value: val, isSet: true}
}

func (v NullableSmsPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
