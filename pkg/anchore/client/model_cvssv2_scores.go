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

// Cvssv2Scores struct for Cvssv2Scores
type Cvssv2Scores struct {
	BaseScore           *float32 `json:"base_score,omitempty"`
	ExploitabilityScore *float32 `json:"exploitability_score,omitempty"`
	ImpactScore         *float32 `json:"impact_score,omitempty"`
}
