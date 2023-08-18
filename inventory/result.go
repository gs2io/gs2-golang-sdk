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

package inventory

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

type DescribeInventoryModelMastersResult struct {
    Items []InventoryModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeInventoryModelMastersAsyncResult struct {
	result *DescribeInventoryModelMastersResult
	err    error
}

func NewDescribeInventoryModelMastersResultFromJson(data string) DescribeInventoryModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoryModelMastersResultFromDict(dict)
}

func NewDescribeInventoryModelMastersResultFromDict(data map[string]interface{}) DescribeInventoryModelMastersResult {
    return DescribeInventoryModelMastersResult {
        Items: CastInventoryModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeInventoryModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastInventoryModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeInventoryModelMastersResult) Pointer() *DescribeInventoryModelMastersResult {
    return &p
}

type CreateInventoryModelMasterResult struct {
    Item *InventoryModelMaster `json:"item"`
}

type CreateInventoryModelMasterAsyncResult struct {
	result *CreateInventoryModelMasterResult
	err    error
}

func NewCreateInventoryModelMasterResultFromJson(data string) CreateInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateInventoryModelMasterResultFromDict(dict)
}

func NewCreateInventoryModelMasterResultFromDict(data map[string]interface{}) CreateInventoryModelMasterResult {
    return CreateInventoryModelMasterResult {
        Item: NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateInventoryModelMasterResult) Pointer() *CreateInventoryModelMasterResult {
    return &p
}

type GetInventoryModelMasterResult struct {
    Item *InventoryModelMaster `json:"item"`
}

type GetInventoryModelMasterAsyncResult struct {
	result *GetInventoryModelMasterResult
	err    error
}

func NewGetInventoryModelMasterResultFromJson(data string) GetInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInventoryModelMasterResultFromDict(dict)
}

func NewGetInventoryModelMasterResultFromDict(data map[string]interface{}) GetInventoryModelMasterResult {
    return GetInventoryModelMasterResult {
        Item: NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetInventoryModelMasterResult) Pointer() *GetInventoryModelMasterResult {
    return &p
}

type UpdateInventoryModelMasterResult struct {
    Item *InventoryModelMaster `json:"item"`
}

type UpdateInventoryModelMasterAsyncResult struct {
	result *UpdateInventoryModelMasterResult
	err    error
}

func NewUpdateInventoryModelMasterResultFromJson(data string) UpdateInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateInventoryModelMasterResultFromDict(dict)
}

func NewUpdateInventoryModelMasterResultFromDict(data map[string]interface{}) UpdateInventoryModelMasterResult {
    return UpdateInventoryModelMasterResult {
        Item: NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateInventoryModelMasterResult) Pointer() *UpdateInventoryModelMasterResult {
    return &p
}

type DeleteInventoryModelMasterResult struct {
    Item *InventoryModelMaster `json:"item"`
}

type DeleteInventoryModelMasterAsyncResult struct {
	result *DeleteInventoryModelMasterResult
	err    error
}

func NewDeleteInventoryModelMasterResultFromJson(data string) DeleteInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteInventoryModelMasterResultFromDict(dict)
}

func NewDeleteInventoryModelMasterResultFromDict(data map[string]interface{}) DeleteInventoryModelMasterResult {
    return DeleteInventoryModelMasterResult {
        Item: NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteInventoryModelMasterResult) Pointer() *DeleteInventoryModelMasterResult {
    return &p
}

type DescribeInventoryModelsResult struct {
    Items []InventoryModel `json:"items"`
}

type DescribeInventoryModelsAsyncResult struct {
	result *DescribeInventoryModelsResult
	err    error
}

func NewDescribeInventoryModelsResultFromJson(data string) DescribeInventoryModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoryModelsResultFromDict(dict)
}

func NewDescribeInventoryModelsResultFromDict(data map[string]interface{}) DescribeInventoryModelsResult {
    return DescribeInventoryModelsResult {
        Items: CastInventoryModels(core.CastArray(data["items"])),
    }
}

func (p DescribeInventoryModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastInventoryModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeInventoryModelsResult) Pointer() *DescribeInventoryModelsResult {
    return &p
}

type GetInventoryModelResult struct {
    Item *InventoryModel `json:"item"`
}

type GetInventoryModelAsyncResult struct {
	result *GetInventoryModelResult
	err    error
}

func NewGetInventoryModelResultFromJson(data string) GetInventoryModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInventoryModelResultFromDict(dict)
}

func NewGetInventoryModelResultFromDict(data map[string]interface{}) GetInventoryModelResult {
    return GetInventoryModelResult {
        Item: NewInventoryModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetInventoryModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetInventoryModelResult) Pointer() *GetInventoryModelResult {
    return &p
}

type DescribeItemModelMastersResult struct {
    Items []ItemModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeItemModelMastersAsyncResult struct {
	result *DescribeItemModelMastersResult
	err    error
}

func NewDescribeItemModelMastersResultFromJson(data string) DescribeItemModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeItemModelMastersResultFromDict(dict)
}

func NewDescribeItemModelMastersResultFromDict(data map[string]interface{}) DescribeItemModelMastersResult {
    return DescribeItemModelMastersResult {
        Items: CastItemModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeItemModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeItemModelMastersResult) Pointer() *DescribeItemModelMastersResult {
    return &p
}

type CreateItemModelMasterResult struct {
    Item *ItemModelMaster `json:"item"`
}

type CreateItemModelMasterAsyncResult struct {
	result *CreateItemModelMasterResult
	err    error
}

func NewCreateItemModelMasterResultFromJson(data string) CreateItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateItemModelMasterResultFromDict(dict)
}

func NewCreateItemModelMasterResultFromDict(data map[string]interface{}) CreateItemModelMasterResult {
    return CreateItemModelMasterResult {
        Item: NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateItemModelMasterResult) Pointer() *CreateItemModelMasterResult {
    return &p
}

type GetItemModelMasterResult struct {
    Item *ItemModelMaster `json:"item"`
}

type GetItemModelMasterAsyncResult struct {
	result *GetItemModelMasterResult
	err    error
}

func NewGetItemModelMasterResultFromJson(data string) GetItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemModelMasterResultFromDict(dict)
}

func NewGetItemModelMasterResultFromDict(data map[string]interface{}) GetItemModelMasterResult {
    return GetItemModelMasterResult {
        Item: NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetItemModelMasterResult) Pointer() *GetItemModelMasterResult {
    return &p
}

type UpdateItemModelMasterResult struct {
    Item *ItemModelMaster `json:"item"`
}

type UpdateItemModelMasterAsyncResult struct {
	result *UpdateItemModelMasterResult
	err    error
}

func NewUpdateItemModelMasterResultFromJson(data string) UpdateItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateItemModelMasterResultFromDict(dict)
}

func NewUpdateItemModelMasterResultFromDict(data map[string]interface{}) UpdateItemModelMasterResult {
    return UpdateItemModelMasterResult {
        Item: NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateItemModelMasterResult) Pointer() *UpdateItemModelMasterResult {
    return &p
}

type DeleteItemModelMasterResult struct {
    Item *ItemModelMaster `json:"item"`
}

type DeleteItemModelMasterAsyncResult struct {
	result *DeleteItemModelMasterResult
	err    error
}

func NewDeleteItemModelMasterResultFromJson(data string) DeleteItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteItemModelMasterResultFromDict(dict)
}

func NewDeleteItemModelMasterResultFromDict(data map[string]interface{}) DeleteItemModelMasterResult {
    return DeleteItemModelMasterResult {
        Item: NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteItemModelMasterResult) Pointer() *DeleteItemModelMasterResult {
    return &p
}

type DescribeItemModelsResult struct {
    Items []ItemModel `json:"items"`
}

type DescribeItemModelsAsyncResult struct {
	result *DescribeItemModelsResult
	err    error
}

func NewDescribeItemModelsResultFromJson(data string) DescribeItemModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeItemModelsResultFromDict(dict)
}

func NewDescribeItemModelsResultFromDict(data map[string]interface{}) DescribeItemModelsResult {
    return DescribeItemModelsResult {
        Items: CastItemModels(core.CastArray(data["items"])),
    }
}

func (p DescribeItemModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeItemModelsResult) Pointer() *DescribeItemModelsResult {
    return &p
}

type GetItemModelResult struct {
    Item *ItemModel `json:"item"`
}

type GetItemModelAsyncResult struct {
	result *GetItemModelResult
	err    error
}

func NewGetItemModelResultFromJson(data string) GetItemModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemModelResultFromDict(dict)
}

