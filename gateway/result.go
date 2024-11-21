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

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
	Items         []Namespace          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Status   *string              `json:"status"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Url      *string              `json:"url"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Url      *string              `json:"url"`
	Metadata *core.ResultMetadata `json:"metadata"`
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

type DescribeWebSocketSessionsResult struct {
	Items         []WebSocketSession   `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeWebSocketSessionsAsyncResult struct {
	result *DescribeWebSocketSessionsResult
	err    error
}

func NewDescribeWebSocketSessionsResultFromJson(data string) DescribeWebSocketSessionsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeWebSocketSessionsResultFromDict(dict)
}

func NewDescribeWebSocketSessionsResultFromDict(data map[string]interface{}) DescribeWebSocketSessionsResult {
	return DescribeWebSocketSessionsResult{
		Items: func() []WebSocketSession {
			if data["items"] == nil {
				return nil
			}
			return CastWebSocketSessions(core.CastArray(data["items"]))
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

func (p DescribeWebSocketSessionsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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
	Items         []WebSocketSession   `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeWebSocketSessionsByUserIdAsyncResult struct {
	result *DescribeWebSocketSessionsByUserIdResult
	err    error
}

func NewDescribeWebSocketSessionsByUserIdResultFromJson(data string) DescribeWebSocketSessionsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeWebSocketSessionsByUserIdResultFromDict(dict)
}

