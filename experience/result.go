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

type DescribeExperienceModelMastersResult struct {
    /** 経験値の種類マスターのリスト */
	Items         []ExperienceModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeExperienceModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]ExperienceModelMaster, 0)
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

type DescribeExperienceModelMastersAsyncResult struct {
	result *DescribeExperienceModelMastersResult
	err    error
}

type CreateExperienceModelMasterResult struct {
    /** 作成した経験値の種類マスター */
	Item         *ExperienceModelMaster	`json:"item"`
}

func (p *CreateExperienceModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateExperienceModelMasterAsyncResult struct {
	result *CreateExperienceModelMasterResult
	err    error
}

type GetExperienceModelMasterResult struct {
    /** 経験値の種類マスター */
	Item         *ExperienceModelMaster	`json:"item"`
}

func (p *GetExperienceModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetExperienceModelMasterAsyncResult struct {
	result *GetExperienceModelMasterResult
	err    error
}

type UpdateExperienceModelMasterResult struct {
    /** 更新した経験値の種類マスター */
	Item         *ExperienceModelMaster	`json:"item"`
}

func (p *UpdateExperienceModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateExperienceModelMasterAsyncResult struct {
	result *UpdateExperienceModelMasterResult
	err    error
}

type DeleteExperienceModelMasterResult struct {
    /** 削除した経験値の種類マスター */
	Item         *ExperienceModelMaster	`json:"item"`
}

func (p *DeleteExperienceModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteExperienceModelMasterAsyncResult struct {
	result *DeleteExperienceModelMasterResult
	err    error
}

type DescribeExperienceModelsResult struct {
    /** 経験値・ランクアップ閾値モデルのリスト */
	Items         []ExperienceModel	`json:"items"`
}

func (p *DescribeExperienceModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]ExperienceModel, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeExperienceModelsAsyncResult struct {
	result *DescribeExperienceModelsResult
	err    error
}

type GetExperienceModelResult struct {
    /** 経験値・ランクアップ閾値モデル */
	Item         *ExperienceModel	`json:"item"`
}

func (p *GetExperienceModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetExperienceModelAsyncResult struct {
	result *GetExperienceModelResult
	err    error
}

type DescribeThresholdMastersResult struct {
    /** ランクアップ閾値マスターのリスト */
	Items         []ThresholdMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeThresholdMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]ThresholdMaster, 0)
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

type DescribeThresholdMastersAsyncResult struct {
	result *DescribeThresholdMastersResult
	err    error
}

type CreateThresholdMasterResult struct {
    /** 作成したランクアップ閾値マスター */
	Item         *ThresholdMaster	`json:"item"`
}

func (p *CreateThresholdMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateThresholdMasterAsyncResult struct {
	result *CreateThresholdMasterResult
	err    error
}

type GetThresholdMasterResult struct {
    /** ランクアップ閾値マスター */
	Item         *ThresholdMaster	`json:"item"`
}

func (p *GetThresholdMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetThresholdMasterAsyncResult struct {
	result *GetThresholdMasterResult
	err    error
}

type UpdateThresholdMasterResult struct {
    /** 更新したランクアップ閾値マスター */
	Item         *ThresholdMaster	`json:"item"`
}

func (p *UpdateThresholdMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateThresholdMasterAsyncResult struct {
	result *UpdateThresholdMasterResult
	err    error
}

type DeleteThresholdMasterResult struct {
    /** 削除したランクアップ閾値マスター */
	Item         *ThresholdMaster	`json:"item"`
}

func (p *DeleteThresholdMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteThresholdMasterAsyncResult struct {
	result *DeleteThresholdMasterResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効な経験値設定 */
	Item         *CurrentExperienceMaster	`json:"item"`
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

type GetCurrentExperienceMasterResult struct {
    /** 現在有効な経験値設定 */
	Item         *CurrentExperienceMaster	`json:"item"`
}

func (p *GetCurrentExperienceMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentExperienceMasterAsyncResult struct {
	result *GetCurrentExperienceMasterResult
	err    error
}

type UpdateCurrentExperienceMasterResult struct {
    /** 更新した現在有効な経験値設定 */
	Item         *CurrentExperienceMaster	`json:"item"`
}

func (p *UpdateCurrentExperienceMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentExperienceMasterAsyncResult struct {
	result *UpdateCurrentExperienceMasterResult
	err    error
}

type UpdateCurrentExperienceMasterFromGitHubResult struct {
    /** 更新した現在有効な経験値設定 */
	Item         *CurrentExperienceMaster	`json:"item"`
}

func (p *UpdateCurrentExperienceMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentExperienceMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentExperienceMasterFromGitHubResult
	err    error
}

type DescribeStatusesResult struct {
    /** ステータスのリスト */
	Items         []Status	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeStatusesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Status, 0)
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

type DescribeStatusesAsyncResult struct {
	result *DescribeStatusesResult
	err    error
}

type DescribeStatusesByUserIdResult struct {
    /** ステータスのリスト */
	Items         []Status	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeStatusesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Status, 0)
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

type DescribeStatusesByUserIdAsyncResult struct {
	result *DescribeStatusesByUserIdResult
	err    error
}

type GetStatusResult struct {
    /** ステータス */
	Item         *Status	`json:"item"`
}

func (p *GetStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetStatusAsyncResult struct {
	result *GetStatusResult
	err    error
}

type GetStatusByUserIdResult struct {
    /** ステータス */
	Item         *Status	`json:"item"`
}

func (p *GetStatusByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetStatusByUserIdAsyncResult struct {
	result *GetStatusByUserIdResult
	err    error
}

type GetStatusWithSignatureResult struct {
    /** ステータス */
	Item         *Status	`json:"item"`
    /** 検証対象のオブジェクト */
	Body         *string	`json:"body"`
    /** 署名 */
	Signature         *string	`json:"signature"`
}

func (p *GetStatusWithSignatureResult) ToDict() *map[string]interface{} {
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

type GetStatusWithSignatureAsyncResult struct {
	result *GetStatusWithSignatureResult
	err    error
}

type AddExperienceByUserIdResult struct {
    /** 加算後のステータス */
	Item         *Status	`json:"item"`
}

func (p *AddExperienceByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type AddExperienceByUserIdAsyncResult struct {
	result *AddExperienceByUserIdResult
	err    error
}

type SetExperienceByUserIdResult struct {
    /** 更新後のステータス */
	Item         *Status	`json:"item"`
}

func (p *SetExperienceByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SetExperienceByUserIdAsyncResult struct {
	result *SetExperienceByUserIdResult
	err    error
}

type AddRankCapByUserIdResult struct {
    /** 加算後のステータス */
	Item         *Status	`json:"item"`
}

func (p *AddRankCapByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type AddRankCapByUserIdAsyncResult struct {
	result *AddRankCapByUserIdResult
	err    error
}

type SetRankCapByUserIdResult struct {
    /** 更新後のステータス */
	Item         *Status	`json:"item"`
}

func (p *SetRankCapByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SetRankCapByUserIdAsyncResult struct {
	result *SetRankCapByUserIdResult
	err    error
}

type DeleteStatusByUserIdResult struct {
    /** ステータス */
	Item         *Status	`json:"item"`
}

func (p *DeleteStatusByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteStatusByUserIdAsyncResult struct {
	result *DeleteStatusByUserIdResult
	err    error
}

type AddExperienceByStampSheetResult struct {
    /** 加算後のステータス */
	Item         *Status	`json:"item"`
}

func (p *AddExperienceByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type AddExperienceByStampSheetAsyncResult struct {
	result *AddExperienceByStampSheetResult
	err    error
}

type AddRankCapByStampSheetResult struct {
    /** 加算後のステータス */
	Item         *Status	`json:"item"`
}

func (p *AddRankCapByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type AddRankCapByStampSheetAsyncResult struct {
	result *AddRankCapByStampSheetResult
	err    error
}

type SetRankCapByStampSheetResult struct {
    /** 更新後のステータス */
	Item         *Status	`json:"item"`
}

func (p *SetRankCapByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SetRankCapByStampSheetAsyncResult struct {
	result *SetRankCapByStampSheetResult
	err    error
}
