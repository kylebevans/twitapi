# \SearchApi

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TweetsFullarchiveSearch**](SearchApi.md#TweetsFullarchiveSearch) | **Get** /2/tweets/search/all | Returns Tweets that match a search query.
[**TweetsRecentSearch**](SearchApi.md#TweetsRecentSearch) | **Get** /2/tweets/search/recent | Returns Tweets from the last 7 days that match a search query.



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

