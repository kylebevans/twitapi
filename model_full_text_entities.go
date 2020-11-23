/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitclient
// FullTextEntities struct for FullTextEntities
type FullTextEntities struct {
	Urls []UrlEntity `json:"urls,omitempty"`
	Hashtags []HashtagEntity `json:"hashtags,omitempty"`
	Mentions []MentionEntity `json:"mentions,omitempty"`
	Cashtags []CashtagEntity `json:"cashtags,omitempty"`
}
