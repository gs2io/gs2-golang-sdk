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
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeCompletesResult struct {
	Items         []Complete `json:"items"`
	NextPageToken *string    `json:"nextPageToken"`
}

type DescribeCompletesAsyncResult struct {
	result *DescribeCompletesResult
	err    error
}

func NewDescribeCompletesResultFromJson(data string) DescribeCompletesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCompletesResultFromDict(dict)
}

func NewDescribeCompletesResultFromDict(data map[string]interface{}) DescribeCompletesResult {
	return DescribeCompletesResult{
		Items:         CastCompletes(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeCompletesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCompletesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCompletesResult) Pointer() *DescribeCompletesResult {
	return &p
}

type DescribeCompletesByUserIdResult struct {
	Items         []Complete `json:"items"`
	NextPageToken *string    `json:"nextPageToken"`
}

type DescribeCompletesByUserIdAsyncResult struct {
	result *DescribeCompletesByUserIdResult
	err    error
}

func NewDescribeCompletesByUserIdResultFromJson(data string) DescribeCompletesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCompletesByUserIdResultFromDict(dict)
}

func NewDescribeCompletesByUserIdResultFromDict(data map[string]interface{}) DescribeCompletesByUserIdResult {
	return DescribeCompletesByUserIdResult{
		Items:         CastCompletes(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeCompletesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCompletesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCompletesByUserIdResult) Pointer() *DescribeCompletesByUserIdResult {
	return &p
}

type CompleteResult struct {
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type CompleteAsyncResult struct {
	result *CompleteResult
	err    error
}

func NewCompleteResultFromJson(data string) CompleteResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCompleteResultFromDict(dict)
}

func NewCompleteResultFromDict(data map[string]interface{}) CompleteResult {
	return CompleteResult{
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p CompleteResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p CompleteResult) Pointer() *CompleteResult {
	return &p
}

type CompleteByUserIdResult struct {
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type CompleteByUserIdAsyncResult struct {
	result *CompleteByUserIdResult
	err    error
}

func NewCompleteByUserIdResultFromJson(data string) CompleteByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCompleteByUserIdResultFromDict(dict)
}

func NewCompleteByUserIdResultFromDict(data map[string]interface{}) CompleteByUserIdResult {
	return CompleteByUserIdResult{
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p CompleteByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p CompleteByUserIdResult) Pointer() *CompleteByUserIdResult {
	return &p
}

type ReceiveByUserIdResult struct {
	Item *Complete `json:"item"`
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
		Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ReceiveByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p ReceiveByUserIdResult) Pointer() *ReceiveByUserIdResult {
	return &p
}

type RevertReceiveByUserIdResult struct {
	Item *Complete `json:"item"`
}

type RevertReceiveByUserIdAsyncResult struct {
	result *RevertReceiveByUserIdResult
	err    error
}

func NewRevertReceiveByUserIdResultFromJson(data string) RevertReceiveByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertReceiveByUserIdResultFromDict(dict)
}

func NewRevertReceiveByUserIdResultFromDict(data map[string]interface{}) RevertReceiveByUserIdResult {
	return RevertReceiveByUserIdResult{
		Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p RevertReceiveByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p RevertReceiveByUserIdResult) Pointer() *RevertReceiveByUserIdResult {
	return &p
}

type GetCompleteResult struct {
	Item *Complete `json:"item"`
}

type GetCompleteAsyncResult struct {
	result *GetCompleteResult
	err    error
}

func NewGetCompleteResultFromJson(data string) GetCompleteResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCompleteResultFromDict(dict)
}

func NewGetCompleteResultFromDict(data map[string]interface{}) GetCompleteResult {
	return GetCompleteResult{
		Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCompleteResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCompleteResult) Pointer() *GetCompleteResult {
	return &p
}

type GetCompleteByUserIdResult struct {
	Item *Complete `json:"item"`
}

type GetCompleteByUserIdAsyncResult struct {
	result *GetCompleteByUserIdResult
	err    error
}

func NewGetCompleteByUserIdResultFromJson(data string) GetCompleteByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCompleteByUserIdResultFromDict(dict)
}

func NewGetCompleteByUserIdResultFromDict(data map[string]interface{}) GetCompleteByUserIdResult {
	return GetCompleteByUserIdResult{
		Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCompleteByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCompleteByUserIdResult) Pointer() *GetCompleteByUserIdResult {
	return &p
}

type DeleteCompleteByUserIdResult struct {
	Item *Complete `json:"item"`
}

type DeleteCompleteByUserIdAsyncResult struct {
	result *DeleteCompleteByUserIdResult
	err    error
}

func NewDeleteCompleteByUserIdResultFromJson(data string) DeleteCompleteByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCompleteByUserIdResultFromDict(dict)
}

func NewDeleteCompleteByUserIdResultFromDict(data map[string]interface{}) DeleteCompleteByUserIdResult {
	return DeleteCompleteByUserIdResult{
		Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteCompleteByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteCompleteByUserIdResult) Pointer() *DeleteCompleteByUserIdResult {
	return &p
}

type ReceiveByStampTaskResult struct {
	Item            *Complete `json:"item"`
	NewContextStack *string   `json:"newContextStack"`
}

type ReceiveByStampTaskAsyncResult struct {
	result *ReceiveByStampTaskResult
	err    error
}

func NewReceiveByStampTaskResultFromJson(data string) ReceiveByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveByStampTaskResultFromDict(dict)
}

func NewReceiveByStampTaskResultFromDict(data map[string]interface{}) ReceiveByStampTaskResult {
	return ReceiveByStampTaskResult{
		Item:            NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p ReceiveByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":            p.Item.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p ReceiveByStampTaskResult) Pointer() *ReceiveByStampTaskResult {
	return &p
}

type RevertReceiveByStampSheetResult struct {
	Item *Complete `json:"item"`
}

type RevertReceiveByStampSheetAsyncResult struct {
	result *RevertReceiveByStampSheetResult
	err    error
}

func NewRevertReceiveByStampSheetResultFromJson(data string) RevertReceiveByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertReceiveByStampSheetResultFromDict(dict)
}

func NewRevertReceiveByStampSheetResultFromDict(data map[string]interface{}) RevertReceiveByStampSheetResult {
	return RevertReceiveByStampSheetResult{
		Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p RevertReceiveByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p RevertReceiveByStampSheetResult) Pointer() *RevertReceiveByStampSheetResult {
	return &p
}

type DescribeCounterModelMastersResult struct {
	Items         []CounterModelMaster `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
}

type DescribeCounterModelMastersAsyncResult struct {
	result *DescribeCounterModelMastersResult
	err    error
}

func NewDescribeCounterModelMastersResultFromJson(data string) DescribeCounterModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCounterModelMastersResultFromDict(dict)
}

func NewDescribeCounterModelMastersResultFromDict(data map[string]interface{}) DescribeCounterModelMastersResult {
	return DescribeCounterModelMastersResult{
		Items:         CastCounterModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeCounterModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCounterModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCounterModelMastersResult) Pointer() *DescribeCounterModelMastersResult {
	return &p
}

type CreateCounterModelMasterResult struct {
	Item *CounterModelMaster `json:"item"`
}

type CreateCounterModelMasterAsyncResult struct {
	result *CreateCounterModelMasterResult
	err    error
}

func NewCreateCounterModelMasterResultFromJson(data string) CreateCounterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateCounterModelMasterResultFromDict(dict)
}

func NewCreateCounterModelMasterResultFromDict(data map[string]interface{}) CreateCounterModelMasterResult {
	return CreateCounterModelMasterResult{
		Item: NewCounterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateCounterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateCounterModelMasterResult) Pointer() *CreateCounterModelMasterResult {
	return &p
}

type GetCounterModelMasterResult struct {
	Item *CounterModelMaster `json:"item"`
}

type GetCounterModelMasterAsyncResult struct {
	result *GetCounterModelMasterResult
	err    error
}

func NewGetCounterModelMasterResultFromJson(data string) GetCounterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterModelMasterResultFromDict(dict)
}

func NewGetCounterModelMasterResultFromDict(data map[string]interface{}) GetCounterModelMasterResult {
	return GetCounterModelMasterResult{
		Item: NewCounterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCounterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCounterModelMasterResult) Pointer() *GetCounterModelMasterResult {
	return &p
}

type UpdateCounterModelMasterResult struct {
	Item *CounterModelMaster `json:"item"`
}

type UpdateCounterModelMasterAsyncResult struct {
	result *UpdateCounterModelMasterResult
	err    error
}

func NewUpdateCounterModelMasterResultFromJson(data string) UpdateCounterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCounterModelMasterResultFromDict(dict)
}

func NewUpdateCounterModelMasterResultFromDict(data map[string]interface{}) UpdateCounterModelMasterResult {
	return UpdateCounterModelMasterResult{
		Item: NewCounterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCounterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCounterModelMasterResult) Pointer() *UpdateCounterModelMasterResult {
	return &p
}

type DeleteCounterModelMasterResult struct {
	Item *CounterModelMaster `json:"item"`
}

type DeleteCounterModelMasterAsyncResult struct {
	result *DeleteCounterModelMasterResult
	err    error
}

func NewDeleteCounterModelMasterResultFromJson(data string) DeleteCounterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCounterModelMasterResultFromDict(dict)
}

func NewDeleteCounterModelMasterResultFromDict(data map[string]interface{}) DeleteCounterModelMasterResult {
	return DeleteCounterModelMasterResult{
		Item: NewCounterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteCounterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteCounterModelMasterResult) Pointer() *DeleteCounterModelMasterResult {
	return &p
}

type DescribeMissionGroupModelMastersResult struct {
	Items         []MissionGroupModelMaster `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
}

type DescribeMissionGroupModelMastersAsyncResult struct {
	result *DescribeMissionGroupModelMastersResult
	err    error
}

func NewDescribeMissionGroupModelMastersResultFromJson(data string) DescribeMissionGroupModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionGroupModelMastersResultFromDict(dict)
}

func NewDescribeMissionGroupModelMastersResultFromDict(data map[string]interface{}) DescribeMissionGroupModelMastersResult {
	return DescribeMissionGroupModelMastersResult{
		Items:         CastMissionGroupModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeMissionGroupModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMissionGroupModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeMissionGroupModelMastersResult) Pointer() *DescribeMissionGroupModelMastersResult {
	return &p
}

type CreateMissionGroupModelMasterResult struct {
	Item *MissionGroupModelMaster `json:"item"`
}

type CreateMissionGroupModelMasterAsyncResult struct {
	result *CreateMissionGroupModelMasterResult
	err    error
}

func NewCreateMissionGroupModelMasterResultFromJson(data string) CreateMissionGroupModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMissionGroupModelMasterResultFromDict(dict)
}

func NewCreateMissionGroupModelMasterResultFromDict(data map[string]interface{}) CreateMissionGroupModelMasterResult {
	return CreateMissionGroupModelMasterResult{
		Item: NewMissionGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateMissionGroupModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateMissionGroupModelMasterResult) Pointer() *CreateMissionGroupModelMasterResult {
	return &p
}

type GetMissionGroupModelMasterResult struct {
	Item *MissionGroupModelMaster `json:"item"`
}

type GetMissionGroupModelMasterAsyncResult struct {
	result *GetMissionGroupModelMasterResult
	err    error
}

func NewGetMissionGroupModelMasterResultFromJson(data string) GetMissionGroupModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionGroupModelMasterResultFromDict(dict)
}

func NewGetMissionGroupModelMasterResultFromDict(data map[string]interface{}) GetMissionGroupModelMasterResult {
	return GetMissionGroupModelMasterResult{
		Item: NewMissionGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMissionGroupModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMissionGroupModelMasterResult) Pointer() *GetMissionGroupModelMasterResult {
	return &p
}

type UpdateMissionGroupModelMasterResult struct {
	Item *MissionGroupModelMaster `json:"item"`
}

type UpdateMissionGroupModelMasterAsyncResult struct {
	result *UpdateMissionGroupModelMasterResult
	err    error
}

func NewUpdateMissionGroupModelMasterResultFromJson(data string) UpdateMissionGroupModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMissionGroupModelMasterResultFromDict(dict)
}

func NewUpdateMissionGroupModelMasterResultFromDict(data map[string]interface{}) UpdateMissionGroupModelMasterResult {
	return UpdateMissionGroupModelMasterResult{
		Item: NewMissionGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateMissionGroupModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateMissionGroupModelMasterResult) Pointer() *UpdateMissionGroupModelMasterResult {
	return &p
}

type DeleteMissionGroupModelMasterResult struct {
	Item *MissionGroupModelMaster `json:"item"`
}

type DeleteMissionGroupModelMasterAsyncResult struct {
	result *DeleteMissionGroupModelMasterResult
	err    error
}

func NewDeleteMissionGroupModelMasterResultFromJson(data string) DeleteMissionGroupModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMissionGroupModelMasterResultFromDict(dict)
}

func NewDeleteMissionGroupModelMasterResultFromDict(data map[string]interface{}) DeleteMissionGroupModelMasterResult {
	return DeleteMissionGroupModelMasterResult{
		Item: NewMissionGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteMissionGroupModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteMissionGroupModelMasterResult) Pointer() *DeleteMissionGroupModelMasterResult {
	return &p
}

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

type DumpUserDataByUserIdResult struct {
}

type DumpUserDataByUserIdAsyncResult struct {
	result *DumpUserDataByUserIdResult
	err    error
}

func NewDumpUserDataByUserIdResultFromJson(data string) DumpUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdResultFromDict(dict)
}

func NewDumpUserDataByUserIdResultFromDict(data map[string]interface{}) DumpUserDataByUserIdResult {
	return DumpUserDataByUserIdResult{}
}

func (p DumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DumpUserDataByUserIdResult) Pointer() *DumpUserDataByUserIdResult {
	return &p
}

type CheckDumpUserDataByUserIdResult struct {
	Url *string `json:"url"`
}

type CheckDumpUserDataByUserIdAsyncResult struct {
	result *CheckDumpUserDataByUserIdResult
	err    error
}

func NewCheckDumpUserDataByUserIdResultFromJson(data string) CheckDumpUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdResultFromDict(dict)
}

func NewCheckDumpUserDataByUserIdResultFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdResult {
	return CheckDumpUserDataByUserIdResult{
		Url: core.CastString(data["url"]),
	}
}

func (p CheckDumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
	}
}

func (p CheckDumpUserDataByUserIdResult) Pointer() *CheckDumpUserDataByUserIdResult {
	return &p
}

type CleanUserDataByUserIdResult struct {
}

type CleanUserDataByUserIdAsyncResult struct {
	result *CleanUserDataByUserIdResult
	err    error
}

func NewCleanUserDataByUserIdResultFromJson(data string) CleanUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdResultFromDict(dict)
}

func NewCleanUserDataByUserIdResultFromDict(data map[string]interface{}) CleanUserDataByUserIdResult {
	return CleanUserDataByUserIdResult{}
}

func (p CleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p CleanUserDataByUserIdResult) Pointer() *CleanUserDataByUserIdResult {
	return &p
}

type CheckCleanUserDataByUserIdResult struct {
}

type CheckCleanUserDataByUserIdAsyncResult struct {
	result *CheckCleanUserDataByUserIdResult
	err    error
}

func NewCheckCleanUserDataByUserIdResultFromJson(data string) CheckCleanUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdResultFromDict(dict)
}

func NewCheckCleanUserDataByUserIdResultFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdResult {
	return CheckCleanUserDataByUserIdResult{}
}

func (p CheckCleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p CheckCleanUserDataByUserIdResult) Pointer() *CheckCleanUserDataByUserIdResult {
	return &p
}

type PrepareImportUserDataByUserIdResult struct {
	UploadToken *string `json:"uploadToken"`
	UploadUrl   *string `json:"uploadUrl"`
}

type PrepareImportUserDataByUserIdAsyncResult struct {
	result *PrepareImportUserDataByUserIdResult
	err    error
}

func NewPrepareImportUserDataByUserIdResultFromJson(data string) PrepareImportUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdResultFromDict(dict)
}

func NewPrepareImportUserDataByUserIdResultFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdResult {
	return PrepareImportUserDataByUserIdResult{
		UploadToken: core.CastString(data["uploadToken"]),
		UploadUrl:   core.CastString(data["uploadUrl"]),
	}
}

func (p PrepareImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
	}
}

func (p PrepareImportUserDataByUserIdResult) Pointer() *PrepareImportUserDataByUserIdResult {
	return &p
}

type ImportUserDataByUserIdResult struct {
}

type ImportUserDataByUserIdAsyncResult struct {
	result *ImportUserDataByUserIdResult
	err    error
}

