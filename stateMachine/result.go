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

package stateMachine

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

type DescribeStateMachineMastersResult struct {
	Items         []StateMachineMaster `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
}

type DescribeStateMachineMastersAsyncResult struct {
	result *DescribeStateMachineMastersResult
	err    error
}

func NewDescribeStateMachineMastersResultFromJson(data string) DescribeStateMachineMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStateMachineMastersResultFromDict(dict)
}

func NewDescribeStateMachineMastersResultFromDict(data map[string]interface{}) DescribeStateMachineMastersResult {
	return DescribeStateMachineMastersResult{
		Items:         CastStateMachineMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStateMachineMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStateMachineMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStateMachineMastersResult) Pointer() *DescribeStateMachineMastersResult {
	return &p
}

type UpdateStateMachineMasterResult struct {
	Item *StateMachineMaster `json:"item"`
}

type UpdateStateMachineMasterAsyncResult struct {
	result *UpdateStateMachineMasterResult
	err    error
}

func NewUpdateStateMachineMasterResultFromJson(data string) UpdateStateMachineMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStateMachineMasterResultFromDict(dict)
}

func NewUpdateStateMachineMasterResultFromDict(data map[string]interface{}) UpdateStateMachineMasterResult {
	return UpdateStateMachineMasterResult{
		Item: NewStateMachineMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateStateMachineMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateStateMachineMasterResult) Pointer() *UpdateStateMachineMasterResult {
	return &p
}

type GetStateMachineMasterResult struct {
	Item *StateMachineMaster `json:"item"`
}

type GetStateMachineMasterAsyncResult struct {
	result *GetStateMachineMasterResult
	err    error
}

func NewGetStateMachineMasterResultFromJson(data string) GetStateMachineMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStateMachineMasterResultFromDict(dict)
}

func NewGetStateMachineMasterResultFromDict(data map[string]interface{}) GetStateMachineMasterResult {
	return GetStateMachineMasterResult{
		Item: NewStateMachineMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStateMachineMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetStateMachineMasterResult) Pointer() *GetStateMachineMasterResult {
	return &p
}

type DeleteStateMachineMasterResult struct {
	Item *StateMachineMaster `json:"item"`
}

type DeleteStateMachineMasterAsyncResult struct {
	result *DeleteStateMachineMasterResult
	err    error
}

func NewDeleteStateMachineMasterResultFromJson(data string) DeleteStateMachineMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStateMachineMasterResultFromDict(dict)
}

func NewDeleteStateMachineMasterResultFromDict(data map[string]interface{}) DeleteStateMachineMasterResult {
	return DeleteStateMachineMasterResult{
		Item: NewStateMachineMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteStateMachineMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteStateMachineMasterResult) Pointer() *DeleteStateMachineMasterResult {
	return &p
}

type DescribeStatusesResult struct {
	Items         []Status `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
}

type DescribeStatusesAsyncResult struct {
	result *DescribeStatusesResult
	err    error
}

func NewDescribeStatusesResultFromJson(data string) DescribeStatusesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesResultFromDict(dict)
}

