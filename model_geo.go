/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitapi
// Geo struct for Geo
type Geo struct {
	Type string `json:"type"`
	Bbox []float64 `json:"bbox"`
	Geometry Point `json:"geometry,omitempty"`
	Properties map[string]interface{} `json:"properties"`
}
