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

package enhance

type Namespace struct {
	/** ネームスペース */
	NamespaceId *string `json:"namespaceId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ネームスペース名 */
	Name *string `json:"name"`
	/** ネームスペースの説明 */
	Description *string `json:"description"`
	/** DirectEnhance を利用できるようにするか */
	EnableDirectEnhance *bool `json:"enableDirectEnhance"`
	/** 交換処理をジョブとして追加するキューのネームスペース のGRN */
	QueueNamespaceId *string `json:"queueNamespaceId"`
	/** 交換処理のスタンプシートで使用する暗号鍵GRN */
	KeyId *string `json:"keyId"`
	/** ログの出力設定 */
	LogSetting *LogSetting `json:"logSetting"`
	/** None */
	Status *string `json:"status"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *Namespace) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["namespaceId"] = p.NamespaceId
	data["ownerId"] = p.OwnerId
	data["name"] = p.Name
	data["description"] = p.Description
	data["enableDirectEnhance"] = p.EnableDirectEnhance
	data["queueNamespaceId"] = p.QueueNamespaceId
	data["keyId"] = p.KeyId
	if p.LogSetting != nil {
		data["logSetting"] = *p.LogSetting.ToDict()
	}
	data["status"] = p.Status
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type RateModel struct {
	/** 強化レートモデル */
	RateModelId *string `json:"rateModelId"`
	/** 強化レート名 */
	Name *string `json:"name"`
	/** 強化レートの説明 */
	Description *string `json:"description"`
	/** 強化レートのメタデータ */
	Metadata *string `json:"metadata"`
	/** 強化対象に使用できるインベントリモデル のGRN */
	TargetInventoryModelId *string `json:"targetInventoryModelId"`
	/** GS2-Experience で入手した経験値を格納する プロパティID に付与するサフィックス */
	AcquireExperienceSuffix *string `json:"acquireExperienceSuffix"`
	/** 強化素材に使用できるインベントリモデル のGRN */
	MaterialInventoryModelId *string `json:"materialInventoryModelId"`
	/** 入手経験値を格納しているメタデータのJSON階層 */
	AcquireExperienceHierarchy *[]string `json:"acquireExperienceHierarchy"`
	/** 獲得できる経験値の種類マスター のGRN */
	ExperienceModelId *string `json:"experienceModelId"`
	/** 経験値獲得量ボーナス */
	BonusRates *[]*BonusRate `json:"bonusRates"`
}

func (p *RateModel) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["rateModelId"] = p.RateModelId
	data["name"] = p.Name
	data["description"] = p.Description
	data["metadata"] = p.Metadata
	data["targetInventoryModelId"] = p.TargetInventoryModelId
	data["acquireExperienceSuffix"] = p.AcquireExperienceSuffix
	data["materialInventoryModelId"] = p.MaterialInventoryModelId
	if p.AcquireExperienceHierarchy != nil {
		var _acquireExperienceHierarchy []string
		for _, item := range *p.AcquireExperienceHierarchy {
			_acquireExperienceHierarchy = append(_acquireExperienceHierarchy, item)
		}
		data["acquireExperienceHierarchy"] = &_acquireExperienceHierarchy
	}
	data["experienceModelId"] = p.ExperienceModelId
	if p.BonusRates != nil {
		var _bonusRates []*map[string]interface{}
		for _, item := range *p.BonusRates {
			_bonusRates = append(_bonusRates, item.ToDict())
		}
		data["bonusRates"] = &_bonusRates
	}
	return &data
}

