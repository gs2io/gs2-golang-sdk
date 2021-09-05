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

package ranking

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
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
    DuplicationAvoider *string `json:"duplicationAvoider"`
    Name *string `json:"name"`
    Description *string `json:"description"`
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
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
    return &p
}

type GetNamespaceStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
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
    DuplicationAvoider *string `json:"duplicationAvoider"`
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
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    Description *string `json:"description"`
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
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
    return &p
}

type DeleteNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
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

type DescribeCategoryModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeCategoryModelsRequestFromJson(data string) DescribeCategoryModelsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeCategoryModelsRequestFromDict(dict)
}

func NewDescribeCategoryModelsRequestFromDict(data map[string]interface{}) DescribeCategoryModelsRequest {
    return DescribeCategoryModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeCategoryModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeCategoryModelsRequest) Pointer() *DescribeCategoryModelsRequest {
    return &p
}

type GetCategoryModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
}

func NewGetCategoryModelRequestFromJson(data string) GetCategoryModelRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCategoryModelRequestFromDict(dict)
}

func NewGetCategoryModelRequestFromDict(data map[string]interface{}) GetCategoryModelRequest {
    return GetCategoryModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
    }
}

func (p GetCategoryModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
    }
}

func (p GetCategoryModelRequest) Pointer() *GetCategoryModelRequest {
    return &p
}

type DescribeCategoryModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeCategoryModelMastersRequestFromJson(data string) DescribeCategoryModelMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeCategoryModelMastersRequestFromDict(dict)
}

func NewDescribeCategoryModelMastersRequestFromDict(data map[string]interface{}) DescribeCategoryModelMastersRequest {
    return DescribeCategoryModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeCategoryModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeCategoryModelMastersRequest) Pointer() *DescribeCategoryModelMastersRequest {
    return &p
}

type CreateCategoryModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    MinimumValue *int64 `json:"minimumValue"`
    MaximumValue *int64 `json:"maximumValue"`
    OrderDirection *string `json:"orderDirection"`
    Scope *string `json:"scope"`
    UniqueByUserId *bool `json:"uniqueByUserId"`
    CalculateFixedTimingHour *int32 `json:"calculateFixedTimingHour"`
    CalculateFixedTimingMinute *int32 `json:"calculateFixedTimingMinute"`
    CalculateIntervalMinutes *int32 `json:"calculateIntervalMinutes"`
    EntryPeriodEventId *string `json:"entryPeriodEventId"`
    AccessPeriodEventId *string `json:"accessPeriodEventId"`
    Generation *string `json:"generation"`
}

func NewCreateCategoryModelMasterRequestFromJson(data string) CreateCategoryModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateCategoryModelMasterRequestFromDict(dict)
}

func NewCreateCategoryModelMasterRequestFromDict(data map[string]interface{}) CreateCategoryModelMasterRequest {
    return CreateCategoryModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        MinimumValue: core.CastInt64(data["minimumValue"]),
        MaximumValue: core.CastInt64(data["maximumValue"]),
        OrderDirection: core.CastString(data["orderDirection"]),
        Scope: core.CastString(data["scope"]),
        UniqueByUserId: core.CastBool(data["uniqueByUserId"]),
        CalculateFixedTimingHour: core.CastInt32(data["calculateFixedTimingHour"]),
        CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
        CalculateIntervalMinutes: core.CastInt32(data["calculateIntervalMinutes"]),
        EntryPeriodEventId: core.CastString(data["entryPeriodEventId"]),
        AccessPeriodEventId: core.CastString(data["accessPeriodEventId"]),
        Generation: core.CastString(data["generation"]),
    }
}

func (p CreateCategoryModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "minimumValue": p.MinimumValue,
        "maximumValue": p.MaximumValue,
        "orderDirection": p.OrderDirection,
        "scope": p.Scope,
        "uniqueByUserId": p.UniqueByUserId,
        "calculateFixedTimingHour": p.CalculateFixedTimingHour,
        "calculateFixedTimingMinute": p.CalculateFixedTimingMinute,
        "calculateIntervalMinutes": p.CalculateIntervalMinutes,
        "entryPeriodEventId": p.EntryPeriodEventId,
        "accessPeriodEventId": p.AccessPeriodEventId,
        "generation": p.Generation,
    }
}

