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

package distributor

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

type DescribeDistributorModelMastersResult struct {
    Items []DistributorModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeDistributorModelMastersAsyncResult struct {
	result *DescribeDistributorModelMastersResult
	err    error
}

func NewDescribeDistributorModelMastersResultFromDict(data map[string]interface{}) DescribeDistributorModelMastersResult {
    return DescribeDistributorModelMastersResult {
        Items: CastDistributorModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeDistributorModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastDistributorModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeDistributorModelMastersResult) Pointer() *DescribeDistributorModelMastersResult {
    return &p
}

type CreateDistributorModelMasterResult struct {
    Item *DistributorModelMaster `json:"item"`
}

type CreateDistributorModelMasterAsyncResult struct {
	result *CreateDistributorModelMasterResult
	err    error
}

func NewCreateDistributorModelMasterResultFromDict(data map[string]interface{}) CreateDistributorModelMasterResult {
    return CreateDistributorModelMasterResult {
        Item: NewDistributorModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateDistributorModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateDistributorModelMasterResult) Pointer() *CreateDistributorModelMasterResult {
    return &p
}

type GetDistributorModelMasterResult struct {
    Item *DistributorModelMaster `json:"item"`
}

type GetDistributorModelMasterAsyncResult struct {
	result *GetDistributorModelMasterResult
	err    error
}

func NewGetDistributorModelMasterResultFromDict(data map[string]interface{}) GetDistributorModelMasterResult {
    return GetDistributorModelMasterResult {
        Item: NewDistributorModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetDistributorModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetDistributorModelMasterResult) Pointer() *GetDistributorModelMasterResult {
    return &p
}

type UpdateDistributorModelMasterResult struct {
    Item *DistributorModelMaster `json:"item"`
}

type UpdateDistributorModelMasterAsyncResult struct {
	result *UpdateDistributorModelMasterResult
	err    error
}

func NewUpdateDistributorModelMasterResultFromDict(data map[string]interface{}) UpdateDistributorModelMasterResult {
    return UpdateDistributorModelMasterResult {
        Item: NewDistributorModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateDistributorModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateDistributorModelMasterResult) Pointer() *UpdateDistributorModelMasterResult {
    return &p
}

type DeleteDistributorModelMasterResult struct {
    Item *DistributorModelMaster `json:"item"`
}

type DeleteDistributorModelMasterAsyncResult struct {
	result *DeleteDistributorModelMasterResult
	err    error
}

func NewDeleteDistributorModelMasterResultFromDict(data map[string]interface{}) DeleteDistributorModelMasterResult {
    return DeleteDistributorModelMasterResult {
        Item: NewDistributorModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteDistributorModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteDistributorModelMasterResult) Pointer() *DeleteDistributorModelMasterResult {
    return &p
}

type DescribeDistributorModelsResult struct {
    Items []DistributorModel `json:"items"`
}

type DescribeDistributorModelsAsyncResult struct {
	result *DescribeDistributorModelsResult
	err    error
}

func NewDescribeDistributorModelsResultFromDict(data map[string]interface{}) DescribeDistributorModelsResult {
    return DescribeDistributorModelsResult {
        Items: CastDistributorModels(core.CastArray(data["items"])),
    }
}

func (p DescribeDistributorModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastDistributorModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeDistributorModelsResult) Pointer() *DescribeDistributorModelsResult {
    return &p
}

type GetDistributorModelResult struct {
    Item *DistributorModel `json:"item"`
}

type GetDistributorModelAsyncResult struct {
	result *GetDistributorModelResult
	err    error
}

func NewGetDistributorModelResultFromDict(data map[string]interface{}) GetDistributorModelResult {
    return GetDistributorModelResult {
        Item: NewDistributorModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetDistributorModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetDistributorModelResult) Pointer() *GetDistributorModelResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentDistributorMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentDistributorMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentDistributorMasterResult struct {
    Item *CurrentDistributorMaster `json:"item"`
}

type GetCurrentDistributorMasterAsyncResult struct {
	result *GetCurrentDistributorMasterResult
	err    error
}

func NewGetCurrentDistributorMasterResultFromDict(data map[string]interface{}) GetCurrentDistributorMasterResult {
    return GetCurrentDistributorMasterResult {
        Item: NewCurrentDistributorMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentDistributorMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentDistributorMasterResult) Pointer() *GetCurrentDistributorMasterResult {
    return &p
}

type UpdateCurrentDistributorMasterResult struct {
    Item *CurrentDistributorMaster `json:"item"`
}

type UpdateCurrentDistributorMasterAsyncResult struct {
	result *UpdateCurrentDistributorMasterResult
	err    error
}

func NewUpdateCurrentDistributorMasterResultFromDict(data map[string]interface{}) UpdateCurrentDistributorMasterResult {
    return UpdateCurrentDistributorMasterResult {
        Item: NewCurrentDistributorMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentDistributorMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentDistributorMasterResult) Pointer() *UpdateCurrentDistributorMasterResult {
    return &p
}

type UpdateCurrentDistributorMasterFromGitHubResult struct {
    Item *CurrentDistributorMaster `json:"item"`
}

type UpdateCurrentDistributorMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentDistributorMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentDistributorMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentDistributorMasterFromGitHubResult {
    return UpdateCurrentDistributorMasterFromGitHubResult {
        Item: NewCurrentDistributorMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentDistributorMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentDistributorMasterFromGitHubResult) Pointer() *UpdateCurrentDistributorMasterFromGitHubResult {
    return &p
}

type DistributeResult struct {
    DistributeResource *DistributeResource `json:"distributeResource"`
    InboxNamespaceId *string `json:"inboxNamespaceId"`
    Result *string `json:"result"`
}

type DistributeAsyncResult struct {
	result *DistributeResult
	err    error
}

func NewDistributeResultFromDict(data map[string]interface{}) DistributeResult {
    return DistributeResult {
        DistributeResource: NewDistributeResourceFromDict(core.CastMap(data["distributeResource"])).Pointer(),
        InboxNamespaceId: core.CastString(data["inboxNamespaceId"]),
        Result: core.CastString(data["result"]),
    }
}

func (p DistributeResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "distributeResource": p.DistributeResource.ToDict(),
        "inboxNamespaceId": p.InboxNamespaceId,
        "result": p.Result,
    }
}

func (p DistributeResult) Pointer() *DistributeResult {
    return &p
}

type DistributeWithoutOverflowProcessResult struct {
    DistributeResource *DistributeResource `json:"distributeResource"`
    Result *string `json:"result"`
}

type DistributeWithoutOverflowProcessAsyncResult struct {
	result *DistributeWithoutOverflowProcessResult
	err    error
}

func NewDistributeWithoutOverflowProcessResultFromDict(data map[string]interface{}) DistributeWithoutOverflowProcessResult {
    return DistributeWithoutOverflowProcessResult {
        DistributeResource: NewDistributeResourceFromDict(core.CastMap(data["distributeResource"])).Pointer(),
        Result: core.CastString(data["result"]),
    }
}

func (p DistributeWithoutOverflowProcessResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "distributeResource": p.DistributeResource.ToDict(),
        "result": p.Result,
    }
}

func (p DistributeWithoutOverflowProcessResult) Pointer() *DistributeWithoutOverflowProcessResult {
    return &p
}

type RunStampTaskResult struct {
    ContextStack *string `json:"contextStack"`
    Result *string `json:"result"`
}

type RunStampTaskAsyncResult struct {
	result *RunStampTaskResult
	err    error
}

func NewRunStampTaskResultFromDict(data map[string]interface{}) RunStampTaskResult {
    return RunStampTaskResult {
        ContextStack: core.CastString(data["contextStack"]),
        Result: core.CastString(data["result"]),
    }
}

func (p RunStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "contextStack": p.ContextStack,
        "result": p.Result,
    }
}

func (p RunStampTaskResult) Pointer() *RunStampTaskResult {
    return &p
}

type RunStampSheetResult struct {
    Result *string `json:"result"`
}

type RunStampSheetAsyncResult struct {
	result *RunStampSheetResult
	err    error
}

func NewRunStampSheetResultFromDict(data map[string]interface{}) RunStampSheetResult {
    return RunStampSheetResult {
        Result: core.CastString(data["result"]),
    }
}

func (p RunStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "result": p.Result,
    }
}

func (p RunStampSheetResult) Pointer() *RunStampSheetResult {
    return &p
}

type RunStampSheetExpressResult struct {
    TaskResults []string `json:"taskResults"`
    SheetResult *string `json:"sheetResult"`
}

type RunStampSheetExpressAsyncResult struct {
	result *RunStampSheetExpressResult
	err    error
}

func NewRunStampSheetExpressResultFromDict(data map[string]interface{}) RunStampSheetExpressResult {
    return RunStampSheetExpressResult {
        TaskResults: core.CastStrings(core.CastArray(data["taskResults"])),
        SheetResult: core.CastString(data["sheetResult"]),
    }
}

func (p RunStampSheetExpressResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "taskResults": core.CastStringsFromDict(
            p.TaskResults,
        ),
        "sheetResult": p.SheetResult,
    }
}

func (p RunStampSheetExpressResult) Pointer() *RunStampSheetExpressResult {
    return &p
}

type RunStampTaskWithoutNamespaceResult struct {
    ContextStack *string `json:"contextStack"`
    Result *string `json:"result"`
}

type RunStampTaskWithoutNamespaceAsyncResult struct {
	result *RunStampTaskWithoutNamespaceResult
	err    error
}

func NewRunStampTaskWithoutNamespaceResultFromDict(data map[string]interface{}) RunStampTaskWithoutNamespaceResult {
    return RunStampTaskWithoutNamespaceResult {
        ContextStack: core.CastString(data["contextStack"]),
        Result: core.CastString(data["result"]),
    }
}

func (p RunStampTaskWithoutNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "contextStack": p.ContextStack,
        "result": p.Result,
    }
}

func (p RunStampTaskWithoutNamespaceResult) Pointer() *RunStampTaskWithoutNamespaceResult {
    return &p
}

type RunStampSheetWithoutNamespaceResult struct {
    Result *string `json:"result"`
}

type RunStampSheetWithoutNamespaceAsyncResult struct {
	result *RunStampSheetWithoutNamespaceResult
	err    error
}

func NewRunStampSheetWithoutNamespaceResultFromDict(data map[string]interface{}) RunStampSheetWithoutNamespaceResult {
    return RunStampSheetWithoutNamespaceResult {
        Result: core.CastString(data["result"]),
    }
}

func (p RunStampSheetWithoutNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "result": p.Result,
    }
}

func (p RunStampSheetWithoutNamespaceResult) Pointer() *RunStampSheetWithoutNamespaceResult {
    return &p
}

type RunStampSheetExpressWithoutNamespaceResult struct {
    TaskResults []string `json:"taskResults"`
    SheetResult *string `json:"sheetResult"`
}

type RunStampSheetExpressWithoutNamespaceAsyncResult struct {
	result *RunStampSheetExpressWithoutNamespaceResult
	err    error
}

func NewRunStampSheetExpressWithoutNamespaceResultFromDict(data map[string]interface{}) RunStampSheetExpressWithoutNamespaceResult {
    return RunStampSheetExpressWithoutNamespaceResult {
        TaskResults: core.CastStrings(core.CastArray(data["taskResults"])),
        SheetResult: core.CastString(data["sheetResult"]),
    }
}

func (p RunStampSheetExpressWithoutNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "taskResults": core.CastStringsFromDict(
            p.TaskResults,
        ),
        "sheetResult": p.SheetResult,
    }
}

func (p RunStampSheetExpressWithoutNamespaceResult) Pointer() *RunStampSheetExpressWithoutNamespaceResult {
    return &p
}