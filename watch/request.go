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

package watch

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type GetChartRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Measure *string `json:"measure"`
    Grn *string `json:"grn"`
    Round *string `json:"round"`
    Filters []Filter `json:"filters"`
    GroupBys []*string `json:"groupBys"`
    CountBy *string `json:"countBy"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewGetChartRequestFromJson(data string) GetChartRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetChartRequestFromDict(dict)
}

func NewGetChartRequestFromDict(data map[string]interface{}) GetChartRequest {
    return GetChartRequest {
        Measure: core.CastString(data["measure"]),
        Grn: core.CastString(data["grn"]),
        Round: core.CastString(data["round"]),
        Filters: CastFilters(core.CastArray(data["filters"])),
        GroupBys: core.CastStrings(core.CastArray(data["groupBys"])),
        CountBy: core.CastString(data["countBy"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p GetChartRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "measure": p.Measure,
        "grn": p.Grn,
        "round": p.Round,
        "filters": CastFiltersFromDict(
            p.Filters,
        ),
        "groupBys": core.CastStringsFromDict(
            p.GroupBys,
        ),
        "countBy": p.CountBy,
        "begin": p.Begin,
        "end": p.End,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p GetChartRequest) Pointer() *GetChartRequest {
    return &p
}

type GetDistributionRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Measure *string `json:"measure"`
    Grn *string `json:"grn"`
    Filters []Filter `json:"filters"`
    GroupBys []*string `json:"groupBys"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewGetDistributionRequestFromJson(data string) GetDistributionRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetDistributionRequestFromDict(dict)
}

func NewGetDistributionRequestFromDict(data map[string]interface{}) GetDistributionRequest {
    return GetDistributionRequest {
        Measure: core.CastString(data["measure"]),
        Grn: core.CastString(data["grn"]),
        Filters: CastFilters(core.CastArray(data["filters"])),
        GroupBys: core.CastStrings(core.CastArray(data["groupBys"])),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p GetDistributionRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "measure": p.Measure,
        "grn": p.Grn,
        "filters": CastFiltersFromDict(
            p.Filters,
        ),
        "groupBys": core.CastStringsFromDict(
            p.GroupBys,
        ),
        "begin": p.Begin,
        "end": p.End,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p GetDistributionRequest) Pointer() *GetDistributionRequest {
    return &p
}

type GetCumulativeRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Name *string `json:"name"`
    ResourceGrn *string `json:"resourceGrn"`
}

func NewGetCumulativeRequestFromJson(data string) GetCumulativeRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCumulativeRequestFromDict(dict)
}

func NewGetCumulativeRequestFromDict(data map[string]interface{}) GetCumulativeRequest {
    return GetCumulativeRequest {
        Name: core.CastString(data["name"]),
        ResourceGrn: core.CastString(data["resourceGrn"]),
    }
}

func (p GetCumulativeRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "resourceGrn": p.ResourceGrn,
    }
}

func (p GetCumulativeRequest) Pointer() *GetCumulativeRequest {
    return &p
}

type DescribeBillingActivitiesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Year *int32 `json:"year"`
    Month *int32 `json:"month"`
    Service *string `json:"service"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeBillingActivitiesRequestFromJson(data string) DescribeBillingActivitiesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBillingActivitiesRequestFromDict(dict)
}

func NewDescribeBillingActivitiesRequestFromDict(data map[string]interface{}) DescribeBillingActivitiesRequest {
    return DescribeBillingActivitiesRequest {
        Year: core.CastInt32(data["year"]),
        Month: core.CastInt32(data["month"]),
        Service: core.CastString(data["service"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeBillingActivitiesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "year": p.Year,
        "month": p.Month,
        "service": p.Service,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeBillingActivitiesRequest) Pointer() *DescribeBillingActivitiesRequest {
    return &p
}

type GetBillingActivityRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Year *int32 `json:"year"`
    Month *int32 `json:"month"`
    Service *string `json:"service"`
    ActivityType *string `json:"activityType"`
}

func NewGetBillingActivityRequestFromJson(data string) GetBillingActivityRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBillingActivityRequestFromDict(dict)
}

func NewGetBillingActivityRequestFromDict(data map[string]interface{}) GetBillingActivityRequest {
    return GetBillingActivityRequest {
        Year: core.CastInt32(data["year"]),
        Month: core.CastInt32(data["month"]),
        Service: core.CastString(data["service"]),
        ActivityType: core.CastString(data["activityType"]),
    }
}

func (p GetBillingActivityRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "year": p.Year,
        "month": p.Month,
        "service": p.Service,
        "activityType": p.ActivityType,
    }
}

func (p GetBillingActivityRequest) Pointer() *GetBillingActivityRequest {
    return &p
}

type GetGeneralMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewGetGeneralMetricsRequestFromJson(data string) GetGeneralMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetGeneralMetricsRequestFromDict(dict)
}

func NewGetGeneralMetricsRequestFromDict(data map[string]interface{}) GetGeneralMetricsRequest {
    return GetGeneralMetricsRequest {
    }
}

func (p GetGeneralMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p GetGeneralMetricsRequest) Pointer() *GetGeneralMetricsRequest {
    return &p
}

type DescribeAccountNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeAccountNamespaceMetricsRequestFromJson(data string) DescribeAccountNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeAccountNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeAccountNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeAccountNamespaceMetricsRequest {
    return DescribeAccountNamespaceMetricsRequest {
    }
}

func (p DescribeAccountNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeAccountNamespaceMetricsRequest) Pointer() *DescribeAccountNamespaceMetricsRequest {
    return &p
}

type GetAccountNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetAccountNamespaceMetricsRequestFromJson(data string) GetAccountNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetAccountNamespaceMetricsRequestFromDict(dict)
}

func NewGetAccountNamespaceMetricsRequestFromDict(data map[string]interface{}) GetAccountNamespaceMetricsRequest {
    return GetAccountNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetAccountNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetAccountNamespaceMetricsRequest) Pointer() *GetAccountNamespaceMetricsRequest {
    return &p
}

type DescribeChatNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeChatNamespaceMetricsRequestFromJson(data string) DescribeChatNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeChatNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeChatNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeChatNamespaceMetricsRequest {
    return DescribeChatNamespaceMetricsRequest {
    }
}

func (p DescribeChatNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeChatNamespaceMetricsRequest) Pointer() *DescribeChatNamespaceMetricsRequest {
    return &p
}

type GetChatNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetChatNamespaceMetricsRequestFromJson(data string) GetChatNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetChatNamespaceMetricsRequestFromDict(dict)
}

func NewGetChatNamespaceMetricsRequestFromDict(data map[string]interface{}) GetChatNamespaceMetricsRequest {
    return GetChatNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetChatNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetChatNamespaceMetricsRequest) Pointer() *GetChatNamespaceMetricsRequest {
    return &p
}

type DescribeDatastoreNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeDatastoreNamespaceMetricsRequestFromJson(data string) DescribeDatastoreNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeDatastoreNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeDatastoreNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeDatastoreNamespaceMetricsRequest {
    return DescribeDatastoreNamespaceMetricsRequest {
    }
}

func (p DescribeDatastoreNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeDatastoreNamespaceMetricsRequest) Pointer() *DescribeDatastoreNamespaceMetricsRequest {
    return &p
}

type GetDatastoreNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetDatastoreNamespaceMetricsRequestFromJson(data string) GetDatastoreNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetDatastoreNamespaceMetricsRequestFromDict(dict)
}

func NewGetDatastoreNamespaceMetricsRequestFromDict(data map[string]interface{}) GetDatastoreNamespaceMetricsRequest {
    return GetDatastoreNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetDatastoreNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetDatastoreNamespaceMetricsRequest) Pointer() *GetDatastoreNamespaceMetricsRequest {
    return &p
}

type DescribeDictionaryNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeDictionaryNamespaceMetricsRequestFromJson(data string) DescribeDictionaryNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeDictionaryNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeDictionaryNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeDictionaryNamespaceMetricsRequest {
    return DescribeDictionaryNamespaceMetricsRequest {
    }
}

func (p DescribeDictionaryNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeDictionaryNamespaceMetricsRequest) Pointer() *DescribeDictionaryNamespaceMetricsRequest {
    return &p
}

type GetDictionaryNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetDictionaryNamespaceMetricsRequestFromJson(data string) GetDictionaryNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetDictionaryNamespaceMetricsRequestFromDict(dict)
}

func NewGetDictionaryNamespaceMetricsRequestFromDict(data map[string]interface{}) GetDictionaryNamespaceMetricsRequest {
    return GetDictionaryNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetDictionaryNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetDictionaryNamespaceMetricsRequest) Pointer() *GetDictionaryNamespaceMetricsRequest {
    return &p
}

type DescribeExchangeRateModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeExchangeRateModelMetricsRequestFromJson(data string) DescribeExchangeRateModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeExchangeRateModelMetricsRequestFromDict(dict)
}

func NewDescribeExchangeRateModelMetricsRequestFromDict(data map[string]interface{}) DescribeExchangeRateModelMetricsRequest {
    return DescribeExchangeRateModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeExchangeRateModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeExchangeRateModelMetricsRequest) Pointer() *DescribeExchangeRateModelMetricsRequest {
    return &p
}

type GetExchangeRateModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RateName *string `json:"rateName"`
}

func NewGetExchangeRateModelMetricsRequestFromJson(data string) GetExchangeRateModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetExchangeRateModelMetricsRequestFromDict(dict)
}

func NewGetExchangeRateModelMetricsRequestFromDict(data map[string]interface{}) GetExchangeRateModelMetricsRequest {
    return GetExchangeRateModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RateName: core.CastString(data["rateName"]),
    }
}

func (p GetExchangeRateModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "rateName": p.RateName,
    }
}

func (p GetExchangeRateModelMetricsRequest) Pointer() *GetExchangeRateModelMetricsRequest {
    return &p
}

type DescribeExchangeNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeExchangeNamespaceMetricsRequestFromJson(data string) DescribeExchangeNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeExchangeNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeExchangeNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeExchangeNamespaceMetricsRequest {
    return DescribeExchangeNamespaceMetricsRequest {
    }
}

func (p DescribeExchangeNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeExchangeNamespaceMetricsRequest) Pointer() *DescribeExchangeNamespaceMetricsRequest {
    return &p
}

type GetExchangeNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetExchangeNamespaceMetricsRequestFromJson(data string) GetExchangeNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetExchangeNamespaceMetricsRequestFromDict(dict)
}

func NewGetExchangeNamespaceMetricsRequestFromDict(data map[string]interface{}) GetExchangeNamespaceMetricsRequest {
    return GetExchangeNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetExchangeNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetExchangeNamespaceMetricsRequest) Pointer() *GetExchangeNamespaceMetricsRequest {
    return &p
}

type DescribeExperienceStatusMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ExperienceName *string `json:"experienceName"`
}

func NewDescribeExperienceStatusMetricsRequestFromJson(data string) DescribeExperienceStatusMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeExperienceStatusMetricsRequestFromDict(dict)
}

func NewDescribeExperienceStatusMetricsRequestFromDict(data map[string]interface{}) DescribeExperienceStatusMetricsRequest {
    return DescribeExperienceStatusMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ExperienceName: core.CastString(data["experienceName"]),
    }
}

func (p DescribeExperienceStatusMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "experienceName": p.ExperienceName,
    }
}

func (p DescribeExperienceStatusMetricsRequest) Pointer() *DescribeExperienceStatusMetricsRequest {
    return &p
}

type DescribeExperienceExperienceModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeExperienceExperienceModelMetricsRequestFromJson(data string) DescribeExperienceExperienceModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeExperienceExperienceModelMetricsRequestFromDict(dict)
}

func NewDescribeExperienceExperienceModelMetricsRequestFromDict(data map[string]interface{}) DescribeExperienceExperienceModelMetricsRequest {
    return DescribeExperienceExperienceModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeExperienceExperienceModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeExperienceExperienceModelMetricsRequest) Pointer() *DescribeExperienceExperienceModelMetricsRequest {
    return &p
}

type GetExperienceExperienceModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ExperienceName *string `json:"experienceName"`
}

func NewGetExperienceExperienceModelMetricsRequestFromJson(data string) GetExperienceExperienceModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetExperienceExperienceModelMetricsRequestFromDict(dict)
}

func NewGetExperienceExperienceModelMetricsRequestFromDict(data map[string]interface{}) GetExperienceExperienceModelMetricsRequest {
    return GetExperienceExperienceModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ExperienceName: core.CastString(data["experienceName"]),
    }
}

func (p GetExperienceExperienceModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "experienceName": p.ExperienceName,
    }
}

func (p GetExperienceExperienceModelMetricsRequest) Pointer() *GetExperienceExperienceModelMetricsRequest {
    return &p
}

type DescribeExperienceNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeExperienceNamespaceMetricsRequestFromJson(data string) DescribeExperienceNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeExperienceNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeExperienceNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeExperienceNamespaceMetricsRequest {
    return DescribeExperienceNamespaceMetricsRequest {
    }
}

func (p DescribeExperienceNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeExperienceNamespaceMetricsRequest) Pointer() *DescribeExperienceNamespaceMetricsRequest {
    return &p
}

type GetExperienceNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetExperienceNamespaceMetricsRequestFromJson(data string) GetExperienceNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetExperienceNamespaceMetricsRequestFromDict(dict)
}

func NewGetExperienceNamespaceMetricsRequestFromDict(data map[string]interface{}) GetExperienceNamespaceMetricsRequest {
    return GetExperienceNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetExperienceNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetExperienceNamespaceMetricsRequest) Pointer() *GetExperienceNamespaceMetricsRequest {
    return &p
}

type DescribeFormationFormMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    MoldName *string `json:"moldName"`
}

func NewDescribeFormationFormMetricsRequestFromJson(data string) DescribeFormationFormMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeFormationFormMetricsRequestFromDict(dict)
}

func NewDescribeFormationFormMetricsRequestFromDict(data map[string]interface{}) DescribeFormationFormMetricsRequest {
    return DescribeFormationFormMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        MoldName: core.CastString(data["moldName"]),
    }
}

func (p DescribeFormationFormMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "moldName": p.MoldName,
    }
}

func (p DescribeFormationFormMetricsRequest) Pointer() *DescribeFormationFormMetricsRequest {
    return &p
}

type DescribeFormationMoldMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeFormationMoldMetricsRequestFromJson(data string) DescribeFormationMoldMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeFormationMoldMetricsRequestFromDict(dict)
}

func NewDescribeFormationMoldMetricsRequestFromDict(data map[string]interface{}) DescribeFormationMoldMetricsRequest {
    return DescribeFormationMoldMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeFormationMoldMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeFormationMoldMetricsRequest) Pointer() *DescribeFormationMoldMetricsRequest {
    return &p
}

type DescribeFormationNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeFormationNamespaceMetricsRequestFromJson(data string) DescribeFormationNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeFormationNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeFormationNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeFormationNamespaceMetricsRequest {
    return DescribeFormationNamespaceMetricsRequest {
    }
}

func (p DescribeFormationNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeFormationNamespaceMetricsRequest) Pointer() *DescribeFormationNamespaceMetricsRequest {
    return &p
}

type GetFormationNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetFormationNamespaceMetricsRequestFromJson(data string) GetFormationNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetFormationNamespaceMetricsRequestFromDict(dict)
}

func NewGetFormationNamespaceMetricsRequestFromDict(data map[string]interface{}) GetFormationNamespaceMetricsRequest {
    return GetFormationNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetFormationNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetFormationNamespaceMetricsRequest) Pointer() *GetFormationNamespaceMetricsRequest {
    return &p
}

type DescribeFriendNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeFriendNamespaceMetricsRequestFromJson(data string) DescribeFriendNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeFriendNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeFriendNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeFriendNamespaceMetricsRequest {
    return DescribeFriendNamespaceMetricsRequest {
    }
}

func (p DescribeFriendNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeFriendNamespaceMetricsRequest) Pointer() *DescribeFriendNamespaceMetricsRequest {
    return &p
}

type GetFriendNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetFriendNamespaceMetricsRequestFromJson(data string) GetFriendNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetFriendNamespaceMetricsRequestFromDict(dict)
}

func NewGetFriendNamespaceMetricsRequestFromDict(data map[string]interface{}) GetFriendNamespaceMetricsRequest {
    return GetFriendNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetFriendNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetFriendNamespaceMetricsRequest) Pointer() *GetFriendNamespaceMetricsRequest {
    return &p
}

type DescribeInboxNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeInboxNamespaceMetricsRequestFromJson(data string) DescribeInboxNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInboxNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeInboxNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeInboxNamespaceMetricsRequest {
    return DescribeInboxNamespaceMetricsRequest {
    }
}

func (p DescribeInboxNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeInboxNamespaceMetricsRequest) Pointer() *DescribeInboxNamespaceMetricsRequest {
    return &p
}

type GetInboxNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetInboxNamespaceMetricsRequestFromJson(data string) GetInboxNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInboxNamespaceMetricsRequestFromDict(dict)
}

func NewGetInboxNamespaceMetricsRequestFromDict(data map[string]interface{}) GetInboxNamespaceMetricsRequest {
    return GetInboxNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetInboxNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetInboxNamespaceMetricsRequest) Pointer() *GetInboxNamespaceMetricsRequest {
    return &p
}

type DescribeInventoryItemSetMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    InventoryName *string `json:"inventoryName"`
}

func NewDescribeInventoryItemSetMetricsRequestFromJson(data string) DescribeInventoryItemSetMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoryItemSetMetricsRequestFromDict(dict)
}

func NewDescribeInventoryItemSetMetricsRequestFromDict(data map[string]interface{}) DescribeInventoryItemSetMetricsRequest {
    return DescribeInventoryItemSetMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        InventoryName: core.CastString(data["inventoryName"]),
    }
}

func (p DescribeInventoryItemSetMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "inventoryName": p.InventoryName,
    }
}

func (p DescribeInventoryItemSetMetricsRequest) Pointer() *DescribeInventoryItemSetMetricsRequest {
    return &p
}

type DescribeInventoryInventoryMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeInventoryInventoryMetricsRequestFromJson(data string) DescribeInventoryInventoryMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoryInventoryMetricsRequestFromDict(dict)
}

func NewDescribeInventoryInventoryMetricsRequestFromDict(data map[string]interface{}) DescribeInventoryInventoryMetricsRequest {
    return DescribeInventoryInventoryMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeInventoryInventoryMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeInventoryInventoryMetricsRequest) Pointer() *DescribeInventoryInventoryMetricsRequest {
    return &p
}

type DescribeInventoryNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeInventoryNamespaceMetricsRequestFromJson(data string) DescribeInventoryNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeInventoryNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeInventoryNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeInventoryNamespaceMetricsRequest {
    return DescribeInventoryNamespaceMetricsRequest {
    }
}

func (p DescribeInventoryNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeInventoryNamespaceMetricsRequest) Pointer() *DescribeInventoryNamespaceMetricsRequest {
    return &p
}

type GetInventoryNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetInventoryNamespaceMetricsRequestFromJson(data string) GetInventoryNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetInventoryNamespaceMetricsRequestFromDict(dict)
}

func NewGetInventoryNamespaceMetricsRequestFromDict(data map[string]interface{}) GetInventoryNamespaceMetricsRequest {
    return GetInventoryNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetInventoryNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetInventoryNamespaceMetricsRequest) Pointer() *GetInventoryNamespaceMetricsRequest {
    return &p
}

type DescribeKeyNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeKeyNamespaceMetricsRequestFromJson(data string) DescribeKeyNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeKeyNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeKeyNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeKeyNamespaceMetricsRequest {
    return DescribeKeyNamespaceMetricsRequest {
    }
}

func (p DescribeKeyNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeKeyNamespaceMetricsRequest) Pointer() *DescribeKeyNamespaceMetricsRequest {
    return &p
}

type GetKeyNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetKeyNamespaceMetricsRequestFromJson(data string) GetKeyNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetKeyNamespaceMetricsRequestFromDict(dict)
}

func NewGetKeyNamespaceMetricsRequestFromDict(data map[string]interface{}) GetKeyNamespaceMetricsRequest {
    return GetKeyNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetKeyNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetKeyNamespaceMetricsRequest) Pointer() *GetKeyNamespaceMetricsRequest {
    return &p
}

type DescribeLimitCounterMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
}

func NewDescribeLimitCounterMetricsRequestFromJson(data string) DescribeLimitCounterMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeLimitCounterMetricsRequestFromDict(dict)
}

func NewDescribeLimitCounterMetricsRequestFromDict(data map[string]interface{}) DescribeLimitCounterMetricsRequest {
    return DescribeLimitCounterMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
    }
}

func (p DescribeLimitCounterMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
    }
}

func (p DescribeLimitCounterMetricsRequest) Pointer() *DescribeLimitCounterMetricsRequest {
    return &p
}

type DescribeLimitLimitModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeLimitLimitModelMetricsRequestFromJson(data string) DescribeLimitLimitModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeLimitLimitModelMetricsRequestFromDict(dict)
}

func NewDescribeLimitLimitModelMetricsRequestFromDict(data map[string]interface{}) DescribeLimitLimitModelMetricsRequest {
    return DescribeLimitLimitModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeLimitLimitModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeLimitLimitModelMetricsRequest) Pointer() *DescribeLimitLimitModelMetricsRequest {
    return &p
}

type GetLimitLimitModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    LimitName *string `json:"limitName"`
}

func NewGetLimitLimitModelMetricsRequestFromJson(data string) GetLimitLimitModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetLimitLimitModelMetricsRequestFromDict(dict)
}

func NewGetLimitLimitModelMetricsRequestFromDict(data map[string]interface{}) GetLimitLimitModelMetricsRequest {
    return GetLimitLimitModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LimitName: core.CastString(data["limitName"]),
    }
}

func (p GetLimitLimitModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "limitName": p.LimitName,
    }
}

func (p GetLimitLimitModelMetricsRequest) Pointer() *GetLimitLimitModelMetricsRequest {
    return &p
}

type DescribeLimitNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeLimitNamespaceMetricsRequestFromJson(data string) DescribeLimitNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeLimitNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeLimitNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeLimitNamespaceMetricsRequest {
    return DescribeLimitNamespaceMetricsRequest {
    }
}

func (p DescribeLimitNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeLimitNamespaceMetricsRequest) Pointer() *DescribeLimitNamespaceMetricsRequest {
    return &p
}

type GetLimitNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetLimitNamespaceMetricsRequestFromJson(data string) GetLimitNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetLimitNamespaceMetricsRequestFromDict(dict)
}

func NewGetLimitNamespaceMetricsRequestFromDict(data map[string]interface{}) GetLimitNamespaceMetricsRequest {
    return GetLimitNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetLimitNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetLimitNamespaceMetricsRequest) Pointer() *GetLimitNamespaceMetricsRequest {
    return &p
}

type DescribeLotteryLotteryMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeLotteryLotteryMetricsRequestFromJson(data string) DescribeLotteryLotteryMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeLotteryLotteryMetricsRequestFromDict(dict)
}

func NewDescribeLotteryLotteryMetricsRequestFromDict(data map[string]interface{}) DescribeLotteryLotteryMetricsRequest {
    return DescribeLotteryLotteryMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeLotteryLotteryMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeLotteryLotteryMetricsRequest) Pointer() *DescribeLotteryLotteryMetricsRequest {
    return &p
}

type GetLotteryLotteryMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    LotteryName *string `json:"lotteryName"`
}

func NewGetLotteryLotteryMetricsRequestFromJson(data string) GetLotteryLotteryMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetLotteryLotteryMetricsRequestFromDict(dict)
}

func NewGetLotteryLotteryMetricsRequestFromDict(data map[string]interface{}) GetLotteryLotteryMetricsRequest {
    return GetLotteryLotteryMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        LotteryName: core.CastString(data["lotteryName"]),
    }
}

func (p GetLotteryLotteryMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "lotteryName": p.LotteryName,
    }
}

func (p GetLotteryLotteryMetricsRequest) Pointer() *GetLotteryLotteryMetricsRequest {
    return &p
}

type DescribeLotteryNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeLotteryNamespaceMetricsRequestFromJson(data string) DescribeLotteryNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeLotteryNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeLotteryNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeLotteryNamespaceMetricsRequest {
    return DescribeLotteryNamespaceMetricsRequest {
    }
}

func (p DescribeLotteryNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeLotteryNamespaceMetricsRequest) Pointer() *DescribeLotteryNamespaceMetricsRequest {
    return &p
}

type GetLotteryNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetLotteryNamespaceMetricsRequestFromJson(data string) GetLotteryNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetLotteryNamespaceMetricsRequestFromDict(dict)
}

func NewGetLotteryNamespaceMetricsRequestFromDict(data map[string]interface{}) GetLotteryNamespaceMetricsRequest {
    return GetLotteryNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetLotteryNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetLotteryNamespaceMetricsRequest) Pointer() *GetLotteryNamespaceMetricsRequest {
    return &p
}

type DescribeMatchmakingNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeMatchmakingNamespaceMetricsRequestFromJson(data string) DescribeMatchmakingNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeMatchmakingNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeMatchmakingNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeMatchmakingNamespaceMetricsRequest {
    return DescribeMatchmakingNamespaceMetricsRequest {
    }
}

func (p DescribeMatchmakingNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeMatchmakingNamespaceMetricsRequest) Pointer() *DescribeMatchmakingNamespaceMetricsRequest {
    return &p
}

type GetMatchmakingNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetMatchmakingNamespaceMetricsRequestFromJson(data string) GetMatchmakingNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetMatchmakingNamespaceMetricsRequestFromDict(dict)
}

func NewGetMatchmakingNamespaceMetricsRequestFromDict(data map[string]interface{}) GetMatchmakingNamespaceMetricsRequest {
    return GetMatchmakingNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetMatchmakingNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetMatchmakingNamespaceMetricsRequest) Pointer() *GetMatchmakingNamespaceMetricsRequest {
    return &p
}

type DescribeMissionCounterMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeMissionCounterMetricsRequestFromJson(data string) DescribeMissionCounterMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeMissionCounterMetricsRequestFromDict(dict)
}

func NewDescribeMissionCounterMetricsRequestFromDict(data map[string]interface{}) DescribeMissionCounterMetricsRequest {
    return DescribeMissionCounterMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeMissionCounterMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeMissionCounterMetricsRequest) Pointer() *DescribeMissionCounterMetricsRequest {
    return &p
}

type DescribeMissionMissionGroupModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeMissionMissionGroupModelMetricsRequestFromJson(data string) DescribeMissionMissionGroupModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeMissionMissionGroupModelMetricsRequestFromDict(dict)
}

func NewDescribeMissionMissionGroupModelMetricsRequestFromDict(data map[string]interface{}) DescribeMissionMissionGroupModelMetricsRequest {
    return DescribeMissionMissionGroupModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeMissionMissionGroupModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeMissionMissionGroupModelMetricsRequest) Pointer() *DescribeMissionMissionGroupModelMetricsRequest {
    return &p
}

type GetMissionMissionGroupModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    MissionGroupName *string `json:"missionGroupName"`
}

func NewGetMissionMissionGroupModelMetricsRequestFromJson(data string) GetMissionMissionGroupModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetMissionMissionGroupModelMetricsRequestFromDict(dict)
}

func NewGetMissionMissionGroupModelMetricsRequestFromDict(data map[string]interface{}) GetMissionMissionGroupModelMetricsRequest {
    return GetMissionMissionGroupModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        MissionGroupName: core.CastString(data["missionGroupName"]),
    }
}

func (p GetMissionMissionGroupModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "missionGroupName": p.MissionGroupName,
    }
}

func (p GetMissionMissionGroupModelMetricsRequest) Pointer() *GetMissionMissionGroupModelMetricsRequest {
    return &p
}

type DescribeMissionNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeMissionNamespaceMetricsRequestFromJson(data string) DescribeMissionNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeMissionNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeMissionNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeMissionNamespaceMetricsRequest {
    return DescribeMissionNamespaceMetricsRequest {
    }
}

func (p DescribeMissionNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeMissionNamespaceMetricsRequest) Pointer() *DescribeMissionNamespaceMetricsRequest {
    return &p
}

type GetMissionNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetMissionNamespaceMetricsRequestFromJson(data string) GetMissionNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetMissionNamespaceMetricsRequestFromDict(dict)
}

func NewGetMissionNamespaceMetricsRequestFromDict(data map[string]interface{}) GetMissionNamespaceMetricsRequest {
    return GetMissionNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetMissionNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetMissionNamespaceMetricsRequest) Pointer() *GetMissionNamespaceMetricsRequest {
    return &p
}

type DescribeMoneyWalletMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeMoneyWalletMetricsRequestFromJson(data string) DescribeMoneyWalletMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeMoneyWalletMetricsRequestFromDict(dict)
}

func NewDescribeMoneyWalletMetricsRequestFromDict(data map[string]interface{}) DescribeMoneyWalletMetricsRequest {
    return DescribeMoneyWalletMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeMoneyWalletMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeMoneyWalletMetricsRequest) Pointer() *DescribeMoneyWalletMetricsRequest {
    return &p
}

type DescribeMoneyReceiptMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeMoneyReceiptMetricsRequestFromJson(data string) DescribeMoneyReceiptMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeMoneyReceiptMetricsRequestFromDict(dict)
}

func NewDescribeMoneyReceiptMetricsRequestFromDict(data map[string]interface{}) DescribeMoneyReceiptMetricsRequest {
    return DescribeMoneyReceiptMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeMoneyReceiptMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeMoneyReceiptMetricsRequest) Pointer() *DescribeMoneyReceiptMetricsRequest {
    return &p
}

type DescribeMoneyNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeMoneyNamespaceMetricsRequestFromJson(data string) DescribeMoneyNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeMoneyNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeMoneyNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeMoneyNamespaceMetricsRequest {
    return DescribeMoneyNamespaceMetricsRequest {
    }
}

func (p DescribeMoneyNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeMoneyNamespaceMetricsRequest) Pointer() *DescribeMoneyNamespaceMetricsRequest {
    return &p
}

type GetMoneyNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetMoneyNamespaceMetricsRequestFromJson(data string) GetMoneyNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetMoneyNamespaceMetricsRequestFromDict(dict)
}

func NewGetMoneyNamespaceMetricsRequestFromDict(data map[string]interface{}) GetMoneyNamespaceMetricsRequest {
    return GetMoneyNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetMoneyNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetMoneyNamespaceMetricsRequest) Pointer() *GetMoneyNamespaceMetricsRequest {
    return &p
}

type DescribeQuestQuestModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
}

func NewDescribeQuestQuestModelMetricsRequestFromJson(data string) DescribeQuestQuestModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeQuestQuestModelMetricsRequestFromDict(dict)
}

func NewDescribeQuestQuestModelMetricsRequestFromDict(data map[string]interface{}) DescribeQuestQuestModelMetricsRequest {
    return DescribeQuestQuestModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
    }
}

func (p DescribeQuestQuestModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
    }
}

func (p DescribeQuestQuestModelMetricsRequest) Pointer() *DescribeQuestQuestModelMetricsRequest {
    return &p
}

type GetQuestQuestModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
    QuestName *string `json:"questName"`
}

func NewGetQuestQuestModelMetricsRequestFromJson(data string) GetQuestQuestModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetQuestQuestModelMetricsRequestFromDict(dict)
}

func NewGetQuestQuestModelMetricsRequestFromDict(data map[string]interface{}) GetQuestQuestModelMetricsRequest {
    return GetQuestQuestModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        QuestName: core.CastString(data["questName"]),
    }
}

func (p GetQuestQuestModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
        "questName": p.QuestName,
    }
}

func (p GetQuestQuestModelMetricsRequest) Pointer() *GetQuestQuestModelMetricsRequest {
    return &p
}

type DescribeQuestQuestGroupModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeQuestQuestGroupModelMetricsRequestFromJson(data string) DescribeQuestQuestGroupModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeQuestQuestGroupModelMetricsRequestFromDict(dict)
}

func NewDescribeQuestQuestGroupModelMetricsRequestFromDict(data map[string]interface{}) DescribeQuestQuestGroupModelMetricsRequest {
    return DescribeQuestQuestGroupModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeQuestQuestGroupModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeQuestQuestGroupModelMetricsRequest) Pointer() *DescribeQuestQuestGroupModelMetricsRequest {
    return &p
}

type GetQuestQuestGroupModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    QuestGroupName *string `json:"questGroupName"`
}

func NewGetQuestQuestGroupModelMetricsRequestFromJson(data string) GetQuestQuestGroupModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetQuestQuestGroupModelMetricsRequestFromDict(dict)
}

func NewGetQuestQuestGroupModelMetricsRequestFromDict(data map[string]interface{}) GetQuestQuestGroupModelMetricsRequest {
    return GetQuestQuestGroupModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
    }
}

func (p GetQuestQuestGroupModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "questGroupName": p.QuestGroupName,
    }
}

func (p GetQuestQuestGroupModelMetricsRequest) Pointer() *GetQuestQuestGroupModelMetricsRequest {
    return &p
}

type DescribeQuestNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeQuestNamespaceMetricsRequestFromJson(data string) DescribeQuestNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeQuestNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeQuestNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeQuestNamespaceMetricsRequest {
    return DescribeQuestNamespaceMetricsRequest {
    }
}

func (p DescribeQuestNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeQuestNamespaceMetricsRequest) Pointer() *DescribeQuestNamespaceMetricsRequest {
    return &p
}

type GetQuestNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetQuestNamespaceMetricsRequestFromJson(data string) GetQuestNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetQuestNamespaceMetricsRequestFromDict(dict)
}

func NewGetQuestNamespaceMetricsRequestFromDict(data map[string]interface{}) GetQuestNamespaceMetricsRequest {
    return GetQuestNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetQuestNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetQuestNamespaceMetricsRequest) Pointer() *GetQuestNamespaceMetricsRequest {
    return &p
}

type DescribeRankingCategoryModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeRankingCategoryModelMetricsRequestFromJson(data string) DescribeRankingCategoryModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRankingCategoryModelMetricsRequestFromDict(dict)
}

func NewDescribeRankingCategoryModelMetricsRequestFromDict(data map[string]interface{}) DescribeRankingCategoryModelMetricsRequest {
    return DescribeRankingCategoryModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeRankingCategoryModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeRankingCategoryModelMetricsRequest) Pointer() *DescribeRankingCategoryModelMetricsRequest {
    return &p
}

type GetRankingCategoryModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    CategoryName *string `json:"categoryName"`
}

func NewGetRankingCategoryModelMetricsRequestFromJson(data string) GetRankingCategoryModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRankingCategoryModelMetricsRequestFromDict(dict)
}

func NewGetRankingCategoryModelMetricsRequestFromDict(data map[string]interface{}) GetRankingCategoryModelMetricsRequest {
    return GetRankingCategoryModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CategoryName: core.CastString(data["categoryName"]),
    }
}

func (p GetRankingCategoryModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "categoryName": p.CategoryName,
    }
}

func (p GetRankingCategoryModelMetricsRequest) Pointer() *GetRankingCategoryModelMetricsRequest {
    return &p
}

type DescribeRankingNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeRankingNamespaceMetricsRequestFromJson(data string) DescribeRankingNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRankingNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeRankingNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeRankingNamespaceMetricsRequest {
    return DescribeRankingNamespaceMetricsRequest {
    }
}

func (p DescribeRankingNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeRankingNamespaceMetricsRequest) Pointer() *DescribeRankingNamespaceMetricsRequest {
    return &p
}

type GetRankingNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetRankingNamespaceMetricsRequestFromJson(data string) GetRankingNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRankingNamespaceMetricsRequestFromDict(dict)
}

func NewGetRankingNamespaceMetricsRequestFromDict(data map[string]interface{}) GetRankingNamespaceMetricsRequest {
    return GetRankingNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetRankingNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetRankingNamespaceMetricsRequest) Pointer() *GetRankingNamespaceMetricsRequest {
    return &p
}

type DescribeShowcaseDisplayItemMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
}

func NewDescribeShowcaseDisplayItemMetricsRequestFromJson(data string) DescribeShowcaseDisplayItemMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeShowcaseDisplayItemMetricsRequestFromDict(dict)
}

func NewDescribeShowcaseDisplayItemMetricsRequestFromDict(data map[string]interface{}) DescribeShowcaseDisplayItemMetricsRequest {
    return DescribeShowcaseDisplayItemMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
    }
}

func (p DescribeShowcaseDisplayItemMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
    }
}

func (p DescribeShowcaseDisplayItemMetricsRequest) Pointer() *DescribeShowcaseDisplayItemMetricsRequest {
    return &p
}

type GetShowcaseDisplayItemMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
    DisplayItemId *string `json:"displayItemId"`
}

func NewGetShowcaseDisplayItemMetricsRequestFromJson(data string) GetShowcaseDisplayItemMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetShowcaseDisplayItemMetricsRequestFromDict(dict)
}

func NewGetShowcaseDisplayItemMetricsRequestFromDict(data map[string]interface{}) GetShowcaseDisplayItemMetricsRequest {
    return GetShowcaseDisplayItemMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
        DisplayItemId: core.CastString(data["displayItemId"]),
    }
}

func (p GetShowcaseDisplayItemMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
        "displayItemId": p.DisplayItemId,
    }
}

func (p GetShowcaseDisplayItemMetricsRequest) Pointer() *GetShowcaseDisplayItemMetricsRequest {
    return &p
}

type DescribeShowcaseShowcaseMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeShowcaseShowcaseMetricsRequestFromJson(data string) DescribeShowcaseShowcaseMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeShowcaseShowcaseMetricsRequestFromDict(dict)
}

func NewDescribeShowcaseShowcaseMetricsRequestFromDict(data map[string]interface{}) DescribeShowcaseShowcaseMetricsRequest {
    return DescribeShowcaseShowcaseMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeShowcaseShowcaseMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeShowcaseShowcaseMetricsRequest) Pointer() *DescribeShowcaseShowcaseMetricsRequest {
    return &p
}

type GetShowcaseShowcaseMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    ShowcaseName *string `json:"showcaseName"`
}

func NewGetShowcaseShowcaseMetricsRequestFromJson(data string) GetShowcaseShowcaseMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetShowcaseShowcaseMetricsRequestFromDict(dict)
}

func NewGetShowcaseShowcaseMetricsRequestFromDict(data map[string]interface{}) GetShowcaseShowcaseMetricsRequest {
    return GetShowcaseShowcaseMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        ShowcaseName: core.CastString(data["showcaseName"]),
    }
}

func (p GetShowcaseShowcaseMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "showcaseName": p.ShowcaseName,
    }
}

func (p GetShowcaseShowcaseMetricsRequest) Pointer() *GetShowcaseShowcaseMetricsRequest {
    return &p
}

type DescribeShowcaseNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeShowcaseNamespaceMetricsRequestFromJson(data string) DescribeShowcaseNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeShowcaseNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeShowcaseNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeShowcaseNamespaceMetricsRequest {
    return DescribeShowcaseNamespaceMetricsRequest {
    }
}

func (p DescribeShowcaseNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeShowcaseNamespaceMetricsRequest) Pointer() *DescribeShowcaseNamespaceMetricsRequest {
    return &p
}

type GetShowcaseNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetShowcaseNamespaceMetricsRequestFromJson(data string) GetShowcaseNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetShowcaseNamespaceMetricsRequestFromDict(dict)
}

func NewGetShowcaseNamespaceMetricsRequestFromDict(data map[string]interface{}) GetShowcaseNamespaceMetricsRequest {
    return GetShowcaseNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetShowcaseNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetShowcaseNamespaceMetricsRequest) Pointer() *GetShowcaseNamespaceMetricsRequest {
    return &p
}

type DescribeStaminaStaminaModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeStaminaStaminaModelMetricsRequestFromJson(data string) DescribeStaminaStaminaModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeStaminaStaminaModelMetricsRequestFromDict(dict)
}

func NewDescribeStaminaStaminaModelMetricsRequestFromDict(data map[string]interface{}) DescribeStaminaStaminaModelMetricsRequest {
    return DescribeStaminaStaminaModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeStaminaStaminaModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeStaminaStaminaModelMetricsRequest) Pointer() *DescribeStaminaStaminaModelMetricsRequest {
    return &p
}

type GetStaminaStaminaModelMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    StaminaName *string `json:"staminaName"`
}

func NewGetStaminaStaminaModelMetricsRequestFromJson(data string) GetStaminaStaminaModelMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetStaminaStaminaModelMetricsRequestFromDict(dict)
}

func NewGetStaminaStaminaModelMetricsRequestFromDict(data map[string]interface{}) GetStaminaStaminaModelMetricsRequest {
    return GetStaminaStaminaModelMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        StaminaName: core.CastString(data["staminaName"]),
    }
}

func (p GetStaminaStaminaModelMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "staminaName": p.StaminaName,
    }
}

func (p GetStaminaStaminaModelMetricsRequest) Pointer() *GetStaminaStaminaModelMetricsRequest {
    return &p
}

type DescribeStaminaNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
}

func NewDescribeStaminaNamespaceMetricsRequestFromJson(data string) DescribeStaminaNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeStaminaNamespaceMetricsRequestFromDict(dict)
}

func NewDescribeStaminaNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeStaminaNamespaceMetricsRequest {
    return DescribeStaminaNamespaceMetricsRequest {
    }
}

func (p DescribeStaminaNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DescribeStaminaNamespaceMetricsRequest) Pointer() *DescribeStaminaNamespaceMetricsRequest {
    return &p
}

type GetStaminaNamespaceMetricsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetStaminaNamespaceMetricsRequestFromJson(data string) GetStaminaNamespaceMetricsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetStaminaNamespaceMetricsRequestFromDict(dict)
}

func NewGetStaminaNamespaceMetricsRequestFromDict(data map[string]interface{}) GetStaminaNamespaceMetricsRequest {
    return GetStaminaNamespaceMetricsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetStaminaNamespaceMetricsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetStaminaNamespaceMetricsRequest) Pointer() *GetStaminaNamespaceMetricsRequest {
    return &p
}