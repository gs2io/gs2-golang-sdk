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

package lock

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
    Item *Namespace `json:"item"`
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
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

type DescribeMutexesResult struct {
    Items []Mutex `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeMutexesAsyncResult struct {
	result *DescribeMutexesResult
	err    error
}

func NewDescribeMutexesResultFromDict(data map[string]interface{}) DescribeMutexesResult {
    return DescribeMutexesResult {
        Items: CastMutexes(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeMutexesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastMutexesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeMutexesResult) Pointer() *DescribeMutexesResult {
    return &p
}

type DescribeMutexesByUserIdResult struct {
    Items []Mutex `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeMutexesByUserIdAsyncResult struct {
	result *DescribeMutexesByUserIdResult
	err    error
}

func NewDescribeMutexesByUserIdResultFromDict(data map[string]interface{}) DescribeMutexesByUserIdResult {
    return DescribeMutexesByUserIdResult {
        Items: CastMutexes(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeMutexesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastMutexesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeMutexesByUserIdResult) Pointer() *DescribeMutexesByUserIdResult {
    return &p
}

type LockResult struct {
    Item *Mutex `json:"item"`
}

type LockAsyncResult struct {
	result *LockResult
	err    error
}

func NewLockResultFromDict(data map[string]interface{}) LockResult {
    return LockResult {
        Item: NewMutexFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p LockResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p LockResult) Pointer() *LockResult {
    return &p
}

type LockByUserIdResult struct {
    Item *Mutex `json:"item"`
}

type LockByUserIdAsyncResult struct {
	result *LockByUserIdResult
	err    error
}

func NewLockByUserIdResultFromDict(data map[string]interface{}) LockByUserIdResult {
    return LockByUserIdResult {
        Item: NewMutexFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p LockByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p LockByUserIdResult) Pointer() *LockByUserIdResult {
    return &p
}

type UnlockResult struct {
    Item *Mutex `json:"item"`
}

type UnlockAsyncResult struct {
	result *UnlockResult
	err    error
}

func NewUnlockResultFromDict(data map[string]interface{}) UnlockResult {
    return UnlockResult {
        Item: NewMutexFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UnlockResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UnlockResult) Pointer() *UnlockResult {
    return &p
}

type UnlockByUserIdResult struct {
    Item *Mutex `json:"item"`
}

type UnlockByUserIdAsyncResult struct {
	result *UnlockByUserIdResult
	err    error
}

func NewUnlockByUserIdResultFromDict(data map[string]interface{}) UnlockByUserIdResult {
    return UnlockByUserIdResult {
        Item: NewMutexFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UnlockByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UnlockByUserIdResult) Pointer() *UnlockByUserIdResult {
    return &p
}

type GetMutexResult struct {
    Item *Mutex `json:"item"`
}

type GetMutexAsyncResult struct {
	result *GetMutexResult
	err    error
}

func NewGetMutexResultFromDict(data map[string]interface{}) GetMutexResult {
    return GetMutexResult {
        Item: NewMutexFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetMutexResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetMutexResult) Pointer() *GetMutexResult {
    return &p
}

type GetMutexByUserIdResult struct {
    Item *Mutex `json:"item"`
}

type GetMutexByUserIdAsyncResult struct {
	result *GetMutexByUserIdResult
	err    error
}

func NewGetMutexByUserIdResultFromDict(data map[string]interface{}) GetMutexByUserIdResult {
    return GetMutexByUserIdResult {
        Item: NewMutexFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetMutexByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetMutexByUserIdResult) Pointer() *GetMutexByUserIdResult {
    return &p
}

type DeleteMutexByUserIdResult struct {
    Item *Mutex `json:"item"`
}

type DeleteMutexByUserIdAsyncResult struct {
	result *DeleteMutexByUserIdResult
	err    error
}

func NewDeleteMutexByUserIdResultFromDict(data map[string]interface{}) DeleteMutexByUserIdResult {
    return DeleteMutexByUserIdResult {
        Item: NewMutexFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteMutexByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteMutexByUserIdResult) Pointer() *DeleteMutexByUserIdResult {
    return &p
}