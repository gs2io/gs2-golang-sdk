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

package schedule

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
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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

type DescribeEventMastersResult struct {
	Items         []EventMaster `json:"items"`
	NextPageToken *string       `json:"nextPageToken"`
}

type DescribeEventMastersAsyncResult struct {
	result *DescribeEventMastersResult
	err    error
}

func NewDescribeEventMastersResultFromJson(data string) DescribeEventMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEventMastersResultFromDict(dict)
}

func NewDescribeEventMastersResultFromDict(data map[string]interface{}) DescribeEventMastersResult {
	return DescribeEventMastersResult{
		Items:         CastEventMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeEventMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEventMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeEventMastersResult) Pointer() *DescribeEventMastersResult {
	return &p
}

type CreateEventMasterResult struct {
	Item *EventMaster `json:"item"`
}

type CreateEventMasterAsyncResult struct {
	result *CreateEventMasterResult
	err    error
}

func NewCreateEventMasterResultFromJson(data string) CreateEventMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateEventMasterResultFromDict(dict)
}

func NewCreateEventMasterResultFromDict(data map[string]interface{}) CreateEventMasterResult {
	return CreateEventMasterResult{
		Item: NewEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateEventMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateEventMasterResult) Pointer() *CreateEventMasterResult {
	return &p
}

type GetEventMasterResult struct {
	Item *EventMaster `json:"item"`
}

type GetEventMasterAsyncResult struct {
	result *GetEventMasterResult
	err    error
}

func NewGetEventMasterResultFromJson(data string) GetEventMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEventMasterResultFromDict(dict)
}

func NewGetEventMasterResultFromDict(data map[string]interface{}) GetEventMasterResult {
	return GetEventMasterResult{
		Item: NewEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetEventMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetEventMasterResult) Pointer() *GetEventMasterResult {
	return &p
}

type UpdateEventMasterResult struct {
	Item *EventMaster `json:"item"`
}

type UpdateEventMasterAsyncResult struct {
	result *UpdateEventMasterResult
	err    error
}

func NewUpdateEventMasterResultFromJson(data string) UpdateEventMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateEventMasterResultFromDict(dict)
}

func NewUpdateEventMasterResultFromDict(data map[string]interface{}) UpdateEventMasterResult {
	return UpdateEventMasterResult{
		Item: NewEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateEventMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateEventMasterResult) Pointer() *UpdateEventMasterResult {
	return &p
}

type DeleteEventMasterResult struct {
	Item *EventMaster `json:"item"`
}

type DeleteEventMasterAsyncResult struct {
	result *DeleteEventMasterResult
	err    error
}

func NewDeleteEventMasterResultFromJson(data string) DeleteEventMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteEventMasterResultFromDict(dict)
}

