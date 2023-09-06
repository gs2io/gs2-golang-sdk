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

package loginReward

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeNamespacesRequestFromDict(dict)
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
    TransactionSetting *TransactionSetting `json:"transactionSetting"`
    ReceiveScript *ScriptSetting `json:"receiveScript"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
        ReceiveScript: NewScriptSettingFromDict(core.CastMap(data["receiveScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "transactionSetting": p.TransactionSetting.ToDict(),
        "receiveScript": p.ReceiveScript.ToDict(),
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

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceStatusRequestFromDict(dict)
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

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceRequestFromDict(dict)
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
    TransactionSetting *TransactionSetting `json:"transactionSetting"`
    ReceiveScript *ScriptSetting `json:"receiveScript"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
        ReceiveScript: NewScriptSettingFromDict(core.CastMap(data["receiveScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "transactionSetting": p.TransactionSetting.ToDict(),
        "receiveScript": p.ReceiveScript.ToDict(),
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

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteNamespaceRequestFromDict(dict)
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

type DescribeBonusModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeBonusModelMastersRequestFromJson(data string) DescribeBonusModelMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBonusModelMastersRequestFromDict(dict)
}

func NewDescribeBonusModelMastersRequestFromDict(data map[string]interface{}) DescribeBonusModelMastersRequest {
    return DescribeBonusModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeBonusModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeBonusModelMastersRequest) Pointer() *DescribeBonusModelMastersRequest {
    return &p
}

type CreateBonusModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    Mode *string `json:"mode"`
    PeriodEventId *string `json:"periodEventId"`
    ResetHour *int32 `json:"resetHour"`
    Repeat *string `json:"repeat"`
    Rewards []Reward `json:"rewards"`
    MissedReceiveRelief *string `json:"missedReceiveRelief"`
    MissedReceiveReliefConsumeActions []ConsumeAction `json:"missedReceiveReliefConsumeActions"`
}

func NewCreateBonusModelMasterRequestFromJson(data string) CreateBonusModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateBonusModelMasterRequestFromDict(dict)
}

func NewCreateBonusModelMasterRequestFromDict(data map[string]interface{}) CreateBonusModelMasterRequest {
    return CreateBonusModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Mode: core.CastString(data["mode"]),
        PeriodEventId: core.CastString(data["periodEventId"]),
        ResetHour: core.CastInt32(data["resetHour"]),
        Repeat: core.CastString(data["repeat"]),
        Rewards: CastRewards(core.CastArray(data["rewards"])),
        MissedReceiveRelief: core.CastString(data["missedReceiveRelief"]),
        MissedReceiveReliefConsumeActions: CastConsumeActions(core.CastArray(data["missedReceiveReliefConsumeActions"])),
    }
}

func (p CreateBonusModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "mode": p.Mode,
        "periodEventId": p.PeriodEventId,
        "resetHour": p.ResetHour,
        "repeat": p.Repeat,
        "rewards": CastRewardsFromDict(
            p.Rewards,
        ),
        "missedReceiveRelief": p.MissedReceiveRelief,
        "missedReceiveReliefConsumeActions": CastConsumeActionsFromDict(
            p.MissedReceiveReliefConsumeActions,
        ),
    }
}

func (p CreateBonusModelMasterRequest) Pointer() *CreateBonusModelMasterRequest {
    return &p
}

type GetBonusModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
}

func NewGetBonusModelMasterRequestFromJson(data string) GetBonusModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBonusModelMasterRequestFromDict(dict)
}

func NewGetBonusModelMasterRequestFromDict(data map[string]interface{}) GetBonusModelMasterRequest {
    return GetBonusModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
    }
}

func (p GetBonusModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
    }
}

func (p GetBonusModelMasterRequest) Pointer() *GetBonusModelMasterRequest {
    return &p
}

type UpdateBonusModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    Mode *string `json:"mode"`
    PeriodEventId *string `json:"periodEventId"`
    ResetHour *int32 `json:"resetHour"`
    Repeat *string `json:"repeat"`
    Rewards []Reward `json:"rewards"`
    MissedReceiveRelief *string `json:"missedReceiveRelief"`
    MissedReceiveReliefConsumeActions []ConsumeAction `json:"missedReceiveReliefConsumeActions"`
}

func NewUpdateBonusModelMasterRequestFromJson(data string) UpdateBonusModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateBonusModelMasterRequestFromDict(dict)
}

