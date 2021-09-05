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
    return map[string]interface{} {
        "chartId": p.ChartId,
        "embedId": p.EmbedId,
        "html": p.Html,
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
    return map[string]interface{} {
        "cumulativeId": p.CumulativeId,
        "resourceGrn": p.ResourceGrn,
        "name": p.Name,
        "value": p.Value,
        "updatedAt": p.UpdatedAt,
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
    return map[string]interface{} {
        "billingActivityId": p.BillingActivityId,
        "year": p.Year,
        "month": p.Month,
        "service": p.Service,
        "activityType": p.ActivityType,
        "value": p.Value,
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