func (p CreateCategoryModelMasterRequest) Pointer() *CreateCategoryModelMasterRequest {
    return &p
}

type GetCategoryModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
}

func NewGetCategoryModelMasterRequestFromJson(data string) GetCategoryModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCategoryModelMasterRequestFromDict(dict)
}

func NewGetCategoryModelMasterRequestFromDict(data map[string]interface{}) GetCategoryModelMasterRequest {
    return GetCategoryModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
    }
}

func (p GetCategoryModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
    }
}

func (p GetCategoryModelMasterRequest) Pointer() *GetCategoryModelMasterRequest {
    return &p
}

type UpdateCategoryModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    MinimumValue *int64 `json:"minimumValue"`
    MaximumValue *int64 `json:"maximumValue"`
    OrderDirection *string `json:"orderDirection"`
    Scope *string `json:"scope"`
    UniqueByUserId *bool `json:"uniqueByUserId"`
    CalculateFixedTimingHour *int32 `json:"calculateFixedTimingHour"`
    CalculateFixedTimingMinute *int32 `json:"calculateFixedTimingMinute"`
    CalculateIntervalMinutes *int32 `json:"calculateIntervalMinutes"`
    EntryPeriodEventId *string `json:"entryPeriodEventId"`
    AccessPeriodEventId *string `json:"accessPeriodEventId"`
    Generation *string `json:"generation"`
}

func NewUpdateCategoryModelMasterRequestFromJson(data string) UpdateCategoryModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCategoryModelMasterRequestFromDict(dict)
}

func NewUpdateCategoryModelMasterRequestFromDict(data map[string]interface{}) UpdateCategoryModelMasterRequest {
    return UpdateCategoryModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        MinimumValue: core.CastInt64(data["minimumValue"]),
        MaximumValue: core.CastInt64(data["maximumValue"]),
        OrderDirection: core.CastString(data["orderDirection"]),
        Scope: core.CastString(data["scope"]),
        UniqueByUserId: core.CastBool(data["uniqueByUserId"]),
        CalculateFixedTimingHour: core.CastInt32(data["calculateFixedTimingHour"]),
        CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
        CalculateIntervalMinutes: core.CastInt32(data["calculateIntervalMinutes"]),
        EntryPeriodEventId: core.CastString(data["entryPeriodEventId"]),
        AccessPeriodEventId: core.CastString(data["accessPeriodEventId"]),
        Generation: core.CastString(data["generation"]),
    }
}

func (p UpdateCategoryModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "description": p.Description,
        "metadata": p.Metadata,
        "minimumValue": p.MinimumValue,
        "maximumValue": p.MaximumValue,
        "orderDirection": p.OrderDirection,
        "scope": p.Scope,
        "uniqueByUserId": p.UniqueByUserId,
        "calculateFixedTimingHour": p.CalculateFixedTimingHour,
        "calculateFixedTimingMinute": p.CalculateFixedTimingMinute,
        "calculateIntervalMinutes": p.CalculateIntervalMinutes,
        "entryPeriodEventId": p.EntryPeriodEventId,
        "accessPeriodEventId": p.AccessPeriodEventId,
        "generation": p.Generation,
    }
}

func (p UpdateCategoryModelMasterRequest) Pointer() *UpdateCategoryModelMasterRequest {
    return &p
}

type DeleteCategoryModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
}

func NewDeleteCategoryModelMasterRequestFromJson(data string) DeleteCategoryModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteCategoryModelMasterRequestFromDict(dict)
}

func NewDeleteCategoryModelMasterRequestFromDict(data map[string]interface{}) DeleteCategoryModelMasterRequest {
    return DeleteCategoryModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
    }
}

func (p DeleteCategoryModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
    }
}

func (p DeleteCategoryModelMasterRequest) Pointer() *DeleteCategoryModelMasterRequest {
    return &p
}

type DescribeSubscribesByCategoryNameRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    AccessToken *string `json:"accessToken"`
}

func NewDescribeSubscribesByCategoryNameRequestFromJson(data string) DescribeSubscribesByCategoryNameRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSubscribesByCategoryNameRequestFromDict(dict)
}

func NewDescribeSubscribesByCategoryNameRequestFromDict(data map[string]interface{}) DescribeSubscribesByCategoryNameRequest {
    return DescribeSubscribesByCategoryNameRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p DescribeSubscribesByCategoryNameRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "accessToken": p.AccessToken,
    }
}

