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

package distributor

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

type DescribeDistributorModelMastersResult struct {
	Items         []DistributorModelMaster `json:"items"`
	NextPageToken *string                  `json:"nextPageToken"`
	Metadata      *core.ResultMetadata     `json:"metadata"`
}

type DescribeDistributorModelMastersAsyncResult struct {
	result *DescribeDistributorModelMastersResult
	err    error
}

func NewDescribeDistributorModelMastersResultFromJson(data string) DescribeDistributorModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDistributorModelMastersResultFromDict(dict)
}

func NewDescribeDistributorModelMastersResultFromDict(data map[string]interface{}) DescribeDistributorModelMastersResult {
	return DescribeDistributorModelMastersResult{
		Items: func() []DistributorModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastDistributorModelMasters(core.CastArray(data["items"]))
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

func (p DescribeDistributorModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDistributorModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeDistributorModelMastersResult) Pointer() *DescribeDistributorModelMastersResult {
	return &p
}

type CreateDistributorModelMasterResult struct {
	Item     *DistributorModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type CreateDistributorModelMasterAsyncResult struct {
	result *CreateDistributorModelMasterResult
	err    error
}

func NewCreateDistributorModelMasterResultFromJson(data string) CreateDistributorModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateDistributorModelMasterResultFromDict(dict)
}

func NewCreateDistributorModelMasterResultFromDict(data map[string]interface{}) CreateDistributorModelMasterResult {
	return CreateDistributorModelMasterResult{
		Item: func() *DistributorModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDistributorModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateDistributorModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateDistributorModelMasterResult) Pointer() *CreateDistributorModelMasterResult {
	return &p
}

type GetDistributorModelMasterResult struct {
	Item     *DistributorModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type GetDistributorModelMasterAsyncResult struct {
	result *GetDistributorModelMasterResult
	err    error
}

func NewGetDistributorModelMasterResultFromJson(data string) GetDistributorModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDistributorModelMasterResultFromDict(dict)
}

func NewGetDistributorModelMasterResultFromDict(data map[string]interface{}) GetDistributorModelMasterResult {
	return GetDistributorModelMasterResult{
		Item: func() *DistributorModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDistributorModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetDistributorModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetDistributorModelMasterResult) Pointer() *GetDistributorModelMasterResult {
	return &p
}

type UpdateDistributorModelMasterResult struct {
	Item     *DistributorModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type UpdateDistributorModelMasterAsyncResult struct {
	result *UpdateDistributorModelMasterResult
	err    error
}

func NewUpdateDistributorModelMasterResultFromJson(data string) UpdateDistributorModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateDistributorModelMasterResultFromDict(dict)
}

func NewUpdateDistributorModelMasterResultFromDict(data map[string]interface{}) UpdateDistributorModelMasterResult {
	return UpdateDistributorModelMasterResult{
		Item: func() *DistributorModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDistributorModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateDistributorModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateDistributorModelMasterResult) Pointer() *UpdateDistributorModelMasterResult {
	return &p
}

type DeleteDistributorModelMasterResult struct {
	Item     *DistributorModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type DeleteDistributorModelMasterAsyncResult struct {
	result *DeleteDistributorModelMasterResult
	err    error
}

func NewDeleteDistributorModelMasterResultFromJson(data string) DeleteDistributorModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteDistributorModelMasterResultFromDict(dict)
}

func NewDeleteDistributorModelMasterResultFromDict(data map[string]interface{}) DeleteDistributorModelMasterResult {
	return DeleteDistributorModelMasterResult{
		Item: func() *DistributorModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDistributorModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteDistributorModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteDistributorModelMasterResult) Pointer() *DeleteDistributorModelMasterResult {
	return &p
}

type DescribeDistributorModelsResult struct {
	Items    []DistributorModel   `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeDistributorModelsAsyncResult struct {
	result *DescribeDistributorModelsResult
	err    error
}

func NewDescribeDistributorModelsResultFromJson(data string) DescribeDistributorModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDistributorModelsResultFromDict(dict)
}

func NewDescribeDistributorModelsResultFromDict(data map[string]interface{}) DescribeDistributorModelsResult {
	return DescribeDistributorModelsResult{
		Items: func() []DistributorModel {
			if data["items"] == nil {
				return nil
			}
			return CastDistributorModels(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeDistributorModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDistributorModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeDistributorModelsResult) Pointer() *DescribeDistributorModelsResult {
	return &p
}

type GetDistributorModelResult struct {
	Item     *DistributorModel    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetDistributorModelAsyncResult struct {
	result *GetDistributorModelResult
	err    error
}

func NewGetDistributorModelResultFromJson(data string) GetDistributorModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDistributorModelResultFromDict(dict)
}

func NewGetDistributorModelResultFromDict(data map[string]interface{}) GetDistributorModelResult {
	return GetDistributorModelResult{
		Item: func() *DistributorModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDistributorModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetDistributorModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetDistributorModelResult) Pointer() *GetDistributorModelResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentDistributorMaster `json:"item"`
	Metadata *core.ResultMetadata      `json:"metadata"`
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
		Item: func() *CurrentDistributorMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentDistributorMasterFromDict(core.CastMap(data["item"])).Pointer()
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

type GetCurrentDistributorMasterResult struct {
	Item     *CurrentDistributorMaster `json:"item"`
	Metadata *core.ResultMetadata      `json:"metadata"`
}

type GetCurrentDistributorMasterAsyncResult struct {
	result *GetCurrentDistributorMasterResult
	err    error
}

func NewGetCurrentDistributorMasterResultFromJson(data string) GetCurrentDistributorMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentDistributorMasterResultFromDict(dict)
}

func NewGetCurrentDistributorMasterResultFromDict(data map[string]interface{}) GetCurrentDistributorMasterResult {
	return GetCurrentDistributorMasterResult{
		Item: func() *CurrentDistributorMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentDistributorMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCurrentDistributorMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentDistributorMasterResult) Pointer() *GetCurrentDistributorMasterResult {
	return &p
}

type UpdateCurrentDistributorMasterResult struct {
	Item     *CurrentDistributorMaster `json:"item"`
	Metadata *core.ResultMetadata      `json:"metadata"`
}

type UpdateCurrentDistributorMasterAsyncResult struct {
	result *UpdateCurrentDistributorMasterResult
	err    error
}

func NewUpdateCurrentDistributorMasterResultFromJson(data string) UpdateCurrentDistributorMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentDistributorMasterResultFromDict(dict)
}

func NewUpdateCurrentDistributorMasterResultFromDict(data map[string]interface{}) UpdateCurrentDistributorMasterResult {
	return UpdateCurrentDistributorMasterResult{
		Item: func() *CurrentDistributorMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentDistributorMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentDistributorMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentDistributorMasterResult) Pointer() *UpdateCurrentDistributorMasterResult {
	return &p
}

type UpdateCurrentDistributorMasterFromGitHubResult struct {
	Item     *CurrentDistributorMaster `json:"item"`
	Metadata *core.ResultMetadata      `json:"metadata"`
}

type UpdateCurrentDistributorMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentDistributorMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentDistributorMasterFromGitHubResultFromJson(data string) UpdateCurrentDistributorMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentDistributorMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentDistributorMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentDistributorMasterFromGitHubResult {
	return UpdateCurrentDistributorMasterFromGitHubResult{
		Item: func() *CurrentDistributorMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentDistributorMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentDistributorMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentDistributorMasterFromGitHubResult) Pointer() *UpdateCurrentDistributorMasterFromGitHubResult {
	return &p
}

type DistributeResult struct {
	DistributeResource *DistributeResource  `json:"distributeResource"`
	InboxNamespaceId   *string              `json:"inboxNamespaceId"`
	Result             *string              `json:"result"`
	Metadata           *core.ResultMetadata `json:"metadata"`
}

type DistributeAsyncResult struct {
	result *DistributeResult
	err    error
}

func NewDistributeResultFromJson(data string) DistributeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDistributeResultFromDict(dict)
}

func NewDistributeResultFromDict(data map[string]interface{}) DistributeResult {
	return DistributeResult{
		DistributeResource: func() *DistributeResource {
			v, ok := data["distributeResource"]
			if !ok || v == nil {
				return nil
			}
			return NewDistributeResourceFromDict(core.CastMap(data["distributeResource"])).Pointer()
		}(),
		InboxNamespaceId: func() *string {
			v, ok := data["inboxNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inboxNamespaceId"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
	}
}

func (p DistributeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"distributeResource": func() map[string]interface{} {
			if p.DistributeResource == nil {
				return nil
			}
			return p.DistributeResource.ToDict()
		}(),
		"inboxNamespaceId": p.InboxNamespaceId,
		"result":           p.Result,
	}
}

func (p DistributeResult) Pointer() *DistributeResult {
	return &p
}

type DistributeWithoutOverflowProcessResult struct {
	DistributeResource *DistributeResource  `json:"distributeResource"`
	Result             *string              `json:"result"`
	Metadata           *core.ResultMetadata `json:"metadata"`
}

type DistributeWithoutOverflowProcessAsyncResult struct {
	result *DistributeWithoutOverflowProcessResult
	err    error
}

func NewDistributeWithoutOverflowProcessResultFromJson(data string) DistributeWithoutOverflowProcessResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDistributeWithoutOverflowProcessResultFromDict(dict)
}

func NewDistributeWithoutOverflowProcessResultFromDict(data map[string]interface{}) DistributeWithoutOverflowProcessResult {
	return DistributeWithoutOverflowProcessResult{
		DistributeResource: func() *DistributeResource {
			v, ok := data["distributeResource"]
			if !ok || v == nil {
				return nil
			}
			return NewDistributeResourceFromDict(core.CastMap(data["distributeResource"])).Pointer()
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
	}
}

func (p DistributeWithoutOverflowProcessResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"distributeResource": func() map[string]interface{} {
			if p.DistributeResource == nil {
				return nil
			}
			return p.DistributeResource.ToDict()
		}(),
		"result": p.Result,
	}
}

func (p DistributeWithoutOverflowProcessResult) Pointer() *DistributeWithoutOverflowProcessResult {
	return &p
}

type RunVerifyTaskResult struct {
	ContextStack *string              `json:"contextStack"`
	StatusCode   *int32               `json:"statusCode"`
	Result       *string              `json:"result"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type RunVerifyTaskAsyncResult struct {
	result *RunVerifyTaskResult
	err    error
}

func NewRunVerifyTaskResultFromJson(data string) RunVerifyTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunVerifyTaskResultFromDict(dict)
}

func NewRunVerifyTaskResultFromDict(data map[string]interface{}) RunVerifyTaskResult {
	return RunVerifyTaskResult{
		ContextStack: func() *string {
			v, ok := data["contextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contextStack"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
	}
}

func (p RunVerifyTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"contextStack": p.ContextStack,
		"statusCode":   p.StatusCode,
		"result":       p.Result,
	}
}

func (p RunVerifyTaskResult) Pointer() *RunVerifyTaskResult {
	return &p
}

type RunStampTaskResult struct {
	ContextStack *string              `json:"contextStack"`
	StatusCode   *int32               `json:"statusCode"`
	Result       *string              `json:"result"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type RunStampTaskAsyncResult struct {
	result *RunStampTaskResult
	err    error
}

func NewRunStampTaskResultFromJson(data string) RunStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunStampTaskResultFromDict(dict)
}

func NewRunStampTaskResultFromDict(data map[string]interface{}) RunStampTaskResult {
	return RunStampTaskResult{
		ContextStack: func() *string {
			v, ok := data["contextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contextStack"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
	}
}

func (p RunStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"contextStack": p.ContextStack,
		"statusCode":   p.StatusCode,
		"result":       p.Result,
	}
}

func (p RunStampTaskResult) Pointer() *RunStampTaskResult {
	return &p
}

type RunStampSheetResult struct {
	StatusCode *int32               `json:"statusCode"`
	Result     *string              `json:"result"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type RunStampSheetAsyncResult struct {
	result *RunStampSheetResult
	err    error
}

func NewRunStampSheetResultFromJson(data string) RunStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunStampSheetResultFromDict(dict)
}

func NewRunStampSheetResultFromDict(data map[string]interface{}) RunStampSheetResult {
	return RunStampSheetResult{
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
	}
}

func (p RunStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"statusCode": p.StatusCode,
		"result":     p.Result,
	}
}

func (p RunStampSheetResult) Pointer() *RunStampSheetResult {
	return &p
}

type RunStampSheetExpressResult struct {
	VerifyTaskResultCodes []*int32             `json:"verifyTaskResultCodes"`
	VerifyTaskResults     []*string            `json:"verifyTaskResults"`
	TaskResultCodes       []*int32             `json:"taskResultCodes"`
	TaskResults           []*string            `json:"taskResults"`
	SheetResultCode       *int32               `json:"sheetResultCode"`
	SheetResult           *string              `json:"sheetResult"`
	Metadata              *core.ResultMetadata `json:"metadata"`
}

type RunStampSheetExpressAsyncResult struct {
	result *RunStampSheetExpressResult
	err    error
}

func NewRunStampSheetExpressResultFromJson(data string) RunStampSheetExpressResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunStampSheetExpressResultFromDict(dict)
}

func NewRunStampSheetExpressResultFromDict(data map[string]interface{}) RunStampSheetExpressResult {
	return RunStampSheetExpressResult{
		VerifyTaskResultCodes: func() []*int32 {
			v, ok := data["verifyTaskResultCodes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		VerifyTaskResults: func() []*string {
			v, ok := data["verifyTaskResults"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		TaskResultCodes: func() []*int32 {
			v, ok := data["taskResultCodes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		TaskResults: func() []*string {
			v, ok := data["taskResults"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		SheetResultCode: func() *int32 {
			v, ok := data["sheetResultCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["sheetResultCode"])
		}(),
		SheetResult: func() *string {
			v, ok := data["sheetResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sheetResult"])
		}(),
	}
}

func (p RunStampSheetExpressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"verifyTaskResultCodes": core.CastInt32sFromDict(
			p.VerifyTaskResultCodes,
		),
		"verifyTaskResults": core.CastStringsFromDict(
			p.VerifyTaskResults,
		),
		"taskResultCodes": core.CastInt32sFromDict(
			p.TaskResultCodes,
		),
		"taskResults": core.CastStringsFromDict(
			p.TaskResults,
		),
		"sheetResultCode": p.SheetResultCode,
		"sheetResult":     p.SheetResult,
	}
}

func (p RunStampSheetExpressResult) Pointer() *RunStampSheetExpressResult {
	return &p
}

type RunVerifyTaskWithoutNamespaceResult struct {
	ContextStack *string              `json:"contextStack"`
	StatusCode   *int32               `json:"statusCode"`
	Result       *string              `json:"result"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type RunVerifyTaskWithoutNamespaceAsyncResult struct {
	result *RunVerifyTaskWithoutNamespaceResult
	err    error
}

func NewRunVerifyTaskWithoutNamespaceResultFromJson(data string) RunVerifyTaskWithoutNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunVerifyTaskWithoutNamespaceResultFromDict(dict)
}

func NewRunVerifyTaskWithoutNamespaceResultFromDict(data map[string]interface{}) RunVerifyTaskWithoutNamespaceResult {
	return RunVerifyTaskWithoutNamespaceResult{
		ContextStack: func() *string {
			v, ok := data["contextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contextStack"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
	}
}

func (p RunVerifyTaskWithoutNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"contextStack": p.ContextStack,
		"statusCode":   p.StatusCode,
		"result":       p.Result,
	}
}

func (p RunVerifyTaskWithoutNamespaceResult) Pointer() *RunVerifyTaskWithoutNamespaceResult {
	return &p
}

type RunStampTaskWithoutNamespaceResult struct {
	ContextStack *string              `json:"contextStack"`
	StatusCode   *int32               `json:"statusCode"`
	Result       *string              `json:"result"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type RunStampTaskWithoutNamespaceAsyncResult struct {
	result *RunStampTaskWithoutNamespaceResult
	err    error
}

func NewRunStampTaskWithoutNamespaceResultFromJson(data string) RunStampTaskWithoutNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunStampTaskWithoutNamespaceResultFromDict(dict)
}

func NewRunStampTaskWithoutNamespaceResultFromDict(data map[string]interface{}) RunStampTaskWithoutNamespaceResult {
	return RunStampTaskWithoutNamespaceResult{
		ContextStack: func() *string {
			v, ok := data["contextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contextStack"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
	}
}

func (p RunStampTaskWithoutNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"contextStack": p.ContextStack,
		"statusCode":   p.StatusCode,
		"result":       p.Result,
	}
}

func (p RunStampTaskWithoutNamespaceResult) Pointer() *RunStampTaskWithoutNamespaceResult {
	return &p
}

type RunStampSheetWithoutNamespaceResult struct {
	StatusCode *int32               `json:"statusCode"`
	Result     *string              `json:"result"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type RunStampSheetWithoutNamespaceAsyncResult struct {
	result *RunStampSheetWithoutNamespaceResult
	err    error
}

func NewRunStampSheetWithoutNamespaceResultFromJson(data string) RunStampSheetWithoutNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunStampSheetWithoutNamespaceResultFromDict(dict)
}

func NewRunStampSheetWithoutNamespaceResultFromDict(data map[string]interface{}) RunStampSheetWithoutNamespaceResult {
	return RunStampSheetWithoutNamespaceResult{
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
	}
}

func (p RunStampSheetWithoutNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"statusCode": p.StatusCode,
		"result":     p.Result,
	}
}

func (p RunStampSheetWithoutNamespaceResult) Pointer() *RunStampSheetWithoutNamespaceResult {
	return &p
}

type RunStampSheetExpressWithoutNamespaceResult struct {
	VerifyTaskResultCodes []*int32             `json:"verifyTaskResultCodes"`
	VerifyTaskResults     []*string            `json:"verifyTaskResults"`
	TaskResultCodes       []*int32             `json:"taskResultCodes"`
	TaskResults           []*string            `json:"taskResults"`
	SheetResultCode       *int32               `json:"sheetResultCode"`
	SheetResult           *string              `json:"sheetResult"`
	Metadata              *core.ResultMetadata `json:"metadata"`
}

type RunStampSheetExpressWithoutNamespaceAsyncResult struct {
	result *RunStampSheetExpressWithoutNamespaceResult
	err    error
}

func NewRunStampSheetExpressWithoutNamespaceResultFromJson(data string) RunStampSheetExpressWithoutNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunStampSheetExpressWithoutNamespaceResultFromDict(dict)
}

func NewRunStampSheetExpressWithoutNamespaceResultFromDict(data map[string]interface{}) RunStampSheetExpressWithoutNamespaceResult {
	return RunStampSheetExpressWithoutNamespaceResult{
		VerifyTaskResultCodes: func() []*int32 {
			v, ok := data["verifyTaskResultCodes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		VerifyTaskResults: func() []*string {
			v, ok := data["verifyTaskResults"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		TaskResultCodes: func() []*int32 {
			v, ok := data["taskResultCodes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		TaskResults: func() []*string {
			v, ok := data["taskResults"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		SheetResultCode: func() *int32 {
			v, ok := data["sheetResultCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["sheetResultCode"])
		}(),
		SheetResult: func() *string {
			v, ok := data["sheetResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sheetResult"])
		}(),
	}
}

func (p RunStampSheetExpressWithoutNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"verifyTaskResultCodes": core.CastInt32sFromDict(
			p.VerifyTaskResultCodes,
		),
		"verifyTaskResults": core.CastStringsFromDict(
			p.VerifyTaskResults,
		),
		"taskResultCodes": core.CastInt32sFromDict(
			p.TaskResultCodes,
		),
		"taskResults": core.CastStringsFromDict(
			p.TaskResults,
		),
		"sheetResultCode": p.SheetResultCode,
		"sheetResult":     p.SheetResult,
	}
}

func (p RunStampSheetExpressWithoutNamespaceResult) Pointer() *RunStampSheetExpressWithoutNamespaceResult {
	return &p
}

type SetTransactionDefaultConfigResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type SetTransactionDefaultConfigAsyncResult struct {
	result *SetTransactionDefaultConfigResult
	err    error
}

func NewSetTransactionDefaultConfigResultFromJson(data string) SetTransactionDefaultConfigResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetTransactionDefaultConfigResultFromDict(dict)
}

func NewSetTransactionDefaultConfigResultFromDict(data map[string]interface{}) SetTransactionDefaultConfigResult {
	return SetTransactionDefaultConfigResult{
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p SetTransactionDefaultConfigResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p SetTransactionDefaultConfigResult) Pointer() *SetTransactionDefaultConfigResult {
	return &p
}

type SetTransactionDefaultConfigByUserIdResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type SetTransactionDefaultConfigByUserIdAsyncResult struct {
	result *SetTransactionDefaultConfigByUserIdResult
	err    error
}

func NewSetTransactionDefaultConfigByUserIdResultFromJson(data string) SetTransactionDefaultConfigByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetTransactionDefaultConfigByUserIdResultFromDict(dict)
}

func NewSetTransactionDefaultConfigByUserIdResultFromDict(data map[string]interface{}) SetTransactionDefaultConfigByUserIdResult {
	return SetTransactionDefaultConfigByUserIdResult{
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p SetTransactionDefaultConfigByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p SetTransactionDefaultConfigByUserIdResult) Pointer() *SetTransactionDefaultConfigByUserIdResult {
	return &p
}

type FreezeMasterDataResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type FreezeMasterDataAsyncResult struct {
	result *FreezeMasterDataResult
	err    error
}

func NewFreezeMasterDataResultFromJson(data string) FreezeMasterDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFreezeMasterDataResultFromDict(dict)
}

func NewFreezeMasterDataResultFromDict(data map[string]interface{}) FreezeMasterDataResult {
	return FreezeMasterDataResult{
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p FreezeMasterDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p FreezeMasterDataResult) Pointer() *FreezeMasterDataResult {
	return &p
}

type FreezeMasterDataByUserIdResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type FreezeMasterDataByUserIdAsyncResult struct {
	result *FreezeMasterDataByUserIdResult
	err    error
}

func NewFreezeMasterDataByUserIdResultFromJson(data string) FreezeMasterDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFreezeMasterDataByUserIdResultFromDict(dict)
}

func NewFreezeMasterDataByUserIdResultFromDict(data map[string]interface{}) FreezeMasterDataByUserIdResult {
	return FreezeMasterDataByUserIdResult{
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p FreezeMasterDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p FreezeMasterDataByUserIdResult) Pointer() *FreezeMasterDataByUserIdResult {
	return &p
}

type IfExpressionByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type IfExpressionByUserIdAsyncResult struct {
	result *IfExpressionByUserIdResult
	err    error
}

func NewIfExpressionByUserIdResultFromJson(data string) IfExpressionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIfExpressionByUserIdResultFromDict(dict)
}

func NewIfExpressionByUserIdResultFromDict(data map[string]interface{}) IfExpressionByUserIdResult {
	return IfExpressionByUserIdResult{}
}

func (p IfExpressionByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p IfExpressionByUserIdResult) Pointer() *IfExpressionByUserIdResult {
	return &p
}

type AndExpressionByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AndExpressionByUserIdAsyncResult struct {
	result *AndExpressionByUserIdResult
	err    error
}

func NewAndExpressionByUserIdResultFromJson(data string) AndExpressionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAndExpressionByUserIdResultFromDict(dict)
}

func NewAndExpressionByUserIdResultFromDict(data map[string]interface{}) AndExpressionByUserIdResult {
	return AndExpressionByUserIdResult{}
}

func (p AndExpressionByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p AndExpressionByUserIdResult) Pointer() *AndExpressionByUserIdResult {
	return &p
}

type OrExpressionByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type OrExpressionByUserIdAsyncResult struct {
	result *OrExpressionByUserIdResult
	err    error
}

func NewOrExpressionByUserIdResultFromJson(data string) OrExpressionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewOrExpressionByUserIdResultFromDict(dict)
}

func NewOrExpressionByUserIdResultFromDict(data map[string]interface{}) OrExpressionByUserIdResult {
	return OrExpressionByUserIdResult{}
}

func (p OrExpressionByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p OrExpressionByUserIdResult) Pointer() *OrExpressionByUserIdResult {
	return &p
}

type IfExpressionByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type IfExpressionByStampTaskAsyncResult struct {
	result *IfExpressionByStampTaskResult
	err    error
}

func NewIfExpressionByStampTaskResultFromJson(data string) IfExpressionByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIfExpressionByStampTaskResultFromDict(dict)
}

func NewIfExpressionByStampTaskResultFromDict(data map[string]interface{}) IfExpressionByStampTaskResult {
	return IfExpressionByStampTaskResult{
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p IfExpressionByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p IfExpressionByStampTaskResult) Pointer() *IfExpressionByStampTaskResult {
	return &p
}

type AndExpressionByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type AndExpressionByStampTaskAsyncResult struct {
	result *AndExpressionByStampTaskResult
	err    error
}

func NewAndExpressionByStampTaskResultFromJson(data string) AndExpressionByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAndExpressionByStampTaskResultFromDict(dict)
}

func NewAndExpressionByStampTaskResultFromDict(data map[string]interface{}) AndExpressionByStampTaskResult {
	return AndExpressionByStampTaskResult{
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p AndExpressionByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p AndExpressionByStampTaskResult) Pointer() *AndExpressionByStampTaskResult {
	return &p
}

type OrExpressionByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type OrExpressionByStampTaskAsyncResult struct {
	result *OrExpressionByStampTaskResult
	err    error
}

func NewOrExpressionByStampTaskResultFromJson(data string) OrExpressionByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewOrExpressionByStampTaskResultFromDict(dict)
}

func NewOrExpressionByStampTaskResultFromDict(data map[string]interface{}) OrExpressionByStampTaskResult {
	return OrExpressionByStampTaskResult{
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p OrExpressionByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
	}
}

func (p OrExpressionByStampTaskResult) Pointer() *OrExpressionByStampTaskResult {
	return &p
}

type GetStampSheetResultResult struct {
	Item     *StampSheetResult    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStampSheetResultAsyncResult struct {
	result *GetStampSheetResultResult
	err    error
}

func NewGetStampSheetResultResultFromJson(data string) GetStampSheetResultResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStampSheetResultResultFromDict(dict)
}

func NewGetStampSheetResultResultFromDict(data map[string]interface{}) GetStampSheetResultResult {
	return GetStampSheetResultResult{
		Item: func() *StampSheetResult {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStampSheetResultFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetStampSheetResultResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetStampSheetResultResult) Pointer() *GetStampSheetResultResult {
	return &p
}

type GetStampSheetResultByUserIdResult struct {
	Item     *StampSheetResult    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStampSheetResultByUserIdAsyncResult struct {
	result *GetStampSheetResultByUserIdResult
	err    error
}

func NewGetStampSheetResultByUserIdResultFromJson(data string) GetStampSheetResultByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStampSheetResultByUserIdResultFromDict(dict)
}

func NewGetStampSheetResultByUserIdResultFromDict(data map[string]interface{}) GetStampSheetResultByUserIdResult {
	return GetStampSheetResultByUserIdResult{
		Item: func() *StampSheetResult {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStampSheetResultFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetStampSheetResultByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetStampSheetResultByUserIdResult) Pointer() *GetStampSheetResultByUserIdResult {
	return &p
}

type RunTransactionResult struct {
	Item     *TransactionResult   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type RunTransactionAsyncResult struct {
	result *RunTransactionResult
	err    error
}

func NewRunTransactionResultFromJson(data string) RunTransactionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunTransactionResultFromDict(dict)
}

func NewRunTransactionResultFromDict(data map[string]interface{}) RunTransactionResult {
	return RunTransactionResult{
		Item: func() *TransactionResult {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionResultFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p RunTransactionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p RunTransactionResult) Pointer() *RunTransactionResult {
	return &p
}

type GetTransactionResultResult struct {
	Item     *TransactionResult   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetTransactionResultAsyncResult struct {
	result *GetTransactionResultResult
	err    error
}

func NewGetTransactionResultResultFromJson(data string) GetTransactionResultResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTransactionResultResultFromDict(dict)
}

func NewGetTransactionResultResultFromDict(data map[string]interface{}) GetTransactionResultResult {
	return GetTransactionResultResult{
		Item: func() *TransactionResult {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionResultFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetTransactionResultResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetTransactionResultResult) Pointer() *GetTransactionResultResult {
	return &p
}

type GetTransactionResultByUserIdResult struct {
	Item     *TransactionResult   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetTransactionResultByUserIdAsyncResult struct {
	result *GetTransactionResultByUserIdResult
	err    error
}

func NewGetTransactionResultByUserIdResultFromJson(data string) GetTransactionResultByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTransactionResultByUserIdResultFromDict(dict)
}

func NewGetTransactionResultByUserIdResultFromDict(data map[string]interface{}) GetTransactionResultByUserIdResult {
	return GetTransactionResultByUserIdResult{
		Item: func() *TransactionResult {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionResultFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetTransactionResultByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetTransactionResultByUserIdResult) Pointer() *GetTransactionResultByUserIdResult {
	return &p
}
