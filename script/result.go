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

package script

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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNamespacesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"status": p.Status,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
	return &p
}

type DescribeScriptsResult struct {
	Items         []Script             `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeScriptsAsyncResult struct {
	result *DescribeScriptsResult
	err    error
}

func NewDescribeScriptsResultFromJson(data string) DescribeScriptsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeScriptsResultFromDict(dict)
}

func NewDescribeScriptsResultFromDict(data map[string]interface{}) DescribeScriptsResult {
	return DescribeScriptsResult{
		Items: func() []Script {
			if data["items"] == nil {
				return nil
			}
			return CastScripts(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeScriptsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastScriptsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeScriptsResult) Pointer() *DescribeScriptsResult {
	return &p
}

type CreateScriptResult struct {
	Item     *Script              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateScriptAsyncResult struct {
	result *CreateScriptResult
	err    error
}

func NewCreateScriptResultFromJson(data string) CreateScriptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateScriptResultFromDict(dict)
}

func NewCreateScriptResultFromDict(data map[string]interface{}) CreateScriptResult {
	return CreateScriptResult{
		Item: func() *Script {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateScriptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateScriptResult) Pointer() *CreateScriptResult {
	return &p
}

type CreateScriptFromGitHubResult struct {
	Item     *Script              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateScriptFromGitHubAsyncResult struct {
	result *CreateScriptFromGitHubResult
	err    error
}

func NewCreateScriptFromGitHubResultFromJson(data string) CreateScriptFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateScriptFromGitHubResultFromDict(dict)
}

func NewCreateScriptFromGitHubResultFromDict(data map[string]interface{}) CreateScriptFromGitHubResult {
	return CreateScriptFromGitHubResult{
		Item: func() *Script {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateScriptFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateScriptFromGitHubResult) Pointer() *CreateScriptFromGitHubResult {
	return &p
}

type GetScriptResult struct {
	Item     *Script              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetScriptAsyncResult struct {
	result *GetScriptResult
	err    error
}

func NewGetScriptResultFromJson(data string) GetScriptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetScriptResultFromDict(dict)
}

func NewGetScriptResultFromDict(data map[string]interface{}) GetScriptResult {
	return GetScriptResult{
		Item: func() *Script {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetScriptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetScriptResult) Pointer() *GetScriptResult {
	return &p
}

type UpdateScriptResult struct {
	Item     *Script              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateScriptAsyncResult struct {
	result *UpdateScriptResult
	err    error
}

func NewUpdateScriptResultFromJson(data string) UpdateScriptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateScriptResultFromDict(dict)
}

func NewUpdateScriptResultFromDict(data map[string]interface{}) UpdateScriptResult {
	return UpdateScriptResult{
		Item: func() *Script {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateScriptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateScriptResult) Pointer() *UpdateScriptResult {
	return &p
}

type UpdateScriptFromGitHubResult struct {
	Item     *Script              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateScriptFromGitHubAsyncResult struct {
	result *UpdateScriptFromGitHubResult
	err    error
}

func NewUpdateScriptFromGitHubResultFromJson(data string) UpdateScriptFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateScriptFromGitHubResultFromDict(dict)
}

func NewUpdateScriptFromGitHubResultFromDict(data map[string]interface{}) UpdateScriptFromGitHubResult {
	return UpdateScriptFromGitHubResult{
		Item: func() *Script {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateScriptFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateScriptFromGitHubResult) Pointer() *UpdateScriptFromGitHubResult {
	return &p
}

type DeleteScriptResult struct {
	Item     *Script              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteScriptAsyncResult struct {
	result *DeleteScriptResult
	err    error
}

func NewDeleteScriptResultFromJson(data string) DeleteScriptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteScriptResultFromDict(dict)
}

func NewDeleteScriptResultFromDict(data map[string]interface{}) DeleteScriptResult {
	return DeleteScriptResult{
		Item: func() *Script {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteScriptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteScriptResult) Pointer() *DeleteScriptResult {
	return &p
}

type InvokeScriptResult struct {
	Code              *int32                  `json:"code"`
	Result            *string                 `json:"result"`
	Transaction       *Transaction            `json:"transaction"`
	RandomStatus      *RandomStatus           `json:"randomStatus"`
	AtomicCommit      *bool                   `json:"atomicCommit"`
	TransactionResult *core.TransactionResult `json:"transactionResult"`
	ExecuteTime       *int32                  `json:"executeTime"`
	Charged           *int32                  `json:"charged"`
	Output            []*string               `json:"output"`
	Metadata          *core.ResultMetadata    `json:"metadata"`
}

type InvokeScriptAsyncResult struct {
	result *InvokeScriptResult
	err    error
}

func NewInvokeScriptResultFromJson(data string) InvokeScriptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInvokeScriptResultFromDict(dict)
}

func NewInvokeScriptResultFromDict(data map[string]interface{}) InvokeScriptResult {
	return InvokeScriptResult{
		Code: func() *int32 {
			v, ok := data["code"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["code"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
		Transaction: func() *Transaction {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionFromDict(core.CastMap(data["transaction"])).Pointer()
		}(),
		RandomStatus: func() *RandomStatus {
			v, ok := data["randomStatus"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomStatusFromDict(core.CastMap(data["randomStatus"])).Pointer()
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		ExecuteTime: func() *int32 {
			v, ok := data["executeTime"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["executeTime"])
		}(),
		Charged: func() *int32 {
			v, ok := data["charged"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["charged"])
		}(),
		Output: func() []*string {
			v, ok := data["output"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p InvokeScriptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"code":   p.Code,
		"result": p.Result,
		"transaction": func() map[string]interface{} {
			if p.Transaction == nil {
				return nil
			}
			return p.Transaction.ToDict()
		}(),
		"randomStatus": func() map[string]interface{} {
			if p.RandomStatus == nil {
				return nil
			}
			return p.RandomStatus.ToDict()
		}(),
		"atomicCommit": p.AtomicCommit,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"executeTime": p.ExecuteTime,
		"charged":     p.Charged,
		"output": core.CastStringsFromDict(
			p.Output,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p InvokeScriptResult) Pointer() *InvokeScriptResult {
	return &p
}

type DebugInvokeResult struct {
	Code              *int32                  `json:"code"`
	Result            *string                 `json:"result"`
	Transaction       *Transaction            `json:"transaction"`
	RandomStatus      *RandomStatus           `json:"randomStatus"`
	AtomicCommit      *bool                   `json:"atomicCommit"`
	TransactionResult *core.TransactionResult `json:"transactionResult"`
	ExecuteTime       *int32                  `json:"executeTime"`
	Charged           *int32                  `json:"charged"`
	Output            []*string               `json:"output"`
	Metadata          *core.ResultMetadata    `json:"metadata"`
}

type DebugInvokeAsyncResult struct {
	result *DebugInvokeResult
	err    error
}

func NewDebugInvokeResultFromJson(data string) DebugInvokeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDebugInvokeResultFromDict(dict)
}

func NewDebugInvokeResultFromDict(data map[string]interface{}) DebugInvokeResult {
	return DebugInvokeResult{
		Code: func() *int32 {
			v, ok := data["code"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["code"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
		Transaction: func() *Transaction {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionFromDict(core.CastMap(data["transaction"])).Pointer()
		}(),
		RandomStatus: func() *RandomStatus {
			v, ok := data["randomStatus"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomStatusFromDict(core.CastMap(data["randomStatus"])).Pointer()
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		ExecuteTime: func() *int32 {
			v, ok := data["executeTime"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["executeTime"])
		}(),
		Charged: func() *int32 {
			v, ok := data["charged"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["charged"])
		}(),
		Output: func() []*string {
			v, ok := data["output"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DebugInvokeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"code":   p.Code,
		"result": p.Result,
		"transaction": func() map[string]interface{} {
			if p.Transaction == nil {
				return nil
			}
			return p.Transaction.ToDict()
		}(),
		"randomStatus": func() map[string]interface{} {
			if p.RandomStatus == nil {
				return nil
			}
			return p.RandomStatus.ToDict()
		}(),
		"atomicCommit": p.AtomicCommit,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"executeTime": p.ExecuteTime,
		"charged":     p.Charged,
		"output": core.CastStringsFromDict(
			p.Output,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DebugInvokeResult) Pointer() *DebugInvokeResult {
	return &p
}

type InvokeByStampSheetResult struct {
	Code              *int32                  `json:"code"`
	Result            *string                 `json:"result"`
	Transaction       *Transaction            `json:"transaction"`
	RandomStatus      *RandomStatus           `json:"randomStatus"`
	AtomicCommit      *bool                   `json:"atomicCommit"`
	TransactionResult *core.TransactionResult `json:"transactionResult"`
	ExecuteTime       *int32                  `json:"executeTime"`
	Charged           *int32                  `json:"charged"`
	Output            []*string               `json:"output"`
	Metadata          *core.ResultMetadata    `json:"metadata"`
}

type InvokeByStampSheetAsyncResult struct {
	result *InvokeByStampSheetResult
	err    error
}

func NewInvokeByStampSheetResultFromJson(data string) InvokeByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInvokeByStampSheetResultFromDict(dict)
}

func NewInvokeByStampSheetResultFromDict(data map[string]interface{}) InvokeByStampSheetResult {
	return InvokeByStampSheetResult{
		Code: func() *int32 {
			v, ok := data["code"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["code"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
		Transaction: func() *Transaction {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionFromDict(core.CastMap(data["transaction"])).Pointer()
		}(),
		RandomStatus: func() *RandomStatus {
			v, ok := data["randomStatus"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomStatusFromDict(core.CastMap(data["randomStatus"])).Pointer()
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		ExecuteTime: func() *int32 {
			v, ok := data["executeTime"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["executeTime"])
		}(),
		Charged: func() *int32 {
			v, ok := data["charged"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["charged"])
		}(),
		Output: func() []*string {
			v, ok := data["output"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p InvokeByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"code":   p.Code,
		"result": p.Result,
		"transaction": func() map[string]interface{} {
			if p.Transaction == nil {
				return nil
			}
			return p.Transaction.ToDict()
		}(),
		"randomStatus": func() map[string]interface{} {
			if p.RandomStatus == nil {
				return nil
			}
			return p.RandomStatus.ToDict()
		}(),
		"atomicCommit": p.AtomicCommit,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"executeTime": p.ExecuteTime,
		"charged":     p.Charged,
		"output": core.CastStringsFromDict(
			p.Output,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p InvokeByStampSheetResult) Pointer() *InvokeByStampSheetResult {
	return &p
}
