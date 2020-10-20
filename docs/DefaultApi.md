# \DefaultApi

All URIs are relative to *http://galactus.player/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRoomByCode**](DefaultApi.md#GetRoomByCode) | **Get** /room/{code} | 



## GetRoomByCode

> Room GetRoomByCode(ctx, code)



Returns a Room

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**| ID of Room to return | 

### Return type

[**Room**](Room.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

