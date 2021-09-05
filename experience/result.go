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

package experience

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

type DescribeExperienceModelMastersResult struct {
    Items []ExperienceModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeExperienceModelMastersAsyncResult struct {
	result *DescribeExperienceModelMastersResult
	err    error
}

func NewDescribeExperienceModelMastersResultFromJson(data string) DescribeExperienceModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeExperienceModelMastersResultFromDict(dict)
}

func NewDescribeExperienceModelMastersResultFromDict(data map[string]interface{}) DescribeExperienceModelMastersResult {
    return DescribeExperienceModelMastersResult {
        Items: CastExperienceModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeExperienceModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastExperienceModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeExperienceModelMastersResult) Pointer() *DescribeExperienceModelMastersResult {
    return &p
}

type CreateExperienceModelMasterResult struct {
    Item *ExperienceModelMaster `json:"item"`
}

type CreateExperienceModelMasterAsyncResult struct {
	result *CreateExperienceModelMasterResult
	err    error
}

func NewCreateExperienceModelMasterResultFromJson(data string) CreateExperienceModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateExperienceModelMasterResultFromDict(dict)
}

func NewCreateExperienceModelMasterResultFromDict(data map[string]interface{}) CreateExperienceModelMasterResult {
    return CreateExperienceModelMasterResult {
        Item: NewExperienceModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateExperienceModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateExperienceModelMasterResult) Pointer() *CreateExperienceModelMasterResult {
    return &p
}

type GetExperienceModelMasterResult struct {
    Item *ExperienceModelMaster `json:"item"`
}

type GetExperienceModelMasterAsyncResult struct {
	result *GetExperienceModelMasterResult
	err    error
}

func NewGetExperienceModelMasterResultFromJson(data string) GetExperienceModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetExperienceModelMasterResultFromDict(dict)
}

func NewGetExperienceModelMasterResultFromDict(data map[string]interface{}) GetExperienceModelMasterResult {
    return GetExperienceModelMasterResult {
        Item: NewExperienceModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetExperienceModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetExperienceModelMasterResult) Pointer() *GetExperienceModelMasterResult {
    return &p
}

type UpdateExperienceModelMasterResult struct {
    Item *ExperienceModelMaster `json:"item"`
}

type UpdateExperienceModelMasterAsyncResult struct {
	result *UpdateExperienceModelMasterResult
	err    error
}

func NewUpdateExperienceModelMasterResultFromJson(data string) UpdateExperienceModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateExperienceModelMasterResultFromDict(dict)
}

func NewUpdateExperienceModelMasterResultFromDict(data map[string]interface{}) UpdateExperienceModelMasterResult {
    return UpdateExperienceModelMasterResult {
        Item: NewExperienceModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateExperienceModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateExperienceModelMasterResult) Pointer() *UpdateExperienceModelMasterResult {
    return &p
}

type DeleteExperienceModelMasterResult struct {
    Item *ExperienceModelMaster `json:"item"`
}

type DeleteExperienceModelMasterAsyncResult struct {
	result *DeleteExperienceModelMasterResult
	err    error
}

func NewDeleteExperienceModelMasterResultFromJson(data string) DeleteExperienceModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteExperienceModelMasterResultFromDict(dict)
}

func NewDeleteExperienceModelMasterResultFromDict(data map[string]interface{}) DeleteExperienceModelMasterResult {
    return DeleteExperienceModelMasterResult {
        Item: NewExperienceModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteExperienceModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteExperienceModelMasterResult) Pointer() *DeleteExperienceModelMasterResult {
    return &p
}

type DescribeExperienceModelsResult struct {
    Items []ExperienceModel `json:"items"`
}

type DescribeExperienceModelsAsyncResult struct {
	result *DescribeExperienceModelsResult
	err    error
}

func NewDescribeExperienceModelsResultFromJson(data string) DescribeExperienceModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeExperienceModelsResultFromDict(dict)
}

func NewDescribeExperienceModelsResultFromDict(data map[string]interface{}) DescribeExperienceModelsResult {
    return DescribeExperienceModelsResult {
        Items: CastExperienceModels(core.CastArray(data["items"])),
    }
}

func (p DescribeExperienceModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastExperienceModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeExperienceModelsResult) Pointer() *DescribeExperienceModelsResult {
    return &p
}

