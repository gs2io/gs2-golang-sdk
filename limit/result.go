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

package limit

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

type DescribeCountersResult struct {
    Items []Counter `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeCountersAsyncResult struct {
	result *DescribeCountersResult
	err    error
}

func NewDescribeCountersResultFromDict(data map[string]interface{}) DescribeCountersResult {
    return DescribeCountersResult {
        Items: CastCounters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeCountersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCountersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeCountersResult) Pointer() *DescribeCountersResult {
    return &p
}

type DescribeCountersByUserIdResult struct {
    Items []Counter `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeCountersByUserIdAsyncResult struct {
	result *DescribeCountersByUserIdResult
	err    error
}

func NewDescribeCountersByUserIdResultFromDict(data map[string]interface{}) DescribeCountersByUserIdResult {
    return DescribeCountersByUserIdResult {
        Items: CastCounters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeCountersByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCountersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeCountersByUserIdResult) Pointer() *DescribeCountersByUserIdResult {
    return &p
}

type GetCounterResult struct {
    Item *Counter `json:"item"`
}

type GetCounterAsyncResult struct {
	result *GetCounterResult
	err    error
}

func NewGetCounterResultFromDict(data map[string]interface{}) GetCounterResult {
    return GetCounterResult {
        Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCounterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCounterResult) Pointer() *GetCounterResult {
    return &p
}

type GetCounterByUserIdResult struct {
    Item *Counter `json:"item"`
}

type GetCounterByUserIdAsyncResult struct {
	result *GetCounterByUserIdResult
	err    error
}

func NewGetCounterByUserIdResultFromDict(data map[string]interface{}) GetCounterByUserIdResult {
    return GetCounterByUserIdResult {
        Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCounterByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCounterByUserIdResult) Pointer() *GetCounterByUserIdResult {
    return &p
}

type CountUpResult struct {
    Item *Counter `json:"item"`
}

type CountUpAsyncResult struct {
	result *CountUpResult
	err    error
}

func NewCountUpResultFromDict(data map[string]interface{}) CountUpResult {
    return CountUpResult {
        Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CountUpResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CountUpResult) Pointer() *CountUpResult {
    return &p
}

type CountUpByUserIdResult struct {
    Item *Counter `json:"item"`
}

type CountUpByUserIdAsyncResult struct {
	result *CountUpByUserIdResult
	err    error
}

func NewCountUpByUserIdResultFromDict(data map[string]interface{}) CountUpByUserIdResult {
    return CountUpByUserIdResult {
        Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CountUpByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CountUpByUserIdResult) Pointer() *CountUpByUserIdResult {
    return &p
}

type DeleteCounterByUserIdResult struct {
    Item *Counter `json:"item"`
}

type DeleteCounterByUserIdAsyncResult struct {
	result *DeleteCounterByUserIdResult
	err    error
}

func NewDeleteCounterByUserIdResultFromDict(data map[string]interface{}) DeleteCounterByUserIdResult {
    return DeleteCounterByUserIdResult {
        Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteCounterByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteCounterByUserIdResult) Pointer() *DeleteCounterByUserIdResult {
    return &p
}

type CountUpByStampTaskResult struct {
    Item *Counter `json:"item"`
    NewContextStack *string `json:"newContextStack"`
}

type CountUpByStampTaskAsyncResult struct {
	result *CountUpByStampTaskResult
	err    error
}

func NewCountUpByStampTaskResultFromDict(data map[string]interface{}) CountUpByStampTaskResult {
    return CountUpByStampTaskResult {
        Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p CountUpByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p CountUpByStampTaskResult) Pointer() *CountUpByStampTaskResult {
    return &p
}

type DeleteByStampSheetResult struct {
    Item *Counter `json:"item"`
}

type DeleteByStampSheetAsyncResult struct {
	result *DeleteByStampSheetResult
	err    error
}

func NewDeleteByStampSheetResultFromDict(data map[string]interface{}) DeleteByStampSheetResult {
    return DeleteByStampSheetResult {
        Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteByStampSheetResult) Pointer() *DeleteByStampSheetResult {
    return &p
}

type DescribeLimitModelMastersResult struct {
    Items []LimitModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeLimitModelMastersAsyncResult struct {
	result *DescribeLimitModelMastersResult
	err    error
}

func NewDescribeLimitModelMastersResultFromDict(data map[string]interface{}) DescribeLimitModelMastersResult {
    return DescribeLimitModelMastersResult {
        Items: CastLimitModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeLimitModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastLimitModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeLimitModelMastersResult) Pointer() *DescribeLimitModelMastersResult {
    return &p
}

type CreateLimitModelMasterResult struct {
    Item *LimitModelMaster `json:"item"`
}

type CreateLimitModelMasterAsyncResult struct {
	result *CreateLimitModelMasterResult
	err    error
}

func NewCreateLimitModelMasterResultFromDict(data map[string]interface{}) CreateLimitModelMasterResult {
    return CreateLimitModelMasterResult {
        Item: NewLimitModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateLimitModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateLimitModelMasterResult) Pointer() *CreateLimitModelMasterResult {
    return &p
}

type GetLimitModelMasterResult struct {
    Item *LimitModelMaster `json:"item"`
}

type GetLimitModelMasterAsyncResult struct {
	result *GetLimitModelMasterResult
	err    error
}

func NewGetLimitModelMasterResultFromDict(data map[string]interface{}) GetLimitModelMasterResult {
    return GetLimitModelMasterResult {
        Item: NewLimitModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetLimitModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetLimitModelMasterResult) Pointer() *GetLimitModelMasterResult {
    return &p
}

type UpdateLimitModelMasterResult struct {
    Item *LimitModelMaster `json:"item"`
}

type UpdateLimitModelMasterAsyncResult struct {
	result *UpdateLimitModelMasterResult
	err    error
}

func NewUpdateLimitModelMasterResultFromDict(data map[string]interface{}) UpdateLimitModelMasterResult {
    return UpdateLimitModelMasterResult {
        Item: NewLimitModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateLimitModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateLimitModelMasterResult) Pointer() *UpdateLimitModelMasterResult {
    return &p
}

type DeleteLimitModelMasterResult struct {
    Item *LimitModelMaster `json:"item"`
}

type DeleteLimitModelMasterAsyncResult struct {
	result *DeleteLimitModelMasterResult
	err    error
}

func NewDeleteLimitModelMasterResultFromDict(data map[string]interface{}) DeleteLimitModelMasterResult {
    return DeleteLimitModelMasterResult {
        Item: NewLimitModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteLimitModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteLimitModelMasterResult) Pointer() *DeleteLimitModelMasterResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentLimitMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentLimitMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentLimitMasterResult struct {
    Item *CurrentLimitMaster `json:"item"`
}

type GetCurrentLimitMasterAsyncResult struct {
	result *GetCurrentLimitMasterResult
	err    error
}

func NewGetCurrentLimitMasterResultFromDict(data map[string]interface{}) GetCurrentLimitMasterResult {
    return GetCurrentLimitMasterResult {
        Item: NewCurrentLimitMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentLimitMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentLimitMasterResult) Pointer() *GetCurrentLimitMasterResult {
    return &p
}

type UpdateCurrentLimitMasterResult struct {
    Item *CurrentLimitMaster `json:"item"`
}

type UpdateCurrentLimitMasterAsyncResult struct {
	result *UpdateCurrentLimitMasterResult
	err    error
}

func NewUpdateCurrentLimitMasterResultFromDict(data map[string]interface{}) UpdateCurrentLimitMasterResult {
    return UpdateCurrentLimitMasterResult {
        Item: NewCurrentLimitMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentLimitMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentLimitMasterResult) Pointer() *UpdateCurrentLimitMasterResult {
    return &p
}

type UpdateCurrentLimitMasterFromGitHubResult struct {
    Item *CurrentLimitMaster `json:"item"`
}

type UpdateCurrentLimitMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentLimitMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentLimitMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentLimitMasterFromGitHubResult {
    return UpdateCurrentLimitMasterFromGitHubResult {
        Item: NewCurrentLimitMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentLimitMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentLimitMasterFromGitHubResult) Pointer() *UpdateCurrentLimitMasterFromGitHubResult {
    return &p
}

type DescribeLimitModelsResult struct {
    Items []LimitModel `json:"items"`
}

type DescribeLimitModelsAsyncResult struct {
	result *DescribeLimitModelsResult
	err    error
}

func NewDescribeLimitModelsResultFromDict(data map[string]interface{}) DescribeLimitModelsResult {
    return DescribeLimitModelsResult {
        Items: CastLimitModels(core.CastArray(data["items"])),
    }
}

func (p DescribeLimitModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastLimitModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeLimitModelsResult) Pointer() *DescribeLimitModelsResult {
    return &p
}

type GetLimitModelResult struct {
    Item *LimitModel `json:"item"`
}

type GetLimitModelAsyncResult struct {
	result *GetLimitModelResult
	err    error
}

func NewGetLimitModelResultFromDict(data map[string]interface{}) GetLimitModelResult {
    return GetLimitModelResult {
        Item: NewLimitModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetLimitModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetLimitModelResult) Pointer() *GetLimitModelResult {
    return &p
}