# RepoSyncEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Status** | **string** |  | [optional] 
**PullRequestUrl** | **string** | URL of the pull request created on export | [optional] 
**AutoImport** | **bool** | Whether the import was triggered by the repo push event | [optional] 
**Errors** | [**[]RepoSyncEventErrorsInner**](RepoSyncEventErrorsInner.md) | List of error messages, in case of failure | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