type GetExperienceModelResult struct {
    Item *ExperienceModel `json:"item"`
}

type GetExperienceModelAsyncResult struct {
	result *GetExperienceModelResult
	err    error
}

func NewGetExperienceModelResultFromJson(data string) GetExperienceModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetExperienceModelResultFromDict(dict)
}

func NewGetExperienceModelResultFromDict(data map[string]interface{}) GetExperienceModelResult {
    return GetExperienceModelResult {
        Item: NewExperienceModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetExperienceModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetExperienceModelResult) Pointer() *GetExperienceModelResult {
    return &p
}

type DescribeThresholdMastersResult struct {
    Items []ThresholdMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeThresholdMastersAsyncResult struct {
	result *DescribeThresholdMastersResult
	err    error
}

func NewDescribeThresholdMastersResultFromJson(data string) DescribeThresholdMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeThresholdMastersResultFromDict(dict)
}

func NewDescribeThresholdMastersResultFromDict(data map[string]interface{}) DescribeThresholdMastersResult {
    return DescribeThresholdMastersResult {
        Items: CastThresholdMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeThresholdMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastThresholdMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeThresholdMastersResult) Pointer() *DescribeThresholdMastersResult {
    return &p
}

type CreateThresholdMasterResult struct {
    Item *ThresholdMaster `json:"item"`
}

type CreateThresholdMasterAsyncResult struct {
	result *CreateThresholdMasterResult
	err    error
}

func NewCreateThresholdMasterResultFromJson(data string) CreateThresholdMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateThresholdMasterResultFromDict(dict)
}

func NewCreateThresholdMasterResultFromDict(data map[string]interface{}) CreateThresholdMasterResult {
    return CreateThresholdMasterResult {
        Item: NewThresholdMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateThresholdMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateThresholdMasterResult) Pointer() *CreateThresholdMasterResult {
    return &p
}

type GetThresholdMasterResult struct {
    Item *ThresholdMaster `json:"item"`
}

type GetThresholdMasterAsyncResult struct {
	result *GetThresholdMasterResult
	err    error
}

func NewGetThresholdMasterResultFromJson(data string) GetThresholdMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetThresholdMasterResultFromDict(dict)
}

func NewGetThresholdMasterResultFromDict(data map[string]interface{}) GetThresholdMasterResult {
    return GetThresholdMasterResult {
        Item: NewThresholdMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetThresholdMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetThresholdMasterResult) Pointer() *GetThresholdMasterResult {
    return &p
}

type UpdateThresholdMasterResult struct {
    Item *ThresholdMaster `json:"item"`
}

type UpdateThresholdMasterAsyncResult struct {
	result *UpdateThresholdMasterResult
	err    error
}

func NewUpdateThresholdMasterResultFromJson(data string) UpdateThresholdMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateThresholdMasterResultFromDict(dict)
}

func NewUpdateThresholdMasterResultFromDict(data map[string]interface{}) UpdateThresholdMasterResult {
    return UpdateThresholdMasterResult {
        Item: NewThresholdMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateThresholdMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateThresholdMasterResult) Pointer() *UpdateThresholdMasterResult {
    return &p
}

type DeleteThresholdMasterResult struct {
    Item *ThresholdMaster `json:"item"`
}

type DeleteThresholdMasterAsyncResult struct {
	result *DeleteThresholdMasterResult
	err    error
}

func NewDeleteThresholdMasterResultFromJson(data string) DeleteThresholdMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteThresholdMasterResultFromDict(dict)
}

func NewDeleteThresholdMasterResultFromDict(data map[string]interface{}) DeleteThresholdMasterResult {
    return DeleteThresholdMasterResult {
        Item: NewThresholdMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteThresholdMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteThresholdMasterResult) Pointer() *DeleteThresholdMasterResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentExperienceMaster `json:"item"`
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
        Item: NewCurrentExperienceMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentExperienceMasterResult struct {
    Item *CurrentExperienceMaster `json:"item"`
}

type GetCurrentExperienceMasterAsyncResult struct {
	result *GetCurrentExperienceMasterResult
	err    error
}

func NewGetCurrentExperienceMasterResultFromJson(data string) GetCurrentExperienceMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentExperienceMasterResultFromDict(dict)
}

func NewGetCurrentExperienceMasterResultFromDict(data map[string]interface{}) GetCurrentExperienceMasterResult {
    return GetCurrentExperienceMasterResult {
        Item: NewCurrentExperienceMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentExperienceMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentExperienceMasterResult) Pointer() *GetCurrentExperienceMasterResult {
    return &p
}

type UpdateCurrentExperienceMasterResult struct {
    Item *CurrentExperienceMaster `json:"item"`
}

type UpdateCurrentExperienceMasterAsyncResult struct {
	result *UpdateCurrentExperienceMasterResult
	err    error
}

func NewUpdateCurrentExperienceMasterResultFromJson(data string) UpdateCurrentExperienceMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentExperienceMasterResultFromDict(dict)
}

func NewUpdateCurrentExperienceMasterResultFromDict(data map[string]interface{}) UpdateCurrentExperienceMasterResult {
    return UpdateCurrentExperienceMasterResult {
        Item: NewCurrentExperienceMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentExperienceMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentExperienceMasterResult) Pointer() *UpdateCurrentExperienceMasterResult {
    return &p
}

type UpdateCurrentExperienceMasterFromGitHubResult struct {
    Item *CurrentExperienceMaster `json:"item"`
}

type UpdateCurrentExperienceMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentExperienceMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentExperienceMasterFromGitHubResultFromJson(data string) UpdateCurrentExperienceMasterFromGitHubResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentExperienceMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentExperienceMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentExperienceMasterFromGitHubResult {
    return UpdateCurrentExperienceMasterFromGitHubResult {
        Item: NewCurrentExperienceMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentExperienceMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentExperienceMasterFromGitHubResult) Pointer() *UpdateCurrentExperienceMasterFromGitHubResult {
    return &p
}

type DescribeStatusesResult struct {
    Items []Status `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
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
    return DescribeStatusesResult {
        Items: CastStatuses(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeStatusesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    Items []Status `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
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
    return DescribeStatusesByUserIdResult {
        Items: CastStatuses(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeStatusesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return GetStatusResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetStatusResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return GetStatusByUserIdResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetStatusByUserIdResult) Pointer() *GetStatusByUserIdResult {
    return &p
}

type GetStatusWithSignatureResult struct {
    Item *Status `json:"item"`
    Body *string `json:"body"`
    Signature *string `json:"signature"`
}

type GetStatusWithSignatureAsyncResult struct {
	result *GetStatusWithSignatureResult
	err    error
}

func NewGetStatusWithSignatureResultFromJson(data string) GetStatusWithSignatureResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetStatusWithSignatureResultFromDict(dict)
}

func NewGetStatusWithSignatureResultFromDict(data map[string]interface{}) GetStatusWithSignatureResult {
    return GetStatusWithSignatureResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Body: core.CastString(data["body"]),
        Signature: core.CastString(data["signature"]),
    }
}

func (p GetStatusWithSignatureResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "body": p.Body,
        "signature": p.Signature,
    }
}

func (p GetStatusWithSignatureResult) Pointer() *GetStatusWithSignatureResult {
    return &p
}

type GetStatusWithSignatureByUserIdResult struct {
    Item *Status `json:"item"`
    Body *string `json:"body"`
    Signature *string `json:"signature"`
}

type GetStatusWithSignatureByUserIdAsyncResult struct {
	result *GetStatusWithSignatureByUserIdResult
	err    error
}

func NewGetStatusWithSignatureByUserIdResultFromJson(data string) GetStatusWithSignatureByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetStatusWithSignatureByUserIdResultFromDict(dict)
}

func NewGetStatusWithSignatureByUserIdResultFromDict(data map[string]interface{}) GetStatusWithSignatureByUserIdResult {
    return GetStatusWithSignatureByUserIdResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Body: core.CastString(data["body"]),
        Signature: core.CastString(data["signature"]),
    }
}

