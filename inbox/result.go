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

package inbox

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
		Items: func() []Namespace {
			if data["items"] == nil {
				return nil
			}
			return CastNamespaces(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		Url: func() *string {
			v, ok := data["url"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["url"])
		}(),
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
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
		UploadUrl: func() *string {
			v, ok := data["uploadUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadUrl"])
		}(),
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
		Url: func() *string {
			v, ok := data["url"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["url"])
		}(),
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

type DescribeMessagesResult struct {
	Items         []Message `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
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
		Items: func() []Message {
			if data["items"] == nil {
				return nil
			}
			return CastMessages(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeMessagesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMessagesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeMessagesResult) Pointer() *DescribeMessagesResult {
	return &p
}

type DescribeMessagesByUserIdResult struct {
	Items         []Message `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
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
		Items: func() []Message {
			if data["items"] == nil {
				return nil
			}
			return CastMessages(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeMessagesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMessagesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeMessagesByUserIdResult) Pointer() *DescribeMessagesByUserIdResult {
	return &p
}

type SendMessageByUserIdResult struct {
	Item *Message `json:"item"`
}

type SendMessageByUserIdAsyncResult struct {
	result *SendMessageByUserIdResult
	err    error
}

func NewSendMessageByUserIdResultFromJson(data string) SendMessageByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendMessageByUserIdResultFromDict(dict)
}

func NewSendMessageByUserIdResultFromDict(data map[string]interface{}) SendMessageByUserIdResult {
	return SendMessageByUserIdResult{
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SendMessageByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SendMessageByUserIdResult) Pointer() *SendMessageByUserIdResult {
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
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetMessageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetMessageByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetMessageByUserIdResult) Pointer() *GetMessageByUserIdResult {
	return &p
}

type ReceiveGlobalMessageResult struct {
	Item []Message `json:"item"`
}

type ReceiveGlobalMessageAsyncResult struct {
	result *ReceiveGlobalMessageResult
	err    error
}

func NewReceiveGlobalMessageResultFromJson(data string) ReceiveGlobalMessageResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveGlobalMessageResultFromDict(dict)
}

func NewReceiveGlobalMessageResultFromDict(data map[string]interface{}) ReceiveGlobalMessageResult {
	return ReceiveGlobalMessageResult{
		Item: func() []Message {
			if data["item"] == nil {
				return nil
			}
			return CastMessages(core.CastArray(data["item"]))
		}(),
	}
}

func (p ReceiveGlobalMessageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": CastMessagesFromDict(
			p.Item,
		),
	}
}

func (p ReceiveGlobalMessageResult) Pointer() *ReceiveGlobalMessageResult {
	return &p
}

type ReceiveGlobalMessageByUserIdResult struct {
	Item []Message `json:"item"`
}

type ReceiveGlobalMessageByUserIdAsyncResult struct {
	result *ReceiveGlobalMessageByUserIdResult
	err    error
}

func NewReceiveGlobalMessageByUserIdResultFromJson(data string) ReceiveGlobalMessageByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveGlobalMessageByUserIdResultFromDict(dict)
}

func NewReceiveGlobalMessageByUserIdResultFromDict(data map[string]interface{}) ReceiveGlobalMessageByUserIdResult {
	return ReceiveGlobalMessageByUserIdResult{
		Item: func() []Message {
			if data["item"] == nil {
				return nil
			}
			return CastMessages(core.CastArray(data["item"]))
		}(),
	}
}

func (p ReceiveGlobalMessageByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": CastMessagesFromDict(
			p.Item,
		),
	}
}

func (p ReceiveGlobalMessageByUserIdResult) Pointer() *ReceiveGlobalMessageByUserIdResult {
	return &p
}

type OpenMessageResult struct {
	Item *Message `json:"item"`
}

type OpenMessageAsyncResult struct {
	result *OpenMessageResult
	err    error
}

func NewOpenMessageResultFromJson(data string) OpenMessageResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewOpenMessageResultFromDict(dict)
}

func NewOpenMessageResultFromDict(data map[string]interface{}) OpenMessageResult {
	return OpenMessageResult{
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p OpenMessageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p OpenMessageResult) Pointer() *OpenMessageResult {
	return &p
}

type OpenMessageByUserIdResult struct {
	Item *Message `json:"item"`
}

type OpenMessageByUserIdAsyncResult struct {
	result *OpenMessageByUserIdResult
	err    error
}

func NewOpenMessageByUserIdResultFromJson(data string) OpenMessageByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewOpenMessageByUserIdResultFromDict(dict)
}