type RateModelMaster struct {
	/** 強化レートマスター */
	RateModelId *string `json:"rateModelId"`
	/** 強化レート名 */
	Name *string `json:"name"`
	/** 強化レートマスターの説明 */
	Description *string `json:"description"`
	/** 強化レートのメタデータ */
	Metadata *string `json:"metadata"`
	/** 強化対象に使用できるインベントリモデル のGRN */
	TargetInventoryModelId *string `json:"targetInventoryModelId"`
	/** GS2-Experience で入手した経験値を格納する プロパティID に付与するサフィックス */
	AcquireExperienceSuffix *string `json:"acquireExperienceSuffix"`
	/** 強化素材に使用できるインベントリモデル のGRN */
	MaterialInventoryModelId *string `json:"materialInventoryModelId"`
	/** 入手経験値を格納しているメタデータのJSON階層 */
	AcquireExperienceHierarchy *[]string `json:"acquireExperienceHierarchy"`
	/** 獲得できる経験値の種類マスター のGRN */
	ExperienceModelId *string `json:"experienceModelId"`
	/** 経験値獲得量ボーナス */
	BonusRates *[]*BonusRate `json:"bonusRates"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *RateModelMaster) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["rateModelId"] = p.RateModelId
	data["name"] = p.Name
	data["description"] = p.Description
	data["metadata"] = p.Metadata
	data["targetInventoryModelId"] = p.TargetInventoryModelId
	data["acquireExperienceSuffix"] = p.AcquireExperienceSuffix
	data["materialInventoryModelId"] = p.MaterialInventoryModelId
	if p.AcquireExperienceHierarchy != nil {
		var _acquireExperienceHierarchy []string
		for _, item := range *p.AcquireExperienceHierarchy {
			_acquireExperienceHierarchy = append(_acquireExperienceHierarchy, item)
		}
		data["acquireExperienceHierarchy"] = &_acquireExperienceHierarchy
	}
	data["experienceModelId"] = p.ExperienceModelId
	if p.BonusRates != nil {
		var _bonusRates []*map[string]interface{}
		for _, item := range *p.BonusRates {
			_bonusRates = append(_bonusRates, item.ToDict())
		}
		data["bonusRates"] = &_bonusRates
	}
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type Progress struct {
	/** 強化実行 */
	ProgressId *string `json:"progressId"`
	/** ユーザーID */
	UserId *string `json:"userId"`
	/** レートモデル名 */
	RateName *string `json:"rateName"`
	/** 強化対象のプロパティID */
	PropertyId *string `json:"propertyId"`
	/** 入手できる経験値 */
	ExperienceValue *int32 `json:"experienceValue"`
	/** 経験値倍率 */
	Rate *float32 `json:"rate"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *Progress) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["progressId"] = p.ProgressId
	data["userId"] = p.UserId
	data["rateName"] = p.RateName
	data["propertyId"] = p.PropertyId
	data["experienceValue"] = p.ExperienceValue
	data["rate"] = p.Rate
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type CurrentRateMaster struct {
	/** ネームスペース名 */
	NamespaceName *string `json:"namespaceName"`
	/** マスターデータ */
	Settings *string `json:"settings"`
}

func (p *CurrentRateMaster) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["namespaceName"] = p.NamespaceName
	data["settings"] = p.Settings
	return &data
}

type ResponseCache struct {
	/** None */
	Region *string `json:"region"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** レスポンスキャッシュ のGRN */
	ResponseCacheId *string `json:"responseCacheId"`
	/** None */
	RequestHash *string `json:"requestHash"`
	/** APIの応答内容 */
	Result *string `json:"result"`
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

type Config struct {
	/** 名前 */
	Key *string `json:"key"`
	/** 値 */
	Value *string `json:"value"`
}

func (p *Config) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["key"] = p.Key
	data["value"] = p.Value
	return &data
}

type GitHubCheckoutSetting struct {
	/** リソースの取得に使用するGitHub のAPIキー のGRN */
	GitHubApiKeyId *string `json:"gitHubApiKeyId"`
	/** リポジトリ名 */
	RepositoryName *string `json:"repositoryName"`
	/** ソースコードのファイルパス */
	SourcePath *string `json:"sourcePath"`
	/** コードの取得元 */
	ReferenceType *string `json:"referenceType"`
	/** コミットハッシュ */
	CommitHash *string `json:"commitHash"`
	/** ブランチ名 */
	BranchName *string `json:"branchName"`
	/** タグ名 */
	TagName *string `json:"tagName"`
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
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["loggingNamespaceId"] = p.LoggingNamespaceId
	return &data
}

type BonusRate struct {
	/** 経験値ボーナスの倍率(1.0=ボーナスなし) */
	Rate *float32 `json:"rate"`
	/** 抽選重み */
	Weight *int32 `json:"weight"`
}

func (p *BonusRate) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["rate"] = p.Rate
	data["weight"] = p.Weight
	return &data
}

type Material struct {
	/** 強化対象の GS2-Inventory アイテムセットGRN */
	MaterialItemSetId *string `json:"materialItemSetId"`
	/** 消費数量 */
	Count *int32 `json:"count"`
}

func (p *Material) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["materialItemSetId"] = p.MaterialItemSetId
	data["count"] = p.Count
	return &data
}
