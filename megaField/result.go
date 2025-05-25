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

package megaField

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

type DescribeAreaModelsResult struct {
	Items    []AreaModel          `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeAreaModelsAsyncResult struct {
	result *DescribeAreaModelsResult
	err    error
}

func NewDescribeAreaModelsResultFromJson(data string) DescribeAreaModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAreaModelsResultFromDict(dict)
}

func NewDescribeAreaModelsResultFromDict(data map[string]interface{}) DescribeAreaModelsResult {
	return DescribeAreaModelsResult{
		Items: func() []AreaModel {
			if data["items"] == nil {
				return nil
			}
			return CastAreaModels(core.CastArray(data["items"]))
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

func (p DescribeAreaModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAreaModelsFromDict(
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

func (p DescribeAreaModelsResult) Pointer() *DescribeAreaModelsResult {
	return &p
}

type GetAreaModelResult struct {
	Item     *AreaModel           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetAreaModelAsyncResult struct {
	result *GetAreaModelResult
	err    error
}

func NewGetAreaModelResultFromJson(data string) GetAreaModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAreaModelResultFromDict(dict)
}

func NewGetAreaModelResultFromDict(data map[string]interface{}) GetAreaModelResult {
	return GetAreaModelResult{
		Item: func() *AreaModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAreaModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetAreaModelResult) ToDict() map[string]interface{} {
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

func (p GetAreaModelResult) Pointer() *GetAreaModelResult {
	return &p
}

type DescribeAreaModelMastersResult struct {
	Items         []AreaModelMaster    `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeAreaModelMastersAsyncResult struct {
	result *DescribeAreaModelMastersResult
	err    error
}

func NewDescribeAreaModelMastersResultFromJson(data string) DescribeAreaModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAreaModelMastersResultFromDict(dict)
}

