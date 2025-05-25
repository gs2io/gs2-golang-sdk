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

package news

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

type GetServiceVersionResult struct {
	Item     *string              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetServiceVersionAsyncResult struct {
	result *GetServiceVersionResult
	err    error
}

func NewGetServiceVersionResultFromJson(data string) GetServiceVersionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetServiceVersionResultFromDict(dict)
}

func NewGetServiceVersionResultFromDict(data map[string]interface{}) GetServiceVersionResult {
	return GetServiceVersionResult{
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
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

func (p GetServiceVersionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetServiceVersionResult) Pointer() *GetServiceVersionResult {
	return &p
}

type DescribeProgressesResult struct {
	Items         []Progress           `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeProgressesAsyncResult struct {
	result *DescribeProgressesResult
	err    error
}

func NewDescribeProgressesResultFromJson(data string) DescribeProgressesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeProgressesResultFromDict(dict)
}

func NewDescribeProgressesResultFromDict(data map[string]interface{}) DescribeProgressesResult {
	return DescribeProgressesResult{
		Items: func() []Progress {
			if data["items"] == nil {
				return nil
			}
			return CastProgresses(core.CastArray(data["items"]))
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

func (p DescribeProgressesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastProgressesFromDict(
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

func (p DescribeProgressesResult) Pointer() *DescribeProgressesResult {
	return &p
}

type GetProgressResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetProgressAsyncResult struct {
	result *GetProgressResult
	err    error
}

func NewGetProgressResultFromJson(data string) GetProgressResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProgressResultFromDict(dict)
}

func NewGetProgressResultFromDict(data map[string]interface{}) GetProgressResult {
	return GetProgressResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetProgressResult) ToDict() map[string]interface{} {
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

func (p GetProgressResult) Pointer() *GetProgressResult {
	return &p
}

type DescribeOutputsResult struct {
	Items         []Output             `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeOutputsAsyncResult struct {
	result *DescribeOutputsResult
	err    error
}

func NewDescribeOutputsResultFromJson(data string) DescribeOutputsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeOutputsResultFromDict(dict)
}

func NewDescribeOutputsResultFromDict(data map[string]interface{}) DescribeOutputsResult {
	return DescribeOutputsResult{
		Items: func() []Output {
			if data["items"] == nil {
				return nil
			}
			return CastOutputs(core.CastArray(data["items"]))
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

func (p DescribeOutputsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastOutputsFromDict(
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

func (p DescribeOutputsResult) Pointer() *DescribeOutputsResult {
	return &p
}

type GetOutputResult struct {
	Item     *Output              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetOutputAsyncResult struct {
	result *GetOutputResult
	err    error
}

func NewGetOutputResultFromJson(data string) GetOutputResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetOutputResultFromDict(dict)
}

func NewGetOutputResultFromDict(data map[string]interface{}) GetOutputResult {
	return GetOutputResult{
		Item: func() *Output {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewOutputFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetOutputResult) ToDict() map[string]interface{} {
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

func (p GetOutputResult) Pointer() *GetOutputResult {
	return &p
}

type PrepareUpdateCurrentNewsMasterResult struct {
	UploadToken       *string              `json:"uploadToken"`
	TemplateUploadUrl *string              `json:"templateUploadUrl"`
	Metadata          *core.ResultMetadata `json:"metadata"`
}

type PrepareUpdateCurrentNewsMasterAsyncResult struct {
	result *PrepareUpdateCurrentNewsMasterResult
	err    error
}

func NewPrepareUpdateCurrentNewsMasterResultFromJson(data string) PrepareUpdateCurrentNewsMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareUpdateCurrentNewsMasterResultFromDict(dict)
}

func NewPrepareUpdateCurrentNewsMasterResultFromDict(data map[string]interface{}) PrepareUpdateCurrentNewsMasterResult {
	return PrepareUpdateCurrentNewsMasterResult{
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
		TemplateUploadUrl: func() *string {
			v, ok := data["templateUploadUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["templateUploadUrl"])
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

func (p PrepareUpdateCurrentNewsMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken":       p.UploadToken,
		"templateUploadUrl": p.TemplateUploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareUpdateCurrentNewsMasterResult) Pointer() *PrepareUpdateCurrentNewsMasterResult {
	return &p
}

type UpdateCurrentNewsMasterResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentNewsMasterAsyncResult struct {
	result *UpdateCurrentNewsMasterResult
	err    error
}

func NewUpdateCurrentNewsMasterResultFromJson(data string) UpdateCurrentNewsMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentNewsMasterResultFromDict(dict)
}

func NewUpdateCurrentNewsMasterResultFromDict(data map[string]interface{}) UpdateCurrentNewsMasterResult {
	return UpdateCurrentNewsMasterResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentNewsMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateCurrentNewsMasterResult) Pointer() *UpdateCurrentNewsMasterResult {
	return &p
}

type PrepareUpdateCurrentNewsMasterFromGitHubResult struct {
	UploadToken *string              `json:"uploadToken"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PrepareUpdateCurrentNewsMasterFromGitHubAsyncResult struct {
	result *PrepareUpdateCurrentNewsMasterFromGitHubResult
	err    error
}

func NewPrepareUpdateCurrentNewsMasterFromGitHubResultFromJson(data string) PrepareUpdateCurrentNewsMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareUpdateCurrentNewsMasterFromGitHubResultFromDict(dict)
}

func NewPrepareUpdateCurrentNewsMasterFromGitHubResultFromDict(data map[string]interface{}) PrepareUpdateCurrentNewsMasterFromGitHubResult {
	return PrepareUpdateCurrentNewsMasterFromGitHubResult{
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
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

func (p PrepareUpdateCurrentNewsMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareUpdateCurrentNewsMasterFromGitHubResult) Pointer() *PrepareUpdateCurrentNewsMasterFromGitHubResult {
	return &p
}

type DescribeNewsResult struct {
	Items        []News               `json:"items"`
	ContentHash  *string              `json:"contentHash"`
	TemplateHash *string              `json:"templateHash"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type DescribeNewsAsyncResult struct {
	result *DescribeNewsResult
	err    error
}

func NewDescribeNewsResultFromJson(data string) DescribeNewsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNewsResultFromDict(dict)
}

func NewDescribeNewsResultFromDict(data map[string]interface{}) DescribeNewsResult {
	return DescribeNewsResult{
		Items: func() []News {
			if data["items"] == nil {
				return nil
			}
			return CastNewses(core.CastArray(data["items"]))
		}(),
		ContentHash: func() *string {
			v, ok := data["contentHash"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentHash"])
		}(),
		TemplateHash: func() *string {
			v, ok := data["templateHash"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["templateHash"])
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

func (p DescribeNewsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNewsesFromDict(
			p.Items,
		),
		"contentHash":  p.ContentHash,
		"templateHash": p.TemplateHash,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeNewsResult) Pointer() *DescribeNewsResult {
	return &p
}

type DescribeNewsByUserIdResult struct {
	Items        []News               `json:"items"`
	ContentHash  *string              `json:"contentHash"`
	TemplateHash *string              `json:"templateHash"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type DescribeNewsByUserIdAsyncResult struct {
	result *DescribeNewsByUserIdResult
	err    error
}

func NewDescribeNewsByUserIdResultFromJson(data string) DescribeNewsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNewsByUserIdResultFromDict(dict)
}

func NewDescribeNewsByUserIdResultFromDict(data map[string]interface{}) DescribeNewsByUserIdResult {
	return DescribeNewsByUserIdResult{
		Items: func() []News {
			if data["items"] == nil {
				return nil
			}
			return CastNewses(core.CastArray(data["items"]))
		}(),
		ContentHash: func() *string {
			v, ok := data["contentHash"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentHash"])
		}(),
		TemplateHash: func() *string {
			v, ok := data["templateHash"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["templateHash"])
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

func (p DescribeNewsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNewsesFromDict(
			p.Items,
		),
		"contentHash":  p.ContentHash,
		"templateHash": p.TemplateHash,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeNewsByUserIdResult) Pointer() *DescribeNewsByUserIdResult {
	return &p
}

type WantGrantResult struct {
	Items      []SetCookieRequestEntry `json:"items"`
	BrowserUrl *string                 `json:"browserUrl"`
	ZipUrl     *string                 `json:"zipUrl"`
	Metadata   *core.ResultMetadata    `json:"metadata"`
}

type WantGrantAsyncResult struct {
	result *WantGrantResult
	err    error
}

func NewWantGrantResultFromJson(data string) WantGrantResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWantGrantResultFromDict(dict)
}

func NewWantGrantResultFromDict(data map[string]interface{}) WantGrantResult {
	return WantGrantResult{
		Items: func() []SetCookieRequestEntry {
			if data["items"] == nil {
				return nil
			}
			return CastSetCookieRequestEntries(core.CastArray(data["items"]))
		}(),
		BrowserUrl: func() *string {
			v, ok := data["browserUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["browserUrl"])
		}(),
		ZipUrl: func() *string {
			v, ok := data["zipUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["zipUrl"])
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

func (p WantGrantResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSetCookieRequestEntriesFromDict(
			p.Items,
		),
		"browserUrl": p.BrowserUrl,
		"zipUrl":     p.ZipUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p WantGrantResult) Pointer() *WantGrantResult {
	return &p
}

type WantGrantByUserIdResult struct {
	Items      []SetCookieRequestEntry `json:"items"`
	BrowserUrl *string                 `json:"browserUrl"`
	ZipUrl     *string                 `json:"zipUrl"`
	Metadata   *core.ResultMetadata    `json:"metadata"`
}

type WantGrantByUserIdAsyncResult struct {
	result *WantGrantByUserIdResult
	err    error
}

func NewWantGrantByUserIdResultFromJson(data string) WantGrantByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWantGrantByUserIdResultFromDict(dict)
}

func NewWantGrantByUserIdResultFromDict(data map[string]interface{}) WantGrantByUserIdResult {
	return WantGrantByUserIdResult{
		Items: func() []SetCookieRequestEntry {
			if data["items"] == nil {
				return nil
			}
			return CastSetCookieRequestEntries(core.CastArray(data["items"]))
		}(),
		BrowserUrl: func() *string {
			v, ok := data["browserUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["browserUrl"])
		}(),
		ZipUrl: func() *string {
			v, ok := data["zipUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["zipUrl"])
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

func (p WantGrantByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSetCookieRequestEntriesFromDict(
			p.Items,
		),
		"browserUrl": p.BrowserUrl,
		"zipUrl":     p.ZipUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p WantGrantByUserIdResult) Pointer() *WantGrantByUserIdResult {
	return &p
}