func NewGetItemModelResultFromDict(data map[string]interface{}) GetItemModelResult {
    return GetItemModelResult {
        Item: NewItemModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetItemModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetItemModelResult) Pointer() *GetItemModelResult {
    return &p
}

type DescribeSimpleInventoryModelMastersResult struct {
    Items []SimpleInventoryModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSimpleInventoryModelMastersAsyncResult struct {
	result *DescribeSimpleInventoryModelMastersResult
	err    error
}

func NewDescribeSimpleInventoryModelMastersResultFromJson(data string) DescribeSimpleInventoryModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSimpleInventoryModelMastersResultFromDict(dict)
}

func NewDescribeSimpleInventoryModelMastersResultFromDict(data map[string]interface{}) DescribeSimpleInventoryModelMastersResult {
    return DescribeSimpleInventoryModelMastersResult {
        Items: CastSimpleInventoryModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSimpleInventoryModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleInventoryModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSimpleInventoryModelMastersResult) Pointer() *DescribeSimpleInventoryModelMastersResult {
    return &p
}

type CreateSimpleInventoryModelMasterResult struct {
    Item *SimpleInventoryModelMaster `json:"item"`
}

type CreateSimpleInventoryModelMasterAsyncResult struct {
	result *CreateSimpleInventoryModelMasterResult
	err    error
}

func NewCreateSimpleInventoryModelMasterResultFromJson(data string) CreateSimpleInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateSimpleInventoryModelMasterResultFromDict(dict)
}

func NewCreateSimpleInventoryModelMasterResultFromDict(data map[string]interface{}) CreateSimpleInventoryModelMasterResult {
    return CreateSimpleInventoryModelMasterResult {
        Item: NewSimpleInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateSimpleInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateSimpleInventoryModelMasterResult) Pointer() *CreateSimpleInventoryModelMasterResult {
    return &p
}

type GetSimpleInventoryModelMasterResult struct {
    Item *SimpleInventoryModelMaster `json:"item"`
}

type GetSimpleInventoryModelMasterAsyncResult struct {
	result *GetSimpleInventoryModelMasterResult
	err    error
}

func NewGetSimpleInventoryModelMasterResultFromJson(data string) GetSimpleInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSimpleInventoryModelMasterResultFromDict(dict)
}

func NewGetSimpleInventoryModelMasterResultFromDict(data map[string]interface{}) GetSimpleInventoryModelMasterResult {
    return GetSimpleInventoryModelMasterResult {
        Item: NewSimpleInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSimpleInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSimpleInventoryModelMasterResult) Pointer() *GetSimpleInventoryModelMasterResult {
    return &p
}

type UpdateSimpleInventoryModelMasterResult struct {
    Item *SimpleInventoryModelMaster `json:"item"`
}

type UpdateSimpleInventoryModelMasterAsyncResult struct {
	result *UpdateSimpleInventoryModelMasterResult
	err    error
}

func NewUpdateSimpleInventoryModelMasterResultFromJson(data string) UpdateSimpleInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateSimpleInventoryModelMasterResultFromDict(dict)
}

func NewUpdateSimpleInventoryModelMasterResultFromDict(data map[string]interface{}) UpdateSimpleInventoryModelMasterResult {
    return UpdateSimpleInventoryModelMasterResult {
        Item: NewSimpleInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateSimpleInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateSimpleInventoryModelMasterResult) Pointer() *UpdateSimpleInventoryModelMasterResult {
    return &p
}

type DeleteSimpleInventoryModelMasterResult struct {
    Item *SimpleInventoryModelMaster `json:"item"`
}

type DeleteSimpleInventoryModelMasterAsyncResult struct {
	result *DeleteSimpleInventoryModelMasterResult
	err    error
}

func NewDeleteSimpleInventoryModelMasterResultFromJson(data string) DeleteSimpleInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteSimpleInventoryModelMasterResultFromDict(dict)
}

func NewDeleteSimpleInventoryModelMasterResultFromDict(data map[string]interface{}) DeleteSimpleInventoryModelMasterResult {
    return DeleteSimpleInventoryModelMasterResult {
        Item: NewSimpleInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteSimpleInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteSimpleInventoryModelMasterResult) Pointer() *DeleteSimpleInventoryModelMasterResult {
    return &p
}

type DescribeSimpleInventoryModelsResult struct {
    Items []SimpleInventoryModel `json:"items"`
}

type DescribeSimpleInventoryModelsAsyncResult struct {
	result *DescribeSimpleInventoryModelsResult
	err    error
}

func NewDescribeSimpleInventoryModelsResultFromJson(data string) DescribeSimpleInventoryModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSimpleInventoryModelsResultFromDict(dict)
}

func NewDescribeSimpleInventoryModelsResultFromDict(data map[string]interface{}) DescribeSimpleInventoryModelsResult {
    return DescribeSimpleInventoryModelsResult {
        Items: CastSimpleInventoryModels(core.CastArray(data["items"])),
    }
}

func (p DescribeSimpleInventoryModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleInventoryModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeSimpleInventoryModelsResult) Pointer() *DescribeSimpleInventoryModelsResult {
    return &p
}

type GetSimpleInventoryModelResult struct {
    Item *SimpleInventoryModel `json:"item"`
}

type GetSimpleInventoryModelAsyncResult struct {
	result *GetSimpleInventoryModelResult
	err    error
}

func NewGetSimpleInventoryModelResultFromJson(data string) GetSimpleInventoryModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSimpleInventoryModelResultFromDict(dict)
}

func NewGetSimpleInventoryModelResultFromDict(data map[string]interface{}) GetSimpleInventoryModelResult {
    return GetSimpleInventoryModelResult {
        Item: NewSimpleInventoryModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSimpleInventoryModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSimpleInventoryModelResult) Pointer() *GetSimpleInventoryModelResult {
    return &p
}

type DescribeSimpleItemModelMastersResult struct {
    Items []SimpleItemModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSimpleItemModelMastersAsyncResult struct {
	result *DescribeSimpleItemModelMastersResult
	err    error
}

func NewDescribeSimpleItemModelMastersResultFromJson(data string) DescribeSimpleItemModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSimpleItemModelMastersResultFromDict(dict)
}

func NewDescribeSimpleItemModelMastersResultFromDict(data map[string]interface{}) DescribeSimpleItemModelMastersResult {
    return DescribeSimpleItemModelMastersResult {
        Items: CastSimpleItemModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSimpleItemModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleItemModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSimpleItemModelMastersResult) Pointer() *DescribeSimpleItemModelMastersResult {
    return &p
}

type CreateSimpleItemModelMasterResult struct {
    Item *SimpleItemModelMaster `json:"item"`
}

type CreateSimpleItemModelMasterAsyncResult struct {
	result *CreateSimpleItemModelMasterResult
	err    error
}

func NewCreateSimpleItemModelMasterResultFromJson(data string) CreateSimpleItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateSimpleItemModelMasterResultFromDict(dict)
}

func NewCreateSimpleItemModelMasterResultFromDict(data map[string]interface{}) CreateSimpleItemModelMasterResult {
    return CreateSimpleItemModelMasterResult {
        Item: NewSimpleItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateSimpleItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateSimpleItemModelMasterResult) Pointer() *CreateSimpleItemModelMasterResult {
    return &p
}

type GetSimpleItemModelMasterResult struct {
    Item *SimpleItemModelMaster `json:"item"`
}

type GetSimpleItemModelMasterAsyncResult struct {
	result *GetSimpleItemModelMasterResult
	err    error
}

func NewGetSimpleItemModelMasterResultFromJson(data string) GetSimpleItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSimpleItemModelMasterResultFromDict(dict)
}

func NewGetSimpleItemModelMasterResultFromDict(data map[string]interface{}) GetSimpleItemModelMasterResult {
    return GetSimpleItemModelMasterResult {
        Item: NewSimpleItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSimpleItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSimpleItemModelMasterResult) Pointer() *GetSimpleItemModelMasterResult {
    return &p
}

type UpdateSimpleItemModelMasterResult struct {
    Item *SimpleItemModelMaster `json:"item"`
}

type UpdateSimpleItemModelMasterAsyncResult struct {
	result *UpdateSimpleItemModelMasterResult
	err    error
}

func NewUpdateSimpleItemModelMasterResultFromJson(data string) UpdateSimpleItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateSimpleItemModelMasterResultFromDict(dict)
}

func NewUpdateSimpleItemModelMasterResultFromDict(data map[string]interface{}) UpdateSimpleItemModelMasterResult {
    return UpdateSimpleItemModelMasterResult {
        Item: NewSimpleItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateSimpleItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateSimpleItemModelMasterResult) Pointer() *UpdateSimpleItemModelMasterResult {
    return &p
}

type DeleteSimpleItemModelMasterResult struct {
    Item *SimpleItemModelMaster `json:"item"`
}

type DeleteSimpleItemModelMasterAsyncResult struct {
	result *DeleteSimpleItemModelMasterResult
	err    error
}

func NewDeleteSimpleItemModelMasterResultFromJson(data string) DeleteSimpleItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteSimpleItemModelMasterResultFromDict(dict)
}

func NewDeleteSimpleItemModelMasterResultFromDict(data map[string]interface{}) DeleteSimpleItemModelMasterResult {
    return DeleteSimpleItemModelMasterResult {
        Item: NewSimpleItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteSimpleItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteSimpleItemModelMasterResult) Pointer() *DeleteSimpleItemModelMasterResult {
    return &p
}

type DescribeSimpleItemModelsResult struct {
    Items []SimpleItemModel `json:"items"`
}

type DescribeSimpleItemModelsAsyncResult struct {
	result *DescribeSimpleItemModelsResult
	err    error
}

func NewDescribeSimpleItemModelsResultFromJson(data string) DescribeSimpleItemModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSimpleItemModelsResultFromDict(dict)
}

func NewDescribeSimpleItemModelsResultFromDict(data map[string]interface{}) DescribeSimpleItemModelsResult {
    return DescribeSimpleItemModelsResult {
        Items: CastSimpleItemModels(core.CastArray(data["items"])),
    }
}

func (p DescribeSimpleItemModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleItemModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeSimpleItemModelsResult) Pointer() *DescribeSimpleItemModelsResult {
    return &p
}

type GetSimpleItemModelResult struct {
    Item *SimpleItemModel `json:"item"`
}

type GetSimpleItemModelAsyncResult struct {
	result *GetSimpleItemModelResult
	err    error
}

func NewGetSimpleItemModelResultFromJson(data string) GetSimpleItemModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSimpleItemModelResultFromDict(dict)
}