func NewImportUserDataByUserIdResultFromJson(data string) ImportUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdResultFromDict(dict)
}

func NewImportUserDataByUserIdResultFromDict(data map[string]interface{}) ImportUserDataByUserIdResult {
	return ImportUserDataByUserIdResult{}
}

func (p ImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p ImportUserDataByUserIdResult) Pointer() *ImportUserDataByUserIdResult {
	return &p
}

type CheckImportUserDataByUserIdResult struct {
	Url *string `json:"url"`
}

type CheckImportUserDataByUserIdAsyncResult struct {
	result *CheckImportUserDataByUserIdResult
	err    error
}

func NewCheckImportUserDataByUserIdResultFromJson(data string) CheckImportUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdResultFromDict(dict)
}

func NewCheckImportUserDataByUserIdResultFromDict(data map[string]interface{}) CheckImportUserDataByUserIdResult {
	return CheckImportUserDataByUserIdResult{
		Url: core.CastString(data["url"]),
	}
}

func (p CheckImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
	}
}

func (p CheckImportUserDataByUserIdResult) Pointer() *CheckImportUserDataByUserIdResult {
	return &p
}

type DescribeCountersResult struct {
	Items         []Counter `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeCountersAsyncResult struct {
	result *DescribeCountersResult
	err    error
}

func NewDescribeCountersResultFromJson(data string) DescribeCountersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCountersResultFromDict(dict)
}

func NewDescribeCountersResultFromDict(data map[string]interface{}) DescribeCountersResult {
	return DescribeCountersResult{
		Items:         CastCounters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeCountersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCountersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCountersResult) Pointer() *DescribeCountersResult {
	return &p
}

type DescribeCountersByUserIdResult struct {
	Items         []Counter `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeCountersByUserIdAsyncResult struct {
	result *DescribeCountersByUserIdResult
	err    error
}

func NewDescribeCountersByUserIdResultFromJson(data string) DescribeCountersByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCountersByUserIdResultFromDict(dict)
}

