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

type Namespace struct {
    /** ネームスペース */
	NamespaceId *core.String   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ネームスペース名 */
	Name *core.String   `json:"name"`
    /** 説明文 */
	Description *core.String   `json:"description"`
    /** スタミナオーバーフロー上限に当たって回復できなかったスタミナを通知する スクリプト のGRN */
	OverflowTriggerScriptId *core.String   `json:"overflowTriggerScriptId"`
    /** スタミナオーバーフロー上限に当たって回復できなかったスタミナを追加する ネームスペース のGRN */
	OverflowTriggerNamespaceId *core.String   `json:"overflowTriggerNamespaceId"`
    /** ログの出力設定 */
	LogSetting *LogSetting   `json:"logSetting"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Namespace) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["ownerId"] = p.OwnerId
    data["name"] = p.Name
    data["description"] = p.Description
    data["overflowTriggerScriptId"] = p.OverflowTriggerScriptId
    data["overflowTriggerNamespaceId"] = p.OverflowTriggerNamespaceId
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type StaminaModelMaster struct {
    /** スタミナモデルマスター */
	StaminaModelId *core.String   `json:"staminaModelId"`
    /** スタミナの種類名 */
	Name *core.String   `json:"name"`
    /** スタミナの種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** スタミナモデルマスターの説明 */
	Description *core.String   `json:"description"`
    /** スタミナを回復する速度(分) */
	RecoverIntervalMinutes *int32   `json:"recoverIntervalMinutes"`
    /** 時間経過後に回復する量 */
	RecoverValue *int32   `json:"recoverValue"`
    /** スタミナの最大値の初期値 */
	InitialCapacity *int32   `json:"initialCapacity"`
    /** 最大値を超えて回復するか */
	IsOverflow *bool   `json:"isOverflow"`
    /** 溢れた状況での最大値 */
	MaxCapacity *int32   `json:"maxCapacity"`
    /** GS2-Experience のランクによって最大スタミナ値を決定するスタミナ最大値テーブル名 */
	MaxStaminaTableName *core.String   `json:"maxStaminaTableName"`
    /** GS2-Experience のランクによってスタミナの回復間隔を決定する回復間隔テーブル名 */
	RecoverIntervalTableName *core.String   `json:"recoverIntervalTableName"`
    /** GS2-Experience のランクによってスタミナの回復量を決定する回復量テーブル名 */
	RecoverValueTableName *core.String   `json:"recoverValueTableName"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *StaminaModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["staminaModelId"] = p.StaminaModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    data["recoverIntervalMinutes"] = p.RecoverIntervalMinutes
    data["recoverValue"] = p.RecoverValue
    data["initialCapacity"] = p.InitialCapacity
    data["isOverflow"] = p.IsOverflow
    data["maxCapacity"] = p.MaxCapacity
    data["maxStaminaTableName"] = p.MaxStaminaTableName
    data["recoverIntervalTableName"] = p.RecoverIntervalTableName
    data["recoverValueTableName"] = p.RecoverValueTableName
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type MaxStaminaTableMaster struct {
    /** スタミナの最大値テーブルマスター */
	MaxStaminaTableId *core.String   `json:"maxStaminaTableId"`
    /** 最大スタミナ値テーブル名 */
	Name *core.String   `json:"name"`
    /** 最大スタミナ値テーブルのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** スタミナの最大値テーブルマスターの説明 */
	Description *core.String   `json:"description"`
    /** 経験値の種類マスター のGRN */
	ExperienceModelId *core.String   `json:"experienceModelId"`
    /** ランク毎のスタミナの最大値テーブル */
	Values *[]int32   `json:"values"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *MaxStaminaTableMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["maxStaminaTableId"] = p.MaxStaminaTableId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    data["experienceModelId"] = p.ExperienceModelId
    if p.Values != nil {
        var _values []int32
        for _, item := range *p.Values {
            _values = append(_values, item)
        }
        data["values"] = &_values
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type RecoverIntervalTableMaster struct {
    /** スタミナ回復間隔テーブルマスター */
	RecoverIntervalTableId *core.String   `json:"recoverIntervalTableId"`
    /** スタミナ回復間隔テーブル名 */
	Name *core.String   `json:"name"`
    /** スタミナ回復間隔テーブルのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** スタミナ回復間隔テーブルマスターの説明 */
	Description *core.String   `json:"description"`
    /** 経験値の種類マスター のGRN */
	ExperienceModelId *core.String   `json:"experienceModelId"`
    /** ランク毎のスタミナ回復間隔テーブル */
	Values *[]int32   `json:"values"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *RecoverIntervalTableMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["recoverIntervalTableId"] = p.RecoverIntervalTableId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    data["experienceModelId"] = p.ExperienceModelId
    if p.Values != nil {
        var _values []int32
        for _, item := range *p.Values {
            _values = append(_values, item)
        }
        data["values"] = &_values
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type RecoverValueTableMaster struct {
    /** スタミナ回復量テーブルマスター */
	RecoverValueTableId *core.String   `json:"recoverValueTableId"`
    /** スタミナ回復量テーブル名 */
	Name *core.String   `json:"name"`
    /** スタミナ回復量テーブルのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** スタミナ回復量テーブルマスターの説明 */
	Description *core.String   `json:"description"`
    /** 経験値の種類マスター のGRN */
	ExperienceModelId *core.String   `json:"experienceModelId"`
    /** ランク毎のスタミナ回復量テーブル */
	Values *[]int32   `json:"values"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *RecoverValueTableMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["recoverValueTableId"] = p.RecoverValueTableId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    data["experienceModelId"] = p.ExperienceModelId
    if p.Values != nil {
        var _values []int32
        for _, item := range *p.Values {
            _values = append(_values, item)
        }
        data["values"] = &_values
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CurrentStaminaMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentStaminaMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
    data["settings"] = p.Settings
    return &data
}

type StaminaModel struct {
    /** スタミナモデルマスター */
	StaminaModelId *core.String   `json:"staminaModelId"`
    /** スタミナの種類名 */
	Name *core.String   `json:"name"`
    /** スタミナの種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** スタミナを回復する速度(分) */
	RecoverIntervalMinutes *int32   `json:"recoverIntervalMinutes"`
    /** 時間経過後に回復する量 */
	RecoverValue *int32   `json:"recoverValue"`
    /** スタミナの最大値の初期値 */
	InitialCapacity *int32   `json:"initialCapacity"`
    /** 最大値を超えて回復するか */
	IsOverflow *bool   `json:"isOverflow"`
    /** 溢れた状況での最大値 */
	MaxCapacity *int32   `json:"maxCapacity"`
    /** GS2-Experience と連携する際に使用するスタミナ最大値テーブル */
	MaxStaminaTable *MaxStaminaTable   `json:"maxStaminaTable"`
    /** GS2-Experience と連携する際に使用する回復間隔テーブル */
	RecoverIntervalTable *RecoverIntervalTable   `json:"recoverIntervalTable"`
    /** GS2-Experience と連携する際に使用する回復量テーブル */
	RecoverValueTable *RecoverValueTable   `json:"recoverValueTable"`
}

func (p *StaminaModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["staminaModelId"] = p.StaminaModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["recoverIntervalMinutes"] = p.RecoverIntervalMinutes
    data["recoverValue"] = p.RecoverValue
    data["initialCapacity"] = p.InitialCapacity
    data["isOverflow"] = p.IsOverflow
    data["maxCapacity"] = p.MaxCapacity
    if p.MaxStaminaTable != nil {
        data["maxStaminaTable"] = *p.MaxStaminaTable.ToDict()
    }
    if p.RecoverIntervalTable != nil {
        data["recoverIntervalTable"] = *p.RecoverIntervalTable.ToDict()
    }
    if p.RecoverValueTable != nil {
        data["recoverValueTable"] = *p.RecoverValueTable.ToDict()
    }
    return &data
}

type MaxStaminaTable struct {
    /** スタミナの最大値テーブルマスター */
	MaxStaminaTableId *core.String   `json:"maxStaminaTableId"`
    /** 最大スタミナ値テーブル名 */
	Name *core.String   `json:"name"`
    /** 最大スタミナ値テーブルのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 経験値の種類マスター のGRN */
	ExperienceModelId *core.String   `json:"experienceModelId"`
    /** ランク毎のスタミナの最大値テーブル */
	Values *[]int32   `json:"values"`
}

func (p *MaxStaminaTable) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["maxStaminaTableId"] = p.MaxStaminaTableId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["experienceModelId"] = p.ExperienceModelId
    if p.Values != nil {
        var _values []int32
        for _, item := range *p.Values {
            _values = append(_values, item)
        }
        data["values"] = &_values
    }
    return &data
}

type RecoverIntervalTable struct {
    /** スタミナ回復間隔テーブルマスター */
	RecoverIntervalTableId *core.String   `json:"recoverIntervalTableId"`
    /** スタミナ回復間隔テーブル名 */
	Name *core.String   `json:"name"`
    /** スタミナ回復間隔テーブルのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 経験値の種類マスター のGRN */
	ExperienceModelId *core.String   `json:"experienceModelId"`
    /** ランク毎のスタミナ回復間隔テーブル */
	Values *[]int32   `json:"values"`
}

func (p *RecoverIntervalTable) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["recoverIntervalTableId"] = p.RecoverIntervalTableId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["experienceModelId"] = p.ExperienceModelId
    if p.Values != nil {
        var _values []int32
        for _, item := range *p.Values {
            _values = append(_values, item)
        }
        data["values"] = &_values
    }
    return &data
}

type RecoverValueTable struct {
    /** スタミナ回復量テーブルマスター */
	RecoverValueTableId *core.String   `json:"recoverValueTableId"`
    /** スタミナ回復量テーブル名 */
	Name *core.String   `json:"name"`
    /** スタミナ回復量テーブルのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 経験値の種類マスター のGRN */
	ExperienceModelId *core.String   `json:"experienceModelId"`
    /** ランク毎のスタミナ回復量テーブル */
	Values *[]int32   `json:"values"`
}

func (p *RecoverValueTable) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["recoverValueTableId"] = p.RecoverValueTableId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["experienceModelId"] = p.ExperienceModelId
    if p.Values != nil {
        var _values []int32
        for _, item := range *p.Values {
            _values = append(_values, item)
        }
        data["values"] = &_values
    }
    return &data
}

type Stamina struct {
    /** スタミナ */
	StaminaId *core.String   `json:"staminaId"`
    /** スタミナモデルの名前 */
	StaminaName *core.String   `json:"staminaName"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 最終更新時におけるスタミナ値 */
	Value *int32   `json:"value"`
    /** スタミナの最大値 */
	MaxValue *int32   `json:"maxValue"`
    /** スタミナの回復間隔(分) */
	RecoverIntervalMinutes *int32   `json:"recoverIntervalMinutes"`
    /** スタミナの回復量 */
	RecoverValue *int32   `json:"recoverValue"`
    /** スタミナの最大値を超えて格納されているスタミナ値 */
	OverflowValue *int32   `json:"overflowValue"`
    /** 次回スタミナが回復する時間 */
	NextRecoverAt *int64   `json:"nextRecoverAt"`
    /** 作成日時 */
	LastRecoveredAt *int64   `json:"lastRecoveredAt"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Stamina) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["staminaId"] = p.StaminaId
    data["staminaName"] = p.StaminaName
    data["userId"] = p.UserId
    data["value"] = p.Value
    data["maxValue"] = p.MaxValue
    data["recoverIntervalMinutes"] = p.RecoverIntervalMinutes
    data["recoverValue"] = p.RecoverValue
    data["overflowValue"] = p.OverflowValue
    data["nextRecoverAt"] = p.NextRecoverAt
    data["lastRecoveredAt"] = p.LastRecoveredAt
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type ResponseCache struct {
    /** None */
	Region *core.String   `json:"region"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** レスポンスキャッシュ のGRN */
	ResponseCacheId *core.String   `json:"responseCacheId"`
    /** None */
	RequestHash *core.String   `json:"requestHash"`
    /** APIの応答内容 */
	Result *core.String   `json:"result"`
}

func (p *ResponseCache) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["region"] = p.Region
    data["ownerId"] = p.OwnerId
    data["responseCacheId"] = p.ResponseCacheId
    data["requestHash"] = p.RequestHash
    data["result"] = p.Result
    return &data
}

type GitHubCheckoutSetting struct {
    /** リソースの取得に使用するGitHub のAPIキー のGRN */
	GitHubApiKeyId *core.String   `json:"gitHubApiKeyId"`
    /** リポジトリ名 */
	RepositoryName *core.String   `json:"repositoryName"`
    /** ソースコードのファイルパス */
	SourcePath *core.String   `json:"sourcePath"`
    /** コードの取得元 */
	ReferenceType *core.String   `json:"referenceType"`
    /** コミットハッシュ */
	CommitHash *core.String   `json:"commitHash"`
    /** ブランチ名 */
	BranchName *core.String   `json:"branchName"`
    /** タグ名 */
	TagName *core.String   `json:"tagName"`
}

func (p *GitHubCheckoutSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["gitHubApiKeyId"] = p.GitHubApiKeyId
    data["repositoryName"] = p.RepositoryName
    data["sourcePath"] = p.SourcePath
    data["referenceType"] = p.ReferenceType
    data["commitHash"] = p.CommitHash
    data["branchName"] = p.BranchName
    data["tagName"] = p.TagName
    return &data
}

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *core.String   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}
