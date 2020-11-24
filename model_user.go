/*
 * Early Access
 *
 * API Reference — v2
 *
 * API version: 2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twitapi
import (
	"time"
)
// User The Twitter User object
type User struct {
	// Unique identifier of this User. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers.
	Id string `json:"id"`
	// Creation time of this user.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The friendly name of this user, as shown on their profile.
	Name string `json:"name"`
	// The Twitter handle (screen name) of this user.
	Username string `json:"username"`
	// Indicates if this user has chosen to protect their Tweets (in other words, if this user's Tweets are private).
	Protected bool `json:"protected,omitempty"`
	// Indicate if this user is a verified Twitter User.
	Verified bool `json:"verified,omitempty"`
	Withheld UserWithheld `json:"withheld,omitempty"`
	// The URL to the profile image for this user.
	ProfileImageUrl string `json:"profile_image_url,omitempty"`
	// The location specified in the user's profile, if the user provided one. As this is a freeform value, it may not indicate a valid location, but it may be fuzzily evaluated when performing searches with location queries.
	Location string `json:"location,omitempty"`
	// The URL specified in the user's profile.
	Url string `json:"url,omitempty"`
	// The text of this user's profile description (also known as bio), if the user provided one.
	Description string `json:"description,omitempty"`
	Entities UserEntities `json:"entities,omitempty"`
	// Unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers.
	PinnedTweetId string `json:"pinned_tweet_id,omitempty"`
	PublicMetrics UserPublicMetrics `json:"public_metrics,omitempty"`
}
