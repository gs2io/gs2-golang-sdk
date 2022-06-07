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

package showcase

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
    TransactionSetting *TransactionSetting `json:"transactionSetting"`
    QueueNamespaceId *string `json:"queueNamespaceId"`
    KeyId *string `json:"keyId"`
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
        TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "transactionSetting": p.TransactionSetting.ToDict(),
        "queueNamespaceId": p.QueueNamespaceId,
        "keyId": p.KeyId,
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
    TransactionSetting *TransactionSetting `json:"transactionSetting"`
    LogSetting *LogSetting `json:"logSetting"`
    QueueNamespaceId *string `json:"queueNamespaceId"`
    KeyId *string `json:"keyId"`
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
        TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "transactionSetting": p.TransactionSetting.ToDict(),
        "logSetting": p.LogSetting.ToDict(),
        "queueNamespaceId": p.QueueNamespaceId,
        "keyId": p.KeyId,
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

type DescribeSalesItemMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeSalesItemMastersRequestFromJson(data string) DescribeSalesItemMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSalesItemMastersRequestFromDict(dict)
}

func NewDescribeSalesItemMastersRequestFromDict(data map[string]interface{}) DescribeSalesItemMastersRequest {
    return DescribeSalesItemMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeSalesItemMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeSalesItemMastersRequest) Pointer() *DescribeSalesItemMastersRequest {
    return &p
}

type CreateSalesItemMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    ConsumeActions []ConsumeAction `json:"consumeActions"`
    AcquireActions []AcquireAction `json:"acquireActions"`
}

func NewCreateSalesItemMasterRequestFromJson(data string) CreateSalesItemMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateSalesItemMasterRequestFromDict(dict)
}

func NewCreateSalesItemMasterRequestFromDict(data map[string]interface{}) CreateSalesItemMasterRequest {
    return CreateSalesItemMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
        AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
    }
}

func (p CreateSalesItemMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "consumeActions": CastConsumeActionsFromDict(
            p.ConsumeActions,
        ),
        "acquireActions": CastAcquireActionsFromDict(
            p.AcquireActions,
        ),
    }
}

func (p CreateSalesItemMasterRequest) Pointer() *CreateSalesItemMasterRequest {
    return &p
}

type GetSalesItemMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    SalesItemName *string `json:"salesItemName"`
}

func NewGetSalesItemMasterRequestFromJson(data string) GetSalesItemMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSalesItemMasterRequestFromDict(dict)
}

func NewGetSalesItemMasterRequestFromDict(data map[string]interface{}) GetSalesItemMasterRequest {
    return GetSalesItemMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        SalesItemName: core.CastString(data["salesItemName"]),
    }
}

func (p GetSalesItemMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "salesItemName": p.SalesItemName,
    }
}

func (p GetSalesItemMasterRequest) Pointer() *GetSalesItemMasterRequest {
    return &p
}

type UpdateSalesItemMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    SalesItemName *string `json:"salesItemName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    ConsumeActions []ConsumeAction `json:"consumeActions"`
    AcquireActions []AcquireAction `json:"acquireActions"`
}

func NewUpdateSalesItemMasterRequestFromJson(data string) UpdateSalesItemMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateSalesItemMasterRequestFromDict(dict)
}

func NewUpdateSalesItemMasterRequestFromDict(data map[string]interface{}) UpdateSalesItemMasterRequest {
    return UpdateSalesItemMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        SalesItemName: core.CastString(data["salesItemName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
        AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
    }
}

func (p UpdateSalesItemMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "salesItemName": p.SalesItemName,
        "description": p.Description,
        "metadata": p.Metadata,
        "consumeActions": CastConsumeActionsFromDict(
            p.ConsumeActions,
        ),
        "acquireActions": CastAcquireActionsFromDict(
            p.AcquireActions,
        ),
    }
}

func (p UpdateSalesItemMasterRequest) Pointer() *UpdateSalesItemMasterRequest {
    return &p
}

type DeleteSalesItemMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    SalesItemName *string `json:"salesItemName"`
}

func NewDeleteSalesItemMasterRequestFromJson(data string) DeleteSalesItemMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteSalesItemMasterRequestFromDict(dict)
}

func NewDeleteSalesItemMasterRequestFromDict(data map[string]interface{}) DeleteSalesItemMasterRequest {
    return DeleteSalesItemMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        SalesItemName: core.CastString(data["salesItemName"]),
    }
}

func (p DeleteSalesItemMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "salesItemName": p.SalesItemName,
    }
}

func (p DeleteSalesItemMasterRequest) Pointer() *DeleteSalesItemMasterRequest {
    return &p
}

type DescribeSalesItemGroupMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeSalesItemGroupMastersRequestFromJson(data string) DescribeSalesItemGroupMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSalesItemGroupMastersRequestFromDict(dict)
}

func NewDescribeSalesItemGroupMastersRequestFromDict(data map[string]interface{}) DescribeSalesItemGroupMastersRequest {
    return DescribeSalesItemGroupMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeSalesItemGroupMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeSalesItemGroupMastersRequest) Pointer() *DescribeSalesItemGroupMastersRequest {
    return &p
}

type CreateSalesItemGroupMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    SalesItemNames []string `json:"salesItemNames"`
}

func NewCreateSalesItemGroupMasterRequestFromJson(data string) CreateSalesItemGroupMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateSalesItemGroupMasterRequestFromDict(dict)
}

func NewCreateSalesItemGroupMasterRequestFromDict(data map[string]interface{}) CreateSalesItemGroupMasterRequest {
    return CreateSalesItemGroupMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        SalesItemNames: core.CastStrings(core.CastArray(data["salesItemNames"])),
    }
}

func (p CreateSalesItemGroupMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "salesItemNames": core.CastStringsFromDict(
            p.SalesItemNames,
        ),
    }
}

func (p CreateSalesItemGroupMasterRequest) Pointer() *CreateSalesItemGroupMasterRequest {
    return &p
}

type GetSalesItemGroupMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    SalesItemGroupName *string `json:"salesItemGroupName"`
}

func NewGetSalesItemGroupMasterRequestFromJson(data string) GetSalesItemGroupMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSalesItemGroupMasterRequestFromDict(dict)
}

func NewGetSalesItemGroupMasterRequestFromDict(data map[string]interface{}) GetSalesItemGroupMasterRequest {
    return GetSalesItemGroupMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        SalesItemGroupName: core.CastString(data["salesItemGroupName"]),
    }
}

func (p GetSalesItemGroupMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "salesItemGroupName": p.SalesItemGroupName,
    }
}

func (p GetSalesItemGroupMasterRequest) Pointer() *GetSalesItemGroupMasterRequest {
    return &p
}

type UpdateSalesItemGroupMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    SalesItemGroupName *string `json:"salesItemGroupName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    SalesItemNames []string `json:"salesItemNames"`
}

func NewUpdateSalesItemGroupMasterRequestFromJson(data string) UpdateSalesItemGroupMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateSalesItemGroupMasterRequestFromDict(dict)
}

func NewUpdateSalesItemGroupMasterRequestFromDict(data map[string]interface{}) UpdateSalesItemGroupMasterRequest {
    return UpdateSalesItemGroupMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        SalesItemGroupName: core.CastString(data["salesItemGroupName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        SalesItemNames: core.CastStrings(core.CastArray(data["salesItemNames"])),
    }
}

func (p UpdateSalesItemGroupMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "salesItemGroupName": p.SalesItemGroupName,
        "description": p.Description,
        "metadata": p.Metadata,
        "salesItemNames": core.CastStringsFromDict(
            p.SalesItemNames,
        ),
    }
}

func (p UpdateSalesItemGroupMasterRequest) Pointer() *UpdateSalesItemGroupMasterRequest {
    return &p
}

type DeleteSalesItemGroupMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    SalesItemGroupName *string `json:"salesItemGroupName"`
}

func NewDeleteSalesItemGroupMasterRequestFromJson(data string) DeleteSalesItemGroupMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteSalesItemGroupMasterRequestFromDict(dict)
}

func NewDeleteSalesItemGroupMasterRequestFromDict(data map[string]interface{}) DeleteSalesItemGroupMasterRequest {
    return DeleteSalesItemGroupMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        SalesItemGroupName: core.CastString(data["salesItemGroupName"]),
    }
}

func (p DeleteSalesItemGroupMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "salesItemGroupName": p.SalesItemGroupName,
    }
}