func (p GetStatusWithSignatureByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "body": p.Body,
        "signature": p.Signature,
    }
}

func (p GetStatusWithSignatureByUserIdResult) Pointer() *GetStatusWithSignatureByUserIdResult {
    return &p
}

type AddExperienceByUserIdResult struct {
    Item *Status `json:"item"`
}

type AddExperienceByUserIdAsyncResult struct {
	result *AddExperienceByUserIdResult
	err    error
}

func NewAddExperienceByUserIdResultFromJson(data string) AddExperienceByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddExperienceByUserIdResultFromDict(dict)
}

func NewAddExperienceByUserIdResultFromDict(data map[string]interface{}) AddExperienceByUserIdResult {
    return AddExperienceByUserIdResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AddExperienceByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AddExperienceByUserIdResult) Pointer() *AddExperienceByUserIdResult {
    return &p
}

type SetExperienceByUserIdResult struct {
    Item *Status `json:"item"`
}

type SetExperienceByUserIdAsyncResult struct {
	result *SetExperienceByUserIdResult
	err    error
}

func NewSetExperienceByUserIdResultFromJson(data string) SetExperienceByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetExperienceByUserIdResultFromDict(dict)
}

func NewSetExperienceByUserIdResultFromDict(data map[string]interface{}) SetExperienceByUserIdResult {
    return SetExperienceByUserIdResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SetExperienceByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SetExperienceByUserIdResult) Pointer() *SetExperienceByUserIdResult {
    return &p
}