func NewDescribeStatusesResultFromDict(data map[string]interface{}) DescribeStatusesResult {
	return DescribeStatusesResult{
		Items:         CastStatuses(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStatusesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStatusesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStatusesResult) Pointer() *DescribeStatusesResult {
	return &p
}

type DescribeStatusesByUserIdResult struct {
	Items         []Status `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
}

type DescribeStatusesByUserIdAsyncResult struct {
	result *DescribeStatusesByUserIdResult
	err    error
}

func NewDescribeStatusesByUserIdResultFromJson(data string) DescribeStatusesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesByUserIdResultFromDict(dict)
}

func NewDescribeStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeStatusesByUserIdResult {
	return DescribeStatusesByUserIdResult{
		Items:         CastStatuses(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStatusesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStatusesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStatusesByUserIdResult) Pointer() *DescribeStatusesByUserIdResult {
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

type StartStateMachineByUserIdResult struct {
	Item *Status `json:"item"`
}

type StartStateMachineByUserIdAsyncResult struct {
	result *StartStateMachineByUserIdResult
	err    error
}

func NewStartStateMachineByUserIdResultFromJson(data string) StartStateMachineByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStartStateMachineByUserIdResultFromDict(dict)
}

func NewStartStateMachineByUserIdResultFromDict(data map[string]interface{}) StartStateMachineByUserIdResult {
	return StartStateMachineByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p StartStateMachineByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p StartStateMachineByUserIdResult) Pointer() *StartStateMachineByUserIdResult {
	return &p
}

type StartStateMachineByStampSheetResult struct {
	Item *Status `json:"item"`
}

type StartStateMachineByStampSheetAsyncResult struct {
	result *StartStateMachineByStampSheetResult
	err    error
}

func NewStartStateMachineByStampSheetResultFromJson(data string) StartStateMachineByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStartStateMachineByStampSheetResultFromDict(dict)
}

func NewStartStateMachineByStampSheetResultFromDict(data map[string]interface{}) StartStateMachineByStampSheetResult {
	return StartStateMachineByStampSheetResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p StartStateMachineByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p StartStateMachineByStampSheetResult) Pointer() *StartStateMachineByStampSheetResult {
	return &p
}

type EmitResult struct {
	Item *Status `json:"item"`
}

type EmitAsyncResult struct {
	result *EmitResult
	err    error
}

func NewEmitResultFromJson(data string) EmitResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEmitResultFromDict(dict)
}

func NewEmitResultFromDict(data map[string]interface{}) EmitResult {
	return EmitResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p EmitResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p EmitResult) Pointer() *EmitResult {
	return &p
}

type EmitByUserIdResult struct {
	Item *Status `json:"item"`
}

type EmitByUserIdAsyncResult struct {
	result *EmitByUserIdResult
	err    error
}

func NewEmitByUserIdResultFromJson(data string) EmitByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEmitByUserIdResultFromDict(dict)
}

func NewEmitByUserIdResultFromDict(data map[string]interface{}) EmitByUserIdResult {
	return EmitByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p EmitByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p EmitByUserIdResult) Pointer() *EmitByUserIdResult {
	return &p
}

type DeleteStatusByUserIdResult struct {
	Item *Status `json:"item"`
}

type DeleteStatusByUserIdAsyncResult struct {
	result *DeleteStatusByUserIdResult
	err    error
}

func NewDeleteStatusByUserIdResultFromJson(data string) DeleteStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStatusByUserIdResultFromDict(dict)
}

func NewDeleteStatusByUserIdResultFromDict(data map[string]interface{}) DeleteStatusByUserIdResult {
	return DeleteStatusByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteStatusByUserIdResult) Pointer() *DeleteStatusByUserIdResult {
	return &p
}

type ExitStateMachineResult struct {
	Item *Status `json:"item"`
}

type ExitStateMachineAsyncResult struct {
	result *ExitStateMachineResult
	err    error
}

func NewExitStateMachineResultFromJson(data string) ExitStateMachineResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExitStateMachineResultFromDict(dict)
}

func NewExitStateMachineResultFromDict(data map[string]interface{}) ExitStateMachineResult {
	return ExitStateMachineResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ExitStateMachineResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p ExitStateMachineResult) Pointer() *ExitStateMachineResult {
	return &p
}

type ExitStateMachineByUserIdResult struct {
	Item *Status `json:"item"`
}

type ExitStateMachineByUserIdAsyncResult struct {
	result *ExitStateMachineByUserIdResult
	err    error
}

func NewExitStateMachineByUserIdResultFromJson(data string) ExitStateMachineByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExitStateMachineByUserIdResultFromDict(dict)
}

func NewExitStateMachineByUserIdResultFromDict(data map[string]interface{}) ExitStateMachineByUserIdResult {
	return ExitStateMachineByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ExitStateMachineByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p ExitStateMachineByUserIdResult) Pointer() *ExitStateMachineByUserIdResult {
	return &p
}
