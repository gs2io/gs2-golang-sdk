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

type DescribeQuestGroupModelMastersResult struct {
    Items []QuestGroupModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeQuestGroupModelMastersAsyncResult struct {
	result *DescribeQuestGroupModelMastersResult
	err    error
}

func NewDescribeQuestGroupModelMastersResultFromDict(data map[string]interface{}) DescribeQuestGroupModelMastersResult {
    return DescribeQuestGroupModelMastersResult {
        Items: CastQuestGroupModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeQuestGroupModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    Item *QuestGroupModelMaster `json:"item"`
}

type CreateQuestGroupModelMasterAsyncResult struct {
	result *CreateQuestGroupModelMasterResult
	err    error
}

func NewCreateQuestGroupModelMasterResultFromDict(data map[string]interface{}) CreateQuestGroupModelMasterResult {
    return CreateQuestGroupModelMasterResult {
        Item: NewQuestGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateQuestGroupModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateQuestGroupModelMasterResult) Pointer() *CreateQuestGroupModelMasterResult {
    return &p
}

type GetQuestGroupModelMasterResult struct {
    Item *QuestGroupModelMaster `json:"item"`
}

type GetQuestGroupModelMasterAsyncResult struct {
	result *GetQuestGroupModelMasterResult
	err    error
}

func NewGetQuestGroupModelMasterResultFromDict(data map[string]interface{}) GetQuestGroupModelMasterResult {
    return GetQuestGroupModelMasterResult {
        Item: NewQuestGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetQuestGroupModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetQuestGroupModelMasterResult) Pointer() *GetQuestGroupModelMasterResult {
    return &p
}

type UpdateQuestGroupModelMasterResult struct {
    Item *QuestGroupModelMaster `json:"item"`
}

type UpdateQuestGroupModelMasterAsyncResult struct {
	result *UpdateQuestGroupModelMasterResult
	err    error
}

func NewUpdateQuestGroupModelMasterResultFromDict(data map[string]interface{}) UpdateQuestGroupModelMasterResult {
    return UpdateQuestGroupModelMasterResult {
        Item: NewQuestGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateQuestGroupModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateQuestGroupModelMasterResult) Pointer() *UpdateQuestGroupModelMasterResult {
    return &p
}

type DeleteQuestGroupModelMasterResult struct {
    Item *QuestGroupModelMaster `json:"item"`
}

type DeleteQuestGroupModelMasterAsyncResult struct {
	result *DeleteQuestGroupModelMasterResult
	err    error
}

func NewDeleteQuestGroupModelMasterResultFromDict(data map[string]interface{}) DeleteQuestGroupModelMasterResult {
    return DeleteQuestGroupModelMasterResult {
        Item: NewQuestGroupModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteQuestGroupModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteQuestGroupModelMasterResult) Pointer() *DeleteQuestGroupModelMasterResult {
    return &p
}

type DescribeQuestModelMastersResult struct {
    Items []QuestModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeQuestModelMastersAsyncResult struct {
	result *DescribeQuestModelMastersResult
	err    error
}

func NewDescribeQuestModelMastersResultFromDict(data map[string]interface{}) DescribeQuestModelMastersResult {
    return DescribeQuestModelMastersResult {
        Items: CastQuestModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeQuestModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    Item *QuestModelMaster `json:"item"`
}

type CreateQuestModelMasterAsyncResult struct {
	result *CreateQuestModelMasterResult
	err    error
}

func NewCreateQuestModelMasterResultFromDict(data map[string]interface{}) CreateQuestModelMasterResult {
    return CreateQuestModelMasterResult {
        Item: NewQuestModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateQuestModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateQuestModelMasterResult) Pointer() *CreateQuestModelMasterResult {
    return &p
}

type GetQuestModelMasterResult struct {
    Item *QuestModelMaster `json:"item"`
}

type GetQuestModelMasterAsyncResult struct {
	result *GetQuestModelMasterResult
	err    error
}

func NewGetQuestModelMasterResultFromDict(data map[string]interface{}) GetQuestModelMasterResult {
    return GetQuestModelMasterResult {
        Item: NewQuestModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetQuestModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetQuestModelMasterResult) Pointer() *GetQuestModelMasterResult {
    return &p
}

type UpdateQuestModelMasterResult struct {
    Item *QuestModelMaster `json:"item"`
}

type UpdateQuestModelMasterAsyncResult struct {
	result *UpdateQuestModelMasterResult
	err    error
}

func NewUpdateQuestModelMasterResultFromDict(data map[string]interface{}) UpdateQuestModelMasterResult {
    return UpdateQuestModelMasterResult {
        Item: NewQuestModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateQuestModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateQuestModelMasterResult) Pointer() *UpdateQuestModelMasterResult {
    return &p
}

type DeleteQuestModelMasterResult struct {
    Item *QuestModelMaster `json:"item"`
}

type DeleteQuestModelMasterAsyncResult struct {
	result *DeleteQuestModelMasterResult
	err    error
}

func NewDeleteQuestModelMasterResultFromDict(data map[string]interface{}) DeleteQuestModelMasterResult {
    return DeleteQuestModelMasterResult {
        Item: NewQuestModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteQuestModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteQuestModelMasterResult) Pointer() *DeleteQuestModelMasterResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentQuestMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentQuestMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentQuestMasterResult struct {
    Item *CurrentQuestMaster `json:"item"`
}

type GetCurrentQuestMasterAsyncResult struct {
	result *GetCurrentQuestMasterResult
	err    error
}

func NewGetCurrentQuestMasterResultFromDict(data map[string]interface{}) GetCurrentQuestMasterResult {
    return GetCurrentQuestMasterResult {
        Item: NewCurrentQuestMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentQuestMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentQuestMasterResult) Pointer() *GetCurrentQuestMasterResult {
    return &p
}

type UpdateCurrentQuestMasterResult struct {
    Item *CurrentQuestMaster `json:"item"`
}

type UpdateCurrentQuestMasterAsyncResult struct {
	result *UpdateCurrentQuestMasterResult
	err    error
}

func NewUpdateCurrentQuestMasterResultFromDict(data map[string]interface{}) UpdateCurrentQuestMasterResult {
    return UpdateCurrentQuestMasterResult {
        Item: NewCurrentQuestMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentQuestMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentQuestMasterResult) Pointer() *UpdateCurrentQuestMasterResult {
    return &p
}

type UpdateCurrentQuestMasterFromGitHubResult struct {
    Item *CurrentQuestMaster `json:"item"`
}

type UpdateCurrentQuestMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentQuestMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentQuestMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentQuestMasterFromGitHubResult {
    return UpdateCurrentQuestMasterFromGitHubResult {
        Item: NewCurrentQuestMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentQuestMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentQuestMasterFromGitHubResult) Pointer() *UpdateCurrentQuestMasterFromGitHubResult {
    return &p
}

type DescribeProgressesByUserIdResult struct {
    Items []Progress `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeProgressesByUserIdAsyncResult struct {
	result *DescribeProgressesByUserIdResult
	err    error
}

func NewDescribeProgressesByUserIdResultFromDict(data map[string]interface{}) DescribeProgressesByUserIdResult {
    return DescribeProgressesByUserIdResult {
        Items: CastProgresses(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeProgressesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    Item *Progress `json:"item"`
}

type CreateProgressByUserIdAsyncResult struct {
	result *CreateProgressByUserIdResult
	err    error
}

func NewCreateProgressByUserIdResultFromDict(data map[string]interface{}) CreateProgressByUserIdResult {
    return CreateProgressByUserIdResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateProgressByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateProgressByUserIdResult) Pointer() *CreateProgressByUserIdResult {
    return &p
}

type GetProgressResult struct {
    Item *Progress `json:"item"`
    QuestGroup *QuestGroupModel `json:"questGroup"`
    Quest *QuestModel `json:"quest"`
}

type GetProgressAsyncResult struct {
	result *GetProgressResult
	err    error
}

func NewGetProgressResultFromDict(data map[string]interface{}) GetProgressResult {
    return GetProgressResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
        QuestGroup: NewQuestGroupModelFromDict(core.CastMap(data["questGroup"])).Pointer(),
        Quest: NewQuestModelFromDict(core.CastMap(data["quest"])).Pointer(),
    }
}

func (p GetProgressResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "questGroup": p.QuestGroup.ToDict(),
        "quest": p.Quest.ToDict(),
    }
}

func (p GetProgressResult) Pointer() *GetProgressResult {
    return &p
}

type GetProgressByUserIdResult struct {
    Item *Progress `json:"item"`
    QuestGroup *QuestGroupModel `json:"questGroup"`
    Quest *QuestModel `json:"quest"`
}

type GetProgressByUserIdAsyncResult struct {
	result *GetProgressByUserIdResult
	err    error
}

func NewGetProgressByUserIdResultFromDict(data map[string]interface{}) GetProgressByUserIdResult {
    return GetProgressByUserIdResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
        QuestGroup: NewQuestGroupModelFromDict(core.CastMap(data["questGroup"])).Pointer(),
        Quest: NewQuestModelFromDict(core.CastMap(data["quest"])).Pointer(),
    }
}

func (p GetProgressByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "questGroup": p.QuestGroup.ToDict(),
        "quest": p.Quest.ToDict(),
    }
}

func (p GetProgressByUserIdResult) Pointer() *GetProgressByUserIdResult {
    return &p
}

type StartResult struct {
    Item *Progress `json:"item"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type StartAsyncResult struct {
	result *StartResult
	err    error
}

func NewStartResultFromDict(data map[string]interface{}) StartResult {
    return StartResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p StartResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p StartResult) Pointer() *StartResult {
    return &p
}

type StartByUserIdResult struct {
    Item *Progress `json:"item"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type StartByUserIdAsyncResult struct {
	result *StartByUserIdResult
	err    error
}

func NewStartByUserIdResultFromDict(data map[string]interface{}) StartByUserIdResult {
    return StartByUserIdResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p StartByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p StartByUserIdResult) Pointer() *StartByUserIdResult {
    return &p
}

type EndResult struct {
    Item *Progress `json:"item"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type EndAsyncResult struct {
	result *EndResult
	err    error
}

func NewEndResultFromDict(data map[string]interface{}) EndResult {
    return EndResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p EndResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p EndResult) Pointer() *EndResult {
    return &p
}

type EndByUserIdResult struct {
    Item *Progress `json:"item"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type EndByUserIdAsyncResult struct {
	result *EndByUserIdResult
	err    error
}

func NewEndByUserIdResultFromDict(data map[string]interface{}) EndByUserIdResult {
    return EndByUserIdResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p EndByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p EndByUserIdResult) Pointer() *EndByUserIdResult {
    return &p
}

type DeleteProgressResult struct {
    Item *Progress `json:"item"`
}

type DeleteProgressAsyncResult struct {
	result *DeleteProgressResult
	err    error
}

func NewDeleteProgressResultFromDict(data map[string]interface{}) DeleteProgressResult {
    return DeleteProgressResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteProgressResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteProgressResult) Pointer() *DeleteProgressResult {
    return &p
}

type DeleteProgressByUserIdResult struct {
    Item *Progress `json:"item"`
}

type DeleteProgressByUserIdAsyncResult struct {
	result *DeleteProgressByUserIdResult
	err    error
}

func NewDeleteProgressByUserIdResultFromDict(data map[string]interface{}) DeleteProgressByUserIdResult {
    return DeleteProgressByUserIdResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteProgressByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteProgressByUserIdResult) Pointer() *DeleteProgressByUserIdResult {
    return &p
}

type CreateProgressByStampSheetResult struct {
    Item *Progress `json:"item"`
}

type CreateProgressByStampSheetAsyncResult struct {
	result *CreateProgressByStampSheetResult
	err    error
}

func NewCreateProgressByStampSheetResultFromDict(data map[string]interface{}) CreateProgressByStampSheetResult {
    return CreateProgressByStampSheetResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateProgressByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateProgressByStampSheetResult) Pointer() *CreateProgressByStampSheetResult {
    return &p
}

type DeleteProgressByStampTaskResult struct {
    Item *Progress `json:"item"`
    NewContextStack *string `json:"newContextStack"`
}

type DeleteProgressByStampTaskAsyncResult struct {
	result *DeleteProgressByStampTaskResult
	err    error
}

func NewDeleteProgressByStampTaskResultFromDict(data map[string]interface{}) DeleteProgressByStampTaskResult {
    return DeleteProgressByStampTaskResult {
        Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p DeleteProgressByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p DeleteProgressByStampTaskResult) Pointer() *DeleteProgressByStampTaskResult {
    return &p
}

type DescribeCompletedQuestListsResult struct {
    Items []CompletedQuestList `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeCompletedQuestListsAsyncResult struct {
	result *DescribeCompletedQuestListsResult
	err    error
}

func NewDescribeCompletedQuestListsResultFromDict(data map[string]interface{}) DescribeCompletedQuestListsResult {
    return DescribeCompletedQuestListsResult {
        Items: CastCompletedQuestLists(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeCompletedQuestListsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCompletedQuestListsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeCompletedQuestListsResult) Pointer() *DescribeCompletedQuestListsResult {
    return &p
}

type DescribeCompletedQuestListsByUserIdResult struct {
    Items []CompletedQuestList `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeCompletedQuestListsByUserIdAsyncResult struct {
	result *DescribeCompletedQuestListsByUserIdResult
	err    error
}

func NewDescribeCompletedQuestListsByUserIdResultFromDict(data map[string]interface{}) DescribeCompletedQuestListsByUserIdResult {
    return DescribeCompletedQuestListsByUserIdResult {
        Items: CastCompletedQuestLists(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeCompletedQuestListsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastCompletedQuestListsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeCompletedQuestListsByUserIdResult) Pointer() *DescribeCompletedQuestListsByUserIdResult {
    return &p
}

type GetCompletedQuestListResult struct {
    Item *CompletedQuestList `json:"item"`
}

type GetCompletedQuestListAsyncResult struct {
	result *GetCompletedQuestListResult
	err    error
}

func NewGetCompletedQuestListResultFromDict(data map[string]interface{}) GetCompletedQuestListResult {
    return GetCompletedQuestListResult {
        Item: NewCompletedQuestListFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCompletedQuestListResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCompletedQuestListResult) Pointer() *GetCompletedQuestListResult {
    return &p
}

type GetCompletedQuestListByUserIdResult struct {
    Item *CompletedQuestList `json:"item"`
}

type GetCompletedQuestListByUserIdAsyncResult struct {
	result *GetCompletedQuestListByUserIdResult
	err    error
}

func NewGetCompletedQuestListByUserIdResultFromDict(data map[string]interface{}) GetCompletedQuestListByUserIdResult {
    return GetCompletedQuestListByUserIdResult {
        Item: NewCompletedQuestListFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCompletedQuestListByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCompletedQuestListByUserIdResult) Pointer() *GetCompletedQuestListByUserIdResult {
    return &p
}

type DeleteCompletedQuestListByUserIdResult struct {
    Item *CompletedQuestList `json:"item"`
}

type DeleteCompletedQuestListByUserIdAsyncResult struct {
	result *DeleteCompletedQuestListByUserIdResult
	err    error
}

func NewDeleteCompletedQuestListByUserIdResultFromDict(data map[string]interface{}) DeleteCompletedQuestListByUserIdResult {
    return DeleteCompletedQuestListByUserIdResult {
        Item: NewCompletedQuestListFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteCompletedQuestListByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteCompletedQuestListByUserIdResult) Pointer() *DeleteCompletedQuestListByUserIdResult {
    return &p
}

type DescribeQuestGroupModelsResult struct {
    Items []QuestGroupModel `json:"items"`
}

type DescribeQuestGroupModelsAsyncResult struct {
	result *DescribeQuestGroupModelsResult
	err    error
}

func NewDescribeQuestGroupModelsResultFromDict(data map[string]interface{}) DescribeQuestGroupModelsResult {
    return DescribeQuestGroupModelsResult {
        Items: CastQuestGroupModels(core.CastArray(data["items"])),
    }
}

func (p DescribeQuestGroupModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastQuestGroupModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeQuestGroupModelsResult) Pointer() *DescribeQuestGroupModelsResult {
    return &p
}

type GetQuestGroupModelResult struct {
    Item *QuestGroupModel `json:"item"`
}

type GetQuestGroupModelAsyncResult struct {
	result *GetQuestGroupModelResult
	err    error
}

func NewGetQuestGroupModelResultFromDict(data map[string]interface{}) GetQuestGroupModelResult {
    return GetQuestGroupModelResult {
        Item: NewQuestGroupModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetQuestGroupModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetQuestGroupModelResult) Pointer() *GetQuestGroupModelResult {
    return &p
}

type DescribeQuestModelsResult struct {
    Items []QuestModel `json:"items"`
}

type DescribeQuestModelsAsyncResult struct {
	result *DescribeQuestModelsResult
	err    error
}

func NewDescribeQuestModelsResultFromDict(data map[string]interface{}) DescribeQuestModelsResult {
    return DescribeQuestModelsResult {
        Items: CastQuestModels(core.CastArray(data["items"])),
    }
}

func (p DescribeQuestModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastQuestModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeQuestModelsResult) Pointer() *DescribeQuestModelsResult {
    return &p
}

type GetQuestModelResult struct {
    Item *QuestModel `json:"item"`
}

type GetQuestModelAsyncResult struct {
	result *GetQuestModelResult
	err    error
}

func NewGetQuestModelResultFromDict(data map[string]interface{}) GetQuestModelResult {
    return GetQuestModelResult {
        Item: NewQuestModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetQuestModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetQuestModelResult) Pointer() *GetQuestModelResult {
    return &p
}