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

package serialKey

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
	return DumpUserDataByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CheckDumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
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
	return CleanUserDataByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
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
	return CheckCleanUserDataByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CheckCleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareImportUserDataByUserIdResult) ToDict() map[string]interface{} {
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
	return ImportUserDataByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CheckImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CheckImportUserDataByUserIdResult) Pointer() *CheckImportUserDataByUserIdResult {
	return &p
}

type DescribeIssueJobsResult struct {
	Items         []IssueJob           `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeIssueJobsAsyncResult struct {
	result *DescribeIssueJobsResult
	err    error
}

func NewDescribeIssueJobsResultFromJson(data string) DescribeIssueJobsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeIssueJobsResultFromDict(dict)
}

func NewDescribeIssueJobsResultFromDict(data map[string]interface{}) DescribeIssueJobsResult {
	return DescribeIssueJobsResult{
		Items: func() []IssueJob {
			if data["items"] == nil {
				return nil
			}
			return CastIssueJobs(core.CastArray(data["items"]))
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

func (p DescribeIssueJobsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastIssueJobsFromDict(
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

func (p DescribeIssueJobsResult) Pointer() *DescribeIssueJobsResult {
	return &p
}

type GetIssueJobResult struct {
	Item     *IssueJob            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetIssueJobAsyncResult struct {
	result *GetIssueJobResult
	err    error
}

func NewGetIssueJobResultFromJson(data string) GetIssueJobResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetIssueJobResultFromDict(dict)
}

func NewGetIssueJobResultFromDict(data map[string]interface{}) GetIssueJobResult {
	return GetIssueJobResult{
		Item: func() *IssueJob {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIssueJobFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetIssueJobResult) ToDict() map[string]interface{} {
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

func (p GetIssueJobResult) Pointer() *GetIssueJobResult {
	return &p
}

type IssueResult struct {
	Item     *IssueJob            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type IssueAsyncResult struct {
	result *IssueResult
	err    error
}

func NewIssueResultFromJson(data string) IssueResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIssueResultFromDict(dict)
}

func NewIssueResultFromDict(data map[string]interface{}) IssueResult {
	return IssueResult{
		Item: func() *IssueJob {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIssueJobFromDict(core.CastMap(data["item"])).Pointer()
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

func (p IssueResult) ToDict() map[string]interface{} {
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

func (p IssueResult) Pointer() *IssueResult {
	return &p
}

type DescribeSerialKeysResult struct {
	Items         []SerialKey          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSerialKeysAsyncResult struct {
	result *DescribeSerialKeysResult
	err    error
}

func NewDescribeSerialKeysResultFromJson(data string) DescribeSerialKeysResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSerialKeysResultFromDict(dict)
}

func NewDescribeSerialKeysResultFromDict(data map[string]interface{}) DescribeSerialKeysResult {
	return DescribeSerialKeysResult{
		Items: func() []SerialKey {
			if data["items"] == nil {
				return nil
			}
			return CastSerialKeys(core.CastArray(data["items"]))
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

func (p DescribeSerialKeysResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSerialKeysFromDict(
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

func (p DescribeSerialKeysResult) Pointer() *DescribeSerialKeysResult {
	return &p
}

type DownloadSerialCodesResult struct {
	Url      *string              `json:"url"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DownloadSerialCodesAsyncResult struct {
	result *DownloadSerialCodesResult
	err    error
}

func NewDownloadSerialCodesResultFromJson(data string) DownloadSerialCodesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDownloadSerialCodesResultFromDict(dict)
}

