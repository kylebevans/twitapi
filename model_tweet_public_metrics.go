/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitapi
// TweetPublicMetrics Engagement metrics for the Tweet at the time of the request.
type TweetPublicMetrics struct {
	// Number of times this Tweet has been Retweeted.
	RetweetCount int32 `json:"retweet_count"`
	// Number of times this Tweet has been replied to.
	ReplyCount int32 `json:"reply_count"`
	// Number of times this Tweet has been liked.
	LikeCount int32 `json:"like_count"`
	// Number of times this Tweet has been quoted.
	QuoteCount int32 `json:"quote_count,omitempty"`
}
