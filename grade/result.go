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

package grade

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
	"github.com/gs2io/gs2-golang-sdk/experience"
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

type DescribeGradeModelMastersResult struct {
	Items         []GradeModelMaster `json:"items"`
	NextPageToken *string            `json:"nextPageToken"`
}

type DescribeGradeModelMastersAsyncResult struct {
	result *DescribeGradeModelMastersResult
	err    error
}

func NewDescribeGradeModelMastersResultFromJson(data string) DescribeGradeModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGradeModelMastersResultFromDict(dict)
}

func NewDescribeGradeModelMastersResultFromDict(data map[string]interface{}) DescribeGradeModelMastersResult {
	return DescribeGradeModelMastersResult{
		Items:         CastGradeModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeGradeModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGradeModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGradeModelMastersResult) Pointer() *DescribeGradeModelMastersResult {
	return &p
}

type CreateGradeModelMasterResult struct {
	Item *GradeModelMaster `json:"item"`
}

type CreateGradeModelMasterAsyncResult struct {
	result *CreateGradeModelMasterResult
	err    error
}

func NewCreateGradeModelMasterResultFromJson(data string) CreateGradeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGradeModelMasterResultFromDict(dict)
}

func NewCreateGradeModelMasterResultFromDict(data map[string]interface{}) CreateGradeModelMasterResult {
	return CreateGradeModelMasterResult{
		Item: NewGradeModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateGradeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateGradeModelMasterResult) Pointer() *CreateGradeModelMasterResult {
	return &p
}

type GetGradeModelMasterResult struct {
	Item *GradeModelMaster `json:"item"`
}

type GetGradeModelMasterAsyncResult struct {
	result *GetGradeModelMasterResult
	err    error
}

func NewGetGradeModelMasterResultFromJson(data string) GetGradeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGradeModelMasterResultFromDict(dict)
}

func NewGetGradeModelMasterResultFromDict(data map[string]interface{}) GetGradeModelMasterResult {
	return GetGradeModelMasterResult{
		Item: NewGradeModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetGradeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetGradeModelMasterResult) Pointer() *GetGradeModelMasterResult {
	return &p
}

type UpdateGradeModelMasterResult struct {
	Item *GradeModelMaster `json:"item"`
}

type UpdateGradeModelMasterAsyncResult struct {
	result *UpdateGradeModelMasterResult
	err    error
}

func NewUpdateGradeModelMasterResultFromJson(data string) UpdateGradeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGradeModelMasterResultFromDict(dict)
}

func NewUpdateGradeModelMasterResultFromDict(data map[string]interface{}) UpdateGradeModelMasterResult {
	return UpdateGradeModelMasterResult{
		Item: NewGradeModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateGradeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateGradeModelMasterResult) Pointer() *UpdateGradeModelMasterResult {
	return &p
}

type DeleteGradeModelMasterResult struct {
	Item *GradeModelMaster `json:"item"`
}

type DeleteGradeModelMasterAsyncResult struct {
	result *DeleteGradeModelMasterResult
	err    error
}

func NewDeleteGradeModelMasterResultFromJson(data string) DeleteGradeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGradeModelMasterResultFromDict(dict)
}

func NewDeleteGradeModelMasterResultFromDict(data map[string]interface{}) DeleteGradeModelMasterResult {
	return DeleteGradeModelMasterResult{
		Item: NewGradeModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteGradeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteGradeModelMasterResult) Pointer() *DeleteGradeModelMasterResult {
	return &p
}

type DescribeGradeModelsResult struct {
	Items []GradeModel `json:"items"`
}

type DescribeGradeModelsAsyncResult struct {
	result *DescribeGradeModelsResult
	err    error
}

func NewDescribeGradeModelsResultFromJson(data string) DescribeGradeModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGradeModelsResultFromDict(dict)
}

func NewDescribeGradeModelsResultFromDict(data map[string]interface{}) DescribeGradeModelsResult {
	return DescribeGradeModelsResult{
		Items: CastGradeModels(core.CastArray(data["items"])),
	}
}

func (p DescribeGradeModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGradeModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeGradeModelsResult) Pointer() *DescribeGradeModelsResult {
	return &p
}

type GetGradeModelResult struct {
	Item *GradeModel `json:"item"`
}

type GetGradeModelAsyncResult struct {
	result *GetGradeModelResult
	err    error
}

