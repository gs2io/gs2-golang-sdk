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

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
    return DescribeNamespacesRequest {
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
    return &p
}

type CreateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    StartQuestScript *ScriptSetting `json:"startQuestScript"`
    CompleteQuestScript *ScriptSetting `json:"completeQuestScript"`
    FailedQuestScript *ScriptSetting `json:"failedQuestScript"`
    QueueNamespaceId *string `json:"queueNamespaceId"`
    KeyId *string `json:"keyId"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        StartQuestScript: NewScriptSettingFromDict(core.CastMap(data["startQuestScript"])).Pointer(),
        CompleteQuestScript: NewScriptSettingFromDict(core.CastMap(data["completeQuestScript"])).Pointer(),
        FailedQuestScript: NewScriptSettingFromDict(core.CastMap(data["failedQuestScript"])).Pointer(),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "startQuestScript": p.StartQuestScript.ToDict(),
        "completeQuestScript": p.CompleteQuestScript.ToDict(),
        "failedQuestScript": p.FailedQuestScript.ToDict(),
        "queueNamespaceId": p.QueueNamespaceId,
        "keyId": p.KeyId,
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
    return &p
}

type GetNamespaceStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
    return GetNamespaceStatusRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
    return &p
}

type GetNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
    return GetNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
    return &p
}

type UpdateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Description *string `json:"description"`
    StartQuestScript *ScriptSetting `json:"startQuestScript"`
    CompleteQuestScript *ScriptSetting `json:"completeQuestScript"`
    FailedQuestScript *ScriptSetting `json:"failedQuestScript"`
    QueueNamespaceId *string `json:"queueNamespaceId"`
    KeyId *string `json:"keyId"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        StartQuestScript: NewScriptSettingFromDict(core.CastMap(data["startQuestScript"])).Pointer(),
        CompleteQuestScript: NewScriptSettingFromDict(core.CastMap(data["completeQuestScript"])).Pointer(),
        FailedQuestScript: NewScriptSettingFromDict(core.CastMap(data["failedQuestScript"])).Pointer(),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "startQuestScript": p.StartQuestScript.ToDict(),
        "completeQuestScript": p.CompleteQuestScript.ToDict(),
        "failedQuestScript": p.FailedQuestScript.ToDict(),
        "queueNamespaceId": p.QueueNamespaceId,
        "keyId": p.KeyId,
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
    return &p
}

type DeleteNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
    return DeleteNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
    return &p
}

type DescribeQuestGroupModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeQuestGroupModelMastersRequestFromDict(data map[string]interface{}) DescribeQuestGroupModelMastersRequest {
    return DescribeQuestGroupModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeQuestGroupModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeQuestGroupModelMastersRequest) Pointer() *DescribeQuestGroupModelMastersRequest {
    return &p
}

type CreateQuestGroupModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    ChallengePeriodEventId *string `json:"challengePeriodEventId"`
}

func NewCreateQuestGroupModelMasterRequestFromDict(data map[string]interface{}) CreateQuestGroupModelMasterRequest {
    return CreateQuestGroupModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
    }
}

func (p CreateQuestGroupModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "challengePeriodEventId": p.ChallengePeriodEventId,
    }
}

func (p CreateQuestGroupModelMasterRequest) Pointer() *CreateQuestGroupModelMasterRequest {
    return &p
}

type GetQuestGroupModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
}

func NewGetQuestGroupModelMasterRequestFromDict(data map[string]interface{}) GetQuestGroupModelMasterRequest {
    return GetQuestGroupModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
    }
}

func (p GetQuestGroupModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
    }
}

func (p GetQuestGroupModelMasterRequest) Pointer() *GetQuestGroupModelMasterRequest {
    return &p
}

type UpdateQuestGroupModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    ChallengePeriodEventId *string `json:"challengePeriodEventId"`
}

func NewUpdateQuestGroupModelMasterRequestFromDict(data map[string]interface{}) UpdateQuestGroupModelMasterRequest {
    return UpdateQuestGroupModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
    }
}

func (p UpdateQuestGroupModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "description": p.Description,
        "metadata": p.Metadata,
        "challengePeriodEventId": p.ChallengePeriodEventId,
    }
}

func (p UpdateQuestGroupModelMasterRequest) Pointer() *UpdateQuestGroupModelMasterRequest {
    return &p
}

type DeleteQuestGroupModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
}

func NewDeleteQuestGroupModelMasterRequestFromDict(data map[string]interface{}) DeleteQuestGroupModelMasterRequest {
    return DeleteQuestGroupModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
    }
}

func (p DeleteQuestGroupModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
    }
}

func (p DeleteQuestGroupModelMasterRequest) Pointer() *DeleteQuestGroupModelMasterRequest {
    return &p
}

type DescribeQuestModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeQuestModelMastersRequestFromDict(data map[string]interface{}) DescribeQuestModelMastersRequest {
    return DescribeQuestModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeQuestModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeQuestModelMastersRequest) Pointer() *DescribeQuestModelMastersRequest {
    return &p
}

type CreateQuestModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    Contents []Contents `json:"contents"`
    ChallengePeriodEventId *string `json:"challengePeriodEventId"`
    ConsumeActions []ConsumeAction `json:"consumeActions"`
    FailedAcquireActions []AcquireAction `json:"failedAcquireActions"`
    PremiseQuestNames []string `json:"premiseQuestNames"`
}

func NewCreateQuestModelMasterRequestFromDict(data map[string]interface{}) CreateQuestModelMasterRequest {
    return CreateQuestModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Contents: CastContentses(core.CastArray(data["contents"])),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
        ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
        FailedAcquireActions: CastAcquireActions(core.CastArray(data["failedAcquireActions"])),
        PremiseQuestNames: core.CastStrings(core.CastArray(data["premiseQuestNames"])),
    }
}

func (p CreateQuestModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "contents": CastContentsesFromDict(
            p.Contents,
        ),
        "challengePeriodEventId": p.ChallengePeriodEventId,
        "consumeActions": CastConsumeActionsFromDict(
            p.ConsumeActions,
        ),
        "failedAcquireActions": CastAcquireActionsFromDict(
            p.FailedAcquireActions,
        ),
        "premiseQuestNames": core.CastStringsFromDict(
            p.PremiseQuestNames,
        ),
    }
}

func (p CreateQuestModelMasterRequest) Pointer() *CreateQuestModelMasterRequest {
    return &p
}

type GetQuestModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    QuestName *string `json:"questName"`
}

func NewGetQuestModelMasterRequestFromDict(data map[string]interface{}) GetQuestModelMasterRequest {
    return GetQuestModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        QuestName: core.CastString(data["questName"]),
    }
}

func (p GetQuestModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "questName": p.QuestName,
    }
}

func (p GetQuestModelMasterRequest) Pointer() *GetQuestModelMasterRequest {
    return &p
}

type UpdateQuestModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    QuestName *string `json:"questName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    Contents []Contents `json:"contents"`
    ChallengePeriodEventId *string `json:"challengePeriodEventId"`
    ConsumeActions []ConsumeAction `json:"consumeActions"`
    FailedAcquireActions []AcquireAction `json:"failedAcquireActions"`
    PremiseQuestNames []string `json:"premiseQuestNames"`
}

func NewUpdateQuestModelMasterRequestFromDict(data map[string]interface{}) UpdateQuestModelMasterRequest {
    return UpdateQuestModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        QuestName: core.CastString(data["questName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Contents: CastContentses(core.CastArray(data["contents"])),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
        ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
        FailedAcquireActions: CastAcquireActions(core.CastArray(data["failedAcquireActions"])),
        PremiseQuestNames: core.CastStrings(core.CastArray(data["premiseQuestNames"])),
    }
}

func (p UpdateQuestModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "questName": p.QuestName,
        "description": p.Description,
        "metadata": p.Metadata,
        "contents": CastContentsesFromDict(
            p.Contents,
        ),
        "challengePeriodEventId": p.ChallengePeriodEventId,
        "consumeActions": CastConsumeActionsFromDict(
            p.ConsumeActions,
        ),
        "failedAcquireActions": CastAcquireActionsFromDict(
            p.FailedAcquireActions,
        ),
        "premiseQuestNames": core.CastStringsFromDict(
            p.PremiseQuestNames,
        ),
    }
}

func (p UpdateQuestModelMasterRequest) Pointer() *UpdateQuestModelMasterRequest {
    return &p
}

type DeleteQuestModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    QuestName *string `json:"questName"`
}

func NewDeleteQuestModelMasterRequestFromDict(data map[string]interface{}) DeleteQuestModelMasterRequest {
    return DeleteQuestModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        QuestName: core.CastString(data["questName"]),
    }
}

func (p DeleteQuestModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "questName": p.QuestName,
    }
}

func (p DeleteQuestModelMasterRequest) Pointer() *DeleteQuestModelMasterRequest {
    return &p
}

type ExportMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewExportMasterRequestFromDict(data map[string]interface{}) ExportMasterRequest {
    return ExportMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p ExportMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p ExportMasterRequest) Pointer() *ExportMasterRequest {
    return &p
}

type GetCurrentQuestMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentQuestMasterRequestFromDict(data map[string]interface{}) GetCurrentQuestMasterRequest {
    return GetCurrentQuestMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentQuestMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentQuestMasterRequest) Pointer() *GetCurrentQuestMasterRequest {
    return &p
}

type UpdateCurrentQuestMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentQuestMasterRequestFromDict(data map[string]interface{}) UpdateCurrentQuestMasterRequest {
    return UpdateCurrentQuestMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentQuestMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentQuestMasterRequest) Pointer() *UpdateCurrentQuestMasterRequest {
    return &p
}

type UpdateCurrentQuestMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentQuestMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentQuestMasterFromGitHubRequest {
    return UpdateCurrentQuestMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentQuestMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentQuestMasterFromGitHubRequest) Pointer() *UpdateCurrentQuestMasterFromGitHubRequest {
    return &p
}

type DescribeProgressesByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeProgressesByUserIdRequestFromDict(data map[string]interface{}) DescribeProgressesByUserIdRequest {
    return DescribeProgressesByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeProgressesByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeProgressesByUserIdRequest) Pointer() *DescribeProgressesByUserIdRequest {
    return &p
}

type CreateProgressByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    QuestModelId *string `json:"questModelId"`
    Force *bool `json:"force"`
    Config []Config `json:"config"`
}

func NewCreateProgressByUserIdRequestFromDict(data map[string]interface{}) CreateProgressByUserIdRequest {
    return CreateProgressByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        QuestModelId: core.CastString(data["questModelId"]),
        Force: core.CastBool(data["force"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p CreateProgressByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "questModelId": p.QuestModelId,
        "force": p.Force,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p CreateProgressByUserIdRequest) Pointer() *CreateProgressByUserIdRequest {
    return &p
}

type GetProgressRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
}

func NewGetProgressRequestFromDict(data map[string]interface{}) GetProgressRequest {
    return GetProgressRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p GetProgressRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
    }
}

func (p GetProgressRequest) Pointer() *GetProgressRequest {
    return &p
}

type GetProgressByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
}

func NewGetProgressByUserIdRequestFromDict(data map[string]interface{}) GetProgressByUserIdRequest {
    return GetProgressByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p GetProgressByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
    }
}

func (p GetProgressByUserIdRequest) Pointer() *GetProgressByUserIdRequest {
    return &p
}

type StartRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    QuestName *string `json:"questName"`
    AccessToken *string `json:"accessToken"`
    Force *bool `json:"force"`
    Config []Config `json:"config"`
}

func NewStartRequestFromDict(data map[string]interface{}) StartRequest {
    return StartRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        QuestName: core.CastString(data["questName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Force: core.CastBool(data["force"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p StartRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "questName": p.QuestName,
        "accessToken": p.AccessToken,
        "force": p.Force,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p StartRequest) Pointer() *StartRequest {
    return &p
}

type StartByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    QuestName *string `json:"questName"`
    UserId *string `json:"userId"`
    Force *bool `json:"force"`
    Config []Config `json:"config"`
}

func NewStartByUserIdRequestFromDict(data map[string]interface{}) StartByUserIdRequest {
    return StartByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        QuestName: core.CastString(data["questName"]),
        UserId: core.CastString(data["userId"]),
        Force: core.CastBool(data["force"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p StartByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "questName": p.QuestName,
        "userId": p.UserId,
        "force": p.Force,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p StartByUserIdRequest) Pointer() *StartByUserIdRequest {
    return &p
}

type EndRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TransactionId *string `json:"transactionId"`
    Rewards []Reward `json:"rewards"`
    IsComplete *bool `json:"isComplete"`
    Config []Config `json:"config"`
}

func NewEndRequestFromDict(data map[string]interface{}) EndRequest {
    return EndRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TransactionId: core.CastString(data["transactionId"]),
        Rewards: CastRewards(core.CastArray(data["rewards"])),
        IsComplete: core.CastBool(data["isComplete"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p EndRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "transactionId": p.TransactionId,
        "rewards": CastRewardsFromDict(
            p.Rewards,
        ),
        "isComplete": p.IsComplete,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p EndRequest) Pointer() *EndRequest {
    return &p
}

type EndByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TransactionId *string `json:"transactionId"`
    Rewards []Reward `json:"rewards"`
    IsComplete *bool `json:"isComplete"`
    Config []Config `json:"config"`
}

func NewEndByUserIdRequestFromDict(data map[string]interface{}) EndByUserIdRequest {
    return EndByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TransactionId: core.CastString(data["transactionId"]),
        Rewards: CastRewards(core.CastArray(data["rewards"])),
        IsComplete: core.CastBool(data["isComplete"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p EndByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "transactionId": p.TransactionId,
        "rewards": CastRewardsFromDict(
            p.Rewards,
        ),
        "isComplete": p.IsComplete,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p EndByUserIdRequest) Pointer() *EndByUserIdRequest {
    return &p
}

type DeleteProgressRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
}

func NewDeleteProgressRequestFromDict(data map[string]interface{}) DeleteProgressRequest {
    return DeleteProgressRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p DeleteProgressRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
    }
}

func (p DeleteProgressRequest) Pointer() *DeleteProgressRequest {
    return &p
}

type DeleteProgressByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
}

func NewDeleteProgressByUserIdRequestFromDict(data map[string]interface{}) DeleteProgressByUserIdRequest {
    return DeleteProgressByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p DeleteProgressByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
    }
}

func (p DeleteProgressByUserIdRequest) Pointer() *DeleteProgressByUserIdRequest {
    return &p
}

type CreateProgressByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewCreateProgressByStampSheetRequestFromDict(data map[string]interface{}) CreateProgressByStampSheetRequest {
    return CreateProgressByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p CreateProgressByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p CreateProgressByStampSheetRequest) Pointer() *CreateProgressByStampSheetRequest {
    return &p
}

type DeleteProgressByStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewDeleteProgressByStampTaskRequestFromDict(data map[string]interface{}) DeleteProgressByStampTaskRequest {
    return DeleteProgressByStampTaskRequest {
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p DeleteProgressByStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p DeleteProgressByStampTaskRequest) Pointer() *DeleteProgressByStampTaskRequest {
    return &p
}

type DescribeCompletedQuestListsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeCompletedQuestListsRequestFromDict(data map[string]interface{}) DescribeCompletedQuestListsRequest {
    return DescribeCompletedQuestListsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeCompletedQuestListsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeCompletedQuestListsRequest) Pointer() *DescribeCompletedQuestListsRequest {
    return &p
}

type DescribeCompletedQuestListsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeCompletedQuestListsByUserIdRequestFromDict(data map[string]interface{}) DescribeCompletedQuestListsByUserIdRequest {
    return DescribeCompletedQuestListsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeCompletedQuestListsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeCompletedQuestListsByUserIdRequest) Pointer() *DescribeCompletedQuestListsByUserIdRequest {
    return &p
}

type GetCompletedQuestListRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    AccessToken *string `json:"accessToken"`
}

func NewGetCompletedQuestListRequestFromDict(data map[string]interface{}) GetCompletedQuestListRequest {
    return GetCompletedQuestListRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p GetCompletedQuestListRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "accessToken": p.AccessToken,
    }
}

func (p GetCompletedQuestListRequest) Pointer() *GetCompletedQuestListRequest {
    return &p
}

type GetCompletedQuestListByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    UserId *string `json:"userId"`
}

func NewGetCompletedQuestListByUserIdRequestFromDict(data map[string]interface{}) GetCompletedQuestListByUserIdRequest {
    return GetCompletedQuestListByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p GetCompletedQuestListByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "userId": p.UserId,
    }
}

func (p GetCompletedQuestListByUserIdRequest) Pointer() *GetCompletedQuestListByUserIdRequest {
    return &p
}

type DeleteCompletedQuestListByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    UserId *string `json:"userId"`
}

func NewDeleteCompletedQuestListByUserIdRequestFromDict(data map[string]interface{}) DeleteCompletedQuestListByUserIdRequest {
    return DeleteCompletedQuestListByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p DeleteCompletedQuestListByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "userId": p.UserId,
    }
}

func (p DeleteCompletedQuestListByUserIdRequest) Pointer() *DeleteCompletedQuestListByUserIdRequest {
    return &p
}

type DescribeQuestGroupModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeQuestGroupModelsRequestFromDict(data map[string]interface{}) DescribeQuestGroupModelsRequest {
    return DescribeQuestGroupModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeQuestGroupModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeQuestGroupModelsRequest) Pointer() *DescribeQuestGroupModelsRequest {
    return &p
}

type GetQuestGroupModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
}

func NewGetQuestGroupModelRequestFromDict(data map[string]interface{}) GetQuestGroupModelRequest {
    return GetQuestGroupModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
    }
}

func (p GetQuestGroupModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
    }
}

func (p GetQuestGroupModelRequest) Pointer() *GetQuestGroupModelRequest {
    return &p
}

type DescribeQuestModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
}

func NewDescribeQuestModelsRequestFromDict(data map[string]interface{}) DescribeQuestModelsRequest {
    return DescribeQuestModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
    }
}

func (p DescribeQuestModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
    }
}

func (p DescribeQuestModelsRequest) Pointer() *DescribeQuestModelsRequest {
    return &p
}

type GetQuestModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    QuestName *string `json:"questName"`
}

func NewGetQuestModelRequestFromDict(data map[string]interface{}) GetQuestModelRequest {
    return GetQuestModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        QuestName: core.CastString(data["questName"]),
    }
}

func (p GetQuestModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "questName": p.QuestName,
    }
}

func (p GetQuestModelRequest) Pointer() *GetQuestModelRequest {
    return &p
}