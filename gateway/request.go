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

package gateway

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
    FirebaseSecret *string `json:"firebaseSecret"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        FirebaseSecret: core.CastString(data["firebaseSecret"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "firebaseSecret": p.FirebaseSecret,
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
    FirebaseSecret *string `json:"firebaseSecret"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        FirebaseSecret: core.CastString(data["firebaseSecret"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "firebaseSecret": p.FirebaseSecret,
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

type DescribeWebSocketSessionsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeWebSocketSessionsRequestFromDict(data map[string]interface{}) DescribeWebSocketSessionsRequest {
    return DescribeWebSocketSessionsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeWebSocketSessionsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeWebSocketSessionsRequest) Pointer() *DescribeWebSocketSessionsRequest {
    return &p
}

type DescribeWebSocketSessionsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeWebSocketSessionsByUserIdRequestFromDict(data map[string]interface{}) DescribeWebSocketSessionsByUserIdRequest {
    return DescribeWebSocketSessionsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeWebSocketSessionsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeWebSocketSessionsByUserIdRequest) Pointer() *DescribeWebSocketSessionsByUserIdRequest {
    return &p
}

type SetUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    AllowConcurrentAccess *bool `json:"allowConcurrentAccess"`
}

func NewSetUserIdRequestFromDict(data map[string]interface{}) SetUserIdRequest {
    return SetUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        AllowConcurrentAccess: core.CastBool(data["allowConcurrentAccess"]),
    }
}

func (p SetUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "allowConcurrentAccess": p.AllowConcurrentAccess,
    }
}

func (p SetUserIdRequest) Pointer() *SetUserIdRequest {
    return &p
}

type SetUserIdByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    AllowConcurrentAccess *bool `json:"allowConcurrentAccess"`
}

func NewSetUserIdByUserIdRequestFromDict(data map[string]interface{}) SetUserIdByUserIdRequest {
    return SetUserIdByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        AllowConcurrentAccess: core.CastBool(data["allowConcurrentAccess"]),
    }
}

func (p SetUserIdByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "allowConcurrentAccess": p.AllowConcurrentAccess,
    }
}

func (p SetUserIdByUserIdRequest) Pointer() *SetUserIdByUserIdRequest {
    return &p
}

type SendNotificationRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Subject *string `json:"subject"`
    Payload *string `json:"payload"`
    EnableTransferMobileNotification *bool `json:"enableTransferMobileNotification"`
    Sound *string `json:"sound"`
}

func NewSendNotificationRequestFromDict(data map[string]interface{}) SendNotificationRequest {
    return SendNotificationRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Subject: core.CastString(data["subject"]),
        Payload: core.CastString(data["payload"]),
        EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
        Sound: core.CastString(data["sound"]),
    }
}

func (p SendNotificationRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "subject": p.Subject,
        "payload": p.Payload,
        "enableTransferMobileNotification": p.EnableTransferMobileNotification,
        "sound": p.Sound,
    }
}

func (p SendNotificationRequest) Pointer() *SendNotificationRequest {
    return &p
}

type SetFirebaseTokenRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    Token *string `json:"token"`
}

func NewSetFirebaseTokenRequestFromDict(data map[string]interface{}) SetFirebaseTokenRequest {
    return SetFirebaseTokenRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Token: core.CastString(data["token"]),
    }
}

func (p SetFirebaseTokenRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "token": p.Token,
    }
}

func (p SetFirebaseTokenRequest) Pointer() *SetFirebaseTokenRequest {
    return &p
}

type SetFirebaseTokenByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Token *string `json:"token"`
}

func NewSetFirebaseTokenByUserIdRequestFromDict(data map[string]interface{}) SetFirebaseTokenByUserIdRequest {
    return SetFirebaseTokenByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Token: core.CastString(data["token"]),
    }
}

func (p SetFirebaseTokenByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "token": p.Token,
    }
}

func (p SetFirebaseTokenByUserIdRequest) Pointer() *SetFirebaseTokenByUserIdRequest {
    return &p
}

type GetFirebaseTokenRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
}

func NewGetFirebaseTokenRequestFromDict(data map[string]interface{}) GetFirebaseTokenRequest {
    return GetFirebaseTokenRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p GetFirebaseTokenRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
    }
}

func (p GetFirebaseTokenRequest) Pointer() *GetFirebaseTokenRequest {
    return &p
}

type GetFirebaseTokenByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
}

func NewGetFirebaseTokenByUserIdRequestFromDict(data map[string]interface{}) GetFirebaseTokenByUserIdRequest {
    return GetFirebaseTokenByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p GetFirebaseTokenByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
    }
}

func (p GetFirebaseTokenByUserIdRequest) Pointer() *GetFirebaseTokenByUserIdRequest {
    return &p
}

type DeleteFirebaseTokenRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
}

func NewDeleteFirebaseTokenRequestFromDict(data map[string]interface{}) DeleteFirebaseTokenRequest {
    return DeleteFirebaseTokenRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p DeleteFirebaseTokenRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
    }
}

func (p DeleteFirebaseTokenRequest) Pointer() *DeleteFirebaseTokenRequest {
    return &p
}

type DeleteFirebaseTokenByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
}

func NewDeleteFirebaseTokenByUserIdRequestFromDict(data map[string]interface{}) DeleteFirebaseTokenByUserIdRequest {
    return DeleteFirebaseTokenByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p DeleteFirebaseTokenByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
    }
}

func (p DeleteFirebaseTokenByUserIdRequest) Pointer() *DeleteFirebaseTokenByUserIdRequest {
    return &p
}

type SendMobileNotificationByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Subject *string `json:"subject"`
    Payload *string `json:"payload"`
    Sound *string `json:"sound"`
}

func NewSendMobileNotificationByUserIdRequestFromDict(data map[string]interface{}) SendMobileNotificationByUserIdRequest {
    return SendMobileNotificationByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Subject: core.CastString(data["subject"]),
        Payload: core.CastString(data["payload"]),
        Sound: core.CastString(data["sound"]),
    }
}

func (p SendMobileNotificationByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "subject": p.Subject,
        "payload": p.Payload,
        "sound": p.Sound,
    }
}

func (p SendMobileNotificationByUserIdRequest) Pointer() *SendMobileNotificationByUserIdRequest {
    return &p
}