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

package lock

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

type DescribeMutexesResult struct {
	/** ミューテックスのリスト */
	Items *[]*Mutex `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeMutexesResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeMutexesAsyncResult struct {
	result *DescribeMutexesResult
	err    error
}

type DescribeMutexesByUserIdResult struct {
	/** ミューテックスのリスト */
	Items *[]*Mutex `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeMutexesByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeMutexesByUserIdAsyncResult struct {
	result *DescribeMutexesByUserIdResult
	err    error
}

type LockResult struct {
	/** ミューテックス */
	Item *Mutex `json:"item"`
}

func (p *LockResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type LockAsyncResult struct {
	result *LockResult
	err    error
}

type LockByUserIdResult struct {
	/** ミューテックス */
	Item *Mutex `json:"item"`
}

func (p *LockByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type LockByUserIdAsyncResult struct {
	result *LockByUserIdResult
	err    error
}

type UnlockResult struct {
	/** ミューテックス */
	Item *Mutex `json:"item"`
}

func (p *UnlockResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UnlockAsyncResult struct {
	result *UnlockResult
	err    error
}

type UnlockByUserIdResult struct {
	/** ミューテックス */
	Item *Mutex `json:"item"`
}

func (p *UnlockByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UnlockByUserIdAsyncResult struct {
	result *UnlockByUserIdResult
	err    error
}

type GetMutexResult struct {
	/** ミューテックス */
	Item *Mutex `json:"item"`
}

func (p *GetMutexResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetMutexAsyncResult struct {
	result *GetMutexResult
	err    error
}

type GetMutexByUserIdResult struct {
	/** ミューテックス */
	Item *Mutex `json:"item"`
}

func (p *GetMutexByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetMutexByUserIdAsyncResult struct {
	result *GetMutexByUserIdResult
	err    error
}

type DeleteMutexByUserIdResult struct {
	/** ミューテックス */
	Item *Mutex `json:"item"`
}

func (p *DeleteMutexByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteMutexByUserIdAsyncResult struct {
	result *DeleteMutexByUserIdResult
	err    error
}