func NewGetSimpleItemModelResultFromDict(data map[string]interface{}) GetSimpleItemModelResult {
    return GetSimpleItemModelResult {
        Item: NewSimpleItemModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSimpleItemModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSimpleItemModelResult) Pointer() *GetSimpleItemModelResult {
    return &p
}

type DescribeBigInventoryModelMastersResult struct {
    Items []BigInventoryModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeBigInventoryModelMastersAsyncResult struct {
	result *DescribeBigInventoryModelMastersResult
	err    error
}

func NewDescribeBigInventoryModelMastersResultFromJson(data string) DescribeBigInventoryModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBigInventoryModelMastersResultFromDict(dict)
}

func NewDescribeBigInventoryModelMastersResultFromDict(data map[string]interface{}) DescribeBigInventoryModelMastersResult {
    return DescribeBigInventoryModelMastersResult {
        Items: CastBigInventoryModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeBigInventoryModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBigInventoryModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeBigInventoryModelMastersResult) Pointer() *DescribeBigInventoryModelMastersResult {
    return &p
}

type CreateBigInventoryModelMasterResult struct {
    Item *BigInventoryModelMaster `json:"item"`
}

type CreateBigInventoryModelMasterAsyncResult struct {
	result *CreateBigInventoryModelMasterResult
	err    error
}

func NewCreateBigInventoryModelMasterResultFromJson(data string) CreateBigInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateBigInventoryModelMasterResultFromDict(dict)
}

func NewCreateBigInventoryModelMasterResultFromDict(data map[string]interface{}) CreateBigInventoryModelMasterResult {
    return CreateBigInventoryModelMasterResult {
        Item: NewBigInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateBigInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateBigInventoryModelMasterResult) Pointer() *CreateBigInventoryModelMasterResult {
    return &p
}

type GetBigInventoryModelMasterResult struct {
    Item *BigInventoryModelMaster `json:"item"`
}

type GetBigInventoryModelMasterAsyncResult struct {
	result *GetBigInventoryModelMasterResult
	err    error
}

func NewGetBigInventoryModelMasterResultFromJson(data string) GetBigInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBigInventoryModelMasterResultFromDict(dict)
}

func NewGetBigInventoryModelMasterResultFromDict(data map[string]interface{}) GetBigInventoryModelMasterResult {
    return GetBigInventoryModelMasterResult {
        Item: NewBigInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBigInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBigInventoryModelMasterResult) Pointer() *GetBigInventoryModelMasterResult {
    return &p
}

type UpdateBigInventoryModelMasterResult struct {
    Item *BigInventoryModelMaster `json:"item"`
}

type UpdateBigInventoryModelMasterAsyncResult struct {
	result *UpdateBigInventoryModelMasterResult
	err    error
}

func NewUpdateBigInventoryModelMasterResultFromJson(data string) UpdateBigInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateBigInventoryModelMasterResultFromDict(dict)
}

func NewUpdateBigInventoryModelMasterResultFromDict(data map[string]interface{}) UpdateBigInventoryModelMasterResult {
    return UpdateBigInventoryModelMasterResult {
        Item: NewBigInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateBigInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateBigInventoryModelMasterResult) Pointer() *UpdateBigInventoryModelMasterResult {
    return &p
}

type DeleteBigInventoryModelMasterResult struct {
    Item *BigInventoryModelMaster `json:"item"`
}

type DeleteBigInventoryModelMasterAsyncResult struct {
	result *DeleteBigInventoryModelMasterResult
	err    error
}

func NewDeleteBigInventoryModelMasterResultFromJson(data string) DeleteBigInventoryModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteBigInventoryModelMasterResultFromDict(dict)
}

func NewDeleteBigInventoryModelMasterResultFromDict(data map[string]interface{}) DeleteBigInventoryModelMasterResult {
    return DeleteBigInventoryModelMasterResult {
        Item: NewBigInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteBigInventoryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteBigInventoryModelMasterResult) Pointer() *DeleteBigInventoryModelMasterResult {
    return &p
}

type DescribeBigInventoryModelsResult struct {
    Items []BigInventoryModel `json:"items"`
}

type DescribeBigInventoryModelsAsyncResult struct {
	result *DescribeBigInventoryModelsResult
	err    error
}

func NewDescribeBigInventoryModelsResultFromJson(data string) DescribeBigInventoryModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBigInventoryModelsResultFromDict(dict)
}

func NewDescribeBigInventoryModelsResultFromDict(data map[string]interface{}) DescribeBigInventoryModelsResult {
    return DescribeBigInventoryModelsResult {
        Items: CastBigInventoryModels(core.CastArray(data["items"])),
    }
}

func (p DescribeBigInventoryModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBigInventoryModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeBigInventoryModelsResult) Pointer() *DescribeBigInventoryModelsResult {
    return &p
}

type GetBigInventoryModelResult struct {
    Item *BigInventoryModel `json:"item"`
}

type GetBigInventoryModelAsyncResult struct {
	result *GetBigInventoryModelResult
	err    error
}

func NewGetBigInventoryModelResultFromJson(data string) GetBigInventoryModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBigInventoryModelResultFromDict(dict)
}

func NewGetBigInventoryModelResultFromDict(data map[string]interface{}) GetBigInventoryModelResult {
    return GetBigInventoryModelResult {
        Item: NewBigInventoryModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBigInventoryModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBigInventoryModelResult) Pointer() *GetBigInventoryModelResult {
    return &p
}

type DescribeBigItemModelMastersResult struct {
    Items []BigItemModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeBigItemModelMastersAsyncResult struct {
	result *DescribeBigItemModelMastersResult
	err    error
}

func NewDescribeBigItemModelMastersResultFromJson(data string) DescribeBigItemModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBigItemModelMastersResultFromDict(dict)
}

func NewDescribeBigItemModelMastersResultFromDict(data map[string]interface{}) DescribeBigItemModelMastersResult {
    return DescribeBigItemModelMastersResult {
        Items: CastBigItemModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeBigItemModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBigItemModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeBigItemModelMastersResult) Pointer() *DescribeBigItemModelMastersResult {
    return &p
}

type CreateBigItemModelMasterResult struct {
    Item *BigItemModelMaster `json:"item"`
}

type CreateBigItemModelMasterAsyncResult struct {
	result *CreateBigItemModelMasterResult
	err    error
}

func NewCreateBigItemModelMasterResultFromJson(data string) CreateBigItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateBigItemModelMasterResultFromDict(dict)
}

func NewCreateBigItemModelMasterResultFromDict(data map[string]interface{}) CreateBigItemModelMasterResult {
    return CreateBigItemModelMasterResult {
        Item: NewBigItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateBigItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateBigItemModelMasterResult) Pointer() *CreateBigItemModelMasterResult {
    return &p
}

type GetBigItemModelMasterResult struct {
    Item *BigItemModelMaster `json:"item"`
}

type GetBigItemModelMasterAsyncResult struct {
	result *GetBigItemModelMasterResult
	err    error
}

func NewGetBigItemModelMasterResultFromJson(data string) GetBigItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBigItemModelMasterResultFromDict(dict)
}

func NewGetBigItemModelMasterResultFromDict(data map[string]interface{}) GetBigItemModelMasterResult {
    return GetBigItemModelMasterResult {
        Item: NewBigItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBigItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBigItemModelMasterResult) Pointer() *GetBigItemModelMasterResult {
    return &p
}

type UpdateBigItemModelMasterResult struct {
    Item *BigItemModelMaster `json:"item"`
}

type UpdateBigItemModelMasterAsyncResult struct {
	result *UpdateBigItemModelMasterResult
	err    error
}

func NewUpdateBigItemModelMasterResultFromJson(data string) UpdateBigItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateBigItemModelMasterResultFromDict(dict)
}

func NewUpdateBigItemModelMasterResultFromDict(data map[string]interface{}) UpdateBigItemModelMasterResult {
    return UpdateBigItemModelMasterResult {
        Item: NewBigItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateBigItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateBigItemModelMasterResult) Pointer() *UpdateBigItemModelMasterResult {
    return &p
}

type DeleteBigItemModelMasterResult struct {
    Item *BigItemModelMaster `json:"item"`
}

type DeleteBigItemModelMasterAsyncResult struct {
	result *DeleteBigItemModelMasterResult
	err    error
}

func NewDeleteBigItemModelMasterResultFromJson(data string) DeleteBigItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteBigItemModelMasterResultFromDict(dict)
}

func NewDeleteBigItemModelMasterResultFromDict(data map[string]interface{}) DeleteBigItemModelMasterResult {
    return DeleteBigItemModelMasterResult {
        Item: NewBigItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteBigItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteBigItemModelMasterResult) Pointer() *DeleteBigItemModelMasterResult {
    return &p
}

type DescribeBigItemModelsResult struct {
    Items []BigItemModel `json:"items"`
}

type DescribeBigItemModelsAsyncResult struct {
	result *DescribeBigItemModelsResult
	err    error
}

func NewDescribeBigItemModelsResultFromJson(data string) DescribeBigItemModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBigItemModelsResultFromDict(dict)
}

func NewDescribeBigItemModelsResultFromDict(data map[string]interface{}) DescribeBigItemModelsResult {
    return DescribeBigItemModelsResult {
        Items: CastBigItemModels(core.CastArray(data["items"])),
    }
}

func (p DescribeBigItemModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBigItemModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeBigItemModelsResult) Pointer() *DescribeBigItemModelsResult {
    return &p
}

type GetBigItemModelResult struct {
    Item *BigItemModel `json:"item"`
}

type GetBigItemModelAsyncResult struct {
	result *GetBigItemModelResult
	err    error
}

func NewGetBigItemModelResultFromJson(data string) GetBigItemModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBigItemModelResultFromDict(dict)
}

