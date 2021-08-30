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

import "core"

type GetChartRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	Metrics            *string  `json:"metrics"`
	Grn                *string  `json:"grn"`
	Queries            []string `json:"queries"`
	By                 *string  `json:"by"`
	Timeframe          *string  `json:"timeframe"`
	Size               *string  `json:"size"`
	Format             *string  `json:"format"`
	Aggregator         *string  `json:"aggregator"`
	Style              *string  `json:"style"`
	Title              *string  `json:"title"`
}

func NewGetChartRequestFromDict(data map[string]interface{}) GetChartRequest {
	return GetChartRequest{
		Metrics:    core.CastString(data["metrics"]),
		Grn:        core.CastString(data["grn"]),
		Queries:    core.CastStrings(core.CastArray(data["queries"])),
		By:         core.CastString(data["by"]),
		Timeframe:  core.CastString(data["timeframe"]),
		Size:       core.CastString(data["size"]),
		Format:     core.CastString(data["format"]),
		Aggregator: core.CastString(data["aggregator"]),
		Style:      core.CastString(data["style"]),
		Title:      core.CastString(data["title"]),
	}
}

func (p GetChartRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metrics": p.Metrics,
		"grn":     p.Grn,
		"queries": core.CastStringsFromDict(
			p.Queries,
		),
		"by":         p.By,
		"timeframe":  p.Timeframe,
		"size":       p.Size,
		"format":     p.Format,
		"aggregator": p.Aggregator,
		"style":      p.Style,
		"title":      p.Title,
	}
}

func (p GetChartRequest) Pointer() *GetChartRequest {
	return &p
}

type GetCumulativeRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	Name               *string `json:"name"`
	ResourceGrn        *string `json:"resourceGrn"`
}

func NewGetCumulativeRequestFromDict(data map[string]interface{}) GetCumulativeRequest {
	return GetCumulativeRequest{
		Name:        core.CastString(data["name"]),
		ResourceGrn: core.CastString(data["resourceGrn"]),
	}
}

func (p GetCumulativeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"resourceGrn": p.ResourceGrn,
	}
}

func (p GetCumulativeRequest) Pointer() *GetCumulativeRequest {
	return &p
}

type DescribeBillingActivitiesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	Year               *int32  `json:"year"`
	Month              *int32  `json:"month"`
	Service            *string `json:"service"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeBillingActivitiesRequestFromDict(data map[string]interface{}) DescribeBillingActivitiesRequest {
	return DescribeBillingActivitiesRequest{
		Year:      core.CastInt32(data["year"]),
		Month:     core.CastInt32(data["month"]),
		Service:   core.CastString(data["service"]),
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeBillingActivitiesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"year":      p.Year,
		"month":     p.Month,
		"service":   p.Service,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeBillingActivitiesRequest) Pointer() *DescribeBillingActivitiesRequest {
	return &p
}

type GetBillingActivityRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	Year               *int32  `json:"year"`
	Month              *int32  `json:"month"`
	Service            *string `json:"service"`
	ActivityType       *string `json:"activityType"`
}

func NewGetBillingActivityRequestFromDict(data map[string]interface{}) GetBillingActivityRequest {
	return GetBillingActivityRequest{
		Year:         core.CastInt32(data["year"]),
		Month:        core.CastInt32(data["month"]),
		Service:      core.CastString(data["service"]),
		ActivityType: core.CastString(data["activityType"]),
	}
}

func (p GetBillingActivityRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"year":         p.Year,
		"month":        p.Month,
		"service":      p.Service,
		"activityType": p.ActivityType,
	}
}

func (p GetBillingActivityRequest) Pointer() *GetBillingActivityRequest {
	return &p
}
