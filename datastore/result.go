/*
Copyright 2016 Game Server Services, Inc. or its affiliates. All Rights
Reserved.

Licensed under the Apache License, Version 2.0 (the "License").
You may not use this file except in compliance with the License.
A copy of the License is located at

 http://www.apache.org/licenses/LICENSE-2.0

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/

package datastore

import "github.com/gs2io/gs2-golang-sdk/core"

type DescribeNamespacesResult struct {
	Items         []Namespace `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromDict(data map[string]interface{}) DescribeNamespacesResult {
	return DescribeNamespacesResult{
		Items:         CastNamespaces(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNamespacesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeNamespacesResult) Pointer() *DescribeNamespacesResult {
	return &p
}

type CreateNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

func NewCreateNamespaceResultFromDict(data map[string]interface{}) CreateNamespaceResult {
	return CreateNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateNamespaceResult) Pointer() *CreateNamespaceResult {
	return &p
}

type GetNamespaceStatusResult struct {
	Status *string `json:"status"`
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

func NewGetNamespaceStatusResultFromDict(data map[string]interface{}) GetNamespaceStatusResult {
	return GetNamespaceStatusResult{
		Status: core.CastString(data["status"]),
	}
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"status": p.Status,
	}
}

func (p GetNamespaceStatusResult) Pointer() *GetNamespaceStatusResult {
	return &p
}

type GetNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

func NewGetNamespaceResultFromDict(data map[string]interface{}) GetNamespaceResult {
	return GetNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetNamespaceResult) Pointer() *GetNamespaceResult {
	return &p
}

type UpdateNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

func NewUpdateNamespaceResultFromDict(data map[string]interface{}) UpdateNamespaceResult {
	return UpdateNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateNamespaceResult) Pointer() *UpdateNamespaceResult {
	return &p
}

type DeleteNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
	return DeleteNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
	return &p
}

type DescribeDataObjectsResult struct {
	Items         []DataObject `json:"items"`
	NextPageToken *string      `json:"nextPageToken"`
}

type DescribeDataObjectsAsyncResult struct {
	result *DescribeDataObjectsResult
	err    error
}

func NewDescribeDataObjectsResultFromDict(data map[string]interface{}) DescribeDataObjectsResult {
	return DescribeDataObjectsResult{
		Items:         CastDataObjects(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeDataObjectsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDataObjectsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeDataObjectsResult) Pointer() *DescribeDataObjectsResult {
	return &p
}

type DescribeDataObjectsByUserIdResult struct {
	Items         []DataObject `json:"items"`
	NextPageToken *string      `json:"nextPageToken"`
}

type DescribeDataObjectsByUserIdAsyncResult struct {
	result *DescribeDataObjectsByUserIdResult
	err    error
}

func NewDescribeDataObjectsByUserIdResultFromDict(data map[string]interface{}) DescribeDataObjectsByUserIdResult {
	return DescribeDataObjectsByUserIdResult{
		Items:         CastDataObjects(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeDataObjectsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDataObjectsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeDataObjectsByUserIdResult) Pointer() *DescribeDataObjectsByUserIdResult {
	return &p
}

type PrepareUploadResult struct {
	Item      *DataObject `json:"item"`
	UploadUrl *string     `json:"uploadUrl"`
}

type PrepareUploadAsyncResult struct {
	result *PrepareUploadResult
	err    error
}

func NewPrepareUploadResultFromDict(data map[string]interface{}) PrepareUploadResult {
	return PrepareUploadResult{
		Item:      NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		UploadUrl: core.CastString(data["uploadUrl"]),
	}
}

func (p PrepareUploadResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"uploadUrl": p.UploadUrl,
	}
}

func (p PrepareUploadResult) Pointer() *PrepareUploadResult {
	return &p
}

type PrepareUploadByUserIdResult struct {
	Item      *DataObject `json:"item"`
	UploadUrl *string     `json:"uploadUrl"`
}

type PrepareUploadByUserIdAsyncResult struct {
	result *PrepareUploadByUserIdResult
	err    error
}

func NewPrepareUploadByUserIdResultFromDict(data map[string]interface{}) PrepareUploadByUserIdResult {
	return PrepareUploadByUserIdResult{
		Item:      NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		UploadUrl: core.CastString(data["uploadUrl"]),
	}
}

func (p PrepareUploadByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"uploadUrl": p.UploadUrl,
	}
}

func (p PrepareUploadByUserIdResult) Pointer() *PrepareUploadByUserIdResult {
	return &p
}

type UpdateDataObjectResult struct {
	Item *DataObject `json:"item"`
}

type UpdateDataObjectAsyncResult struct {
	result *UpdateDataObjectResult
	err    error
}

func NewUpdateDataObjectResultFromDict(data map[string]interface{}) UpdateDataObjectResult {
	return UpdateDataObjectResult{
		Item: NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateDataObjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateDataObjectResult) Pointer() *UpdateDataObjectResult {
	return &p
}

type UpdateDataObjectByUserIdResult struct {
	Item *DataObject `json:"item"`
}

type UpdateDataObjectByUserIdAsyncResult struct {
	result *UpdateDataObjectByUserIdResult
	err    error
}

func NewUpdateDataObjectByUserIdResultFromDict(data map[string]interface{}) UpdateDataObjectByUserIdResult {
	return UpdateDataObjectByUserIdResult{
		Item: NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateDataObjectByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateDataObjectByUserIdResult) Pointer() *UpdateDataObjectByUserIdResult {
	return &p
}

type PrepareReUploadResult struct {
	Item      *DataObject `json:"item"`
	UploadUrl *string     `json:"uploadUrl"`
}

type PrepareReUploadAsyncResult struct {
	result *PrepareReUploadResult
	err    error
}

func NewPrepareReUploadResultFromDict(data map[string]interface{}) PrepareReUploadResult {
	return PrepareReUploadResult{
		Item:      NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		UploadUrl: core.CastString(data["uploadUrl"]),
	}
}

func (p PrepareReUploadResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"uploadUrl": p.UploadUrl,
	}
}

func (p PrepareReUploadResult) Pointer() *PrepareReUploadResult {
	return &p
}

type PrepareReUploadByUserIdResult struct {
	Item      *DataObject `json:"item"`
	UploadUrl *string     `json:"uploadUrl"`
}

type PrepareReUploadByUserIdAsyncResult struct {
	result *PrepareReUploadByUserIdResult
	err    error
}

func NewPrepareReUploadByUserIdResultFromDict(data map[string]interface{}) PrepareReUploadByUserIdResult {
	return PrepareReUploadByUserIdResult{
		Item:      NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		UploadUrl: core.CastString(data["uploadUrl"]),
	}
}

func (p PrepareReUploadByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"uploadUrl": p.UploadUrl,
	}
}

func (p PrepareReUploadByUserIdResult) Pointer() *PrepareReUploadByUserIdResult {
	return &p
}

type DoneUploadResult struct {
	Item *DataObject `json:"item"`
}

type DoneUploadAsyncResult struct {
	result *DoneUploadResult
	err    error
}

func NewDoneUploadResultFromDict(data map[string]interface{}) DoneUploadResult {
	return DoneUploadResult{
		Item: NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DoneUploadResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DoneUploadResult) Pointer() *DoneUploadResult {
	return &p
}

type DoneUploadByUserIdResult struct {
	Item *DataObject `json:"item"`
}

type DoneUploadByUserIdAsyncResult struct {
	result *DoneUploadByUserIdResult
	err    error
}

func NewDoneUploadByUserIdResultFromDict(data map[string]interface{}) DoneUploadByUserIdResult {
	return DoneUploadByUserIdResult{
		Item: NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DoneUploadByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DoneUploadByUserIdResult) Pointer() *DoneUploadByUserIdResult {
	return &p
}

type DeleteDataObjectResult struct {
	Item *DataObject `json:"item"`
}

type DeleteDataObjectAsyncResult struct {
	result *DeleteDataObjectResult
	err    error
}

func NewDeleteDataObjectResultFromDict(data map[string]interface{}) DeleteDataObjectResult {
	return DeleteDataObjectResult{
		Item: NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteDataObjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteDataObjectResult) Pointer() *DeleteDataObjectResult {
	return &p
}

type DeleteDataObjectByUserIdResult struct {
	Item *DataObject `json:"item"`
}

type DeleteDataObjectByUserIdAsyncResult struct {
	result *DeleteDataObjectByUserIdResult
	err    error
}

func NewDeleteDataObjectByUserIdResultFromDict(data map[string]interface{}) DeleteDataObjectByUserIdResult {
	return DeleteDataObjectByUserIdResult{
		Item: NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteDataObjectByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteDataObjectByUserIdResult) Pointer() *DeleteDataObjectByUserIdResult {
	return &p
}

type PrepareDownloadResult struct {
	Item          *DataObject `json:"item"`
	FileUrl       *string     `json:"fileUrl"`
	ContentLength *int64      `json:"contentLength"`
}

type PrepareDownloadAsyncResult struct {
	result *PrepareDownloadResult
	err    error
}

func NewPrepareDownloadResultFromDict(data map[string]interface{}) PrepareDownloadResult {
	return PrepareDownloadResult{
		Item:          NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		FileUrl:       core.CastString(data["fileUrl"]),
		ContentLength: core.CastInt64(data["contentLength"]),
	}
}

func (p PrepareDownloadResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
	}
}

func (p PrepareDownloadResult) Pointer() *PrepareDownloadResult {
	return &p
}

type PrepareDownloadByUserIdResult struct {
	Item          *DataObject `json:"item"`
	FileUrl       *string     `json:"fileUrl"`
	ContentLength *int64      `json:"contentLength"`
}

type PrepareDownloadByUserIdAsyncResult struct {
	result *PrepareDownloadByUserIdResult
	err    error
}

func NewPrepareDownloadByUserIdResultFromDict(data map[string]interface{}) PrepareDownloadByUserIdResult {
	return PrepareDownloadByUserIdResult{
		Item:          NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		FileUrl:       core.CastString(data["fileUrl"]),
		ContentLength: core.CastInt64(data["contentLength"]),
	}
}

func (p PrepareDownloadByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
	}
}

func (p PrepareDownloadByUserIdResult) Pointer() *PrepareDownloadByUserIdResult {
	return &p
}

type PrepareDownloadByGenerationResult struct {
	Item          *DataObject `json:"item"`
	FileUrl       *string     `json:"fileUrl"`
	ContentLength *int64      `json:"contentLength"`
}

type PrepareDownloadByGenerationAsyncResult struct {
	result *PrepareDownloadByGenerationResult
	err    error
}

func NewPrepareDownloadByGenerationResultFromDict(data map[string]interface{}) PrepareDownloadByGenerationResult {
	return PrepareDownloadByGenerationResult{
		Item:          NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		FileUrl:       core.CastString(data["fileUrl"]),
		ContentLength: core.CastInt64(data["contentLength"]),
	}
}

func (p PrepareDownloadByGenerationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
	}
}

func (p PrepareDownloadByGenerationResult) Pointer() *PrepareDownloadByGenerationResult {
	return &p
}

type PrepareDownloadByGenerationAndUserIdResult struct {
	Item          *DataObject `json:"item"`
	FileUrl       *string     `json:"fileUrl"`
	ContentLength *int64      `json:"contentLength"`
}

type PrepareDownloadByGenerationAndUserIdAsyncResult struct {
	result *PrepareDownloadByGenerationAndUserIdResult
	err    error
}

func NewPrepareDownloadByGenerationAndUserIdResultFromDict(data map[string]interface{}) PrepareDownloadByGenerationAndUserIdResult {
	return PrepareDownloadByGenerationAndUserIdResult{
		Item:          NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		FileUrl:       core.CastString(data["fileUrl"]),
		ContentLength: core.CastInt64(data["contentLength"]),
	}
}

func (p PrepareDownloadByGenerationAndUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
	}
}

func (p PrepareDownloadByGenerationAndUserIdResult) Pointer() *PrepareDownloadByGenerationAndUserIdResult {
	return &p
}

type PrepareDownloadOwnDataResult struct {
	Item          *DataObject `json:"item"`
	FileUrl       *string     `json:"fileUrl"`
	ContentLength *int64      `json:"contentLength"`
}

type PrepareDownloadOwnDataAsyncResult struct {
	result *PrepareDownloadOwnDataResult
	err    error
}

func NewPrepareDownloadOwnDataResultFromDict(data map[string]interface{}) PrepareDownloadOwnDataResult {
	return PrepareDownloadOwnDataResult{
		Item:          NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		FileUrl:       core.CastString(data["fileUrl"]),
		ContentLength: core.CastInt64(data["contentLength"]),
	}
}

func (p PrepareDownloadOwnDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
	}
}

func (p PrepareDownloadOwnDataResult) Pointer() *PrepareDownloadOwnDataResult {
	return &p
}

type PrepareDownloadByUserIdAndDataObjectNameResult struct {
	Item          *DataObject `json:"item"`
	FileUrl       *string     `json:"fileUrl"`
	ContentLength *int64      `json:"contentLength"`
}

type PrepareDownloadByUserIdAndDataObjectNameAsyncResult struct {
	result *PrepareDownloadByUserIdAndDataObjectNameResult
	err    error
}

func NewPrepareDownloadByUserIdAndDataObjectNameResultFromDict(data map[string]interface{}) PrepareDownloadByUserIdAndDataObjectNameResult {
	return PrepareDownloadByUserIdAndDataObjectNameResult{
		Item:          NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		FileUrl:       core.CastString(data["fileUrl"]),
		ContentLength: core.CastInt64(data["contentLength"]),
	}
}

func (p PrepareDownloadByUserIdAndDataObjectNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
	}
}

func (p PrepareDownloadByUserIdAndDataObjectNameResult) Pointer() *PrepareDownloadByUserIdAndDataObjectNameResult {
	return &p
}

type PrepareDownloadOwnDataByGenerationResult struct {
	Item          *DataObject `json:"item"`
	FileUrl       *string     `json:"fileUrl"`
	ContentLength *int64      `json:"contentLength"`
}

type PrepareDownloadOwnDataByGenerationAsyncResult struct {
	result *PrepareDownloadOwnDataByGenerationResult
	err    error
}

func NewPrepareDownloadOwnDataByGenerationResultFromDict(data map[string]interface{}) PrepareDownloadOwnDataByGenerationResult {
	return PrepareDownloadOwnDataByGenerationResult{
		Item:          NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		FileUrl:       core.CastString(data["fileUrl"]),
		ContentLength: core.CastInt64(data["contentLength"]),
	}
}

func (p PrepareDownloadOwnDataByGenerationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
	}
}

func (p PrepareDownloadOwnDataByGenerationResult) Pointer() *PrepareDownloadOwnDataByGenerationResult {
	return &p
}

type PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult struct {
	Item          *DataObject `json:"item"`
	FileUrl       *string     `json:"fileUrl"`
	ContentLength *int64      `json:"contentLength"`
}

type PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult struct {
	result *PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult
	err    error
}

func NewPrepareDownloadByUserIdAndDataObjectNameAndGenerationResultFromDict(data map[string]interface{}) PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult {
	return PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult{
		Item:          NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
		FileUrl:       core.CastString(data["fileUrl"]),
		ContentLength: core.CastInt64(data["contentLength"]),
	}
}

func (p PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
	}
}

func (p PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult) Pointer() *PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult {
	return &p
}

type RestoreDataObjectResult struct {
	Item *DataObject `json:"item"`
}

type RestoreDataObjectAsyncResult struct {
	result *RestoreDataObjectResult
	err    error
}

func NewRestoreDataObjectResultFromDict(data map[string]interface{}) RestoreDataObjectResult {
	return RestoreDataObjectResult{
		Item: NewDataObjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p RestoreDataObjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p RestoreDataObjectResult) Pointer() *RestoreDataObjectResult {
	return &p
}

type DescribeDataObjectHistoriesResult struct {
	Items         []DataObjectHistory `json:"items"`
	NextPageToken *string             `json:"nextPageToken"`
}

type DescribeDataObjectHistoriesAsyncResult struct {
	result *DescribeDataObjectHistoriesResult
	err    error
}

func NewDescribeDataObjectHistoriesResultFromDict(data map[string]interface{}) DescribeDataObjectHistoriesResult {
	return DescribeDataObjectHistoriesResult{
		Items:         CastDataObjectHistories(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeDataObjectHistoriesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDataObjectHistoriesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeDataObjectHistoriesResult) Pointer() *DescribeDataObjectHistoriesResult {
	return &p
}

type DescribeDataObjectHistoriesByUserIdResult struct {
	Items         []DataObjectHistory `json:"items"`
	NextPageToken *string             `json:"nextPageToken"`
}

type DescribeDataObjectHistoriesByUserIdAsyncResult struct {
	result *DescribeDataObjectHistoriesByUserIdResult
	err    error
}

func NewDescribeDataObjectHistoriesByUserIdResultFromDict(data map[string]interface{}) DescribeDataObjectHistoriesByUserIdResult {
	return DescribeDataObjectHistoriesByUserIdResult{
		Items:         CastDataObjectHistories(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeDataObjectHistoriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDataObjectHistoriesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeDataObjectHistoriesByUserIdResult) Pointer() *DescribeDataObjectHistoriesByUserIdResult {
	return &p
}

type GetDataObjectHistoryResult struct {
	Item *DataObjectHistory `json:"item"`
}

type GetDataObjectHistoryAsyncResult struct {
	result *GetDataObjectHistoryResult
	err    error
}

func NewGetDataObjectHistoryResultFromDict(data map[string]interface{}) GetDataObjectHistoryResult {
	return GetDataObjectHistoryResult{
		Item: NewDataObjectHistoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetDataObjectHistoryResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetDataObjectHistoryResult) Pointer() *GetDataObjectHistoryResult {
	return &p
}

type GetDataObjectHistoryByUserIdResult struct {
	Item *DataObjectHistory `json:"item"`
}

type GetDataObjectHistoryByUserIdAsyncResult struct {
	result *GetDataObjectHistoryByUserIdResult
	err    error
}

func NewGetDataObjectHistoryByUserIdResultFromDict(data map[string]interface{}) GetDataObjectHistoryByUserIdResult {
	return GetDataObjectHistoryByUserIdResult{
		Item: NewDataObjectHistoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetDataObjectHistoryByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetDataObjectHistoryByUserIdResult) Pointer() *GetDataObjectHistoryByUserIdResult {
	return &p
}