func (p DescribeSubscribesByCategoryNameRequest) Pointer() *DescribeSubscribesByCategoryNameRequest {
    return &p
}

type DescribeSubscribesByCategoryNameAndUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    UserId *string `json:"userId"`
}

func NewDescribeSubscribesByCategoryNameAndUserIdRequestFromJson(data string) DescribeSubscribesByCategoryNameAndUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSubscribesByCategoryNameAndUserIdRequestFromDict(dict)
}

func NewDescribeSubscribesByCategoryNameAndUserIdRequestFromDict(data map[string]interface{}) DescribeSubscribesByCategoryNameAndUserIdRequest {
    return DescribeSubscribesByCategoryNameAndUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p DescribeSubscribesByCategoryNameAndUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "userId": p.UserId,
    }
}

func (p DescribeSubscribesByCategoryNameAndUserIdRequest) Pointer() *DescribeSubscribesByCategoryNameAndUserIdRequest {
    return &p
}

type SubscribeRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewSubscribeRequestFromJson(data string) SubscribeRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSubscribeRequestFromDict(dict)
}

func NewSubscribeRequestFromDict(data map[string]interface{}) SubscribeRequest {
    return SubscribeRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p SubscribeRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p SubscribeRequest) Pointer() *SubscribeRequest {
    return &p
}

type SubscribeByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewSubscribeByUserIdRequestFromJson(data string) SubscribeByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSubscribeByUserIdRequestFromDict(dict)
}

func NewSubscribeByUserIdRequestFromDict(data map[string]interface{}) SubscribeByUserIdRequest {
    return SubscribeByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p SubscribeByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p SubscribeByUserIdRequest) Pointer() *SubscribeByUserIdRequest {
    return &p
}

type GetSubscribeRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewGetSubscribeRequestFromJson(data string) GetSubscribeRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSubscribeRequestFromDict(dict)
}

func NewGetSubscribeRequestFromDict(data map[string]interface{}) GetSubscribeRequest {
    return GetSubscribeRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p GetSubscribeRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p GetSubscribeRequest) Pointer() *GetSubscribeRequest {
    return &p
}

type GetSubscribeByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewGetSubscribeByUserIdRequestFromJson(data string) GetSubscribeByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSubscribeByUserIdRequestFromDict(dict)
}

func NewGetSubscribeByUserIdRequestFromDict(data map[string]interface{}) GetSubscribeByUserIdRequest {
    return GetSubscribeByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p GetSubscribeByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p GetSubscribeByUserIdRequest) Pointer() *GetSubscribeByUserIdRequest {
    return &p
}

type UnsubscribeRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewUnsubscribeRequestFromJson(data string) UnsubscribeRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUnsubscribeRequestFromDict(dict)
}

func NewUnsubscribeRequestFromDict(data map[string]interface{}) UnsubscribeRequest {
    return UnsubscribeRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p UnsubscribeRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p UnsubscribeRequest) Pointer() *UnsubscribeRequest {
    return &p
}

type UnsubscribeByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewUnsubscribeByUserIdRequestFromJson(data string) UnsubscribeByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUnsubscribeByUserIdRequestFromDict(dict)
}

func NewUnsubscribeByUserIdRequestFromDict(data map[string]interface{}) UnsubscribeByUserIdRequest {
    return UnsubscribeByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p UnsubscribeByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p UnsubscribeByUserIdRequest) Pointer() *UnsubscribeByUserIdRequest {
    return &p
}

type DescribeScoresRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    AccessToken *string `json:"accessToken"`
    ScorerUserId *string `json:"scorerUserId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeScoresRequestFromJson(data string) DescribeScoresRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeScoresRequestFromDict(dict)
}

func NewDescribeScoresRequestFromDict(data map[string]interface{}) DescribeScoresRequest {
    return DescribeScoresRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ScorerUserId: core.CastString(data["scorerUserId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeScoresRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "accessToken": p.AccessToken,
        "scorerUserId": p.ScorerUserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeScoresRequest) Pointer() *DescribeScoresRequest {
    return &p
}

type DescribeScoresByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    UserId *string `json:"userId"`
    ScorerUserId *string `json:"scorerUserId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeScoresByUserIdRequestFromJson(data string) DescribeScoresByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeScoresByUserIdRequestFromDict(dict)
}

