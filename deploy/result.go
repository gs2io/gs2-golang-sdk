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

package deploy

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeStacksResult struct {
	Items         []Stack              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeStacksAsyncResult struct {
	result *DescribeStacksResult
	err    error
}

func NewDescribeStacksResultFromJson(data string) DescribeStacksResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStacksResultFromDict(dict)
}

func NewDescribeStacksResultFromDict(data map[string]interface{}) DescribeStacksResult {
	return DescribeStacksResult{
		Items: func() []Stack {
			if data["items"] == nil {
				return nil
			}
			return CastStacks(core.CastArray(data["items"]))
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

func (p DescribeStacksResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStacksFromDict(
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

func (p DescribeStacksResult) Pointer() *DescribeStacksResult {
	return &p
}

type PreCreateStackResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreCreateStackAsyncResult struct {
	result *PreCreateStackResult
	err    error
}

func NewPreCreateStackResultFromJson(data string) PreCreateStackResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreCreateStackResultFromDict(dict)
}

func NewPreCreateStackResultFromDict(data map[string]interface{}) PreCreateStackResult {
	return PreCreateStackResult{
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PreCreateStackResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PreCreateStackResult) Pointer() *PreCreateStackResult {
	return &p
}

type CreateStackResult struct {
	Item     *Stack               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateStackAsyncResult struct {
	result *CreateStackResult
	err    error
}

func NewCreateStackResultFromJson(data string) CreateStackResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateStackResultFromDict(dict)
}

func NewCreateStackResultFromDict(data map[string]interface{}) CreateStackResult {
	return CreateStackResult{
		Item: func() *Stack {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStackFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateStackResult) ToDict() map[string]interface{} {
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

func (p CreateStackResult) Pointer() *CreateStackResult {
	return &p
}

type CreateStackFromGitHubResult struct {
	Item     *Stack               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateStackFromGitHubAsyncResult struct {
	result *CreateStackFromGitHubResult
	err    error
}

func NewCreateStackFromGitHubResultFromJson(data string) CreateStackFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateStackFromGitHubResultFromDict(dict)
}

func NewCreateStackFromGitHubResultFromDict(data map[string]interface{}) CreateStackFromGitHubResult {
	return CreateStackFromGitHubResult{
		Item: func() *Stack {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStackFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateStackFromGitHubResult) ToDict() map[string]interface{} {
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

func (p CreateStackFromGitHubResult) Pointer() *CreateStackFromGitHubResult {
	return &p
}

type PreValidateResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreValidateAsyncResult struct {
	result *PreValidateResult
	err    error
}

func NewPreValidateResultFromJson(data string) PreValidateResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreValidateResultFromDict(dict)
}

func NewPreValidateResultFromDict(data map[string]interface{}) PreValidateResult {
	return PreValidateResult{
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PreValidateResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PreValidateResult) Pointer() *PreValidateResult {
	return &p
}

type ValidateResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ValidateAsyncResult struct {
	result *ValidateResult
	err    error
}

func NewValidateResultFromJson(data string) ValidateResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewValidateResultFromDict(dict)
}

func NewValidateResultFromDict(data map[string]interface{}) ValidateResult {
	return ValidateResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ValidateResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ValidateResult) Pointer() *ValidateResult {
	return &p
}

type GetStackStatusResult struct {
	Status   *string              `json:"status"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStackStatusAsyncResult struct {
	result *GetStackStatusResult
	err    error
}

func NewGetStackStatusResultFromJson(data string) GetStackStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStackStatusResultFromDict(dict)
}

func NewGetStackStatusResultFromDict(data map[string]interface{}) GetStackStatusResult {
	return GetStackStatusResult{
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

func (p GetStackStatusResult) ToDict() map[string]interface{} {
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

func (p GetStackStatusResult) Pointer() *GetStackStatusResult {
	return &p
}

type GetStackResult struct {
	Item     *Stack               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStackAsyncResult struct {
	result *GetStackResult
	err    error
}

func NewGetStackResultFromJson(data string) GetStackResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStackResultFromDict(dict)
}

func NewGetStackResultFromDict(data map[string]interface{}) GetStackResult {
	return GetStackResult{
		Item: func() *Stack {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStackFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetStackResult) ToDict() map[string]interface{} {
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

func (p GetStackResult) Pointer() *GetStackResult {
	return &p
}

type PreUpdateStackResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreUpdateStackAsyncResult struct {
	result *PreUpdateStackResult
	err    error
}

func NewPreUpdateStackResultFromJson(data string) PreUpdateStackResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreUpdateStackResultFromDict(dict)
}

func NewPreUpdateStackResultFromDict(data map[string]interface{}) PreUpdateStackResult {
	return PreUpdateStackResult{
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PreUpdateStackResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PreUpdateStackResult) Pointer() *PreUpdateStackResult {
	return &p
}

type UpdateStackResult struct {
	Item     *Stack               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateStackAsyncResult struct {
	result *UpdateStackResult
	err    error
}

func NewUpdateStackResultFromJson(data string) UpdateStackResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStackResultFromDict(dict)
}

func NewUpdateStackResultFromDict(data map[string]interface{}) UpdateStackResult {
	return UpdateStackResult{
		Item: func() *Stack {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStackFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateStackResult) ToDict() map[string]interface{} {
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

func (p UpdateStackResult) Pointer() *UpdateStackResult {
	return &p
}

type PreChangeSetResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreChangeSetAsyncResult struct {
	result *PreChangeSetResult
	err    error
}

func NewPreChangeSetResultFromJson(data string) PreChangeSetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreChangeSetResultFromDict(dict)
}

func NewPreChangeSetResultFromDict(data map[string]interface{}) PreChangeSetResult {
	return PreChangeSetResult{
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PreChangeSetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PreChangeSetResult) Pointer() *PreChangeSetResult {
	return &p
}

type ChangeSetResult struct {
	Items    []ChangeSet          `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ChangeSetAsyncResult struct {
	result *ChangeSetResult
	err    error
}

func NewChangeSetResultFromJson(data string) ChangeSetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChangeSetResultFromDict(dict)
}

func NewChangeSetResultFromDict(data map[string]interface{}) ChangeSetResult {
	return ChangeSetResult{
		Items: func() []ChangeSet {
			if data["items"] == nil {
				return nil
			}
			return CastChangeSets(core.CastArray(data["items"]))
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

func (p ChangeSetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastChangeSetsFromDict(
			p.Items,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ChangeSetResult) Pointer() *ChangeSetResult {
	return &p
}

type UpdateStackFromGitHubResult struct {
	Item     *Stack               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateStackFromGitHubAsyncResult struct {
	result *UpdateStackFromGitHubResult
	err    error
}

func NewUpdateStackFromGitHubResultFromJson(data string) UpdateStackFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStackFromGitHubResultFromDict(dict)
}

func NewUpdateStackFromGitHubResultFromDict(data map[string]interface{}) UpdateStackFromGitHubResult {
	return UpdateStackFromGitHubResult{
		Item: func() *Stack {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStackFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateStackFromGitHubResult) ToDict() map[string]interface{} {
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

func (p UpdateStackFromGitHubResult) Pointer() *UpdateStackFromGitHubResult {
	return &p
}

type DeleteStackResult struct {
	Item     *Stack               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteStackAsyncResult struct {
	result *DeleteStackResult
	err    error
}

func NewDeleteStackResultFromJson(data string) DeleteStackResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStackResultFromDict(dict)
}

func NewDeleteStackResultFromDict(data map[string]interface{}) DeleteStackResult {
	return DeleteStackResult{
		Item: func() *Stack {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStackFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteStackResult) ToDict() map[string]interface{} {
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

func (p DeleteStackResult) Pointer() *DeleteStackResult {
	return &p
}

type ForceDeleteStackResult struct {
	Item     *Stack               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ForceDeleteStackAsyncResult struct {
	result *ForceDeleteStackResult
	err    error
}

func NewForceDeleteStackResultFromJson(data string) ForceDeleteStackResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewForceDeleteStackResultFromDict(dict)
}

func NewForceDeleteStackResultFromDict(data map[string]interface{}) ForceDeleteStackResult {
	return ForceDeleteStackResult{
		Item: func() *Stack {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStackFromDict(core.CastMap(data["item"])).Pointer()
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

func (p ForceDeleteStackResult) ToDict() map[string]interface{} {
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

func (p ForceDeleteStackResult) Pointer() *ForceDeleteStackResult {
	return &p
}

type DeleteStackResourcesResult struct {
	Item     *Stack               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteStackResourcesAsyncResult struct {
	result *DeleteStackResourcesResult
	err    error
}

func NewDeleteStackResourcesResultFromJson(data string) DeleteStackResourcesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStackResourcesResultFromDict(dict)
}

func NewDeleteStackResourcesResultFromDict(data map[string]interface{}) DeleteStackResourcesResult {
	return DeleteStackResourcesResult{
		Item: func() *Stack {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStackFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteStackResourcesResult) ToDict() map[string]interface{} {
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

func (p DeleteStackResourcesResult) Pointer() *DeleteStackResourcesResult {
	return &p
}

type DeleteStackEntityResult struct {
	Item     *Stack               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteStackEntityAsyncResult struct {
	result *DeleteStackEntityResult
	err    error
}

func NewDeleteStackEntityResultFromJson(data string) DeleteStackEntityResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStackEntityResultFromDict(dict)
}

func NewDeleteStackEntityResultFromDict(data map[string]interface{}) DeleteStackEntityResult {
	return DeleteStackEntityResult{
		Item: func() *Stack {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStackFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteStackEntityResult) ToDict() map[string]interface{} {
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

func (p DeleteStackEntityResult) Pointer() *DeleteStackEntityResult {
	return &p
}

type DescribeResourcesResult struct {
	Items         []Resource           `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeResourcesAsyncResult struct {
	result *DescribeResourcesResult
	err    error
}

func NewDescribeResourcesResultFromJson(data string) DescribeResourcesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeResourcesResultFromDict(dict)
}

func NewDescribeResourcesResultFromDict(data map[string]interface{}) DescribeResourcesResult {
	return DescribeResourcesResult{
		Items: func() []Resource {
			if data["items"] == nil {
				return nil
			}
			return CastResources(core.CastArray(data["items"]))
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

func (p DescribeResourcesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastResourcesFromDict(
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

func (p DescribeResourcesResult) Pointer() *DescribeResourcesResult {
	return &p
}

type GetResourceResult struct {
	Item     *Resource            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetResourceAsyncResult struct {
	result *GetResourceResult
	err    error
}

func NewGetResourceResultFromJson(data string) GetResourceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetResourceResultFromDict(dict)
}

func NewGetResourceResultFromDict(data map[string]interface{}) GetResourceResult {
	return GetResourceResult{
		Item: func() *Resource {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewResourceFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetResourceResult) ToDict() map[string]interface{} {
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

func (p GetResourceResult) Pointer() *GetResourceResult {
	return &p
}

type DescribeEventsResult struct {
	Items         []Event              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeEventsAsyncResult struct {
	result *DescribeEventsResult
	err    error
}

func NewDescribeEventsResultFromJson(data string) DescribeEventsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEventsResultFromDict(dict)
}

func NewDescribeEventsResultFromDict(data map[string]interface{}) DescribeEventsResult {
	return DescribeEventsResult{
		Items: func() []Event {
			if data["items"] == nil {
				return nil
			}
			return CastEvents(core.CastArray(data["items"]))
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

func (p DescribeEventsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEventsFromDict(
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

func (p DescribeEventsResult) Pointer() *DescribeEventsResult {
	return &p
}

type GetEventResult struct {
	Item     *Event               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetEventAsyncResult struct {
	result *GetEventResult
	err    error
}

func NewGetEventResultFromJson(data string) GetEventResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEventResultFromDict(dict)
}

func NewGetEventResultFromDict(data map[string]interface{}) GetEventResult {
	return GetEventResult{
		Item: func() *Event {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEventFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetEventResult) ToDict() map[string]interface{} {
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

func (p GetEventResult) Pointer() *GetEventResult {
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
