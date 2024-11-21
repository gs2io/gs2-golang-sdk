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

type LockResult struct {
	Item     *Mutex               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type LockAsyncResult struct {
	result *LockResult
	err    error
}

func NewLockResultFromJson(data string) LockResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLockResultFromDict(dict)
}

func NewLockResultFromDict(data map[string]interface{}) LockResult {
	return LockResult{
		Item: func() *Mutex {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMutexFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p LockResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p LockResult) Pointer() *LockResult {
	return &p
}

type LockByUserIdResult struct {
	Item     *Mutex               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type LockByUserIdAsyncResult struct {
	result *LockByUserIdResult
	err    error
}

func NewLockByUserIdResultFromJson(data string) LockByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLockByUserIdResultFromDict(dict)
}

func NewLockByUserIdResultFromDict(data map[string]interface{}) LockByUserIdResult {
	return LockByUserIdResult{
		Item: func() *Mutex {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMutexFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p LockByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p LockByUserIdResult) Pointer() *LockByUserIdResult {
	return &p
}

type UnlockResult struct {
	Item     *Mutex               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UnlockAsyncResult struct {
	result *UnlockResult
	err    error
}

func NewUnlockResultFromJson(data string) UnlockResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnlockResultFromDict(dict)
}

func NewUnlockResultFromDict(data map[string]interface{}) UnlockResult {
	return UnlockResult{
		Item: func() *Mutex {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMutexFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UnlockResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UnlockResult) Pointer() *UnlockResult {
	return &p
}

type UnlockByUserIdResult struct {
	Item     *Mutex               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UnlockByUserIdAsyncResult struct {
	result *UnlockByUserIdResult
	err    error
}

func NewUnlockByUserIdResultFromJson(data string) UnlockByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnlockByUserIdResultFromDict(dict)
}

func NewUnlockByUserIdResultFromDict(data map[string]interface{}) UnlockByUserIdResult {
	return UnlockByUserIdResult{
		Item: func() *Mutex {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMutexFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UnlockByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UnlockByUserIdResult) Pointer() *UnlockByUserIdResult {
	return &p
}

type GetMutexResult struct {
	Item     *Mutex               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetMutexAsyncResult struct {
	result *GetMutexResult
	err    error
}

func NewGetMutexResultFromJson(data string) GetMutexResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMutexResultFromDict(dict)
}

func NewGetMutexResultFromDict(data map[string]interface{}) GetMutexResult {
	return GetMutexResult{
		Item: func() *Mutex {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMutexFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetMutexResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetMutexResult) Pointer() *GetMutexResult {
	return &p
}

type GetMutexByUserIdResult struct {
	Item     *Mutex               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetMutexByUserIdAsyncResult struct {
	result *GetMutexByUserIdResult
	err    error
}

func NewGetMutexByUserIdResultFromJson(data string) GetMutexByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMutexByUserIdResultFromDict(dict)
}

func NewGetMutexByUserIdResultFromDict(data map[string]interface{}) GetMutexByUserIdResult {
	return GetMutexByUserIdResult{
		Item: func() *Mutex {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMutexFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetMutexByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetMutexByUserIdResult) Pointer() *GetMutexByUserIdResult {
	return &p
}

type DeleteMutexByUserIdResult struct {
	Item     *Mutex               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteMutexByUserIdAsyncResult struct {
	result *DeleteMutexByUserIdResult
	err    error
}

func NewDeleteMutexByUserIdResultFromJson(data string) DeleteMutexByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMutexByUserIdResultFromDict(dict)
}

func NewDeleteMutexByUserIdResultFromDict(data map[string]interface{}) DeleteMutexByUserIdResult {
	return DeleteMutexByUserIdResult{
		Item: func() *Mutex {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMutexFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteMutexByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteMutexByUserIdResult) Pointer() *DeleteMutexByUserIdResult {
	return &p
}
