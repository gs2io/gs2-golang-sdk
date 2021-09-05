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

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
	Items         []Namespace `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromJson(data string) DescribeNamespacesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNamespacesResultFromDict(dict)
}

func NewDescribeNamespacesResultFromDict(data map[string]interface{}) DescribeNamespacesResult {
	return DescribeNamespacesResult{
		Items:         CastNamespaces(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNamespacesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeNamespacesResult) Pointer() *DescribeNamespacesResult {
	return &p
}

type CreateNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

func NewCreateNamespaceResultFromJson(data string) CreateNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceResultFromDict(dict)
}

func NewCreateNamespaceResultFromDict(data map[string]interface{}) CreateNamespaceResult {
	return CreateNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateNamespaceResult) Pointer() *CreateNamespaceResult {
	return &p
}

type GetNamespaceStatusResult struct {
	Status *string `json:"status"`
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

func NewGetNamespaceStatusResultFromJson(data string) GetNamespaceStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceStatusResultFromDict(dict)
}

func NewGetNamespaceStatusResultFromDict(data map[string]interface{}) GetNamespaceStatusResult {
	return GetNamespaceStatusResult{
		Status: core.CastString(data["status"]),
	}
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"status": p.Status,
	}
}

func (p GetNamespaceStatusResult) Pointer() *GetNamespaceStatusResult {
	return &p
}

type GetNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

func NewGetNamespaceResultFromJson(data string) GetNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceResultFromDict(dict)
}

func NewGetNamespaceResultFromDict(data map[string]interface{}) GetNamespaceResult {
	return GetNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetNamespaceResult) Pointer() *GetNamespaceResult {
	return &p
}

type UpdateNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

func NewUpdateNamespaceResultFromJson(data string) UpdateNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceResultFromDict(dict)
}

func NewUpdateNamespaceResultFromDict(data map[string]interface{}) UpdateNamespaceResult {
	return UpdateNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateNamespaceResult) Pointer() *UpdateNamespaceResult {
	return &p
}

type DeleteNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromJson(data string) DeleteNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteNamespaceResultFromDict(dict)
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
	return DeleteNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
	return &p
}

type DescribeRateModelsResult struct {
	Items []RateModel `json:"items"`
}

type DescribeRateModelsAsyncResult struct {
	result *DescribeRateModelsResult
	err    error
}

func NewDescribeRateModelsResultFromJson(data string) DescribeRateModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRateModelsResultFromDict(dict)
}

func NewDescribeRateModelsResultFromDict(data map[string]interface{}) DescribeRateModelsResult {
	return DescribeRateModelsResult{
		Items: CastRateModels(core.CastArray(data["items"])),
	}
}

func (p DescribeRateModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRateModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeRateModelsResult) Pointer() *DescribeRateModelsResult {
	return &p
}

type GetRateModelResult struct {
	Item *RateModel `json:"item"`
}

type GetRateModelAsyncResult struct {
	result *GetRateModelResult
	err    error
}

func NewGetRateModelResultFromJson(data string) GetRateModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRateModelResultFromDict(dict)
}

func NewGetRateModelResultFromDict(data map[string]interface{}) GetRateModelResult {
	return GetRateModelResult{
		Item: NewRateModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRateModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRateModelResult) Pointer() *GetRateModelResult {
	return &p
}

type DescribeRateModelMastersResult struct {
	Items         []RateModelMaster `json:"items"`
	NextPageToken *string           `json:"nextPageToken"`
}

type DescribeRateModelMastersAsyncResult struct {
	result *DescribeRateModelMastersResult
	err    error
}

func NewDescribeRateModelMastersResultFromJson(data string) DescribeRateModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRateModelMastersResultFromDict(dict)
}

func NewDescribeRateModelMastersResultFromDict(data map[string]interface{}) DescribeRateModelMastersResult {
	return DescribeRateModelMastersResult{
		Items:         CastRateModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRateModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRateModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRateModelMastersResult) Pointer() *DescribeRateModelMastersResult {
	return &p
}

type CreateRateModelMasterResult struct {
	Item *RateModelMaster `json:"item"`
}

type CreateRateModelMasterAsyncResult struct {
	result *CreateRateModelMasterResult
	err    error
}

func NewCreateRateModelMasterResultFromJson(data string) CreateRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRateModelMasterResultFromDict(dict)
}

func NewCreateRateModelMasterResultFromDict(data map[string]interface{}) CreateRateModelMasterResult {
	return CreateRateModelMasterResult{
		Item: NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateRateModelMasterResult) Pointer() *CreateRateModelMasterResult {
	return &p
}

type GetRateModelMasterResult struct {
	Item *RateModelMaster `json:"item"`
}

type GetRateModelMasterAsyncResult struct {
	result *GetRateModelMasterResult
	err    error
}

func NewGetRateModelMasterResultFromJson(data string) GetRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRateModelMasterResultFromDict(dict)
}

func NewGetRateModelMasterResultFromDict(data map[string]interface{}) GetRateModelMasterResult {
	return GetRateModelMasterResult{
		Item: NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRateModelMasterResult) Pointer() *GetRateModelMasterResult {
	return &p
}

type UpdateRateModelMasterResult struct {
	Item *RateModelMaster `json:"item"`
}

type UpdateRateModelMasterAsyncResult struct {
	result *UpdateRateModelMasterResult
	err    error
}

func NewUpdateRateModelMasterResultFromJson(data string) UpdateRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRateModelMasterResultFromDict(dict)
}

func NewUpdateRateModelMasterResultFromDict(data map[string]interface{}) UpdateRateModelMasterResult {
	return UpdateRateModelMasterResult{
		Item: NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateRateModelMasterResult) Pointer() *UpdateRateModelMasterResult {
	return &p
}

type DeleteRateModelMasterResult struct {
	Item *RateModelMaster `json:"item"`
}

type DeleteRateModelMasterAsyncResult struct {
	result *DeleteRateModelMasterResult
	err    error
}

func NewDeleteRateModelMasterResultFromJson(data string) DeleteRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRateModelMasterResultFromDict(dict)
}

func NewDeleteRateModelMasterResultFromDict(data map[string]interface{}) DeleteRateModelMasterResult {
	return DeleteRateModelMasterResult{
		Item: NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteRateModelMasterResult) Pointer() *DeleteRateModelMasterResult {
	return &p
}

type DirectEnhanceResult struct {
	Item                      *RateModel `json:"item"`
	StampSheet                *string    `json:"stampSheet"`
	StampSheetEncryptionKeyId *string    `json:"stampSheetEncryptionKeyId"`
	AcquireExperience         *int64     `json:"acquireExperience"`
	BonusRate                 *float32   `json:"bonusRate"`
}

type DirectEnhanceAsyncResult struct {
	result *DirectEnhanceResult
	err    error
}

func NewDirectEnhanceResultFromJson(data string) DirectEnhanceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDirectEnhanceResultFromDict(dict)
}

func NewDirectEnhanceResultFromDict(data map[string]interface{}) DirectEnhanceResult {
	return DirectEnhanceResult{
		Item:                      NewRateModelFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AcquireExperience:         core.CastInt64(data["acquireExperience"]),
		BonusRate:                 core.CastFloat32(data["bonusRate"]),
	}
}

func (p DirectEnhanceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"acquireExperience":         p.AcquireExperience,
		"bonusRate":                 p.BonusRate,
	}
}

func (p DirectEnhanceResult) Pointer() *DirectEnhanceResult {
	return &p
}

type DirectEnhanceByUserIdResult struct {
	Item                      *RateModel `json:"item"`
	StampSheet                *string    `json:"stampSheet"`
	StampSheetEncryptionKeyId *string    `json:"stampSheetEncryptionKeyId"`
	AcquireExperience         *int64     `json:"acquireExperience"`
	BonusRate                 *float32   `json:"bonusRate"`
}

type DirectEnhanceByUserIdAsyncResult struct {
	result *DirectEnhanceByUserIdResult
	err    error
}

func NewDirectEnhanceByUserIdResultFromJson(data string) DirectEnhanceByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDirectEnhanceByUserIdResultFromDict(dict)
}

func NewDirectEnhanceByUserIdResultFromDict(data map[string]interface{}) DirectEnhanceByUserIdResult {
	return DirectEnhanceByUserIdResult{
		Item:                      NewRateModelFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AcquireExperience:         core.CastInt64(data["acquireExperience"]),
		BonusRate:                 core.CastFloat32(data["bonusRate"]),
	}
}

func (p DirectEnhanceByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"acquireExperience":         p.AcquireExperience,
		"bonusRate":                 p.BonusRate,
	}
}

func (p DirectEnhanceByUserIdResult) Pointer() *DirectEnhanceByUserIdResult {
	return &p
}

type DirectEnhanceByStampSheetResult struct {
	Item                      *RateModel `json:"item"`
	StampSheet                *string    `json:"stampSheet"`
	StampSheetEncryptionKeyId *string    `json:"stampSheetEncryptionKeyId"`
	AcquireExperience         *int64     `json:"acquireExperience"`
	BonusRate                 *float32   `json:"bonusRate"`
}

type DirectEnhanceByStampSheetAsyncResult struct {
	result *DirectEnhanceByStampSheetResult
	err    error
}

func NewDirectEnhanceByStampSheetResultFromJson(data string) DirectEnhanceByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDirectEnhanceByStampSheetResultFromDict(dict)
}

func NewDirectEnhanceByStampSheetResultFromDict(data map[string]interface{}) DirectEnhanceByStampSheetResult {
	return DirectEnhanceByStampSheetResult{
		Item:                      NewRateModelFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AcquireExperience:         core.CastInt64(data["acquireExperience"]),
		BonusRate:                 core.CastFloat32(data["bonusRate"]),
	}
}

func (p DirectEnhanceByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"acquireExperience":         p.AcquireExperience,
		"bonusRate":                 p.BonusRate,
	}
}

func (p DirectEnhanceByStampSheetResult) Pointer() *DirectEnhanceByStampSheetResult {
	return &p
}

type DescribeProgressesByUserIdResult struct {
	Items         []Progress `json:"items"`
	NextPageToken *string    `json:"nextPageToken"`
}

type DescribeProgressesByUserIdAsyncResult struct {
	result *DescribeProgressesByUserIdResult
	err    error
}

func NewDescribeProgressesByUserIdResultFromJson(data string) DescribeProgressesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeProgressesByUserIdResultFromDict(dict)
}

func NewDescribeProgressesByUserIdResultFromDict(data map[string]interface{}) DescribeProgressesByUserIdResult {
	return DescribeProgressesByUserIdResult{
		Items:         CastProgresses(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeProgressesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastProgressesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeProgressesByUserIdResult) Pointer() *DescribeProgressesByUserIdResult {
	return &p
}

type CreateProgressByUserIdResult struct {
	Item *Progress `json:"item"`
}

type CreateProgressByUserIdAsyncResult struct {
	result *CreateProgressByUserIdResult
	err    error
}

func NewCreateProgressByUserIdResultFromJson(data string) CreateProgressByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateProgressByUserIdResultFromDict(dict)
}

func NewCreateProgressByUserIdResultFromDict(data map[string]interface{}) CreateProgressByUserIdResult {
	return CreateProgressByUserIdResult{
		Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateProgressByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateProgressByUserIdResult) Pointer() *CreateProgressByUserIdResult {
	return &p
}

type GetProgressResult struct {
	Item      *Progress  `json:"item"`
	RateModel *RateModel `json:"rateModel"`
}

type GetProgressAsyncResult struct {
	result *GetProgressResult
	err    error
}

func NewGetProgressResultFromJson(data string) GetProgressResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProgressResultFromDict(dict)
}

func NewGetProgressResultFromDict(data map[string]interface{}) GetProgressResult {
	return GetProgressResult{
		Item:      NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
		RateModel: NewRateModelFromDict(core.CastMap(data["rateModel"])).Pointer(),
	}
}

func (p GetProgressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"rateModel": p.RateModel.ToDict(),
	}
}

func (p GetProgressResult) Pointer() *GetProgressResult {
	return &p
}

type GetProgressByUserIdResult struct {
	Item      *Progress  `json:"item"`
	RateModel *RateModel `json:"rateModel"`
}

type GetProgressByUserIdAsyncResult struct {
	result *GetProgressByUserIdResult
	err    error
}

func NewGetProgressByUserIdResultFromJson(data string) GetProgressByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProgressByUserIdResultFromDict(dict)
}

func NewGetProgressByUserIdResultFromDict(data map[string]interface{}) GetProgressByUserIdResult {
	return GetProgressByUserIdResult{
		Item:      NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
		RateModel: NewRateModelFromDict(core.CastMap(data["rateModel"])).Pointer(),
	}
}

func (p GetProgressByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"rateModel": p.RateModel.ToDict(),
	}
}

func (p GetProgressByUserIdResult) Pointer() *GetProgressByUserIdResult {
	return &p
}

type StartResult struct {
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type StartAsyncResult struct {
	result *StartResult
	err    error
}

func NewStartResultFromJson(data string) StartResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStartResultFromDict(dict)
}

func NewStartResultFromDict(data map[string]interface{}) StartResult {
	return StartResult{
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p StartResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p StartResult) Pointer() *StartResult {
	return &p
}

type StartByUserIdResult struct {
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type StartByUserIdAsyncResult struct {
	result *StartByUserIdResult
	err    error
}

func NewStartByUserIdResultFromJson(data string) StartByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStartByUserIdResultFromDict(dict)
}

func NewStartByUserIdResultFromDict(data map[string]interface{}) StartByUserIdResult {
	return StartByUserIdResult{
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p StartByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p StartByUserIdResult) Pointer() *StartByUserIdResult {
	return &p
}

type EndResult struct {
	Item                      *Progress `json:"item"`
	StampSheet                *string   `json:"stampSheet"`
	StampSheetEncryptionKeyId *string   `json:"stampSheetEncryptionKeyId"`
	AcquireExperience         *int64    `json:"acquireExperience"`
	BonusRate                 *float32  `json:"bonusRate"`
}

type EndAsyncResult struct {
	result *EndResult
	err    error
}

func NewEndResultFromJson(data string) EndResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEndResultFromDict(dict)
}

func NewEndResultFromDict(data map[string]interface{}) EndResult {
	return EndResult{
		Item:                      NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AcquireExperience:         core.CastInt64(data["acquireExperience"]),
		BonusRate:                 core.CastFloat32(data["bonusRate"]),
	}
}

func (p EndResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"acquireExperience":         p.AcquireExperience,
		"bonusRate":                 p.BonusRate,
	}
}

func (p EndResult) Pointer() *EndResult {
	return &p
}

type EndByUserIdResult struct {
	Item                      *Progress `json:"item"`
	StampSheet                *string   `json:"stampSheet"`
	StampSheetEncryptionKeyId *string   `json:"stampSheetEncryptionKeyId"`
	AcquireExperience         *int64    `json:"acquireExperience"`
	BonusRate                 *float32  `json:"bonusRate"`
}

type EndByUserIdAsyncResult struct {
	result *EndByUserIdResult
	err    error
}

func NewEndByUserIdResultFromJson(data string) EndByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEndByUserIdResultFromDict(dict)
}

func NewEndByUserIdResultFromDict(data map[string]interface{}) EndByUserIdResult {
	return EndByUserIdResult{
		Item:                      NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AcquireExperience:         core.CastInt64(data["acquireExperience"]),
		BonusRate:                 core.CastFloat32(data["bonusRate"]),
	}
}

func (p EndByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"acquireExperience":         p.AcquireExperience,
		"bonusRate":                 p.BonusRate,
	}
}

func (p EndByUserIdResult) Pointer() *EndByUserIdResult {
	return &p
}

type DeleteProgressResult struct {
	Item *Progress `json:"item"`
}

type DeleteProgressAsyncResult struct {
	result *DeleteProgressResult
	err    error
}

func NewDeleteProgressResultFromJson(data string) DeleteProgressResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteProgressResultFromDict(dict)
}

func NewDeleteProgressResultFromDict(data map[string]interface{}) DeleteProgressResult {
	return DeleteProgressResult{
		Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteProgressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteProgressResult) Pointer() *DeleteProgressResult {
	return &p
}

type DeleteProgressByUserIdResult struct {
	Item *Progress `json:"item"`
}

type DeleteProgressByUserIdAsyncResult struct {
	result *DeleteProgressByUserIdResult
	err    error
}

func NewDeleteProgressByUserIdResultFromJson(data string) DeleteProgressByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteProgressByUserIdResultFromDict(dict)
}

func NewDeleteProgressByUserIdResultFromDict(data map[string]interface{}) DeleteProgressByUserIdResult {
	return DeleteProgressByUserIdResult{
		Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteProgressByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteProgressByUserIdResult) Pointer() *DeleteProgressByUserIdResult {
	return &p
}

type CreateProgressByStampSheetResult struct {
	Item *Progress `json:"item"`
}

type CreateProgressByStampSheetAsyncResult struct {
	result *CreateProgressByStampSheetResult
	err    error
}

func NewCreateProgressByStampSheetResultFromJson(data string) CreateProgressByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateProgressByStampSheetResultFromDict(dict)
}

func NewCreateProgressByStampSheetResultFromDict(data map[string]interface{}) CreateProgressByStampSheetResult {
	return CreateProgressByStampSheetResult{
		Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateProgressByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateProgressByStampSheetResult) Pointer() *CreateProgressByStampSheetResult {
	return &p
}

type DeleteProgressByStampTaskResult struct {
	Item            *Progress `json:"item"`
	NewContextStack *string   `json:"newContextStack"`
}

type DeleteProgressByStampTaskAsyncResult struct {
	result *DeleteProgressByStampTaskResult
	err    error
}

func NewDeleteProgressByStampTaskResultFromJson(data string) DeleteProgressByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteProgressByStampTaskResultFromDict(dict)
}

func NewDeleteProgressByStampTaskResultFromDict(data map[string]interface{}) DeleteProgressByStampTaskResult {
	return DeleteProgressByStampTaskResult{
		Item:            NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p DeleteProgressByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":            p.Item.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p DeleteProgressByStampTaskResult) Pointer() *DeleteProgressByStampTaskResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentRateMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromJson(data string) ExportMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExportMasterResultFromDict(dict)
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
	return ExportMasterResult{
		Item: NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentRateMasterResult struct {
	Item *CurrentRateMaster `json:"item"`
}

type GetCurrentRateMasterAsyncResult struct {
	result *GetCurrentRateMasterResult
	err    error
}

func NewGetCurrentRateMasterResultFromJson(data string) GetCurrentRateMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentRateMasterResultFromDict(dict)
}

func NewGetCurrentRateMasterResultFromDict(data map[string]interface{}) GetCurrentRateMasterResult {
	return GetCurrentRateMasterResult{
		Item: NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentRateMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentRateMasterResult) Pointer() *GetCurrentRateMasterResult {
	return &p
}

type UpdateCurrentRateMasterResult struct {
	Item *CurrentRateMaster `json:"item"`
}

type UpdateCurrentRateMasterAsyncResult struct {
	result *UpdateCurrentRateMasterResult
	err    error
}

func NewUpdateCurrentRateMasterResultFromJson(data string) UpdateCurrentRateMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRateMasterResultFromDict(dict)
}

func NewUpdateCurrentRateMasterResultFromDict(data map[string]interface{}) UpdateCurrentRateMasterResult {
	return UpdateCurrentRateMasterResult{
		Item: NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentRateMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentRateMasterResult) Pointer() *UpdateCurrentRateMasterResult {
	return &p
}

type UpdateCurrentRateMasterFromGitHubResult struct {
	Item *CurrentRateMaster `json:"item"`
}

type UpdateCurrentRateMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRateMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentRateMasterFromGitHubResultFromJson(data string) UpdateCurrentRateMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRateMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentRateMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentRateMasterFromGitHubResult {
	return UpdateCurrentRateMasterFromGitHubResult{
		Item: NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentRateMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentRateMasterFromGitHubResult) Pointer() *UpdateCurrentRateMasterFromGitHubResult {
	return &p
}
