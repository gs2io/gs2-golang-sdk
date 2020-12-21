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

package mission

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeMissionTaskModelMastersResult struct {
    /** ミッションタスクマスターのリスト */
	Items         []MissionTaskModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeMissionTaskModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]MissionTaskModelMaster, 0)
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

type DescribeMissionTaskModelMastersAsyncResult struct {
	result *DescribeMissionTaskModelMastersResult
	err    error
}

type CreateMissionTaskModelMasterResult struct {
    /** 作成したミッションタスクマスター */
	Item         *MissionTaskModelMaster	`json:"item"`
}

func (p *CreateMissionTaskModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateMissionTaskModelMasterAsyncResult struct {
	result *CreateMissionTaskModelMasterResult
	err    error
}

type GetMissionTaskModelMasterResult struct {
    /** ミッションタスクマスター */
	Item         *MissionTaskModelMaster	`json:"item"`
}

func (p *GetMissionTaskModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetMissionTaskModelMasterAsyncResult struct {
	result *GetMissionTaskModelMasterResult
	err    error
}

type UpdateMissionTaskModelMasterResult struct {
    /** 更新したミッションタスクマスター */
	Item         *MissionTaskModelMaster	`json:"item"`
}

func (p *UpdateMissionTaskModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateMissionTaskModelMasterAsyncResult struct {
	result *UpdateMissionTaskModelMasterResult
	err    error
}

type DeleteMissionTaskModelMasterResult struct {
    /** 削除したミッションタスク */
	Item         *MissionTaskModelMaster	`json:"item"`
}

func (p *DeleteMissionTaskModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteMissionTaskModelMasterAsyncResult struct {
	result *DeleteMissionTaskModelMasterResult
	err    error
}

type DescribeCounterModelsResult struct {
    /** カウンターの種類のリスト */
	Items         []CounterModel	`json:"items"`
}

func (p *DescribeCounterModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]CounterModel, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeCounterModelsAsyncResult struct {
	result *DescribeCounterModelsResult
	err    error
}

type GetCounterModelResult struct {
    /** カウンターの種類 */
	Item         *CounterModel	`json:"item"`
}

func (p *GetCounterModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCounterModelAsyncResult struct {
	result *GetCounterModelResult
	err    error
}

type DescribeCountersResult struct {
    /** カウンターのリスト */
	Items         []Counter	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCountersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Counter, 0)
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

type DescribeCountersAsyncResult struct {
	result *DescribeCountersResult
	err    error
}

type DescribeCountersByUserIdResult struct {
    /** カウンターのリスト */
	Items         []Counter	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCountersByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Counter, 0)
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

type DescribeCountersByUserIdAsyncResult struct {
	result *DescribeCountersByUserIdResult
	err    error
}

type IncreaseCounterByUserIdResult struct {
    /** 作成したカウンター */
	Item         *Counter	`json:"item"`
}

func (p *IncreaseCounterByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type IncreaseCounterByUserIdAsyncResult struct {
	result *IncreaseCounterByUserIdResult
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

type DeleteCounterByUserIdResult struct {
    /** 削除したカウンター */
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

type IncreaseByStampSheetResult struct {
    /** カウンター加算後のカウンター */
	Item         *Counter	`json:"item"`
}

func (p *IncreaseByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type IncreaseByStampSheetAsyncResult struct {
	result *IncreaseByStampSheetResult
	err    error
}

type DescribeCounterModelMastersResult struct {
    /** カウンターの種類マスターのリスト */
	Items         []CounterModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCounterModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]CounterModelMaster, 0)
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

type DescribeCounterModelMastersAsyncResult struct {
	result *DescribeCounterModelMastersResult
	err    error
}

type CreateCounterModelMasterResult struct {
    /** 作成したカウンターの種類マスター */
	Item         *CounterModelMaster	`json:"item"`
}

func (p *CreateCounterModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateCounterModelMasterAsyncResult struct {
	result *CreateCounterModelMasterResult
	err    error
}

type GetCounterModelMasterResult struct {
    /** カウンターの種類マスター */
	Item         *CounterModelMaster	`json:"item"`
}

func (p *GetCounterModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCounterModelMasterAsyncResult struct {
	result *GetCounterModelMasterResult
	err    error
}

type UpdateCounterModelMasterResult struct {
    /** 更新したカウンターの種類マスター */
	Item         *CounterModelMaster	`json:"item"`
}

func (p *UpdateCounterModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCounterModelMasterAsyncResult struct {
	result *UpdateCounterModelMasterResult
	err    error
}

type DeleteCounterModelMasterResult struct {
    /** 削除したカウンターの種類マスター */
	Item         *CounterModelMaster	`json:"item"`
}

func (p *DeleteCounterModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteCounterModelMasterAsyncResult struct {
	result *DeleteCounterModelMasterResult
	err    error
}

type DescribeMissionGroupModelsResult struct {
    /** ミッショングループのリスト */
	Items         []MissionGroupModel	`json:"items"`
}

func (p *DescribeMissionGroupModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]MissionGroupModel, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeMissionGroupModelsAsyncResult struct {
	result *DescribeMissionGroupModelsResult
	err    error
}

type GetMissionGroupModelResult struct {
    /** ミッショングループ */
	Item         *MissionGroupModel	`json:"item"`
}

func (p *GetMissionGroupModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetMissionGroupModelAsyncResult struct {
	result *GetMissionGroupModelResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なミッション */
	Item         *CurrentMissionMaster	`json:"item"`
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

type GetCurrentMissionMasterResult struct {
    /** 現在有効なミッション */
	Item         *CurrentMissionMaster	`json:"item"`
}

func (p *GetCurrentMissionMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentMissionMasterAsyncResult struct {
	result *GetCurrentMissionMasterResult
	err    error
}

type UpdateCurrentMissionMasterResult struct {
    /** 更新した現在有効なミッション */
	Item         *CurrentMissionMaster	`json:"item"`
}

func (p *UpdateCurrentMissionMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentMissionMasterAsyncResult struct {
	result *UpdateCurrentMissionMasterResult
	err    error
}

type UpdateCurrentMissionMasterFromGitHubResult struct {
    /** 更新した現在有効なミッション */
	Item         *CurrentMissionMaster	`json:"item"`
}

func (p *UpdateCurrentMissionMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentMissionMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentMissionMasterFromGitHubResult
	err    error
}

type DescribeMissionGroupModelMastersResult struct {
    /** ミッショングループマスターのリスト */
	Items         []MissionGroupModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeMissionGroupModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]MissionGroupModelMaster, 0)
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

type DescribeMissionGroupModelMastersAsyncResult struct {
	result *DescribeMissionGroupModelMastersResult
	err    error
}

type CreateMissionGroupModelMasterResult struct {
    /** 作成したミッショングループマスター */
	Item         *MissionGroupModelMaster	`json:"item"`
}

func (p *CreateMissionGroupModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateMissionGroupModelMasterAsyncResult struct {
	result *CreateMissionGroupModelMasterResult
	err    error
}

type GetMissionGroupModelMasterResult struct {
    /** ミッショングループマスター */
	Item         *MissionGroupModelMaster	`json:"item"`
}

func (p *GetMissionGroupModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetMissionGroupModelMasterAsyncResult struct {
	result *GetMissionGroupModelMasterResult
	err    error
}

type UpdateMissionGroupModelMasterResult struct {
    /** 更新したミッショングループマスター */
	Item         *MissionGroupModelMaster	`json:"item"`
}

func (p *UpdateMissionGroupModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateMissionGroupModelMasterAsyncResult struct {
	result *UpdateMissionGroupModelMasterResult
	err    error
}

type DeleteMissionGroupModelMasterResult struct {
    /** 削除したミッショングループ */
	Item         *MissionGroupModelMaster	`json:"item"`
}

func (p *DeleteMissionGroupModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteMissionGroupModelMasterAsyncResult struct {
	result *DeleteMissionGroupModelMasterResult
	err    error
}

type DescribeCompletesResult struct {
    /** 達成状況のリスト */
	Items         []Complete	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCompletesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Complete, 0)
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

type DescribeCompletesAsyncResult struct {
	result *DescribeCompletesResult
	err    error
}

type DescribeCompletesByUserIdResult struct {
    /** 達成状況のリスト */
	Items         []Complete	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCompletesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Complete, 0)
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

type DescribeCompletesByUserIdAsyncResult struct {
	result *DescribeCompletesByUserIdResult
	err    error
}

type CompleteResult struct {
    /** ミッションの達成報酬を受領するスタンプシート */
	StampSheet         *core.String	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *core.String	`json:"stampSheetEncryptionKeyId"`
}

func (p *CompleteResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type CompleteAsyncResult struct {
	result *CompleteResult
	err    error
}

type CompleteByUserIdResult struct {
    /** ミッションの達成報酬を受領するスタンプシート */
	StampSheet         *core.String	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *core.String	`json:"stampSheetEncryptionKeyId"`
}

func (p *CompleteByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type CompleteByUserIdAsyncResult struct {
	result *CompleteByUserIdResult
	err    error
}

type ReceiveByUserIdResult struct {
    /** 受領した達成状況 */
	Item         *Complete	`json:"item"`
}

func (p *ReceiveByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type ReceiveByUserIdAsyncResult struct {
	result *ReceiveByUserIdResult
	err    error
}

type GetCompleteResult struct {
    /** 達成状況 */
	Item         *Complete	`json:"item"`
}

func (p *GetCompleteResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCompleteAsyncResult struct {
	result *GetCompleteResult
	err    error
}

type GetCompleteByUserIdResult struct {
    /** 達成状況 */
	Item         *Complete	`json:"item"`
}

func (p *GetCompleteByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCompleteByUserIdAsyncResult struct {
	result *GetCompleteByUserIdResult
	err    error
}

type DeleteCompleteByUserIdResult struct {
    /** 削除した達成状況 */
	Item         *Complete	`json:"item"`
}

func (p *DeleteCompleteByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteCompleteByUserIdAsyncResult struct {
	result *DeleteCompleteByUserIdResult
	err    error
}

type ReceiveByStampTaskResult struct {
    /** 達成状況 */
	Item         *Complete	`json:"item"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *core.String	`json:"newContextStack"`
}

func (p *ReceiveByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.NewContextStack != nil {
        data["newContextStack"] = p.NewContextStack
    }
    return &data
}

type ReceiveByStampTaskAsyncResult struct {
	result *ReceiveByStampTaskResult
	err    error
}

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         []Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
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

type DescribeMissionTaskModelsResult struct {
    /** ミッションタスクのリスト */
	Items         []MissionTaskModel	`json:"items"`
}

func (p *DescribeMissionTaskModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]MissionTaskModel, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeMissionTaskModelsAsyncResult struct {
	result *DescribeMissionTaskModelsResult
	err    error
}

type GetMissionTaskModelResult struct {
    /** ミッションタスク */
	Item         *MissionTaskModel	`json:"item"`
}

func (p *GetMissionTaskModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetMissionTaskModelAsyncResult struct {
	result *GetMissionTaskModelResult
	err    error
}