func NewDescribeCountersByUserIdResultFromDict(data map[string]interface{}) DescribeCountersByUserIdResult {
	return DescribeCountersByUserIdResult{
		Items:         CastCounters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeCountersByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCountersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCountersByUserIdResult) Pointer() *DescribeCountersByUserIdResult {
	return &p
}

type IncreaseCounterByUserIdResult struct {
	Item             *Counter   `json:"item"`
	ChangedCompletes []Complete `json:"changedCompletes"`
}

type IncreaseCounterByUserIdAsyncResult struct {
	result *IncreaseCounterByUserIdResult
	err    error
}

func NewIncreaseCounterByUserIdResultFromJson(data string) IncreaseCounterByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseCounterByUserIdResultFromDict(dict)
}

func NewIncreaseCounterByUserIdResultFromDict(data map[string]interface{}) IncreaseCounterByUserIdResult {
	return IncreaseCounterByUserIdResult{
		Item:             NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
		ChangedCompletes: CastCompletes(core.CastArray(data["changedCompletes"])),
	}
}

func (p IncreaseCounterByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
		"changedCompletes": CastCompletesFromDict(
			p.ChangedCompletes,
		),
	}
}

func (p IncreaseCounterByUserIdResult) Pointer() *IncreaseCounterByUserIdResult {
	return &p
}

