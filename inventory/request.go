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

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeNamespacesRequestFromDict(dict)
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
    return DescribeNamespacesRequest {
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
    return &p
}

type CreateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    AcquireScript *ScriptSetting `json:"acquireScript"`
    OverflowScript *ScriptSetting `json:"overflowScript"`
    ConsumeScript *ScriptSetting `json:"consumeScript"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        AcquireScript: NewScriptSettingFromDict(core.CastMap(data["acquireScript"])).Pointer(),
        OverflowScript: NewScriptSettingFromDict(core.CastMap(data["overflowScript"])).Pointer(),
        ConsumeScript: NewScriptSettingFromDict(core.CastMap(data["consumeScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "acquireScript": p.AcquireScript.ToDict(),
        "overflowScript": p.OverflowScript.ToDict(),
        "consumeScript": p.ConsumeScript.ToDict(),
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
    return &p
}

type GetNamespaceStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceStatusRequestFromDict(dict)
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
    return GetNamespaceStatusRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
    return &p
}

type GetNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceRequestFromDict(dict)
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
    return GetNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
    return &p
}

type UpdateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Description *string `json:"description"`
    AcquireScript *ScriptSetting `json:"acquireScript"`
    OverflowScript *ScriptSetting `json:"overflowScript"`
    ConsumeScript *ScriptSetting `json:"consumeScript"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        AcquireScript: NewScriptSettingFromDict(core.CastMap(data["acquireScript"])).Pointer(),
        OverflowScript: NewScriptSettingFromDict(core.CastMap(data["overflowScript"])).Pointer(),
        ConsumeScript: NewScriptSettingFromDict(core.CastMap(data["consumeScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "acquireScript": p.AcquireScript.ToDict(),
        "overflowScript": p.OverflowScript.ToDict(),
        "consumeScript": p.ConsumeScript.ToDict(),
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
    return &p
}

type DeleteNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteNamespaceRequestFromDict(dict)
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
    return DeleteNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
    return &p
}

type DescribeInventoryModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeInventoryModelMastersRequestFromJson(data string) DescribeInventoryModelMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoryModelMastersRequestFromDict(dict)
}

func NewDescribeInventoryModelMastersRequestFromDict(data map[string]interface{}) DescribeInventoryModelMastersRequest {
    return DescribeInventoryModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeInventoryModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeInventoryModelMastersRequest) Pointer() *DescribeInventoryModelMastersRequest {
    return &p
}

type CreateInventoryModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    InitialCapacity *int32 `json:"initialCapacity"`
    MaxCapacity *int32 `json:"maxCapacity"`
    ProtectReferencedItem *bool `json:"protectReferencedItem"`
}

func NewCreateInventoryModelMasterRequestFromJson(data string) CreateInventoryModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateInventoryModelMasterRequestFromDict(dict)
}

func NewCreateInventoryModelMasterRequestFromDict(data map[string]interface{}) CreateInventoryModelMasterRequest {
    return CreateInventoryModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        InitialCapacity: core.CastInt32(data["initialCapacity"]),
        MaxCapacity: core.CastInt32(data["maxCapacity"]),
        ProtectReferencedItem: core.CastBool(data["protectReferencedItem"]),
    }
}

func (p CreateInventoryModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "initialCapacity": p.InitialCapacity,
        "maxCapacity": p.MaxCapacity,
        "protectReferencedItem": p.ProtectReferencedItem,
    }
}

func (p CreateInventoryModelMasterRequest) Pointer() *CreateInventoryModelMasterRequest {
    return &p
}

type GetInventoryModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
}

func NewGetInventoryModelMasterRequestFromJson(data string) GetInventoryModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInventoryModelMasterRequestFromDict(dict)
}

func NewGetInventoryModelMasterRequestFromDict(data map[string]interface{}) GetInventoryModelMasterRequest {
    return GetInventoryModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
    }
}

func (p GetInventoryModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
    }
}

func (p GetInventoryModelMasterRequest) Pointer() *GetInventoryModelMasterRequest {
    return &p
}

type UpdateInventoryModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    InitialCapacity *int32 `json:"initialCapacity"`
    MaxCapacity *int32 `json:"maxCapacity"`
    ProtectReferencedItem *bool `json:"protectReferencedItem"`
}

func NewUpdateInventoryModelMasterRequestFromJson(data string) UpdateInventoryModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateInventoryModelMasterRequestFromDict(dict)
}

