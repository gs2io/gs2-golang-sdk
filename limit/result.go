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

package limit

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

type DescribeCountersResult struct {
	Items         []Counter `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeCountersAsyncResult struct {
	result *DescribeCountersResult
	err    error
}

func NewDescribeCountersResultFromJson(data string) DescribeCountersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCountersResultFromDict(dict)
}

func NewDescribeCountersResultFromDict(data map[string]interface{}) DescribeCountersResult {
	return DescribeCountersResult{
		Items: func() []Counter {
			if data["items"] == nil {
				return nil
			}
			return CastCounters(core.CastArray(data["items"]))
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

func (p DescribeCountersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCountersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCountersResult) Pointer() *DescribeCountersResult {
	return &p
}

type DescribeCountersByUserIdResult struct {
	Items         []Counter `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeCountersByUserIdAsyncResult struct {
	result *DescribeCountersByUserIdResult
	err    error
}

func NewDescribeCountersByUserIdResultFromJson(data string) DescribeCountersByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCountersByUserIdResultFromDict(dict)
}

func NewDescribeCountersByUserIdResultFromDict(data map[string]interface{}) DescribeCountersByUserIdResult {
	return DescribeCountersByUserIdResult{
		Items: func() []Counter {
			if data["items"] == nil {
				return nil
			}
			return CastCounters(core.CastArray(data["items"]))
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

func (p DescribeCountersByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCountersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCountersByUserIdResult) Pointer() *DescribeCountersByUserIdResult {
	return &p
}

type GetCounterResult struct {
	Item *Counter `json:"item"`
}

type GetCounterAsyncResult struct {
	result *GetCounterResult
	err    error
}

func NewGetCounterResultFromJson(data string) GetCounterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterResultFromDict(dict)
}

func NewGetCounterResultFromDict(data map[string]interface{}) GetCounterResult {
	return GetCounterResult{
		Item: func() *Counter {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCounterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCounterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCounterResult) Pointer() *GetCounterResult {
	return &p
}

type GetCounterByUserIdResult struct {
	Item *Counter `json:"item"`
}

type GetCounterByUserIdAsyncResult struct {
	result *GetCounterByUserIdResult
	err    error
}

func NewGetCounterByUserIdResultFromJson(data string) GetCounterByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterByUserIdResultFromDict(dict)
}

func NewGetCounterByUserIdResultFromDict(data map[string]interface{}) GetCounterByUserIdResult {
	return GetCounterByUserIdResult{
		Item: func() *Counter {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCounterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCounterByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCounterByUserIdResult) Pointer() *GetCounterByUserIdResult {
	return &p
}

type CountUpResult struct {
	Item *Counter `json:"item"`
}

type CountUpAsyncResult struct {
	result *CountUpResult
	err    error
}

func NewCountUpResultFromJson(data string) CountUpResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountUpResultFromDict(dict)
}

func NewCountUpResultFromDict(data map[string]interface{}) CountUpResult {
	return CountUpResult{
		Item: func() *Counter {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCounterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CountUpResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CountUpResult) Pointer() *CountUpResult {
	return &p
}

type CountUpByUserIdResult struct {
	Item *Counter `json:"item"`
}

type CountUpByUserIdAsyncResult struct {
	result *CountUpByUserIdResult
	err    error
}

func NewCountUpByUserIdResultFromJson(data string) CountUpByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountUpByUserIdResultFromDict(dict)
}

func NewCountUpByUserIdResultFromDict(data map[string]interface{}) CountUpByUserIdResult {
	return CountUpByUserIdResult{
		Item: func() *Counter {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCounterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CountUpByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CountUpByUserIdResult) Pointer() *CountUpByUserIdResult {
	return &p
}

type CountDownByUserIdResult struct {
	Item *Counter `json:"item"`
}

type CountDownByUserIdAsyncResult struct {
	result *CountDownByUserIdResult
	err    error
}

func NewCountDownByUserIdResultFromJson(data string) CountDownByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountDownByUserIdResultFromDict(dict)
}

func NewCountDownByUserIdResultFromDict(data map[string]interface{}) CountDownByUserIdResult {
	return CountDownByUserIdResult{
		Item: func() *Counter {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCounterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CountDownByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CountDownByUserIdResult) Pointer() *CountDownByUserIdResult {
	return &p
}

type DeleteCounterByUserIdResult struct {
	Item *Counter `json:"item"`
}

type DeleteCounterByUserIdAsyncResult struct {
	result *DeleteCounterByUserIdResult
	err    error
}

func NewDeleteCounterByUserIdResultFromJson(data string) DeleteCounterByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCounterByUserIdResultFromDict(dict)
}

func NewDeleteCounterByUserIdResultFromDict(data map[string]interface{}) DeleteCounterByUserIdResult {
	return DeleteCounterByUserIdResult{
		Item: func() *Counter {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCounterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteCounterByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteCounterByUserIdResult) Pointer() *DeleteCounterByUserIdResult {
	return &p
}

type VerifyCounterResult struct {
}

type VerifyCounterAsyncResult struct {
	result *VerifyCounterResult
	err    error
}

func NewVerifyCounterResultFromJson(data string) VerifyCounterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyCounterResultFromDict(dict)
}

func NewVerifyCounterResultFromDict(data map[string]interface{}) VerifyCounterResult {
	return VerifyCounterResult{}
}

func (p VerifyCounterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyCounterResult) Pointer() *VerifyCounterResult {
	return &p
}

type VerifyCounterByUserIdResult struct {
}

type VerifyCounterByUserIdAsyncResult struct {
	result *VerifyCounterByUserIdResult
	err    error
}

func NewVerifyCounterByUserIdResultFromJson(data string) VerifyCounterByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyCounterByUserIdResultFromDict(dict)
}

func NewVerifyCounterByUserIdResultFromDict(data map[string]interface{}) VerifyCounterByUserIdResult {
	return VerifyCounterByUserIdResult{}
}

func (p VerifyCounterByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p VerifyCounterByUserIdResult) Pointer() *VerifyCounterByUserIdResult {
	return &p
}

type CountUpByStampTaskResult struct {
	Item            *Counter `json:"item"`
	NewContextStack *string  `json:"newContextStack"`
}

type CountUpByStampTaskAsyncResult struct {
	result *CountUpByStampTaskResult
	err    error
}

func NewCountUpByStampTaskResultFromJson(data string) CountUpByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountUpByStampTaskResultFromDict(dict)
}

func NewCountUpByStampTaskResultFromDict(data map[string]interface{}) CountUpByStampTaskResult {
	return CountUpByStampTaskResult{
		Item: func() *Counter {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCounterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CountUpByStampTaskResult) ToDict() map[string]interface{} {
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

func (p CountUpByStampTaskResult) Pointer() *CountUpByStampTaskResult {
	return &p
}

type CountDownByStampSheetResult struct {
	Item *Counter `json:"item"`
}

type CountDownByStampSheetAsyncResult struct {
	result *CountDownByStampSheetResult
	err    error
}

func NewCountDownByStampSheetResultFromJson(data string) CountDownByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCountDownByStampSheetResultFromDict(dict)
}

func NewCountDownByStampSheetResultFromDict(data map[string]interface{}) CountDownByStampSheetResult {
	return CountDownByStampSheetResult{
		Item: func() *Counter {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCounterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CountDownByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CountDownByStampSheetResult) Pointer() *CountDownByStampSheetResult {
	return &p
}

type DeleteByStampSheetResult struct {
	Item *Counter `json:"item"`
}

type DeleteByStampSheetAsyncResult struct {
	result *DeleteByStampSheetResult
	err    error
}

func NewDeleteByStampSheetResultFromJson(data string) DeleteByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteByStampSheetResultFromDict(dict)
}

func NewDeleteByStampSheetResultFromDict(data map[string]interface{}) DeleteByStampSheetResult {
	return DeleteByStampSheetResult{
		Item: func() *Counter {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCounterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteByStampSheetResult) Pointer() *DeleteByStampSheetResult {
	return &p
}

type VerifyCounterByStampTaskResult struct {
	NewContextStack *string `json:"newContextStack"`
}

type VerifyCounterByStampTaskAsyncResult struct {
	result *VerifyCounterByStampTaskResult
	err    error
}

func NewVerifyCounterByStampTaskResultFromJson(data string) VerifyCounterByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyCounterByStampTaskResultFromDict(dict)
}

func NewVerifyCounterByStampTaskResultFromDict(data map[string]interface{}) VerifyCounterByStampTaskResult {
	return VerifyCounterByStampTaskResult{
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p VerifyCounterByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p VerifyCounterByStampTaskResult) Pointer() *VerifyCounterByStampTaskResult {
	return &p
}

type DescribeLimitModelMastersResult struct {
	Items         []LimitModelMaster `json:"items"`
	NextPageToken *string            `json:"nextPageToken"`
}

type DescribeLimitModelMastersAsyncResult struct {
	result *DescribeLimitModelMastersResult
	err    error
}

func NewDescribeLimitModelMastersResultFromJson(data string) DescribeLimitModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLimitModelMastersResultFromDict(dict)
}

func NewDescribeLimitModelMastersResultFromDict(data map[string]interface{}) DescribeLimitModelMastersResult {
	return DescribeLimitModelMastersResult{
		Items: func() []LimitModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastLimitModelMasters(core.CastArray(data["items"]))
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

func (p DescribeLimitModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLimitModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeLimitModelMastersResult) Pointer() *DescribeLimitModelMastersResult {
	return &p
}

type CreateLimitModelMasterResult struct {
	Item *LimitModelMaster `json:"item"`
}

type CreateLimitModelMasterAsyncResult struct {
	result *CreateLimitModelMasterResult
	err    error
}

func NewCreateLimitModelMasterResultFromJson(data string) CreateLimitModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateLimitModelMasterResultFromDict(dict)
}

func NewCreateLimitModelMasterResultFromDict(data map[string]interface{}) CreateLimitModelMasterResult {
	return CreateLimitModelMasterResult{
		Item: func() *LimitModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLimitModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateLimitModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateLimitModelMasterResult) Pointer() *CreateLimitModelMasterResult {
	return &p
}

type GetLimitModelMasterResult struct {
	Item *LimitModelMaster `json:"item"`
}

type GetLimitModelMasterAsyncResult struct {
	result *GetLimitModelMasterResult
	err    error
}

func NewGetLimitModelMasterResultFromJson(data string) GetLimitModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLimitModelMasterResultFromDict(dict)
}

func NewGetLimitModelMasterResultFromDict(data map[string]interface{}) GetLimitModelMasterResult {
	return GetLimitModelMasterResult{
		Item: func() *LimitModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLimitModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetLimitModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetLimitModelMasterResult) Pointer() *GetLimitModelMasterResult {
	return &p
}

type UpdateLimitModelMasterResult struct {
	Item *LimitModelMaster `json:"item"`
}

type UpdateLimitModelMasterAsyncResult struct {
	result *UpdateLimitModelMasterResult
	err    error
}

func NewUpdateLimitModelMasterResultFromJson(data string) UpdateLimitModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateLimitModelMasterResultFromDict(dict)
}

func NewUpdateLimitModelMasterResultFromDict(data map[string]interface{}) UpdateLimitModelMasterResult {
	return UpdateLimitModelMasterResult{
		Item: func() *LimitModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLimitModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateLimitModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateLimitModelMasterResult) Pointer() *UpdateLimitModelMasterResult {
	return &p
}

type DeleteLimitModelMasterResult struct {
	Item *LimitModelMaster `json:"item"`
}

type DeleteLimitModelMasterAsyncResult struct {
	result *DeleteLimitModelMasterResult
	err    error
}

func NewDeleteLimitModelMasterResultFromJson(data string) DeleteLimitModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteLimitModelMasterResultFromDict(dict)
}

func NewDeleteLimitModelMasterResultFromDict(data map[string]interface{}) DeleteLimitModelMasterResult {
	return DeleteLimitModelMasterResult{
		Item: func() *LimitModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLimitModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteLimitModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteLimitModelMasterResult) Pointer() *DeleteLimitModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentLimitMaster `json:"item"`
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
		Item: func() *CurrentLimitMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentLimitMasterFromDict(core.CastMap(data["item"])).Pointer()
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
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentLimitMasterResult struct {
	Item *CurrentLimitMaster `json:"item"`
}

type GetCurrentLimitMasterAsyncResult struct {
	result *GetCurrentLimitMasterResult
	err    error
}

func NewGetCurrentLimitMasterResultFromJson(data string) GetCurrentLimitMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentLimitMasterResultFromDict(dict)
}

func NewGetCurrentLimitMasterResultFromDict(data map[string]interface{}) GetCurrentLimitMasterResult {
	return GetCurrentLimitMasterResult{
		Item: func() *CurrentLimitMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentLimitMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCurrentLimitMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentLimitMasterResult) Pointer() *GetCurrentLimitMasterResult {
	return &p
}

type UpdateCurrentLimitMasterResult struct {
	Item *CurrentLimitMaster `json:"item"`
}

type UpdateCurrentLimitMasterAsyncResult struct {
	result *UpdateCurrentLimitMasterResult
	err    error
}

func NewUpdateCurrentLimitMasterResultFromJson(data string) UpdateCurrentLimitMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentLimitMasterResultFromDict(dict)
}

func NewUpdateCurrentLimitMasterResultFromDict(data map[string]interface{}) UpdateCurrentLimitMasterResult {
	return UpdateCurrentLimitMasterResult{
		Item: func() *CurrentLimitMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentLimitMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentLimitMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentLimitMasterResult) Pointer() *UpdateCurrentLimitMasterResult {
	return &p
}

type UpdateCurrentLimitMasterFromGitHubResult struct {
	Item *CurrentLimitMaster `json:"item"`
}

type UpdateCurrentLimitMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentLimitMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentLimitMasterFromGitHubResultFromJson(data string) UpdateCurrentLimitMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentLimitMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentLimitMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentLimitMasterFromGitHubResult {
	return UpdateCurrentLimitMasterFromGitHubResult{
		Item: func() *CurrentLimitMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentLimitMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentLimitMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentLimitMasterFromGitHubResult) Pointer() *UpdateCurrentLimitMasterFromGitHubResult {
	return &p
}

type DescribeLimitModelsResult struct {
	Items []LimitModel `json:"items"`
}

type DescribeLimitModelsAsyncResult struct {
	result *DescribeLimitModelsResult
	err    error
}

func NewDescribeLimitModelsResultFromJson(data string) DescribeLimitModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLimitModelsResultFromDict(dict)
}

func NewDescribeLimitModelsResultFromDict(data map[string]interface{}) DescribeLimitModelsResult {
	return DescribeLimitModelsResult{
		Items: func() []LimitModel {
			if data["items"] == nil {
				return nil
			}
			return CastLimitModels(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeLimitModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLimitModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeLimitModelsResult) Pointer() *DescribeLimitModelsResult {
	return &p
}

type GetLimitModelResult struct {
	Item *LimitModel `json:"item"`
}

type GetLimitModelAsyncResult struct {
	result *GetLimitModelResult
	err    error
}

func NewGetLimitModelResultFromJson(data string) GetLimitModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLimitModelResultFromDict(dict)
}

func NewGetLimitModelResultFromDict(data map[string]interface{}) GetLimitModelResult {
	return GetLimitModelResult{
		Item: func() *LimitModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLimitModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetLimitModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetLimitModelResult) Pointer() *GetLimitModelResult {
	return &p
}