type DecreaseCounterByUserIdResult struct {
	Item *Counter `json:"item"`
}

type DecreaseCounterByUserIdAsyncResult struct {
	result *DecreaseCounterByUserIdResult
	err    error
}

func NewDecreaseCounterByUserIdResultFromJson(data string) DecreaseCounterByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseCounterByUserIdResultFromDict(dict)
}

func NewDecreaseCounterByUserIdResultFromDict(data map[string]interface{}) DecreaseCounterByUserIdResult {
	return DecreaseCounterByUserIdResult{
		Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DecreaseCounterByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DecreaseCounterByUserIdResult) Pointer() *DecreaseCounterByUserIdResult {
	return &p
}

type GetCounterResult struct {
	Item *Counter `json:"item"`
}

type GetCounterAsyncResult struct {
	result *GetCounterResult
	err    error
}

func NewGetCounterResultFromJson(data string) GetCounterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterResultFromDict(dict)
}

func NewGetCounterResultFromDict(data map[string]interface{}) GetCounterResult {
	return GetCounterResult{
		Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCounterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCounterResult) Pointer() *GetCounterResult {
	return &p
}

type GetCounterByUserIdResult struct {
	Item *Counter `json:"item"`
}

type GetCounterByUserIdAsyncResult struct {
	result *GetCounterByUserIdResult
	err    error
}

func NewGetCounterByUserIdResultFromJson(data string) GetCounterByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterByUserIdResultFromDict(dict)
}

