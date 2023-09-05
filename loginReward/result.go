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

package loginReward

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

type DescribeBonusModelMastersResult struct {
	Items         []BonusModelMaster `json:"items"`
	NextPageToken *string            `json:"nextPageToken"`
}

type DescribeBonusModelMastersAsyncResult struct {
	result *DescribeBonusModelMastersResult
	err    error
}

func NewDescribeBonusModelMastersResultFromJson(data string) DescribeBonusModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBonusModelMastersResultFromDict(dict)
}

func NewDescribeBonusModelMastersResultFromDict(data map[string]interface{}) DescribeBonusModelMastersResult {
	return DescribeBonusModelMastersResult{
		Items:         CastBonusModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeBonusModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBonusModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeBonusModelMastersResult) Pointer() *DescribeBonusModelMastersResult {
	return &p
}

type CreateBonusModelMasterResult struct {
	Item *BonusModelMaster `json:"item"`
}

type CreateBonusModelMasterAsyncResult struct {
	result *CreateBonusModelMasterResult
	err    error
}

func NewCreateBonusModelMasterResultFromJson(data string) CreateBonusModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBonusModelMasterResultFromDict(dict)
}

func NewCreateBonusModelMasterResultFromDict(data map[string]interface{}) CreateBonusModelMasterResult {
	return CreateBonusModelMasterResult{
		Item: NewBonusModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateBonusModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateBonusModelMasterResult) Pointer() *CreateBonusModelMasterResult {
	return &p
}

type GetBonusModelMasterResult struct {
	Item *BonusModelMaster `json:"item"`
}

type GetBonusModelMasterAsyncResult struct {
	result *GetBonusModelMasterResult
	err    error
}

func NewGetBonusModelMasterResultFromJson(data string) GetBonusModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBonusModelMasterResultFromDict(dict)
}

func NewGetBonusModelMasterResultFromDict(data map[string]interface{}) GetBonusModelMasterResult {
	return GetBonusModelMasterResult{
		Item: NewBonusModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetBonusModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetBonusModelMasterResult) Pointer() *GetBonusModelMasterResult {
	return &p
}

type UpdateBonusModelMasterResult struct {
	Item *BonusModelMaster `json:"item"`
}

type UpdateBonusModelMasterAsyncResult struct {
	result *UpdateBonusModelMasterResult
	err    error
}

func NewUpdateBonusModelMasterResultFromJson(data string) UpdateBonusModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBonusModelMasterResultFromDict(dict)
}

func NewUpdateBonusModelMasterResultFromDict(data map[string]interface{}) UpdateBonusModelMasterResult {
	return UpdateBonusModelMasterResult{
		Item: NewBonusModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateBonusModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateBonusModelMasterResult) Pointer() *UpdateBonusModelMasterResult {
	return &p
}

type DeleteBonusModelMasterResult struct {
	Item *BonusModelMaster `json:"item"`
}

type DeleteBonusModelMasterAsyncResult struct {
	result *DeleteBonusModelMasterResult
	err    error
}

func NewDeleteBonusModelMasterResultFromJson(data string) DeleteBonusModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBonusModelMasterResultFromDict(dict)
}

func NewDeleteBonusModelMasterResultFromDict(data map[string]interface{}) DeleteBonusModelMasterResult {
	return DeleteBonusModelMasterResult{
		Item: NewBonusModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteBonusModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteBonusModelMasterResult) Pointer() *DeleteBonusModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentBonusMaster `json:"item"`
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
		Item: NewCurrentBonusMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentBonusMasterResult struct {
	Item *CurrentBonusMaster `json:"item"`
}

type GetCurrentBonusMasterAsyncResult struct {
	result *GetCurrentBonusMasterResult
	err    error
}

func NewGetCurrentBonusMasterResultFromJson(data string) GetCurrentBonusMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentBonusMasterResultFromDict(dict)
}

