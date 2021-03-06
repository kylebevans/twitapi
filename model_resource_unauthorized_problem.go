/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitapi
// ResourceUnauthorizedProblem A problem that indicates you are not allowed to see a particular Tweet, User, etc.
type ResourceUnauthorizedProblem struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Type string `json:"type,omitempty"`
	Value string `json:"value"`
	Parameter string `json:"parameter"`
	Section string `json:"section"`
	ResourceType string `json:"resource_type"`
}
