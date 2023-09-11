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

package skillTree

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

type DescribeNodeModelsResult struct {
	Items []NodeModel `json:"items"`
}

type DescribeNodeModelsAsyncResult struct {
	result *DescribeNodeModelsResult
	err    error
}

func NewDescribeNodeModelsResultFromJson(data string) DescribeNodeModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNodeModelsResultFromDict(dict)
}

func NewDescribeNodeModelsResultFromDict(data map[string]interface{}) DescribeNodeModelsResult {
	return DescribeNodeModelsResult{
		Items: CastNodeModels(core.CastArray(data["items"])),
	}
}

func (p DescribeNodeModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNodeModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeNodeModelsResult) Pointer() *DescribeNodeModelsResult {
	return &p
}

type GetNodeModelResult struct {
	Item *NodeModel `json:"item"`
}

type GetNodeModelAsyncResult struct {
	result *GetNodeModelResult
	err    error
}

func NewGetNodeModelResultFromJson(data string) GetNodeModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNodeModelResultFromDict(dict)
}

func NewGetNodeModelResultFromDict(data map[string]interface{}) GetNodeModelResult {
	return GetNodeModelResult{
		Item: NewNodeModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetNodeModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetNodeModelResult) Pointer() *GetNodeModelResult {
	return &p
}

type DescribeNodeModelMastersResult struct {
	Items         []NodeModelMaster `json:"items"`
	NextPageToken *string           `json:"nextPageToken"`
}

type DescribeNodeModelMastersAsyncResult struct {
	result *DescribeNodeModelMastersResult
	err    error
}

func NewDescribeNodeModelMastersResultFromJson(data string) DescribeNodeModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNodeModelMastersResultFromDict(dict)
}

func NewDescribeNodeModelMastersResultFromDict(data map[string]interface{}) DescribeNodeModelMastersResult {
	return DescribeNodeModelMastersResult{
		Items:         CastNodeModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeNodeModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNodeModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeNodeModelMastersResult) Pointer() *DescribeNodeModelMastersResult {
	return &p
}

type CreateNodeModelMasterResult struct {
	Item *NodeModelMaster `json:"item"`
}

type CreateNodeModelMasterAsyncResult struct {
	result *CreateNodeModelMasterResult
	err    error
}

func NewCreateNodeModelMasterResultFromJson(data string) CreateNodeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNodeModelMasterResultFromDict(dict)
}

func NewCreateNodeModelMasterResultFromDict(data map[string]interface{}) CreateNodeModelMasterResult {
	return CreateNodeModelMasterResult{
		Item: NewNodeModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateNodeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateNodeModelMasterResult) Pointer() *CreateNodeModelMasterResult {
	return &p
}

type GetNodeModelMasterResult struct {
	Item *NodeModelMaster `json:"item"`
}

type GetNodeModelMasterAsyncResult struct {
	result *GetNodeModelMasterResult
	err    error
}

func NewGetNodeModelMasterResultFromJson(data string) GetNodeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNodeModelMasterResultFromDict(dict)
}

func NewGetNodeModelMasterResultFromDict(data map[string]interface{}) GetNodeModelMasterResult {
	return GetNodeModelMasterResult{
		Item: NewNodeModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetNodeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetNodeModelMasterResult) Pointer() *GetNodeModelMasterResult {
	return &p
}

type UpdateNodeModelMasterResult struct {
	Item *NodeModelMaster `json:"item"`
}

type UpdateNodeModelMasterAsyncResult struct {
	result *UpdateNodeModelMasterResult
	err    error
}

func NewUpdateNodeModelMasterResultFromJson(data string) UpdateNodeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNodeModelMasterResultFromDict(dict)
}

func NewUpdateNodeModelMasterResultFromDict(data map[string]interface{}) UpdateNodeModelMasterResult {
	return UpdateNodeModelMasterResult{
		Item: NewNodeModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateNodeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateNodeModelMasterResult) Pointer() *UpdateNodeModelMasterResult {
	return &p
}

type DeleteNodeModelMasterResult struct {
	Item *NodeModelMaster `json:"item"`
}

type DeleteNodeModelMasterAsyncResult struct {
	result *DeleteNodeModelMasterResult
	err    error
}

func NewDeleteNodeModelMasterResultFromJson(data string) DeleteNodeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteNodeModelMasterResultFromDict(dict)
}