func NewGetCounterByUserIdResultFromDict(data map[string]interface{}) GetCounterByUserIdResult {
	return GetCounterByUserIdResult{
		Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCounterByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCounterByUserIdResult) Pointer() *GetCounterByUserIdResult {
	return &p
}

type DeleteCounterByUserIdResult struct {
	Item *Counter `json:"item"`
}

type DeleteCounterByUserIdAsyncResult struct {
	result *DeleteCounterByUserIdResult
	err    error
}

func NewDeleteCounterByUserIdResultFromJson(data string) DeleteCounterByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCounterByUserIdResultFromDict(dict)
}

func NewDeleteCounterByUserIdResultFromDict(data map[string]interface{}) DeleteCounterByUserIdResult {
	return DeleteCounterByUserIdResult{
		Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteCounterByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteCounterByUserIdResult) Pointer() *DeleteCounterByUserIdResult {
	return &p
}

type IncreaseByStampSheetResult struct {
	Item             *Counter   `json:"item"`
	ChangedCompletes []Complete `json:"changedCompletes"`
}

type IncreaseByStampSheetAsyncResult struct {
	result *IncreaseByStampSheetResult
	err    error
}

func NewIncreaseByStampSheetResultFromJson(data string) IncreaseByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseByStampSheetResultFromDict(dict)
}

func NewIncreaseByStampSheetResultFromDict(data map[string]interface{}) IncreaseByStampSheetResult {
	return IncreaseByStampSheetResult{
		Item:             NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
		ChangedCompletes: CastCompletes(core.CastArray(data["changedCompletes"])),
	}
}

func (p IncreaseByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
		"changedCompletes": CastCompletesFromDict(
			p.ChangedCompletes,
		),
	}
}

func (p IncreaseByStampSheetResult) Pointer() *IncreaseByStampSheetResult {
	return &p
}

type DecreaseByStampTaskResult struct {
	Item            *Counter `json:"item"`
	NewContextStack *string  `json:"newContextStack"`
}

type DecreaseByStampTaskAsyncResult struct {
	result *DecreaseByStampTaskResult
	err    error
}

func NewDecreaseByStampTaskResultFromJson(data string) DecreaseByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseByStampTaskResultFromDict(dict)
}