func NewGetCurrentBonusMasterResultFromDict(data map[string]interface{}) GetCurrentBonusMasterResult {
	return GetCurrentBonusMasterResult{
		Item: NewCurrentBonusMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentBonusMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentBonusMasterResult) Pointer() *GetCurrentBonusMasterResult {
	return &p
}

type UpdateCurrentBonusMasterResult struct {
	Item *CurrentBonusMaster `json:"item"`
}

type UpdateCurrentBonusMasterAsyncResult struct {
	result *UpdateCurrentBonusMasterResult
	err    error
}

func NewUpdateCurrentBonusMasterResultFromJson(data string) UpdateCurrentBonusMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentBonusMasterResultFromDict(dict)
}

func NewUpdateCurrentBonusMasterResultFromDict(data map[string]interface{}) UpdateCurrentBonusMasterResult {
	return UpdateCurrentBonusMasterResult{
		Item: NewCurrentBonusMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentBonusMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentBonusMasterResult) Pointer() *UpdateCurrentBonusMasterResult {
	return &p
}

type UpdateCurrentBonusMasterFromGitHubResult struct {
	Item *CurrentBonusMaster `json:"item"`
}

type UpdateCurrentBonusMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentBonusMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentBonusMasterFromGitHubResultFromJson(data string) UpdateCurrentBonusMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentBonusMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentBonusMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentBonusMasterFromGitHubResult {
	return UpdateCurrentBonusMasterFromGitHubResult{
		Item: NewCurrentBonusMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentBonusMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentBonusMasterFromGitHubResult) Pointer() *UpdateCurrentBonusMasterFromGitHubResult {
	return &p
}

type DescribeBonusModelsResult struct {
	Items []BonusModel `json:"items"`
}

type DescribeBonusModelsAsyncResult struct {
	result *DescribeBonusModelsResult
	err    error
}

func NewDescribeBonusModelsResultFromJson(data string) DescribeBonusModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBonusModelsResultFromDict(dict)
}

func NewDescribeBonusModelsResultFromDict(data map[string]interface{}) DescribeBonusModelsResult {
	return DescribeBonusModelsResult{
		Items: CastBonusModels(core.CastArray(data["items"])),
	}
}

func (p DescribeBonusModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBonusModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeBonusModelsResult) Pointer() *DescribeBonusModelsResult {
	return &p
}

type GetBonusModelResult struct {
	Item *BonusModel `json:"item"`
}

type GetBonusModelAsyncResult struct {
	result *GetBonusModelResult
	err    error
}

func NewGetBonusModelResultFromJson(data string) GetBonusModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBonusModelResultFromDict(dict)
}

func NewGetBonusModelResultFromDict(data map[string]interface{}) GetBonusModelResult {
	return GetBonusModelResult{
		Item: NewBonusModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetBonusModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetBonusModelResult) Pointer() *GetBonusModelResult {
	return &p
}

type ReceiveResult struct {
	Item                      *ReceiveStatus `json:"item"`
	BonusModel                *BonusModel    `json:"bonusModel"`
	TransactionId             *string        `json:"transactionId"`
	StampSheet                *string        `json:"stampSheet"`
	StampSheetEncryptionKeyId *string        `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool          `json:"autoRunStampSheet"`
}

type ReceiveAsyncResult struct {
	result *ReceiveResult
	err    error
}

func NewReceiveResultFromJson(data string) ReceiveResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveResultFromDict(dict)
}

func NewReceiveResultFromDict(data map[string]interface{}) ReceiveResult {
	return ReceiveResult{
		Item:                      NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel:                NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p ReceiveResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"bonusModel":                p.BonusModel.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ReceiveResult) Pointer() *ReceiveResult {
	return &p
}

type ReceiveByUserIdResult struct {
	Item                      *ReceiveStatus `json:"item"`
	BonusModel                *BonusModel    `json:"bonusModel"`
	TransactionId             *string        `json:"transactionId"`
	StampSheet                *string        `json:"stampSheet"`
	StampSheetEncryptionKeyId *string        `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool          `json:"autoRunStampSheet"`
}

type ReceiveByUserIdAsyncResult struct {
	result *ReceiveByUserIdResult
	err    error
}

func NewReceiveByUserIdResultFromJson(data string) ReceiveByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveByUserIdResultFromDict(dict)
}

