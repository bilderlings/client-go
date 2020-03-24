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

// RegexContentMatch Match of a named regex on a file
type RegexContentMatch struct {
	// The name associated with the regular expression
	Name string `json:"name,omitempty"`
	// The regular expression used for the match
	Regex string `json:"regex,omitempty"`
	// A list of line numbers in the file that matched the regex
	Lines []int32 `json:"lines,omitempty"`
}
