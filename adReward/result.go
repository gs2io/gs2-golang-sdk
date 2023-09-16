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

package adReward

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
	return UpdateNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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
	return DeleteNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
	return &p
}

type GetPointResult struct {
	Item *Point `json:"item"`
}

type GetPointAsyncResult struct {
	result *GetPointResult
	err    error
}

func NewGetPointResultFromJson(data string) GetPointResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPointResultFromDict(dict)
}

func NewGetPointResultFromDict(data map[string]interface{}) GetPointResult {
	return GetPointResult{
		Item: NewPointFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetPointResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetPointResult) Pointer() *GetPointResult {
	return &p
}

type GetPointByUserIdResult struct {
	Item *Point `json:"item"`
}

type GetPointByUserIdAsyncResult struct {
	result *GetPointByUserIdResult
	err    error
}

func NewGetPointByUserIdResultFromJson(data string) GetPointByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPointByUserIdResultFromDict(dict)
}

func NewGetPointByUserIdResultFromDict(data map[string]interface{}) GetPointByUserIdResult {
	return GetPointByUserIdResult{
		Item: NewPointFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetPointByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetPointByUserIdResult) Pointer() *GetPointByUserIdResult {
	return &p
}

type AcquirePointByUserIdResult struct {
	Item *Point `json:"item"`
}

type AcquirePointByUserIdAsyncResult struct {
	result *AcquirePointByUserIdResult
	err    error
}

func NewAcquirePointByUserIdResultFromJson(data string) AcquirePointByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquirePointByUserIdResultFromDict(dict)
}

func NewAcquirePointByUserIdResultFromDict(data map[string]interface{}) AcquirePointByUserIdResult {
	return AcquirePointByUserIdResult{
		Item: NewPointFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p AcquirePointByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p AcquirePointByUserIdResult) Pointer() *AcquirePointByUserIdResult {
	return &p
}

type ConsumePointResult struct {
	Item *Point `json:"item"`
}

type ConsumePointAsyncResult struct {
	result *ConsumePointResult
	err    error
}

func NewConsumePointResultFromJson(data string) ConsumePointResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumePointResultFromDict(dict)
}

func NewConsumePointResultFromDict(data map[string]interface{}) ConsumePointResult {
	return ConsumePointResult{
		Item: NewPointFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ConsumePointResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p ConsumePointResult) Pointer() *ConsumePointResult {
	return &p
}

type ConsumePointByUserIdResult struct {
	Item *Point `json:"item"`
}

type ConsumePointByUserIdAsyncResult struct {
	result *ConsumePointByUserIdResult
	err    error
}

func NewConsumePointByUserIdResultFromJson(data string) ConsumePointByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumePointByUserIdResultFromDict(dict)
}

func NewConsumePointByUserIdResultFromDict(data map[string]interface{}) ConsumePointByUserIdResult {
	return ConsumePointByUserIdResult{
		Item: NewPointFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ConsumePointByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p ConsumePointByUserIdResult) Pointer() *ConsumePointByUserIdResult {
	return &p
}

type DeletePointByUserIdResult struct {
	Item *Point `json:"item"`
}

type DeletePointByUserIdAsyncResult struct {
	result *DeletePointByUserIdResult
	err    error
}

func NewDeletePointByUserIdResultFromJson(data string) DeletePointByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePointByUserIdResultFromDict(dict)
}

func NewDeletePointByUserIdResultFromDict(data map[string]interface{}) DeletePointByUserIdResult {
	return DeletePointByUserIdResult{
		Item: NewPointFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeletePointByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeletePointByUserIdResult) Pointer() *DeletePointByUserIdResult {
	return &p
}

type ConsumePointByStampTaskResult struct {
	Item            *Point  `json:"item"`
	NewContextStack *string `json:"newContextStack"`
}

type ConsumePointByStampTaskAsyncResult struct {
	result *ConsumePointByStampTaskResult
	err    error
}

func NewConsumePointByStampTaskResultFromJson(data string) ConsumePointByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumePointByStampTaskResultFromDict(dict)
}

func NewConsumePointByStampTaskResultFromDict(data map[string]interface{}) ConsumePointByStampTaskResult {
	return ConsumePointByStampTaskResult{
		Item:            NewPointFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p ConsumePointByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":            p.Item.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p ConsumePointByStampTaskResult) Pointer() *ConsumePointByStampTaskResult {
	return &p
}

type AcquirePointByStampSheetResult struct {
	Item *Point `json:"item"`
}

type AcquirePointByStampSheetAsyncResult struct {
	result *AcquirePointByStampSheetResult
	err    error
}

func NewAcquirePointByStampSheetResultFromJson(data string) AcquirePointByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquirePointByStampSheetResultFromDict(dict)
}

func NewAcquirePointByStampSheetResultFromDict(data map[string]interface{}) AcquirePointByStampSheetResult {
	return AcquirePointByStampSheetResult{
		Item: NewPointFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p AcquirePointByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p AcquirePointByStampSheetResult) Pointer() *AcquirePointByStampSheetResult {
	return &p
}