func NewUpdateBonusModelMasterRequestFromDict(data map[string]interface{}) UpdateBonusModelMasterRequest {
    return UpdateBonusModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Mode: core.CastString(data["mode"]),
        PeriodEventId: core.CastString(data["periodEventId"]),
        ResetHour: core.CastInt32(data["resetHour"]),
        Repeat: core.CastString(data["repeat"]),
        Rewards: CastRewards(core.CastArray(data["rewards"])),
        MissedReceiveRelief: core.CastString(data["missedReceiveRelief"]),
        MissedReceiveReliefConsumeActions: CastConsumeActions(core.CastArray(data["missedReceiveReliefConsumeActions"])),
    }
}

func (p UpdateBonusModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "description": p.Description,
        "metadata": p.Metadata,
        "mode": p.Mode,
        "periodEventId": p.PeriodEventId,
        "resetHour": p.ResetHour,
        "repeat": p.Repeat,
        "rewards": CastRewardsFromDict(
            p.Rewards,
        ),
        "missedReceiveRelief": p.MissedReceiveRelief,
        "missedReceiveReliefConsumeActions": CastConsumeActionsFromDict(
            p.MissedReceiveReliefConsumeActions,
        ),
    }
}

func (p UpdateBonusModelMasterRequest) Pointer() *UpdateBonusModelMasterRequest {
    return &p
}

type DeleteBonusModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
}

func NewDeleteBonusModelMasterRequestFromJson(data string) DeleteBonusModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteBonusModelMasterRequestFromDict(dict)
}

func NewDeleteBonusModelMasterRequestFromDict(data map[string]interface{}) DeleteBonusModelMasterRequest {
    return DeleteBonusModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
    }
}

func (p DeleteBonusModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
    }
}

func (p DeleteBonusModelMasterRequest) Pointer() *DeleteBonusModelMasterRequest {
    return &p
}

type ExportMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewExportMasterRequestFromJson(data string) ExportMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExportMasterRequestFromDict(dict)
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

type GetCurrentBonusMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentBonusMasterRequestFromJson(data string) GetCurrentBonusMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentBonusMasterRequestFromDict(dict)
}

func NewGetCurrentBonusMasterRequestFromDict(data map[string]interface{}) GetCurrentBonusMasterRequest {
    return GetCurrentBonusMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentBonusMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentBonusMasterRequest) Pointer() *GetCurrentBonusMasterRequest {
    return &p
}

type UpdateCurrentBonusMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentBonusMasterRequestFromJson(data string) UpdateCurrentBonusMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentBonusMasterRequestFromDict(dict)
}

func NewUpdateCurrentBonusMasterRequestFromDict(data map[string]interface{}) UpdateCurrentBonusMasterRequest {
    return UpdateCurrentBonusMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentBonusMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentBonusMasterRequest) Pointer() *UpdateCurrentBonusMasterRequest {
    return &p
}

type UpdateCurrentBonusMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentBonusMasterFromGitHubRequestFromJson(data string) UpdateCurrentBonusMasterFromGitHubRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentBonusMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentBonusMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentBonusMasterFromGitHubRequest {
    return UpdateCurrentBonusMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentBonusMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentBonusMasterFromGitHubRequest) Pointer() *UpdateCurrentBonusMasterFromGitHubRequest {
    return &p
}

type DescribeBonusModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeBonusModelsRequestFromJson(data string) DescribeBonusModelsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBonusModelsRequestFromDict(dict)
}

func NewDescribeBonusModelsRequestFromDict(data map[string]interface{}) DescribeBonusModelsRequest {
    return DescribeBonusModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeBonusModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeBonusModelsRequest) Pointer() *DescribeBonusModelsRequest {
    return &p
}

type GetBonusModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
}

func NewGetBonusModelRequestFromJson(data string) GetBonusModelRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBonusModelRequestFromDict(dict)
}

func NewGetBonusModelRequestFromDict(data map[string]interface{}) GetBonusModelRequest {
    return GetBonusModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
    }
}

func (p GetBonusModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
    }
}

func (p GetBonusModelRequest) Pointer() *GetBonusModelRequest {
    return &p
}

type ReceiveRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    AccessToken *string `json:"accessToken"`
    Config []Config `json:"config"`
}

func NewReceiveRequestFromJson(data string) ReceiveRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewReceiveRequestFromDict(dict)
}

func NewReceiveRequestFromDict(data map[string]interface{}) ReceiveRequest {
    return ReceiveRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p ReceiveRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "accessToken": p.AccessToken,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p ReceiveRequest) Pointer() *ReceiveRequest {
    return &p
}

type ReceiveByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    UserId *string `json:"userId"`
    Config []Config `json:"config"`
}