func NewOpenMessageByUserIdResultFromDict(data map[string]interface{}) OpenMessageByUserIdResult {
	return OpenMessageByUserIdResult{
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p OpenMessageByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p OpenMessageByUserIdResult) Pointer() *OpenMessageByUserIdResult {
	return &p
}

type ReadMessageResult struct {
	Item                      *Message `json:"item"`
	TransactionId             *string  `json:"transactionId"`
	StampSheet                *string  `json:"stampSheet"`
	StampSheetEncryptionKeyId *string  `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool    `json:"autoRunStampSheet"`
}

type ReadMessageAsyncResult struct {
	result *ReadMessageResult
	err    error
}

func NewReadMessageResultFromJson(data string) ReadMessageResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReadMessageResultFromDict(dict)
}

func NewReadMessageResultFromDict(data map[string]interface{}) ReadMessageResult {
	return ReadMessageResult{
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
	}
}

func (p ReadMessageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ReadMessageResult) Pointer() *ReadMessageResult {
	return &p
}

type ReadMessageByUserIdResult struct {
	Item                      *Message `json:"item"`
	TransactionId             *string  `json:"transactionId"`
	StampSheet                *string  `json:"stampSheet"`
	StampSheetEncryptionKeyId *string  `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool    `json:"autoRunStampSheet"`
}

type ReadMessageByUserIdAsyncResult struct {
	result *ReadMessageByUserIdResult
	err    error
}

func NewReadMessageByUserIdResultFromJson(data string) ReadMessageByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReadMessageByUserIdResultFromDict(dict)
}

func NewReadMessageByUserIdResultFromDict(data map[string]interface{}) ReadMessageByUserIdResult {
	return ReadMessageByUserIdResult{
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
	}
}

func (p ReadMessageByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ReadMessageByUserIdResult) Pointer() *ReadMessageByUserIdResult {
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
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteMessageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteMessageResult) Pointer() *DeleteMessageResult {
	return &p
}

type DeleteMessageByUserIdResult struct {
	Item *Message `json:"item"`
}

type DeleteMessageByUserIdAsyncResult struct {
	result *DeleteMessageByUserIdResult
	err    error
}

func NewDeleteMessageByUserIdResultFromJson(data string) DeleteMessageByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMessageByUserIdResultFromDict(dict)
}

func NewDeleteMessageByUserIdResultFromDict(data map[string]interface{}) DeleteMessageByUserIdResult {
	return DeleteMessageByUserIdResult{
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteMessageByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteMessageByUserIdResult) Pointer() *DeleteMessageByUserIdResult {
	return &p
}

type SendByStampSheetResult struct {
	Item *Message `json:"item"`
}

type SendByStampSheetAsyncResult struct {
	result *SendByStampSheetResult
	err    error
}

func NewSendByStampSheetResultFromJson(data string) SendByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendByStampSheetResultFromDict(dict)
}

func NewSendByStampSheetResultFromDict(data map[string]interface{}) SendByStampSheetResult {
	return SendByStampSheetResult{
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SendByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SendByStampSheetResult) Pointer() *SendByStampSheetResult {
	return &p
}

type OpenByStampTaskResult struct {
	Item            *Message `json:"item"`
	NewContextStack *string  `json:"newContextStack"`
}

type OpenByStampTaskAsyncResult struct {
	result *OpenByStampTaskResult
	err    error
}

func NewOpenByStampTaskResultFromJson(data string) OpenByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewOpenByStampTaskResultFromDict(dict)
}

func NewOpenByStampTaskResultFromDict(data map[string]interface{}) OpenByStampTaskResult {
	return OpenByStampTaskResult{
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p OpenByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
	}
}

func (p OpenByStampTaskResult) Pointer() *OpenByStampTaskResult {
	return &p
}

type DeleteMessageByStampTaskResult struct {
	Item            *Message `json:"item"`
	NewContextStack *string  `json:"newContextStack"`
}

type DeleteMessageByStampTaskAsyncResult struct {
	result *DeleteMessageByStampTaskResult
	err    error
}

func NewDeleteMessageByStampTaskResultFromJson(data string) DeleteMessageByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMessageByStampTaskResultFromDict(dict)
}

