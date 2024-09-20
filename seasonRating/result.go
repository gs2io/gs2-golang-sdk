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

package seasonRating

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

type DescribeMatchSessionsResult struct {
	Items         []MatchSession `json:"items"`
	NextPageToken *string        `json:"nextPageToken"`
}

type DescribeMatchSessionsAsyncResult struct {
	result *DescribeMatchSessionsResult
	err    error
}

func NewDescribeMatchSessionsResultFromJson(data string) DescribeMatchSessionsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMatchSessionsResultFromDict(dict)
}

func NewDescribeMatchSessionsResultFromDict(data map[string]interface{}) DescribeMatchSessionsResult {
	return DescribeMatchSessionsResult{
		Items:         CastMatchSessions(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeMatchSessionsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMatchSessionsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeMatchSessionsResult) Pointer() *DescribeMatchSessionsResult {
	return &p
}

type CreateMatchSessionResult struct {
	Item *MatchSession `json:"item"`
}

type CreateMatchSessionAsyncResult struct {
	result *CreateMatchSessionResult
	err    error
}

func NewCreateMatchSessionResultFromJson(data string) CreateMatchSessionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMatchSessionResultFromDict(dict)
}

func NewCreateMatchSessionResultFromDict(data map[string]interface{}) CreateMatchSessionResult {
	return CreateMatchSessionResult{
		Item: NewMatchSessionFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateMatchSessionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateMatchSessionResult) Pointer() *CreateMatchSessionResult {
	return &p
}

type GetMatchSessionResult struct {
	Item *MatchSession `json:"item"`
}

type GetMatchSessionAsyncResult struct {
	result *GetMatchSessionResult
	err    error
}

func NewGetMatchSessionResultFromJson(data string) GetMatchSessionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMatchSessionResultFromDict(dict)
}

func NewGetMatchSessionResultFromDict(data map[string]interface{}) GetMatchSessionResult {
	return GetMatchSessionResult{
		Item: NewMatchSessionFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMatchSessionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetMatchSessionResult) Pointer() *GetMatchSessionResult {
	return &p
}

type DeleteMatchSessionResult struct {
	Item *MatchSession `json:"item"`
}

type DeleteMatchSessionAsyncResult struct {
	result *DeleteMatchSessionResult
	err    error
}

func NewDeleteMatchSessionResultFromJson(data string) DeleteMatchSessionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMatchSessionResultFromDict(dict)
}

func NewDeleteMatchSessionResultFromDict(data map[string]interface{}) DeleteMatchSessionResult {
	return DeleteMatchSessionResult{
		Item: NewMatchSessionFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteMatchSessionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteMatchSessionResult) Pointer() *DeleteMatchSessionResult {
	return &p
}

type DescribeSeasonModelMastersResult struct {
	Items         []SeasonModelMaster `json:"items"`
	NextPageToken *string             `json:"nextPageToken"`
}

type DescribeSeasonModelMastersAsyncResult struct {
	result *DescribeSeasonModelMastersResult
	err    error
}

func NewDescribeSeasonModelMastersResultFromJson(data string) DescribeSeasonModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSeasonModelMastersResultFromDict(dict)
}

func NewDescribeSeasonModelMastersResultFromDict(data map[string]interface{}) DescribeSeasonModelMastersResult {
	return DescribeSeasonModelMastersResult{
		Items:         CastSeasonModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeSeasonModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSeasonModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeSeasonModelMastersResult) Pointer() *DescribeSeasonModelMastersResult {
	return &p
}

type CreateSeasonModelMasterResult struct {
	Item *SeasonModelMaster `json:"item"`
}

type CreateSeasonModelMasterAsyncResult struct {
	result *CreateSeasonModelMasterResult
	err    error
}

func NewCreateSeasonModelMasterResultFromJson(data string) CreateSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSeasonModelMasterResultFromDict(dict)
}

func NewCreateSeasonModelMasterResultFromDict(data map[string]interface{}) CreateSeasonModelMasterResult {
	return CreateSeasonModelMasterResult{
		Item: NewSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateSeasonModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateSeasonModelMasterResult) Pointer() *CreateSeasonModelMasterResult {
	return &p
}

type GetSeasonModelMasterResult struct {
	Item *SeasonModelMaster `json:"item"`
}

type GetSeasonModelMasterAsyncResult struct {
	result *GetSeasonModelMasterResult
	err    error
}

func NewGetSeasonModelMasterResultFromJson(data string) GetSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSeasonModelMasterResultFromDict(dict)
}

func NewGetSeasonModelMasterResultFromDict(data map[string]interface{}) GetSeasonModelMasterResult {
	return GetSeasonModelMasterResult{
		Item: NewSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetSeasonModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSeasonModelMasterResult) Pointer() *GetSeasonModelMasterResult {
	return &p
}

type UpdateSeasonModelMasterResult struct {
	Item *SeasonModelMaster `json:"item"`
}

type UpdateSeasonModelMasterAsyncResult struct {
	result *UpdateSeasonModelMasterResult
	err    error
}

func NewUpdateSeasonModelMasterResultFromJson(data string) UpdateSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSeasonModelMasterResultFromDict(dict)
}

func NewUpdateSeasonModelMasterResultFromDict(data map[string]interface{}) UpdateSeasonModelMasterResult {
	return UpdateSeasonModelMasterResult{
		Item: NewSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateSeasonModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateSeasonModelMasterResult) Pointer() *UpdateSeasonModelMasterResult {
	return &p
}

type DeleteSeasonModelMasterResult struct {
	Item *SeasonModelMaster `json:"item"`
}

type DeleteSeasonModelMasterAsyncResult struct {
	result *DeleteSeasonModelMasterResult
	err    error
}

func NewDeleteSeasonModelMasterResultFromJson(data string) DeleteSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSeasonModelMasterResultFromDict(dict)
}

func NewDeleteSeasonModelMasterResultFromDict(data map[string]interface{}) DeleteSeasonModelMasterResult {
	return DeleteSeasonModelMasterResult{
		Item: NewSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteSeasonModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteSeasonModelMasterResult) Pointer() *DeleteSeasonModelMasterResult {
	return &p
}

type DescribeSeasonModelsResult struct {
	Items []SeasonModel `json:"items"`
}

type DescribeSeasonModelsAsyncResult struct {
	result *DescribeSeasonModelsResult
	err    error
}

func NewDescribeSeasonModelsResultFromJson(data string) DescribeSeasonModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSeasonModelsResultFromDict(dict)
}

func NewDescribeSeasonModelsResultFromDict(data map[string]interface{}) DescribeSeasonModelsResult {
	return DescribeSeasonModelsResult{
		Items: CastSeasonModels(core.CastArray(data["items"])),
	}
}

func (p DescribeSeasonModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSeasonModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeSeasonModelsResult) Pointer() *DescribeSeasonModelsResult {
	return &p
}

type GetSeasonModelResult struct {
	Item *SeasonModel `json:"item"`
}

type GetSeasonModelAsyncResult struct {
	result *GetSeasonModelResult
	err    error
}

func NewGetSeasonModelResultFromJson(data string) GetSeasonModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSeasonModelResultFromDict(dict)
}

