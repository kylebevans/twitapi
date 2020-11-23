# Tweet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | 
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time of the Tweet. | [optional] 
**Text** | **string** | The content of the Tweet. | 
**AuthorId** | **string** | Unique identifier of this User. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | [optional] 
**InReplyToUserId** | **string** | Unique identifier of this User. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | [optional] 
**ConversationId** | **string** | Unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers. | [optional] 
**ConversationControl** | [**ConversationControl**](ConversationControl.md) |  | [optional] 
**ReferencedTweets** | [**[]TweetReferencedTweets**](Tweet_referenced_tweets.md) | A list of Tweets this Tweet refers to. For example, if the parent Tweet is a Retweet, a Quoted Tweet or a Reply, it will include the related Tweet referenced to by its parent. | [optional] 
**Attachments** | [**TweetAttachments**](Tweet_attachments.md) |  | [optional] 
**ContextAnnotations** | [**[]ContextAnnotation**](ContextAnnotation.md) |  | [optional] 
**Withheld** | [**TweetWithheld**](TweetWithheld.md) |  | [optional] 
**Geo** | [**TweetGeo**](Tweet_geo.md) |  | [optional] 
**Entities** | [**FullTextEntities**](FullTextEntities.md) |  | [optional] 
**PublicMetrics** | [**TweetPublicMetrics**](Tweet_public_metrics.md) |  | [optional] 
**PossiblySensitive** | **bool** | Indicates if this Tweet contains URLs marked as sensitive, for example content suitable for mature audiences. | [optional] 
**Lang** | **string** | Language of the Tweet, if detected by Twitter. Returned as a BCP47 language tag. | [optional] 
**Source** | **string** | The name of the app the user Tweeted from. | [optional] 
**NonPublicMetrics** | [**TweetNonPublicMetrics**](Tweet_non_public_metrics.md) |  | [optional] 
**PromotedMetrics** | [**TweetPromotedMetrics**](Tweet_promoted_metrics.md) |  | [optional] 
**OrganicMetrics** | [**TweetOrganicMetrics**](Tweet_organic_metrics.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


