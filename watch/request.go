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
	"github.com/gs2io/gs2-golang-sdk/core"
)

type GetChartRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Metrics *string	`json:"metrics"`
	Grn *string	`json:"grn"`
	Queries []string	`json:"queries"`
	By *string	`json:"by"`
	Timeframe *string	`json:"timeframe"`
	Size *string	`json:"size"`
	Format *string	`json:"format"`
	Aggregator *string	`json:"aggregator"`
	Style *string	`json:"style"`
	Title *string	`json:"title"`
}

type GetCumulativeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *string	`json:"name"`
	ResourceGrn *string	`json:"resourceGrn"`
}

type DescribeBillingActivitiesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Year *int32	`json:"year"`
	Month *int32	`json:"month"`
	Service *string	`json:"service"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type GetBillingActivityRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Year *int32	`json:"year"`
	Month *int32	`json:"month"`
	Service *string	`json:"service"`
	ActivityType *string	`json:"activityType"`
}
