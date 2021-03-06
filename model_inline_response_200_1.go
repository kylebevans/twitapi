/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitapi
// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	Data []Tweet `json:"data,omitempty"`
	Includes Expansions `json:"includes,omitempty"`
	Errors []Problem `json:"errors,omitempty"`
	Meta InlineResponse2001Meta `json:"meta,omitempty"`
}
