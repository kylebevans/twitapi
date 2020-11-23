/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitclient
// UserPublicMetrics A list of metrics for this user
type UserPublicMetrics struct {
	// Number of users who are following this user.
	FollowersCount int32 `json:"followers_count"`
	// Number of users this user is following.
	FollowingCount int32 `json:"following_count"`
	// The number of Tweets (including Retweets) posted by this user.
	TweetCount int32 `json:"tweet_count"`
	// The number of lists that include this user.
	ListedCount int32 `json:"listed_count"`
}
