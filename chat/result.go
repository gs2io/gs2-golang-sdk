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

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
	Items         []Namespace `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromJson(data string) DescribeNamespacesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNamespacesResultFromDict(dict)
}

func NewDescribeNamespacesResultFromDict(data map[string]interface{}) DescribeNamespacesResult {
	return DescribeNamespacesResult{
		Items:         CastNamespaces(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewCreateNamespaceResultFromJson(data string) CreateNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceResultFromDict(dict)
}

func NewCreateNamespaceResultFromDict(data map[string]interface{}) CreateNamespaceResult {
	return CreateNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewGetNamespaceStatusResultFromJson(data string) GetNamespaceStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceStatusResultFromDict(dict)
}

func NewGetNamespaceStatusResultFromDict(data map[string]interface{}) GetNamespaceStatusResult {
	return GetNamespaceStatusResult{
		Status: core.CastString(data["status"]),
	}
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewGetNamespaceResultFromJson(data string) GetNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceResultFromDict(dict)
}

func NewGetNamespaceResultFromDict(data map[string]interface{}) GetNamespaceResult {
	return GetNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewUpdateNamespaceResultFromJson(data string) UpdateNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceResultFromDict(dict)
}

func NewUpdateNamespaceResultFromDict(data map[string]interface{}) UpdateNamespaceResult {
	return UpdateNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewDeleteNamespaceResultFromJson(data string) DeleteNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteNamespaceResultFromDict(dict)
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
	return DeleteNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
	return &p
}

type DumpUserDataByUserIdResult struct {
}

type DumpUserDataByUserIdAsyncResult struct {
	result *DumpUserDataByUserIdResult
	err    error
}

func NewDumpUserDataByUserIdResultFromJson(data string) DumpUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdResultFromDict(dict)
}

func NewDumpUserDataByUserIdResultFromDict(data map[string]interface{}) DumpUserDataByUserIdResult {
	return DumpUserDataByUserIdResult{}
}

func (p DumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DumpUserDataByUserIdResult) Pointer() *DumpUserDataByUserIdResult {
	return &p
}

type CheckDumpUserDataByUserIdResult struct {
	Url *string `json:"url"`
}

type CheckDumpUserDataByUserIdAsyncResult struct {
	result *CheckDumpUserDataByUserIdResult
	err    error
}

func NewCheckDumpUserDataByUserIdResultFromJson(data string) CheckDumpUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdResultFromDict(dict)
}

func NewCheckDumpUserDataByUserIdResultFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdResult {
	return CheckDumpUserDataByUserIdResult{
		Url: core.CastString(data["url"]),
	}
}

func (p CheckDumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
	}
}

func (p CheckDumpUserDataByUserIdResult) Pointer() *CheckDumpUserDataByUserIdResult {
	return &p
}

type CleanUserDataByUserIdResult struct {
}

type CleanUserDataByUserIdAsyncResult struct {
	result *CleanUserDataByUserIdResult
	err    error
}

func NewCleanUserDataByUserIdResultFromJson(data string) CleanUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdResultFromDict(dict)
}

func NewCleanUserDataByUserIdResultFromDict(data map[string]interface{}) CleanUserDataByUserIdResult {
	return CleanUserDataByUserIdResult{}
}

func (p CleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p CleanUserDataByUserIdResult) Pointer() *CleanUserDataByUserIdResult {
	return &p
}

type CheckCleanUserDataByUserIdResult struct {
}

type CheckCleanUserDataByUserIdAsyncResult struct {
	result *CheckCleanUserDataByUserIdResult
	err    error
}

func NewCheckCleanUserDataByUserIdResultFromJson(data string) CheckCleanUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdResultFromDict(dict)
}

func NewCheckCleanUserDataByUserIdResultFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdResult {
	return CheckCleanUserDataByUserIdResult{}
}

func (p CheckCleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p CheckCleanUserDataByUserIdResult) Pointer() *CheckCleanUserDataByUserIdResult {
	return &p
}

type PrepareImportUserDataByUserIdResult struct {
	UploadToken *string `json:"uploadToken"`
	UploadUrl   *string `json:"uploadUrl"`
}

type PrepareImportUserDataByUserIdAsyncResult struct {
	result *PrepareImportUserDataByUserIdResult
	err    error
}

func NewPrepareImportUserDataByUserIdResultFromJson(data string) PrepareImportUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdResultFromDict(dict)
}

func NewPrepareImportUserDataByUserIdResultFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdResult {
	return PrepareImportUserDataByUserIdResult{
		UploadToken: core.CastString(data["uploadToken"]),
		UploadUrl:   core.CastString(data["uploadUrl"]),
	}
}

func (p PrepareImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
	}
}

func (p PrepareImportUserDataByUserIdResult) Pointer() *PrepareImportUserDataByUserIdResult {
	return &p
}

type ImportUserDataByUserIdResult struct {
}

type ImportUserDataByUserIdAsyncResult struct {
	result *ImportUserDataByUserIdResult
	err    error
}

func NewImportUserDataByUserIdResultFromJson(data string) ImportUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdResultFromDict(dict)
}

func NewImportUserDataByUserIdResultFromDict(data map[string]interface{}) ImportUserDataByUserIdResult {
	return ImportUserDataByUserIdResult{}
}

func (p ImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p ImportUserDataByUserIdResult) Pointer() *ImportUserDataByUserIdResult {
	return &p
}

type CheckImportUserDataByUserIdResult struct {
	Url *string `json:"url"`
}

type CheckImportUserDataByUserIdAsyncResult struct {
	result *CheckImportUserDataByUserIdResult
	err    error
}

func NewCheckImportUserDataByUserIdResultFromJson(data string) CheckImportUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdResultFromDict(dict)
}

func NewCheckImportUserDataByUserIdResultFromDict(data map[string]interface{}) CheckImportUserDataByUserIdResult {
	return CheckImportUserDataByUserIdResult{
		Url: core.CastString(data["url"]),
	}
}

func (p CheckImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
	}
}

func (p CheckImportUserDataByUserIdResult) Pointer() *CheckImportUserDataByUserIdResult {
	return &p
}

type DescribeRoomsResult struct {
	Items         []Room  `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeRoomsAsyncResult struct {
	result *DescribeRoomsResult
	err    error
}