func NewGetBigItemModelResultFromDict(data map[string]interface{}) GetBigItemModelResult {
    return GetBigItemModelResult {
        Item: NewBigItemModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBigItemModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBigItemModelResult) Pointer() *GetBigItemModelResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentItemModelMaster `json:"item"`
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
        Item: NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentItemModelMasterResult struct {
    Item *CurrentItemModelMaster `json:"item"`
}

type GetCurrentItemModelMasterAsyncResult struct {
	result *GetCurrentItemModelMasterResult
	err    error
}

func NewGetCurrentItemModelMasterResultFromJson(data string) GetCurrentItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentItemModelMasterResultFromDict(dict)
}

func NewGetCurrentItemModelMasterResultFromDict(data map[string]interface{}) GetCurrentItemModelMasterResult {
    return GetCurrentItemModelMasterResult {
        Item: NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentItemModelMasterResult) Pointer() *GetCurrentItemModelMasterResult {
    return &p
}

type UpdateCurrentItemModelMasterResult struct {
    Item *CurrentItemModelMaster `json:"item"`
}

type UpdateCurrentItemModelMasterAsyncResult struct {
	result *UpdateCurrentItemModelMasterResult
	err    error
}

func NewUpdateCurrentItemModelMasterResultFromJson(data string) UpdateCurrentItemModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentItemModelMasterResultFromDict(dict)
}

func NewUpdateCurrentItemModelMasterResultFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterResult {
    return UpdateCurrentItemModelMasterResult {
        Item: NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentItemModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentItemModelMasterResult) Pointer() *UpdateCurrentItemModelMasterResult {
    return &p
}

type UpdateCurrentItemModelMasterFromGitHubResult struct {
    Item *CurrentItemModelMaster `json:"item"`
}

type UpdateCurrentItemModelMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentItemModelMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentItemModelMasterFromGitHubResultFromJson(data string) UpdateCurrentItemModelMasterFromGitHubResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentItemModelMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentItemModelMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterFromGitHubResult {
    return UpdateCurrentItemModelMasterFromGitHubResult {
        Item: NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentItemModelMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentItemModelMasterFromGitHubResult) Pointer() *UpdateCurrentItemModelMasterFromGitHubResult {
    return &p
}

type DescribeInventoriesResult struct {
    Items []Inventory `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeInventoriesAsyncResult struct {
	result *DescribeInventoriesResult
	err    error
}

func NewDescribeInventoriesResultFromJson(data string) DescribeInventoriesResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoriesResultFromDict(dict)
}

func NewDescribeInventoriesResultFromDict(data map[string]interface{}) DescribeInventoriesResult {
    return DescribeInventoriesResult {
        Items: CastInventories(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeInventoriesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastInventoriesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeInventoriesResult) Pointer() *DescribeInventoriesResult {
    return &p
}

type DescribeInventoriesByUserIdResult struct {
    Items []Inventory `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeInventoriesByUserIdAsyncResult struct {
	result *DescribeInventoriesByUserIdResult
	err    error
}

func NewDescribeInventoriesByUserIdResultFromJson(data string) DescribeInventoriesByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoriesByUserIdResultFromDict(dict)
}

func NewDescribeInventoriesByUserIdResultFromDict(data map[string]interface{}) DescribeInventoriesByUserIdResult {
    return DescribeInventoriesByUserIdResult {
        Items: CastInventories(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeInventoriesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastInventoriesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeInventoriesByUserIdResult) Pointer() *DescribeInventoriesByUserIdResult {
    return &p
}

type GetInventoryResult struct {
    Item *Inventory `json:"item"`
}

type GetInventoryAsyncResult struct {
	result *GetInventoryResult
	err    error
}

func NewGetInventoryResultFromJson(data string) GetInventoryResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInventoryResultFromDict(dict)
}

func NewGetInventoryResultFromDict(data map[string]interface{}) GetInventoryResult {
    return GetInventoryResult {
        Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetInventoryResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetInventoryResult) Pointer() *GetInventoryResult {
    return &p
}

type GetInventoryByUserIdResult struct {
    Item *Inventory `json:"item"`
}

type GetInventoryByUserIdAsyncResult struct {
	result *GetInventoryByUserIdResult
	err    error
}

func NewGetInventoryByUserIdResultFromJson(data string) GetInventoryByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInventoryByUserIdResultFromDict(dict)
}

func NewGetInventoryByUserIdResultFromDict(data map[string]interface{}) GetInventoryByUserIdResult {
    return GetInventoryByUserIdResult {
        Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetInventoryByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetInventoryByUserIdResult) Pointer() *GetInventoryByUserIdResult {
    return &p
}

type AddCapacityByUserIdResult struct {
    Item *Inventory `json:"item"`
}

type AddCapacityByUserIdAsyncResult struct {
	result *AddCapacityByUserIdResult
	err    error
}

func NewAddCapacityByUserIdResultFromJson(data string) AddCapacityByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddCapacityByUserIdResultFromDict(dict)
}

func NewAddCapacityByUserIdResultFromDict(data map[string]interface{}) AddCapacityByUserIdResult {
    return AddCapacityByUserIdResult {
        Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AddCapacityByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AddCapacityByUserIdResult) Pointer() *AddCapacityByUserIdResult {
    return &p
}

type SetCapacityByUserIdResult struct {
    Item *Inventory `json:"item"`
}

type SetCapacityByUserIdAsyncResult struct {
	result *SetCapacityByUserIdResult
	err    error
}

func NewSetCapacityByUserIdResultFromJson(data string) SetCapacityByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetCapacityByUserIdResultFromDict(dict)
}

func NewSetCapacityByUserIdResultFromDict(data map[string]interface{}) SetCapacityByUserIdResult {
    return SetCapacityByUserIdResult {
        Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SetCapacityByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SetCapacityByUserIdResult) Pointer() *SetCapacityByUserIdResult {
    return &p
}

type DeleteInventoryByUserIdResult struct {
    Item *Inventory `json:"item"`
}

type DeleteInventoryByUserIdAsyncResult struct {
	result *DeleteInventoryByUserIdResult
	err    error
}

func NewDeleteInventoryByUserIdResultFromJson(data string) DeleteInventoryByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteInventoryByUserIdResultFromDict(dict)
}

func NewDeleteInventoryByUserIdResultFromDict(data map[string]interface{}) DeleteInventoryByUserIdResult {
    return DeleteInventoryByUserIdResult {
        Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteInventoryByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteInventoryByUserIdResult) Pointer() *DeleteInventoryByUserIdResult {
    return &p
}

type AddCapacityByStampSheetResult struct {
    Item *Inventory `json:"item"`
}

type AddCapacityByStampSheetAsyncResult struct {
	result *AddCapacityByStampSheetResult
	err    error
}

func NewAddCapacityByStampSheetResultFromJson(data string) AddCapacityByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddCapacityByStampSheetResultFromDict(dict)
}

func NewAddCapacityByStampSheetResultFromDict(data map[string]interface{}) AddCapacityByStampSheetResult {
    return AddCapacityByStampSheetResult {
        Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AddCapacityByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AddCapacityByStampSheetResult) Pointer() *AddCapacityByStampSheetResult {
    return &p
}

type SetCapacityByStampSheetResult struct {
    Item *Inventory `json:"item"`
}

type SetCapacityByStampSheetAsyncResult struct {
	result *SetCapacityByStampSheetResult
	err    error
}

func NewSetCapacityByStampSheetResultFromJson(data string) SetCapacityByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetCapacityByStampSheetResultFromDict(dict)
}

func NewSetCapacityByStampSheetResultFromDict(data map[string]interface{}) SetCapacityByStampSheetResult {
    return SetCapacityByStampSheetResult {
        Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SetCapacityByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SetCapacityByStampSheetResult) Pointer() *SetCapacityByStampSheetResult {
    return &p
}

type DescribeItemSetsResult struct {
    Items []ItemSet `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeItemSetsAsyncResult struct {
	result *DescribeItemSetsResult
	err    error
}

func NewDescribeItemSetsResultFromJson(data string) DescribeItemSetsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeItemSetsResultFromDict(dict)
}

func NewDescribeItemSetsResultFromDict(data map[string]interface{}) DescribeItemSetsResult {
    return DescribeItemSetsResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeItemSetsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeItemSetsResult) Pointer() *DescribeItemSetsResult {
    return &p
}

type DescribeItemSetsByUserIdResult struct {
    Items []ItemSet `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeItemSetsByUserIdAsyncResult struct {
	result *DescribeItemSetsByUserIdResult
	err    error
}

func NewDescribeItemSetsByUserIdResultFromJson(data string) DescribeItemSetsByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeItemSetsByUserIdResultFromDict(dict)
}

func NewDescribeItemSetsByUserIdResultFromDict(data map[string]interface{}) DescribeItemSetsByUserIdResult {
    return DescribeItemSetsByUserIdResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeItemSetsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeItemSetsByUserIdResult) Pointer() *DescribeItemSetsByUserIdResult {
    return &p
}

type GetItemSetResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type GetItemSetAsyncResult struct {
	result *GetItemSetResult
	err    error
}

func NewGetItemSetResultFromJson(data string) GetItemSetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemSetResultFromDict(dict)
}

func NewGetItemSetResultFromDict(data map[string]interface{}) GetItemSetResult {
    return GetItemSetResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p GetItemSetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p GetItemSetResult) Pointer() *GetItemSetResult {
    return &p
}

type GetItemSetByUserIdResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type GetItemSetByUserIdAsyncResult struct {
	result *GetItemSetByUserIdResult
	err    error
}

func NewGetItemSetByUserIdResultFromJson(data string) GetItemSetByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemSetByUserIdResultFromDict(dict)
}

func NewGetItemSetByUserIdResultFromDict(data map[string]interface{}) GetItemSetByUserIdResult {
    return GetItemSetByUserIdResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p GetItemSetByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p GetItemSetByUserIdResult) Pointer() *GetItemSetByUserIdResult {
    return &p
}

type GetItemWithSignatureResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
    Body *string `json:"body"`
    Signature *string `json:"signature"`
}

type GetItemWithSignatureAsyncResult struct {
	result *GetItemWithSignatureResult
	err    error
}

func NewGetItemWithSignatureResultFromJson(data string) GetItemWithSignatureResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemWithSignatureResultFromDict(dict)
}