func NewUpdateInventoryModelMasterRequestFromDict(data map[string]interface{}) UpdateInventoryModelMasterRequest {
    return UpdateInventoryModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        InitialCapacity: core.CastInt32(data["initialCapacity"]),
        MaxCapacity: core.CastInt32(data["maxCapacity"]),
        ProtectReferencedItem: core.CastBool(data["protectReferencedItem"]),
    }
}

func (p UpdateInventoryModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "description": p.Description,
        "metadata": p.Metadata,
        "initialCapacity": p.InitialCapacity,
        "maxCapacity": p.MaxCapacity,
        "protectReferencedItem": p.ProtectReferencedItem,
    }
}

func (p UpdateInventoryModelMasterRequest) Pointer() *UpdateInventoryModelMasterRequest {
    return &p
}

type DeleteInventoryModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
}

func NewDeleteInventoryModelMasterRequestFromJson(data string) DeleteInventoryModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteInventoryModelMasterRequestFromDict(dict)
}

func NewDeleteInventoryModelMasterRequestFromDict(data map[string]interface{}) DeleteInventoryModelMasterRequest {
    return DeleteInventoryModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
    }
}

func (p DeleteInventoryModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
    }
}

func (p DeleteInventoryModelMasterRequest) Pointer() *DeleteInventoryModelMasterRequest {
    return &p
}

type DescribeInventoryModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeInventoryModelsRequestFromJson(data string) DescribeInventoryModelsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoryModelsRequestFromDict(dict)
}

func NewDescribeInventoryModelsRequestFromDict(data map[string]interface{}) DescribeInventoryModelsRequest {
    return DescribeInventoryModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeInventoryModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeInventoryModelsRequest) Pointer() *DescribeInventoryModelsRequest {
    return &p
}

type GetInventoryModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
}

func NewGetInventoryModelRequestFromJson(data string) GetInventoryModelRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInventoryModelRequestFromDict(dict)
}

func NewGetInventoryModelRequestFromDict(data map[string]interface{}) GetInventoryModelRequest {
    return GetInventoryModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
    }
}

func (p GetInventoryModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
    }
}

func (p GetInventoryModelRequest) Pointer() *GetInventoryModelRequest {
    return &p
}

type DescribeItemModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeItemModelMastersRequestFromJson(data string) DescribeItemModelMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeItemModelMastersRequestFromDict(dict)
}

func NewDescribeItemModelMastersRequestFromDict(data map[string]interface{}) DescribeItemModelMastersRequest {
    return DescribeItemModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeItemModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeItemModelMastersRequest) Pointer() *DescribeItemModelMastersRequest {
    return &p
}

type CreateItemModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    StackingLimit *int64 `json:"stackingLimit"`
    AllowMultipleStacks *bool `json:"allowMultipleStacks"`
    SortValue *int32 `json:"sortValue"`
}

func NewCreateItemModelMasterRequestFromJson(data string) CreateItemModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateItemModelMasterRequestFromDict(dict)
}

func NewCreateItemModelMasterRequestFromDict(data map[string]interface{}) CreateItemModelMasterRequest {
    return CreateItemModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        StackingLimit: core.CastInt64(data["stackingLimit"]),
        AllowMultipleStacks: core.CastBool(data["allowMultipleStacks"]),
        SortValue: core.CastInt32(data["sortValue"]),
    }
}

func (p CreateItemModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "stackingLimit": p.StackingLimit,
        "allowMultipleStacks": p.AllowMultipleStacks,
        "sortValue": p.SortValue,
    }
}

func (p CreateItemModelMasterRequest) Pointer() *CreateItemModelMasterRequest {
    return &p
}

type GetItemModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    ItemName *string `json:"itemName"`
}

func NewGetItemModelMasterRequestFromJson(data string) GetItemModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemModelMasterRequestFromDict(dict)
}

func NewGetItemModelMasterRequestFromDict(data map[string]interface{}) GetItemModelMasterRequest {
    return GetItemModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        ItemName: core.CastString(data["itemName"]),
    }
}

func (p GetItemModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "itemName": p.ItemName,
    }
}

func (p GetItemModelMasterRequest) Pointer() *GetItemModelMasterRequest {
    return &p
}

type UpdateItemModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    ItemName *string `json:"itemName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    StackingLimit *int64 `json:"stackingLimit"`
    AllowMultipleStacks *bool `json:"allowMultipleStacks"`
    SortValue *int32 `json:"sortValue"`
}

func NewUpdateItemModelMasterRequestFromJson(data string) UpdateItemModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateItemModelMasterRequestFromDict(dict)
}

func NewUpdateItemModelMasterRequestFromDict(data map[string]interface{}) UpdateItemModelMasterRequest {
    return UpdateItemModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        ItemName: core.CastString(data["itemName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        StackingLimit: core.CastInt64(data["stackingLimit"]),
        AllowMultipleStacks: core.CastBool(data["allowMultipleStacks"]),
        SortValue: core.CastInt32(data["sortValue"]),
    }
}

func (p UpdateItemModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "itemName": p.ItemName,
        "description": p.Description,
        "metadata": p.Metadata,
        "stackingLimit": p.StackingLimit,
        "allowMultipleStacks": p.AllowMultipleStacks,
        "sortValue": p.SortValue,
    }
}

func (p UpdateItemModelMasterRequest) Pointer() *UpdateItemModelMasterRequest {
    return &p
}

type DeleteItemModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    ItemName *string `json:"itemName"`
}

func NewDeleteItemModelMasterRequestFromJson(data string) DeleteItemModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteItemModelMasterRequestFromDict(dict)
}

func NewDeleteItemModelMasterRequestFromDict(data map[string]interface{}) DeleteItemModelMasterRequest {
    return DeleteItemModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        ItemName: core.CastString(data["itemName"]),
    }
}

func (p DeleteItemModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "itemName": p.ItemName,
    }
}

func (p DeleteItemModelMasterRequest) Pointer() *DeleteItemModelMasterRequest {
    return &p
}

type DescribeItemModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
}

func NewDescribeItemModelsRequestFromJson(data string) DescribeItemModelsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeItemModelsRequestFromDict(dict)
}

func NewDescribeItemModelsRequestFromDict(data map[string]interface{}) DescribeItemModelsRequest {
    return DescribeItemModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
    }
}

func (p DescribeItemModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
    }
}

func (p DescribeItemModelsRequest) Pointer() *DescribeItemModelsRequest {
    return &p
}

type GetItemModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    ItemName *string `json:"itemName"`
}

func NewGetItemModelRequestFromJson(data string) GetItemModelRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemModelRequestFromDict(dict)
}

func NewGetItemModelRequestFromDict(data map[string]interface{}) GetItemModelRequest {
    return GetItemModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        ItemName: core.CastString(data["itemName"]),
    }
}

func (p GetItemModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "itemName": p.ItemName,
    }
}

func (p GetItemModelRequest) Pointer() *GetItemModelRequest {
    return &p
}

type ExportMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewExportMasterRequestFromJson(data string) ExportMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExportMasterRequestFromDict(dict)
}

func NewExportMasterRequestFromDict(data map[string]interface{}) ExportMasterRequest {
    return ExportMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p ExportMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p ExportMasterRequest) Pointer() *ExportMasterRequest {
    return &p
}

type GetCurrentItemModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentItemModelMasterRequestFromJson(data string) GetCurrentItemModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentItemModelMasterRequestFromDict(dict)
}

func NewGetCurrentItemModelMasterRequestFromDict(data map[string]interface{}) GetCurrentItemModelMasterRequest {
    return GetCurrentItemModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentItemModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentItemModelMasterRequest) Pointer() *GetCurrentItemModelMasterRequest {
    return &p
}

type UpdateCurrentItemModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentItemModelMasterRequestFromJson(data string) UpdateCurrentItemModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentItemModelMasterRequestFromDict(dict)
}

func NewUpdateCurrentItemModelMasterRequestFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterRequest {
    return UpdateCurrentItemModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentItemModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentItemModelMasterRequest) Pointer() *UpdateCurrentItemModelMasterRequest {
    return &p
}

type UpdateCurrentItemModelMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentItemModelMasterFromGitHubRequestFromJson(data string) UpdateCurrentItemModelMasterFromGitHubRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentItemModelMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentItemModelMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterFromGitHubRequest {
    return UpdateCurrentItemModelMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentItemModelMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentItemModelMasterFromGitHubRequest) Pointer() *UpdateCurrentItemModelMasterFromGitHubRequest {
    return &p
}

type DescribeInventoriesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeInventoriesRequestFromJson(data string) DescribeInventoriesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoriesRequestFromDict(dict)
}

func NewDescribeInventoriesRequestFromDict(data map[string]interface{}) DescribeInventoriesRequest {
    return DescribeInventoriesRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeInventoriesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeInventoriesRequest) Pointer() *DescribeInventoriesRequest {
    return &p
}

type DescribeInventoriesByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeInventoriesByUserIdRequestFromJson(data string) DescribeInventoriesByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoriesByUserIdRequestFromDict(dict)
}

func NewDescribeInventoriesByUserIdRequestFromDict(data map[string]interface{}) DescribeInventoriesByUserIdRequest {
    return DescribeInventoriesByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeInventoriesByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeInventoriesByUserIdRequest) Pointer() *DescribeInventoriesByUserIdRequest {
    return &p
}

type GetInventoryRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
}

func NewGetInventoryRequestFromJson(data string) GetInventoryRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInventoryRequestFromDict(dict)
}

func NewGetInventoryRequestFromDict(data map[string]interface{}) GetInventoryRequest {
    return GetInventoryRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p GetInventoryRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
    }
}

func (p GetInventoryRequest) Pointer() *GetInventoryRequest {
    return &p
}

type GetInventoryByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
}

func NewGetInventoryByUserIdRequestFromJson(data string) GetInventoryByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInventoryByUserIdRequestFromDict(dict)
}

func NewGetInventoryByUserIdRequestFromDict(data map[string]interface{}) GetInventoryByUserIdRequest {
    return GetInventoryByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p GetInventoryByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
    }
}

func (p GetInventoryByUserIdRequest) Pointer() *GetInventoryByUserIdRequest {
    return &p
}

type AddCapacityByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    AddCapacityValue *int32 `json:"addCapacityValue"`
}

func NewAddCapacityByUserIdRequestFromJson(data string) AddCapacityByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddCapacityByUserIdRequestFromDict(dict)
}

func NewAddCapacityByUserIdRequestFromDict(data map[string]interface{}) AddCapacityByUserIdRequest {
    return AddCapacityByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        AddCapacityValue: core.CastInt32(data["addCapacityValue"]),
    }
}

func (p AddCapacityByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "addCapacityValue": p.AddCapacityValue,
    }
}

func (p AddCapacityByUserIdRequest) Pointer() *AddCapacityByUserIdRequest {
    return &p
}

type SetCapacityByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    NewCapacityValue *int32 `json:"newCapacityValue"`
}

func NewSetCapacityByUserIdRequestFromJson(data string) SetCapacityByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetCapacityByUserIdRequestFromDict(dict)
}

func NewSetCapacityByUserIdRequestFromDict(data map[string]interface{}) SetCapacityByUserIdRequest {
    return SetCapacityByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        NewCapacityValue: core.CastInt32(data["newCapacityValue"]),
    }
}

func (p SetCapacityByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "newCapacityValue": p.NewCapacityValue,
    }
}

func (p SetCapacityByUserIdRequest) Pointer() *SetCapacityByUserIdRequest {
    return &p
}

type DeleteInventoryByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
}

func NewDeleteInventoryByUserIdRequestFromJson(data string) DeleteInventoryByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteInventoryByUserIdRequestFromDict(dict)
}

func NewDeleteInventoryByUserIdRequestFromDict(data map[string]interface{}) DeleteInventoryByUserIdRequest {
    return DeleteInventoryByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p DeleteInventoryByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
    }
}

func (p DeleteInventoryByUserIdRequest) Pointer() *DeleteInventoryByUserIdRequest {
    return &p
}

type AddCapacityByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewAddCapacityByStampSheetRequestFromJson(data string) AddCapacityByStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddCapacityByStampSheetRequestFromDict(dict)
}

func NewAddCapacityByStampSheetRequestFromDict(data map[string]interface{}) AddCapacityByStampSheetRequest {
    return AddCapacityByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p AddCapacityByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p AddCapacityByStampSheetRequest) Pointer() *AddCapacityByStampSheetRequest {
    return &p
}

type SetCapacityByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewSetCapacityByStampSheetRequestFromJson(data string) SetCapacityByStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetCapacityByStampSheetRequestFromDict(dict)
}

func NewSetCapacityByStampSheetRequestFromDict(data map[string]interface{}) SetCapacityByStampSheetRequest {
    return SetCapacityByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p SetCapacityByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p SetCapacityByStampSheetRequest) Pointer() *SetCapacityByStampSheetRequest {
    return &p
}

type DescribeItemSetsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeItemSetsRequestFromJson(data string) DescribeItemSetsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeItemSetsRequestFromDict(dict)
}

func NewDescribeItemSetsRequestFromDict(data map[string]interface{}) DescribeItemSetsRequest {
    return DescribeItemSetsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeItemSetsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeItemSetsRequest) Pointer() *DescribeItemSetsRequest {
    return &p
}

type DescribeItemSetsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeItemSetsByUserIdRequestFromJson(data string) DescribeItemSetsByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeItemSetsByUserIdRequestFromDict(dict)
}

func NewDescribeItemSetsByUserIdRequestFromDict(data map[string]interface{}) DescribeItemSetsByUserIdRequest {
    return DescribeItemSetsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeItemSetsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeItemSetsByUserIdRequest) Pointer() *DescribeItemSetsByUserIdRequest {
    return &p
}

type GetItemSetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
}

func NewGetItemSetRequestFromJson(data string) GetItemSetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemSetRequestFromDict(dict)
}

func NewGetItemSetRequestFromDict(data map[string]interface{}) GetItemSetRequest {
    return GetItemSetRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
    }
}

func (p GetItemSetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
    }
}

func (p GetItemSetRequest) Pointer() *GetItemSetRequest {
    return &p
}

type GetItemSetByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
}

func NewGetItemSetByUserIdRequestFromJson(data string) GetItemSetByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemSetByUserIdRequestFromDict(dict)
}

func NewGetItemSetByUserIdRequestFromDict(data map[string]interface{}) GetItemSetByUserIdRequest {
    return GetItemSetByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
    }
}

func (p GetItemSetByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
    }
}

func (p GetItemSetByUserIdRequest) Pointer() *GetItemSetByUserIdRequest {
    return &p
}

type GetItemWithSignatureRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    KeyId *string `json:"keyId"`
}

func NewGetItemWithSignatureRequestFromJson(data string) GetItemWithSignatureRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemWithSignatureRequestFromDict(dict)
}

func NewGetItemWithSignatureRequestFromDict(data map[string]interface{}) GetItemWithSignatureRequest {
    return GetItemWithSignatureRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p GetItemWithSignatureRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "keyId": p.KeyId,
    }
}

func (p GetItemWithSignatureRequest) Pointer() *GetItemWithSignatureRequest {
    return &p
}

type GetItemWithSignatureByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    KeyId *string `json:"keyId"`
}

func NewGetItemWithSignatureByUserIdRequestFromJson(data string) GetItemWithSignatureByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetItemWithSignatureByUserIdRequestFromDict(dict)
}

func NewGetItemWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetItemWithSignatureByUserIdRequest {
    return GetItemWithSignatureByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p GetItemWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "keyId": p.KeyId,
    }
}

func (p GetItemWithSignatureByUserIdRequest) Pointer() *GetItemWithSignatureByUserIdRequest {
    return &p
}

type AcquireItemSetByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    ItemName *string `json:"itemName"`
    UserId *string `json:"userId"`
    AcquireCount *int64 `json:"acquireCount"`
    ExpiresAt *int64 `json:"expiresAt"`
    CreateNewItemSet *bool `json:"createNewItemSet"`
    ItemSetName *string `json:"itemSetName"`
}

func NewAcquireItemSetByUserIdRequestFromJson(data string) AcquireItemSetByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcquireItemSetByUserIdRequestFromDict(dict)
}

func NewAcquireItemSetByUserIdRequestFromDict(data map[string]interface{}) AcquireItemSetByUserIdRequest {
    return AcquireItemSetByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        ItemName: core.CastString(data["itemName"]),
        UserId: core.CastString(data["userId"]),
        AcquireCount: core.CastInt64(data["acquireCount"]),
        ExpiresAt: core.CastInt64(data["expiresAt"]),
        CreateNewItemSet: core.CastBool(data["createNewItemSet"]),
        ItemSetName: core.CastString(data["itemSetName"]),
    }
}

func (p AcquireItemSetByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "itemName": p.ItemName,
        "userId": p.UserId,
        "acquireCount": p.AcquireCount,
        "expiresAt": p.ExpiresAt,
        "createNewItemSet": p.CreateNewItemSet,
        "itemSetName": p.ItemSetName,
    }
}

func (p AcquireItemSetByUserIdRequest) Pointer() *AcquireItemSetByUserIdRequest {
    return &p
}

type ConsumeItemSetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
    ItemName *string `json:"itemName"`
    ConsumeCount *int64 `json:"consumeCount"`
    ItemSetName *string `json:"itemSetName"`
}

func NewConsumeItemSetRequestFromJson(data string) ConsumeItemSetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeItemSetRequestFromDict(dict)
}

func NewConsumeItemSetRequestFromDict(data map[string]interface{}) ConsumeItemSetRequest {
    return ConsumeItemSetRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ItemName: core.CastString(data["itemName"]),
        ConsumeCount: core.CastInt64(data["consumeCount"]),
        ItemSetName: core.CastString(data["itemSetName"]),
    }
}

func (p ConsumeItemSetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
        "itemName": p.ItemName,
        "consumeCount": p.ConsumeCount,
        "itemSetName": p.ItemSetName,
    }
}

func (p ConsumeItemSetRequest) Pointer() *ConsumeItemSetRequest {
    return &p
}

type ConsumeItemSetByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    ItemName *string `json:"itemName"`
    ConsumeCount *int64 `json:"consumeCount"`
    ItemSetName *string `json:"itemSetName"`
}

func NewConsumeItemSetByUserIdRequestFromJson(data string) ConsumeItemSetByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeItemSetByUserIdRequestFromDict(dict)
}

func NewConsumeItemSetByUserIdRequestFromDict(data map[string]interface{}) ConsumeItemSetByUserIdRequest {
    return ConsumeItemSetByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        ItemName: core.CastString(data["itemName"]),
        ConsumeCount: core.CastInt64(data["consumeCount"]),
        ItemSetName: core.CastString(data["itemSetName"]),
    }
}

func (p ConsumeItemSetByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "itemName": p.ItemName,
        "consumeCount": p.ConsumeCount,
        "itemSetName": p.ItemSetName,
    }
}

func (p ConsumeItemSetByUserIdRequest) Pointer() *ConsumeItemSetByUserIdRequest {
    return &p
}

type DeleteItemSetByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
}

func NewDeleteItemSetByUserIdRequestFromJson(data string) DeleteItemSetByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteItemSetByUserIdRequestFromDict(dict)
}

func NewDeleteItemSetByUserIdRequestFromDict(data map[string]interface{}) DeleteItemSetByUserIdRequest {
    return DeleteItemSetByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
    }
}

func (p DeleteItemSetByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
    }
}

func (p DeleteItemSetByUserIdRequest) Pointer() *DeleteItemSetByUserIdRequest {
    return &p
}

type AcquireItemSetByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewAcquireItemSetByStampSheetRequestFromJson(data string) AcquireItemSetByStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcquireItemSetByStampSheetRequestFromDict(dict)
}

func NewAcquireItemSetByStampSheetRequestFromDict(data map[string]interface{}) AcquireItemSetByStampSheetRequest {
    return AcquireItemSetByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p AcquireItemSetByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p AcquireItemSetByStampSheetRequest) Pointer() *AcquireItemSetByStampSheetRequest {
    return &p
}

type ConsumeItemSetByStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewConsumeItemSetByStampTaskRequestFromJson(data string) ConsumeItemSetByStampTaskRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeItemSetByStampTaskRequestFromDict(dict)
}

func NewConsumeItemSetByStampTaskRequestFromDict(data map[string]interface{}) ConsumeItemSetByStampTaskRequest {
    return ConsumeItemSetByStampTaskRequest {
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p ConsumeItemSetByStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p ConsumeItemSetByStampTaskRequest) Pointer() *ConsumeItemSetByStampTaskRequest {
    return &p
}

type DescribeReferenceOfRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
}

func NewDescribeReferenceOfRequestFromJson(data string) DescribeReferenceOfRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeReferenceOfRequestFromDict(dict)
}

func NewDescribeReferenceOfRequestFromDict(data map[string]interface{}) DescribeReferenceOfRequest {
    return DescribeReferenceOfRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
    }
}

func (p DescribeReferenceOfRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
    }
}

func (p DescribeReferenceOfRequest) Pointer() *DescribeReferenceOfRequest {
    return &p
}

type DescribeReferenceOfByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
}

func NewDescribeReferenceOfByUserIdRequestFromJson(data string) DescribeReferenceOfByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeReferenceOfByUserIdRequestFromDict(dict)
}

func NewDescribeReferenceOfByUserIdRequestFromDict(data map[string]interface{}) DescribeReferenceOfByUserIdRequest {
    return DescribeReferenceOfByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
    }
}

func (p DescribeReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
    }
}

func (p DescribeReferenceOfByUserIdRequest) Pointer() *DescribeReferenceOfByUserIdRequest {
    return &p
}

type GetReferenceOfRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    ReferenceOf *string `json:"referenceOf"`
}

func NewGetReferenceOfRequestFromJson(data string) GetReferenceOfRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetReferenceOfRequestFromDict(dict)
}

func NewGetReferenceOfRequestFromDict(data map[string]interface{}) GetReferenceOfRequest {
    return GetReferenceOfRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        ReferenceOf: core.CastString(data["referenceOf"]),
    }
}

func (p GetReferenceOfRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "referenceOf": p.ReferenceOf,
    }
}

func (p GetReferenceOfRequest) Pointer() *GetReferenceOfRequest {
    return &p
}

type GetReferenceOfByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    ReferenceOf *string `json:"referenceOf"`
}

func NewGetReferenceOfByUserIdRequestFromJson(data string) GetReferenceOfByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetReferenceOfByUserIdRequestFromDict(dict)
}

func NewGetReferenceOfByUserIdRequestFromDict(data map[string]interface{}) GetReferenceOfByUserIdRequest {
    return GetReferenceOfByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        ReferenceOf: core.CastString(data["referenceOf"]),
    }
}

func (p GetReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "referenceOf": p.ReferenceOf,
    }
}

func (p GetReferenceOfByUserIdRequest) Pointer() *GetReferenceOfByUserIdRequest {
    return &p
}

type VerifyReferenceOfRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    ReferenceOf *string `json:"referenceOf"`
    VerifyType *string `json:"verifyType"`
}

func NewVerifyReferenceOfRequestFromJson(data string) VerifyReferenceOfRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyReferenceOfRequestFromDict(dict)
}

func NewVerifyReferenceOfRequestFromDict(data map[string]interface{}) VerifyReferenceOfRequest {
    return VerifyReferenceOfRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        ReferenceOf: core.CastString(data["referenceOf"]),
        VerifyType: core.CastString(data["verifyType"]),
    }
}

func (p VerifyReferenceOfRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "referenceOf": p.ReferenceOf,
        "verifyType": p.VerifyType,
    }
}

func (p VerifyReferenceOfRequest) Pointer() *VerifyReferenceOfRequest {
    return &p
}

type VerifyReferenceOfByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    ReferenceOf *string `json:"referenceOf"`
    VerifyType *string `json:"verifyType"`
}

func NewVerifyReferenceOfByUserIdRequestFromJson(data string) VerifyReferenceOfByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyReferenceOfByUserIdRequestFromDict(dict)
}

func NewVerifyReferenceOfByUserIdRequestFromDict(data map[string]interface{}) VerifyReferenceOfByUserIdRequest {
    return VerifyReferenceOfByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        ReferenceOf: core.CastString(data["referenceOf"]),
        VerifyType: core.CastString(data["verifyType"]),
    }
}

func (p VerifyReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "referenceOf": p.ReferenceOf,
        "verifyType": p.VerifyType,
    }
}

func (p VerifyReferenceOfByUserIdRequest) Pointer() *VerifyReferenceOfByUserIdRequest {
    return &p
}

type AddReferenceOfRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    ReferenceOf *string `json:"referenceOf"`
}

func NewAddReferenceOfRequestFromJson(data string) AddReferenceOfRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddReferenceOfRequestFromDict(dict)
}

func NewAddReferenceOfRequestFromDict(data map[string]interface{}) AddReferenceOfRequest {
    return AddReferenceOfRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        ReferenceOf: core.CastString(data["referenceOf"]),
    }
}

func (p AddReferenceOfRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "referenceOf": p.ReferenceOf,
    }
}

func (p AddReferenceOfRequest) Pointer() *AddReferenceOfRequest {
    return &p
}

type AddReferenceOfByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    ReferenceOf *string `json:"referenceOf"`
}

func NewAddReferenceOfByUserIdRequestFromJson(data string) AddReferenceOfByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddReferenceOfByUserIdRequestFromDict(dict)
}

func NewAddReferenceOfByUserIdRequestFromDict(data map[string]interface{}) AddReferenceOfByUserIdRequest {
    return AddReferenceOfByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        ReferenceOf: core.CastString(data["referenceOf"]),
    }
}

func (p AddReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "referenceOf": p.ReferenceOf,
    }
}

func (p AddReferenceOfByUserIdRequest) Pointer() *AddReferenceOfByUserIdRequest {
    return &p
}

type DeleteReferenceOfRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    AccessToken *string `json:"accessToken"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    ReferenceOf *string `json:"referenceOf"`
}

