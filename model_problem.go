/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitapi
// Problem An HTTP Problem Details object, as defined in IETF RFC 7807 (https://tools.ietf.org/html/rfc7807).
type Problem struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Type string `json:"type"`
	Status int32 `json:"status"`
	Errors []InvalidRequestProblemAllOfErrors `json:"errors,omitempty"`
	Reason string `json:"reason,omitempty"`
	RegistrationUrl string `json:"registration_url,omitempty"`
	Parameter string `json:"parameter"`
	Value string `json:"value"`
	ResourceType string `json:"resource_type"`
	Section string `json:"section"`
	Field string `json:"field"`
	ResourceId string `json:"resource_id"`
	Period string `json:"period,omitempty"`
	Scope string `json:"scope,omitempty"`
	ConnectionIssue string `json:"connection_issue,omitempty"`
	DisconnectType string `json:"disconnect_type,omitempty"`
}
