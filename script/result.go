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
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
    return DeleteNamespaceResult {
    }
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
    return &p
}

type DescribeScriptsResult struct {
    Items []Script `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeScriptsAsyncResult struct {
	result *DescribeScriptsResult
	err    error
}

func NewDescribeScriptsResultFromDict(data map[string]interface{}) DescribeScriptsResult {
    return DescribeScriptsResult {
        Items: CastScripts(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeScriptsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastScriptsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeScriptsResult) Pointer() *DescribeScriptsResult {
    return &p
}

type CreateScriptResult struct {
    Item *Script `json:"item"`
}

type CreateScriptAsyncResult struct {
	result *CreateScriptResult
	err    error
}

func NewCreateScriptResultFromDict(data map[string]interface{}) CreateScriptResult {
    return CreateScriptResult {
        Item: NewScriptFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateScriptResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateScriptResult) Pointer() *CreateScriptResult {
    return &p
}

type CreateScriptFromGitHubResult struct {
    Item *Script `json:"item"`
}

type CreateScriptFromGitHubAsyncResult struct {
	result *CreateScriptFromGitHubResult
	err    error
}

func NewCreateScriptFromGitHubResultFromDict(data map[string]interface{}) CreateScriptFromGitHubResult {
    return CreateScriptFromGitHubResult {
        Item: NewScriptFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateScriptFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateScriptFromGitHubResult) Pointer() *CreateScriptFromGitHubResult {
    return &p
}

type GetScriptResult struct {
    Item *Script `json:"item"`
}

type GetScriptAsyncResult struct {
	result *GetScriptResult
	err    error
}

func NewGetScriptResultFromDict(data map[string]interface{}) GetScriptResult {
    return GetScriptResult {
        Item: NewScriptFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetScriptResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetScriptResult) Pointer() *GetScriptResult {
    return &p
}

type UpdateScriptResult struct {
    Item *Script `json:"item"`
}

type UpdateScriptAsyncResult struct {
	result *UpdateScriptResult
	err    error
}

func NewUpdateScriptResultFromDict(data map[string]interface{}) UpdateScriptResult {
    return UpdateScriptResult {
        Item: NewScriptFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateScriptResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateScriptResult) Pointer() *UpdateScriptResult {
    return &p
}

type UpdateScriptFromGitHubResult struct {
    Item *Script `json:"item"`
}

type UpdateScriptFromGitHubAsyncResult struct {
	result *UpdateScriptFromGitHubResult
	err    error
}

func NewUpdateScriptFromGitHubResultFromDict(data map[string]interface{}) UpdateScriptFromGitHubResult {
    return UpdateScriptFromGitHubResult {
        Item: NewScriptFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateScriptFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateScriptFromGitHubResult) Pointer() *UpdateScriptFromGitHubResult {
    return &p
}

type DeleteScriptResult struct {
}

type DeleteScriptAsyncResult struct {
	result *DeleteScriptResult
	err    error
}

func NewDeleteScriptResultFromDict(data map[string]interface{}) DeleteScriptResult {
    return DeleteScriptResult {
    }
}

func (p DeleteScriptResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteScriptResult) Pointer() *DeleteScriptResult {
    return &p
}

type InvokeScriptResult struct {
    Code *int32 `json:"code"`
    Result *string `json:"result"`
    ExecuteTime *int32 `json:"executeTime"`
    Charged *int32 `json:"charged"`
    Output []string `json:"output"`
}

type InvokeScriptAsyncResult struct {
	result *InvokeScriptResult
	err    error
}

func NewInvokeScriptResultFromDict(data map[string]interface{}) InvokeScriptResult {
    return InvokeScriptResult {
        Code: core.CastInt32(data["code"]),
        Result: core.CastString(data["result"]),
        ExecuteTime: core.CastInt32(data["executeTime"]),
        Charged: core.CastInt32(data["charged"]),
        Output: core.CastStrings(core.CastArray(data["output"])),
    }
}

func (p InvokeScriptResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "code": p.Code,
        "result": p.Result,
        "executeTime": p.ExecuteTime,
        "charged": p.Charged,
        "output": core.CastStringsFromDict(
            p.Output,
        ),
    }
}

func (p InvokeScriptResult) Pointer() *InvokeScriptResult {
    return &p
}

type DebugInvokeResult struct {
    Code *int32 `json:"code"`
    Result *string `json:"result"`
    ExecuteTime *int32 `json:"executeTime"`
    Charged *int32 `json:"charged"`
    Output []string `json:"output"`
}

type DebugInvokeAsyncResult struct {
	result *DebugInvokeResult
	err    error
}

func NewDebugInvokeResultFromDict(data map[string]interface{}) DebugInvokeResult {
    return DebugInvokeResult {
        Code: core.CastInt32(data["code"]),
        Result: core.CastString(data["result"]),
        ExecuteTime: core.CastInt32(data["executeTime"]),
        Charged: core.CastInt32(data["charged"]),
        Output: core.CastStrings(core.CastArray(data["output"])),
    }
}

func (p DebugInvokeResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "code": p.Code,
        "result": p.Result,
        "executeTime": p.ExecuteTime,
        "charged": p.Charged,
        "output": core.CastStringsFromDict(
            p.Output,
        ),
    }
}

func (p DebugInvokeResult) Pointer() *DebugInvokeResult {
    return &p
}