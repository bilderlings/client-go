# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | The timestamp when the account was created | [optional] 
**Email** | **string** | Optional email address associated with the account | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) | The timestamp of the last update to the account metadata itself (not users or creds) | [optional] 
**Name** | **string** | The account identifier, not updatable after creation | 
**State** | **string** | State of the account. Disabled accounts prevent member users from logging in, deleting accounts are disabled and pending deletion and will be removed once all owned resources are garbage collected by the system | [optional] 
**Type** | **string** | The user type (admin vs user). If not specified in a POST request, &#39;user&#39; is default | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


