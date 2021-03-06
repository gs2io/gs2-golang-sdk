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

package mission

import "core"

type DescribeCompletesResult struct {
    Items []Complete `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeCompletesAsyncResult struct {
	result *DescribeCompletesResult
	err    error
}

func NewDescribeCompletesResultFromDict(data map[string]interface{}) DescribeCompletesResult {
    return DescribeCompletesResult {
        Items: CastCompletes(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeCompletesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCompletesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeCompletesResult) Pointer() *DescribeCompletesResult {
    return &p
}

type DescribeCompletesByUserIdResult struct {
    Items []Complete `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeCompletesByUserIdAsyncResult struct {
	result *DescribeCompletesByUserIdResult
	err    error
}

func NewDescribeCompletesByUserIdResultFromDict(data map[string]interface{}) DescribeCompletesByUserIdResult {
    return DescribeCompletesByUserIdResult {
        Items: CastCompletes(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeCompletesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCompletesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeCompletesByUserIdResult) Pointer() *DescribeCompletesByUserIdResult {
    return &p
}

type CompleteResult struct {
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type CompleteAsyncResult struct {
	result *CompleteResult
	err    error
}

func NewCompleteResultFromDict(data map[string]interface{}) CompleteResult {
    return CompleteResult {
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p CompleteResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p CompleteResult) Pointer() *CompleteResult {
    return &p
}

type CompleteByUserIdResult struct {
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type CompleteByUserIdAsyncResult struct {
	result *CompleteByUserIdResult
	err    error
}

func NewCompleteByUserIdResultFromDict(data map[string]interface{}) CompleteByUserIdResult {
    return CompleteByUserIdResult {
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p CompleteByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p CompleteByUserIdResult) Pointer() *CompleteByUserIdResult {
    return &p
}

type ReceiveByUserIdResult struct {
    Item *Complete `json:"item"`
}

type ReceiveByUserIdAsyncResult struct {
	result *ReceiveByUserIdResult
	err    error
}

func NewReceiveByUserIdResultFromDict(data map[string]interface{}) ReceiveByUserIdResult {
    return ReceiveByUserIdResult {
        Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p ReceiveByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p ReceiveByUserIdResult) Pointer() *ReceiveByUserIdResult {
    return &p
}

type GetCompleteResult struct {
    Item *Complete `json:"item"`
}

type GetCompleteAsyncResult struct {
	result *GetCompleteResult
	err    error
}

func NewGetCompleteResultFromDict(data map[string]interface{}) GetCompleteResult {
    return GetCompleteResult {
        Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCompleteResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCompleteResult) Pointer() *GetCompleteResult {
    return &p
}

type GetCompleteByUserIdResult struct {
    Item *Complete `json:"item"`
}

type GetCompleteByUserIdAsyncResult struct {
	result *GetCompleteByUserIdResult
	err    error
}

func NewGetCompleteByUserIdResultFromDict(data map[string]interface{}) GetCompleteByUserIdResult {
    return GetCompleteByUserIdResult {
        Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCompleteByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCompleteByUserIdResult) Pointer() *GetCompleteByUserIdResult {
    return &p
}

type DeleteCompleteByUserIdResult struct {
    Item *Complete `json:"item"`
}

type DeleteCompleteByUserIdAsyncResult struct {
	result *DeleteCompleteByUserIdResult
	err    error
}

func NewDeleteCompleteByUserIdResultFromDict(data map[string]interface{}) DeleteCompleteByUserIdResult {
    return DeleteCompleteByUserIdResult {
        Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteCompleteByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteCompleteByUserIdResult) Pointer() *DeleteCompleteByUserIdResult {
    return &p
}

type ReceiveByStampTaskResult struct {
    Item *Complete `json:"item"`
    NewContextStack *string `json:"newContextStack"`
}

type ReceiveByStampTaskAsyncResult struct {
	result *ReceiveByStampTaskResult
	err    error
}

func NewReceiveByStampTaskResultFromDict(data map[string]interface{}) ReceiveByStampTaskResult {
    return ReceiveByStampTaskResult {
        Item: NewCompleteFromDict(core.CastMap(data["item"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p ReceiveByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p ReceiveByStampTaskResult) Pointer() *ReceiveByStampTaskResult {
    return &p
}

type DescribeCounterModelMastersResult struct {
    Items []CounterModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeCounterModelMastersAsyncResult struct {
	result *DescribeCounterModelMastersResult
	err    error
}

func NewDescribeCounterModelMastersResultFromDict(data map[string]interface{}) DescribeCounterModelMastersResult {
    return DescribeCounterModelMastersResult {
        Items: CastCounterModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeCounterModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCounterModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeCounterModelMastersResult) Pointer() *DescribeCounterModelMastersResult {
    return &p
}

type CreateCounterModelMasterResult struct {
    Item *CounterModelMaster `json:"item"`
}

type CreateCounterModelMasterAsyncResult struct {
	result *CreateCounterModelMasterResult
	err    error
}

func NewCreateCounterModelMasterResultFromDict(data map[string]interface{}) CreateCounterModelMasterResult {
    return CreateCounterModelMasterResult {
        Item: NewCounterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateCounterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateCounterModelMasterResult) Pointer() *CreateCounterModelMasterResult {
    return &p
}

type GetCounterModelMasterResult struct {
    Item *CounterModelMaster `json:"item"`
}

type GetCounterModelMasterAsyncResult struct {
	result *GetCounterModelMasterResult
	err    error
}

func NewGetCounterModelMasterResultFromDict(data map[string]interface{}) GetCounterModelMasterResult {
    return GetCounterModelMasterResult {
        Item: NewCounterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCounterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCounterModelMasterResult) Pointer() *GetCounterModelMasterResult {
    return &p
}

type UpdateCounterModelMasterResult struct {
    Item *CounterModelMaster `json:"item"`
}

type UpdateCounterModelMasterAsyncResult struct {
	result *UpdateCounterModelMasterResult
	err    error
}

func NewUpdateCounterModelMasterResultFromDict(data map[string]interface{}) UpdateCounterModelMasterResult {
    return UpdateCounterModelMasterResult {
        Item: NewCounterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCounterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCounterModelMasterResult) Pointer() *UpdateCounterModelMasterResult {
    return &p
}

type DeleteCounterModelMasterResult struct {
    Item *CounterModelMaster `json:"item"`
}

type DeleteCounterModelMasterAsyncResult struct {
	result *DeleteCounterModelMasterResult
	err    error
}

func NewDeleteCounterModelMasterResultFromDict(data map[string]interface{}) DeleteCounterModelMasterResult {
    return DeleteCounterModelMasterResult {
        Item: NewCounterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteCounterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteCounterModelMasterResult) Pointer() *DeleteCounterModelMasterResult {
    return &p
}

type DescribeMissionGroupModelMastersResult struct {
    Items []MissionGroupModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeMissionGroupModelMastersAsyncResult struct {
	result *DescribeMissionGroupModelMastersResult
	err    error
}

func NewDescribeMissionGroupModelMastersResultFromDict(data map[string]interface{}) DescribeMissionGroupModelMastersResult {
    return DescribeMissionGroupModelMastersResult {
        Items: CastMissionGroupModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeMissionGroupModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastMissionGroupModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeMissionGroupModelMastersResult) Pointer() *DescribeMissionGroupModelMastersResult {
    return &p
}

type CreateMissionGroupModelMasterResult struct {
    Item *MissionGroupModelMaster `json:"item"`
}

type CreateMissionGroupModelMasterAsyncResult struct {
	result *CreateMissionGroupModelMasterResult
	err    error
}

func NewCreateMissionGroupModelMasterResultFromDict(data map[string]interface{}) CreateMissionGroupModelMasterResult {
    return CreateMissionGroupModelMasterResult {
        Item: NewMissionGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateMissionGroupModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateMissionGroupModelMasterResult) Pointer() *CreateMissionGroupModelMasterResult {
    return &p
}

type GetMissionGroupModelMasterResult struct {
    Item *MissionGroupModelMaster `json:"item"`
}

type GetMissionGroupModelMasterAsyncResult struct {
	result *GetMissionGroupModelMasterResult
	err    error
}

func NewGetMissionGroupModelMasterResultFromDict(data map[string]interface{}) GetMissionGroupModelMasterResult {
    return GetMissionGroupModelMasterResult {
        Item: NewMissionGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetMissionGroupModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetMissionGroupModelMasterResult) Pointer() *GetMissionGroupModelMasterResult {
    return &p
}

type UpdateMissionGroupModelMasterResult struct {
    Item *MissionGroupModelMaster `json:"item"`
}

type UpdateMissionGroupModelMasterAsyncResult struct {
	result *UpdateMissionGroupModelMasterResult
	err    error
}

func NewUpdateMissionGroupModelMasterResultFromDict(data map[string]interface{}) UpdateMissionGroupModelMasterResult {
    return UpdateMissionGroupModelMasterResult {
        Item: NewMissionGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateMissionGroupModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateMissionGroupModelMasterResult) Pointer() *UpdateMissionGroupModelMasterResult {
    return &p
}

type DeleteMissionGroupModelMasterResult struct {
    Item *MissionGroupModelMaster `json:"item"`
}

type DeleteMissionGroupModelMasterAsyncResult struct {
	result *DeleteMissionGroupModelMasterResult
	err    error
}

func NewDeleteMissionGroupModelMasterResultFromDict(data map[string]interface{}) DeleteMissionGroupModelMasterResult {
    return DeleteMissionGroupModelMasterResult {
        Item: NewMissionGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteMissionGroupModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteMissionGroupModelMasterResult) Pointer() *DeleteMissionGroupModelMasterResult {
    return &p
}

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

type IncreaseCounterByUserIdResult struct {
    Item *Counter `json:"item"`
}

type IncreaseCounterByUserIdAsyncResult struct {
	result *IncreaseCounterByUserIdResult
	err    error
}

func NewIncreaseCounterByUserIdResultFromDict(data map[string]interface{}) IncreaseCounterByUserIdResult {
    return IncreaseCounterByUserIdResult {
        Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p IncreaseCounterByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p IncreaseCounterByUserIdResult) Pointer() *IncreaseCounterByUserIdResult {
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

type IncreaseByStampSheetResult struct {
    Item *Counter `json:"item"`
}

type IncreaseByStampSheetAsyncResult struct {
	result *IncreaseByStampSheetResult
	err    error
}

func NewIncreaseByStampSheetResultFromDict(data map[string]interface{}) IncreaseByStampSheetResult {
    return IncreaseByStampSheetResult {
        Item: NewCounterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p IncreaseByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p IncreaseByStampSheetResult) Pointer() *IncreaseByStampSheetResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentMissionMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentMissionMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentMissionMasterResult struct {
    Item *CurrentMissionMaster `json:"item"`
}

type GetCurrentMissionMasterAsyncResult struct {
	result *GetCurrentMissionMasterResult
	err    error
}

func NewGetCurrentMissionMasterResultFromDict(data map[string]interface{}) GetCurrentMissionMasterResult {
    return GetCurrentMissionMasterResult {
        Item: NewCurrentMissionMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentMissionMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentMissionMasterResult) Pointer() *GetCurrentMissionMasterResult {
    return &p
}

type UpdateCurrentMissionMasterResult struct {
    Item *CurrentMissionMaster `json:"item"`
}

type UpdateCurrentMissionMasterAsyncResult struct {
	result *UpdateCurrentMissionMasterResult
	err    error
}

func NewUpdateCurrentMissionMasterResultFromDict(data map[string]interface{}) UpdateCurrentMissionMasterResult {
    return UpdateCurrentMissionMasterResult {
        Item: NewCurrentMissionMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentMissionMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentMissionMasterResult) Pointer() *UpdateCurrentMissionMasterResult {
    return &p
}

type UpdateCurrentMissionMasterFromGitHubResult struct {
    Item *CurrentMissionMaster `json:"item"`
}

type UpdateCurrentMissionMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentMissionMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentMissionMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentMissionMasterFromGitHubResult {
    return UpdateCurrentMissionMasterFromGitHubResult {
        Item: NewCurrentMissionMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentMissionMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentMissionMasterFromGitHubResult) Pointer() *UpdateCurrentMissionMasterFromGitHubResult {
    return &p
}

type DescribeCounterModelsResult struct {
    Items []CounterModel `json:"items"`
}

type DescribeCounterModelsAsyncResult struct {
	result *DescribeCounterModelsResult
	err    error
}

func NewDescribeCounterModelsResultFromDict(data map[string]interface{}) DescribeCounterModelsResult {
    return DescribeCounterModelsResult {
        Items: CastCounterModels(core.CastArray(data["items"])),
    }
}

func (p DescribeCounterModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCounterModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeCounterModelsResult) Pointer() *DescribeCounterModelsResult {
    return &p
}

type GetCounterModelResult struct {
    Item *CounterModel `json:"item"`
}

type GetCounterModelAsyncResult struct {
	result *GetCounterModelResult
	err    error
}

func NewGetCounterModelResultFromDict(data map[string]interface{}) GetCounterModelResult {
    return GetCounterModelResult {
        Item: NewCounterModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCounterModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCounterModelResult) Pointer() *GetCounterModelResult {
    return &p
}

type DescribeMissionGroupModelsResult struct {
    Items []MissionGroupModel `json:"items"`
}

type DescribeMissionGroupModelsAsyncResult struct {
	result *DescribeMissionGroupModelsResult
	err    error
}

func NewDescribeMissionGroupModelsResultFromDict(data map[string]interface{}) DescribeMissionGroupModelsResult {
    return DescribeMissionGroupModelsResult {
        Items: CastMissionGroupModels(core.CastArray(data["items"])),
    }
}

func (p DescribeMissionGroupModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastMissionGroupModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeMissionGroupModelsResult) Pointer() *DescribeMissionGroupModelsResult {
    return &p
}

type GetMissionGroupModelResult struct {
    Item *MissionGroupModel `json:"item"`
}

type GetMissionGroupModelAsyncResult struct {
	result *GetMissionGroupModelResult
	err    error
}

func NewGetMissionGroupModelResultFromDict(data map[string]interface{}) GetMissionGroupModelResult {
    return GetMissionGroupModelResult {
        Item: NewMissionGroupModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetMissionGroupModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetMissionGroupModelResult) Pointer() *GetMissionGroupModelResult {
    return &p
}

type DescribeMissionTaskModelsResult struct {
    Items []MissionTaskModel `json:"items"`
}

type DescribeMissionTaskModelsAsyncResult struct {
	result *DescribeMissionTaskModelsResult
	err    error
}

func NewDescribeMissionTaskModelsResultFromDict(data map[string]interface{}) DescribeMissionTaskModelsResult {
    return DescribeMissionTaskModelsResult {
        Items: CastMissionTaskModels(core.CastArray(data["items"])),
    }
}

func (p DescribeMissionTaskModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastMissionTaskModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeMissionTaskModelsResult) Pointer() *DescribeMissionTaskModelsResult {
    return &p
}

type GetMissionTaskModelResult struct {
    Item *MissionTaskModel `json:"item"`
}

type GetMissionTaskModelAsyncResult struct {
	result *GetMissionTaskModelResult
	err    error
}

func NewGetMissionTaskModelResultFromDict(data map[string]interface{}) GetMissionTaskModelResult {
    return GetMissionTaskModelResult {
        Item: NewMissionTaskModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetMissionTaskModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetMissionTaskModelResult) Pointer() *GetMissionTaskModelResult {
    return &p
}

type DescribeMissionTaskModelMastersResult struct {
    Items []MissionTaskModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeMissionTaskModelMastersAsyncResult struct {
	result *DescribeMissionTaskModelMastersResult
	err    error
}

func NewDescribeMissionTaskModelMastersResultFromDict(data map[string]interface{}) DescribeMissionTaskModelMastersResult {
    return DescribeMissionTaskModelMastersResult {
        Items: CastMissionTaskModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeMissionTaskModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastMissionTaskModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeMissionTaskModelMastersResult) Pointer() *DescribeMissionTaskModelMastersResult {
    return &p
}

type CreateMissionTaskModelMasterResult struct {
    Item *MissionTaskModelMaster `json:"item"`
}

type CreateMissionTaskModelMasterAsyncResult struct {
	result *CreateMissionTaskModelMasterResult
	err    error
}

func NewCreateMissionTaskModelMasterResultFromDict(data map[string]interface{}) CreateMissionTaskModelMasterResult {
    return CreateMissionTaskModelMasterResult {
        Item: NewMissionTaskModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateMissionTaskModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateMissionTaskModelMasterResult) Pointer() *CreateMissionTaskModelMasterResult {
    return &p
}

type GetMissionTaskModelMasterResult struct {
    Item *MissionTaskModelMaster `json:"item"`
}

type GetMissionTaskModelMasterAsyncResult struct {
	result *GetMissionTaskModelMasterResult
	err    error
}

func NewGetMissionTaskModelMasterResultFromDict(data map[string]interface{}) GetMissionTaskModelMasterResult {
    return GetMissionTaskModelMasterResult {
        Item: NewMissionTaskModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetMissionTaskModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetMissionTaskModelMasterResult) Pointer() *GetMissionTaskModelMasterResult {
    return &p
}

type UpdateMissionTaskModelMasterResult struct {
    Item *MissionTaskModelMaster `json:"item"`
}

type UpdateMissionTaskModelMasterAsyncResult struct {
	result *UpdateMissionTaskModelMasterResult
	err    error
}

func NewUpdateMissionTaskModelMasterResultFromDict(data map[string]interface{}) UpdateMissionTaskModelMasterResult {
    return UpdateMissionTaskModelMasterResult {
        Item: NewMissionTaskModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateMissionTaskModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateMissionTaskModelMasterResult) Pointer() *UpdateMissionTaskModelMasterResult {
    return &p
}

type DeleteMissionTaskModelMasterResult struct {
    Item *MissionTaskModelMaster `json:"item"`
}

type DeleteMissionTaskModelMasterAsyncResult struct {
	result *DeleteMissionTaskModelMasterResult
	err    error
}

func NewDeleteMissionTaskModelMasterResultFromDict(data map[string]interface{}) DeleteMissionTaskModelMasterResult {
    return DeleteMissionTaskModelMasterResult {
        Item: NewMissionTaskModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteMissionTaskModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteMissionTaskModelMasterResult) Pointer() *DeleteMissionTaskModelMasterResult {
    return &p
}