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

package formation

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

type DescribeFormModelMastersResult struct {
    /** フォームマスターのリスト */
	Items         *[]*FormModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeFormModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeFormModelMastersAsyncResult struct {
	result *DescribeFormModelMastersResult
	err    error
}

type CreateFormModelMasterResult struct {
    /** 作成したフォームマスター */
	Item         *FormModelMaster	`json:"item"`
}

func (p *CreateFormModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateFormModelMasterAsyncResult struct {
	result *CreateFormModelMasterResult
	err    error
}

type GetFormModelMasterResult struct {
    /** フォームマスター */
	Item         *FormModelMaster	`json:"item"`
}

func (p *GetFormModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetFormModelMasterAsyncResult struct {
	result *GetFormModelMasterResult
	err    error
}

type UpdateFormModelMasterResult struct {
    /** 更新したフォームマスター */
	Item         *FormModelMaster	`json:"item"`
}

func (p *UpdateFormModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateFormModelMasterAsyncResult struct {
	result *UpdateFormModelMasterResult
	err    error
}

type DeleteFormModelMasterResult struct {
    /** 削除したフォームマスター */
	Item         *FormModelMaster	`json:"item"`
}

func (p *DeleteFormModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteFormModelMasterAsyncResult struct {
	result *DeleteFormModelMasterResult
	err    error
}

type DescribeMoldModelsResult struct {
    /** フォームの保存領域のリスト */
	Items         *[]*MoldModel	`json:"items"`
}

func (p *DescribeMoldModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    return &data
}

type DescribeMoldModelsAsyncResult struct {
	result *DescribeMoldModelsResult
	err    error
}

type GetMoldModelResult struct {
    /** フォームの保存領域 */
	Item         *MoldModel	`json:"item"`
}

func (p *GetMoldModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetMoldModelAsyncResult struct {
	result *GetMoldModelResult
	err    error
}

type DescribeMoldModelMastersResult struct {
    /** フォームの保存領域マスターのリスト */
	Items         *[]*MoldModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeMoldModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeMoldModelMastersAsyncResult struct {
	result *DescribeMoldModelMastersResult
	err    error
}

type CreateMoldModelMasterResult struct {
    /** 作成したフォームの保存領域マスター */
	Item         *MoldModelMaster	`json:"item"`
}

func (p *CreateMoldModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateMoldModelMasterAsyncResult struct {
	result *CreateMoldModelMasterResult
	err    error
}

type GetMoldModelMasterResult struct {
    /** フォームの保存領域マスター */
	Item         *MoldModelMaster	`json:"item"`
}

func (p *GetMoldModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetMoldModelMasterAsyncResult struct {
	result *GetMoldModelMasterResult
	err    error
}

type UpdateMoldModelMasterResult struct {
    /** 更新したフォームの保存領域マスター */
	Item         *MoldModelMaster	`json:"item"`
}

func (p *UpdateMoldModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateMoldModelMasterAsyncResult struct {
	result *UpdateMoldModelMasterResult
	err    error
}

type DeleteMoldModelMasterResult struct {
    /** 削除したフォームの保存領域マスター */
	Item         *MoldModelMaster	`json:"item"`
}

func (p *DeleteMoldModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteMoldModelMasterAsyncResult struct {
	result *DeleteMoldModelMasterResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なフォーム設定 */
	Item         *CurrentFormMaster	`json:"item"`
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

type GetCurrentFormMasterResult struct {
    /** 現在有効なフォーム設定 */
	Item         *CurrentFormMaster	`json:"item"`
}

func (p *GetCurrentFormMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentFormMasterAsyncResult struct {
	result *GetCurrentFormMasterResult
	err    error
}

type UpdateCurrentFormMasterResult struct {
    /** 更新した現在有効なフォーム設定 */
	Item         *CurrentFormMaster	`json:"item"`
}

func (p *UpdateCurrentFormMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentFormMasterAsyncResult struct {
	result *UpdateCurrentFormMasterResult
	err    error
}

type UpdateCurrentFormMasterFromGitHubResult struct {
    /** 更新した現在有効なフォーム設定 */
	Item         *CurrentFormMaster	`json:"item"`
}

func (p *UpdateCurrentFormMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentFormMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentFormMasterFromGitHubResult
	err    error
}

type DescribeMoldsResult struct {
    /** 保存したフォームのリスト */
	Items         *[]*Mold	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeMoldsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeMoldsAsyncResult struct {
	result *DescribeMoldsResult
	err    error
}

type DescribeMoldsByUserIdResult struct {
    /** 保存したフォームのリスト */
	Items         *[]*Mold	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeMoldsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeMoldsByUserIdAsyncResult struct {
	result *DescribeMoldsByUserIdResult
	err    error
}

type GetMoldResult struct {
    /** 保存したフォーム */
	Item         *Mold	`json:"item"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
}

func (p *GetMoldResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    return &data
}

type GetMoldAsyncResult struct {
	result *GetMoldResult
	err    error
}

type GetMoldByUserIdResult struct {
    /** 保存したフォーム */
	Item         *Mold	`json:"item"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
}

func (p *GetMoldByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    return &data
}

type GetMoldByUserIdAsyncResult struct {
	result *GetMoldByUserIdResult
	err    error
}

type SetMoldCapacityByUserIdResult struct {
    /** キャパシティを更新した保存したフォーム */
	Item         *Mold	`json:"item"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
}

func (p *SetMoldCapacityByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    return &data
}

type SetMoldCapacityByUserIdAsyncResult struct {
	result *SetMoldCapacityByUserIdResult
	err    error
}

type AddMoldCapacityByUserIdResult struct {
    /** キャパシティを更新した保存したフォーム */
	Item         *Mold	`json:"item"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
}

func (p *AddMoldCapacityByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    return &data
}

type AddMoldCapacityByUserIdAsyncResult struct {
	result *AddMoldCapacityByUserIdResult
	err    error
}

type DeleteMoldResult struct {
    /** 保存したフォーム */
	Item         *Mold	`json:"item"`
}

func (p *DeleteMoldResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteMoldAsyncResult struct {
	result *DeleteMoldResult
	err    error
}

type DeleteMoldByUserIdResult struct {
    /** 保存したフォーム */
	Item         *Mold	`json:"item"`
}

func (p *DeleteMoldByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteMoldByUserIdAsyncResult struct {
	result *DeleteMoldByUserIdResult
	err    error
}

type AddCapacityByStampSheetResult struct {
    /** キャパシティ加算後の保存したフォーム */
	Item         *Mold	`json:"item"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
}

func (p *AddCapacityByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    return &data
}

type AddCapacityByStampSheetAsyncResult struct {
	result *AddCapacityByStampSheetResult
	err    error
}

type SetCapacityByStampSheetResult struct {
    /** 更新後の保存したフォーム */
	Item         *Mold	`json:"item"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
}

func (p *SetCapacityByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    return &data
}

type SetCapacityByStampSheetAsyncResult struct {
	result *SetCapacityByStampSheetResult
	err    error
}

type DescribeFormsResult struct {
    /** フォームのリスト */
	Items         *[]*Form	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeFormsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeFormsAsyncResult struct {
	result *DescribeFormsResult
	err    error
}

type DescribeFormsByUserIdResult struct {
    /** フォームのリスト */
	Items         *[]*Form	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeFormsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeFormsByUserIdAsyncResult struct {
	result *DescribeFormsByUserIdResult
	err    error
}

type GetFormResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
    /** フォームモデル */
	FormModel         *FormModel	`json:"formModel"`
}

func (p *GetFormResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    if p.FormModel != nil {
     data["FormModel"] = p.FormModel.ToDict()
    }
    return &data
}

type GetFormAsyncResult struct {
	result *GetFormResult
	err    error
}

type GetFormByUserIdResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
    /** フォームモデル */
	FormModel         *FormModel	`json:"formModel"`
}

func (p *GetFormByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    if p.FormModel != nil {
     data["FormModel"] = p.FormModel.ToDict()
    }
    return &data
}

type GetFormByUserIdAsyncResult struct {
	result *GetFormByUserIdResult
	err    error
}

type GetFormWithSignatureResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 署名対象の値 */
	Body         *core.String	`json:"body"`
    /** 署名 */
	Signature         *core.String	`json:"signature"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
    /** フォームモデル */
	FormModel         *FormModel	`json:"formModel"`
}

func (p *GetFormWithSignatureResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Body != nil {
     data["Body"] = p.Body
    }
    if p.Signature != nil {
     data["Signature"] = p.Signature
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    if p.FormModel != nil {
     data["FormModel"] = p.FormModel.ToDict()
    }
    return &data
}

type GetFormWithSignatureAsyncResult struct {
	result *GetFormWithSignatureResult
	err    error
}

type GetFormWithSignatureByUserIdResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 署名対象の値 */
	Body         *core.String	`json:"body"`
    /** 署名 */
	Signature         *core.String	`json:"signature"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
    /** フォームモデル */
	FormModel         *FormModel	`json:"formModel"`
}

func (p *GetFormWithSignatureByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Body != nil {
     data["Body"] = p.Body
    }
    if p.Signature != nil {
     data["Signature"] = p.Signature
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    if p.FormModel != nil {
     data["FormModel"] = p.FormModel.ToDict()
    }
    return &data
}

type GetFormWithSignatureByUserIdAsyncResult struct {
	result *GetFormWithSignatureByUserIdResult
	err    error
}

type SetFormByUserIdResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
    /** フォームモデル */
	FormModel         *FormModel	`json:"formModel"`
}

func (p *SetFormByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    if p.FormModel != nil {
     data["FormModel"] = p.FormModel.ToDict()
    }
    return &data
}

type SetFormByUserIdAsyncResult struct {
	result *SetFormByUserIdResult
	err    error
}

type SetFormWithSignatureResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
    /** フォームモデル */
	FormModel         *FormModel	`json:"formModel"`
}

func (p *SetFormWithSignatureResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    if p.FormModel != nil {
     data["FormModel"] = p.FormModel.ToDict()
    }
    return &data
}

type SetFormWithSignatureAsyncResult struct {
	result *SetFormWithSignatureResult
	err    error
}

type AcquireActionsToFormPropertiesResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** スタンプシート */
	StampSheet         *core.String	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *core.String	`json:"stampSheetEncryptionKeyId"`
}

func (p *AcquireActionsToFormPropertiesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.StampSheet != nil {
     data["StampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
     data["StampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type AcquireActionsToFormPropertiesAsyncResult struct {
	result *AcquireActionsToFormPropertiesResult
	err    error
}

type DeleteFormResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
    /** フォームモデル */
	FormModel         *FormModel	`json:"formModel"`
}

func (p *DeleteFormResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    if p.FormModel != nil {
     data["FormModel"] = p.FormModel.ToDict()
    }
    return &data
}

type DeleteFormAsyncResult struct {
	result *DeleteFormResult
	err    error
}

type DeleteFormByUserIdResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** フォームの保存領域 */
	MoldModel         *MoldModel	`json:"moldModel"`
    /** フォームモデル */
	FormModel         *FormModel	`json:"formModel"`
}

func (p *DeleteFormByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.MoldModel != nil {
     data["MoldModel"] = p.MoldModel.ToDict()
    }
    if p.FormModel != nil {
     data["FormModel"] = p.FormModel.ToDict()
    }
    return &data
}

type DeleteFormByUserIdAsyncResult struct {
	result *DeleteFormByUserIdResult
	err    error
}

type AcquireActionToFormPropertiesByStampSheetResult struct {
    /** フォーム */
	Item         *Form	`json:"item"`
    /** 保存したフォーム */
	Mold         *Mold	`json:"mold"`
    /** スタンプシート */
	StampSheet         *core.String	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *core.String	`json:"stampSheetEncryptionKeyId"`
}

func (p *AcquireActionToFormPropertiesByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Mold != nil {
     data["Mold"] = p.Mold.ToDict()
    }
    if p.StampSheet != nil {
     data["StampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
     data["StampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type AcquireActionToFormPropertiesByStampSheetAsyncResult struct {
	result *AcquireActionToFormPropertiesByStampSheetResult
	err    error
}