func NewDeleteNodeModelMasterResultFromDict(data map[string]interface{}) DeleteNodeModelMasterResult {
	return DeleteNodeModelMasterResult{
		Item: NewNodeModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteNodeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteNodeModelMasterResult) Pointer() *DeleteNodeModelMasterResult {
	return &p
}

type MarkReleaseByUserIdResult struct {
	Item *Status `json:"item"`
}

type MarkReleaseByUserIdAsyncResult struct {
	result *MarkReleaseByUserIdResult
	err    error
}

func NewMarkReleaseByUserIdResultFromJson(data string) MarkReleaseByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReleaseByUserIdResultFromDict(dict)
}

func NewMarkReleaseByUserIdResultFromDict(data map[string]interface{}) MarkReleaseByUserIdResult {
	return MarkReleaseByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p MarkReleaseByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p MarkReleaseByUserIdResult) Pointer() *MarkReleaseByUserIdResult {
	return &p
}

type ReleaseResult struct {
	Item                      *Status `json:"item"`
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type ReleaseAsyncResult struct {
	result *ReleaseResult
	err    error
}

func NewReleaseResultFromJson(data string) ReleaseResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReleaseResultFromDict(dict)
}

func NewReleaseResultFromDict(data map[string]interface{}) ReleaseResult {
	return ReleaseResult{
		Item:                      NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p ReleaseResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ReleaseResult) Pointer() *ReleaseResult {
	return &p
}

type ReleaseByUserIdResult struct {
	Item                      *Status `json:"item"`
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type ReleaseByUserIdAsyncResult struct {
	result *ReleaseByUserIdResult
	err    error
}

func NewReleaseByUserIdResultFromJson(data string) ReleaseByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReleaseByUserIdResultFromDict(dict)
}

func NewReleaseByUserIdResultFromDict(data map[string]interface{}) ReleaseByUserIdResult {
	return ReleaseByUserIdResult{
		Item:                      NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p ReleaseByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ReleaseByUserIdResult) Pointer() *ReleaseByUserIdResult {
	return &p
}

type MarkRestrainByUserIdResult struct {
	Item *Status `json:"item"`
}

type MarkRestrainByUserIdAsyncResult struct {
	result *MarkRestrainByUserIdResult
	err    error
}

func NewMarkRestrainByUserIdResultFromJson(data string) MarkRestrainByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkRestrainByUserIdResultFromDict(dict)
}

func NewMarkRestrainByUserIdResultFromDict(data map[string]interface{}) MarkRestrainByUserIdResult {
	return MarkRestrainByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p MarkRestrainByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p MarkRestrainByUserIdResult) Pointer() *MarkRestrainByUserIdResult {
	return &p
}

type RestrainResult struct {
	Item                      *Status `json:"item"`
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type RestrainAsyncResult struct {
	result *RestrainResult
	err    error
}

func NewRestrainResultFromJson(data string) RestrainResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRestrainResultFromDict(dict)
}

func NewRestrainResultFromDict(data map[string]interface{}) RestrainResult {
	return RestrainResult{
		Item:                      NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p RestrainResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p RestrainResult) Pointer() *RestrainResult {
	return &p
}

type RestrainByUserIdResult struct {
	Item                      *Status `json:"item"`
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type RestrainByUserIdAsyncResult struct {
	result *RestrainByUserIdResult
	err    error
}

func NewRestrainByUserIdResultFromJson(data string) RestrainByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRestrainByUserIdResultFromDict(dict)
}

func NewRestrainByUserIdResultFromDict(data map[string]interface{}) RestrainByUserIdResult {
	return RestrainByUserIdResult{
		Item:                      NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p RestrainByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p RestrainByUserIdResult) Pointer() *RestrainByUserIdResult {
	return &p
}

type GetStatusResult struct {
	Item *Status `json:"item"`
}

type GetStatusAsyncResult struct {
	result *GetStatusResult
	err    error
}

func NewGetStatusResultFromJson(data string) GetStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusResultFromDict(dict)
}

func NewGetStatusResultFromDict(data map[string]interface{}) GetStatusResult {
	return GetStatusResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetStatusResult) Pointer() *GetStatusResult {
	return &p
}

type GetStatusByUserIdResult struct {
	Item *Status `json:"item"`
}

type GetStatusByUserIdAsyncResult struct {
	result *GetStatusByUserIdResult
	err    error
}

func NewGetStatusByUserIdResultFromJson(data string) GetStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusByUserIdResultFromDict(dict)
}

func NewGetStatusByUserIdResultFromDict(data map[string]interface{}) GetStatusByUserIdResult {
	return GetStatusByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetStatusByUserIdResult) Pointer() *GetStatusByUserIdResult {
	return &p
}

type ResetResult struct {
	Item                      *Status `json:"item"`
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type ResetAsyncResult struct {
	result *ResetResult
	err    error
}

func NewResetResultFromJson(data string) ResetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetResultFromDict(dict)
}

func NewResetResultFromDict(data map[string]interface{}) ResetResult {
	return ResetResult{
		Item:                      NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p ResetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ResetResult) Pointer() *ResetResult {
	return &p
}

type ResetByUserIdResult struct {
	Item                      *Status `json:"item"`
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type ResetByUserIdAsyncResult struct {
	result *ResetByUserIdResult
	err    error
}

func NewResetByUserIdResultFromJson(data string) ResetByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetByUserIdResultFromDict(dict)
}

func NewResetByUserIdResultFromDict(data map[string]interface{}) ResetByUserIdResult {
	return ResetByUserIdResult{
		Item:                      NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p ResetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ResetByUserIdResult) Pointer() *ResetByUserIdResult {
	return &p
}

type MarkReleaseByStampSheetResult struct {
	Item *Status `json:"item"`
}

type MarkReleaseByStampSheetAsyncResult struct {
	result *MarkReleaseByStampSheetResult
	err    error
}

func NewMarkReleaseByStampSheetResultFromJson(data string) MarkReleaseByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReleaseByStampSheetResultFromDict(dict)
}

func NewMarkReleaseByStampSheetResultFromDict(data map[string]interface{}) MarkReleaseByStampSheetResult {
	return MarkReleaseByStampSheetResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p MarkReleaseByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p MarkReleaseByStampSheetResult) Pointer() *MarkReleaseByStampSheetResult {
	return &p
}

type MarkRestrainByStampTaskResult struct {
	Item            *Status `json:"item"`
	NewContextStack *string `json:"newContextStack"`
}

type MarkRestrainByStampTaskAsyncResult struct {
	result *MarkRestrainByStampTaskResult
	err    error
}

func NewMarkRestrainByStampTaskResultFromJson(data string) MarkRestrainByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkRestrainByStampTaskResultFromDict(dict)
}

func NewMarkRestrainByStampTaskResultFromDict(data map[string]interface{}) MarkRestrainByStampTaskResult {
	return MarkRestrainByStampTaskResult{
		Item:            NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p MarkRestrainByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":            p.Item.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p MarkRestrainByStampTaskResult) Pointer() *MarkRestrainByStampTaskResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentTreeMaster `json:"item"`
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
		Item: NewCurrentTreeMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentTreeMasterResult struct {
	Item *CurrentTreeMaster `json:"item"`
}

type GetCurrentTreeMasterAsyncResult struct {
	result *GetCurrentTreeMasterResult
	err    error
}

func NewGetCurrentTreeMasterResultFromJson(data string) GetCurrentTreeMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentTreeMasterResultFromDict(dict)
}

func NewGetCurrentTreeMasterResultFromDict(data map[string]interface{}) GetCurrentTreeMasterResult {
	return GetCurrentTreeMasterResult{
		Item: NewCurrentTreeMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentTreeMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentTreeMasterResult) Pointer() *GetCurrentTreeMasterResult {
	return &p
}

type UpdateCurrentTreeMasterResult struct {
	Item *CurrentTreeMaster `json:"item"`
}

type UpdateCurrentTreeMasterAsyncResult struct {
	result *UpdateCurrentTreeMasterResult
	err    error
}

func NewUpdateCurrentTreeMasterResultFromJson(data string) UpdateCurrentTreeMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentTreeMasterResultFromDict(dict)
}

func NewUpdateCurrentTreeMasterResultFromDict(data map[string]interface{}) UpdateCurrentTreeMasterResult {
	return UpdateCurrentTreeMasterResult{
		Item: NewCurrentTreeMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentTreeMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentTreeMasterResult) Pointer() *UpdateCurrentTreeMasterResult {
	return &p
}

type UpdateCurrentTreeMasterFromGitHubResult struct {
	Item *CurrentTreeMaster `json:"item"`
}

type UpdateCurrentTreeMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentTreeMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentTreeMasterFromGitHubResultFromJson(data string) UpdateCurrentTreeMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentTreeMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentTreeMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentTreeMasterFromGitHubResult {
	return UpdateCurrentTreeMasterFromGitHubResult{
		Item: NewCurrentTreeMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentTreeMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentTreeMasterFromGitHubResult) Pointer() *UpdateCurrentTreeMasterFromGitHubResult {
	return &p
}