func NewReceiveByUserIdResultFromDict(data map[string]interface{}) ReceiveByUserIdResult {
	return ReceiveByUserIdResult{
		Item:                      NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel:                NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p ReceiveByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"bonusModel":                p.BonusModel.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ReceiveByUserIdResult) Pointer() *ReceiveByUserIdResult {
	return &p
}

type MissedReceiveResult struct {
	Item                      *ReceiveStatus `json:"item"`
	BonusModel                *BonusModel    `json:"bonusModel"`
	TransactionId             *string        `json:"transactionId"`
	StampSheet                *string        `json:"stampSheet"`
	StampSheetEncryptionKeyId *string        `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool          `json:"autoRunStampSheet"`
}

type MissedReceiveAsyncResult struct {
	result *MissedReceiveResult
	err    error
}

func NewMissedReceiveResultFromJson(data string) MissedReceiveResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissedReceiveResultFromDict(dict)
}

func NewMissedReceiveResultFromDict(data map[string]interface{}) MissedReceiveResult {
	return MissedReceiveResult{
		Item:                      NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel:                NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p MissedReceiveResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"bonusModel":                p.BonusModel.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p MissedReceiveResult) Pointer() *MissedReceiveResult {
	return &p
}

type MissedReceiveByUserIdResult struct {
	Item                      *ReceiveStatus `json:"item"`
	BonusModel                *BonusModel    `json:"bonusModel"`
	TransactionId             *string        `json:"transactionId"`
	StampSheet                *string        `json:"stampSheet"`
	StampSheetEncryptionKeyId *string        `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool          `json:"autoRunStampSheet"`
}

type MissedReceiveByUserIdAsyncResult struct {
	result *MissedReceiveByUserIdResult
	err    error
}

func NewMissedReceiveByUserIdResultFromJson(data string) MissedReceiveByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissedReceiveByUserIdResultFromDict(dict)
}

func NewMissedReceiveByUserIdResultFromDict(data map[string]interface{}) MissedReceiveByUserIdResult {
	return MissedReceiveByUserIdResult{
		Item:                      NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel:                NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p MissedReceiveByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"bonusModel":                p.BonusModel.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p MissedReceiveByUserIdResult) Pointer() *MissedReceiveByUserIdResult {
	return &p
}

type DescribeReceiveStatusesResult struct {
	Items         []ReceiveStatus `json:"items"`
	NextPageToken *string         `json:"nextPageToken"`
}

type DescribeReceiveStatusesAsyncResult struct {
	result *DescribeReceiveStatusesResult
	err    error
}

func NewDescribeReceiveStatusesResultFromJson(data string) DescribeReceiveStatusesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiveStatusesResultFromDict(dict)
}

func NewDescribeReceiveStatusesResultFromDict(data map[string]interface{}) DescribeReceiveStatusesResult {
	return DescribeReceiveStatusesResult{
		Items:         CastReceiveStatuses(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeReceiveStatusesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastReceiveStatusesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeReceiveStatusesResult) Pointer() *DescribeReceiveStatusesResult {
	return &p
}

type DescribeReceiveStatusesByUserIdResult struct {
	Items         []ReceiveStatus `json:"items"`
	NextPageToken *string         `json:"nextPageToken"`
}

type DescribeReceiveStatusesByUserIdAsyncResult struct {
	result *DescribeReceiveStatusesByUserIdResult
	err    error
}

func NewDescribeReceiveStatusesByUserIdResultFromJson(data string) DescribeReceiveStatusesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiveStatusesByUserIdResultFromDict(dict)
}

func NewDescribeReceiveStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeReceiveStatusesByUserIdResult {
	return DescribeReceiveStatusesByUserIdResult{
		Items:         CastReceiveStatuses(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeReceiveStatusesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastReceiveStatusesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeReceiveStatusesByUserIdResult) Pointer() *DescribeReceiveStatusesByUserIdResult {
	return &p
}

type GetReceiveStatusResult struct {
	Item       *ReceiveStatus `json:"item"`
	BonusModel *BonusModel    `json:"bonusModel"`
}

type GetReceiveStatusAsyncResult struct {
	result *GetReceiveStatusResult
	err    error
}

func NewGetReceiveStatusResultFromJson(data string) GetReceiveStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceiveStatusResultFromDict(dict)
}

func NewGetReceiveStatusResultFromDict(data map[string]interface{}) GetReceiveStatusResult {
	return GetReceiveStatusResult{
		Item:       NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel: NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
	}
}

func (p GetReceiveStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":       p.Item.ToDict(),
		"bonusModel": p.BonusModel.ToDict(),
	}
}

func (p GetReceiveStatusResult) Pointer() *GetReceiveStatusResult {
	return &p
}

type GetReceiveStatusByUserIdResult struct {
	Item       *ReceiveStatus `json:"item"`
	BonusModel *BonusModel    `json:"bonusModel"`
}

type GetReceiveStatusByUserIdAsyncResult struct {
	result *GetReceiveStatusByUserIdResult
	err    error
}

func NewGetReceiveStatusByUserIdResultFromJson(data string) GetReceiveStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceiveStatusByUserIdResultFromDict(dict)
}

