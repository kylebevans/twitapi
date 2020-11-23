/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitclient
// Place struct for Place
type Place struct {
	// The identifier for this place
	Id string `json:"id"`
	// The human readable name of this place.
	Name string `json:"name,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	PlaceType PlaceType `json:"place_type,omitempty"`
	FullName string `json:"full_name"`
	Country string `json:"country,omitempty"`
	ContainedWithin []string `json:"contained_within,omitempty"`
	Geo Geo `json:"geo,omitempty"`
}
