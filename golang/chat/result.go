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

package chat

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

type DescribeRoomsResult struct {
    Items []Room `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeRoomsAsyncResult struct {
	result *DescribeRoomsResult
	err    error
}

func NewDescribeRoomsResultFromDict(data map[string]interface{}) DescribeRoomsResult {
    return DescribeRoomsResult {
        Items: CastRooms(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeRoomsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRoomsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeRoomsResult) Pointer() *DescribeRoomsResult {
    return &p
}

type CreateRoomResult struct {
    Item *Room `json:"item"`
}

type CreateRoomAsyncResult struct {
	result *CreateRoomResult
	err    error
}

func NewCreateRoomResultFromDict(data map[string]interface{}) CreateRoomResult {
    return CreateRoomResult {
        Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateRoomResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateRoomResult) Pointer() *CreateRoomResult {
    return &p
}

type CreateRoomFromBackendResult struct {
    Item *Room `json:"item"`
}

type CreateRoomFromBackendAsyncResult struct {
	result *CreateRoomFromBackendResult
	err    error
}

func NewCreateRoomFromBackendResultFromDict(data map[string]interface{}) CreateRoomFromBackendResult {
    return CreateRoomFromBackendResult {
        Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateRoomFromBackendResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateRoomFromBackendResult) Pointer() *CreateRoomFromBackendResult {
    return &p
}

type GetRoomResult struct {
    Item *Room `json:"item"`
}

type GetRoomAsyncResult struct {
	result *GetRoomResult
	err    error
}

func NewGetRoomResultFromDict(data map[string]interface{}) GetRoomResult {
    return GetRoomResult {
        Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRoomResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRoomResult) Pointer() *GetRoomResult {
    return &p
}

type UpdateRoomResult struct {
    Item *Room `json:"item"`
}

type UpdateRoomAsyncResult struct {
	result *UpdateRoomResult
	err    error
}

func NewUpdateRoomResultFromDict(data map[string]interface{}) UpdateRoomResult {
    return UpdateRoomResult {
        Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateRoomResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateRoomResult) Pointer() *UpdateRoomResult {
    return &p
}

type DeleteRoomResult struct {
    Item *Room `json:"item"`
}

type DeleteRoomAsyncResult struct {
	result *DeleteRoomResult
	err    error
}

func NewDeleteRoomResultFromDict(data map[string]interface{}) DeleteRoomResult {
    return DeleteRoomResult {
        Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteRoomResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteRoomResult) Pointer() *DeleteRoomResult {
    return &p
}

type DeleteRoomFromBackendResult struct {
    Item *Room `json:"item"`
}

type DeleteRoomFromBackendAsyncResult struct {
	result *DeleteRoomFromBackendResult
	err    error
}

func NewDeleteRoomFromBackendResultFromDict(data map[string]interface{}) DeleteRoomFromBackendResult {
    return DeleteRoomFromBackendResult {
        Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteRoomFromBackendResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteRoomFromBackendResult) Pointer() *DeleteRoomFromBackendResult {
    return &p
}

type DescribeMessagesResult struct {
    Items []Message `json:"items"`
}

type DescribeMessagesAsyncResult struct {
	result *DescribeMessagesResult
	err    error
}

func NewDescribeMessagesResultFromDict(data map[string]interface{}) DescribeMessagesResult {
    return DescribeMessagesResult {
        Items: CastMessages(core.CastArray(data["items"])),
    }
}

func (p DescribeMessagesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastMessagesFromDict(
            p.Items,
        ),
    }
}

func (p DescribeMessagesResult) Pointer() *DescribeMessagesResult {
    return &p
}

type PostResult struct {
    Item *Message `json:"item"`
}

type PostAsyncResult struct {
	result *PostResult
	err    error
}

func NewPostResultFromDict(data map[string]interface{}) PostResult {
    return PostResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p PostResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p PostResult) Pointer() *PostResult {
    return &p
}

type PostByUserIdResult struct {
    Item *Message `json:"item"`
}

type PostByUserIdAsyncResult struct {
	result *PostByUserIdResult
	err    error
}

func NewPostByUserIdResultFromDict(data map[string]interface{}) PostByUserIdResult {
    return PostByUserIdResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p PostByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p PostByUserIdResult) Pointer() *PostByUserIdResult {
    return &p
}

type GetMessageResult struct {
    Item *Message `json:"item"`
}

type GetMessageAsyncResult struct {
	result *GetMessageResult
	err    error
}

func NewGetMessageResultFromDict(data map[string]interface{}) GetMessageResult {
    return GetMessageResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetMessageResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetMessageResult) Pointer() *GetMessageResult {
    return &p
}

type DeleteMessageResult struct {
    Item *Message `json:"item"`
}

type DeleteMessageAsyncResult struct {
	result *DeleteMessageResult
	err    error
}

func NewDeleteMessageResultFromDict(data map[string]interface{}) DeleteMessageResult {
    return DeleteMessageResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteMessageResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteMessageResult) Pointer() *DeleteMessageResult {
    return &p
}

type DescribeSubscribesResult struct {
    Items []Subscribe `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSubscribesAsyncResult struct {
	result *DescribeSubscribesResult
	err    error
}

func NewDescribeSubscribesResultFromDict(data map[string]interface{}) DescribeSubscribesResult {
    return DescribeSubscribesResult {
        Items: CastSubscribes(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSubscribesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSubscribesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSubscribesResult) Pointer() *DescribeSubscribesResult {
    return &p
}

type DescribeSubscribesByUserIdResult struct {
    Items []Subscribe `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSubscribesByUserIdAsyncResult struct {
	result *DescribeSubscribesByUserIdResult
	err    error
}

func NewDescribeSubscribesByUserIdResultFromDict(data map[string]interface{}) DescribeSubscribesByUserIdResult {
    return DescribeSubscribesByUserIdResult {
        Items: CastSubscribes(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSubscribesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSubscribesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSubscribesByUserIdResult) Pointer() *DescribeSubscribesByUserIdResult {
    return &p
}

type DescribeSubscribesByRoomNameResult struct {
    Items []Subscribe `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSubscribesByRoomNameAsyncResult struct {
	result *DescribeSubscribesByRoomNameResult
	err    error
}

func NewDescribeSubscribesByRoomNameResultFromDict(data map[string]interface{}) DescribeSubscribesByRoomNameResult {
    return DescribeSubscribesByRoomNameResult {
        Items: CastSubscribes(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSubscribesByRoomNameResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSubscribesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSubscribesByRoomNameResult) Pointer() *DescribeSubscribesByRoomNameResult {
    return &p
}

type SubscribeResult struct {
    Item *Subscribe `json:"item"`
}

type SubscribeAsyncResult struct {
	result *SubscribeResult
	err    error
}

func NewSubscribeResultFromDict(data map[string]interface{}) SubscribeResult {
    return SubscribeResult {
        Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SubscribeResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SubscribeResult) Pointer() *SubscribeResult {
    return &p
}

type SubscribeByUserIdResult struct {
    Item *Subscribe `json:"item"`
}

type SubscribeByUserIdAsyncResult struct {
	result *SubscribeByUserIdResult
	err    error
}

func NewSubscribeByUserIdResultFromDict(data map[string]interface{}) SubscribeByUserIdResult {
    return SubscribeByUserIdResult {
        Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SubscribeByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SubscribeByUserIdResult) Pointer() *SubscribeByUserIdResult {
    return &p
}

type GetSubscribeResult struct {
    Item *Subscribe `json:"item"`
}

type GetSubscribeAsyncResult struct {
	result *GetSubscribeResult
	err    error
}

func NewGetSubscribeResultFromDict(data map[string]interface{}) GetSubscribeResult {
    return GetSubscribeResult {
        Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSubscribeResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSubscribeResult) Pointer() *GetSubscribeResult {
    return &p
}

type GetSubscribeByUserIdResult struct {
    Item *Subscribe `json:"item"`
}

type GetSubscribeByUserIdAsyncResult struct {
	result *GetSubscribeByUserIdResult
	err    error
}

func NewGetSubscribeByUserIdResultFromDict(data map[string]interface{}) GetSubscribeByUserIdResult {
    return GetSubscribeByUserIdResult {
        Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSubscribeByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSubscribeByUserIdResult) Pointer() *GetSubscribeByUserIdResult {
    return &p
}

type UpdateNotificationTypeResult struct {
    Item *Subscribe `json:"item"`
}

type UpdateNotificationTypeAsyncResult struct {
	result *UpdateNotificationTypeResult
	err    error
}

func NewUpdateNotificationTypeResultFromDict(data map[string]interface{}) UpdateNotificationTypeResult {
    return UpdateNotificationTypeResult {
        Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateNotificationTypeResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateNotificationTypeResult) Pointer() *UpdateNotificationTypeResult {
    return &p
}

type UpdateNotificationTypeByUserIdResult struct {
    Item *Subscribe `json:"item"`
}

type UpdateNotificationTypeByUserIdAsyncResult struct {
	result *UpdateNotificationTypeByUserIdResult
	err    error
}

func NewUpdateNotificationTypeByUserIdResultFromDict(data map[string]interface{}) UpdateNotificationTypeByUserIdResult {
    return UpdateNotificationTypeByUserIdResult {
        Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateNotificationTypeByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateNotificationTypeByUserIdResult) Pointer() *UpdateNotificationTypeByUserIdResult {
    return &p
}

type UnsubscribeResult struct {
    Item *Subscribe `json:"item"`
}

type UnsubscribeAsyncResult struct {
	result *UnsubscribeResult
	err    error
}

func NewUnsubscribeResultFromDict(data map[string]interface{}) UnsubscribeResult {
    return UnsubscribeResult {
        Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UnsubscribeResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UnsubscribeResult) Pointer() *UnsubscribeResult {
    return &p
}

type UnsubscribeByUserIdResult struct {
    Item *Subscribe `json:"item"`
}

type UnsubscribeByUserIdAsyncResult struct {
	result *UnsubscribeByUserIdResult
	err    error
}

func NewUnsubscribeByUserIdResultFromDict(data map[string]interface{}) UnsubscribeByUserIdResult {
    return UnsubscribeByUserIdResult {
        Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UnsubscribeByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UnsubscribeByUserIdResult) Pointer() *UnsubscribeByUserIdResult {
    return &p
}