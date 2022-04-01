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

type Chart struct {
	ChartId *string `json:"chartId"`
	EmbedId *string `json:"embedId"`
	Html *string `json:"html"`
}

func NewChartFromJson(data string) Chart {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewChartFromDict(dict)
}

func NewChartFromDict(data map[string]interface{}) Chart {
    return Chart {
        ChartId: core.CastString(data["chartId"]),
        EmbedId: core.CastString(data["embedId"]),
        Html: core.CastString(data["html"]),
    }
}

func (p Chart) ToDict() map[string]interface{} {
    
    var chartId *string
    if p.ChartId != nil {
        chartId = p.ChartId
    }
    var embedId *string
    if p.EmbedId != nil {
        embedId = p.EmbedId
    }
    var html *string
    if p.Html != nil {
        html = p.Html
    }
    return map[string]interface{} {
        "chartId": chartId,
        "embedId": embedId,
        "html": html,
    }
}

func (p Chart) Pointer() *Chart {
    return &p
}

func CastCharts(data []interface{}) []Chart {
	v := make([]Chart, 0)
	for _, d := range data {
		v = append(v, NewChartFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChartsFromDict(data []Chart) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Cumulative struct {
	CumulativeId *string `json:"cumulativeId"`
	ResourceGrn *string `json:"resourceGrn"`
	Name *string `json:"name"`
	Value *int64 `json:"value"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewCumulativeFromJson(data string) Cumulative {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCumulativeFromDict(dict)
}

func NewCumulativeFromDict(data map[string]interface{}) Cumulative {
    return Cumulative {
        CumulativeId: core.CastString(data["cumulativeId"]),
        ResourceGrn: core.CastString(data["resourceGrn"]),
        Name: core.CastString(data["name"]),
        Value: core.CastInt64(data["value"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Cumulative) ToDict() map[string]interface{} {
    
    var cumulativeId *string
    if p.CumulativeId != nil {
        cumulativeId = p.CumulativeId
    }
    var resourceGrn *string
    if p.ResourceGrn != nil {
        resourceGrn = p.ResourceGrn
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var value *int64
    if p.Value != nil {
        value = p.Value
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "cumulativeId": cumulativeId,
        "resourceGrn": resourceGrn,
        "name": name,
        "value": value,
        "updatedAt": updatedAt,
    }
}

func (p Cumulative) Pointer() *Cumulative {
    return &p
}

func CastCumulatives(data []interface{}) []Cumulative {
	v := make([]Cumulative, 0)
	for _, d := range data {
		v = append(v, NewCumulativeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCumulativesFromDict(data []Cumulative) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type BillingActivity struct {
	BillingActivityId *string `json:"billingActivityId"`
	Year *int32 `json:"year"`
	Month *int32 `json:"month"`
	Service *string `json:"service"`
	ActivityType *string `json:"activityType"`
	Value *int64 `json:"value"`
}

func NewBillingActivityFromJson(data string) BillingActivity {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBillingActivityFromDict(dict)
}

func NewBillingActivityFromDict(data map[string]interface{}) BillingActivity {
    return BillingActivity {
        BillingActivityId: core.CastString(data["billingActivityId"]),
        Year: core.CastInt32(data["year"]),
        Month: core.CastInt32(data["month"]),
        Service: core.CastString(data["service"]),
        ActivityType: core.CastString(data["activityType"]),
        Value: core.CastInt64(data["value"]),
    }
}

func (p BillingActivity) ToDict() map[string]interface{} {
    
    var billingActivityId *string
    if p.BillingActivityId != nil {
        billingActivityId = p.BillingActivityId
    }
    var year *int32
    if p.Year != nil {
        year = p.Year
    }
    var month *int32
    if p.Month != nil {
        month = p.Month
    }
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var activityType *string
    if p.ActivityType != nil {
        activityType = p.ActivityType
    }
    var value *int64
    if p.Value != nil {
        value = p.Value
    }
    return map[string]interface{} {
        "billingActivityId": billingActivityId,
        "year": year,
        "month": month,
        "service": service,
        "activityType": activityType,
        "value": value,
    }
}

func (p BillingActivity) Pointer() *BillingActivity {
    return &p
}

func CastBillingActivities(data []interface{}) []BillingActivity {
	v := make([]BillingActivity, 0)
	for _, d := range data {
		v = append(v, NewBillingActivityFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBillingActivitiesFromDict(data []BillingActivity) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type StatsEvent struct {
	Grn *string `json:"grn"`
	Service *string `json:"service"`
	Method *string `json:"method"`
	Metric *string `json:"metric"`
	Cumulative *bool `json:"cumulative"`
	Value *float64 `json:"value"`
	Tags []string `json:"tags"`
	CallAt *int64 `json:"callAt"`
}

func NewStatsEventFromJson(data string) StatsEvent {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStatsEventFromDict(dict)
}

func NewStatsEventFromDict(data map[string]interface{}) StatsEvent {
    return StatsEvent {
        Grn: core.CastString(data["grn"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        Metric: core.CastString(data["metric"]),
        Cumulative: core.CastBool(data["cumulative"]),
        Value: core.CastFloat64(data["value"]),
        Tags: core.CastStrings(core.CastArray(data["tags"])),
        CallAt: core.CastInt64(data["callAt"]),
    }
}

func (p StatsEvent) ToDict() map[string]interface{} {
    
    var grn *string
    if p.Grn != nil {
        grn = p.Grn
    }
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var metric *string
    if p.Metric != nil {
        metric = p.Metric
    }
    var cumulative *bool
    if p.Cumulative != nil {
        cumulative = p.Cumulative
    }
    var value *float64
    if p.Value != nil {
        value = p.Value
    }
    var tags []interface{}
    if p.Tags != nil {
        tags = core.CastStringsFromDict(
            p.Tags,
        )
    }
    var callAt *int64
    if p.CallAt != nil {
        callAt = p.CallAt
    }
    return map[string]interface{} {
        "grn": grn,
        "service": service,
        "method": method,
        "metric": metric,
        "cumulative": cumulative,
        "value": value,
        "tags": tags,
        "callAt": callAt,
    }
}

func (p StatsEvent) Pointer() *StatsEvent {
    return &p
}

func CastStatsEvents(data []interface{}) []StatsEvent {
	v := make([]StatsEvent, 0)
	for _, d := range data {
		v = append(v, NewStatsEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStatsEventsFromDict(data []StatsEvent) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}