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

// FeedSyncResult The result of a sync of a single feed
type FeedSyncResult struct {
	// The name of the feed synced
	Feed string `json:"feed,omitempty"`
	// The result of the sync operations, either co
	Status string `json:"status,omitempty"`
	// The duratin, in seconds, of the sync of the feed, the sum of all the group syncs
	TotalTimeSeconds float32 `json:"total_time_seconds,omitempty"`
	// Array of group sync results
	Groups []GroupSyncResult `json:"groups,omitempty"`
}
