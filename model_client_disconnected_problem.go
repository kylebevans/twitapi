/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitclient
// ClientDisconnectedProblem Your client has gone away.
type ClientDisconnectedProblem struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Type string `json:"type,omitempty"`
}
