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

// PaginatedVulnerableImageListAllOf struct for PaginatedVulnerableImageListAllOf
type PaginatedVulnerableImageListAllOf struct {
	Images []VulnerableImage `json:"images,omitempty"`
}
