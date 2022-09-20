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

package serialKey

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    Items []Namespace `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
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
    return DescribeNamespacesResult {
        Items: CastNamespaces(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return CreateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return GetNamespaceStatusResult {
        Status: core.CastString(data["status"]),
    }
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return GetNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return UpdateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return DeleteNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
    return &p
}

type DescribeIssueJobsResult struct {
    Items []IssueJob `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeIssueJobsAsyncResult struct {
	result *DescribeIssueJobsResult
	err    error
}

func NewDescribeIssueJobsResultFromJson(data string) DescribeIssueJobsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeIssueJobsResultFromDict(dict)
}

func NewDescribeIssueJobsResultFromDict(data map[string]interface{}) DescribeIssueJobsResult {
    return DescribeIssueJobsResult {
        Items: CastIssueJobs(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeIssueJobsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastIssueJobsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeIssueJobsResult) Pointer() *DescribeIssueJobsResult {
    return &p
}

type GetIssueJobResult struct {
    Item *IssueJob `json:"item"`
}

type GetIssueJobAsyncResult struct {
	result *GetIssueJobResult
	err    error
}

func NewGetIssueJobResultFromJson(data string) GetIssueJobResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetIssueJobResultFromDict(dict)
}

func NewGetIssueJobResultFromDict(data map[string]interface{}) GetIssueJobResult {
    return GetIssueJobResult {
        Item: NewIssueJobFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetIssueJobResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetIssueJobResult) Pointer() *GetIssueJobResult {
    return &p
}

type IssueResult struct {
    Item *IssueJob `json:"item"`
}

type IssueAsyncResult struct {
	result *IssueResult
	err    error
}

func NewIssueResultFromJson(data string) IssueResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewIssueResultFromDict(dict)
}

func NewIssueResultFromDict(data map[string]interface{}) IssueResult {
    return IssueResult {
        Item: NewIssueJobFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p IssueResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p IssueResult) Pointer() *IssueResult {
    return &p
}

type DescribeSerialCodesResult struct {
    Items []string `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSerialCodesAsyncResult struct {
	result *DescribeSerialCodesResult
	err    error
}

func NewDescribeSerialCodesResultFromJson(data string) DescribeSerialCodesResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSerialCodesResultFromDict(dict)
}

func NewDescribeSerialCodesResultFromDict(data map[string]interface{}) DescribeSerialCodesResult {
    return DescribeSerialCodesResult {
        Items: core.CastStrings(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSerialCodesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": core.CastStringsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSerialCodesResult) Pointer() *DescribeSerialCodesResult {
    return &p
}

type UseResult struct {
    Item *SerialKey `json:"item"`
    CampaignModel *CampaignModel `json:"campaignModel"`
}

type UseAsyncResult struct {
	result *UseResult
	err    error
}

func NewUseResultFromJson(data string) UseResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUseResultFromDict(dict)
}

func NewUseResultFromDict(data map[string]interface{}) UseResult {
    return UseResult {
        Item: NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer(),
        CampaignModel: NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer(),
    }
}

func (p UseResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "campaignModel": p.CampaignModel.ToDict(),
    }
}

func (p UseResult) Pointer() *UseResult {
    return &p
}

type UseByUserIdResult struct {
    Item *SerialKey `json:"item"`
    CampaignModel *CampaignModel `json:"campaignModel"`
}

type UseByUserIdAsyncResult struct {
	result *UseByUserIdResult
	err    error
}

func NewUseByUserIdResultFromJson(data string) UseByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUseByUserIdResultFromDict(dict)
}

func NewUseByUserIdResultFromDict(data map[string]interface{}) UseByUserIdResult {
    return UseByUserIdResult {
        Item: NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer(),
        CampaignModel: NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer(),
    }
}

func (p UseByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "campaignModel": p.CampaignModel.ToDict(),
    }
}

func (p UseByUserIdResult) Pointer() *UseByUserIdResult {
    return &p
}

type UseByStampTaskResult struct {
    Item *SerialKey `json:"item"`
    CampaignModel *CampaignModel `json:"campaignModel"`
    NewContextStack *string `json:"newContextStack"`
}

type UseByStampTaskAsyncResult struct {
	result *UseByStampTaskResult
	err    error
}

func NewUseByStampTaskResultFromJson(data string) UseByStampTaskResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUseByStampTaskResultFromDict(dict)
}

func NewUseByStampTaskResultFromDict(data map[string]interface{}) UseByStampTaskResult {
    return UseByStampTaskResult {
        Item: NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer(),
        CampaignModel: NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p UseByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "campaignModel": p.CampaignModel.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p UseByStampTaskResult) Pointer() *UseByStampTaskResult {
    return &p
}

type DescribeCampaignModelsResult struct {
    Items []CampaignModel `json:"items"`
}

type DescribeCampaignModelsAsyncResult struct {
	result *DescribeCampaignModelsResult
	err    error
}

func NewDescribeCampaignModelsResultFromJson(data string) DescribeCampaignModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeCampaignModelsResultFromDict(dict)
}

func NewDescribeCampaignModelsResultFromDict(data map[string]interface{}) DescribeCampaignModelsResult {
    return DescribeCampaignModelsResult {
        Items: CastCampaignModels(core.CastArray(data["items"])),
    }
}

func (p DescribeCampaignModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCampaignModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeCampaignModelsResult) Pointer() *DescribeCampaignModelsResult {
    return &p
}

type GetCampaignModelResult struct {
    Item *CampaignModel `json:"item"`
}

type GetCampaignModelAsyncResult struct {
	result *GetCampaignModelResult
	err    error
}

func NewGetCampaignModelResultFromJson(data string) GetCampaignModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCampaignModelResultFromDict(dict)
}

func NewGetCampaignModelResultFromDict(data map[string]interface{}) GetCampaignModelResult {
    return GetCampaignModelResult {
        Item: NewCampaignModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCampaignModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCampaignModelResult) Pointer() *GetCampaignModelResult {
    return &p
}

type DescribeCampaignModelMastersResult struct {
    Items []CampaignModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeCampaignModelMastersAsyncResult struct {
	result *DescribeCampaignModelMastersResult
	err    error
}

func NewDescribeCampaignModelMastersResultFromJson(data string) DescribeCampaignModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeCampaignModelMastersResultFromDict(dict)
}

func NewDescribeCampaignModelMastersResultFromDict(data map[string]interface{}) DescribeCampaignModelMastersResult {
    return DescribeCampaignModelMastersResult {
        Items: CastCampaignModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeCampaignModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCampaignModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeCampaignModelMastersResult) Pointer() *DescribeCampaignModelMastersResult {
    return &p
}

type CreateCampaignModelMasterResult struct {
    Item *CampaignModelMaster `json:"item"`
}

type CreateCampaignModelMasterAsyncResult struct {
	result *CreateCampaignModelMasterResult
	err    error
}

func NewCreateCampaignModelMasterResultFromJson(data string) CreateCampaignModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateCampaignModelMasterResultFromDict(dict)
}

func NewCreateCampaignModelMasterResultFromDict(data map[string]interface{}) CreateCampaignModelMasterResult {
    return CreateCampaignModelMasterResult {
        Item: NewCampaignModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateCampaignModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateCampaignModelMasterResult) Pointer() *CreateCampaignModelMasterResult {
    return &p
}

type GetCampaignModelMasterResult struct {
    Item *CampaignModelMaster `json:"item"`
}

type GetCampaignModelMasterAsyncResult struct {
	result *GetCampaignModelMasterResult
	err    error
}

func NewGetCampaignModelMasterResultFromJson(data string) GetCampaignModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCampaignModelMasterResultFromDict(dict)
}

func NewGetCampaignModelMasterResultFromDict(data map[string]interface{}) GetCampaignModelMasterResult {
    return GetCampaignModelMasterResult {
        Item: NewCampaignModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCampaignModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCampaignModelMasterResult) Pointer() *GetCampaignModelMasterResult {
    return &p
}

type UpdateCampaignModelMasterResult struct {
    Item *CampaignModelMaster `json:"item"`
}

type UpdateCampaignModelMasterAsyncResult struct {
	result *UpdateCampaignModelMasterResult
	err    error
}

func NewUpdateCampaignModelMasterResultFromJson(data string) UpdateCampaignModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCampaignModelMasterResultFromDict(dict)
}

func NewUpdateCampaignModelMasterResultFromDict(data map[string]interface{}) UpdateCampaignModelMasterResult {
    return UpdateCampaignModelMasterResult {
        Item: NewCampaignModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCampaignModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCampaignModelMasterResult) Pointer() *UpdateCampaignModelMasterResult {
    return &p
}

type DeleteCampaignModelMasterResult struct {
    Item *CampaignModelMaster `json:"item"`
}

type DeleteCampaignModelMasterAsyncResult struct {
	result *DeleteCampaignModelMasterResult
	err    error
}

func NewDeleteCampaignModelMasterResultFromJson(data string) DeleteCampaignModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteCampaignModelMasterResultFromDict(dict)
}

func NewDeleteCampaignModelMasterResultFromDict(data map[string]interface{}) DeleteCampaignModelMasterResult {
    return DeleteCampaignModelMasterResult {
        Item: NewCampaignModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteCampaignModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteCampaignModelMasterResult) Pointer() *DeleteCampaignModelMasterResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentCampaignMaster `json:"item"`
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
    return ExportMasterResult {
        Item: NewCurrentCampaignMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
    return &p
}

type GetCurrentCampaignMasterResult struct {
    Item *CurrentCampaignMaster `json:"item"`
}

type GetCurrentCampaignMasterAsyncResult struct {
	result *GetCurrentCampaignMasterResult
	err    error
}

func NewGetCurrentCampaignMasterResultFromJson(data string) GetCurrentCampaignMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentCampaignMasterResultFromDict(dict)
}

func NewGetCurrentCampaignMasterResultFromDict(data map[string]interface{}) GetCurrentCampaignMasterResult {
    return GetCurrentCampaignMasterResult {
        Item: NewCurrentCampaignMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentCampaignMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentCampaignMasterResult) Pointer() *GetCurrentCampaignMasterResult {
    return &p
}

type UpdateCurrentCampaignMasterResult struct {
    Item *CurrentCampaignMaster `json:"item"`
}

type UpdateCurrentCampaignMasterAsyncResult struct {
	result *UpdateCurrentCampaignMasterResult
	err    error
}

func NewUpdateCurrentCampaignMasterResultFromJson(data string) UpdateCurrentCampaignMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentCampaignMasterResultFromDict(dict)
}

func NewUpdateCurrentCampaignMasterResultFromDict(data map[string]interface{}) UpdateCurrentCampaignMasterResult {
    return UpdateCurrentCampaignMasterResult {
        Item: NewCurrentCampaignMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentCampaignMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentCampaignMasterResult) Pointer() *UpdateCurrentCampaignMasterResult {
    return &p
}

type UpdateCurrentCampaignMasterFromGitHubResult struct {
    Item *CurrentCampaignMaster `json:"item"`
}

type UpdateCurrentCampaignMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentCampaignMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentCampaignMasterFromGitHubResultFromJson(data string) UpdateCurrentCampaignMasterFromGitHubResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentCampaignMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentCampaignMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentCampaignMasterFromGitHubResult {
    return UpdateCurrentCampaignMasterFromGitHubResult {
        Item: NewCurrentCampaignMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentCampaignMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentCampaignMasterFromGitHubResult) Pointer() *UpdateCurrentCampaignMasterFromGitHubResult {
    return &p
}