func NewDescribeRoomsResultFromJson(data string) DescribeRoomsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRoomsResultFromDict(dict)
}

func NewDescribeRoomsResultFromDict(data map[string]interface{}) DescribeRoomsResult {
	return DescribeRoomsResult{
		Items:         CastRooms(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRoomsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewCreateRoomResultFromJson(data string) CreateRoomResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRoomResultFromDict(dict)
}

func NewCreateRoomResultFromDict(data map[string]interface{}) CreateRoomResult {
	return CreateRoomResult{
		Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateRoomResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewCreateRoomFromBackendResultFromJson(data string) CreateRoomFromBackendResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRoomFromBackendResultFromDict(dict)
}

func NewCreateRoomFromBackendResultFromDict(data map[string]interface{}) CreateRoomFromBackendResult {
	return CreateRoomFromBackendResult{
		Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateRoomFromBackendResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewGetRoomResultFromJson(data string) GetRoomResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRoomResultFromDict(dict)
}

func NewGetRoomResultFromDict(data map[string]interface{}) GetRoomResult {
	return GetRoomResult{
		Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRoomResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewUpdateRoomResultFromJson(data string) UpdateRoomResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRoomResultFromDict(dict)
}

func NewUpdateRoomResultFromDict(data map[string]interface{}) UpdateRoomResult {
	return UpdateRoomResult{
		Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateRoomResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateRoomResult) Pointer() *UpdateRoomResult {
	return &p
}

type UpdateRoomFromBackendResult struct {
	Item *Room `json:"item"`
}

type UpdateRoomFromBackendAsyncResult struct {
	result *UpdateRoomFromBackendResult
	err    error
}

func NewUpdateRoomFromBackendResultFromJson(data string) UpdateRoomFromBackendResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRoomFromBackendResultFromDict(dict)
}

func NewUpdateRoomFromBackendResultFromDict(data map[string]interface{}) UpdateRoomFromBackendResult {
	return UpdateRoomFromBackendResult{
		Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateRoomFromBackendResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateRoomFromBackendResult) Pointer() *UpdateRoomFromBackendResult {
	return &p
}

type DeleteRoomResult struct {
	Item *Room `json:"item"`
}

type DeleteRoomAsyncResult struct {
	result *DeleteRoomResult
	err    error
}

func NewDeleteRoomResultFromJson(data string) DeleteRoomResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRoomResultFromDict(dict)
}

func NewDeleteRoomResultFromDict(data map[string]interface{}) DeleteRoomResult {
	return DeleteRoomResult{
		Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteRoomResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewDeleteRoomFromBackendResultFromJson(data string) DeleteRoomFromBackendResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRoomFromBackendResultFromDict(dict)
}

func NewDeleteRoomFromBackendResultFromDict(data map[string]interface{}) DeleteRoomFromBackendResult {
	return DeleteRoomFromBackendResult{
		Item: NewRoomFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteRoomFromBackendResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewDescribeMessagesResultFromJson(data string) DescribeMessagesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMessagesResultFromDict(dict)
}

func NewDescribeMessagesResultFromDict(data map[string]interface{}) DescribeMessagesResult {
	return DescribeMessagesResult{
		Items: CastMessages(core.CastArray(data["items"])),
	}
}

func (p DescribeMessagesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMessagesFromDict(
			p.Items,
		),
	}
}

func (p DescribeMessagesResult) Pointer() *DescribeMessagesResult {
	return &p
}

type DescribeMessagesByUserIdResult struct {
	Items []Message `json:"items"`
}

type DescribeMessagesByUserIdAsyncResult struct {
	result *DescribeMessagesByUserIdResult
	err    error
}

func NewDescribeMessagesByUserIdResultFromJson(data string) DescribeMessagesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMessagesByUserIdResultFromDict(dict)
}

func NewDescribeMessagesByUserIdResultFromDict(data map[string]interface{}) DescribeMessagesByUserIdResult {
	return DescribeMessagesByUserIdResult{
		Items: CastMessages(core.CastArray(data["items"])),
	}
}

func (p DescribeMessagesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMessagesFromDict(
			p.Items,
		),
	}
}

func (p DescribeMessagesByUserIdResult) Pointer() *DescribeMessagesByUserIdResult {
	return &p
}

type PostResult struct {
	Item *Message `json:"item"`
}

type PostAsyncResult struct {
	result *PostResult
	err    error
}

func NewPostResultFromJson(data string) PostResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPostResultFromDict(dict)
}

func NewPostResultFromDict(data map[string]interface{}) PostResult {
	return PostResult{
		Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p PostResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewPostByUserIdResultFromJson(data string) PostByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPostByUserIdResultFromDict(dict)
}

func NewPostByUserIdResultFromDict(data map[string]interface{}) PostByUserIdResult {
	return PostByUserIdResult{
		Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p PostByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewGetMessageResultFromJson(data string) GetMessageResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMessageResultFromDict(dict)
}

func NewGetMessageResultFromDict(data map[string]interface{}) GetMessageResult {
	return GetMessageResult{
		Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMessageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMessageResult) Pointer() *GetMessageResult {
	return &p
}

type GetMessageByUserIdResult struct {
	Item *Message `json:"item"`
}

type GetMessageByUserIdAsyncResult struct {
	result *GetMessageByUserIdResult
	err    error
}

func NewGetMessageByUserIdResultFromJson(data string) GetMessageByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMessageByUserIdResultFromDict(dict)
}

func NewGetMessageByUserIdResultFromDict(data map[string]interface{}) GetMessageByUserIdResult {
	return GetMessageByUserIdResult{
		Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMessageByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMessageByUserIdResult) Pointer() *GetMessageByUserIdResult {
	return &p
}

type DeleteMessageResult struct {
	Item *Message `json:"item"`
}

type DeleteMessageAsyncResult struct {
	result *DeleteMessageResult
	err    error
}

func NewDeleteMessageResultFromJson(data string) DeleteMessageResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMessageResultFromDict(dict)
}

func NewDeleteMessageResultFromDict(data map[string]interface{}) DeleteMessageResult {
	return DeleteMessageResult{
		Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteMessageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteMessageResult) Pointer() *DeleteMessageResult {
	return &p
}

type DescribeSubscribesResult struct {
	Items         []Subscribe `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeSubscribesAsyncResult struct {
	result *DescribeSubscribesResult
	err    error
}

func NewDescribeSubscribesResultFromJson(data string) DescribeSubscribesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesResultFromDict(dict)
}

func NewDescribeSubscribesResultFromDict(data map[string]interface{}) DescribeSubscribesResult {
	return DescribeSubscribesResult{
		Items:         CastSubscribes(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeSubscribesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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
	Items         []Subscribe `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeSubscribesByUserIdAsyncResult struct {
	result *DescribeSubscribesByUserIdResult
	err    error
}

func NewDescribeSubscribesByUserIdResultFromJson(data string) DescribeSubscribesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesByUserIdResultFromDict(dict)
}

func NewDescribeSubscribesByUserIdResultFromDict(data map[string]interface{}) DescribeSubscribesByUserIdResult {
	return DescribeSubscribesByUserIdResult{
		Items:         CastSubscribes(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeSubscribesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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
	Items         []Subscribe `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeSubscribesByRoomNameAsyncResult struct {
	result *DescribeSubscribesByRoomNameResult
	err    error
}

func NewDescribeSubscribesByRoomNameResultFromJson(data string) DescribeSubscribesByRoomNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesByRoomNameResultFromDict(dict)
}

func NewDescribeSubscribesByRoomNameResultFromDict(data map[string]interface{}) DescribeSubscribesByRoomNameResult {
	return DescribeSubscribesByRoomNameResult{
		Items:         CastSubscribes(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeSubscribesByRoomNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewSubscribeResultFromJson(data string) SubscribeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubscribeResultFromDict(dict)
}

func NewSubscribeResultFromDict(data map[string]interface{}) SubscribeResult {
	return SubscribeResult{
		Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p SubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewSubscribeByUserIdResultFromJson(data string) SubscribeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubscribeByUserIdResultFromDict(dict)
}

func NewSubscribeByUserIdResultFromDict(data map[string]interface{}) SubscribeByUserIdResult {
	return SubscribeByUserIdResult{
		Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p SubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewGetSubscribeResultFromJson(data string) GetSubscribeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeResultFromDict(dict)
}

func NewGetSubscribeResultFromDict(data map[string]interface{}) GetSubscribeResult {
	return GetSubscribeResult{
		Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetSubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewGetSubscribeByUserIdResultFromJson(data string) GetSubscribeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeByUserIdResultFromDict(dict)
}

func NewGetSubscribeByUserIdResultFromDict(data map[string]interface{}) GetSubscribeByUserIdResult {
	return GetSubscribeByUserIdResult{
		Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetSubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewUpdateNotificationTypeResultFromJson(data string) UpdateNotificationTypeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNotificationTypeResultFromDict(dict)
}

func NewUpdateNotificationTypeResultFromDict(data map[string]interface{}) UpdateNotificationTypeResult {
	return UpdateNotificationTypeResult{
		Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateNotificationTypeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewUpdateNotificationTypeByUserIdResultFromJson(data string) UpdateNotificationTypeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNotificationTypeByUserIdResultFromDict(dict)
}

func NewUpdateNotificationTypeByUserIdResultFromDict(data map[string]interface{}) UpdateNotificationTypeByUserIdResult {
	return UpdateNotificationTypeByUserIdResult{
		Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateNotificationTypeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewUnsubscribeResultFromJson(data string) UnsubscribeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnsubscribeResultFromDict(dict)
}

func NewUnsubscribeResultFromDict(data map[string]interface{}) UnsubscribeResult {
	return UnsubscribeResult{
		Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UnsubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewUnsubscribeByUserIdResultFromJson(data string) UnsubscribeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnsubscribeByUserIdResultFromDict(dict)
}

func NewUnsubscribeByUserIdResultFromDict(data map[string]interface{}) UnsubscribeByUserIdResult {
	return UnsubscribeByUserIdResult{
		Item: NewSubscribeFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UnsubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UnsubscribeByUserIdResult) Pointer() *UnsubscribeByUserIdResult {
	return &p
}
