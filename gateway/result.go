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
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
    return DeleteNamespaceResult {
    }
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
    return &p
}

type DescribeWebSocketSessionsResult struct {
    Items []WebSocketSession `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeWebSocketSessionsAsyncResult struct {
	result *DescribeWebSocketSessionsResult
	err    error
}

func NewDescribeWebSocketSessionsResultFromDict(data map[string]interface{}) DescribeWebSocketSessionsResult {
    return DescribeWebSocketSessionsResult {
        Items: CastWebSocketSessions(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeWebSocketSessionsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastWebSocketSessionsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeWebSocketSessionsResult) Pointer() *DescribeWebSocketSessionsResult {
    return &p
}

type DescribeWebSocketSessionsByUserIdResult struct {
    Items []WebSocketSession `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeWebSocketSessionsByUserIdAsyncResult struct {
	result *DescribeWebSocketSessionsByUserIdResult
	err    error
}

func NewDescribeWebSocketSessionsByUserIdResultFromDict(data map[string]interface{}) DescribeWebSocketSessionsByUserIdResult {
    return DescribeWebSocketSessionsByUserIdResult {
        Items: CastWebSocketSessions(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeWebSocketSessionsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastWebSocketSessionsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeWebSocketSessionsByUserIdResult) Pointer() *DescribeWebSocketSessionsByUserIdResult {
    return &p
}

type SetUserIdResult struct {
    Item *WebSocketSession `json:"item"`
}

type SetUserIdAsyncResult struct {
	result *SetUserIdResult
	err    error
}

func NewSetUserIdResultFromDict(data map[string]interface{}) SetUserIdResult {
    return SetUserIdResult {
        Item: NewWebSocketSessionFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SetUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SetUserIdResult) Pointer() *SetUserIdResult {
    return &p
}

type SetUserIdByUserIdResult struct {
    Item *WebSocketSession `json:"item"`
}

type SetUserIdByUserIdAsyncResult struct {
	result *SetUserIdByUserIdResult
	err    error
}

func NewSetUserIdByUserIdResultFromDict(data map[string]interface{}) SetUserIdByUserIdResult {
    return SetUserIdByUserIdResult {
        Item: NewWebSocketSessionFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SetUserIdByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SetUserIdByUserIdResult) Pointer() *SetUserIdByUserIdResult {
    return &p
}

type SendNotificationResult struct {
    Protocol *string `json:"protocol"`
}

type SendNotificationAsyncResult struct {
	result *SendNotificationResult
	err    error
}

func NewSendNotificationResultFromDict(data map[string]interface{}) SendNotificationResult {
    return SendNotificationResult {
        Protocol: core.CastString(data["protocol"]),
    }
}

func (p SendNotificationResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "protocol": p.Protocol,
    }
}

func (p SendNotificationResult) Pointer() *SendNotificationResult {
    return &p
}

type SetFirebaseTokenResult struct {
    Item *FirebaseToken `json:"item"`
}

type SetFirebaseTokenAsyncResult struct {
	result *SetFirebaseTokenResult
	err    error
}

func NewSetFirebaseTokenResultFromDict(data map[string]interface{}) SetFirebaseTokenResult {
    return SetFirebaseTokenResult {
        Item: NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SetFirebaseTokenResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SetFirebaseTokenResult) Pointer() *SetFirebaseTokenResult {
    return &p
}

type SetFirebaseTokenByUserIdResult struct {
    Item *FirebaseToken `json:"item"`
}

type SetFirebaseTokenByUserIdAsyncResult struct {
	result *SetFirebaseTokenByUserIdResult
	err    error
}

func NewSetFirebaseTokenByUserIdResultFromDict(data map[string]interface{}) SetFirebaseTokenByUserIdResult {
    return SetFirebaseTokenByUserIdResult {
        Item: NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SetFirebaseTokenByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SetFirebaseTokenByUserIdResult) Pointer() *SetFirebaseTokenByUserIdResult {
    return &p
}

type GetFirebaseTokenResult struct {
    Item *FirebaseToken `json:"item"`
}

type GetFirebaseTokenAsyncResult struct {
	result *GetFirebaseTokenResult
	err    error
}

func NewGetFirebaseTokenResultFromDict(data map[string]interface{}) GetFirebaseTokenResult {
    return GetFirebaseTokenResult {
        Item: NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetFirebaseTokenResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetFirebaseTokenResult) Pointer() *GetFirebaseTokenResult {
    return &p
}

type GetFirebaseTokenByUserIdResult struct {
    Item *FirebaseToken `json:"item"`
}

type GetFirebaseTokenByUserIdAsyncResult struct {
	result *GetFirebaseTokenByUserIdResult
	err    error
}

func NewGetFirebaseTokenByUserIdResultFromDict(data map[string]interface{}) GetFirebaseTokenByUserIdResult {
    return GetFirebaseTokenByUserIdResult {
        Item: NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetFirebaseTokenByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetFirebaseTokenByUserIdResult) Pointer() *GetFirebaseTokenByUserIdResult {
    return &p
}

type DeleteFirebaseTokenResult struct {
    Item *FirebaseToken `json:"item"`
}

type DeleteFirebaseTokenAsyncResult struct {
	result *DeleteFirebaseTokenResult
	err    error
}

func NewDeleteFirebaseTokenResultFromDict(data map[string]interface{}) DeleteFirebaseTokenResult {
    return DeleteFirebaseTokenResult {
        Item: NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteFirebaseTokenResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteFirebaseTokenResult) Pointer() *DeleteFirebaseTokenResult {
    return &p
}

type DeleteFirebaseTokenByUserIdResult struct {
    Item *FirebaseToken `json:"item"`
}

type DeleteFirebaseTokenByUserIdAsyncResult struct {
	result *DeleteFirebaseTokenByUserIdResult
	err    error
}

func NewDeleteFirebaseTokenByUserIdResultFromDict(data map[string]interface{}) DeleteFirebaseTokenByUserIdResult {
    return DeleteFirebaseTokenByUserIdResult {
        Item: NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteFirebaseTokenByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteFirebaseTokenByUserIdResult) Pointer() *DeleteFirebaseTokenByUserIdResult {
    return &p
}

type SendMobileNotificationByUserIdResult struct {
}

type SendMobileNotificationByUserIdAsyncResult struct {
	result *SendMobileNotificationByUserIdResult
	err    error
}

func NewSendMobileNotificationByUserIdResultFromDict(data map[string]interface{}) SendMobileNotificationByUserIdResult {
    return SendMobileNotificationByUserIdResult {
    }
}

func (p SendMobileNotificationByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p SendMobileNotificationByUserIdResult) Pointer() *SendMobileNotificationByUserIdResult {
    return &p
}