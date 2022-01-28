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
    Items []Namespace `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
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

func NewCreateNamespaceResultFromJson(data string) CreateNamespaceResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateNamespaceResultFromDict(dict)
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

func NewGetNamespaceStatusResultFromJson(data string) GetNamespaceStatusResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceStatusResultFromDict(dict)
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

func NewGetNamespaceResultFromJson(data string) GetNamespaceResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceResultFromDict(dict)
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

func NewUpdateNamespaceResultFromJson(data string) UpdateNamespaceResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateNamespaceResultFromDict(dict)
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

func NewDeleteNamespaceResultFromJson(data string) DeleteNamespaceResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteNamespaceResultFromDict(dict)
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

type DescribeJobsByUserIdResult struct {
    Items []Job `json:"items"`
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
    return DescribeJobsByUserIdResult {
        Items: CastJobs(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeJobsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return GetJobByUserIdResult {
        Item: NewJobFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetJobByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetJobByUserIdResult) Pointer() *GetJobByUserIdResult {
    return &p
}

type PushByUserIdResult struct {
    Items []Job `json:"items"`
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
    return PushByUserIdResult {
        Items: CastJobs(core.CastArray(data["items"])),
    }
}

func (p PushByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastJobsFromDict(
            p.Items,
        ),
    }
}

func (p PushByUserIdResult) Pointer() *PushByUserIdResult {
    return &p
}

type RunResult struct {
    Item *Job `json:"item"`
    Result *JobResultBody `json:"result"`
    IsLastJob *bool `json:"isLastJob"`
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
    return RunResult {
        Item: NewJobFromDict(core.CastMap(data["item"])).Pointer(),
        Result: NewJobResultBodyFromDict(core.CastMap(data["result"])).Pointer(),
        IsLastJob: core.CastBool(data["isLastJob"]),
    }
}

func (p RunResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "result": p.Result.ToDict(),
        "isLastJob": p.IsLastJob,
    }
}

func (p RunResult) Pointer() *RunResult {
    return &p
}

type RunByUserIdResult struct {
    Item *Job `json:"item"`
    Result *JobResultBody `json:"result"`
    IsLastJob *bool `json:"isLastJob"`
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
    return RunByUserIdResult {
        Item: NewJobFromDict(core.CastMap(data["item"])).Pointer(),
        Result: NewJobResultBodyFromDict(core.CastMap(data["result"])).Pointer(),
        IsLastJob: core.CastBool(data["isLastJob"]),
    }
}

func (p RunByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "result": p.Result.ToDict(),
        "isLastJob": p.IsLastJob,
    }
}

func (p RunByUserIdResult) Pointer() *RunByUserIdResult {
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
    return DeleteJobByUserIdResult {
        Item: NewJobFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteJobByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteJobByUserIdResult) Pointer() *DeleteJobByUserIdResult {
    return &p
}

type PushByStampSheetResult struct {
    Items []Job `json:"items"`
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
    return PushByStampSheetResult {
        Items: CastJobs(core.CastArray(data["items"])),
    }
}

func (p PushByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastJobsFromDict(
            p.Items,
        ),
    }
}

func (p PushByStampSheetResult) Pointer() *PushByStampSheetResult {
    return &p
}

type DescribeDeadLetterJobsByUserIdResult struct {
    Items []DeadLetterJob `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
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
    return DescribeDeadLetterJobsByUserIdResult {
        Items: CastDeadLetterJobs(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeDeadLetterJobsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return GetDeadLetterJobByUserIdResult {
        Item: NewDeadLetterJobFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetDeadLetterJobByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
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
    return DeleteDeadLetterJobByUserIdResult {
        Item: NewDeadLetterJobFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteDeadLetterJobByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteDeadLetterJobByUserIdResult) Pointer() *DeleteDeadLetterJobByUserIdResult {
    return &p
}