func NewDownloadSerialCodesResultFromDict(data map[string]interface{}) DownloadSerialCodesResult {
	return DownloadSerialCodesResult{
		Url: func() *string {
			v, ok := data["url"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["url"])
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

func (p DownloadSerialCodesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DownloadSerialCodesResult) Pointer() *DownloadSerialCodesResult {
	return &p
}

type IssueOnceResult struct {
	Item     *SerialKey           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type IssueOnceAsyncResult struct {
	result *IssueOnceResult
	err    error
}

func NewIssueOnceResultFromJson(data string) IssueOnceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIssueOnceResultFromDict(dict)
}

func NewIssueOnceResultFromDict(data map[string]interface{}) IssueOnceResult {
	return IssueOnceResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p IssueOnceResult) ToDict() map[string]interface{} {
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

func (p IssueOnceResult) Pointer() *IssueOnceResult {
	return &p
}

type GetSerialKeyResult struct {
	Item          *SerialKey           `json:"item"`
	CampaignModel *CampaignModel       `json:"campaignModel"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type GetSerialKeyAsyncResult struct {
	result *GetSerialKeyResult
	err    error
}

func NewGetSerialKeyResultFromJson(data string) GetSerialKeyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSerialKeyResultFromDict(dict)
}

func NewGetSerialKeyResultFromDict(data map[string]interface{}) GetSerialKeyResult {
	return GetSerialKeyResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		CampaignModel: func() *CampaignModel {
			v, ok := data["campaignModel"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer()
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

func (p GetSerialKeyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"campaignModel": func() map[string]interface{} {
			if p.CampaignModel == nil {
				return nil
			}
			return p.CampaignModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetSerialKeyResult) Pointer() *GetSerialKeyResult {
	return &p
}

type VerifyCodeResult struct {
	Item          *SerialKey           `json:"item"`
	CampaignModel *CampaignModel       `json:"campaignModel"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type VerifyCodeAsyncResult struct {
	result *VerifyCodeResult
	err    error
}

func NewVerifyCodeResultFromJson(data string) VerifyCodeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyCodeResultFromDict(dict)
}

func NewVerifyCodeResultFromDict(data map[string]interface{}) VerifyCodeResult {
	return VerifyCodeResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		CampaignModel: func() *CampaignModel {
			v, ok := data["campaignModel"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer()
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

func (p VerifyCodeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"campaignModel": func() map[string]interface{} {
			if p.CampaignModel == nil {
				return nil
			}
			return p.CampaignModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyCodeResult) Pointer() *VerifyCodeResult {
	return &p
}

type VerifyCodeByUserIdResult struct {
	Item          *SerialKey           `json:"item"`
	CampaignModel *CampaignModel       `json:"campaignModel"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type VerifyCodeByUserIdAsyncResult struct {
	result *VerifyCodeByUserIdResult
	err    error
}

func NewVerifyCodeByUserIdResultFromJson(data string) VerifyCodeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyCodeByUserIdResultFromDict(dict)
}

func NewVerifyCodeByUserIdResultFromDict(data map[string]interface{}) VerifyCodeByUserIdResult {
	return VerifyCodeByUserIdResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		CampaignModel: func() *CampaignModel {
			v, ok := data["campaignModel"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer()
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

func (p VerifyCodeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"campaignModel": func() map[string]interface{} {
			if p.CampaignModel == nil {
				return nil
			}
			return p.CampaignModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyCodeByUserIdResult) Pointer() *VerifyCodeByUserIdResult {
	return &p
}

type UseResult struct {
	Item          *SerialKey           `json:"item"`
	CampaignModel *CampaignModel       `json:"campaignModel"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type UseAsyncResult struct {
	result *UseResult
	err    error
}

func NewUseResultFromJson(data string) UseResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUseResultFromDict(dict)
}

func NewUseResultFromDict(data map[string]interface{}) UseResult {
	return UseResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		CampaignModel: func() *CampaignModel {
			v, ok := data["campaignModel"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer()
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

func (p UseResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"campaignModel": func() map[string]interface{} {
			if p.CampaignModel == nil {
				return nil
			}
			return p.CampaignModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UseResult) Pointer() *UseResult {
	return &p
}

type UseByUserIdResult struct {
	Item          *SerialKey           `json:"item"`
	CampaignModel *CampaignModel       `json:"campaignModel"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type UseByUserIdAsyncResult struct {
	result *UseByUserIdResult
	err    error
}

func NewUseByUserIdResultFromJson(data string) UseByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUseByUserIdResultFromDict(dict)
}

func NewUseByUserIdResultFromDict(data map[string]interface{}) UseByUserIdResult {
	return UseByUserIdResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		CampaignModel: func() *CampaignModel {
			v, ok := data["campaignModel"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer()
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

func (p UseByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"campaignModel": func() map[string]interface{} {
			if p.CampaignModel == nil {
				return nil
			}
			return p.CampaignModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UseByUserIdResult) Pointer() *UseByUserIdResult {
	return &p
}

type RevertUseByUserIdResult struct {
	Item          *SerialKey           `json:"item"`
	CampaignModel *CampaignModel       `json:"campaignModel"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type RevertUseByUserIdAsyncResult struct {
	result *RevertUseByUserIdResult
	err    error
}

func NewRevertUseByUserIdResultFromJson(data string) RevertUseByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertUseByUserIdResultFromDict(dict)
}

func NewRevertUseByUserIdResultFromDict(data map[string]interface{}) RevertUseByUserIdResult {
	return RevertUseByUserIdResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		CampaignModel: func() *CampaignModel {
			v, ok := data["campaignModel"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer()
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

func (p RevertUseByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"campaignModel": func() map[string]interface{} {
			if p.CampaignModel == nil {
				return nil
			}
			return p.CampaignModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RevertUseByUserIdResult) Pointer() *RevertUseByUserIdResult {
	return &p
}

type UseByStampTaskResult struct {
	Item            *SerialKey           `json:"item"`
	CampaignModel   *CampaignModel       `json:"campaignModel"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type UseByStampTaskAsyncResult struct {
	result *UseByStampTaskResult
	err    error
}

func NewUseByStampTaskResultFromJson(data string) UseByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUseByStampTaskResultFromDict(dict)
}

func NewUseByStampTaskResultFromDict(data map[string]interface{}) UseByStampTaskResult {
	return UseByStampTaskResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		CampaignModel: func() *CampaignModel {
			v, ok := data["campaignModel"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
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

func (p UseByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"campaignModel": func() map[string]interface{} {
			if p.CampaignModel == nil {
				return nil
			}
			return p.CampaignModel.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UseByStampTaskResult) Pointer() *UseByStampTaskResult {
	return &p
}

type RevertUseByStampSheetResult struct {
	Item          *SerialKey           `json:"item"`
	CampaignModel *CampaignModel       `json:"campaignModel"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type RevertUseByStampSheetAsyncResult struct {
	result *RevertUseByStampSheetResult
	err    error
}

func NewRevertUseByStampSheetResultFromJson(data string) RevertUseByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertUseByStampSheetResultFromDict(dict)
}

func NewRevertUseByStampSheetResultFromDict(data map[string]interface{}) RevertUseByStampSheetResult {
	return RevertUseByStampSheetResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		CampaignModel: func() *CampaignModel {
			v, ok := data["campaignModel"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer()
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

func (p RevertUseByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"campaignModel": func() map[string]interface{} {
			if p.CampaignModel == nil {
				return nil
			}
			return p.CampaignModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RevertUseByStampSheetResult) Pointer() *RevertUseByStampSheetResult {
	return &p
}

type VerifyByStampTaskResult struct {
	Item            *SerialKey           `json:"item"`
	CampaignModel   *CampaignModel       `json:"campaignModel"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyByStampTaskAsyncResult struct {
	result *VerifyByStampTaskResult
	err    error
}

func NewVerifyByStampTaskResultFromJson(data string) VerifyByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyByStampTaskResultFromDict(dict)
}

func NewVerifyByStampTaskResultFromDict(data map[string]interface{}) VerifyByStampTaskResult {
	return VerifyByStampTaskResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		CampaignModel: func() *CampaignModel {
			v, ok := data["campaignModel"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["campaignModel"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
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

func (p VerifyByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"campaignModel": func() map[string]interface{} {
			if p.CampaignModel == nil {
				return nil
			}
			return p.CampaignModel.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyByStampTaskResult) Pointer() *VerifyByStampTaskResult {
	return &p
}

type IssueOnceByStampSheetResult struct {
	Item     *SerialKey           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type IssueOnceByStampSheetAsyncResult struct {
	result *IssueOnceByStampSheetResult
	err    error
}

func NewIssueOnceByStampSheetResultFromJson(data string) IssueOnceByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIssueOnceByStampSheetResultFromDict(dict)
}

func NewIssueOnceByStampSheetResultFromDict(data map[string]interface{}) IssueOnceByStampSheetResult {
	return IssueOnceByStampSheetResult{
		Item: func() *SerialKey {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSerialKeyFromDict(core.CastMap(data["item"])).Pointer()
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

func (p IssueOnceByStampSheetResult) ToDict() map[string]interface{} {
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

func (p IssueOnceByStampSheetResult) Pointer() *IssueOnceByStampSheetResult {
	return &p
}

type DescribeCampaignModelsResult struct {
	Items    []CampaignModel      `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeCampaignModelsAsyncResult struct {
	result *DescribeCampaignModelsResult
	err    error
}

func NewDescribeCampaignModelsResultFromJson(data string) DescribeCampaignModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCampaignModelsResultFromDict(dict)
}

func NewDescribeCampaignModelsResultFromDict(data map[string]interface{}) DescribeCampaignModelsResult {
	return DescribeCampaignModelsResult{
		Items: func() []CampaignModel {
			if data["items"] == nil {
				return nil
			}
			return CastCampaignModels(core.CastArray(data["items"]))
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

func (p DescribeCampaignModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCampaignModelsFromDict(
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

func (p DescribeCampaignModelsResult) Pointer() *DescribeCampaignModelsResult {
	return &p
}

type GetCampaignModelResult struct {
	Item     *CampaignModel       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCampaignModelAsyncResult struct {
	result *GetCampaignModelResult
	err    error
}

func NewGetCampaignModelResultFromJson(data string) GetCampaignModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCampaignModelResultFromDict(dict)
}

func NewGetCampaignModelResultFromDict(data map[string]interface{}) GetCampaignModelResult {
	return GetCampaignModelResult{
		Item: func() *CampaignModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetCampaignModelResult) ToDict() map[string]interface{} {
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

func (p GetCampaignModelResult) Pointer() *GetCampaignModelResult {
	return &p
}

type DescribeCampaignModelMastersResult struct {
	Items         []CampaignModelMaster `json:"items"`
	NextPageToken *string               `json:"nextPageToken"`
	Metadata      *core.ResultMetadata  `json:"metadata"`
}

type DescribeCampaignModelMastersAsyncResult struct {
	result *DescribeCampaignModelMastersResult
	err    error
}

func NewDescribeCampaignModelMastersResultFromJson(data string) DescribeCampaignModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCampaignModelMastersResultFromDict(dict)
}

func NewDescribeCampaignModelMastersResultFromDict(data map[string]interface{}) DescribeCampaignModelMastersResult {
	return DescribeCampaignModelMastersResult{
		Items: func() []CampaignModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastCampaignModelMasters(core.CastArray(data["items"]))
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

func (p DescribeCampaignModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCampaignModelMastersFromDict(
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

func (p DescribeCampaignModelMastersResult) Pointer() *DescribeCampaignModelMastersResult {
	return &p
}

type CreateCampaignModelMasterResult struct {
	Item     *CampaignModelMaster `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateCampaignModelMasterAsyncResult struct {
	result *CreateCampaignModelMasterResult
	err    error
}

func NewCreateCampaignModelMasterResultFromJson(data string) CreateCampaignModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateCampaignModelMasterResultFromDict(dict)
}

func NewCreateCampaignModelMasterResultFromDict(data map[string]interface{}) CreateCampaignModelMasterResult {
	return CreateCampaignModelMasterResult{
		Item: func() *CampaignModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateCampaignModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateCampaignModelMasterResult) Pointer() *CreateCampaignModelMasterResult {
	return &p
}

type GetCampaignModelMasterResult struct {
	Item     *CampaignModelMaster `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCampaignModelMasterAsyncResult struct {
	result *GetCampaignModelMasterResult
	err    error
}

func NewGetCampaignModelMasterResultFromJson(data string) GetCampaignModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCampaignModelMasterResultFromDict(dict)
}

func NewGetCampaignModelMasterResultFromDict(data map[string]interface{}) GetCampaignModelMasterResult {
	return GetCampaignModelMasterResult{
		Item: func() *CampaignModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetCampaignModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetCampaignModelMasterResult) Pointer() *GetCampaignModelMasterResult {
	return &p
}

type UpdateCampaignModelMasterResult struct {
	Item     *CampaignModelMaster `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCampaignModelMasterAsyncResult struct {
	result *UpdateCampaignModelMasterResult
	err    error
}

func NewUpdateCampaignModelMasterResultFromJson(data string) UpdateCampaignModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCampaignModelMasterResultFromDict(dict)
}

func NewUpdateCampaignModelMasterResultFromDict(data map[string]interface{}) UpdateCampaignModelMasterResult {
	return UpdateCampaignModelMasterResult{
		Item: func() *CampaignModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCampaignModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateCampaignModelMasterResult) Pointer() *UpdateCampaignModelMasterResult {
	return &p
}

type DeleteCampaignModelMasterResult struct {
	Item     *CampaignModelMaster `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteCampaignModelMasterAsyncResult struct {
	result *DeleteCampaignModelMasterResult
	err    error
}

func NewDeleteCampaignModelMasterResultFromJson(data string) DeleteCampaignModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCampaignModelMasterResultFromDict(dict)
}

func NewDeleteCampaignModelMasterResultFromDict(data map[string]interface{}) DeleteCampaignModelMasterResult {
	return DeleteCampaignModelMasterResult{
		Item: func() *CampaignModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCampaignModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteCampaignModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteCampaignModelMasterResult) Pointer() *DeleteCampaignModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentCampaignMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
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
		Item: func() *CurrentCampaignMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentCampaignMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p ExportMasterResult) ToDict() map[string]interface{} {
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

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentCampaignMasterResult struct {
	Item     *CurrentCampaignMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetCurrentCampaignMasterAsyncResult struct {
	result *GetCurrentCampaignMasterResult
	err    error
}

func NewGetCurrentCampaignMasterResultFromJson(data string) GetCurrentCampaignMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentCampaignMasterResultFromDict(dict)
}

func NewGetCurrentCampaignMasterResultFromDict(data map[string]interface{}) GetCurrentCampaignMasterResult {
	return GetCurrentCampaignMasterResult{
		Item: func() *CurrentCampaignMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentCampaignMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetCurrentCampaignMasterResult) ToDict() map[string]interface{} {
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

func (p GetCurrentCampaignMasterResult) Pointer() *GetCurrentCampaignMasterResult {
	return &p
}

type PreUpdateCurrentCampaignMasterResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreUpdateCurrentCampaignMasterAsyncResult struct {
	result *PreUpdateCurrentCampaignMasterResult
	err    error
}

func NewPreUpdateCurrentCampaignMasterResultFromJson(data string) PreUpdateCurrentCampaignMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreUpdateCurrentCampaignMasterResultFromDict(dict)
}

func NewPreUpdateCurrentCampaignMasterResultFromDict(data map[string]interface{}) PreUpdateCurrentCampaignMasterResult {
	return PreUpdateCurrentCampaignMasterResult{
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

func (p PreUpdateCurrentCampaignMasterResult) ToDict() map[string]interface{} {
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

func (p PreUpdateCurrentCampaignMasterResult) Pointer() *PreUpdateCurrentCampaignMasterResult {
	return &p
}

type UpdateCurrentCampaignMasterResult struct {
	Item     *CurrentCampaignMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type UpdateCurrentCampaignMasterAsyncResult struct {
	result *UpdateCurrentCampaignMasterResult
	err    error
}

func NewUpdateCurrentCampaignMasterResultFromJson(data string) UpdateCurrentCampaignMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentCampaignMasterResultFromDict(dict)
}

func NewUpdateCurrentCampaignMasterResultFromDict(data map[string]interface{}) UpdateCurrentCampaignMasterResult {
	return UpdateCurrentCampaignMasterResult{
		Item: func() *CurrentCampaignMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentCampaignMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentCampaignMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentCampaignMasterResult) Pointer() *UpdateCurrentCampaignMasterResult {
	return &p
}

type UpdateCurrentCampaignMasterFromGitHubResult struct {
	Item     *CurrentCampaignMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type UpdateCurrentCampaignMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentCampaignMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentCampaignMasterFromGitHubResultFromJson(data string) UpdateCurrentCampaignMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentCampaignMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentCampaignMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentCampaignMasterFromGitHubResult {
	return UpdateCurrentCampaignMasterFromGitHubResult{
		Item: func() *CurrentCampaignMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentCampaignMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentCampaignMasterFromGitHubResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentCampaignMasterFromGitHubResult) Pointer() *UpdateCurrentCampaignMasterFromGitHubResult {
	return &p
}
