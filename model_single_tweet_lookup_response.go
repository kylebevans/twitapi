/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitapi
// SingleTweetLookupResponse struct for SingleTweetLookupResponse
type SingleTweetLookupResponse struct {
	Data Tweet `json:"data,omitempty"`
	Includes Expansions `json:"includes,omitempty"`
	Errors []Problem `json:"errors,omitempty"`
}