func NewReceiveByUserIdRequestFromJson(data string) ReceiveByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewReceiveByUserIdRequestFromDict(dict)
}

func NewReceiveByUserIdRequestFromDict(data map[string]interface{}) ReceiveByUserIdRequest {
    return ReceiveByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        UserId: core.CastString(data["userId"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p ReceiveByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "userId": p.UserId,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p ReceiveByUserIdRequest) Pointer() *ReceiveByUserIdRequest {
    return &p
}

type MissedReceiveRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    AccessToken *string `json:"accessToken"`
    StepNumber *int32 `json:"stepNumber"`
    Config []Config `json:"config"`
}

func NewMissedReceiveRequestFromJson(data string) MissedReceiveRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMissedReceiveRequestFromDict(dict)
}

func NewMissedReceiveRequestFromDict(data map[string]interface{}) MissedReceiveRequest {
    return MissedReceiveRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        AccessToken: core.CastString(data["accessToken"]),
        StepNumber: core.CastInt32(data["stepNumber"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p MissedReceiveRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "accessToken": p.AccessToken,
        "stepNumber": p.StepNumber,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p MissedReceiveRequest) Pointer() *MissedReceiveRequest {
    return &p
}

type MissedReceiveByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    UserId *string `json:"userId"`
    StepNumber *int32 `json:"stepNumber"`
    Config []Config `json:"config"`
}

func NewMissedReceiveByUserIdRequestFromJson(data string) MissedReceiveByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMissedReceiveByUserIdRequestFromDict(dict)
}

func NewMissedReceiveByUserIdRequestFromDict(data map[string]interface{}) MissedReceiveByUserIdRequest {
    return MissedReceiveByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        UserId: core.CastString(data["userId"]),
        StepNumber: core.CastInt32(data["stepNumber"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p MissedReceiveByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "userId": p.UserId,
        "stepNumber": p.StepNumber,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p MissedReceiveByUserIdRequest) Pointer() *MissedReceiveByUserIdRequest {
    return &p
}

type DescribeReceiveStatusesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeReceiveStatusesRequestFromJson(data string) DescribeReceiveStatusesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeReceiveStatusesRequestFromDict(dict)
}

func NewDescribeReceiveStatusesRequestFromDict(data map[string]interface{}) DescribeReceiveStatusesRequest {
    return DescribeReceiveStatusesRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeReceiveStatusesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeReceiveStatusesRequest) Pointer() *DescribeReceiveStatusesRequest {
    return &p
}

type DescribeReceiveStatusesByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeReceiveStatusesByUserIdRequestFromJson(data string) DescribeReceiveStatusesByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeReceiveStatusesByUserIdRequestFromDict(dict)
}

func NewDescribeReceiveStatusesByUserIdRequestFromDict(data map[string]interface{}) DescribeReceiveStatusesByUserIdRequest {
    return DescribeReceiveStatusesByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeReceiveStatusesByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeReceiveStatusesByUserIdRequest) Pointer() *DescribeReceiveStatusesByUserIdRequest {
    return &p
}

type GetReceiveStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    AccessToken *string `json:"accessToken"`
}

func NewGetReceiveStatusRequestFromJson(data string) GetReceiveStatusRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetReceiveStatusRequestFromDict(dict)
}

func NewGetReceiveStatusRequestFromDict(data map[string]interface{}) GetReceiveStatusRequest {
    return GetReceiveStatusRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p GetReceiveStatusRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "accessToken": p.AccessToken,
    }
}

func (p GetReceiveStatusRequest) Pointer() *GetReceiveStatusRequest {
    return &p
}

type GetReceiveStatusByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    UserId *string `json:"userId"`
}

func NewGetReceiveStatusByUserIdRequestFromJson(data string) GetReceiveStatusByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetReceiveStatusByUserIdRequestFromDict(dict)
}

func NewGetReceiveStatusByUserIdRequestFromDict(data map[string]interface{}) GetReceiveStatusByUserIdRequest {
    return GetReceiveStatusByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p GetReceiveStatusByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "userId": p.UserId,
    }
}

func (p GetReceiveStatusByUserIdRequest) Pointer() *GetReceiveStatusByUserIdRequest {
    return &p
}

type DeleteReceiveStatusByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    UserId *string `json:"userId"`
}

func NewDeleteReceiveStatusByUserIdRequestFromJson(data string) DeleteReceiveStatusByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteReceiveStatusByUserIdRequestFromDict(dict)
}