func NewGetSeasonModelResultFromDict(data map[string]interface{}) GetSeasonModelResult {
	return GetSeasonModelResult{
		Item: NewSeasonModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetSeasonModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSeasonModelResult) Pointer() *GetSeasonModelResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentSeasonModelMaster `json:"item"`
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
		Item: NewCurrentSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentSeasonModelMasterResult struct {
	Item *CurrentSeasonModelMaster `json:"item"`
}

type GetCurrentSeasonModelMasterAsyncResult struct {
	result *GetCurrentSeasonModelMasterResult
	err    error
}

func NewGetCurrentSeasonModelMasterResultFromJson(data string) GetCurrentSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentSeasonModelMasterResultFromDict(dict)
}

func NewGetCurrentSeasonModelMasterResultFromDict(data map[string]interface{}) GetCurrentSeasonModelMasterResult {
	return GetCurrentSeasonModelMasterResult{
		Item: NewCurrentSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentSeasonModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentSeasonModelMasterResult) Pointer() *GetCurrentSeasonModelMasterResult {
	return &p
}

type UpdateCurrentSeasonModelMasterResult struct {
	Item *CurrentSeasonModelMaster `json:"item"`
}

type UpdateCurrentSeasonModelMasterAsyncResult struct {
	result *UpdateCurrentSeasonModelMasterResult
	err    error
}

func NewUpdateCurrentSeasonModelMasterResultFromJson(data string) UpdateCurrentSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentSeasonModelMasterResultFromDict(dict)
}