func NewDescribeScoresByUserIdRequestFromDict(data map[string]interface{}) DescribeScoresByUserIdRequest {
    return DescribeScoresByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        ScorerUserId: core.CastString(data["scorerUserId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeScoresByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "userId": p.UserId,
        "scorerUserId": p.ScorerUserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeScoresByUserIdRequest) Pointer() *DescribeScoresByUserIdRequest {
    return &p
}

type GetScoreRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    AccessToken *string `json:"accessToken"`
    ScorerUserId *string `json:"scorerUserId"`
    UniqueId *string `json:"uniqueId"`
}

func NewGetScoreRequestFromJson(data string) GetScoreRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetScoreRequestFromDict(dict)
}

func NewGetScoreRequestFromDict(data map[string]interface{}) GetScoreRequest {
    return GetScoreRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ScorerUserId: core.CastString(data["scorerUserId"]),
        UniqueId: core.CastString(data["uniqueId"]),
    }
}

func (p GetScoreRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "accessToken": p.AccessToken,
        "scorerUserId": p.ScorerUserId,
        "uniqueId": p.UniqueId,
    }
}

func (p GetScoreRequest) Pointer() *GetScoreRequest {
    return &p
}

type GetScoreByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    UserId *string `json:"userId"`
    ScorerUserId *string `json:"scorerUserId"`
    UniqueId *string `json:"uniqueId"`
}

func NewGetScoreByUserIdRequestFromJson(data string) GetScoreByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetScoreByUserIdRequestFromDict(dict)
}

func NewGetScoreByUserIdRequestFromDict(data map[string]interface{}) GetScoreByUserIdRequest {
    return GetScoreByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        ScorerUserId: core.CastString(data["scorerUserId"]),
        UniqueId: core.CastString(data["uniqueId"]),
    }
}

func (p GetScoreByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "userId": p.UserId,
        "scorerUserId": p.ScorerUserId,
        "uniqueId": p.UniqueId,
    }
}

func (p GetScoreByUserIdRequest) Pointer() *GetScoreByUserIdRequest {
    return &p
}

type DescribeRankingsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    AccessToken *string `json:"accessToken"`
    StartIndex *int64 `json:"startIndex"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeRankingsRequestFromJson(data string) DescribeRankingsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRankingsRequestFromDict(dict)
}

func NewDescribeRankingsRequestFromDict(data map[string]interface{}) DescribeRankingsRequest {
    return DescribeRankingsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        StartIndex: core.CastInt64(data["startIndex"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeRankingsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "accessToken": p.AccessToken,
        "startIndex": p.StartIndex,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeRankingsRequest) Pointer() *DescribeRankingsRequest {
    return &p
}

type DescribeRankingssByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    UserId *string `json:"userId"`
    StartIndex *int64 `json:"startIndex"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeRankingssByUserIdRequestFromJson(data string) DescribeRankingssByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRankingssByUserIdRequestFromDict(dict)
}

func NewDescribeRankingssByUserIdRequestFromDict(data map[string]interface{}) DescribeRankingssByUserIdRequest {
    return DescribeRankingssByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        StartIndex: core.CastInt64(data["startIndex"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeRankingssByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "userId": p.UserId,
        "startIndex": p.StartIndex,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeRankingssByUserIdRequest) Pointer() *DescribeRankingssByUserIdRequest {
    return &p
}

type DescribeNearRankingsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    Score *int64 `json:"score"`
}

func NewDescribeNearRankingsRequestFromJson(data string) DescribeNearRankingsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeNearRankingsRequestFromDict(dict)
}

func NewDescribeNearRankingsRequestFromDict(data map[string]interface{}) DescribeNearRankingsRequest {
    return DescribeNearRankingsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        Score: core.CastInt64(data["score"]),
    }
}

func (p DescribeNearRankingsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "score": p.Score,
    }
}

func (p DescribeNearRankingsRequest) Pointer() *DescribeNearRankingsRequest {
    return &p
}

type GetRankingRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    AccessToken *string `json:"accessToken"`
    ScorerUserId *string `json:"scorerUserId"`
    UniqueId *string `json:"uniqueId"`
}