func NewDecreaseByStampTaskResultFromDict(data map[string]interface{}) DecreaseByStampTaskResult {
	return DecreaseByStampTaskResult{
		Item:            NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p DecreaseByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":            p.Item.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p DecreaseByStampTaskResult) Pointer() *DecreaseByStampTaskResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentMissionMaster `json:"item"`
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
		Item: NewCurrentMissionMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentMissionMasterResult struct {
	Item *CurrentMissionMaster `json:"item"`
}

type GetCurrentMissionMasterAsyncResult struct {
	result *GetCurrentMissionMasterResult
	err    error
}

func NewGetCurrentMissionMasterResultFromJson(data string) GetCurrentMissionMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentMissionMasterResultFromDict(dict)
}

func NewGetCurrentMissionMasterResultFromDict(data map[string]interface{}) GetCurrentMissionMasterResult {
	return GetCurrentMissionMasterResult{
		Item: NewCurrentMissionMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentMissionMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentMissionMasterResult) Pointer() *GetCurrentMissionMasterResult {
	return &p
}

type UpdateCurrentMissionMasterResult struct {
	Item *CurrentMissionMaster `json:"item"`
}

type UpdateCurrentMissionMasterAsyncResult struct {
	result *UpdateCurrentMissionMasterResult
	err    error
}

func NewUpdateCurrentMissionMasterResultFromJson(data string) UpdateCurrentMissionMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentMissionMasterResultFromDict(dict)
}

func NewUpdateCurrentMissionMasterResultFromDict(data map[string]interface{}) UpdateCurrentMissionMasterResult {
	return UpdateCurrentMissionMasterResult{
		Item: NewCurrentMissionMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentMissionMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentMissionMasterResult) Pointer() *UpdateCurrentMissionMasterResult {
	return &p
}

type UpdateCurrentMissionMasterFromGitHubResult struct {
	Item *CurrentMissionMaster `json:"item"`
}

type UpdateCurrentMissionMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentMissionMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentMissionMasterFromGitHubResultFromJson(data string) UpdateCurrentMissionMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentMissionMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentMissionMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentMissionMasterFromGitHubResult {
	return UpdateCurrentMissionMasterFromGitHubResult{
		Item: NewCurrentMissionMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentMissionMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentMissionMasterFromGitHubResult) Pointer() *UpdateCurrentMissionMasterFromGitHubResult {
	return &p
}

type DescribeCounterModelsResult struct {
	Items []CounterModel `json:"items"`
}

type DescribeCounterModelsAsyncResult struct {
	result *DescribeCounterModelsResult
	err    error
}

func NewDescribeCounterModelsResultFromJson(data string) DescribeCounterModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCounterModelsResultFromDict(dict)
}

func NewDescribeCounterModelsResultFromDict(data map[string]interface{}) DescribeCounterModelsResult {
	return DescribeCounterModelsResult{
		Items: CastCounterModels(core.CastArray(data["items"])),
	}
}

func (p DescribeCounterModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCounterModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeCounterModelsResult) Pointer() *DescribeCounterModelsResult {
	return &p
}

type GetCounterModelResult struct {
	Item *CounterModel `json:"item"`
}

type GetCounterModelAsyncResult struct {
	result *GetCounterModelResult
	err    error
}

func NewGetCounterModelResultFromJson(data string) GetCounterModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterModelResultFromDict(dict)
}

func NewGetCounterModelResultFromDict(data map[string]interface{}) GetCounterModelResult {
	return GetCounterModelResult{
		Item: NewCounterModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCounterModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCounterModelResult) Pointer() *GetCounterModelResult {
	return &p
}

type DescribeMissionGroupModelsResult struct {
	Items []MissionGroupModel `json:"items"`
}

type DescribeMissionGroupModelsAsyncResult struct {
	result *DescribeMissionGroupModelsResult
	err    error
}

func NewDescribeMissionGroupModelsResultFromJson(data string) DescribeMissionGroupModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionGroupModelsResultFromDict(dict)
}

func NewDescribeMissionGroupModelsResultFromDict(data map[string]interface{}) DescribeMissionGroupModelsResult {
	return DescribeMissionGroupModelsResult{
		Items: CastMissionGroupModels(core.CastArray(data["items"])),
	}
}

func (p DescribeMissionGroupModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMissionGroupModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeMissionGroupModelsResult) Pointer() *DescribeMissionGroupModelsResult {
	return &p
}

type GetMissionGroupModelResult struct {
	Item *MissionGroupModel `json:"item"`
}

type GetMissionGroupModelAsyncResult struct {
	result *GetMissionGroupModelResult
	err    error
}

func NewGetMissionGroupModelResultFromJson(data string) GetMissionGroupModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionGroupModelResultFromDict(dict)
}

func NewGetMissionGroupModelResultFromDict(data map[string]interface{}) GetMissionGroupModelResult {
	return GetMissionGroupModelResult{
		Item: NewMissionGroupModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMissionGroupModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMissionGroupModelResult) Pointer() *GetMissionGroupModelResult {
	return &p
}

type DescribeMissionTaskModelsResult struct {
	Items []MissionTaskModel `json:"items"`
}

type DescribeMissionTaskModelsAsyncResult struct {
	result *DescribeMissionTaskModelsResult
	err    error
}

func NewDescribeMissionTaskModelsResultFromJson(data string) DescribeMissionTaskModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionTaskModelsResultFromDict(dict)
}

func NewDescribeMissionTaskModelsResultFromDict(data map[string]interface{}) DescribeMissionTaskModelsResult {
	return DescribeMissionTaskModelsResult{
		Items: CastMissionTaskModels(core.CastArray(data["items"])),
	}
}

func (p DescribeMissionTaskModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMissionTaskModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeMissionTaskModelsResult) Pointer() *DescribeMissionTaskModelsResult {
	return &p
}

type GetMissionTaskModelResult struct {
	Item *MissionTaskModel `json:"item"`
}

type GetMissionTaskModelAsyncResult struct {
	result *GetMissionTaskModelResult
	err    error
}

