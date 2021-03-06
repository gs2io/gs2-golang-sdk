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

type DescribeMessagesResult struct {
    Items []Message `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeMessagesAsyncResult struct {
	result *DescribeMessagesResult
	err    error
}

func NewDescribeMessagesResultFromDict(data map[string]interface{}) DescribeMessagesResult {
    return DescribeMessagesResult {
        Items: CastMessages(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeMessagesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    Items []Message `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeMessagesByUserIdAsyncResult struct {
	result *DescribeMessagesByUserIdResult
	err    error
}

func NewDescribeMessagesByUserIdResultFromDict(data map[string]interface{}) DescribeMessagesByUserIdResult {
    return DescribeMessagesByUserIdResult {
        Items: CastMessages(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeMessagesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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

func NewSendMessageByUserIdResultFromDict(data map[string]interface{}) SendMessageByUserIdResult {
    return SendMessageByUserIdResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SendMessageByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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

type GetMessageByUserIdResult struct {
    Item *Message `json:"item"`
}

type GetMessageByUserIdAsyncResult struct {
	result *GetMessageByUserIdResult
	err    error
}

func NewGetMessageByUserIdResultFromDict(data map[string]interface{}) GetMessageByUserIdResult {
    return GetMessageByUserIdResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetMessageByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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

func NewReceiveGlobalMessageResultFromDict(data map[string]interface{}) ReceiveGlobalMessageResult {
    return ReceiveGlobalMessageResult {
        Item: CastMessages(core.CastArray(data["item"])),
    }
}

func (p ReceiveGlobalMessageResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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

func NewReceiveGlobalMessageByUserIdResultFromDict(data map[string]interface{}) ReceiveGlobalMessageByUserIdResult {
    return ReceiveGlobalMessageByUserIdResult {
        Item: CastMessages(core.CastArray(data["item"])),
    }
}

func (p ReceiveGlobalMessageByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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

func NewOpenMessageResultFromDict(data map[string]interface{}) OpenMessageResult {
    return OpenMessageResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p OpenMessageResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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

func NewOpenMessageByUserIdResultFromDict(data map[string]interface{}) OpenMessageByUserIdResult {
    return OpenMessageByUserIdResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p OpenMessageByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p OpenMessageByUserIdResult) Pointer() *OpenMessageByUserIdResult {
    return &p
}

type ReadMessageResult struct {
    Item *Message `json:"item"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type ReadMessageAsyncResult struct {
	result *ReadMessageResult
	err    error
}

func NewReadMessageResultFromDict(data map[string]interface{}) ReadMessageResult {
    return ReadMessageResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p ReadMessageResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p ReadMessageResult) Pointer() *ReadMessageResult {
    return &p
}

type ReadMessageByUserIdResult struct {
    Item *Message `json:"item"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type ReadMessageByUserIdAsyncResult struct {
	result *ReadMessageByUserIdResult
	err    error
}

func NewReadMessageByUserIdResultFromDict(data map[string]interface{}) ReadMessageByUserIdResult {
    return ReadMessageByUserIdResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p ReadMessageByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p ReadMessageByUserIdResult) Pointer() *ReadMessageByUserIdResult {
    return &p
}

type DeleteMessageResult struct {
}

type DeleteMessageAsyncResult struct {
	result *DeleteMessageResult
	err    error
}

func NewDeleteMessageResultFromDict(data map[string]interface{}) DeleteMessageResult {
    return DeleteMessageResult {
    }
}

func (p DeleteMessageResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteMessageResult) Pointer() *DeleteMessageResult {
    return &p
}

type DeleteMessageByUserIdResult struct {
}

type DeleteMessageByUserIdAsyncResult struct {
	result *DeleteMessageByUserIdResult
	err    error
}

func NewDeleteMessageByUserIdResultFromDict(data map[string]interface{}) DeleteMessageByUserIdResult {
    return DeleteMessageByUserIdResult {
    }
}

func (p DeleteMessageByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteMessageByUserIdResult) Pointer() *DeleteMessageByUserIdResult {
    return &p
}

type OpenByStampTaskResult struct {
    Item *Message `json:"item"`
    NewContextStack *string `json:"newContextStack"`
}

type OpenByStampTaskAsyncResult struct {
	result *OpenByStampTaskResult
	err    error
}

func NewOpenByStampTaskResultFromDict(data map[string]interface{}) OpenByStampTaskResult {
    return OpenByStampTaskResult {
        Item: NewMessageFromDict(core.CastMap(data["item"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p OpenByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p OpenByStampTaskResult) Pointer() *OpenByStampTaskResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentMessageMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentMessageMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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

func NewGetCurrentMessageMasterResultFromDict(data map[string]interface{}) GetCurrentMessageMasterResult {
    return GetCurrentMessageMasterResult {
        Item: NewCurrentMessageMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentMessageMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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

func NewUpdateCurrentMessageMasterResultFromDict(data map[string]interface{}) UpdateCurrentMessageMasterResult {
    return UpdateCurrentMessageMasterResult {
        Item: NewCurrentMessageMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentMessageMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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

func NewUpdateCurrentMessageMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentMessageMasterFromGitHubResult {
    return UpdateCurrentMessageMasterFromGitHubResult {
        Item: NewCurrentMessageMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentMessageMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentMessageMasterFromGitHubResult) Pointer() *UpdateCurrentMessageMasterFromGitHubResult {
    return &p
}

type DescribeGlobalMessageMastersResult struct {
    Items []GlobalMessageMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeGlobalMessageMastersAsyncResult struct {
	result *DescribeGlobalMessageMastersResult
	err    error
}

func NewDescribeGlobalMessageMastersResultFromDict(data map[string]interface{}) DescribeGlobalMessageMastersResult {
    return DescribeGlobalMessageMastersResult {
        Items: CastGlobalMessageMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeGlobalMessageMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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

func NewCreateGlobalMessageMasterResultFromDict(data map[string]interface{}) CreateGlobalMessageMasterResult {
    return CreateGlobalMessageMasterResult {
        Item: NewGlobalMessageMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateGlobalMessageMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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

func NewGetGlobalMessageMasterResultFromDict(data map[string]interface{}) GetGlobalMessageMasterResult {
    return GetGlobalMessageMasterResult {
        Item: NewGlobalMessageMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetGlobalMessageMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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

func NewUpdateGlobalMessageMasterResultFromDict(data map[string]interface{}) UpdateGlobalMessageMasterResult {
    return UpdateGlobalMessageMasterResult {
        Item: NewGlobalMessageMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateGlobalMessageMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateGlobalMessageMasterResult) Pointer() *UpdateGlobalMessageMasterResult {
    return &p
}

type DeleteGlobalMessageMasterResult struct {
}

type DeleteGlobalMessageMasterAsyncResult struct {
	result *DeleteGlobalMessageMasterResult
	err    error
}

func NewDeleteGlobalMessageMasterResultFromDict(data map[string]interface{}) DeleteGlobalMessageMasterResult {
    return DeleteGlobalMessageMasterResult {
    }
}

func (p DeleteGlobalMessageMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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

func NewDescribeGlobalMessagesResultFromDict(data map[string]interface{}) DescribeGlobalMessagesResult {
    return DescribeGlobalMessagesResult {
        Items: CastGlobalMessages(core.CastArray(data["items"])),
    }
}

func (p DescribeGlobalMessagesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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

func NewGetGlobalMessageResultFromDict(data map[string]interface{}) GetGlobalMessageResult {
    return GetGlobalMessageResult {
        Item: NewGlobalMessageFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetGlobalMessageResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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

func NewGetReceivedByUserIdResultFromDict(data map[string]interface{}) GetReceivedByUserIdResult {
    return GetReceivedByUserIdResult {
        Item: NewReceivedFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetReceivedByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetReceivedByUserIdResult) Pointer() *GetReceivedByUserIdResult {
    return &p
}

type UpdateReceivedByUserIdResult struct {
}

type UpdateReceivedByUserIdAsyncResult struct {
	result *UpdateReceivedByUserIdResult
	err    error
}

func NewUpdateReceivedByUserIdResultFromDict(data map[string]interface{}) UpdateReceivedByUserIdResult {
    return UpdateReceivedByUserIdResult {
    }
}

func (p UpdateReceivedByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p UpdateReceivedByUserIdResult) Pointer() *UpdateReceivedByUserIdResult {
    return &p
}

type DeleteReceivedByUserIdResult struct {
}

type DeleteReceivedByUserIdAsyncResult struct {
	result *DeleteReceivedByUserIdResult
	err    error
}

func NewDeleteReceivedByUserIdResultFromDict(data map[string]interface{}) DeleteReceivedByUserIdResult {
    return DeleteReceivedByUserIdResult {
    }
}

func (p DeleteReceivedByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteReceivedByUserIdResult) Pointer() *DeleteReceivedByUserIdResult {
    return &p
}