/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitclient
// DisallowedResourceProblem A problem that indicates that the resource requested violates the precepts of this API.
type DisallowedResourceProblem struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Type string `json:"type,omitempty"`
	ResourceId string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	Section string `json:"section"`
}
