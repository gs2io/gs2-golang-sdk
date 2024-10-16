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

package jobQueue

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

type DescribeJobsByUserIdResult struct {
	Items         []Job   `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeJobsByUserIdAsyncResult struct {
	result *DescribeJobsByUserIdResult
	err    error
}

func NewDescribeJobsByUserIdResultFromJson(data string) DescribeJobsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeJobsByUserIdResultFromDict(dict)
}

func NewDescribeJobsByUserIdResultFromDict(data map[string]interface{}) DescribeJobsByUserIdResult {
	return DescribeJobsByUserIdResult{
		Items: func() []Job {
			if data["items"] == nil {
				return nil
			}
			return CastJobs(core.CastArray(data["items"]))
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

func (p DescribeJobsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastJobsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeJobsByUserIdResult) Pointer() *DescribeJobsByUserIdResult {
	return &p
}

type GetJobByUserIdResult struct {
	Item *Job `json:"item"`
}

type GetJobByUserIdAsyncResult struct {
	result *GetJobByUserIdResult
	err    error
}

func NewGetJobByUserIdResultFromJson(data string) GetJobByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJobByUserIdResultFromDict(dict)
}

func NewGetJobByUserIdResultFromDict(data map[string]interface{}) GetJobByUserIdResult {
	return GetJobByUserIdResult{
		Item: func() *Job {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJobFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetJobByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetJobByUserIdResult) Pointer() *GetJobByUserIdResult {
	return &p
}

type PushByUserIdResult struct {
	Items   []Job `json:"items"`
	AutoRun *bool `json:"autoRun"`
}

type PushByUserIdAsyncResult struct {
	result *PushByUserIdResult
	err    error
}

func NewPushByUserIdResultFromJson(data string) PushByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPushByUserIdResultFromDict(dict)
}

func NewPushByUserIdResultFromDict(data map[string]interface{}) PushByUserIdResult {
	return PushByUserIdResult{
		Items: func() []Job {
			if data["items"] == nil {
				return nil
			}
			return CastJobs(core.CastArray(data["items"]))
		}(),
		AutoRun: func() *bool {
			v, ok := data["autoRun"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRun"])
		}(),
	}
}

func (p PushByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastJobsFromDict(
			p.Items,
		),
		"autoRun": p.AutoRun,
	}
}

func (p PushByUserIdResult) Pointer() *PushByUserIdResult {
	return &p
}

type RunResult struct {
	Item      *Job           `json:"item"`
	Result    *JobResultBody `json:"result"`
	IsLastJob *bool          `json:"isLastJob"`
}

type RunAsyncResult struct {
	result *RunResult
	err    error
}

func NewRunResultFromJson(data string) RunResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunResultFromDict(dict)
}

func NewRunResultFromDict(data map[string]interface{}) RunResult {
	return RunResult{
		Item: func() *Job {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJobFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Result: func() *JobResultBody {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return NewJobResultBodyFromDict(core.CastMap(data["result"])).Pointer()
		}(),
		IsLastJob: func() *bool {
			v, ok := data["isLastJob"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["isLastJob"])
		}(),
	}
}

func (p RunResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"result": func() map[string]interface{} {
			if p.Result == nil {
				return nil
			}
			return p.Result.ToDict()
		}(),
		"isLastJob": p.IsLastJob,
	}
}

func (p RunResult) Pointer() *RunResult {
	return &p
}

type RunByUserIdResult struct {
	Item      *Job           `json:"item"`
	Result    *JobResultBody `json:"result"`
	IsLastJob *bool          `json:"isLastJob"`
}

type RunByUserIdAsyncResult struct {
	result *RunByUserIdResult
	err    error
}

func NewRunByUserIdResultFromJson(data string) RunByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunByUserIdResultFromDict(dict)
}

func NewRunByUserIdResultFromDict(data map[string]interface{}) RunByUserIdResult {
	return RunByUserIdResult{
		Item: func() *Job {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJobFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Result: func() *JobResultBody {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return NewJobResultBodyFromDict(core.CastMap(data["result"])).Pointer()
		}(),
		IsLastJob: func() *bool {
			v, ok := data["isLastJob"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["isLastJob"])
		}(),
	}
}

func (p RunByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"result": func() map[string]interface{} {
			if p.Result == nil {
				return nil
			}
			return p.Result.ToDict()
		}(),
		"isLastJob": p.IsLastJob,
	}
}

func (p RunByUserIdResult) Pointer() *RunByUserIdResult {
	return &p
}

type DeleteJobResult struct {
	Item *Job `json:"item"`
}

type DeleteJobAsyncResult struct {
	result *DeleteJobResult
	err    error
}

func NewDeleteJobResultFromJson(data string) DeleteJobResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteJobResultFromDict(dict)
}

func NewDeleteJobResultFromDict(data map[string]interface{}) DeleteJobResult {
	return DeleteJobResult{
		Item: func() *Job {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJobFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteJobResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteJobResult) Pointer() *DeleteJobResult {
	return &p
}

type DeleteJobByUserIdResult struct {
	Item *Job `json:"item"`
}

type DeleteJobByUserIdAsyncResult struct {
	result *DeleteJobByUserIdResult
	err    error
}

func NewDeleteJobByUserIdResultFromJson(data string) DeleteJobByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteJobByUserIdResultFromDict(dict)
}

func NewDeleteJobByUserIdResultFromDict(data map[string]interface{}) DeleteJobByUserIdResult {
	return DeleteJobByUserIdResult{
		Item: func() *Job {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJobFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteJobByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteJobByUserIdResult) Pointer() *DeleteJobByUserIdResult {
	return &p
}

type PushByStampSheetResult struct {
	Items   []Job `json:"items"`
	AutoRun *bool `json:"autoRun"`
}

type PushByStampSheetAsyncResult struct {
	result *PushByStampSheetResult
	err    error
}

func NewPushByStampSheetResultFromJson(data string) PushByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPushByStampSheetResultFromDict(dict)
}

func NewPushByStampSheetResultFromDict(data map[string]interface{}) PushByStampSheetResult {
	return PushByStampSheetResult{
		Items: func() []Job {
			if data["items"] == nil {
				return nil
			}
			return CastJobs(core.CastArray(data["items"]))
		}(),
		AutoRun: func() *bool {
			v, ok := data["autoRun"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRun"])
		}(),
	}
}

func (p PushByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastJobsFromDict(
			p.Items,
		),
		"autoRun": p.AutoRun,
	}
}

func (p PushByStampSheetResult) Pointer() *PushByStampSheetResult {
	return &p
}

type DeleteByStampTaskResult struct {
	Item            *Job    `json:"item"`
	NewContextStack *string `json:"newContextStack"`
}

type DeleteByStampTaskAsyncResult struct {
	result *DeleteByStampTaskResult
	err    error
}

func NewDeleteByStampTaskResultFromJson(data string) DeleteByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteByStampTaskResultFromDict(dict)
}

func NewDeleteByStampTaskResultFromDict(data map[string]interface{}) DeleteByStampTaskResult {
	return DeleteByStampTaskResult{
		Item: func() *Job {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJobFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteByStampTaskResult) ToDict() map[string]interface{} {
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

func (p DeleteByStampTaskResult) Pointer() *DeleteByStampTaskResult {
	return &p
}

type GetJobResultResult struct {
	Item *JobResult `json:"item"`
}

type GetJobResultAsyncResult struct {
	result *GetJobResultResult
	err    error
}

func NewGetJobResultResultFromJson(data string) GetJobResultResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJobResultResultFromDict(dict)
}

func NewGetJobResultResultFromDict(data map[string]interface{}) GetJobResultResult {
	return GetJobResultResult{
		Item: func() *JobResult {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJobResultFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetJobResultResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetJobResultResult) Pointer() *GetJobResultResult {
	return &p
}

type GetJobResultByUserIdResult struct {
	Item *JobResult `json:"item"`
}

type GetJobResultByUserIdAsyncResult struct {
	result *GetJobResultByUserIdResult
	err    error
}

func NewGetJobResultByUserIdResultFromJson(data string) GetJobResultByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJobResultByUserIdResultFromDict(dict)
}

func NewGetJobResultByUserIdResultFromDict(data map[string]interface{}) GetJobResultByUserIdResult {
	return GetJobResultByUserIdResult{
		Item: func() *JobResult {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJobResultFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetJobResultByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetJobResultByUserIdResult) Pointer() *GetJobResultByUserIdResult {
	return &p
}

type DescribeDeadLetterJobsByUserIdResult struct {
	Items         []DeadLetterJob `json:"items"`
	NextPageToken *string         `json:"nextPageToken"`
}

type DescribeDeadLetterJobsByUserIdAsyncResult struct {
	result *DescribeDeadLetterJobsByUserIdResult
	err    error
}

func NewDescribeDeadLetterJobsByUserIdResultFromJson(data string) DescribeDeadLetterJobsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDeadLetterJobsByUserIdResultFromDict(dict)
}

func NewDescribeDeadLetterJobsByUserIdResultFromDict(data map[string]interface{}) DescribeDeadLetterJobsByUserIdResult {
	return DescribeDeadLetterJobsByUserIdResult{
		Items: func() []DeadLetterJob {
			if data["items"] == nil {
				return nil
			}
			return CastDeadLetterJobs(core.CastArray(data["items"]))
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

func (p DescribeDeadLetterJobsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDeadLetterJobsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeDeadLetterJobsByUserIdResult) Pointer() *DescribeDeadLetterJobsByUserIdResult {
	return &p
}

type GetDeadLetterJobByUserIdResult struct {
	Item *DeadLetterJob `json:"item"`
}

type GetDeadLetterJobByUserIdAsyncResult struct {
	result *GetDeadLetterJobByUserIdResult
	err    error
}

func NewGetDeadLetterJobByUserIdResultFromJson(data string) GetDeadLetterJobByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDeadLetterJobByUserIdResultFromDict(dict)
}

func NewGetDeadLetterJobByUserIdResultFromDict(data map[string]interface{}) GetDeadLetterJobByUserIdResult {
	return GetDeadLetterJobByUserIdResult{
		Item: func() *DeadLetterJob {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDeadLetterJobFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetDeadLetterJobByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetDeadLetterJobByUserIdResult) Pointer() *GetDeadLetterJobByUserIdResult {
	return &p
}

type DeleteDeadLetterJobByUserIdResult struct {
	Item *DeadLetterJob `json:"item"`
}

type DeleteDeadLetterJobByUserIdAsyncResult struct {
	result *DeleteDeadLetterJobByUserIdResult
	err    error
}

func NewDeleteDeadLetterJobByUserIdResultFromJson(data string) DeleteDeadLetterJobByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteDeadLetterJobByUserIdResultFromDict(dict)
}

func NewDeleteDeadLetterJobByUserIdResultFromDict(data map[string]interface{}) DeleteDeadLetterJobByUserIdResult {
	return DeleteDeadLetterJobByUserIdResult{
		Item: func() *DeadLetterJob {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDeadLetterJobFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteDeadLetterJobByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteDeadLetterJobByUserIdResult) Pointer() *DeleteDeadLetterJobByUserIdResult {
	return &p
}
