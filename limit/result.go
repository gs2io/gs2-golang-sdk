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

package limit

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

type DescribeCountersResult struct {
    /** カウンターのリスト */
	Items         *[]*Counter	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCountersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Counter, 0)
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

type DescribeCountersAsyncResult struct {
	result *DescribeCountersResult
	err    error
}

type DescribeCountersByUserIdResult struct {
    /** カウンターのリスト */
	Items         *[]*Counter	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCountersByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Counter, 0)
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

type DescribeCountersByUserIdAsyncResult struct {
	result *DescribeCountersByUserIdResult
	err    error
}

type GetCounterResult struct {
    /** カウンター */
	Item         *Counter	`json:"item"`
}

func (p *GetCounterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCounterAsyncResult struct {
	result *GetCounterResult
	err    error
}

type GetCounterByUserIdResult struct {
    /** カウンター */
	Item         *Counter	`json:"item"`
}

func (p *GetCounterByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCounterByUserIdAsyncResult struct {
	result *GetCounterByUserIdResult
	err    error
}

type CountUpResult struct {
    /** カウントを増やしたカウンター */
	Item         *Counter	`json:"item"`
}

func (p *CountUpResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CountUpAsyncResult struct {
	result *CountUpResult
	err    error
}

type CountUpByUserIdResult struct {
    /** カウントを増やしたカウンター */
	Item         *Counter	`json:"item"`
}

func (p *CountUpByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CountUpByUserIdAsyncResult struct {
	result *CountUpByUserIdResult
	err    error
}

type DeleteCounterByUserIdResult struct {
    /** カウンター */
	Item         *Counter	`json:"item"`
}

func (p *DeleteCounterByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteCounterByUserIdAsyncResult struct {
	result *DeleteCounterByUserIdResult
	err    error
}

type CountUpByStampTaskResult struct {
    /** カウントを増やしたカウンター */
	Item         *Counter	`json:"item"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *core.String	`json:"newContextStack"`
}

func (p *CountUpByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.NewContextStack != nil {
        data["newContextStack"] = p.NewContextStack
    }
    return &data
}

type CountUpByStampTaskAsyncResult struct {
	result *CountUpByStampTaskResult
	err    error
}

type DeleteByStampSheetResult struct {
    /** カウンター */
	Item         *Counter	`json:"item"`
}

func (p *DeleteByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteByStampSheetAsyncResult struct {
	result *DeleteByStampSheetResult
	err    error
}

type DescribeLimitModelMastersResult struct {
    /** 回数制限の種類マスターのリスト */
	Items         *[]*LimitModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeLimitModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*LimitModelMaster, 0)
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

type DescribeLimitModelMastersAsyncResult struct {
	result *DescribeLimitModelMastersResult
	err    error
}

type CreateLimitModelMasterResult struct {
    /** 作成した回数制限の種類マスター */
	Item         *LimitModelMaster	`json:"item"`
}

func (p *CreateLimitModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateLimitModelMasterAsyncResult struct {
	result *CreateLimitModelMasterResult
	err    error
}

type GetLimitModelMasterResult struct {
    /** 回数制限の種類マスター */
	Item         *LimitModelMaster	`json:"item"`
}

func (p *GetLimitModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetLimitModelMasterAsyncResult struct {
	result *GetLimitModelMasterResult
	err    error
}

type UpdateLimitModelMasterResult struct {
    /** 更新した回数制限の種類マスター */
	Item         *LimitModelMaster	`json:"item"`
}

func (p *UpdateLimitModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateLimitModelMasterAsyncResult struct {
	result *UpdateLimitModelMasterResult
	err    error
}

type DeleteLimitModelMasterResult struct {
    /** 削除した回数制限の種類マスター */
	Item         *LimitModelMaster	`json:"item"`
}

func (p *DeleteLimitModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteLimitModelMasterAsyncResult struct {
	result *DeleteLimitModelMasterResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効な回数制限設定 */
	Item         *CurrentLimitMaster	`json:"item"`
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

type GetCurrentLimitMasterResult struct {
    /** 現在有効な回数制限設定 */
	Item         *CurrentLimitMaster	`json:"item"`
}

func (p *GetCurrentLimitMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentLimitMasterAsyncResult struct {
	result *GetCurrentLimitMasterResult
	err    error
}

type UpdateCurrentLimitMasterResult struct {
    /** 更新した現在有効な回数制限設定 */
	Item         *CurrentLimitMaster	`json:"item"`
}

func (p *UpdateCurrentLimitMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentLimitMasterAsyncResult struct {
	result *UpdateCurrentLimitMasterResult
	err    error
}

type UpdateCurrentLimitMasterFromGitHubResult struct {
    /** 更新した現在有効な回数制限設定 */
	Item         *CurrentLimitMaster	`json:"item"`
}

func (p *UpdateCurrentLimitMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentLimitMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentLimitMasterFromGitHubResult
	err    error
}

type DescribeLimitModelsResult struct {
    /** 回数制限の種類のリスト */
	Items         *[]*LimitModel	`json:"items"`
}

func (p *DescribeLimitModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*LimitModel, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeLimitModelsAsyncResult struct {
	result *DescribeLimitModelsResult
	err    error
}

type GetLimitModelResult struct {
    /** 回数制限の種類 */
	Item         *LimitModel	`json:"item"`
}

func (p *GetLimitModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetLimitModelAsyncResult struct {
	result *GetLimitModelResult
	err    error
}
