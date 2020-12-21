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

package dictionary

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         *[]*Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Namespace, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

type CreateNamespaceResult struct {
    /** 作成したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *CreateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

type GetNamespaceStatusResult struct {
    /** None */
	Status         *core.String	`json:"status"`
}

func (p *GetNamespaceStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Status != nil {
        data["status"] = p.Status
    }
    return &data
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

type GetNamespaceResult struct {
    /** ネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *GetNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

type UpdateNamespaceResult struct {
    /** 更新したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *UpdateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

type DeleteNamespaceResult struct {
    /** 削除したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeEntryModelsResult struct {
    /** エントリーモデルのリスト */
	Items         *[]*EntryModel	`json:"items"`
}

func (p *DescribeEntryModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*EntryModel, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeEntryModelsAsyncResult struct {
	result *DescribeEntryModelsResult
	err    error
}

type GetEntryModelResult struct {
    /** エントリーモデル */
	Item         *EntryModel	`json:"item"`
}

func (p *GetEntryModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetEntryModelAsyncResult struct {
	result *GetEntryModelResult
	err    error
}

type DescribeEntryModelMastersResult struct {
    /** エントリーモデルマスターのリスト */
	Items         *[]*EntryModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeEntryModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*EntryModelMaster, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeEntryModelMastersAsyncResult struct {
	result *DescribeEntryModelMastersResult
	err    error
}

type CreateEntryModelMasterResult struct {
    /** 作成したエントリーモデルマスター */
	Item         *EntryModelMaster	`json:"item"`
}

func (p *CreateEntryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateEntryModelMasterAsyncResult struct {
	result *CreateEntryModelMasterResult
	err    error
}

type GetEntryModelMasterResult struct {
    /** エントリーモデルマスター */
	Item         *EntryModelMaster	`json:"item"`
}

func (p *GetEntryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetEntryModelMasterAsyncResult struct {
	result *GetEntryModelMasterResult
	err    error
}

type UpdateEntryModelMasterResult struct {
    /** 更新したエントリーモデルマスター */
	Item         *EntryModelMaster	`json:"item"`
}

func (p *UpdateEntryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateEntryModelMasterAsyncResult struct {
	result *UpdateEntryModelMasterResult
	err    error
}

type DeleteEntryModelMasterResult struct {
    /** 削除したエントリーモデルマスター */
	Item         *EntryModelMaster	`json:"item"`
}

func (p *DeleteEntryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteEntryModelMasterAsyncResult struct {
	result *DeleteEntryModelMasterResult
	err    error
}

type DescribeEntriesResult struct {
    /** エントリーのリスト */
	Items         *[]*Entry	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeEntriesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Entry, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeEntriesAsyncResult struct {
	result *DescribeEntriesResult
	err    error
}

type DescribeEntriesByUserIdResult struct {
    /** エントリーのリスト */
	Items         *[]*Entry	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeEntriesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Entry, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeEntriesByUserIdAsyncResult struct {
	result *DescribeEntriesByUserIdResult
	err    error
}

type AddEntriesByUserIdResult struct {
    /** 登録した{model_name}のリスト */
	Items         *[]*Entry	`json:"items"`
}

func (p *AddEntriesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Entry, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type AddEntriesByUserIdAsyncResult struct {
	result *AddEntriesByUserIdResult
	err    error
}

type GetEntryResult struct {
    /** エントリー */
	Item         *Entry	`json:"item"`
}

func (p *GetEntryResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetEntryAsyncResult struct {
	result *GetEntryResult
	err    error
}

type GetEntryByUserIdResult struct {
    /** エントリー */
	Item         *Entry	`json:"item"`
}

func (p *GetEntryByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetEntryByUserIdAsyncResult struct {
	result *GetEntryByUserIdResult
	err    error
}

type GetEntryWithSignatureResult struct {
    /** エントリー */
	Item         *Entry	`json:"item"`
    /** 署名対象のエントリー情報 */
	Body         *core.String	`json:"body"`
    /** 署名 */
	Signature         *core.String	`json:"signature"`
}

func (p *GetEntryWithSignatureResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.Body != nil {
        data["body"] = p.Body
    }
    if p.Signature != nil {
        data["signature"] = p.Signature
    }
    return &data
}

type GetEntryWithSignatureAsyncResult struct {
	result *GetEntryWithSignatureResult
	err    error
}

type GetEntryWithSignatureByUserIdResult struct {
    /** エントリー */
	Item         *Entry	`json:"item"`
    /** 署名対象のエントリー情報 */
	Body         *core.String	`json:"body"`
    /** 署名 */
	Signature         *core.String	`json:"signature"`
}

func (p *GetEntryWithSignatureByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.Body != nil {
        data["body"] = p.Body
    }
    if p.Signature != nil {
        data["signature"] = p.Signature
    }
    return &data
}

type GetEntryWithSignatureByUserIdAsyncResult struct {
	result *GetEntryWithSignatureByUserIdResult
	err    error
}

type ResetByUserIdResult struct {
}

func (p *ResetByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type ResetByUserIdAsyncResult struct {
	result *ResetByUserIdResult
	err    error
}

type AddEntriesByStampSheetResult struct {
    /** 追加後のエントリーのリスト */
	Items         *[]*Entry	`json:"items"`
}

func (p *AddEntriesByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Entry, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type AddEntriesByStampSheetAsyncResult struct {
	result *AddEntriesByStampSheetResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なエントリー設定 */
	Item         *CurrentEntryMaster	`json:"item"`
}

func (p *ExportMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

type GetCurrentEntryMasterResult struct {
    /** 現在有効なエントリー設定 */
	Item         *CurrentEntryMaster	`json:"item"`
}

func (p *GetCurrentEntryMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentEntryMasterAsyncResult struct {
	result *GetCurrentEntryMasterResult
	err    error
}

type UpdateCurrentEntryMasterResult struct {
    /** 更新した現在有効なエントリー設定 */
	Item         *CurrentEntryMaster	`json:"item"`
}

func (p *UpdateCurrentEntryMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentEntryMasterAsyncResult struct {
	result *UpdateCurrentEntryMasterResult
	err    error
}

type UpdateCurrentEntryMasterFromGitHubResult struct {
    /** 更新した現在有効なエントリー設定 */
	Item         *CurrentEntryMaster	`json:"item"`
}

func (p *UpdateCurrentEntryMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentEntryMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentEntryMasterFromGitHubResult
	err    error
}
