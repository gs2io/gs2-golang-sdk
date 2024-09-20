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
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
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
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
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

type DescribeProgressesResult struct {
	Items         []Progress `json:"items"`
	NextPageToken *string    `json:"nextPageToken"`
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
		Items:         CastProgresses(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeProgressesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastProgressesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeProgressesResult) Pointer() *DescribeProgressesResult {
	return &p
}

type GetProgressResult struct {
	Item *Progress `json:"item"`
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
		Item: NewProgressFromDict(core.CastMap(data["item"])).Pointer(),
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
	}
}

func (p GetProgressResult) Pointer() *GetProgressResult {
	return &p
}

type DescribeOutputsResult struct {
	Items         []Output `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
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
		Items:         CastOutputs(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeOutputsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastOutputsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeOutputsResult) Pointer() *DescribeOutputsResult {
	return &p
}

type GetOutputResult struct {
	Item *Output `json:"item"`
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
		Item: NewOutputFromDict(core.CastMap(data["item"])).Pointer(),
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
	}
}

func (p GetOutputResult) Pointer() *GetOutputResult {
	return &p
}

type PrepareUpdateCurrentNewsMasterResult struct {
	UploadToken       *string `json:"uploadToken"`
	TemplateUploadUrl *string `json:"templateUploadUrl"`
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
		UploadToken:       core.CastString(data["uploadToken"]),
		TemplateUploadUrl: core.CastString(data["templateUploadUrl"]),
	}
}

func (p PrepareUpdateCurrentNewsMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken":       p.UploadToken,
		"templateUploadUrl": p.TemplateUploadUrl,
	}
}

func (p PrepareUpdateCurrentNewsMasterResult) Pointer() *PrepareUpdateCurrentNewsMasterResult {
	return &p
}

type UpdateCurrentNewsMasterResult struct {
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
	return UpdateCurrentNewsMasterResult{}
}

func (p UpdateCurrentNewsMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p UpdateCurrentNewsMasterResult) Pointer() *UpdateCurrentNewsMasterResult {
	return &p
}

type PrepareUpdateCurrentNewsMasterFromGitHubResult struct {
	UploadToken *string `json:"uploadToken"`
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
		UploadToken: core.CastString(data["uploadToken"]),
	}
}

func (p PrepareUpdateCurrentNewsMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
	}
}

func (p PrepareUpdateCurrentNewsMasterFromGitHubResult) Pointer() *PrepareUpdateCurrentNewsMasterFromGitHubResult {
	return &p
}

type DescribeNewsResult struct {
	Items        []News  `json:"items"`
	ContentHash  *string `json:"contentHash"`
	TemplateHash *string `json:"templateHash"`
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
		Items:        CastNewses(core.CastArray(data["items"])),
		ContentHash:  core.CastString(data["contentHash"]),
		TemplateHash: core.CastString(data["templateHash"]),
	}
}

func (p DescribeNewsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNewsesFromDict(
			p.Items,
		),
		"contentHash":  p.ContentHash,
		"templateHash": p.TemplateHash,
	}
}

func (p DescribeNewsResult) Pointer() *DescribeNewsResult {
	return &p
}

type DescribeNewsByUserIdResult struct {
	Items        []News  `json:"items"`
	ContentHash  *string `json:"contentHash"`
	TemplateHash *string `json:"templateHash"`
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
		Items:        CastNewses(core.CastArray(data["items"])),
		ContentHash:  core.CastString(data["contentHash"]),
		TemplateHash: core.CastString(data["templateHash"]),
	}
}

func (p DescribeNewsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNewsesFromDict(
			p.Items,
		),
		"contentHash":  p.ContentHash,
		"templateHash": p.TemplateHash,
	}
}

func (p DescribeNewsByUserIdResult) Pointer() *DescribeNewsByUserIdResult {
	return &p
}

type WantGrantResult struct {
	Items      []SetCookieRequestEntry `json:"items"`
	BrowserUrl *string                 `json:"browserUrl"`
	ZipUrl     *string                 `json:"zipUrl"`
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
		Items:      CastSetCookieRequestEntries(core.CastArray(data["items"])),
		BrowserUrl: core.CastString(data["browserUrl"]),
		ZipUrl:     core.CastString(data["zipUrl"]),
	}
}

func (p WantGrantResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSetCookieRequestEntriesFromDict(
			p.Items,
		),
		"browserUrl": p.BrowserUrl,
		"zipUrl":     p.ZipUrl,
	}
}

func (p WantGrantResult) Pointer() *WantGrantResult {
	return &p
}

type WantGrantByUserIdResult struct {
	Items      []SetCookieRequestEntry `json:"items"`
	BrowserUrl *string                 `json:"browserUrl"`
	ZipUrl     *string                 `json:"zipUrl"`
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
		Items:      CastSetCookieRequestEntries(core.CastArray(data["items"])),
		BrowserUrl: core.CastString(data["browserUrl"]),
		ZipUrl:     core.CastString(data["zipUrl"]),
	}
}

func (p WantGrantByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSetCookieRequestEntriesFromDict(
			p.Items,
		),
		"browserUrl": p.BrowserUrl,
		"zipUrl":     p.ZipUrl,
	}
}

func (p WantGrantByUserIdResult) Pointer() *WantGrantByUserIdResult {
	return &p
}
