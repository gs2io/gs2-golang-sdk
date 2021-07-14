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

package version

import "core"

type DescribeNamespacesResult struct {
    Items []Namespace `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
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

type DescribeVersionModelMastersResult struct {
    Items []VersionModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeVersionModelMastersAsyncResult struct {
	result *DescribeVersionModelMastersResult
	err    error
}

func NewDescribeVersionModelMastersResultFromDict(data map[string]interface{}) DescribeVersionModelMastersResult {
    return DescribeVersionModelMastersResult {
        Items: CastVersionModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeVersionModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastVersionModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeVersionModelMastersResult) Pointer() *DescribeVersionModelMastersResult {
    return &p
}

type CreateVersionModelMasterResult struct {
    Item *VersionModelMaster `json:"item"`
}

type CreateVersionModelMasterAsyncResult struct {
	result *CreateVersionModelMasterResult
	err    error
}

func NewCreateVersionModelMasterResultFromDict(data map[string]interface{}) CreateVersionModelMasterResult {
    return CreateVersionModelMasterResult {
        Item: NewVersionModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateVersionModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateVersionModelMasterResult) Pointer() *CreateVersionModelMasterResult {
    return &p
}

type GetVersionModelMasterResult struct {
    Item *VersionModelMaster `json:"item"`
}

type GetVersionModelMasterAsyncResult struct {
	result *GetVersionModelMasterResult
	err    error
}

func NewGetVersionModelMasterResultFromDict(data map[string]interface{}) GetVersionModelMasterResult {
    return GetVersionModelMasterResult {
        Item: NewVersionModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetVersionModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetVersionModelMasterResult) Pointer() *GetVersionModelMasterResult {
    return &p
}

type UpdateVersionModelMasterResult struct {
    Item *VersionModelMaster `json:"item"`
}

type UpdateVersionModelMasterAsyncResult struct {
	result *UpdateVersionModelMasterResult
	err    error
}

func NewUpdateVersionModelMasterResultFromDict(data map[string]interface{}) UpdateVersionModelMasterResult {
    return UpdateVersionModelMasterResult {
        Item: NewVersionModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateVersionModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateVersionModelMasterResult) Pointer() *UpdateVersionModelMasterResult {
    return &p
}

type DeleteVersionModelMasterResult struct {
    Item *VersionModelMaster `json:"item"`
}

type DeleteVersionModelMasterAsyncResult struct {
	result *DeleteVersionModelMasterResult
	err    error
}

func NewDeleteVersionModelMasterResultFromDict(data map[string]interface{}) DeleteVersionModelMasterResult {
    return DeleteVersionModelMasterResult {
        Item: NewVersionModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteVersionModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteVersionModelMasterResult) Pointer() *DeleteVersionModelMasterResult {
    return &p
}

type DescribeVersionModelsResult struct {
    Items []VersionModel `json:"items"`
}

type DescribeVersionModelsAsyncResult struct {
	result *DescribeVersionModelsResult
	err    error
}

func NewDescribeVersionModelsResultFromDict(data map[string]interface{}) DescribeVersionModelsResult {
    return DescribeVersionModelsResult {
        Items: CastVersionModels(core.CastArray(data["items"])),
    }
}

func (p DescribeVersionModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastVersionModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeVersionModelsResult) Pointer() *DescribeVersionModelsResult {
    return &p
}

type GetVersionModelResult struct {
    Item *VersionModel `json:"item"`
}

type GetVersionModelAsyncResult struct {
	result *GetVersionModelResult
	err    error
}

func NewGetVersionModelResultFromDict(data map[string]interface{}) GetVersionModelResult {
    return GetVersionModelResult {
        Item: NewVersionModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetVersionModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetVersionModelResult) Pointer() *GetVersionModelResult {
    return &p
}

type DescribeAcceptVersionsResult struct {
    Items []AcceptVersion `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeAcceptVersionsAsyncResult struct {
	result *DescribeAcceptVersionsResult
	err    error
}

func NewDescribeAcceptVersionsResultFromDict(data map[string]interface{}) DescribeAcceptVersionsResult {
    return DescribeAcceptVersionsResult {
        Items: CastAcceptVersions(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeAcceptVersionsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastAcceptVersionsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeAcceptVersionsResult) Pointer() *DescribeAcceptVersionsResult {
    return &p
}

type DescribeAcceptVersionsByUserIdResult struct {
    Items []AcceptVersion `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeAcceptVersionsByUserIdAsyncResult struct {
	result *DescribeAcceptVersionsByUserIdResult
	err    error
}

func NewDescribeAcceptVersionsByUserIdResultFromDict(data map[string]interface{}) DescribeAcceptVersionsByUserIdResult {
    return DescribeAcceptVersionsByUserIdResult {
        Items: CastAcceptVersions(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeAcceptVersionsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastAcceptVersionsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeAcceptVersionsByUserIdResult) Pointer() *DescribeAcceptVersionsByUserIdResult {
    return &p
}

type AcceptResult struct {
    Item *AcceptVersion `json:"item"`
}

type AcceptAsyncResult struct {
	result *AcceptResult
	err    error
}

func NewAcceptResultFromDict(data map[string]interface{}) AcceptResult {
    return AcceptResult {
        Item: NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AcceptResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AcceptResult) Pointer() *AcceptResult {
    return &p
}

type AcceptByUserIdResult struct {
    Item *AcceptVersion `json:"item"`
}

type AcceptByUserIdAsyncResult struct {
	result *AcceptByUserIdResult
	err    error
}

func NewAcceptByUserIdResultFromDict(data map[string]interface{}) AcceptByUserIdResult {
    return AcceptByUserIdResult {
        Item: NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AcceptByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AcceptByUserIdResult) Pointer() *AcceptByUserIdResult {
    return &p
}

type GetAcceptVersionResult struct {
    Item *AcceptVersion `json:"item"`
}

type GetAcceptVersionAsyncResult struct {
	result *GetAcceptVersionResult
	err    error
}

func NewGetAcceptVersionResultFromDict(data map[string]interface{}) GetAcceptVersionResult {
    return GetAcceptVersionResult {
        Item: NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetAcceptVersionResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetAcceptVersionResult) Pointer() *GetAcceptVersionResult {
    return &p
}

type GetAcceptVersionByUserIdResult struct {
    Item *AcceptVersion `json:"item"`
}

type GetAcceptVersionByUserIdAsyncResult struct {
	result *GetAcceptVersionByUserIdResult
	err    error
}

func NewGetAcceptVersionByUserIdResultFromDict(data map[string]interface{}) GetAcceptVersionByUserIdResult {
    return GetAcceptVersionByUserIdResult {
        Item: NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetAcceptVersionByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetAcceptVersionByUserIdResult) Pointer() *GetAcceptVersionByUserIdResult {
    return &p
}

type DeleteAcceptVersionResult struct {
}

type DeleteAcceptVersionAsyncResult struct {
	result *DeleteAcceptVersionResult
	err    error
}

func NewDeleteAcceptVersionResultFromDict(data map[string]interface{}) DeleteAcceptVersionResult {
    return DeleteAcceptVersionResult {
    }
}

func (p DeleteAcceptVersionResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteAcceptVersionResult) Pointer() *DeleteAcceptVersionResult {
    return &p
}

type DeleteAcceptVersionByUserIdResult struct {
}

type DeleteAcceptVersionByUserIdAsyncResult struct {
	result *DeleteAcceptVersionByUserIdResult
	err    error
}

func NewDeleteAcceptVersionByUserIdResultFromDict(data map[string]interface{}) DeleteAcceptVersionByUserIdResult {
    return DeleteAcceptVersionByUserIdResult {
    }
}

func (p DeleteAcceptVersionByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteAcceptVersionByUserIdResult) Pointer() *DeleteAcceptVersionByUserIdResult {
    return &p
}

type CheckVersionResult struct {
    ProjectToken *string `json:"projectToken"`
    Warnings []Status `json:"warnings"`
    Errors []Status `json:"errors"`
}

type CheckVersionAsyncResult struct {
	result *CheckVersionResult
	err    error
}

func NewCheckVersionResultFromDict(data map[string]interface{}) CheckVersionResult {
    return CheckVersionResult {
        ProjectToken: core.CastString(data["projectToken"]),
        Warnings: CastStatuses(core.CastArray(data["warnings"])),
        Errors: CastStatuses(core.CastArray(data["errors"])),
    }
}

func (p CheckVersionResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "projectToken": p.ProjectToken,
        "warnings": CastStatusesFromDict(
            p.Warnings,
        ),
        "errors": CastStatusesFromDict(
            p.Errors,
        ),
    }
}

func (p CheckVersionResult) Pointer() *CheckVersionResult {
    return &p
}

type CheckVersionByUserIdResult struct {
    ProjectToken *string `json:"projectToken"`
    Warnings []Status `json:"warnings"`
    Errors []Status `json:"errors"`
}

type CheckVersionByUserIdAsyncResult struct {
	result *CheckVersionByUserIdResult
	err    error
}

func NewCheckVersionByUserIdResultFromDict(data map[string]interface{}) CheckVersionByUserIdResult {
    return CheckVersionByUserIdResult {
        ProjectToken: core.CastString(data["projectToken"]),
        Warnings: CastStatuses(core.CastArray(data["warnings"])),
        Errors: CastStatuses(core.CastArray(data["errors"])),
    }
}

func (p CheckVersionByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "projectToken": p.ProjectToken,
        "warnings": CastStatusesFromDict(
            p.Warnings,
        ),
        "errors": CastStatusesFromDict(
            p.Errors,
        ),
    }
}

func (p CheckVersionByUserIdResult) Pointer() *CheckVersionByUserIdResult {
    return &p
}

type CalculateSignatureResult struct {
    Body *string `json:"body"`
    Signature *string `json:"signature"`
}

type CalculateSignatureAsyncResult struct {
	result *CalculateSignatureResult
	err    error
}

func NewCalculateSignatureResultFromDict(data map[string]interface{}) CalculateSignatureResult {
    return CalculateSignatureResult {
        Body: core.CastString(data["body"]),
        Signature: core.CastString(data["signature"]),
    }
}

func (p CalculateSignatureResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "body": p.Body,
        "signature": p.Signature,
    }
}

func (p CalculateSignatureResult) Pointer() *CalculateSignatureResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentVersionMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentVersionMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentVersionMasterResult struct {
    Item *CurrentVersionMaster `json:"item"`
}

type GetCurrentVersionMasterAsyncResult struct {
	result *GetCurrentVersionMasterResult
	err    error
}

func NewGetCurrentVersionMasterResultFromDict(data map[string]interface{}) GetCurrentVersionMasterResult {
    return GetCurrentVersionMasterResult {
        Item: NewCurrentVersionMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentVersionMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentVersionMasterResult) Pointer() *GetCurrentVersionMasterResult {
    return &p
}

type UpdateCurrentVersionMasterResult struct {
    Item *CurrentVersionMaster `json:"item"`
}

type UpdateCurrentVersionMasterAsyncResult struct {
	result *UpdateCurrentVersionMasterResult
	err    error
}

func NewUpdateCurrentVersionMasterResultFromDict(data map[string]interface{}) UpdateCurrentVersionMasterResult {
    return UpdateCurrentVersionMasterResult {
        Item: NewCurrentVersionMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentVersionMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentVersionMasterResult) Pointer() *UpdateCurrentVersionMasterResult {
    return &p
}

type UpdateCurrentVersionMasterFromGitHubResult struct {
    Item *CurrentVersionMaster `json:"item"`
}

type UpdateCurrentVersionMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentVersionMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentVersionMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentVersionMasterFromGitHubResult {
    return UpdateCurrentVersionMasterFromGitHubResult {
        Item: NewCurrentVersionMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentVersionMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentVersionMasterFromGitHubResult) Pointer() *UpdateCurrentVersionMasterFromGitHubResult {
    return &p
}