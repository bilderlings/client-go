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

// AnchoreImageTagSummary A unique image in the engine.
type AnchoreImageTagSummary struct {
	ImageDigest    string `json:"imageDigest,omitempty"`
	ParentDigest   string `json:"parentDigest,omitempty"`
	ImageId        string `json:"imageId,omitempty"`
	AnalysisStatus string `json:"analysis_status,omitempty"`
	Fulltag        string `json:"fulltag,omitempty"`
	CreatedAt      int32  `json:"created_at,omitempty"`
	AnalyzedAt     int32  `json:"analyzed_at,omitempty"`
	TagDetectedAt  int32  `json:"tag_detected_at,omitempty"`
}