func NewGetReceiveStatusByUserIdResultFromDict(data map[string]interface{}) GetReceiveStatusByUserIdResult {
	return GetReceiveStatusByUserIdResult{
		Item:       NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel: NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
	}
}

func (p GetReceiveStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":       p.Item.ToDict(),
		"bonusModel": p.BonusModel.ToDict(),
	}
}

func (p GetReceiveStatusByUserIdResult) Pointer() *GetReceiveStatusByUserIdResult {
	return &p
}

type DeleteReceiveStatusByUserIdResult struct {
	Item       *ReceiveStatus `json:"item"`
	BonusModel *BonusModel    `json:"bonusModel"`
}

type DeleteReceiveStatusByUserIdAsyncResult struct {
	result *DeleteReceiveStatusByUserIdResult
	err    error
}

func NewDeleteReceiveStatusByUserIdResultFromJson(data string) DeleteReceiveStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReceiveStatusByUserIdResultFromDict(dict)
}

func NewDeleteReceiveStatusByUserIdResultFromDict(data map[string]interface{}) DeleteReceiveStatusByUserIdResult {
	return DeleteReceiveStatusByUserIdResult{
		Item:       NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel: NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
	}
}

func (p DeleteReceiveStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":       p.Item.ToDict(),
		"bonusModel": p.BonusModel.ToDict(),
	}
}

func (p DeleteReceiveStatusByUserIdResult) Pointer() *DeleteReceiveStatusByUserIdResult {
	return &p
}

type DeleteReceiveStatusByStampSheetResult struct {
	Item       *ReceiveStatus `json:"item"`
	BonusModel *BonusModel    `json:"bonusModel"`
}

type DeleteReceiveStatusByStampSheetAsyncResult struct {
	result *DeleteReceiveStatusByStampSheetResult
	err    error
}

func NewDeleteReceiveStatusByStampSheetResultFromJson(data string) DeleteReceiveStatusByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReceiveStatusByStampSheetResultFromDict(dict)
}

func NewDeleteReceiveStatusByStampSheetResultFromDict(data map[string]interface{}) DeleteReceiveStatusByStampSheetResult {
	return DeleteReceiveStatusByStampSheetResult{
		Item:       NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel: NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
	}
}

func (p DeleteReceiveStatusByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":       p.Item.ToDict(),
		"bonusModel": p.BonusModel.ToDict(),
	}
}

func (p DeleteReceiveStatusByStampSheetResult) Pointer() *DeleteReceiveStatusByStampSheetResult {
	return &p
}

type MarkReceivedResult struct {
	Item       *ReceiveStatus `json:"item"`
	BonusModel *BonusModel    `json:"bonusModel"`
}

type MarkReceivedAsyncResult struct {
	result *MarkReceivedResult
	err    error
}

func NewMarkReceivedResultFromJson(data string) MarkReceivedResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReceivedResultFromDict(dict)
}

func NewMarkReceivedResultFromDict(data map[string]interface{}) MarkReceivedResult {
	return MarkReceivedResult{
		Item:       NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel: NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
	}
}

