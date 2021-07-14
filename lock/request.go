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

package lock

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
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
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
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
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

type DescribeMutexesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeMutexesRequestFromDict(data map[string]interface{}) DescribeMutexesRequest {
    return DescribeMutexesRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeMutexesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeMutexesRequest) Pointer() *DescribeMutexesRequest {
    return &p
}

type DescribeMutexesByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeMutexesByUserIdRequestFromDict(data map[string]interface{}) DescribeMutexesByUserIdRequest {
    return DescribeMutexesByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeMutexesByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeMutexesByUserIdRequest) Pointer() *DescribeMutexesByUserIdRequest {
    return &p
}

type LockRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PropertyId *string `json:"propertyId"`
    AccessToken *string `json:"accessToken"`
    TransactionId *string `json:"transactionId"`
    Ttl *int64 `json:"ttl"`
}

func NewLockRequestFromDict(data map[string]interface{}) LockRequest {
    return LockRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PropertyId: core.CastString(data["propertyId"]),
        AccessToken: core.CastString(data["accessToken"]),
        TransactionId: core.CastString(data["transactionId"]),
        Ttl: core.CastInt64(data["ttl"]),
    }
}

func (p LockRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "propertyId": p.PropertyId,
        "accessToken": p.AccessToken,
        "transactionId": p.TransactionId,
        "ttl": p.Ttl,
    }
}

func (p LockRequest) Pointer() *LockRequest {
    return &p
}

type LockByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PropertyId *string `json:"propertyId"`
    UserId *string `json:"userId"`
    TransactionId *string `json:"transactionId"`
    Ttl *int64 `json:"ttl"`
}

func NewLockByUserIdRequestFromDict(data map[string]interface{}) LockByUserIdRequest {
    return LockByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PropertyId: core.CastString(data["propertyId"]),
        UserId: core.CastString(data["userId"]),
        TransactionId: core.CastString(data["transactionId"]),
        Ttl: core.CastInt64(data["ttl"]),
    }
}

func (p LockByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "propertyId": p.PropertyId,
        "userId": p.UserId,
        "transactionId": p.TransactionId,
        "ttl": p.Ttl,
    }
}

func (p LockByUserIdRequest) Pointer() *LockByUserIdRequest {
    return &p
}

type UnlockRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PropertyId *string `json:"propertyId"`
    AccessToken *string `json:"accessToken"`
    TransactionId *string `json:"transactionId"`
}

func NewUnlockRequestFromDict(data map[string]interface{}) UnlockRequest {
    return UnlockRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PropertyId: core.CastString(data["propertyId"]),
        AccessToken: core.CastString(data["accessToken"]),
        TransactionId: core.CastString(data["transactionId"]),
    }
}

func (p UnlockRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "propertyId": p.PropertyId,
        "accessToken": p.AccessToken,
        "transactionId": p.TransactionId,
    }
}

func (p UnlockRequest) Pointer() *UnlockRequest {
    return &p
}

type UnlockByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PropertyId *string `json:"propertyId"`
    UserId *string `json:"userId"`
    TransactionId *string `json:"transactionId"`
}

func NewUnlockByUserIdRequestFromDict(data map[string]interface{}) UnlockByUserIdRequest {
    return UnlockByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PropertyId: core.CastString(data["propertyId"]),
        UserId: core.CastString(data["userId"]),
        TransactionId: core.CastString(data["transactionId"]),
    }
}

func (p UnlockByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "propertyId": p.PropertyId,
        "userId": p.UserId,
        "transactionId": p.TransactionId,
    }
}

func (p UnlockByUserIdRequest) Pointer() *UnlockByUserIdRequest {
    return &p
}

type GetMutexRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PropertyId *string `json:"propertyId"`
}

func NewGetMutexRequestFromDict(data map[string]interface{}) GetMutexRequest {
    return GetMutexRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PropertyId: core.CastString(data["propertyId"]),
    }
}

func (p GetMutexRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "propertyId": p.PropertyId,
    }
}

func (p GetMutexRequest) Pointer() *GetMutexRequest {
    return &p
}

type GetMutexByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PropertyId *string `json:"propertyId"`
}

func NewGetMutexByUserIdRequestFromDict(data map[string]interface{}) GetMutexByUserIdRequest {
    return GetMutexByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PropertyId: core.CastString(data["propertyId"]),
    }
}

func (p GetMutexByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "propertyId": p.PropertyId,
    }
}

func (p GetMutexByUserIdRequest) Pointer() *GetMutexByUserIdRequest {
    return &p
}

type DeleteMutexByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PropertyId *string `json:"propertyId"`
}

func NewDeleteMutexByUserIdRequestFromDict(data map[string]interface{}) DeleteMutexByUserIdRequest {
    return DeleteMutexByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PropertyId: core.CastString(data["propertyId"]),
    }
}

func (p DeleteMutexByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "propertyId": p.PropertyId,
    }
}

func (p DeleteMutexByUserIdRequest) Pointer() *DeleteMutexByUserIdRequest {
    return &p
}