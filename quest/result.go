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

package quest

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
	Items         []Namespace          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []Namespace {
			if data["items"] == nil {
				return nil
			}
			return CastNamespaces(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Status   *string              `json:"status"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Url      *string              `json:"url"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Url: func() *string {
			v, ok := data["url"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["url"])
		}(),
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
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
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
		UploadUrl: func() *string {
			v, ok := data["uploadUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadUrl"])
		}(),
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Url      *string              `json:"url"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Url: func() *string {
			v, ok := data["url"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["url"])
		}(),
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

type DescribeQuestGroupModelMastersResult struct {
	Items         []QuestGroupModelMaster `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
	Metadata      *core.ResultMetadata    `json:"metadata"`
}

type DescribeQuestGroupModelMastersAsyncResult struct {
	result *DescribeQuestGroupModelMastersResult
	err    error
}

func NewDescribeQuestGroupModelMastersResultFromJson(data string) DescribeQuestGroupModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeQuestGroupModelMastersResultFromDict(dict)
}

func NewDescribeQuestGroupModelMastersResultFromDict(data map[string]interface{}) DescribeQuestGroupModelMastersResult {
	return DescribeQuestGroupModelMastersResult{
		Items: func() []QuestGroupModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastQuestGroupModelMasters(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeQuestGroupModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastQuestGroupModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeQuestGroupModelMastersResult) Pointer() *DescribeQuestGroupModelMastersResult {
	return &p
}

type CreateQuestGroupModelMasterResult struct {
	Item     *QuestGroupModelMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type CreateQuestGroupModelMasterAsyncResult struct {
	result *CreateQuestGroupModelMasterResult
	err    error
}

func NewCreateQuestGroupModelMasterResultFromJson(data string) CreateQuestGroupModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateQuestGroupModelMasterResultFromDict(dict)
}

func NewCreateQuestGroupModelMasterResultFromDict(data map[string]interface{}) CreateQuestGroupModelMasterResult {
	return CreateQuestGroupModelMasterResult{
		Item: func() *QuestGroupModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateQuestGroupModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateQuestGroupModelMasterResult) Pointer() *CreateQuestGroupModelMasterResult {
	return &p
}

type GetQuestGroupModelMasterResult struct {
	Item     *QuestGroupModelMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetQuestGroupModelMasterAsyncResult struct {
	result *GetQuestGroupModelMasterResult
	err    error
}

func NewGetQuestGroupModelMasterResultFromJson(data string) GetQuestGroupModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetQuestGroupModelMasterResultFromDict(dict)
}

func NewGetQuestGroupModelMasterResultFromDict(data map[string]interface{}) GetQuestGroupModelMasterResult {
	return GetQuestGroupModelMasterResult{
		Item: func() *QuestGroupModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetQuestGroupModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetQuestGroupModelMasterResult) Pointer() *GetQuestGroupModelMasterResult {
	return &p
}

type UpdateQuestGroupModelMasterResult struct {
	Item     *QuestGroupModelMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type UpdateQuestGroupModelMasterAsyncResult struct {
	result *UpdateQuestGroupModelMasterResult
	err    error
}

func NewUpdateQuestGroupModelMasterResultFromJson(data string) UpdateQuestGroupModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateQuestGroupModelMasterResultFromDict(dict)
}

func NewUpdateQuestGroupModelMasterResultFromDict(data map[string]interface{}) UpdateQuestGroupModelMasterResult {
	return UpdateQuestGroupModelMasterResult{
		Item: func() *QuestGroupModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateQuestGroupModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateQuestGroupModelMasterResult) Pointer() *UpdateQuestGroupModelMasterResult {
	return &p
}

type DeleteQuestGroupModelMasterResult struct {
	Item     *QuestGroupModelMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type DeleteQuestGroupModelMasterAsyncResult struct {
	result *DeleteQuestGroupModelMasterResult
	err    error
}

func NewDeleteQuestGroupModelMasterResultFromJson(data string) DeleteQuestGroupModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteQuestGroupModelMasterResultFromDict(dict)
}

func NewDeleteQuestGroupModelMasterResultFromDict(data map[string]interface{}) DeleteQuestGroupModelMasterResult {
	return DeleteQuestGroupModelMasterResult{
		Item: func() *QuestGroupModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteQuestGroupModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteQuestGroupModelMasterResult) Pointer() *DeleteQuestGroupModelMasterResult {
	return &p
}

type DescribeQuestModelMastersResult struct {
	Items         []QuestModelMaster   `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeQuestModelMastersAsyncResult struct {
	result *DescribeQuestModelMastersResult
	err    error
}

func NewDescribeQuestModelMastersResultFromJson(data string) DescribeQuestModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeQuestModelMastersResultFromDict(dict)
}

func NewDescribeQuestModelMastersResultFromDict(data map[string]interface{}) DescribeQuestModelMastersResult {
	return DescribeQuestModelMastersResult{
		Items: func() []QuestModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastQuestModelMasters(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeQuestModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastQuestModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeQuestModelMastersResult) Pointer() *DescribeQuestModelMastersResult {
	return &p
}

type CreateQuestModelMasterResult struct {
	Item     *QuestModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateQuestModelMasterAsyncResult struct {
	result *CreateQuestModelMasterResult
	err    error
}

func NewCreateQuestModelMasterResultFromJson(data string) CreateQuestModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateQuestModelMasterResultFromDict(dict)
}

func NewCreateQuestModelMasterResultFromDict(data map[string]interface{}) CreateQuestModelMasterResult {
	return CreateQuestModelMasterResult{
		Item: func() *QuestModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateQuestModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateQuestModelMasterResult) Pointer() *CreateQuestModelMasterResult {
	return &p
}

type GetQuestModelMasterResult struct {
	Item     *QuestModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetQuestModelMasterAsyncResult struct {
	result *GetQuestModelMasterResult
	err    error
}

func NewGetQuestModelMasterResultFromJson(data string) GetQuestModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetQuestModelMasterResultFromDict(dict)
}

func NewGetQuestModelMasterResultFromDict(data map[string]interface{}) GetQuestModelMasterResult {
	return GetQuestModelMasterResult{
		Item: func() *QuestModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetQuestModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetQuestModelMasterResult) Pointer() *GetQuestModelMasterResult {
	return &p
}

type UpdateQuestModelMasterResult struct {
	Item     *QuestModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateQuestModelMasterAsyncResult struct {
	result *UpdateQuestModelMasterResult
	err    error
}

func NewUpdateQuestModelMasterResultFromJson(data string) UpdateQuestModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateQuestModelMasterResultFromDict(dict)
}

func NewUpdateQuestModelMasterResultFromDict(data map[string]interface{}) UpdateQuestModelMasterResult {
	return UpdateQuestModelMasterResult{
		Item: func() *QuestModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateQuestModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateQuestModelMasterResult) Pointer() *UpdateQuestModelMasterResult {
	return &p
}

type DeleteQuestModelMasterResult struct {
	Item     *QuestModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteQuestModelMasterAsyncResult struct {
	result *DeleteQuestModelMasterResult
	err    error
}

func NewDeleteQuestModelMasterResultFromJson(data string) DeleteQuestModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteQuestModelMasterResultFromDict(dict)
}

func NewDeleteQuestModelMasterResultFromDict(data map[string]interface{}) DeleteQuestModelMasterResult {
	return DeleteQuestModelMasterResult{
		Item: func() *QuestModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteQuestModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteQuestModelMasterResult) Pointer() *DeleteQuestModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentQuestMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *CurrentQuestMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentQuestMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type GetCurrentQuestMasterResult struct {
	Item     *CurrentQuestMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCurrentQuestMasterAsyncResult struct {
	result *GetCurrentQuestMasterResult
	err    error
}

func NewGetCurrentQuestMasterResultFromJson(data string) GetCurrentQuestMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentQuestMasterResultFromDict(dict)
}

func NewGetCurrentQuestMasterResultFromDict(data map[string]interface{}) GetCurrentQuestMasterResult {
	return GetCurrentQuestMasterResult{
		Item: func() *CurrentQuestMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentQuestMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCurrentQuestMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentQuestMasterResult) Pointer() *GetCurrentQuestMasterResult {
	return &p
}

type UpdateCurrentQuestMasterResult struct {
	Item     *CurrentQuestMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentQuestMasterAsyncResult struct {
	result *UpdateCurrentQuestMasterResult
	err    error
}

func NewUpdateCurrentQuestMasterResultFromJson(data string) UpdateCurrentQuestMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentQuestMasterResultFromDict(dict)
}

func NewUpdateCurrentQuestMasterResultFromDict(data map[string]interface{}) UpdateCurrentQuestMasterResult {
	return UpdateCurrentQuestMasterResult{
		Item: func() *CurrentQuestMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentQuestMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentQuestMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentQuestMasterResult) Pointer() *UpdateCurrentQuestMasterResult {
	return &p
}

type UpdateCurrentQuestMasterFromGitHubResult struct {
	Item     *CurrentQuestMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentQuestMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentQuestMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentQuestMasterFromGitHubResultFromJson(data string) UpdateCurrentQuestMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentQuestMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentQuestMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentQuestMasterFromGitHubResult {
	return UpdateCurrentQuestMasterFromGitHubResult{
		Item: func() *CurrentQuestMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentQuestMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentQuestMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentQuestMasterFromGitHubResult) Pointer() *UpdateCurrentQuestMasterFromGitHubResult {
	return &p
}

type DescribeProgressesByUserIdResult struct {
	Items         []Progress           `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []Progress {
			if data["items"] == nil {
				return nil
			}
			return CastProgresses(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
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
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateProgressByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateProgressByUserIdResult) Pointer() *CreateProgressByUserIdResult {
	return &p
}

type GetProgressResult struct {
	Item       *Progress            `json:"item"`
	QuestGroup *QuestGroupModel     `json:"questGroup"`
	Quest      *QuestModel          `json:"quest"`
	Metadata   *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		QuestGroup: func() *QuestGroupModel {
			v, ok := data["questGroup"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestGroupModelFromDict(core.CastMap(data["questGroup"])).Pointer()
		}(),
		Quest: func() *QuestModel {
			v, ok := data["quest"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestModelFromDict(core.CastMap(data["quest"])).Pointer()
		}(),
	}
}

func (p GetProgressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"questGroup": func() map[string]interface{} {
			if p.QuestGroup == nil {
				return nil
			}
			return p.QuestGroup.ToDict()
		}(),
		"quest": func() map[string]interface{} {
			if p.Quest == nil {
				return nil
			}
			return p.Quest.ToDict()
		}(),
	}
}

func (p GetProgressResult) Pointer() *GetProgressResult {
	return &p
}

type GetProgressByUserIdResult struct {
	Item       *Progress            `json:"item"`
	QuestGroup *QuestGroupModel     `json:"questGroup"`
	Quest      *QuestModel          `json:"quest"`
	Metadata   *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		QuestGroup: func() *QuestGroupModel {
			v, ok := data["questGroup"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestGroupModelFromDict(core.CastMap(data["questGroup"])).Pointer()
		}(),
		Quest: func() *QuestModel {
			v, ok := data["quest"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestModelFromDict(core.CastMap(data["quest"])).Pointer()
		}(),
	}
}

func (p GetProgressByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"questGroup": func() map[string]interface{} {
			if p.QuestGroup == nil {
				return nil
			}
			return p.QuestGroup.ToDict()
		}(),
		"quest": func() map[string]interface{} {
			if p.Quest == nil {
				return nil
			}
			return p.Quest.ToDict()
		}(),
	}
}

func (p GetProgressByUserIdResult) Pointer() *GetProgressByUserIdResult {
	return &p
}

type StartResult struct {
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
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
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
	}
}

func (p StartResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
	}
}

func (p StartResult) Pointer() *StartResult {
	return &p
}

type StartByUserIdResult struct {
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
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
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
	}
}

func (p StartByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
	}
}

func (p StartByUserIdResult) Pointer() *StartByUserIdResult {
	return &p
}

type EndResult struct {
	Item                      *Progress               `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
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
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
	}
}

func (p EndResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
	}
}

func (p EndResult) Pointer() *EndResult {
	return &p
}

type EndByUserIdResult struct {
	Item                      *Progress               `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
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
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
	}
}

func (p EndByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
	}
}

func (p EndByUserIdResult) Pointer() *EndByUserIdResult {
	return &p
}

type DeleteProgressResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteProgressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteProgressResult) Pointer() *DeleteProgressResult {
	return &p
}

type DeleteProgressByUserIdResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteProgressByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteProgressByUserIdResult) Pointer() *DeleteProgressByUserIdResult {
	return &p
}

type CreateProgressByStampSheetResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateProgressByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateProgressByStampSheetResult) Pointer() *CreateProgressByStampSheetResult {
	return &p
}

type DeleteProgressByStampTaskResult struct {
	Item            *Progress            `json:"item"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p DeleteProgressByStampTaskResult) ToDict() map[string]interface{} {
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

func (p DeleteProgressByStampTaskResult) Pointer() *DeleteProgressByStampTaskResult {
	return &p
}

type DescribeCompletedQuestListsResult struct {
	Items         []CompletedQuestList `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeCompletedQuestListsAsyncResult struct {
	result *DescribeCompletedQuestListsResult
	err    error
}

func NewDescribeCompletedQuestListsResultFromJson(data string) DescribeCompletedQuestListsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCompletedQuestListsResultFromDict(dict)
}

func NewDescribeCompletedQuestListsResultFromDict(data map[string]interface{}) DescribeCompletedQuestListsResult {
	return DescribeCompletedQuestListsResult{
		Items: func() []CompletedQuestList {
			if data["items"] == nil {
				return nil
			}
			return CastCompletedQuestList(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeCompletedQuestListsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCompletedQuestListFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCompletedQuestListsResult) Pointer() *DescribeCompletedQuestListsResult {
	return &p
}

type DescribeCompletedQuestListsByUserIdResult struct {
	Items         []CompletedQuestList `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeCompletedQuestListsByUserIdAsyncResult struct {
	result *DescribeCompletedQuestListsByUserIdResult
	err    error
}

func NewDescribeCompletedQuestListsByUserIdResultFromJson(data string) DescribeCompletedQuestListsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCompletedQuestListsByUserIdResultFromDict(dict)
}

func NewDescribeCompletedQuestListsByUserIdResultFromDict(data map[string]interface{}) DescribeCompletedQuestListsByUserIdResult {
	return DescribeCompletedQuestListsByUserIdResult{
		Items: func() []CompletedQuestList {
			if data["items"] == nil {
				return nil
			}
			return CastCompletedQuestList(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeCompletedQuestListsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCompletedQuestListFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCompletedQuestListsByUserIdResult) Pointer() *DescribeCompletedQuestListsByUserIdResult {
	return &p
}

type GetCompletedQuestListResult struct {
	Item     *CompletedQuestList  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCompletedQuestListAsyncResult struct {
	result *GetCompletedQuestListResult
	err    error
}

func NewGetCompletedQuestListResultFromJson(data string) GetCompletedQuestListResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCompletedQuestListResultFromDict(dict)
}

func NewGetCompletedQuestListResultFromDict(data map[string]interface{}) GetCompletedQuestListResult {
	return GetCompletedQuestListResult{
		Item: func() *CompletedQuestList {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCompletedQuestListFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCompletedQuestListResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCompletedQuestListResult) Pointer() *GetCompletedQuestListResult {
	return &p
}

type GetCompletedQuestListByUserIdResult struct {
	Item     *CompletedQuestList  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCompletedQuestListByUserIdAsyncResult struct {
	result *GetCompletedQuestListByUserIdResult
	err    error
}

func NewGetCompletedQuestListByUserIdResultFromJson(data string) GetCompletedQuestListByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCompletedQuestListByUserIdResultFromDict(dict)
}

func NewGetCompletedQuestListByUserIdResultFromDict(data map[string]interface{}) GetCompletedQuestListByUserIdResult {
	return GetCompletedQuestListByUserIdResult{
		Item: func() *CompletedQuestList {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCompletedQuestListFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCompletedQuestListByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCompletedQuestListByUserIdResult) Pointer() *GetCompletedQuestListByUserIdResult {
	return &p
}

type DeleteCompletedQuestListByUserIdResult struct {
	Item     *CompletedQuestList  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteCompletedQuestListByUserIdAsyncResult struct {
	result *DeleteCompletedQuestListByUserIdResult
	err    error
}

func NewDeleteCompletedQuestListByUserIdResultFromJson(data string) DeleteCompletedQuestListByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCompletedQuestListByUserIdResultFromDict(dict)
}

func NewDeleteCompletedQuestListByUserIdResultFromDict(data map[string]interface{}) DeleteCompletedQuestListByUserIdResult {
	return DeleteCompletedQuestListByUserIdResult{
		Item: func() *CompletedQuestList {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCompletedQuestListFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteCompletedQuestListByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteCompletedQuestListByUserIdResult) Pointer() *DeleteCompletedQuestListByUserIdResult {
	return &p
}

type DescribeQuestGroupModelsResult struct {
	Items    []QuestGroupModel    `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeQuestGroupModelsAsyncResult struct {
	result *DescribeQuestGroupModelsResult
	err    error
}

func NewDescribeQuestGroupModelsResultFromJson(data string) DescribeQuestGroupModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeQuestGroupModelsResultFromDict(dict)
}

func NewDescribeQuestGroupModelsResultFromDict(data map[string]interface{}) DescribeQuestGroupModelsResult {
	return DescribeQuestGroupModelsResult{
		Items: func() []QuestGroupModel {
			if data["items"] == nil {
				return nil
			}
			return CastQuestGroupModels(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeQuestGroupModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastQuestGroupModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeQuestGroupModelsResult) Pointer() *DescribeQuestGroupModelsResult {
	return &p
}

type GetQuestGroupModelResult struct {
	Item     *QuestGroupModel     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetQuestGroupModelAsyncResult struct {
	result *GetQuestGroupModelResult
	err    error
}

func NewGetQuestGroupModelResultFromJson(data string) GetQuestGroupModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetQuestGroupModelResultFromDict(dict)
}

func NewGetQuestGroupModelResultFromDict(data map[string]interface{}) GetQuestGroupModelResult {
	return GetQuestGroupModelResult{
		Item: func() *QuestGroupModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestGroupModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetQuestGroupModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetQuestGroupModelResult) Pointer() *GetQuestGroupModelResult {
	return &p
}

type DescribeQuestModelsResult struct {
	Items    []QuestModel         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeQuestModelsAsyncResult struct {
	result *DescribeQuestModelsResult
	err    error
}

func NewDescribeQuestModelsResultFromJson(data string) DescribeQuestModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeQuestModelsResultFromDict(dict)
}

func NewDescribeQuestModelsResultFromDict(data map[string]interface{}) DescribeQuestModelsResult {
	return DescribeQuestModelsResult{
		Items: func() []QuestModel {
			if data["items"] == nil {
				return nil
			}
			return CastQuestModels(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeQuestModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastQuestModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeQuestModelsResult) Pointer() *DescribeQuestModelsResult {
	return &p
}

type GetQuestModelResult struct {
	Item     *QuestModel          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetQuestModelAsyncResult struct {
	result *GetQuestModelResult
	err    error
}

func NewGetQuestModelResultFromJson(data string) GetQuestModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetQuestModelResultFromDict(dict)
}

func NewGetQuestModelResultFromDict(data map[string]interface{}) GetQuestModelResult {
	return GetQuestModelResult{
		Item: func() *QuestModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewQuestModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetQuestModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetQuestModelResult) Pointer() *GetQuestModelResult {
	return &p
}