func NewDeleteEventMasterResultFromDict(data map[string]interface{}) DeleteEventMasterResult {
	return DeleteEventMasterResult{
		Item: NewEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteEventMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteEventMasterResult) Pointer() *DeleteEventMasterResult {
	return &p
}

type DescribeTriggersResult struct {
	Items         []Trigger `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeTriggersAsyncResult struct {
	result *DescribeTriggersResult
	err    error
}

func NewDescribeTriggersResultFromJson(data string) DescribeTriggersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTriggersResultFromDict(dict)
}

func NewDescribeTriggersResultFromDict(data map[string]interface{}) DescribeTriggersResult {
	return DescribeTriggersResult{
		Items:         CastTriggers(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeTriggersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastTriggersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeTriggersResult) Pointer() *DescribeTriggersResult {
	return &p
}

type DescribeTriggersByUserIdResult struct {
	Items         []Trigger `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeTriggersByUserIdAsyncResult struct {
	result *DescribeTriggersByUserIdResult
	err    error
}

func NewDescribeTriggersByUserIdResultFromJson(data string) DescribeTriggersByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTriggersByUserIdResultFromDict(dict)
}

func NewDescribeTriggersByUserIdResultFromDict(data map[string]interface{}) DescribeTriggersByUserIdResult {
	return DescribeTriggersByUserIdResult{
		Items:         CastTriggers(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeTriggersByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastTriggersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeTriggersByUserIdResult) Pointer() *DescribeTriggersByUserIdResult {
	return &p
}

type GetTriggerResult struct {
	Item *Trigger `json:"item"`
}

type GetTriggerAsyncResult struct {
	result *GetTriggerResult
	err    error
}

func NewGetTriggerResultFromJson(data string) GetTriggerResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTriggerResultFromDict(dict)
}

func NewGetTriggerResultFromDict(data map[string]interface{}) GetTriggerResult {
	return GetTriggerResult{
		Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetTriggerResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetTriggerResult) Pointer() *GetTriggerResult {
	return &p
}

type GetTriggerByUserIdResult struct {
	Item *Trigger `json:"item"`
}

type GetTriggerByUserIdAsyncResult struct {
	result *GetTriggerByUserIdResult
	err    error
}

func NewGetTriggerByUserIdResultFromJson(data string) GetTriggerByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTriggerByUserIdResultFromDict(dict)
}

func NewGetTriggerByUserIdResultFromDict(data map[string]interface{}) GetTriggerByUserIdResult {
	return GetTriggerByUserIdResult{
		Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetTriggerByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetTriggerByUserIdResult) Pointer() *GetTriggerByUserIdResult {
	return &p
}

type TriggerByUserIdResult struct {
	Item *Trigger `json:"item"`
}

type TriggerByUserIdAsyncResult struct {
	result *TriggerByUserIdResult
	err    error
}

func NewTriggerByUserIdResultFromJson(data string) TriggerByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTriggerByUserIdResultFromDict(dict)
}

func NewTriggerByUserIdResultFromDict(data map[string]interface{}) TriggerByUserIdResult {
	return TriggerByUserIdResult{
		Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p TriggerByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p TriggerByUserIdResult) Pointer() *TriggerByUserIdResult {
	return &p
}

type TriggerByStampSheetResult struct {
	Item *Trigger `json:"item"`
}

type TriggerByStampSheetAsyncResult struct {
	result *TriggerByStampSheetResult
	err    error
}

func NewTriggerByStampSheetResultFromJson(data string) TriggerByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTriggerByStampSheetResultFromDict(dict)
}

func NewTriggerByStampSheetResultFromDict(data map[string]interface{}) TriggerByStampSheetResult {
	return TriggerByStampSheetResult{
		Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p TriggerByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p TriggerByStampSheetResult) Pointer() *TriggerByStampSheetResult {
	return &p
}

type DeleteTriggerResult struct {
	Item *Trigger `json:"item"`
}

type DeleteTriggerAsyncResult struct {
	result *DeleteTriggerResult
	err    error
}

func NewDeleteTriggerResultFromJson(data string) DeleteTriggerResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTriggerResultFromDict(dict)
}

func NewDeleteTriggerResultFromDict(data map[string]interface{}) DeleteTriggerResult {
	return DeleteTriggerResult{
		Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteTriggerResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteTriggerResult) Pointer() *DeleteTriggerResult {
	return &p
}

type DeleteTriggerByUserIdResult struct {
	Item *Trigger `json:"item"`
}

type DeleteTriggerByUserIdAsyncResult struct {
	result *DeleteTriggerByUserIdResult
	err    error
}

func NewDeleteTriggerByUserIdResultFromJson(data string) DeleteTriggerByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTriggerByUserIdResultFromDict(dict)
}

func NewDeleteTriggerByUserIdResultFromDict(data map[string]interface{}) DeleteTriggerByUserIdResult {
	return DeleteTriggerByUserIdResult{
		Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteTriggerByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteTriggerByUserIdResult) Pointer() *DeleteTriggerByUserIdResult {
	return &p
}

type VerifyTriggerResult struct {
}

type VerifyTriggerAsyncResult struct {
	result *VerifyTriggerResult
	err    error
}

func NewVerifyTriggerResultFromJson(data string) VerifyTriggerResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyTriggerResultFromDict(dict)
}

func NewVerifyTriggerResultFromDict(data map[string]interface{}) VerifyTriggerResult {
	return VerifyTriggerResult{}
}

func (p VerifyTriggerResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyTriggerResult) Pointer() *VerifyTriggerResult {
	return &p
}

type VerifyTriggerByUserIdResult struct {
}

type VerifyTriggerByUserIdAsyncResult struct {
	result *VerifyTriggerByUserIdResult
	err    error
}

func NewVerifyTriggerByUserIdResultFromJson(data string) VerifyTriggerByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyTriggerByUserIdResultFromDict(dict)
}

func NewVerifyTriggerByUserIdResultFromDict(data map[string]interface{}) VerifyTriggerByUserIdResult {
	return VerifyTriggerByUserIdResult{}
}

func (p VerifyTriggerByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyTriggerByUserIdResult) Pointer() *VerifyTriggerByUserIdResult {
	return &p
}

type DeleteTriggerByStampTaskResult struct {
	Item            *Trigger `json:"item"`
	NewContextStack *string  `json:"newContextStack"`
}

type DeleteTriggerByStampTaskAsyncResult struct {
	result *DeleteTriggerByStampTaskResult
	err    error
}

func NewDeleteTriggerByStampTaskResultFromJson(data string) DeleteTriggerByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTriggerByStampTaskResultFromDict(dict)
}