func NewDescribeWebSocketSessionsByUserIdResultFromDict(data map[string]interface{}) DescribeWebSocketSessionsByUserIdResult {
	return DescribeWebSocketSessionsByUserIdResult{
		Items: func() []WebSocketSession {
			if data["items"] == nil {
				return nil
			}
			return CastWebSocketSessions(core.CastArray(data["items"]))
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

func (p DescribeWebSocketSessionsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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
	Item     *WebSocketSession    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetUserIdAsyncResult struct {
	result *SetUserIdResult
	err    error
}

func NewSetUserIdResultFromJson(data string) SetUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetUserIdResultFromDict(dict)
}

func NewSetUserIdResultFromDict(data map[string]interface{}) SetUserIdResult {
	return SetUserIdResult{
		Item: func() *WebSocketSession {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewWebSocketSessionFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SetUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SetUserIdResult) Pointer() *SetUserIdResult {
	return &p
}

type SetUserIdByUserIdResult struct {
	Item     *WebSocketSession    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetUserIdByUserIdAsyncResult struct {
	result *SetUserIdByUserIdResult
	err    error
}

func NewSetUserIdByUserIdResultFromJson(data string) SetUserIdByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetUserIdByUserIdResultFromDict(dict)
}

func NewSetUserIdByUserIdResultFromDict(data map[string]interface{}) SetUserIdByUserIdResult {
	return SetUserIdByUserIdResult{
		Item: func() *WebSocketSession {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewWebSocketSessionFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SetUserIdByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SetUserIdByUserIdResult) Pointer() *SetUserIdByUserIdResult {
	return &p
}

type SendNotificationResult struct {
	Protocol          *string              `json:"protocol"`
	SendConnectionIds []*string            `json:"sendConnectionIds"`
	Metadata          *core.ResultMetadata `json:"metadata"`
}

type SendNotificationAsyncResult struct {
	result *SendNotificationResult
	err    error
}

func NewSendNotificationResultFromJson(data string) SendNotificationResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendNotificationResultFromDict(dict)
}

func NewSendNotificationResultFromDict(data map[string]interface{}) SendNotificationResult {
	return SendNotificationResult{
		Protocol: func() *string {
			v, ok := data["protocol"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["protocol"])
		}(),
		SendConnectionIds: func() []*string {
			v, ok := data["sendConnectionIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
	}
}

func (p SendNotificationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"protocol": p.Protocol,
		"sendConnectionIds": core.CastStringsFromDict(
			p.SendConnectionIds,
		),
	}
}

func (p SendNotificationResult) Pointer() *SendNotificationResult {
	return &p
}

type DisconnectByUserIdResult struct {
	Items    []WebSocketSession   `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DisconnectByUserIdAsyncResult struct {
	result *DisconnectByUserIdResult
	err    error
}

func NewDisconnectByUserIdResultFromJson(data string) DisconnectByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDisconnectByUserIdResultFromDict(dict)
}

func NewDisconnectByUserIdResultFromDict(data map[string]interface{}) DisconnectByUserIdResult {
	return DisconnectByUserIdResult{
		Items: func() []WebSocketSession {
			if data["items"] == nil {
				return nil
			}
			return CastWebSocketSessions(core.CastArray(data["items"]))
		}(),
	}
}

func (p DisconnectByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastWebSocketSessionsFromDict(
			p.Items,
		),
	}
}

func (p DisconnectByUserIdResult) Pointer() *DisconnectByUserIdResult {
	return &p
}

type DisconnectAllResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DisconnectAllAsyncResult struct {
	result *DisconnectAllResult
	err    error
}

func NewDisconnectAllResultFromJson(data string) DisconnectAllResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDisconnectAllResultFromDict(dict)
}

func NewDisconnectAllResultFromDict(data map[string]interface{}) DisconnectAllResult {
	return DisconnectAllResult{}
}

func (p DisconnectAllResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DisconnectAllResult) Pointer() *DisconnectAllResult {
	return &p
}

type SetFirebaseTokenResult struct {
	Item     *FirebaseToken       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetFirebaseTokenAsyncResult struct {
	result *SetFirebaseTokenResult
	err    error
}

func NewSetFirebaseTokenResultFromJson(data string) SetFirebaseTokenResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetFirebaseTokenResultFromDict(dict)
}

func NewSetFirebaseTokenResultFromDict(data map[string]interface{}) SetFirebaseTokenResult {
	return SetFirebaseTokenResult{
		Item: func() *FirebaseToken {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SetFirebaseTokenResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SetFirebaseTokenResult) Pointer() *SetFirebaseTokenResult {
	return &p
}

type SetFirebaseTokenByUserIdResult struct {
	Item     *FirebaseToken       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetFirebaseTokenByUserIdAsyncResult struct {
	result *SetFirebaseTokenByUserIdResult
	err    error
}

func NewSetFirebaseTokenByUserIdResultFromJson(data string) SetFirebaseTokenByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetFirebaseTokenByUserIdResultFromDict(dict)
}

func NewSetFirebaseTokenByUserIdResultFromDict(data map[string]interface{}) SetFirebaseTokenByUserIdResult {
	return SetFirebaseTokenByUserIdResult{
		Item: func() *FirebaseToken {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SetFirebaseTokenByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SetFirebaseTokenByUserIdResult) Pointer() *SetFirebaseTokenByUserIdResult {
	return &p
}

type GetFirebaseTokenResult struct {
	Item     *FirebaseToken       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetFirebaseTokenAsyncResult struct {
	result *GetFirebaseTokenResult
	err    error
}

func NewGetFirebaseTokenResultFromJson(data string) GetFirebaseTokenResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFirebaseTokenResultFromDict(dict)
}

func NewGetFirebaseTokenResultFromDict(data map[string]interface{}) GetFirebaseTokenResult {
	return GetFirebaseTokenResult{
		Item: func() *FirebaseToken {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetFirebaseTokenResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetFirebaseTokenResult) Pointer() *GetFirebaseTokenResult {
	return &p
}

type GetFirebaseTokenByUserIdResult struct {
	Item     *FirebaseToken       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetFirebaseTokenByUserIdAsyncResult struct {
	result *GetFirebaseTokenByUserIdResult
	err    error
}

func NewGetFirebaseTokenByUserIdResultFromJson(data string) GetFirebaseTokenByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFirebaseTokenByUserIdResultFromDict(dict)
}

func NewGetFirebaseTokenByUserIdResultFromDict(data map[string]interface{}) GetFirebaseTokenByUserIdResult {
	return GetFirebaseTokenByUserIdResult{
		Item: func() *FirebaseToken {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetFirebaseTokenByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetFirebaseTokenByUserIdResult) Pointer() *GetFirebaseTokenByUserIdResult {
	return &p
}

type DeleteFirebaseTokenResult struct {
	Item     *FirebaseToken       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteFirebaseTokenAsyncResult struct {
	result *DeleteFirebaseTokenResult
	err    error
}

func NewDeleteFirebaseTokenResultFromJson(data string) DeleteFirebaseTokenResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFirebaseTokenResultFromDict(dict)
}

func NewDeleteFirebaseTokenResultFromDict(data map[string]interface{}) DeleteFirebaseTokenResult {
	return DeleteFirebaseTokenResult{
		Item: func() *FirebaseToken {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteFirebaseTokenResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteFirebaseTokenResult) Pointer() *DeleteFirebaseTokenResult {
	return &p
}

type DeleteFirebaseTokenByUserIdResult struct {
	Item     *FirebaseToken       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteFirebaseTokenByUserIdAsyncResult struct {
	result *DeleteFirebaseTokenByUserIdResult
	err    error
}

func NewDeleteFirebaseTokenByUserIdResultFromJson(data string) DeleteFirebaseTokenByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFirebaseTokenByUserIdResultFromDict(dict)
}

func NewDeleteFirebaseTokenByUserIdResultFromDict(data map[string]interface{}) DeleteFirebaseTokenByUserIdResult {
	return DeleteFirebaseTokenByUserIdResult{
		Item: func() *FirebaseToken {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFirebaseTokenFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteFirebaseTokenByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteFirebaseTokenByUserIdResult) Pointer() *DeleteFirebaseTokenByUserIdResult {
	return &p
}

type SendMobileNotificationByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SendMobileNotificationByUserIdAsyncResult struct {
	result *SendMobileNotificationByUserIdResult
	err    error
}

func NewSendMobileNotificationByUserIdResultFromJson(data string) SendMobileNotificationByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendMobileNotificationByUserIdResultFromDict(dict)
}

func NewSendMobileNotificationByUserIdResultFromDict(data map[string]interface{}) SendMobileNotificationByUserIdResult {
	return SendMobileNotificationByUserIdResult{}
}

func (p SendMobileNotificationByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p SendMobileNotificationByUserIdResult) Pointer() *SendMobileNotificationByUserIdResult {
	return &p
}
