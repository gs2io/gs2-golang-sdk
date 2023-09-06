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

type DescribeCountersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    LimitName *string `json:"limitName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeCountersRequestFromJson(data string) DescribeCountersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeCountersRequestFromDict(dict)
}

func NewDescribeCountersRequestFromDict(data map[string]interface{}) DescribeCountersRequest {
    return DescribeCountersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        LimitName: core.CastString(data["limitName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeCountersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "limitName": p.LimitName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeCountersRequest) Pointer() *DescribeCountersRequest {
    return &p
}

type DescribeCountersByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    LimitName *string `json:"limitName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeCountersByUserIdRequestFromJson(data string) DescribeCountersByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeCountersByUserIdRequestFromDict(dict)
}

func NewDescribeCountersByUserIdRequestFromDict(data map[string]interface{}) DescribeCountersByUserIdRequest {
    return DescribeCountersByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        LimitName: core.CastString(data["limitName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeCountersByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "limitName": p.LimitName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeCountersByUserIdRequest) Pointer() *DescribeCountersByUserIdRequest {
    return &p
}

type GetCounterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
    AccessToken *string `json:"accessToken"`
    CounterName *string `json:"counterName"`
}

func NewGetCounterRequestFromJson(data string) GetCounterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCounterRequestFromDict(dict)
}

func NewGetCounterRequestFromDict(data map[string]interface{}) GetCounterRequest {
    return GetCounterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
        AccessToken: core.CastString(data["accessToken"]),
        CounterName: core.CastString(data["counterName"]),
    }
}

func (p GetCounterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
        "accessToken": p.AccessToken,
        "counterName": p.CounterName,
    }
}

func (p GetCounterRequest) Pointer() *GetCounterRequest {
    return &p
}

type GetCounterByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
    UserId *string `json:"userId"`
    CounterName *string `json:"counterName"`
}

func NewGetCounterByUserIdRequestFromJson(data string) GetCounterByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCounterByUserIdRequestFromDict(dict)
}

func NewGetCounterByUserIdRequestFromDict(data map[string]interface{}) GetCounterByUserIdRequest {
    return GetCounterByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
        UserId: core.CastString(data["userId"]),
        CounterName: core.CastString(data["counterName"]),
    }
}

func (p GetCounterByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
        "userId": p.UserId,
        "counterName": p.CounterName,
    }
}

func (p GetCounterByUserIdRequest) Pointer() *GetCounterByUserIdRequest {
    return &p
}

type CountUpRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
    CounterName *string `json:"counterName"`
    AccessToken *string `json:"accessToken"`
    CountUpValue *int32 `json:"countUpValue"`
    MaxValue *int32 `json:"maxValue"`
}

func NewCountUpRequestFromJson(data string) CountUpRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCountUpRequestFromDict(dict)
}

func NewCountUpRequestFromDict(data map[string]interface{}) CountUpRequest {
    return CountUpRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
        CounterName: core.CastString(data["counterName"]),
        AccessToken: core.CastString(data["accessToken"]),
        CountUpValue: core.CastInt32(data["countUpValue"]),
        MaxValue: core.CastInt32(data["maxValue"]),
    }
}

func (p CountUpRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
        "counterName": p.CounterName,
        "accessToken": p.AccessToken,
        "countUpValue": p.CountUpValue,
        "maxValue": p.MaxValue,
    }
}

func (p CountUpRequest) Pointer() *CountUpRequest {
    return &p
}

type CountUpByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
    CounterName *string `json:"counterName"`
    UserId *string `json:"userId"`
    CountUpValue *int32 `json:"countUpValue"`
    MaxValue *int32 `json:"maxValue"`
}

func NewCountUpByUserIdRequestFromJson(data string) CountUpByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCountUpByUserIdRequestFromDict(dict)
}

func NewCountUpByUserIdRequestFromDict(data map[string]interface{}) CountUpByUserIdRequest {
    return CountUpByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
        CounterName: core.CastString(data["counterName"]),
        UserId: core.CastString(data["userId"]),
        CountUpValue: core.CastInt32(data["countUpValue"]),
        MaxValue: core.CastInt32(data["maxValue"]),
    }
}

func (p CountUpByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
        "counterName": p.CounterName,
        "userId": p.UserId,
        "countUpValue": p.CountUpValue,
        "maxValue": p.MaxValue,
    }
}

func (p CountUpByUserIdRequest) Pointer() *CountUpByUserIdRequest {
    return &p
}

type CountDownByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
    CounterName *string `json:"counterName"`
    UserId *string `json:"userId"`
    CountDownValue *int32 `json:"countDownValue"`
}

func NewCountDownByUserIdRequestFromJson(data string) CountDownByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCountDownByUserIdRequestFromDict(dict)
}

func NewCountDownByUserIdRequestFromDict(data map[string]interface{}) CountDownByUserIdRequest {
    return CountDownByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
        CounterName: core.CastString(data["counterName"]),
        UserId: core.CastString(data["userId"]),
        CountDownValue: core.CastInt32(data["countDownValue"]),
    }
}

func (p CountDownByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
        "counterName": p.CounterName,
        "userId": p.UserId,
        "countDownValue": p.CountDownValue,
    }
}

func (p CountDownByUserIdRequest) Pointer() *CountDownByUserIdRequest {
    return &p
}

type DeleteCounterByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
    UserId *string `json:"userId"`
    CounterName *string `json:"counterName"`
}

func NewDeleteCounterByUserIdRequestFromJson(data string) DeleteCounterByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteCounterByUserIdRequestFromDict(dict)
}

func NewDeleteCounterByUserIdRequestFromDict(data map[string]interface{}) DeleteCounterByUserIdRequest {
    return DeleteCounterByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
        UserId: core.CastString(data["userId"]),
        CounterName: core.CastString(data["counterName"]),
    }
}

func (p DeleteCounterByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
        "userId": p.UserId,
        "counterName": p.CounterName,
    }
}

func (p DeleteCounterByUserIdRequest) Pointer() *DeleteCounterByUserIdRequest {
    return &p
}

type CountUpByStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewCountUpByStampTaskRequestFromJson(data string) CountUpByStampTaskRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCountUpByStampTaskRequestFromDict(dict)
}

func NewCountUpByStampTaskRequestFromDict(data map[string]interface{}) CountUpByStampTaskRequest {
    return CountUpByStampTaskRequest {
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p CountUpByStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p CountUpByStampTaskRequest) Pointer() *CountUpByStampTaskRequest {
    return &p
}

type CountDownByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewCountDownByStampSheetRequestFromJson(data string) CountDownByStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCountDownByStampSheetRequestFromDict(dict)
}

func NewCountDownByStampSheetRequestFromDict(data map[string]interface{}) CountDownByStampSheetRequest {
    return CountDownByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p CountDownByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p CountDownByStampSheetRequest) Pointer() *CountDownByStampSheetRequest {
    return &p
}

type DeleteByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewDeleteByStampSheetRequestFromJson(data string) DeleteByStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteByStampSheetRequestFromDict(dict)
}

func NewDeleteByStampSheetRequestFromDict(data map[string]interface{}) DeleteByStampSheetRequest {
    return DeleteByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p DeleteByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p DeleteByStampSheetRequest) Pointer() *DeleteByStampSheetRequest {
    return &p
}

type DescribeLimitModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeLimitModelMastersRequestFromJson(data string) DescribeLimitModelMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeLimitModelMastersRequestFromDict(dict)
}

func NewDescribeLimitModelMastersRequestFromDict(data map[string]interface{}) DescribeLimitModelMastersRequest {
    return DescribeLimitModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeLimitModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeLimitModelMastersRequest) Pointer() *DescribeLimitModelMastersRequest {
    return &p
}

type CreateLimitModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    ResetType *string `json:"resetType"`
    ResetDayOfMonth *int32 `json:"resetDayOfMonth"`
    ResetDayOfWeek *string `json:"resetDayOfWeek"`
    ResetHour *int32 `json:"resetHour"`
}

func NewCreateLimitModelMasterRequestFromJson(data string) CreateLimitModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateLimitModelMasterRequestFromDict(dict)
}

func NewCreateLimitModelMasterRequestFromDict(data map[string]interface{}) CreateLimitModelMasterRequest {
    return CreateLimitModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        ResetType: core.CastString(data["resetType"]),
        ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
        ResetDayOfWeek: core.CastString(data["resetDayOfWeek"]),
        ResetHour: core.CastInt32(data["resetHour"]),
    }
}

func (p CreateLimitModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "resetType": p.ResetType,
        "resetDayOfMonth": p.ResetDayOfMonth,
        "resetDayOfWeek": p.ResetDayOfWeek,
        "resetHour": p.ResetHour,
    }
}

