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

// ImageAnalysisRequest A request to add an image to be watched and analyzed by the engine. Optionally include the dockerfile content. Either source, digest or tag must be present.
type ImageAnalysisRequest struct {
	// Optional. The type of image this is adding, defaults to \"docker\". This can be ommitted until multiple image types are supported.
	ImageType string `json:"image_type,omitempty"`
	// Annotations to be associated with the added image in key/value form
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	Source      ImageSource            `json:"source,omitempty"`
}
