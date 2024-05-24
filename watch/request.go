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
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type GetChartRequest struct {
	SourceRequestId *string   `json:"sourceRequestId"`
	RequestId       *string   `json:"requestId"`
	ContextStack    *string   `json:"contextStack"`
	Measure         *string   `json:"measure"`
	Grn             *string   `json:"grn"`
	Round           *string   `json:"round"`
	Filters         []Filter  `json:"filters"`
	GroupBys        []*string `json:"groupBys"`
	CountBy         *string   `json:"countBy"`
	Begin           *int64    `json:"begin"`
	End             *int64    `json:"end"`
	PageToken       *string   `json:"pageToken"`
	Limit           *int32    `json:"limit"`
}

func (p *GetChartRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetChartRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetChartRequest{}
	} else {
		*p = GetChartRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["measure"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Measure = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Measure = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Measure = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Measure = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Measure = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Measure)
				}
			}
		}
		if v, ok := d["grn"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Grn = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Grn = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Grn = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Grn = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Grn = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Grn)
				}
			}
		}
		if v, ok := d["round"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Round = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Round = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Round = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Round = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Round = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Round)
				}
			}
		}
		if v, ok := d["filters"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Filters)
		}
		if v, ok := d["groupBys"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.GroupBys = l
			}
		}
		if v, ok := d["countBy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CountBy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CountBy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CountBy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CountBy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CountBy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CountBy)
				}
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Begin)
		}
		if v, ok := d["end"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.End)
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewGetChartRequestFromJson(data string) (GetChartRequest, error) {
	req := GetChartRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetChartRequest{}, err
	}
	return req, nil
}

func NewGetChartRequestFromDict(data map[string]interface{}) GetChartRequest {
	return GetChartRequest{
		Measure:   core.CastString(data["measure"]),
		Grn:       core.CastString(data["grn"]),
		Round:     core.CastString(data["round"]),
		Filters:   CastFilters(core.CastArray(data["filters"])),
		GroupBys:  core.CastStrings(core.CastArray(data["groupBys"])),
		CountBy:   core.CastString(data["countBy"]),
		Begin:     core.CastInt64(data["begin"]),
		End:       core.CastInt64(data["end"]),
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p GetChartRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"measure": p.Measure,
		"grn":     p.Grn,
		"round":   p.Round,
		"filters": CastFiltersFromDict(
			p.Filters,
		),
		"groupBys": core.CastStringsFromDict(
			p.GroupBys,
		),
		"countBy":   p.CountBy,
		"begin":     p.Begin,
		"end":       p.End,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p GetChartRequest) Pointer() *GetChartRequest {
	return &p
}

type GetDistributionRequest struct {
	SourceRequestId *string   `json:"sourceRequestId"`
	RequestId       *string   `json:"requestId"`
	ContextStack    *string   `json:"contextStack"`
	Measure         *string   `json:"measure"`
	Grn             *string   `json:"grn"`
	Filters         []Filter  `json:"filters"`
	GroupBys        []*string `json:"groupBys"`
	Begin           *int64    `json:"begin"`
	End             *int64    `json:"end"`
	PageToken       *string   `json:"pageToken"`
	Limit           *int32    `json:"limit"`
}

func (p *GetDistributionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetDistributionRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetDistributionRequest{}
	} else {
		*p = GetDistributionRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["measure"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Measure = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Measure = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Measure = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Measure = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Measure = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Measure)
				}
			}
		}
		if v, ok := d["grn"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Grn = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Grn = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Grn = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Grn = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Grn = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Grn)
				}
			}
		}
		if v, ok := d["filters"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Filters)
		}
		if v, ok := d["groupBys"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.GroupBys = l
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Begin)
		}
		if v, ok := d["end"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.End)
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewGetDistributionRequestFromJson(data string) (GetDistributionRequest, error) {
	req := GetDistributionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetDistributionRequest{}, err
	}
	return req, nil
}

func NewGetDistributionRequestFromDict(data map[string]interface{}) GetDistributionRequest {
	return GetDistributionRequest{
		Measure:   core.CastString(data["measure"]),
		Grn:       core.CastString(data["grn"]),
		Filters:   CastFilters(core.CastArray(data["filters"])),
		GroupBys:  core.CastStrings(core.CastArray(data["groupBys"])),
		Begin:     core.CastInt64(data["begin"]),
		End:       core.CastInt64(data["end"]),
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p GetDistributionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"measure": p.Measure,
		"grn":     p.Grn,
		"filters": CastFiltersFromDict(
			p.Filters,
		),
		"groupBys": core.CastStringsFromDict(
			p.GroupBys,
		),
		"begin":     p.Begin,
		"end":       p.End,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p GetDistributionRequest) Pointer() *GetDistributionRequest {
	return &p
}

type GetCumulativeRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Name            *string `json:"name"`
	ResourceGrn     *string `json:"resourceGrn"`
}

func (p *GetCumulativeRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCumulativeRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetCumulativeRequest{}
	} else {
		*p = GetCumulativeRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["resourceGrn"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResourceGrn = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResourceGrn = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResourceGrn = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResourceGrn = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResourceGrn = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResourceGrn)
				}
			}
		}
	}
	return nil
}

