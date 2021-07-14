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

package exchange

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
    EnableAwaitExchange *bool `json:"enableAwaitExchange"`
    EnableDirectExchange *bool `json:"enableDirectExchange"`
    QueueNamespaceId *string `json:"queueNamespaceId"`
    KeyId *string `json:"keyId"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        EnableAwaitExchange: core.CastBool(data["enableAwaitExchange"]),
        EnableDirectExchange: core.CastBool(data["enableDirectExchange"]),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "enableAwaitExchange": p.EnableAwaitExchange,
        "enableDirectExchange": p.EnableDirectExchange,
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
    EnableAwaitExchange *bool `json:"enableAwaitExchange"`
    EnableDirectExchange *bool `json:"enableDirectExchange"`
    QueueNamespaceId *string `json:"queueNamespaceId"`
    KeyId *string `json:"keyId"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        EnableAwaitExchange: core.CastBool(data["enableAwaitExchange"]),
        EnableDirectExchange: core.CastBool(data["enableDirectExchange"]),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "enableAwaitExchange": p.EnableAwaitExchange,
        "enableDirectExchange": p.EnableDirectExchange,
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

type DescribeRateModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeRateModelsRequestFromDict(data map[string]interface{}) DescribeRateModelsRequest {
    return DescribeRateModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeRateModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeRateModelsRequest) Pointer() *DescribeRateModelsRequest {
    return &p
}

type GetRateModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RateName *string `json:"rateName"`
}

func NewGetRateModelRequestFromDict(data map[string]interface{}) GetRateModelRequest {
    return GetRateModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RateName: core.CastString(data["rateName"]),
    }
}

func (p GetRateModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "rateName": p.RateName,
    }
}

func (p GetRateModelRequest) Pointer() *GetRateModelRequest {
    return &p
}

type DescribeRateModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeRateModelMastersRequestFromDict(data map[string]interface{}) DescribeRateModelMastersRequest {
    return DescribeRateModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeRateModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeRateModelMastersRequest) Pointer() *DescribeRateModelMastersRequest {
    return &p
}

type CreateRateModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    TimingType *string `json:"timingType"`
    LockTime *int32 `json:"lockTime"`
    EnableSkip *bool `json:"enableSkip"`
    SkipConsumeActions []ConsumeAction `json:"skipConsumeActions"`
    AcquireActions []AcquireAction `json:"acquireActions"`
    ConsumeActions []ConsumeAction `json:"consumeActions"`
}

func NewCreateRateModelMasterRequestFromDict(data map[string]interface{}) CreateRateModelMasterRequest {
    return CreateRateModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        TimingType: core.CastString(data["timingType"]),
        LockTime: core.CastInt32(data["lockTime"]),
        EnableSkip: core.CastBool(data["enableSkip"]),
        SkipConsumeActions: CastConsumeActions(core.CastArray(data["skipConsumeActions"])),
        AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
        ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
    }
}

func (p CreateRateModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "timingType": p.TimingType,
        "lockTime": p.LockTime,
        "enableSkip": p.EnableSkip,
        "skipConsumeActions": CastConsumeActionsFromDict(
            p.SkipConsumeActions,
        ),
        "acquireActions": CastAcquireActionsFromDict(
            p.AcquireActions,
        ),
        "consumeActions": CastConsumeActionsFromDict(
            p.ConsumeActions,
        ),
    }
}

func (p CreateRateModelMasterRequest) Pointer() *CreateRateModelMasterRequest {
    return &p
}

type GetRateModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RateName *string `json:"rateName"`
}

func NewGetRateModelMasterRequestFromDict(data map[string]interface{}) GetRateModelMasterRequest {
    return GetRateModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RateName: core.CastString(data["rateName"]),
    }
}

func (p GetRateModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "rateName": p.RateName,
    }
}

func (p GetRateModelMasterRequest) Pointer() *GetRateModelMasterRequest {
    return &p
}

type UpdateRateModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RateName *string `json:"rateName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    TimingType *string `json:"timingType"`
    LockTime *int32 `json:"lockTime"`
    EnableSkip *bool `json:"enableSkip"`
    SkipConsumeActions []ConsumeAction `json:"skipConsumeActions"`
    AcquireActions []AcquireAction `json:"acquireActions"`
    ConsumeActions []ConsumeAction `json:"consumeActions"`
}

func NewUpdateRateModelMasterRequestFromDict(data map[string]interface{}) UpdateRateModelMasterRequest {
    return UpdateRateModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RateName: core.CastString(data["rateName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        TimingType: core.CastString(data["timingType"]),
        LockTime: core.CastInt32(data["lockTime"]),
        EnableSkip: core.CastBool(data["enableSkip"]),
        SkipConsumeActions: CastConsumeActions(core.CastArray(data["skipConsumeActions"])),
        AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
        ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
    }
}

func (p UpdateRateModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "rateName": p.RateName,
        "description": p.Description,
        "metadata": p.Metadata,
        "timingType": p.TimingType,
        "lockTime": p.LockTime,
        "enableSkip": p.EnableSkip,
        "skipConsumeActions": CastConsumeActionsFromDict(
            p.SkipConsumeActions,
        ),
        "acquireActions": CastAcquireActionsFromDict(
            p.AcquireActions,
        ),
        "consumeActions": CastConsumeActionsFromDict(
            p.ConsumeActions,
        ),
    }
}

func (p UpdateRateModelMasterRequest) Pointer() *UpdateRateModelMasterRequest {
    return &p
}

type DeleteRateModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RateName *string `json:"rateName"`
}

func NewDeleteRateModelMasterRequestFromDict(data map[string]interface{}) DeleteRateModelMasterRequest {
    return DeleteRateModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RateName: core.CastString(data["rateName"]),
    }
}

func (p DeleteRateModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "rateName": p.RateName,
    }
}

func (p DeleteRateModelMasterRequest) Pointer() *DeleteRateModelMasterRequest {
    return &p
}

type ExchangeRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RateName *string `json:"rateName"`
    AccessToken *string `json:"accessToken"`
    Count *int32 `json:"count"`
    Config []Config `json:"config"`
}

func NewExchangeRequestFromDict(data map[string]interface{}) ExchangeRequest {
    return ExchangeRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RateName: core.CastString(data["rateName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Count: core.CastInt32(data["count"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p ExchangeRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "rateName": p.RateName,
        "accessToken": p.AccessToken,
        "count": p.Count,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p ExchangeRequest) Pointer() *ExchangeRequest {
    return &p
}

type ExchangeByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RateName *string `json:"rateName"`
    UserId *string `json:"userId"`
    Count *int32 `json:"count"`
    Config []Config `json:"config"`
}

func NewExchangeByUserIdRequestFromDict(data map[string]interface{}) ExchangeByUserIdRequest {
    return ExchangeByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RateName: core.CastString(data["rateName"]),
        UserId: core.CastString(data["userId"]),
        Count: core.CastInt32(data["count"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p ExchangeByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "rateName": p.RateName,
        "userId": p.UserId,
        "count": p.Count,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p ExchangeByUserIdRequest) Pointer() *ExchangeByUserIdRequest {
    return &p
}

type ExchangeByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewExchangeByStampSheetRequestFromDict(data map[string]interface{}) ExchangeByStampSheetRequest {
    return ExchangeByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p ExchangeByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p ExchangeByStampSheetRequest) Pointer() *ExchangeByStampSheetRequest {
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

type GetCurrentRateMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentRateMasterRequestFromDict(data map[string]interface{}) GetCurrentRateMasterRequest {
    return GetCurrentRateMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentRateMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentRateMasterRequest) Pointer() *GetCurrentRateMasterRequest {
    return &p
}

type UpdateCurrentRateMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentRateMasterRequestFromDict(data map[string]interface{}) UpdateCurrentRateMasterRequest {
    return UpdateCurrentRateMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentRateMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentRateMasterRequest) Pointer() *UpdateCurrentRateMasterRequest {
    return &p
}

type UpdateCurrentRateMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentRateMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentRateMasterFromGitHubRequest {
    return UpdateCurrentRateMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentRateMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentRateMasterFromGitHubRequest) Pointer() *UpdateCurrentRateMasterFromGitHubRequest {
    return &p
}

type CreateAwaitByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    RateName *string `json:"rateName"`
    Count *int32 `json:"count"`
}

func NewCreateAwaitByUserIdRequestFromDict(data map[string]interface{}) CreateAwaitByUserIdRequest {
    return CreateAwaitByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        RateName: core.CastString(data["rateName"]),
        Count: core.CastInt32(data["count"]),
    }
}

func (p CreateAwaitByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "rateName": p.RateName,
        "count": p.Count,
    }
}

func (p CreateAwaitByUserIdRequest) Pointer() *CreateAwaitByUserIdRequest {
    return &p
}

type DescribeAwaitsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    RateName *string `json:"rateName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeAwaitsRequestFromDict(data map[string]interface{}) DescribeAwaitsRequest {
    return DescribeAwaitsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        RateName: core.CastString(data["rateName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeAwaitsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "rateName": p.RateName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeAwaitsRequest) Pointer() *DescribeAwaitsRequest {
    return &p
}

type DescribeAwaitsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    RateName *string `json:"rateName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeAwaitsByUserIdRequestFromDict(data map[string]interface{}) DescribeAwaitsByUserIdRequest {
    return DescribeAwaitsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        RateName: core.CastString(data["rateName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeAwaitsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "rateName": p.RateName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeAwaitsByUserIdRequest) Pointer() *DescribeAwaitsByUserIdRequest {
    return &p
}

type GetAwaitRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    RateName *string `json:"rateName"`
    AwaitName *string `json:"awaitName"`
}

func NewGetAwaitRequestFromDict(data map[string]interface{}) GetAwaitRequest {
    return GetAwaitRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        RateName: core.CastString(data["rateName"]),
        AwaitName: core.CastString(data["awaitName"]),
    }
}

func (p GetAwaitRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "rateName": p.RateName,
        "awaitName": p.AwaitName,
    }
}

func (p GetAwaitRequest) Pointer() *GetAwaitRequest {
    return &p
}

type GetAwaitByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    RateName *string `json:"rateName"`
    AwaitName *string `json:"awaitName"`
}

func NewGetAwaitByUserIdRequestFromDict(data map[string]interface{}) GetAwaitByUserIdRequest {
    return GetAwaitByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        RateName: core.CastString(data["rateName"]),
        AwaitName: core.CastString(data["awaitName"]),
    }
}

func (p GetAwaitByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "rateName": p.RateName,
        "awaitName": p.AwaitName,
    }
}

func (p GetAwaitByUserIdRequest) Pointer() *GetAwaitByUserIdRequest {
    return &p
}

type AcquireRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    RateName *string `json:"rateName"`
    AwaitName *string `json:"awaitName"`
    Config []Config `json:"config"`
}

func NewAcquireRequestFromDict(data map[string]interface{}) AcquireRequest {
    return AcquireRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        RateName: core.CastString(data["rateName"]),
        AwaitName: core.CastString(data["awaitName"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p AcquireRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "rateName": p.RateName,
        "awaitName": p.AwaitName,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p AcquireRequest) Pointer() *AcquireRequest {
    return &p
}

type AcquireByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    RateName *string `json:"rateName"`
    AwaitName *string `json:"awaitName"`
    Config []Config `json:"config"`
}

func NewAcquireByUserIdRequestFromDict(data map[string]interface{}) AcquireByUserIdRequest {
    return AcquireByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        RateName: core.CastString(data["rateName"]),
        AwaitName: core.CastString(data["awaitName"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p AcquireByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "rateName": p.RateName,
        "awaitName": p.AwaitName,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p AcquireByUserIdRequest) Pointer() *AcquireByUserIdRequest {
    return &p
}

type AcquireForceByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    RateName *string `json:"rateName"`
    AwaitName *string `json:"awaitName"`
    Config []Config `json:"config"`
}

func NewAcquireForceByUserIdRequestFromDict(data map[string]interface{}) AcquireForceByUserIdRequest {
    return AcquireForceByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        RateName: core.CastString(data["rateName"]),
        AwaitName: core.CastString(data["awaitName"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p AcquireForceByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "rateName": p.RateName,
        "awaitName": p.AwaitName,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p AcquireForceByUserIdRequest) Pointer() *AcquireForceByUserIdRequest {
    return &p
}

type SkipRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    RateName *string `json:"rateName"`
    AwaitName *string `json:"awaitName"`
    Config []Config `json:"config"`
}

func NewSkipRequestFromDict(data map[string]interface{}) SkipRequest {
    return SkipRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        RateName: core.CastString(data["rateName"]),
        AwaitName: core.CastString(data["awaitName"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p SkipRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "rateName": p.RateName,
        "awaitName": p.AwaitName,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p SkipRequest) Pointer() *SkipRequest {
    return &p
}

type SkipByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    RateName *string `json:"rateName"`
    AwaitName *string `json:"awaitName"`
    Config []Config `json:"config"`
}

func NewSkipByUserIdRequestFromDict(data map[string]interface{}) SkipByUserIdRequest {
    return SkipByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        RateName: core.CastString(data["rateName"]),
        AwaitName: core.CastString(data["awaitName"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p SkipByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "rateName": p.RateName,
        "awaitName": p.AwaitName,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p SkipByUserIdRequest) Pointer() *SkipByUserIdRequest {
    return &p
}

type DeleteAwaitRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    RateName *string `json:"rateName"`
    AwaitName *string `json:"awaitName"`
}

func NewDeleteAwaitRequestFromDict(data map[string]interface{}) DeleteAwaitRequest {
    return DeleteAwaitRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        RateName: core.CastString(data["rateName"]),
        AwaitName: core.CastString(data["awaitName"]),
    }
}

func (p DeleteAwaitRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "rateName": p.RateName,
        "awaitName": p.AwaitName,
    }
}

func (p DeleteAwaitRequest) Pointer() *DeleteAwaitRequest {
    return &p
}

type DeleteAwaitByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    RateName *string `json:"rateName"`
    AwaitName *string `json:"awaitName"`
}

func NewDeleteAwaitByUserIdRequestFromDict(data map[string]interface{}) DeleteAwaitByUserIdRequest {
    return DeleteAwaitByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        RateName: core.CastString(data["rateName"]),
        AwaitName: core.CastString(data["awaitName"]),
    }
}

func (p DeleteAwaitByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "rateName": p.RateName,
        "awaitName": p.AwaitName,
    }
}

func (p DeleteAwaitByUserIdRequest) Pointer() *DeleteAwaitByUserIdRequest {
    return &p
}

type CreateAwaitByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewCreateAwaitByStampSheetRequestFromDict(data map[string]interface{}) CreateAwaitByStampSheetRequest {
    return CreateAwaitByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p CreateAwaitByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p CreateAwaitByStampSheetRequest) Pointer() *CreateAwaitByStampSheetRequest {
    return &p
}

type DeleteAwaitByStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewDeleteAwaitByStampTaskRequestFromDict(data map[string]interface{}) DeleteAwaitByStampTaskRequest {
    return DeleteAwaitByStampTaskRequest {
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p DeleteAwaitByStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p DeleteAwaitByStampTaskRequest) Pointer() *DeleteAwaitByStampTaskRequest {
    return &p
}