func NewGetGradeModelResultFromJson(data string) GetGradeModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGradeModelResultFromDict(dict)
}

func NewGetGradeModelResultFromDict(data map[string]interface{}) GetGradeModelResult {
	return GetGradeModelResult{
		Item: NewGradeModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetGradeModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetGradeModelResult) Pointer() *GetGradeModelResult {
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

type AddGradeByUserIdResult struct {
	Item                    *Status            `json:"item"`
	ExperienceNamespaceName *string            `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status `json:"experienceStatus"`
}

type AddGradeByUserIdAsyncResult struct {
	result *AddGradeByUserIdResult
	err    error
}

func NewAddGradeByUserIdResultFromJson(data string) AddGradeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddGradeByUserIdResultFromDict(dict)
}

func NewAddGradeByUserIdResultFromDict(data map[string]interface{}) AddGradeByUserIdResult {
	return AddGradeByUserIdResult{
		Item:                    NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		ExperienceNamespaceName: core.CastString(data["experienceNamespaceName"]),
		ExperienceStatus:        experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer(),
	}
}

func (p AddGradeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus":        p.ExperienceStatus.ToDict(),
	}
}

func (p AddGradeByUserIdResult) Pointer() *AddGradeByUserIdResult {
	return &p
}

type SubGradeResult struct {
	Item                    *Status            `json:"item"`
	ExperienceNamespaceName *string            `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status `json:"experienceStatus"`
}

type SubGradeAsyncResult struct {
	result *SubGradeResult
	err    error
}

func NewSubGradeResultFromJson(data string) SubGradeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubGradeResultFromDict(dict)
}

func NewSubGradeResultFromDict(data map[string]interface{}) SubGradeResult {
	return SubGradeResult{
		Item:                    NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		ExperienceNamespaceName: core.CastString(data["experienceNamespaceName"]),
		ExperienceStatus:        experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer(),
	}
}

func (p SubGradeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus":        p.ExperienceStatus.ToDict(),
	}
}

func (p SubGradeResult) Pointer() *SubGradeResult {
	return &p
}

type SubGradeByUserIdResult struct {
	Item                    *Status            `json:"item"`
	ExperienceNamespaceName *string            `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status `json:"experienceStatus"`
}

type SubGradeByUserIdAsyncResult struct {
	result *SubGradeByUserIdResult
	err    error
}

func NewSubGradeByUserIdResultFromJson(data string) SubGradeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubGradeByUserIdResultFromDict(dict)
}

func NewSubGradeByUserIdResultFromDict(data map[string]interface{}) SubGradeByUserIdResult {
	return SubGradeByUserIdResult{
		Item:                    NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		ExperienceNamespaceName: core.CastString(data["experienceNamespaceName"]),
		ExperienceStatus:        experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer(),
	}
}

func (p SubGradeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus":        p.ExperienceStatus.ToDict(),
	}
}

func (p SubGradeByUserIdResult) Pointer() *SubGradeByUserIdResult {
	return &p
}

type SetGradeByUserIdResult struct {
	Item                    *Status            `json:"item"`
	Old                     *Status            `json:"old"`
	ExperienceNamespaceName *string            `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status `json:"experienceStatus"`
}

type SetGradeByUserIdAsyncResult struct {
	result *SetGradeByUserIdResult
	err    error
}

func NewSetGradeByUserIdResultFromJson(data string) SetGradeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetGradeByUserIdResultFromDict(dict)
}

func NewSetGradeByUserIdResultFromDict(data map[string]interface{}) SetGradeByUserIdResult {
	return SetGradeByUserIdResult{
		Item:                    NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		Old:                     NewStatusFromDict(core.CastMap(data["old"])).Pointer(),
		ExperienceNamespaceName: core.CastString(data["experienceNamespaceName"]),
		ExperienceStatus:        experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer(),
	}
}

func (p SetGradeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"old":                     p.Old.ToDict(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus":        p.ExperienceStatus.ToDict(),
	}
}

func (p SetGradeByUserIdResult) Pointer() *SetGradeByUserIdResult {
	return &p
}