func (p DeleteSalesItemGroupMasterRequest) Pointer() *DeleteSalesItemGroupMasterRequest {
    return &p
}

type DescribeShowcaseMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeShowcaseMastersRequestFromJson(data string) DescribeShowcaseMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeShowcaseMastersRequestFromDict(dict)
}

func NewDescribeShowcaseMastersRequestFromDict(data map[string]interface{}) DescribeShowcaseMastersRequest {
    return DescribeShowcaseMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeShowcaseMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeShowcaseMastersRequest) Pointer() *DescribeShowcaseMastersRequest {
    return &p
}

type CreateShowcaseMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    DisplayItems []DisplayItemMaster `json:"displayItems"`
    SalesPeriodEventId *string `json:"salesPeriodEventId"`
}

func NewCreateShowcaseMasterRequestFromJson(data string) CreateShowcaseMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateShowcaseMasterRequestFromDict(dict)
}

func NewCreateShowcaseMasterRequestFromDict(data map[string]interface{}) CreateShowcaseMasterRequest {
    return CreateShowcaseMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        DisplayItems: CastDisplayItemMasters(core.CastArray(data["displayItems"])),
        SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
    }
}

func (p CreateShowcaseMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "displayItems": CastDisplayItemMastersFromDict(
            p.DisplayItems,
        ),
        "salesPeriodEventId": p.SalesPeriodEventId,
    }
}

func (p CreateShowcaseMasterRequest) Pointer() *CreateShowcaseMasterRequest {
    return &p
}

type GetShowcaseMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
}

func NewGetShowcaseMasterRequestFromJson(data string) GetShowcaseMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetShowcaseMasterRequestFromDict(dict)
}

func NewGetShowcaseMasterRequestFromDict(data map[string]interface{}) GetShowcaseMasterRequest {
    return GetShowcaseMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
    }
}

func (p GetShowcaseMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
    }
}

func (p GetShowcaseMasterRequest) Pointer() *GetShowcaseMasterRequest {
    return &p
}

type UpdateShowcaseMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    DisplayItems []DisplayItemMaster `json:"displayItems"`
    SalesPeriodEventId *string `json:"salesPeriodEventId"`
}

func NewUpdateShowcaseMasterRequestFromJson(data string) UpdateShowcaseMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateShowcaseMasterRequestFromDict(dict)
}

func NewUpdateShowcaseMasterRequestFromDict(data map[string]interface{}) UpdateShowcaseMasterRequest {
    return UpdateShowcaseMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        DisplayItems: CastDisplayItemMasters(core.CastArray(data["displayItems"])),
        SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
    }
}

func (p UpdateShowcaseMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
        "description": p.Description,
        "metadata": p.Metadata,
        "displayItems": CastDisplayItemMastersFromDict(
            p.DisplayItems,
        ),
        "salesPeriodEventId": p.SalesPeriodEventId,
    }
}

func (p UpdateShowcaseMasterRequest) Pointer() *UpdateShowcaseMasterRequest {
    return &p
}

type DeleteShowcaseMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
}

func NewDeleteShowcaseMasterRequestFromJson(data string) DeleteShowcaseMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteShowcaseMasterRequestFromDict(dict)
}

func NewDeleteShowcaseMasterRequestFromDict(data map[string]interface{}) DeleteShowcaseMasterRequest {
    return DeleteShowcaseMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
    }
}

func (p DeleteShowcaseMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
    }
}

func (p DeleteShowcaseMasterRequest) Pointer() *DeleteShowcaseMasterRequest {
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

type GetCurrentShowcaseMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentShowcaseMasterRequestFromJson(data string) GetCurrentShowcaseMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentShowcaseMasterRequestFromDict(dict)
}

func NewGetCurrentShowcaseMasterRequestFromDict(data map[string]interface{}) GetCurrentShowcaseMasterRequest {
    return GetCurrentShowcaseMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentShowcaseMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentShowcaseMasterRequest) Pointer() *GetCurrentShowcaseMasterRequest {
    return &p
}

type UpdateCurrentShowcaseMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentShowcaseMasterRequestFromJson(data string) UpdateCurrentShowcaseMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentShowcaseMasterRequestFromDict(dict)
}

