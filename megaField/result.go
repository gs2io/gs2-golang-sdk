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

type DescribeAreaModelsResult struct {
    Items []AreaModel `json:"items"`
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
    return DescribeAreaModelsResult {
        Items: CastAreaModels(core.CastArray(data["items"])),
    }
}

func (p DescribeAreaModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastAreaModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeAreaModelsResult) Pointer() *DescribeAreaModelsResult {
    return &p
}

type GetAreaModelResult struct {
    Item *AreaModel `json:"item"`
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
    return GetAreaModelResult {
        Item: NewAreaModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetAreaModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetAreaModelResult) Pointer() *GetAreaModelResult {
    return &p
}

type DescribeAreaModelMastersResult struct {
    Items []AreaModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
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
    return DescribeAreaModelMastersResult {
        Items: CastAreaModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeAreaModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastAreaModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeAreaModelMastersResult) Pointer() *DescribeAreaModelMastersResult {
    return &p
}

type CreateAreaModelMasterResult struct {
    Item *AreaModelMaster `json:"item"`
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
    return CreateAreaModelMasterResult {
        Item: NewAreaModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateAreaModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateAreaModelMasterResult) Pointer() *CreateAreaModelMasterResult {
    return &p
}

type GetAreaModelMasterResult struct {
    Item *AreaModelMaster `json:"item"`
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
    return GetAreaModelMasterResult {
        Item: NewAreaModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetAreaModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetAreaModelMasterResult) Pointer() *GetAreaModelMasterResult {
    return &p
}

type UpdateAreaModelMasterResult struct {
    Item *AreaModelMaster `json:"item"`
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
    return UpdateAreaModelMasterResult {
        Item: NewAreaModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateAreaModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateAreaModelMasterResult) Pointer() *UpdateAreaModelMasterResult {
    return &p
}

type DeleteAreaModelMasterResult struct {
    Item *AreaModelMaster `json:"item"`
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
    return DeleteAreaModelMasterResult {
        Item: NewAreaModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteAreaModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteAreaModelMasterResult) Pointer() *DeleteAreaModelMasterResult {
    return &p
}

type DescribeLayerModelsResult struct {
    Items []LayerModel `json:"items"`
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
    return DescribeLayerModelsResult {
        Items: CastLayerModels(core.CastArray(data["items"])),
    }
}

func (p DescribeLayerModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastLayerModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeLayerModelsResult) Pointer() *DescribeLayerModelsResult {
    return &p
}

type GetLayerModelResult struct {
    Item *LayerModel `json:"item"`
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
    return GetLayerModelResult {
        Item: NewLayerModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetLayerModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetLayerModelResult) Pointer() *GetLayerModelResult {
    return &p
}

type DescribeLayerModelMastersResult struct {
    Items []LayerModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
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
    return DescribeLayerModelMastersResult {
        Items: CastLayerModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeLayerModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastLayerModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeLayerModelMastersResult) Pointer() *DescribeLayerModelMastersResult {
    return &p
}

type CreateLayerModelMasterResult struct {
    Item *LayerModelMaster `json:"item"`
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
    return CreateLayerModelMasterResult {
        Item: NewLayerModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateLayerModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateLayerModelMasterResult) Pointer() *CreateLayerModelMasterResult {
    return &p
}

type GetLayerModelMasterResult struct {
    Item *LayerModelMaster `json:"item"`
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
    return GetLayerModelMasterResult {
        Item: NewLayerModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetLayerModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetLayerModelMasterResult) Pointer() *GetLayerModelMasterResult {
    return &p
}

type UpdateLayerModelMasterResult struct {
    Item *LayerModelMaster `json:"item"`
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
    return UpdateLayerModelMasterResult {
        Item: NewLayerModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateLayerModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateLayerModelMasterResult) Pointer() *UpdateLayerModelMasterResult {
    return &p
}

type DeleteLayerModelMasterResult struct {
    Item *LayerModelMaster `json:"item"`
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
    return DeleteLayerModelMasterResult {
        Item: NewLayerModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteLayerModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteLayerModelMasterResult) Pointer() *DeleteLayerModelMasterResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentFieldMaster `json:"item"`
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
    return ExportMasterResult {
        Item: NewCurrentFieldMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
    return &p
}

type GetCurrentFieldMasterResult struct {
    Item *CurrentFieldMaster `json:"item"`
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
    return GetCurrentFieldMasterResult {
        Item: NewCurrentFieldMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentFieldMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentFieldMasterResult) Pointer() *GetCurrentFieldMasterResult {
    return &p
}

type UpdateCurrentFieldMasterResult struct {
    Item *CurrentFieldMaster `json:"item"`
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
    return UpdateCurrentFieldMasterResult {
        Item: NewCurrentFieldMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentFieldMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentFieldMasterResult) Pointer() *UpdateCurrentFieldMasterResult {
    return &p
}

type UpdateCurrentFieldMasterFromGitHubResult struct {
    Item *CurrentFieldMaster `json:"item"`
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
    return UpdateCurrentFieldMasterFromGitHubResult {
        Item: NewCurrentFieldMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentFieldMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentFieldMasterFromGitHubResult) Pointer() *UpdateCurrentFieldMasterFromGitHubResult {
    return &p
}

type PutPositionResult struct {
    Item *Spatial `json:"item"`
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
    return PutPositionResult {
        Item: NewSpatialFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p PutPositionResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p PutPositionResult) Pointer() *PutPositionResult {
    return &p
}

type PutPositionByUserIdResult struct {
    Item *Spatial `json:"item"`
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
    return PutPositionByUserIdResult {
        Item: NewSpatialFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p PutPositionByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p PutPositionByUserIdResult) Pointer() *PutPositionByUserIdResult {
    return &p
}

type FetchPositionResult struct {
    Items []Spatial `json:"items"`
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
    return FetchPositionResult {
        Items: CastSpatials(core.CastArray(data["items"])),
    }
}

func (p FetchPositionResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSpatialsFromDict(
            p.Items,
        ),
    }
}

func (p FetchPositionResult) Pointer() *FetchPositionResult {
    return &p
}

type FetchPositionFromSystemResult struct {
    Items []Spatial `json:"items"`
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
    return FetchPositionFromSystemResult {
        Items: CastSpatials(core.CastArray(data["items"])),
    }
}

func (p FetchPositionFromSystemResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSpatialsFromDict(
            p.Items,
        ),
    }
}

func (p FetchPositionFromSystemResult) Pointer() *FetchPositionFromSystemResult {
    return &p
}

type NearUserIdsResult struct {
    Items []*string `json:"items"`
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
    return NearUserIdsResult {
        Items: core.CastStrings(core.CastArray(data["items"])),
    }
}

func (p NearUserIdsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": core.CastStringsFromDict(
            p.Items,
        ),
    }
}

func (p NearUserIdsResult) Pointer() *NearUserIdsResult {
    return &p
}

type NearUserIdsFromSystemResult struct {
    Items []*string `json:"items"`
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
    return NearUserIdsFromSystemResult {
        Items: core.CastStrings(core.CastArray(data["items"])),
    }
}

func (p NearUserIdsFromSystemResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": core.CastStringsFromDict(
            p.Items,
        ),
    }
}

func (p NearUserIdsFromSystemResult) Pointer() *NearUserIdsFromSystemResult {
    return &p
}

type ActionResult struct {
    Items []Spatial `json:"items"`
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
    return ActionResult {
        Items: CastSpatials(core.CastArray(data["items"])),
    }
}

func (p ActionResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSpatialsFromDict(
            p.Items,
        ),
    }
}

func (p ActionResult) Pointer() *ActionResult {
    return &p
}

type ActionByUserIdResult struct {
    Items []Spatial `json:"items"`
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
    return ActionByUserIdResult {
        Items: CastSpatials(core.CastArray(data["items"])),
    }
}

func (p ActionByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSpatialsFromDict(
            p.Items,
        ),
    }
}

func (p ActionByUserIdResult) Pointer() *ActionByUserIdResult {
    return &p
}