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

type DescribeNamespacesResult struct {
	/** ネームスペースのリスト */
	Items *[]*Namespace `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

type CreateNamespaceResult struct {
	/** 作成したネームスペース */
	Item *Namespace `json:"item"`
}

func (p *CreateNamespaceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

type GetNamespaceStatusResult struct {
	/** None */
	Status *string `json:"status"`
}

func (p *GetNamespaceStatusResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Status != nil {
		data["Status"] = p.Status
	}
	return &data
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

type GetNamespaceResult struct {
	/** ネームスペース */
	Item *Namespace `json:"item"`
}

func (p *GetNamespaceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

type UpdateNamespaceResult struct {
	/** 更新したネームスペース */
	Item *Namespace `json:"item"`
}

func (p *UpdateNamespaceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

type DeleteNamespaceResult struct {
	/** 削除したネームスペース */
	Item *Namespace `json:"item"`
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeDataObjectsResult struct {
	/** データオブジェクトのリスト */
	Items *[]*DataObject `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeDataObjectsResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeDataObjectsAsyncResult struct {
	result *DescribeDataObjectsResult
	err    error
}

type DescribeDataObjectsByUserIdResult struct {
	/** データオブジェクトのリスト */
	Items *[]*DataObject `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeDataObjectsByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeDataObjectsByUserIdAsyncResult struct {
	result *DescribeDataObjectsByUserIdResult
	err    error
}

type PrepareUploadResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** アップロード処理の実行に使用するURL */
	UploadUrl *string `json:"uploadUrl"`
}

func (p *PrepareUploadResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.UploadUrl != nil {
		data["UploadUrl"] = p.UploadUrl
	}
	return &data
}

type PrepareUploadAsyncResult struct {
	result *PrepareUploadResult
	err    error
}

type PrepareUploadByUserIdResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** アップロード処理の実行に使用するURL */
	UploadUrl *string `json:"uploadUrl"`
}

func (p *PrepareUploadByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.UploadUrl != nil {
		data["UploadUrl"] = p.UploadUrl
	}
	return &data
}

type PrepareUploadByUserIdAsyncResult struct {
	result *PrepareUploadByUserIdResult
	err    error
}

type UpdateDataObjectResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
}

func (p *UpdateDataObjectResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateDataObjectAsyncResult struct {
	result *UpdateDataObjectResult
	err    error
}

type UpdateDataObjectByUserIdResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
}

func (p *UpdateDataObjectByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateDataObjectByUserIdAsyncResult struct {
	result *UpdateDataObjectByUserIdResult
	err    error
}

type PrepareReUploadResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** アップロード処理の実行に使用するURL */
	UploadUrl *string `json:"uploadUrl"`
}

func (p *PrepareReUploadResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.UploadUrl != nil {
		data["UploadUrl"] = p.UploadUrl
	}
	return &data
}

type PrepareReUploadAsyncResult struct {
	result *PrepareReUploadResult
	err    error
}

type PrepareReUploadByUserIdResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** アップロード処理の実行に使用するURL */
	UploadUrl *string `json:"uploadUrl"`
}

func (p *PrepareReUploadByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.UploadUrl != nil {
		data["UploadUrl"] = p.UploadUrl
	}
	return &data
}

type PrepareReUploadByUserIdAsyncResult struct {
	result *PrepareReUploadByUserIdResult
	err    error
}

type DoneUploadResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
}

func (p *DoneUploadResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DoneUploadAsyncResult struct {
	result *DoneUploadResult
	err    error
}

type DoneUploadByUserIdResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
}

func (p *DoneUploadByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DoneUploadByUserIdAsyncResult struct {
	result *DoneUploadByUserIdResult
	err    error
}

type DeleteDataObjectResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
}

func (p *DeleteDataObjectResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteDataObjectAsyncResult struct {
	result *DeleteDataObjectResult
	err    error
}

type DeleteDataObjectByUserIdResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
}

func (p *DeleteDataObjectByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteDataObjectByUserIdAsyncResult struct {
	result *DeleteDataObjectByUserIdResult
	err    error
}

type PrepareDownloadResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** ファイルをダウンロードするためのURL */
	FileUrl *string `json:"fileUrl"`
	/** ファイルの容量 */
	ContentLength *int64 `json:"contentLength"`
}

func (p *PrepareDownloadResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.FileUrl != nil {
		data["FileUrl"] = p.FileUrl
	}
	if p.ContentLength != nil {
		data["ContentLength"] = p.ContentLength
	}
	return &data
}

type PrepareDownloadAsyncResult struct {
	result *PrepareDownloadResult
	err    error
}

type PrepareDownloadByUserIdResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** ファイルをダウンロードするためのURL */
	FileUrl *string `json:"fileUrl"`
	/** ファイルの容量 */
	ContentLength *int64 `json:"contentLength"`
}

func (p *PrepareDownloadByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.FileUrl != nil {
		data["FileUrl"] = p.FileUrl
	}
	if p.ContentLength != nil {
		data["ContentLength"] = p.ContentLength
	}
	return &data
}

type PrepareDownloadByUserIdAsyncResult struct {
	result *PrepareDownloadByUserIdResult
	err    error
}

type PrepareDownloadByGenerationResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** ファイルをダウンロードするためのURL */
	FileUrl *string `json:"fileUrl"`
	/** ファイルの容量 */
	ContentLength *int64 `json:"contentLength"`
}

func (p *PrepareDownloadByGenerationResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.FileUrl != nil {
		data["FileUrl"] = p.FileUrl
	}
	if p.ContentLength != nil {
		data["ContentLength"] = p.ContentLength
	}
	return &data
}

type PrepareDownloadByGenerationAsyncResult struct {
	result *PrepareDownloadByGenerationResult
	err    error
}

type PrepareDownloadByGenerationAndUserIdResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** ファイルをダウンロードするためのURL */
	FileUrl *string `json:"fileUrl"`
	/** ファイルの容量 */
	ContentLength *int64 `json:"contentLength"`
}

func (p *PrepareDownloadByGenerationAndUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.FileUrl != nil {
		data["FileUrl"] = p.FileUrl
	}
	if p.ContentLength != nil {
		data["ContentLength"] = p.ContentLength
	}
	return &data
}

type PrepareDownloadByGenerationAndUserIdAsyncResult struct {
	result *PrepareDownloadByGenerationAndUserIdResult
	err    error
}

type PrepareDownloadOwnDataResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** ファイルをダウンロードするためのURL */
	FileUrl *string `json:"fileUrl"`
	/** ファイルの容量 */
	ContentLength *int64 `json:"contentLength"`
}

func (p *PrepareDownloadOwnDataResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.FileUrl != nil {
		data["FileUrl"] = p.FileUrl
	}
	if p.ContentLength != nil {
		data["ContentLength"] = p.ContentLength
	}
	return &data
}

type PrepareDownloadOwnDataAsyncResult struct {
	result *PrepareDownloadOwnDataResult
	err    error
}

type PrepareDownloadByUserIdAndDataObjectNameResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** ファイルをダウンロードするためのURL */
	FileUrl *string `json:"fileUrl"`
	/** ファイルの容量 */
	ContentLength *int64 `json:"contentLength"`
}

func (p *PrepareDownloadByUserIdAndDataObjectNameResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.FileUrl != nil {
		data["FileUrl"] = p.FileUrl
	}
	if p.ContentLength != nil {
		data["ContentLength"] = p.ContentLength
	}
	return &data
}

type PrepareDownloadByUserIdAndDataObjectNameAsyncResult struct {
	result *PrepareDownloadByUserIdAndDataObjectNameResult
	err    error
}

type PrepareDownloadOwnDataByGenerationResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** ファイルをダウンロードするためのURL */
	FileUrl *string `json:"fileUrl"`
	/** ファイルの容量 */
	ContentLength *int64 `json:"contentLength"`
}

func (p *PrepareDownloadOwnDataByGenerationResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.FileUrl != nil {
		data["FileUrl"] = p.FileUrl
	}
	if p.ContentLength != nil {
		data["ContentLength"] = p.ContentLength
	}
	return &data
}

type PrepareDownloadOwnDataByGenerationAsyncResult struct {
	result *PrepareDownloadOwnDataByGenerationResult
	err    error
}

type PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult struct {
	/** データオブジェクト */
	Item *DataObject `json:"item"`
	/** ファイルをダウンロードするためのURL */
	FileUrl *string `json:"fileUrl"`
	/** ファイルの容量 */
	ContentLength *int64 `json:"contentLength"`
}

func (p *PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.FileUrl != nil {
		data["FileUrl"] = p.FileUrl
	}
	if p.ContentLength != nil {
		data["ContentLength"] = p.ContentLength
	}
	return &data
}

type PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult struct {
	result *PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult
	err    error
}

type DescribeDataObjectHistoriesResult struct {
	/** データオブジェクト履歴のリスト */
	Items *[]*DataObjectHistory `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeDataObjectHistoriesResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeDataObjectHistoriesAsyncResult struct {
	result *DescribeDataObjectHistoriesResult
	err    error
}

type DescribeDataObjectHistoriesByUserIdResult struct {
	/** データオブジェクト履歴のリスト */
	Items *[]*DataObjectHistory `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeDataObjectHistoriesByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeDataObjectHistoriesByUserIdAsyncResult struct {
	result *DescribeDataObjectHistoriesByUserIdResult
	err    error
}

type GetDataObjectHistoryResult struct {
	/** データオブジェクト履歴 */
	Item *DataObjectHistory `json:"item"`
}

func (p *GetDataObjectHistoryResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetDataObjectHistoryAsyncResult struct {
	result *GetDataObjectHistoryResult
	err    error
}

type GetDataObjectHistoryByUserIdResult struct {
	/** データオブジェクト履歴 */
	Item *DataObjectHistory `json:"item"`
}

func (p *GetDataObjectHistoryByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetDataObjectHistoryByUserIdAsyncResult struct {
	result *GetDataObjectHistoryByUserIdResult
	err    error
}