func NewDeleteMessageByStampTaskResultFromDict(data map[string]interface{}) DeleteMessageByStampTaskResult {
	return DeleteMessageByStampTaskResult{
		Item: func() *Message {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p DeleteMessageByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
	}
}

func (p DeleteMessageByStampTaskResult) Pointer() *DeleteMessageByStampTaskResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentMessageMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromJson(data string) ExportMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExportMasterResultFromDict(dict)
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
	return ExportMasterResult{
		Item: func() *CurrentMessageMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentMessageMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentMessageMasterResult struct {
	Item *CurrentMessageMaster `json:"item"`
}

type GetCurrentMessageMasterAsyncResult struct {
	result *GetCurrentMessageMasterResult
	err    error
}

func NewGetCurrentMessageMasterResultFromJson(data string) GetCurrentMessageMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentMessageMasterResultFromDict(dict)
}

func NewGetCurrentMessageMasterResultFromDict(data map[string]interface{}) GetCurrentMessageMasterResult {
	return GetCurrentMessageMasterResult{
		Item: func() *CurrentMessageMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentMessageMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCurrentMessageMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentMessageMasterResult) Pointer() *GetCurrentMessageMasterResult {
	return &p
}

type UpdateCurrentMessageMasterResult struct {
	Item *CurrentMessageMaster `json:"item"`
}

type UpdateCurrentMessageMasterAsyncResult struct {
	result *UpdateCurrentMessageMasterResult
	err    error
}

func NewUpdateCurrentMessageMasterResultFromJson(data string) UpdateCurrentMessageMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentMessageMasterResultFromDict(dict)
}

func NewUpdateCurrentMessageMasterResultFromDict(data map[string]interface{}) UpdateCurrentMessageMasterResult {
	return UpdateCurrentMessageMasterResult{
		Item: func() *CurrentMessageMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentMessageMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentMessageMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentMessageMasterResult) Pointer() *UpdateCurrentMessageMasterResult {
	return &p
}

type UpdateCurrentMessageMasterFromGitHubResult struct {
	Item *CurrentMessageMaster `json:"item"`
}

type UpdateCurrentMessageMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentMessageMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentMessageMasterFromGitHubResultFromJson(data string) UpdateCurrentMessageMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentMessageMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentMessageMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentMessageMasterFromGitHubResult {
	return UpdateCurrentMessageMasterFromGitHubResult{
		Item: func() *CurrentMessageMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentMessageMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentMessageMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentMessageMasterFromGitHubResult) Pointer() *UpdateCurrentMessageMasterFromGitHubResult {
	return &p
}

type DescribeGlobalMessageMastersResult struct {
	Items         []GlobalMessageMaster `json:"items"`
	NextPageToken *string               `json:"nextPageToken"`
}

type DescribeGlobalMessageMastersAsyncResult struct {
	result *DescribeGlobalMessageMastersResult
	err    error
}

func NewDescribeGlobalMessageMastersResultFromJson(data string) DescribeGlobalMessageMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalMessageMastersResultFromDict(dict)
}

func NewDescribeGlobalMessageMastersResultFromDict(data map[string]interface{}) DescribeGlobalMessageMastersResult {
	return DescribeGlobalMessageMastersResult{
		Items: func() []GlobalMessageMaster {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalMessageMasters(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeGlobalMessageMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalMessageMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGlobalMessageMastersResult) Pointer() *DescribeGlobalMessageMastersResult {
	return &p
}

type CreateGlobalMessageMasterResult struct {
	Item *GlobalMessageMaster `json:"item"`
}

type CreateGlobalMessageMasterAsyncResult struct {
	result *CreateGlobalMessageMasterResult
	err    error
}

func NewCreateGlobalMessageMasterResultFromJson(data string) CreateGlobalMessageMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGlobalMessageMasterResultFromDict(dict)
}

func NewCreateGlobalMessageMasterResultFromDict(data map[string]interface{}) CreateGlobalMessageMasterResult {
	return CreateGlobalMessageMasterResult{
		Item: func() *GlobalMessageMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalMessageMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateGlobalMessageMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateGlobalMessageMasterResult) Pointer() *CreateGlobalMessageMasterResult {
	return &p
}

type GetGlobalMessageMasterResult struct {
	Item *GlobalMessageMaster `json:"item"`
}

type GetGlobalMessageMasterAsyncResult struct {
	result *GetGlobalMessageMasterResult
	err    error
}

func NewGetGlobalMessageMasterResultFromJson(data string) GetGlobalMessageMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalMessageMasterResultFromDict(dict)
}

func NewGetGlobalMessageMasterResultFromDict(data map[string]interface{}) GetGlobalMessageMasterResult {
	return GetGlobalMessageMasterResult{
		Item: func() *GlobalMessageMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalMessageMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalMessageMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalMessageMasterResult) Pointer() *GetGlobalMessageMasterResult {
	return &p
}

type UpdateGlobalMessageMasterResult struct {
	Item *GlobalMessageMaster `json:"item"`
}

type UpdateGlobalMessageMasterAsyncResult struct {
	result *UpdateGlobalMessageMasterResult
	err    error
}

func NewUpdateGlobalMessageMasterResultFromJson(data string) UpdateGlobalMessageMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGlobalMessageMasterResultFromDict(dict)
}

func NewUpdateGlobalMessageMasterResultFromDict(data map[string]interface{}) UpdateGlobalMessageMasterResult {
	return UpdateGlobalMessageMasterResult{
		Item: func() *GlobalMessageMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalMessageMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateGlobalMessageMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateGlobalMessageMasterResult) Pointer() *UpdateGlobalMessageMasterResult {
	return &p
}

type DeleteGlobalMessageMasterResult struct {
	Item *GlobalMessageMaster `json:"item"`
}

type DeleteGlobalMessageMasterAsyncResult struct {
	result *DeleteGlobalMessageMasterResult
	err    error
}

func NewDeleteGlobalMessageMasterResultFromJson(data string) DeleteGlobalMessageMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGlobalMessageMasterResultFromDict(dict)
}

func NewDeleteGlobalMessageMasterResultFromDict(data map[string]interface{}) DeleteGlobalMessageMasterResult {
	return DeleteGlobalMessageMasterResult{
		Item: func() *GlobalMessageMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalMessageMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteGlobalMessageMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteGlobalMessageMasterResult) Pointer() *DeleteGlobalMessageMasterResult {
	return &p
}

type DescribeGlobalMessagesResult struct {
	Items []GlobalMessage `json:"items"`
}

type DescribeGlobalMessagesAsyncResult struct {
	result *DescribeGlobalMessagesResult
	err    error
}

func NewDescribeGlobalMessagesResultFromJson(data string) DescribeGlobalMessagesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalMessagesResultFromDict(dict)
}

func NewDescribeGlobalMessagesResultFromDict(data map[string]interface{}) DescribeGlobalMessagesResult {
	return DescribeGlobalMessagesResult{
		Items: func() []GlobalMessage {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalMessages(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeGlobalMessagesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalMessagesFromDict(
			p.Items,
		),
	}
}

func (p DescribeGlobalMessagesResult) Pointer() *DescribeGlobalMessagesResult {
	return &p
}

type GetGlobalMessageResult struct {
	Item *GlobalMessage `json:"item"`
}

type GetGlobalMessageAsyncResult struct {
	result *GetGlobalMessageResult
	err    error
}

func NewGetGlobalMessageResultFromJson(data string) GetGlobalMessageResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalMessageResultFromDict(dict)
}

func NewGetGlobalMessageResultFromDict(data map[string]interface{}) GetGlobalMessageResult {
	return GetGlobalMessageResult{
		Item: func() *GlobalMessage {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalMessageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalMessageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalMessageResult) Pointer() *GetGlobalMessageResult {
	return &p
}

type GetReceivedByUserIdResult struct {
	Item *Received `json:"item"`
}

type GetReceivedByUserIdAsyncResult struct {
	result *GetReceivedByUserIdResult
	err    error
}

func NewGetReceivedByUserIdResultFromJson(data string) GetReceivedByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceivedByUserIdResultFromDict(dict)
}

func NewGetReceivedByUserIdResultFromDict(data map[string]interface{}) GetReceivedByUserIdResult {
	return GetReceivedByUserIdResult{
		Item: func() *Received {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceivedFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetReceivedByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetReceivedByUserIdResult) Pointer() *GetReceivedByUserIdResult {
	return &p
}

type UpdateReceivedByUserIdResult struct {
	Item *Received `json:"item"`
}

type UpdateReceivedByUserIdAsyncResult struct {
	result *UpdateReceivedByUserIdResult
	err    error
}

func NewUpdateReceivedByUserIdResultFromJson(data string) UpdateReceivedByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateReceivedByUserIdResultFromDict(dict)
}

func NewUpdateReceivedByUserIdResultFromDict(data map[string]interface{}) UpdateReceivedByUserIdResult {
	return UpdateReceivedByUserIdResult{
		Item: func() *Received {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceivedFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateReceivedByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateReceivedByUserIdResult) Pointer() *UpdateReceivedByUserIdResult {
	return &p
}

type DeleteReceivedByUserIdResult struct {
	Item *Received `json:"item"`
}

type DeleteReceivedByUserIdAsyncResult struct {
	result *DeleteReceivedByUserIdResult
	err    error
}

func NewDeleteReceivedByUserIdResultFromJson(data string) DeleteReceivedByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReceivedByUserIdResultFromDict(dict)
}

func NewDeleteReceivedByUserIdResultFromDict(data map[string]interface{}) DeleteReceivedByUserIdResult {
	return DeleteReceivedByUserIdResult{
		Item: func() *Received {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceivedFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteReceivedByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteReceivedByUserIdResult) Pointer() *DeleteReceivedByUserIdResult {
	return &p
}