func NewDeleteReceiveStatusByUserIdRequestFromDict(data map[string]interface{}) DeleteReceiveStatusByUserIdRequest {
    return DeleteReceiveStatusByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p DeleteReceiveStatusByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "userId": p.UserId,
    }
}

func (p DeleteReceiveStatusByUserIdRequest) Pointer() *DeleteReceiveStatusByUserIdRequest {
    return &p
}

type DeleteReceiveStatusByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewDeleteReceiveStatusByStampSheetRequestFromJson(data string) DeleteReceiveStatusByStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteReceiveStatusByStampSheetRequestFromDict(dict)
}

func NewDeleteReceiveStatusByStampSheetRequestFromDict(data map[string]interface{}) DeleteReceiveStatusByStampSheetRequest {
    return DeleteReceiveStatusByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p DeleteReceiveStatusByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p DeleteReceiveStatusByStampSheetRequest) Pointer() *DeleteReceiveStatusByStampSheetRequest {
    return &p
}

type MarkReceivedRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    AccessToken *string `json:"accessToken"`
    StepNumber *int32 `json:"stepNumber"`
}

func NewMarkReceivedRequestFromJson(data string) MarkReceivedRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMarkReceivedRequestFromDict(dict)
}

func NewMarkReceivedRequestFromDict(data map[string]interface{}) MarkReceivedRequest {
    return MarkReceivedRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        AccessToken: core.CastString(data["accessToken"]),
        StepNumber: core.CastInt32(data["stepNumber"]),
    }
}

func (p MarkReceivedRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "accessToken": p.AccessToken,
        "stepNumber": p.StepNumber,
    }
}

func (p MarkReceivedRequest) Pointer() *MarkReceivedRequest {
    return &p
}

type MarkReceivedByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    UserId *string `json:"userId"`
    StepNumber *int32 `json:"stepNumber"`
}

func NewMarkReceivedByUserIdRequestFromJson(data string) MarkReceivedByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMarkReceivedByUserIdRequestFromDict(dict)
}

func NewMarkReceivedByUserIdRequestFromDict(data map[string]interface{}) MarkReceivedByUserIdRequest {
    return MarkReceivedByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        UserId: core.CastString(data["userId"]),
        StepNumber: core.CastInt32(data["stepNumber"]),
    }
}

func (p MarkReceivedByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "userId": p.UserId,
        "stepNumber": p.StepNumber,
    }
}

func (p MarkReceivedByUserIdRequest) Pointer() *MarkReceivedByUserIdRequest {
    return &p
}

type UnmarkReceivedByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    BonusModelName *string `json:"bonusModelName"`
    UserId *string `json:"userId"`
    StepNumber *int32 `json:"stepNumber"`
}

func NewUnmarkReceivedByUserIdRequestFromJson(data string) UnmarkReceivedByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUnmarkReceivedByUserIdRequestFromDict(dict)
}

func NewUnmarkReceivedByUserIdRequestFromDict(data map[string]interface{}) UnmarkReceivedByUserIdRequest {
    return UnmarkReceivedByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BonusModelName: core.CastString(data["bonusModelName"]),
        UserId: core.CastString(data["userId"]),
        StepNumber: core.CastInt32(data["stepNumber"]),
    }
}

func (p UnmarkReceivedByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "bonusModelName": p.BonusModelName,
        "userId": p.UserId,
        "stepNumber": p.StepNumber,
    }
}

func (p UnmarkReceivedByUserIdRequest) Pointer() *UnmarkReceivedByUserIdRequest {
    return &p
}

type MarkReceivedByStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewMarkReceivedByStampTaskRequestFromJson(data string) MarkReceivedByStampTaskRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMarkReceivedByStampTaskRequestFromDict(dict)
}

func NewMarkReceivedByStampTaskRequestFromDict(data map[string]interface{}) MarkReceivedByStampTaskRequest {
    return MarkReceivedByStampTaskRequest {
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p MarkReceivedByStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p MarkReceivedByStampTaskRequest) Pointer() *MarkReceivedByStampTaskRequest {
    return &p
}

type UnmarkReceivedByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewUnmarkReceivedByStampSheetRequestFromJson(data string) UnmarkReceivedByStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUnmarkReceivedByStampSheetRequestFromDict(dict)
}

func NewUnmarkReceivedByStampSheetRequestFromDict(data map[string]interface{}) UnmarkReceivedByStampSheetRequest {
    return UnmarkReceivedByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p UnmarkReceivedByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p UnmarkReceivedByStampSheetRequest) Pointer() *UnmarkReceivedByStampSheetRequest {
    return &p
}