func NewGetRankingRequestFromJson(data string) GetRankingRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRankingRequestFromDict(dict)
}

func NewGetRankingRequestFromDict(data map[string]interface{}) GetRankingRequest {
    return GetRankingRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ScorerUserId: core.CastString(data["scorerUserId"]),
        UniqueId: core.CastString(data["uniqueId"]),
    }
}

func (p GetRankingRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "accessToken": p.AccessToken,
        "scorerUserId": p.ScorerUserId,
        "uniqueId": p.UniqueId,
    }
}

func (p GetRankingRequest) Pointer() *GetRankingRequest {
    return &p
}

type GetRankingByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    UserId *string `json:"userId"`
    ScorerUserId *string `json:"scorerUserId"`
    UniqueId *string `json:"uniqueId"`
}

func NewGetRankingByUserIdRequestFromJson(data string) GetRankingByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRankingByUserIdRequestFromDict(dict)
}

func NewGetRankingByUserIdRequestFromDict(data map[string]interface{}) GetRankingByUserIdRequest {
    return GetRankingByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        ScorerUserId: core.CastString(data["scorerUserId"]),
        UniqueId: core.CastString(data["uniqueId"]),
    }
}

func (p GetRankingByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "userId": p.UserId,
        "scorerUserId": p.ScorerUserId,
        "uniqueId": p.UniqueId,
    }
}

func (p GetRankingByUserIdRequest) Pointer() *GetRankingByUserIdRequest {
    return &p
}

type PutScoreRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    AccessToken *string `json:"accessToken"`
    Score *int64 `json:"score"`
    Metadata *string `json:"metadata"`
}

func NewPutScoreRequestFromJson(data string) PutScoreRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPutScoreRequestFromDict(dict)
}

func NewPutScoreRequestFromDict(data map[string]interface{}) PutScoreRequest {
    return PutScoreRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Score: core.CastInt64(data["score"]),
        Metadata: core.CastString(data["metadata"]),
    }
}

func (p PutScoreRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "accessToken": p.AccessToken,
        "score": p.Score,
        "metadata": p.Metadata,
    }
}

func (p PutScoreRequest) Pointer() *PutScoreRequest {
    return &p
}

type PutScoreByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
    UserId *string `json:"userId"`
    Score *int64 `json:"score"`
    Metadata *string `json:"metadata"`
}

func NewPutScoreByUserIdRequestFromJson(data string) PutScoreByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPutScoreByUserIdRequestFromDict(dict)
}

func NewPutScoreByUserIdRequestFromDict(data map[string]interface{}) PutScoreByUserIdRequest {
    return PutScoreByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        Score: core.CastInt64(data["score"]),
        Metadata: core.CastString(data["metadata"]),
    }
}

func (p PutScoreByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
        "userId": p.UserId,
        "score": p.Score,
        "metadata": p.Metadata,
    }
}

func (p PutScoreByUserIdRequest) Pointer() *PutScoreByUserIdRequest {
    return &p
}

type ExportMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
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

type GetCurrentRankingMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentRankingMasterRequestFromJson(data string) GetCurrentRankingMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentRankingMasterRequestFromDict(dict)
}

func NewGetCurrentRankingMasterRequestFromDict(data map[string]interface{}) GetCurrentRankingMasterRequest {
    return GetCurrentRankingMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentRankingMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentRankingMasterRequest) Pointer() *GetCurrentRankingMasterRequest {
    return &p
}

type UpdateCurrentRankingMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentRankingMasterRequestFromJson(data string) UpdateCurrentRankingMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentRankingMasterRequestFromDict(dict)
}

func NewUpdateCurrentRankingMasterRequestFromDict(data map[string]interface{}) UpdateCurrentRankingMasterRequest {
    return UpdateCurrentRankingMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentRankingMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentRankingMasterRequest) Pointer() *UpdateCurrentRankingMasterRequest {
    return &p
}

type UpdateCurrentRankingMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentRankingMasterFromGitHubRequestFromJson(data string) UpdateCurrentRankingMasterFromGitHubRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentRankingMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentRankingMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentRankingMasterFromGitHubRequest {
    return UpdateCurrentRankingMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentRankingMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentRankingMasterFromGitHubRequest) Pointer() *UpdateCurrentRankingMasterFromGitHubRequest {
    return &p
}