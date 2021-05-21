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

package experience

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** ランクキャップ取得時 に実行されるスクリプト のGRN */
	ExperienceCapScriptId *string   `json:"experienceCapScriptId"`
    /** 経験値変化したときに実行するスクリプト */
	ChangeExperienceScript *ScriptSetting   `json:"changeExperienceScript"`
    /** ランク変化したときに実行するスクリプト */
	ChangeRankScript *ScriptSetting   `json:"changeRankScript"`
    /** ランクキャップ変化したときに実行するスクリプト */
	ChangeRankCapScript *ScriptSetting   `json:"changeRankCapScript"`
    /** 経験値あふれしたときに実行するスクリプト */
	OverflowExperienceScript *ScriptSetting   `json:"overflowExperienceScript"`
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
    data["experienceCapScriptId"] = p.ExperienceCapScriptId
    if p.ChangeExperienceScript != nil {
        data["changeExperienceScript"] = *p.ChangeExperienceScript.ToDict()
    }
    if p.ChangeRankScript != nil {
        data["changeRankScript"] = *p.ChangeRankScript.ToDict()
    }
    if p.ChangeRankCapScript != nil {
        data["changeRankCapScript"] = *p.ChangeRankCapScript.ToDict()
    }
    if p.OverflowExperienceScript != nil {
        data["overflowExperienceScript"] = *p.OverflowExperienceScript.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type ExperienceModelMaster struct {
    /** 経験値の種類マスター */
	ExperienceModelId *string   `json:"experienceModelId"`
    /** 経験値の種類名 */
	Name *string   `json:"name"`
    /** 経験値の種類マスターの説明 */
	Description *string   `json:"description"`
    /** 経験値の種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** 経験値の初期値 */
	DefaultExperience *int64   `json:"defaultExperience"`
    /** ランクキャップの初期値 */
	DefaultRankCap *int64   `json:"defaultRankCap"`
    /** ランクキャップの最大値 */
	MaxRankCap *int64   `json:"maxRankCap"`
    /** ランク計算に用いる */
	RankThresholdId *string   `json:"rankThresholdId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *ExperienceModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["experienceModelId"] = p.ExperienceModelId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    data["defaultExperience"] = p.DefaultExperience
    data["defaultRankCap"] = p.DefaultRankCap
    data["maxRankCap"] = p.MaxRankCap
    data["rankThresholdId"] = p.RankThresholdId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type ExperienceModel struct {
    /** 経験値の種類マスター */
	ExperienceModelId *string   `json:"experienceModelId"`
    /** 経験値の種類名 */
	Name *string   `json:"name"`
    /** 経験値の種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** 経験値の初期値 */
	DefaultExperience *int64   `json:"defaultExperience"`
    /** ランクキャップの初期値 */
	DefaultRankCap *int64   `json:"defaultRankCap"`
    /** ランクキャップの最大値 */
	MaxRankCap *int64   `json:"maxRankCap"`
    /** ランクアップ閾値 */
	RankThreshold *Threshold   `json:"rankThreshold"`
}

func (p *ExperienceModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["experienceModelId"] = p.ExperienceModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["defaultExperience"] = p.DefaultExperience
    data["defaultRankCap"] = p.DefaultRankCap
    data["maxRankCap"] = p.MaxRankCap
    if p.RankThreshold != nil {
        data["rankThreshold"] = *p.RankThreshold.ToDict()
    }
    return &data
}

type ThresholdMaster struct {
    /** ランクアップ閾値マスター */
	ThresholdId *string   `json:"thresholdId"`
    /** ランクアップ閾値名 */
	Name *string   `json:"name"`
    /** ランクアップ閾値マスターの説明 */
	Description *string   `json:"description"`
    /** ランクアップ閾値のメタデータ */
	Metadata *string   `json:"metadata"`
    /** ランクアップ経験値閾値リスト */
	Values []int64   `json:"values"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *ThresholdMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["thresholdId"] = p.ThresholdId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    if p.Values != nil {
        var _values []int64
        for _, item := range p.Values {
            _values = append(_values, item)
        }
        data["values"] = &_values
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Threshold struct {
    /** ランクアップ閾値のメタデータ */
	Metadata *string   `json:"metadata"`
    /** ランクアップ経験値閾値リスト */
	Values []int64   `json:"values"`
}

func (p *Threshold) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["metadata"] = p.Metadata
    if p.Values != nil {
        var _values []int64
        for _, item := range p.Values {
            _values = append(_values, item)
        }
        data["values"] = &_values
    }
    return &data
}

type CurrentExperienceMaster struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentExperienceMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["settings"] = p.Settings
    return &data
}

type Status struct {
    /** ステータス */
	StatusId *string   `json:"statusId"`
    /** 経験値の種類の名前 */
	ExperienceName *string   `json:"experienceName"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** プロパティID */
	PropertyId *string   `json:"propertyId"`
    /** 累計獲得経験値 */
	ExperienceValue *int64   `json:"experienceValue"`
    /** 現在のランク */
	RankValue *int64   `json:"rankValue"`
    /** 現在のランクキャップ */
	RankCapValue *int64   `json:"rankCapValue"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Status) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["statusId"] = p.StatusId
    data["experienceName"] = p.ExperienceName
    data["userId"] = p.UserId
    data["propertyId"] = p.PropertyId
    data["experienceValue"] = p.ExperienceValue
    data["rankValue"] = p.RankValue
    data["rankCapValue"] = p.RankCapValue
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type ResponseCache struct {
    /** None */
	Region *string   `json:"region"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** レスポンスキャッシュ のGRN */
	ResponseCacheId *string   `json:"responseCacheId"`
    /** None */
	RequestHash *string   `json:"requestHash"`
    /** APIの応答内容 */
	Result *string   `json:"result"`
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
	ApiKeyId *string   `json:"apiKeyId"`
    /** リポジトリ名 */
	RepositoryName *string   `json:"repositoryName"`
    /** ソースコードのファイルパス */
	SourcePath *string   `json:"sourcePath"`
    /** コードの取得元 */
	ReferenceType *string   `json:"referenceType"`
    /** コミットハッシュ */
	CommitHash *string   `json:"commitHash"`
    /** ブランチ名 */
	BranchName *string   `json:"branchName"`
    /** タグ名 */
	TagName *string   `json:"tagName"`
}

func (p *GitHubCheckoutSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["apiKeyId"] = p.ApiKeyId
    data["repositoryName"] = p.RepositoryName
    data["sourcePath"] = p.SourcePath
    data["referenceType"] = p.ReferenceType
    data["commitHash"] = p.CommitHash
    data["branchName"] = p.BranchName
    data["tagName"] = p.TagName
    return &data
}

type ScriptSetting struct {
    /** 実行前に使用する GS2-Script のスクリプト のGRN */
	TriggerScriptId *string   `json:"triggerScriptId"`
    /** 完了通知の通知先 */
	DoneTriggerTargetType *string   `json:"doneTriggerTargetType"`
    /** 完了時に使用する GS2-Script のスクリプト のGRN */
	DoneTriggerScriptId *string   `json:"doneTriggerScriptId"`
    /** 完了時に使用する GS2-JobQueue のネームスペース のGRN */
	DoneTriggerQueueNamespaceId *string   `json:"doneTriggerQueueNamespaceId"`
}

func (p *ScriptSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["triggerScriptId"] = p.TriggerScriptId
    data["doneTriggerTargetType"] = p.DoneTriggerTargetType
    data["doneTriggerScriptId"] = p.DoneTriggerScriptId
    data["doneTriggerQueueNamespaceId"] = p.DoneTriggerQueueNamespaceId
    return &data
}

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *string   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}