func NewGetCumulativeRequestFromJson(data string) (GetCumulativeRequest, error) {
	req := GetCumulativeRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCumulativeRequest{}, err
	}
	return req, nil
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Year            *int32  `json:"year"`
	Month           *int32  `json:"month"`
	Service         *string `json:"service"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeBillingActivitiesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeBillingActivitiesRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeBillingActivitiesRequest{}
	} else {
		*p = DescribeBillingActivitiesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["year"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Year)
		}
		if v, ok := d["month"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Month)
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeBillingActivitiesRequestFromJson(data string) (DescribeBillingActivitiesRequest, error) {
	req := DescribeBillingActivitiesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeBillingActivitiesRequest{}, err
	}
	return req, nil
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Year            *int32  `json:"year"`
	Month           *int32  `json:"month"`
	Service         *string `json:"service"`
	ActivityType    *string `json:"activityType"`
}

func (p *GetBillingActivityRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBillingActivityRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetBillingActivityRequest{}
	} else {
		*p = GetBillingActivityRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["year"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Year)
		}
		if v, ok := d["month"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Month)
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["activityType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ActivityType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ActivityType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ActivityType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ActivityType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ActivityType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ActivityType)
				}
			}
		}
	}
	return nil
}

func NewGetBillingActivityRequestFromJson(data string) (GetBillingActivityRequest, error) {
	req := GetBillingActivityRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBillingActivityRequest{}, err
	}
	return req, nil
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

type GetGeneralMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *GetGeneralMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGeneralMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGeneralMetricsRequest{}
	} else {
		*p = GetGeneralMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewGetGeneralMetricsRequestFromJson(data string) (GetGeneralMetricsRequest, error) {
	req := GetGeneralMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGeneralMetricsRequest{}, err
	}
	return req, nil
}

func NewGetGeneralMetricsRequestFromDict(data map[string]interface{}) GetGeneralMetricsRequest {
	return GetGeneralMetricsRequest{}
}

func (p GetGeneralMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p GetGeneralMetricsRequest) Pointer() *GetGeneralMetricsRequest {
	return &p
}

type DescribeAccountNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeAccountNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeAccountNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeAccountNamespaceMetricsRequest{}
	} else {
		*p = DescribeAccountNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeAccountNamespaceMetricsRequestFromJson(data string) (DescribeAccountNamespaceMetricsRequest, error) {
	req := DescribeAccountNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeAccountNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeAccountNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeAccountNamespaceMetricsRequest {
	return DescribeAccountNamespaceMetricsRequest{}
}

func (p DescribeAccountNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeAccountNamespaceMetricsRequest) Pointer() *DescribeAccountNamespaceMetricsRequest {
	return &p
}

type GetAccountNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetAccountNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetAccountNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetAccountNamespaceMetricsRequest{}
	} else {
		*p = GetAccountNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetAccountNamespaceMetricsRequestFromJson(data string) (GetAccountNamespaceMetricsRequest, error) {
	req := GetAccountNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetAccountNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetAccountNamespaceMetricsRequestFromDict(data map[string]interface{}) GetAccountNamespaceMetricsRequest {
	return GetAccountNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetAccountNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetAccountNamespaceMetricsRequest) Pointer() *GetAccountNamespaceMetricsRequest {
	return &p
}

type DescribeChatNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeChatNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeChatNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeChatNamespaceMetricsRequest{}
	} else {
		*p = DescribeChatNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeChatNamespaceMetricsRequestFromJson(data string) (DescribeChatNamespaceMetricsRequest, error) {
	req := DescribeChatNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeChatNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeChatNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeChatNamespaceMetricsRequest {
	return DescribeChatNamespaceMetricsRequest{}
}

func (p DescribeChatNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeChatNamespaceMetricsRequest) Pointer() *DescribeChatNamespaceMetricsRequest {
	return &p
}

type GetChatNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetChatNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetChatNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetChatNamespaceMetricsRequest{}
	} else {
		*p = GetChatNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetChatNamespaceMetricsRequestFromJson(data string) (GetChatNamespaceMetricsRequest, error) {
	req := GetChatNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetChatNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetChatNamespaceMetricsRequestFromDict(data map[string]interface{}) GetChatNamespaceMetricsRequest {
	return GetChatNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetChatNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetChatNamespaceMetricsRequest) Pointer() *GetChatNamespaceMetricsRequest {
	return &p
}

type DescribeDatastoreNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeDatastoreNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeDatastoreNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeDatastoreNamespaceMetricsRequest{}
	} else {
		*p = DescribeDatastoreNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeDatastoreNamespaceMetricsRequestFromJson(data string) (DescribeDatastoreNamespaceMetricsRequest, error) {
	req := DescribeDatastoreNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeDatastoreNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeDatastoreNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeDatastoreNamespaceMetricsRequest {
	return DescribeDatastoreNamespaceMetricsRequest{}
}

func (p DescribeDatastoreNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeDatastoreNamespaceMetricsRequest) Pointer() *DescribeDatastoreNamespaceMetricsRequest {
	return &p
}

type GetDatastoreNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetDatastoreNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetDatastoreNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetDatastoreNamespaceMetricsRequest{}
	} else {
		*p = GetDatastoreNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetDatastoreNamespaceMetricsRequestFromJson(data string) (GetDatastoreNamespaceMetricsRequest, error) {
	req := GetDatastoreNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetDatastoreNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetDatastoreNamespaceMetricsRequestFromDict(data map[string]interface{}) GetDatastoreNamespaceMetricsRequest {
	return GetDatastoreNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetDatastoreNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetDatastoreNamespaceMetricsRequest) Pointer() *GetDatastoreNamespaceMetricsRequest {
	return &p
}

type DescribeDictionaryNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeDictionaryNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeDictionaryNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeDictionaryNamespaceMetricsRequest{}
	} else {
		*p = DescribeDictionaryNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeDictionaryNamespaceMetricsRequestFromJson(data string) (DescribeDictionaryNamespaceMetricsRequest, error) {
	req := DescribeDictionaryNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeDictionaryNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeDictionaryNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeDictionaryNamespaceMetricsRequest {
	return DescribeDictionaryNamespaceMetricsRequest{}
}

func (p DescribeDictionaryNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeDictionaryNamespaceMetricsRequest) Pointer() *DescribeDictionaryNamespaceMetricsRequest {
	return &p
}

type GetDictionaryNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetDictionaryNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetDictionaryNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetDictionaryNamespaceMetricsRequest{}
	} else {
		*p = GetDictionaryNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetDictionaryNamespaceMetricsRequestFromJson(data string) (GetDictionaryNamespaceMetricsRequest, error) {
	req := GetDictionaryNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetDictionaryNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetDictionaryNamespaceMetricsRequestFromDict(data map[string]interface{}) GetDictionaryNamespaceMetricsRequest {
	return GetDictionaryNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetDictionaryNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetDictionaryNamespaceMetricsRequest) Pointer() *GetDictionaryNamespaceMetricsRequest {
	return &p
}

type DescribeExchangeRateModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeExchangeRateModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeExchangeRateModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeExchangeRateModelMetricsRequest{}
	} else {
		*p = DescribeExchangeRateModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeExchangeRateModelMetricsRequestFromJson(data string) (DescribeExchangeRateModelMetricsRequest, error) {
	req := DescribeExchangeRateModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeExchangeRateModelMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeExchangeRateModelMetricsRequestFromDict(data map[string]interface{}) DescribeExchangeRateModelMetricsRequest {
	return DescribeExchangeRateModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeExchangeRateModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeExchangeRateModelMetricsRequest) Pointer() *DescribeExchangeRateModelMetricsRequest {
	return &p
}

type GetExchangeRateModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RateName        *string `json:"rateName"`
}

func (p *GetExchangeRateModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetExchangeRateModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetExchangeRateModelMetricsRequest{}
	} else {
		*p = GetExchangeRateModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["rateName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RateName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RateName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RateName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RateName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RateName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RateName)
				}
			}
		}
	}
	return nil
}

func NewGetExchangeRateModelMetricsRequestFromJson(data string) (GetExchangeRateModelMetricsRequest, error) {
	req := GetExchangeRateModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetExchangeRateModelMetricsRequest{}, err
	}
	return req, nil
}

func NewGetExchangeRateModelMetricsRequestFromDict(data map[string]interface{}) GetExchangeRateModelMetricsRequest {
	return GetExchangeRateModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RateName:      core.CastString(data["rateName"]),
	}
}

func (p GetExchangeRateModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rateName":      p.RateName,
	}
}

func (p GetExchangeRateModelMetricsRequest) Pointer() *GetExchangeRateModelMetricsRequest {
	return &p
}

type DescribeExchangeNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeExchangeNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeExchangeNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeExchangeNamespaceMetricsRequest{}
	} else {
		*p = DescribeExchangeNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeExchangeNamespaceMetricsRequestFromJson(data string) (DescribeExchangeNamespaceMetricsRequest, error) {
	req := DescribeExchangeNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeExchangeNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeExchangeNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeExchangeNamespaceMetricsRequest {
	return DescribeExchangeNamespaceMetricsRequest{}
}

func (p DescribeExchangeNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeExchangeNamespaceMetricsRequest) Pointer() *DescribeExchangeNamespaceMetricsRequest {
	return &p
}

type GetExchangeNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetExchangeNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetExchangeNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetExchangeNamespaceMetricsRequest{}
	} else {
		*p = GetExchangeNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetExchangeNamespaceMetricsRequestFromJson(data string) (GetExchangeNamespaceMetricsRequest, error) {
	req := GetExchangeNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetExchangeNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetExchangeNamespaceMetricsRequestFromDict(data map[string]interface{}) GetExchangeNamespaceMetricsRequest {
	return GetExchangeNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetExchangeNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetExchangeNamespaceMetricsRequest) Pointer() *GetExchangeNamespaceMetricsRequest {
	return &p
}

type DescribeExperienceStatusMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ExperienceName  *string `json:"experienceName"`
}

func (p *DescribeExperienceStatusMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeExperienceStatusMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeExperienceStatusMetricsRequest{}
	} else {
		*p = DescribeExperienceStatusMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeExperienceStatusMetricsRequestFromJson(data string) (DescribeExperienceStatusMetricsRequest, error) {
	req := DescribeExperienceStatusMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeExperienceStatusMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeExperienceStatusMetricsRequestFromDict(data map[string]interface{}) DescribeExperienceStatusMetricsRequest {
	return DescribeExperienceStatusMetricsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		ExperienceName: core.CastString(data["experienceName"]),
	}
}

func (p DescribeExperienceStatusMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
	}
}

func (p DescribeExperienceStatusMetricsRequest) Pointer() *DescribeExperienceStatusMetricsRequest {
	return &p
}

type DescribeExperienceExperienceModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeExperienceExperienceModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeExperienceExperienceModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeExperienceExperienceModelMetricsRequest{}
	} else {
		*p = DescribeExperienceExperienceModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeExperienceExperienceModelMetricsRequestFromJson(data string) (DescribeExperienceExperienceModelMetricsRequest, error) {
	req := DescribeExperienceExperienceModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeExperienceExperienceModelMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeExperienceExperienceModelMetricsRequestFromDict(data map[string]interface{}) DescribeExperienceExperienceModelMetricsRequest {
	return DescribeExperienceExperienceModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeExperienceExperienceModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeExperienceExperienceModelMetricsRequest) Pointer() *DescribeExperienceExperienceModelMetricsRequest {
	return &p
}

type GetExperienceExperienceModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ExperienceName  *string `json:"experienceName"`
}

func (p *GetExperienceExperienceModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetExperienceExperienceModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetExperienceExperienceModelMetricsRequest{}
	} else {
		*p = GetExperienceExperienceModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
	}
	return nil
}

func NewGetExperienceExperienceModelMetricsRequestFromJson(data string) (GetExperienceExperienceModelMetricsRequest, error) {
	req := GetExperienceExperienceModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetExperienceExperienceModelMetricsRequest{}, err
	}
	return req, nil
}

func NewGetExperienceExperienceModelMetricsRequestFromDict(data map[string]interface{}) GetExperienceExperienceModelMetricsRequest {
	return GetExperienceExperienceModelMetricsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		ExperienceName: core.CastString(data["experienceName"]),
	}
}

func (p GetExperienceExperienceModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
	}
}

func (p GetExperienceExperienceModelMetricsRequest) Pointer() *GetExperienceExperienceModelMetricsRequest {
	return &p
}

type DescribeExperienceNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeExperienceNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeExperienceNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeExperienceNamespaceMetricsRequest{}
	} else {
		*p = DescribeExperienceNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeExperienceNamespaceMetricsRequestFromJson(data string) (DescribeExperienceNamespaceMetricsRequest, error) {
	req := DescribeExperienceNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeExperienceNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeExperienceNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeExperienceNamespaceMetricsRequest {
	return DescribeExperienceNamespaceMetricsRequest{}
}

func (p DescribeExperienceNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeExperienceNamespaceMetricsRequest) Pointer() *DescribeExperienceNamespaceMetricsRequest {
	return &p
}

type GetExperienceNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetExperienceNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetExperienceNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetExperienceNamespaceMetricsRequest{}
	} else {
		*p = GetExperienceNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetExperienceNamespaceMetricsRequestFromJson(data string) (GetExperienceNamespaceMetricsRequest, error) {
	req := GetExperienceNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetExperienceNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetExperienceNamespaceMetricsRequestFromDict(data map[string]interface{}) GetExperienceNamespaceMetricsRequest {
	return GetExperienceNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetExperienceNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetExperienceNamespaceMetricsRequest) Pointer() *GetExperienceNamespaceMetricsRequest {
	return &p
}

type DescribeFormationFormMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	MoldModelName   *string `json:"moldModelName"`
}

func (p *DescribeFormationFormMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeFormationFormMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeFormationFormMetricsRequest{}
	} else {
		*p = DescribeFormationFormMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["moldModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MoldModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MoldModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MoldModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MoldModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MoldModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MoldModelName)
				}
			}
		}
	}
	return nil
}

func NewDescribeFormationFormMetricsRequestFromJson(data string) (DescribeFormationFormMetricsRequest, error) {
	req := DescribeFormationFormMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeFormationFormMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeFormationFormMetricsRequestFromDict(data map[string]interface{}) DescribeFormationFormMetricsRequest {
	return DescribeFormationFormMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		MoldModelName: core.CastString(data["moldModelName"]),
	}
}

func (p DescribeFormationFormMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"moldModelName": p.MoldModelName,
	}
}

func (p DescribeFormationFormMetricsRequest) Pointer() *DescribeFormationFormMetricsRequest {
	return &p
}

type DescribeFormationMoldMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeFormationMoldMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeFormationMoldMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeFormationMoldMetricsRequest{}
	} else {
		*p = DescribeFormationMoldMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeFormationMoldMetricsRequestFromJson(data string) (DescribeFormationMoldMetricsRequest, error) {
	req := DescribeFormationMoldMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeFormationMoldMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeFormationMoldMetricsRequestFromDict(data map[string]interface{}) DescribeFormationMoldMetricsRequest {
	return DescribeFormationMoldMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeFormationMoldMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeFormationMoldMetricsRequest) Pointer() *DescribeFormationMoldMetricsRequest {
	return &p
}

type DescribeFormationNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeFormationNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeFormationNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeFormationNamespaceMetricsRequest{}
	} else {
		*p = DescribeFormationNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeFormationNamespaceMetricsRequestFromJson(data string) (DescribeFormationNamespaceMetricsRequest, error) {
	req := DescribeFormationNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeFormationNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeFormationNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeFormationNamespaceMetricsRequest {
	return DescribeFormationNamespaceMetricsRequest{}
}

func (p DescribeFormationNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeFormationNamespaceMetricsRequest) Pointer() *DescribeFormationNamespaceMetricsRequest {
	return &p
}

type GetFormationNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetFormationNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetFormationNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetFormationNamespaceMetricsRequest{}
	} else {
		*p = GetFormationNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetFormationNamespaceMetricsRequestFromJson(data string) (GetFormationNamespaceMetricsRequest, error) {
	req := GetFormationNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetFormationNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetFormationNamespaceMetricsRequestFromDict(data map[string]interface{}) GetFormationNamespaceMetricsRequest {
	return GetFormationNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetFormationNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetFormationNamespaceMetricsRequest) Pointer() *GetFormationNamespaceMetricsRequest {
	return &p
}

type DescribeFriendNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeFriendNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeFriendNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeFriendNamespaceMetricsRequest{}
	} else {
		*p = DescribeFriendNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeFriendNamespaceMetricsRequestFromJson(data string) (DescribeFriendNamespaceMetricsRequest, error) {
	req := DescribeFriendNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeFriendNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeFriendNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeFriendNamespaceMetricsRequest {
	return DescribeFriendNamespaceMetricsRequest{}
}

func (p DescribeFriendNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeFriendNamespaceMetricsRequest) Pointer() *DescribeFriendNamespaceMetricsRequest {
	return &p
}

type GetFriendNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetFriendNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetFriendNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetFriendNamespaceMetricsRequest{}
	} else {
		*p = GetFriendNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetFriendNamespaceMetricsRequestFromJson(data string) (GetFriendNamespaceMetricsRequest, error) {
	req := GetFriendNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetFriendNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetFriendNamespaceMetricsRequestFromDict(data map[string]interface{}) GetFriendNamespaceMetricsRequest {
	return GetFriendNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetFriendNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetFriendNamespaceMetricsRequest) Pointer() *GetFriendNamespaceMetricsRequest {
	return &p
}

type DescribeInboxNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeInboxNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeInboxNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeInboxNamespaceMetricsRequest{}
	} else {
		*p = DescribeInboxNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeInboxNamespaceMetricsRequestFromJson(data string) (DescribeInboxNamespaceMetricsRequest, error) {
	req := DescribeInboxNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeInboxNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeInboxNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeInboxNamespaceMetricsRequest {
	return DescribeInboxNamespaceMetricsRequest{}
}

func (p DescribeInboxNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeInboxNamespaceMetricsRequest) Pointer() *DescribeInboxNamespaceMetricsRequest {
	return &p
}

type GetInboxNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetInboxNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetInboxNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetInboxNamespaceMetricsRequest{}
	} else {
		*p = GetInboxNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetInboxNamespaceMetricsRequestFromJson(data string) (GetInboxNamespaceMetricsRequest, error) {
	req := GetInboxNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetInboxNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetInboxNamespaceMetricsRequestFromDict(data map[string]interface{}) GetInboxNamespaceMetricsRequest {
	return GetInboxNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetInboxNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetInboxNamespaceMetricsRequest) Pointer() *GetInboxNamespaceMetricsRequest {
	return &p
}

type DescribeInventoryItemSetMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *DescribeInventoryItemSetMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeInventoryItemSetMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeInventoryItemSetMetricsRequest{}
	} else {
		*p = DescribeInventoryItemSetMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewDescribeInventoryItemSetMetricsRequestFromJson(data string) (DescribeInventoryItemSetMetricsRequest, error) {
	req := DescribeInventoryItemSetMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeInventoryItemSetMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeInventoryItemSetMetricsRequestFromDict(data map[string]interface{}) DescribeInventoryItemSetMetricsRequest {
	return DescribeInventoryItemSetMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DescribeInventoryItemSetMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DescribeInventoryItemSetMetricsRequest) Pointer() *DescribeInventoryItemSetMetricsRequest {
	return &p
}

type DescribeInventoryInventoryMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeInventoryInventoryMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeInventoryInventoryMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeInventoryInventoryMetricsRequest{}
	} else {
		*p = DescribeInventoryInventoryMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeInventoryInventoryMetricsRequestFromJson(data string) (DescribeInventoryInventoryMetricsRequest, error) {
	req := DescribeInventoryInventoryMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeInventoryInventoryMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeInventoryInventoryMetricsRequestFromDict(data map[string]interface{}) DescribeInventoryInventoryMetricsRequest {
	return DescribeInventoryInventoryMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeInventoryInventoryMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeInventoryInventoryMetricsRequest) Pointer() *DescribeInventoryInventoryMetricsRequest {
	return &p
}

type DescribeInventoryNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeInventoryNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeInventoryNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeInventoryNamespaceMetricsRequest{}
	} else {
		*p = DescribeInventoryNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeInventoryNamespaceMetricsRequestFromJson(data string) (DescribeInventoryNamespaceMetricsRequest, error) {
	req := DescribeInventoryNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeInventoryNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeInventoryNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeInventoryNamespaceMetricsRequest {
	return DescribeInventoryNamespaceMetricsRequest{}
}

func (p DescribeInventoryNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeInventoryNamespaceMetricsRequest) Pointer() *DescribeInventoryNamespaceMetricsRequest {
	return &p
}

type GetInventoryNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetInventoryNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetInventoryNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetInventoryNamespaceMetricsRequest{}
	} else {
		*p = GetInventoryNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetInventoryNamespaceMetricsRequestFromJson(data string) (GetInventoryNamespaceMetricsRequest, error) {
	req := GetInventoryNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetInventoryNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetInventoryNamespaceMetricsRequestFromDict(data map[string]interface{}) GetInventoryNamespaceMetricsRequest {
	return GetInventoryNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetInventoryNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetInventoryNamespaceMetricsRequest) Pointer() *GetInventoryNamespaceMetricsRequest {
	return &p
}

type DescribeKeyNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeKeyNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeKeyNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeKeyNamespaceMetricsRequest{}
	} else {
		*p = DescribeKeyNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeKeyNamespaceMetricsRequestFromJson(data string) (DescribeKeyNamespaceMetricsRequest, error) {
	req := DescribeKeyNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeKeyNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeKeyNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeKeyNamespaceMetricsRequest {
	return DescribeKeyNamespaceMetricsRequest{}
}

func (p DescribeKeyNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeKeyNamespaceMetricsRequest) Pointer() *DescribeKeyNamespaceMetricsRequest {
	return &p
}

type GetKeyNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetKeyNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetKeyNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetKeyNamespaceMetricsRequest{}
	} else {
		*p = GetKeyNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetKeyNamespaceMetricsRequestFromJson(data string) (GetKeyNamespaceMetricsRequest, error) {
	req := GetKeyNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetKeyNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetKeyNamespaceMetricsRequestFromDict(data map[string]interface{}) GetKeyNamespaceMetricsRequest {
	return GetKeyNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetKeyNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetKeyNamespaceMetricsRequest) Pointer() *GetKeyNamespaceMetricsRequest {
	return &p
}

type DescribeLimitCounterMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	LimitName       *string `json:"limitName"`
}

func (p *DescribeLimitCounterMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeLimitCounterMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeLimitCounterMetricsRequest{}
	} else {
		*p = DescribeLimitCounterMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["limitName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LimitName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LimitName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LimitName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LimitName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LimitName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LimitName)
				}
			}
		}
	}
	return nil
}

func NewDescribeLimitCounterMetricsRequestFromJson(data string) (DescribeLimitCounterMetricsRequest, error) {
	req := DescribeLimitCounterMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeLimitCounterMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeLimitCounterMetricsRequestFromDict(data map[string]interface{}) DescribeLimitCounterMetricsRequest {
	return DescribeLimitCounterMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
	}
}

func (p DescribeLimitCounterMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
	}
}

func (p DescribeLimitCounterMetricsRequest) Pointer() *DescribeLimitCounterMetricsRequest {
	return &p
}

type DescribeLimitLimitModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeLimitLimitModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeLimitLimitModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeLimitLimitModelMetricsRequest{}
	} else {
		*p = DescribeLimitLimitModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeLimitLimitModelMetricsRequestFromJson(data string) (DescribeLimitLimitModelMetricsRequest, error) {
	req := DescribeLimitLimitModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeLimitLimitModelMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeLimitLimitModelMetricsRequestFromDict(data map[string]interface{}) DescribeLimitLimitModelMetricsRequest {
	return DescribeLimitLimitModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeLimitLimitModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeLimitLimitModelMetricsRequest) Pointer() *DescribeLimitLimitModelMetricsRequest {
	return &p
}

type GetLimitLimitModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	LimitName       *string `json:"limitName"`
}

func (p *GetLimitLimitModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetLimitLimitModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetLimitLimitModelMetricsRequest{}
	} else {
		*p = GetLimitLimitModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["limitName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LimitName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LimitName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LimitName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LimitName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LimitName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LimitName)
				}
			}
		}
	}
	return nil
}

func NewGetLimitLimitModelMetricsRequestFromJson(data string) (GetLimitLimitModelMetricsRequest, error) {
	req := GetLimitLimitModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetLimitLimitModelMetricsRequest{}, err
	}
	return req, nil
}

func NewGetLimitLimitModelMetricsRequestFromDict(data map[string]interface{}) GetLimitLimitModelMetricsRequest {
	return GetLimitLimitModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LimitName:     core.CastString(data["limitName"]),
	}
}

func (p GetLimitLimitModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"limitName":     p.LimitName,
	}
}

func (p GetLimitLimitModelMetricsRequest) Pointer() *GetLimitLimitModelMetricsRequest {
	return &p
}

type DescribeLimitNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeLimitNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeLimitNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeLimitNamespaceMetricsRequest{}
	} else {
		*p = DescribeLimitNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeLimitNamespaceMetricsRequestFromJson(data string) (DescribeLimitNamespaceMetricsRequest, error) {
	req := DescribeLimitNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeLimitNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeLimitNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeLimitNamespaceMetricsRequest {
	return DescribeLimitNamespaceMetricsRequest{}
}

func (p DescribeLimitNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeLimitNamespaceMetricsRequest) Pointer() *DescribeLimitNamespaceMetricsRequest {
	return &p
}

type GetLimitNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetLimitNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetLimitNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetLimitNamespaceMetricsRequest{}
	} else {
		*p = GetLimitNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetLimitNamespaceMetricsRequestFromJson(data string) (GetLimitNamespaceMetricsRequest, error) {
	req := GetLimitNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetLimitNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetLimitNamespaceMetricsRequestFromDict(data map[string]interface{}) GetLimitNamespaceMetricsRequest {
	return GetLimitNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetLimitNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetLimitNamespaceMetricsRequest) Pointer() *GetLimitNamespaceMetricsRequest {
	return &p
}

type DescribeLotteryLotteryMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeLotteryLotteryMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeLotteryLotteryMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeLotteryLotteryMetricsRequest{}
	} else {
		*p = DescribeLotteryLotteryMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeLotteryLotteryMetricsRequestFromJson(data string) (DescribeLotteryLotteryMetricsRequest, error) {
	req := DescribeLotteryLotteryMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeLotteryLotteryMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeLotteryLotteryMetricsRequestFromDict(data map[string]interface{}) DescribeLotteryLotteryMetricsRequest {
	return DescribeLotteryLotteryMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeLotteryLotteryMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeLotteryLotteryMetricsRequest) Pointer() *DescribeLotteryLotteryMetricsRequest {
	return &p
}

type GetLotteryLotteryMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	LotteryName     *string `json:"lotteryName"`
}

func (p *GetLotteryLotteryMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetLotteryLotteryMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetLotteryLotteryMetricsRequest{}
	} else {
		*p = GetLotteryLotteryMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["lotteryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LotteryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LotteryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LotteryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LotteryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LotteryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LotteryName)
				}
			}
		}
	}
	return nil
}

func NewGetLotteryLotteryMetricsRequestFromJson(data string) (GetLotteryLotteryMetricsRequest, error) {
	req := GetLotteryLotteryMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetLotteryLotteryMetricsRequest{}, err
	}
	return req, nil
}

func NewGetLotteryLotteryMetricsRequestFromDict(data map[string]interface{}) GetLotteryLotteryMetricsRequest {
	return GetLotteryLotteryMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		LotteryName:   core.CastString(data["lotteryName"]),
	}
}

func (p GetLotteryLotteryMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"lotteryName":   p.LotteryName,
	}
}

func (p GetLotteryLotteryMetricsRequest) Pointer() *GetLotteryLotteryMetricsRequest {
	return &p
}

type DescribeLotteryNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeLotteryNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeLotteryNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeLotteryNamespaceMetricsRequest{}
	} else {
		*p = DescribeLotteryNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeLotteryNamespaceMetricsRequestFromJson(data string) (DescribeLotteryNamespaceMetricsRequest, error) {
	req := DescribeLotteryNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeLotteryNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeLotteryNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeLotteryNamespaceMetricsRequest {
	return DescribeLotteryNamespaceMetricsRequest{}
}

func (p DescribeLotteryNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeLotteryNamespaceMetricsRequest) Pointer() *DescribeLotteryNamespaceMetricsRequest {
	return &p
}

type GetLotteryNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetLotteryNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetLotteryNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetLotteryNamespaceMetricsRequest{}
	} else {
		*p = GetLotteryNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetLotteryNamespaceMetricsRequestFromJson(data string) (GetLotteryNamespaceMetricsRequest, error) {
	req := GetLotteryNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetLotteryNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetLotteryNamespaceMetricsRequestFromDict(data map[string]interface{}) GetLotteryNamespaceMetricsRequest {
	return GetLotteryNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetLotteryNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetLotteryNamespaceMetricsRequest) Pointer() *GetLotteryNamespaceMetricsRequest {
	return &p
}

type DescribeMatchmakingNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeMatchmakingNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeMatchmakingNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeMatchmakingNamespaceMetricsRequest{}
	} else {
		*p = DescribeMatchmakingNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeMatchmakingNamespaceMetricsRequestFromJson(data string) (DescribeMatchmakingNamespaceMetricsRequest, error) {
	req := DescribeMatchmakingNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeMatchmakingNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeMatchmakingNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeMatchmakingNamespaceMetricsRequest {
	return DescribeMatchmakingNamespaceMetricsRequest{}
}

func (p DescribeMatchmakingNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeMatchmakingNamespaceMetricsRequest) Pointer() *DescribeMatchmakingNamespaceMetricsRequest {
	return &p
}

type GetMatchmakingNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetMatchmakingNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetMatchmakingNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetMatchmakingNamespaceMetricsRequest{}
	} else {
		*p = GetMatchmakingNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetMatchmakingNamespaceMetricsRequestFromJson(data string) (GetMatchmakingNamespaceMetricsRequest, error) {
	req := GetMatchmakingNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetMatchmakingNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetMatchmakingNamespaceMetricsRequestFromDict(data map[string]interface{}) GetMatchmakingNamespaceMetricsRequest {
	return GetMatchmakingNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetMatchmakingNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetMatchmakingNamespaceMetricsRequest) Pointer() *GetMatchmakingNamespaceMetricsRequest {
	return &p
}

type DescribeMissionCounterMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeMissionCounterMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeMissionCounterMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeMissionCounterMetricsRequest{}
	} else {
		*p = DescribeMissionCounterMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeMissionCounterMetricsRequestFromJson(data string) (DescribeMissionCounterMetricsRequest, error) {
	req := DescribeMissionCounterMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeMissionCounterMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeMissionCounterMetricsRequestFromDict(data map[string]interface{}) DescribeMissionCounterMetricsRequest {
	return DescribeMissionCounterMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeMissionCounterMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeMissionCounterMetricsRequest) Pointer() *DescribeMissionCounterMetricsRequest {
	return &p
}

type DescribeMissionMissionGroupModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeMissionMissionGroupModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeMissionMissionGroupModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeMissionMissionGroupModelMetricsRequest{}
	} else {
		*p = DescribeMissionMissionGroupModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeMissionMissionGroupModelMetricsRequestFromJson(data string) (DescribeMissionMissionGroupModelMetricsRequest, error) {
	req := DescribeMissionMissionGroupModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeMissionMissionGroupModelMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeMissionMissionGroupModelMetricsRequestFromDict(data map[string]interface{}) DescribeMissionMissionGroupModelMetricsRequest {
	return DescribeMissionMissionGroupModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeMissionMissionGroupModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeMissionMissionGroupModelMetricsRequest) Pointer() *DescribeMissionMissionGroupModelMetricsRequest {
	return &p
}

type GetMissionMissionGroupModelMetricsRequest struct {
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
}

func (p *GetMissionMissionGroupModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetMissionMissionGroupModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetMissionMissionGroupModelMetricsRequest{}
	} else {
		*p = GetMissionMissionGroupModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["missionGroupName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MissionGroupName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MissionGroupName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MissionGroupName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MissionGroupName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MissionGroupName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MissionGroupName)
				}
			}
		}
	}
	return nil
}

func NewGetMissionMissionGroupModelMetricsRequestFromJson(data string) (GetMissionMissionGroupModelMetricsRequest, error) {
	req := GetMissionMissionGroupModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetMissionMissionGroupModelMetricsRequest{}, err
	}
	return req, nil
}

func NewGetMissionMissionGroupModelMetricsRequestFromDict(data map[string]interface{}) GetMissionMissionGroupModelMetricsRequest {
	return GetMissionMissionGroupModelMetricsRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
	}
}

func (p GetMissionMissionGroupModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
	}
}

func (p GetMissionMissionGroupModelMetricsRequest) Pointer() *GetMissionMissionGroupModelMetricsRequest {
	return &p
}

type DescribeMissionNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeMissionNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeMissionNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeMissionNamespaceMetricsRequest{}
	} else {
		*p = DescribeMissionNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeMissionNamespaceMetricsRequestFromJson(data string) (DescribeMissionNamespaceMetricsRequest, error) {
	req := DescribeMissionNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeMissionNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeMissionNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeMissionNamespaceMetricsRequest {
	return DescribeMissionNamespaceMetricsRequest{}
}

func (p DescribeMissionNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeMissionNamespaceMetricsRequest) Pointer() *DescribeMissionNamespaceMetricsRequest {
	return &p
}

type GetMissionNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetMissionNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetMissionNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetMissionNamespaceMetricsRequest{}
	} else {
		*p = GetMissionNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetMissionNamespaceMetricsRequestFromJson(data string) (GetMissionNamespaceMetricsRequest, error) {
	req := GetMissionNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetMissionNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetMissionNamespaceMetricsRequestFromDict(data map[string]interface{}) GetMissionNamespaceMetricsRequest {
	return GetMissionNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetMissionNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetMissionNamespaceMetricsRequest) Pointer() *GetMissionNamespaceMetricsRequest {
	return &p
}

type DescribeMoneyWalletMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeMoneyWalletMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeMoneyWalletMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeMoneyWalletMetricsRequest{}
	} else {
		*p = DescribeMoneyWalletMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeMoneyWalletMetricsRequestFromJson(data string) (DescribeMoneyWalletMetricsRequest, error) {
	req := DescribeMoneyWalletMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeMoneyWalletMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeMoneyWalletMetricsRequestFromDict(data map[string]interface{}) DescribeMoneyWalletMetricsRequest {
	return DescribeMoneyWalletMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeMoneyWalletMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeMoneyWalletMetricsRequest) Pointer() *DescribeMoneyWalletMetricsRequest {
	return &p
}

type DescribeMoneyReceiptMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeMoneyReceiptMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeMoneyReceiptMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeMoneyReceiptMetricsRequest{}
	} else {
		*p = DescribeMoneyReceiptMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeMoneyReceiptMetricsRequestFromJson(data string) (DescribeMoneyReceiptMetricsRequest, error) {
	req := DescribeMoneyReceiptMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeMoneyReceiptMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeMoneyReceiptMetricsRequestFromDict(data map[string]interface{}) DescribeMoneyReceiptMetricsRequest {
	return DescribeMoneyReceiptMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeMoneyReceiptMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeMoneyReceiptMetricsRequest) Pointer() *DescribeMoneyReceiptMetricsRequest {
	return &p
}

type DescribeMoneyNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeMoneyNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeMoneyNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeMoneyNamespaceMetricsRequest{}
	} else {
		*p = DescribeMoneyNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeMoneyNamespaceMetricsRequestFromJson(data string) (DescribeMoneyNamespaceMetricsRequest, error) {
	req := DescribeMoneyNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeMoneyNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeMoneyNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeMoneyNamespaceMetricsRequest {
	return DescribeMoneyNamespaceMetricsRequest{}
}

func (p DescribeMoneyNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeMoneyNamespaceMetricsRequest) Pointer() *DescribeMoneyNamespaceMetricsRequest {
	return &p
}

type GetMoneyNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetMoneyNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetMoneyNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetMoneyNamespaceMetricsRequest{}
	} else {
		*p = GetMoneyNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetMoneyNamespaceMetricsRequestFromJson(data string) (GetMoneyNamespaceMetricsRequest, error) {
	req := GetMoneyNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetMoneyNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetMoneyNamespaceMetricsRequestFromDict(data map[string]interface{}) GetMoneyNamespaceMetricsRequest {
	return GetMoneyNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetMoneyNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetMoneyNamespaceMetricsRequest) Pointer() *GetMoneyNamespaceMetricsRequest {
	return &p
}

type DescribeQuestQuestModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	QuestGroupName  *string `json:"questGroupName"`
}

func (p *DescribeQuestQuestModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeQuestQuestModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeQuestQuestModelMetricsRequest{}
	} else {
		*p = DescribeQuestQuestModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["questGroupName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestGroupName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestGroupName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestGroupName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestGroupName)
				}
			}
		}
	}
	return nil
}

func NewDescribeQuestQuestModelMetricsRequestFromJson(data string) (DescribeQuestQuestModelMetricsRequest, error) {
	req := DescribeQuestQuestModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeQuestQuestModelMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeQuestQuestModelMetricsRequestFromDict(data map[string]interface{}) DescribeQuestQuestModelMetricsRequest {
	return DescribeQuestQuestModelMetricsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		QuestGroupName: core.CastString(data["questGroupName"]),
	}
}

func (p DescribeQuestQuestModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"questGroupName": p.QuestGroupName,
	}
}

func (p DescribeQuestQuestModelMetricsRequest) Pointer() *DescribeQuestQuestModelMetricsRequest {
	return &p
}

type GetQuestQuestModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	QuestGroupName  *string `json:"questGroupName"`
	QuestName       *string `json:"questName"`
}

func (p *GetQuestQuestModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetQuestQuestModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetQuestQuestModelMetricsRequest{}
	} else {
		*p = GetQuestQuestModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["questGroupName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestGroupName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestGroupName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestGroupName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestGroupName)
				}
			}
		}
		if v, ok := d["questName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestName)
				}
			}
		}
	}
	return nil
}

func NewGetQuestQuestModelMetricsRequestFromJson(data string) (GetQuestQuestModelMetricsRequest, error) {
	req := GetQuestQuestModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetQuestQuestModelMetricsRequest{}, err
	}
	return req, nil
}

func NewGetQuestQuestModelMetricsRequestFromDict(data map[string]interface{}) GetQuestQuestModelMetricsRequest {
	return GetQuestQuestModelMetricsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		QuestGroupName: core.CastString(data["questGroupName"]),
		QuestName:      core.CastString(data["questName"]),
	}
}

func (p GetQuestQuestModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"questGroupName": p.QuestGroupName,
		"questName":      p.QuestName,
	}
}

func (p GetQuestQuestModelMetricsRequest) Pointer() *GetQuestQuestModelMetricsRequest {
	return &p
}

type DescribeQuestQuestGroupModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeQuestQuestGroupModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeQuestQuestGroupModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeQuestQuestGroupModelMetricsRequest{}
	} else {
		*p = DescribeQuestQuestGroupModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeQuestQuestGroupModelMetricsRequestFromJson(data string) (DescribeQuestQuestGroupModelMetricsRequest, error) {
	req := DescribeQuestQuestGroupModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeQuestQuestGroupModelMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeQuestQuestGroupModelMetricsRequestFromDict(data map[string]interface{}) DescribeQuestQuestGroupModelMetricsRequest {
	return DescribeQuestQuestGroupModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeQuestQuestGroupModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeQuestQuestGroupModelMetricsRequest) Pointer() *DescribeQuestQuestGroupModelMetricsRequest {
	return &p
}

type GetQuestQuestGroupModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	QuestGroupName  *string `json:"questGroupName"`
}

func (p *GetQuestQuestGroupModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetQuestQuestGroupModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetQuestQuestGroupModelMetricsRequest{}
	} else {
		*p = GetQuestQuestGroupModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["questGroupName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestGroupName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestGroupName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestGroupName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestGroupName)
				}
			}
		}
	}
	return nil
}

func NewGetQuestQuestGroupModelMetricsRequestFromJson(data string) (GetQuestQuestGroupModelMetricsRequest, error) {
	req := GetQuestQuestGroupModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetQuestQuestGroupModelMetricsRequest{}, err
	}
	return req, nil
}

func NewGetQuestQuestGroupModelMetricsRequestFromDict(data map[string]interface{}) GetQuestQuestGroupModelMetricsRequest {
	return GetQuestQuestGroupModelMetricsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		QuestGroupName: core.CastString(data["questGroupName"]),
	}
}

func (p GetQuestQuestGroupModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"questGroupName": p.QuestGroupName,
	}
}

func (p GetQuestQuestGroupModelMetricsRequest) Pointer() *GetQuestQuestGroupModelMetricsRequest {
	return &p
}

type DescribeQuestNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeQuestNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeQuestNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeQuestNamespaceMetricsRequest{}
	} else {
		*p = DescribeQuestNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeQuestNamespaceMetricsRequestFromJson(data string) (DescribeQuestNamespaceMetricsRequest, error) {
	req := DescribeQuestNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeQuestNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeQuestNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeQuestNamespaceMetricsRequest {
	return DescribeQuestNamespaceMetricsRequest{}
}

func (p DescribeQuestNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeQuestNamespaceMetricsRequest) Pointer() *DescribeQuestNamespaceMetricsRequest {
	return &p
}

type GetQuestNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetQuestNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetQuestNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetQuestNamespaceMetricsRequest{}
	} else {
		*p = GetQuestNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetQuestNamespaceMetricsRequestFromJson(data string) (GetQuestNamespaceMetricsRequest, error) {
	req := GetQuestNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetQuestNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetQuestNamespaceMetricsRequestFromDict(data map[string]interface{}) GetQuestNamespaceMetricsRequest {
	return GetQuestNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetQuestNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetQuestNamespaceMetricsRequest) Pointer() *GetQuestNamespaceMetricsRequest {
	return &p
}

type DescribeRankingCategoryModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeRankingCategoryModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRankingCategoryModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeRankingCategoryModelMetricsRequest{}
	} else {
		*p = DescribeRankingCategoryModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeRankingCategoryModelMetricsRequestFromJson(data string) (DescribeRankingCategoryModelMetricsRequest, error) {
	req := DescribeRankingCategoryModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRankingCategoryModelMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeRankingCategoryModelMetricsRequestFromDict(data map[string]interface{}) DescribeRankingCategoryModelMetricsRequest {
	return DescribeRankingCategoryModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeRankingCategoryModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeRankingCategoryModelMetricsRequest) Pointer() *DescribeRankingCategoryModelMetricsRequest {
	return &p
}

type GetRankingCategoryModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CategoryName    *string `json:"categoryName"`
}

func (p *GetRankingCategoryModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetRankingCategoryModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetRankingCategoryModelMetricsRequest{}
	} else {
		*p = GetRankingCategoryModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["categoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CategoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CategoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CategoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CategoryName)
				}
			}
		}
	}
	return nil
}

func NewGetRankingCategoryModelMetricsRequestFromJson(data string) (GetRankingCategoryModelMetricsRequest, error) {
	req := GetRankingCategoryModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetRankingCategoryModelMetricsRequest{}, err
	}
	return req, nil
}

func NewGetRankingCategoryModelMetricsRequestFromDict(data map[string]interface{}) GetRankingCategoryModelMetricsRequest {
	return GetRankingCategoryModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CategoryName:  core.CastString(data["categoryName"]),
	}
}

func (p GetRankingCategoryModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"categoryName":  p.CategoryName,
	}
}

func (p GetRankingCategoryModelMetricsRequest) Pointer() *GetRankingCategoryModelMetricsRequest {
	return &p
}

type DescribeRankingNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeRankingNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRankingNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeRankingNamespaceMetricsRequest{}
	} else {
		*p = DescribeRankingNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeRankingNamespaceMetricsRequestFromJson(data string) (DescribeRankingNamespaceMetricsRequest, error) {
	req := DescribeRankingNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRankingNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeRankingNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeRankingNamespaceMetricsRequest {
	return DescribeRankingNamespaceMetricsRequest{}
}

func (p DescribeRankingNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeRankingNamespaceMetricsRequest) Pointer() *DescribeRankingNamespaceMetricsRequest {
	return &p
}

type GetRankingNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetRankingNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetRankingNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetRankingNamespaceMetricsRequest{}
	} else {
		*p = GetRankingNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetRankingNamespaceMetricsRequestFromJson(data string) (GetRankingNamespaceMetricsRequest, error) {
	req := GetRankingNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetRankingNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetRankingNamespaceMetricsRequestFromDict(data map[string]interface{}) GetRankingNamespaceMetricsRequest {
	return GetRankingNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetRankingNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetRankingNamespaceMetricsRequest) Pointer() *GetRankingNamespaceMetricsRequest {
	return &p
}

type DescribeShowcaseDisplayItemMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
}

func (p *DescribeShowcaseDisplayItemMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeShowcaseDisplayItemMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeShowcaseDisplayItemMetricsRequest{}
	} else {
		*p = DescribeShowcaseDisplayItemMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["showcaseName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ShowcaseName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ShowcaseName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ShowcaseName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ShowcaseName)
				}
			}
		}
	}
	return nil
}

func NewDescribeShowcaseDisplayItemMetricsRequestFromJson(data string) (DescribeShowcaseDisplayItemMetricsRequest, error) {
	req := DescribeShowcaseDisplayItemMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeShowcaseDisplayItemMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeShowcaseDisplayItemMetricsRequestFromDict(data map[string]interface{}) DescribeShowcaseDisplayItemMetricsRequest {
	return DescribeShowcaseDisplayItemMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
	}
}

func (p DescribeShowcaseDisplayItemMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
	}
}

func (p DescribeShowcaseDisplayItemMetricsRequest) Pointer() *DescribeShowcaseDisplayItemMetricsRequest {
	return &p
}

type GetShowcaseDisplayItemMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
	DisplayItemId   *string `json:"displayItemId"`
}

func (p *GetShowcaseDisplayItemMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetShowcaseDisplayItemMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetShowcaseDisplayItemMetricsRequest{}
	} else {
		*p = GetShowcaseDisplayItemMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["showcaseName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ShowcaseName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ShowcaseName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ShowcaseName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ShowcaseName)
				}
			}
		}
		if v, ok := d["displayItemId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayItemId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayItemId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayItemId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayItemId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayItemId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayItemId)
				}
			}
		}
	}
	return nil
}

func NewGetShowcaseDisplayItemMetricsRequestFromJson(data string) (GetShowcaseDisplayItemMetricsRequest, error) {
	req := GetShowcaseDisplayItemMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetShowcaseDisplayItemMetricsRequest{}, err
	}
	return req, nil
}

func NewGetShowcaseDisplayItemMetricsRequestFromDict(data map[string]interface{}) GetShowcaseDisplayItemMetricsRequest {
	return GetShowcaseDisplayItemMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
		DisplayItemId: core.CastString(data["displayItemId"]),
	}
}

func (p GetShowcaseDisplayItemMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
		"displayItemId": p.DisplayItemId,
	}
}

func (p GetShowcaseDisplayItemMetricsRequest) Pointer() *GetShowcaseDisplayItemMetricsRequest {
	return &p
}

type DescribeShowcaseShowcaseMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeShowcaseShowcaseMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeShowcaseShowcaseMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeShowcaseShowcaseMetricsRequest{}
	} else {
		*p = DescribeShowcaseShowcaseMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeShowcaseShowcaseMetricsRequestFromJson(data string) (DescribeShowcaseShowcaseMetricsRequest, error) {
	req := DescribeShowcaseShowcaseMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeShowcaseShowcaseMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeShowcaseShowcaseMetricsRequestFromDict(data map[string]interface{}) DescribeShowcaseShowcaseMetricsRequest {
	return DescribeShowcaseShowcaseMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeShowcaseShowcaseMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeShowcaseShowcaseMetricsRequest) Pointer() *DescribeShowcaseShowcaseMetricsRequest {
	return &p
}

type GetShowcaseShowcaseMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ShowcaseName    *string `json:"showcaseName"`
}

func (p *GetShowcaseShowcaseMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetShowcaseShowcaseMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetShowcaseShowcaseMetricsRequest{}
	} else {
		*p = GetShowcaseShowcaseMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["showcaseName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ShowcaseName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ShowcaseName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ShowcaseName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ShowcaseName)
				}
			}
		}
	}
	return nil
}

func NewGetShowcaseShowcaseMetricsRequestFromJson(data string) (GetShowcaseShowcaseMetricsRequest, error) {
	req := GetShowcaseShowcaseMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetShowcaseShowcaseMetricsRequest{}, err
	}
	return req, nil
}

func NewGetShowcaseShowcaseMetricsRequestFromDict(data map[string]interface{}) GetShowcaseShowcaseMetricsRequest {
	return GetShowcaseShowcaseMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
	}
}

func (p GetShowcaseShowcaseMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"showcaseName":  p.ShowcaseName,
	}
}

func (p GetShowcaseShowcaseMetricsRequest) Pointer() *GetShowcaseShowcaseMetricsRequest {
	return &p
}

type DescribeShowcaseNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeShowcaseNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeShowcaseNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeShowcaseNamespaceMetricsRequest{}
	} else {
		*p = DescribeShowcaseNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeShowcaseNamespaceMetricsRequestFromJson(data string) (DescribeShowcaseNamespaceMetricsRequest, error) {
	req := DescribeShowcaseNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeShowcaseNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeShowcaseNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeShowcaseNamespaceMetricsRequest {
	return DescribeShowcaseNamespaceMetricsRequest{}
}

func (p DescribeShowcaseNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeShowcaseNamespaceMetricsRequest) Pointer() *DescribeShowcaseNamespaceMetricsRequest {
	return &p
}

type GetShowcaseNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetShowcaseNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetShowcaseNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetShowcaseNamespaceMetricsRequest{}
	} else {
		*p = GetShowcaseNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetShowcaseNamespaceMetricsRequestFromJson(data string) (GetShowcaseNamespaceMetricsRequest, error) {
	req := GetShowcaseNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetShowcaseNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetShowcaseNamespaceMetricsRequestFromDict(data map[string]interface{}) GetShowcaseNamespaceMetricsRequest {
	return GetShowcaseNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetShowcaseNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetShowcaseNamespaceMetricsRequest) Pointer() *GetShowcaseNamespaceMetricsRequest {
	return &p
}

type DescribeStaminaStaminaModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeStaminaStaminaModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStaminaStaminaModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeStaminaStaminaModelMetricsRequest{}
	} else {
		*p = DescribeStaminaStaminaModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeStaminaStaminaModelMetricsRequestFromJson(data string) (DescribeStaminaStaminaModelMetricsRequest, error) {
	req := DescribeStaminaStaminaModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStaminaStaminaModelMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeStaminaStaminaModelMetricsRequestFromDict(data map[string]interface{}) DescribeStaminaStaminaModelMetricsRequest {
	return DescribeStaminaStaminaModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeStaminaStaminaModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeStaminaStaminaModelMetricsRequest) Pointer() *DescribeStaminaStaminaModelMetricsRequest {
	return &p
}

type GetStaminaStaminaModelMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	StaminaName     *string `json:"staminaName"`
}

func (p *GetStaminaStaminaModelMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStaminaStaminaModelMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStaminaStaminaModelMetricsRequest{}
	} else {
		*p = GetStaminaStaminaModelMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StaminaName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StaminaName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StaminaName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StaminaName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StaminaName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StaminaName)
				}
			}
		}
	}
	return nil
}

func NewGetStaminaStaminaModelMetricsRequestFromJson(data string) (GetStaminaStaminaModelMetricsRequest, error) {
	req := GetStaminaStaminaModelMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStaminaStaminaModelMetricsRequest{}, err
	}
	return req, nil
}

func NewGetStaminaStaminaModelMetricsRequestFromDict(data map[string]interface{}) GetStaminaStaminaModelMetricsRequest {
	return GetStaminaStaminaModelMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
	}
}

func (p GetStaminaStaminaModelMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
	}
}

func (p GetStaminaStaminaModelMetricsRequest) Pointer() *GetStaminaStaminaModelMetricsRequest {
	return &p
}

type DescribeStaminaNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
}

func (p *DescribeStaminaNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStaminaNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeStaminaNamespaceMetricsRequest{}
	} else {
		*p = DescribeStaminaNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeStaminaNamespaceMetricsRequestFromJson(data string) (DescribeStaminaNamespaceMetricsRequest, error) {
	req := DescribeStaminaNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStaminaNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewDescribeStaminaNamespaceMetricsRequestFromDict(data map[string]interface{}) DescribeStaminaNamespaceMetricsRequest {
	return DescribeStaminaNamespaceMetricsRequest{}
}

func (p DescribeStaminaNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeStaminaNamespaceMetricsRequest) Pointer() *DescribeStaminaNamespaceMetricsRequest {
	return &p
}

type GetStaminaNamespaceMetricsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetStaminaNamespaceMetricsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStaminaNamespaceMetricsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStaminaNamespaceMetricsRequest{}
	} else {
		*p = GetStaminaNamespaceMetricsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetStaminaNamespaceMetricsRequestFromJson(data string) (GetStaminaNamespaceMetricsRequest, error) {
	req := GetStaminaNamespaceMetricsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStaminaNamespaceMetricsRequest{}, err
	}
	return req, nil
}

func NewGetStaminaNamespaceMetricsRequestFromDict(data map[string]interface{}) GetStaminaNamespaceMetricsRequest {
	return GetStaminaNamespaceMetricsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetStaminaNamespaceMetricsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetStaminaNamespaceMetricsRequest) Pointer() *GetStaminaNamespaceMetricsRequest {
	return &p
}
