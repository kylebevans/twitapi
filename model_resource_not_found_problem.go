/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitclient
// ResourceNotFoundProblem A problem that indicates that a given Tweet, User, etc. does not exist.
type ResourceNotFoundProblem struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Type string `json:"type,omitempty"`
	Parameter string `json:"parameter"`
	// Value will match the schema of the field.
	Value interface{} `json:"value"`
	ResourceType string `json:"resource_type"`
}
