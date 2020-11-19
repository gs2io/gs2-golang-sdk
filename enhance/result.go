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

type DescribeRateModelsResult struct {
	/** 強化レートモデルのリスト */
	Items *[]*RateModel `json:"items"`
}

func (p *DescribeRateModelsResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	return &data
}

type DescribeRateModelsAsyncResult struct {
	result *DescribeRateModelsResult
	err    error
}

type GetRateModelResult struct {
	/** 強化レートモデル */
	Item *RateModel `json:"item"`
}

func (p *GetRateModelResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetRateModelAsyncResult struct {
	result *GetRateModelResult
	err    error
}

type DescribeRateModelMastersResult struct {
	/** 強化レートマスターのリスト */
	Items *[]*RateModelMaster `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeRateModelMastersResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeRateModelMastersAsyncResult struct {
	result *DescribeRateModelMastersResult
	err    error
}

type CreateRateModelMasterResult struct {
	/** 作成した強化レートマスター */
	Item *RateModelMaster `json:"item"`
}

func (p *CreateRateModelMasterResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type CreateRateModelMasterAsyncResult struct {
	result *CreateRateModelMasterResult
	err    error
}

type GetRateModelMasterResult struct {
	/** 強化レートマスター */
	Item *RateModelMaster `json:"item"`
}

func (p *GetRateModelMasterResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetRateModelMasterAsyncResult struct {
	result *GetRateModelMasterResult
	err    error
}

type UpdateRateModelMasterResult struct {
	/** 更新した強化レートマスター */
	Item *RateModelMaster `json:"item"`
}

func (p *UpdateRateModelMasterResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateRateModelMasterAsyncResult struct {
	result *UpdateRateModelMasterResult
	err    error
}

type DeleteRateModelMasterResult struct {
	/** 削除した強化レートマスター */
	Item *RateModelMaster `json:"item"`
}

func (p *DeleteRateModelMasterResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteRateModelMasterAsyncResult struct {
	result *DeleteRateModelMasterResult
	err    error
}

type DirectEnhanceResult struct {
	/** 強化レートモデル */
	Item *RateModel `json:"item"`
	/** 強化処理の実行に使用するスタンプシート */
	StampSheet *string `json:"stampSheet"`
	/** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	/** 獲得経験値量 */
	AcquireExperience *int64 `json:"acquireExperience"`
	/** 経験値ボーナスの倍率(1.0=ボーナスなし) */
	BonusRate *float32 `json:"bonusRate"`
}

func (p *DirectEnhanceResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.StampSheet != nil {
		data["StampSheet"] = p.StampSheet
	}
	if p.StampSheetEncryptionKeyId != nil {
		data["StampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
	}
	if p.AcquireExperience != nil {
		data["AcquireExperience"] = p.AcquireExperience
	}
	if p.BonusRate != nil {
		data["BonusRate"] = p.BonusRate
	}
	return &data
}

type DirectEnhanceAsyncResult struct {
	result *DirectEnhanceResult
	err    error
}

type DirectEnhanceByUserIdResult struct {
	/** 強化レートモデル */
	Item *RateModel `json:"item"`
	/** 強化処理の実行に使用するスタンプシート */
	StampSheet *string `json:"stampSheet"`
	/** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	/** 獲得経験値量 */
	AcquireExperience *int64 `json:"acquireExperience"`
	/** 経験値ボーナスの倍率(1.0=ボーナスなし) */
	BonusRate *float32 `json:"bonusRate"`
}

func (p *DirectEnhanceByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.StampSheet != nil {
		data["StampSheet"] = p.StampSheet
	}
	if p.StampSheetEncryptionKeyId != nil {
		data["StampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
	}
	if p.AcquireExperience != nil {
		data["AcquireExperience"] = p.AcquireExperience
	}
	if p.BonusRate != nil {
		data["BonusRate"] = p.BonusRate
	}
	return &data
}

type DirectEnhanceByUserIdAsyncResult struct {
	result *DirectEnhanceByUserIdResult
	err    error
}

type DirectEnhanceByStampSheetResult struct {
	/** 強化レートモデル */
	Item *RateModel `json:"item"`
	/** 強化処理の実行に使用するスタンプシート */
	StampSheet *string `json:"stampSheet"`
	/** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	/** 獲得経験値量 */
	AcquireExperience *int64 `json:"acquireExperience"`
	/** 経験値ボーナスの倍率(1.0=ボーナスなし) */
	BonusRate *float32 `json:"bonusRate"`
}

func (p *DirectEnhanceByStampSheetResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.StampSheet != nil {
		data["StampSheet"] = p.StampSheet
	}
	if p.StampSheetEncryptionKeyId != nil {
		data["StampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
	}
	if p.AcquireExperience != nil {
		data["AcquireExperience"] = p.AcquireExperience
	}
	if p.BonusRate != nil {
		data["BonusRate"] = p.BonusRate
	}
	return &data
}

type DirectEnhanceByStampSheetAsyncResult struct {
	result *DirectEnhanceByStampSheetResult
	err    error
}

type DescribeProgressesByUserIdResult struct {
	/** 強化実行のリスト */
	Items *[]*Progress `json:"items"`
	/** リストの続きを取得するためのページトークン */
	NextPageToken *string `json:"nextPageToken"`
}

func (p *DescribeProgressesByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Items != nil {
	}
	if p.NextPageToken != nil {
		data["NextPageToken"] = p.NextPageToken
	}
	return &data
}

type DescribeProgressesByUserIdAsyncResult struct {
	result *DescribeProgressesByUserIdResult
	err    error
}

type CreateProgressByUserIdResult struct {
	/** 強化実行 */
	Item *Progress `json:"item"`
}

func (p *CreateProgressByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type CreateProgressByUserIdAsyncResult struct {
	result *CreateProgressByUserIdResult
	err    error
}

type GetProgressResult struct {
	/** 強化実行 */
	Item *Progress `json:"item"`
	/** 強化レートモデル */
	RateModel *RateModel `json:"rateModel"`
}

func (p *GetProgressResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.RateModel != nil {
		data["RateModel"] = p.RateModel.ToDict()
	}
	return &data
}

type GetProgressAsyncResult struct {
	result *GetProgressResult
	err    error
}

type GetProgressByUserIdResult struct {
	/** 強化実行 */
	Item *Progress `json:"item"`
	/** 強化レートモデル */
	RateModel *RateModel `json:"rateModel"`
}

func (p *GetProgressByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.RateModel != nil {
		data["RateModel"] = p.RateModel.ToDict()
	}
	return &data
}

type GetProgressByUserIdAsyncResult struct {
	result *GetProgressByUserIdResult
	err    error
}

type StartResult struct {
	/** 強化の開始処理の実行に使用するスタンプシート */
	StampSheet *string `json:"stampSheet"`
	/** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

func (p *StartResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.StampSheet != nil {
		data["StampSheet"] = p.StampSheet
	}
	if p.StampSheetEncryptionKeyId != nil {
		data["StampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
	}
	return &data
}

type StartAsyncResult struct {
	result *StartResult
	err    error
}

type StartByUserIdResult struct {
	/** 強化の開始処理の実行に使用するスタンプシート */
	StampSheet *string `json:"stampSheet"`
	/** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

func (p *StartByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.StampSheet != nil {
		data["StampSheet"] = p.StampSheet
	}
	if p.StampSheetEncryptionKeyId != nil {
		data["StampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
	}
	return &data
}

type StartByUserIdAsyncResult struct {
	result *StartByUserIdResult
	err    error
}

type EndResult struct {
	/** 強化実行 */
	Item *Progress `json:"item"`
	/** 報酬付与処理の実行に使用するスタンプシート */
	StampSheet *string `json:"stampSheet"`
	/** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	/** 獲得経験値量 */
	AcquireExperience *int64 `json:"acquireExperience"`
	/** 経験値ボーナスの倍率(1.0=ボーナスなし) */
	BonusRate *float32 `json:"bonusRate"`
}

func (p *EndResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.StampSheet != nil {
		data["StampSheet"] = p.StampSheet
	}
	if p.StampSheetEncryptionKeyId != nil {
		data["StampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
	}
	if p.AcquireExperience != nil {
		data["AcquireExperience"] = p.AcquireExperience
	}
	if p.BonusRate != nil {
		data["BonusRate"] = p.BonusRate
	}
	return &data
}

type EndAsyncResult struct {
	result *EndResult
	err    error
}

type EndByUserIdResult struct {
	/** 強化実行 */
	Item *Progress `json:"item"`
	/** 報酬付与処理の実行に使用するスタンプシート */
	StampSheet *string `json:"stampSheet"`
	/** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	/** 獲得経験値量 */
	AcquireExperience *int64 `json:"acquireExperience"`
	/** 経験値ボーナスの倍率(1.0=ボーナスなし) */
	BonusRate *float32 `json:"bonusRate"`
}

func (p *EndByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.StampSheet != nil {
		data["StampSheet"] = p.StampSheet
	}
	if p.StampSheetEncryptionKeyId != nil {
		data["StampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
	}
	if p.AcquireExperience != nil {
		data["AcquireExperience"] = p.AcquireExperience
	}
	if p.BonusRate != nil {
		data["BonusRate"] = p.BonusRate
	}
	return &data
}

type EndByUserIdAsyncResult struct {
	result *EndByUserIdResult
	err    error
}

type DeleteProgressResult struct {
	/** 強化実行 */
	Item *Progress `json:"item"`
}

func (p *DeleteProgressResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteProgressAsyncResult struct {
	result *DeleteProgressResult
	err    error
}

type DeleteProgressByUserIdResult struct {
	/** 強化実行 */
	Item *Progress `json:"item"`
}

func (p *DeleteProgressByUserIdResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type DeleteProgressByUserIdAsyncResult struct {
	result *DeleteProgressByUserIdResult
	err    error
}

type CreateProgressByStampSheetResult struct {
	/** 強化実行 */
	Item *Progress `json:"item"`
}

func (p *CreateProgressByStampSheetResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type CreateProgressByStampSheetAsyncResult struct {
	result *CreateProgressByStampSheetResult
	err    error
}

type DeleteProgressByStampTaskResult struct {
	/** 強化実行 */
	Item *Progress `json:"item"`
	/** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack *string `json:"newContextStack"`
}

func (p *DeleteProgressByStampTaskResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	if p.NewContextStack != nil {
		data["NewContextStack"] = p.NewContextStack
	}
	return &data
}

type DeleteProgressByStampTaskAsyncResult struct {
	result *DeleteProgressByStampTaskResult
	err    error
}

type ExportMasterResult struct {
	/** 現在有効な強化レート設定 */
	Item *CurrentRateMaster `json:"item"`
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

type GetCurrentRateMasterResult struct {
	/** 現在有効な強化レート設定 */
	Item *CurrentRateMaster `json:"item"`
}

func (p *GetCurrentRateMasterResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type GetCurrentRateMasterAsyncResult struct {
	result *GetCurrentRateMasterResult
	err    error
}

type UpdateCurrentRateMasterResult struct {
	/** 更新した現在有効な強化レート設定 */
	Item *CurrentRateMaster `json:"item"`
}

func (p *UpdateCurrentRateMasterResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateCurrentRateMasterAsyncResult struct {
	result *UpdateCurrentRateMasterResult
	err    error
}

type UpdateCurrentRateMasterFromGitHubResult struct {
	/** 更新した現在有効な強化レート設定 */
	Item *CurrentRateMaster `json:"item"`
}

func (p *UpdateCurrentRateMasterFromGitHubResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Item != nil {
		data["Item"] = p.Item.ToDict()
	}
	return &data
}

type UpdateCurrentRateMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRateMasterFromGitHubResult
	err    error
}