func NewDeleteTriggerByStampTaskResultFromDict(data map[string]interface{}) DeleteTriggerByStampTaskResult {
	return DeleteTriggerByStampTaskResult{
		Item:            NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p DeleteTriggerByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
	}
}

func (p DeleteTriggerByStampTaskResult) Pointer() *DeleteTriggerByStampTaskResult {
	return &p
}

type VerifyTriggerByStampTaskResult struct {
	NewContextStack *string `json:"newContextStack"`
}

type VerifyTriggerByStampTaskAsyncResult struct {
	result *VerifyTriggerByStampTaskResult
	err    error
}

func NewVerifyTriggerByStampTaskResultFromJson(data string) VerifyTriggerByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyTriggerByStampTaskResultFromDict(dict)
}

func NewVerifyTriggerByStampTaskResultFromDict(data map[string]interface{}) VerifyTriggerByStampTaskResult {
	return VerifyTriggerByStampTaskResult{
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p VerifyTriggerByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p VerifyTriggerByStampTaskResult) Pointer() *VerifyTriggerByStampTaskResult {
	return &p
}

type DescribeEventsResult struct {
	Items []Event `json:"items"`
}

type DescribeEventsAsyncResult struct {
	result *DescribeEventsResult
	err    error
}

func NewDescribeEventsResultFromJson(data string) DescribeEventsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEventsResultFromDict(dict)
}

func NewDescribeEventsResultFromDict(data map[string]interface{}) DescribeEventsResult {
	return DescribeEventsResult{
		Items: CastEvents(core.CastArray(data["items"])),
	}
}

func (p DescribeEventsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEventsFromDict(
			p.Items,
		),
	}
}

func (p DescribeEventsResult) Pointer() *DescribeEventsResult {
	return &p
}

type DescribeEventsByUserIdResult struct {
	Items []Event `json:"items"`
}

type DescribeEventsByUserIdAsyncResult struct {
	result *DescribeEventsByUserIdResult
	err    error
}

func NewDescribeEventsByUserIdResultFromJson(data string) DescribeEventsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEventsByUserIdResultFromDict(dict)
}

func NewDescribeEventsByUserIdResultFromDict(data map[string]interface{}) DescribeEventsByUserIdResult {
	return DescribeEventsByUserIdResult{
		Items: CastEvents(core.CastArray(data["items"])),
	}
}

func (p DescribeEventsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEventsFromDict(
			p.Items,
		),
	}
}

func (p DescribeEventsByUserIdResult) Pointer() *DescribeEventsByUserIdResult {
	return &p
}

type DescribeRawEventsResult struct {
	Items []Event `json:"items"`
}

