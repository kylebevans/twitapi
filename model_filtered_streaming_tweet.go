/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitapi
// FilteredStreamingTweet A tweet or error that can be returned by the streaming tweet API
type FilteredStreamingTweet struct {
	Data Tweet `json:"data"`
	// The list of rules which matched the tweet
	MatchingRules []FilteredStreamingTweetOneOfMatchingRules `json:"matching_rules"`
	Includes Expansions `json:"includes,omitempty"`
	Errors []Problem `json:"errors"`
}
