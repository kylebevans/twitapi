# \UsersApi

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindUserById**](UsersApi.md#FindUserById) | **Get** /2/users/{id} | Return details for the specified users
[**FindUserByUsername**](UsersApi.md#FindUserByUsername) | **Get** /2/users/by/username/{username} | Return details for the specified users
[**FindUsersById**](UsersApi.md#FindUsersById) | **Get** /2/users | Return details for the specified users
[**FindUsersByUsername**](UsersApi.md#FindUsersByUsername) | **Get** /2/users/by | Return details for the specified users



## FindUserById

> SingleUserLookupResponse FindUserById(ctx, id, optional)

Return details for the specified users

This endpoint returns information about a user. Specify user by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| Required. A User ID. | 
 **optional** | ***FindUserByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindUserByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**SingleUserLookupResponse**](SingleUserLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUserByUsername

> SingleUserLookupResponse FindUserByUsername(ctx, username, optional)

Return details for the specified users

This endpoint returns information about a user. Specify user by username.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string**| Required. A username. | 
 **optional** | ***FindUserByUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindUserByUsernameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**SingleUserLookupResponse**](SingleUserLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsersById

> UserLookupResponse FindUsersById(ctx, ids, optional)

Return details for the specified users

This endpoint returns information about users. Specify users by their ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | [**[]string**](string.md)| Required. A list of User IDs, comma-separated. You can specify up to 100 IDs. | 
 **optional** | ***FindUsersByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindUsersByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**UserLookupResponse**](UserLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsersByUsername

> UserLookupResponse FindUsersByUsername(ctx, usernames, optional)

Return details for the specified users

This endpoint returns information about users. Specify users by their username.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usernames** | [**[]string**](string.md)| Required . A list of usernames, comma-separated. You can specify up to 100 usernames. | 
 **optional** | ***FindUsersByUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindUsersByUsernameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expansions** | [**optional.Interface of []string**](string.md)| A comma separated list of fields to expand. | 
 **tweetFields** | [**optional.Interface of []string**](string.md)| A comma separated list of Tweet fields to display. | 
 **userFields** | [**optional.Interface of []string**](string.md)| A comma separated list of User fields to display. | 

### Return type

[**UserLookupResponse**](UserLookupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

