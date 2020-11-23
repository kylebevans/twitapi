/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitclient
// GenericProblem A generic problem with no additional information beyond that provided by the HTTP status code.
type GenericProblem struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Type string `json:"type,omitempty"`
	Status int32 `json:"status"`
}