type DescribeRawEventsAsyncResult struct {
	result *DescribeRawEventsResult
	err    error
}

func NewDescribeRawEventsResultFromJson(data string) DescribeRawEventsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRawEventsResultFromDict(dict)
}

func NewDescribeRawEventsResultFromDict(data map[string]interface{}) DescribeRawEventsResult {
	return DescribeRawEventsResult{
		Items: CastEvents(core.CastArray(data["items"])),
	}
}

func (p DescribeRawEventsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEventsFromDict(
			p.Items,
		),
	}
}

func (p DescribeRawEventsResult) Pointer() *DescribeRawEventsResult {
	return &p
}

type GetEventResult struct {
	Item            *Event          `json:"item"`
	InSchedule      *bool           `json:"inSchedule"`
	ScheduleStartAt *int64          `json:"scheduleStartAt"`
	ScheduleEndAt   *int64          `json:"scheduleEndAt"`
	RepeatSchedule  *RepeatSchedule `json:"repeatSchedule"`
}

type GetEventAsyncResult struct {
	result *GetEventResult
	err    error
}

func NewGetEventResultFromJson(data string) GetEventResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEventResultFromDict(dict)
}

func NewGetEventResultFromDict(data map[string]interface{}) GetEventResult {
	return GetEventResult{
		Item:            NewEventFromDict(core.CastMap(data["item"])).Pointer(),
		InSchedule:      core.CastBool(data["inSchedule"]),
		ScheduleStartAt: core.CastInt64(data["scheduleStartAt"]),
		ScheduleEndAt:   core.CastInt64(data["scheduleEndAt"]),
		RepeatSchedule:  NewRepeatScheduleFromDict(core.CastMap(data["repeatSchedule"])).Pointer(),
	}
}

func (p GetEventResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"inSchedule":      p.InSchedule,
		"scheduleStartAt": p.ScheduleStartAt,
		"scheduleEndAt":   p.ScheduleEndAt,
		"repeatSchedule": func() map[string]interface{} {
			if p.RepeatSchedule == nil {
				return nil
			}
			return p.RepeatSchedule.ToDict()
		}(),
	}
}

func (p GetEventResult) Pointer() *GetEventResult {
	return &p
}

type GetEventByUserIdResult struct {
	Item            *Event          `json:"item"`
	InSchedule      *bool           `json:"inSchedule"`
	ScheduleStartAt *int64          `json:"scheduleStartAt"`
	ScheduleEndAt   *int64          `json:"scheduleEndAt"`
	RepeatSchedule  *RepeatSchedule `json:"repeatSchedule"`
}

type GetEventByUserIdAsyncResult struct {
	result *GetEventByUserIdResult
	err    error
}

func NewGetEventByUserIdResultFromJson(data string) GetEventByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEventByUserIdResultFromDict(dict)
}

func NewGetEventByUserIdResultFromDict(data map[string]interface{}) GetEventByUserIdResult {
	return GetEventByUserIdResult{
		Item:            NewEventFromDict(core.CastMap(data["item"])).Pointer(),
		InSchedule:      core.CastBool(data["inSchedule"]),
		ScheduleStartAt: core.CastInt64(data["scheduleStartAt"]),
		ScheduleEndAt:   core.CastInt64(data["scheduleEndAt"]),
		RepeatSchedule:  NewRepeatScheduleFromDict(core.CastMap(data["repeatSchedule"])).Pointer(),
	}
}

func (p GetEventByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"inSchedule":      p.InSchedule,
		"scheduleStartAt": p.ScheduleStartAt,
		"scheduleEndAt":   p.ScheduleEndAt,
		"repeatSchedule": func() map[string]interface{} {
			if p.RepeatSchedule == nil {
				return nil
			}
			return p.RepeatSchedule.ToDict()
		}(),
	}
}

func (p GetEventByUserIdResult) Pointer() *GetEventByUserIdResult {
	return &p
}

type GetRawEventResult struct {
	Item *Event `json:"item"`
}

type GetRawEventAsyncResult struct {
	result *GetRawEventResult
	err    error
}

