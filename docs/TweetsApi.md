# \TweetsApi

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrDeleteRules**](TweetsApi.md#AddOrDeleteRules) | **Post** /2/tweets/search/stream/rules | Add or delete rules from a user&#39;s active rule set.
[**FindTweetById**](TweetsApi.md#FindTweetById) | **Get** /2/tweets/{id} | Returns hydrated Tweet objects
[**FindTweetsById**](TweetsApi.md#FindTweetsById) | **Get** /2/tweets | Returns hydrated Tweet objects
[**GetRules**](TweetsApi.md#GetRules) | **Get** /2/tweets/search/stream/rules | Returns rules from a user&#39;s active rule set.
[**HideReplyById**](TweetsApi.md#HideReplyById) | **Put** /2/tweets/{id}/hidden | Hides or unhides a reply to an owned conversation.
[**SampleStream**](TweetsApi.md#SampleStream) | **Get** /2/tweets/sample/stream | Streams a deterministic 1% of public tweets.
[**SearchStream**](TweetsApi.md#SearchStream) | **Get** /2/tweets/search/stream | Streams tweets matching a user&#39;s active rule set.
[**TweetsFullarchiveSearch**](TweetsApi.md#TweetsFullarchiveSearch) | **Get** /2/tweets/search/all | Returns Tweets that match a search query.
[**TweetsRecentSearch**](TweetsApi.md#TweetsRecentSearch) | **Get** /2/tweets/search/recent | Returns Tweets from the last 7 days that match a search query.



## AddOrDeleteRules

> AddOrDeleteRulesResponse AddOrDeleteRules(ctx, addOrDeleteRulesRequest, optional)

Add or delete rules from a user's active rule set.

Add or delete rules from a user's active rule set. Users can provide unique, optionally tagged rules to add. Users can delete their entire rule set or a subset specified by rule ids or values.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addOrDeleteRulesRequest** | [**AddOrDeleteRulesRequest**](AddOrDeleteRulesRequest.md)|  | 
 **optional** | ***AddOrDeleteRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOrDeleteRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **optional.Bool**| Dry Run can be used with both the add and delete action, with the expected result given, but without actually taking any action in the system (meaning the end state will always be as it was when the request was submitted). This is particularly useful to validate rule changes. | 

### Return type

[**AddOrDeleteRulesResponse**](AddOrDeleteRulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTweetById

> SingleTweetLookupResponse FindTweetById(ctx, id, optional)

Returns hydrated Tweet objects

Returns a variety of information about the Tweet specified by the requested ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| A single Tweet ID. | 
 **optional** | ***FindTweetByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindTweetByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**SingleTweetLookupResponse**](SingleTweetLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTweetsById

> TweetLookupResponse FindTweetsById(ctx, ids, optional)

Returns hydrated Tweet objects

Returns a variety of information about the Tweet specified by the requested ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | [**[]string**](string.md)| A comma separated list of Tweet IDs. Up to 100 are allowed in a single request. | 
 **optional** | ***FindTweetsByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindTweetsByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**TweetLookupResponse**](TweetLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRules

> InlineResponse2002 GetRules(ctx, optional)

Returns rules from a user's active rule set.

Returns rules from a user's active rule set. Users can fetch all of their rules or a subset, specified by the provided rule ids.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)| A comma-separated list of Rule IDs. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideReplyById

> InlineResponse200 HideReplyById(ctx, id, optional)

Hides or unhides a reply to an owned conversation.

Tweet ID in the path is that of the reply to hide or unhide.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The ID of the reply that you want to hide or unhide. | 
 **optional** | ***HideReplyByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HideReplyByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SampleStream

> StreamingTweet SampleStream(ctx, optional)

Streams a deterministic 1% of public tweets.

Streams a deterministic 1% of public tweets.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SampleStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SampleStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**StreamingTweet**](StreamingTweet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStream

> FilteredStreamingTweet SearchStream(ctx, optional)

Streams tweets matching a user's active rule set.

Streams tweets matching a user's active rule set.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**FilteredStreamingTweet**](FilteredStreamingTweet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TweetsFullarchiveSearch

> InlineResponse2001 TweetsFullarchiveSearch(ctx, query, optional)

Returns Tweets that match a search query.

Returns Tweets that match a search query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | **string**| One query/rule/filter for matching Tweets. Up to 1024 characters. | 
 **optional** | ***TweetsFullarchiveSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TweetsFullarchiveSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | **optional.String**| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | **optional.String**| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **maxResults** | **optional.Int32**| The maximum number of search results to be returned by a request. | [default to 10]
 **nextToken** | **optional.String**| This parameter is used to get the next &#39;page&#39; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TweetsRecentSearch

> TweetSearchResponse TweetsRecentSearch(ctx, query, optional)

Returns Tweets from the last 7 days that match a search query.

Returns Tweets from the last 7 days that match a search query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | **string**| One query/rule/filter for matching Tweets. Up to 512 characters. | 
 **optional** | ***TweetsRecentSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TweetsRecentSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The oldest UTC timestamp (from most recent 7 days) from which the Tweets will be provided. Timestamp is in second granularity and is inclusive (i.e. 12:00:01 includes the first second of the minute). | 
 **endTime** | **optional.Time**| YYYY-MM-DDTHH:mm:ssZ. The newest, most recent UTC timestamp to which the Tweets will be provided. Timestamp is in second granularity and is exclusive (i.e. 12:00:01 excludes the first second of the minute). | 
 **sinceId** | **optional.String**| Returns results with a Tweet ID greater than (that is, more recent than) the specified ID. | 
 **untilId** | **optional.String**| Returns results with a Tweet ID less than (that is, older than) the specified ID. | 
 **maxResults** | **optional.Int32**| The maximum number of search results to be returned by a request. | [default to 10]
 **nextToken** | **optional.String**| This parameter is used to get the next &#39;page&#39; of results. The value used with the parameter is pulled directly from the response provided by the API, and should not be modified. | 
 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 
 **mediaFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Media fields to display. | 
 **placeFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Place fields to display. | 
 **pollFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Poll fields to display. | 

### Return type

[**TweetSearchResponse**](TweetSearchResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

