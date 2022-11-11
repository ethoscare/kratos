/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.
 *
 * API version:
 * Contact: office@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// VerificationFlowState The state represents the state of the verification flow.  choose_method: ask the user to choose a method (e.g. recover account via email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the recovery challenge was passed.
type VerificationFlowState string

// List of verificationFlowState
const (
	VERIFICATIONFLOWSTATE_CHOOSE_METHOD    VerificationFlowState = "choose_method"
	VERIFICATIONFLOWSTATE_SENT_EMAIL       VerificationFlowState = "sent_email"
	VERIFICATIONFLOWSTATE_PASSED_CHALLENGE VerificationFlowState = "passed_challenge"
)

func (v *VerificationFlowState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerificationFlowState(value)
	for _, existing := range []VerificationFlowState{"choose_method", "sent_email", "passed_challenge"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerificationFlowState", value)
}

// Ptr returns reference to verificationFlowState value
func (v VerificationFlowState) Ptr() *VerificationFlowState {
	return &v
}

type NullableVerificationFlowState struct {
	value *VerificationFlowState
	isSet bool
}

func (v NullableVerificationFlowState) Get() *VerificationFlowState {
	return v.value
}

func (v *NullableVerificationFlowState) Set(val *VerificationFlowState) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationFlowState) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationFlowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationFlowState(val *VerificationFlowState) *NullableVerificationFlowState {
	return &NullableVerificationFlowState{value: val, isSet: true}
}

func (v NullableVerificationFlowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationFlowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}