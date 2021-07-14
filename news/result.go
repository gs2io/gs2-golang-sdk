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

type PrepareUpdateCurrentNewsMasterResult struct {
    UploadToken *string `json:"uploadToken"`
    TemplateUploadUrl *string `json:"templateUploadUrl"`
}

type PrepareUpdateCurrentNewsMasterAsyncResult struct {
	result *PrepareUpdateCurrentNewsMasterResult
	err    error
}

func NewPrepareUpdateCurrentNewsMasterResultFromDict(data map[string]interface{}) PrepareUpdateCurrentNewsMasterResult {
    return PrepareUpdateCurrentNewsMasterResult {
        UploadToken: core.CastString(data["uploadToken"]),
        TemplateUploadUrl: core.CastString(data["templateUploadUrl"]),
    }
}

func (p PrepareUpdateCurrentNewsMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "uploadToken": p.UploadToken,
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

func NewUpdateCurrentNewsMasterResultFromDict(data map[string]interface{}) UpdateCurrentNewsMasterResult {
    return UpdateCurrentNewsMasterResult {
    }
}

func (p UpdateCurrentNewsMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
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

func NewPrepareUpdateCurrentNewsMasterFromGitHubResultFromDict(data map[string]interface{}) PrepareUpdateCurrentNewsMasterFromGitHubResult {
    return PrepareUpdateCurrentNewsMasterFromGitHubResult {
        UploadToken: core.CastString(data["uploadToken"]),
    }
}

func (p PrepareUpdateCurrentNewsMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "uploadToken": p.UploadToken,
    }
}

func (p PrepareUpdateCurrentNewsMasterFromGitHubResult) Pointer() *PrepareUpdateCurrentNewsMasterFromGitHubResult {
    return &p
}

type DescribeNewsResult struct {
    Items []News `json:"items"`
    ContentHash *string `json:"contentHash"`
    TemplateHash *string `json:"templateHash"`
}

type DescribeNewsAsyncResult struct {
	result *DescribeNewsResult
	err    error
}

func NewDescribeNewsResultFromDict(data map[string]interface{}) DescribeNewsResult {
    return DescribeNewsResult {
        Items: CastNewses(core.CastArray(data["items"])),
        ContentHash: core.CastString(data["contentHash"]),
        TemplateHash: core.CastString(data["templateHash"]),
    }
}

func (p DescribeNewsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastNewsesFromDict(
            p.Items,
        ),
        "contentHash": p.ContentHash,
        "templateHash": p.TemplateHash,
    }
}

func (p DescribeNewsResult) Pointer() *DescribeNewsResult {
    return &p
}

type DescribeNewsByUserIdResult struct {
    Items []News `json:"items"`
    ContentHash *string `json:"contentHash"`
    TemplateHash *string `json:"templateHash"`
}

type DescribeNewsByUserIdAsyncResult struct {
	result *DescribeNewsByUserIdResult
	err    error
}

func NewDescribeNewsByUserIdResultFromDict(data map[string]interface{}) DescribeNewsByUserIdResult {
    return DescribeNewsByUserIdResult {
        Items: CastNewses(core.CastArray(data["items"])),
        ContentHash: core.CastString(data["contentHash"]),
        TemplateHash: core.CastString(data["templateHash"]),
    }
}

func (p DescribeNewsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastNewsesFromDict(
            p.Items,
        ),
        "contentHash": p.ContentHash,
        "templateHash": p.TemplateHash,
    }
}

func (p DescribeNewsByUserIdResult) Pointer() *DescribeNewsByUserIdResult {
    return &p
}

type WantGrantResult struct {
    Items []SetCookieRequestEntry `json:"items"`
    BrowserUrl *string `json:"browserUrl"`
    ZipUrl *string `json:"zipUrl"`
}

type WantGrantAsyncResult struct {
	result *WantGrantResult
	err    error
}

func NewWantGrantResultFromDict(data map[string]interface{}) WantGrantResult {
    return WantGrantResult {
        Items: CastSetCookieRequestEntries(core.CastArray(data["items"])),
        BrowserUrl: core.CastString(data["browserUrl"]),
        ZipUrl: core.CastString(data["zipUrl"]),
    }
}

func (p WantGrantResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSetCookieRequestEntriesFromDict(
            p.Items,
        ),
        "browserUrl": p.BrowserUrl,
        "zipUrl": p.ZipUrl,
    }
}

func (p WantGrantResult) Pointer() *WantGrantResult {
    return &p
}

type WantGrantByUserIdResult struct {
    Items []SetCookieRequestEntry `json:"items"`
    BrowserUrl *string `json:"browserUrl"`
    ZipUrl *string `json:"zipUrl"`
}

type WantGrantByUserIdAsyncResult struct {
	result *WantGrantByUserIdResult
	err    error
}

func NewWantGrantByUserIdResultFromDict(data map[string]interface{}) WantGrantByUserIdResult {
    return WantGrantByUserIdResult {
        Items: CastSetCookieRequestEntries(core.CastArray(data["items"])),
        BrowserUrl: core.CastString(data["browserUrl"]),
        ZipUrl: core.CastString(data["zipUrl"]),
    }
}

func (p WantGrantByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSetCookieRequestEntriesFromDict(
            p.Items,
        ),
        "browserUrl": p.BrowserUrl,
        "zipUrl": p.ZipUrl,
    }
}

func (p WantGrantByUserIdResult) Pointer() *WantGrantByUserIdResult {
    return &p
}