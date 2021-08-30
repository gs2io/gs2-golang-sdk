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

import "github.com/gs2io/gs2-golang-sdk/core"

type GetChartResult struct {
	Item *Chart `json:"item"`
}

type GetChartAsyncResult struct {
	result *GetChartResult
	err    error
}

func NewGetChartResultFromDict(data map[string]interface{}) GetChartResult {
	return GetChartResult{
		Item: NewChartFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetChartResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetChartResult) Pointer() *GetChartResult {
	return &p
}

type GetCumulativeResult struct {
	Item *Cumulative `json:"item"`
}

type GetCumulativeAsyncResult struct {
	result *GetCumulativeResult
	err    error
}

func NewGetCumulativeResultFromDict(data map[string]interface{}) GetCumulativeResult {
	return GetCumulativeResult{
		Item: NewCumulativeFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCumulativeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCumulativeResult) Pointer() *GetCumulativeResult {
	return &p
}

type DescribeBillingActivitiesResult struct {
	Items         []BillingActivity `json:"items"`
	NextPageToken *string           `json:"nextPageToken"`
}

type DescribeBillingActivitiesAsyncResult struct {
	result *DescribeBillingActivitiesResult
	err    error
}

func NewDescribeBillingActivitiesResultFromDict(data map[string]interface{}) DescribeBillingActivitiesResult {
	return DescribeBillingActivitiesResult{
		Items:         CastBillingActivities(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeBillingActivitiesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBillingActivitiesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeBillingActivitiesResult) Pointer() *DescribeBillingActivitiesResult {
	return &p
}

type GetBillingActivityResult struct {
	Item *BillingActivity `json:"item"`
}

type GetBillingActivityAsyncResult struct {
	result *GetBillingActivityResult
	err    error
}

func NewGetBillingActivityResultFromDict(data map[string]interface{}) GetBillingActivityResult {
	return GetBillingActivityResult{
		Item: NewBillingActivityFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetBillingActivityResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetBillingActivityResult) Pointer() *GetBillingActivityResult {
	return &p
}