type ApplyRankCapResult struct {
	Item                    *Status            `json:"item"`
	ExperienceNamespaceName *string            `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status `json:"experienceStatus"`
}

type ApplyRankCapAsyncResult struct {
	result *ApplyRankCapResult
	err    error
}

func NewApplyRankCapResultFromJson(data string) ApplyRankCapResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyRankCapResultFromDict(dict)
}

func NewApplyRankCapResultFromDict(data map[string]interface{}) ApplyRankCapResult {
	return ApplyRankCapResult{
		Item:                    NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		ExperienceNamespaceName: core.CastString(data["experienceNamespaceName"]),
		ExperienceStatus:        experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer(),
	}
}

func (p ApplyRankCapResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus":        p.ExperienceStatus.ToDict(),
	}
}

func (p ApplyRankCapResult) Pointer() *ApplyRankCapResult {
	return &p
}

type ApplyRankCapByUserIdResult struct {
	Item                    *Status            `json:"item"`
	ExperienceNamespaceName *string            `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status `json:"experienceStatus"`
}

type ApplyRankCapByUserIdAsyncResult struct {
	result *ApplyRankCapByUserIdResult
	err    error
}

func NewApplyRankCapByUserIdResultFromJson(data string) ApplyRankCapByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyRankCapByUserIdResultFromDict(dict)
}

func NewApplyRankCapByUserIdResultFromDict(data map[string]interface{}) ApplyRankCapByUserIdResult {
	return ApplyRankCapByUserIdResult{
		Item:                    NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		ExperienceNamespaceName: core.CastString(data["experienceNamespaceName"]),
		ExperienceStatus:        experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer(),
	}
}

func (p ApplyRankCapByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus":        p.ExperienceStatus.ToDict(),
	}
}

func (p ApplyRankCapByUserIdResult) Pointer() *ApplyRankCapByUserIdResult {
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

type VerifyGradeResult struct {
}

type VerifyGradeAsyncResult struct {
	result *VerifyGradeResult
	err    error
}

func NewVerifyGradeResultFromJson(data string) VerifyGradeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeResultFromDict(dict)
}

func NewVerifyGradeResultFromDict(data map[string]interface{}) VerifyGradeResult {
	return VerifyGradeResult{}
}

func (p VerifyGradeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyGradeResult) Pointer() *VerifyGradeResult {
	return &p
}

type VerifyGradeByUserIdResult struct {
}

type VerifyGradeByUserIdAsyncResult struct {
	result *VerifyGradeByUserIdResult
	err    error
}

func NewVerifyGradeByUserIdResultFromJson(data string) VerifyGradeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeByUserIdResultFromDict(dict)
}

func NewVerifyGradeByUserIdResultFromDict(data map[string]interface{}) VerifyGradeByUserIdResult {
	return VerifyGradeByUserIdResult{}
}

func (p VerifyGradeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyGradeByUserIdResult) Pointer() *VerifyGradeByUserIdResult {
	return &p
}

type VerifyGradeUpMaterialResult struct {
}

type VerifyGradeUpMaterialAsyncResult struct {
	result *VerifyGradeUpMaterialResult
	err    error
}

func NewVerifyGradeUpMaterialResultFromJson(data string) VerifyGradeUpMaterialResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeUpMaterialResultFromDict(dict)
}

func NewVerifyGradeUpMaterialResultFromDict(data map[string]interface{}) VerifyGradeUpMaterialResult {
	return VerifyGradeUpMaterialResult{}
}

func (p VerifyGradeUpMaterialResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyGradeUpMaterialResult) Pointer() *VerifyGradeUpMaterialResult {
	return &p
}

type VerifyGradeUpMaterialByUserIdResult struct {
}

type VerifyGradeUpMaterialByUserIdAsyncResult struct {
	result *VerifyGradeUpMaterialByUserIdResult
	err    error
}

func NewVerifyGradeUpMaterialByUserIdResultFromJson(data string) VerifyGradeUpMaterialByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeUpMaterialByUserIdResultFromDict(dict)
}

func NewVerifyGradeUpMaterialByUserIdResultFromDict(data map[string]interface{}) VerifyGradeUpMaterialByUserIdResult {
	return VerifyGradeUpMaterialByUserIdResult{}
}

func (p VerifyGradeUpMaterialByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyGradeUpMaterialByUserIdResult) Pointer() *VerifyGradeUpMaterialByUserIdResult {
	return &p
}

