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

package distributor

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         []Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Namespace, 0)
    	for _, item := range p.Items {
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
	Status         *string	`json:"status"`
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

type DescribeDistributorModelMastersResult struct {
    /** 配信設定マスターのリスト */
	Items         []DistributorModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeDistributorModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]DistributorModelMaster, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeDistributorModelMastersAsyncResult struct {
	result *DescribeDistributorModelMastersResult
	err    error
}

type CreateDistributorModelMasterResult struct {
    /** 作成した配信設定マスター */
	Item         *DistributorModelMaster	`json:"item"`
}

func (p *CreateDistributorModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateDistributorModelMasterAsyncResult struct {
	result *CreateDistributorModelMasterResult
	err    error
}

type GetDistributorModelMasterResult struct {
    /** 配信設定マスター */
	Item         *DistributorModelMaster	`json:"item"`
}

func (p *GetDistributorModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetDistributorModelMasterAsyncResult struct {
	result *GetDistributorModelMasterResult
	err    error
}

type UpdateDistributorModelMasterResult struct {
    /** 更新した配信設定マスター */
	Item         *DistributorModelMaster	`json:"item"`
}

func (p *UpdateDistributorModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateDistributorModelMasterAsyncResult struct {
	result *UpdateDistributorModelMasterResult
	err    error
}

type DeleteDistributorModelMasterResult struct {
    /** 削除した配信設定マスター */
	Item         *DistributorModelMaster	`json:"item"`
}

func (p *DeleteDistributorModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteDistributorModelMasterAsyncResult struct {
	result *DeleteDistributorModelMasterResult
	err    error
}

type DescribeDistributorModelsResult struct {
    /** 配信設定のリスト */
	Items         []DistributorModel	`json:"items"`
}

func (p *DescribeDistributorModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]DistributorModel, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeDistributorModelsAsyncResult struct {
	result *DescribeDistributorModelsResult
	err    error
}

type GetDistributorModelResult struct {
    /** 配信設定 */
	Item         *DistributorModel	`json:"item"`
}

func (p *GetDistributorModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetDistributorModelAsyncResult struct {
	result *GetDistributorModelResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効な配信設定 */
	Item         *CurrentDistributorMaster	`json:"item"`
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

type GetCurrentDistributorMasterResult struct {
    /** 現在有効な配信設定 */
	Item         *CurrentDistributorMaster	`json:"item"`
}

func (p *GetCurrentDistributorMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentDistributorMasterAsyncResult struct {
	result *GetCurrentDistributorMasterResult
	err    error
}

type UpdateCurrentDistributorMasterResult struct {
    /** 更新した現在有効な配信設定 */
	Item         *CurrentDistributorMaster	`json:"item"`
}

func (p *UpdateCurrentDistributorMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentDistributorMasterAsyncResult struct {
	result *UpdateCurrentDistributorMasterResult
	err    error
}

type UpdateCurrentDistributorMasterFromGitHubResult struct {
    /** 更新した現在有効な配信設定 */
	Item         *CurrentDistributorMaster	`json:"item"`
}

func (p *UpdateCurrentDistributorMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentDistributorMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentDistributorMasterFromGitHubResult
	err    error
}

type DistributeResult struct {
    /** 処理した DistributeResource */
	DistributeResource         *DistributeResource	`json:"distributeResource"`
    /** 所持品がキャパシティをオーバーしたときに転送するプレゼントボックスのネームスペース のGRN */
	InboxNamespaceId         *string	`json:"inboxNamespaceId"`
    /** レスポンス内容 */
	Result         *string	`json:"result"`
}

func (p *DistributeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.DistributeResource != nil {
        data["distributeResource"] = p.DistributeResource.ToDict()
    }
    if p.InboxNamespaceId != nil {
        data["inboxNamespaceId"] = p.InboxNamespaceId
    }
    if p.Result != nil {
        data["result"] = p.Result
    }
    return &data
}

type DistributeAsyncResult struct {
	result *DistributeResult
	err    error
}

type DistributeWithoutOverflowProcessResult struct {
    /** 処理した DistributeResource */
	DistributeResource         *DistributeResource	`json:"distributeResource"`
    /** レスポンス内容 */
	Result         *string	`json:"result"`
}

func (p *DistributeWithoutOverflowProcessResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.DistributeResource != nil {
        data["distributeResource"] = p.DistributeResource.ToDict()
    }
    if p.Result != nil {
        data["result"] = p.Result
    }
    return &data
}

type DistributeWithoutOverflowProcessAsyncResult struct {
	result *DistributeWithoutOverflowProcessResult
	err    error
}

type RunStampTaskResult struct {
    /** タスクの実行結果を反映したコンテキストスタック */
	ContextStack         *string	`json:"contextStack"`
    /** レスポンス内容 */
	Result         *string	`json:"result"`
}

func (p *RunStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.ContextStack != nil {
        data["contextStack"] = p.ContextStack
    }
    if p.Result != nil {
        data["result"] = p.Result
    }
    return &data
}

type RunStampTaskAsyncResult struct {
	result *RunStampTaskResult
	err    error
}

type RunStampSheetResult struct {
    /** レスポンス内容 */
	Result         *string	`json:"result"`
}

func (p *RunStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Result != nil {
        data["result"] = p.Result
    }
    return &data
}

type RunStampSheetAsyncResult struct {
	result *RunStampSheetResult
	err    error
}

type RunStampSheetExpressResult struct {
    /** スタンプタスクの実行結果 */
	TaskResults         []string	`json:"taskResults"`
    /** スタンプシートの実行結果レスポンス内容 */
	SheetResult         *string	`json:"sheetResult"`
}

func (p *RunStampSheetExpressResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.TaskResults != nil {
    	items := make([]string, 0)
    	for _, item := range p.TaskResults {
			items = append(items, item)
		}
		data["taskResults"] = items
    }
    if p.SheetResult != nil {
        data["sheetResult"] = p.SheetResult
    }
    return &data
}

type RunStampSheetExpressAsyncResult struct {
	result *RunStampSheetExpressResult
	err    error
}

type RunStampTaskWithoutNamespaceResult struct {
    /** タスクの実行結果を反映したコンテキストスタック */
	ContextStack         *string	`json:"contextStack"`
    /** レスポンス内容 */
	Result         *string	`json:"result"`
}

func (p *RunStampTaskWithoutNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.ContextStack != nil {
        data["contextStack"] = p.ContextStack
    }
    if p.Result != nil {
        data["result"] = p.Result
    }
    return &data
}

type RunStampTaskWithoutNamespaceAsyncResult struct {
	result *RunStampTaskWithoutNamespaceResult
	err    error
}

type RunStampSheetWithoutNamespaceResult struct {
    /** レスポンス内容 */
	Result         *string	`json:"result"`
}

func (p *RunStampSheetWithoutNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Result != nil {
        data["result"] = p.Result
    }
    return &data
}

type RunStampSheetWithoutNamespaceAsyncResult struct {
	result *RunStampSheetWithoutNamespaceResult
	err    error
}

type RunStampSheetExpressWithoutNamespaceResult struct {
    /** スタンプタスクの実行結果 */
	TaskResults         []string	`json:"taskResults"`
    /** スタンプシートの実行結果レスポンス内容 */
	SheetResult         *string	`json:"sheetResult"`
}

func (p *RunStampSheetExpressWithoutNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.TaskResults != nil {
    	items := make([]string, 0)
    	for _, item := range p.TaskResults {
			items = append(items, item)
		}
		data["taskResults"] = items
    }
    if p.SheetResult != nil {
        data["sheetResult"] = p.SheetResult
    }
    return &data
}

type RunStampSheetExpressWithoutNamespaceAsyncResult struct {
	result *RunStampSheetExpressWithoutNamespaceResult
	err    error
}
