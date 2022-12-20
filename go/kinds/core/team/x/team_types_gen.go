package team

import (
	"encoding/json"
	"fmt"
)

// Defines values for Permission.
const (
	PermissionN1 Permission = 1

	PermissionN2 Permission = 2

	PermissionN4 Permission = 4
)

// Permission defines model for Permission.
type Permission int

// Team defines model for team.
type Team struct {
	// AccessControl metadata associated with a given resource.
	AccessControl _AccessControl `json:"accessControl"`

	// AvatarUrl is the team's avatar URL.
	AvatarUrl *string `json:"avatarUrl,omitempty"`

	// Created indicates when the team was created.
	Created int64 `json:"created"`

	// Email of the team.
	Email *string `json:"email,omitempty"`

	// MemberCount is the number of the team members.
	MemberCount int64 `json:"memberCount"`

	// Name of the team.
	Name string `json:"name"`

	// OrgId is the ID of an organisation the team belongs to.
	OrgId      int64      `json:"orgId"`
	Permission Permission `json:"permission"`

	// Updated indicates when the team was updated.
	Updated int64 `json:"updated"`
}

// AccessControl metadata associated with a given resource.
type _AccessControl struct {
	AdditionalProperties map[string]bool `json:"-"`
}

// Getter for additional properties for _AccessControl. Returns the specified
// element and whether it was found
func (a _AccessControl) Get(fieldName string) (value bool, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for _AccessControl
func (a *_AccessControl) Set(fieldName string, value bool) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]bool)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for _AccessControl to handle AdditionalProperties
func (a *_AccessControl) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]bool)
		for fieldName, fieldBuf := range object {
			var fieldVal bool
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for _AccessControl to handle AdditionalProperties
func (a _AccessControl) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