type AddGradeByStampSheetResult struct {
	Item                    *Status            `json:"item"`
	ExperienceNamespaceName *string            `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status `json:"experienceStatus"`
}

type AddGradeByStampSheetAsyncResult struct {
	result *AddGradeByStampSheetResult
	err    error
}

func NewAddGradeByStampSheetResultFromJson(data string) AddGradeByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddGradeByStampSheetResultFromDict(dict)
}

func NewAddGradeByStampSheetResultFromDict(data map[string]interface{}) AddGradeByStampSheetResult {
	return AddGradeByStampSheetResult{
		Item:                    NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		ExperienceNamespaceName: core.CastString(data["experienceNamespaceName"]),
		ExperienceStatus:        experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer(),
	}
}

func (p AddGradeByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus":        p.ExperienceStatus.ToDict(),
	}
}

func (p AddGradeByStampSheetResult) Pointer() *AddGradeByStampSheetResult {
	return &p
}

type ApplyRankCapByStampSheetResult struct {
	Item                    *Status            `json:"item"`
	ExperienceNamespaceName *string            `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status `json:"experienceStatus"`
}

type ApplyRankCapByStampSheetAsyncResult struct {
	result *ApplyRankCapByStampSheetResult
	err    error
}

func NewApplyRankCapByStampSheetResultFromJson(data string) ApplyRankCapByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyRankCapByStampSheetResultFromDict(dict)
}

func NewApplyRankCapByStampSheetResultFromDict(data map[string]interface{}) ApplyRankCapByStampSheetResult {
	return ApplyRankCapByStampSheetResult{
		Item:                    NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		ExperienceNamespaceName: core.CastString(data["experienceNamespaceName"]),
		ExperienceStatus:        experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer(),
	}
}

func (p ApplyRankCapByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus":        p.ExperienceStatus.ToDict(),
	}
}

func (p ApplyRankCapByStampSheetResult) Pointer() *ApplyRankCapByStampSheetResult {
	return &p
}

type SubGradeByStampTaskResult struct {
	Item                    *Status            `json:"item"`
	NewContextStack         *string            `json:"newContextStack"`
	ExperienceNamespaceName *string            `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status `json:"experienceStatus"`
}

type SubGradeByStampTaskAsyncResult struct {
	result *SubGradeByStampTaskResult
	err    error
}

func NewSubGradeByStampTaskResultFromJson(data string) SubGradeByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubGradeByStampTaskResultFromDict(dict)
}

func NewSubGradeByStampTaskResultFromDict(data map[string]interface{}) SubGradeByStampTaskResult {
	return SubGradeByStampTaskResult{
		Item:                    NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack:         core.CastString(data["newContextStack"]),
		ExperienceNamespaceName: core.CastString(data["experienceNamespaceName"]),
		ExperienceStatus:        experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer(),
	}
}

func (p SubGradeByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"newContextStack":         p.NewContextStack,
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus":        p.ExperienceStatus.ToDict(),
	}
}

func (p SubGradeByStampTaskResult) Pointer() *SubGradeByStampTaskResult {
	return &p
}

type MultiplyAcquireActionsByUserIdResult struct {
	Items                     []AcquireAction `json:"items"`
	TransactionId             *string         `json:"transactionId"`
	StampSheet                *string         `json:"stampSheet"`
	StampSheetEncryptionKeyId *string         `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool           `json:"autoRunStampSheet"`
}

type MultiplyAcquireActionsByUserIdAsyncResult struct {
	result *MultiplyAcquireActionsByUserIdResult
	err    error
}

func NewMultiplyAcquireActionsByUserIdResultFromJson(data string) MultiplyAcquireActionsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMultiplyAcquireActionsByUserIdResultFromDict(dict)
}