func NewGetRawEventResultFromJson(data string) GetRawEventResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRawEventResultFromDict(dict)
}

func NewGetRawEventResultFromDict(data map[string]interface{}) GetRawEventResult {
	return GetRawEventResult{
		Item: NewEventFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRawEventResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetRawEventResult) Pointer() *GetRawEventResult {
	return &p
}

type VerifyEventResult struct {
}

type VerifyEventAsyncResult struct {
	result *VerifyEventResult
	err    error
}

func NewVerifyEventResultFromJson(data string) VerifyEventResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyEventResultFromDict(dict)
}

func NewVerifyEventResultFromDict(data map[string]interface{}) VerifyEventResult {
	return VerifyEventResult{}
}

func (p VerifyEventResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyEventResult) Pointer() *VerifyEventResult {
	return &p
}

type VerifyEventByUserIdResult struct {
}

type VerifyEventByUserIdAsyncResult struct {
	result *VerifyEventByUserIdResult
	err    error
}

func NewVerifyEventByUserIdResultFromJson(data string) VerifyEventByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyEventByUserIdResultFromDict(dict)
}

func NewVerifyEventByUserIdResultFromDict(data map[string]interface{}) VerifyEventByUserIdResult {
	return VerifyEventByUserIdResult{}
}

func (p VerifyEventByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyEventByUserIdResult) Pointer() *VerifyEventByUserIdResult {
	return &p
}

type VerifyEventByStampTaskResult struct {
	NewContextStack *string `json:"newContextStack"`
}

type VerifyEventByStampTaskAsyncResult struct {
	result *VerifyEventByStampTaskResult
	err    error
}

func NewVerifyEventByStampTaskResultFromJson(data string) VerifyEventByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyEventByStampTaskResultFromDict(dict)
}

func NewVerifyEventByStampTaskResultFromDict(data map[string]interface{}) VerifyEventByStampTaskResult {
	return VerifyEventByStampTaskResult{
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p VerifyEventByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p VerifyEventByStampTaskResult) Pointer() *VerifyEventByStampTaskResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentEventMaster `json:"item"`
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
		Item: NewCurrentEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentEventMasterResult struct {
	Item *CurrentEventMaster `json:"item"`
}

type GetCurrentEventMasterAsyncResult struct {
	result *GetCurrentEventMasterResult
	err    error
}

func NewGetCurrentEventMasterResultFromJson(data string) GetCurrentEventMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentEventMasterResultFromDict(dict)
}

func NewGetCurrentEventMasterResultFromDict(data map[string]interface{}) GetCurrentEventMasterResult {
	return GetCurrentEventMasterResult{
		Item: NewCurrentEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentEventMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentEventMasterResult) Pointer() *GetCurrentEventMasterResult {
	return &p
}

type UpdateCurrentEventMasterResult struct {
	Item *CurrentEventMaster `json:"item"`
}

type UpdateCurrentEventMasterAsyncResult struct {
	result *UpdateCurrentEventMasterResult
	err    error
}

func NewUpdateCurrentEventMasterResultFromJson(data string) UpdateCurrentEventMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentEventMasterResultFromDict(dict)
}

func NewUpdateCurrentEventMasterResultFromDict(data map[string]interface{}) UpdateCurrentEventMasterResult {
	return UpdateCurrentEventMasterResult{
		Item: NewCurrentEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentEventMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentEventMasterResult) Pointer() *UpdateCurrentEventMasterResult {
	return &p
}

type UpdateCurrentEventMasterFromGitHubResult struct {
	Item *CurrentEventMaster `json:"item"`
}

type UpdateCurrentEventMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentEventMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentEventMasterFromGitHubResultFromJson(data string) UpdateCurrentEventMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentEventMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentEventMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentEventMasterFromGitHubResult {
	return UpdateCurrentEventMasterFromGitHubResult{
		Item: NewCurrentEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentEventMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentEventMasterFromGitHubResult) Pointer() *UpdateCurrentEventMasterFromGitHubResult {
	return &p
}