func NewDeleteReferenceOfRequestFromJson(data string) DeleteReferenceOfRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteReferenceOfRequestFromDict(dict)
}

func NewDeleteReferenceOfRequestFromDict(data map[string]interface{}) DeleteReferenceOfRequest {
    return DeleteReferenceOfRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        ReferenceOf: core.CastString(data["referenceOf"]),
    }
}

func (p DeleteReferenceOfRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "accessToken": p.AccessToken,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "referenceOf": p.ReferenceOf,
    }
}

func (p DeleteReferenceOfRequest) Pointer() *DeleteReferenceOfRequest {
    return &p
}

type DeleteReferenceOfByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
    UserId *string `json:"userId"`
    ItemName *string `json:"itemName"`
    ItemSetName *string `json:"itemSetName"`
    ReferenceOf *string `json:"referenceOf"`
}

func NewDeleteReferenceOfByUserIdRequestFromJson(data string) DeleteReferenceOfByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteReferenceOfByUserIdRequestFromDict(dict)
}

func NewDeleteReferenceOfByUserIdRequestFromDict(data map[string]interface{}) DeleteReferenceOfByUserIdRequest {
    return DeleteReferenceOfByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
        UserId: core.CastString(data["userId"]),
        ItemName: core.CastString(data["itemName"]),
        ItemSetName: core.CastString(data["itemSetName"]),
        ReferenceOf: core.CastString(data["referenceOf"]),
    }
}