func NewGetItemWithSignatureResultFromDict(data map[string]interface{}) GetItemWithSignatureResult {
    return GetItemWithSignatureResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
        Body: core.CastString(data["body"]),
        Signature: core.CastString(data["signature"]),
    }
}

func (p GetItemWithSignatureResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
        "body": p.Body,
        "signature": p.Signature,
    }
}

func (p GetItemWithSignatureResult) Pointer() *GetItemWithSignatureResult {
    return &p
}

type GetItemWithSignatureByUserIdResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
    Body *string `json:"body"`
    Signature *string `json:"signature"`
}

type GetItemWithSignatureByUserIdAsyncResult struct {
	result *GetItemWithSignatureByUserIdResult
	err    error
}

func NewGetItemWithSignatureByUserIdResultFromJson(data string) GetItemWithSignatureByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemWithSignatureByUserIdResultFromDict(dict)
}

func NewGetItemWithSignatureByUserIdResultFromDict(data map[string]interface{}) GetItemWithSignatureByUserIdResult {
    return GetItemWithSignatureByUserIdResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
        Body: core.CastString(data["body"]),
        Signature: core.CastString(data["signature"]),
    }
}

func (p GetItemWithSignatureByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
        "body": p.Body,
        "signature": p.Signature,
    }
}

func (p GetItemWithSignatureByUserIdResult) Pointer() *GetItemWithSignatureByUserIdResult {
    return &p
}

type AcquireItemSetByUserIdResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
    OverflowCount *int64 `json:"overflowCount"`
}

type AcquireItemSetByUserIdAsyncResult struct {
	result *AcquireItemSetByUserIdResult
	err    error
}

func NewAcquireItemSetByUserIdResultFromJson(data string) AcquireItemSetByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcquireItemSetByUserIdResultFromDict(dict)
}

func NewAcquireItemSetByUserIdResultFromDict(data map[string]interface{}) AcquireItemSetByUserIdResult {
    return AcquireItemSetByUserIdResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
        OverflowCount: core.CastInt64(data["overflowCount"]),
    }
}

func (p AcquireItemSetByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
        "overflowCount": p.OverflowCount,
    }
}

func (p AcquireItemSetByUserIdResult) Pointer() *AcquireItemSetByUserIdResult {
    return &p
}

type ConsumeItemSetResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type ConsumeItemSetAsyncResult struct {
	result *ConsumeItemSetResult
	err    error
}

func NewConsumeItemSetResultFromJson(data string) ConsumeItemSetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeItemSetResultFromDict(dict)
}

func NewConsumeItemSetResultFromDict(data map[string]interface{}) ConsumeItemSetResult {
    return ConsumeItemSetResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p ConsumeItemSetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p ConsumeItemSetResult) Pointer() *ConsumeItemSetResult {
    return &p
}

type ConsumeItemSetByUserIdResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type ConsumeItemSetByUserIdAsyncResult struct {
	result *ConsumeItemSetByUserIdResult
	err    error
}

func NewConsumeItemSetByUserIdResultFromJson(data string) ConsumeItemSetByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeItemSetByUserIdResultFromDict(dict)
}

func NewConsumeItemSetByUserIdResultFromDict(data map[string]interface{}) ConsumeItemSetByUserIdResult {
    return ConsumeItemSetByUserIdResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p ConsumeItemSetByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p ConsumeItemSetByUserIdResult) Pointer() *ConsumeItemSetByUserIdResult {
    return &p
}

type DeleteItemSetByUserIdResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type DeleteItemSetByUserIdAsyncResult struct {
	result *DeleteItemSetByUserIdResult
	err    error
}

func NewDeleteItemSetByUserIdResultFromJson(data string) DeleteItemSetByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteItemSetByUserIdResultFromDict(dict)
}

func NewDeleteItemSetByUserIdResultFromDict(data map[string]interface{}) DeleteItemSetByUserIdResult {
    return DeleteItemSetByUserIdResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p DeleteItemSetByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p DeleteItemSetByUserIdResult) Pointer() *DeleteItemSetByUserIdResult {
    return &p
}

type AcquireItemSetByStampSheetResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
    OverflowCount *int64 `json:"overflowCount"`
}

type AcquireItemSetByStampSheetAsyncResult struct {
	result *AcquireItemSetByStampSheetResult
	err    error
}

func NewAcquireItemSetByStampSheetResultFromJson(data string) AcquireItemSetByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcquireItemSetByStampSheetResultFromDict(dict)
}

func NewAcquireItemSetByStampSheetResultFromDict(data map[string]interface{}) AcquireItemSetByStampSheetResult {
    return AcquireItemSetByStampSheetResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
        OverflowCount: core.CastInt64(data["overflowCount"]),
    }
}

func (p AcquireItemSetByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
        "overflowCount": p.OverflowCount,
    }
}

func (p AcquireItemSetByStampSheetResult) Pointer() *AcquireItemSetByStampSheetResult {
    return &p
}

type ConsumeItemSetByStampTaskResult struct {
    Items []ItemSet `json:"items"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
    NewContextStack *string `json:"newContextStack"`
}

type ConsumeItemSetByStampTaskAsyncResult struct {
	result *ConsumeItemSetByStampTaskResult
	err    error
}

func NewConsumeItemSetByStampTaskResultFromJson(data string) ConsumeItemSetByStampTaskResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeItemSetByStampTaskResultFromDict(dict)
}

func NewConsumeItemSetByStampTaskResultFromDict(data map[string]interface{}) ConsumeItemSetByStampTaskResult {
    return ConsumeItemSetByStampTaskResult {
        Items: CastItemSets(core.CastArray(data["items"])),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p ConsumeItemSetByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastItemSetsFromDict(
            p.Items,
        ),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p ConsumeItemSetByStampTaskResult) Pointer() *ConsumeItemSetByStampTaskResult {
    return &p
}

type DescribeReferenceOfResult struct {
    Items []*string `json:"items"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type DescribeReferenceOfAsyncResult struct {
	result *DescribeReferenceOfResult
	err    error
}

func NewDescribeReferenceOfResultFromJson(data string) DescribeReferenceOfResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeReferenceOfResultFromDict(dict)
}