func (p MarkReceivedResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":       p.Item.ToDict(),
		"bonusModel": p.BonusModel.ToDict(),
	}
}

func (p MarkReceivedResult) Pointer() *MarkReceivedResult {
	return &p
}

type MarkReceivedByUserIdResult struct {
	Item       *ReceiveStatus `json:"item"`
	BonusModel *BonusModel    `json:"bonusModel"`
}

type MarkReceivedByUserIdAsyncResult struct {
	result *MarkReceivedByUserIdResult
	err    error
}

func NewMarkReceivedByUserIdResultFromJson(data string) MarkReceivedByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReceivedByUserIdResultFromDict(dict)
}

func NewMarkReceivedByUserIdResultFromDict(data map[string]interface{}) MarkReceivedByUserIdResult {
	return MarkReceivedByUserIdResult{
		Item:       NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel: NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
	}
}

func (p MarkReceivedByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":       p.Item.ToDict(),
		"bonusModel": p.BonusModel.ToDict(),
	}
}

func (p MarkReceivedByUserIdResult) Pointer() *MarkReceivedByUserIdResult {
	return &p
}

type UnmarkReceivedByUserIdResult struct {
	Item       *ReceiveStatus `json:"item"`
	BonusModel *BonusModel    `json:"bonusModel"`
}

type UnmarkReceivedByUserIdAsyncResult struct {
	result *UnmarkReceivedByUserIdResult
	err    error
}

func NewUnmarkReceivedByUserIdResultFromJson(data string) UnmarkReceivedByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnmarkReceivedByUserIdResultFromDict(dict)
}

func NewUnmarkReceivedByUserIdResultFromDict(data map[string]interface{}) UnmarkReceivedByUserIdResult {
	return UnmarkReceivedByUserIdResult{
		Item:       NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel: NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
	}
}

func (p UnmarkReceivedByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":       p.Item.ToDict(),
		"bonusModel": p.BonusModel.ToDict(),
	}
}

func (p UnmarkReceivedByUserIdResult) Pointer() *UnmarkReceivedByUserIdResult {
	return &p
}

type MarkReceivedByStampTaskResult struct {
	Item            *ReceiveStatus `json:"item"`
	BonusModel      *BonusModel    `json:"bonusModel"`
	NewContextStack *string        `json:"newContextStack"`
}

type MarkReceivedByStampTaskAsyncResult struct {
	result *MarkReceivedByStampTaskResult
	err    error
}

func NewMarkReceivedByStampTaskResultFromJson(data string) MarkReceivedByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReceivedByStampTaskResultFromDict(dict)
}

func NewMarkReceivedByStampTaskResultFromDict(data map[string]interface{}) MarkReceivedByStampTaskResult {
	return MarkReceivedByStampTaskResult{
		Item:            NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel:      NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p MarkReceivedByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":            p.Item.ToDict(),
		"bonusModel":      p.BonusModel.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p MarkReceivedByStampTaskResult) Pointer() *MarkReceivedByStampTaskResult {
	return &p
}

type UnmarkReceivedByStampSheetResult struct {
	Item       *ReceiveStatus `json:"item"`
	BonusModel *BonusModel    `json:"bonusModel"`
}

type UnmarkReceivedByStampSheetAsyncResult struct {
	result *UnmarkReceivedByStampSheetResult
	err    error
}

func NewUnmarkReceivedByStampSheetResultFromJson(data string) UnmarkReceivedByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnmarkReceivedByStampSheetResultFromDict(dict)
}

func NewUnmarkReceivedByStampSheetResultFromDict(data map[string]interface{}) UnmarkReceivedByStampSheetResult {
	return UnmarkReceivedByStampSheetResult{
		Item:       NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer(),
		BonusModel: NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer(),
	}
}

func (p UnmarkReceivedByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":       p.Item.ToDict(),
		"bonusModel": p.BonusModel.ToDict(),
	}
}

func (p UnmarkReceivedByStampSheetResult) Pointer() *UnmarkReceivedByStampSheetResult {
	return &p
}
