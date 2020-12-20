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

type Chart struct {
    /** Datadog のJSON 形式のグラフ定義 */
	ChartId *core.String   `json:"chartId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** Datadog から払い出された組み込みID */
	EmbedId *core.String   `json:"embedId"`
    /** Datadog から払い出された組み込み用HTML */
	Html *core.String   `json:"html"`
}

func (p *Chart) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["chartId"] = p.ChartId
    data["ownerId"] = p.OwnerId
    data["embedId"] = p.EmbedId
    data["html"] = p.Html
    return &data
}

type Cumulative struct {
    /** 累積値 */
	CumulativeId *core.String   `json:"cumulativeId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** リソースのGRN */
	ResourceGrn *core.String   `json:"resourceGrn"`
    /** 名前 */
	Name *core.String   `json:"name"`
    /** 累積値 */
	Value *int64   `json:"value"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Cumulative) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["cumulativeId"] = p.CumulativeId
    data["ownerId"] = p.OwnerId
    data["resourceGrn"] = p.ResourceGrn
    data["name"] = p.Name
    data["value"] = p.Value
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type BillingActivity struct {
    /** 請求にまつわるアクティビティ */
	BillingActivityId *core.String   `json:"billingActivityId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** イベントの発生年 */
	Year *int32   `json:"year"`
    /** イベントの発生月 */
	Month *int32   `json:"month"`
    /** サービスの種類 */
	Service *core.String   `json:"service"`
    /** イベントの種類 */
	ActivityType *core.String   `json:"activityType"`
    /** イベントの値 */
	Value *int64   `json:"value"`
}

func (p *BillingActivity) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["billingActivityId"] = p.BillingActivityId
    data["ownerId"] = p.OwnerId
    data["year"] = p.Year
    data["month"] = p.Month
    data["service"] = p.Service
    data["activityType"] = p.ActivityType
    data["value"] = p.Value
    return &data
}