func NewDescribeAreaModelMastersResultFromDict(data map[string]interface{}) DescribeAreaModelMastersResult {
	return DescribeAreaModelMastersResult{
		Items: func() []AreaModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastAreaModelMasters(core.CastArray(data["items"]))
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

func (p DescribeAreaModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAreaModelMastersFromDict(
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

func (p DescribeAreaModelMastersResult) Pointer() *DescribeAreaModelMastersResult {
	return &p
}

type CreateAreaModelMasterResult struct {
	Item     *AreaModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateAreaModelMasterAsyncResult struct {
	result *CreateAreaModelMasterResult
	err    error
}

func NewCreateAreaModelMasterResultFromJson(data string) CreateAreaModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAreaModelMasterResultFromDict(dict)
}

func NewCreateAreaModelMasterResultFromDict(data map[string]interface{}) CreateAreaModelMasterResult {
	return CreateAreaModelMasterResult{
		Item: func() *AreaModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAreaModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateAreaModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateAreaModelMasterResult) Pointer() *CreateAreaModelMasterResult {
	return &p
}

type GetAreaModelMasterResult struct {
	Item     *AreaModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetAreaModelMasterAsyncResult struct {
	result *GetAreaModelMasterResult
	err    error
}

func NewGetAreaModelMasterResultFromJson(data string) GetAreaModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAreaModelMasterResultFromDict(dict)
}

func NewGetAreaModelMasterResultFromDict(data map[string]interface{}) GetAreaModelMasterResult {
	return GetAreaModelMasterResult{
		Item: func() *AreaModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAreaModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetAreaModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetAreaModelMasterResult) Pointer() *GetAreaModelMasterResult {
	return &p
}

type UpdateAreaModelMasterResult struct {
	Item     *AreaModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateAreaModelMasterAsyncResult struct {
	result *UpdateAreaModelMasterResult
	err    error
}

func NewUpdateAreaModelMasterResultFromJson(data string) UpdateAreaModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateAreaModelMasterResultFromDict(dict)
}

func NewUpdateAreaModelMasterResultFromDict(data map[string]interface{}) UpdateAreaModelMasterResult {
	return UpdateAreaModelMasterResult{
		Item: func() *AreaModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAreaModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateAreaModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateAreaModelMasterResult) Pointer() *UpdateAreaModelMasterResult {
	return &p
}

type DeleteAreaModelMasterResult struct {
	Item     *AreaModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteAreaModelMasterAsyncResult struct {
	result *DeleteAreaModelMasterResult
	err    error
}

func NewDeleteAreaModelMasterResultFromJson(data string) DeleteAreaModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAreaModelMasterResultFromDict(dict)
}

func NewDeleteAreaModelMasterResultFromDict(data map[string]interface{}) DeleteAreaModelMasterResult {
	return DeleteAreaModelMasterResult{
		Item: func() *AreaModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAreaModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteAreaModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteAreaModelMasterResult) Pointer() *DeleteAreaModelMasterResult {
	return &p
}

type DescribeLayerModelsResult struct {
	Items    []LayerModel         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeLayerModelsAsyncResult struct {
	result *DescribeLayerModelsResult
	err    error
}

func NewDescribeLayerModelsResultFromJson(data string) DescribeLayerModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLayerModelsResultFromDict(dict)
}

func NewDescribeLayerModelsResultFromDict(data map[string]interface{}) DescribeLayerModelsResult {
	return DescribeLayerModelsResult{
		Items: func() []LayerModel {
			if data["items"] == nil {
				return nil
			}
			return CastLayerModels(core.CastArray(data["items"]))
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

func (p DescribeLayerModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLayerModelsFromDict(
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

func (p DescribeLayerModelsResult) Pointer() *DescribeLayerModelsResult {
	return &p
}

type GetLayerModelResult struct {
	Item     *LayerModel          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetLayerModelAsyncResult struct {
	result *GetLayerModelResult
	err    error
}

func NewGetLayerModelResultFromJson(data string) GetLayerModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLayerModelResultFromDict(dict)
}

func NewGetLayerModelResultFromDict(data map[string]interface{}) GetLayerModelResult {
	return GetLayerModelResult{
		Item: func() *LayerModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLayerModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetLayerModelResult) ToDict() map[string]interface{} {
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

func (p GetLayerModelResult) Pointer() *GetLayerModelResult {
	return &p
}

type DescribeLayerModelMastersResult struct {
	Items         []LayerModelMaster   `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeLayerModelMastersAsyncResult struct {
	result *DescribeLayerModelMastersResult
	err    error
}

func NewDescribeLayerModelMastersResultFromJson(data string) DescribeLayerModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLayerModelMastersResultFromDict(dict)
}

func NewDescribeLayerModelMastersResultFromDict(data map[string]interface{}) DescribeLayerModelMastersResult {
	return DescribeLayerModelMastersResult{
		Items: func() []LayerModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastLayerModelMasters(core.CastArray(data["items"]))
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

func (p DescribeLayerModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLayerModelMastersFromDict(
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

func (p DescribeLayerModelMastersResult) Pointer() *DescribeLayerModelMastersResult {
	return &p
}

type CreateLayerModelMasterResult struct {
	Item     *LayerModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateLayerModelMasterAsyncResult struct {
	result *CreateLayerModelMasterResult
	err    error
}

func NewCreateLayerModelMasterResultFromJson(data string) CreateLayerModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateLayerModelMasterResultFromDict(dict)
}

func NewCreateLayerModelMasterResultFromDict(data map[string]interface{}) CreateLayerModelMasterResult {
	return CreateLayerModelMasterResult{
		Item: func() *LayerModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLayerModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateLayerModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateLayerModelMasterResult) Pointer() *CreateLayerModelMasterResult {
	return &p
}

type GetLayerModelMasterResult struct {
	Item     *LayerModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetLayerModelMasterAsyncResult struct {
	result *GetLayerModelMasterResult
	err    error
}

func NewGetLayerModelMasterResultFromJson(data string) GetLayerModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLayerModelMasterResultFromDict(dict)
}

func NewGetLayerModelMasterResultFromDict(data map[string]interface{}) GetLayerModelMasterResult {
	return GetLayerModelMasterResult{
		Item: func() *LayerModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLayerModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetLayerModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetLayerModelMasterResult) Pointer() *GetLayerModelMasterResult {
	return &p
}

type UpdateLayerModelMasterResult struct {
	Item     *LayerModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateLayerModelMasterAsyncResult struct {
	result *UpdateLayerModelMasterResult
	err    error
}

func NewUpdateLayerModelMasterResultFromJson(data string) UpdateLayerModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateLayerModelMasterResultFromDict(dict)
}

func NewUpdateLayerModelMasterResultFromDict(data map[string]interface{}) UpdateLayerModelMasterResult {
	return UpdateLayerModelMasterResult{
		Item: func() *LayerModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLayerModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateLayerModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateLayerModelMasterResult) Pointer() *UpdateLayerModelMasterResult {
	return &p
}

type DeleteLayerModelMasterResult struct {
	Item     *LayerModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteLayerModelMasterAsyncResult struct {
	result *DeleteLayerModelMasterResult
	err    error
}

func NewDeleteLayerModelMasterResultFromJson(data string) DeleteLayerModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteLayerModelMasterResultFromDict(dict)
}

func NewDeleteLayerModelMasterResultFromDict(data map[string]interface{}) DeleteLayerModelMasterResult {
	return DeleteLayerModelMasterResult{
		Item: func() *LayerModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLayerModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteLayerModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteLayerModelMasterResult) Pointer() *DeleteLayerModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentFieldMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *CurrentFieldMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentFieldMasterFromDict(core.CastMap(data["item"])).Pointer()
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

type GetCurrentFieldMasterResult struct {
	Item     *CurrentFieldMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCurrentFieldMasterAsyncResult struct {
	result *GetCurrentFieldMasterResult
	err    error
}

func NewGetCurrentFieldMasterResultFromJson(data string) GetCurrentFieldMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentFieldMasterResultFromDict(dict)
}

func NewGetCurrentFieldMasterResultFromDict(data map[string]interface{}) GetCurrentFieldMasterResult {
	return GetCurrentFieldMasterResult{
		Item: func() *CurrentFieldMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentFieldMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetCurrentFieldMasterResult) ToDict() map[string]interface{} {
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

func (p GetCurrentFieldMasterResult) Pointer() *GetCurrentFieldMasterResult {
	return &p
}

type PreUpdateCurrentFieldMasterResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreUpdateCurrentFieldMasterAsyncResult struct {
	result *PreUpdateCurrentFieldMasterResult
	err    error
}

func NewPreUpdateCurrentFieldMasterResultFromJson(data string) PreUpdateCurrentFieldMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreUpdateCurrentFieldMasterResultFromDict(dict)
}

func NewPreUpdateCurrentFieldMasterResultFromDict(data map[string]interface{}) PreUpdateCurrentFieldMasterResult {
	return PreUpdateCurrentFieldMasterResult{
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

func (p PreUpdateCurrentFieldMasterResult) ToDict() map[string]interface{} {
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

func (p PreUpdateCurrentFieldMasterResult) Pointer() *PreUpdateCurrentFieldMasterResult {
	return &p
}

type UpdateCurrentFieldMasterResult struct {
	Item     *CurrentFieldMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentFieldMasterAsyncResult struct {
	result *UpdateCurrentFieldMasterResult
	err    error
}

func NewUpdateCurrentFieldMasterResultFromJson(data string) UpdateCurrentFieldMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentFieldMasterResultFromDict(dict)
}

func NewUpdateCurrentFieldMasterResultFromDict(data map[string]interface{}) UpdateCurrentFieldMasterResult {
	return UpdateCurrentFieldMasterResult{
		Item: func() *CurrentFieldMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentFieldMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentFieldMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentFieldMasterResult) Pointer() *UpdateCurrentFieldMasterResult {
	return &p
}

type UpdateCurrentFieldMasterFromGitHubResult struct {
	Item     *CurrentFieldMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentFieldMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentFieldMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentFieldMasterFromGitHubResultFromJson(data string) UpdateCurrentFieldMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentFieldMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentFieldMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentFieldMasterFromGitHubResult {
	return UpdateCurrentFieldMasterFromGitHubResult{
		Item: func() *CurrentFieldMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentFieldMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentFieldMasterFromGitHubResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentFieldMasterFromGitHubResult) Pointer() *UpdateCurrentFieldMasterFromGitHubResult {
	return &p
}

type PutPositionResult struct {
	Item     *Spatial             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PutPositionAsyncResult struct {
	result *PutPositionResult
	err    error
}

func NewPutPositionResultFromJson(data string) PutPositionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutPositionResultFromDict(dict)
}

func NewPutPositionResultFromDict(data map[string]interface{}) PutPositionResult {
	return PutPositionResult{
		Item: func() *Spatial {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSpatialFromDict(core.CastMap(data["item"])).Pointer()
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

func (p PutPositionResult) ToDict() map[string]interface{} {
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

func (p PutPositionResult) Pointer() *PutPositionResult {
	return &p
}

type PutPositionByUserIdResult struct {
	Item     *Spatial             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PutPositionByUserIdAsyncResult struct {
	result *PutPositionByUserIdResult
	err    error
}

func NewPutPositionByUserIdResultFromJson(data string) PutPositionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutPositionByUserIdResultFromDict(dict)
}

func NewPutPositionByUserIdResultFromDict(data map[string]interface{}) PutPositionByUserIdResult {
	return PutPositionByUserIdResult{
		Item: func() *Spatial {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSpatialFromDict(core.CastMap(data["item"])).Pointer()
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

func (p PutPositionByUserIdResult) ToDict() map[string]interface{} {
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

func (p PutPositionByUserIdResult) Pointer() *PutPositionByUserIdResult {
	return &p
}

type FetchPositionResult struct {
	Items    []Spatial            `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type FetchPositionAsyncResult struct {
	result *FetchPositionResult
	err    error
}

func NewFetchPositionResultFromJson(data string) FetchPositionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFetchPositionResultFromDict(dict)
}

func NewFetchPositionResultFromDict(data map[string]interface{}) FetchPositionResult {
	return FetchPositionResult{
		Items: func() []Spatial {
			if data["items"] == nil {
				return nil
			}
			return CastSpatials(core.CastArray(data["items"]))
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

func (p FetchPositionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSpatialsFromDict(
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

func (p FetchPositionResult) Pointer() *FetchPositionResult {
	return &p
}

type FetchPositionFromSystemResult struct {
	Items    []Spatial            `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type FetchPositionFromSystemAsyncResult struct {
	result *FetchPositionFromSystemResult
	err    error
}

func NewFetchPositionFromSystemResultFromJson(data string) FetchPositionFromSystemResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFetchPositionFromSystemResultFromDict(dict)
}

func NewFetchPositionFromSystemResultFromDict(data map[string]interface{}) FetchPositionFromSystemResult {
	return FetchPositionFromSystemResult{
		Items: func() []Spatial {
			if data["items"] == nil {
				return nil
			}
			return CastSpatials(core.CastArray(data["items"]))
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

func (p FetchPositionFromSystemResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSpatialsFromDict(
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

func (p FetchPositionFromSystemResult) Pointer() *FetchPositionFromSystemResult {
	return &p
}

type NearUserIdsResult struct {
	Items    []*string            `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type NearUserIdsAsyncResult struct {
	result *NearUserIdsResult
	err    error
}

func NewNearUserIdsResultFromJson(data string) NearUserIdsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNearUserIdsResultFromDict(dict)
}

func NewNearUserIdsResultFromDict(data map[string]interface{}) NearUserIdsResult {
	return NearUserIdsResult{
		Items: func() []*string {
			v, ok := data["items"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p NearUserIdsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": core.CastStringsFromDict(
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

func (p NearUserIdsResult) Pointer() *NearUserIdsResult {
	return &p
}

type NearUserIdsFromSystemResult struct {
	Items    []*string            `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type NearUserIdsFromSystemAsyncResult struct {
	result *NearUserIdsFromSystemResult
	err    error
}

func NewNearUserIdsFromSystemResultFromJson(data string) NearUserIdsFromSystemResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNearUserIdsFromSystemResultFromDict(dict)
}

func NewNearUserIdsFromSystemResultFromDict(data map[string]interface{}) NearUserIdsFromSystemResult {
	return NearUserIdsFromSystemResult{
		Items: func() []*string {
			v, ok := data["items"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p NearUserIdsFromSystemResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": core.CastStringsFromDict(
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

func (p NearUserIdsFromSystemResult) Pointer() *NearUserIdsFromSystemResult {
	return &p
}

type ActionResult struct {
	Items    []Spatial            `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ActionAsyncResult struct {
	result *ActionResult
	err    error
}

func NewActionResultFromJson(data string) ActionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewActionResultFromDict(dict)
}

func NewActionResultFromDict(data map[string]interface{}) ActionResult {
	return ActionResult{
		Items: func() []Spatial {
			if data["items"] == nil {
				return nil
			}
			return CastSpatials(core.CastArray(data["items"]))
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

func (p ActionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSpatialsFromDict(
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

func (p ActionResult) Pointer() *ActionResult {
	return &p
}

type ActionByUserIdResult struct {
	Items    []Spatial            `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ActionByUserIdAsyncResult struct {
	result *ActionByUserIdResult
	err    error
}

func NewActionByUserIdResultFromJson(data string) ActionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewActionByUserIdResultFromDict(dict)
}

func NewActionByUserIdResultFromDict(data map[string]interface{}) ActionByUserIdResult {
	return ActionByUserIdResult{
		Items: func() []Spatial {
			if data["items"] == nil {
				return nil
			}
			return CastSpatials(core.CastArray(data["items"]))
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

func (p ActionByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSpatialsFromDict(
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

func (p ActionByUserIdResult) Pointer() *ActionByUserIdResult {
	return &p
}
