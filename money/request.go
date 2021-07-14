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

package money

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
    Priority *string `json:"priority"`
    ShareFree *bool `json:"shareFree"`
    Currency *string `json:"currency"`
    AppleKey *string `json:"appleKey"`
    GoogleKey *string `json:"googleKey"`
    EnableFakeReceipt *bool `json:"enableFakeReceipt"`
    CreateWalletScript *ScriptSetting `json:"createWalletScript"`
    DepositScript *ScriptSetting `json:"depositScript"`
    WithdrawScript *ScriptSetting `json:"withdrawScript"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Priority: core.CastString(data["priority"]),
        ShareFree: core.CastBool(data["shareFree"]),
        Currency: core.CastString(data["currency"]),
        AppleKey: core.CastString(data["appleKey"]),
        GoogleKey: core.CastString(data["googleKey"]),
        EnableFakeReceipt: core.CastBool(data["enableFakeReceipt"]),
        CreateWalletScript: NewScriptSettingFromDict(core.CastMap(data["createWalletScript"])).Pointer(),
        DepositScript: NewScriptSettingFromDict(core.CastMap(data["depositScript"])).Pointer(),
        WithdrawScript: NewScriptSettingFromDict(core.CastMap(data["withdrawScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "priority": p.Priority,
        "shareFree": p.ShareFree,
        "currency": p.Currency,
        "appleKey": p.AppleKey,
        "googleKey": p.GoogleKey,
        "enableFakeReceipt": p.EnableFakeReceipt,
        "createWalletScript": p.CreateWalletScript.ToDict(),
        "depositScript": p.DepositScript.ToDict(),
        "withdrawScript": p.WithdrawScript.ToDict(),
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
    Priority *string `json:"priority"`
    AppleKey *string `json:"appleKey"`
    GoogleKey *string `json:"googleKey"`
    EnableFakeReceipt *bool `json:"enableFakeReceipt"`
    CreateWalletScript *ScriptSetting `json:"createWalletScript"`
    DepositScript *ScriptSetting `json:"depositScript"`
    WithdrawScript *ScriptSetting `json:"withdrawScript"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        Priority: core.CastString(data["priority"]),
        AppleKey: core.CastString(data["appleKey"]),
        GoogleKey: core.CastString(data["googleKey"]),
        EnableFakeReceipt: core.CastBool(data["enableFakeReceipt"]),
        CreateWalletScript: NewScriptSettingFromDict(core.CastMap(data["createWalletScript"])).Pointer(),
        DepositScript: NewScriptSettingFromDict(core.CastMap(data["depositScript"])).Pointer(),
        WithdrawScript: NewScriptSettingFromDict(core.CastMap(data["withdrawScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "priority": p.Priority,
        "appleKey": p.AppleKey,
        "googleKey": p.GoogleKey,
        "enableFakeReceipt": p.EnableFakeReceipt,
        "createWalletScript": p.CreateWalletScript.ToDict(),
        "depositScript": p.DepositScript.ToDict(),
        "withdrawScript": p.WithdrawScript.ToDict(),
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

type DescribeWalletsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeWalletsRequestFromDict(data map[string]interface{}) DescribeWalletsRequest {
    return DescribeWalletsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeWalletsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeWalletsRequest) Pointer() *DescribeWalletsRequest {
    return &p
}

type DescribeWalletsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeWalletsByUserIdRequestFromDict(data map[string]interface{}) DescribeWalletsByUserIdRequest {
    return DescribeWalletsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeWalletsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeWalletsByUserIdRequest) Pointer() *DescribeWalletsByUserIdRequest {
    return &p
}

type GetWalletRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    Slot *int32 `json:"slot"`
}

func NewGetWalletRequestFromDict(data map[string]interface{}) GetWalletRequest {
    return GetWalletRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Slot: core.CastInt32(data["slot"]),
    }
}

func (p GetWalletRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "slot": p.Slot,
    }
}

func (p GetWalletRequest) Pointer() *GetWalletRequest {
    return &p
}

type GetWalletByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Slot *int32 `json:"slot"`
}

func NewGetWalletByUserIdRequestFromDict(data map[string]interface{}) GetWalletByUserIdRequest {
    return GetWalletByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Slot: core.CastInt32(data["slot"]),
    }
}

func (p GetWalletByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "slot": p.Slot,
    }
}

func (p GetWalletByUserIdRequest) Pointer() *GetWalletByUserIdRequest {
    return &p
}

type DepositByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Slot *int32 `json:"slot"`
    Price *float32 `json:"price"`
    Count *int32 `json:"count"`
}

