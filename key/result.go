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

package key

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

type DescribeKeysResult struct {
	Items         []Key                `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeKeysAsyncResult struct {
	result *DescribeKeysResult
	err    error
}

func NewDescribeKeysResultFromJson(data string) DescribeKeysResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeKeysResultFromDict(dict)
}

func NewDescribeKeysResultFromDict(data map[string]interface{}) DescribeKeysResult {
	return DescribeKeysResult{
		Items: func() []Key {
			if data["items"] == nil {
				return nil
			}
			return CastKeys(core.CastArray(data["items"]))
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

func (p DescribeKeysResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastKeysFromDict(
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

func (p DescribeKeysResult) Pointer() *DescribeKeysResult {
	return &p
}

type CreateKeyResult struct {
	Item     *Key                 `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateKeyAsyncResult struct {
	result *CreateKeyResult
	err    error
}

func NewCreateKeyResultFromJson(data string) CreateKeyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateKeyResultFromDict(dict)
}

func NewCreateKeyResultFromDict(data map[string]interface{}) CreateKeyResult {
	return CreateKeyResult{
		Item: func() *Key {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateKeyResult) ToDict() map[string]interface{} {
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

func (p CreateKeyResult) Pointer() *CreateKeyResult {
	return &p
}

type UpdateKeyResult struct {
	Item     *Key                 `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateKeyAsyncResult struct {
	result *UpdateKeyResult
	err    error
}

func NewUpdateKeyResultFromJson(data string) UpdateKeyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateKeyResultFromDict(dict)
}

func NewUpdateKeyResultFromDict(data map[string]interface{}) UpdateKeyResult {
	return UpdateKeyResult{
		Item: func() *Key {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateKeyResult) ToDict() map[string]interface{} {
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

func (p UpdateKeyResult) Pointer() *UpdateKeyResult {
	return &p
}

type GetKeyResult struct {
	Item     *Key                 `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetKeyAsyncResult struct {
	result *GetKeyResult
	err    error
}

func NewGetKeyResultFromJson(data string) GetKeyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetKeyResultFromDict(dict)
}

func NewGetKeyResultFromDict(data map[string]interface{}) GetKeyResult {
	return GetKeyResult{
		Item: func() *Key {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetKeyResult) ToDict() map[string]interface{} {
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

func (p GetKeyResult) Pointer() *GetKeyResult {
	return &p
}

type DeleteKeyResult struct {
	Item     *Key                 `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteKeyAsyncResult struct {
	result *DeleteKeyResult
	err    error
}

func NewDeleteKeyResultFromJson(data string) DeleteKeyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteKeyResultFromDict(dict)
}

func NewDeleteKeyResultFromDict(data map[string]interface{}) DeleteKeyResult {
	return DeleteKeyResult{
		Item: func() *Key {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteKeyResult) ToDict() map[string]interface{} {
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

func (p DeleteKeyResult) Pointer() *DeleteKeyResult {
	return &p
}

type EncryptResult struct {
	Data     *string              `json:"data"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type EncryptAsyncResult struct {
	result *EncryptResult
	err    error
}

func NewEncryptResultFromJson(data string) EncryptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEncryptResultFromDict(dict)
}

func NewEncryptResultFromDict(data map[string]interface{}) EncryptResult {
	return EncryptResult{
		Data: func() *string {
			v, ok := data["data"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["data"])
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

func (p EncryptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"data": p.Data,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p EncryptResult) Pointer() *EncryptResult {
	return &p
}

type DecryptResult struct {
	Data     *string              `json:"data"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DecryptAsyncResult struct {
	result *DecryptResult
	err    error
}

func NewDecryptResultFromJson(data string) DecryptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecryptResultFromDict(dict)
}

func NewDecryptResultFromDict(data map[string]interface{}) DecryptResult {
	return DecryptResult{
		Data: func() *string {
			v, ok := data["data"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["data"])
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

func (p DecryptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"data": p.Data,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DecryptResult) Pointer() *DecryptResult {
	return &p
}

type DescribeGitHubApiKeysResult struct {
	Items         []GitHubApiKey       `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeGitHubApiKeysAsyncResult struct {
	result *DescribeGitHubApiKeysResult
	err    error
}

func NewDescribeGitHubApiKeysResultFromJson(data string) DescribeGitHubApiKeysResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGitHubApiKeysResultFromDict(dict)
}

func NewDescribeGitHubApiKeysResultFromDict(data map[string]interface{}) DescribeGitHubApiKeysResult {
	return DescribeGitHubApiKeysResult{
		Items: func() []GitHubApiKey {
			if data["items"] == nil {
				return nil
			}
			return CastGitHubApiKeys(core.CastArray(data["items"]))
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

func (p DescribeGitHubApiKeysResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGitHubApiKeysFromDict(
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

func (p DescribeGitHubApiKeysResult) Pointer() *DescribeGitHubApiKeysResult {
	return &p
}

type CreateGitHubApiKeyResult struct {
	Item     *GitHubApiKey        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateGitHubApiKeyAsyncResult struct {
	result *CreateGitHubApiKeyResult
	err    error
}

func NewCreateGitHubApiKeyResultFromJson(data string) CreateGitHubApiKeyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGitHubApiKeyResultFromDict(dict)
}

func NewCreateGitHubApiKeyResultFromDict(data map[string]interface{}) CreateGitHubApiKeyResult {
	return CreateGitHubApiKeyResult{
		Item: func() *GitHubApiKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGitHubApiKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateGitHubApiKeyResult) ToDict() map[string]interface{} {
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

func (p CreateGitHubApiKeyResult) Pointer() *CreateGitHubApiKeyResult {
	return &p
}

type UpdateGitHubApiKeyResult struct {
	Item     *GitHubApiKey        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateGitHubApiKeyAsyncResult struct {
	result *UpdateGitHubApiKeyResult
	err    error
}

func NewUpdateGitHubApiKeyResultFromJson(data string) UpdateGitHubApiKeyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGitHubApiKeyResultFromDict(dict)
}

func NewUpdateGitHubApiKeyResultFromDict(data map[string]interface{}) UpdateGitHubApiKeyResult {
	return UpdateGitHubApiKeyResult{
		Item: func() *GitHubApiKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGitHubApiKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateGitHubApiKeyResult) ToDict() map[string]interface{} {
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

func (p UpdateGitHubApiKeyResult) Pointer() *UpdateGitHubApiKeyResult {
	return &p
}

type GetGitHubApiKeyResult struct {
	Item     *GitHubApiKey        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGitHubApiKeyAsyncResult struct {
	result *GetGitHubApiKeyResult
	err    error
}

func NewGetGitHubApiKeyResultFromJson(data string) GetGitHubApiKeyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGitHubApiKeyResultFromDict(dict)
}

func NewGetGitHubApiKeyResultFromDict(data map[string]interface{}) GetGitHubApiKeyResult {
	return GetGitHubApiKeyResult{
		Item: func() *GitHubApiKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGitHubApiKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetGitHubApiKeyResult) ToDict() map[string]interface{} {
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

func (p GetGitHubApiKeyResult) Pointer() *GetGitHubApiKeyResult {
	return &p
}

type DeleteGitHubApiKeyResult struct {
	Item     *GitHubApiKey        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteGitHubApiKeyAsyncResult struct {
	result *DeleteGitHubApiKeyResult
	err    error
}

func NewDeleteGitHubApiKeyResultFromJson(data string) DeleteGitHubApiKeyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGitHubApiKeyResultFromDict(dict)
}

func NewDeleteGitHubApiKeyResultFromDict(data map[string]interface{}) DeleteGitHubApiKeyResult {
	return DeleteGitHubApiKeyResult{
		Item: func() *GitHubApiKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGitHubApiKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteGitHubApiKeyResult) ToDict() map[string]interface{} {
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

func (p DeleteGitHubApiKeyResult) Pointer() *DeleteGitHubApiKeyResult {
	return &p
}
