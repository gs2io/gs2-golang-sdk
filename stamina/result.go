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

package stamina

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
	Item         *Namespace	`json:"item"`
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
	Status         *core.String	`json:"status"`
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
	Item         *Namespace	`json:"item"`
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
	Item         *Namespace	`json:"item"`
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
	Item         *Namespace	`json:"item"`
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

type DescribeStaminaModelMastersResult struct {
    /** スタミナモデルマスターのリスト */
	Items         *[]*StaminaModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeStaminaModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeStaminaModelMastersAsyncResult struct {
	result *DescribeStaminaModelMastersResult
	err    error
}

type CreateStaminaModelMasterResult struct {
    /** 作成したスタミナモデルマスター */
	Item         *StaminaModelMaster	`json:"item"`
}

func (p *CreateStaminaModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateStaminaModelMasterAsyncResult struct {
	result *CreateStaminaModelMasterResult
	err    error
}

type GetStaminaModelMasterResult struct {
    /** スタミナモデルマスター */
	Item         *StaminaModelMaster	`json:"item"`
}

func (p *GetStaminaModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetStaminaModelMasterAsyncResult struct {
	result *GetStaminaModelMasterResult
	err    error
}

type UpdateStaminaModelMasterResult struct {
    /** 更新したスタミナモデルマスター */
	Item         *StaminaModelMaster	`json:"item"`
}

func (p *UpdateStaminaModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateStaminaModelMasterAsyncResult struct {
	result *UpdateStaminaModelMasterResult
	err    error
}

type DeleteStaminaModelMasterResult struct {
    /** 削除したスタミナモデルマスター */
	Item         *StaminaModelMaster	`json:"item"`
}

func (p *DeleteStaminaModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteStaminaModelMasterAsyncResult struct {
	result *DeleteStaminaModelMasterResult
	err    error
}

type DescribeMaxStaminaTableMastersResult struct {
    /** スタミナの最大値テーブルマスターのリスト */
	Items         *[]*MaxStaminaTableMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeMaxStaminaTableMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeMaxStaminaTableMastersAsyncResult struct {
	result *DescribeMaxStaminaTableMastersResult
	err    error
}

type CreateMaxStaminaTableMasterResult struct {
    /** 作成したスタミナの最大値テーブルマスター */
	Item         *MaxStaminaTableMaster	`json:"item"`
}

func (p *CreateMaxStaminaTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateMaxStaminaTableMasterAsyncResult struct {
	result *CreateMaxStaminaTableMasterResult
	err    error
}

type GetMaxStaminaTableMasterResult struct {
    /** スタミナの最大値テーブルマスター */
	Item         *MaxStaminaTableMaster	`json:"item"`
}

func (p *GetMaxStaminaTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetMaxStaminaTableMasterAsyncResult struct {
	result *GetMaxStaminaTableMasterResult
	err    error
}

type UpdateMaxStaminaTableMasterResult struct {
    /** 更新したスタミナの最大値テーブルマスター */
	Item         *MaxStaminaTableMaster	`json:"item"`
}

func (p *UpdateMaxStaminaTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateMaxStaminaTableMasterAsyncResult struct {
	result *UpdateMaxStaminaTableMasterResult
	err    error
}

type DeleteMaxStaminaTableMasterResult struct {
    /** 削除したスタミナの最大値テーブルマスター */
	Item         *MaxStaminaTableMaster	`json:"item"`
}

func (p *DeleteMaxStaminaTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteMaxStaminaTableMasterAsyncResult struct {
	result *DeleteMaxStaminaTableMasterResult
	err    error
}

type DescribeRecoverIntervalTableMastersResult struct {
    /** スタミナ回復間隔テーブルマスターのリスト */
	Items         *[]*RecoverIntervalTableMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeRecoverIntervalTableMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeRecoverIntervalTableMastersAsyncResult struct {
	result *DescribeRecoverIntervalTableMastersResult
	err    error
}

type CreateRecoverIntervalTableMasterResult struct {
    /** 作成したスタミナ回復間隔テーブルマスター */
	Item         *RecoverIntervalTableMaster	`json:"item"`
}

func (p *CreateRecoverIntervalTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateRecoverIntervalTableMasterAsyncResult struct {
	result *CreateRecoverIntervalTableMasterResult
	err    error
}

type GetRecoverIntervalTableMasterResult struct {
    /** スタミナ回復間隔テーブルマスター */
	Item         *RecoverIntervalTableMaster	`json:"item"`
}

func (p *GetRecoverIntervalTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetRecoverIntervalTableMasterAsyncResult struct {
	result *GetRecoverIntervalTableMasterResult
	err    error
}

type UpdateRecoverIntervalTableMasterResult struct {
    /** 更新したスタミナ回復間隔テーブルマスター */
	Item         *RecoverIntervalTableMaster	`json:"item"`
}

func (p *UpdateRecoverIntervalTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateRecoverIntervalTableMasterAsyncResult struct {
	result *UpdateRecoverIntervalTableMasterResult
	err    error
}

type DeleteRecoverIntervalTableMasterResult struct {
    /** 削除したスタミナ回復間隔テーブルマスター */
	Item         *RecoverIntervalTableMaster	`json:"item"`
}

func (p *DeleteRecoverIntervalTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteRecoverIntervalTableMasterAsyncResult struct {
	result *DeleteRecoverIntervalTableMasterResult
	err    error
}

type DescribeRecoverValueTableMastersResult struct {
    /** スタミナ回復量テーブルマスターのリスト */
	Items         *[]*RecoverValueTableMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeRecoverValueTableMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeRecoverValueTableMastersAsyncResult struct {
	result *DescribeRecoverValueTableMastersResult
	err    error
}

type CreateRecoverValueTableMasterResult struct {
    /** 作成したスタミナ回復量テーブルマスター */
	Item         *RecoverValueTableMaster	`json:"item"`
}

func (p *CreateRecoverValueTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateRecoverValueTableMasterAsyncResult struct {
	result *CreateRecoverValueTableMasterResult
	err    error
}

type GetRecoverValueTableMasterResult struct {
    /** スタミナ回復量テーブルマスター */
	Item         *RecoverValueTableMaster	`json:"item"`
}

func (p *GetRecoverValueTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetRecoverValueTableMasterAsyncResult struct {
	result *GetRecoverValueTableMasterResult
	err    error
}

type UpdateRecoverValueTableMasterResult struct {
    /** 更新したスタミナ回復量テーブルマスター */
	Item         *RecoverValueTableMaster	`json:"item"`
}

func (p *UpdateRecoverValueTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateRecoverValueTableMasterAsyncResult struct {
	result *UpdateRecoverValueTableMasterResult
	err    error
}

type DeleteRecoverValueTableMasterResult struct {
    /** 削除したスタミナ回復量テーブルマスター */
	Item         *RecoverValueTableMaster	`json:"item"`
}

func (p *DeleteRecoverValueTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteRecoverValueTableMasterAsyncResult struct {
	result *DeleteRecoverValueTableMasterResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なスタミナマスター */
	Item         *CurrentStaminaMaster	`json:"item"`
}

func (p *ExportMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

type GetCurrentStaminaMasterResult struct {
    /** 現在有効なスタミナマスター */
	Item         *CurrentStaminaMaster	`json:"item"`
}

func (p *GetCurrentStaminaMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentStaminaMasterAsyncResult struct {
	result *GetCurrentStaminaMasterResult
	err    error
}

type UpdateCurrentStaminaMasterResult struct {
    /** 更新した現在有効なスタミナマスター */
	Item         *CurrentStaminaMaster	`json:"item"`
}

func (p *UpdateCurrentStaminaMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentStaminaMasterAsyncResult struct {
	result *UpdateCurrentStaminaMasterResult
	err    error
}

type UpdateCurrentStaminaMasterFromGitHubResult struct {
    /** 更新した現在有効なスタミナマスター */
	Item         *CurrentStaminaMaster	`json:"item"`
}

func (p *UpdateCurrentStaminaMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentStaminaMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentStaminaMasterFromGitHubResult
	err    error
}

type DescribeStaminaModelsResult struct {
    /** スタミナモデルのリスト */
	Items         *[]*StaminaModel	`json:"items"`
}

func (p *DescribeStaminaModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    return &data
}

type DescribeStaminaModelsAsyncResult struct {
	result *DescribeStaminaModelsResult
	err    error
}

type GetStaminaModelResult struct {
    /** スタミナモデル */
	Item         *StaminaModel	`json:"item"`
}

func (p *GetStaminaModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetStaminaModelAsyncResult struct {
	result *GetStaminaModelResult
	err    error
}

type DescribeStaminasResult struct {
    /** スタミナのリスト */
	Items         *[]*Stamina	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeStaminasResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeStaminasAsyncResult struct {
	result *DescribeStaminasResult
	err    error
}

type DescribeStaminasByUserIdResult struct {
    /** スタミナのリスト */
	Items         *[]*Stamina	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeStaminasByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeStaminasByUserIdAsyncResult struct {
	result *DescribeStaminasByUserIdResult
	err    error
}

type GetStaminaResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *GetStaminaResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type GetStaminaAsyncResult struct {
	result *GetStaminaResult
	err    error
}

type GetStaminaByUserIdResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *GetStaminaByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type GetStaminaByUserIdAsyncResult struct {
	result *GetStaminaByUserIdResult
	err    error
}

type UpdateStaminaByUserIdResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *UpdateStaminaByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type UpdateStaminaByUserIdAsyncResult struct {
	result *UpdateStaminaByUserIdResult
	err    error
}

type ConsumeStaminaResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *ConsumeStaminaResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type ConsumeStaminaAsyncResult struct {
	result *ConsumeStaminaResult
	err    error
}

type ConsumeStaminaByUserIdResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *ConsumeStaminaByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type ConsumeStaminaByUserIdAsyncResult struct {
	result *ConsumeStaminaByUserIdResult
	err    error
}

type RecoverStaminaByUserIdResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
    /** スタミナ値の上限を超えて受け取れずに GS2-Inbox に転送したスタミナ値 */
	OverflowValue         *int64	`json:"overflowValue"`
}

func (p *RecoverStaminaByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    if p.OverflowValue != nil {
     data["OverflowValue"] = p.OverflowValue
    }
    return &data
}

type RecoverStaminaByUserIdAsyncResult struct {
	result *RecoverStaminaByUserIdResult
	err    error
}

type RaiseMaxValueByUserIdResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *RaiseMaxValueByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type RaiseMaxValueByUserIdAsyncResult struct {
	result *RaiseMaxValueByUserIdResult
	err    error
}

type SetMaxValueByUserIdResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *SetMaxValueByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type SetMaxValueByUserIdAsyncResult struct {
	result *SetMaxValueByUserIdResult
	err    error
}

type SetRecoverIntervalByUserIdResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *SetRecoverIntervalByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type SetRecoverIntervalByUserIdAsyncResult struct {
	result *SetRecoverIntervalByUserIdResult
	err    error
}

type SetRecoverValueByUserIdResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *SetRecoverValueByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type SetRecoverValueByUserIdAsyncResult struct {
	result *SetRecoverValueByUserIdResult
	err    error
}

type SetMaxValueByStatusResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *SetMaxValueByStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type SetMaxValueByStatusAsyncResult struct {
	result *SetMaxValueByStatusResult
	err    error
}

type SetRecoverIntervalByStatusResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *SetRecoverIntervalByStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type SetRecoverIntervalByStatusAsyncResult struct {
	result *SetRecoverIntervalByStatusResult
	err    error
}

type SetRecoverValueByStatusResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *SetRecoverValueByStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type SetRecoverValueByStatusAsyncResult struct {
	result *SetRecoverValueByStatusResult
	err    error
}

type DeleteStaminaByUserIdResult struct {
}

func (p *DeleteStaminaByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteStaminaByUserIdAsyncResult struct {
	result *DeleteStaminaByUserIdResult
	err    error
}

type RecoverStaminaByStampSheetResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
    /** スタミナ値の上限を超えて受け取れずに GS2-Inbox に転送したスタミナ値 */
	OverflowValue         *int64	`json:"overflowValue"`
}

func (p *RecoverStaminaByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    if p.OverflowValue != nil {
     data["OverflowValue"] = p.OverflowValue
    }
    return &data
}

type RecoverStaminaByStampSheetAsyncResult struct {
	result *RecoverStaminaByStampSheetResult
	err    error
}

type RaiseMaxValueByStampSheetResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *RaiseMaxValueByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type RaiseMaxValueByStampSheetAsyncResult struct {
	result *RaiseMaxValueByStampSheetResult
	err    error
}

type SetMaxValueByStampSheetResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *SetMaxValueByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type SetMaxValueByStampSheetAsyncResult struct {
	result *SetMaxValueByStampSheetResult
	err    error
}

type SetRecoverIntervalByStampSheetResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *SetRecoverIntervalByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type SetRecoverIntervalByStampSheetAsyncResult struct {
	result *SetRecoverIntervalByStampSheetResult
	err    error
}

type SetRecoverValueByStampSheetResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
}

func (p *SetRecoverValueByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    return &data
}

type SetRecoverValueByStampSheetAsyncResult struct {
	result *SetRecoverValueByStampSheetResult
	err    error
}

type ConsumeStaminaByStampTaskResult struct {
    /** スタミナ */
	Item         *Stamina	`json:"item"`
    /** スタミナモデル */
	StaminaModel         *StaminaModel	`json:"staminaModel"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *core.String	`json:"newContextStack"`
}

func (p *ConsumeStaminaByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.StaminaModel != nil {
     data["StaminaModel"] = p.StaminaModel.ToDict()
    }
    if p.NewContextStack != nil {
     data["NewContextStack"] = p.NewContextStack
    }
    return &data
}

type ConsumeStaminaByStampTaskAsyncResult struct {
	result *ConsumeStaminaByStampTaskResult
	err    error
}