func NewDescribeReferenceOfResultFromDict(data map[string]interface{}) DescribeReferenceOfResult {
    return DescribeReferenceOfResult {
        Items: core.CastStrings(core.CastArray(data["items"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p DescribeReferenceOfResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": core.CastStringsFromDict(
            p.Items,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p DescribeReferenceOfResult) Pointer() *DescribeReferenceOfResult {
    return &p
}

type DescribeReferenceOfByUserIdResult struct {
    Items []*string `json:"items"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type DescribeReferenceOfByUserIdAsyncResult struct {
	result *DescribeReferenceOfByUserIdResult
	err    error
}

func NewDescribeReferenceOfByUserIdResultFromJson(data string) DescribeReferenceOfByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeReferenceOfByUserIdResultFromDict(dict)
}

func NewDescribeReferenceOfByUserIdResultFromDict(data map[string]interface{}) DescribeReferenceOfByUserIdResult {
    return DescribeReferenceOfByUserIdResult {
        Items: core.CastStrings(core.CastArray(data["items"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p DescribeReferenceOfByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": core.CastStringsFromDict(
            p.Items,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p DescribeReferenceOfByUserIdResult) Pointer() *DescribeReferenceOfByUserIdResult {
    return &p
}

type GetReferenceOfResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type GetReferenceOfAsyncResult struct {
	result *GetReferenceOfResult
	err    error
}

func NewGetReferenceOfResultFromJson(data string) GetReferenceOfResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetReferenceOfResultFromDict(dict)
}

func NewGetReferenceOfResultFromDict(data map[string]interface{}) GetReferenceOfResult {
    return GetReferenceOfResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p GetReferenceOfResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p GetReferenceOfResult) Pointer() *GetReferenceOfResult {
    return &p
}

type GetReferenceOfByUserIdResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type GetReferenceOfByUserIdAsyncResult struct {
	result *GetReferenceOfByUserIdResult
	err    error
}

func NewGetReferenceOfByUserIdResultFromJson(data string) GetReferenceOfByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetReferenceOfByUserIdResultFromDict(dict)
}

func NewGetReferenceOfByUserIdResultFromDict(data map[string]interface{}) GetReferenceOfByUserIdResult {
    return GetReferenceOfByUserIdResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p GetReferenceOfByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p GetReferenceOfByUserIdResult) Pointer() *GetReferenceOfByUserIdResult {
    return &p
}

type VerifyReferenceOfResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type VerifyReferenceOfAsyncResult struct {
	result *VerifyReferenceOfResult
	err    error
}

func NewVerifyReferenceOfResultFromJson(data string) VerifyReferenceOfResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyReferenceOfResultFromDict(dict)
}

func NewVerifyReferenceOfResultFromDict(data map[string]interface{}) VerifyReferenceOfResult {
    return VerifyReferenceOfResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p VerifyReferenceOfResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p VerifyReferenceOfResult) Pointer() *VerifyReferenceOfResult {
    return &p
}

type VerifyReferenceOfByUserIdResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type VerifyReferenceOfByUserIdAsyncResult struct {
	result *VerifyReferenceOfByUserIdResult
	err    error
}

func NewVerifyReferenceOfByUserIdResultFromJson(data string) VerifyReferenceOfByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyReferenceOfByUserIdResultFromDict(dict)
}

func NewVerifyReferenceOfByUserIdResultFromDict(data map[string]interface{}) VerifyReferenceOfByUserIdResult {
    return VerifyReferenceOfByUserIdResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p VerifyReferenceOfByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p VerifyReferenceOfByUserIdResult) Pointer() *VerifyReferenceOfByUserIdResult {
    return &p
}

type AddReferenceOfResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type AddReferenceOfAsyncResult struct {
	result *AddReferenceOfResult
	err    error
}

func NewAddReferenceOfResultFromJson(data string) AddReferenceOfResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddReferenceOfResultFromDict(dict)
}

func NewAddReferenceOfResultFromDict(data map[string]interface{}) AddReferenceOfResult {
    return AddReferenceOfResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p AddReferenceOfResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p AddReferenceOfResult) Pointer() *AddReferenceOfResult {
    return &p
}

type AddReferenceOfByUserIdResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type AddReferenceOfByUserIdAsyncResult struct {
	result *AddReferenceOfByUserIdResult
	err    error
}

func NewAddReferenceOfByUserIdResultFromJson(data string) AddReferenceOfByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddReferenceOfByUserIdResultFromDict(dict)
}

func NewAddReferenceOfByUserIdResultFromDict(data map[string]interface{}) AddReferenceOfByUserIdResult {
    return AddReferenceOfByUserIdResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p AddReferenceOfByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p AddReferenceOfByUserIdResult) Pointer() *AddReferenceOfByUserIdResult {
    return &p
}

type DeleteReferenceOfResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type DeleteReferenceOfAsyncResult struct {
	result *DeleteReferenceOfResult
	err    error
}

func NewDeleteReferenceOfResultFromJson(data string) DeleteReferenceOfResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteReferenceOfResultFromDict(dict)
}

func NewDeleteReferenceOfResultFromDict(data map[string]interface{}) DeleteReferenceOfResult {
    return DeleteReferenceOfResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p DeleteReferenceOfResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p DeleteReferenceOfResult) Pointer() *DeleteReferenceOfResult {
    return &p
}

type DeleteReferenceOfByUserIdResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type DeleteReferenceOfByUserIdAsyncResult struct {
	result *DeleteReferenceOfByUserIdResult
	err    error
}

func NewDeleteReferenceOfByUserIdResultFromJson(data string) DeleteReferenceOfByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteReferenceOfByUserIdResultFromDict(dict)
}

func NewDeleteReferenceOfByUserIdResultFromDict(data map[string]interface{}) DeleteReferenceOfByUserIdResult {
    return DeleteReferenceOfByUserIdResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p DeleteReferenceOfByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p DeleteReferenceOfByUserIdResult) Pointer() *DeleteReferenceOfByUserIdResult {
    return &p
}

type AddReferenceOfItemSetByStampSheetResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type AddReferenceOfItemSetByStampSheetAsyncResult struct {
	result *AddReferenceOfItemSetByStampSheetResult
	err    error
}

func NewAddReferenceOfItemSetByStampSheetResultFromJson(data string) AddReferenceOfItemSetByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddReferenceOfItemSetByStampSheetResultFromDict(dict)
}

func NewAddReferenceOfItemSetByStampSheetResultFromDict(data map[string]interface{}) AddReferenceOfItemSetByStampSheetResult {
    return AddReferenceOfItemSetByStampSheetResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p AddReferenceOfItemSetByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p AddReferenceOfItemSetByStampSheetResult) Pointer() *AddReferenceOfItemSetByStampSheetResult {
    return &p
}

type DeleteReferenceOfItemSetByStampSheetResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
}

type DeleteReferenceOfItemSetByStampSheetAsyncResult struct {
	result *DeleteReferenceOfItemSetByStampSheetResult
	err    error
}

func NewDeleteReferenceOfItemSetByStampSheetResultFromJson(data string) DeleteReferenceOfItemSetByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteReferenceOfItemSetByStampSheetResultFromDict(dict)
}

func NewDeleteReferenceOfItemSetByStampSheetResultFromDict(data map[string]interface{}) DeleteReferenceOfItemSetByStampSheetResult {
    return DeleteReferenceOfItemSetByStampSheetResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
    }
}

func (p DeleteReferenceOfItemSetByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
    }
}

func (p DeleteReferenceOfItemSetByStampSheetResult) Pointer() *DeleteReferenceOfItemSetByStampSheetResult {
    return &p
}

type VerifyReferenceOfByStampTaskResult struct {
    Item []*string `json:"item"`
    ItemSet *ItemSet `json:"itemSet"`
    ItemModel *ItemModel `json:"itemModel"`
    Inventory *Inventory `json:"inventory"`
    NewContextStack *string `json:"newContextStack"`
}

type VerifyReferenceOfByStampTaskAsyncResult struct {
	result *VerifyReferenceOfByStampTaskResult
	err    error
}

func NewVerifyReferenceOfByStampTaskResultFromJson(data string) VerifyReferenceOfByStampTaskResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyReferenceOfByStampTaskResultFromDict(dict)
}

func NewVerifyReferenceOfByStampTaskResultFromDict(data map[string]interface{}) VerifyReferenceOfByStampTaskResult {
    return VerifyReferenceOfByStampTaskResult {
        Item: core.CastStrings(core.CastArray(data["item"])),
        ItemSet: NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
        ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
        Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p VerifyReferenceOfByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": core.CastStringsFromDict(
            p.Item,
        ),
        "itemSet": p.ItemSet.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
        "inventory": p.Inventory.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p VerifyReferenceOfByStampTaskResult) Pointer() *VerifyReferenceOfByStampTaskResult {
    return &p
}

type DescribeSimpleItemsResult struct {
    Items []SimpleItem `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSimpleItemsAsyncResult struct {
	result *DescribeSimpleItemsResult
	err    error
}

func NewDescribeSimpleItemsResultFromJson(data string) DescribeSimpleItemsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSimpleItemsResultFromDict(dict)
}

func NewDescribeSimpleItemsResultFromDict(data map[string]interface{}) DescribeSimpleItemsResult {
    return DescribeSimpleItemsResult {
        Items: CastSimpleItems(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSimpleItemsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleItemsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSimpleItemsResult) Pointer() *DescribeSimpleItemsResult {
    return &p
}

type DescribeSimpleItemsByUserIdResult struct {
    Items []SimpleItem `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSimpleItemsByUserIdAsyncResult struct {
	result *DescribeSimpleItemsByUserIdResult
	err    error
}

func NewDescribeSimpleItemsByUserIdResultFromJson(data string) DescribeSimpleItemsByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSimpleItemsByUserIdResultFromDict(dict)
}

func NewDescribeSimpleItemsByUserIdResultFromDict(data map[string]interface{}) DescribeSimpleItemsByUserIdResult {
    return DescribeSimpleItemsByUserIdResult {
        Items: CastSimpleItems(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSimpleItemsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleItemsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSimpleItemsByUserIdResult) Pointer() *DescribeSimpleItemsByUserIdResult {
    return &p
}

type GetSimpleItemResult struct {
    Item *SimpleItem `json:"item"`
    ItemModel *SimpleItemModel `json:"itemModel"`
}

type GetSimpleItemAsyncResult struct {
	result *GetSimpleItemResult
	err    error
}

func NewGetSimpleItemResultFromJson(data string) GetSimpleItemResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSimpleItemResultFromDict(dict)
}

func NewGetSimpleItemResultFromDict(data map[string]interface{}) GetSimpleItemResult {
    return GetSimpleItemResult {
        Item: NewSimpleItemFromDict(core.CastMap(data["item"])).Pointer(),
        ItemModel: NewSimpleItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
    }
}

func (p GetSimpleItemResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
    }
}

func (p GetSimpleItemResult) Pointer() *GetSimpleItemResult {
    return &p
}

type GetSimpleItemByUserIdResult struct {
    Item *SimpleItem `json:"item"`
    ItemModel *SimpleItemModel `json:"itemModel"`
}

type GetSimpleItemByUserIdAsyncResult struct {
	result *GetSimpleItemByUserIdResult
	err    error
}

func NewGetSimpleItemByUserIdResultFromJson(data string) GetSimpleItemByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSimpleItemByUserIdResultFromDict(dict)
}

func NewGetSimpleItemByUserIdResultFromDict(data map[string]interface{}) GetSimpleItemByUserIdResult {
    return GetSimpleItemByUserIdResult {
        Item: NewSimpleItemFromDict(core.CastMap(data["item"])).Pointer(),
        ItemModel: NewSimpleItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
    }
}

func (p GetSimpleItemByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
    }
}

func (p GetSimpleItemByUserIdResult) Pointer() *GetSimpleItemByUserIdResult {
    return &p
}

type GetSimpleItemWithSignatureResult struct {
    Item *SimpleItem `json:"item"`
    SimpleItemModel *SimpleItemModel `json:"simpleItemModel"`
    Body *string `json:"body"`
    Signature *string `json:"signature"`
}

type GetSimpleItemWithSignatureAsyncResult struct {
	result *GetSimpleItemWithSignatureResult
	err    error
}

func NewGetSimpleItemWithSignatureResultFromJson(data string) GetSimpleItemWithSignatureResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSimpleItemWithSignatureResultFromDict(dict)
}

func NewGetSimpleItemWithSignatureResultFromDict(data map[string]interface{}) GetSimpleItemWithSignatureResult {
    return GetSimpleItemWithSignatureResult {
        Item: NewSimpleItemFromDict(core.CastMap(data["item"])).Pointer(),
        SimpleItemModel: NewSimpleItemModelFromDict(core.CastMap(data["simpleItemModel"])).Pointer(),
        Body: core.CastString(data["body"]),
        Signature: core.CastString(data["signature"]),
    }
}

func (p GetSimpleItemWithSignatureResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "simpleItemModel": p.SimpleItemModel.ToDict(),
        "body": p.Body,
        "signature": p.Signature,
    }
}

func (p GetSimpleItemWithSignatureResult) Pointer() *GetSimpleItemWithSignatureResult {
    return &p
}

type GetSimpleItemWithSignatureByUserIdResult struct {
    Item *SimpleItem `json:"item"`
    SimpleItemModel *SimpleItemModel `json:"simpleItemModel"`
    Body *string `json:"body"`
    Signature *string `json:"signature"`
}

type GetSimpleItemWithSignatureByUserIdAsyncResult struct {
	result *GetSimpleItemWithSignatureByUserIdResult
	err    error
}

func NewGetSimpleItemWithSignatureByUserIdResultFromJson(data string) GetSimpleItemWithSignatureByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSimpleItemWithSignatureByUserIdResultFromDict(dict)
}

func NewGetSimpleItemWithSignatureByUserIdResultFromDict(data map[string]interface{}) GetSimpleItemWithSignatureByUserIdResult {
    return GetSimpleItemWithSignatureByUserIdResult {
        Item: NewSimpleItemFromDict(core.CastMap(data["item"])).Pointer(),
        SimpleItemModel: NewSimpleItemModelFromDict(core.CastMap(data["simpleItemModel"])).Pointer(),
        Body: core.CastString(data["body"]),
        Signature: core.CastString(data["signature"]),
    }
}

func (p GetSimpleItemWithSignatureByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "simpleItemModel": p.SimpleItemModel.ToDict(),
        "body": p.Body,
        "signature": p.Signature,
    }
}

func (p GetSimpleItemWithSignatureByUserIdResult) Pointer() *GetSimpleItemWithSignatureByUserIdResult {
    return &p
}

type AcquireSimpleItemsByUserIdResult struct {
    Items []SimpleItem `json:"items"`
}

type AcquireSimpleItemsByUserIdAsyncResult struct {
	result *AcquireSimpleItemsByUserIdResult
	err    error
}

func NewAcquireSimpleItemsByUserIdResultFromJson(data string) AcquireSimpleItemsByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcquireSimpleItemsByUserIdResultFromDict(dict)
}

func NewAcquireSimpleItemsByUserIdResultFromDict(data map[string]interface{}) AcquireSimpleItemsByUserIdResult {
    return AcquireSimpleItemsByUserIdResult {
        Items: CastSimpleItems(core.CastArray(data["items"])),
    }
}

func (p AcquireSimpleItemsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleItemsFromDict(
            p.Items,
        ),
    }
}

func (p AcquireSimpleItemsByUserIdResult) Pointer() *AcquireSimpleItemsByUserIdResult {
    return &p
}

type ConsumeSimpleItemsResult struct {
    Items []SimpleItem `json:"items"`
}

type ConsumeSimpleItemsAsyncResult struct {
	result *ConsumeSimpleItemsResult
	err    error
}

func NewConsumeSimpleItemsResultFromJson(data string) ConsumeSimpleItemsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeSimpleItemsResultFromDict(dict)
}

func NewConsumeSimpleItemsResultFromDict(data map[string]interface{}) ConsumeSimpleItemsResult {
    return ConsumeSimpleItemsResult {
        Items: CastSimpleItems(core.CastArray(data["items"])),
    }
}

func (p ConsumeSimpleItemsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleItemsFromDict(
            p.Items,
        ),
    }
}

func (p ConsumeSimpleItemsResult) Pointer() *ConsumeSimpleItemsResult {
    return &p
}

type ConsumeSimpleItemsByUserIdResult struct {
    Items []SimpleItem `json:"items"`
}

type ConsumeSimpleItemsByUserIdAsyncResult struct {
	result *ConsumeSimpleItemsByUserIdResult
	err    error
}

func NewConsumeSimpleItemsByUserIdResultFromJson(data string) ConsumeSimpleItemsByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeSimpleItemsByUserIdResultFromDict(dict)
}

func NewConsumeSimpleItemsByUserIdResultFromDict(data map[string]interface{}) ConsumeSimpleItemsByUserIdResult {
    return ConsumeSimpleItemsByUserIdResult {
        Items: CastSimpleItems(core.CastArray(data["items"])),
    }
}

func (p ConsumeSimpleItemsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleItemsFromDict(
            p.Items,
        ),
    }
}

func (p ConsumeSimpleItemsByUserIdResult) Pointer() *ConsumeSimpleItemsByUserIdResult {
    return &p
}

type DeleteSimpleItemsByUserIdResult struct {
}

type DeleteSimpleItemsByUserIdAsyncResult struct {
	result *DeleteSimpleItemsByUserIdResult
	err    error
}

func NewDeleteSimpleItemsByUserIdResultFromJson(data string) DeleteSimpleItemsByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteSimpleItemsByUserIdResultFromDict(dict)
}

func NewDeleteSimpleItemsByUserIdResultFromDict(data map[string]interface{}) DeleteSimpleItemsByUserIdResult {
    return DeleteSimpleItemsByUserIdResult {
    }
}

func (p DeleteSimpleItemsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteSimpleItemsByUserIdResult) Pointer() *DeleteSimpleItemsByUserIdResult {
    return &p
}

type AcquireSimpleItemsByStampSheetResult struct {
    Items []SimpleItem `json:"items"`
}

type AcquireSimpleItemsByStampSheetAsyncResult struct {
	result *AcquireSimpleItemsByStampSheetResult
	err    error
}

func NewAcquireSimpleItemsByStampSheetResultFromJson(data string) AcquireSimpleItemsByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcquireSimpleItemsByStampSheetResultFromDict(dict)
}

func NewAcquireSimpleItemsByStampSheetResultFromDict(data map[string]interface{}) AcquireSimpleItemsByStampSheetResult {
    return AcquireSimpleItemsByStampSheetResult {
        Items: CastSimpleItems(core.CastArray(data["items"])),
    }
}

func (p AcquireSimpleItemsByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleItemsFromDict(
            p.Items,
        ),
    }
}

func (p AcquireSimpleItemsByStampSheetResult) Pointer() *AcquireSimpleItemsByStampSheetResult {
    return &p
}

type ConsumeSimpleItemsByStampTaskResult struct {
    Items []SimpleItem `json:"items"`
    NewContextStack *string `json:"newContextStack"`
}

type ConsumeSimpleItemsByStampTaskAsyncResult struct {
	result *ConsumeSimpleItemsByStampTaskResult
	err    error
}

func NewConsumeSimpleItemsByStampTaskResultFromJson(data string) ConsumeSimpleItemsByStampTaskResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeSimpleItemsByStampTaskResultFromDict(dict)
}

func NewConsumeSimpleItemsByStampTaskResultFromDict(data map[string]interface{}) ConsumeSimpleItemsByStampTaskResult {
    return ConsumeSimpleItemsByStampTaskResult {
        Items: CastSimpleItems(core.CastArray(data["items"])),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p ConsumeSimpleItemsByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSimpleItemsFromDict(
            p.Items,
        ),
        "newContextStack": p.NewContextStack,
    }
}

func (p ConsumeSimpleItemsByStampTaskResult) Pointer() *ConsumeSimpleItemsByStampTaskResult {
    return &p
}

type DescribeBigItemsResult struct {
    Items []BigItem `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeBigItemsAsyncResult struct {
	result *DescribeBigItemsResult
	err    error
}

func NewDescribeBigItemsResultFromJson(data string) DescribeBigItemsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBigItemsResultFromDict(dict)
}

func NewDescribeBigItemsResultFromDict(data map[string]interface{}) DescribeBigItemsResult {
    return DescribeBigItemsResult {
        Items: CastBigItems(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeBigItemsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBigItemsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeBigItemsResult) Pointer() *DescribeBigItemsResult {
    return &p
}

type DescribeBigItemsByUserIdResult struct {
    Items []BigItem `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeBigItemsByUserIdAsyncResult struct {
	result *DescribeBigItemsByUserIdResult
	err    error
}

func NewDescribeBigItemsByUserIdResultFromJson(data string) DescribeBigItemsByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBigItemsByUserIdResultFromDict(dict)
}

func NewDescribeBigItemsByUserIdResultFromDict(data map[string]interface{}) DescribeBigItemsByUserIdResult {
    return DescribeBigItemsByUserIdResult {
        Items: CastBigItems(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeBigItemsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBigItemsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeBigItemsByUserIdResult) Pointer() *DescribeBigItemsByUserIdResult {
    return &p
}

type GetBigItemResult struct {
    Item *BigItem `json:"item"`
    ItemModel *BigItemModel `json:"itemModel"`
}

type GetBigItemAsyncResult struct {
	result *GetBigItemResult
	err    error
}

func NewGetBigItemResultFromJson(data string) GetBigItemResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBigItemResultFromDict(dict)
}

func NewGetBigItemResultFromDict(data map[string]interface{}) GetBigItemResult {
    return GetBigItemResult {
        Item: NewBigItemFromDict(core.CastMap(data["item"])).Pointer(),
        ItemModel: NewBigItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
    }
}

func (p GetBigItemResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
    }
}

func (p GetBigItemResult) Pointer() *GetBigItemResult {
    return &p
}

type GetBigItemByUserIdResult struct {
    Item *BigItem `json:"item"`
    ItemModel *BigItemModel `json:"itemModel"`
}

type GetBigItemByUserIdAsyncResult struct {
	result *GetBigItemByUserIdResult
	err    error
}

func NewGetBigItemByUserIdResultFromJson(data string) GetBigItemByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBigItemByUserIdResultFromDict(dict)
}

func NewGetBigItemByUserIdResultFromDict(data map[string]interface{}) GetBigItemByUserIdResult {
    return GetBigItemByUserIdResult {
        Item: NewBigItemFromDict(core.CastMap(data["item"])).Pointer(),
        ItemModel: NewBigItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
    }
}

func (p GetBigItemByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "itemModel": p.ItemModel.ToDict(),
    }
}

func (p GetBigItemByUserIdResult) Pointer() *GetBigItemByUserIdResult {
    return &p
}

type AcquireBigItemByUserIdResult struct {
    Item *BigItem `json:"item"`
}

type AcquireBigItemByUserIdAsyncResult struct {
	result *AcquireBigItemByUserIdResult
	err    error
}

func NewAcquireBigItemByUserIdResultFromJson(data string) AcquireBigItemByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcquireBigItemByUserIdResultFromDict(dict)
}

func NewAcquireBigItemByUserIdResultFromDict(data map[string]interface{}) AcquireBigItemByUserIdResult {
    return AcquireBigItemByUserIdResult {
        Item: NewBigItemFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AcquireBigItemByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AcquireBigItemByUserIdResult) Pointer() *AcquireBigItemByUserIdResult {
    return &p
}

type ConsumeBigItemResult struct {
    Item *BigItem `json:"item"`
}

type ConsumeBigItemAsyncResult struct {
	result *ConsumeBigItemResult
	err    error
}

func NewConsumeBigItemResultFromJson(data string) ConsumeBigItemResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeBigItemResultFromDict(dict)
}

func NewConsumeBigItemResultFromDict(data map[string]interface{}) ConsumeBigItemResult {
    return ConsumeBigItemResult {
        Item: NewBigItemFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p ConsumeBigItemResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p ConsumeBigItemResult) Pointer() *ConsumeBigItemResult {
    return &p
}

type ConsumeBigItemByUserIdResult struct {
    Item *BigItem `json:"item"`
}

type ConsumeBigItemByUserIdAsyncResult struct {
	result *ConsumeBigItemByUserIdResult
	err    error
}

func NewConsumeBigItemByUserIdResultFromJson(data string) ConsumeBigItemByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeBigItemByUserIdResultFromDict(dict)
}

func NewConsumeBigItemByUserIdResultFromDict(data map[string]interface{}) ConsumeBigItemByUserIdResult {
    return ConsumeBigItemByUserIdResult {
        Item: NewBigItemFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p ConsumeBigItemByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p ConsumeBigItemByUserIdResult) Pointer() *ConsumeBigItemByUserIdResult {
    return &p
}

type DeleteBigItemByUserIdResult struct {
    Item *BigItem `json:"item"`
}

type DeleteBigItemByUserIdAsyncResult struct {
	result *DeleteBigItemByUserIdResult
	err    error
}

func NewDeleteBigItemByUserIdResultFromJson(data string) DeleteBigItemByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteBigItemByUserIdResultFromDict(dict)
}

func NewDeleteBigItemByUserIdResultFromDict(data map[string]interface{}) DeleteBigItemByUserIdResult {
    return DeleteBigItemByUserIdResult {
        Item: NewBigItemFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteBigItemByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteBigItemByUserIdResult) Pointer() *DeleteBigItemByUserIdResult {
    return &p
}

type AcquireBigItemByStampSheetResult struct {
    Item *BigItem `json:"item"`
}

type AcquireBigItemByStampSheetAsyncResult struct {
	result *AcquireBigItemByStampSheetResult
	err    error
}

func NewAcquireBigItemByStampSheetResultFromJson(data string) AcquireBigItemByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcquireBigItemByStampSheetResultFromDict(dict)
}

func NewAcquireBigItemByStampSheetResultFromDict(data map[string]interface{}) AcquireBigItemByStampSheetResult {
    return AcquireBigItemByStampSheetResult {
        Item: NewBigItemFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AcquireBigItemByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AcquireBigItemByStampSheetResult) Pointer() *AcquireBigItemByStampSheetResult {
    return &p
}

type ConsumeBigItemByStampTaskResult struct {
    Item *BigItem `json:"item"`
    NewContextStack *string `json:"newContextStack"`
}

type ConsumeBigItemByStampTaskAsyncResult struct {
	result *ConsumeBigItemByStampTaskResult
	err    error
}

func NewConsumeBigItemByStampTaskResultFromJson(data string) ConsumeBigItemByStampTaskResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeBigItemByStampTaskResultFromDict(dict)
}

func NewConsumeBigItemByStampTaskResultFromDict(data map[string]interface{}) ConsumeBigItemByStampTaskResult {
    return ConsumeBigItemByStampTaskResult {
        Item: NewBigItemFromDict(core.CastMap(data["item"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p ConsumeBigItemByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p ConsumeBigItemByStampTaskResult) Pointer() *ConsumeBigItemByStampTaskResult {
    return &p
}