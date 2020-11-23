# \RulesApi

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrDeleteRules**](RulesApi.md#AddOrDeleteRules) | **Post** /2/tweets/search/stream/rules | Add or delete rules from a user&#39;s active rule set.
[**GetRules**](RulesApi.md#GetRules) | **Get** /2/tweets/search/stream/rules | Returns rules from a user&#39;s active rule set.



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