func (p CreateLimitModelMasterRequest) Pointer() *CreateLimitModelMasterRequest {
    return &p
}

type GetLimitModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
}

func NewGetLimitModelMasterRequestFromJson(data string) GetLimitModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetLimitModelMasterRequestFromDict(dict)
}

func NewGetLimitModelMasterRequestFromDict(data map[string]interface{}) GetLimitModelMasterRequest {
    return GetLimitModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
    }
}

func (p GetLimitModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
    }
}

func (p GetLimitModelMasterRequest) Pointer() *GetLimitModelMasterRequest {
    return &p
}

type UpdateLimitModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    ResetType *string `json:"resetType"`
    ResetDayOfMonth *int32 `json:"resetDayOfMonth"`
    ResetDayOfWeek *string `json:"resetDayOfWeek"`
    ResetHour *int32 `json:"resetHour"`
}

func NewUpdateLimitModelMasterRequestFromJson(data string) UpdateLimitModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateLimitModelMasterRequestFromDict(dict)
}

func NewUpdateLimitModelMasterRequestFromDict(data map[string]interface{}) UpdateLimitModelMasterRequest {
    return UpdateLimitModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        ResetType: core.CastString(data["resetType"]),
        ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
        ResetDayOfWeek: core.CastString(data["resetDayOfWeek"]),
        ResetHour: core.CastInt32(data["resetHour"]),
    }
}

func (p UpdateLimitModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
        "description": p.Description,
        "metadata": p.Metadata,
        "resetType": p.ResetType,
        "resetDayOfMonth": p.ResetDayOfMonth,
        "resetDayOfWeek": p.ResetDayOfWeek,
        "resetHour": p.ResetHour,
    }
}

func (p UpdateLimitModelMasterRequest) Pointer() *UpdateLimitModelMasterRequest {
    return &p
}

type DeleteLimitModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
}

func NewDeleteLimitModelMasterRequestFromJson(data string) DeleteLimitModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteLimitModelMasterRequestFromDict(dict)
}

func NewDeleteLimitModelMasterRequestFromDict(data map[string]interface{}) DeleteLimitModelMasterRequest {
    return DeleteLimitModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
    }
}

func (p DeleteLimitModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
    }
}

func (p DeleteLimitModelMasterRequest) Pointer() *DeleteLimitModelMasterRequest {
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

type GetCurrentLimitMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentLimitMasterRequestFromJson(data string) GetCurrentLimitMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentLimitMasterRequestFromDict(dict)
}

func NewGetCurrentLimitMasterRequestFromDict(data map[string]interface{}) GetCurrentLimitMasterRequest {
    return GetCurrentLimitMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentLimitMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentLimitMasterRequest) Pointer() *GetCurrentLimitMasterRequest {
    return &p
}

type UpdateCurrentLimitMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentLimitMasterRequestFromJson(data string) UpdateCurrentLimitMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentLimitMasterRequestFromDict(dict)
}

func NewUpdateCurrentLimitMasterRequestFromDict(data map[string]interface{}) UpdateCurrentLimitMasterRequest {
    return UpdateCurrentLimitMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentLimitMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentLimitMasterRequest) Pointer() *UpdateCurrentLimitMasterRequest {
    return &p
}

type UpdateCurrentLimitMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentLimitMasterFromGitHubRequestFromJson(data string) UpdateCurrentLimitMasterFromGitHubRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentLimitMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentLimitMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentLimitMasterFromGitHubRequest {
    return UpdateCurrentLimitMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentLimitMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentLimitMasterFromGitHubRequest) Pointer() *UpdateCurrentLimitMasterFromGitHubRequest {
    return &p
}

type DescribeLimitModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeLimitModelsRequestFromJson(data string) DescribeLimitModelsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeLimitModelsRequestFromDict(dict)
}

func NewDescribeLimitModelsRequestFromDict(data map[string]interface{}) DescribeLimitModelsRequest {
    return DescribeLimitModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeLimitModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeLimitModelsRequest) Pointer() *DescribeLimitModelsRequest {
    return &p
}

type GetLimitModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
}

func NewGetLimitModelRequestFromJson(data string) GetLimitModelRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetLimitModelRequestFromDict(dict)
}

func NewGetLimitModelRequestFromDict(data map[string]interface{}) GetLimitModelRequest {
    return GetLimitModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
    }
}

func (p GetLimitModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
    }
}

func (p GetLimitModelRequest) Pointer() *GetLimitModelRequest {
    return &p
}