func NewUpdateCurrentShowcaseMasterRequestFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterRequest {
    return UpdateCurrentShowcaseMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentShowcaseMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentShowcaseMasterRequest) Pointer() *UpdateCurrentShowcaseMasterRequest {
    return &p
}

type UpdateCurrentShowcaseMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentShowcaseMasterFromGitHubRequestFromJson(data string) UpdateCurrentShowcaseMasterFromGitHubRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentShowcaseMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentShowcaseMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterFromGitHubRequest {
    return UpdateCurrentShowcaseMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentShowcaseMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentShowcaseMasterFromGitHubRequest) Pointer() *UpdateCurrentShowcaseMasterFromGitHubRequest {
    return &p
}

type DescribeShowcasesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
}

func NewDescribeShowcasesRequestFromJson(data string) DescribeShowcasesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeShowcasesRequestFromDict(dict)
}

func NewDescribeShowcasesRequestFromDict(data map[string]interface{}) DescribeShowcasesRequest {
    return DescribeShowcasesRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p DescribeShowcasesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
    }
}

func (p DescribeShowcasesRequest) Pointer() *DescribeShowcasesRequest {
    return &p
}

type DescribeShowcasesByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
}

func NewDescribeShowcasesByUserIdRequestFromJson(data string) DescribeShowcasesByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeShowcasesByUserIdRequestFromDict(dict)
}

func NewDescribeShowcasesByUserIdRequestFromDict(data map[string]interface{}) DescribeShowcasesByUserIdRequest {
    return DescribeShowcasesByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p DescribeShowcasesByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
    }
}

func (p DescribeShowcasesByUserIdRequest) Pointer() *DescribeShowcasesByUserIdRequest {
    return &p
}

type GetShowcaseRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
    AccessToken *string `json:"accessToken"`
}

func NewGetShowcaseRequestFromJson(data string) GetShowcaseRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetShowcaseRequestFromDict(dict)
}

func NewGetShowcaseRequestFromDict(data map[string]interface{}) GetShowcaseRequest {
    return GetShowcaseRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p GetShowcaseRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
        "accessToken": p.AccessToken,
    }
}

func (p GetShowcaseRequest) Pointer() *GetShowcaseRequest {
    return &p
}

type GetShowcaseByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
    UserId *string `json:"userId"`
}

func NewGetShowcaseByUserIdRequestFromJson(data string) GetShowcaseByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetShowcaseByUserIdRequestFromDict(dict)
}

func NewGetShowcaseByUserIdRequestFromDict(data map[string]interface{}) GetShowcaseByUserIdRequest {
    return GetShowcaseByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p GetShowcaseByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
        "userId": p.UserId,
    }
}

func (p GetShowcaseByUserIdRequest) Pointer() *GetShowcaseByUserIdRequest {
    return &p
}

type BuyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
    DisplayItemId *string `json:"displayItemId"`
    AccessToken *string `json:"accessToken"`
    Config []Config `json:"config"`
}

func NewBuyRequestFromJson(data string) BuyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBuyRequestFromDict(dict)
}

func NewBuyRequestFromDict(data map[string]interface{}) BuyRequest {
    return BuyRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
        DisplayItemId: core.CastString(data["displayItemId"]),
        AccessToken: core.CastString(data["accessToken"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p BuyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
        "displayItemId": p.DisplayItemId,
        "accessToken": p.AccessToken,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p BuyRequest) Pointer() *BuyRequest {
    return &p
}

type BuyByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
    DisplayItemId *string `json:"displayItemId"`
    UserId *string `json:"userId"`
    Config []Config `json:"config"`
}

func NewBuyByUserIdRequestFromJson(data string) BuyByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBuyByUserIdRequestFromDict(dict)
}

func NewBuyByUserIdRequestFromDict(data map[string]interface{}) BuyByUserIdRequest {
    return BuyByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
        DisplayItemId: core.CastString(data["displayItemId"]),
        UserId: core.CastString(data["userId"]),
        Config: CastConfigs(core.CastArray(data["config"])),
    }
}

func (p BuyByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
        "displayItemId": p.DisplayItemId,
        "userId": p.UserId,
        "config": CastConfigsFromDict(
            p.Config,
        ),
    }
}

func (p BuyByUserIdRequest) Pointer() *BuyByUserIdRequest {
    return &p
}