func NewMultiplyAcquireActionsByUserIdResultFromDict(data map[string]interface{}) MultiplyAcquireActionsByUserIdResult {
	return MultiplyAcquireActionsByUserIdResult{
		Items:                     CastAcquireActions(core.CastArray(data["items"])),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p MultiplyAcquireActionsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcquireActionsFromDict(
			p.Items,
		),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p MultiplyAcquireActionsByUserIdResult) Pointer() *MultiplyAcquireActionsByUserIdResult {
	return &p
}

type MultiplyAcquireActionsByStampSheetResult struct {
	Items                     []AcquireAction `json:"items"`
	TransactionId             *string         `json:"transactionId"`
	StampSheet                *string         `json:"stampSheet"`
	StampSheetEncryptionKeyId *string         `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool           `json:"autoRunStampSheet"`
}

type MultiplyAcquireActionsByStampSheetAsyncResult struct {
	result *MultiplyAcquireActionsByStampSheetResult
	err    error
}

func NewMultiplyAcquireActionsByStampSheetResultFromJson(data string) MultiplyAcquireActionsByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMultiplyAcquireActionsByStampSheetResultFromDict(dict)
}

func NewMultiplyAcquireActionsByStampSheetResultFromDict(data map[string]interface{}) MultiplyAcquireActionsByStampSheetResult {
	return MultiplyAcquireActionsByStampSheetResult{
		Items:                     CastAcquireActions(core.CastArray(data["items"])),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p MultiplyAcquireActionsByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcquireActionsFromDict(
			p.Items,
		),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p MultiplyAcquireActionsByStampSheetResult) Pointer() *MultiplyAcquireActionsByStampSheetResult {
	return &p
}

type VerifyGradeByStampTaskResult struct {
	NewContextStack *string `json:"newContextStack"`
}

type VerifyGradeByStampTaskAsyncResult struct {
	result *VerifyGradeByStampTaskResult
	err    error
}

func NewVerifyGradeByStampTaskResultFromJson(data string) VerifyGradeByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeByStampTaskResultFromDict(dict)
}

func NewVerifyGradeByStampTaskResultFromDict(data map[string]interface{}) VerifyGradeByStampTaskResult {
	return VerifyGradeByStampTaskResult{
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p VerifyGradeByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p VerifyGradeByStampTaskResult) Pointer() *VerifyGradeByStampTaskResult {
	return &p
}

type VerifyGradeUpMaterialByStampTaskResult struct {
	NewContextStack *string `json:"newContextStack"`
}

type VerifyGradeUpMaterialByStampTaskAsyncResult struct {
	result *VerifyGradeUpMaterialByStampTaskResult
	err    error
}

func NewVerifyGradeUpMaterialByStampTaskResultFromJson(data string) VerifyGradeUpMaterialByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeUpMaterialByStampTaskResultFromDict(dict)
}

func NewVerifyGradeUpMaterialByStampTaskResultFromDict(data map[string]interface{}) VerifyGradeUpMaterialByStampTaskResult {
	return VerifyGradeUpMaterialByStampTaskResult{
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p VerifyGradeUpMaterialByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p VerifyGradeUpMaterialByStampTaskResult) Pointer() *VerifyGradeUpMaterialByStampTaskResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentGradeMaster `json:"item"`
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
		Item: NewCurrentGradeMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentGradeMasterResult struct {
	Item *CurrentGradeMaster `json:"item"`
}

type GetCurrentGradeMasterAsyncResult struct {
	result *GetCurrentGradeMasterResult
	err    error
}

func NewGetCurrentGradeMasterResultFromJson(data string) GetCurrentGradeMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentGradeMasterResultFromDict(dict)
}

func NewGetCurrentGradeMasterResultFromDict(data map[string]interface{}) GetCurrentGradeMasterResult {
	return GetCurrentGradeMasterResult{
		Item: NewCurrentGradeMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentGradeMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentGradeMasterResult) Pointer() *GetCurrentGradeMasterResult {
	return &p
}

type UpdateCurrentGradeMasterResult struct {
	Item *CurrentGradeMaster `json:"item"`
}

type UpdateCurrentGradeMasterAsyncResult struct {
	result *UpdateCurrentGradeMasterResult
	err    error
}

func NewUpdateCurrentGradeMasterResultFromJson(data string) UpdateCurrentGradeMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGradeMasterResultFromDict(dict)
}

func NewUpdateCurrentGradeMasterResultFromDict(data map[string]interface{}) UpdateCurrentGradeMasterResult {
	return UpdateCurrentGradeMasterResult{
		Item: NewCurrentGradeMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentGradeMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentGradeMasterResult) Pointer() *UpdateCurrentGradeMasterResult {
	return &p
}

type UpdateCurrentGradeMasterFromGitHubResult struct {
	Item *CurrentGradeMaster `json:"item"`
}

type UpdateCurrentGradeMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentGradeMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentGradeMasterFromGitHubResultFromJson(data string) UpdateCurrentGradeMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGradeMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentGradeMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentGradeMasterFromGitHubResult {
	return UpdateCurrentGradeMasterFromGitHubResult{
		Item: NewCurrentGradeMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentGradeMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentGradeMasterFromGitHubResult) Pointer() *UpdateCurrentGradeMasterFromGitHubResult {
	return &p
}