func NewUpdateCurrentSeasonModelMasterResultFromDict(data map[string]interface{}) UpdateCurrentSeasonModelMasterResult {
	return UpdateCurrentSeasonModelMasterResult{
		Item: NewCurrentSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentSeasonModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentSeasonModelMasterResult) Pointer() *UpdateCurrentSeasonModelMasterResult {
	return &p
}

type UpdateCurrentSeasonModelMasterFromGitHubResult struct {
	Item *CurrentSeasonModelMaster `json:"item"`
}

type UpdateCurrentSeasonModelMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentSeasonModelMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentSeasonModelMasterFromGitHubResultFromJson(data string) UpdateCurrentSeasonModelMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentSeasonModelMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentSeasonModelMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentSeasonModelMasterFromGitHubResult {
	return UpdateCurrentSeasonModelMasterFromGitHubResult{
		Item: NewCurrentSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentSeasonModelMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentSeasonModelMasterFromGitHubResult) Pointer() *UpdateCurrentSeasonModelMasterFromGitHubResult {
	return &p
}

type GetBallotResult struct {
	Item      *Ballot `json:"item"`
	Body      *string `json:"body"`
	Signature *string `json:"signature"`
}

type GetBallotAsyncResult struct {
	result *GetBallotResult
	err    error
}

func NewGetBallotResultFromJson(data string) GetBallotResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBallotResultFromDict(dict)
}

func NewGetBallotResultFromDict(data map[string]interface{}) GetBallotResult {
	return GetBallotResult{
		Item:      NewBallotFromDict(core.CastMap(data["item"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p GetBallotResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p GetBallotResult) Pointer() *GetBallotResult {
	return &p
}

type GetBallotByUserIdResult struct {
	Item      *Ballot `json:"item"`
	Body      *string `json:"body"`
	Signature *string `json:"signature"`
}

type GetBallotByUserIdAsyncResult struct {
	result *GetBallotByUserIdResult
	err    error
}

func NewGetBallotByUserIdResultFromJson(data string) GetBallotByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBallotByUserIdResultFromDict(dict)
}

func NewGetBallotByUserIdResultFromDict(data map[string]interface{}) GetBallotByUserIdResult {
	return GetBallotByUserIdResult{
		Item:      NewBallotFromDict(core.CastMap(data["item"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p GetBallotByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p GetBallotByUserIdResult) Pointer() *GetBallotByUserIdResult {
	return &p
}

type VoteResult struct {
	Item *Ballot `json:"item"`
}

type VoteAsyncResult struct {
	result *VoteResult
	err    error
}

func NewVoteResultFromJson(data string) VoteResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVoteResultFromDict(dict)
}

func NewVoteResultFromDict(data map[string]interface{}) VoteResult {
	return VoteResult{
		Item: NewBallotFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p VoteResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VoteResult) Pointer() *VoteResult {
	return &p
}

type VoteMultipleResult struct {
	Item *Ballot `json:"item"`
}

type VoteMultipleAsyncResult struct {
	result *VoteMultipleResult
	err    error
}

func NewVoteMultipleResultFromJson(data string) VoteMultipleResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVoteMultipleResultFromDict(dict)
}

func NewVoteMultipleResultFromDict(data map[string]interface{}) VoteMultipleResult {
	return VoteMultipleResult{
		Item: NewBallotFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p VoteMultipleResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VoteMultipleResult) Pointer() *VoteMultipleResult {
	return &p
}

type CommitVoteResult struct {
}

type CommitVoteAsyncResult struct {
	result *CommitVoteResult
	err    error
}

func NewCommitVoteResultFromJson(data string) CommitVoteResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCommitVoteResultFromDict(dict)
}

func NewCommitVoteResultFromDict(data map[string]interface{}) CommitVoteResult {
	return CommitVoteResult{}
}

func (p CommitVoteResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p CommitVoteResult) Pointer() *CommitVoteResult {
	return &p
}
