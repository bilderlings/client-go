/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.14
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

// GateSpec A description of the set of gates available in this engine and the triggers and parameters supported
type GateSpec struct {
	// Gate name, as it would appear in a policy document
	Name string `json:"name,omitempty"`
	// Description of the gate
	Description string `json:"description,omitempty"`
	// State of the gate and transitively all triggers it contains if not 'active'
	State string `json:"state,omitempty"`
	// The name of another trigger that supercedes this on functionally if this is deprecated
	SupercededBy *string `json:"superceded_by,omitempty"`
	// List of the triggers that can fire for this Gate
	Triggers []TriggerSpec `json:"triggers,omitempty"`
}
