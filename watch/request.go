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
	Metrics *core.String	`json:"metrics"`
	Grn *core.String	`json:"grn"`
	Queries *[]core.String	`json:"queries"`
	By *core.String	`json:"by"`
	Timeframe *core.String	`json:"timeframe"`
	Size *core.String	`json:"size"`
	Format *core.String	`json:"format"`
	Aggregator *core.String	`json:"aggregator"`
	Style *core.String	`json:"style"`
	Title *core.String	`json:"title"`
}

type GetCumulativeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *core.String	`json:"name"`
	ResourceGrn *core.String	`json:"resourceGrn"`
}

type DescribeBillingActivitiesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Year *int32	`json:"year"`
	Month *int32	`json:"month"`
	Service *core.String	`json:"service"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type GetBillingActivityRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Year *int32	`json:"year"`
	Month *int32	`json:"month"`
	Service *core.String	`json:"service"`
	ActivityType *core.String	`json:"activityType"`
}
