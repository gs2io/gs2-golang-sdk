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

type GetChartResult struct {
    /** チャート */
	Item         *Chart	`json:"item"`
}

func (p *GetChartResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetChartAsyncResult struct {
	result *GetChartResult
	err    error
}

type GetCumulativeResult struct {
    /** 累積値 */
	Item         *Cumulative	`json:"item"`
}

func (p *GetCumulativeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetCumulativeAsyncResult struct {
	result *GetCumulativeResult
	err    error
}

type DescribeBillingActivitiesResult struct {
    /** 請求にまつわるアクティビティのリスト */
	Items         *[]*BillingActivity	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeBillingActivitiesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeBillingActivitiesAsyncResult struct {
	result *DescribeBillingActivitiesResult
	err    error
}

type GetBillingActivityResult struct {
    /** 請求にまつわるアクティビティ */
	Item         *BillingActivity	`json:"item"`
}

func (p *GetBillingActivityResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetBillingActivityAsyncResult struct {
	result *GetBillingActivityResult
	err    error
}