func (p DeleteReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
        "userId": p.UserId,
        "itemName": p.ItemName,
        "itemSetName": p.ItemSetName,
        "referenceOf": p.ReferenceOf,
    }
}

func (p DeleteReferenceOfByUserIdRequest) Pointer() *DeleteReferenceOfByUserIdRequest {
    return &p
}

type AddReferenceOfItemSetByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewAddReferenceOfItemSetByStampSheetRequestFromJson(data string) AddReferenceOfItemSetByStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddReferenceOfItemSetByStampSheetRequestFromDict(dict)
}

func NewAddReferenceOfItemSetByStampSheetRequestFromDict(data map[string]interface{}) AddReferenceOfItemSetByStampSheetRequest {
    return AddReferenceOfItemSetByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p AddReferenceOfItemSetByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p AddReferenceOfItemSetByStampSheetRequest) Pointer() *AddReferenceOfItemSetByStampSheetRequest {
    return &p
}

type DeleteReferenceOfItemSetByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewDeleteReferenceOfItemSetByStampSheetRequestFromJson(data string) DeleteReferenceOfItemSetByStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteReferenceOfItemSetByStampSheetRequestFromDict(dict)
}

func NewDeleteReferenceOfItemSetByStampSheetRequestFromDict(data map[string]interface{}) DeleteReferenceOfItemSetByStampSheetRequest {
    return DeleteReferenceOfItemSetByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p DeleteReferenceOfItemSetByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p DeleteReferenceOfItemSetByStampSheetRequest) Pointer() *DeleteReferenceOfItemSetByStampSheetRequest {
    return &p
}

type VerifyReferenceOfByStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewVerifyReferenceOfByStampTaskRequestFromJson(data string) VerifyReferenceOfByStampTaskRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyReferenceOfByStampTaskRequestFromDict(dict)
}

func NewVerifyReferenceOfByStampTaskRequestFromDict(data map[string]interface{}) VerifyReferenceOfByStampTaskRequest {
    return VerifyReferenceOfByStampTaskRequest {
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p VerifyReferenceOfByStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p VerifyReferenceOfByStampTaskRequest) Pointer() *VerifyReferenceOfByStampTaskRequest {
    return &p
}