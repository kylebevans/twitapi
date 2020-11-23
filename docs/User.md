# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of this User. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | 
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time of this user. | [optional] 
**Name** | **string** | The friendly name of this user, as shown on their profile. | 
**Username** | **string** | The Twitter handle (screen name) of this user. | 
**Protected** | **bool** | Indicates if this user has chosen to protect their Tweets (in other words, if this user&#39;s Tweets are private). | [optional] 
**Verified** | **bool** | Indicate if this user is a verified Twitter User. | [optional] 
**Withheld** | [**UserWithheld**](UserWithheld.md) |  | [optional] 
**ProfileImageUrl** | **string** | The URL to the profile image for this user. | [optional] 
**Location** | **string** | The location specified in the user&#39;s profile, if the user provided one. As this is a freeform value, it may not indicate a valid location, but it may be fuzzily evaluated when performing searches with location queries. | [optional] 
**Url** | **string** | The URL specified in the user&#39;s profile. | [optional] 
**Description** | **string** | The text of this user&#39;s profile description (also known as bio), if the user provided one. | [optional] 
**Entities** | [**UserEntities**](User_entities.md) |  | [optional] 
**PinnedTweetId** | **string** | Unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | [optional] 
**PublicMetrics** | [**UserPublicMetrics**](User_public_metrics.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