func NewGetMissionTaskModelResultFromJson(data string) GetMissionTaskModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionTaskModelResultFromDict(dict)
}

func NewGetMissionTaskModelResultFromDict(data map[string]interface{}) GetMissionTaskModelResult {
	return GetMissionTaskModelResult{
		Item: NewMissionTaskModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMissionTaskModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMissionTaskModelResult) Pointer() *GetMissionTaskModelResult {
	return &p
}

type DescribeMissionTaskModelMastersResult struct {
	Items         []MissionTaskModelMaster `json:"items"`
	NextPageToken *string                  `json:"nextPageToken"`
}

type DescribeMissionTaskModelMastersAsyncResult struct {
	result *DescribeMissionTaskModelMastersResult
	err    error
}

func NewDescribeMissionTaskModelMastersResultFromJson(data string) DescribeMissionTaskModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionTaskModelMastersResultFromDict(dict)
}

func NewDescribeMissionTaskModelMastersResultFromDict(data map[string]interface{}) DescribeMissionTaskModelMastersResult {
	return DescribeMissionTaskModelMastersResult{
		Items:         CastMissionTaskModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeMissionTaskModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMissionTaskModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeMissionTaskModelMastersResult) Pointer() *DescribeMissionTaskModelMastersResult {
	return &p
}

type CreateMissionTaskModelMasterResult struct {
	Item *MissionTaskModelMaster `json:"item"`
}

type CreateMissionTaskModelMasterAsyncResult struct {
	result *CreateMissionTaskModelMasterResult
	err    error
}

func NewCreateMissionTaskModelMasterResultFromJson(data string) CreateMissionTaskModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMissionTaskModelMasterResultFromDict(dict)
}

func NewCreateMissionTaskModelMasterResultFromDict(data map[string]interface{}) CreateMissionTaskModelMasterResult {
	return CreateMissionTaskModelMasterResult{
		Item: NewMissionTaskModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateMissionTaskModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateMissionTaskModelMasterResult) Pointer() *CreateMissionTaskModelMasterResult {
	return &p
}

type GetMissionTaskModelMasterResult struct {
	Item *MissionTaskModelMaster `json:"item"`
}

type GetMissionTaskModelMasterAsyncResult struct {
	result *GetMissionTaskModelMasterResult
	err    error
}

func NewGetMissionTaskModelMasterResultFromJson(data string) GetMissionTaskModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionTaskModelMasterResultFromDict(dict)
}

func NewGetMissionTaskModelMasterResultFromDict(data map[string]interface{}) GetMissionTaskModelMasterResult {
	return GetMissionTaskModelMasterResult{
		Item: NewMissionTaskModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMissionTaskModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMissionTaskModelMasterResult) Pointer() *GetMissionTaskModelMasterResult {
	return &p
}

type UpdateMissionTaskModelMasterResult struct {
	Item *MissionTaskModelMaster `json:"item"`
}

type UpdateMissionTaskModelMasterAsyncResult struct {
	result *UpdateMissionTaskModelMasterResult
	err    error
}

func NewUpdateMissionTaskModelMasterResultFromJson(data string) UpdateMissionTaskModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMissionTaskModelMasterResultFromDict(dict)
}

func NewUpdateMissionTaskModelMasterResultFromDict(data map[string]interface{}) UpdateMissionTaskModelMasterResult {
	return UpdateMissionTaskModelMasterResult{
		Item: NewMissionTaskModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateMissionTaskModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateMissionTaskModelMasterResult) Pointer() *UpdateMissionTaskModelMasterResult {
	return &p
}

type DeleteMissionTaskModelMasterResult struct {
	Item *MissionTaskModelMaster `json:"item"`
}

type DeleteMissionTaskModelMasterAsyncResult struct {
	result *DeleteMissionTaskModelMasterResult
	err    error
}

func NewDeleteMissionTaskModelMasterResultFromJson(data string) DeleteMissionTaskModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMissionTaskModelMasterResultFromDict(dict)
}

func NewDeleteMissionTaskModelMasterResultFromDict(data map[string]interface{}) DeleteMissionTaskModelMasterResult {
	return DeleteMissionTaskModelMasterResult{
		Item: NewMissionTaskModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteMissionTaskModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteMissionTaskModelMasterResult) Pointer() *DeleteMissionTaskModelMasterResult {
	return &p
}
