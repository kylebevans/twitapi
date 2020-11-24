/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitapi
// EntityIndices Represent a boundary range (start and end index) for a recognized entity (for example a hashtag or a mention). `start` must be smaller than `end`.
type EntityIndices struct {
	// Index (zero-based) at which position this entity starts.
	Start int32 `json:"start"`
	// Index (zero-based) at which position this entity ends.
	End int32 `json:"end"`
}