func NewDepositByUserIdRequestFromDict(data map[string]interface{}) DepositByUserIdRequest {
    return DepositByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Slot: core.CastInt32(data["slot"]),
        Price: core.CastFloat32(data["price"]),
        Count: core.CastInt32(data["count"]),
    }
}

func (p DepositByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "slot": p.Slot,
        "price": p.Price,
        "count": p.Count,
    }
}

func (p DepositByUserIdRequest) Pointer() *DepositByUserIdRequest {
    return &p
}

type WithdrawRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    Slot *int32 `json:"slot"`
    Count *int32 `json:"count"`
    PaidOnly *bool `json:"paidOnly"`
}

func NewWithdrawRequestFromDict(data map[string]interface{}) WithdrawRequest {
    return WithdrawRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Slot: core.CastInt32(data["slot"]),
        Count: core.CastInt32(data["count"]),
        PaidOnly: core.CastBool(data["paidOnly"]),
    }
}

func (p WithdrawRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "slot": p.Slot,
        "count": p.Count,
        "paidOnly": p.PaidOnly,
    }
}

func (p WithdrawRequest) Pointer() *WithdrawRequest {
    return &p
}

type WithdrawByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Slot *int32 `json:"slot"`
    Count *int32 `json:"count"`
    PaidOnly *bool `json:"paidOnly"`
}

func NewWithdrawByUserIdRequestFromDict(data map[string]interface{}) WithdrawByUserIdRequest {
    return WithdrawByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Slot: core.CastInt32(data["slot"]),
        Count: core.CastInt32(data["count"]),
        PaidOnly: core.CastBool(data["paidOnly"]),
    }
}

func (p WithdrawByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "slot": p.Slot,
        "count": p.Count,
        "paidOnly": p.PaidOnly,
    }
}

func (p WithdrawByUserIdRequest) Pointer() *WithdrawByUserIdRequest {
    return &p
}

type DepositByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewDepositByStampSheetRequestFromDict(data map[string]interface{}) DepositByStampSheetRequest {
    return DepositByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p DepositByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p DepositByStampSheetRequest) Pointer() *DepositByStampSheetRequest {
    return &p
}

type WithdrawByStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewWithdrawByStampTaskRequestFromDict(data map[string]interface{}) WithdrawByStampTaskRequest {
    return WithdrawByStampTaskRequest {
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p WithdrawByStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p WithdrawByStampTaskRequest) Pointer() *WithdrawByStampTaskRequest {
    return &p
}

type DescribeReceiptsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Slot *int32 `json:"slot"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeReceiptsRequestFromDict(data map[string]interface{}) DescribeReceiptsRequest {
    return DescribeReceiptsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Slot: core.CastInt32(data["slot"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeReceiptsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "slot": p.Slot,
        "begin": p.Begin,
        "end": p.End,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeReceiptsRequest) Pointer() *DescribeReceiptsRequest {
    return &p
}

type GetByUserIdAndTransactionIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TransactionId *string `json:"transactionId"`
}

func NewGetByUserIdAndTransactionIdRequestFromDict(data map[string]interface{}) GetByUserIdAndTransactionIdRequest {
    return GetByUserIdAndTransactionIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TransactionId: core.CastString(data["transactionId"]),
    }
}

func (p GetByUserIdAndTransactionIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "transactionId": p.TransactionId,
    }
}

func (p GetByUserIdAndTransactionIdRequest) Pointer() *GetByUserIdAndTransactionIdRequest {
    return &p
}

type RecordReceiptRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    ContentsId *string `json:"contentsId"`
    Receipt *string `json:"receipt"`
}

func NewRecordReceiptRequestFromDict(data map[string]interface{}) RecordReceiptRequest {
    return RecordReceiptRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        ContentsId: core.CastString(data["contentsId"]),
        Receipt: core.CastString(data["receipt"]),
    }
}

func (p RecordReceiptRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "contentsId": p.ContentsId,
        "receipt": p.Receipt,
    }
}

func (p RecordReceiptRequest) Pointer() *RecordReceiptRequest {
    return &p
}

type RecordReceiptByStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewRecordReceiptByStampTaskRequestFromDict(data map[string]interface{}) RecordReceiptByStampTaskRequest {
    return RecordReceiptByStampTaskRequest {
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p RecordReceiptByStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p RecordReceiptByStampTaskRequest) Pointer() *RecordReceiptByStampTaskRequest {
    return &p
}