type AddRankCapByUserIdResult struct {
    Item *Status `json:"item"`
}

type AddRankCapByUserIdAsyncResult struct {
	result *AddRankCapByUserIdResult
	err    error
}

func NewAddRankCapByUserIdResultFromJson(data string) AddRankCapByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddRankCapByUserIdResultFromDict(dict)
}

func NewAddRankCapByUserIdResultFromDict(data map[string]interface{}) AddRankCapByUserIdResult {
    return AddRankCapByUserIdResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AddRankCapByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AddRankCapByUserIdResult) Pointer() *AddRankCapByUserIdResult {
    return &p
}

type SetRankCapByUserIdResult struct {
    Item *Status `json:"item"`
}

type SetRankCapByUserIdAsyncResult struct {
	result *SetRankCapByUserIdResult
	err    error
}

func NewSetRankCapByUserIdResultFromJson(data string) SetRankCapByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetRankCapByUserIdResultFromDict(dict)
}

func NewSetRankCapByUserIdResultFromDict(data map[string]interface{}) SetRankCapByUserIdResult {
    return SetRankCapByUserIdResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SetRankCapByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SetRankCapByUserIdResult) Pointer() *SetRankCapByUserIdResult {
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
    return DeleteStatusByUserIdResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteStatusByUserIdResult) Pointer() *DeleteStatusByUserIdResult {
    return &p
}

type AddExperienceByStampSheetResult struct {
    Item *Status `json:"item"`
}

type AddExperienceByStampSheetAsyncResult struct {
	result *AddExperienceByStampSheetResult
	err    error
}

func NewAddExperienceByStampSheetResultFromJson(data string) AddExperienceByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddExperienceByStampSheetResultFromDict(dict)
}

func NewAddExperienceByStampSheetResultFromDict(data map[string]interface{}) AddExperienceByStampSheetResult {
    return AddExperienceByStampSheetResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AddExperienceByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AddExperienceByStampSheetResult) Pointer() *AddExperienceByStampSheetResult {
    return &p
}

type AddRankCapByStampSheetResult struct {
    Item *Status `json:"item"`
}

type AddRankCapByStampSheetAsyncResult struct {
	result *AddRankCapByStampSheetResult
	err    error
}

func NewAddRankCapByStampSheetResultFromJson(data string) AddRankCapByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddRankCapByStampSheetResultFromDict(dict)
}

func NewAddRankCapByStampSheetResultFromDict(data map[string]interface{}) AddRankCapByStampSheetResult {
    return AddRankCapByStampSheetResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AddRankCapByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AddRankCapByStampSheetResult) Pointer() *AddRankCapByStampSheetResult {
    return &p
}

type SetRankCapByStampSheetResult struct {
    Item *Status `json:"item"`
}

type SetRankCapByStampSheetAsyncResult struct {
	result *SetRankCapByStampSheetResult
	err    error
}

func NewSetRankCapByStampSheetResultFromJson(data string) SetRankCapByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetRankCapByStampSheetResultFromDict(dict)
}

func NewSetRankCapByStampSheetResultFromDict(data map[string]interface{}) SetRankCapByStampSheetResult {
    return SetRankCapByStampSheetResult {
        Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SetRankCapByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SetRankCapByStampSheetResult) Pointer() *SetRankCapByStampSheetResult {
    return &p
}