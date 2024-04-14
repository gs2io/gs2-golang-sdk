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
	Timestamp *int64    `json:"timestamp"`
	Value     *int64    `json:"value"`
	GroupBys  []*string `json:"groupBys"`
}

func NewChartFromJson(data string) Chart {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChartFromDict(dict)
}

func NewChartFromDict(data map[string]interface{}) Chart {
	return Chart{
		Timestamp: core.CastInt64(data["timestamp"]),
		Value:     core.CastInt64(data["value"]),
		GroupBys:  core.CastStrings(core.CastArray(data["groupBys"])),
	}
}

func (p Chart) ToDict() map[string]interface{} {

	var timestamp *int64
	if p.Timestamp != nil {
		timestamp = p.Timestamp
	}
	var value *int64
	if p.Value != nil {
		value = p.Value
	}
	var groupBys []interface{}
	if p.GroupBys != nil {
		groupBys = core.CastStringsFromDict(
			p.GroupBys,
		)
	}
	return map[string]interface{}{
		"timestamp": timestamp,
		"value":     value,
		"groupBys":  groupBys,
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

type Distribution struct {
	Value    *int64    `json:"value"`
	Count    *int64    `json:"count"`
	GroupBys []*string `json:"groupBys"`
}

func NewDistributionFromJson(data string) Distribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDistributionFromDict(dict)
}

func NewDistributionFromDict(data map[string]interface{}) Distribution {
	return Distribution{
		Value:    core.CastInt64(data["value"]),
		Count:    core.CastInt64(data["count"]),
		GroupBys: core.CastStrings(core.CastArray(data["groupBys"])),
	}
}

func (p Distribution) ToDict() map[string]interface{} {

	var value *int64
	if p.Value != nil {
		value = p.Value
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var groupBys []interface{}
	if p.GroupBys != nil {
		groupBys = core.CastStringsFromDict(
			p.GroupBys,
		)
	}
	return map[string]interface{}{
		"value":    value,
		"count":    count,
		"groupBys": groupBys,
	}
}

func (p Distribution) Pointer() *Distribution {
	return &p
}

func CastDistributions(data []interface{}) []Distribution {
	v := make([]Distribution, 0)
	for _, d := range data {
		v = append(v, NewDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDistributionsFromDict(data []Distribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Cumulative struct {
	CumulativeId *string `json:"cumulativeId"`
	ResourceGrn  *string `json:"resourceGrn"`
	Name         *string `json:"name"`
	Value        *int64  `json:"value"`
	UpdatedAt    *int64  `json:"updatedAt"`
	Revision     *int64  `json:"revision"`
}

func NewCumulativeFromJson(data string) Cumulative {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCumulativeFromDict(dict)
}

func NewCumulativeFromDict(data map[string]interface{}) Cumulative {
	return Cumulative{
		CumulativeId: core.CastString(data["cumulativeId"]),
		ResourceGrn:  core.CastString(data["resourceGrn"]),
		Name:         core.CastString(data["name"]),
		Value:        core.CastInt64(data["value"]),
		UpdatedAt:    core.CastInt64(data["updatedAt"]),
		Revision:     core.CastInt64(data["revision"]),
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"cumulativeId": cumulativeId,
		"resourceGrn":  resourceGrn,
		"name":         name,
		"value":        value,
		"updatedAt":    updatedAt,
		"revision":     revision,
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
	Year              *int32  `json:"year"`
	Month             *int32  `json:"month"`
	Service           *string `json:"service"`
	ActivityType      *string `json:"activityType"`
	Value             *int64  `json:"value"`
	Revision          *int64  `json:"revision"`
}

func NewBillingActivityFromJson(data string) BillingActivity {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBillingActivityFromDict(dict)
}

func NewBillingActivityFromDict(data map[string]interface{}) BillingActivity {
	return BillingActivity{
		BillingActivityId: core.CastString(data["billingActivityId"]),
		Year:              core.CastInt32(data["year"]),
		Month:             core.CastInt32(data["month"]),
		Service:           core.CastString(data["service"]),
		ActivityType:      core.CastString(data["activityType"]),
		Value:             core.CastInt64(data["value"]),
		Revision:          core.CastInt64(data["revision"]),
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"billingActivityId": billingActivityId,
		"year":              year,
		"month":             month,
		"service":           service,
		"activityType":      activityType,
		"value":             value,
		"revision":          revision,
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
	Grn        *string   `json:"grn"`
	Service    *string   `json:"service"`
	Method     *string   `json:"method"`
	Metric     *string   `json:"metric"`
	Cumulative *bool     `json:"cumulative"`
	Value      *float64  `json:"value"`
	Tags       []*string `json:"tags"`
	CallAt     *int64    `json:"callAt"`
}

func NewStatsEventFromJson(data string) StatsEvent {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStatsEventFromDict(dict)
}

func NewStatsEventFromDict(data map[string]interface{}) StatsEvent {
	return StatsEvent{
		Grn:        core.CastString(data["grn"]),
		Service:    core.CastString(data["service"]),
		Method:     core.CastString(data["method"]),
		Metric:     core.CastString(data["metric"]),
		Cumulative: core.CastBool(data["cumulative"]),
		Value:      core.CastFloat64(data["value"]),
		Tags:       core.CastStrings(core.CastArray(data["tags"])),
		CallAt:     core.CastInt64(data["callAt"]),
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
	return map[string]interface{}{
		"grn":        grn,
		"service":    service,
		"method":     method,
		"metric":     metric,
		"cumulative": cumulative,
		"value":      value,
		"tags":       tags,
		"callAt":     callAt,
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

type Filter struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

func NewFilterFromJson(data string) Filter {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFilterFromDict(dict)
}

func NewFilterFromDict(data map[string]interface{}) Filter {
	return Filter{
		Key:   core.CastString(data["key"]),
		Value: core.CastString(data["value"]),
	}
}

func (p Filter) ToDict() map[string]interface{} {

	var key *string
	if p.Key != nil {
		key = p.Key
	}
	var value *string
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"key":   key,
		"value": value,
	}
}

func (p Filter) Pointer() *Filter {
	return &p
}

func CastFilters(data []interface{}) []Filter {
	v := make([]Filter, 0)
	for _, d := range data {
		v = append(v, NewFilterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFiltersFromDict(data []Filter) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GeneralDauWauMauHistory struct {
	Date                *string `json:"date"`
	Dau                 *int64  `json:"dau"`
	WauLast7Days        *int64  `json:"wauLast7Days"`
	WauTargetWeekSunday *int64  `json:"wauTargetWeekSunday"`
	WauTargetWeekMonday *int64  `json:"wauTargetWeekMonday"`
	MauLast30Days       *int64  `json:"mauLast30Days"`
	MauTargetMonth      *int64  `json:"mauTargetMonth"`
}

func NewGeneralDauWauMauHistoryFromJson(data string) GeneralDauWauMauHistory {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGeneralDauWauMauHistoryFromDict(dict)
}

func NewGeneralDauWauMauHistoryFromDict(data map[string]interface{}) GeneralDauWauMauHistory {
	return GeneralDauWauMauHistory{
		Date:                core.CastString(data["date"]),
		Dau:                 core.CastInt64(data["dau"]),
		WauLast7Days:        core.CastInt64(data["wauLast7Days"]),
		WauTargetWeekSunday: core.CastInt64(data["wauTargetWeekSunday"]),
		WauTargetWeekMonday: core.CastInt64(data["wauTargetWeekMonday"]),
		MauLast30Days:       core.CastInt64(data["mauLast30Days"]),
		MauTargetMonth:      core.CastInt64(data["mauTargetMonth"]),
	}
}

func (p GeneralDauWauMauHistory) ToDict() map[string]interface{} {

	var date *string
	if p.Date != nil {
		date = p.Date
	}
	var dau *int64
	if p.Dau != nil {
		dau = p.Dau
	}
	var wauLast7Days *int64
	if p.WauLast7Days != nil {
		wauLast7Days = p.WauLast7Days
	}
	var wauTargetWeekSunday *int64
	if p.WauTargetWeekSunday != nil {
		wauTargetWeekSunday = p.WauTargetWeekSunday
	}
	var wauTargetWeekMonday *int64
	if p.WauTargetWeekMonday != nil {
		wauTargetWeekMonday = p.WauTargetWeekMonday
	}
	var mauLast30Days *int64
	if p.MauLast30Days != nil {
		mauLast30Days = p.MauLast30Days
	}
	var mauTargetMonth *int64
	if p.MauTargetMonth != nil {
		mauTargetMonth = p.MauTargetMonth
	}
	return map[string]interface{}{
		"date":                date,
		"dau":                 dau,
		"wauLast7Days":        wauLast7Days,
		"wauTargetWeekSunday": wauTargetWeekSunday,
		"wauTargetWeekMonday": wauTargetWeekMonday,
		"mauLast30Days":       mauLast30Days,
		"mauTargetMonth":      mauTargetMonth,
	}
}

func (p GeneralDauWauMauHistory) Pointer() *GeneralDauWauMauHistory {
	return &p
}

func CastGeneralDauWauMauHistories(data []interface{}) []GeneralDauWauMauHistory {
	v := make([]GeneralDauWauMauHistory, 0)
	for _, d := range data {
		v = append(v, NewGeneralDauWauMauHistoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGeneralDauWauMauHistoriesFromDict(data []GeneralDauWauMauHistory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GeneralDauWauMauAverage struct {
	Dau *float32 `json:"dau"`
	Wau *float32 `json:"wau"`
	Mau *float32 `json:"mau"`
}

func NewGeneralDauWauMauAverageFromJson(data string) GeneralDauWauMauAverage {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGeneralDauWauMauAverageFromDict(dict)
}

func NewGeneralDauWauMauAverageFromDict(data map[string]interface{}) GeneralDauWauMauAverage {
	return GeneralDauWauMauAverage{
		Dau: core.CastFloat32(data["dau"]),
		Wau: core.CastFloat32(data["wau"]),
		Mau: core.CastFloat32(data["mau"]),
	}
}

func (p GeneralDauWauMauAverage) ToDict() map[string]interface{} {

	var dau *float32
	if p.Dau != nil {
		dau = p.Dau
	}
	var wau *float32
	if p.Wau != nil {
		wau = p.Wau
	}
	var mau *float32
	if p.Mau != nil {
		mau = p.Mau
	}
	return map[string]interface{}{
		"dau": dau,
		"wau": wau,
		"mau": mau,
	}
}

func (p GeneralDauWauMauAverage) Pointer() *GeneralDauWauMauAverage {
	return &p
}

func CastGeneralDauWauMauAverages(data []interface{}) []GeneralDauWauMauAverage {
	v := make([]GeneralDauWauMauAverage, 0)
	for _, d := range data {
		v = append(v, NewGeneralDauWauMauAverageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGeneralDauWauMauAveragesFromDict(data []GeneralDauWauMauAverage) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GeneralDauWauMau struct {
	History []GeneralDauWauMauHistory `json:"history"`
	Avg     *GeneralDauWauMauAverage  `json:"avg"`
}

func NewGeneralDauWauMauFromJson(data string) GeneralDauWauMau {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGeneralDauWauMauFromDict(dict)
}

func NewGeneralDauWauMauFromDict(data map[string]interface{}) GeneralDauWauMau {
	return GeneralDauWauMau{
		History: CastGeneralDauWauMauHistories(core.CastArray(data["history"])),
		Avg:     NewGeneralDauWauMauAverageFromDict(core.CastMap(data["avg"])).Pointer(),
	}
}

func (p GeneralDauWauMau) ToDict() map[string]interface{} {

	var history []interface{}
	if p.History != nil {
		history = CastGeneralDauWauMauHistoriesFromDict(
			p.History,
		)
	}
	var avg map[string]interface{}
	if p.Avg != nil {
		avg = p.Avg.ToDict()
	}
	return map[string]interface{}{
		"history": history,
		"avg":     avg,
	}
}

func (p GeneralDauWauMau) Pointer() *GeneralDauWauMau {
	return &p
}

func CastGeneralDauWauMaus(data []interface{}) []GeneralDauWauMau {
	v := make([]GeneralDauWauMau, 0)
	for _, d := range data {
		v = append(v, NewGeneralDauWauMauFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGeneralDauWauMausFromDict(data []GeneralDauWauMau) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FirstEngagementStatisticsLoginDays struct {
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFirstEngagementStatisticsLoginDaysFromJson(data string) FirstEngagementStatisticsLoginDays {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFirstEngagementStatisticsLoginDaysFromDict(dict)
}

func NewFirstEngagementStatisticsLoginDaysFromDict(data map[string]interface{}) FirstEngagementStatisticsLoginDays {
	return FirstEngagementStatisticsLoginDays{
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FirstEngagementStatisticsLoginDays) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FirstEngagementStatisticsLoginDays) Pointer() *FirstEngagementStatisticsLoginDays {
	return &p
}

func CastFirstEngagementStatisticsLoginDayses(data []interface{}) []FirstEngagementStatisticsLoginDays {
	v := make([]FirstEngagementStatisticsLoginDays, 0)
	for _, d := range data {
		v = append(v, NewFirstEngagementStatisticsLoginDaysFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFirstEngagementStatisticsLoginDaysesFromDict(data []FirstEngagementStatisticsLoginDays) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FirstEngagementStatistics struct {
	LoginDays *FirstEngagementStatisticsLoginDays `json:"loginDays"`
}

func NewFirstEngagementStatisticsFromJson(data string) FirstEngagementStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFirstEngagementStatisticsFromDict(dict)
}

func NewFirstEngagementStatisticsFromDict(data map[string]interface{}) FirstEngagementStatistics {
	return FirstEngagementStatistics{
		LoginDays: NewFirstEngagementStatisticsLoginDaysFromDict(core.CastMap(data["loginDays"])).Pointer(),
	}
}

func (p FirstEngagementStatistics) ToDict() map[string]interface{} {

	var loginDays map[string]interface{}
	if p.LoginDays != nil {
		loginDays = p.LoginDays.ToDict()
	}
	return map[string]interface{}{
		"loginDays": loginDays,
	}
}

func (p FirstEngagementStatistics) Pointer() *FirstEngagementStatistics {
	return &p
}

func CastFirstEngagementStatisticses(data []interface{}) []FirstEngagementStatistics {
	v := make([]FirstEngagementStatistics, 0)
	for _, d := range data {
		v = append(v, NewFirstEngagementStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFirstEngagementStatisticsesFromDict(data []FirstEngagementStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FirstEngagementDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewFirstEngagementDistributionSegmentFromJson(data string) FirstEngagementDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFirstEngagementDistributionSegmentFromDict(dict)
}

func NewFirstEngagementDistributionSegmentFromDict(data map[string]interface{}) FirstEngagementDistributionSegment {
	return FirstEngagementDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p FirstEngagementDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p FirstEngagementDistributionSegment) Pointer() *FirstEngagementDistributionSegment {
	return &p
}

func CastFirstEngagementDistributionSegments(data []interface{}) []FirstEngagementDistributionSegment {
	v := make([]FirstEngagementDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFirstEngagementDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFirstEngagementDistributionSegmentsFromDict(data []FirstEngagementDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FirstEngagement struct {
	Statistics   *FirstEngagementStatistics           `json:"statistics"`
	Distribution []FirstEngagementDistributionSegment `json:"distribution"`
}

func NewFirstEngagementFromJson(data string) FirstEngagement {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFirstEngagementFromDict(dict)
}

func NewFirstEngagementFromDict(data map[string]interface{}) FirstEngagement {
	return FirstEngagement{
		Statistics:   NewFirstEngagementStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFirstEngagementDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FirstEngagement) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFirstEngagementDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FirstEngagement) Pointer() *FirstEngagement {
	return &p
}

func CastFirstEngagements(data []interface{}) []FirstEngagement {
	v := make([]FirstEngagement, 0)
	for _, d := range data {
		v = append(v, NewFirstEngagementFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFirstEngagementsFromDict(data []FirstEngagement) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SessionDurationStatistics struct {
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewSessionDurationStatisticsFromJson(data string) SessionDurationStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSessionDurationStatisticsFromDict(dict)
}

func NewSessionDurationStatisticsFromDict(data map[string]interface{}) SessionDurationStatistics {
	return SessionDurationStatistics{
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p SessionDurationStatistics) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p SessionDurationStatistics) Pointer() *SessionDurationStatistics {
	return &p
}

func CastSessionDurationStatisticses(data []interface{}) []SessionDurationStatistics {
	v := make([]SessionDurationStatistics, 0)
	for _, d := range data {
		v = append(v, NewSessionDurationStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSessionDurationStatisticsesFromDict(data []SessionDurationStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SessionDurationDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewSessionDurationDistributionSegmentFromJson(data string) SessionDurationDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSessionDurationDistributionSegmentFromDict(dict)
}

func NewSessionDurationDistributionSegmentFromDict(data map[string]interface{}) SessionDurationDistributionSegment {
	return SessionDurationDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p SessionDurationDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p SessionDurationDistributionSegment) Pointer() *SessionDurationDistributionSegment {
	return &p
}

func CastSessionDurationDistributionSegments(data []interface{}) []SessionDurationDistributionSegment {
	v := make([]SessionDurationDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewSessionDurationDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSessionDurationDistributionSegmentsFromDict(data []SessionDurationDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type EngagementHistory struct {
	Date            *string  `json:"date"`
	NewUserRate     *float32 `json:"newUserRate"`
	ReturnUserRate  *float32 `json:"returnUserRate"`
	EngagedUserRate *float32 `json:"engagedUserRate"`
}

func NewEngagementHistoryFromJson(data string) EngagementHistory {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEngagementHistoryFromDict(dict)
}

func NewEngagementHistoryFromDict(data map[string]interface{}) EngagementHistory {
	return EngagementHistory{
		Date:            core.CastString(data["date"]),
		NewUserRate:     core.CastFloat32(data["newUserRate"]),
		ReturnUserRate:  core.CastFloat32(data["returnUserRate"]),
		EngagedUserRate: core.CastFloat32(data["engagedUserRate"]),
	}
}

func (p EngagementHistory) ToDict() map[string]interface{} {

	var date *string
	if p.Date != nil {
		date = p.Date
	}
	var newUserRate *float32
	if p.NewUserRate != nil {
		newUserRate = p.NewUserRate
	}
	var returnUserRate *float32
	if p.ReturnUserRate != nil {
		returnUserRate = p.ReturnUserRate
	}
	var engagedUserRate *float32
	if p.EngagedUserRate != nil {
		engagedUserRate = p.EngagedUserRate
	}
	return map[string]interface{}{
		"date":            date,
		"newUserRate":     newUserRate,
		"returnUserRate":  returnUserRate,
		"engagedUserRate": engagedUserRate,
	}
}

func (p EngagementHistory) Pointer() *EngagementHistory {
	return &p
}

func CastEngagementHistories(data []interface{}) []EngagementHistory {
	v := make([]EngagementHistory, 0)
	for _, d := range data {
		v = append(v, NewEngagementHistoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEngagementHistoriesFromDict(data []EngagementHistory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type EngagementAverage struct {
	NewUserRate     *float32 `json:"newUserRate"`
	ReturnUserRate  *float32 `json:"returnUserRate"`
	EngagedUserRate *float32 `json:"engagedUserRate"`
}

func NewEngagementAverageFromJson(data string) EngagementAverage {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEngagementAverageFromDict(dict)
}

func NewEngagementAverageFromDict(data map[string]interface{}) EngagementAverage {
	return EngagementAverage{
		NewUserRate:     core.CastFloat32(data["newUserRate"]),
		ReturnUserRate:  core.CastFloat32(data["returnUserRate"]),
		EngagedUserRate: core.CastFloat32(data["engagedUserRate"]),
	}
}

func (p EngagementAverage) ToDict() map[string]interface{} {

	var newUserRate *float32
	if p.NewUserRate != nil {
		newUserRate = p.NewUserRate
	}
	var returnUserRate *float32
	if p.ReturnUserRate != nil {
		returnUserRate = p.ReturnUserRate
	}
	var engagedUserRate *float32
	if p.EngagedUserRate != nil {
		engagedUserRate = p.EngagedUserRate
	}
	return map[string]interface{}{
		"newUserRate":     newUserRate,
		"returnUserRate":  returnUserRate,
		"engagedUserRate": engagedUserRate,
	}
}

func (p EngagementAverage) Pointer() *EngagementAverage {
	return &p
}

func CastEngagementAverages(data []interface{}) []EngagementAverage {
	v := make([]EngagementAverage, 0)
	for _, d := range data {
		v = append(v, NewEngagementAverageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEngagementAveragesFromDict(data []EngagementAverage) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Engagements struct {
	History []EngagementHistory `json:"history"`
	Avg     *EngagementAverage  `json:"avg"`
}

func NewEngagementsFromJson(data string) Engagements {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEngagementsFromDict(dict)
}

func NewEngagementsFromDict(data map[string]interface{}) Engagements {
	return Engagements{
		History: CastEngagementHistories(core.CastArray(data["history"])),
		Avg:     NewEngagementAverageFromDict(core.CastMap(data["avg"])).Pointer(),
	}
}

func (p Engagements) ToDict() map[string]interface{} {

	var history []interface{}
	if p.History != nil {
		history = CastEngagementHistoriesFromDict(
			p.History,
		)
	}
	var avg map[string]interface{}
	if p.Avg != nil {
		avg = p.Avg.ToDict()
	}
	return map[string]interface{}{
		"history": history,
		"avg":     avg,
	}
}

func (p Engagements) Pointer() *Engagements {
	return &p
}

func CastEngagementses(data []interface{}) []Engagements {
	v := make([]Engagements, 0)
	for _, d := range data {
		v = append(v, NewEngagementsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEngagementsesFromDict(data []Engagements) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChurnRateHistory struct {
	Date      *string  `json:"date"`
	ChurnRate *float32 `json:"churnRate"`
}

func NewChurnRateHistoryFromJson(data string) ChurnRateHistory {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChurnRateHistoryFromDict(dict)
}

func NewChurnRateHistoryFromDict(data map[string]interface{}) ChurnRateHistory {
	return ChurnRateHistory{
		Date:      core.CastString(data["date"]),
		ChurnRate: core.CastFloat32(data["churnRate"]),
	}
}

func (p ChurnRateHistory) ToDict() map[string]interface{} {

	var date *string
	if p.Date != nil {
		date = p.Date
	}
	var churnRate *float32
	if p.ChurnRate != nil {
		churnRate = p.ChurnRate
	}
	return map[string]interface{}{
		"date":      date,
		"churnRate": churnRate,
	}
}

func (p ChurnRateHistory) Pointer() *ChurnRateHistory {
	return &p
}

func CastChurnRateHistories(data []interface{}) []ChurnRateHistory {
	v := make([]ChurnRateHistory, 0)
	for _, d := range data {
		v = append(v, NewChurnRateHistoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChurnRateHistoriesFromDict(data []ChurnRateHistory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChurnRateAverage struct {
	ChurnRate *float32 `json:"churnRate"`
}

func NewChurnRateAverageFromJson(data string) ChurnRateAverage {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChurnRateAverageFromDict(dict)
}

func NewChurnRateAverageFromDict(data map[string]interface{}) ChurnRateAverage {
	return ChurnRateAverage{
		ChurnRate: core.CastFloat32(data["churnRate"]),
	}
}

func (p ChurnRateAverage) ToDict() map[string]interface{} {

	var churnRate *float32
	if p.ChurnRate != nil {
		churnRate = p.ChurnRate
	}
	return map[string]interface{}{
		"churnRate": churnRate,
	}
}

func (p ChurnRateAverage) Pointer() *ChurnRateAverage {
	return &p
}

func CastChurnRateAverages(data []interface{}) []ChurnRateAverage {
	v := make([]ChurnRateAverage, 0)
	for _, d := range data {
		v = append(v, NewChurnRateAverageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChurnRateAveragesFromDict(data []ChurnRateAverage) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChurnRateAggregate struct {
	History []ChurnRateHistory `json:"history"`
	Avg     *ChurnRateAverage  `json:"avg"`
}

func NewChurnRateAggregateFromJson(data string) ChurnRateAggregate {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChurnRateAggregateFromDict(dict)
}

func NewChurnRateAggregateFromDict(data map[string]interface{}) ChurnRateAggregate {
	return ChurnRateAggregate{
		History: CastChurnRateHistories(core.CastArray(data["history"])),
		Avg:     NewChurnRateAverageFromDict(core.CastMap(data["avg"])).Pointer(),
	}
}

func (p ChurnRateAggregate) ToDict() map[string]interface{} {

	var history []interface{}
	if p.History != nil {
		history = CastChurnRateHistoriesFromDict(
			p.History,
		)
	}
	var avg map[string]interface{}
	if p.Avg != nil {
		avg = p.Avg.ToDict()
	}
	return map[string]interface{}{
		"history": history,
		"avg":     avg,
	}
}

func (p ChurnRateAggregate) Pointer() *ChurnRateAggregate {
	return &p
}

func CastChurnRateAggregates(data []interface{}) []ChurnRateAggregate {
	v := make([]ChurnRateAggregate, 0)
	for _, d := range data {
		v = append(v, NewChurnRateAggregateFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChurnRateAggregatesFromDict(data []ChurnRateAggregate) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SessionDuration struct {
	Statistics   *SessionDurationStatistics           `json:"statistics"`
	Distribution []SessionDurationDistributionSegment `json:"distribution"`
}

func NewSessionDurationFromJson(data string) SessionDuration {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSessionDurationFromDict(dict)
}

func NewSessionDurationFromDict(data map[string]interface{}) SessionDuration {
	return SessionDuration{
		Statistics:   NewSessionDurationStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastSessionDurationDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p SessionDuration) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastSessionDurationDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p SessionDuration) Pointer() *SessionDuration {
	return &p
}

func CastSessionDurations(data []interface{}) []SessionDuration {
	v := make([]SessionDuration, 0)
	for _, d := range data {
		v = append(v, NewSessionDurationFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSessionDurationsFromDict(data []SessionDuration) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type UseServices struct {
	Account     *bool `json:"account"`
	Chat        *bool `json:"chat"`
	Datastore   *bool `json:"datastore"`
	Dictionary  *bool `json:"dictionary"`
	Exchange    *bool `json:"exchange"`
	Experience  *bool `json:"experience"`
	Formation   *bool `json:"formation"`
	Friend      *bool `json:"friend"`
	Inbox       *bool `json:"inbox"`
	Inventory   *bool `json:"inventory"`
	Key         *bool `json:"key"`
	Limit       *bool `json:"limit"`
	Lottery     *bool `json:"lottery"`
	Matchmaking *bool `json:"matchmaking"`
	Mission     *bool `json:"mission"`
	Money       *bool `json:"money"`
	Quest       *bool `json:"quest"`
	Ranking     *bool `json:"ranking"`
	Showcase    *bool `json:"showcase"`
	Stamina     *bool `json:"stamina"`
}

func NewUseServicesFromJson(data string) UseServices {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUseServicesFromDict(dict)
}

func NewUseServicesFromDict(data map[string]interface{}) UseServices {
	return UseServices{
		Account:     core.CastBool(data["account"]),
		Chat:        core.CastBool(data["chat"]),
		Datastore:   core.CastBool(data["datastore"]),
		Dictionary:  core.CastBool(data["dictionary"]),
		Exchange:    core.CastBool(data["exchange"]),
		Experience:  core.CastBool(data["experience"]),
		Formation:   core.CastBool(data["formation"]),
		Friend:      core.CastBool(data["friend"]),
		Inbox:       core.CastBool(data["inbox"]),
		Inventory:   core.CastBool(data["inventory"]),
		Key:         core.CastBool(data["key"]),
		Limit:       core.CastBool(data["limit"]),
		Lottery:     core.CastBool(data["lottery"]),
		Matchmaking: core.CastBool(data["matchmaking"]),
		Mission:     core.CastBool(data["mission"]),
		Money:       core.CastBool(data["money"]),
		Quest:       core.CastBool(data["quest"]),
		Ranking:     core.CastBool(data["ranking"]),
		Showcase:    core.CastBool(data["showcase"]),
		Stamina:     core.CastBool(data["stamina"]),
	}
}

func (p UseServices) ToDict() map[string]interface{} {

	var account *bool
	if p.Account != nil {
		account = p.Account
	}
	var chat *bool
	if p.Chat != nil {
		chat = p.Chat
	}
	var datastore *bool
	if p.Datastore != nil {
		datastore = p.Datastore
	}
	var dictionary *bool
	if p.Dictionary != nil {
		dictionary = p.Dictionary
	}
	var exchange *bool
	if p.Exchange != nil {
		exchange = p.Exchange
	}
	var experience *bool
	if p.Experience != nil {
		experience = p.Experience
	}
	var formation *bool
	if p.Formation != nil {
		formation = p.Formation
	}
	var friend *bool
	if p.Friend != nil {
		friend = p.Friend
	}
	var inbox *bool
	if p.Inbox != nil {
		inbox = p.Inbox
	}
	var inventory *bool
	if p.Inventory != nil {
		inventory = p.Inventory
	}
	var key *bool
	if p.Key != nil {
		key = p.Key
	}
	var limit *bool
	if p.Limit != nil {
		limit = p.Limit
	}
	var lottery *bool
	if p.Lottery != nil {
		lottery = p.Lottery
	}
	var matchmaking *bool
	if p.Matchmaking != nil {
		matchmaking = p.Matchmaking
	}
	var mission *bool
	if p.Mission != nil {
		mission = p.Mission
	}
	var money *bool
	if p.Money != nil {
		money = p.Money
	}
	var quest *bool
	if p.Quest != nil {
		quest = p.Quest
	}
	var ranking *bool
	if p.Ranking != nil {
		ranking = p.Ranking
	}
	var showcase *bool
	if p.Showcase != nil {
		showcase = p.Showcase
	}
	var stamina *bool
	if p.Stamina != nil {
		stamina = p.Stamina
	}
	return map[string]interface{}{
		"account":     account,
		"chat":        chat,
		"datastore":   datastore,
		"dictionary":  dictionary,
		"exchange":    exchange,
		"experience":  experience,
		"formation":   formation,
		"friend":      friend,
		"inbox":       inbox,
		"inventory":   inventory,
		"key":         key,
		"limit":       limit,
		"lottery":     lottery,
		"matchmaking": matchmaking,
		"mission":     mission,
		"money":       money,
		"quest":       quest,
		"ranking":     ranking,
		"showcase":    showcase,
		"stamina":     stamina,
	}
}

func (p UseServices) Pointer() *UseServices {
	return &p
}

func CastUseServiceses(data []interface{}) []UseServices {
	v := make([]UseServices, 0)
	for _, d := range data {
		v = append(v, NewUseServicesFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastUseServicesesFromDict(data []UseServices) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GeneralMetrics struct {
	DauWauMau       *GeneralDauWauMau   `json:"dauWauMau"`
	SessionDuration *SessionDuration    `json:"sessionDuration"`
	FirstEngagement *FirstEngagement    `json:"firstEngagement"`
	Engagements     *Engagements        `json:"engagements"`
	ChurnRates      *ChurnRateAggregate `json:"churnRates"`
	UseServices     *UseServices        `json:"useServices"`
}

func NewGeneralMetricsFromJson(data string) GeneralMetrics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGeneralMetricsFromDict(dict)
}

func NewGeneralMetricsFromDict(data map[string]interface{}) GeneralMetrics {
	return GeneralMetrics{
		DauWauMau:       NewGeneralDauWauMauFromDict(core.CastMap(data["dauWauMau"])).Pointer(),
		SessionDuration: NewSessionDurationFromDict(core.CastMap(data["sessionDuration"])).Pointer(),
		FirstEngagement: NewFirstEngagementFromDict(core.CastMap(data["firstEngagement"])).Pointer(),
		Engagements:     NewEngagementsFromDict(core.CastMap(data["engagements"])).Pointer(),
		ChurnRates:      NewChurnRateAggregateFromDict(core.CastMap(data["churnRates"])).Pointer(),
		UseServices:     NewUseServicesFromDict(core.CastMap(data["useServices"])).Pointer(),
	}
}

func (p GeneralMetrics) ToDict() map[string]interface{} {

	var dauWauMau map[string]interface{}
	if p.DauWauMau != nil {
		dauWauMau = p.DauWauMau.ToDict()
	}
	var sessionDuration map[string]interface{}
	if p.SessionDuration != nil {
		sessionDuration = p.SessionDuration.ToDict()
	}
	var firstEngagement map[string]interface{}
	if p.FirstEngagement != nil {
		firstEngagement = p.FirstEngagement.ToDict()
	}
	var engagements map[string]interface{}
	if p.Engagements != nil {
		engagements = p.Engagements.ToDict()
	}
	var churnRates map[string]interface{}
	if p.ChurnRates != nil {
		churnRates = p.ChurnRates.ToDict()
	}
	var useServices map[string]interface{}
	if p.UseServices != nil {
		useServices = p.UseServices.ToDict()
	}
	return map[string]interface{}{
		"dauWauMau":       dauWauMau,
		"sessionDuration": sessionDuration,
		"firstEngagement": firstEngagement,
		"engagements":     engagements,
		"churnRates":      churnRates,
		"useServices":     useServices,
	}
}

func (p GeneralMetrics) Pointer() *GeneralMetrics {
	return &p
}

func CastGeneralMetricses(data []interface{}) []GeneralMetrics {
	v := make([]GeneralMetrics, 0)
	for _, d := range data {
		v = append(v, NewGeneralMetricsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGeneralMetricsesFromDict(data []GeneralMetrics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AccountNamespaceStatistics struct {
	Signup             *int64 `json:"signup"`
	Authentication     *int64 `json:"authentication"`
	RegisteredTakeOver *int64 `json:"registeredTakeOver"`
	RemoveTakeOver     *int64 `json:"removeTakeOver"`
	ExecuteTakeOver    *int64 `json:"executeTakeOver"`
}

func NewAccountNamespaceStatisticsFromJson(data string) AccountNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAccountNamespaceStatisticsFromDict(dict)
}

func NewAccountNamespaceStatisticsFromDict(data map[string]interface{}) AccountNamespaceStatistics {
	return AccountNamespaceStatistics{
		Signup:             core.CastInt64(data["signup"]),
		Authentication:     core.CastInt64(data["authentication"]),
		RegisteredTakeOver: core.CastInt64(data["registeredTakeOver"]),
		RemoveTakeOver:     core.CastInt64(data["removeTakeOver"]),
		ExecuteTakeOver:    core.CastInt64(data["executeTakeOver"]),
	}
}

func (p AccountNamespaceStatistics) ToDict() map[string]interface{} {

	var signup *int64
	if p.Signup != nil {
		signup = p.Signup
	}
	var authentication *int64
	if p.Authentication != nil {
		authentication = p.Authentication
	}
	var registeredTakeOver *int64
	if p.RegisteredTakeOver != nil {
		registeredTakeOver = p.RegisteredTakeOver
	}
	var removeTakeOver *int64
	if p.RemoveTakeOver != nil {
		removeTakeOver = p.RemoveTakeOver
	}
	var executeTakeOver *int64
	if p.ExecuteTakeOver != nil {
		executeTakeOver = p.ExecuteTakeOver
	}
	return map[string]interface{}{
		"signup":             signup,
		"authentication":     authentication,
		"registeredTakeOver": registeredTakeOver,
		"removeTakeOver":     removeTakeOver,
		"executeTakeOver":    executeTakeOver,
	}
}

func (p AccountNamespaceStatistics) Pointer() *AccountNamespaceStatistics {
	return &p
}

func CastAccountNamespaceStatisticses(data []interface{}) []AccountNamespaceStatistics {
	v := make([]AccountNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewAccountNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccountNamespaceStatisticsesFromDict(data []AccountNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AccountNamespaceTypeDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewAccountNamespaceTypeDistributionStatisticsFromJson(data string) AccountNamespaceTypeDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAccountNamespaceTypeDistributionStatisticsFromDict(dict)
}

func NewAccountNamespaceTypeDistributionStatisticsFromDict(data map[string]interface{}) AccountNamespaceTypeDistributionStatistics {
	return AccountNamespaceTypeDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p AccountNamespaceTypeDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p AccountNamespaceTypeDistributionStatistics) Pointer() *AccountNamespaceTypeDistributionStatistics {
	return &p
}

func CastAccountNamespaceTypeDistributionStatisticses(data []interface{}) []AccountNamespaceTypeDistributionStatistics {
	v := make([]AccountNamespaceTypeDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewAccountNamespaceTypeDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccountNamespaceTypeDistributionStatisticsesFromDict(data []AccountNamespaceTypeDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AccountNamespaceTypeDistributionSegment struct {
	Type  *int64 `json:"type"`
	Count *int64 `json:"count"`
}

func NewAccountNamespaceTypeDistributionSegmentFromJson(data string) AccountNamespaceTypeDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAccountNamespaceTypeDistributionSegmentFromDict(dict)
}

func NewAccountNamespaceTypeDistributionSegmentFromDict(data map[string]interface{}) AccountNamespaceTypeDistributionSegment {
	return AccountNamespaceTypeDistributionSegment{
		Type:  core.CastInt64(data["type"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p AccountNamespaceTypeDistributionSegment) ToDict() map[string]interface{} {

	var _type *int64
	if p.Type != nil {
		_type = p.Type
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"type":  _type,
		"count": count,
	}
}

func (p AccountNamespaceTypeDistributionSegment) Pointer() *AccountNamespaceTypeDistributionSegment {
	return &p
}

func CastAccountNamespaceTypeDistributionSegments(data []interface{}) []AccountNamespaceTypeDistributionSegment {
	v := make([]AccountNamespaceTypeDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewAccountNamespaceTypeDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccountNamespaceTypeDistributionSegmentsFromDict(data []AccountNamespaceTypeDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AccountNamespaceTypeDistribution struct {
	Statistics   *AccountNamespaceTypeDistributionStatistics `json:"statistics"`
	Distribution []AccountNamespaceTypeDistributionSegment   `json:"distribution"`
}

func NewAccountNamespaceTypeDistributionFromJson(data string) AccountNamespaceTypeDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAccountNamespaceTypeDistributionFromDict(dict)
}

func NewAccountNamespaceTypeDistributionFromDict(data map[string]interface{}) AccountNamespaceTypeDistribution {
	return AccountNamespaceTypeDistribution{
		Statistics:   NewAccountNamespaceTypeDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastAccountNamespaceTypeDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p AccountNamespaceTypeDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastAccountNamespaceTypeDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p AccountNamespaceTypeDistribution) Pointer() *AccountNamespaceTypeDistribution {
	return &p
}

func CastAccountNamespaceTypeDistributions(data []interface{}) []AccountNamespaceTypeDistribution {
	v := make([]AccountNamespaceTypeDistribution, 0)
	for _, d := range data {
		v = append(v, NewAccountNamespaceTypeDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccountNamespaceTypeDistributionsFromDict(data []AccountNamespaceTypeDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AccountNamespaceDistributions struct {
	Type *AccountNamespaceTypeDistribution `json:"type"`
}

func NewAccountNamespaceDistributionsFromJson(data string) AccountNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAccountNamespaceDistributionsFromDict(dict)
}

func NewAccountNamespaceDistributionsFromDict(data map[string]interface{}) AccountNamespaceDistributions {
	return AccountNamespaceDistributions{
		Type: NewAccountNamespaceTypeDistributionFromDict(core.CastMap(data["type"])).Pointer(),
	}
}

func (p AccountNamespaceDistributions) ToDict() map[string]interface{} {

	var _type map[string]interface{}
	if p.Type != nil {
		_type = p.Type.ToDict()
	}
	return map[string]interface{}{
		"type": _type,
	}
}

func (p AccountNamespaceDistributions) Pointer() *AccountNamespaceDistributions {
	return &p
}

func CastAccountNamespaceDistributionses(data []interface{}) []AccountNamespaceDistributions {
	v := make([]AccountNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewAccountNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccountNamespaceDistributionsesFromDict(data []AccountNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AccountNamespace struct {
	NamespaceId   *string                        `json:"namespaceId"`
	Year          *int32                         `json:"year"`
	Month         *int32                         `json:"month"`
	Day           *int32                         `json:"day"`
	NamespaceName *string                        `json:"namespaceName"`
	Statistics    *AccountNamespaceStatistics    `json:"statistics"`
	Distributions *AccountNamespaceDistributions `json:"distributions"`
}

func NewAccountNamespaceFromJson(data string) AccountNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAccountNamespaceFromDict(dict)
}

func NewAccountNamespaceFromDict(data map[string]interface{}) AccountNamespace {
	return AccountNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewAccountNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewAccountNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p AccountNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p AccountNamespace) Pointer() *AccountNamespace {
	return &p
}

func CastAccountNamespaces(data []interface{}) []AccountNamespace {
	v := make([]AccountNamespace, 0)
	for _, d := range data {
		v = append(v, NewAccountNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccountNamespacesFromDict(data []AccountNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespaceStatistics struct {
	Post            *int64 `json:"post"`
	CreateRoom      *int64 `json:"createRoom"`
	DeleteRoom      *int64 `json:"deleteRoom"`
	CreateSubscribe *int64 `json:"createSubscribe"`
	DeleteSubscribe *int64 `json:"deleteSubscribe"`
}

func NewChatNamespaceStatisticsFromJson(data string) ChatNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespaceStatisticsFromDict(dict)
}

func NewChatNamespaceStatisticsFromDict(data map[string]interface{}) ChatNamespaceStatistics {
	return ChatNamespaceStatistics{
		Post:            core.CastInt64(data["post"]),
		CreateRoom:      core.CastInt64(data["createRoom"]),
		DeleteRoom:      core.CastInt64(data["deleteRoom"]),
		CreateSubscribe: core.CastInt64(data["createSubscribe"]),
		DeleteSubscribe: core.CastInt64(data["deleteSubscribe"]),
	}
}

func (p ChatNamespaceStatistics) ToDict() map[string]interface{} {

	var post *int64
	if p.Post != nil {
		post = p.Post
	}
	var createRoom *int64
	if p.CreateRoom != nil {
		createRoom = p.CreateRoom
	}
	var deleteRoom *int64
	if p.DeleteRoom != nil {
		deleteRoom = p.DeleteRoom
	}
	var createSubscribe *int64
	if p.CreateSubscribe != nil {
		createSubscribe = p.CreateSubscribe
	}
	var deleteSubscribe *int64
	if p.DeleteSubscribe != nil {
		deleteSubscribe = p.DeleteSubscribe
	}
	return map[string]interface{}{
		"post":            post,
		"createRoom":      createRoom,
		"deleteRoom":      deleteRoom,
		"createSubscribe": createSubscribe,
		"deleteSubscribe": deleteSubscribe,
	}
}

func (p ChatNamespaceStatistics) Pointer() *ChatNamespaceStatistics {
	return &p
}

func CastChatNamespaceStatisticses(data []interface{}) []ChatNamespaceStatistics {
	v := make([]ChatNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewChatNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespaceStatisticsesFromDict(data []ChatNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespacePostByRoomDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewChatNamespacePostByRoomDistributionStatisticsFromJson(data string) ChatNamespacePostByRoomDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespacePostByRoomDistributionStatisticsFromDict(dict)
}

func NewChatNamespacePostByRoomDistributionStatisticsFromDict(data map[string]interface{}) ChatNamespacePostByRoomDistributionStatistics {
	return ChatNamespacePostByRoomDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ChatNamespacePostByRoomDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ChatNamespacePostByRoomDistributionStatistics) Pointer() *ChatNamespacePostByRoomDistributionStatistics {
	return &p
}

func CastChatNamespacePostByRoomDistributionStatisticses(data []interface{}) []ChatNamespacePostByRoomDistributionStatistics {
	v := make([]ChatNamespacePostByRoomDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewChatNamespacePostByRoomDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacePostByRoomDistributionStatisticsesFromDict(data []ChatNamespacePostByRoomDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespacePostByRoomDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewChatNamespacePostByRoomDistributionSegmentFromJson(data string) ChatNamespacePostByRoomDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespacePostByRoomDistributionSegmentFromDict(dict)
}

func NewChatNamespacePostByRoomDistributionSegmentFromDict(data map[string]interface{}) ChatNamespacePostByRoomDistributionSegment {
	return ChatNamespacePostByRoomDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ChatNamespacePostByRoomDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ChatNamespacePostByRoomDistributionSegment) Pointer() *ChatNamespacePostByRoomDistributionSegment {
	return &p
}

func CastChatNamespacePostByRoomDistributionSegments(data []interface{}) []ChatNamespacePostByRoomDistributionSegment {
	v := make([]ChatNamespacePostByRoomDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewChatNamespacePostByRoomDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacePostByRoomDistributionSegmentsFromDict(data []ChatNamespacePostByRoomDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespacePostByRoomDistribution struct {
	Statistics   *ChatNamespacePostByRoomDistributionStatistics `json:"statistics"`
	Distribution []ChatNamespacePostByRoomDistributionSegment   `json:"distribution"`
}

func NewChatNamespacePostByRoomDistributionFromJson(data string) ChatNamespacePostByRoomDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespacePostByRoomDistributionFromDict(dict)
}

func NewChatNamespacePostByRoomDistributionFromDict(data map[string]interface{}) ChatNamespacePostByRoomDistribution {
	return ChatNamespacePostByRoomDistribution{
		Statistics:   NewChatNamespacePostByRoomDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastChatNamespacePostByRoomDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ChatNamespacePostByRoomDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastChatNamespacePostByRoomDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ChatNamespacePostByRoomDistribution) Pointer() *ChatNamespacePostByRoomDistribution {
	return &p
}

func CastChatNamespacePostByRoomDistributions(data []interface{}) []ChatNamespacePostByRoomDistribution {
	v := make([]ChatNamespacePostByRoomDistribution, 0)
	for _, d := range data {
		v = append(v, NewChatNamespacePostByRoomDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacePostByRoomDistributionsFromDict(data []ChatNamespacePostByRoomDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespacePostByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewChatNamespacePostByUserDistributionStatisticsFromJson(data string) ChatNamespacePostByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespacePostByUserDistributionStatisticsFromDict(dict)
}

func NewChatNamespacePostByUserDistributionStatisticsFromDict(data map[string]interface{}) ChatNamespacePostByUserDistributionStatistics {
	return ChatNamespacePostByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ChatNamespacePostByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ChatNamespacePostByUserDistributionStatistics) Pointer() *ChatNamespacePostByUserDistributionStatistics {
	return &p
}

func CastChatNamespacePostByUserDistributionStatisticses(data []interface{}) []ChatNamespacePostByUserDistributionStatistics {
	v := make([]ChatNamespacePostByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewChatNamespacePostByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacePostByUserDistributionStatisticsesFromDict(data []ChatNamespacePostByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespacePostByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewChatNamespacePostByUserDistributionSegmentFromJson(data string) ChatNamespacePostByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespacePostByUserDistributionSegmentFromDict(dict)
}

func NewChatNamespacePostByUserDistributionSegmentFromDict(data map[string]interface{}) ChatNamespacePostByUserDistributionSegment {
	return ChatNamespacePostByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ChatNamespacePostByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ChatNamespacePostByUserDistributionSegment) Pointer() *ChatNamespacePostByUserDistributionSegment {
	return &p
}

func CastChatNamespacePostByUserDistributionSegments(data []interface{}) []ChatNamespacePostByUserDistributionSegment {
	v := make([]ChatNamespacePostByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewChatNamespacePostByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacePostByUserDistributionSegmentsFromDict(data []ChatNamespacePostByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespacePostByUserDistribution struct {
	Statistics   *ChatNamespacePostByUserDistributionStatistics `json:"statistics"`
	Distribution []ChatNamespacePostByUserDistributionSegment   `json:"distribution"`
}

func NewChatNamespacePostByUserDistributionFromJson(data string) ChatNamespacePostByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespacePostByUserDistributionFromDict(dict)
}

func NewChatNamespacePostByUserDistributionFromDict(data map[string]interface{}) ChatNamespacePostByUserDistribution {
	return ChatNamespacePostByUserDistribution{
		Statistics:   NewChatNamespacePostByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastChatNamespacePostByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ChatNamespacePostByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastChatNamespacePostByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ChatNamespacePostByUserDistribution) Pointer() *ChatNamespacePostByUserDistribution {
	return &p
}

func CastChatNamespacePostByUserDistributions(data []interface{}) []ChatNamespacePostByUserDistribution {
	v := make([]ChatNamespacePostByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewChatNamespacePostByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacePostByUserDistributionsFromDict(data []ChatNamespacePostByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespacePostByCategoryDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewChatNamespacePostByCategoryDistributionStatisticsFromJson(data string) ChatNamespacePostByCategoryDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespacePostByCategoryDistributionStatisticsFromDict(dict)
}

func NewChatNamespacePostByCategoryDistributionStatisticsFromDict(data map[string]interface{}) ChatNamespacePostByCategoryDistributionStatistics {
	return ChatNamespacePostByCategoryDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ChatNamespacePostByCategoryDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ChatNamespacePostByCategoryDistributionStatistics) Pointer() *ChatNamespacePostByCategoryDistributionStatistics {
	return &p
}

func CastChatNamespacePostByCategoryDistributionStatisticses(data []interface{}) []ChatNamespacePostByCategoryDistributionStatistics {
	v := make([]ChatNamespacePostByCategoryDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewChatNamespacePostByCategoryDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacePostByCategoryDistributionStatisticsesFromDict(data []ChatNamespacePostByCategoryDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespacePostByCategoryDistributionSegment struct {
	Category *int64 `json:"category"`
	Count    *int64 `json:"count"`
}

func NewChatNamespacePostByCategoryDistributionSegmentFromJson(data string) ChatNamespacePostByCategoryDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespacePostByCategoryDistributionSegmentFromDict(dict)
}

func NewChatNamespacePostByCategoryDistributionSegmentFromDict(data map[string]interface{}) ChatNamespacePostByCategoryDistributionSegment {
	return ChatNamespacePostByCategoryDistributionSegment{
		Category: core.CastInt64(data["category"]),
		Count:    core.CastInt64(data["count"]),
	}
}

func (p ChatNamespacePostByCategoryDistributionSegment) ToDict() map[string]interface{} {

	var category *int64
	if p.Category != nil {
		category = p.Category
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"category": category,
		"count":    count,
	}
}

func (p ChatNamespacePostByCategoryDistributionSegment) Pointer() *ChatNamespacePostByCategoryDistributionSegment {
	return &p
}

func CastChatNamespacePostByCategoryDistributionSegments(data []interface{}) []ChatNamespacePostByCategoryDistributionSegment {
	v := make([]ChatNamespacePostByCategoryDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewChatNamespacePostByCategoryDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacePostByCategoryDistributionSegmentsFromDict(data []ChatNamespacePostByCategoryDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespacePostByCategoryDistribution struct {
	Statistics   *ChatNamespacePostByCategoryDistributionStatistics `json:"statistics"`
	Distribution []ChatNamespacePostByCategoryDistributionSegment   `json:"distribution"`
}

func NewChatNamespacePostByCategoryDistributionFromJson(data string) ChatNamespacePostByCategoryDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespacePostByCategoryDistributionFromDict(dict)
}

func NewChatNamespacePostByCategoryDistributionFromDict(data map[string]interface{}) ChatNamespacePostByCategoryDistribution {
	return ChatNamespacePostByCategoryDistribution{
		Statistics:   NewChatNamespacePostByCategoryDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastChatNamespacePostByCategoryDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ChatNamespacePostByCategoryDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastChatNamespacePostByCategoryDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ChatNamespacePostByCategoryDistribution) Pointer() *ChatNamespacePostByCategoryDistribution {
	return &p
}

func CastChatNamespacePostByCategoryDistributions(data []interface{}) []ChatNamespacePostByCategoryDistribution {
	v := make([]ChatNamespacePostByCategoryDistribution, 0)
	for _, d := range data {
		v = append(v, NewChatNamespacePostByCategoryDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacePostByCategoryDistributionsFromDict(data []ChatNamespacePostByCategoryDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespaceDistributions struct {
	PostByRoom     *ChatNamespacePostByRoomDistribution     `json:"postByRoom"`
	PostByUser     *ChatNamespacePostByUserDistribution     `json:"postByUser"`
	PostByCategory *ChatNamespacePostByCategoryDistribution `json:"postByCategory"`
}

func NewChatNamespaceDistributionsFromJson(data string) ChatNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespaceDistributionsFromDict(dict)
}

func NewChatNamespaceDistributionsFromDict(data map[string]interface{}) ChatNamespaceDistributions {
	return ChatNamespaceDistributions{
		PostByRoom:     NewChatNamespacePostByRoomDistributionFromDict(core.CastMap(data["postByRoom"])).Pointer(),
		PostByUser:     NewChatNamespacePostByUserDistributionFromDict(core.CastMap(data["postByUser"])).Pointer(),
		PostByCategory: NewChatNamespacePostByCategoryDistributionFromDict(core.CastMap(data["postByCategory"])).Pointer(),
	}
}

func (p ChatNamespaceDistributions) ToDict() map[string]interface{} {

	var postByRoom map[string]interface{}
	if p.PostByRoom != nil {
		postByRoom = p.PostByRoom.ToDict()
	}
	var postByUser map[string]interface{}
	if p.PostByUser != nil {
		postByUser = p.PostByUser.ToDict()
	}
	var postByCategory map[string]interface{}
	if p.PostByCategory != nil {
		postByCategory = p.PostByCategory.ToDict()
	}
	return map[string]interface{}{
		"postByRoom":     postByRoom,
		"postByUser":     postByUser,
		"postByCategory": postByCategory,
	}
}

func (p ChatNamespaceDistributions) Pointer() *ChatNamespaceDistributions {
	return &p
}

func CastChatNamespaceDistributionses(data []interface{}) []ChatNamespaceDistributions {
	v := make([]ChatNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewChatNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespaceDistributionsesFromDict(data []ChatNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChatNamespace struct {
	NamespaceId   *string                     `json:"namespaceId"`
	Year          *int32                      `json:"year"`
	Month         *int32                      `json:"month"`
	Day           *int32                      `json:"day"`
	NamespaceName *string                     `json:"namespaceName"`
	Statistics    *ChatNamespaceStatistics    `json:"statistics"`
	Distributions *ChatNamespaceDistributions `json:"distributions"`
}

func NewChatNamespaceFromJson(data string) ChatNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChatNamespaceFromDict(dict)
}

func NewChatNamespaceFromDict(data map[string]interface{}) ChatNamespace {
	return ChatNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewChatNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewChatNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p ChatNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p ChatNamespace) Pointer() *ChatNamespace {
	return &p
}

func CastChatNamespaces(data []interface{}) []ChatNamespace {
	v := make([]ChatNamespace, 0)
	for _, d := range data {
		v = append(v, NewChatNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChatNamespacesFromDict(data []ChatNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceStatistics struct {
	Upload   *int64 `json:"upload"`
	Download *int64 `json:"download"`
}

func NewDatastoreNamespaceStatisticsFromJson(data string) DatastoreNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceStatisticsFromDict(dict)
}

func NewDatastoreNamespaceStatisticsFromDict(data map[string]interface{}) DatastoreNamespaceStatistics {
	return DatastoreNamespaceStatistics{
		Upload:   core.CastInt64(data["upload"]),
		Download: core.CastInt64(data["download"]),
	}
}

func (p DatastoreNamespaceStatistics) ToDict() map[string]interface{} {

	var upload *int64
	if p.Upload != nil {
		upload = p.Upload
	}
	var download *int64
	if p.Download != nil {
		download = p.Download
	}
	return map[string]interface{}{
		"upload":   upload,
		"download": download,
	}
}

func (p DatastoreNamespaceStatistics) Pointer() *DatastoreNamespaceStatistics {
	return &p
}

func CastDatastoreNamespaceStatisticses(data []interface{}) []DatastoreNamespaceStatistics {
	v := make([]DatastoreNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceStatisticsesFromDict(data []DatastoreNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceDownloadByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewDatastoreNamespaceDownloadByUserDistributionStatisticsFromJson(data string) DatastoreNamespaceDownloadByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceDownloadByUserDistributionStatisticsFromDict(dict)
}

func NewDatastoreNamespaceDownloadByUserDistributionStatisticsFromDict(data map[string]interface{}) DatastoreNamespaceDownloadByUserDistributionStatistics {
	return DatastoreNamespaceDownloadByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p DatastoreNamespaceDownloadByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p DatastoreNamespaceDownloadByUserDistributionStatistics) Pointer() *DatastoreNamespaceDownloadByUserDistributionStatistics {
	return &p
}

func CastDatastoreNamespaceDownloadByUserDistributionStatisticses(data []interface{}) []DatastoreNamespaceDownloadByUserDistributionStatistics {
	v := make([]DatastoreNamespaceDownloadByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceDownloadByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceDownloadByUserDistributionStatisticsesFromDict(data []DatastoreNamespaceDownloadByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceDownloadByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewDatastoreNamespaceDownloadByUserDistributionSegmentFromJson(data string) DatastoreNamespaceDownloadByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceDownloadByUserDistributionSegmentFromDict(dict)
}

func NewDatastoreNamespaceDownloadByUserDistributionSegmentFromDict(data map[string]interface{}) DatastoreNamespaceDownloadByUserDistributionSegment {
	return DatastoreNamespaceDownloadByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p DatastoreNamespaceDownloadByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p DatastoreNamespaceDownloadByUserDistributionSegment) Pointer() *DatastoreNamespaceDownloadByUserDistributionSegment {
	return &p
}

func CastDatastoreNamespaceDownloadByUserDistributionSegments(data []interface{}) []DatastoreNamespaceDownloadByUserDistributionSegment {
	v := make([]DatastoreNamespaceDownloadByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceDownloadByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceDownloadByUserDistributionSegmentsFromDict(data []DatastoreNamespaceDownloadByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceDownloadByUserDistribution struct {
	Statistics   *DatastoreNamespaceDownloadByUserDistributionStatistics `json:"statistics"`
	Distribution []DatastoreNamespaceDownloadByUserDistributionSegment   `json:"distribution"`
}

func NewDatastoreNamespaceDownloadByUserDistributionFromJson(data string) DatastoreNamespaceDownloadByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceDownloadByUserDistributionFromDict(dict)
}

func NewDatastoreNamespaceDownloadByUserDistributionFromDict(data map[string]interface{}) DatastoreNamespaceDownloadByUserDistribution {
	return DatastoreNamespaceDownloadByUserDistribution{
		Statistics:   NewDatastoreNamespaceDownloadByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastDatastoreNamespaceDownloadByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p DatastoreNamespaceDownloadByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastDatastoreNamespaceDownloadByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p DatastoreNamespaceDownloadByUserDistribution) Pointer() *DatastoreNamespaceDownloadByUserDistribution {
	return &p
}

func CastDatastoreNamespaceDownloadByUserDistributions(data []interface{}) []DatastoreNamespaceDownloadByUserDistribution {
	v := make([]DatastoreNamespaceDownloadByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceDownloadByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceDownloadByUserDistributionsFromDict(data []DatastoreNamespaceDownloadByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceUploadByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewDatastoreNamespaceUploadByUserDistributionStatisticsFromJson(data string) DatastoreNamespaceUploadByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceUploadByUserDistributionStatisticsFromDict(dict)
}

func NewDatastoreNamespaceUploadByUserDistributionStatisticsFromDict(data map[string]interface{}) DatastoreNamespaceUploadByUserDistributionStatistics {
	return DatastoreNamespaceUploadByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p DatastoreNamespaceUploadByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p DatastoreNamespaceUploadByUserDistributionStatistics) Pointer() *DatastoreNamespaceUploadByUserDistributionStatistics {
	return &p
}

func CastDatastoreNamespaceUploadByUserDistributionStatisticses(data []interface{}) []DatastoreNamespaceUploadByUserDistributionStatistics {
	v := make([]DatastoreNamespaceUploadByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceUploadByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceUploadByUserDistributionStatisticsesFromDict(data []DatastoreNamespaceUploadByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceUploadByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewDatastoreNamespaceUploadByUserDistributionSegmentFromJson(data string) DatastoreNamespaceUploadByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceUploadByUserDistributionSegmentFromDict(dict)
}

func NewDatastoreNamespaceUploadByUserDistributionSegmentFromDict(data map[string]interface{}) DatastoreNamespaceUploadByUserDistributionSegment {
	return DatastoreNamespaceUploadByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p DatastoreNamespaceUploadByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p DatastoreNamespaceUploadByUserDistributionSegment) Pointer() *DatastoreNamespaceUploadByUserDistributionSegment {
	return &p
}

func CastDatastoreNamespaceUploadByUserDistributionSegments(data []interface{}) []DatastoreNamespaceUploadByUserDistributionSegment {
	v := make([]DatastoreNamespaceUploadByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceUploadByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceUploadByUserDistributionSegmentsFromDict(data []DatastoreNamespaceUploadByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceUploadByUserDistribution struct {
	Statistics   *DatastoreNamespaceUploadByUserDistributionStatistics `json:"statistics"`
	Distribution []DatastoreNamespaceUploadByUserDistributionSegment   `json:"distribution"`
}

func NewDatastoreNamespaceUploadByUserDistributionFromJson(data string) DatastoreNamespaceUploadByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceUploadByUserDistributionFromDict(dict)
}

func NewDatastoreNamespaceUploadByUserDistributionFromDict(data map[string]interface{}) DatastoreNamespaceUploadByUserDistribution {
	return DatastoreNamespaceUploadByUserDistribution{
		Statistics:   NewDatastoreNamespaceUploadByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastDatastoreNamespaceUploadByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p DatastoreNamespaceUploadByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastDatastoreNamespaceUploadByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p DatastoreNamespaceUploadByUserDistribution) Pointer() *DatastoreNamespaceUploadByUserDistribution {
	return &p
}

func CastDatastoreNamespaceUploadByUserDistributions(data []interface{}) []DatastoreNamespaceUploadByUserDistribution {
	v := make([]DatastoreNamespaceUploadByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceUploadByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceUploadByUserDistributionsFromDict(data []DatastoreNamespaceUploadByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceDataSizeDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewDatastoreNamespaceDataSizeDistributionStatisticsFromJson(data string) DatastoreNamespaceDataSizeDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceDataSizeDistributionStatisticsFromDict(dict)
}

func NewDatastoreNamespaceDataSizeDistributionStatisticsFromDict(data map[string]interface{}) DatastoreNamespaceDataSizeDistributionStatistics {
	return DatastoreNamespaceDataSizeDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p DatastoreNamespaceDataSizeDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p DatastoreNamespaceDataSizeDistributionStatistics) Pointer() *DatastoreNamespaceDataSizeDistributionStatistics {
	return &p
}

func CastDatastoreNamespaceDataSizeDistributionStatisticses(data []interface{}) []DatastoreNamespaceDataSizeDistributionStatistics {
	v := make([]DatastoreNamespaceDataSizeDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceDataSizeDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceDataSizeDistributionStatisticsesFromDict(data []DatastoreNamespaceDataSizeDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceDataSizeDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewDatastoreNamespaceDataSizeDistributionSegmentFromJson(data string) DatastoreNamespaceDataSizeDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceDataSizeDistributionSegmentFromDict(dict)
}

func NewDatastoreNamespaceDataSizeDistributionSegmentFromDict(data map[string]interface{}) DatastoreNamespaceDataSizeDistributionSegment {
	return DatastoreNamespaceDataSizeDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p DatastoreNamespaceDataSizeDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p DatastoreNamespaceDataSizeDistributionSegment) Pointer() *DatastoreNamespaceDataSizeDistributionSegment {
	return &p
}

func CastDatastoreNamespaceDataSizeDistributionSegments(data []interface{}) []DatastoreNamespaceDataSizeDistributionSegment {
	v := make([]DatastoreNamespaceDataSizeDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceDataSizeDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceDataSizeDistributionSegmentsFromDict(data []DatastoreNamespaceDataSizeDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceDataSizeDistribution struct {
	Statistics   *DatastoreNamespaceDataSizeDistributionStatistics `json:"statistics"`
	Distribution []DatastoreNamespaceDataSizeDistributionSegment   `json:"distribution"`
}

func NewDatastoreNamespaceDataSizeDistributionFromJson(data string) DatastoreNamespaceDataSizeDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceDataSizeDistributionFromDict(dict)
}

func NewDatastoreNamespaceDataSizeDistributionFromDict(data map[string]interface{}) DatastoreNamespaceDataSizeDistribution {
	return DatastoreNamespaceDataSizeDistribution{
		Statistics:   NewDatastoreNamespaceDataSizeDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastDatastoreNamespaceDataSizeDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p DatastoreNamespaceDataSizeDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastDatastoreNamespaceDataSizeDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p DatastoreNamespaceDataSizeDistribution) Pointer() *DatastoreNamespaceDataSizeDistribution {
	return &p
}

func CastDatastoreNamespaceDataSizeDistributions(data []interface{}) []DatastoreNamespaceDataSizeDistribution {
	v := make([]DatastoreNamespaceDataSizeDistribution, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceDataSizeDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceDataSizeDistributionsFromDict(data []DatastoreNamespaceDataSizeDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespaceDistributions struct {
	DownloadByUser *DatastoreNamespaceDownloadByUserDistribution `json:"downloadByUser"`
	UploadByUser   *DatastoreNamespaceUploadByUserDistribution   `json:"uploadByUser"`
	DataSize       *DatastoreNamespaceDataSizeDistribution       `json:"dataSize"`
}

func NewDatastoreNamespaceDistributionsFromJson(data string) DatastoreNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceDistributionsFromDict(dict)
}

func NewDatastoreNamespaceDistributionsFromDict(data map[string]interface{}) DatastoreNamespaceDistributions {
	return DatastoreNamespaceDistributions{
		DownloadByUser: NewDatastoreNamespaceDownloadByUserDistributionFromDict(core.CastMap(data["downloadByUser"])).Pointer(),
		UploadByUser:   NewDatastoreNamespaceUploadByUserDistributionFromDict(core.CastMap(data["uploadByUser"])).Pointer(),
		DataSize:       NewDatastoreNamespaceDataSizeDistributionFromDict(core.CastMap(data["dataSize"])).Pointer(),
	}
}

func (p DatastoreNamespaceDistributions) ToDict() map[string]interface{} {

	var downloadByUser map[string]interface{}
	if p.DownloadByUser != nil {
		downloadByUser = p.DownloadByUser.ToDict()
	}
	var uploadByUser map[string]interface{}
	if p.UploadByUser != nil {
		uploadByUser = p.UploadByUser.ToDict()
	}
	var dataSize map[string]interface{}
	if p.DataSize != nil {
		dataSize = p.DataSize.ToDict()
	}
	return map[string]interface{}{
		"downloadByUser": downloadByUser,
		"uploadByUser":   uploadByUser,
		"dataSize":       dataSize,
	}
}

func (p DatastoreNamespaceDistributions) Pointer() *DatastoreNamespaceDistributions {
	return &p
}

func CastDatastoreNamespaceDistributionses(data []interface{}) []DatastoreNamespaceDistributions {
	v := make([]DatastoreNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespaceDistributionsesFromDict(data []DatastoreNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DatastoreNamespace struct {
	NamespaceId   *string                          `json:"namespaceId"`
	Year          *int32                           `json:"year"`
	Month         *int32                           `json:"month"`
	Day           *int32                           `json:"day"`
	NamespaceName *string                          `json:"namespaceName"`
	Statistics    *DatastoreNamespaceStatistics    `json:"statistics"`
	Distributions *DatastoreNamespaceDistributions `json:"distributions"`
}

func NewDatastoreNamespaceFromJson(data string) DatastoreNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDatastoreNamespaceFromDict(dict)
}

func NewDatastoreNamespaceFromDict(data map[string]interface{}) DatastoreNamespace {
	return DatastoreNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewDatastoreNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewDatastoreNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p DatastoreNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p DatastoreNamespace) Pointer() *DatastoreNamespace {
	return &p
}

func CastDatastoreNamespaces(data []interface{}) []DatastoreNamespace {
	v := make([]DatastoreNamespace, 0)
	for _, d := range data {
		v = append(v, NewDatastoreNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDatastoreNamespacesFromDict(data []DatastoreNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryNamespaceStatistics struct {
	Register *int64 `json:"register"`
}

func NewDictionaryNamespaceStatisticsFromJson(data string) DictionaryNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryNamespaceStatisticsFromDict(dict)
}

func NewDictionaryNamespaceStatisticsFromDict(data map[string]interface{}) DictionaryNamespaceStatistics {
	return DictionaryNamespaceStatistics{
		Register: core.CastInt64(data["register"]),
	}
}

func (p DictionaryNamespaceStatistics) ToDict() map[string]interface{} {

	var register *int64
	if p.Register != nil {
		register = p.Register
	}
	return map[string]interface{}{
		"register": register,
	}
}

func (p DictionaryNamespaceStatistics) Pointer() *DictionaryNamespaceStatistics {
	return &p
}

func CastDictionaryNamespaceStatisticses(data []interface{}) []DictionaryNamespaceStatistics {
	v := make([]DictionaryNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewDictionaryNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryNamespaceStatisticsesFromDict(data []DictionaryNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryNamespaceEntryByNameDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewDictionaryNamespaceEntryByNameDistributionStatisticsFromJson(data string) DictionaryNamespaceEntryByNameDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryNamespaceEntryByNameDistributionStatisticsFromDict(dict)
}

func NewDictionaryNamespaceEntryByNameDistributionStatisticsFromDict(data map[string]interface{}) DictionaryNamespaceEntryByNameDistributionStatistics {
	return DictionaryNamespaceEntryByNameDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p DictionaryNamespaceEntryByNameDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p DictionaryNamespaceEntryByNameDistributionStatistics) Pointer() *DictionaryNamespaceEntryByNameDistributionStatistics {
	return &p
}

func CastDictionaryNamespaceEntryByNameDistributionStatisticses(data []interface{}) []DictionaryNamespaceEntryByNameDistributionStatistics {
	v := make([]DictionaryNamespaceEntryByNameDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewDictionaryNamespaceEntryByNameDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryNamespaceEntryByNameDistributionStatisticsesFromDict(data []DictionaryNamespaceEntryByNameDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryNamespaceEntryByNameDistributionSegment struct {
	EntryModelName *string `json:"entryModelName"`
	Count          *int64  `json:"count"`
}

func NewDictionaryNamespaceEntryByNameDistributionSegmentFromJson(data string) DictionaryNamespaceEntryByNameDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryNamespaceEntryByNameDistributionSegmentFromDict(dict)
}

func NewDictionaryNamespaceEntryByNameDistributionSegmentFromDict(data map[string]interface{}) DictionaryNamespaceEntryByNameDistributionSegment {
	return DictionaryNamespaceEntryByNameDistributionSegment{
		EntryModelName: core.CastString(data["entryModelName"]),
		Count:          core.CastInt64(data["count"]),
	}
}

func (p DictionaryNamespaceEntryByNameDistributionSegment) ToDict() map[string]interface{} {

	var entryModelName *string
	if p.EntryModelName != nil {
		entryModelName = p.EntryModelName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"entryModelName": entryModelName,
		"count":          count,
	}
}

func (p DictionaryNamespaceEntryByNameDistributionSegment) Pointer() *DictionaryNamespaceEntryByNameDistributionSegment {
	return &p
}

func CastDictionaryNamespaceEntryByNameDistributionSegments(data []interface{}) []DictionaryNamespaceEntryByNameDistributionSegment {
	v := make([]DictionaryNamespaceEntryByNameDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewDictionaryNamespaceEntryByNameDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryNamespaceEntryByNameDistributionSegmentsFromDict(data []DictionaryNamespaceEntryByNameDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryNamespaceEntryByNameDistribution struct {
	Statistics   *DictionaryNamespaceEntryByNameDistributionStatistics `json:"statistics"`
	Distribution []DictionaryNamespaceEntryByNameDistributionSegment   `json:"distribution"`
}

func NewDictionaryNamespaceEntryByNameDistributionFromJson(data string) DictionaryNamespaceEntryByNameDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryNamespaceEntryByNameDistributionFromDict(dict)
}

func NewDictionaryNamespaceEntryByNameDistributionFromDict(data map[string]interface{}) DictionaryNamespaceEntryByNameDistribution {
	return DictionaryNamespaceEntryByNameDistribution{
		Statistics:   NewDictionaryNamespaceEntryByNameDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastDictionaryNamespaceEntryByNameDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p DictionaryNamespaceEntryByNameDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastDictionaryNamespaceEntryByNameDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p DictionaryNamespaceEntryByNameDistribution) Pointer() *DictionaryNamespaceEntryByNameDistribution {
	return &p
}

func CastDictionaryNamespaceEntryByNameDistributions(data []interface{}) []DictionaryNamespaceEntryByNameDistribution {
	v := make([]DictionaryNamespaceEntryByNameDistribution, 0)
	for _, d := range data {
		v = append(v, NewDictionaryNamespaceEntryByNameDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryNamespaceEntryByNameDistributionsFromDict(data []DictionaryNamespaceEntryByNameDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryNamespaceEntryByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewDictionaryNamespaceEntryByUserDistributionStatisticsFromJson(data string) DictionaryNamespaceEntryByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryNamespaceEntryByUserDistributionStatisticsFromDict(dict)
}

func NewDictionaryNamespaceEntryByUserDistributionStatisticsFromDict(data map[string]interface{}) DictionaryNamespaceEntryByUserDistributionStatistics {
	return DictionaryNamespaceEntryByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p DictionaryNamespaceEntryByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p DictionaryNamespaceEntryByUserDistributionStatistics) Pointer() *DictionaryNamespaceEntryByUserDistributionStatistics {
	return &p
}

func CastDictionaryNamespaceEntryByUserDistributionStatisticses(data []interface{}) []DictionaryNamespaceEntryByUserDistributionStatistics {
	v := make([]DictionaryNamespaceEntryByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewDictionaryNamespaceEntryByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryNamespaceEntryByUserDistributionStatisticsesFromDict(data []DictionaryNamespaceEntryByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryNamespaceEntryByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewDictionaryNamespaceEntryByUserDistributionSegmentFromJson(data string) DictionaryNamespaceEntryByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryNamespaceEntryByUserDistributionSegmentFromDict(dict)
}

func NewDictionaryNamespaceEntryByUserDistributionSegmentFromDict(data map[string]interface{}) DictionaryNamespaceEntryByUserDistributionSegment {
	return DictionaryNamespaceEntryByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p DictionaryNamespaceEntryByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p DictionaryNamespaceEntryByUserDistributionSegment) Pointer() *DictionaryNamespaceEntryByUserDistributionSegment {
	return &p
}

func CastDictionaryNamespaceEntryByUserDistributionSegments(data []interface{}) []DictionaryNamespaceEntryByUserDistributionSegment {
	v := make([]DictionaryNamespaceEntryByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewDictionaryNamespaceEntryByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryNamespaceEntryByUserDistributionSegmentsFromDict(data []DictionaryNamespaceEntryByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryNamespaceEntryByUserDistribution struct {
	Statistics   *DictionaryNamespaceEntryByUserDistributionStatistics `json:"statistics"`
	Distribution []DictionaryNamespaceEntryByUserDistributionSegment   `json:"distribution"`
}

func NewDictionaryNamespaceEntryByUserDistributionFromJson(data string) DictionaryNamespaceEntryByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryNamespaceEntryByUserDistributionFromDict(dict)
}

func NewDictionaryNamespaceEntryByUserDistributionFromDict(data map[string]interface{}) DictionaryNamespaceEntryByUserDistribution {
	return DictionaryNamespaceEntryByUserDistribution{
		Statistics:   NewDictionaryNamespaceEntryByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastDictionaryNamespaceEntryByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p DictionaryNamespaceEntryByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastDictionaryNamespaceEntryByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p DictionaryNamespaceEntryByUserDistribution) Pointer() *DictionaryNamespaceEntryByUserDistribution {
	return &p
}

func CastDictionaryNamespaceEntryByUserDistributions(data []interface{}) []DictionaryNamespaceEntryByUserDistribution {
	v := make([]DictionaryNamespaceEntryByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewDictionaryNamespaceEntryByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryNamespaceEntryByUserDistributionsFromDict(data []DictionaryNamespaceEntryByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryNamespaceDistributions struct {
	EntryByName *DictionaryNamespaceEntryByNameDistribution `json:"entryByName"`
	EntryByUser *DictionaryNamespaceEntryByUserDistribution `json:"entryByUser"`
}

func NewDictionaryNamespaceDistributionsFromJson(data string) DictionaryNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryNamespaceDistributionsFromDict(dict)
}

func NewDictionaryNamespaceDistributionsFromDict(data map[string]interface{}) DictionaryNamespaceDistributions {
	return DictionaryNamespaceDistributions{
		EntryByName: NewDictionaryNamespaceEntryByNameDistributionFromDict(core.CastMap(data["entryByName"])).Pointer(),
		EntryByUser: NewDictionaryNamespaceEntryByUserDistributionFromDict(core.CastMap(data["entryByUser"])).Pointer(),
	}
}

func (p DictionaryNamespaceDistributions) ToDict() map[string]interface{} {

	var entryByName map[string]interface{}
	if p.EntryByName != nil {
		entryByName = p.EntryByName.ToDict()
	}
	var entryByUser map[string]interface{}
	if p.EntryByUser != nil {
		entryByUser = p.EntryByUser.ToDict()
	}
	return map[string]interface{}{
		"entryByName": entryByName,
		"entryByUser": entryByUser,
	}
}

func (p DictionaryNamespaceDistributions) Pointer() *DictionaryNamespaceDistributions {
	return &p
}

func CastDictionaryNamespaceDistributionses(data []interface{}) []DictionaryNamespaceDistributions {
	v := make([]DictionaryNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewDictionaryNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryNamespaceDistributionsesFromDict(data []DictionaryNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryNamespace struct {
	NamespaceId   *string                           `json:"namespaceId"`
	Year          *int32                            `json:"year"`
	Month         *int32                            `json:"month"`
	Day           *int32                            `json:"day"`
	NamespaceName *string                           `json:"namespaceName"`
	Statistics    *DictionaryNamespaceStatistics    `json:"statistics"`
	Distributions *DictionaryNamespaceDistributions `json:"distributions"`
}

func NewDictionaryNamespaceFromJson(data string) DictionaryNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryNamespaceFromDict(dict)
}

func NewDictionaryNamespaceFromDict(data map[string]interface{}) DictionaryNamespace {
	return DictionaryNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewDictionaryNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewDictionaryNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p DictionaryNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p DictionaryNamespace) Pointer() *DictionaryNamespace {
	return &p
}

func CastDictionaryNamespaces(data []interface{}) []DictionaryNamespace {
	v := make([]DictionaryNamespace, 0)
	for _, d := range data {
		v = append(v, NewDictionaryNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryNamespacesFromDict(data []DictionaryNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DictionaryEntryModel struct {
	EntryModelModelId *string `json:"entryModelModelId"`
	EntryName         *string `json:"entryName"`
}

func NewDictionaryEntryModelFromJson(data string) DictionaryEntryModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDictionaryEntryModelFromDict(dict)
}

func NewDictionaryEntryModelFromDict(data map[string]interface{}) DictionaryEntryModel {
	return DictionaryEntryModel{
		EntryModelModelId: core.CastString(data["entryModelModelId"]),
		EntryName:         core.CastString(data["entryName"]),
	}
}

func (p DictionaryEntryModel) ToDict() map[string]interface{} {

	var entryModelModelId *string
	if p.EntryModelModelId != nil {
		entryModelModelId = p.EntryModelModelId
	}
	var entryName *string
	if p.EntryName != nil {
		entryName = p.EntryName
	}
	return map[string]interface{}{
		"entryModelModelId": entryModelModelId,
		"entryName":         entryName,
	}
}

func (p DictionaryEntryModel) Pointer() *DictionaryEntryModel {
	return &p
}

func CastDictionaryEntryModels(data []interface{}) []DictionaryEntryModel {
	v := make([]DictionaryEntryModel, 0)
	for _, d := range data {
		v = append(v, NewDictionaryEntryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDictionaryEntryModelsFromDict(data []DictionaryEntryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeRateModelStatistics struct {
	Exchange *int64 `json:"exchange"`
	Amount   *int64 `json:"amount"`
}

func NewExchangeRateModelStatisticsFromJson(data string) ExchangeRateModelStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeRateModelStatisticsFromDict(dict)
}

func NewExchangeRateModelStatisticsFromDict(data map[string]interface{}) ExchangeRateModelStatistics {
	return ExchangeRateModelStatistics{
		Exchange: core.CastInt64(data["exchange"]),
		Amount:   core.CastInt64(data["amount"]),
	}
}

func (p ExchangeRateModelStatistics) ToDict() map[string]interface{} {

	var exchange *int64
	if p.Exchange != nil {
		exchange = p.Exchange
	}
	var amount *int64
	if p.Amount != nil {
		amount = p.Amount
	}
	return map[string]interface{}{
		"exchange": exchange,
		"amount":   amount,
	}
}

func (p ExchangeRateModelStatistics) Pointer() *ExchangeRateModelStatistics {
	return &p
}

func CastExchangeRateModelStatisticses(data []interface{}) []ExchangeRateModelStatistics {
	v := make([]ExchangeRateModelStatistics, 0)
	for _, d := range data {
		v = append(v, NewExchangeRateModelStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeRateModelStatisticsesFromDict(data []ExchangeRateModelStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeRateModelAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExchangeRateModelAmountDistributionStatisticsFromJson(data string) ExchangeRateModelAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeRateModelAmountDistributionStatisticsFromDict(dict)
}

func NewExchangeRateModelAmountDistributionStatisticsFromDict(data map[string]interface{}) ExchangeRateModelAmountDistributionStatistics {
	return ExchangeRateModelAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExchangeRateModelAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExchangeRateModelAmountDistributionStatistics) Pointer() *ExchangeRateModelAmountDistributionStatistics {
	return &p
}

func CastExchangeRateModelAmountDistributionStatisticses(data []interface{}) []ExchangeRateModelAmountDistributionStatistics {
	v := make([]ExchangeRateModelAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExchangeRateModelAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeRateModelAmountDistributionStatisticsesFromDict(data []ExchangeRateModelAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeRateModelAmountDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewExchangeRateModelAmountDistributionSegmentFromJson(data string) ExchangeRateModelAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeRateModelAmountDistributionSegmentFromDict(dict)
}

func NewExchangeRateModelAmountDistributionSegmentFromDict(data map[string]interface{}) ExchangeRateModelAmountDistributionSegment {
	return ExchangeRateModelAmountDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ExchangeRateModelAmountDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ExchangeRateModelAmountDistributionSegment) Pointer() *ExchangeRateModelAmountDistributionSegment {
	return &p
}

func CastExchangeRateModelAmountDistributionSegments(data []interface{}) []ExchangeRateModelAmountDistributionSegment {
	v := make([]ExchangeRateModelAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExchangeRateModelAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeRateModelAmountDistributionSegmentsFromDict(data []ExchangeRateModelAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeRateModelAmountDistribution struct {
	Statistics   *ExchangeRateModelAmountDistributionStatistics `json:"statistics"`
	Distribution []ExchangeRateModelAmountDistributionSegment   `json:"distribution"`
}

func NewExchangeRateModelAmountDistributionFromJson(data string) ExchangeRateModelAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeRateModelAmountDistributionFromDict(dict)
}

func NewExchangeRateModelAmountDistributionFromDict(data map[string]interface{}) ExchangeRateModelAmountDistribution {
	return ExchangeRateModelAmountDistribution{
		Statistics:   NewExchangeRateModelAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExchangeRateModelAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExchangeRateModelAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExchangeRateModelAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExchangeRateModelAmountDistribution) Pointer() *ExchangeRateModelAmountDistribution {
	return &p
}

func CastExchangeRateModelAmountDistributions(data []interface{}) []ExchangeRateModelAmountDistribution {
	v := make([]ExchangeRateModelAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewExchangeRateModelAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeRateModelAmountDistributionsFromDict(data []ExchangeRateModelAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeRateModelDistributions struct {
	Amount *ExchangeRateModelAmountDistribution `json:"amount"`
}

func NewExchangeRateModelDistributionsFromJson(data string) ExchangeRateModelDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeRateModelDistributionsFromDict(dict)
}

func NewExchangeRateModelDistributionsFromDict(data map[string]interface{}) ExchangeRateModelDistributions {
	return ExchangeRateModelDistributions{
		Amount: NewExchangeRateModelAmountDistributionFromDict(core.CastMap(data["amount"])).Pointer(),
	}
}

func (p ExchangeRateModelDistributions) ToDict() map[string]interface{} {

	var amount map[string]interface{}
	if p.Amount != nil {
		amount = p.Amount.ToDict()
	}
	return map[string]interface{}{
		"amount": amount,
	}
}

func (p ExchangeRateModelDistributions) Pointer() *ExchangeRateModelDistributions {
	return &p
}

func CastExchangeRateModelDistributionses(data []interface{}) []ExchangeRateModelDistributions {
	v := make([]ExchangeRateModelDistributions, 0)
	for _, d := range data {
		v = append(v, NewExchangeRateModelDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeRateModelDistributionsesFromDict(data []ExchangeRateModelDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeRateModel struct {
	RateModelId   *string                         `json:"rateModelId"`
	RateName      *string                         `json:"rateName"`
	Statistics    *ExchangeRateModelStatistics    `json:"statistics"`
	Distributions *ExchangeRateModelDistributions `json:"distributions"`
}

func NewExchangeRateModelFromJson(data string) ExchangeRateModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeRateModelFromDict(dict)
}

func NewExchangeRateModelFromDict(data map[string]interface{}) ExchangeRateModel {
	return ExchangeRateModel{
		RateModelId:   core.CastString(data["rateModelId"]),
		RateName:      core.CastString(data["rateName"]),
		Statistics:    NewExchangeRateModelStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewExchangeRateModelDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p ExchangeRateModel) ToDict() map[string]interface{} {

	var rateModelId *string
	if p.RateModelId != nil {
		rateModelId = p.RateModelId
	}
	var rateName *string
	if p.RateName != nil {
		rateName = p.RateName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"rateModelId":   rateModelId,
		"rateName":      rateName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p ExchangeRateModel) Pointer() *ExchangeRateModel {
	return &p
}

func CastExchangeRateModels(data []interface{}) []ExchangeRateModel {
	v := make([]ExchangeRateModel, 0)
	for _, d := range data {
		v = append(v, NewExchangeRateModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeRateModelsFromDict(data []ExchangeRateModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceStatistics struct {
	Exchange *int64 `json:"exchange"`
	Amount   *int64 `json:"amount"`
}

func NewExchangeNamespaceStatisticsFromJson(data string) ExchangeNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceStatisticsFromDict(dict)
}

func NewExchangeNamespaceStatisticsFromDict(data map[string]interface{}) ExchangeNamespaceStatistics {
	return ExchangeNamespaceStatistics{
		Exchange: core.CastInt64(data["exchange"]),
		Amount:   core.CastInt64(data["amount"]),
	}
}

func (p ExchangeNamespaceStatistics) ToDict() map[string]interface{} {

	var exchange *int64
	if p.Exchange != nil {
		exchange = p.Exchange
	}
	var amount *int64
	if p.Amount != nil {
		amount = p.Amount
	}
	return map[string]interface{}{
		"exchange": exchange,
		"amount":   amount,
	}
}

func (p ExchangeNamespaceStatistics) Pointer() *ExchangeNamespaceStatistics {
	return &p
}

func CastExchangeNamespaceStatisticses(data []interface{}) []ExchangeNamespaceStatistics {
	v := make([]ExchangeNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceStatisticsesFromDict(data []ExchangeNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExchangeNamespaceExchangeDistributionStatisticsFromJson(data string) ExchangeNamespaceExchangeDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeDistributionStatisticsFromDict(dict)
}

func NewExchangeNamespaceExchangeDistributionStatisticsFromDict(data map[string]interface{}) ExchangeNamespaceExchangeDistributionStatistics {
	return ExchangeNamespaceExchangeDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExchangeNamespaceExchangeDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExchangeNamespaceExchangeDistributionStatistics) Pointer() *ExchangeNamespaceExchangeDistributionStatistics {
	return &p
}

func CastExchangeNamespaceExchangeDistributionStatisticses(data []interface{}) []ExchangeNamespaceExchangeDistributionStatistics {
	v := make([]ExchangeNamespaceExchangeDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeDistributionStatisticsesFromDict(data []ExchangeNamespaceExchangeDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeDistributionSegment struct {
	RateName *string `json:"rateName"`
	Count    *int64  `json:"count"`
}

func NewExchangeNamespaceExchangeDistributionSegmentFromJson(data string) ExchangeNamespaceExchangeDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeDistributionSegmentFromDict(dict)
}

func NewExchangeNamespaceExchangeDistributionSegmentFromDict(data map[string]interface{}) ExchangeNamespaceExchangeDistributionSegment {
	return ExchangeNamespaceExchangeDistributionSegment{
		RateName: core.CastString(data["rateName"]),
		Count:    core.CastInt64(data["count"]),
	}
}

func (p ExchangeNamespaceExchangeDistributionSegment) ToDict() map[string]interface{} {

	var rateName *string
	if p.RateName != nil {
		rateName = p.RateName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"rateName": rateName,
		"count":    count,
	}
}

func (p ExchangeNamespaceExchangeDistributionSegment) Pointer() *ExchangeNamespaceExchangeDistributionSegment {
	return &p
}

func CastExchangeNamespaceExchangeDistributionSegments(data []interface{}) []ExchangeNamespaceExchangeDistributionSegment {
	v := make([]ExchangeNamespaceExchangeDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeDistributionSegmentsFromDict(data []ExchangeNamespaceExchangeDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeDistribution struct {
	Statistics   *ExchangeNamespaceExchangeDistributionStatistics `json:"statistics"`
	Distribution []ExchangeNamespaceExchangeDistributionSegment   `json:"distribution"`
}

func NewExchangeNamespaceExchangeDistributionFromJson(data string) ExchangeNamespaceExchangeDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeDistributionFromDict(dict)
}

func NewExchangeNamespaceExchangeDistributionFromDict(data map[string]interface{}) ExchangeNamespaceExchangeDistribution {
	return ExchangeNamespaceExchangeDistribution{
		Statistics:   NewExchangeNamespaceExchangeDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExchangeNamespaceExchangeDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExchangeNamespaceExchangeDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExchangeNamespaceExchangeDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExchangeNamespaceExchangeDistribution) Pointer() *ExchangeNamespaceExchangeDistribution {
	return &p
}

func CastExchangeNamespaceExchangeDistributions(data []interface{}) []ExchangeNamespaceExchangeDistribution {
	v := make([]ExchangeNamespaceExchangeDistribution, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeDistributionsFromDict(data []ExchangeNamespaceExchangeDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExchangeNamespaceExchangeAmountDistributionStatisticsFromJson(data string) ExchangeNamespaceExchangeAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeAmountDistributionStatisticsFromDict(dict)
}

func NewExchangeNamespaceExchangeAmountDistributionStatisticsFromDict(data map[string]interface{}) ExchangeNamespaceExchangeAmountDistributionStatistics {
	return ExchangeNamespaceExchangeAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExchangeNamespaceExchangeAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExchangeNamespaceExchangeAmountDistributionStatistics) Pointer() *ExchangeNamespaceExchangeAmountDistributionStatistics {
	return &p
}

func CastExchangeNamespaceExchangeAmountDistributionStatisticses(data []interface{}) []ExchangeNamespaceExchangeAmountDistributionStatistics {
	v := make([]ExchangeNamespaceExchangeAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeAmountDistributionStatisticsesFromDict(data []ExchangeNamespaceExchangeAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeAmountDistributionSegment struct {
	RateName *string `json:"rateName"`
	Sum      *int64  `json:"sum"`
}

func NewExchangeNamespaceExchangeAmountDistributionSegmentFromJson(data string) ExchangeNamespaceExchangeAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeAmountDistributionSegmentFromDict(dict)
}

func NewExchangeNamespaceExchangeAmountDistributionSegmentFromDict(data map[string]interface{}) ExchangeNamespaceExchangeAmountDistributionSegment {
	return ExchangeNamespaceExchangeAmountDistributionSegment{
		RateName: core.CastString(data["rateName"]),
		Sum:      core.CastInt64(data["sum"]),
	}
}

func (p ExchangeNamespaceExchangeAmountDistributionSegment) ToDict() map[string]interface{} {

	var rateName *string
	if p.RateName != nil {
		rateName = p.RateName
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"rateName": rateName,
		"sum":      sum,
	}
}

func (p ExchangeNamespaceExchangeAmountDistributionSegment) Pointer() *ExchangeNamespaceExchangeAmountDistributionSegment {
	return &p
}

func CastExchangeNamespaceExchangeAmountDistributionSegments(data []interface{}) []ExchangeNamespaceExchangeAmountDistributionSegment {
	v := make([]ExchangeNamespaceExchangeAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeAmountDistributionSegmentsFromDict(data []ExchangeNamespaceExchangeAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeAmountDistribution struct {
	Statistics   *ExchangeNamespaceExchangeAmountDistributionStatistics `json:"statistics"`
	Distribution []ExchangeNamespaceExchangeAmountDistributionSegment   `json:"distribution"`
}

func NewExchangeNamespaceExchangeAmountDistributionFromJson(data string) ExchangeNamespaceExchangeAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeAmountDistributionFromDict(dict)
}

func NewExchangeNamespaceExchangeAmountDistributionFromDict(data map[string]interface{}) ExchangeNamespaceExchangeAmountDistribution {
	return ExchangeNamespaceExchangeAmountDistribution{
		Statistics:   NewExchangeNamespaceExchangeAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExchangeNamespaceExchangeAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExchangeNamespaceExchangeAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExchangeNamespaceExchangeAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExchangeNamespaceExchangeAmountDistribution) Pointer() *ExchangeNamespaceExchangeAmountDistribution {
	return &p
}

func CastExchangeNamespaceExchangeAmountDistributions(data []interface{}) []ExchangeNamespaceExchangeAmountDistribution {
	v := make([]ExchangeNamespaceExchangeAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeAmountDistributionsFromDict(data []ExchangeNamespaceExchangeAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExchangeNamespaceExchangeByUserDistributionStatisticsFromJson(data string) ExchangeNamespaceExchangeByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeByUserDistributionStatisticsFromDict(dict)
}

func NewExchangeNamespaceExchangeByUserDistributionStatisticsFromDict(data map[string]interface{}) ExchangeNamespaceExchangeByUserDistributionStatistics {
	return ExchangeNamespaceExchangeByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExchangeNamespaceExchangeByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExchangeNamespaceExchangeByUserDistributionStatistics) Pointer() *ExchangeNamespaceExchangeByUserDistributionStatistics {
	return &p
}

func CastExchangeNamespaceExchangeByUserDistributionStatisticses(data []interface{}) []ExchangeNamespaceExchangeByUserDistributionStatistics {
	v := make([]ExchangeNamespaceExchangeByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeByUserDistributionStatisticsesFromDict(data []ExchangeNamespaceExchangeByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewExchangeNamespaceExchangeByUserDistributionSegmentFromJson(data string) ExchangeNamespaceExchangeByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeByUserDistributionSegmentFromDict(dict)
}

func NewExchangeNamespaceExchangeByUserDistributionSegmentFromDict(data map[string]interface{}) ExchangeNamespaceExchangeByUserDistributionSegment {
	return ExchangeNamespaceExchangeByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ExchangeNamespaceExchangeByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ExchangeNamespaceExchangeByUserDistributionSegment) Pointer() *ExchangeNamespaceExchangeByUserDistributionSegment {
	return &p
}

func CastExchangeNamespaceExchangeByUserDistributionSegments(data []interface{}) []ExchangeNamespaceExchangeByUserDistributionSegment {
	v := make([]ExchangeNamespaceExchangeByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeByUserDistributionSegmentsFromDict(data []ExchangeNamespaceExchangeByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeByUserDistribution struct {
	Statistics   *ExchangeNamespaceExchangeByUserDistributionStatistics `json:"statistics"`
	Distribution []ExchangeNamespaceExchangeByUserDistributionSegment   `json:"distribution"`
}

func NewExchangeNamespaceExchangeByUserDistributionFromJson(data string) ExchangeNamespaceExchangeByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeByUserDistributionFromDict(dict)
}

func NewExchangeNamespaceExchangeByUserDistributionFromDict(data map[string]interface{}) ExchangeNamespaceExchangeByUserDistribution {
	return ExchangeNamespaceExchangeByUserDistribution{
		Statistics:   NewExchangeNamespaceExchangeByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExchangeNamespaceExchangeByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExchangeNamespaceExchangeByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExchangeNamespaceExchangeByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExchangeNamespaceExchangeByUserDistribution) Pointer() *ExchangeNamespaceExchangeByUserDistribution {
	return &p
}

func CastExchangeNamespaceExchangeByUserDistributions(data []interface{}) []ExchangeNamespaceExchangeByUserDistribution {
	v := make([]ExchangeNamespaceExchangeByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeByUserDistributionsFromDict(data []ExchangeNamespaceExchangeByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeAmountByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExchangeNamespaceExchangeAmountByUserDistributionStatisticsFromJson(data string) ExchangeNamespaceExchangeAmountByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeAmountByUserDistributionStatisticsFromDict(dict)
}

func NewExchangeNamespaceExchangeAmountByUserDistributionStatisticsFromDict(data map[string]interface{}) ExchangeNamespaceExchangeAmountByUserDistributionStatistics {
	return ExchangeNamespaceExchangeAmountByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExchangeNamespaceExchangeAmountByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExchangeNamespaceExchangeAmountByUserDistributionStatistics) Pointer() *ExchangeNamespaceExchangeAmountByUserDistributionStatistics {
	return &p
}

func CastExchangeNamespaceExchangeAmountByUserDistributionStatisticses(data []interface{}) []ExchangeNamespaceExchangeAmountByUserDistributionStatistics {
	v := make([]ExchangeNamespaceExchangeAmountByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeAmountByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeAmountByUserDistributionStatisticsesFromDict(data []ExchangeNamespaceExchangeAmountByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeAmountByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewExchangeNamespaceExchangeAmountByUserDistributionSegmentFromJson(data string) ExchangeNamespaceExchangeAmountByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeAmountByUserDistributionSegmentFromDict(dict)
}

func NewExchangeNamespaceExchangeAmountByUserDistributionSegmentFromDict(data map[string]interface{}) ExchangeNamespaceExchangeAmountByUserDistributionSegment {
	return ExchangeNamespaceExchangeAmountByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ExchangeNamespaceExchangeAmountByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ExchangeNamespaceExchangeAmountByUserDistributionSegment) Pointer() *ExchangeNamespaceExchangeAmountByUserDistributionSegment {
	return &p
}

func CastExchangeNamespaceExchangeAmountByUserDistributionSegments(data []interface{}) []ExchangeNamespaceExchangeAmountByUserDistributionSegment {
	v := make([]ExchangeNamespaceExchangeAmountByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeAmountByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeAmountByUserDistributionSegmentsFromDict(data []ExchangeNamespaceExchangeAmountByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceExchangeAmountByUserDistribution struct {
	Statistics   *ExchangeNamespaceExchangeAmountByUserDistributionStatistics `json:"statistics"`
	Distribution []ExchangeNamespaceExchangeAmountByUserDistributionSegment   `json:"distribution"`
}

func NewExchangeNamespaceExchangeAmountByUserDistributionFromJson(data string) ExchangeNamespaceExchangeAmountByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceExchangeAmountByUserDistributionFromDict(dict)
}

func NewExchangeNamespaceExchangeAmountByUserDistributionFromDict(data map[string]interface{}) ExchangeNamespaceExchangeAmountByUserDistribution {
	return ExchangeNamespaceExchangeAmountByUserDistribution{
		Statistics:   NewExchangeNamespaceExchangeAmountByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExchangeNamespaceExchangeAmountByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExchangeNamespaceExchangeAmountByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExchangeNamespaceExchangeAmountByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExchangeNamespaceExchangeAmountByUserDistribution) Pointer() *ExchangeNamespaceExchangeAmountByUserDistribution {
	return &p
}

func CastExchangeNamespaceExchangeAmountByUserDistributions(data []interface{}) []ExchangeNamespaceExchangeAmountByUserDistribution {
	v := make([]ExchangeNamespaceExchangeAmountByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceExchangeAmountByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceExchangeAmountByUserDistributionsFromDict(data []ExchangeNamespaceExchangeAmountByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespaceDistributions struct {
	Exchange             *ExchangeNamespaceExchangeDistribution             `json:"exchange"`
	ExchangeAmount       *ExchangeNamespaceExchangeAmountDistribution       `json:"exchangeAmount"`
	ExchangeByUser       *ExchangeNamespaceExchangeByUserDistribution       `json:"exchangeByUser"`
	ExchangeAmountByUser *ExchangeNamespaceExchangeAmountByUserDistribution `json:"exchangeAmountByUser"`
}

func NewExchangeNamespaceDistributionsFromJson(data string) ExchangeNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceDistributionsFromDict(dict)
}

func NewExchangeNamespaceDistributionsFromDict(data map[string]interface{}) ExchangeNamespaceDistributions {
	return ExchangeNamespaceDistributions{
		Exchange:             NewExchangeNamespaceExchangeDistributionFromDict(core.CastMap(data["exchange"])).Pointer(),
		ExchangeAmount:       NewExchangeNamespaceExchangeAmountDistributionFromDict(core.CastMap(data["exchangeAmount"])).Pointer(),
		ExchangeByUser:       NewExchangeNamespaceExchangeByUserDistributionFromDict(core.CastMap(data["exchangeByUser"])).Pointer(),
		ExchangeAmountByUser: NewExchangeNamespaceExchangeAmountByUserDistributionFromDict(core.CastMap(data["exchangeAmountByUser"])).Pointer(),
	}
}

func (p ExchangeNamespaceDistributions) ToDict() map[string]interface{} {

	var exchange map[string]interface{}
	if p.Exchange != nil {
		exchange = p.Exchange.ToDict()
	}
	var exchangeAmount map[string]interface{}
	if p.ExchangeAmount != nil {
		exchangeAmount = p.ExchangeAmount.ToDict()
	}
	var exchangeByUser map[string]interface{}
	if p.ExchangeByUser != nil {
		exchangeByUser = p.ExchangeByUser.ToDict()
	}
	var exchangeAmountByUser map[string]interface{}
	if p.ExchangeAmountByUser != nil {
		exchangeAmountByUser = p.ExchangeAmountByUser.ToDict()
	}
	return map[string]interface{}{
		"exchange":             exchange,
		"exchangeAmount":       exchangeAmount,
		"exchangeByUser":       exchangeByUser,
		"exchangeAmountByUser": exchangeAmountByUser,
	}
}

func (p ExchangeNamespaceDistributions) Pointer() *ExchangeNamespaceDistributions {
	return &p
}

func CastExchangeNamespaceDistributionses(data []interface{}) []ExchangeNamespaceDistributions {
	v := make([]ExchangeNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespaceDistributionsesFromDict(data []ExchangeNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExchangeNamespace struct {
	NamespaceId   *string                         `json:"namespaceId"`
	Year          *int32                          `json:"year"`
	Month         *int32                          `json:"month"`
	Day           *int32                          `json:"day"`
	NamespaceName *string                         `json:"namespaceName"`
	Statistics    *ExchangeNamespaceStatistics    `json:"statistics"`
	Distributions *ExchangeNamespaceDistributions `json:"distributions"`
	RateModels    []ExchangeRateModel             `json:"rateModels"`
}

func NewExchangeNamespaceFromJson(data string) ExchangeNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeNamespaceFromDict(dict)
}

func NewExchangeNamespaceFromDict(data map[string]interface{}) ExchangeNamespace {
	return ExchangeNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewExchangeNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewExchangeNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		RateModels:    CastExchangeRateModels(core.CastArray(data["rateModels"])),
	}
}

func (p ExchangeNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var rateModels []interface{}
	if p.RateModels != nil {
		rateModels = CastExchangeRateModelsFromDict(
			p.RateModels,
		)
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
		"rateModels":    rateModels,
	}
}

func (p ExchangeNamespace) Pointer() *ExchangeNamespace {
	return &p
}

func CastExchangeNamespaces(data []interface{}) []ExchangeNamespace {
	v := make([]ExchangeNamespace, 0)
	for _, d := range data {
		v = append(v, NewExchangeNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExchangeNamespacesFromDict(data []ExchangeNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceStatusRankDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExperienceStatusRankDistributionStatisticsFromJson(data string) ExperienceStatusRankDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceStatusRankDistributionStatisticsFromDict(dict)
}

func NewExperienceStatusRankDistributionStatisticsFromDict(data map[string]interface{}) ExperienceStatusRankDistributionStatistics {
	return ExperienceStatusRankDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExperienceStatusRankDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExperienceStatusRankDistributionStatistics) Pointer() *ExperienceStatusRankDistributionStatistics {
	return &p
}

func CastExperienceStatusRankDistributionStatisticses(data []interface{}) []ExperienceStatusRankDistributionStatistics {
	v := make([]ExperienceStatusRankDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceStatusRankDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceStatusRankDistributionStatisticsesFromDict(data []ExperienceStatusRankDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceStatusRankDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewExperienceStatusRankDistributionSegmentFromJson(data string) ExperienceStatusRankDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceStatusRankDistributionSegmentFromDict(dict)
}

func NewExperienceStatusRankDistributionSegmentFromDict(data map[string]interface{}) ExperienceStatusRankDistributionSegment {
	return ExperienceStatusRankDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ExperienceStatusRankDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ExperienceStatusRankDistributionSegment) Pointer() *ExperienceStatusRankDistributionSegment {
	return &p
}

func CastExperienceStatusRankDistributionSegments(data []interface{}) []ExperienceStatusRankDistributionSegment {
	v := make([]ExperienceStatusRankDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExperienceStatusRankDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceStatusRankDistributionSegmentsFromDict(data []ExperienceStatusRankDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceStatusRankDistribution struct {
	Statistics   *ExperienceStatusRankDistributionStatistics `json:"statistics"`
	Distribution []ExperienceStatusRankDistributionSegment   `json:"distribution"`
}

func NewExperienceStatusRankDistributionFromJson(data string) ExperienceStatusRankDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceStatusRankDistributionFromDict(dict)
}

func NewExperienceStatusRankDistributionFromDict(data map[string]interface{}) ExperienceStatusRankDistribution {
	return ExperienceStatusRankDistribution{
		Statistics:   NewExperienceStatusRankDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExperienceStatusRankDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExperienceStatusRankDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExperienceStatusRankDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExperienceStatusRankDistribution) Pointer() *ExperienceStatusRankDistribution {
	return &p
}

func CastExperienceStatusRankDistributions(data []interface{}) []ExperienceStatusRankDistribution {
	v := make([]ExperienceStatusRankDistribution, 0)
	for _, d := range data {
		v = append(v, NewExperienceStatusRankDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceStatusRankDistributionsFromDict(data []ExperienceStatusRankDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceStatusRankCapDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExperienceStatusRankCapDistributionStatisticsFromJson(data string) ExperienceStatusRankCapDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceStatusRankCapDistributionStatisticsFromDict(dict)
}

func NewExperienceStatusRankCapDistributionStatisticsFromDict(data map[string]interface{}) ExperienceStatusRankCapDistributionStatistics {
	return ExperienceStatusRankCapDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExperienceStatusRankCapDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExperienceStatusRankCapDistributionStatistics) Pointer() *ExperienceStatusRankCapDistributionStatistics {
	return &p
}

func CastExperienceStatusRankCapDistributionStatisticses(data []interface{}) []ExperienceStatusRankCapDistributionStatistics {
	v := make([]ExperienceStatusRankCapDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceStatusRankCapDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceStatusRankCapDistributionStatisticsesFromDict(data []ExperienceStatusRankCapDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceStatusRankCapDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewExperienceStatusRankCapDistributionSegmentFromJson(data string) ExperienceStatusRankCapDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceStatusRankCapDistributionSegmentFromDict(dict)
}

func NewExperienceStatusRankCapDistributionSegmentFromDict(data map[string]interface{}) ExperienceStatusRankCapDistributionSegment {
	return ExperienceStatusRankCapDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ExperienceStatusRankCapDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ExperienceStatusRankCapDistributionSegment) Pointer() *ExperienceStatusRankCapDistributionSegment {
	return &p
}

func CastExperienceStatusRankCapDistributionSegments(data []interface{}) []ExperienceStatusRankCapDistributionSegment {
	v := make([]ExperienceStatusRankCapDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExperienceStatusRankCapDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceStatusRankCapDistributionSegmentsFromDict(data []ExperienceStatusRankCapDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceStatusRankCapDistribution struct {
	Statistics   *ExperienceStatusRankCapDistributionStatistics `json:"statistics"`
	Distribution []ExperienceStatusRankCapDistributionSegment   `json:"distribution"`
}

func NewExperienceStatusRankCapDistributionFromJson(data string) ExperienceStatusRankCapDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceStatusRankCapDistributionFromDict(dict)
}

func NewExperienceStatusRankCapDistributionFromDict(data map[string]interface{}) ExperienceStatusRankCapDistribution {
	return ExperienceStatusRankCapDistribution{
		Statistics:   NewExperienceStatusRankCapDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExperienceStatusRankCapDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExperienceStatusRankCapDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExperienceStatusRankCapDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExperienceStatusRankCapDistribution) Pointer() *ExperienceStatusRankCapDistribution {
	return &p
}

func CastExperienceStatusRankCapDistributions(data []interface{}) []ExperienceStatusRankCapDistribution {
	v := make([]ExperienceStatusRankCapDistribution, 0)
	for _, d := range data {
		v = append(v, NewExperienceStatusRankCapDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceStatusRankCapDistributionsFromDict(data []ExperienceStatusRankCapDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceStatusDistributions struct {
	Rank    *ExperienceStatusRankDistribution    `json:"rank"`
	RankCap *ExperienceStatusRankCapDistribution `json:"rankCap"`
}

func NewExperienceStatusDistributionsFromJson(data string) ExperienceStatusDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceStatusDistributionsFromDict(dict)
}

func NewExperienceStatusDistributionsFromDict(data map[string]interface{}) ExperienceStatusDistributions {
	return ExperienceStatusDistributions{
		Rank:    NewExperienceStatusRankDistributionFromDict(core.CastMap(data["rank"])).Pointer(),
		RankCap: NewExperienceStatusRankCapDistributionFromDict(core.CastMap(data["rankCap"])).Pointer(),
	}
}

func (p ExperienceStatusDistributions) ToDict() map[string]interface{} {

	var rank map[string]interface{}
	if p.Rank != nil {
		rank = p.Rank.ToDict()
	}
	var rankCap map[string]interface{}
	if p.RankCap != nil {
		rankCap = p.RankCap.ToDict()
	}
	return map[string]interface{}{
		"rank":    rank,
		"rankCap": rankCap,
	}
}

func (p ExperienceStatusDistributions) Pointer() *ExperienceStatusDistributions {
	return &p
}

func CastExperienceStatusDistributionses(data []interface{}) []ExperienceStatusDistributions {
	v := make([]ExperienceStatusDistributions, 0)
	for _, d := range data {
		v = append(v, NewExperienceStatusDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceStatusDistributionsesFromDict(data []ExperienceStatusDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceStatus struct {
	StatusId       *string                        `json:"statusId"`
	ExperienceName *string                        `json:"experienceName"`
	PropertyId     *string                        `json:"propertyId"`
	Distributions  *ExperienceStatusDistributions `json:"distributions"`
}

func NewExperienceStatusFromJson(data string) ExperienceStatus {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceStatusFromDict(dict)
}

func NewExperienceStatusFromDict(data map[string]interface{}) ExperienceStatus {
	return ExperienceStatus{
		StatusId:       core.CastString(data["statusId"]),
		ExperienceName: core.CastString(data["experienceName"]),
		PropertyId:     core.CastString(data["propertyId"]),
		Distributions:  NewExperienceStatusDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p ExperienceStatus) ToDict() map[string]interface{} {

	var statusId *string
	if p.StatusId != nil {
		statusId = p.StatusId
	}
	var experienceName *string
	if p.ExperienceName != nil {
		experienceName = p.ExperienceName
	}
	var propertyId *string
	if p.PropertyId != nil {
		propertyId = p.PropertyId
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"statusId":       statusId,
		"experienceName": experienceName,
		"propertyId":     propertyId,
		"distributions":  distributions,
	}
}

func (p ExperienceStatus) Pointer() *ExperienceStatus {
	return &p
}

func CastExperienceStatuses(data []interface{}) []ExperienceStatus {
	v := make([]ExperienceStatus, 0)
	for _, d := range data {
		v = append(v, NewExperienceStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceStatusesFromDict(data []ExperienceStatus) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelStatistics struct {
	Experience       *int64 `json:"experience"`
	ExperienceAmount *int64 `json:"experienceAmount"`
	RankCap          *int64 `json:"rankCap"`
	RankCapAmount    *int64 `json:"rankCapAmount"`
}

func NewExperienceExperienceModelStatisticsFromJson(data string) ExperienceExperienceModelStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelStatisticsFromDict(dict)
}

func NewExperienceExperienceModelStatisticsFromDict(data map[string]interface{}) ExperienceExperienceModelStatistics {
	return ExperienceExperienceModelStatistics{
		Experience:       core.CastInt64(data["experience"]),
		ExperienceAmount: core.CastInt64(data["experienceAmount"]),
		RankCap:          core.CastInt64(data["rankCap"]),
		RankCapAmount:    core.CastInt64(data["rankCapAmount"]),
	}
}

func (p ExperienceExperienceModelStatistics) ToDict() map[string]interface{} {

	var experience *int64
	if p.Experience != nil {
		experience = p.Experience
	}
	var experienceAmount *int64
	if p.ExperienceAmount != nil {
		experienceAmount = p.ExperienceAmount
	}
	var rankCap *int64
	if p.RankCap != nil {
		rankCap = p.RankCap
	}
	var rankCapAmount *int64
	if p.RankCapAmount != nil {
		rankCapAmount = p.RankCapAmount
	}
	return map[string]interface{}{
		"experience":       experience,
		"experienceAmount": experienceAmount,
		"rankCap":          rankCap,
		"rankCapAmount":    rankCapAmount,
	}
}

func (p ExperienceExperienceModelStatistics) Pointer() *ExperienceExperienceModelStatistics {
	return &p
}

func CastExperienceExperienceModelStatisticses(data []interface{}) []ExperienceExperienceModelStatistics {
	v := make([]ExperienceExperienceModelStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelStatisticsesFromDict(data []ExperienceExperienceModelStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddExperienceByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExperienceExperienceModelAddExperienceByUserDistributionStatisticsFromJson(data string) ExperienceExperienceModelAddExperienceByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddExperienceByUserDistributionStatisticsFromDict(dict)
}

func NewExperienceExperienceModelAddExperienceByUserDistributionStatisticsFromDict(data map[string]interface{}) ExperienceExperienceModelAddExperienceByUserDistributionStatistics {
	return ExperienceExperienceModelAddExperienceByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExperienceExperienceModelAddExperienceByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExperienceExperienceModelAddExperienceByUserDistributionStatistics) Pointer() *ExperienceExperienceModelAddExperienceByUserDistributionStatistics {
	return &p
}

func CastExperienceExperienceModelAddExperienceByUserDistributionStatisticses(data []interface{}) []ExperienceExperienceModelAddExperienceByUserDistributionStatistics {
	v := make([]ExperienceExperienceModelAddExperienceByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddExperienceByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddExperienceByUserDistributionStatisticsesFromDict(data []ExperienceExperienceModelAddExperienceByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddExperienceByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewExperienceExperienceModelAddExperienceByUserDistributionSegmentFromJson(data string) ExperienceExperienceModelAddExperienceByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddExperienceByUserDistributionSegmentFromDict(dict)
}

func NewExperienceExperienceModelAddExperienceByUserDistributionSegmentFromDict(data map[string]interface{}) ExperienceExperienceModelAddExperienceByUserDistributionSegment {
	return ExperienceExperienceModelAddExperienceByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ExperienceExperienceModelAddExperienceByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ExperienceExperienceModelAddExperienceByUserDistributionSegment) Pointer() *ExperienceExperienceModelAddExperienceByUserDistributionSegment {
	return &p
}

func CastExperienceExperienceModelAddExperienceByUserDistributionSegments(data []interface{}) []ExperienceExperienceModelAddExperienceByUserDistributionSegment {
	v := make([]ExperienceExperienceModelAddExperienceByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddExperienceByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddExperienceByUserDistributionSegmentsFromDict(data []ExperienceExperienceModelAddExperienceByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddExperienceByUserDistribution struct {
	Statistics   *ExperienceExperienceModelAddExperienceByUserDistributionStatistics `json:"statistics"`
	Distribution []ExperienceExperienceModelAddExperienceByUserDistributionSegment   `json:"distribution"`
}

func NewExperienceExperienceModelAddExperienceByUserDistributionFromJson(data string) ExperienceExperienceModelAddExperienceByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddExperienceByUserDistributionFromDict(dict)
}

func NewExperienceExperienceModelAddExperienceByUserDistributionFromDict(data map[string]interface{}) ExperienceExperienceModelAddExperienceByUserDistribution {
	return ExperienceExperienceModelAddExperienceByUserDistribution{
		Statistics:   NewExperienceExperienceModelAddExperienceByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExperienceExperienceModelAddExperienceByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExperienceExperienceModelAddExperienceByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExperienceExperienceModelAddExperienceByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExperienceExperienceModelAddExperienceByUserDistribution) Pointer() *ExperienceExperienceModelAddExperienceByUserDistribution {
	return &p
}

func CastExperienceExperienceModelAddExperienceByUserDistributions(data []interface{}) []ExperienceExperienceModelAddExperienceByUserDistribution {
	v := make([]ExperienceExperienceModelAddExperienceByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddExperienceByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddExperienceByUserDistributionsFromDict(data []ExperienceExperienceModelAddExperienceByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExperienceExperienceModelAddExperienceAmountByUserDistributionStatisticsFromJson(data string) ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddExperienceAmountByUserDistributionStatisticsFromDict(dict)
}

func NewExperienceExperienceModelAddExperienceAmountByUserDistributionStatisticsFromDict(data map[string]interface{}) ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics {
	return ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics) Pointer() *ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics {
	return &p
}

func CastExperienceExperienceModelAddExperienceAmountByUserDistributionStatisticses(data []interface{}) []ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics {
	v := make([]ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddExperienceAmountByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddExperienceAmountByUserDistributionStatisticsesFromDict(data []ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewExperienceExperienceModelAddExperienceAmountByUserDistributionSegmentFromJson(data string) ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddExperienceAmountByUserDistributionSegmentFromDict(dict)
}

func NewExperienceExperienceModelAddExperienceAmountByUserDistributionSegmentFromDict(data map[string]interface{}) ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment {
	return ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment) Pointer() *ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment {
	return &p
}

func CastExperienceExperienceModelAddExperienceAmountByUserDistributionSegments(data []interface{}) []ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment {
	v := make([]ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddExperienceAmountByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddExperienceAmountByUserDistributionSegmentsFromDict(data []ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddExperienceAmountByUserDistribution struct {
	Statistics   *ExperienceExperienceModelAddExperienceAmountByUserDistributionStatistics `json:"statistics"`
	Distribution []ExperienceExperienceModelAddExperienceAmountByUserDistributionSegment   `json:"distribution"`
}

func NewExperienceExperienceModelAddExperienceAmountByUserDistributionFromJson(data string) ExperienceExperienceModelAddExperienceAmountByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddExperienceAmountByUserDistributionFromDict(dict)
}

func NewExperienceExperienceModelAddExperienceAmountByUserDistributionFromDict(data map[string]interface{}) ExperienceExperienceModelAddExperienceAmountByUserDistribution {
	return ExperienceExperienceModelAddExperienceAmountByUserDistribution{
		Statistics:   NewExperienceExperienceModelAddExperienceAmountByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExperienceExperienceModelAddExperienceAmountByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExperienceExperienceModelAddExperienceAmountByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExperienceExperienceModelAddExperienceAmountByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExperienceExperienceModelAddExperienceAmountByUserDistribution) Pointer() *ExperienceExperienceModelAddExperienceAmountByUserDistribution {
	return &p
}

func CastExperienceExperienceModelAddExperienceAmountByUserDistributions(data []interface{}) []ExperienceExperienceModelAddExperienceAmountByUserDistribution {
	v := make([]ExperienceExperienceModelAddExperienceAmountByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddExperienceAmountByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddExperienceAmountByUserDistributionsFromDict(data []ExperienceExperienceModelAddExperienceAmountByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddRankCapByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExperienceExperienceModelAddRankCapByUserDistributionStatisticsFromJson(data string) ExperienceExperienceModelAddRankCapByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddRankCapByUserDistributionStatisticsFromDict(dict)
}

func NewExperienceExperienceModelAddRankCapByUserDistributionStatisticsFromDict(data map[string]interface{}) ExperienceExperienceModelAddRankCapByUserDistributionStatistics {
	return ExperienceExperienceModelAddRankCapByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExperienceExperienceModelAddRankCapByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExperienceExperienceModelAddRankCapByUserDistributionStatistics) Pointer() *ExperienceExperienceModelAddRankCapByUserDistributionStatistics {
	return &p
}

func CastExperienceExperienceModelAddRankCapByUserDistributionStatisticses(data []interface{}) []ExperienceExperienceModelAddRankCapByUserDistributionStatistics {
	v := make([]ExperienceExperienceModelAddRankCapByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddRankCapByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddRankCapByUserDistributionStatisticsesFromDict(data []ExperienceExperienceModelAddRankCapByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddRankCapByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewExperienceExperienceModelAddRankCapByUserDistributionSegmentFromJson(data string) ExperienceExperienceModelAddRankCapByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddRankCapByUserDistributionSegmentFromDict(dict)
}

func NewExperienceExperienceModelAddRankCapByUserDistributionSegmentFromDict(data map[string]interface{}) ExperienceExperienceModelAddRankCapByUserDistributionSegment {
	return ExperienceExperienceModelAddRankCapByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ExperienceExperienceModelAddRankCapByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ExperienceExperienceModelAddRankCapByUserDistributionSegment) Pointer() *ExperienceExperienceModelAddRankCapByUserDistributionSegment {
	return &p
}

func CastExperienceExperienceModelAddRankCapByUserDistributionSegments(data []interface{}) []ExperienceExperienceModelAddRankCapByUserDistributionSegment {
	v := make([]ExperienceExperienceModelAddRankCapByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddRankCapByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddRankCapByUserDistributionSegmentsFromDict(data []ExperienceExperienceModelAddRankCapByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddRankCapByUserDistribution struct {
	Statistics   *ExperienceExperienceModelAddRankCapByUserDistributionStatistics `json:"statistics"`
	Distribution []ExperienceExperienceModelAddRankCapByUserDistributionSegment   `json:"distribution"`
}

func NewExperienceExperienceModelAddRankCapByUserDistributionFromJson(data string) ExperienceExperienceModelAddRankCapByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddRankCapByUserDistributionFromDict(dict)
}

func NewExperienceExperienceModelAddRankCapByUserDistributionFromDict(data map[string]interface{}) ExperienceExperienceModelAddRankCapByUserDistribution {
	return ExperienceExperienceModelAddRankCapByUserDistribution{
		Statistics:   NewExperienceExperienceModelAddRankCapByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExperienceExperienceModelAddRankCapByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExperienceExperienceModelAddRankCapByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExperienceExperienceModelAddRankCapByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExperienceExperienceModelAddRankCapByUserDistribution) Pointer() *ExperienceExperienceModelAddRankCapByUserDistribution {
	return &p
}

func CastExperienceExperienceModelAddRankCapByUserDistributions(data []interface{}) []ExperienceExperienceModelAddRankCapByUserDistribution {
	v := make([]ExperienceExperienceModelAddRankCapByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddRankCapByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddRankCapByUserDistributionsFromDict(data []ExperienceExperienceModelAddRankCapByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExperienceExperienceModelAddRankCapAmountByUserDistributionStatisticsFromJson(data string) ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddRankCapAmountByUserDistributionStatisticsFromDict(dict)
}

func NewExperienceExperienceModelAddRankCapAmountByUserDistributionStatisticsFromDict(data map[string]interface{}) ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics {
	return ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics) Pointer() *ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics {
	return &p
}

func CastExperienceExperienceModelAddRankCapAmountByUserDistributionStatisticses(data []interface{}) []ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics {
	v := make([]ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddRankCapAmountByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddRankCapAmountByUserDistributionStatisticsesFromDict(data []ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewExperienceExperienceModelAddRankCapAmountByUserDistributionSegmentFromJson(data string) ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddRankCapAmountByUserDistributionSegmentFromDict(dict)
}

func NewExperienceExperienceModelAddRankCapAmountByUserDistributionSegmentFromDict(data map[string]interface{}) ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment {
	return ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment) Pointer() *ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment {
	return &p
}

func CastExperienceExperienceModelAddRankCapAmountByUserDistributionSegments(data []interface{}) []ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment {
	v := make([]ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddRankCapAmountByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddRankCapAmountByUserDistributionSegmentsFromDict(data []ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelAddRankCapAmountByUserDistribution struct {
	Statistics   *ExperienceExperienceModelAddRankCapAmountByUserDistributionStatistics `json:"statistics"`
	Distribution []ExperienceExperienceModelAddRankCapAmountByUserDistributionSegment   `json:"distribution"`
}

func NewExperienceExperienceModelAddRankCapAmountByUserDistributionFromJson(data string) ExperienceExperienceModelAddRankCapAmountByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelAddRankCapAmountByUserDistributionFromDict(dict)
}

func NewExperienceExperienceModelAddRankCapAmountByUserDistributionFromDict(data map[string]interface{}) ExperienceExperienceModelAddRankCapAmountByUserDistribution {
	return ExperienceExperienceModelAddRankCapAmountByUserDistribution{
		Statistics:   NewExperienceExperienceModelAddRankCapAmountByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExperienceExperienceModelAddRankCapAmountByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExperienceExperienceModelAddRankCapAmountByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExperienceExperienceModelAddRankCapAmountByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExperienceExperienceModelAddRankCapAmountByUserDistribution) Pointer() *ExperienceExperienceModelAddRankCapAmountByUserDistribution {
	return &p
}

func CastExperienceExperienceModelAddRankCapAmountByUserDistributions(data []interface{}) []ExperienceExperienceModelAddRankCapAmountByUserDistribution {
	v := make([]ExperienceExperienceModelAddRankCapAmountByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelAddRankCapAmountByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelAddRankCapAmountByUserDistributionsFromDict(data []ExperienceExperienceModelAddRankCapAmountByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModelDistributions struct {
	AddExperienceByUser       *ExperienceExperienceModelAddExperienceByUserDistribution       `json:"addExperienceByUser"`
	AddExperienceAmountByUser *ExperienceExperienceModelAddExperienceAmountByUserDistribution `json:"addExperienceAmountByUser"`
	AddRankCapByUser          *ExperienceExperienceModelAddRankCapByUserDistribution          `json:"addRankCapByUser"`
	AddRankCapAmountByUser    *ExperienceExperienceModelAddRankCapAmountByUserDistribution    `json:"addRankCapAmountByUser"`
}

func NewExperienceExperienceModelDistributionsFromJson(data string) ExperienceExperienceModelDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelDistributionsFromDict(dict)
}

func NewExperienceExperienceModelDistributionsFromDict(data map[string]interface{}) ExperienceExperienceModelDistributions {
	return ExperienceExperienceModelDistributions{
		AddExperienceByUser:       NewExperienceExperienceModelAddExperienceByUserDistributionFromDict(core.CastMap(data["addExperienceByUser"])).Pointer(),
		AddExperienceAmountByUser: NewExperienceExperienceModelAddExperienceAmountByUserDistributionFromDict(core.CastMap(data["addExperienceAmountByUser"])).Pointer(),
		AddRankCapByUser:          NewExperienceExperienceModelAddRankCapByUserDistributionFromDict(core.CastMap(data["addRankCapByUser"])).Pointer(),
		AddRankCapAmountByUser:    NewExperienceExperienceModelAddRankCapAmountByUserDistributionFromDict(core.CastMap(data["addRankCapAmountByUser"])).Pointer(),
	}
}

func (p ExperienceExperienceModelDistributions) ToDict() map[string]interface{} {

	var addExperienceByUser map[string]interface{}
	if p.AddExperienceByUser != nil {
		addExperienceByUser = p.AddExperienceByUser.ToDict()
	}
	var addExperienceAmountByUser map[string]interface{}
	if p.AddExperienceAmountByUser != nil {
		addExperienceAmountByUser = p.AddExperienceAmountByUser.ToDict()
	}
	var addRankCapByUser map[string]interface{}
	if p.AddRankCapByUser != nil {
		addRankCapByUser = p.AddRankCapByUser.ToDict()
	}
	var addRankCapAmountByUser map[string]interface{}
	if p.AddRankCapAmountByUser != nil {
		addRankCapAmountByUser = p.AddRankCapAmountByUser.ToDict()
	}
	return map[string]interface{}{
		"addExperienceByUser":       addExperienceByUser,
		"addExperienceAmountByUser": addExperienceAmountByUser,
		"addRankCapByUser":          addRankCapByUser,
		"addRankCapAmountByUser":    addRankCapAmountByUser,
	}
}

func (p ExperienceExperienceModelDistributions) Pointer() *ExperienceExperienceModelDistributions {
	return &p
}

func CastExperienceExperienceModelDistributionses(data []interface{}) []ExperienceExperienceModelDistributions {
	v := make([]ExperienceExperienceModelDistributions, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelDistributionsesFromDict(data []ExperienceExperienceModelDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceExperienceModel struct {
	ExperienceModelId *string                                 `json:"experienceModelId"`
	ExperienceName    *string                                 `json:"experienceName"`
	Statistics        *ExperienceExperienceModelStatistics    `json:"statistics"`
	Distributions     *ExperienceExperienceModelDistributions `json:"distributions"`
	Statuses          []ExperienceStatus                      `json:"statuses"`
}

func NewExperienceExperienceModelFromJson(data string) ExperienceExperienceModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceExperienceModelFromDict(dict)
}

func NewExperienceExperienceModelFromDict(data map[string]interface{}) ExperienceExperienceModel {
	return ExperienceExperienceModel{
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		ExperienceName:    core.CastString(data["experienceName"]),
		Statistics:        NewExperienceExperienceModelStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions:     NewExperienceExperienceModelDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		Statuses:          CastExperienceStatuses(core.CastArray(data["statuses"])),
	}
}

func (p ExperienceExperienceModel) ToDict() map[string]interface{} {

	var experienceModelId *string
	if p.ExperienceModelId != nil {
		experienceModelId = p.ExperienceModelId
	}
	var experienceName *string
	if p.ExperienceName != nil {
		experienceName = p.ExperienceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var statuses []interface{}
	if p.Statuses != nil {
		statuses = CastExperienceStatusesFromDict(
			p.Statuses,
		)
	}
	return map[string]interface{}{
		"experienceModelId": experienceModelId,
		"experienceName":    experienceName,
		"statistics":        statistics,
		"distributions":     distributions,
		"statuses":          statuses,
	}
}

func (p ExperienceExperienceModel) Pointer() *ExperienceExperienceModel {
	return &p
}

func CastExperienceExperienceModels(data []interface{}) []ExperienceExperienceModel {
	v := make([]ExperienceExperienceModel, 0)
	for _, d := range data {
		v = append(v, NewExperienceExperienceModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceExperienceModelsFromDict(data []ExperienceExperienceModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceNamespaceStatistics struct {
	Experience *int64 `json:"experience"`
	RankCap    *int64 `json:"rankCap"`
}

func NewExperienceNamespaceStatisticsFromJson(data string) ExperienceNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceNamespaceStatisticsFromDict(dict)
}

func NewExperienceNamespaceStatisticsFromDict(data map[string]interface{}) ExperienceNamespaceStatistics {
	return ExperienceNamespaceStatistics{
		Experience: core.CastInt64(data["experience"]),
		RankCap:    core.CastInt64(data["rankCap"]),
	}
}

func (p ExperienceNamespaceStatistics) ToDict() map[string]interface{} {

	var experience *int64
	if p.Experience != nil {
		experience = p.Experience
	}
	var rankCap *int64
	if p.RankCap != nil {
		rankCap = p.RankCap
	}
	return map[string]interface{}{
		"experience": experience,
		"rankCap":    rankCap,
	}
}

func (p ExperienceNamespaceStatistics) Pointer() *ExperienceNamespaceStatistics {
	return &p
}

func CastExperienceNamespaceStatisticses(data []interface{}) []ExperienceNamespaceStatistics {
	v := make([]ExperienceNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceNamespaceStatisticsesFromDict(data []ExperienceNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceNamespaceAddExperienceByExperienceDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExperienceNamespaceAddExperienceByExperienceDistributionStatisticsFromJson(data string) ExperienceNamespaceAddExperienceByExperienceDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceNamespaceAddExperienceByExperienceDistributionStatisticsFromDict(dict)
}

func NewExperienceNamespaceAddExperienceByExperienceDistributionStatisticsFromDict(data map[string]interface{}) ExperienceNamespaceAddExperienceByExperienceDistributionStatistics {
	return ExperienceNamespaceAddExperienceByExperienceDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExperienceNamespaceAddExperienceByExperienceDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExperienceNamespaceAddExperienceByExperienceDistributionStatistics) Pointer() *ExperienceNamespaceAddExperienceByExperienceDistributionStatistics {
	return &p
}

func CastExperienceNamespaceAddExperienceByExperienceDistributionStatisticses(data []interface{}) []ExperienceNamespaceAddExperienceByExperienceDistributionStatistics {
	v := make([]ExperienceNamespaceAddExperienceByExperienceDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceNamespaceAddExperienceByExperienceDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceNamespaceAddExperienceByExperienceDistributionStatisticsesFromDict(data []ExperienceNamespaceAddExperienceByExperienceDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceNamespaceAddExperienceByExperienceDistributionSegment struct {
	ExperienceName *string `json:"experienceName"`
	Count          *int64  `json:"count"`
}

func NewExperienceNamespaceAddExperienceByExperienceDistributionSegmentFromJson(data string) ExperienceNamespaceAddExperienceByExperienceDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceNamespaceAddExperienceByExperienceDistributionSegmentFromDict(dict)
}

func NewExperienceNamespaceAddExperienceByExperienceDistributionSegmentFromDict(data map[string]interface{}) ExperienceNamespaceAddExperienceByExperienceDistributionSegment {
	return ExperienceNamespaceAddExperienceByExperienceDistributionSegment{
		ExperienceName: core.CastString(data["experienceName"]),
		Count:          core.CastInt64(data["count"]),
	}
}

func (p ExperienceNamespaceAddExperienceByExperienceDistributionSegment) ToDict() map[string]interface{} {

	var experienceName *string
	if p.ExperienceName != nil {
		experienceName = p.ExperienceName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"experienceName": experienceName,
		"count":          count,
	}
}

func (p ExperienceNamespaceAddExperienceByExperienceDistributionSegment) Pointer() *ExperienceNamespaceAddExperienceByExperienceDistributionSegment {
	return &p
}

func CastExperienceNamespaceAddExperienceByExperienceDistributionSegments(data []interface{}) []ExperienceNamespaceAddExperienceByExperienceDistributionSegment {
	v := make([]ExperienceNamespaceAddExperienceByExperienceDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExperienceNamespaceAddExperienceByExperienceDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceNamespaceAddExperienceByExperienceDistributionSegmentsFromDict(data []ExperienceNamespaceAddExperienceByExperienceDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceNamespaceAddExperienceByExperienceDistribution struct {
	Statistics   *ExperienceNamespaceAddExperienceByExperienceDistributionStatistics `json:"statistics"`
	Distribution []ExperienceNamespaceAddExperienceByExperienceDistributionSegment   `json:"distribution"`
}

func NewExperienceNamespaceAddExperienceByExperienceDistributionFromJson(data string) ExperienceNamespaceAddExperienceByExperienceDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceNamespaceAddExperienceByExperienceDistributionFromDict(dict)
}

func NewExperienceNamespaceAddExperienceByExperienceDistributionFromDict(data map[string]interface{}) ExperienceNamespaceAddExperienceByExperienceDistribution {
	return ExperienceNamespaceAddExperienceByExperienceDistribution{
		Statistics:   NewExperienceNamespaceAddExperienceByExperienceDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExperienceNamespaceAddExperienceByExperienceDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExperienceNamespaceAddExperienceByExperienceDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExperienceNamespaceAddExperienceByExperienceDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExperienceNamespaceAddExperienceByExperienceDistribution) Pointer() *ExperienceNamespaceAddExperienceByExperienceDistribution {
	return &p
}

func CastExperienceNamespaceAddExperienceByExperienceDistributions(data []interface{}) []ExperienceNamespaceAddExperienceByExperienceDistribution {
	v := make([]ExperienceNamespaceAddExperienceByExperienceDistribution, 0)
	for _, d := range data {
		v = append(v, NewExperienceNamespaceAddExperienceByExperienceDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceNamespaceAddExperienceByExperienceDistributionsFromDict(data []ExperienceNamespaceAddExperienceByExperienceDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceNamespaceAddRankCapByExperienceDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewExperienceNamespaceAddRankCapByExperienceDistributionStatisticsFromJson(data string) ExperienceNamespaceAddRankCapByExperienceDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceNamespaceAddRankCapByExperienceDistributionStatisticsFromDict(dict)
}

func NewExperienceNamespaceAddRankCapByExperienceDistributionStatisticsFromDict(data map[string]interface{}) ExperienceNamespaceAddRankCapByExperienceDistributionStatistics {
	return ExperienceNamespaceAddRankCapByExperienceDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ExperienceNamespaceAddRankCapByExperienceDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ExperienceNamespaceAddRankCapByExperienceDistributionStatistics) Pointer() *ExperienceNamespaceAddRankCapByExperienceDistributionStatistics {
	return &p
}

func CastExperienceNamespaceAddRankCapByExperienceDistributionStatisticses(data []interface{}) []ExperienceNamespaceAddRankCapByExperienceDistributionStatistics {
	v := make([]ExperienceNamespaceAddRankCapByExperienceDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewExperienceNamespaceAddRankCapByExperienceDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceNamespaceAddRankCapByExperienceDistributionStatisticsesFromDict(data []ExperienceNamespaceAddRankCapByExperienceDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceNamespaceAddRankCapByExperienceDistributionSegment struct {
	ExperienceName *string `json:"experienceName"`
	Count          *int64  `json:"count"`
}

func NewExperienceNamespaceAddRankCapByExperienceDistributionSegmentFromJson(data string) ExperienceNamespaceAddRankCapByExperienceDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceNamespaceAddRankCapByExperienceDistributionSegmentFromDict(dict)
}

func NewExperienceNamespaceAddRankCapByExperienceDistributionSegmentFromDict(data map[string]interface{}) ExperienceNamespaceAddRankCapByExperienceDistributionSegment {
	return ExperienceNamespaceAddRankCapByExperienceDistributionSegment{
		ExperienceName: core.CastString(data["experienceName"]),
		Count:          core.CastInt64(data["count"]),
	}
}

func (p ExperienceNamespaceAddRankCapByExperienceDistributionSegment) ToDict() map[string]interface{} {

	var experienceName *string
	if p.ExperienceName != nil {
		experienceName = p.ExperienceName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"experienceName": experienceName,
		"count":          count,
	}
}

func (p ExperienceNamespaceAddRankCapByExperienceDistributionSegment) Pointer() *ExperienceNamespaceAddRankCapByExperienceDistributionSegment {
	return &p
}

func CastExperienceNamespaceAddRankCapByExperienceDistributionSegments(data []interface{}) []ExperienceNamespaceAddRankCapByExperienceDistributionSegment {
	v := make([]ExperienceNamespaceAddRankCapByExperienceDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewExperienceNamespaceAddRankCapByExperienceDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceNamespaceAddRankCapByExperienceDistributionSegmentsFromDict(data []ExperienceNamespaceAddRankCapByExperienceDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceNamespaceAddRankCapByExperienceDistribution struct {
	Statistics   *ExperienceNamespaceAddRankCapByExperienceDistributionStatistics `json:"statistics"`
	Distribution []ExperienceNamespaceAddRankCapByExperienceDistributionSegment   `json:"distribution"`
}

func NewExperienceNamespaceAddRankCapByExperienceDistributionFromJson(data string) ExperienceNamespaceAddRankCapByExperienceDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceNamespaceAddRankCapByExperienceDistributionFromDict(dict)
}

func NewExperienceNamespaceAddRankCapByExperienceDistributionFromDict(data map[string]interface{}) ExperienceNamespaceAddRankCapByExperienceDistribution {
	return ExperienceNamespaceAddRankCapByExperienceDistribution{
		Statistics:   NewExperienceNamespaceAddRankCapByExperienceDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastExperienceNamespaceAddRankCapByExperienceDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ExperienceNamespaceAddRankCapByExperienceDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastExperienceNamespaceAddRankCapByExperienceDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ExperienceNamespaceAddRankCapByExperienceDistribution) Pointer() *ExperienceNamespaceAddRankCapByExperienceDistribution {
	return &p
}

func CastExperienceNamespaceAddRankCapByExperienceDistributions(data []interface{}) []ExperienceNamespaceAddRankCapByExperienceDistribution {
	v := make([]ExperienceNamespaceAddRankCapByExperienceDistribution, 0)
	for _, d := range data {
		v = append(v, NewExperienceNamespaceAddRankCapByExperienceDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceNamespaceAddRankCapByExperienceDistributionsFromDict(data []ExperienceNamespaceAddRankCapByExperienceDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceNamespaceDistributions struct {
	AddExperienceByExperience *ExperienceNamespaceAddExperienceByExperienceDistribution `json:"addExperienceByExperience"`
	AddRankCapByExperience    *ExperienceNamespaceAddRankCapByExperienceDistribution    `json:"addRankCapByExperience"`
}

func NewExperienceNamespaceDistributionsFromJson(data string) ExperienceNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceNamespaceDistributionsFromDict(dict)
}

func NewExperienceNamespaceDistributionsFromDict(data map[string]interface{}) ExperienceNamespaceDistributions {
	return ExperienceNamespaceDistributions{
		AddExperienceByExperience: NewExperienceNamespaceAddExperienceByExperienceDistributionFromDict(core.CastMap(data["addExperienceByExperience"])).Pointer(),
		AddRankCapByExperience:    NewExperienceNamespaceAddRankCapByExperienceDistributionFromDict(core.CastMap(data["addRankCapByExperience"])).Pointer(),
	}
}

func (p ExperienceNamespaceDistributions) ToDict() map[string]interface{} {

	var addExperienceByExperience map[string]interface{}
	if p.AddExperienceByExperience != nil {
		addExperienceByExperience = p.AddExperienceByExperience.ToDict()
	}
	var addRankCapByExperience map[string]interface{}
	if p.AddRankCapByExperience != nil {
		addRankCapByExperience = p.AddRankCapByExperience.ToDict()
	}
	return map[string]interface{}{
		"addExperienceByExperience": addExperienceByExperience,
		"addRankCapByExperience":    addRankCapByExperience,
	}
}

func (p ExperienceNamespaceDistributions) Pointer() *ExperienceNamespaceDistributions {
	return &p
}

func CastExperienceNamespaceDistributionses(data []interface{}) []ExperienceNamespaceDistributions {
	v := make([]ExperienceNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewExperienceNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceNamespaceDistributionsesFromDict(data []ExperienceNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceNamespace struct {
	NamespaceId      *string                           `json:"namespaceId"`
	Year             *int32                            `json:"year"`
	Month            *int32                            `json:"month"`
	Day              *int32                            `json:"day"`
	NamespaceName    *string                           `json:"namespaceName"`
	Statistics       *ExperienceNamespaceStatistics    `json:"statistics"`
	Distributions    *ExperienceNamespaceDistributions `json:"distributions"`
	ExperienceModels []ExperienceExperienceModel       `json:"experienceModels"`
}

func NewExperienceNamespaceFromJson(data string) ExperienceNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceNamespaceFromDict(dict)
}

func NewExperienceNamespaceFromDict(data map[string]interface{}) ExperienceNamespace {
	return ExperienceNamespace{
		NamespaceId:      core.CastString(data["namespaceId"]),
		Year:             core.CastInt32(data["year"]),
		Month:            core.CastInt32(data["month"]),
		Day:              core.CastInt32(data["day"]),
		NamespaceName:    core.CastString(data["namespaceName"]),
		Statistics:       NewExperienceNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions:    NewExperienceNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		ExperienceModels: CastExperienceExperienceModels(core.CastArray(data["experienceModels"])),
	}
}

func (p ExperienceNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var experienceModels []interface{}
	if p.ExperienceModels != nil {
		experienceModels = CastExperienceExperienceModelsFromDict(
			p.ExperienceModels,
		)
	}
	return map[string]interface{}{
		"namespaceId":      namespaceId,
		"year":             year,
		"month":            month,
		"day":              day,
		"namespaceName":    namespaceName,
		"statistics":       statistics,
		"distributions":    distributions,
		"experienceModels": experienceModels,
	}
}

func (p ExperienceNamespace) Pointer() *ExperienceNamespace {
	return &p
}

func CastExperienceNamespaces(data []interface{}) []ExperienceNamespace {
	v := make([]ExperienceNamespace, 0)
	for _, d := range data {
		v = append(v, NewExperienceNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceNamespacesFromDict(data []ExperienceNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationFormStatistics struct {
	Update *int64 `json:"update"`
}

func NewFormationFormStatisticsFromJson(data string) FormationFormStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationFormStatisticsFromDict(dict)
}

func NewFormationFormStatisticsFromDict(data map[string]interface{}) FormationFormStatistics {
	return FormationFormStatistics{
		Update: core.CastInt64(data["update"]),
	}
}

func (p FormationFormStatistics) ToDict() map[string]interface{} {

	var update *int64
	if p.Update != nil {
		update = p.Update
	}
	return map[string]interface{}{
		"update": update,
	}
}

func (p FormationFormStatistics) Pointer() *FormationFormStatistics {
	return &p
}

func CastFormationFormStatisticses(data []interface{}) []FormationFormStatistics {
	v := make([]FormationFormStatistics, 0)
	for _, d := range data {
		v = append(v, NewFormationFormStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationFormStatisticsesFromDict(data []FormationFormStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationFormSlotDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFormationFormSlotDistributionStatisticsFromJson(data string) FormationFormSlotDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationFormSlotDistributionStatisticsFromDict(dict)
}

func NewFormationFormSlotDistributionStatisticsFromDict(data map[string]interface{}) FormationFormSlotDistributionStatistics {
	return FormationFormSlotDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FormationFormSlotDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FormationFormSlotDistributionStatistics) Pointer() *FormationFormSlotDistributionStatistics {
	return &p
}

func CastFormationFormSlotDistributionStatisticses(data []interface{}) []FormationFormSlotDistributionStatistics {
	v := make([]FormationFormSlotDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFormationFormSlotDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationFormSlotDistributionStatisticsesFromDict(data []FormationFormSlotDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationFormSlotDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewFormationFormSlotDistributionSegmentFromJson(data string) FormationFormSlotDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationFormSlotDistributionSegmentFromDict(dict)
}

func NewFormationFormSlotDistributionSegmentFromDict(data map[string]interface{}) FormationFormSlotDistributionSegment {
	return FormationFormSlotDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p FormationFormSlotDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p FormationFormSlotDistributionSegment) Pointer() *FormationFormSlotDistributionSegment {
	return &p
}

func CastFormationFormSlotDistributionSegments(data []interface{}) []FormationFormSlotDistributionSegment {
	v := make([]FormationFormSlotDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFormationFormSlotDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationFormSlotDistributionSegmentsFromDict(data []FormationFormSlotDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationFormSlotDistribution struct {
	Statistics   *FormationFormSlotDistributionStatistics `json:"statistics"`
	Distribution []FormationFormSlotDistributionSegment   `json:"distribution"`
}

func NewFormationFormSlotDistributionFromJson(data string) FormationFormSlotDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationFormSlotDistributionFromDict(dict)
}

func NewFormationFormSlotDistributionFromDict(data map[string]interface{}) FormationFormSlotDistribution {
	return FormationFormSlotDistribution{
		Statistics:   NewFormationFormSlotDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFormationFormSlotDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FormationFormSlotDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFormationFormSlotDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FormationFormSlotDistribution) Pointer() *FormationFormSlotDistribution {
	return &p
}

func CastFormationFormSlotDistributions(data []interface{}) []FormationFormSlotDistribution {
	v := make([]FormationFormSlotDistribution, 0)
	for _, d := range data {
		v = append(v, NewFormationFormSlotDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationFormSlotDistributionsFromDict(data []FormationFormSlotDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationFormUsageDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFormationFormUsageDistributionStatisticsFromJson(data string) FormationFormUsageDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationFormUsageDistributionStatisticsFromDict(dict)
}

func NewFormationFormUsageDistributionStatisticsFromDict(data map[string]interface{}) FormationFormUsageDistributionStatistics {
	return FormationFormUsageDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FormationFormUsageDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FormationFormUsageDistributionStatistics) Pointer() *FormationFormUsageDistributionStatistics {
	return &p
}

func CastFormationFormUsageDistributionStatisticses(data []interface{}) []FormationFormUsageDistributionStatistics {
	v := make([]FormationFormUsageDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFormationFormUsageDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationFormUsageDistributionStatisticsesFromDict(data []FormationFormUsageDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationFormUsageDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewFormationFormUsageDistributionSegmentFromJson(data string) FormationFormUsageDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationFormUsageDistributionSegmentFromDict(dict)
}

func NewFormationFormUsageDistributionSegmentFromDict(data map[string]interface{}) FormationFormUsageDistributionSegment {
	return FormationFormUsageDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p FormationFormUsageDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p FormationFormUsageDistributionSegment) Pointer() *FormationFormUsageDistributionSegment {
	return &p
}

func CastFormationFormUsageDistributionSegments(data []interface{}) []FormationFormUsageDistributionSegment {
	v := make([]FormationFormUsageDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFormationFormUsageDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationFormUsageDistributionSegmentsFromDict(data []FormationFormUsageDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationFormUsageDistribution struct {
	Statistics   *FormationFormUsageDistributionStatistics `json:"statistics"`
	Distribution []FormationFormUsageDistributionSegment   `json:"distribution"`
}

func NewFormationFormUsageDistributionFromJson(data string) FormationFormUsageDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationFormUsageDistributionFromDict(dict)
}

func NewFormationFormUsageDistributionFromDict(data map[string]interface{}) FormationFormUsageDistribution {
	return FormationFormUsageDistribution{
		Statistics:   NewFormationFormUsageDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFormationFormUsageDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FormationFormUsageDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFormationFormUsageDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FormationFormUsageDistribution) Pointer() *FormationFormUsageDistribution {
	return &p
}

func CastFormationFormUsageDistributions(data []interface{}) []FormationFormUsageDistribution {
	v := make([]FormationFormUsageDistribution, 0)
	for _, d := range data {
		v = append(v, NewFormationFormUsageDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationFormUsageDistributionsFromDict(data []FormationFormUsageDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationFormDistributions struct {
	Slot  *FormationFormSlotDistribution  `json:"slot"`
	Usage *FormationFormUsageDistribution `json:"usage"`
}

func NewFormationFormDistributionsFromJson(data string) FormationFormDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationFormDistributionsFromDict(dict)
}

func NewFormationFormDistributionsFromDict(data map[string]interface{}) FormationFormDistributions {
	return FormationFormDistributions{
		Slot:  NewFormationFormSlotDistributionFromDict(core.CastMap(data["slot"])).Pointer(),
		Usage: NewFormationFormUsageDistributionFromDict(core.CastMap(data["usage"])).Pointer(),
	}
}

func (p FormationFormDistributions) ToDict() map[string]interface{} {

	var slot map[string]interface{}
	if p.Slot != nil {
		slot = p.Slot.ToDict()
	}
	var usage map[string]interface{}
	if p.Usage != nil {
		usage = p.Usage.ToDict()
	}
	return map[string]interface{}{
		"slot":  slot,
		"usage": usage,
	}
}

func (p FormationFormDistributions) Pointer() *FormationFormDistributions {
	return &p
}

func CastFormationFormDistributionses(data []interface{}) []FormationFormDistributions {
	v := make([]FormationFormDistributions, 0)
	for _, d := range data {
		v = append(v, NewFormationFormDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationFormDistributionsesFromDict(data []FormationFormDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationForm struct {
	FormId        *string                     `json:"formId"`
	Index         *int32                      `json:"index"`
	Statistics    *FormationFormStatistics    `json:"statistics"`
	Distributions *FormationFormDistributions `json:"distributions"`
}

func NewFormationFormFromJson(data string) FormationForm {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationFormFromDict(dict)
}

func NewFormationFormFromDict(data map[string]interface{}) FormationForm {
	return FormationForm{
		FormId:        core.CastString(data["formId"]),
		Index:         core.CastInt32(data["index"]),
		Statistics:    NewFormationFormStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewFormationFormDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p FormationForm) ToDict() map[string]interface{} {

	var formId *string
	if p.FormId != nil {
		formId = p.FormId
	}
	var index *int32
	if p.Index != nil {
		index = p.Index
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"formId":        formId,
		"index":         index,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p FormationForm) Pointer() *FormationForm {
	return &p
}

func CastFormationForms(data []interface{}) []FormationForm {
	v := make([]FormationForm, 0)
	for _, d := range data {
		v = append(v, NewFormationFormFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationFormsFromDict(data []FormationForm) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationMoldStatistics struct {
	IncreaseCapacity       *int64 `json:"increaseCapacity"`
	IncreaseCapacityAmount *int64 `json:"increaseCapacityAmount"`
}

func NewFormationMoldStatisticsFromJson(data string) FormationMoldStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationMoldStatisticsFromDict(dict)
}

func NewFormationMoldStatisticsFromDict(data map[string]interface{}) FormationMoldStatistics {
	return FormationMoldStatistics{
		IncreaseCapacity:       core.CastInt64(data["increaseCapacity"]),
		IncreaseCapacityAmount: core.CastInt64(data["increaseCapacityAmount"]),
	}
}

func (p FormationMoldStatistics) ToDict() map[string]interface{} {

	var increaseCapacity *int64
	if p.IncreaseCapacity != nil {
		increaseCapacity = p.IncreaseCapacity
	}
	var increaseCapacityAmount *int64
	if p.IncreaseCapacityAmount != nil {
		increaseCapacityAmount = p.IncreaseCapacityAmount
	}
	return map[string]interface{}{
		"increaseCapacity":       increaseCapacity,
		"increaseCapacityAmount": increaseCapacityAmount,
	}
}

func (p FormationMoldStatistics) Pointer() *FormationMoldStatistics {
	return &p
}

func CastFormationMoldStatisticses(data []interface{}) []FormationMoldStatistics {
	v := make([]FormationMoldStatistics, 0)
	for _, d := range data {
		v = append(v, NewFormationMoldStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationMoldStatisticsesFromDict(data []FormationMoldStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationMoldCapacityDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFormationMoldCapacityDistributionStatisticsFromJson(data string) FormationMoldCapacityDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationMoldCapacityDistributionStatisticsFromDict(dict)
}

func NewFormationMoldCapacityDistributionStatisticsFromDict(data map[string]interface{}) FormationMoldCapacityDistributionStatistics {
	return FormationMoldCapacityDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FormationMoldCapacityDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FormationMoldCapacityDistributionStatistics) Pointer() *FormationMoldCapacityDistributionStatistics {
	return &p
}

func CastFormationMoldCapacityDistributionStatisticses(data []interface{}) []FormationMoldCapacityDistributionStatistics {
	v := make([]FormationMoldCapacityDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFormationMoldCapacityDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationMoldCapacityDistributionStatisticsesFromDict(data []FormationMoldCapacityDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationMoldCapacityDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewFormationMoldCapacityDistributionSegmentFromJson(data string) FormationMoldCapacityDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationMoldCapacityDistributionSegmentFromDict(dict)
}

func NewFormationMoldCapacityDistributionSegmentFromDict(data map[string]interface{}) FormationMoldCapacityDistributionSegment {
	return FormationMoldCapacityDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p FormationMoldCapacityDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p FormationMoldCapacityDistributionSegment) Pointer() *FormationMoldCapacityDistributionSegment {
	return &p
}

func CastFormationMoldCapacityDistributionSegments(data []interface{}) []FormationMoldCapacityDistributionSegment {
	v := make([]FormationMoldCapacityDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFormationMoldCapacityDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationMoldCapacityDistributionSegmentsFromDict(data []FormationMoldCapacityDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationMoldCapacityDistribution struct {
	Statistics   *FormationMoldCapacityDistributionStatistics `json:"statistics"`
	Distribution []FormationMoldCapacityDistributionSegment   `json:"distribution"`
}

func NewFormationMoldCapacityDistributionFromJson(data string) FormationMoldCapacityDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationMoldCapacityDistributionFromDict(dict)
}

func NewFormationMoldCapacityDistributionFromDict(data map[string]interface{}) FormationMoldCapacityDistribution {
	return FormationMoldCapacityDistribution{
		Statistics:   NewFormationMoldCapacityDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFormationMoldCapacityDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FormationMoldCapacityDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFormationMoldCapacityDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FormationMoldCapacityDistribution) Pointer() *FormationMoldCapacityDistribution {
	return &p
}

func CastFormationMoldCapacityDistributions(data []interface{}) []FormationMoldCapacityDistribution {
	v := make([]FormationMoldCapacityDistribution, 0)
	for _, d := range data {
		v = append(v, NewFormationMoldCapacityDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationMoldCapacityDistributionsFromDict(data []FormationMoldCapacityDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationMoldUpdateByIndexDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFormationMoldUpdateByIndexDistributionStatisticsFromJson(data string) FormationMoldUpdateByIndexDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationMoldUpdateByIndexDistributionStatisticsFromDict(dict)
}

func NewFormationMoldUpdateByIndexDistributionStatisticsFromDict(data map[string]interface{}) FormationMoldUpdateByIndexDistributionStatistics {
	return FormationMoldUpdateByIndexDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FormationMoldUpdateByIndexDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FormationMoldUpdateByIndexDistributionStatistics) Pointer() *FormationMoldUpdateByIndexDistributionStatistics {
	return &p
}

func CastFormationMoldUpdateByIndexDistributionStatisticses(data []interface{}) []FormationMoldUpdateByIndexDistributionStatistics {
	v := make([]FormationMoldUpdateByIndexDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFormationMoldUpdateByIndexDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationMoldUpdateByIndexDistributionStatisticsesFromDict(data []FormationMoldUpdateByIndexDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationMoldUpdateByIndexDistributionSegment struct {
	Index *int64 `json:"index"`
	Count *int64 `json:"count"`
}

func NewFormationMoldUpdateByIndexDistributionSegmentFromJson(data string) FormationMoldUpdateByIndexDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationMoldUpdateByIndexDistributionSegmentFromDict(dict)
}

func NewFormationMoldUpdateByIndexDistributionSegmentFromDict(data map[string]interface{}) FormationMoldUpdateByIndexDistributionSegment {
	return FormationMoldUpdateByIndexDistributionSegment{
		Index: core.CastInt64(data["index"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p FormationMoldUpdateByIndexDistributionSegment) ToDict() map[string]interface{} {

	var index *int64
	if p.Index != nil {
		index = p.Index
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"index": index,
		"count": count,
	}
}

func (p FormationMoldUpdateByIndexDistributionSegment) Pointer() *FormationMoldUpdateByIndexDistributionSegment {
	return &p
}

func CastFormationMoldUpdateByIndexDistributionSegments(data []interface{}) []FormationMoldUpdateByIndexDistributionSegment {
	v := make([]FormationMoldUpdateByIndexDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFormationMoldUpdateByIndexDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationMoldUpdateByIndexDistributionSegmentsFromDict(data []FormationMoldUpdateByIndexDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationMoldUpdateByIndexDistribution struct {
	Statistics   *FormationMoldUpdateByIndexDistributionStatistics `json:"statistics"`
	Distribution []FormationMoldUpdateByIndexDistributionSegment   `json:"distribution"`
}

func NewFormationMoldUpdateByIndexDistributionFromJson(data string) FormationMoldUpdateByIndexDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationMoldUpdateByIndexDistributionFromDict(dict)
}

func NewFormationMoldUpdateByIndexDistributionFromDict(data map[string]interface{}) FormationMoldUpdateByIndexDistribution {
	return FormationMoldUpdateByIndexDistribution{
		Statistics:   NewFormationMoldUpdateByIndexDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFormationMoldUpdateByIndexDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FormationMoldUpdateByIndexDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFormationMoldUpdateByIndexDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FormationMoldUpdateByIndexDistribution) Pointer() *FormationMoldUpdateByIndexDistribution {
	return &p
}

func CastFormationMoldUpdateByIndexDistributions(data []interface{}) []FormationMoldUpdateByIndexDistribution {
	v := make([]FormationMoldUpdateByIndexDistribution, 0)
	for _, d := range data {
		v = append(v, NewFormationMoldUpdateByIndexDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationMoldUpdateByIndexDistributionsFromDict(data []FormationMoldUpdateByIndexDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationMoldDistributions struct {
	Capacity      *FormationMoldCapacityDistribution      `json:"capacity"`
	UpdateByIndex *FormationMoldUpdateByIndexDistribution `json:"updateByIndex"`
}

func NewFormationMoldDistributionsFromJson(data string) FormationMoldDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationMoldDistributionsFromDict(dict)
}

func NewFormationMoldDistributionsFromDict(data map[string]interface{}) FormationMoldDistributions {
	return FormationMoldDistributions{
		Capacity:      NewFormationMoldCapacityDistributionFromDict(core.CastMap(data["capacity"])).Pointer(),
		UpdateByIndex: NewFormationMoldUpdateByIndexDistributionFromDict(core.CastMap(data["updateByIndex"])).Pointer(),
	}
}

func (p FormationMoldDistributions) ToDict() map[string]interface{} {

	var capacity map[string]interface{}
	if p.Capacity != nil {
		capacity = p.Capacity.ToDict()
	}
	var updateByIndex map[string]interface{}
	if p.UpdateByIndex != nil {
		updateByIndex = p.UpdateByIndex.ToDict()
	}
	return map[string]interface{}{
		"capacity":      capacity,
		"updateByIndex": updateByIndex,
	}
}

func (p FormationMoldDistributions) Pointer() *FormationMoldDistributions {
	return &p
}

func CastFormationMoldDistributionses(data []interface{}) []FormationMoldDistributions {
	v := make([]FormationMoldDistributions, 0)
	for _, d := range data {
		v = append(v, NewFormationMoldDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationMoldDistributionsesFromDict(data []FormationMoldDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationMold struct {
	MoldId        *string                     `json:"moldId"`
	MoldModelName *string                     `json:"moldModelName"`
	Statistics    *FormationMoldStatistics    `json:"statistics"`
	Distributions *FormationMoldDistributions `json:"distributions"`
	Forms         []FormationForm             `json:"forms"`
}

func NewFormationMoldFromJson(data string) FormationMold {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationMoldFromDict(dict)
}

func NewFormationMoldFromDict(data map[string]interface{}) FormationMold {
	return FormationMold{
		MoldId:        core.CastString(data["moldId"]),
		MoldModelName: core.CastString(data["moldModelName"]),
		Statistics:    NewFormationMoldStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewFormationMoldDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		Forms:         CastFormationForms(core.CastArray(data["forms"])),
	}
}

func (p FormationMold) ToDict() map[string]interface{} {

	var moldId *string
	if p.MoldId != nil {
		moldId = p.MoldId
	}
	var moldModelName *string
	if p.MoldModelName != nil {
		moldModelName = p.MoldModelName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var forms []interface{}
	if p.Forms != nil {
		forms = CastFormationFormsFromDict(
			p.Forms,
		)
	}
	return map[string]interface{}{
		"moldId":        moldId,
		"moldModelName": moldModelName,
		"statistics":    statistics,
		"distributions": distributions,
		"forms":         forms,
	}
}

func (p FormationMold) Pointer() *FormationMold {
	return &p
}

func CastFormationMolds(data []interface{}) []FormationMold {
	v := make([]FormationMold, 0)
	for _, d := range data {
		v = append(v, NewFormationMoldFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationMoldsFromDict(data []FormationMold) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationNamespaceStatistics struct {
	Update   *int64 `json:"update"`
	Increase *int64 `json:"increase"`
}

func NewFormationNamespaceStatisticsFromJson(data string) FormationNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationNamespaceStatisticsFromDict(dict)
}

func NewFormationNamespaceStatisticsFromDict(data map[string]interface{}) FormationNamespaceStatistics {
	return FormationNamespaceStatistics{
		Update:   core.CastInt64(data["update"]),
		Increase: core.CastInt64(data["increase"]),
	}
}

func (p FormationNamespaceStatistics) ToDict() map[string]interface{} {

	var update *int64
	if p.Update != nil {
		update = p.Update
	}
	var increase *int64
	if p.Increase != nil {
		increase = p.Increase
	}
	return map[string]interface{}{
		"update":   update,
		"increase": increase,
	}
}

func (p FormationNamespaceStatistics) Pointer() *FormationNamespaceStatistics {
	return &p
}

func CastFormationNamespaceStatisticses(data []interface{}) []FormationNamespaceStatistics {
	v := make([]FormationNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewFormationNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationNamespaceStatisticsesFromDict(data []FormationNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationNamespaceUpdateByMoldDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFormationNamespaceUpdateByMoldDistributionStatisticsFromJson(data string) FormationNamespaceUpdateByMoldDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationNamespaceUpdateByMoldDistributionStatisticsFromDict(dict)
}

func NewFormationNamespaceUpdateByMoldDistributionStatisticsFromDict(data map[string]interface{}) FormationNamespaceUpdateByMoldDistributionStatistics {
	return FormationNamespaceUpdateByMoldDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FormationNamespaceUpdateByMoldDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FormationNamespaceUpdateByMoldDistributionStatistics) Pointer() *FormationNamespaceUpdateByMoldDistributionStatistics {
	return &p
}

func CastFormationNamespaceUpdateByMoldDistributionStatisticses(data []interface{}) []FormationNamespaceUpdateByMoldDistributionStatistics {
	v := make([]FormationNamespaceUpdateByMoldDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFormationNamespaceUpdateByMoldDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationNamespaceUpdateByMoldDistributionStatisticsesFromDict(data []FormationNamespaceUpdateByMoldDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationNamespaceUpdateByMoldDistributionSegment struct {
	MoldModelName *string `json:"moldModelName"`
	Count         *int64  `json:"count"`
}

func NewFormationNamespaceUpdateByMoldDistributionSegmentFromJson(data string) FormationNamespaceUpdateByMoldDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationNamespaceUpdateByMoldDistributionSegmentFromDict(dict)
}

func NewFormationNamespaceUpdateByMoldDistributionSegmentFromDict(data map[string]interface{}) FormationNamespaceUpdateByMoldDistributionSegment {
	return FormationNamespaceUpdateByMoldDistributionSegment{
		MoldModelName: core.CastString(data["moldModelName"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p FormationNamespaceUpdateByMoldDistributionSegment) ToDict() map[string]interface{} {

	var moldModelName *string
	if p.MoldModelName != nil {
		moldModelName = p.MoldModelName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"moldModelName": moldModelName,
		"count":         count,
	}
}

func (p FormationNamespaceUpdateByMoldDistributionSegment) Pointer() *FormationNamespaceUpdateByMoldDistributionSegment {
	return &p
}

func CastFormationNamespaceUpdateByMoldDistributionSegments(data []interface{}) []FormationNamespaceUpdateByMoldDistributionSegment {
	v := make([]FormationNamespaceUpdateByMoldDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFormationNamespaceUpdateByMoldDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationNamespaceUpdateByMoldDistributionSegmentsFromDict(data []FormationNamespaceUpdateByMoldDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationNamespaceUpdateByMoldDistribution struct {
	Statistics   *FormationNamespaceUpdateByMoldDistributionStatistics `json:"statistics"`
	Distribution []FormationNamespaceUpdateByMoldDistributionSegment   `json:"distribution"`
}

func NewFormationNamespaceUpdateByMoldDistributionFromJson(data string) FormationNamespaceUpdateByMoldDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationNamespaceUpdateByMoldDistributionFromDict(dict)
}

func NewFormationNamespaceUpdateByMoldDistributionFromDict(data map[string]interface{}) FormationNamespaceUpdateByMoldDistribution {
	return FormationNamespaceUpdateByMoldDistribution{
		Statistics:   NewFormationNamespaceUpdateByMoldDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFormationNamespaceUpdateByMoldDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FormationNamespaceUpdateByMoldDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFormationNamespaceUpdateByMoldDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FormationNamespaceUpdateByMoldDistribution) Pointer() *FormationNamespaceUpdateByMoldDistribution {
	return &p
}

func CastFormationNamespaceUpdateByMoldDistributions(data []interface{}) []FormationNamespaceUpdateByMoldDistribution {
	v := make([]FormationNamespaceUpdateByMoldDistribution, 0)
	for _, d := range data {
		v = append(v, NewFormationNamespaceUpdateByMoldDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationNamespaceUpdateByMoldDistributionsFromDict(data []FormationNamespaceUpdateByMoldDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationNamespaceIncreaseCapacityByMoldDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFormationNamespaceIncreaseCapacityByMoldDistributionStatisticsFromJson(data string) FormationNamespaceIncreaseCapacityByMoldDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationNamespaceIncreaseCapacityByMoldDistributionStatisticsFromDict(dict)
}

func NewFormationNamespaceIncreaseCapacityByMoldDistributionStatisticsFromDict(data map[string]interface{}) FormationNamespaceIncreaseCapacityByMoldDistributionStatistics {
	return FormationNamespaceIncreaseCapacityByMoldDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FormationNamespaceIncreaseCapacityByMoldDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FormationNamespaceIncreaseCapacityByMoldDistributionStatistics) Pointer() *FormationNamespaceIncreaseCapacityByMoldDistributionStatistics {
	return &p
}

func CastFormationNamespaceIncreaseCapacityByMoldDistributionStatisticses(data []interface{}) []FormationNamespaceIncreaseCapacityByMoldDistributionStatistics {
	v := make([]FormationNamespaceIncreaseCapacityByMoldDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFormationNamespaceIncreaseCapacityByMoldDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationNamespaceIncreaseCapacityByMoldDistributionStatisticsesFromDict(data []FormationNamespaceIncreaseCapacityByMoldDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationNamespaceIncreaseCapacityByMoldDistributionSegment struct {
	MoldModelName *string `json:"moldModelName"`
	Count         *int64  `json:"count"`
}

func NewFormationNamespaceIncreaseCapacityByMoldDistributionSegmentFromJson(data string) FormationNamespaceIncreaseCapacityByMoldDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationNamespaceIncreaseCapacityByMoldDistributionSegmentFromDict(dict)
}

func NewFormationNamespaceIncreaseCapacityByMoldDistributionSegmentFromDict(data map[string]interface{}) FormationNamespaceIncreaseCapacityByMoldDistributionSegment {
	return FormationNamespaceIncreaseCapacityByMoldDistributionSegment{
		MoldModelName: core.CastString(data["moldModelName"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p FormationNamespaceIncreaseCapacityByMoldDistributionSegment) ToDict() map[string]interface{} {

	var moldModelName *string
	if p.MoldModelName != nil {
		moldModelName = p.MoldModelName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"moldModelName": moldModelName,
		"count":         count,
	}
}

func (p FormationNamespaceIncreaseCapacityByMoldDistributionSegment) Pointer() *FormationNamespaceIncreaseCapacityByMoldDistributionSegment {
	return &p
}

func CastFormationNamespaceIncreaseCapacityByMoldDistributionSegments(data []interface{}) []FormationNamespaceIncreaseCapacityByMoldDistributionSegment {
	v := make([]FormationNamespaceIncreaseCapacityByMoldDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFormationNamespaceIncreaseCapacityByMoldDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationNamespaceIncreaseCapacityByMoldDistributionSegmentsFromDict(data []FormationNamespaceIncreaseCapacityByMoldDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationNamespaceIncreaseCapacityByMoldDistribution struct {
	Statistics   *FormationNamespaceIncreaseCapacityByMoldDistributionStatistics `json:"statistics"`
	Distribution []FormationNamespaceIncreaseCapacityByMoldDistributionSegment   `json:"distribution"`
}

func NewFormationNamespaceIncreaseCapacityByMoldDistributionFromJson(data string) FormationNamespaceIncreaseCapacityByMoldDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationNamespaceIncreaseCapacityByMoldDistributionFromDict(dict)
}

func NewFormationNamespaceIncreaseCapacityByMoldDistributionFromDict(data map[string]interface{}) FormationNamespaceIncreaseCapacityByMoldDistribution {
	return FormationNamespaceIncreaseCapacityByMoldDistribution{
		Statistics:   NewFormationNamespaceIncreaseCapacityByMoldDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFormationNamespaceIncreaseCapacityByMoldDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FormationNamespaceIncreaseCapacityByMoldDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFormationNamespaceIncreaseCapacityByMoldDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FormationNamespaceIncreaseCapacityByMoldDistribution) Pointer() *FormationNamespaceIncreaseCapacityByMoldDistribution {
	return &p
}

func CastFormationNamespaceIncreaseCapacityByMoldDistributions(data []interface{}) []FormationNamespaceIncreaseCapacityByMoldDistribution {
	v := make([]FormationNamespaceIncreaseCapacityByMoldDistribution, 0)
	for _, d := range data {
		v = append(v, NewFormationNamespaceIncreaseCapacityByMoldDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationNamespaceIncreaseCapacityByMoldDistributionsFromDict(data []FormationNamespaceIncreaseCapacityByMoldDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationNamespaceDistributions struct {
	UpdateByMold           *FormationNamespaceUpdateByMoldDistribution           `json:"updateByMold"`
	IncreaseCapacityByMold *FormationNamespaceIncreaseCapacityByMoldDistribution `json:"increaseCapacityByMold"`
}

func NewFormationNamespaceDistributionsFromJson(data string) FormationNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationNamespaceDistributionsFromDict(dict)
}

func NewFormationNamespaceDistributionsFromDict(data map[string]interface{}) FormationNamespaceDistributions {
	return FormationNamespaceDistributions{
		UpdateByMold:           NewFormationNamespaceUpdateByMoldDistributionFromDict(core.CastMap(data["updateByMold"])).Pointer(),
		IncreaseCapacityByMold: NewFormationNamespaceIncreaseCapacityByMoldDistributionFromDict(core.CastMap(data["increaseCapacityByMold"])).Pointer(),
	}
}

func (p FormationNamespaceDistributions) ToDict() map[string]interface{} {

	var updateByMold map[string]interface{}
	if p.UpdateByMold != nil {
		updateByMold = p.UpdateByMold.ToDict()
	}
	var increaseCapacityByMold map[string]interface{}
	if p.IncreaseCapacityByMold != nil {
		increaseCapacityByMold = p.IncreaseCapacityByMold.ToDict()
	}
	return map[string]interface{}{
		"updateByMold":           updateByMold,
		"increaseCapacityByMold": increaseCapacityByMold,
	}
}

func (p FormationNamespaceDistributions) Pointer() *FormationNamespaceDistributions {
	return &p
}

func CastFormationNamespaceDistributionses(data []interface{}) []FormationNamespaceDistributions {
	v := make([]FormationNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewFormationNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationNamespaceDistributionsesFromDict(data []FormationNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormationNamespace struct {
	NamespaceId   *string                          `json:"namespaceId"`
	Year          *int32                           `json:"year"`
	Month         *int32                           `json:"month"`
	Day           *int32                           `json:"day"`
	NamespaceName *string                          `json:"namespaceName"`
	Statistics    *FormationNamespaceStatistics    `json:"statistics"`
	Distributions *FormationNamespaceDistributions `json:"distributions"`
	Molds         []FormationMold                  `json:"molds"`
}

func NewFormationNamespaceFromJson(data string) FormationNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormationNamespaceFromDict(dict)
}

func NewFormationNamespaceFromDict(data map[string]interface{}) FormationNamespace {
	return FormationNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewFormationNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewFormationNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		Molds:         CastFormationMolds(core.CastArray(data["molds"])),
	}
}

func (p FormationNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var molds []interface{}
	if p.Molds != nil {
		molds = CastFormationMoldsFromDict(
			p.Molds,
		)
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
		"molds":         molds,
	}
}

func (p FormationNamespace) Pointer() *FormationNamespace {
	return &p
}

func CastFormationNamespaces(data []interface{}) []FormationNamespace {
	v := make([]FormationNamespace, 0)
	for _, d := range data {
		v = append(v, NewFormationNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormationNamespacesFromDict(data []FormationNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceStatistics struct {
	Accept      *int64 `json:"accept"`
	Reject      *int64 `json:"reject"`
	SendRequest *int64 `json:"sendRequest"`
	Follow      *int64 `json:"follow"`
}

func NewFriendNamespaceStatisticsFromJson(data string) FriendNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceStatisticsFromDict(dict)
}

func NewFriendNamespaceStatisticsFromDict(data map[string]interface{}) FriendNamespaceStatistics {
	return FriendNamespaceStatistics{
		Accept:      core.CastInt64(data["accept"]),
		Reject:      core.CastInt64(data["reject"]),
		SendRequest: core.CastInt64(data["sendRequest"]),
		Follow:      core.CastInt64(data["follow"]),
	}
}

func (p FriendNamespaceStatistics) ToDict() map[string]interface{} {

	var accept *int64
	if p.Accept != nil {
		accept = p.Accept
	}
	var reject *int64
	if p.Reject != nil {
		reject = p.Reject
	}
	var sendRequest *int64
	if p.SendRequest != nil {
		sendRequest = p.SendRequest
	}
	var follow *int64
	if p.Follow != nil {
		follow = p.Follow
	}
	return map[string]interface{}{
		"accept":      accept,
		"reject":      reject,
		"sendRequest": sendRequest,
		"follow":      follow,
	}
}

func (p FriendNamespaceStatistics) Pointer() *FriendNamespaceStatistics {
	return &p
}

func CastFriendNamespaceStatisticses(data []interface{}) []FriendNamespaceStatistics {
	v := make([]FriendNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceStatisticsesFromDict(data []FriendNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceAcceptByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFriendNamespaceAcceptByUserDistributionStatisticsFromJson(data string) FriendNamespaceAcceptByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceAcceptByUserDistributionStatisticsFromDict(dict)
}

func NewFriendNamespaceAcceptByUserDistributionStatisticsFromDict(data map[string]interface{}) FriendNamespaceAcceptByUserDistributionStatistics {
	return FriendNamespaceAcceptByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FriendNamespaceAcceptByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FriendNamespaceAcceptByUserDistributionStatistics) Pointer() *FriendNamespaceAcceptByUserDistributionStatistics {
	return &p
}

func CastFriendNamespaceAcceptByUserDistributionStatisticses(data []interface{}) []FriendNamespaceAcceptByUserDistributionStatistics {
	v := make([]FriendNamespaceAcceptByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceAcceptByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceAcceptByUserDistributionStatisticsesFromDict(data []FriendNamespaceAcceptByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceAcceptByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewFriendNamespaceAcceptByUserDistributionSegmentFromJson(data string) FriendNamespaceAcceptByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceAcceptByUserDistributionSegmentFromDict(dict)
}

func NewFriendNamespaceAcceptByUserDistributionSegmentFromDict(data map[string]interface{}) FriendNamespaceAcceptByUserDistributionSegment {
	return FriendNamespaceAcceptByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p FriendNamespaceAcceptByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p FriendNamespaceAcceptByUserDistributionSegment) Pointer() *FriendNamespaceAcceptByUserDistributionSegment {
	return &p
}

func CastFriendNamespaceAcceptByUserDistributionSegments(data []interface{}) []FriendNamespaceAcceptByUserDistributionSegment {
	v := make([]FriendNamespaceAcceptByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceAcceptByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceAcceptByUserDistributionSegmentsFromDict(data []FriendNamespaceAcceptByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceAcceptByUserDistribution struct {
	Statistics   *FriendNamespaceAcceptByUserDistributionStatistics `json:"statistics"`
	Distribution []FriendNamespaceAcceptByUserDistributionSegment   `json:"distribution"`
}

func NewFriendNamespaceAcceptByUserDistributionFromJson(data string) FriendNamespaceAcceptByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceAcceptByUserDistributionFromDict(dict)
}

func NewFriendNamespaceAcceptByUserDistributionFromDict(data map[string]interface{}) FriendNamespaceAcceptByUserDistribution {
	return FriendNamespaceAcceptByUserDistribution{
		Statistics:   NewFriendNamespaceAcceptByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFriendNamespaceAcceptByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FriendNamespaceAcceptByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFriendNamespaceAcceptByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FriendNamespaceAcceptByUserDistribution) Pointer() *FriendNamespaceAcceptByUserDistribution {
	return &p
}

func CastFriendNamespaceAcceptByUserDistributions(data []interface{}) []FriendNamespaceAcceptByUserDistribution {
	v := make([]FriendNamespaceAcceptByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceAcceptByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceAcceptByUserDistributionsFromDict(data []FriendNamespaceAcceptByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceRejectByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFriendNamespaceRejectByUserDistributionStatisticsFromJson(data string) FriendNamespaceRejectByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceRejectByUserDistributionStatisticsFromDict(dict)
}

func NewFriendNamespaceRejectByUserDistributionStatisticsFromDict(data map[string]interface{}) FriendNamespaceRejectByUserDistributionStatistics {
	return FriendNamespaceRejectByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FriendNamespaceRejectByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FriendNamespaceRejectByUserDistributionStatistics) Pointer() *FriendNamespaceRejectByUserDistributionStatistics {
	return &p
}

func CastFriendNamespaceRejectByUserDistributionStatisticses(data []interface{}) []FriendNamespaceRejectByUserDistributionStatistics {
	v := make([]FriendNamespaceRejectByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceRejectByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceRejectByUserDistributionStatisticsesFromDict(data []FriendNamespaceRejectByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceRejectByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewFriendNamespaceRejectByUserDistributionSegmentFromJson(data string) FriendNamespaceRejectByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceRejectByUserDistributionSegmentFromDict(dict)
}

func NewFriendNamespaceRejectByUserDistributionSegmentFromDict(data map[string]interface{}) FriendNamespaceRejectByUserDistributionSegment {
	return FriendNamespaceRejectByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p FriendNamespaceRejectByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p FriendNamespaceRejectByUserDistributionSegment) Pointer() *FriendNamespaceRejectByUserDistributionSegment {
	return &p
}

func CastFriendNamespaceRejectByUserDistributionSegments(data []interface{}) []FriendNamespaceRejectByUserDistributionSegment {
	v := make([]FriendNamespaceRejectByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceRejectByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceRejectByUserDistributionSegmentsFromDict(data []FriendNamespaceRejectByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceRejectByUserDistribution struct {
	Statistics   *FriendNamespaceRejectByUserDistributionStatistics `json:"statistics"`
	Distribution []FriendNamespaceRejectByUserDistributionSegment   `json:"distribution"`
}

func NewFriendNamespaceRejectByUserDistributionFromJson(data string) FriendNamespaceRejectByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceRejectByUserDistributionFromDict(dict)
}

func NewFriendNamespaceRejectByUserDistributionFromDict(data map[string]interface{}) FriendNamespaceRejectByUserDistribution {
	return FriendNamespaceRejectByUserDistribution{
		Statistics:   NewFriendNamespaceRejectByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFriendNamespaceRejectByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FriendNamespaceRejectByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFriendNamespaceRejectByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FriendNamespaceRejectByUserDistribution) Pointer() *FriendNamespaceRejectByUserDistribution {
	return &p
}

func CastFriendNamespaceRejectByUserDistributions(data []interface{}) []FriendNamespaceRejectByUserDistribution {
	v := make([]FriendNamespaceRejectByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceRejectByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceRejectByUserDistributionsFromDict(data []FriendNamespaceRejectByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceSendRequestByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFriendNamespaceSendRequestByUserDistributionStatisticsFromJson(data string) FriendNamespaceSendRequestByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceSendRequestByUserDistributionStatisticsFromDict(dict)
}

func NewFriendNamespaceSendRequestByUserDistributionStatisticsFromDict(data map[string]interface{}) FriendNamespaceSendRequestByUserDistributionStatistics {
	return FriendNamespaceSendRequestByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FriendNamespaceSendRequestByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FriendNamespaceSendRequestByUserDistributionStatistics) Pointer() *FriendNamespaceSendRequestByUserDistributionStatistics {
	return &p
}

func CastFriendNamespaceSendRequestByUserDistributionStatisticses(data []interface{}) []FriendNamespaceSendRequestByUserDistributionStatistics {
	v := make([]FriendNamespaceSendRequestByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceSendRequestByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceSendRequestByUserDistributionStatisticsesFromDict(data []FriendNamespaceSendRequestByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceSendRequestByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewFriendNamespaceSendRequestByUserDistributionSegmentFromJson(data string) FriendNamespaceSendRequestByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceSendRequestByUserDistributionSegmentFromDict(dict)
}

func NewFriendNamespaceSendRequestByUserDistributionSegmentFromDict(data map[string]interface{}) FriendNamespaceSendRequestByUserDistributionSegment {
	return FriendNamespaceSendRequestByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p FriendNamespaceSendRequestByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p FriendNamespaceSendRequestByUserDistributionSegment) Pointer() *FriendNamespaceSendRequestByUserDistributionSegment {
	return &p
}

func CastFriendNamespaceSendRequestByUserDistributionSegments(data []interface{}) []FriendNamespaceSendRequestByUserDistributionSegment {
	v := make([]FriendNamespaceSendRequestByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceSendRequestByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceSendRequestByUserDistributionSegmentsFromDict(data []FriendNamespaceSendRequestByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceSendRequestByUserDistribution struct {
	Statistics   *FriendNamespaceSendRequestByUserDistributionStatistics `json:"statistics"`
	Distribution []FriendNamespaceSendRequestByUserDistributionSegment   `json:"distribution"`
}

func NewFriendNamespaceSendRequestByUserDistributionFromJson(data string) FriendNamespaceSendRequestByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceSendRequestByUserDistributionFromDict(dict)
}

func NewFriendNamespaceSendRequestByUserDistributionFromDict(data map[string]interface{}) FriendNamespaceSendRequestByUserDistribution {
	return FriendNamespaceSendRequestByUserDistribution{
		Statistics:   NewFriendNamespaceSendRequestByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFriendNamespaceSendRequestByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FriendNamespaceSendRequestByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFriendNamespaceSendRequestByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FriendNamespaceSendRequestByUserDistribution) Pointer() *FriendNamespaceSendRequestByUserDistribution {
	return &p
}

func CastFriendNamespaceSendRequestByUserDistributions(data []interface{}) []FriendNamespaceSendRequestByUserDistribution {
	v := make([]FriendNamespaceSendRequestByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceSendRequestByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceSendRequestByUserDistributionsFromDict(data []FriendNamespaceSendRequestByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceNewFollowByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewFriendNamespaceNewFollowByUserDistributionStatisticsFromJson(data string) FriendNamespaceNewFollowByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceNewFollowByUserDistributionStatisticsFromDict(dict)
}

func NewFriendNamespaceNewFollowByUserDistributionStatisticsFromDict(data map[string]interface{}) FriendNamespaceNewFollowByUserDistributionStatistics {
	return FriendNamespaceNewFollowByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p FriendNamespaceNewFollowByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p FriendNamespaceNewFollowByUserDistributionStatistics) Pointer() *FriendNamespaceNewFollowByUserDistributionStatistics {
	return &p
}

func CastFriendNamespaceNewFollowByUserDistributionStatisticses(data []interface{}) []FriendNamespaceNewFollowByUserDistributionStatistics {
	v := make([]FriendNamespaceNewFollowByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceNewFollowByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceNewFollowByUserDistributionStatisticsesFromDict(data []FriendNamespaceNewFollowByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceNewFollowByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewFriendNamespaceNewFollowByUserDistributionSegmentFromJson(data string) FriendNamespaceNewFollowByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceNewFollowByUserDistributionSegmentFromDict(dict)
}

func NewFriendNamespaceNewFollowByUserDistributionSegmentFromDict(data map[string]interface{}) FriendNamespaceNewFollowByUserDistributionSegment {
	return FriendNamespaceNewFollowByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p FriendNamespaceNewFollowByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p FriendNamespaceNewFollowByUserDistributionSegment) Pointer() *FriendNamespaceNewFollowByUserDistributionSegment {
	return &p
}

func CastFriendNamespaceNewFollowByUserDistributionSegments(data []interface{}) []FriendNamespaceNewFollowByUserDistributionSegment {
	v := make([]FriendNamespaceNewFollowByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceNewFollowByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceNewFollowByUserDistributionSegmentsFromDict(data []FriendNamespaceNewFollowByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceNewFollowByUserDistribution struct {
	Statistics   *FriendNamespaceNewFollowByUserDistributionStatistics `json:"statistics"`
	Distribution []FriendNamespaceNewFollowByUserDistributionSegment   `json:"distribution"`
}

func NewFriendNamespaceNewFollowByUserDistributionFromJson(data string) FriendNamespaceNewFollowByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceNewFollowByUserDistributionFromDict(dict)
}

func NewFriendNamespaceNewFollowByUserDistributionFromDict(data map[string]interface{}) FriendNamespaceNewFollowByUserDistribution {
	return FriendNamespaceNewFollowByUserDistribution{
		Statistics:   NewFriendNamespaceNewFollowByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastFriendNamespaceNewFollowByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p FriendNamespaceNewFollowByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastFriendNamespaceNewFollowByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p FriendNamespaceNewFollowByUserDistribution) Pointer() *FriendNamespaceNewFollowByUserDistribution {
	return &p
}

func CastFriendNamespaceNewFollowByUserDistributions(data []interface{}) []FriendNamespaceNewFollowByUserDistribution {
	v := make([]FriendNamespaceNewFollowByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceNewFollowByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceNewFollowByUserDistributionsFromDict(data []FriendNamespaceNewFollowByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespaceDistributions struct {
	AcceptByUser      *FriendNamespaceAcceptByUserDistribution      `json:"acceptByUser"`
	RejectByUser      *FriendNamespaceRejectByUserDistribution      `json:"rejectByUser"`
	SendRequestByUser *FriendNamespaceSendRequestByUserDistribution `json:"sendRequestByUser"`
	NewFollowByUser   *FriendNamespaceNewFollowByUserDistribution   `json:"newFollowByUser"`
}

func NewFriendNamespaceDistributionsFromJson(data string) FriendNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceDistributionsFromDict(dict)
}

func NewFriendNamespaceDistributionsFromDict(data map[string]interface{}) FriendNamespaceDistributions {
	return FriendNamespaceDistributions{
		AcceptByUser:      NewFriendNamespaceAcceptByUserDistributionFromDict(core.CastMap(data["acceptByUser"])).Pointer(),
		RejectByUser:      NewFriendNamespaceRejectByUserDistributionFromDict(core.CastMap(data["rejectByUser"])).Pointer(),
		SendRequestByUser: NewFriendNamespaceSendRequestByUserDistributionFromDict(core.CastMap(data["sendRequestByUser"])).Pointer(),
		NewFollowByUser:   NewFriendNamespaceNewFollowByUserDistributionFromDict(core.CastMap(data["newFollowByUser"])).Pointer(),
	}
}

func (p FriendNamespaceDistributions) ToDict() map[string]interface{} {

	var acceptByUser map[string]interface{}
	if p.AcceptByUser != nil {
		acceptByUser = p.AcceptByUser.ToDict()
	}
	var rejectByUser map[string]interface{}
	if p.RejectByUser != nil {
		rejectByUser = p.RejectByUser.ToDict()
	}
	var sendRequestByUser map[string]interface{}
	if p.SendRequestByUser != nil {
		sendRequestByUser = p.SendRequestByUser.ToDict()
	}
	var newFollowByUser map[string]interface{}
	if p.NewFollowByUser != nil {
		newFollowByUser = p.NewFollowByUser.ToDict()
	}
	return map[string]interface{}{
		"acceptByUser":      acceptByUser,
		"rejectByUser":      rejectByUser,
		"sendRequestByUser": sendRequestByUser,
		"newFollowByUser":   newFollowByUser,
	}
}

func (p FriendNamespaceDistributions) Pointer() *FriendNamespaceDistributions {
	return &p
}

func CastFriendNamespaceDistributionses(data []interface{}) []FriendNamespaceDistributions {
	v := make([]FriendNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespaceDistributionsesFromDict(data []FriendNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendNamespace struct {
	NamespaceId   *string                       `json:"namespaceId"`
	Year          *int32                        `json:"year"`
	Month         *int32                        `json:"month"`
	Day           *int32                        `json:"day"`
	NamespaceName *string                       `json:"namespaceName"`
	Statistics    *FriendNamespaceStatistics    `json:"statistics"`
	Distributions *FriendNamespaceDistributions `json:"distributions"`
}

func NewFriendNamespaceFromJson(data string) FriendNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendNamespaceFromDict(dict)
}

func NewFriendNamespaceFromDict(data map[string]interface{}) FriendNamespace {
	return FriendNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewFriendNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewFriendNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p FriendNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p FriendNamespace) Pointer() *FriendNamespace {
	return &p
}

func CastFriendNamespaces(data []interface{}) []FriendNamespace {
	v := make([]FriendNamespace, 0)
	for _, d := range data {
		v = append(v, NewFriendNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendNamespacesFromDict(data []FriendNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InboxNamespaceStatistics struct {
	Sent     *int64   `json:"sent"`
	Open     *int64   `json:"open"`
	OpenRate *float32 `json:"openRate"`
}

func NewInboxNamespaceStatisticsFromJson(data string) InboxNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxNamespaceStatisticsFromDict(dict)
}

func NewInboxNamespaceStatisticsFromDict(data map[string]interface{}) InboxNamespaceStatistics {
	return InboxNamespaceStatistics{
		Sent:     core.CastInt64(data["sent"]),
		Open:     core.CastInt64(data["open"]),
		OpenRate: core.CastFloat32(data["openRate"]),
	}
}

func (p InboxNamespaceStatistics) ToDict() map[string]interface{} {

	var sent *int64
	if p.Sent != nil {
		sent = p.Sent
	}
	var open *int64
	if p.Open != nil {
		open = p.Open
	}
	var openRate *float32
	if p.OpenRate != nil {
		openRate = p.OpenRate
	}
	return map[string]interface{}{
		"sent":     sent,
		"open":     open,
		"openRate": openRate,
	}
}

func (p InboxNamespaceStatistics) Pointer() *InboxNamespaceStatistics {
	return &p
}

func CastInboxNamespaceStatisticses(data []interface{}) []InboxNamespaceStatistics {
	v := make([]InboxNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewInboxNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxNamespaceStatisticsesFromDict(data []InboxNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InboxNamespaceSendByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInboxNamespaceSendByUserDistributionStatisticsFromJson(data string) InboxNamespaceSendByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxNamespaceSendByUserDistributionStatisticsFromDict(dict)
}

func NewInboxNamespaceSendByUserDistributionStatisticsFromDict(data map[string]interface{}) InboxNamespaceSendByUserDistributionStatistics {
	return InboxNamespaceSendByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InboxNamespaceSendByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InboxNamespaceSendByUserDistributionStatistics) Pointer() *InboxNamespaceSendByUserDistributionStatistics {
	return &p
}

func CastInboxNamespaceSendByUserDistributionStatisticses(data []interface{}) []InboxNamespaceSendByUserDistributionStatistics {
	v := make([]InboxNamespaceSendByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInboxNamespaceSendByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxNamespaceSendByUserDistributionStatisticsesFromDict(data []InboxNamespaceSendByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InboxNamespaceSendByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewInboxNamespaceSendByUserDistributionSegmentFromJson(data string) InboxNamespaceSendByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxNamespaceSendByUserDistributionSegmentFromDict(dict)
}

func NewInboxNamespaceSendByUserDistributionSegmentFromDict(data map[string]interface{}) InboxNamespaceSendByUserDistributionSegment {
	return InboxNamespaceSendByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p InboxNamespaceSendByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p InboxNamespaceSendByUserDistributionSegment) Pointer() *InboxNamespaceSendByUserDistributionSegment {
	return &p
}

func CastInboxNamespaceSendByUserDistributionSegments(data []interface{}) []InboxNamespaceSendByUserDistributionSegment {
	v := make([]InboxNamespaceSendByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInboxNamespaceSendByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxNamespaceSendByUserDistributionSegmentsFromDict(data []InboxNamespaceSendByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InboxNamespaceSendByUserDistribution struct {
	Statistics   *InboxNamespaceSendByUserDistributionStatistics `json:"statistics"`
	Distribution []InboxNamespaceSendByUserDistributionSegment   `json:"distribution"`
}

func NewInboxNamespaceSendByUserDistributionFromJson(data string) InboxNamespaceSendByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxNamespaceSendByUserDistributionFromDict(dict)
}

func NewInboxNamespaceSendByUserDistributionFromDict(data map[string]interface{}) InboxNamespaceSendByUserDistribution {
	return InboxNamespaceSendByUserDistribution{
		Statistics:   NewInboxNamespaceSendByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInboxNamespaceSendByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InboxNamespaceSendByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInboxNamespaceSendByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InboxNamespaceSendByUserDistribution) Pointer() *InboxNamespaceSendByUserDistribution {
	return &p
}

func CastInboxNamespaceSendByUserDistributions(data []interface{}) []InboxNamespaceSendByUserDistribution {
	v := make([]InboxNamespaceSendByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewInboxNamespaceSendByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxNamespaceSendByUserDistributionsFromDict(data []InboxNamespaceSendByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InboxNamespaceReadElapsedMinutesDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInboxNamespaceReadElapsedMinutesDistributionStatisticsFromJson(data string) InboxNamespaceReadElapsedMinutesDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxNamespaceReadElapsedMinutesDistributionStatisticsFromDict(dict)
}

func NewInboxNamespaceReadElapsedMinutesDistributionStatisticsFromDict(data map[string]interface{}) InboxNamespaceReadElapsedMinutesDistributionStatistics {
	return InboxNamespaceReadElapsedMinutesDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InboxNamespaceReadElapsedMinutesDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InboxNamespaceReadElapsedMinutesDistributionStatistics) Pointer() *InboxNamespaceReadElapsedMinutesDistributionStatistics {
	return &p
}

func CastInboxNamespaceReadElapsedMinutesDistributionStatisticses(data []interface{}) []InboxNamespaceReadElapsedMinutesDistributionStatistics {
	v := make([]InboxNamespaceReadElapsedMinutesDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInboxNamespaceReadElapsedMinutesDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxNamespaceReadElapsedMinutesDistributionStatisticsesFromDict(data []InboxNamespaceReadElapsedMinutesDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InboxNamespaceReadElapsedMinutesDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewInboxNamespaceReadElapsedMinutesDistributionSegmentFromJson(data string) InboxNamespaceReadElapsedMinutesDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxNamespaceReadElapsedMinutesDistributionSegmentFromDict(dict)
}

func NewInboxNamespaceReadElapsedMinutesDistributionSegmentFromDict(data map[string]interface{}) InboxNamespaceReadElapsedMinutesDistributionSegment {
	return InboxNamespaceReadElapsedMinutesDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p InboxNamespaceReadElapsedMinutesDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p InboxNamespaceReadElapsedMinutesDistributionSegment) Pointer() *InboxNamespaceReadElapsedMinutesDistributionSegment {
	return &p
}

func CastInboxNamespaceReadElapsedMinutesDistributionSegments(data []interface{}) []InboxNamespaceReadElapsedMinutesDistributionSegment {
	v := make([]InboxNamespaceReadElapsedMinutesDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInboxNamespaceReadElapsedMinutesDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxNamespaceReadElapsedMinutesDistributionSegmentsFromDict(data []InboxNamespaceReadElapsedMinutesDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InboxNamespaceReadElapsedMinutesDistribution struct {
	Statistics   *InboxNamespaceReadElapsedMinutesDistributionStatistics `json:"statistics"`
	Distribution []InboxNamespaceReadElapsedMinutesDistributionSegment   `json:"distribution"`
}

func NewInboxNamespaceReadElapsedMinutesDistributionFromJson(data string) InboxNamespaceReadElapsedMinutesDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxNamespaceReadElapsedMinutesDistributionFromDict(dict)
}

func NewInboxNamespaceReadElapsedMinutesDistributionFromDict(data map[string]interface{}) InboxNamespaceReadElapsedMinutesDistribution {
	return InboxNamespaceReadElapsedMinutesDistribution{
		Statistics:   NewInboxNamespaceReadElapsedMinutesDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInboxNamespaceReadElapsedMinutesDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InboxNamespaceReadElapsedMinutesDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInboxNamespaceReadElapsedMinutesDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InboxNamespaceReadElapsedMinutesDistribution) Pointer() *InboxNamespaceReadElapsedMinutesDistribution {
	return &p
}

func CastInboxNamespaceReadElapsedMinutesDistributions(data []interface{}) []InboxNamespaceReadElapsedMinutesDistribution {
	v := make([]InboxNamespaceReadElapsedMinutesDistribution, 0)
	for _, d := range data {
		v = append(v, NewInboxNamespaceReadElapsedMinutesDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxNamespaceReadElapsedMinutesDistributionsFromDict(data []InboxNamespaceReadElapsedMinutesDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InboxNamespaceDistributions struct {
	SendByUser         *InboxNamespaceSendByUserDistribution         `json:"sendByUser"`
	ReadElapsedMinutes *InboxNamespaceReadElapsedMinutesDistribution `json:"readElapsedMinutes"`
}

func NewInboxNamespaceDistributionsFromJson(data string) InboxNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxNamespaceDistributionsFromDict(dict)
}

func NewInboxNamespaceDistributionsFromDict(data map[string]interface{}) InboxNamespaceDistributions {
	return InboxNamespaceDistributions{
		SendByUser:         NewInboxNamespaceSendByUserDistributionFromDict(core.CastMap(data["sendByUser"])).Pointer(),
		ReadElapsedMinutes: NewInboxNamespaceReadElapsedMinutesDistributionFromDict(core.CastMap(data["readElapsedMinutes"])).Pointer(),
	}
}

func (p InboxNamespaceDistributions) ToDict() map[string]interface{} {

	var sendByUser map[string]interface{}
	if p.SendByUser != nil {
		sendByUser = p.SendByUser.ToDict()
	}
	var readElapsedMinutes map[string]interface{}
	if p.ReadElapsedMinutes != nil {
		readElapsedMinutes = p.ReadElapsedMinutes.ToDict()
	}
	return map[string]interface{}{
		"sendByUser":         sendByUser,
		"readElapsedMinutes": readElapsedMinutes,
	}
}

func (p InboxNamespaceDistributions) Pointer() *InboxNamespaceDistributions {
	return &p
}

func CastInboxNamespaceDistributionses(data []interface{}) []InboxNamespaceDistributions {
	v := make([]InboxNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewInboxNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxNamespaceDistributionsesFromDict(data []InboxNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InboxNamespace struct {
	NamespaceId   *string                      `json:"namespaceId"`
	Year          *int32                       `json:"year"`
	Month         *int32                       `json:"month"`
	Day           *int32                       `json:"day"`
	NamespaceName *string                      `json:"namespaceName"`
	Statistics    *InboxNamespaceStatistics    `json:"statistics"`
	Distributions *InboxNamespaceDistributions `json:"distributions"`
}

func NewInboxNamespaceFromJson(data string) InboxNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxNamespaceFromDict(dict)
}

func NewInboxNamespaceFromDict(data map[string]interface{}) InboxNamespace {
	return InboxNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewInboxNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewInboxNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p InboxNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p InboxNamespace) Pointer() *InboxNamespace {
	return &p
}

func CastInboxNamespaces(data []interface{}) []InboxNamespace {
	v := make([]InboxNamespace, 0)
	for _, d := range data {
		v = append(v, NewInboxNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxNamespacesFromDict(data []InboxNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryItemSetStatistics struct {
	Acquired       *int64   `json:"acquired"`
	AcquiredAmount *int64   `json:"acquiredAmount"`
	Consumed       *int64   `json:"consumed"`
	ConsumedAmount *int64   `json:"consumedAmount"`
	ConsumedRate   *float32 `json:"consumedRate"`
}

func NewInventoryItemSetStatisticsFromJson(data string) InventoryItemSetStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryItemSetStatisticsFromDict(dict)
}

func NewInventoryItemSetStatisticsFromDict(data map[string]interface{}) InventoryItemSetStatistics {
	return InventoryItemSetStatistics{
		Acquired:       core.CastInt64(data["acquired"]),
		AcquiredAmount: core.CastInt64(data["acquiredAmount"]),
		Consumed:       core.CastInt64(data["consumed"]),
		ConsumedAmount: core.CastInt64(data["consumedAmount"]),
		ConsumedRate:   core.CastFloat32(data["consumedRate"]),
	}
}

func (p InventoryItemSetStatistics) ToDict() map[string]interface{} {

	var acquired *int64
	if p.Acquired != nil {
		acquired = p.Acquired
	}
	var acquiredAmount *int64
	if p.AcquiredAmount != nil {
		acquiredAmount = p.AcquiredAmount
	}
	var consumed *int64
	if p.Consumed != nil {
		consumed = p.Consumed
	}
	var consumedAmount *int64
	if p.ConsumedAmount != nil {
		consumedAmount = p.ConsumedAmount
	}
	var consumedRate *float32
	if p.ConsumedRate != nil {
		consumedRate = p.ConsumedRate
	}
	return map[string]interface{}{
		"acquired":       acquired,
		"acquiredAmount": acquiredAmount,
		"consumed":       consumed,
		"consumedAmount": consumedAmount,
		"consumedRate":   consumedRate,
	}
}

func (p InventoryItemSetStatistics) Pointer() *InventoryItemSetStatistics {
	return &p
}

func CastInventoryItemSetStatisticses(data []interface{}) []InventoryItemSetStatistics {
	v := make([]InventoryItemSetStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryItemSetStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryItemSetStatisticsesFromDict(data []InventoryItemSetStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryItemSetCountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryItemSetCountDistributionStatisticsFromJson(data string) InventoryItemSetCountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryItemSetCountDistributionStatisticsFromDict(dict)
}

func NewInventoryItemSetCountDistributionStatisticsFromDict(data map[string]interface{}) InventoryItemSetCountDistributionStatistics {
	return InventoryItemSetCountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryItemSetCountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryItemSetCountDistributionStatistics) Pointer() *InventoryItemSetCountDistributionStatistics {
	return &p
}

func CastInventoryItemSetCountDistributionStatisticses(data []interface{}) []InventoryItemSetCountDistributionStatistics {
	v := make([]InventoryItemSetCountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryItemSetCountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryItemSetCountDistributionStatisticsesFromDict(data []InventoryItemSetCountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryItemSetCountDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewInventoryItemSetCountDistributionSegmentFromJson(data string) InventoryItemSetCountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryItemSetCountDistributionSegmentFromDict(dict)
}

func NewInventoryItemSetCountDistributionSegmentFromDict(data map[string]interface{}) InventoryItemSetCountDistributionSegment {
	return InventoryItemSetCountDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p InventoryItemSetCountDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p InventoryItemSetCountDistributionSegment) Pointer() *InventoryItemSetCountDistributionSegment {
	return &p
}

func CastInventoryItemSetCountDistributionSegments(data []interface{}) []InventoryItemSetCountDistributionSegment {
	v := make([]InventoryItemSetCountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryItemSetCountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryItemSetCountDistributionSegmentsFromDict(data []InventoryItemSetCountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryItemSetCountDistribution struct {
	Statistics   *InventoryItemSetCountDistributionStatistics `json:"statistics"`
	Distribution []InventoryItemSetCountDistributionSegment   `json:"distribution"`
}

func NewInventoryItemSetCountDistributionFromJson(data string) InventoryItemSetCountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryItemSetCountDistributionFromDict(dict)
}

func NewInventoryItemSetCountDistributionFromDict(data map[string]interface{}) InventoryItemSetCountDistribution {
	return InventoryItemSetCountDistribution{
		Statistics:   NewInventoryItemSetCountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryItemSetCountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryItemSetCountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryItemSetCountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryItemSetCountDistribution) Pointer() *InventoryItemSetCountDistribution {
	return &p
}

func CastInventoryItemSetCountDistributions(data []interface{}) []InventoryItemSetCountDistribution {
	v := make([]InventoryItemSetCountDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryItemSetCountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryItemSetCountDistributionsFromDict(data []InventoryItemSetCountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryItemSetDistributions struct {
	Count *InventoryItemSetCountDistribution `json:"count"`
}

func NewInventoryItemSetDistributionsFromJson(data string) InventoryItemSetDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryItemSetDistributionsFromDict(dict)
}

func NewInventoryItemSetDistributionsFromDict(data map[string]interface{}) InventoryItemSetDistributions {
	return InventoryItemSetDistributions{
		Count: NewInventoryItemSetCountDistributionFromDict(core.CastMap(data["count"])).Pointer(),
	}
}

func (p InventoryItemSetDistributions) ToDict() map[string]interface{} {

	var count map[string]interface{}
	if p.Count != nil {
		count = p.Count.ToDict()
	}
	return map[string]interface{}{
		"count": count,
	}
}

func (p InventoryItemSetDistributions) Pointer() *InventoryItemSetDistributions {
	return &p
}

func CastInventoryItemSetDistributionses(data []interface{}) []InventoryItemSetDistributions {
	v := make([]InventoryItemSetDistributions, 0)
	for _, d := range data {
		v = append(v, NewInventoryItemSetDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryItemSetDistributionsesFromDict(data []InventoryItemSetDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryItemSet struct {
	ItemSetId     *string                        `json:"itemSetId"`
	ItemName      *string                        `json:"itemName"`
	ItemSetName   *string                        `json:"itemSetName"`
	Statistics    *InventoryItemSetStatistics    `json:"statistics"`
	Distributions *InventoryItemSetDistributions `json:"distributions"`
}

func NewInventoryItemSetFromJson(data string) InventoryItemSet {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryItemSetFromDict(dict)
}

func NewInventoryItemSetFromDict(data map[string]interface{}) InventoryItemSet {
	return InventoryItemSet{
		ItemSetId:     core.CastString(data["itemSetId"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		Statistics:    NewInventoryItemSetStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewInventoryItemSetDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p InventoryItemSet) ToDict() map[string]interface{} {

	var itemSetId *string
	if p.ItemSetId != nil {
		itemSetId = p.ItemSetId
	}
	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var itemSetName *string
	if p.ItemSetName != nil {
		itemSetName = p.ItemSetName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"itemSetId":     itemSetId,
		"itemName":      itemName,
		"itemSetName":   itemSetName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p InventoryItemSet) Pointer() *InventoryItemSet {
	return &p
}

func CastInventoryItemSets(data []interface{}) []InventoryItemSet {
	v := make([]InventoryItemSet, 0)
	for _, d := range data {
		v = append(v, NewInventoryItemSetFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryItemSetsFromDict(data []InventoryItemSet) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryStatistics struct {
	Acquired               *int64 `json:"acquired"`
	Consume                *int64 `json:"consume"`
	IncreaseCapacity       *int64 `json:"increaseCapacity"`
	IncreaseCapacityAmount *int64 `json:"increaseCapacityAmount"`
}

func NewInventoryInventoryStatisticsFromJson(data string) InventoryInventoryStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryStatisticsFromDict(dict)
}

func NewInventoryInventoryStatisticsFromDict(data map[string]interface{}) InventoryInventoryStatistics {
	return InventoryInventoryStatistics{
		Acquired:               core.CastInt64(data["acquired"]),
		Consume:                core.CastInt64(data["consume"]),
		IncreaseCapacity:       core.CastInt64(data["increaseCapacity"]),
		IncreaseCapacityAmount: core.CastInt64(data["increaseCapacityAmount"]),
	}
}

func (p InventoryInventoryStatistics) ToDict() map[string]interface{} {

	var acquired *int64
	if p.Acquired != nil {
		acquired = p.Acquired
	}
	var consume *int64
	if p.Consume != nil {
		consume = p.Consume
	}
	var increaseCapacity *int64
	if p.IncreaseCapacity != nil {
		increaseCapacity = p.IncreaseCapacity
	}
	var increaseCapacityAmount *int64
	if p.IncreaseCapacityAmount != nil {
		increaseCapacityAmount = p.IncreaseCapacityAmount
	}
	return map[string]interface{}{
		"acquired":               acquired,
		"consume":                consume,
		"increaseCapacity":       increaseCapacity,
		"increaseCapacityAmount": increaseCapacityAmount,
	}
}

func (p InventoryInventoryStatistics) Pointer() *InventoryInventoryStatistics {
	return &p
}

func CastInventoryInventoryStatisticses(data []interface{}) []InventoryInventoryStatistics {
	v := make([]InventoryInventoryStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryStatisticsesFromDict(data []InventoryInventoryStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryCapacityDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryInventoryCapacityDistributionStatisticsFromJson(data string) InventoryInventoryCapacityDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryCapacityDistributionStatisticsFromDict(dict)
}

func NewInventoryInventoryCapacityDistributionStatisticsFromDict(data map[string]interface{}) InventoryInventoryCapacityDistributionStatistics {
	return InventoryInventoryCapacityDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryInventoryCapacityDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryInventoryCapacityDistributionStatistics) Pointer() *InventoryInventoryCapacityDistributionStatistics {
	return &p
}

func CastInventoryInventoryCapacityDistributionStatisticses(data []interface{}) []InventoryInventoryCapacityDistributionStatistics {
	v := make([]InventoryInventoryCapacityDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryCapacityDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryCapacityDistributionStatisticsesFromDict(data []InventoryInventoryCapacityDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryCapacityDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewInventoryInventoryCapacityDistributionSegmentFromJson(data string) InventoryInventoryCapacityDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryCapacityDistributionSegmentFromDict(dict)
}

func NewInventoryInventoryCapacityDistributionSegmentFromDict(data map[string]interface{}) InventoryInventoryCapacityDistributionSegment {
	return InventoryInventoryCapacityDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p InventoryInventoryCapacityDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p InventoryInventoryCapacityDistributionSegment) Pointer() *InventoryInventoryCapacityDistributionSegment {
	return &p
}

func CastInventoryInventoryCapacityDistributionSegments(data []interface{}) []InventoryInventoryCapacityDistributionSegment {
	v := make([]InventoryInventoryCapacityDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryCapacityDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryCapacityDistributionSegmentsFromDict(data []InventoryInventoryCapacityDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryCapacityDistribution struct {
	Statistics   *InventoryInventoryCapacityDistributionStatistics `json:"statistics"`
	Distribution []InventoryInventoryCapacityDistributionSegment   `json:"distribution"`
}

func NewInventoryInventoryCapacityDistributionFromJson(data string) InventoryInventoryCapacityDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryCapacityDistributionFromDict(dict)
}

func NewInventoryInventoryCapacityDistributionFromDict(data map[string]interface{}) InventoryInventoryCapacityDistribution {
	return InventoryInventoryCapacityDistribution{
		Statistics:   NewInventoryInventoryCapacityDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryInventoryCapacityDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryInventoryCapacityDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryInventoryCapacityDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryInventoryCapacityDistribution) Pointer() *InventoryInventoryCapacityDistribution {
	return &p
}

func CastInventoryInventoryCapacityDistributions(data []interface{}) []InventoryInventoryCapacityDistribution {
	v := make([]InventoryInventoryCapacityDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryCapacityDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryCapacityDistributionsFromDict(data []InventoryInventoryCapacityDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryAcquireDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryInventoryAcquireDistributionStatisticsFromJson(data string) InventoryInventoryAcquireDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryAcquireDistributionStatisticsFromDict(dict)
}

func NewInventoryInventoryAcquireDistributionStatisticsFromDict(data map[string]interface{}) InventoryInventoryAcquireDistributionStatistics {
	return InventoryInventoryAcquireDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryInventoryAcquireDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryInventoryAcquireDistributionStatistics) Pointer() *InventoryInventoryAcquireDistributionStatistics {
	return &p
}

func CastInventoryInventoryAcquireDistributionStatisticses(data []interface{}) []InventoryInventoryAcquireDistributionStatistics {
	v := make([]InventoryInventoryAcquireDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryAcquireDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryAcquireDistributionStatisticsesFromDict(data []InventoryInventoryAcquireDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryAcquireDistributionSegment struct {
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
}

func NewInventoryInventoryAcquireDistributionSegmentFromJson(data string) InventoryInventoryAcquireDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryAcquireDistributionSegmentFromDict(dict)
}

func NewInventoryInventoryAcquireDistributionSegmentFromDict(data map[string]interface{}) InventoryInventoryAcquireDistributionSegment {
	return InventoryInventoryAcquireDistributionSegment{
		ItemName: core.CastString(data["itemName"]),
		Count:    core.CastInt64(data["count"]),
	}
}

func (p InventoryInventoryAcquireDistributionSegment) ToDict() map[string]interface{} {

	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"itemName": itemName,
		"count":    count,
	}
}

func (p InventoryInventoryAcquireDistributionSegment) Pointer() *InventoryInventoryAcquireDistributionSegment {
	return &p
}

func CastInventoryInventoryAcquireDistributionSegments(data []interface{}) []InventoryInventoryAcquireDistributionSegment {
	v := make([]InventoryInventoryAcquireDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryAcquireDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryAcquireDistributionSegmentsFromDict(data []InventoryInventoryAcquireDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryAcquireDistribution struct {
	Statistics   *InventoryInventoryAcquireDistributionStatistics `json:"statistics"`
	Distribution []InventoryInventoryAcquireDistributionSegment   `json:"distribution"`
}

func NewInventoryInventoryAcquireDistributionFromJson(data string) InventoryInventoryAcquireDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryAcquireDistributionFromDict(dict)
}

func NewInventoryInventoryAcquireDistributionFromDict(data map[string]interface{}) InventoryInventoryAcquireDistribution {
	return InventoryInventoryAcquireDistribution{
		Statistics:   NewInventoryInventoryAcquireDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryInventoryAcquireDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryInventoryAcquireDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryInventoryAcquireDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryInventoryAcquireDistribution) Pointer() *InventoryInventoryAcquireDistribution {
	return &p
}

func CastInventoryInventoryAcquireDistributions(data []interface{}) []InventoryInventoryAcquireDistribution {
	v := make([]InventoryInventoryAcquireDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryAcquireDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryAcquireDistributionsFromDict(data []InventoryInventoryAcquireDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryAcquireAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryInventoryAcquireAmountDistributionStatisticsFromJson(data string) InventoryInventoryAcquireAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryAcquireAmountDistributionStatisticsFromDict(dict)
}

func NewInventoryInventoryAcquireAmountDistributionStatisticsFromDict(data map[string]interface{}) InventoryInventoryAcquireAmountDistributionStatistics {
	return InventoryInventoryAcquireAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryInventoryAcquireAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryInventoryAcquireAmountDistributionStatistics) Pointer() *InventoryInventoryAcquireAmountDistributionStatistics {
	return &p
}

func CastInventoryInventoryAcquireAmountDistributionStatisticses(data []interface{}) []InventoryInventoryAcquireAmountDistributionStatistics {
	v := make([]InventoryInventoryAcquireAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryAcquireAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryAcquireAmountDistributionStatisticsesFromDict(data []InventoryInventoryAcquireAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryAcquireAmountDistributionSegment struct {
	ItemName *string `json:"itemName"`
	Sum      *int64  `json:"sum"`
}

func NewInventoryInventoryAcquireAmountDistributionSegmentFromJson(data string) InventoryInventoryAcquireAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryAcquireAmountDistributionSegmentFromDict(dict)
}

func NewInventoryInventoryAcquireAmountDistributionSegmentFromDict(data map[string]interface{}) InventoryInventoryAcquireAmountDistributionSegment {
	return InventoryInventoryAcquireAmountDistributionSegment{
		ItemName: core.CastString(data["itemName"]),
		Sum:      core.CastInt64(data["sum"]),
	}
}

func (p InventoryInventoryAcquireAmountDistributionSegment) ToDict() map[string]interface{} {

	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"itemName": itemName,
		"sum":      sum,
	}
}

func (p InventoryInventoryAcquireAmountDistributionSegment) Pointer() *InventoryInventoryAcquireAmountDistributionSegment {
	return &p
}

func CastInventoryInventoryAcquireAmountDistributionSegments(data []interface{}) []InventoryInventoryAcquireAmountDistributionSegment {
	v := make([]InventoryInventoryAcquireAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryAcquireAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryAcquireAmountDistributionSegmentsFromDict(data []InventoryInventoryAcquireAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryAcquireAmountDistribution struct {
	Statistics   *InventoryInventoryAcquireAmountDistributionStatistics `json:"statistics"`
	Distribution []InventoryInventoryAcquireAmountDistributionSegment   `json:"distribution"`
}

func NewInventoryInventoryAcquireAmountDistributionFromJson(data string) InventoryInventoryAcquireAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryAcquireAmountDistributionFromDict(dict)
}

func NewInventoryInventoryAcquireAmountDistributionFromDict(data map[string]interface{}) InventoryInventoryAcquireAmountDistribution {
	return InventoryInventoryAcquireAmountDistribution{
		Statistics:   NewInventoryInventoryAcquireAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryInventoryAcquireAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryInventoryAcquireAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryInventoryAcquireAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryInventoryAcquireAmountDistribution) Pointer() *InventoryInventoryAcquireAmountDistribution {
	return &p
}

func CastInventoryInventoryAcquireAmountDistributions(data []interface{}) []InventoryInventoryAcquireAmountDistribution {
	v := make([]InventoryInventoryAcquireAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryAcquireAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryAcquireAmountDistributionsFromDict(data []InventoryInventoryAcquireAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryConsumeDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryInventoryConsumeDistributionStatisticsFromJson(data string) InventoryInventoryConsumeDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryConsumeDistributionStatisticsFromDict(dict)
}

func NewInventoryInventoryConsumeDistributionStatisticsFromDict(data map[string]interface{}) InventoryInventoryConsumeDistributionStatistics {
	return InventoryInventoryConsumeDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryInventoryConsumeDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryInventoryConsumeDistributionStatistics) Pointer() *InventoryInventoryConsumeDistributionStatistics {
	return &p
}

func CastInventoryInventoryConsumeDistributionStatisticses(data []interface{}) []InventoryInventoryConsumeDistributionStatistics {
	v := make([]InventoryInventoryConsumeDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryConsumeDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryConsumeDistributionStatisticsesFromDict(data []InventoryInventoryConsumeDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryConsumeDistributionSegment struct {
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
}

func NewInventoryInventoryConsumeDistributionSegmentFromJson(data string) InventoryInventoryConsumeDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryConsumeDistributionSegmentFromDict(dict)
}

func NewInventoryInventoryConsumeDistributionSegmentFromDict(data map[string]interface{}) InventoryInventoryConsumeDistributionSegment {
	return InventoryInventoryConsumeDistributionSegment{
		ItemName: core.CastString(data["itemName"]),
		Count:    core.CastInt64(data["count"]),
	}
}

func (p InventoryInventoryConsumeDistributionSegment) ToDict() map[string]interface{} {

	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"itemName": itemName,
		"count":    count,
	}
}

func (p InventoryInventoryConsumeDistributionSegment) Pointer() *InventoryInventoryConsumeDistributionSegment {
	return &p
}

func CastInventoryInventoryConsumeDistributionSegments(data []interface{}) []InventoryInventoryConsumeDistributionSegment {
	v := make([]InventoryInventoryConsumeDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryConsumeDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryConsumeDistributionSegmentsFromDict(data []InventoryInventoryConsumeDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryConsumeDistribution struct {
	Statistics   *InventoryInventoryConsumeDistributionStatistics `json:"statistics"`
	Distribution []InventoryInventoryConsumeDistributionSegment   `json:"distribution"`
}

func NewInventoryInventoryConsumeDistributionFromJson(data string) InventoryInventoryConsumeDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryConsumeDistributionFromDict(dict)
}

func NewInventoryInventoryConsumeDistributionFromDict(data map[string]interface{}) InventoryInventoryConsumeDistribution {
	return InventoryInventoryConsumeDistribution{
		Statistics:   NewInventoryInventoryConsumeDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryInventoryConsumeDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryInventoryConsumeDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryInventoryConsumeDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryInventoryConsumeDistribution) Pointer() *InventoryInventoryConsumeDistribution {
	return &p
}

func CastInventoryInventoryConsumeDistributions(data []interface{}) []InventoryInventoryConsumeDistribution {
	v := make([]InventoryInventoryConsumeDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryConsumeDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryConsumeDistributionsFromDict(data []InventoryInventoryConsumeDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryConsumeAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryInventoryConsumeAmountDistributionStatisticsFromJson(data string) InventoryInventoryConsumeAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryConsumeAmountDistributionStatisticsFromDict(dict)
}

func NewInventoryInventoryConsumeAmountDistributionStatisticsFromDict(data map[string]interface{}) InventoryInventoryConsumeAmountDistributionStatistics {
	return InventoryInventoryConsumeAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryInventoryConsumeAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryInventoryConsumeAmountDistributionStatistics) Pointer() *InventoryInventoryConsumeAmountDistributionStatistics {
	return &p
}

func CastInventoryInventoryConsumeAmountDistributionStatisticses(data []interface{}) []InventoryInventoryConsumeAmountDistributionStatistics {
	v := make([]InventoryInventoryConsumeAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryConsumeAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryConsumeAmountDistributionStatisticsesFromDict(data []InventoryInventoryConsumeAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryConsumeAmountDistributionSegment struct {
	ItemName *string `json:"itemName"`
	Sum      *int64  `json:"sum"`
}

func NewInventoryInventoryConsumeAmountDistributionSegmentFromJson(data string) InventoryInventoryConsumeAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryConsumeAmountDistributionSegmentFromDict(dict)
}

func NewInventoryInventoryConsumeAmountDistributionSegmentFromDict(data map[string]interface{}) InventoryInventoryConsumeAmountDistributionSegment {
	return InventoryInventoryConsumeAmountDistributionSegment{
		ItemName: core.CastString(data["itemName"]),
		Sum:      core.CastInt64(data["sum"]),
	}
}

func (p InventoryInventoryConsumeAmountDistributionSegment) ToDict() map[string]interface{} {

	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"itemName": itemName,
		"sum":      sum,
	}
}

func (p InventoryInventoryConsumeAmountDistributionSegment) Pointer() *InventoryInventoryConsumeAmountDistributionSegment {
	return &p
}

func CastInventoryInventoryConsumeAmountDistributionSegments(data []interface{}) []InventoryInventoryConsumeAmountDistributionSegment {
	v := make([]InventoryInventoryConsumeAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryConsumeAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryConsumeAmountDistributionSegmentsFromDict(data []InventoryInventoryConsumeAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryConsumeAmountDistribution struct {
	Statistics   *InventoryInventoryConsumeAmountDistributionStatistics `json:"statistics"`
	Distribution []InventoryInventoryConsumeAmountDistributionSegment   `json:"distribution"`
}

func NewInventoryInventoryConsumeAmountDistributionFromJson(data string) InventoryInventoryConsumeAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryConsumeAmountDistributionFromDict(dict)
}

func NewInventoryInventoryConsumeAmountDistributionFromDict(data map[string]interface{}) InventoryInventoryConsumeAmountDistribution {
	return InventoryInventoryConsumeAmountDistribution{
		Statistics:   NewInventoryInventoryConsumeAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryInventoryConsumeAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryInventoryConsumeAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryInventoryConsumeAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryInventoryConsumeAmountDistribution) Pointer() *InventoryInventoryConsumeAmountDistribution {
	return &p
}

func CastInventoryInventoryConsumeAmountDistributions(data []interface{}) []InventoryInventoryConsumeAmountDistribution {
	v := make([]InventoryInventoryConsumeAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryConsumeAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryConsumeAmountDistributionsFromDict(data []InventoryInventoryConsumeAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventoryDistributions struct {
	Capacity      *InventoryInventoryCapacityDistribution      `json:"capacity"`
	Acquire       *InventoryInventoryAcquireDistribution       `json:"acquire"`
	AcquireAmount *InventoryInventoryAcquireAmountDistribution `json:"acquireAmount"`
	Consume       *InventoryInventoryConsumeDistribution       `json:"consume"`
	ConsumeAmount *InventoryInventoryConsumeAmountDistribution `json:"consumeAmount"`
}

func NewInventoryInventoryDistributionsFromJson(data string) InventoryInventoryDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryDistributionsFromDict(dict)
}

func NewInventoryInventoryDistributionsFromDict(data map[string]interface{}) InventoryInventoryDistributions {
	return InventoryInventoryDistributions{
		Capacity:      NewInventoryInventoryCapacityDistributionFromDict(core.CastMap(data["capacity"])).Pointer(),
		Acquire:       NewInventoryInventoryAcquireDistributionFromDict(core.CastMap(data["acquire"])).Pointer(),
		AcquireAmount: NewInventoryInventoryAcquireAmountDistributionFromDict(core.CastMap(data["acquireAmount"])).Pointer(),
		Consume:       NewInventoryInventoryConsumeDistributionFromDict(core.CastMap(data["consume"])).Pointer(),
		ConsumeAmount: NewInventoryInventoryConsumeAmountDistributionFromDict(core.CastMap(data["consumeAmount"])).Pointer(),
	}
}

func (p InventoryInventoryDistributions) ToDict() map[string]interface{} {

	var capacity map[string]interface{}
	if p.Capacity != nil {
		capacity = p.Capacity.ToDict()
	}
	var acquire map[string]interface{}
	if p.Acquire != nil {
		acquire = p.Acquire.ToDict()
	}
	var acquireAmount map[string]interface{}
	if p.AcquireAmount != nil {
		acquireAmount = p.AcquireAmount.ToDict()
	}
	var consume map[string]interface{}
	if p.Consume != nil {
		consume = p.Consume.ToDict()
	}
	var consumeAmount map[string]interface{}
	if p.ConsumeAmount != nil {
		consumeAmount = p.ConsumeAmount.ToDict()
	}
	return map[string]interface{}{
		"capacity":      capacity,
		"acquire":       acquire,
		"acquireAmount": acquireAmount,
		"consume":       consume,
		"consumeAmount": consumeAmount,
	}
}

func (p InventoryInventoryDistributions) Pointer() *InventoryInventoryDistributions {
	return &p
}

func CastInventoryInventoryDistributionses(data []interface{}) []InventoryInventoryDistributions {
	v := make([]InventoryInventoryDistributions, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoryDistributionsesFromDict(data []InventoryInventoryDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryInventory struct {
	InventoryId   *string                          `json:"inventoryId"`
	InventoryName *string                          `json:"inventoryName"`
	Statistics    *InventoryInventoryStatistics    `json:"statistics"`
	Distributions *InventoryInventoryDistributions `json:"distributions"`
	ItemSets      []InventoryItemSet               `json:"itemSets"`
}

func NewInventoryInventoryFromJson(data string) InventoryInventory {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryInventoryFromDict(dict)
}

func NewInventoryInventoryFromDict(data map[string]interface{}) InventoryInventory {
	return InventoryInventory{
		InventoryId:   core.CastString(data["inventoryId"]),
		InventoryName: core.CastString(data["inventoryName"]),
		Statistics:    NewInventoryInventoryStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewInventoryInventoryDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		ItemSets:      CastInventoryItemSets(core.CastArray(data["itemSets"])),
	}
}

func (p InventoryInventory) ToDict() map[string]interface{} {

	var inventoryId *string
	if p.InventoryId != nil {
		inventoryId = p.InventoryId
	}
	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var itemSets []interface{}
	if p.ItemSets != nil {
		itemSets = CastInventoryItemSetsFromDict(
			p.ItemSets,
		)
	}
	return map[string]interface{}{
		"inventoryId":   inventoryId,
		"inventoryName": inventoryName,
		"statistics":    statistics,
		"distributions": distributions,
		"itemSets":      itemSets,
	}
}

func (p InventoryInventory) Pointer() *InventoryInventory {
	return &p
}

func CastInventoryInventories(data []interface{}) []InventoryInventory {
	v := make([]InventoryInventory, 0)
	for _, d := range data {
		v = append(v, NewInventoryInventoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryInventoriesFromDict(data []InventoryInventory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceStatistics struct {
	Acquire          *int64 `json:"acquire"`
	Consume          *int64 `json:"consume"`
	IncreaseCapacity *int64 `json:"increaseCapacity"`
}

func NewInventoryNamespaceStatisticsFromJson(data string) InventoryNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceStatisticsFromDict(dict)
}

func NewInventoryNamespaceStatisticsFromDict(data map[string]interface{}) InventoryNamespaceStatistics {
	return InventoryNamespaceStatistics{
		Acquire:          core.CastInt64(data["acquire"]),
		Consume:          core.CastInt64(data["consume"]),
		IncreaseCapacity: core.CastInt64(data["increaseCapacity"]),
	}
}

func (p InventoryNamespaceStatistics) ToDict() map[string]interface{} {

	var acquire *int64
	if p.Acquire != nil {
		acquire = p.Acquire
	}
	var consume *int64
	if p.Consume != nil {
		consume = p.Consume
	}
	var increaseCapacity *int64
	if p.IncreaseCapacity != nil {
		increaseCapacity = p.IncreaseCapacity
	}
	return map[string]interface{}{
		"acquire":          acquire,
		"consume":          consume,
		"increaseCapacity": increaseCapacity,
	}
}

func (p InventoryNamespaceStatistics) Pointer() *InventoryNamespaceStatistics {
	return &p
}

func CastInventoryNamespaceStatisticses(data []interface{}) []InventoryNamespaceStatistics {
	v := make([]InventoryNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceStatisticsesFromDict(data []InventoryNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceAcquireDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryNamespaceAcquireDistributionStatisticsFromJson(data string) InventoryNamespaceAcquireDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceAcquireDistributionStatisticsFromDict(dict)
}

func NewInventoryNamespaceAcquireDistributionStatisticsFromDict(data map[string]interface{}) InventoryNamespaceAcquireDistributionStatistics {
	return InventoryNamespaceAcquireDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryNamespaceAcquireDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryNamespaceAcquireDistributionStatistics) Pointer() *InventoryNamespaceAcquireDistributionStatistics {
	return &p
}

func CastInventoryNamespaceAcquireDistributionStatisticses(data []interface{}) []InventoryNamespaceAcquireDistributionStatistics {
	v := make([]InventoryNamespaceAcquireDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceAcquireDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceAcquireDistributionStatisticsesFromDict(data []InventoryNamespaceAcquireDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceAcquireDistributionSegment struct {
	InventoryName *string `json:"inventoryName"`
	Count         *int64  `json:"count"`
}

func NewInventoryNamespaceAcquireDistributionSegmentFromJson(data string) InventoryNamespaceAcquireDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceAcquireDistributionSegmentFromDict(dict)
}

func NewInventoryNamespaceAcquireDistributionSegmentFromDict(data map[string]interface{}) InventoryNamespaceAcquireDistributionSegment {
	return InventoryNamespaceAcquireDistributionSegment{
		InventoryName: core.CastString(data["inventoryName"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p InventoryNamespaceAcquireDistributionSegment) ToDict() map[string]interface{} {

	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"inventoryName": inventoryName,
		"count":         count,
	}
}

func (p InventoryNamespaceAcquireDistributionSegment) Pointer() *InventoryNamespaceAcquireDistributionSegment {
	return &p
}

func CastInventoryNamespaceAcquireDistributionSegments(data []interface{}) []InventoryNamespaceAcquireDistributionSegment {
	v := make([]InventoryNamespaceAcquireDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceAcquireDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceAcquireDistributionSegmentsFromDict(data []InventoryNamespaceAcquireDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceAcquireDistribution struct {
	Statistics   *InventoryNamespaceAcquireDistributionStatistics `json:"statistics"`
	Distribution []InventoryNamespaceAcquireDistributionSegment   `json:"distribution"`
}

func NewInventoryNamespaceAcquireDistributionFromJson(data string) InventoryNamespaceAcquireDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceAcquireDistributionFromDict(dict)
}

func NewInventoryNamespaceAcquireDistributionFromDict(data map[string]interface{}) InventoryNamespaceAcquireDistribution {
	return InventoryNamespaceAcquireDistribution{
		Statistics:   NewInventoryNamespaceAcquireDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryNamespaceAcquireDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryNamespaceAcquireDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryNamespaceAcquireDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryNamespaceAcquireDistribution) Pointer() *InventoryNamespaceAcquireDistribution {
	return &p
}

func CastInventoryNamespaceAcquireDistributions(data []interface{}) []InventoryNamespaceAcquireDistribution {
	v := make([]InventoryNamespaceAcquireDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceAcquireDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceAcquireDistributionsFromDict(data []InventoryNamespaceAcquireDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceAcquireAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryNamespaceAcquireAmountDistributionStatisticsFromJson(data string) InventoryNamespaceAcquireAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceAcquireAmountDistributionStatisticsFromDict(dict)
}

func NewInventoryNamespaceAcquireAmountDistributionStatisticsFromDict(data map[string]interface{}) InventoryNamespaceAcquireAmountDistributionStatistics {
	return InventoryNamespaceAcquireAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryNamespaceAcquireAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryNamespaceAcquireAmountDistributionStatistics) Pointer() *InventoryNamespaceAcquireAmountDistributionStatistics {
	return &p
}

func CastInventoryNamespaceAcquireAmountDistributionStatisticses(data []interface{}) []InventoryNamespaceAcquireAmountDistributionStatistics {
	v := make([]InventoryNamespaceAcquireAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceAcquireAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceAcquireAmountDistributionStatisticsesFromDict(data []InventoryNamespaceAcquireAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceAcquireAmountDistributionSegment struct {
	InventoryName *string `json:"inventoryName"`
	Sum           *int64  `json:"sum"`
}

func NewInventoryNamespaceAcquireAmountDistributionSegmentFromJson(data string) InventoryNamespaceAcquireAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceAcquireAmountDistributionSegmentFromDict(dict)
}

func NewInventoryNamespaceAcquireAmountDistributionSegmentFromDict(data map[string]interface{}) InventoryNamespaceAcquireAmountDistributionSegment {
	return InventoryNamespaceAcquireAmountDistributionSegment{
		InventoryName: core.CastString(data["inventoryName"]),
		Sum:           core.CastInt64(data["sum"]),
	}
}

func (p InventoryNamespaceAcquireAmountDistributionSegment) ToDict() map[string]interface{} {

	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"inventoryName": inventoryName,
		"sum":           sum,
	}
}

func (p InventoryNamespaceAcquireAmountDistributionSegment) Pointer() *InventoryNamespaceAcquireAmountDistributionSegment {
	return &p
}

func CastInventoryNamespaceAcquireAmountDistributionSegments(data []interface{}) []InventoryNamespaceAcquireAmountDistributionSegment {
	v := make([]InventoryNamespaceAcquireAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceAcquireAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceAcquireAmountDistributionSegmentsFromDict(data []InventoryNamespaceAcquireAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceAcquireAmountDistribution struct {
	Statistics   *InventoryNamespaceAcquireAmountDistributionStatistics `json:"statistics"`
	Distribution []InventoryNamespaceAcquireAmountDistributionSegment   `json:"distribution"`
}

func NewInventoryNamespaceAcquireAmountDistributionFromJson(data string) InventoryNamespaceAcquireAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceAcquireAmountDistributionFromDict(dict)
}

func NewInventoryNamespaceAcquireAmountDistributionFromDict(data map[string]interface{}) InventoryNamespaceAcquireAmountDistribution {
	return InventoryNamespaceAcquireAmountDistribution{
		Statistics:   NewInventoryNamespaceAcquireAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryNamespaceAcquireAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryNamespaceAcquireAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryNamespaceAcquireAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryNamespaceAcquireAmountDistribution) Pointer() *InventoryNamespaceAcquireAmountDistribution {
	return &p
}

func CastInventoryNamespaceAcquireAmountDistributions(data []interface{}) []InventoryNamespaceAcquireAmountDistribution {
	v := make([]InventoryNamespaceAcquireAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceAcquireAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceAcquireAmountDistributionsFromDict(data []InventoryNamespaceAcquireAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceConsumeDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryNamespaceConsumeDistributionStatisticsFromJson(data string) InventoryNamespaceConsumeDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceConsumeDistributionStatisticsFromDict(dict)
}

func NewInventoryNamespaceConsumeDistributionStatisticsFromDict(data map[string]interface{}) InventoryNamespaceConsumeDistributionStatistics {
	return InventoryNamespaceConsumeDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryNamespaceConsumeDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryNamespaceConsumeDistributionStatistics) Pointer() *InventoryNamespaceConsumeDistributionStatistics {
	return &p
}

func CastInventoryNamespaceConsumeDistributionStatisticses(data []interface{}) []InventoryNamespaceConsumeDistributionStatistics {
	v := make([]InventoryNamespaceConsumeDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceConsumeDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceConsumeDistributionStatisticsesFromDict(data []InventoryNamespaceConsumeDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceConsumeDistributionSegment struct {
	InventoryName *string `json:"inventoryName"`
	Count         *int64  `json:"count"`
}

func NewInventoryNamespaceConsumeDistributionSegmentFromJson(data string) InventoryNamespaceConsumeDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceConsumeDistributionSegmentFromDict(dict)
}

func NewInventoryNamespaceConsumeDistributionSegmentFromDict(data map[string]interface{}) InventoryNamespaceConsumeDistributionSegment {
	return InventoryNamespaceConsumeDistributionSegment{
		InventoryName: core.CastString(data["inventoryName"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p InventoryNamespaceConsumeDistributionSegment) ToDict() map[string]interface{} {

	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"inventoryName": inventoryName,
		"count":         count,
	}
}

func (p InventoryNamespaceConsumeDistributionSegment) Pointer() *InventoryNamespaceConsumeDistributionSegment {
	return &p
}

func CastInventoryNamespaceConsumeDistributionSegments(data []interface{}) []InventoryNamespaceConsumeDistributionSegment {
	v := make([]InventoryNamespaceConsumeDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceConsumeDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceConsumeDistributionSegmentsFromDict(data []InventoryNamespaceConsumeDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceConsumeDistribution struct {
	Statistics   *InventoryNamespaceConsumeDistributionStatistics `json:"statistics"`
	Distribution []InventoryNamespaceConsumeDistributionSegment   `json:"distribution"`
}

func NewInventoryNamespaceConsumeDistributionFromJson(data string) InventoryNamespaceConsumeDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceConsumeDistributionFromDict(dict)
}

func NewInventoryNamespaceConsumeDistributionFromDict(data map[string]interface{}) InventoryNamespaceConsumeDistribution {
	return InventoryNamespaceConsumeDistribution{
		Statistics:   NewInventoryNamespaceConsumeDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryNamespaceConsumeDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryNamespaceConsumeDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryNamespaceConsumeDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryNamespaceConsumeDistribution) Pointer() *InventoryNamespaceConsumeDistribution {
	return &p
}

func CastInventoryNamespaceConsumeDistributions(data []interface{}) []InventoryNamespaceConsumeDistribution {
	v := make([]InventoryNamespaceConsumeDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceConsumeDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceConsumeDistributionsFromDict(data []InventoryNamespaceConsumeDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceConsumeAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryNamespaceConsumeAmountDistributionStatisticsFromJson(data string) InventoryNamespaceConsumeAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceConsumeAmountDistributionStatisticsFromDict(dict)
}

func NewInventoryNamespaceConsumeAmountDistributionStatisticsFromDict(data map[string]interface{}) InventoryNamespaceConsumeAmountDistributionStatistics {
	return InventoryNamespaceConsumeAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryNamespaceConsumeAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryNamespaceConsumeAmountDistributionStatistics) Pointer() *InventoryNamespaceConsumeAmountDistributionStatistics {
	return &p
}

func CastInventoryNamespaceConsumeAmountDistributionStatisticses(data []interface{}) []InventoryNamespaceConsumeAmountDistributionStatistics {
	v := make([]InventoryNamespaceConsumeAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceConsumeAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceConsumeAmountDistributionStatisticsesFromDict(data []InventoryNamespaceConsumeAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceConsumeAmountDistributionSegment struct {
	InventoryName *string `json:"inventoryName"`
	Sum           *int64  `json:"sum"`
}

func NewInventoryNamespaceConsumeAmountDistributionSegmentFromJson(data string) InventoryNamespaceConsumeAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceConsumeAmountDistributionSegmentFromDict(dict)
}

func NewInventoryNamespaceConsumeAmountDistributionSegmentFromDict(data map[string]interface{}) InventoryNamespaceConsumeAmountDistributionSegment {
	return InventoryNamespaceConsumeAmountDistributionSegment{
		InventoryName: core.CastString(data["inventoryName"]),
		Sum:           core.CastInt64(data["sum"]),
	}
}

func (p InventoryNamespaceConsumeAmountDistributionSegment) ToDict() map[string]interface{} {

	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"inventoryName": inventoryName,
		"sum":           sum,
	}
}

func (p InventoryNamespaceConsumeAmountDistributionSegment) Pointer() *InventoryNamespaceConsumeAmountDistributionSegment {
	return &p
}

func CastInventoryNamespaceConsumeAmountDistributionSegments(data []interface{}) []InventoryNamespaceConsumeAmountDistributionSegment {
	v := make([]InventoryNamespaceConsumeAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceConsumeAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceConsumeAmountDistributionSegmentsFromDict(data []InventoryNamespaceConsumeAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceConsumeAmountDistribution struct {
	Statistics   *InventoryNamespaceConsumeAmountDistributionStatistics `json:"statistics"`
	Distribution []InventoryNamespaceConsumeAmountDistributionSegment   `json:"distribution"`
}

func NewInventoryNamespaceConsumeAmountDistributionFromJson(data string) InventoryNamespaceConsumeAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceConsumeAmountDistributionFromDict(dict)
}

func NewInventoryNamespaceConsumeAmountDistributionFromDict(data map[string]interface{}) InventoryNamespaceConsumeAmountDistribution {
	return InventoryNamespaceConsumeAmountDistribution{
		Statistics:   NewInventoryNamespaceConsumeAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryNamespaceConsumeAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryNamespaceConsumeAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryNamespaceConsumeAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryNamespaceConsumeAmountDistribution) Pointer() *InventoryNamespaceConsumeAmountDistribution {
	return &p
}

func CastInventoryNamespaceConsumeAmountDistributions(data []interface{}) []InventoryNamespaceConsumeAmountDistribution {
	v := make([]InventoryNamespaceConsumeAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceConsumeAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceConsumeAmountDistributionsFromDict(data []InventoryNamespaceConsumeAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceIncreaseCapacityDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryNamespaceIncreaseCapacityDistributionStatisticsFromJson(data string) InventoryNamespaceIncreaseCapacityDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceIncreaseCapacityDistributionStatisticsFromDict(dict)
}

func NewInventoryNamespaceIncreaseCapacityDistributionStatisticsFromDict(data map[string]interface{}) InventoryNamespaceIncreaseCapacityDistributionStatistics {
	return InventoryNamespaceIncreaseCapacityDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryNamespaceIncreaseCapacityDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryNamespaceIncreaseCapacityDistributionStatistics) Pointer() *InventoryNamespaceIncreaseCapacityDistributionStatistics {
	return &p
}

func CastInventoryNamespaceIncreaseCapacityDistributionStatisticses(data []interface{}) []InventoryNamespaceIncreaseCapacityDistributionStatistics {
	v := make([]InventoryNamespaceIncreaseCapacityDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceIncreaseCapacityDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceIncreaseCapacityDistributionStatisticsesFromDict(data []InventoryNamespaceIncreaseCapacityDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceIncreaseCapacityDistributionSegment struct {
	InventoryName *string `json:"inventoryName"`
	Count         *int64  `json:"count"`
}

func NewInventoryNamespaceIncreaseCapacityDistributionSegmentFromJson(data string) InventoryNamespaceIncreaseCapacityDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceIncreaseCapacityDistributionSegmentFromDict(dict)
}

func NewInventoryNamespaceIncreaseCapacityDistributionSegmentFromDict(data map[string]interface{}) InventoryNamespaceIncreaseCapacityDistributionSegment {
	return InventoryNamespaceIncreaseCapacityDistributionSegment{
		InventoryName: core.CastString(data["inventoryName"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p InventoryNamespaceIncreaseCapacityDistributionSegment) ToDict() map[string]interface{} {

	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"inventoryName": inventoryName,
		"count":         count,
	}
}

func (p InventoryNamespaceIncreaseCapacityDistributionSegment) Pointer() *InventoryNamespaceIncreaseCapacityDistributionSegment {
	return &p
}

func CastInventoryNamespaceIncreaseCapacityDistributionSegments(data []interface{}) []InventoryNamespaceIncreaseCapacityDistributionSegment {
	v := make([]InventoryNamespaceIncreaseCapacityDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceIncreaseCapacityDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceIncreaseCapacityDistributionSegmentsFromDict(data []InventoryNamespaceIncreaseCapacityDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceIncreaseCapacityDistribution struct {
	Statistics   *InventoryNamespaceIncreaseCapacityDistributionStatistics `json:"statistics"`
	Distribution []InventoryNamespaceIncreaseCapacityDistributionSegment   `json:"distribution"`
}

func NewInventoryNamespaceIncreaseCapacityDistributionFromJson(data string) InventoryNamespaceIncreaseCapacityDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceIncreaseCapacityDistributionFromDict(dict)
}

func NewInventoryNamespaceIncreaseCapacityDistributionFromDict(data map[string]interface{}) InventoryNamespaceIncreaseCapacityDistribution {
	return InventoryNamespaceIncreaseCapacityDistribution{
		Statistics:   NewInventoryNamespaceIncreaseCapacityDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryNamespaceIncreaseCapacityDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryNamespaceIncreaseCapacityDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryNamespaceIncreaseCapacityDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryNamespaceIncreaseCapacityDistribution) Pointer() *InventoryNamespaceIncreaseCapacityDistribution {
	return &p
}

func CastInventoryNamespaceIncreaseCapacityDistributions(data []interface{}) []InventoryNamespaceIncreaseCapacityDistribution {
	v := make([]InventoryNamespaceIncreaseCapacityDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceIncreaseCapacityDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceIncreaseCapacityDistributionsFromDict(data []InventoryNamespaceIncreaseCapacityDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceIncreaseCapacityAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewInventoryNamespaceIncreaseCapacityAmountDistributionStatisticsFromJson(data string) InventoryNamespaceIncreaseCapacityAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceIncreaseCapacityAmountDistributionStatisticsFromDict(dict)
}

func NewInventoryNamespaceIncreaseCapacityAmountDistributionStatisticsFromDict(data map[string]interface{}) InventoryNamespaceIncreaseCapacityAmountDistributionStatistics {
	return InventoryNamespaceIncreaseCapacityAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p InventoryNamespaceIncreaseCapacityAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p InventoryNamespaceIncreaseCapacityAmountDistributionStatistics) Pointer() *InventoryNamespaceIncreaseCapacityAmountDistributionStatistics {
	return &p
}

func CastInventoryNamespaceIncreaseCapacityAmountDistributionStatisticses(data []interface{}) []InventoryNamespaceIncreaseCapacityAmountDistributionStatistics {
	v := make([]InventoryNamespaceIncreaseCapacityAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceIncreaseCapacityAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceIncreaseCapacityAmountDistributionStatisticsesFromDict(data []InventoryNamespaceIncreaseCapacityAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceIncreaseCapacityAmountDistributionSegment struct {
	InventoryName *string `json:"inventoryName"`
	Sum           *int64  `json:"sum"`
}

func NewInventoryNamespaceIncreaseCapacityAmountDistributionSegmentFromJson(data string) InventoryNamespaceIncreaseCapacityAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceIncreaseCapacityAmountDistributionSegmentFromDict(dict)
}

func NewInventoryNamespaceIncreaseCapacityAmountDistributionSegmentFromDict(data map[string]interface{}) InventoryNamespaceIncreaseCapacityAmountDistributionSegment {
	return InventoryNamespaceIncreaseCapacityAmountDistributionSegment{
		InventoryName: core.CastString(data["inventoryName"]),
		Sum:           core.CastInt64(data["sum"]),
	}
}

func (p InventoryNamespaceIncreaseCapacityAmountDistributionSegment) ToDict() map[string]interface{} {

	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"inventoryName": inventoryName,
		"sum":           sum,
	}
}

func (p InventoryNamespaceIncreaseCapacityAmountDistributionSegment) Pointer() *InventoryNamespaceIncreaseCapacityAmountDistributionSegment {
	return &p
}

func CastInventoryNamespaceIncreaseCapacityAmountDistributionSegments(data []interface{}) []InventoryNamespaceIncreaseCapacityAmountDistributionSegment {
	v := make([]InventoryNamespaceIncreaseCapacityAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceIncreaseCapacityAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceIncreaseCapacityAmountDistributionSegmentsFromDict(data []InventoryNamespaceIncreaseCapacityAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceIncreaseCapacityAmountDistribution struct {
	Statistics   *InventoryNamespaceIncreaseCapacityAmountDistributionStatistics `json:"statistics"`
	Distribution []InventoryNamespaceIncreaseCapacityAmountDistributionSegment   `json:"distribution"`
}

func NewInventoryNamespaceIncreaseCapacityAmountDistributionFromJson(data string) InventoryNamespaceIncreaseCapacityAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceIncreaseCapacityAmountDistributionFromDict(dict)
}

func NewInventoryNamespaceIncreaseCapacityAmountDistributionFromDict(data map[string]interface{}) InventoryNamespaceIncreaseCapacityAmountDistribution {
	return InventoryNamespaceIncreaseCapacityAmountDistribution{
		Statistics:   NewInventoryNamespaceIncreaseCapacityAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastInventoryNamespaceIncreaseCapacityAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p InventoryNamespaceIncreaseCapacityAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastInventoryNamespaceIncreaseCapacityAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p InventoryNamespaceIncreaseCapacityAmountDistribution) Pointer() *InventoryNamespaceIncreaseCapacityAmountDistribution {
	return &p
}

func CastInventoryNamespaceIncreaseCapacityAmountDistributions(data []interface{}) []InventoryNamespaceIncreaseCapacityAmountDistribution {
	v := make([]InventoryNamespaceIncreaseCapacityAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceIncreaseCapacityAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceIncreaseCapacityAmountDistributionsFromDict(data []InventoryNamespaceIncreaseCapacityAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespaceDistributions struct {
	Acquire                *InventoryNamespaceAcquireDistribution                `json:"acquire"`
	AcquireAmount          *InventoryNamespaceAcquireAmountDistribution          `json:"acquireAmount"`
	Consume                *InventoryNamespaceConsumeDistribution                `json:"consume"`
	ConsumeAmount          *InventoryNamespaceConsumeAmountDistribution          `json:"consumeAmount"`
	IncreaseCapacity       *InventoryNamespaceIncreaseCapacityDistribution       `json:"increaseCapacity"`
	IncreaseCapacityAmount *InventoryNamespaceIncreaseCapacityAmountDistribution `json:"increaseCapacityAmount"`
}

func NewInventoryNamespaceDistributionsFromJson(data string) InventoryNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceDistributionsFromDict(dict)
}

func NewInventoryNamespaceDistributionsFromDict(data map[string]interface{}) InventoryNamespaceDistributions {
	return InventoryNamespaceDistributions{
		Acquire:                NewInventoryNamespaceAcquireDistributionFromDict(core.CastMap(data["acquire"])).Pointer(),
		AcquireAmount:          NewInventoryNamespaceAcquireAmountDistributionFromDict(core.CastMap(data["acquireAmount"])).Pointer(),
		Consume:                NewInventoryNamespaceConsumeDistributionFromDict(core.CastMap(data["consume"])).Pointer(),
		ConsumeAmount:          NewInventoryNamespaceConsumeAmountDistributionFromDict(core.CastMap(data["consumeAmount"])).Pointer(),
		IncreaseCapacity:       NewInventoryNamespaceIncreaseCapacityDistributionFromDict(core.CastMap(data["increaseCapacity"])).Pointer(),
		IncreaseCapacityAmount: NewInventoryNamespaceIncreaseCapacityAmountDistributionFromDict(core.CastMap(data["increaseCapacityAmount"])).Pointer(),
	}
}

func (p InventoryNamespaceDistributions) ToDict() map[string]interface{} {

	var acquire map[string]interface{}
	if p.Acquire != nil {
		acquire = p.Acquire.ToDict()
	}
	var acquireAmount map[string]interface{}
	if p.AcquireAmount != nil {
		acquireAmount = p.AcquireAmount.ToDict()
	}
	var consume map[string]interface{}
	if p.Consume != nil {
		consume = p.Consume.ToDict()
	}
	var consumeAmount map[string]interface{}
	if p.ConsumeAmount != nil {
		consumeAmount = p.ConsumeAmount.ToDict()
	}
	var increaseCapacity map[string]interface{}
	if p.IncreaseCapacity != nil {
		increaseCapacity = p.IncreaseCapacity.ToDict()
	}
	var increaseCapacityAmount map[string]interface{}
	if p.IncreaseCapacityAmount != nil {
		increaseCapacityAmount = p.IncreaseCapacityAmount.ToDict()
	}
	return map[string]interface{}{
		"acquire":                acquire,
		"acquireAmount":          acquireAmount,
		"consume":                consume,
		"consumeAmount":          consumeAmount,
		"increaseCapacity":       increaseCapacity,
		"increaseCapacityAmount": increaseCapacityAmount,
	}
}

func (p InventoryNamespaceDistributions) Pointer() *InventoryNamespaceDistributions {
	return &p
}

func CastInventoryNamespaceDistributionses(data []interface{}) []InventoryNamespaceDistributions {
	v := make([]InventoryNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespaceDistributionsesFromDict(data []InventoryNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryNamespace struct {
	NamespaceId   *string                          `json:"namespaceId"`
	Year          *int32                           `json:"year"`
	Month         *int32                           `json:"month"`
	Day           *int32                           `json:"day"`
	NamespaceName *string                          `json:"namespaceName"`
	Statistics    *InventoryNamespaceStatistics    `json:"statistics"`
	Distributions *InventoryNamespaceDistributions `json:"distributions"`
	Inventories   []InventoryInventory             `json:"inventories"`
}

func NewInventoryNamespaceFromJson(data string) InventoryNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryNamespaceFromDict(dict)
}

func NewInventoryNamespaceFromDict(data map[string]interface{}) InventoryNamespace {
	return InventoryNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewInventoryNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewInventoryNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		Inventories:   CastInventoryInventories(core.CastArray(data["inventories"])),
	}
}

func (p InventoryNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var inventories []interface{}
	if p.Inventories != nil {
		inventories = CastInventoryInventoriesFromDict(
			p.Inventories,
		)
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
		"inventories":   inventories,
	}
}

func (p InventoryNamespace) Pointer() *InventoryNamespace {
	return &p
}

func CastInventoryNamespaces(data []interface{}) []InventoryNamespace {
	v := make([]InventoryNamespace, 0)
	for _, d := range data {
		v = append(v, NewInventoryNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryNamespacesFromDict(data []InventoryNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyNamespaceStatistics struct {
	Encrypt *int64 `json:"encrypt"`
	Decrypt *int64 `json:"decrypt"`
}

func NewKeyNamespaceStatisticsFromJson(data string) KeyNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyNamespaceStatisticsFromDict(dict)
}

func NewKeyNamespaceStatisticsFromDict(data map[string]interface{}) KeyNamespaceStatistics {
	return KeyNamespaceStatistics{
		Encrypt: core.CastInt64(data["encrypt"]),
		Decrypt: core.CastInt64(data["decrypt"]),
	}
}

func (p KeyNamespaceStatistics) ToDict() map[string]interface{} {

	var encrypt *int64
	if p.Encrypt != nil {
		encrypt = p.Encrypt
	}
	var decrypt *int64
	if p.Decrypt != nil {
		decrypt = p.Decrypt
	}
	return map[string]interface{}{
		"encrypt": encrypt,
		"decrypt": decrypt,
	}
}

func (p KeyNamespaceStatistics) Pointer() *KeyNamespaceStatistics {
	return &p
}

func CastKeyNamespaceStatisticses(data []interface{}) []KeyNamespaceStatistics {
	v := make([]KeyNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewKeyNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyNamespaceStatisticsesFromDict(data []KeyNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyNamespaceEncryptDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewKeyNamespaceEncryptDistributionStatisticsFromJson(data string) KeyNamespaceEncryptDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyNamespaceEncryptDistributionStatisticsFromDict(dict)
}

func NewKeyNamespaceEncryptDistributionStatisticsFromDict(data map[string]interface{}) KeyNamespaceEncryptDistributionStatistics {
	return KeyNamespaceEncryptDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p KeyNamespaceEncryptDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p KeyNamespaceEncryptDistributionStatistics) Pointer() *KeyNamespaceEncryptDistributionStatistics {
	return &p
}

func CastKeyNamespaceEncryptDistributionStatisticses(data []interface{}) []KeyNamespaceEncryptDistributionStatistics {
	v := make([]KeyNamespaceEncryptDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewKeyNamespaceEncryptDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyNamespaceEncryptDistributionStatisticsesFromDict(data []KeyNamespaceEncryptDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyNamespaceEncryptDistributionSegment struct {
	KeyName *string `json:"keyName"`
	Count   *int64  `json:"count"`
}

func NewKeyNamespaceEncryptDistributionSegmentFromJson(data string) KeyNamespaceEncryptDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyNamespaceEncryptDistributionSegmentFromDict(dict)
}

func NewKeyNamespaceEncryptDistributionSegmentFromDict(data map[string]interface{}) KeyNamespaceEncryptDistributionSegment {
	return KeyNamespaceEncryptDistributionSegment{
		KeyName: core.CastString(data["keyName"]),
		Count:   core.CastInt64(data["count"]),
	}
}

func (p KeyNamespaceEncryptDistributionSegment) ToDict() map[string]interface{} {

	var keyName *string
	if p.KeyName != nil {
		keyName = p.KeyName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"keyName": keyName,
		"count":   count,
	}
}

func (p KeyNamespaceEncryptDistributionSegment) Pointer() *KeyNamespaceEncryptDistributionSegment {
	return &p
}

func CastKeyNamespaceEncryptDistributionSegments(data []interface{}) []KeyNamespaceEncryptDistributionSegment {
	v := make([]KeyNamespaceEncryptDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewKeyNamespaceEncryptDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyNamespaceEncryptDistributionSegmentsFromDict(data []KeyNamespaceEncryptDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyNamespaceEncryptDistribution struct {
	Statistics   *KeyNamespaceEncryptDistributionStatistics `json:"statistics"`
	Distribution []KeyNamespaceEncryptDistributionSegment   `json:"distribution"`
}

func NewKeyNamespaceEncryptDistributionFromJson(data string) KeyNamespaceEncryptDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyNamespaceEncryptDistributionFromDict(dict)
}

func NewKeyNamespaceEncryptDistributionFromDict(data map[string]interface{}) KeyNamespaceEncryptDistribution {
	return KeyNamespaceEncryptDistribution{
		Statistics:   NewKeyNamespaceEncryptDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastKeyNamespaceEncryptDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p KeyNamespaceEncryptDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastKeyNamespaceEncryptDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p KeyNamespaceEncryptDistribution) Pointer() *KeyNamespaceEncryptDistribution {
	return &p
}

func CastKeyNamespaceEncryptDistributions(data []interface{}) []KeyNamespaceEncryptDistribution {
	v := make([]KeyNamespaceEncryptDistribution, 0)
	for _, d := range data {
		v = append(v, NewKeyNamespaceEncryptDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyNamespaceEncryptDistributionsFromDict(data []KeyNamespaceEncryptDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyNamespaceDecryptDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewKeyNamespaceDecryptDistributionStatisticsFromJson(data string) KeyNamespaceDecryptDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyNamespaceDecryptDistributionStatisticsFromDict(dict)
}

func NewKeyNamespaceDecryptDistributionStatisticsFromDict(data map[string]interface{}) KeyNamespaceDecryptDistributionStatistics {
	return KeyNamespaceDecryptDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p KeyNamespaceDecryptDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p KeyNamespaceDecryptDistributionStatistics) Pointer() *KeyNamespaceDecryptDistributionStatistics {
	return &p
}

func CastKeyNamespaceDecryptDistributionStatisticses(data []interface{}) []KeyNamespaceDecryptDistributionStatistics {
	v := make([]KeyNamespaceDecryptDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewKeyNamespaceDecryptDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyNamespaceDecryptDistributionStatisticsesFromDict(data []KeyNamespaceDecryptDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyNamespaceDecryptDistributionSegment struct {
	KeyName *string `json:"keyName"`
	Count   *int64  `json:"count"`
}

func NewKeyNamespaceDecryptDistributionSegmentFromJson(data string) KeyNamespaceDecryptDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyNamespaceDecryptDistributionSegmentFromDict(dict)
}

func NewKeyNamespaceDecryptDistributionSegmentFromDict(data map[string]interface{}) KeyNamespaceDecryptDistributionSegment {
	return KeyNamespaceDecryptDistributionSegment{
		KeyName: core.CastString(data["keyName"]),
		Count:   core.CastInt64(data["count"]),
	}
}

func (p KeyNamespaceDecryptDistributionSegment) ToDict() map[string]interface{} {

	var keyName *string
	if p.KeyName != nil {
		keyName = p.KeyName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"keyName": keyName,
		"count":   count,
	}
}

func (p KeyNamespaceDecryptDistributionSegment) Pointer() *KeyNamespaceDecryptDistributionSegment {
	return &p
}

func CastKeyNamespaceDecryptDistributionSegments(data []interface{}) []KeyNamespaceDecryptDistributionSegment {
	v := make([]KeyNamespaceDecryptDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewKeyNamespaceDecryptDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyNamespaceDecryptDistributionSegmentsFromDict(data []KeyNamespaceDecryptDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyNamespaceDecryptDistribution struct {
	Statistics   *KeyNamespaceDecryptDistributionStatistics `json:"statistics"`
	Distribution []KeyNamespaceDecryptDistributionSegment   `json:"distribution"`
}

func NewKeyNamespaceDecryptDistributionFromJson(data string) KeyNamespaceDecryptDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyNamespaceDecryptDistributionFromDict(dict)
}

func NewKeyNamespaceDecryptDistributionFromDict(data map[string]interface{}) KeyNamespaceDecryptDistribution {
	return KeyNamespaceDecryptDistribution{
		Statistics:   NewKeyNamespaceDecryptDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastKeyNamespaceDecryptDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p KeyNamespaceDecryptDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastKeyNamespaceDecryptDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p KeyNamespaceDecryptDistribution) Pointer() *KeyNamespaceDecryptDistribution {
	return &p
}

func CastKeyNamespaceDecryptDistributions(data []interface{}) []KeyNamespaceDecryptDistribution {
	v := make([]KeyNamespaceDecryptDistribution, 0)
	for _, d := range data {
		v = append(v, NewKeyNamespaceDecryptDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyNamespaceDecryptDistributionsFromDict(data []KeyNamespaceDecryptDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyNamespaceDistributions struct {
	Encrypt *KeyNamespaceEncryptDistribution `json:"encrypt"`
	Decrypt *KeyNamespaceDecryptDistribution `json:"decrypt"`
}

func NewKeyNamespaceDistributionsFromJson(data string) KeyNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyNamespaceDistributionsFromDict(dict)
}

func NewKeyNamespaceDistributionsFromDict(data map[string]interface{}) KeyNamespaceDistributions {
	return KeyNamespaceDistributions{
		Encrypt: NewKeyNamespaceEncryptDistributionFromDict(core.CastMap(data["encrypt"])).Pointer(),
		Decrypt: NewKeyNamespaceDecryptDistributionFromDict(core.CastMap(data["decrypt"])).Pointer(),
	}
}

func (p KeyNamespaceDistributions) ToDict() map[string]interface{} {

	var encrypt map[string]interface{}
	if p.Encrypt != nil {
		encrypt = p.Encrypt.ToDict()
	}
	var decrypt map[string]interface{}
	if p.Decrypt != nil {
		decrypt = p.Decrypt.ToDict()
	}
	return map[string]interface{}{
		"encrypt": encrypt,
		"decrypt": decrypt,
	}
}

func (p KeyNamespaceDistributions) Pointer() *KeyNamespaceDistributions {
	return &p
}

func CastKeyNamespaceDistributionses(data []interface{}) []KeyNamespaceDistributions {
	v := make([]KeyNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewKeyNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyNamespaceDistributionsesFromDict(data []KeyNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyNamespace struct {
	NamespaceId   *string                    `json:"namespaceId"`
	Year          *int32                     `json:"year"`
	Month         *int32                     `json:"month"`
	Day           *int32                     `json:"day"`
	NamespaceName *string                    `json:"namespaceName"`
	Statistics    *KeyNamespaceStatistics    `json:"statistics"`
	Distributions *KeyNamespaceDistributions `json:"distributions"`
}

func NewKeyNamespaceFromJson(data string) KeyNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyNamespaceFromDict(dict)
}

func NewKeyNamespaceFromDict(data map[string]interface{}) KeyNamespace {
	return KeyNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewKeyNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewKeyNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p KeyNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p KeyNamespace) Pointer() *KeyNamespace {
	return &p
}

func CastKeyNamespaces(data []interface{}) []KeyNamespace {
	v := make([]KeyNamespace, 0)
	for _, d := range data {
		v = append(v, NewKeyNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyNamespacesFromDict(data []KeyNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type KeyKey struct {
	KeyId   *string `json:"keyId"`
	KeyName *string `json:"keyName"`
}

func NewKeyKeyFromJson(data string) KeyKey {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyKeyFromDict(dict)
}

func NewKeyKeyFromDict(data map[string]interface{}) KeyKey {
	return KeyKey{
		KeyId:   core.CastString(data["keyId"]),
		KeyName: core.CastString(data["keyName"]),
	}
}

func (p KeyKey) ToDict() map[string]interface{} {

	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var keyName *string
	if p.KeyName != nil {
		keyName = p.KeyName
	}
	return map[string]interface{}{
		"keyId":   keyId,
		"keyName": keyName,
	}
}

func (p KeyKey) Pointer() *KeyKey {
	return &p
}

func CastKeyKeys(data []interface{}) []KeyKey {
	v := make([]KeyKey, 0)
	for _, d := range data {
		v = append(v, NewKeyKeyFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeyKeysFromDict(data []KeyKey) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitCounterStatistics struct {
	Increase       *int64 `json:"increase"`
	IncreaseAmount *int64 `json:"increaseAmount"`
}

func NewLimitCounterStatisticsFromJson(data string) LimitCounterStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitCounterStatisticsFromDict(dict)
}

func NewLimitCounterStatisticsFromDict(data map[string]interface{}) LimitCounterStatistics {
	return LimitCounterStatistics{
		Increase:       core.CastInt64(data["increase"]),
		IncreaseAmount: core.CastInt64(data["increaseAmount"]),
	}
}

func (p LimitCounterStatistics) ToDict() map[string]interface{} {

	var increase *int64
	if p.Increase != nil {
		increase = p.Increase
	}
	var increaseAmount *int64
	if p.IncreaseAmount != nil {
		increaseAmount = p.IncreaseAmount
	}
	return map[string]interface{}{
		"increase":       increase,
		"increaseAmount": increaseAmount,
	}
}

func (p LimitCounterStatistics) Pointer() *LimitCounterStatistics {
	return &p
}

func CastLimitCounterStatisticses(data []interface{}) []LimitCounterStatistics {
	v := make([]LimitCounterStatistics, 0)
	for _, d := range data {
		v = append(v, NewLimitCounterStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitCounterStatisticsesFromDict(data []LimitCounterStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitCounterCounterDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLimitCounterCounterDistributionStatisticsFromJson(data string) LimitCounterCounterDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitCounterCounterDistributionStatisticsFromDict(dict)
}

func NewLimitCounterCounterDistributionStatisticsFromDict(data map[string]interface{}) LimitCounterCounterDistributionStatistics {
	return LimitCounterCounterDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LimitCounterCounterDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LimitCounterCounterDistributionStatistics) Pointer() *LimitCounterCounterDistributionStatistics {
	return &p
}

func CastLimitCounterCounterDistributionStatisticses(data []interface{}) []LimitCounterCounterDistributionStatistics {
	v := make([]LimitCounterCounterDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLimitCounterCounterDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitCounterCounterDistributionStatisticsesFromDict(data []LimitCounterCounterDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitCounterCounterDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewLimitCounterCounterDistributionSegmentFromJson(data string) LimitCounterCounterDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitCounterCounterDistributionSegmentFromDict(dict)
}

func NewLimitCounterCounterDistributionSegmentFromDict(data map[string]interface{}) LimitCounterCounterDistributionSegment {
	return LimitCounterCounterDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p LimitCounterCounterDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p LimitCounterCounterDistributionSegment) Pointer() *LimitCounterCounterDistributionSegment {
	return &p
}

func CastLimitCounterCounterDistributionSegments(data []interface{}) []LimitCounterCounterDistributionSegment {
	v := make([]LimitCounterCounterDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLimitCounterCounterDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitCounterCounterDistributionSegmentsFromDict(data []LimitCounterCounterDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitCounterCounterDistribution struct {
	Statistics   *LimitCounterCounterDistributionStatistics `json:"statistics"`
	Distribution []LimitCounterCounterDistributionSegment   `json:"distribution"`
}

func NewLimitCounterCounterDistributionFromJson(data string) LimitCounterCounterDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitCounterCounterDistributionFromDict(dict)
}

func NewLimitCounterCounterDistributionFromDict(data map[string]interface{}) LimitCounterCounterDistribution {
	return LimitCounterCounterDistribution{
		Statistics:   NewLimitCounterCounterDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLimitCounterCounterDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LimitCounterCounterDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLimitCounterCounterDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LimitCounterCounterDistribution) Pointer() *LimitCounterCounterDistribution {
	return &p
}

func CastLimitCounterCounterDistributions(data []interface{}) []LimitCounterCounterDistribution {
	v := make([]LimitCounterCounterDistribution, 0)
	for _, d := range data {
		v = append(v, NewLimitCounterCounterDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitCounterCounterDistributionsFromDict(data []LimitCounterCounterDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitCounterDistributions struct {
	Counter *LimitCounterCounterDistribution `json:"counter"`
}

func NewLimitCounterDistributionsFromJson(data string) LimitCounterDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitCounterDistributionsFromDict(dict)
}

func NewLimitCounterDistributionsFromDict(data map[string]interface{}) LimitCounterDistributions {
	return LimitCounterDistributions{
		Counter: NewLimitCounterCounterDistributionFromDict(core.CastMap(data["counter"])).Pointer(),
	}
}

func (p LimitCounterDistributions) ToDict() map[string]interface{} {

	var counter map[string]interface{}
	if p.Counter != nil {
		counter = p.Counter.ToDict()
	}
	return map[string]interface{}{
		"counter": counter,
	}
}

func (p LimitCounterDistributions) Pointer() *LimitCounterDistributions {
	return &p
}

func CastLimitCounterDistributionses(data []interface{}) []LimitCounterDistributions {
	v := make([]LimitCounterDistributions, 0)
	for _, d := range data {
		v = append(v, NewLimitCounterDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitCounterDistributionsesFromDict(data []LimitCounterDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitCounter struct {
	CounterId     *string                    `json:"counterId"`
	LimitName     *string                    `json:"limitName"`
	CounterName   *string                    `json:"counterName"`
	Statistics    *LimitCounterStatistics    `json:"statistics"`
	Distributions *LimitCounterDistributions `json:"distributions"`
}

func NewLimitCounterFromJson(data string) LimitCounter {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitCounterFromDict(dict)
}

func NewLimitCounterFromDict(data map[string]interface{}) LimitCounter {
	return LimitCounter{
		CounterId:     core.CastString(data["counterId"]),
		LimitName:     core.CastString(data["limitName"]),
		CounterName:   core.CastString(data["counterName"]),
		Statistics:    NewLimitCounterStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewLimitCounterDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p LimitCounter) ToDict() map[string]interface{} {

	var counterId *string
	if p.CounterId != nil {
		counterId = p.CounterId
	}
	var limitName *string
	if p.LimitName != nil {
		limitName = p.LimitName
	}
	var counterName *string
	if p.CounterName != nil {
		counterName = p.CounterName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"counterId":     counterId,
		"limitName":     limitName,
		"counterName":   counterName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p LimitCounter) Pointer() *LimitCounter {
	return &p
}

func CastLimitCounters(data []interface{}) []LimitCounter {
	v := make([]LimitCounter, 0)
	for _, d := range data {
		v = append(v, NewLimitCounterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitCountersFromDict(data []LimitCounter) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelStatistics struct {
	Increase       *int64 `json:"increase"`
	IncreaseAmount *int64 `json:"increaseAmount"`
}

func NewLimitLimitModelStatisticsFromJson(data string) LimitLimitModelStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelStatisticsFromDict(dict)
}

func NewLimitLimitModelStatisticsFromDict(data map[string]interface{}) LimitLimitModelStatistics {
	return LimitLimitModelStatistics{
		Increase:       core.CastInt64(data["increase"]),
		IncreaseAmount: core.CastInt64(data["increaseAmount"]),
	}
}

func (p LimitLimitModelStatistics) ToDict() map[string]interface{} {

	var increase *int64
	if p.Increase != nil {
		increase = p.Increase
	}
	var increaseAmount *int64
	if p.IncreaseAmount != nil {
		increaseAmount = p.IncreaseAmount
	}
	return map[string]interface{}{
		"increase":       increase,
		"increaseAmount": increaseAmount,
	}
}

func (p LimitLimitModelStatistics) Pointer() *LimitLimitModelStatistics {
	return &p
}

func CastLimitLimitModelStatisticses(data []interface{}) []LimitLimitModelStatistics {
	v := make([]LimitLimitModelStatistics, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelStatisticsesFromDict(data []LimitLimitModelStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLimitLimitModelIncreaseDistributionStatisticsFromJson(data string) LimitLimitModelIncreaseDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseDistributionStatisticsFromDict(dict)
}

func NewLimitLimitModelIncreaseDistributionStatisticsFromDict(data map[string]interface{}) LimitLimitModelIncreaseDistributionStatistics {
	return LimitLimitModelIncreaseDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LimitLimitModelIncreaseDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LimitLimitModelIncreaseDistributionStatistics) Pointer() *LimitLimitModelIncreaseDistributionStatistics {
	return &p
}

func CastLimitLimitModelIncreaseDistributionStatisticses(data []interface{}) []LimitLimitModelIncreaseDistributionStatistics {
	v := make([]LimitLimitModelIncreaseDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseDistributionStatisticsesFromDict(data []LimitLimitModelIncreaseDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseDistributionSegment struct {
	CounterName *string `json:"counterName"`
	Count       *int64  `json:"count"`
}

func NewLimitLimitModelIncreaseDistributionSegmentFromJson(data string) LimitLimitModelIncreaseDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseDistributionSegmentFromDict(dict)
}

func NewLimitLimitModelIncreaseDistributionSegmentFromDict(data map[string]interface{}) LimitLimitModelIncreaseDistributionSegment {
	return LimitLimitModelIncreaseDistributionSegment{
		CounterName: core.CastString(data["counterName"]),
		Count:       core.CastInt64(data["count"]),
	}
}

func (p LimitLimitModelIncreaseDistributionSegment) ToDict() map[string]interface{} {

	var counterName *string
	if p.CounterName != nil {
		counterName = p.CounterName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"counterName": counterName,
		"count":       count,
	}
}

func (p LimitLimitModelIncreaseDistributionSegment) Pointer() *LimitLimitModelIncreaseDistributionSegment {
	return &p
}

func CastLimitLimitModelIncreaseDistributionSegments(data []interface{}) []LimitLimitModelIncreaseDistributionSegment {
	v := make([]LimitLimitModelIncreaseDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseDistributionSegmentsFromDict(data []LimitLimitModelIncreaseDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseDistribution struct {
	Statistics   *LimitLimitModelIncreaseDistributionStatistics `json:"statistics"`
	Distribution []LimitLimitModelIncreaseDistributionSegment   `json:"distribution"`
}

func NewLimitLimitModelIncreaseDistributionFromJson(data string) LimitLimitModelIncreaseDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseDistributionFromDict(dict)
}

func NewLimitLimitModelIncreaseDistributionFromDict(data map[string]interface{}) LimitLimitModelIncreaseDistribution {
	return LimitLimitModelIncreaseDistribution{
		Statistics:   NewLimitLimitModelIncreaseDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLimitLimitModelIncreaseDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LimitLimitModelIncreaseDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLimitLimitModelIncreaseDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LimitLimitModelIncreaseDistribution) Pointer() *LimitLimitModelIncreaseDistribution {
	return &p
}

func CastLimitLimitModelIncreaseDistributions(data []interface{}) []LimitLimitModelIncreaseDistribution {
	v := make([]LimitLimitModelIncreaseDistribution, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseDistributionsFromDict(data []LimitLimitModelIncreaseDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLimitLimitModelIncreaseAmountDistributionStatisticsFromJson(data string) LimitLimitModelIncreaseAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseAmountDistributionStatisticsFromDict(dict)
}

func NewLimitLimitModelIncreaseAmountDistributionStatisticsFromDict(data map[string]interface{}) LimitLimitModelIncreaseAmountDistributionStatistics {
	return LimitLimitModelIncreaseAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LimitLimitModelIncreaseAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LimitLimitModelIncreaseAmountDistributionStatistics) Pointer() *LimitLimitModelIncreaseAmountDistributionStatistics {
	return &p
}

func CastLimitLimitModelIncreaseAmountDistributionStatisticses(data []interface{}) []LimitLimitModelIncreaseAmountDistributionStatistics {
	v := make([]LimitLimitModelIncreaseAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseAmountDistributionStatisticsesFromDict(data []LimitLimitModelIncreaseAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseAmountDistributionSegment struct {
	CounterName *string `json:"counterName"`
	Sum         *int64  `json:"sum"`
}

func NewLimitLimitModelIncreaseAmountDistributionSegmentFromJson(data string) LimitLimitModelIncreaseAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseAmountDistributionSegmentFromDict(dict)
}

func NewLimitLimitModelIncreaseAmountDistributionSegmentFromDict(data map[string]interface{}) LimitLimitModelIncreaseAmountDistributionSegment {
	return LimitLimitModelIncreaseAmountDistributionSegment{
		CounterName: core.CastString(data["counterName"]),
		Sum:         core.CastInt64(data["sum"]),
	}
}

func (p LimitLimitModelIncreaseAmountDistributionSegment) ToDict() map[string]interface{} {

	var counterName *string
	if p.CounterName != nil {
		counterName = p.CounterName
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"counterName": counterName,
		"sum":         sum,
	}
}

func (p LimitLimitModelIncreaseAmountDistributionSegment) Pointer() *LimitLimitModelIncreaseAmountDistributionSegment {
	return &p
}

func CastLimitLimitModelIncreaseAmountDistributionSegments(data []interface{}) []LimitLimitModelIncreaseAmountDistributionSegment {
	v := make([]LimitLimitModelIncreaseAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseAmountDistributionSegmentsFromDict(data []LimitLimitModelIncreaseAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseAmountDistribution struct {
	Statistics   *LimitLimitModelIncreaseAmountDistributionStatistics `json:"statistics"`
	Distribution []LimitLimitModelIncreaseAmountDistributionSegment   `json:"distribution"`
}

func NewLimitLimitModelIncreaseAmountDistributionFromJson(data string) LimitLimitModelIncreaseAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseAmountDistributionFromDict(dict)
}

func NewLimitLimitModelIncreaseAmountDistributionFromDict(data map[string]interface{}) LimitLimitModelIncreaseAmountDistribution {
	return LimitLimitModelIncreaseAmountDistribution{
		Statistics:   NewLimitLimitModelIncreaseAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLimitLimitModelIncreaseAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LimitLimitModelIncreaseAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLimitLimitModelIncreaseAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LimitLimitModelIncreaseAmountDistribution) Pointer() *LimitLimitModelIncreaseAmountDistribution {
	return &p
}

func CastLimitLimitModelIncreaseAmountDistributions(data []interface{}) []LimitLimitModelIncreaseAmountDistribution {
	v := make([]LimitLimitModelIncreaseAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseAmountDistributionsFromDict(data []LimitLimitModelIncreaseAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLimitLimitModelIncreaseByUserDistributionStatisticsFromJson(data string) LimitLimitModelIncreaseByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseByUserDistributionStatisticsFromDict(dict)
}

func NewLimitLimitModelIncreaseByUserDistributionStatisticsFromDict(data map[string]interface{}) LimitLimitModelIncreaseByUserDistributionStatistics {
	return LimitLimitModelIncreaseByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LimitLimitModelIncreaseByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LimitLimitModelIncreaseByUserDistributionStatistics) Pointer() *LimitLimitModelIncreaseByUserDistributionStatistics {
	return &p
}

func CastLimitLimitModelIncreaseByUserDistributionStatisticses(data []interface{}) []LimitLimitModelIncreaseByUserDistributionStatistics {
	v := make([]LimitLimitModelIncreaseByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseByUserDistributionStatisticsesFromDict(data []LimitLimitModelIncreaseByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewLimitLimitModelIncreaseByUserDistributionSegmentFromJson(data string) LimitLimitModelIncreaseByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseByUserDistributionSegmentFromDict(dict)
}

func NewLimitLimitModelIncreaseByUserDistributionSegmentFromDict(data map[string]interface{}) LimitLimitModelIncreaseByUserDistributionSegment {
	return LimitLimitModelIncreaseByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p LimitLimitModelIncreaseByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p LimitLimitModelIncreaseByUserDistributionSegment) Pointer() *LimitLimitModelIncreaseByUserDistributionSegment {
	return &p
}

func CastLimitLimitModelIncreaseByUserDistributionSegments(data []interface{}) []LimitLimitModelIncreaseByUserDistributionSegment {
	v := make([]LimitLimitModelIncreaseByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseByUserDistributionSegmentsFromDict(data []LimitLimitModelIncreaseByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseByUserDistribution struct {
	Statistics   *LimitLimitModelIncreaseByUserDistributionStatistics `json:"statistics"`
	Distribution []LimitLimitModelIncreaseByUserDistributionSegment   `json:"distribution"`
}

func NewLimitLimitModelIncreaseByUserDistributionFromJson(data string) LimitLimitModelIncreaseByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseByUserDistributionFromDict(dict)
}

func NewLimitLimitModelIncreaseByUserDistributionFromDict(data map[string]interface{}) LimitLimitModelIncreaseByUserDistribution {
	return LimitLimitModelIncreaseByUserDistribution{
		Statistics:   NewLimitLimitModelIncreaseByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLimitLimitModelIncreaseByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LimitLimitModelIncreaseByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLimitLimitModelIncreaseByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LimitLimitModelIncreaseByUserDistribution) Pointer() *LimitLimitModelIncreaseByUserDistribution {
	return &p
}

func CastLimitLimitModelIncreaseByUserDistributions(data []interface{}) []LimitLimitModelIncreaseByUserDistribution {
	v := make([]LimitLimitModelIncreaseByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseByUserDistributionsFromDict(data []LimitLimitModelIncreaseByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseAmountByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLimitLimitModelIncreaseAmountByUserDistributionStatisticsFromJson(data string) LimitLimitModelIncreaseAmountByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseAmountByUserDistributionStatisticsFromDict(dict)
}

func NewLimitLimitModelIncreaseAmountByUserDistributionStatisticsFromDict(data map[string]interface{}) LimitLimitModelIncreaseAmountByUserDistributionStatistics {
	return LimitLimitModelIncreaseAmountByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LimitLimitModelIncreaseAmountByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LimitLimitModelIncreaseAmountByUserDistributionStatistics) Pointer() *LimitLimitModelIncreaseAmountByUserDistributionStatistics {
	return &p
}

func CastLimitLimitModelIncreaseAmountByUserDistributionStatisticses(data []interface{}) []LimitLimitModelIncreaseAmountByUserDistributionStatistics {
	v := make([]LimitLimitModelIncreaseAmountByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseAmountByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseAmountByUserDistributionStatisticsesFromDict(data []LimitLimitModelIncreaseAmountByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseAmountByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewLimitLimitModelIncreaseAmountByUserDistributionSegmentFromJson(data string) LimitLimitModelIncreaseAmountByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseAmountByUserDistributionSegmentFromDict(dict)
}

func NewLimitLimitModelIncreaseAmountByUserDistributionSegmentFromDict(data map[string]interface{}) LimitLimitModelIncreaseAmountByUserDistributionSegment {
	return LimitLimitModelIncreaseAmountByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p LimitLimitModelIncreaseAmountByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p LimitLimitModelIncreaseAmountByUserDistributionSegment) Pointer() *LimitLimitModelIncreaseAmountByUserDistributionSegment {
	return &p
}

func CastLimitLimitModelIncreaseAmountByUserDistributionSegments(data []interface{}) []LimitLimitModelIncreaseAmountByUserDistributionSegment {
	v := make([]LimitLimitModelIncreaseAmountByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseAmountByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseAmountByUserDistributionSegmentsFromDict(data []LimitLimitModelIncreaseAmountByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelIncreaseAmountByUserDistribution struct {
	Statistics   *LimitLimitModelIncreaseAmountByUserDistributionStatistics `json:"statistics"`
	Distribution []LimitLimitModelIncreaseAmountByUserDistributionSegment   `json:"distribution"`
}

func NewLimitLimitModelIncreaseAmountByUserDistributionFromJson(data string) LimitLimitModelIncreaseAmountByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelIncreaseAmountByUserDistributionFromDict(dict)
}

func NewLimitLimitModelIncreaseAmountByUserDistributionFromDict(data map[string]interface{}) LimitLimitModelIncreaseAmountByUserDistribution {
	return LimitLimitModelIncreaseAmountByUserDistribution{
		Statistics:   NewLimitLimitModelIncreaseAmountByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLimitLimitModelIncreaseAmountByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LimitLimitModelIncreaseAmountByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLimitLimitModelIncreaseAmountByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LimitLimitModelIncreaseAmountByUserDistribution) Pointer() *LimitLimitModelIncreaseAmountByUserDistribution {
	return &p
}

func CastLimitLimitModelIncreaseAmountByUserDistributions(data []interface{}) []LimitLimitModelIncreaseAmountByUserDistribution {
	v := make([]LimitLimitModelIncreaseAmountByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelIncreaseAmountByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelIncreaseAmountByUserDistributionsFromDict(data []LimitLimitModelIncreaseAmountByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModelDistributions struct {
	Increase             *LimitLimitModelIncreaseDistribution             `json:"increase"`
	IncreaseAmount       *LimitLimitModelIncreaseAmountDistribution       `json:"increaseAmount"`
	IncreaseByUser       *LimitLimitModelIncreaseByUserDistribution       `json:"increaseByUser"`
	IncreaseAmountByUser *LimitLimitModelIncreaseAmountByUserDistribution `json:"increaseAmountByUser"`
}

func NewLimitLimitModelDistributionsFromJson(data string) LimitLimitModelDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelDistributionsFromDict(dict)
}

func NewLimitLimitModelDistributionsFromDict(data map[string]interface{}) LimitLimitModelDistributions {
	return LimitLimitModelDistributions{
		Increase:             NewLimitLimitModelIncreaseDistributionFromDict(core.CastMap(data["increase"])).Pointer(),
		IncreaseAmount:       NewLimitLimitModelIncreaseAmountDistributionFromDict(core.CastMap(data["increaseAmount"])).Pointer(),
		IncreaseByUser:       NewLimitLimitModelIncreaseByUserDistributionFromDict(core.CastMap(data["increaseByUser"])).Pointer(),
		IncreaseAmountByUser: NewLimitLimitModelIncreaseAmountByUserDistributionFromDict(core.CastMap(data["increaseAmountByUser"])).Pointer(),
	}
}

func (p LimitLimitModelDistributions) ToDict() map[string]interface{} {

	var increase map[string]interface{}
	if p.Increase != nil {
		increase = p.Increase.ToDict()
	}
	var increaseAmount map[string]interface{}
	if p.IncreaseAmount != nil {
		increaseAmount = p.IncreaseAmount.ToDict()
	}
	var increaseByUser map[string]interface{}
	if p.IncreaseByUser != nil {
		increaseByUser = p.IncreaseByUser.ToDict()
	}
	var increaseAmountByUser map[string]interface{}
	if p.IncreaseAmountByUser != nil {
		increaseAmountByUser = p.IncreaseAmountByUser.ToDict()
	}
	return map[string]interface{}{
		"increase":             increase,
		"increaseAmount":       increaseAmount,
		"increaseByUser":       increaseByUser,
		"increaseAmountByUser": increaseAmountByUser,
	}
}

func (p LimitLimitModelDistributions) Pointer() *LimitLimitModelDistributions {
	return &p
}

func CastLimitLimitModelDistributionses(data []interface{}) []LimitLimitModelDistributions {
	v := make([]LimitLimitModelDistributions, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelDistributionsesFromDict(data []LimitLimitModelDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitLimitModel struct {
	LimitModelId  *string                       `json:"limitModelId"`
	LimitName     *string                       `json:"limitName"`
	Statistics    *LimitLimitModelStatistics    `json:"statistics"`
	Distributions *LimitLimitModelDistributions `json:"distributions"`
	Counters      []LimitCounter                `json:"counters"`
}

func NewLimitLimitModelFromJson(data string) LimitLimitModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitLimitModelFromDict(dict)
}

func NewLimitLimitModelFromDict(data map[string]interface{}) LimitLimitModel {
	return LimitLimitModel{
		LimitModelId:  core.CastString(data["limitModelId"]),
		LimitName:     core.CastString(data["limitName"]),
		Statistics:    NewLimitLimitModelStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewLimitLimitModelDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		Counters:      CastLimitCounters(core.CastArray(data["counters"])),
	}
}

func (p LimitLimitModel) ToDict() map[string]interface{} {

	var limitModelId *string
	if p.LimitModelId != nil {
		limitModelId = p.LimitModelId
	}
	var limitName *string
	if p.LimitName != nil {
		limitName = p.LimitName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var counters []interface{}
	if p.Counters != nil {
		counters = CastLimitCountersFromDict(
			p.Counters,
		)
	}
	return map[string]interface{}{
		"limitModelId":  limitModelId,
		"limitName":     limitName,
		"statistics":    statistics,
		"distributions": distributions,
		"counters":      counters,
	}
}

func (p LimitLimitModel) Pointer() *LimitLimitModel {
	return &p
}

func CastLimitLimitModels(data []interface{}) []LimitLimitModel {
	v := make([]LimitLimitModel, 0)
	for _, d := range data {
		v = append(v, NewLimitLimitModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitLimitModelsFromDict(data []LimitLimitModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitNamespaceStatistics struct {
	Increase *int64 `json:"increase"`
}

func NewLimitNamespaceStatisticsFromJson(data string) LimitNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitNamespaceStatisticsFromDict(dict)
}

func NewLimitNamespaceStatisticsFromDict(data map[string]interface{}) LimitNamespaceStatistics {
	return LimitNamespaceStatistics{
		Increase: core.CastInt64(data["increase"]),
	}
}

func (p LimitNamespaceStatistics) ToDict() map[string]interface{} {

	var increase *int64
	if p.Increase != nil {
		increase = p.Increase
	}
	return map[string]interface{}{
		"increase": increase,
	}
}

func (p LimitNamespaceStatistics) Pointer() *LimitNamespaceStatistics {
	return &p
}

func CastLimitNamespaceStatisticses(data []interface{}) []LimitNamespaceStatistics {
	v := make([]LimitNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewLimitNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitNamespaceStatisticsesFromDict(data []LimitNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitNamespaceIncreaseDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLimitNamespaceIncreaseDistributionStatisticsFromJson(data string) LimitNamespaceIncreaseDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitNamespaceIncreaseDistributionStatisticsFromDict(dict)
}

func NewLimitNamespaceIncreaseDistributionStatisticsFromDict(data map[string]interface{}) LimitNamespaceIncreaseDistributionStatistics {
	return LimitNamespaceIncreaseDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LimitNamespaceIncreaseDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LimitNamespaceIncreaseDistributionStatistics) Pointer() *LimitNamespaceIncreaseDistributionStatistics {
	return &p
}

func CastLimitNamespaceIncreaseDistributionStatisticses(data []interface{}) []LimitNamespaceIncreaseDistributionStatistics {
	v := make([]LimitNamespaceIncreaseDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLimitNamespaceIncreaseDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitNamespaceIncreaseDistributionStatisticsesFromDict(data []LimitNamespaceIncreaseDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitNamespaceIncreaseDistributionSegment struct {
	LimitName *string `json:"limitName"`
	Count     *int64  `json:"count"`
}

func NewLimitNamespaceIncreaseDistributionSegmentFromJson(data string) LimitNamespaceIncreaseDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitNamespaceIncreaseDistributionSegmentFromDict(dict)
}

func NewLimitNamespaceIncreaseDistributionSegmentFromDict(data map[string]interface{}) LimitNamespaceIncreaseDistributionSegment {
	return LimitNamespaceIncreaseDistributionSegment{
		LimitName: core.CastString(data["limitName"]),
		Count:     core.CastInt64(data["count"]),
	}
}

func (p LimitNamespaceIncreaseDistributionSegment) ToDict() map[string]interface{} {

	var limitName *string
	if p.LimitName != nil {
		limitName = p.LimitName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"limitName": limitName,
		"count":     count,
	}
}

func (p LimitNamespaceIncreaseDistributionSegment) Pointer() *LimitNamespaceIncreaseDistributionSegment {
	return &p
}

func CastLimitNamespaceIncreaseDistributionSegments(data []interface{}) []LimitNamespaceIncreaseDistributionSegment {
	v := make([]LimitNamespaceIncreaseDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLimitNamespaceIncreaseDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitNamespaceIncreaseDistributionSegmentsFromDict(data []LimitNamespaceIncreaseDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitNamespaceIncreaseDistribution struct {
	Statistics   *LimitNamespaceIncreaseDistributionStatistics `json:"statistics"`
	Distribution []LimitNamespaceIncreaseDistributionSegment   `json:"distribution"`
}

func NewLimitNamespaceIncreaseDistributionFromJson(data string) LimitNamespaceIncreaseDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitNamespaceIncreaseDistributionFromDict(dict)
}

func NewLimitNamespaceIncreaseDistributionFromDict(data map[string]interface{}) LimitNamespaceIncreaseDistribution {
	return LimitNamespaceIncreaseDistribution{
		Statistics:   NewLimitNamespaceIncreaseDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLimitNamespaceIncreaseDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LimitNamespaceIncreaseDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLimitNamespaceIncreaseDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LimitNamespaceIncreaseDistribution) Pointer() *LimitNamespaceIncreaseDistribution {
	return &p
}

func CastLimitNamespaceIncreaseDistributions(data []interface{}) []LimitNamespaceIncreaseDistribution {
	v := make([]LimitNamespaceIncreaseDistribution, 0)
	for _, d := range data {
		v = append(v, NewLimitNamespaceIncreaseDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitNamespaceIncreaseDistributionsFromDict(data []LimitNamespaceIncreaseDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitNamespaceDistributions struct {
	Increase *LimitNamespaceIncreaseDistribution `json:"increase"`
}

func NewLimitNamespaceDistributionsFromJson(data string) LimitNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitNamespaceDistributionsFromDict(dict)
}

func NewLimitNamespaceDistributionsFromDict(data map[string]interface{}) LimitNamespaceDistributions {
	return LimitNamespaceDistributions{
		Increase: NewLimitNamespaceIncreaseDistributionFromDict(core.CastMap(data["increase"])).Pointer(),
	}
}

func (p LimitNamespaceDistributions) ToDict() map[string]interface{} {

	var increase map[string]interface{}
	if p.Increase != nil {
		increase = p.Increase.ToDict()
	}
	return map[string]interface{}{
		"increase": increase,
	}
}

func (p LimitNamespaceDistributions) Pointer() *LimitNamespaceDistributions {
	return &p
}

func CastLimitNamespaceDistributionses(data []interface{}) []LimitNamespaceDistributions {
	v := make([]LimitNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewLimitNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitNamespaceDistributionsesFromDict(data []LimitNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LimitNamespace struct {
	NamespaceId   *string                      `json:"namespaceId"`
	Year          *int32                       `json:"year"`
	Month         *int32                       `json:"month"`
	Day           *int32                       `json:"day"`
	NamespaceName *string                      `json:"namespaceName"`
	Statistics    *LimitNamespaceStatistics    `json:"statistics"`
	Distributions *LimitNamespaceDistributions `json:"distributions"`
	LimitModels   []LimitLimitModel            `json:"limitModels"`
}

func NewLimitNamespaceFromJson(data string) LimitNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitNamespaceFromDict(dict)
}

func NewLimitNamespaceFromDict(data map[string]interface{}) LimitNamespace {
	return LimitNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewLimitNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewLimitNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		LimitModels:   CastLimitLimitModels(core.CastArray(data["limitModels"])),
	}
}

func (p LimitNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var limitModels []interface{}
	if p.LimitModels != nil {
		limitModels = CastLimitLimitModelsFromDict(
			p.LimitModels,
		)
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
		"limitModels":   limitModels,
	}
}

func (p LimitNamespace) Pointer() *LimitNamespace {
	return &p
}

func CastLimitNamespaces(data []interface{}) []LimitNamespace {
	v := make([]LimitNamespace, 0)
	for _, d := range data {
		v = append(v, NewLimitNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitNamespacesFromDict(data []LimitNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryLotteryStatistics struct {
	Draw       *int64 `json:"draw"`
	DrawAmount *int64 `json:"drawAmount"`
}

func NewLotteryLotteryStatisticsFromJson(data string) LotteryLotteryStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryLotteryStatisticsFromDict(dict)
}

func NewLotteryLotteryStatisticsFromDict(data map[string]interface{}) LotteryLotteryStatistics {
	return LotteryLotteryStatistics{
		Draw:       core.CastInt64(data["draw"]),
		DrawAmount: core.CastInt64(data["drawAmount"]),
	}
}

func (p LotteryLotteryStatistics) ToDict() map[string]interface{} {

	var draw *int64
	if p.Draw != nil {
		draw = p.Draw
	}
	var drawAmount *int64
	if p.DrawAmount != nil {
		drawAmount = p.DrawAmount
	}
	return map[string]interface{}{
		"draw":       draw,
		"drawAmount": drawAmount,
	}
}

func (p LotteryLotteryStatistics) Pointer() *LotteryLotteryStatistics {
	return &p
}

func CastLotteryLotteryStatisticses(data []interface{}) []LotteryLotteryStatistics {
	v := make([]LotteryLotteryStatistics, 0)
	for _, d := range data {
		v = append(v, NewLotteryLotteryStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryLotteryStatisticsesFromDict(data []LotteryLotteryStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryLotteryDrawResultDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLotteryLotteryDrawResultDistributionStatisticsFromJson(data string) LotteryLotteryDrawResultDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryLotteryDrawResultDistributionStatisticsFromDict(dict)
}

func NewLotteryLotteryDrawResultDistributionStatisticsFromDict(data map[string]interface{}) LotteryLotteryDrawResultDistributionStatistics {
	return LotteryLotteryDrawResultDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LotteryLotteryDrawResultDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LotteryLotteryDrawResultDistributionStatistics) Pointer() *LotteryLotteryDrawResultDistributionStatistics {
	return &p
}

func CastLotteryLotteryDrawResultDistributionStatisticses(data []interface{}) []LotteryLotteryDrawResultDistributionStatistics {
	v := make([]LotteryLotteryDrawResultDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLotteryLotteryDrawResultDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryLotteryDrawResultDistributionStatisticsesFromDict(data []LotteryLotteryDrawResultDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryLotteryDrawResultDistributionSegment struct {
	PrizeId *string `json:"prizeId"`
	Count   *int64  `json:"count"`
}

func NewLotteryLotteryDrawResultDistributionSegmentFromJson(data string) LotteryLotteryDrawResultDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryLotteryDrawResultDistributionSegmentFromDict(dict)
}

func NewLotteryLotteryDrawResultDistributionSegmentFromDict(data map[string]interface{}) LotteryLotteryDrawResultDistributionSegment {
	return LotteryLotteryDrawResultDistributionSegment{
		PrizeId: core.CastString(data["prizeId"]),
		Count:   core.CastInt64(data["count"]),
	}
}

func (p LotteryLotteryDrawResultDistributionSegment) ToDict() map[string]interface{} {

	var prizeId *string
	if p.PrizeId != nil {
		prizeId = p.PrizeId
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"prizeId": prizeId,
		"count":   count,
	}
}

func (p LotteryLotteryDrawResultDistributionSegment) Pointer() *LotteryLotteryDrawResultDistributionSegment {
	return &p
}

func CastLotteryLotteryDrawResultDistributionSegments(data []interface{}) []LotteryLotteryDrawResultDistributionSegment {
	v := make([]LotteryLotteryDrawResultDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLotteryLotteryDrawResultDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryLotteryDrawResultDistributionSegmentsFromDict(data []LotteryLotteryDrawResultDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryLotteryDrawResultDistribution struct {
	Statistics   *LotteryLotteryDrawResultDistributionStatistics `json:"statistics"`
	Distribution []LotteryLotteryDrawResultDistributionSegment   `json:"distribution"`
}

func NewLotteryLotteryDrawResultDistributionFromJson(data string) LotteryLotteryDrawResultDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryLotteryDrawResultDistributionFromDict(dict)
}

func NewLotteryLotteryDrawResultDistributionFromDict(data map[string]interface{}) LotteryLotteryDrawResultDistribution {
	return LotteryLotteryDrawResultDistribution{
		Statistics:   NewLotteryLotteryDrawResultDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLotteryLotteryDrawResultDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LotteryLotteryDrawResultDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLotteryLotteryDrawResultDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LotteryLotteryDrawResultDistribution) Pointer() *LotteryLotteryDrawResultDistribution {
	return &p
}

func CastLotteryLotteryDrawResultDistributions(data []interface{}) []LotteryLotteryDrawResultDistribution {
	v := make([]LotteryLotteryDrawResultDistribution, 0)
	for _, d := range data {
		v = append(v, NewLotteryLotteryDrawResultDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryLotteryDrawResultDistributionsFromDict(data []LotteryLotteryDrawResultDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryLotteryDistributions struct {
	DrawResult *LotteryLotteryDrawResultDistribution `json:"drawResult"`
}

func NewLotteryLotteryDistributionsFromJson(data string) LotteryLotteryDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryLotteryDistributionsFromDict(dict)
}

func NewLotteryLotteryDistributionsFromDict(data map[string]interface{}) LotteryLotteryDistributions {
	return LotteryLotteryDistributions{
		DrawResult: NewLotteryLotteryDrawResultDistributionFromDict(core.CastMap(data["drawResult"])).Pointer(),
	}
}

func (p LotteryLotteryDistributions) ToDict() map[string]interface{} {

	var drawResult map[string]interface{}
	if p.DrawResult != nil {
		drawResult = p.DrawResult.ToDict()
	}
	return map[string]interface{}{
		"drawResult": drawResult,
	}
}

func (p LotteryLotteryDistributions) Pointer() *LotteryLotteryDistributions {
	return &p
}

func CastLotteryLotteryDistributionses(data []interface{}) []LotteryLotteryDistributions {
	v := make([]LotteryLotteryDistributions, 0)
	for _, d := range data {
		v = append(v, NewLotteryLotteryDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryLotteryDistributionsesFromDict(data []LotteryLotteryDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryLottery struct {
	LotteryId     *string                      `json:"lotteryId"`
	LotteryName   *string                      `json:"lotteryName"`
	Statistics    *LotteryLotteryStatistics    `json:"statistics"`
	Distributions *LotteryLotteryDistributions `json:"distributions"`
}

func NewLotteryLotteryFromJson(data string) LotteryLottery {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryLotteryFromDict(dict)
}

func NewLotteryLotteryFromDict(data map[string]interface{}) LotteryLottery {
	return LotteryLottery{
		LotteryId:     core.CastString(data["lotteryId"]),
		LotteryName:   core.CastString(data["lotteryName"]),
		Statistics:    NewLotteryLotteryStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewLotteryLotteryDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p LotteryLottery) ToDict() map[string]interface{} {

	var lotteryId *string
	if p.LotteryId != nil {
		lotteryId = p.LotteryId
	}
	var lotteryName *string
	if p.LotteryName != nil {
		lotteryName = p.LotteryName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"lotteryId":     lotteryId,
		"lotteryName":   lotteryName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p LotteryLottery) Pointer() *LotteryLottery {
	return &p
}

func CastLotteryLotteries(data []interface{}) []LotteryLottery {
	v := make([]LotteryLottery, 0)
	for _, d := range data {
		v = append(v, NewLotteryLotteryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryLotteriesFromDict(data []LotteryLottery) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceStatistics struct {
	Draw       *int64 `json:"draw"`
	DrawAmount *int64 `json:"drawAmount"`
}

func NewLotteryNamespaceStatisticsFromJson(data string) LotteryNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceStatisticsFromDict(dict)
}

func NewLotteryNamespaceStatisticsFromDict(data map[string]interface{}) LotteryNamespaceStatistics {
	return LotteryNamespaceStatistics{
		Draw:       core.CastInt64(data["draw"]),
		DrawAmount: core.CastInt64(data["drawAmount"]),
	}
}

func (p LotteryNamespaceStatistics) ToDict() map[string]interface{} {

	var draw *int64
	if p.Draw != nil {
		draw = p.Draw
	}
	var drawAmount *int64
	if p.DrawAmount != nil {
		drawAmount = p.DrawAmount
	}
	return map[string]interface{}{
		"draw":       draw,
		"drawAmount": drawAmount,
	}
}

func (p LotteryNamespaceStatistics) Pointer() *LotteryNamespaceStatistics {
	return &p
}

func CastLotteryNamespaceStatisticses(data []interface{}) []LotteryNamespaceStatistics {
	v := make([]LotteryNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceStatisticsesFromDict(data []LotteryNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLotteryNamespaceDrawDistributionStatisticsFromJson(data string) LotteryNamespaceDrawDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawDistributionStatisticsFromDict(dict)
}

func NewLotteryNamespaceDrawDistributionStatisticsFromDict(data map[string]interface{}) LotteryNamespaceDrawDistributionStatistics {
	return LotteryNamespaceDrawDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LotteryNamespaceDrawDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LotteryNamespaceDrawDistributionStatistics) Pointer() *LotteryNamespaceDrawDistributionStatistics {
	return &p
}

func CastLotteryNamespaceDrawDistributionStatisticses(data []interface{}) []LotteryNamespaceDrawDistributionStatistics {
	v := make([]LotteryNamespaceDrawDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawDistributionStatisticsesFromDict(data []LotteryNamespaceDrawDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawDistributionSegment struct {
	LotteryName *string `json:"lotteryName"`
	Count       *int64  `json:"count"`
}

func NewLotteryNamespaceDrawDistributionSegmentFromJson(data string) LotteryNamespaceDrawDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawDistributionSegmentFromDict(dict)
}

func NewLotteryNamespaceDrawDistributionSegmentFromDict(data map[string]interface{}) LotteryNamespaceDrawDistributionSegment {
	return LotteryNamespaceDrawDistributionSegment{
		LotteryName: core.CastString(data["lotteryName"]),
		Count:       core.CastInt64(data["count"]),
	}
}

func (p LotteryNamespaceDrawDistributionSegment) ToDict() map[string]interface{} {

	var lotteryName *string
	if p.LotteryName != nil {
		lotteryName = p.LotteryName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"lotteryName": lotteryName,
		"count":       count,
	}
}

func (p LotteryNamespaceDrawDistributionSegment) Pointer() *LotteryNamespaceDrawDistributionSegment {
	return &p
}

func CastLotteryNamespaceDrawDistributionSegments(data []interface{}) []LotteryNamespaceDrawDistributionSegment {
	v := make([]LotteryNamespaceDrawDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawDistributionSegmentsFromDict(data []LotteryNamespaceDrawDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawDistribution struct {
	Statistics   *LotteryNamespaceDrawDistributionStatistics `json:"statistics"`
	Distribution []LotteryNamespaceDrawDistributionSegment   `json:"distribution"`
}

func NewLotteryNamespaceDrawDistributionFromJson(data string) LotteryNamespaceDrawDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawDistributionFromDict(dict)
}

func NewLotteryNamespaceDrawDistributionFromDict(data map[string]interface{}) LotteryNamespaceDrawDistribution {
	return LotteryNamespaceDrawDistribution{
		Statistics:   NewLotteryNamespaceDrawDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLotteryNamespaceDrawDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LotteryNamespaceDrawDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLotteryNamespaceDrawDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LotteryNamespaceDrawDistribution) Pointer() *LotteryNamespaceDrawDistribution {
	return &p
}

func CastLotteryNamespaceDrawDistributions(data []interface{}) []LotteryNamespaceDrawDistribution {
	v := make([]LotteryNamespaceDrawDistribution, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawDistributionsFromDict(data []LotteryNamespaceDrawDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLotteryNamespaceDrawAmountDistributionStatisticsFromJson(data string) LotteryNamespaceDrawAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawAmountDistributionStatisticsFromDict(dict)
}

func NewLotteryNamespaceDrawAmountDistributionStatisticsFromDict(data map[string]interface{}) LotteryNamespaceDrawAmountDistributionStatistics {
	return LotteryNamespaceDrawAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LotteryNamespaceDrawAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LotteryNamespaceDrawAmountDistributionStatistics) Pointer() *LotteryNamespaceDrawAmountDistributionStatistics {
	return &p
}

func CastLotteryNamespaceDrawAmountDistributionStatisticses(data []interface{}) []LotteryNamespaceDrawAmountDistributionStatistics {
	v := make([]LotteryNamespaceDrawAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawAmountDistributionStatisticsesFromDict(data []LotteryNamespaceDrawAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawAmountDistributionSegment struct {
	LotteryName *string `json:"lotteryName"`
	Sum         *int64  `json:"sum"`
}

func NewLotteryNamespaceDrawAmountDistributionSegmentFromJson(data string) LotteryNamespaceDrawAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawAmountDistributionSegmentFromDict(dict)
}

func NewLotteryNamespaceDrawAmountDistributionSegmentFromDict(data map[string]interface{}) LotteryNamespaceDrawAmountDistributionSegment {
	return LotteryNamespaceDrawAmountDistributionSegment{
		LotteryName: core.CastString(data["lotteryName"]),
		Sum:         core.CastInt64(data["sum"]),
	}
}

func (p LotteryNamespaceDrawAmountDistributionSegment) ToDict() map[string]interface{} {

	var lotteryName *string
	if p.LotteryName != nil {
		lotteryName = p.LotteryName
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"lotteryName": lotteryName,
		"sum":         sum,
	}
}

func (p LotteryNamespaceDrawAmountDistributionSegment) Pointer() *LotteryNamespaceDrawAmountDistributionSegment {
	return &p
}

func CastLotteryNamespaceDrawAmountDistributionSegments(data []interface{}) []LotteryNamespaceDrawAmountDistributionSegment {
	v := make([]LotteryNamespaceDrawAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawAmountDistributionSegmentsFromDict(data []LotteryNamespaceDrawAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawAmountDistribution struct {
	Statistics   *LotteryNamespaceDrawAmountDistributionStatistics `json:"statistics"`
	Distribution []LotteryNamespaceDrawAmountDistributionSegment   `json:"distribution"`
}

func NewLotteryNamespaceDrawAmountDistributionFromJson(data string) LotteryNamespaceDrawAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawAmountDistributionFromDict(dict)
}

func NewLotteryNamespaceDrawAmountDistributionFromDict(data map[string]interface{}) LotteryNamespaceDrawAmountDistribution {
	return LotteryNamespaceDrawAmountDistribution{
		Statistics:   NewLotteryNamespaceDrawAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLotteryNamespaceDrawAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LotteryNamespaceDrawAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLotteryNamespaceDrawAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LotteryNamespaceDrawAmountDistribution) Pointer() *LotteryNamespaceDrawAmountDistribution {
	return &p
}

func CastLotteryNamespaceDrawAmountDistributions(data []interface{}) []LotteryNamespaceDrawAmountDistribution {
	v := make([]LotteryNamespaceDrawAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawAmountDistributionsFromDict(data []LotteryNamespaceDrawAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLotteryNamespaceDrawByUserDistributionStatisticsFromJson(data string) LotteryNamespaceDrawByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawByUserDistributionStatisticsFromDict(dict)
}

func NewLotteryNamespaceDrawByUserDistributionStatisticsFromDict(data map[string]interface{}) LotteryNamespaceDrawByUserDistributionStatistics {
	return LotteryNamespaceDrawByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LotteryNamespaceDrawByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LotteryNamespaceDrawByUserDistributionStatistics) Pointer() *LotteryNamespaceDrawByUserDistributionStatistics {
	return &p
}

func CastLotteryNamespaceDrawByUserDistributionStatisticses(data []interface{}) []LotteryNamespaceDrawByUserDistributionStatistics {
	v := make([]LotteryNamespaceDrawByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawByUserDistributionStatisticsesFromDict(data []LotteryNamespaceDrawByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewLotteryNamespaceDrawByUserDistributionSegmentFromJson(data string) LotteryNamespaceDrawByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawByUserDistributionSegmentFromDict(dict)
}

func NewLotteryNamespaceDrawByUserDistributionSegmentFromDict(data map[string]interface{}) LotteryNamespaceDrawByUserDistributionSegment {
	return LotteryNamespaceDrawByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p LotteryNamespaceDrawByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p LotteryNamespaceDrawByUserDistributionSegment) Pointer() *LotteryNamespaceDrawByUserDistributionSegment {
	return &p
}

func CastLotteryNamespaceDrawByUserDistributionSegments(data []interface{}) []LotteryNamespaceDrawByUserDistributionSegment {
	v := make([]LotteryNamespaceDrawByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawByUserDistributionSegmentsFromDict(data []LotteryNamespaceDrawByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawByUserDistribution struct {
	Statistics   *LotteryNamespaceDrawByUserDistributionStatistics `json:"statistics"`
	Distribution []LotteryNamespaceDrawByUserDistributionSegment   `json:"distribution"`
}

func NewLotteryNamespaceDrawByUserDistributionFromJson(data string) LotteryNamespaceDrawByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawByUserDistributionFromDict(dict)
}

func NewLotteryNamespaceDrawByUserDistributionFromDict(data map[string]interface{}) LotteryNamespaceDrawByUserDistribution {
	return LotteryNamespaceDrawByUserDistribution{
		Statistics:   NewLotteryNamespaceDrawByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLotteryNamespaceDrawByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LotteryNamespaceDrawByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLotteryNamespaceDrawByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LotteryNamespaceDrawByUserDistribution) Pointer() *LotteryNamespaceDrawByUserDistribution {
	return &p
}

func CastLotteryNamespaceDrawByUserDistributions(data []interface{}) []LotteryNamespaceDrawByUserDistribution {
	v := make([]LotteryNamespaceDrawByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawByUserDistributionsFromDict(data []LotteryNamespaceDrawByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawAmountByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewLotteryNamespaceDrawAmountByUserDistributionStatisticsFromJson(data string) LotteryNamespaceDrawAmountByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawAmountByUserDistributionStatisticsFromDict(dict)
}

func NewLotteryNamespaceDrawAmountByUserDistributionStatisticsFromDict(data map[string]interface{}) LotteryNamespaceDrawAmountByUserDistributionStatistics {
	return LotteryNamespaceDrawAmountByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p LotteryNamespaceDrawAmountByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p LotteryNamespaceDrawAmountByUserDistributionStatistics) Pointer() *LotteryNamespaceDrawAmountByUserDistributionStatistics {
	return &p
}

func CastLotteryNamespaceDrawAmountByUserDistributionStatisticses(data []interface{}) []LotteryNamespaceDrawAmountByUserDistributionStatistics {
	v := make([]LotteryNamespaceDrawAmountByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawAmountByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawAmountByUserDistributionStatisticsesFromDict(data []LotteryNamespaceDrawAmountByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawAmountByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewLotteryNamespaceDrawAmountByUserDistributionSegmentFromJson(data string) LotteryNamespaceDrawAmountByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawAmountByUserDistributionSegmentFromDict(dict)
}

func NewLotteryNamespaceDrawAmountByUserDistributionSegmentFromDict(data map[string]interface{}) LotteryNamespaceDrawAmountByUserDistributionSegment {
	return LotteryNamespaceDrawAmountByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p LotteryNamespaceDrawAmountByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p LotteryNamespaceDrawAmountByUserDistributionSegment) Pointer() *LotteryNamespaceDrawAmountByUserDistributionSegment {
	return &p
}

func CastLotteryNamespaceDrawAmountByUserDistributionSegments(data []interface{}) []LotteryNamespaceDrawAmountByUserDistributionSegment {
	v := make([]LotteryNamespaceDrawAmountByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawAmountByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawAmountByUserDistributionSegmentsFromDict(data []LotteryNamespaceDrawAmountByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDrawAmountByUserDistribution struct {
	Statistics   *LotteryNamespaceDrawAmountByUserDistributionStatistics `json:"statistics"`
	Distribution []LotteryNamespaceDrawAmountByUserDistributionSegment   `json:"distribution"`
}

func NewLotteryNamespaceDrawAmountByUserDistributionFromJson(data string) LotteryNamespaceDrawAmountByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDrawAmountByUserDistributionFromDict(dict)
}

func NewLotteryNamespaceDrawAmountByUserDistributionFromDict(data map[string]interface{}) LotteryNamespaceDrawAmountByUserDistribution {
	return LotteryNamespaceDrawAmountByUserDistribution{
		Statistics:   NewLotteryNamespaceDrawAmountByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastLotteryNamespaceDrawAmountByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p LotteryNamespaceDrawAmountByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastLotteryNamespaceDrawAmountByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p LotteryNamespaceDrawAmountByUserDistribution) Pointer() *LotteryNamespaceDrawAmountByUserDistribution {
	return &p
}

func CastLotteryNamespaceDrawAmountByUserDistributions(data []interface{}) []LotteryNamespaceDrawAmountByUserDistribution {
	v := make([]LotteryNamespaceDrawAmountByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDrawAmountByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDrawAmountByUserDistributionsFromDict(data []LotteryNamespaceDrawAmountByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespaceDistributions struct {
	Draw             *LotteryNamespaceDrawDistribution             `json:"draw"`
	DrawAmount       *LotteryNamespaceDrawAmountDistribution       `json:"drawAmount"`
	DrawByUser       *LotteryNamespaceDrawByUserDistribution       `json:"drawByUser"`
	DrawAmountByUser *LotteryNamespaceDrawAmountByUserDistribution `json:"drawAmountByUser"`
}

func NewLotteryNamespaceDistributionsFromJson(data string) LotteryNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceDistributionsFromDict(dict)
}

func NewLotteryNamespaceDistributionsFromDict(data map[string]interface{}) LotteryNamespaceDistributions {
	return LotteryNamespaceDistributions{
		Draw:             NewLotteryNamespaceDrawDistributionFromDict(core.CastMap(data["draw"])).Pointer(),
		DrawAmount:       NewLotteryNamespaceDrawAmountDistributionFromDict(core.CastMap(data["drawAmount"])).Pointer(),
		DrawByUser:       NewLotteryNamespaceDrawByUserDistributionFromDict(core.CastMap(data["drawByUser"])).Pointer(),
		DrawAmountByUser: NewLotteryNamespaceDrawAmountByUserDistributionFromDict(core.CastMap(data["drawAmountByUser"])).Pointer(),
	}
}

func (p LotteryNamespaceDistributions) ToDict() map[string]interface{} {

	var draw map[string]interface{}
	if p.Draw != nil {
		draw = p.Draw.ToDict()
	}
	var drawAmount map[string]interface{}
	if p.DrawAmount != nil {
		drawAmount = p.DrawAmount.ToDict()
	}
	var drawByUser map[string]interface{}
	if p.DrawByUser != nil {
		drawByUser = p.DrawByUser.ToDict()
	}
	var drawAmountByUser map[string]interface{}
	if p.DrawAmountByUser != nil {
		drawAmountByUser = p.DrawAmountByUser.ToDict()
	}
	return map[string]interface{}{
		"draw":             draw,
		"drawAmount":       drawAmount,
		"drawByUser":       drawByUser,
		"drawAmountByUser": drawAmountByUser,
	}
}

func (p LotteryNamespaceDistributions) Pointer() *LotteryNamespaceDistributions {
	return &p
}

func CastLotteryNamespaceDistributionses(data []interface{}) []LotteryNamespaceDistributions {
	v := make([]LotteryNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespaceDistributionsesFromDict(data []LotteryNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryNamespace struct {
	NamespaceId   *string                        `json:"namespaceId"`
	Year          *int32                         `json:"year"`
	Month         *int32                         `json:"month"`
	Day           *int32                         `json:"day"`
	NamespaceName *string                        `json:"namespaceName"`
	Statistics    *LotteryNamespaceStatistics    `json:"statistics"`
	Distributions *LotteryNamespaceDistributions `json:"distributions"`
	Lotteries     []LotteryLottery               `json:"lotteries"`
}

func NewLotteryNamespaceFromJson(data string) LotteryNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryNamespaceFromDict(dict)
}

func NewLotteryNamespaceFromDict(data map[string]interface{}) LotteryNamespace {
	return LotteryNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewLotteryNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewLotteryNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		Lotteries:     CastLotteryLotteries(core.CastArray(data["lotteries"])),
	}
}

func (p LotteryNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var lotteries []interface{}
	if p.Lotteries != nil {
		lotteries = CastLotteryLotteriesFromDict(
			p.Lotteries,
		)
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
		"lotteries":     lotteries,
	}
}

func (p LotteryNamespace) Pointer() *LotteryNamespace {
	return &p
}

func CastLotteryNamespaces(data []interface{}) []LotteryNamespace {
	v := make([]LotteryNamespace, 0)
	for _, d := range data {
		v = append(v, NewLotteryNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryNamespacesFromDict(data []LotteryNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryLotteryModel struct {
	LotteryModelId *string `json:"lotteryModelId"`
	LotteryName    *string `json:"lotteryName"`
}

func NewLotteryLotteryModelFromJson(data string) LotteryLotteryModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLotteryLotteryModelFromDict(dict)
}

func NewLotteryLotteryModelFromDict(data map[string]interface{}) LotteryLotteryModel {
	return LotteryLotteryModel{
		LotteryModelId: core.CastString(data["lotteryModelId"]),
		LotteryName:    core.CastString(data["lotteryName"]),
	}
}

func (p LotteryLotteryModel) ToDict() map[string]interface{} {

	var lotteryModelId *string
	if p.LotteryModelId != nil {
		lotteryModelId = p.LotteryModelId
	}
	var lotteryName *string
	if p.LotteryName != nil {
		lotteryName = p.LotteryName
	}
	return map[string]interface{}{
		"lotteryModelId": lotteryModelId,
		"lotteryName":    lotteryName,
	}
}

func (p LotteryLotteryModel) Pointer() *LotteryLotteryModel {
	return &p
}

func CastLotteryLotteryModels(data []interface{}) []LotteryLotteryModel {
	v := make([]LotteryLotteryModel, 0)
	for _, d := range data {
		v = append(v, NewLotteryLotteryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryLotteryModelsFromDict(data []LotteryLotteryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MatchmakingNamespaceStatistics struct {
	CreateGathering *int64   `json:"createGathering"`
	Matchmaking     *int64   `json:"matchmaking"`
	CompleteRate    *float32 `json:"completeRate"`
}

func NewMatchmakingNamespaceStatisticsFromJson(data string) MatchmakingNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchmakingNamespaceStatisticsFromDict(dict)
}

func NewMatchmakingNamespaceStatisticsFromDict(data map[string]interface{}) MatchmakingNamespaceStatistics {
	return MatchmakingNamespaceStatistics{
		CreateGathering: core.CastInt64(data["createGathering"]),
		Matchmaking:     core.CastInt64(data["matchmaking"]),
		CompleteRate:    core.CastFloat32(data["completeRate"]),
	}
}

func (p MatchmakingNamespaceStatistics) ToDict() map[string]interface{} {

	var createGathering *int64
	if p.CreateGathering != nil {
		createGathering = p.CreateGathering
	}
	var matchmaking *int64
	if p.Matchmaking != nil {
		matchmaking = p.Matchmaking
	}
	var completeRate *float32
	if p.CompleteRate != nil {
		completeRate = p.CompleteRate
	}
	return map[string]interface{}{
		"createGathering": createGathering,
		"matchmaking":     matchmaking,
		"completeRate":    completeRate,
	}
}

func (p MatchmakingNamespaceStatistics) Pointer() *MatchmakingNamespaceStatistics {
	return &p
}

func CastMatchmakingNamespaceStatisticses(data []interface{}) []MatchmakingNamespaceStatistics {
	v := make([]MatchmakingNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewMatchmakingNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchmakingNamespaceStatisticsesFromDict(data []MatchmakingNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MatchmakingNamespaceResultDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMatchmakingNamespaceResultDistributionStatisticsFromJson(data string) MatchmakingNamespaceResultDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchmakingNamespaceResultDistributionStatisticsFromDict(dict)
}

func NewMatchmakingNamespaceResultDistributionStatisticsFromDict(data map[string]interface{}) MatchmakingNamespaceResultDistributionStatistics {
	return MatchmakingNamespaceResultDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MatchmakingNamespaceResultDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MatchmakingNamespaceResultDistributionStatistics) Pointer() *MatchmakingNamespaceResultDistributionStatistics {
	return &p
}

func CastMatchmakingNamespaceResultDistributionStatisticses(data []interface{}) []MatchmakingNamespaceResultDistributionStatistics {
	v := make([]MatchmakingNamespaceResultDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMatchmakingNamespaceResultDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchmakingNamespaceResultDistributionStatisticsesFromDict(data []MatchmakingNamespaceResultDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MatchmakingNamespaceResultDistributionSegment struct {
	Result *string `json:"result"`
	Count  *int64  `json:"count"`
}

func NewMatchmakingNamespaceResultDistributionSegmentFromJson(data string) MatchmakingNamespaceResultDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchmakingNamespaceResultDistributionSegmentFromDict(dict)
}

func NewMatchmakingNamespaceResultDistributionSegmentFromDict(data map[string]interface{}) MatchmakingNamespaceResultDistributionSegment {
	return MatchmakingNamespaceResultDistributionSegment{
		Result: core.CastString(data["result"]),
		Count:  core.CastInt64(data["count"]),
	}
}

func (p MatchmakingNamespaceResultDistributionSegment) ToDict() map[string]interface{} {

	var result *string
	if p.Result != nil {
		result = p.Result
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"result": result,
		"count":  count,
	}
}

func (p MatchmakingNamespaceResultDistributionSegment) Pointer() *MatchmakingNamespaceResultDistributionSegment {
	return &p
}

func CastMatchmakingNamespaceResultDistributionSegments(data []interface{}) []MatchmakingNamespaceResultDistributionSegment {
	v := make([]MatchmakingNamespaceResultDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMatchmakingNamespaceResultDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchmakingNamespaceResultDistributionSegmentsFromDict(data []MatchmakingNamespaceResultDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MatchmakingNamespaceResultDistribution struct {
	Statistics   *MatchmakingNamespaceResultDistributionStatistics `json:"statistics"`
	Distribution []MatchmakingNamespaceResultDistributionSegment   `json:"distribution"`
}

func NewMatchmakingNamespaceResultDistributionFromJson(data string) MatchmakingNamespaceResultDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchmakingNamespaceResultDistributionFromDict(dict)
}

func NewMatchmakingNamespaceResultDistributionFromDict(data map[string]interface{}) MatchmakingNamespaceResultDistribution {
	return MatchmakingNamespaceResultDistribution{
		Statistics:   NewMatchmakingNamespaceResultDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMatchmakingNamespaceResultDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MatchmakingNamespaceResultDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMatchmakingNamespaceResultDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MatchmakingNamespaceResultDistribution) Pointer() *MatchmakingNamespaceResultDistribution {
	return &p
}

func CastMatchmakingNamespaceResultDistributions(data []interface{}) []MatchmakingNamespaceResultDistribution {
	v := make([]MatchmakingNamespaceResultDistribution, 0)
	for _, d := range data {
		v = append(v, NewMatchmakingNamespaceResultDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchmakingNamespaceResultDistributionsFromDict(data []MatchmakingNamespaceResultDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMatchmakingNamespaceWaitElapsedSecondsDistributionStatisticsFromJson(data string) MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchmakingNamespaceWaitElapsedSecondsDistributionStatisticsFromDict(dict)
}

func NewMatchmakingNamespaceWaitElapsedSecondsDistributionStatisticsFromDict(data map[string]interface{}) MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics {
	return MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics) Pointer() *MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics {
	return &p
}

func CastMatchmakingNamespaceWaitElapsedSecondsDistributionStatisticses(data []interface{}) []MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics {
	v := make([]MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMatchmakingNamespaceWaitElapsedSecondsDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchmakingNamespaceWaitElapsedSecondsDistributionStatisticsesFromDict(data []MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MatchmakingNamespaceWaitElapsedSecondsDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewMatchmakingNamespaceWaitElapsedSecondsDistributionSegmentFromJson(data string) MatchmakingNamespaceWaitElapsedSecondsDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchmakingNamespaceWaitElapsedSecondsDistributionSegmentFromDict(dict)
}

func NewMatchmakingNamespaceWaitElapsedSecondsDistributionSegmentFromDict(data map[string]interface{}) MatchmakingNamespaceWaitElapsedSecondsDistributionSegment {
	return MatchmakingNamespaceWaitElapsedSecondsDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p MatchmakingNamespaceWaitElapsedSecondsDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p MatchmakingNamespaceWaitElapsedSecondsDistributionSegment) Pointer() *MatchmakingNamespaceWaitElapsedSecondsDistributionSegment {
	return &p
}

func CastMatchmakingNamespaceWaitElapsedSecondsDistributionSegments(data []interface{}) []MatchmakingNamespaceWaitElapsedSecondsDistributionSegment {
	v := make([]MatchmakingNamespaceWaitElapsedSecondsDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMatchmakingNamespaceWaitElapsedSecondsDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchmakingNamespaceWaitElapsedSecondsDistributionSegmentsFromDict(data []MatchmakingNamespaceWaitElapsedSecondsDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MatchmakingNamespaceWaitElapsedSecondsDistribution struct {
	Statistics   *MatchmakingNamespaceWaitElapsedSecondsDistributionStatistics `json:"statistics"`
	Distribution []MatchmakingNamespaceWaitElapsedSecondsDistributionSegment   `json:"distribution"`
}

func NewMatchmakingNamespaceWaitElapsedSecondsDistributionFromJson(data string) MatchmakingNamespaceWaitElapsedSecondsDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchmakingNamespaceWaitElapsedSecondsDistributionFromDict(dict)
}

func NewMatchmakingNamespaceWaitElapsedSecondsDistributionFromDict(data map[string]interface{}) MatchmakingNamespaceWaitElapsedSecondsDistribution {
	return MatchmakingNamespaceWaitElapsedSecondsDistribution{
		Statistics:   NewMatchmakingNamespaceWaitElapsedSecondsDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMatchmakingNamespaceWaitElapsedSecondsDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MatchmakingNamespaceWaitElapsedSecondsDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMatchmakingNamespaceWaitElapsedSecondsDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MatchmakingNamespaceWaitElapsedSecondsDistribution) Pointer() *MatchmakingNamespaceWaitElapsedSecondsDistribution {
	return &p
}

func CastMatchmakingNamespaceWaitElapsedSecondsDistributions(data []interface{}) []MatchmakingNamespaceWaitElapsedSecondsDistribution {
	v := make([]MatchmakingNamespaceWaitElapsedSecondsDistribution, 0)
	for _, d := range data {
		v = append(v, NewMatchmakingNamespaceWaitElapsedSecondsDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchmakingNamespaceWaitElapsedSecondsDistributionsFromDict(data []MatchmakingNamespaceWaitElapsedSecondsDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MatchmakingNamespaceDistributions struct {
	Result             *MatchmakingNamespaceResultDistribution             `json:"result"`
	WaitElapsedSeconds *MatchmakingNamespaceWaitElapsedSecondsDistribution `json:"waitElapsedSeconds"`
}

func NewMatchmakingNamespaceDistributionsFromJson(data string) MatchmakingNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchmakingNamespaceDistributionsFromDict(dict)
}

func NewMatchmakingNamespaceDistributionsFromDict(data map[string]interface{}) MatchmakingNamespaceDistributions {
	return MatchmakingNamespaceDistributions{
		Result:             NewMatchmakingNamespaceResultDistributionFromDict(core.CastMap(data["result"])).Pointer(),
		WaitElapsedSeconds: NewMatchmakingNamespaceWaitElapsedSecondsDistributionFromDict(core.CastMap(data["waitElapsedSeconds"])).Pointer(),
	}
}

func (p MatchmakingNamespaceDistributions) ToDict() map[string]interface{} {

	var result map[string]interface{}
	if p.Result != nil {
		result = p.Result.ToDict()
	}
	var waitElapsedSeconds map[string]interface{}
	if p.WaitElapsedSeconds != nil {
		waitElapsedSeconds = p.WaitElapsedSeconds.ToDict()
	}
	return map[string]interface{}{
		"result":             result,
		"waitElapsedSeconds": waitElapsedSeconds,
	}
}

func (p MatchmakingNamespaceDistributions) Pointer() *MatchmakingNamespaceDistributions {
	return &p
}

func CastMatchmakingNamespaceDistributionses(data []interface{}) []MatchmakingNamespaceDistributions {
	v := make([]MatchmakingNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewMatchmakingNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchmakingNamespaceDistributionsesFromDict(data []MatchmakingNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MatchmakingNamespace struct {
	NamespaceId   *string                            `json:"namespaceId"`
	Year          *int32                             `json:"year"`
	Month         *int32                             `json:"month"`
	Day           *int32                             `json:"day"`
	NamespaceName *string                            `json:"namespaceName"`
	Statistics    *MatchmakingNamespaceStatistics    `json:"statistics"`
	Distributions *MatchmakingNamespaceDistributions `json:"distributions"`
}

func NewMatchmakingNamespaceFromJson(data string) MatchmakingNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchmakingNamespaceFromDict(dict)
}

func NewMatchmakingNamespaceFromDict(data map[string]interface{}) MatchmakingNamespace {
	return MatchmakingNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewMatchmakingNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewMatchmakingNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p MatchmakingNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p MatchmakingNamespace) Pointer() *MatchmakingNamespace {
	return &p
}

func CastMatchmakingNamespaces(data []interface{}) []MatchmakingNamespace {
	v := make([]MatchmakingNamespace, 0)
	for _, d := range data {
		v = append(v, NewMatchmakingNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchmakingNamespacesFromDict(data []MatchmakingNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionCounterCounterDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMissionCounterCounterDistributionStatisticsFromJson(data string) MissionCounterCounterDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionCounterCounterDistributionStatisticsFromDict(dict)
}

func NewMissionCounterCounterDistributionStatisticsFromDict(data map[string]interface{}) MissionCounterCounterDistributionStatistics {
	return MissionCounterCounterDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MissionCounterCounterDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MissionCounterCounterDistributionStatistics) Pointer() *MissionCounterCounterDistributionStatistics {
	return &p
}

func CastMissionCounterCounterDistributionStatisticses(data []interface{}) []MissionCounterCounterDistributionStatistics {
	v := make([]MissionCounterCounterDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMissionCounterCounterDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionCounterCounterDistributionStatisticsesFromDict(data []MissionCounterCounterDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionCounterCounterDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewMissionCounterCounterDistributionSegmentFromJson(data string) MissionCounterCounterDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionCounterCounterDistributionSegmentFromDict(dict)
}

func NewMissionCounterCounterDistributionSegmentFromDict(data map[string]interface{}) MissionCounterCounterDistributionSegment {
	return MissionCounterCounterDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p MissionCounterCounterDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p MissionCounterCounterDistributionSegment) Pointer() *MissionCounterCounterDistributionSegment {
	return &p
}

func CastMissionCounterCounterDistributionSegments(data []interface{}) []MissionCounterCounterDistributionSegment {
	v := make([]MissionCounterCounterDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMissionCounterCounterDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionCounterCounterDistributionSegmentsFromDict(data []MissionCounterCounterDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionCounterCounterDistribution struct {
	Statistics   *MissionCounterCounterDistributionStatistics `json:"statistics"`
	Distribution []MissionCounterCounterDistributionSegment   `json:"distribution"`
}

func NewMissionCounterCounterDistributionFromJson(data string) MissionCounterCounterDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionCounterCounterDistributionFromDict(dict)
}

func NewMissionCounterCounterDistributionFromDict(data map[string]interface{}) MissionCounterCounterDistribution {
	return MissionCounterCounterDistribution{
		Statistics:   NewMissionCounterCounterDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMissionCounterCounterDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MissionCounterCounterDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMissionCounterCounterDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MissionCounterCounterDistribution) Pointer() *MissionCounterCounterDistribution {
	return &p
}

func CastMissionCounterCounterDistributions(data []interface{}) []MissionCounterCounterDistribution {
	v := make([]MissionCounterCounterDistribution, 0)
	for _, d := range data {
		v = append(v, NewMissionCounterCounterDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionCounterCounterDistributionsFromDict(data []MissionCounterCounterDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionCounterDistributions struct {
	Counter *MissionCounterCounterDistribution `json:"counter"`
}

func NewMissionCounterDistributionsFromJson(data string) MissionCounterDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionCounterDistributionsFromDict(dict)
}

func NewMissionCounterDistributionsFromDict(data map[string]interface{}) MissionCounterDistributions {
	return MissionCounterDistributions{
		Counter: NewMissionCounterCounterDistributionFromDict(core.CastMap(data["counter"])).Pointer(),
	}
}

func (p MissionCounterDistributions) ToDict() map[string]interface{} {

	var counter map[string]interface{}
	if p.Counter != nil {
		counter = p.Counter.ToDict()
	}
	return map[string]interface{}{
		"counter": counter,
	}
}

func (p MissionCounterDistributions) Pointer() *MissionCounterDistributions {
	return &p
}

func CastMissionCounterDistributionses(data []interface{}) []MissionCounterDistributions {
	v := make([]MissionCounterDistributions, 0)
	for _, d := range data {
		v = append(v, NewMissionCounterDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionCounterDistributionsesFromDict(data []MissionCounterDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionCounter struct {
	CounterId     *string                      `json:"counterId"`
	CounterName   *string                      `json:"counterName"`
	Distributions *MissionCounterDistributions `json:"distributions"`
}

func NewMissionCounterFromJson(data string) MissionCounter {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionCounterFromDict(dict)
}

func NewMissionCounterFromDict(data map[string]interface{}) MissionCounter {
	return MissionCounter{
		CounterId:     core.CastString(data["counterId"]),
		CounterName:   core.CastString(data["counterName"]),
		Distributions: NewMissionCounterDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p MissionCounter) ToDict() map[string]interface{} {

	var counterId *string
	if p.CounterId != nil {
		counterId = p.CounterId
	}
	var counterName *string
	if p.CounterName != nil {
		counterName = p.CounterName
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"counterId":     counterId,
		"counterName":   counterName,
		"distributions": distributions,
	}
}

func (p MissionCounter) Pointer() *MissionCounter {
	return &p
}

func CastMissionCounters(data []interface{}) []MissionCounter {
	v := make([]MissionCounter, 0)
	for _, d := range data {
		v = append(v, NewMissionCounterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionCountersFromDict(data []MissionCounter) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionMissionGroupModelStatistics struct {
	Receive *int64 `json:"receive"`
}

func NewMissionMissionGroupModelStatisticsFromJson(data string) MissionMissionGroupModelStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionMissionGroupModelStatisticsFromDict(dict)
}

func NewMissionMissionGroupModelStatisticsFromDict(data map[string]interface{}) MissionMissionGroupModelStatistics {
	return MissionMissionGroupModelStatistics{
		Receive: core.CastInt64(data["receive"]),
	}
}

func (p MissionMissionGroupModelStatistics) ToDict() map[string]interface{} {

	var receive *int64
	if p.Receive != nil {
		receive = p.Receive
	}
	return map[string]interface{}{
		"receive": receive,
	}
}

func (p MissionMissionGroupModelStatistics) Pointer() *MissionMissionGroupModelStatistics {
	return &p
}

func CastMissionMissionGroupModelStatisticses(data []interface{}) []MissionMissionGroupModelStatistics {
	v := make([]MissionMissionGroupModelStatistics, 0)
	for _, d := range data {
		v = append(v, NewMissionMissionGroupModelStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionMissionGroupModelStatisticsesFromDict(data []MissionMissionGroupModelStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionMissionGroupModelReceiveDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMissionMissionGroupModelReceiveDistributionStatisticsFromJson(data string) MissionMissionGroupModelReceiveDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionMissionGroupModelReceiveDistributionStatisticsFromDict(dict)
}

func NewMissionMissionGroupModelReceiveDistributionStatisticsFromDict(data map[string]interface{}) MissionMissionGroupModelReceiveDistributionStatistics {
	return MissionMissionGroupModelReceiveDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MissionMissionGroupModelReceiveDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MissionMissionGroupModelReceiveDistributionStatistics) Pointer() *MissionMissionGroupModelReceiveDistributionStatistics {
	return &p
}

func CastMissionMissionGroupModelReceiveDistributionStatisticses(data []interface{}) []MissionMissionGroupModelReceiveDistributionStatistics {
	v := make([]MissionMissionGroupModelReceiveDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMissionMissionGroupModelReceiveDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionMissionGroupModelReceiveDistributionStatisticsesFromDict(data []MissionMissionGroupModelReceiveDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionMissionGroupModelReceiveDistributionSegment struct {
	MissionTaskName *string `json:"missionTaskName"`
	Count           *int64  `json:"count"`
}

func NewMissionMissionGroupModelReceiveDistributionSegmentFromJson(data string) MissionMissionGroupModelReceiveDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionMissionGroupModelReceiveDistributionSegmentFromDict(dict)
}

func NewMissionMissionGroupModelReceiveDistributionSegmentFromDict(data map[string]interface{}) MissionMissionGroupModelReceiveDistributionSegment {
	return MissionMissionGroupModelReceiveDistributionSegment{
		MissionTaskName: core.CastString(data["missionTaskName"]),
		Count:           core.CastInt64(data["count"]),
	}
}

func (p MissionMissionGroupModelReceiveDistributionSegment) ToDict() map[string]interface{} {

	var missionTaskName *string
	if p.MissionTaskName != nil {
		missionTaskName = p.MissionTaskName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"missionTaskName": missionTaskName,
		"count":           count,
	}
}

func (p MissionMissionGroupModelReceiveDistributionSegment) Pointer() *MissionMissionGroupModelReceiveDistributionSegment {
	return &p
}

func CastMissionMissionGroupModelReceiveDistributionSegments(data []interface{}) []MissionMissionGroupModelReceiveDistributionSegment {
	v := make([]MissionMissionGroupModelReceiveDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMissionMissionGroupModelReceiveDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionMissionGroupModelReceiveDistributionSegmentsFromDict(data []MissionMissionGroupModelReceiveDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionMissionGroupModelReceiveDistribution struct {
	Statistics   *MissionMissionGroupModelReceiveDistributionStatistics `json:"statistics"`
	Distribution []MissionMissionGroupModelReceiveDistributionSegment   `json:"distribution"`
}

func NewMissionMissionGroupModelReceiveDistributionFromJson(data string) MissionMissionGroupModelReceiveDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionMissionGroupModelReceiveDistributionFromDict(dict)
}

func NewMissionMissionGroupModelReceiveDistributionFromDict(data map[string]interface{}) MissionMissionGroupModelReceiveDistribution {
	return MissionMissionGroupModelReceiveDistribution{
		Statistics:   NewMissionMissionGroupModelReceiveDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMissionMissionGroupModelReceiveDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MissionMissionGroupModelReceiveDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMissionMissionGroupModelReceiveDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MissionMissionGroupModelReceiveDistribution) Pointer() *MissionMissionGroupModelReceiveDistribution {
	return &p
}

func CastMissionMissionGroupModelReceiveDistributions(data []interface{}) []MissionMissionGroupModelReceiveDistribution {
	v := make([]MissionMissionGroupModelReceiveDistribution, 0)
	for _, d := range data {
		v = append(v, NewMissionMissionGroupModelReceiveDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionMissionGroupModelReceiveDistributionsFromDict(data []MissionMissionGroupModelReceiveDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionMissionGroupModelDistributions struct {
	Receive *MissionMissionGroupModelReceiveDistribution `json:"receive"`
}

func NewMissionMissionGroupModelDistributionsFromJson(data string) MissionMissionGroupModelDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionMissionGroupModelDistributionsFromDict(dict)
}

func NewMissionMissionGroupModelDistributionsFromDict(data map[string]interface{}) MissionMissionGroupModelDistributions {
	return MissionMissionGroupModelDistributions{
		Receive: NewMissionMissionGroupModelReceiveDistributionFromDict(core.CastMap(data["receive"])).Pointer(),
	}
}

func (p MissionMissionGroupModelDistributions) ToDict() map[string]interface{} {

	var receive map[string]interface{}
	if p.Receive != nil {
		receive = p.Receive.ToDict()
	}
	return map[string]interface{}{
		"receive": receive,
	}
}

func (p MissionMissionGroupModelDistributions) Pointer() *MissionMissionGroupModelDistributions {
	return &p
}

func CastMissionMissionGroupModelDistributionses(data []interface{}) []MissionMissionGroupModelDistributions {
	v := make([]MissionMissionGroupModelDistributions, 0)
	for _, d := range data {
		v = append(v, NewMissionMissionGroupModelDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionMissionGroupModelDistributionsesFromDict(data []MissionMissionGroupModelDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionMissionGroupModel struct {
	MissionGroupModelId *string                                `json:"missionGroupModelId"`
	MissionGroupName    *string                                `json:"missionGroupName"`
	Statistics          *MissionMissionGroupModelStatistics    `json:"statistics"`
	Distributions       *MissionMissionGroupModelDistributions `json:"distributions"`
}

func NewMissionMissionGroupModelFromJson(data string) MissionMissionGroupModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionMissionGroupModelFromDict(dict)
}

func NewMissionMissionGroupModelFromDict(data map[string]interface{}) MissionMissionGroupModel {
	return MissionMissionGroupModel{
		MissionGroupModelId: core.CastString(data["missionGroupModelId"]),
		MissionGroupName:    core.CastString(data["missionGroupName"]),
		Statistics:          NewMissionMissionGroupModelStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions:       NewMissionMissionGroupModelDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p MissionMissionGroupModel) ToDict() map[string]interface{} {

	var missionGroupModelId *string
	if p.MissionGroupModelId != nil {
		missionGroupModelId = p.MissionGroupModelId
	}
	var missionGroupName *string
	if p.MissionGroupName != nil {
		missionGroupName = p.MissionGroupName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"missionGroupModelId": missionGroupModelId,
		"missionGroupName":    missionGroupName,
		"statistics":          statistics,
		"distributions":       distributions,
	}
}

func (p MissionMissionGroupModel) Pointer() *MissionMissionGroupModel {
	return &p
}

func CastMissionMissionGroupModels(data []interface{}) []MissionMissionGroupModel {
	v := make([]MissionMissionGroupModel, 0)
	for _, d := range data {
		v = append(v, NewMissionMissionGroupModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionMissionGroupModelsFromDict(data []MissionMissionGroupModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionMissionTaskModel struct {
	MissionTaskModelId *string `json:"missionTaskModelId"`
	MissionTaskName    *string `json:"missionTaskName"`
}

func NewMissionMissionTaskModelFromJson(data string) MissionMissionTaskModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionMissionTaskModelFromDict(dict)
}

func NewMissionMissionTaskModelFromDict(data map[string]interface{}) MissionMissionTaskModel {
	return MissionMissionTaskModel{
		MissionTaskModelId: core.CastString(data["missionTaskModelId"]),
		MissionTaskName:    core.CastString(data["missionTaskName"]),
	}
}

func (p MissionMissionTaskModel) ToDict() map[string]interface{} {

	var missionTaskModelId *string
	if p.MissionTaskModelId != nil {
		missionTaskModelId = p.MissionTaskModelId
	}
	var missionTaskName *string
	if p.MissionTaskName != nil {
		missionTaskName = p.MissionTaskName
	}
	return map[string]interface{}{
		"missionTaskModelId": missionTaskModelId,
		"missionTaskName":    missionTaskName,
	}
}

func (p MissionMissionTaskModel) Pointer() *MissionMissionTaskModel {
	return &p
}

func CastMissionMissionTaskModels(data []interface{}) []MissionMissionTaskModel {
	v := make([]MissionMissionTaskModel, 0)
	for _, d := range data {
		v = append(v, NewMissionMissionTaskModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionMissionTaskModelsFromDict(data []MissionMissionTaskModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceStatistics struct {
	Increase       *int64 `json:"increase"`
	IncreaseAmount *int64 `json:"increaseAmount"`
	Receive        *int64 `json:"receive"`
}

func NewMissionNamespaceStatisticsFromJson(data string) MissionNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceStatisticsFromDict(dict)
}

func NewMissionNamespaceStatisticsFromDict(data map[string]interface{}) MissionNamespaceStatistics {
	return MissionNamespaceStatistics{
		Increase:       core.CastInt64(data["increase"]),
		IncreaseAmount: core.CastInt64(data["increaseAmount"]),
		Receive:        core.CastInt64(data["receive"]),
	}
}

func (p MissionNamespaceStatistics) ToDict() map[string]interface{} {

	var increase *int64
	if p.Increase != nil {
		increase = p.Increase
	}
	var increaseAmount *int64
	if p.IncreaseAmount != nil {
		increaseAmount = p.IncreaseAmount
	}
	var receive *int64
	if p.Receive != nil {
		receive = p.Receive
	}
	return map[string]interface{}{
		"increase":       increase,
		"increaseAmount": increaseAmount,
		"receive":        receive,
	}
}

func (p MissionNamespaceStatistics) Pointer() *MissionNamespaceStatistics {
	return &p
}

func CastMissionNamespaceStatisticses(data []interface{}) []MissionNamespaceStatistics {
	v := make([]MissionNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceStatisticsesFromDict(data []MissionNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceIncreaseDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMissionNamespaceIncreaseDistributionStatisticsFromJson(data string) MissionNamespaceIncreaseDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceIncreaseDistributionStatisticsFromDict(dict)
}

func NewMissionNamespaceIncreaseDistributionStatisticsFromDict(data map[string]interface{}) MissionNamespaceIncreaseDistributionStatistics {
	return MissionNamespaceIncreaseDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MissionNamespaceIncreaseDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MissionNamespaceIncreaseDistributionStatistics) Pointer() *MissionNamespaceIncreaseDistributionStatistics {
	return &p
}

func CastMissionNamespaceIncreaseDistributionStatisticses(data []interface{}) []MissionNamespaceIncreaseDistributionStatistics {
	v := make([]MissionNamespaceIncreaseDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceIncreaseDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceIncreaseDistributionStatisticsesFromDict(data []MissionNamespaceIncreaseDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceIncreaseDistributionSegment struct {
	CounterName *string `json:"counterName"`
	Count       *int64  `json:"count"`
}

func NewMissionNamespaceIncreaseDistributionSegmentFromJson(data string) MissionNamespaceIncreaseDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceIncreaseDistributionSegmentFromDict(dict)
}

func NewMissionNamespaceIncreaseDistributionSegmentFromDict(data map[string]interface{}) MissionNamespaceIncreaseDistributionSegment {
	return MissionNamespaceIncreaseDistributionSegment{
		CounterName: core.CastString(data["counterName"]),
		Count:       core.CastInt64(data["count"]),
	}
}

func (p MissionNamespaceIncreaseDistributionSegment) ToDict() map[string]interface{} {

	var counterName *string
	if p.CounterName != nil {
		counterName = p.CounterName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"counterName": counterName,
		"count":       count,
	}
}

func (p MissionNamespaceIncreaseDistributionSegment) Pointer() *MissionNamespaceIncreaseDistributionSegment {
	return &p
}

func CastMissionNamespaceIncreaseDistributionSegments(data []interface{}) []MissionNamespaceIncreaseDistributionSegment {
	v := make([]MissionNamespaceIncreaseDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceIncreaseDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceIncreaseDistributionSegmentsFromDict(data []MissionNamespaceIncreaseDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceIncreaseDistribution struct {
	Statistics   *MissionNamespaceIncreaseDistributionStatistics `json:"statistics"`
	Distribution []MissionNamespaceIncreaseDistributionSegment   `json:"distribution"`
}

func NewMissionNamespaceIncreaseDistributionFromJson(data string) MissionNamespaceIncreaseDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceIncreaseDistributionFromDict(dict)
}

func NewMissionNamespaceIncreaseDistributionFromDict(data map[string]interface{}) MissionNamespaceIncreaseDistribution {
	return MissionNamespaceIncreaseDistribution{
		Statistics:   NewMissionNamespaceIncreaseDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMissionNamespaceIncreaseDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MissionNamespaceIncreaseDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMissionNamespaceIncreaseDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MissionNamespaceIncreaseDistribution) Pointer() *MissionNamespaceIncreaseDistribution {
	return &p
}

func CastMissionNamespaceIncreaseDistributions(data []interface{}) []MissionNamespaceIncreaseDistribution {
	v := make([]MissionNamespaceIncreaseDistribution, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceIncreaseDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceIncreaseDistributionsFromDict(data []MissionNamespaceIncreaseDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceIncreaseAmountDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMissionNamespaceIncreaseAmountDistributionStatisticsFromJson(data string) MissionNamespaceIncreaseAmountDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceIncreaseAmountDistributionStatisticsFromDict(dict)
}

func NewMissionNamespaceIncreaseAmountDistributionStatisticsFromDict(data map[string]interface{}) MissionNamespaceIncreaseAmountDistributionStatistics {
	return MissionNamespaceIncreaseAmountDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MissionNamespaceIncreaseAmountDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MissionNamespaceIncreaseAmountDistributionStatistics) Pointer() *MissionNamespaceIncreaseAmountDistributionStatistics {
	return &p
}

func CastMissionNamespaceIncreaseAmountDistributionStatisticses(data []interface{}) []MissionNamespaceIncreaseAmountDistributionStatistics {
	v := make([]MissionNamespaceIncreaseAmountDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceIncreaseAmountDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceIncreaseAmountDistributionStatisticsesFromDict(data []MissionNamespaceIncreaseAmountDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceIncreaseAmountDistributionSegment struct {
	CounterName *string `json:"counterName"`
	Sum         *int64  `json:"sum"`
}

func NewMissionNamespaceIncreaseAmountDistributionSegmentFromJson(data string) MissionNamespaceIncreaseAmountDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceIncreaseAmountDistributionSegmentFromDict(dict)
}

func NewMissionNamespaceIncreaseAmountDistributionSegmentFromDict(data map[string]interface{}) MissionNamespaceIncreaseAmountDistributionSegment {
	return MissionNamespaceIncreaseAmountDistributionSegment{
		CounterName: core.CastString(data["counterName"]),
		Sum:         core.CastInt64(data["sum"]),
	}
}

func (p MissionNamespaceIncreaseAmountDistributionSegment) ToDict() map[string]interface{} {

	var counterName *string
	if p.CounterName != nil {
		counterName = p.CounterName
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"counterName": counterName,
		"sum":         sum,
	}
}

func (p MissionNamespaceIncreaseAmountDistributionSegment) Pointer() *MissionNamespaceIncreaseAmountDistributionSegment {
	return &p
}

func CastMissionNamespaceIncreaseAmountDistributionSegments(data []interface{}) []MissionNamespaceIncreaseAmountDistributionSegment {
	v := make([]MissionNamespaceIncreaseAmountDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceIncreaseAmountDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceIncreaseAmountDistributionSegmentsFromDict(data []MissionNamespaceIncreaseAmountDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceIncreaseAmountDistribution struct {
	Statistics   *MissionNamespaceIncreaseAmountDistributionStatistics `json:"statistics"`
	Distribution []MissionNamespaceIncreaseAmountDistributionSegment   `json:"distribution"`
}

func NewMissionNamespaceIncreaseAmountDistributionFromJson(data string) MissionNamespaceIncreaseAmountDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceIncreaseAmountDistributionFromDict(dict)
}

func NewMissionNamespaceIncreaseAmountDistributionFromDict(data map[string]interface{}) MissionNamespaceIncreaseAmountDistribution {
	return MissionNamespaceIncreaseAmountDistribution{
		Statistics:   NewMissionNamespaceIncreaseAmountDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMissionNamespaceIncreaseAmountDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MissionNamespaceIncreaseAmountDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMissionNamespaceIncreaseAmountDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MissionNamespaceIncreaseAmountDistribution) Pointer() *MissionNamespaceIncreaseAmountDistribution {
	return &p
}

func CastMissionNamespaceIncreaseAmountDistributions(data []interface{}) []MissionNamespaceIncreaseAmountDistribution {
	v := make([]MissionNamespaceIncreaseAmountDistribution, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceIncreaseAmountDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceIncreaseAmountDistributionsFromDict(data []MissionNamespaceIncreaseAmountDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceReceiveDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMissionNamespaceReceiveDistributionStatisticsFromJson(data string) MissionNamespaceReceiveDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceReceiveDistributionStatisticsFromDict(dict)
}

func NewMissionNamespaceReceiveDistributionStatisticsFromDict(data map[string]interface{}) MissionNamespaceReceiveDistributionStatistics {
	return MissionNamespaceReceiveDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MissionNamespaceReceiveDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MissionNamespaceReceiveDistributionStatistics) Pointer() *MissionNamespaceReceiveDistributionStatistics {
	return &p
}

func CastMissionNamespaceReceiveDistributionStatisticses(data []interface{}) []MissionNamespaceReceiveDistributionStatistics {
	v := make([]MissionNamespaceReceiveDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceReceiveDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceReceiveDistributionStatisticsesFromDict(data []MissionNamespaceReceiveDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceReceiveDistributionSegment struct {
	MissionGroupName *string `json:"missionGroupName"`
	Count            *int64  `json:"count"`
}

func NewMissionNamespaceReceiveDistributionSegmentFromJson(data string) MissionNamespaceReceiveDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceReceiveDistributionSegmentFromDict(dict)
}

func NewMissionNamespaceReceiveDistributionSegmentFromDict(data map[string]interface{}) MissionNamespaceReceiveDistributionSegment {
	return MissionNamespaceReceiveDistributionSegment{
		MissionGroupName: core.CastString(data["missionGroupName"]),
		Count:            core.CastInt64(data["count"]),
	}
}

func (p MissionNamespaceReceiveDistributionSegment) ToDict() map[string]interface{} {

	var missionGroupName *string
	if p.MissionGroupName != nil {
		missionGroupName = p.MissionGroupName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"missionGroupName": missionGroupName,
		"count":            count,
	}
}

func (p MissionNamespaceReceiveDistributionSegment) Pointer() *MissionNamespaceReceiveDistributionSegment {
	return &p
}

func CastMissionNamespaceReceiveDistributionSegments(data []interface{}) []MissionNamespaceReceiveDistributionSegment {
	v := make([]MissionNamespaceReceiveDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceReceiveDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceReceiveDistributionSegmentsFromDict(data []MissionNamespaceReceiveDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceReceiveDistribution struct {
	Statistics   *MissionNamespaceReceiveDistributionStatistics `json:"statistics"`
	Distribution []MissionNamespaceReceiveDistributionSegment   `json:"distribution"`
}

func NewMissionNamespaceReceiveDistributionFromJson(data string) MissionNamespaceReceiveDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceReceiveDistributionFromDict(dict)
}

func NewMissionNamespaceReceiveDistributionFromDict(data map[string]interface{}) MissionNamespaceReceiveDistribution {
	return MissionNamespaceReceiveDistribution{
		Statistics:   NewMissionNamespaceReceiveDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMissionNamespaceReceiveDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MissionNamespaceReceiveDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMissionNamespaceReceiveDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MissionNamespaceReceiveDistribution) Pointer() *MissionNamespaceReceiveDistribution {
	return &p
}

func CastMissionNamespaceReceiveDistributions(data []interface{}) []MissionNamespaceReceiveDistribution {
	v := make([]MissionNamespaceReceiveDistribution, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceReceiveDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceReceiveDistributionsFromDict(data []MissionNamespaceReceiveDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespaceDistributions struct {
	Increase       *MissionNamespaceIncreaseDistribution       `json:"increase"`
	IncreaseAmount *MissionNamespaceIncreaseAmountDistribution `json:"increaseAmount"`
	Receive        *MissionNamespaceReceiveDistribution        `json:"receive"`
}

func NewMissionNamespaceDistributionsFromJson(data string) MissionNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceDistributionsFromDict(dict)
}

func NewMissionNamespaceDistributionsFromDict(data map[string]interface{}) MissionNamespaceDistributions {
	return MissionNamespaceDistributions{
		Increase:       NewMissionNamespaceIncreaseDistributionFromDict(core.CastMap(data["increase"])).Pointer(),
		IncreaseAmount: NewMissionNamespaceIncreaseAmountDistributionFromDict(core.CastMap(data["increaseAmount"])).Pointer(),
		Receive:        NewMissionNamespaceReceiveDistributionFromDict(core.CastMap(data["receive"])).Pointer(),
	}
}

func (p MissionNamespaceDistributions) ToDict() map[string]interface{} {

	var increase map[string]interface{}
	if p.Increase != nil {
		increase = p.Increase.ToDict()
	}
	var increaseAmount map[string]interface{}
	if p.IncreaseAmount != nil {
		increaseAmount = p.IncreaseAmount.ToDict()
	}
	var receive map[string]interface{}
	if p.Receive != nil {
		receive = p.Receive.ToDict()
	}
	return map[string]interface{}{
		"increase":       increase,
		"increaseAmount": increaseAmount,
		"receive":        receive,
	}
}

func (p MissionNamespaceDistributions) Pointer() *MissionNamespaceDistributions {
	return &p
}

func CastMissionNamespaceDistributionses(data []interface{}) []MissionNamespaceDistributions {
	v := make([]MissionNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespaceDistributionsesFromDict(data []MissionNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionNamespace struct {
	NamespaceId        *string                        `json:"namespaceId"`
	Year               *int32                         `json:"year"`
	Month              *int32                         `json:"month"`
	Day                *int32                         `json:"day"`
	NamespaceName      *string                        `json:"namespaceName"`
	Statistics         *MissionNamespaceStatistics    `json:"statistics"`
	Distributions      *MissionNamespaceDistributions `json:"distributions"`
	MissionGroupModels []MissionMissionGroupModel     `json:"missionGroupModels"`
	Counters           []MissionCounter               `json:"counters"`
}

func NewMissionNamespaceFromJson(data string) MissionNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissionNamespaceFromDict(dict)
}

func NewMissionNamespaceFromDict(data map[string]interface{}) MissionNamespace {
	return MissionNamespace{
		NamespaceId:        core.CastString(data["namespaceId"]),
		Year:               core.CastInt32(data["year"]),
		Month:              core.CastInt32(data["month"]),
		Day:                core.CastInt32(data["day"]),
		NamespaceName:      core.CastString(data["namespaceName"]),
		Statistics:         NewMissionNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions:      NewMissionNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		MissionGroupModels: CastMissionMissionGroupModels(core.CastArray(data["missionGroupModels"])),
		Counters:           CastMissionCounters(core.CastArray(data["counters"])),
	}
}

func (p MissionNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var missionGroupModels []interface{}
	if p.MissionGroupModels != nil {
		missionGroupModels = CastMissionMissionGroupModelsFromDict(
			p.MissionGroupModels,
		)
	}
	var counters []interface{}
	if p.Counters != nil {
		counters = CastMissionCountersFromDict(
			p.Counters,
		)
	}
	return map[string]interface{}{
		"namespaceId":        namespaceId,
		"year":               year,
		"month":              month,
		"day":                day,
		"namespaceName":      namespaceName,
		"statistics":         statistics,
		"distributions":      distributions,
		"missionGroupModels": missionGroupModels,
		"counters":           counters,
	}
}

func (p MissionNamespace) Pointer() *MissionNamespace {
	return &p
}

func CastMissionNamespaces(data []interface{}) []MissionNamespace {
	v := make([]MissionNamespace, 0)
	for _, d := range data {
		v = append(v, NewMissionNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionNamespacesFromDict(data []MissionNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyWalletStatistics struct {
	Deposit        *int64 `json:"deposit"`
	DepositAmount  *int64 `json:"depositAmount"`
	Withdraw       *int64 `json:"withdraw"`
	WithdrawAmount *int64 `json:"withdrawAmount"`
	Revenue        *int64 `json:"revenue"`
}

func NewMoneyWalletStatisticsFromJson(data string) MoneyWalletStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyWalletStatisticsFromDict(dict)
}

func NewMoneyWalletStatisticsFromDict(data map[string]interface{}) MoneyWalletStatistics {
	return MoneyWalletStatistics{
		Deposit:        core.CastInt64(data["deposit"]),
		DepositAmount:  core.CastInt64(data["depositAmount"]),
		Withdraw:       core.CastInt64(data["withdraw"]),
		WithdrawAmount: core.CastInt64(data["withdrawAmount"]),
		Revenue:        core.CastInt64(data["revenue"]),
	}
}

func (p MoneyWalletStatistics) ToDict() map[string]interface{} {

	var deposit *int64
	if p.Deposit != nil {
		deposit = p.Deposit
	}
	var depositAmount *int64
	if p.DepositAmount != nil {
		depositAmount = p.DepositAmount
	}
	var withdraw *int64
	if p.Withdraw != nil {
		withdraw = p.Withdraw
	}
	var withdrawAmount *int64
	if p.WithdrawAmount != nil {
		withdrawAmount = p.WithdrawAmount
	}
	var revenue *int64
	if p.Revenue != nil {
		revenue = p.Revenue
	}
	return map[string]interface{}{
		"deposit":        deposit,
		"depositAmount":  depositAmount,
		"withdraw":       withdraw,
		"withdrawAmount": withdrawAmount,
		"revenue":        revenue,
	}
}

func (p MoneyWalletStatistics) Pointer() *MoneyWalletStatistics {
	return &p
}

func CastMoneyWalletStatisticses(data []interface{}) []MoneyWalletStatistics {
	v := make([]MoneyWalletStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyWalletStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyWalletStatisticsesFromDict(data []MoneyWalletStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyWalletFreeDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMoneyWalletFreeDistributionStatisticsFromJson(data string) MoneyWalletFreeDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyWalletFreeDistributionStatisticsFromDict(dict)
}

func NewMoneyWalletFreeDistributionStatisticsFromDict(data map[string]interface{}) MoneyWalletFreeDistributionStatistics {
	return MoneyWalletFreeDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MoneyWalletFreeDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MoneyWalletFreeDistributionStatistics) Pointer() *MoneyWalletFreeDistributionStatistics {
	return &p
}

func CastMoneyWalletFreeDistributionStatisticses(data []interface{}) []MoneyWalletFreeDistributionStatistics {
	v := make([]MoneyWalletFreeDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyWalletFreeDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyWalletFreeDistributionStatisticsesFromDict(data []MoneyWalletFreeDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyWalletFreeDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewMoneyWalletFreeDistributionSegmentFromJson(data string) MoneyWalletFreeDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyWalletFreeDistributionSegmentFromDict(dict)
}

func NewMoneyWalletFreeDistributionSegmentFromDict(data map[string]interface{}) MoneyWalletFreeDistributionSegment {
	return MoneyWalletFreeDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p MoneyWalletFreeDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p MoneyWalletFreeDistributionSegment) Pointer() *MoneyWalletFreeDistributionSegment {
	return &p
}

func CastMoneyWalletFreeDistributionSegments(data []interface{}) []MoneyWalletFreeDistributionSegment {
	v := make([]MoneyWalletFreeDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMoneyWalletFreeDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyWalletFreeDistributionSegmentsFromDict(data []MoneyWalletFreeDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyWalletFreeDistribution struct {
	Statistics   *MoneyWalletFreeDistributionStatistics `json:"statistics"`
	Distribution []MoneyWalletFreeDistributionSegment   `json:"distribution"`
}

func NewMoneyWalletFreeDistributionFromJson(data string) MoneyWalletFreeDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyWalletFreeDistributionFromDict(dict)
}

func NewMoneyWalletFreeDistributionFromDict(data map[string]interface{}) MoneyWalletFreeDistribution {
	return MoneyWalletFreeDistribution{
		Statistics:   NewMoneyWalletFreeDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMoneyWalletFreeDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MoneyWalletFreeDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMoneyWalletFreeDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MoneyWalletFreeDistribution) Pointer() *MoneyWalletFreeDistribution {
	return &p
}

func CastMoneyWalletFreeDistributions(data []interface{}) []MoneyWalletFreeDistribution {
	v := make([]MoneyWalletFreeDistribution, 0)
	for _, d := range data {
		v = append(v, NewMoneyWalletFreeDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyWalletFreeDistributionsFromDict(data []MoneyWalletFreeDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyWalletPaidDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMoneyWalletPaidDistributionStatisticsFromJson(data string) MoneyWalletPaidDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyWalletPaidDistributionStatisticsFromDict(dict)
}

func NewMoneyWalletPaidDistributionStatisticsFromDict(data map[string]interface{}) MoneyWalletPaidDistributionStatistics {
	return MoneyWalletPaidDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MoneyWalletPaidDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MoneyWalletPaidDistributionStatistics) Pointer() *MoneyWalletPaidDistributionStatistics {
	return &p
}

func CastMoneyWalletPaidDistributionStatisticses(data []interface{}) []MoneyWalletPaidDistributionStatistics {
	v := make([]MoneyWalletPaidDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyWalletPaidDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyWalletPaidDistributionStatisticsesFromDict(data []MoneyWalletPaidDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyWalletPaidDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewMoneyWalletPaidDistributionSegmentFromJson(data string) MoneyWalletPaidDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyWalletPaidDistributionSegmentFromDict(dict)
}

func NewMoneyWalletPaidDistributionSegmentFromDict(data map[string]interface{}) MoneyWalletPaidDistributionSegment {
	return MoneyWalletPaidDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p MoneyWalletPaidDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p MoneyWalletPaidDistributionSegment) Pointer() *MoneyWalletPaidDistributionSegment {
	return &p
}

func CastMoneyWalletPaidDistributionSegments(data []interface{}) []MoneyWalletPaidDistributionSegment {
	v := make([]MoneyWalletPaidDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMoneyWalletPaidDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyWalletPaidDistributionSegmentsFromDict(data []MoneyWalletPaidDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyWalletPaidDistribution struct {
	Statistics   *MoneyWalletPaidDistributionStatistics `json:"statistics"`
	Distribution []MoneyWalletPaidDistributionSegment   `json:"distribution"`
}

func NewMoneyWalletPaidDistributionFromJson(data string) MoneyWalletPaidDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyWalletPaidDistributionFromDict(dict)
}

func NewMoneyWalletPaidDistributionFromDict(data map[string]interface{}) MoneyWalletPaidDistribution {
	return MoneyWalletPaidDistribution{
		Statistics:   NewMoneyWalletPaidDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMoneyWalletPaidDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MoneyWalletPaidDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMoneyWalletPaidDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MoneyWalletPaidDistribution) Pointer() *MoneyWalletPaidDistribution {
	return &p
}

func CastMoneyWalletPaidDistributions(data []interface{}) []MoneyWalletPaidDistribution {
	v := make([]MoneyWalletPaidDistribution, 0)
	for _, d := range data {
		v = append(v, NewMoneyWalletPaidDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyWalletPaidDistributionsFromDict(data []MoneyWalletPaidDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyWalletDistributions struct {
	Free *MoneyWalletFreeDistribution `json:"free"`
	Paid *MoneyWalletPaidDistribution `json:"paid"`
}

func NewMoneyWalletDistributionsFromJson(data string) MoneyWalletDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyWalletDistributionsFromDict(dict)
}

func NewMoneyWalletDistributionsFromDict(data map[string]interface{}) MoneyWalletDistributions {
	return MoneyWalletDistributions{
		Free: NewMoneyWalletFreeDistributionFromDict(core.CastMap(data["free"])).Pointer(),
		Paid: NewMoneyWalletPaidDistributionFromDict(core.CastMap(data["paid"])).Pointer(),
	}
}

func (p MoneyWalletDistributions) ToDict() map[string]interface{} {

	var free map[string]interface{}
	if p.Free != nil {
		free = p.Free.ToDict()
	}
	var paid map[string]interface{}
	if p.Paid != nil {
		paid = p.Paid.ToDict()
	}
	return map[string]interface{}{
		"free": free,
		"paid": paid,
	}
}

func (p MoneyWalletDistributions) Pointer() *MoneyWalletDistributions {
	return &p
}

func CastMoneyWalletDistributionses(data []interface{}) []MoneyWalletDistributions {
	v := make([]MoneyWalletDistributions, 0)
	for _, d := range data {
		v = append(v, NewMoneyWalletDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyWalletDistributionsesFromDict(data []MoneyWalletDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyWallet struct {
	WalletId      *string                   `json:"walletId"`
	Slot          *int32                    `json:"slot"`
	Statistics    *MoneyWalletStatistics    `json:"statistics"`
	Distributions *MoneyWalletDistributions `json:"distributions"`
}

func NewMoneyWalletFromJson(data string) MoneyWallet {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyWalletFromDict(dict)
}

func NewMoneyWalletFromDict(data map[string]interface{}) MoneyWallet {
	return MoneyWallet{
		WalletId:      core.CastString(data["walletId"]),
		Slot:          core.CastInt32(data["slot"]),
		Statistics:    NewMoneyWalletStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewMoneyWalletDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p MoneyWallet) ToDict() map[string]interface{} {

	var walletId *string
	if p.WalletId != nil {
		walletId = p.WalletId
	}
	var slot *int32
	if p.Slot != nil {
		slot = p.Slot
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"walletId":      walletId,
		"slot":          slot,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p MoneyWallet) Pointer() *MoneyWallet {
	return &p
}

func CastMoneyWallets(data []interface{}) []MoneyWallet {
	v := make([]MoneyWallet, 0)
	for _, d := range data {
		v = append(v, NewMoneyWalletFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyWalletsFromDict(data []MoneyWallet) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyReceiptStatistics struct {
	Verification *int64 `json:"verification"`
}

func NewMoneyReceiptStatisticsFromJson(data string) MoneyReceiptStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyReceiptStatisticsFromDict(dict)
}

func NewMoneyReceiptStatisticsFromDict(data map[string]interface{}) MoneyReceiptStatistics {
	return MoneyReceiptStatistics{
		Verification: core.CastInt64(data["verification"]),
	}
}

func (p MoneyReceiptStatistics) ToDict() map[string]interface{} {

	var verification *int64
	if p.Verification != nil {
		verification = p.Verification
	}
	return map[string]interface{}{
		"verification": verification,
	}
}

func (p MoneyReceiptStatistics) Pointer() *MoneyReceiptStatistics {
	return &p
}

func CastMoneyReceiptStatisticses(data []interface{}) []MoneyReceiptStatistics {
	v := make([]MoneyReceiptStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyReceiptStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyReceiptStatisticsesFromDict(data []MoneyReceiptStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyReceiptVerificationByUserDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMoneyReceiptVerificationByUserDistributionStatisticsFromJson(data string) MoneyReceiptVerificationByUserDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyReceiptVerificationByUserDistributionStatisticsFromDict(dict)
}

func NewMoneyReceiptVerificationByUserDistributionStatisticsFromDict(data map[string]interface{}) MoneyReceiptVerificationByUserDistributionStatistics {
	return MoneyReceiptVerificationByUserDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MoneyReceiptVerificationByUserDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MoneyReceiptVerificationByUserDistributionStatistics) Pointer() *MoneyReceiptVerificationByUserDistributionStatistics {
	return &p
}

func CastMoneyReceiptVerificationByUserDistributionStatisticses(data []interface{}) []MoneyReceiptVerificationByUserDistributionStatistics {
	v := make([]MoneyReceiptVerificationByUserDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyReceiptVerificationByUserDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyReceiptVerificationByUserDistributionStatisticsesFromDict(data []MoneyReceiptVerificationByUserDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyReceiptVerificationByUserDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewMoneyReceiptVerificationByUserDistributionSegmentFromJson(data string) MoneyReceiptVerificationByUserDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyReceiptVerificationByUserDistributionSegmentFromDict(dict)
}

func NewMoneyReceiptVerificationByUserDistributionSegmentFromDict(data map[string]interface{}) MoneyReceiptVerificationByUserDistributionSegment {
	return MoneyReceiptVerificationByUserDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p MoneyReceiptVerificationByUserDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p MoneyReceiptVerificationByUserDistributionSegment) Pointer() *MoneyReceiptVerificationByUserDistributionSegment {
	return &p
}

func CastMoneyReceiptVerificationByUserDistributionSegments(data []interface{}) []MoneyReceiptVerificationByUserDistributionSegment {
	v := make([]MoneyReceiptVerificationByUserDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMoneyReceiptVerificationByUserDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyReceiptVerificationByUserDistributionSegmentsFromDict(data []MoneyReceiptVerificationByUserDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyReceiptVerificationByUserDistribution struct {
	Statistics   *MoneyReceiptVerificationByUserDistributionStatistics `json:"statistics"`
	Distribution []MoneyReceiptVerificationByUserDistributionSegment   `json:"distribution"`
}

func NewMoneyReceiptVerificationByUserDistributionFromJson(data string) MoneyReceiptVerificationByUserDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyReceiptVerificationByUserDistributionFromDict(dict)
}

func NewMoneyReceiptVerificationByUserDistributionFromDict(data map[string]interface{}) MoneyReceiptVerificationByUserDistribution {
	return MoneyReceiptVerificationByUserDistribution{
		Statistics:   NewMoneyReceiptVerificationByUserDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMoneyReceiptVerificationByUserDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MoneyReceiptVerificationByUserDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMoneyReceiptVerificationByUserDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MoneyReceiptVerificationByUserDistribution) Pointer() *MoneyReceiptVerificationByUserDistribution {
	return &p
}

func CastMoneyReceiptVerificationByUserDistributions(data []interface{}) []MoneyReceiptVerificationByUserDistribution {
	v := make([]MoneyReceiptVerificationByUserDistribution, 0)
	for _, d := range data {
		v = append(v, NewMoneyReceiptVerificationByUserDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyReceiptVerificationByUserDistributionsFromDict(data []MoneyReceiptVerificationByUserDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyReceiptDistributions struct {
	VerificationByUser *MoneyReceiptVerificationByUserDistribution `json:"verificationByUser"`
}

func NewMoneyReceiptDistributionsFromJson(data string) MoneyReceiptDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyReceiptDistributionsFromDict(dict)
}

func NewMoneyReceiptDistributionsFromDict(data map[string]interface{}) MoneyReceiptDistributions {
	return MoneyReceiptDistributions{
		VerificationByUser: NewMoneyReceiptVerificationByUserDistributionFromDict(core.CastMap(data["verificationByUser"])).Pointer(),
	}
}

func (p MoneyReceiptDistributions) ToDict() map[string]interface{} {

	var verificationByUser map[string]interface{}
	if p.VerificationByUser != nil {
		verificationByUser = p.VerificationByUser.ToDict()
	}
	return map[string]interface{}{
		"verificationByUser": verificationByUser,
	}
}

func (p MoneyReceiptDistributions) Pointer() *MoneyReceiptDistributions {
	return &p
}

func CastMoneyReceiptDistributionses(data []interface{}) []MoneyReceiptDistributions {
	v := make([]MoneyReceiptDistributions, 0)
	for _, d := range data {
		v = append(v, NewMoneyReceiptDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyReceiptDistributionsesFromDict(data []MoneyReceiptDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyReceipt struct {
	ReceiptId     *string                    `json:"receiptId"`
	ContentsId    *string                    `json:"contentsId"`
	Statistics    *MoneyReceiptStatistics    `json:"statistics"`
	Distributions *MoneyReceiptDistributions `json:"distributions"`
}

func NewMoneyReceiptFromJson(data string) MoneyReceipt {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyReceiptFromDict(dict)
}

func NewMoneyReceiptFromDict(data map[string]interface{}) MoneyReceipt {
	return MoneyReceipt{
		ReceiptId:     core.CastString(data["receiptId"]),
		ContentsId:    core.CastString(data["contentsId"]),
		Statistics:    NewMoneyReceiptStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewMoneyReceiptDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p MoneyReceipt) ToDict() map[string]interface{} {

	var receiptId *string
	if p.ReceiptId != nil {
		receiptId = p.ReceiptId
	}
	var contentsId *string
	if p.ContentsId != nil {
		contentsId = p.ContentsId
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"receiptId":     receiptId,
		"contentsId":    contentsId,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p MoneyReceipt) Pointer() *MoneyReceipt {
	return &p
}

func CastMoneyReceipts(data []interface{}) []MoneyReceipt {
	v := make([]MoneyReceipt, 0)
	for _, d := range data {
		v = append(v, NewMoneyReceiptFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyReceiptsFromDict(data []MoneyReceipt) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceStatistics struct {
	Verification *int64 `json:"verification"`
	Deposit      *int64 `json:"deposit"`
	Withdraw     *int64 `json:"withdraw"`
	Revenue      *int64 `json:"revenue"`
}

func NewMoneyNamespaceStatisticsFromJson(data string) MoneyNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceStatisticsFromDict(dict)
}

func NewMoneyNamespaceStatisticsFromDict(data map[string]interface{}) MoneyNamespaceStatistics {
	return MoneyNamespaceStatistics{
		Verification: core.CastInt64(data["verification"]),
		Deposit:      core.CastInt64(data["deposit"]),
		Withdraw:     core.CastInt64(data["withdraw"]),
		Revenue:      core.CastInt64(data["revenue"]),
	}
}

func (p MoneyNamespaceStatistics) ToDict() map[string]interface{} {

	var verification *int64
	if p.Verification != nil {
		verification = p.Verification
	}
	var deposit *int64
	if p.Deposit != nil {
		deposit = p.Deposit
	}
	var withdraw *int64
	if p.Withdraw != nil {
		withdraw = p.Withdraw
	}
	var revenue *int64
	if p.Revenue != nil {
		revenue = p.Revenue
	}
	return map[string]interface{}{
		"verification": verification,
		"deposit":      deposit,
		"withdraw":     withdraw,
		"revenue":      revenue,
	}
}

func (p MoneyNamespaceStatistics) Pointer() *MoneyNamespaceStatistics {
	return &p
}

func CastMoneyNamespaceStatisticses(data []interface{}) []MoneyNamespaceStatistics {
	v := make([]MoneyNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceStatisticsesFromDict(data []MoneyNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceVerificationDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMoneyNamespaceVerificationDistributionStatisticsFromJson(data string) MoneyNamespaceVerificationDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceVerificationDistributionStatisticsFromDict(dict)
}

func NewMoneyNamespaceVerificationDistributionStatisticsFromDict(data map[string]interface{}) MoneyNamespaceVerificationDistributionStatistics {
	return MoneyNamespaceVerificationDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MoneyNamespaceVerificationDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MoneyNamespaceVerificationDistributionStatistics) Pointer() *MoneyNamespaceVerificationDistributionStatistics {
	return &p
}

func CastMoneyNamespaceVerificationDistributionStatisticses(data []interface{}) []MoneyNamespaceVerificationDistributionStatistics {
	v := make([]MoneyNamespaceVerificationDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceVerificationDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceVerificationDistributionStatisticsesFromDict(data []MoneyNamespaceVerificationDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceVerificationDistributionSegment struct {
	ContentsId *string `json:"contentsId"`
	Count      *int64  `json:"count"`
}

func NewMoneyNamespaceVerificationDistributionSegmentFromJson(data string) MoneyNamespaceVerificationDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceVerificationDistributionSegmentFromDict(dict)
}

func NewMoneyNamespaceVerificationDistributionSegmentFromDict(data map[string]interface{}) MoneyNamespaceVerificationDistributionSegment {
	return MoneyNamespaceVerificationDistributionSegment{
		ContentsId: core.CastString(data["contentsId"]),
		Count:      core.CastInt64(data["count"]),
	}
}

func (p MoneyNamespaceVerificationDistributionSegment) ToDict() map[string]interface{} {

	var contentsId *string
	if p.ContentsId != nil {
		contentsId = p.ContentsId
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"contentsId": contentsId,
		"count":      count,
	}
}

func (p MoneyNamespaceVerificationDistributionSegment) Pointer() *MoneyNamespaceVerificationDistributionSegment {
	return &p
}

func CastMoneyNamespaceVerificationDistributionSegments(data []interface{}) []MoneyNamespaceVerificationDistributionSegment {
	v := make([]MoneyNamespaceVerificationDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceVerificationDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceVerificationDistributionSegmentsFromDict(data []MoneyNamespaceVerificationDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceVerificationDistribution struct {
	Statistics   *MoneyNamespaceVerificationDistributionStatistics `json:"statistics"`
	Distribution []MoneyNamespaceVerificationDistributionSegment   `json:"distribution"`
}

func NewMoneyNamespaceVerificationDistributionFromJson(data string) MoneyNamespaceVerificationDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceVerificationDistributionFromDict(dict)
}

func NewMoneyNamespaceVerificationDistributionFromDict(data map[string]interface{}) MoneyNamespaceVerificationDistribution {
	return MoneyNamespaceVerificationDistribution{
		Statistics:   NewMoneyNamespaceVerificationDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMoneyNamespaceVerificationDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MoneyNamespaceVerificationDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMoneyNamespaceVerificationDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MoneyNamespaceVerificationDistribution) Pointer() *MoneyNamespaceVerificationDistribution {
	return &p
}

func CastMoneyNamespaceVerificationDistributions(data []interface{}) []MoneyNamespaceVerificationDistribution {
	v := make([]MoneyNamespaceVerificationDistribution, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceVerificationDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceVerificationDistributionsFromDict(data []MoneyNamespaceVerificationDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceDepositDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMoneyNamespaceDepositDistributionStatisticsFromJson(data string) MoneyNamespaceDepositDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceDepositDistributionStatisticsFromDict(dict)
}

func NewMoneyNamespaceDepositDistributionStatisticsFromDict(data map[string]interface{}) MoneyNamespaceDepositDistributionStatistics {
	return MoneyNamespaceDepositDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MoneyNamespaceDepositDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MoneyNamespaceDepositDistributionStatistics) Pointer() *MoneyNamespaceDepositDistributionStatistics {
	return &p
}

func CastMoneyNamespaceDepositDistributionStatisticses(data []interface{}) []MoneyNamespaceDepositDistributionStatistics {
	v := make([]MoneyNamespaceDepositDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceDepositDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceDepositDistributionStatisticsesFromDict(data []MoneyNamespaceDepositDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceDepositDistributionSegment struct {
	Slot  *int32 `json:"slot"`
	Count *int64 `json:"count"`
}

func NewMoneyNamespaceDepositDistributionSegmentFromJson(data string) MoneyNamespaceDepositDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceDepositDistributionSegmentFromDict(dict)
}

func NewMoneyNamespaceDepositDistributionSegmentFromDict(data map[string]interface{}) MoneyNamespaceDepositDistributionSegment {
	return MoneyNamespaceDepositDistributionSegment{
		Slot:  core.CastInt32(data["slot"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p MoneyNamespaceDepositDistributionSegment) ToDict() map[string]interface{} {

	var slot *int32
	if p.Slot != nil {
		slot = p.Slot
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"slot":  slot,
		"count": count,
	}
}

func (p MoneyNamespaceDepositDistributionSegment) Pointer() *MoneyNamespaceDepositDistributionSegment {
	return &p
}

func CastMoneyNamespaceDepositDistributionSegments(data []interface{}) []MoneyNamespaceDepositDistributionSegment {
	v := make([]MoneyNamespaceDepositDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceDepositDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceDepositDistributionSegmentsFromDict(data []MoneyNamespaceDepositDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceDepositDistribution struct {
	Statistics   *MoneyNamespaceDepositDistributionStatistics `json:"statistics"`
	Distribution []MoneyNamespaceDepositDistributionSegment   `json:"distribution"`
}

func NewMoneyNamespaceDepositDistributionFromJson(data string) MoneyNamespaceDepositDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceDepositDistributionFromDict(dict)
}

func NewMoneyNamespaceDepositDistributionFromDict(data map[string]interface{}) MoneyNamespaceDepositDistribution {
	return MoneyNamespaceDepositDistribution{
		Statistics:   NewMoneyNamespaceDepositDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMoneyNamespaceDepositDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MoneyNamespaceDepositDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMoneyNamespaceDepositDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MoneyNamespaceDepositDistribution) Pointer() *MoneyNamespaceDepositDistribution {
	return &p
}

func CastMoneyNamespaceDepositDistributions(data []interface{}) []MoneyNamespaceDepositDistribution {
	v := make([]MoneyNamespaceDepositDistribution, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceDepositDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceDepositDistributionsFromDict(data []MoneyNamespaceDepositDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceWithdrawDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMoneyNamespaceWithdrawDistributionStatisticsFromJson(data string) MoneyNamespaceWithdrawDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceWithdrawDistributionStatisticsFromDict(dict)
}

func NewMoneyNamespaceWithdrawDistributionStatisticsFromDict(data map[string]interface{}) MoneyNamespaceWithdrawDistributionStatistics {
	return MoneyNamespaceWithdrawDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MoneyNamespaceWithdrawDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MoneyNamespaceWithdrawDistributionStatistics) Pointer() *MoneyNamespaceWithdrawDistributionStatistics {
	return &p
}

func CastMoneyNamespaceWithdrawDistributionStatisticses(data []interface{}) []MoneyNamespaceWithdrawDistributionStatistics {
	v := make([]MoneyNamespaceWithdrawDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceWithdrawDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceWithdrawDistributionStatisticsesFromDict(data []MoneyNamespaceWithdrawDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceWithdrawDistributionSegment struct {
	Slot  *int32 `json:"slot"`
	Count *int64 `json:"count"`
}

func NewMoneyNamespaceWithdrawDistributionSegmentFromJson(data string) MoneyNamespaceWithdrawDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceWithdrawDistributionSegmentFromDict(dict)
}

func NewMoneyNamespaceWithdrawDistributionSegmentFromDict(data map[string]interface{}) MoneyNamespaceWithdrawDistributionSegment {
	return MoneyNamespaceWithdrawDistributionSegment{
		Slot:  core.CastInt32(data["slot"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p MoneyNamespaceWithdrawDistributionSegment) ToDict() map[string]interface{} {

	var slot *int32
	if p.Slot != nil {
		slot = p.Slot
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"slot":  slot,
		"count": count,
	}
}

func (p MoneyNamespaceWithdrawDistributionSegment) Pointer() *MoneyNamespaceWithdrawDistributionSegment {
	return &p
}

func CastMoneyNamespaceWithdrawDistributionSegments(data []interface{}) []MoneyNamespaceWithdrawDistributionSegment {
	v := make([]MoneyNamespaceWithdrawDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceWithdrawDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceWithdrawDistributionSegmentsFromDict(data []MoneyNamespaceWithdrawDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceWithdrawDistribution struct {
	Statistics   *MoneyNamespaceWithdrawDistributionStatistics `json:"statistics"`
	Distribution []MoneyNamespaceWithdrawDistributionSegment   `json:"distribution"`
}

func NewMoneyNamespaceWithdrawDistributionFromJson(data string) MoneyNamespaceWithdrawDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceWithdrawDistributionFromDict(dict)
}

func NewMoneyNamespaceWithdrawDistributionFromDict(data map[string]interface{}) MoneyNamespaceWithdrawDistribution {
	return MoneyNamespaceWithdrawDistribution{
		Statistics:   NewMoneyNamespaceWithdrawDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMoneyNamespaceWithdrawDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MoneyNamespaceWithdrawDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMoneyNamespaceWithdrawDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MoneyNamespaceWithdrawDistribution) Pointer() *MoneyNamespaceWithdrawDistribution {
	return &p
}

func CastMoneyNamespaceWithdrawDistributions(data []interface{}) []MoneyNamespaceWithdrawDistribution {
	v := make([]MoneyNamespaceWithdrawDistribution, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceWithdrawDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceWithdrawDistributionsFromDict(data []MoneyNamespaceWithdrawDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceRevenueDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewMoneyNamespaceRevenueDistributionStatisticsFromJson(data string) MoneyNamespaceRevenueDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceRevenueDistributionStatisticsFromDict(dict)
}

func NewMoneyNamespaceRevenueDistributionStatisticsFromDict(data map[string]interface{}) MoneyNamespaceRevenueDistributionStatistics {
	return MoneyNamespaceRevenueDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p MoneyNamespaceRevenueDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p MoneyNamespaceRevenueDistributionStatistics) Pointer() *MoneyNamespaceRevenueDistributionStatistics {
	return &p
}

func CastMoneyNamespaceRevenueDistributionStatisticses(data []interface{}) []MoneyNamespaceRevenueDistributionStatistics {
	v := make([]MoneyNamespaceRevenueDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceRevenueDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceRevenueDistributionStatisticsesFromDict(data []MoneyNamespaceRevenueDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceRevenueDistributionSegment struct {
	Slot *int32 `json:"slot"`
	Sum  *int64 `json:"sum"`
}

func NewMoneyNamespaceRevenueDistributionSegmentFromJson(data string) MoneyNamespaceRevenueDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceRevenueDistributionSegmentFromDict(dict)
}

func NewMoneyNamespaceRevenueDistributionSegmentFromDict(data map[string]interface{}) MoneyNamespaceRevenueDistributionSegment {
	return MoneyNamespaceRevenueDistributionSegment{
		Slot: core.CastInt32(data["slot"]),
		Sum:  core.CastInt64(data["sum"]),
	}
}

func (p MoneyNamespaceRevenueDistributionSegment) ToDict() map[string]interface{} {

	var slot *int32
	if p.Slot != nil {
		slot = p.Slot
	}
	var sum *int64
	if p.Sum != nil {
		sum = p.Sum
	}
	return map[string]interface{}{
		"slot": slot,
		"sum":  sum,
	}
}

func (p MoneyNamespaceRevenueDistributionSegment) Pointer() *MoneyNamespaceRevenueDistributionSegment {
	return &p
}

func CastMoneyNamespaceRevenueDistributionSegments(data []interface{}) []MoneyNamespaceRevenueDistributionSegment {
	v := make([]MoneyNamespaceRevenueDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceRevenueDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceRevenueDistributionSegmentsFromDict(data []MoneyNamespaceRevenueDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceRevenueDistribution struct {
	Statistics   *MoneyNamespaceRevenueDistributionStatistics `json:"statistics"`
	Distribution []MoneyNamespaceRevenueDistributionSegment   `json:"distribution"`
}

func NewMoneyNamespaceRevenueDistributionFromJson(data string) MoneyNamespaceRevenueDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceRevenueDistributionFromDict(dict)
}

func NewMoneyNamespaceRevenueDistributionFromDict(data map[string]interface{}) MoneyNamespaceRevenueDistribution {
	return MoneyNamespaceRevenueDistribution{
		Statistics:   NewMoneyNamespaceRevenueDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastMoneyNamespaceRevenueDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p MoneyNamespaceRevenueDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastMoneyNamespaceRevenueDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p MoneyNamespaceRevenueDistribution) Pointer() *MoneyNamespaceRevenueDistribution {
	return &p
}

func CastMoneyNamespaceRevenueDistributions(data []interface{}) []MoneyNamespaceRevenueDistribution {
	v := make([]MoneyNamespaceRevenueDistribution, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceRevenueDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceRevenueDistributionsFromDict(data []MoneyNamespaceRevenueDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespaceDistributions struct {
	Verification *MoneyNamespaceVerificationDistribution `json:"verification"`
	Deposit      *MoneyNamespaceDepositDistribution      `json:"deposit"`
	Withdraw     *MoneyNamespaceWithdrawDistribution     `json:"withdraw"`
	Revenue      *MoneyNamespaceRevenueDistribution      `json:"revenue"`
}

func NewMoneyNamespaceDistributionsFromJson(data string) MoneyNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceDistributionsFromDict(dict)
}

func NewMoneyNamespaceDistributionsFromDict(data map[string]interface{}) MoneyNamespaceDistributions {
	return MoneyNamespaceDistributions{
		Verification: NewMoneyNamespaceVerificationDistributionFromDict(core.CastMap(data["verification"])).Pointer(),
		Deposit:      NewMoneyNamespaceDepositDistributionFromDict(core.CastMap(data["deposit"])).Pointer(),
		Withdraw:     NewMoneyNamespaceWithdrawDistributionFromDict(core.CastMap(data["withdraw"])).Pointer(),
		Revenue:      NewMoneyNamespaceRevenueDistributionFromDict(core.CastMap(data["revenue"])).Pointer(),
	}
}

func (p MoneyNamespaceDistributions) ToDict() map[string]interface{} {

	var verification map[string]interface{}
	if p.Verification != nil {
		verification = p.Verification.ToDict()
	}
	var deposit map[string]interface{}
	if p.Deposit != nil {
		deposit = p.Deposit.ToDict()
	}
	var withdraw map[string]interface{}
	if p.Withdraw != nil {
		withdraw = p.Withdraw.ToDict()
	}
	var revenue map[string]interface{}
	if p.Revenue != nil {
		revenue = p.Revenue.ToDict()
	}
	return map[string]interface{}{
		"verification": verification,
		"deposit":      deposit,
		"withdraw":     withdraw,
		"revenue":      revenue,
	}
}

func (p MoneyNamespaceDistributions) Pointer() *MoneyNamespaceDistributions {
	return &p
}

func CastMoneyNamespaceDistributionses(data []interface{}) []MoneyNamespaceDistributions {
	v := make([]MoneyNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespaceDistributionsesFromDict(data []MoneyNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoneyNamespace struct {
	NamespaceId   *string                      `json:"namespaceId"`
	Year          *int32                       `json:"year"`
	Month         *int32                       `json:"month"`
	Day           *int32                       `json:"day"`
	NamespaceName *string                      `json:"namespaceName"`
	Statistics    *MoneyNamespaceStatistics    `json:"statistics"`
	Distributions *MoneyNamespaceDistributions `json:"distributions"`
	Wallets       []MoneyWallet                `json:"wallets"`
	Receipts      []MoneyReceipt               `json:"receipts"`
}

func NewMoneyNamespaceFromJson(data string) MoneyNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoneyNamespaceFromDict(dict)
}

func NewMoneyNamespaceFromDict(data map[string]interface{}) MoneyNamespace {
	return MoneyNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewMoneyNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewMoneyNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		Wallets:       CastMoneyWallets(core.CastArray(data["wallets"])),
		Receipts:      CastMoneyReceipts(core.CastArray(data["receipts"])),
	}
}

func (p MoneyNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var wallets []interface{}
	if p.Wallets != nil {
		wallets = CastMoneyWalletsFromDict(
			p.Wallets,
		)
	}
	var receipts []interface{}
	if p.Receipts != nil {
		receipts = CastMoneyReceiptsFromDict(
			p.Receipts,
		)
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
		"wallets":       wallets,
		"receipts":      receipts,
	}
}

func (p MoneyNamespace) Pointer() *MoneyNamespace {
	return &p
}

func CastMoneyNamespaces(data []interface{}) []MoneyNamespace {
	v := make([]MoneyNamespace, 0)
	for _, d := range data {
		v = append(v, NewMoneyNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoneyNamespacesFromDict(data []MoneyNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestModelStatistics struct {
	StartQuest     *int64   `json:"startQuest"`
	EndQuest       *int64   `json:"endQuest"`
	Successful     *int64   `json:"successful"`
	SuccessfulRate *float32 `json:"successfulRate"`
}

func NewQuestQuestModelStatisticsFromJson(data string) QuestQuestModelStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestModelStatisticsFromDict(dict)
}

func NewQuestQuestModelStatisticsFromDict(data map[string]interface{}) QuestQuestModelStatistics {
	return QuestQuestModelStatistics{
		StartQuest:     core.CastInt64(data["startQuest"]),
		EndQuest:       core.CastInt64(data["endQuest"]),
		Successful:     core.CastInt64(data["successful"]),
		SuccessfulRate: core.CastFloat32(data["successfulRate"]),
	}
}

func (p QuestQuestModelStatistics) ToDict() map[string]interface{} {

	var startQuest *int64
	if p.StartQuest != nil {
		startQuest = p.StartQuest
	}
	var endQuest *int64
	if p.EndQuest != nil {
		endQuest = p.EndQuest
	}
	var successful *int64
	if p.Successful != nil {
		successful = p.Successful
	}
	var successfulRate *float32
	if p.SuccessfulRate != nil {
		successfulRate = p.SuccessfulRate
	}
	return map[string]interface{}{
		"startQuest":     startQuest,
		"endQuest":       endQuest,
		"successful":     successful,
		"successfulRate": successfulRate,
	}
}

func (p QuestQuestModelStatistics) Pointer() *QuestQuestModelStatistics {
	return &p
}

func CastQuestQuestModelStatisticses(data []interface{}) []QuestQuestModelStatistics {
	v := make([]QuestQuestModelStatistics, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestModelStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestModelStatisticsesFromDict(data []QuestQuestModelStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestModelPlayTimeSecondsDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewQuestQuestModelPlayTimeSecondsDistributionStatisticsFromJson(data string) QuestQuestModelPlayTimeSecondsDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestModelPlayTimeSecondsDistributionStatisticsFromDict(dict)
}

func NewQuestQuestModelPlayTimeSecondsDistributionStatisticsFromDict(data map[string]interface{}) QuestQuestModelPlayTimeSecondsDistributionStatistics {
	return QuestQuestModelPlayTimeSecondsDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p QuestQuestModelPlayTimeSecondsDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p QuestQuestModelPlayTimeSecondsDistributionStatistics) Pointer() *QuestQuestModelPlayTimeSecondsDistributionStatistics {
	return &p
}

func CastQuestQuestModelPlayTimeSecondsDistributionStatisticses(data []interface{}) []QuestQuestModelPlayTimeSecondsDistributionStatistics {
	v := make([]QuestQuestModelPlayTimeSecondsDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestModelPlayTimeSecondsDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestModelPlayTimeSecondsDistributionStatisticsesFromDict(data []QuestQuestModelPlayTimeSecondsDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestModelPlayTimeSecondsDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewQuestQuestModelPlayTimeSecondsDistributionSegmentFromJson(data string) QuestQuestModelPlayTimeSecondsDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestModelPlayTimeSecondsDistributionSegmentFromDict(dict)
}

func NewQuestQuestModelPlayTimeSecondsDistributionSegmentFromDict(data map[string]interface{}) QuestQuestModelPlayTimeSecondsDistributionSegment {
	return QuestQuestModelPlayTimeSecondsDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p QuestQuestModelPlayTimeSecondsDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p QuestQuestModelPlayTimeSecondsDistributionSegment) Pointer() *QuestQuestModelPlayTimeSecondsDistributionSegment {
	return &p
}

func CastQuestQuestModelPlayTimeSecondsDistributionSegments(data []interface{}) []QuestQuestModelPlayTimeSecondsDistributionSegment {
	v := make([]QuestQuestModelPlayTimeSecondsDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestModelPlayTimeSecondsDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestModelPlayTimeSecondsDistributionSegmentsFromDict(data []QuestQuestModelPlayTimeSecondsDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestModelPlayTimeSecondsDistribution struct {
	Statistics   *QuestQuestModelPlayTimeSecondsDistributionStatistics `json:"statistics"`
	Distribution []QuestQuestModelPlayTimeSecondsDistributionSegment   `json:"distribution"`
}

func NewQuestQuestModelPlayTimeSecondsDistributionFromJson(data string) QuestQuestModelPlayTimeSecondsDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestModelPlayTimeSecondsDistributionFromDict(dict)
}

func NewQuestQuestModelPlayTimeSecondsDistributionFromDict(data map[string]interface{}) QuestQuestModelPlayTimeSecondsDistribution {
	return QuestQuestModelPlayTimeSecondsDistribution{
		Statistics:   NewQuestQuestModelPlayTimeSecondsDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastQuestQuestModelPlayTimeSecondsDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p QuestQuestModelPlayTimeSecondsDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastQuestQuestModelPlayTimeSecondsDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p QuestQuestModelPlayTimeSecondsDistribution) Pointer() *QuestQuestModelPlayTimeSecondsDistribution {
	return &p
}

func CastQuestQuestModelPlayTimeSecondsDistributions(data []interface{}) []QuestQuestModelPlayTimeSecondsDistribution {
	v := make([]QuestQuestModelPlayTimeSecondsDistribution, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestModelPlayTimeSecondsDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestModelPlayTimeSecondsDistributionsFromDict(data []QuestQuestModelPlayTimeSecondsDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestModelDistributions struct {
	PlayTimeSeconds *QuestQuestModelPlayTimeSecondsDistribution `json:"playTimeSeconds"`
}

func NewQuestQuestModelDistributionsFromJson(data string) QuestQuestModelDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestModelDistributionsFromDict(dict)
}

func NewQuestQuestModelDistributionsFromDict(data map[string]interface{}) QuestQuestModelDistributions {
	return QuestQuestModelDistributions{
		PlayTimeSeconds: NewQuestQuestModelPlayTimeSecondsDistributionFromDict(core.CastMap(data["playTimeSeconds"])).Pointer(),
	}
}

func (p QuestQuestModelDistributions) ToDict() map[string]interface{} {

	var playTimeSeconds map[string]interface{}
	if p.PlayTimeSeconds != nil {
		playTimeSeconds = p.PlayTimeSeconds.ToDict()
	}
	return map[string]interface{}{
		"playTimeSeconds": playTimeSeconds,
	}
}

func (p QuestQuestModelDistributions) Pointer() *QuestQuestModelDistributions {
	return &p
}

func CastQuestQuestModelDistributionses(data []interface{}) []QuestQuestModelDistributions {
	v := make([]QuestQuestModelDistributions, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestModelDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestModelDistributionsesFromDict(data []QuestQuestModelDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestModel struct {
	QuestModelId  *string                       `json:"questModelId"`
	QuestName     *string                       `json:"questName"`
	Statistics    *QuestQuestModelStatistics    `json:"statistics"`
	Distributions *QuestQuestModelDistributions `json:"distributions"`
}

func NewQuestQuestModelFromJson(data string) QuestQuestModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestModelFromDict(dict)
}

func NewQuestQuestModelFromDict(data map[string]interface{}) QuestQuestModel {
	return QuestQuestModel{
		QuestModelId:  core.CastString(data["questModelId"]),
		QuestName:     core.CastString(data["questName"]),
		Statistics:    NewQuestQuestModelStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewQuestQuestModelDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p QuestQuestModel) ToDict() map[string]interface{} {

	var questModelId *string
	if p.QuestModelId != nil {
		questModelId = p.QuestModelId
	}
	var questName *string
	if p.QuestName != nil {
		questName = p.QuestName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"questModelId":  questModelId,
		"questName":     questName,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p QuestQuestModel) Pointer() *QuestQuestModel {
	return &p
}

func CastQuestQuestModels(data []interface{}) []QuestQuestModel {
	v := make([]QuestQuestModel, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestModelsFromDict(data []QuestQuestModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestGroupModelStatistics struct {
	StartQuest     *int64   `json:"startQuest"`
	EndQuest       *int64   `json:"endQuest"`
	Successful     *int64   `json:"successful"`
	SuccessfulRate *float32 `json:"successfulRate"`
}

func NewQuestQuestGroupModelStatisticsFromJson(data string) QuestQuestGroupModelStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestGroupModelStatisticsFromDict(dict)
}

func NewQuestQuestGroupModelStatisticsFromDict(data map[string]interface{}) QuestQuestGroupModelStatistics {
	return QuestQuestGroupModelStatistics{
		StartQuest:     core.CastInt64(data["startQuest"]),
		EndQuest:       core.CastInt64(data["endQuest"]),
		Successful:     core.CastInt64(data["successful"]),
		SuccessfulRate: core.CastFloat32(data["successfulRate"]),
	}
}

func (p QuestQuestGroupModelStatistics) ToDict() map[string]interface{} {

	var startQuest *int64
	if p.StartQuest != nil {
		startQuest = p.StartQuest
	}
	var endQuest *int64
	if p.EndQuest != nil {
		endQuest = p.EndQuest
	}
	var successful *int64
	if p.Successful != nil {
		successful = p.Successful
	}
	var successfulRate *float32
	if p.SuccessfulRate != nil {
		successfulRate = p.SuccessfulRate
	}
	return map[string]interface{}{
		"startQuest":     startQuest,
		"endQuest":       endQuest,
		"successful":     successful,
		"successfulRate": successfulRate,
	}
}

func (p QuestQuestGroupModelStatistics) Pointer() *QuestQuestGroupModelStatistics {
	return &p
}

func CastQuestQuestGroupModelStatisticses(data []interface{}) []QuestQuestGroupModelStatistics {
	v := make([]QuestQuestGroupModelStatistics, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestGroupModelStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestGroupModelStatisticsesFromDict(data []QuestQuestGroupModelStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestGroupModelQuestDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewQuestQuestGroupModelQuestDistributionStatisticsFromJson(data string) QuestQuestGroupModelQuestDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestGroupModelQuestDistributionStatisticsFromDict(dict)
}

func NewQuestQuestGroupModelQuestDistributionStatisticsFromDict(data map[string]interface{}) QuestQuestGroupModelQuestDistributionStatistics {
	return QuestQuestGroupModelQuestDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p QuestQuestGroupModelQuestDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p QuestQuestGroupModelQuestDistributionStatistics) Pointer() *QuestQuestGroupModelQuestDistributionStatistics {
	return &p
}

func CastQuestQuestGroupModelQuestDistributionStatisticses(data []interface{}) []QuestQuestGroupModelQuestDistributionStatistics {
	v := make([]QuestQuestGroupModelQuestDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestGroupModelQuestDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestGroupModelQuestDistributionStatisticsesFromDict(data []QuestQuestGroupModelQuestDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestGroupModelQuestDistributionSegment struct {
	QuestName *string `json:"questName"`
	Count     *int64  `json:"count"`
}

func NewQuestQuestGroupModelQuestDistributionSegmentFromJson(data string) QuestQuestGroupModelQuestDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestGroupModelQuestDistributionSegmentFromDict(dict)
}

func NewQuestQuestGroupModelQuestDistributionSegmentFromDict(data map[string]interface{}) QuestQuestGroupModelQuestDistributionSegment {
	return QuestQuestGroupModelQuestDistributionSegment{
		QuestName: core.CastString(data["questName"]),
		Count:     core.CastInt64(data["count"]),
	}
}

func (p QuestQuestGroupModelQuestDistributionSegment) ToDict() map[string]interface{} {

	var questName *string
	if p.QuestName != nil {
		questName = p.QuestName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"questName": questName,
		"count":     count,
	}
}

func (p QuestQuestGroupModelQuestDistributionSegment) Pointer() *QuestQuestGroupModelQuestDistributionSegment {
	return &p
}

func CastQuestQuestGroupModelQuestDistributionSegments(data []interface{}) []QuestQuestGroupModelQuestDistributionSegment {
	v := make([]QuestQuestGroupModelQuestDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestGroupModelQuestDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestGroupModelQuestDistributionSegmentsFromDict(data []QuestQuestGroupModelQuestDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestGroupModelQuestDistribution struct {
	Statistics   *QuestQuestGroupModelQuestDistributionStatistics `json:"statistics"`
	Distribution []QuestQuestGroupModelQuestDistributionSegment   `json:"distribution"`
}

func NewQuestQuestGroupModelQuestDistributionFromJson(data string) QuestQuestGroupModelQuestDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestGroupModelQuestDistributionFromDict(dict)
}

func NewQuestQuestGroupModelQuestDistributionFromDict(data map[string]interface{}) QuestQuestGroupModelQuestDistribution {
	return QuestQuestGroupModelQuestDistribution{
		Statistics:   NewQuestQuestGroupModelQuestDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastQuestQuestGroupModelQuestDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p QuestQuestGroupModelQuestDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastQuestQuestGroupModelQuestDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p QuestQuestGroupModelQuestDistribution) Pointer() *QuestQuestGroupModelQuestDistribution {
	return &p
}

func CastQuestQuestGroupModelQuestDistributions(data []interface{}) []QuestQuestGroupModelQuestDistribution {
	v := make([]QuestQuestGroupModelQuestDistribution, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestGroupModelQuestDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestGroupModelQuestDistributionsFromDict(data []QuestQuestGroupModelQuestDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestGroupModelDistributions struct {
	Quest *QuestQuestGroupModelQuestDistribution `json:"quest"`
}

func NewQuestQuestGroupModelDistributionsFromJson(data string) QuestQuestGroupModelDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestGroupModelDistributionsFromDict(dict)
}

func NewQuestQuestGroupModelDistributionsFromDict(data map[string]interface{}) QuestQuestGroupModelDistributions {
	return QuestQuestGroupModelDistributions{
		Quest: NewQuestQuestGroupModelQuestDistributionFromDict(core.CastMap(data["quest"])).Pointer(),
	}
}

func (p QuestQuestGroupModelDistributions) ToDict() map[string]interface{} {

	var quest map[string]interface{}
	if p.Quest != nil {
		quest = p.Quest.ToDict()
	}
	return map[string]interface{}{
		"quest": quest,
	}
}

func (p QuestQuestGroupModelDistributions) Pointer() *QuestQuestGroupModelDistributions {
	return &p
}

func CastQuestQuestGroupModelDistributionses(data []interface{}) []QuestQuestGroupModelDistributions {
	v := make([]QuestQuestGroupModelDistributions, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestGroupModelDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestGroupModelDistributionsesFromDict(data []QuestQuestGroupModelDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestQuestGroupModel struct {
	QuestGroupModelId *string                            `json:"questGroupModelId"`
	QuestGroupName    *string                            `json:"questGroupName"`
	Statistics        *QuestQuestGroupModelStatistics    `json:"statistics"`
	Distributions     *QuestQuestGroupModelDistributions `json:"distributions"`
	QuestModels       []QuestQuestModel                  `json:"questModels"`
}

func NewQuestQuestGroupModelFromJson(data string) QuestQuestGroupModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestQuestGroupModelFromDict(dict)
}

func NewQuestQuestGroupModelFromDict(data map[string]interface{}) QuestQuestGroupModel {
	return QuestQuestGroupModel{
		QuestGroupModelId: core.CastString(data["questGroupModelId"]),
		QuestGroupName:    core.CastString(data["questGroupName"]),
		Statistics:        NewQuestQuestGroupModelStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions:     NewQuestQuestGroupModelDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		QuestModels:       CastQuestQuestModels(core.CastArray(data["questModels"])),
	}
}

func (p QuestQuestGroupModel) ToDict() map[string]interface{} {

	var questGroupModelId *string
	if p.QuestGroupModelId != nil {
		questGroupModelId = p.QuestGroupModelId
	}
	var questGroupName *string
	if p.QuestGroupName != nil {
		questGroupName = p.QuestGroupName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var questModels []interface{}
	if p.QuestModels != nil {
		questModels = CastQuestQuestModelsFromDict(
			p.QuestModels,
		)
	}
	return map[string]interface{}{
		"questGroupModelId": questGroupModelId,
		"questGroupName":    questGroupName,
		"statistics":        statistics,
		"distributions":     distributions,
		"questModels":       questModels,
	}
}

func (p QuestQuestGroupModel) Pointer() *QuestQuestGroupModel {
	return &p
}

func CastQuestQuestGroupModels(data []interface{}) []QuestQuestGroupModel {
	v := make([]QuestQuestGroupModel, 0)
	for _, d := range data {
		v = append(v, NewQuestQuestGroupModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestQuestGroupModelsFromDict(data []QuestQuestGroupModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestNamespaceStatistics struct {
	StartQuest     *int64   `json:"startQuest"`
	EndQuest       *int64   `json:"endQuest"`
	Successful     *int64   `json:"successful"`
	SuccessfulRate *float32 `json:"successfulRate"`
}

func NewQuestNamespaceStatisticsFromJson(data string) QuestNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestNamespaceStatisticsFromDict(dict)
}

func NewQuestNamespaceStatisticsFromDict(data map[string]interface{}) QuestNamespaceStatistics {
	return QuestNamespaceStatistics{
		StartQuest:     core.CastInt64(data["startQuest"]),
		EndQuest:       core.CastInt64(data["endQuest"]),
		Successful:     core.CastInt64(data["successful"]),
		SuccessfulRate: core.CastFloat32(data["successfulRate"]),
	}
}

func (p QuestNamespaceStatistics) ToDict() map[string]interface{} {

	var startQuest *int64
	if p.StartQuest != nil {
		startQuest = p.StartQuest
	}
	var endQuest *int64
	if p.EndQuest != nil {
		endQuest = p.EndQuest
	}
	var successful *int64
	if p.Successful != nil {
		successful = p.Successful
	}
	var successfulRate *float32
	if p.SuccessfulRate != nil {
		successfulRate = p.SuccessfulRate
	}
	return map[string]interface{}{
		"startQuest":     startQuest,
		"endQuest":       endQuest,
		"successful":     successful,
		"successfulRate": successfulRate,
	}
}

func (p QuestNamespaceStatistics) Pointer() *QuestNamespaceStatistics {
	return &p
}

func CastQuestNamespaceStatisticses(data []interface{}) []QuestNamespaceStatistics {
	v := make([]QuestNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewQuestNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestNamespaceStatisticsesFromDict(data []QuestNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestNamespaceQuestDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewQuestNamespaceQuestDistributionStatisticsFromJson(data string) QuestNamespaceQuestDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestNamespaceQuestDistributionStatisticsFromDict(dict)
}

func NewQuestNamespaceQuestDistributionStatisticsFromDict(data map[string]interface{}) QuestNamespaceQuestDistributionStatistics {
	return QuestNamespaceQuestDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p QuestNamespaceQuestDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p QuestNamespaceQuestDistributionStatistics) Pointer() *QuestNamespaceQuestDistributionStatistics {
	return &p
}

func CastQuestNamespaceQuestDistributionStatisticses(data []interface{}) []QuestNamespaceQuestDistributionStatistics {
	v := make([]QuestNamespaceQuestDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewQuestNamespaceQuestDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestNamespaceQuestDistributionStatisticsesFromDict(data []QuestNamespaceQuestDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestNamespaceQuestDistributionSegment struct {
	QuestGroupName *string `json:"questGroupName"`
	Count          *int64  `json:"count"`
}

func NewQuestNamespaceQuestDistributionSegmentFromJson(data string) QuestNamespaceQuestDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestNamespaceQuestDistributionSegmentFromDict(dict)
}

func NewQuestNamespaceQuestDistributionSegmentFromDict(data map[string]interface{}) QuestNamespaceQuestDistributionSegment {
	return QuestNamespaceQuestDistributionSegment{
		QuestGroupName: core.CastString(data["questGroupName"]),
		Count:          core.CastInt64(data["count"]),
	}
}

func (p QuestNamespaceQuestDistributionSegment) ToDict() map[string]interface{} {

	var questGroupName *string
	if p.QuestGroupName != nil {
		questGroupName = p.QuestGroupName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"questGroupName": questGroupName,
		"count":          count,
	}
}

func (p QuestNamespaceQuestDistributionSegment) Pointer() *QuestNamespaceQuestDistributionSegment {
	return &p
}

func CastQuestNamespaceQuestDistributionSegments(data []interface{}) []QuestNamespaceQuestDistributionSegment {
	v := make([]QuestNamespaceQuestDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewQuestNamespaceQuestDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestNamespaceQuestDistributionSegmentsFromDict(data []QuestNamespaceQuestDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestNamespaceQuestDistribution struct {
	Statistics   *QuestNamespaceQuestDistributionStatistics `json:"statistics"`
	Distribution []QuestNamespaceQuestDistributionSegment   `json:"distribution"`
}

func NewQuestNamespaceQuestDistributionFromJson(data string) QuestNamespaceQuestDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestNamespaceQuestDistributionFromDict(dict)
}

func NewQuestNamespaceQuestDistributionFromDict(data map[string]interface{}) QuestNamespaceQuestDistribution {
	return QuestNamespaceQuestDistribution{
		Statistics:   NewQuestNamespaceQuestDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastQuestNamespaceQuestDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p QuestNamespaceQuestDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastQuestNamespaceQuestDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p QuestNamespaceQuestDistribution) Pointer() *QuestNamespaceQuestDistribution {
	return &p
}

func CastQuestNamespaceQuestDistributions(data []interface{}) []QuestNamespaceQuestDistribution {
	v := make([]QuestNamespaceQuestDistribution, 0)
	for _, d := range data {
		v = append(v, NewQuestNamespaceQuestDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestNamespaceQuestDistributionsFromDict(data []QuestNamespaceQuestDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestNamespaceDistributions struct {
	Quest *QuestNamespaceQuestDistribution `json:"quest"`
}

func NewQuestNamespaceDistributionsFromJson(data string) QuestNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestNamespaceDistributionsFromDict(dict)
}

func NewQuestNamespaceDistributionsFromDict(data map[string]interface{}) QuestNamespaceDistributions {
	return QuestNamespaceDistributions{
		Quest: NewQuestNamespaceQuestDistributionFromDict(core.CastMap(data["quest"])).Pointer(),
	}
}

func (p QuestNamespaceDistributions) ToDict() map[string]interface{} {

	var quest map[string]interface{}
	if p.Quest != nil {
		quest = p.Quest.ToDict()
	}
	return map[string]interface{}{
		"quest": quest,
	}
}

func (p QuestNamespaceDistributions) Pointer() *QuestNamespaceDistributions {
	return &p
}

func CastQuestNamespaceDistributionses(data []interface{}) []QuestNamespaceDistributions {
	v := make([]QuestNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewQuestNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestNamespaceDistributionsesFromDict(data []QuestNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestNamespace struct {
	NamespaceId      *string                      `json:"namespaceId"`
	Year             *int32                       `json:"year"`
	Month            *int32                       `json:"month"`
	Day              *int32                       `json:"day"`
	NamespaceName    *string                      `json:"namespaceName"`
	Statistics       *QuestNamespaceStatistics    `json:"statistics"`
	Distributions    *QuestNamespaceDistributions `json:"distributions"`
	QuestGroupModels []QuestQuestGroupModel       `json:"questGroupModels"`
}

func NewQuestNamespaceFromJson(data string) QuestNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestNamespaceFromDict(dict)
}

func NewQuestNamespaceFromDict(data map[string]interface{}) QuestNamespace {
	return QuestNamespace{
		NamespaceId:      core.CastString(data["namespaceId"]),
		Year:             core.CastInt32(data["year"]),
		Month:            core.CastInt32(data["month"]),
		Day:              core.CastInt32(data["day"]),
		NamespaceName:    core.CastString(data["namespaceName"]),
		Statistics:       NewQuestNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions:    NewQuestNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		QuestGroupModels: CastQuestQuestGroupModels(core.CastArray(data["questGroupModels"])),
	}
}

func (p QuestNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var questGroupModels []interface{}
	if p.QuestGroupModels != nil {
		questGroupModels = CastQuestQuestGroupModelsFromDict(
			p.QuestGroupModels,
		)
	}
	return map[string]interface{}{
		"namespaceId":      namespaceId,
		"year":             year,
		"month":            month,
		"day":              day,
		"namespaceName":    namespaceName,
		"statistics":       statistics,
		"distributions":    distributions,
		"questGroupModels": questGroupModels,
	}
}

func (p QuestNamespace) Pointer() *QuestNamespace {
	return &p
}

func CastQuestNamespaces(data []interface{}) []QuestNamespace {
	v := make([]QuestNamespace, 0)
	for _, d := range data {
		v = append(v, NewQuestNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestNamespacesFromDict(data []QuestNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingCategoryModelStatistics struct {
	Put *int64 `json:"put"`
}

func NewRankingCategoryModelStatisticsFromJson(data string) RankingCategoryModelStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingCategoryModelStatisticsFromDict(dict)
}

func NewRankingCategoryModelStatisticsFromDict(data map[string]interface{}) RankingCategoryModelStatistics {
	return RankingCategoryModelStatistics{
		Put: core.CastInt64(data["put"]),
	}
}

func (p RankingCategoryModelStatistics) ToDict() map[string]interface{} {

	var put *int64
	if p.Put != nil {
		put = p.Put
	}
	return map[string]interface{}{
		"put": put,
	}
}

func (p RankingCategoryModelStatistics) Pointer() *RankingCategoryModelStatistics {
	return &p
}

func CastRankingCategoryModelStatisticses(data []interface{}) []RankingCategoryModelStatistics {
	v := make([]RankingCategoryModelStatistics, 0)
	for _, d := range data {
		v = append(v, NewRankingCategoryModelStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingCategoryModelStatisticsesFromDict(data []RankingCategoryModelStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingCategoryModelScoreDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewRankingCategoryModelScoreDistributionStatisticsFromJson(data string) RankingCategoryModelScoreDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingCategoryModelScoreDistributionStatisticsFromDict(dict)
}

func NewRankingCategoryModelScoreDistributionStatisticsFromDict(data map[string]interface{}) RankingCategoryModelScoreDistributionStatistics {
	return RankingCategoryModelScoreDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p RankingCategoryModelScoreDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p RankingCategoryModelScoreDistributionStatistics) Pointer() *RankingCategoryModelScoreDistributionStatistics {
	return &p
}

func CastRankingCategoryModelScoreDistributionStatisticses(data []interface{}) []RankingCategoryModelScoreDistributionStatistics {
	v := make([]RankingCategoryModelScoreDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewRankingCategoryModelScoreDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingCategoryModelScoreDistributionStatisticsesFromDict(data []RankingCategoryModelScoreDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingCategoryModelScoreDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewRankingCategoryModelScoreDistributionSegmentFromJson(data string) RankingCategoryModelScoreDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingCategoryModelScoreDistributionSegmentFromDict(dict)
}

func NewRankingCategoryModelScoreDistributionSegmentFromDict(data map[string]interface{}) RankingCategoryModelScoreDistributionSegment {
	return RankingCategoryModelScoreDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p RankingCategoryModelScoreDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p RankingCategoryModelScoreDistributionSegment) Pointer() *RankingCategoryModelScoreDistributionSegment {
	return &p
}

func CastRankingCategoryModelScoreDistributionSegments(data []interface{}) []RankingCategoryModelScoreDistributionSegment {
	v := make([]RankingCategoryModelScoreDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewRankingCategoryModelScoreDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingCategoryModelScoreDistributionSegmentsFromDict(data []RankingCategoryModelScoreDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingCategoryModelScoreDistribution struct {
	Statistics   *RankingCategoryModelScoreDistributionStatistics `json:"statistics"`
	Distribution []RankingCategoryModelScoreDistributionSegment   `json:"distribution"`
}

func NewRankingCategoryModelScoreDistributionFromJson(data string) RankingCategoryModelScoreDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingCategoryModelScoreDistributionFromDict(dict)
}

func NewRankingCategoryModelScoreDistributionFromDict(data map[string]interface{}) RankingCategoryModelScoreDistribution {
	return RankingCategoryModelScoreDistribution{
		Statistics:   NewRankingCategoryModelScoreDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastRankingCategoryModelScoreDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p RankingCategoryModelScoreDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastRankingCategoryModelScoreDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p RankingCategoryModelScoreDistribution) Pointer() *RankingCategoryModelScoreDistribution {
	return &p
}

func CastRankingCategoryModelScoreDistributions(data []interface{}) []RankingCategoryModelScoreDistribution {
	v := make([]RankingCategoryModelScoreDistribution, 0)
	for _, d := range data {
		v = append(v, NewRankingCategoryModelScoreDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingCategoryModelScoreDistributionsFromDict(data []RankingCategoryModelScoreDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingCategoryModelDistributions struct {
	Score *RankingCategoryModelScoreDistribution `json:"score"`
}

func NewRankingCategoryModelDistributionsFromJson(data string) RankingCategoryModelDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingCategoryModelDistributionsFromDict(dict)
}

func NewRankingCategoryModelDistributionsFromDict(data map[string]interface{}) RankingCategoryModelDistributions {
	return RankingCategoryModelDistributions{
		Score: NewRankingCategoryModelScoreDistributionFromDict(core.CastMap(data["score"])).Pointer(),
	}
}

func (p RankingCategoryModelDistributions) ToDict() map[string]interface{} {

	var score map[string]interface{}
	if p.Score != nil {
		score = p.Score.ToDict()
	}
	return map[string]interface{}{
		"score": score,
	}
}

func (p RankingCategoryModelDistributions) Pointer() *RankingCategoryModelDistributions {
	return &p
}

func CastRankingCategoryModelDistributionses(data []interface{}) []RankingCategoryModelDistributions {
	v := make([]RankingCategoryModelDistributions, 0)
	for _, d := range data {
		v = append(v, NewRankingCategoryModelDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingCategoryModelDistributionsesFromDict(data []RankingCategoryModelDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingCategoryModel struct {
	CategoryModelId *string                            `json:"categoryModelId"`
	CategoryName    *string                            `json:"categoryName"`
	Statistics      *RankingCategoryModelStatistics    `json:"statistics"`
	Distributions   *RankingCategoryModelDistributions `json:"distributions"`
}

func NewRankingCategoryModelFromJson(data string) RankingCategoryModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingCategoryModelFromDict(dict)
}

func NewRankingCategoryModelFromDict(data map[string]interface{}) RankingCategoryModel {
	return RankingCategoryModel{
		CategoryModelId: core.CastString(data["categoryModelId"]),
		CategoryName:    core.CastString(data["categoryName"]),
		Statistics:      NewRankingCategoryModelStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions:   NewRankingCategoryModelDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p RankingCategoryModel) ToDict() map[string]interface{} {

	var categoryModelId *string
	if p.CategoryModelId != nil {
		categoryModelId = p.CategoryModelId
	}
	var categoryName *string
	if p.CategoryName != nil {
		categoryName = p.CategoryName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"categoryModelId": categoryModelId,
		"categoryName":    categoryName,
		"statistics":      statistics,
		"distributions":   distributions,
	}
}

func (p RankingCategoryModel) Pointer() *RankingCategoryModel {
	return &p
}

func CastRankingCategoryModels(data []interface{}) []RankingCategoryModel {
	v := make([]RankingCategoryModel, 0)
	for _, d := range data {
		v = append(v, NewRankingCategoryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingCategoryModelsFromDict(data []RankingCategoryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingNamespaceStatistics struct {
	Put *int64 `json:"put"`
}

func NewRankingNamespaceStatisticsFromJson(data string) RankingNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingNamespaceStatisticsFromDict(dict)
}

func NewRankingNamespaceStatisticsFromDict(data map[string]interface{}) RankingNamespaceStatistics {
	return RankingNamespaceStatistics{
		Put: core.CastInt64(data["put"]),
	}
}

func (p RankingNamespaceStatistics) ToDict() map[string]interface{} {

	var put *int64
	if p.Put != nil {
		put = p.Put
	}
	return map[string]interface{}{
		"put": put,
	}
}

func (p RankingNamespaceStatistics) Pointer() *RankingNamespaceStatistics {
	return &p
}

func CastRankingNamespaceStatisticses(data []interface{}) []RankingNamespaceStatistics {
	v := make([]RankingNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewRankingNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingNamespaceStatisticsesFromDict(data []RankingNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingNamespacePutDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewRankingNamespacePutDistributionStatisticsFromJson(data string) RankingNamespacePutDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingNamespacePutDistributionStatisticsFromDict(dict)
}

func NewRankingNamespacePutDistributionStatisticsFromDict(data map[string]interface{}) RankingNamespacePutDistributionStatistics {
	return RankingNamespacePutDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p RankingNamespacePutDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p RankingNamespacePutDistributionStatistics) Pointer() *RankingNamespacePutDistributionStatistics {
	return &p
}

func CastRankingNamespacePutDistributionStatisticses(data []interface{}) []RankingNamespacePutDistributionStatistics {
	v := make([]RankingNamespacePutDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewRankingNamespacePutDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingNamespacePutDistributionStatisticsesFromDict(data []RankingNamespacePutDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingNamespacePutDistributionSegment struct {
	CategoryName *string `json:"categoryName"`
	Count        *int64  `json:"count"`
}

func NewRankingNamespacePutDistributionSegmentFromJson(data string) RankingNamespacePutDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingNamespacePutDistributionSegmentFromDict(dict)
}

func NewRankingNamespacePutDistributionSegmentFromDict(data map[string]interface{}) RankingNamespacePutDistributionSegment {
	return RankingNamespacePutDistributionSegment{
		CategoryName: core.CastString(data["categoryName"]),
		Count:        core.CastInt64(data["count"]),
	}
}

func (p RankingNamespacePutDistributionSegment) ToDict() map[string]interface{} {

	var categoryName *string
	if p.CategoryName != nil {
		categoryName = p.CategoryName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"categoryName": categoryName,
		"count":        count,
	}
}

func (p RankingNamespacePutDistributionSegment) Pointer() *RankingNamespacePutDistributionSegment {
	return &p
}

func CastRankingNamespacePutDistributionSegments(data []interface{}) []RankingNamespacePutDistributionSegment {
	v := make([]RankingNamespacePutDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewRankingNamespacePutDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingNamespacePutDistributionSegmentsFromDict(data []RankingNamespacePutDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingNamespacePutDistribution struct {
	Statistics   *RankingNamespacePutDistributionStatistics `json:"statistics"`
	Distribution []RankingNamespacePutDistributionSegment   `json:"distribution"`
}

func NewRankingNamespacePutDistributionFromJson(data string) RankingNamespacePutDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingNamespacePutDistributionFromDict(dict)
}

func NewRankingNamespacePutDistributionFromDict(data map[string]interface{}) RankingNamespacePutDistribution {
	return RankingNamespacePutDistribution{
		Statistics:   NewRankingNamespacePutDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastRankingNamespacePutDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p RankingNamespacePutDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastRankingNamespacePutDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p RankingNamespacePutDistribution) Pointer() *RankingNamespacePutDistribution {
	return &p
}

func CastRankingNamespacePutDistributions(data []interface{}) []RankingNamespacePutDistribution {
	v := make([]RankingNamespacePutDistribution, 0)
	for _, d := range data {
		v = append(v, NewRankingNamespacePutDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingNamespacePutDistributionsFromDict(data []RankingNamespacePutDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingNamespaceDistributions struct {
	Put *RankingNamespacePutDistribution `json:"put"`
}

func NewRankingNamespaceDistributionsFromJson(data string) RankingNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingNamespaceDistributionsFromDict(dict)
}

func NewRankingNamespaceDistributionsFromDict(data map[string]interface{}) RankingNamespaceDistributions {
	return RankingNamespaceDistributions{
		Put: NewRankingNamespacePutDistributionFromDict(core.CastMap(data["put"])).Pointer(),
	}
}

func (p RankingNamespaceDistributions) ToDict() map[string]interface{} {

	var put map[string]interface{}
	if p.Put != nil {
		put = p.Put.ToDict()
	}
	return map[string]interface{}{
		"put": put,
	}
}

func (p RankingNamespaceDistributions) Pointer() *RankingNamespaceDistributions {
	return &p
}

func CastRankingNamespaceDistributionses(data []interface{}) []RankingNamespaceDistributions {
	v := make([]RankingNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewRankingNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingNamespaceDistributionsesFromDict(data []RankingNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingNamespace struct {
	NamespaceId    *string                        `json:"namespaceId"`
	Year           *int32                         `json:"year"`
	Month          *int32                         `json:"month"`
	Day            *int32                         `json:"day"`
	NamespaceName  *string                        `json:"namespaceName"`
	Statistics     *RankingNamespaceStatistics    `json:"statistics"`
	Distributions  *RankingNamespaceDistributions `json:"distributions"`
	CategoryModels []RankingCategoryModel         `json:"categoryModels"`
}

func NewRankingNamespaceFromJson(data string) RankingNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRankingNamespaceFromDict(dict)
}

func NewRankingNamespaceFromDict(data map[string]interface{}) RankingNamespace {
	return RankingNamespace{
		NamespaceId:    core.CastString(data["namespaceId"]),
		Year:           core.CastInt32(data["year"]),
		Month:          core.CastInt32(data["month"]),
		Day:            core.CastInt32(data["day"]),
		NamespaceName:  core.CastString(data["namespaceName"]),
		Statistics:     NewRankingNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions:  NewRankingNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		CategoryModels: CastRankingCategoryModels(core.CastArray(data["categoryModels"])),
	}
}

func (p RankingNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var categoryModels []interface{}
	if p.CategoryModels != nil {
		categoryModels = CastRankingCategoryModelsFromDict(
			p.CategoryModels,
		)
	}
	return map[string]interface{}{
		"namespaceId":    namespaceId,
		"year":           year,
		"month":          month,
		"day":            day,
		"namespaceName":  namespaceName,
		"statistics":     statistics,
		"distributions":  distributions,
		"categoryModels": categoryModels,
	}
}

func (p RankingNamespace) Pointer() *RankingNamespace {
	return &p
}

func CastRankingNamespaces(data []interface{}) []RankingNamespace {
	v := make([]RankingNamespace, 0)
	for _, d := range data {
		v = append(v, NewRankingNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingNamespacesFromDict(data []RankingNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseDisplayItemStatistics struct {
	Buy *int64 `json:"buy"`
}

func NewShowcaseDisplayItemStatisticsFromJson(data string) ShowcaseDisplayItemStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseDisplayItemStatisticsFromDict(dict)
}

func NewShowcaseDisplayItemStatisticsFromDict(data map[string]interface{}) ShowcaseDisplayItemStatistics {
	return ShowcaseDisplayItemStatistics{
		Buy: core.CastInt64(data["buy"]),
	}
}

func (p ShowcaseDisplayItemStatistics) ToDict() map[string]interface{} {

	var buy *int64
	if p.Buy != nil {
		buy = p.Buy
	}
	return map[string]interface{}{
		"buy": buy,
	}
}

func (p ShowcaseDisplayItemStatistics) Pointer() *ShowcaseDisplayItemStatistics {
	return &p
}

func CastShowcaseDisplayItemStatisticses(data []interface{}) []ShowcaseDisplayItemStatistics {
	v := make([]ShowcaseDisplayItemStatistics, 0)
	for _, d := range data {
		v = append(v, NewShowcaseDisplayItemStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseDisplayItemStatisticsesFromDict(data []ShowcaseDisplayItemStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseDisplayItemQuantityDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewShowcaseDisplayItemQuantityDistributionStatisticsFromJson(data string) ShowcaseDisplayItemQuantityDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseDisplayItemQuantityDistributionStatisticsFromDict(dict)
}

func NewShowcaseDisplayItemQuantityDistributionStatisticsFromDict(data map[string]interface{}) ShowcaseDisplayItemQuantityDistributionStatistics {
	return ShowcaseDisplayItemQuantityDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ShowcaseDisplayItemQuantityDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ShowcaseDisplayItemQuantityDistributionStatistics) Pointer() *ShowcaseDisplayItemQuantityDistributionStatistics {
	return &p
}

func CastShowcaseDisplayItemQuantityDistributionStatisticses(data []interface{}) []ShowcaseDisplayItemQuantityDistributionStatistics {
	v := make([]ShowcaseDisplayItemQuantityDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewShowcaseDisplayItemQuantityDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseDisplayItemQuantityDistributionStatisticsesFromDict(data []ShowcaseDisplayItemQuantityDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseDisplayItemQuantityDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewShowcaseDisplayItemQuantityDistributionSegmentFromJson(data string) ShowcaseDisplayItemQuantityDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseDisplayItemQuantityDistributionSegmentFromDict(dict)
}

func NewShowcaseDisplayItemQuantityDistributionSegmentFromDict(data map[string]interface{}) ShowcaseDisplayItemQuantityDistributionSegment {
	return ShowcaseDisplayItemQuantityDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p ShowcaseDisplayItemQuantityDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p ShowcaseDisplayItemQuantityDistributionSegment) Pointer() *ShowcaseDisplayItemQuantityDistributionSegment {
	return &p
}

func CastShowcaseDisplayItemQuantityDistributionSegments(data []interface{}) []ShowcaseDisplayItemQuantityDistributionSegment {
	v := make([]ShowcaseDisplayItemQuantityDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewShowcaseDisplayItemQuantityDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseDisplayItemQuantityDistributionSegmentsFromDict(data []ShowcaseDisplayItemQuantityDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseDisplayItemQuantityDistribution struct {
	Statistics   *ShowcaseDisplayItemQuantityDistributionStatistics `json:"statistics"`
	Distribution []ShowcaseDisplayItemQuantityDistributionSegment   `json:"distribution"`
}

func NewShowcaseDisplayItemQuantityDistributionFromJson(data string) ShowcaseDisplayItemQuantityDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseDisplayItemQuantityDistributionFromDict(dict)
}

func NewShowcaseDisplayItemQuantityDistributionFromDict(data map[string]interface{}) ShowcaseDisplayItemQuantityDistribution {
	return ShowcaseDisplayItemQuantityDistribution{
		Statistics:   NewShowcaseDisplayItemQuantityDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastShowcaseDisplayItemQuantityDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ShowcaseDisplayItemQuantityDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastShowcaseDisplayItemQuantityDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ShowcaseDisplayItemQuantityDistribution) Pointer() *ShowcaseDisplayItemQuantityDistribution {
	return &p
}

func CastShowcaseDisplayItemQuantityDistributions(data []interface{}) []ShowcaseDisplayItemQuantityDistribution {
	v := make([]ShowcaseDisplayItemQuantityDistribution, 0)
	for _, d := range data {
		v = append(v, NewShowcaseDisplayItemQuantityDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseDisplayItemQuantityDistributionsFromDict(data []ShowcaseDisplayItemQuantityDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseDisplayItemDistributions struct {
	Quantity *ShowcaseDisplayItemQuantityDistribution `json:"quantity"`
}

func NewShowcaseDisplayItemDistributionsFromJson(data string) ShowcaseDisplayItemDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseDisplayItemDistributionsFromDict(dict)
}

func NewShowcaseDisplayItemDistributionsFromDict(data map[string]interface{}) ShowcaseDisplayItemDistributions {
	return ShowcaseDisplayItemDistributions{
		Quantity: NewShowcaseDisplayItemQuantityDistributionFromDict(core.CastMap(data["quantity"])).Pointer(),
	}
}

func (p ShowcaseDisplayItemDistributions) ToDict() map[string]interface{} {

	var quantity map[string]interface{}
	if p.Quantity != nil {
		quantity = p.Quantity.ToDict()
	}
	return map[string]interface{}{
		"quantity": quantity,
	}
}

func (p ShowcaseDisplayItemDistributions) Pointer() *ShowcaseDisplayItemDistributions {
	return &p
}

func CastShowcaseDisplayItemDistributionses(data []interface{}) []ShowcaseDisplayItemDistributions {
	v := make([]ShowcaseDisplayItemDistributions, 0)
	for _, d := range data {
		v = append(v, NewShowcaseDisplayItemDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseDisplayItemDistributionsesFromDict(data []ShowcaseDisplayItemDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseDisplayItem struct {
	DisplayItemId *string                           `json:"displayItemId"`
	Statistics    *ShowcaseDisplayItemStatistics    `json:"statistics"`
	Distributions *ShowcaseDisplayItemDistributions `json:"distributions"`
}

func NewShowcaseDisplayItemFromJson(data string) ShowcaseDisplayItem {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseDisplayItemFromDict(dict)
}

func NewShowcaseDisplayItemFromDict(data map[string]interface{}) ShowcaseDisplayItem {
	return ShowcaseDisplayItem{
		DisplayItemId: core.CastString(data["displayItemId"]),
		Statistics:    NewShowcaseDisplayItemStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewShowcaseDisplayItemDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p ShowcaseDisplayItem) ToDict() map[string]interface{} {

	var displayItemId *string
	if p.DisplayItemId != nil {
		displayItemId = p.DisplayItemId
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"displayItemId": displayItemId,
		"statistics":    statistics,
		"distributions": distributions,
	}
}

func (p ShowcaseDisplayItem) Pointer() *ShowcaseDisplayItem {
	return &p
}

func CastShowcaseDisplayItems(data []interface{}) []ShowcaseDisplayItem {
	v := make([]ShowcaseDisplayItem, 0)
	for _, d := range data {
		v = append(v, NewShowcaseDisplayItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseDisplayItemsFromDict(data []ShowcaseDisplayItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseShowcaseStatistics struct {
	Buy *int64 `json:"buy"`
}

func NewShowcaseShowcaseStatisticsFromJson(data string) ShowcaseShowcaseStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseShowcaseStatisticsFromDict(dict)
}

func NewShowcaseShowcaseStatisticsFromDict(data map[string]interface{}) ShowcaseShowcaseStatistics {
	return ShowcaseShowcaseStatistics{
		Buy: core.CastInt64(data["buy"]),
	}
}

func (p ShowcaseShowcaseStatistics) ToDict() map[string]interface{} {

	var buy *int64
	if p.Buy != nil {
		buy = p.Buy
	}
	return map[string]interface{}{
		"buy": buy,
	}
}

func (p ShowcaseShowcaseStatistics) Pointer() *ShowcaseShowcaseStatistics {
	return &p
}

func CastShowcaseShowcaseStatisticses(data []interface{}) []ShowcaseShowcaseStatistics {
	v := make([]ShowcaseShowcaseStatistics, 0)
	for _, d := range data {
		v = append(v, NewShowcaseShowcaseStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseShowcaseStatisticsesFromDict(data []ShowcaseShowcaseStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseShowcaseBuyDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewShowcaseShowcaseBuyDistributionStatisticsFromJson(data string) ShowcaseShowcaseBuyDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseShowcaseBuyDistributionStatisticsFromDict(dict)
}

func NewShowcaseShowcaseBuyDistributionStatisticsFromDict(data map[string]interface{}) ShowcaseShowcaseBuyDistributionStatistics {
	return ShowcaseShowcaseBuyDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ShowcaseShowcaseBuyDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ShowcaseShowcaseBuyDistributionStatistics) Pointer() *ShowcaseShowcaseBuyDistributionStatistics {
	return &p
}

func CastShowcaseShowcaseBuyDistributionStatisticses(data []interface{}) []ShowcaseShowcaseBuyDistributionStatistics {
	v := make([]ShowcaseShowcaseBuyDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewShowcaseShowcaseBuyDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseShowcaseBuyDistributionStatisticsesFromDict(data []ShowcaseShowcaseBuyDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseShowcaseBuyDistributionSegment struct {
	DisplayItemId *string `json:"displayItemId"`
	Count         *int64  `json:"count"`
}

func NewShowcaseShowcaseBuyDistributionSegmentFromJson(data string) ShowcaseShowcaseBuyDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseShowcaseBuyDistributionSegmentFromDict(dict)
}

func NewShowcaseShowcaseBuyDistributionSegmentFromDict(data map[string]interface{}) ShowcaseShowcaseBuyDistributionSegment {
	return ShowcaseShowcaseBuyDistributionSegment{
		DisplayItemId: core.CastString(data["displayItemId"]),
		Count:         core.CastInt64(data["count"]),
	}
}

func (p ShowcaseShowcaseBuyDistributionSegment) ToDict() map[string]interface{} {

	var displayItemId *string
	if p.DisplayItemId != nil {
		displayItemId = p.DisplayItemId
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"displayItemId": displayItemId,
		"count":         count,
	}
}

func (p ShowcaseShowcaseBuyDistributionSegment) Pointer() *ShowcaseShowcaseBuyDistributionSegment {
	return &p
}

func CastShowcaseShowcaseBuyDistributionSegments(data []interface{}) []ShowcaseShowcaseBuyDistributionSegment {
	v := make([]ShowcaseShowcaseBuyDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewShowcaseShowcaseBuyDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseShowcaseBuyDistributionSegmentsFromDict(data []ShowcaseShowcaseBuyDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseShowcaseBuyDistribution struct {
	Statistics   *ShowcaseShowcaseBuyDistributionStatistics `json:"statistics"`
	Distribution []ShowcaseShowcaseBuyDistributionSegment   `json:"distribution"`
}

func NewShowcaseShowcaseBuyDistributionFromJson(data string) ShowcaseShowcaseBuyDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseShowcaseBuyDistributionFromDict(dict)
}

func NewShowcaseShowcaseBuyDistributionFromDict(data map[string]interface{}) ShowcaseShowcaseBuyDistribution {
	return ShowcaseShowcaseBuyDistribution{
		Statistics:   NewShowcaseShowcaseBuyDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastShowcaseShowcaseBuyDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ShowcaseShowcaseBuyDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastShowcaseShowcaseBuyDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ShowcaseShowcaseBuyDistribution) Pointer() *ShowcaseShowcaseBuyDistribution {
	return &p
}

func CastShowcaseShowcaseBuyDistributions(data []interface{}) []ShowcaseShowcaseBuyDistribution {
	v := make([]ShowcaseShowcaseBuyDistribution, 0)
	for _, d := range data {
		v = append(v, NewShowcaseShowcaseBuyDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseShowcaseBuyDistributionsFromDict(data []ShowcaseShowcaseBuyDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseShowcaseDistributions struct {
	Buy *ShowcaseShowcaseBuyDistribution `json:"buy"`
}

func NewShowcaseShowcaseDistributionsFromJson(data string) ShowcaseShowcaseDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseShowcaseDistributionsFromDict(dict)
}

func NewShowcaseShowcaseDistributionsFromDict(data map[string]interface{}) ShowcaseShowcaseDistributions {
	return ShowcaseShowcaseDistributions{
		Buy: NewShowcaseShowcaseBuyDistributionFromDict(core.CastMap(data["buy"])).Pointer(),
	}
}

func (p ShowcaseShowcaseDistributions) ToDict() map[string]interface{} {

	var buy map[string]interface{}
	if p.Buy != nil {
		buy = p.Buy.ToDict()
	}
	return map[string]interface{}{
		"buy": buy,
	}
}

func (p ShowcaseShowcaseDistributions) Pointer() *ShowcaseShowcaseDistributions {
	return &p
}

func CastShowcaseShowcaseDistributionses(data []interface{}) []ShowcaseShowcaseDistributions {
	v := make([]ShowcaseShowcaseDistributions, 0)
	for _, d := range data {
		v = append(v, NewShowcaseShowcaseDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseShowcaseDistributionsesFromDict(data []ShowcaseShowcaseDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseShowcase struct {
	ShowcaseId    *string                        `json:"showcaseId"`
	ShowcaseName  *string                        `json:"showcaseName"`
	Statistics    *ShowcaseShowcaseStatistics    `json:"statistics"`
	Distributions *ShowcaseShowcaseDistributions `json:"distributions"`
	DisplayItems  []ShowcaseDisplayItem          `json:"displayItems"`
}

func NewShowcaseShowcaseFromJson(data string) ShowcaseShowcase {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseShowcaseFromDict(dict)
}

func NewShowcaseShowcaseFromDict(data map[string]interface{}) ShowcaseShowcase {
	return ShowcaseShowcase{
		ShowcaseId:    core.CastString(data["showcaseId"]),
		ShowcaseName:  core.CastString(data["showcaseName"]),
		Statistics:    NewShowcaseShowcaseStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewShowcaseShowcaseDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		DisplayItems:  CastShowcaseDisplayItems(core.CastArray(data["displayItems"])),
	}
}

func (p ShowcaseShowcase) ToDict() map[string]interface{} {

	var showcaseId *string
	if p.ShowcaseId != nil {
		showcaseId = p.ShowcaseId
	}
	var showcaseName *string
	if p.ShowcaseName != nil {
		showcaseName = p.ShowcaseName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var displayItems []interface{}
	if p.DisplayItems != nil {
		displayItems = CastShowcaseDisplayItemsFromDict(
			p.DisplayItems,
		)
	}
	return map[string]interface{}{
		"showcaseId":    showcaseId,
		"showcaseName":  showcaseName,
		"statistics":    statistics,
		"distributions": distributions,
		"displayItems":  displayItems,
	}
}

func (p ShowcaseShowcase) Pointer() *ShowcaseShowcase {
	return &p
}

func CastShowcaseShowcases(data []interface{}) []ShowcaseShowcase {
	v := make([]ShowcaseShowcase, 0)
	for _, d := range data {
		v = append(v, NewShowcaseShowcaseFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseShowcasesFromDict(data []ShowcaseShowcase) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseNamespaceStatistics struct {
	Buy *int64 `json:"buy"`
}

func NewShowcaseNamespaceStatisticsFromJson(data string) ShowcaseNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseNamespaceStatisticsFromDict(dict)
}

func NewShowcaseNamespaceStatisticsFromDict(data map[string]interface{}) ShowcaseNamespaceStatistics {
	return ShowcaseNamespaceStatistics{
		Buy: core.CastInt64(data["buy"]),
	}
}

func (p ShowcaseNamespaceStatistics) ToDict() map[string]interface{} {

	var buy *int64
	if p.Buy != nil {
		buy = p.Buy
	}
	return map[string]interface{}{
		"buy": buy,
	}
}

func (p ShowcaseNamespaceStatistics) Pointer() *ShowcaseNamespaceStatistics {
	return &p
}

func CastShowcaseNamespaceStatisticses(data []interface{}) []ShowcaseNamespaceStatistics {
	v := make([]ShowcaseNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewShowcaseNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseNamespaceStatisticsesFromDict(data []ShowcaseNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseNamespaceBuyDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewShowcaseNamespaceBuyDistributionStatisticsFromJson(data string) ShowcaseNamespaceBuyDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseNamespaceBuyDistributionStatisticsFromDict(dict)
}

func NewShowcaseNamespaceBuyDistributionStatisticsFromDict(data map[string]interface{}) ShowcaseNamespaceBuyDistributionStatistics {
	return ShowcaseNamespaceBuyDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p ShowcaseNamespaceBuyDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p ShowcaseNamespaceBuyDistributionStatistics) Pointer() *ShowcaseNamespaceBuyDistributionStatistics {
	return &p
}

func CastShowcaseNamespaceBuyDistributionStatisticses(data []interface{}) []ShowcaseNamespaceBuyDistributionStatistics {
	v := make([]ShowcaseNamespaceBuyDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewShowcaseNamespaceBuyDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseNamespaceBuyDistributionStatisticsesFromDict(data []ShowcaseNamespaceBuyDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseNamespaceBuyDistributionSegment struct {
	ShowcaseName *string `json:"showcaseName"`
	Count        *int64  `json:"count"`
}

func NewShowcaseNamespaceBuyDistributionSegmentFromJson(data string) ShowcaseNamespaceBuyDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseNamespaceBuyDistributionSegmentFromDict(dict)
}

func NewShowcaseNamespaceBuyDistributionSegmentFromDict(data map[string]interface{}) ShowcaseNamespaceBuyDistributionSegment {
	return ShowcaseNamespaceBuyDistributionSegment{
		ShowcaseName: core.CastString(data["showcaseName"]),
		Count:        core.CastInt64(data["count"]),
	}
}

func (p ShowcaseNamespaceBuyDistributionSegment) ToDict() map[string]interface{} {

	var showcaseName *string
	if p.ShowcaseName != nil {
		showcaseName = p.ShowcaseName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"showcaseName": showcaseName,
		"count":        count,
	}
}

func (p ShowcaseNamespaceBuyDistributionSegment) Pointer() *ShowcaseNamespaceBuyDistributionSegment {
	return &p
}

func CastShowcaseNamespaceBuyDistributionSegments(data []interface{}) []ShowcaseNamespaceBuyDistributionSegment {
	v := make([]ShowcaseNamespaceBuyDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewShowcaseNamespaceBuyDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseNamespaceBuyDistributionSegmentsFromDict(data []ShowcaseNamespaceBuyDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseNamespaceBuyDistribution struct {
	Statistics   *ShowcaseNamespaceBuyDistributionStatistics `json:"statistics"`
	Distribution []ShowcaseNamespaceBuyDistributionSegment   `json:"distribution"`
}

func NewShowcaseNamespaceBuyDistributionFromJson(data string) ShowcaseNamespaceBuyDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseNamespaceBuyDistributionFromDict(dict)
}

func NewShowcaseNamespaceBuyDistributionFromDict(data map[string]interface{}) ShowcaseNamespaceBuyDistribution {
	return ShowcaseNamespaceBuyDistribution{
		Statistics:   NewShowcaseNamespaceBuyDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastShowcaseNamespaceBuyDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p ShowcaseNamespaceBuyDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastShowcaseNamespaceBuyDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p ShowcaseNamespaceBuyDistribution) Pointer() *ShowcaseNamespaceBuyDistribution {
	return &p
}

func CastShowcaseNamespaceBuyDistributions(data []interface{}) []ShowcaseNamespaceBuyDistribution {
	v := make([]ShowcaseNamespaceBuyDistribution, 0)
	for _, d := range data {
		v = append(v, NewShowcaseNamespaceBuyDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseNamespaceBuyDistributionsFromDict(data []ShowcaseNamespaceBuyDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseNamespaceDistributions struct {
	Buy *ShowcaseNamespaceBuyDistribution `json:"buy"`
}

func NewShowcaseNamespaceDistributionsFromJson(data string) ShowcaseNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseNamespaceDistributionsFromDict(dict)
}

func NewShowcaseNamespaceDistributionsFromDict(data map[string]interface{}) ShowcaseNamespaceDistributions {
	return ShowcaseNamespaceDistributions{
		Buy: NewShowcaseNamespaceBuyDistributionFromDict(core.CastMap(data["buy"])).Pointer(),
	}
}

func (p ShowcaseNamespaceDistributions) ToDict() map[string]interface{} {

	var buy map[string]interface{}
	if p.Buy != nil {
		buy = p.Buy.ToDict()
	}
	return map[string]interface{}{
		"buy": buy,
	}
}

func (p ShowcaseNamespaceDistributions) Pointer() *ShowcaseNamespaceDistributions {
	return &p
}

func CastShowcaseNamespaceDistributionses(data []interface{}) []ShowcaseNamespaceDistributions {
	v := make([]ShowcaseNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewShowcaseNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseNamespaceDistributionsesFromDict(data []ShowcaseNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseNamespace struct {
	NamespaceId   *string                         `json:"namespaceId"`
	Year          *int32                          `json:"year"`
	Month         *int32                          `json:"month"`
	Day           *int32                          `json:"day"`
	NamespaceName *string                         `json:"namespaceName"`
	Statistics    *ShowcaseNamespaceStatistics    `json:"statistics"`
	Distributions *ShowcaseNamespaceDistributions `json:"distributions"`
	Showcases     []ShowcaseShowcase              `json:"showcases"`
}

func NewShowcaseNamespaceFromJson(data string) ShowcaseNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseNamespaceFromDict(dict)
}

func NewShowcaseNamespaceFromDict(data map[string]interface{}) ShowcaseNamespace {
	return ShowcaseNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewShowcaseNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewShowcaseNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		Showcases:     CastShowcaseShowcases(core.CastArray(data["showcases"])),
	}
}

func (p ShowcaseNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var showcases []interface{}
	if p.Showcases != nil {
		showcases = CastShowcaseShowcasesFromDict(
			p.Showcases,
		)
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
		"showcases":     showcases,
	}
}

func (p ShowcaseNamespace) Pointer() *ShowcaseNamespace {
	return &p
}

func CastShowcaseNamespaces(data []interface{}) []ShowcaseNamespace {
	v := make([]ShowcaseNamespace, 0)
	for _, d := range data {
		v = append(v, NewShowcaseNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseNamespacesFromDict(data []ShowcaseNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaStaminaModelStatistics struct {
	Consume       *int64 `json:"consume"`
	ConsumeAmount *int64 `json:"consumeAmount"`
	Recover       *int64 `json:"recover"`
	RecoverAmount *int64 `json:"recoverAmount"`
}

func NewStaminaStaminaModelStatisticsFromJson(data string) StaminaStaminaModelStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaStaminaModelStatisticsFromDict(dict)
}

func NewStaminaStaminaModelStatisticsFromDict(data map[string]interface{}) StaminaStaminaModelStatistics {
	return StaminaStaminaModelStatistics{
		Consume:       core.CastInt64(data["consume"]),
		ConsumeAmount: core.CastInt64(data["consumeAmount"]),
		Recover:       core.CastInt64(data["recover"]),
		RecoverAmount: core.CastInt64(data["recoverAmount"]),
	}
}

func (p StaminaStaminaModelStatistics) ToDict() map[string]interface{} {

	var consume *int64
	if p.Consume != nil {
		consume = p.Consume
	}
	var consumeAmount *int64
	if p.ConsumeAmount != nil {
		consumeAmount = p.ConsumeAmount
	}
	var recover *int64
	if p.Recover != nil {
		recover = p.Recover
	}
	var recoverAmount *int64
	if p.RecoverAmount != nil {
		recoverAmount = p.RecoverAmount
	}
	return map[string]interface{}{
		"consume":       consume,
		"consumeAmount": consumeAmount,
		"recover":       recover,
		"recoverAmount": recoverAmount,
	}
}

func (p StaminaStaminaModelStatistics) Pointer() *StaminaStaminaModelStatistics {
	return &p
}

func CastStaminaStaminaModelStatisticses(data []interface{}) []StaminaStaminaModelStatistics {
	v := make([]StaminaStaminaModelStatistics, 0)
	for _, d := range data {
		v = append(v, NewStaminaStaminaModelStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaStaminaModelStatisticsesFromDict(data []StaminaStaminaModelStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaStaminaModelConsumeDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewStaminaStaminaModelConsumeDistributionStatisticsFromJson(data string) StaminaStaminaModelConsumeDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaStaminaModelConsumeDistributionStatisticsFromDict(dict)
}

func NewStaminaStaminaModelConsumeDistributionStatisticsFromDict(data map[string]interface{}) StaminaStaminaModelConsumeDistributionStatistics {
	return StaminaStaminaModelConsumeDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p StaminaStaminaModelConsumeDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p StaminaStaminaModelConsumeDistributionStatistics) Pointer() *StaminaStaminaModelConsumeDistributionStatistics {
	return &p
}

func CastStaminaStaminaModelConsumeDistributionStatisticses(data []interface{}) []StaminaStaminaModelConsumeDistributionStatistics {
	v := make([]StaminaStaminaModelConsumeDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewStaminaStaminaModelConsumeDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaStaminaModelConsumeDistributionStatisticsesFromDict(data []StaminaStaminaModelConsumeDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaStaminaModelConsumeDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewStaminaStaminaModelConsumeDistributionSegmentFromJson(data string) StaminaStaminaModelConsumeDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaStaminaModelConsumeDistributionSegmentFromDict(dict)
}

func NewStaminaStaminaModelConsumeDistributionSegmentFromDict(data map[string]interface{}) StaminaStaminaModelConsumeDistributionSegment {
	return StaminaStaminaModelConsumeDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p StaminaStaminaModelConsumeDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p StaminaStaminaModelConsumeDistributionSegment) Pointer() *StaminaStaminaModelConsumeDistributionSegment {
	return &p
}

func CastStaminaStaminaModelConsumeDistributionSegments(data []interface{}) []StaminaStaminaModelConsumeDistributionSegment {
	v := make([]StaminaStaminaModelConsumeDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewStaminaStaminaModelConsumeDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaStaminaModelConsumeDistributionSegmentsFromDict(data []StaminaStaminaModelConsumeDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaStaminaModelConsumeDistribution struct {
	Statistics   *StaminaStaminaModelConsumeDistributionStatistics `json:"statistics"`
	Distribution []StaminaStaminaModelConsumeDistributionSegment   `json:"distribution"`
}

func NewStaminaStaminaModelConsumeDistributionFromJson(data string) StaminaStaminaModelConsumeDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaStaminaModelConsumeDistributionFromDict(dict)
}

func NewStaminaStaminaModelConsumeDistributionFromDict(data map[string]interface{}) StaminaStaminaModelConsumeDistribution {
	return StaminaStaminaModelConsumeDistribution{
		Statistics:   NewStaminaStaminaModelConsumeDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastStaminaStaminaModelConsumeDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p StaminaStaminaModelConsumeDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastStaminaStaminaModelConsumeDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p StaminaStaminaModelConsumeDistribution) Pointer() *StaminaStaminaModelConsumeDistribution {
	return &p
}

func CastStaminaStaminaModelConsumeDistributions(data []interface{}) []StaminaStaminaModelConsumeDistribution {
	v := make([]StaminaStaminaModelConsumeDistribution, 0)
	for _, d := range data {
		v = append(v, NewStaminaStaminaModelConsumeDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaStaminaModelConsumeDistributionsFromDict(data []StaminaStaminaModelConsumeDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaStaminaModelRecoverDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewStaminaStaminaModelRecoverDistributionStatisticsFromJson(data string) StaminaStaminaModelRecoverDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaStaminaModelRecoverDistributionStatisticsFromDict(dict)
}

func NewStaminaStaminaModelRecoverDistributionStatisticsFromDict(data map[string]interface{}) StaminaStaminaModelRecoverDistributionStatistics {
	return StaminaStaminaModelRecoverDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p StaminaStaminaModelRecoverDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p StaminaStaminaModelRecoverDistributionStatistics) Pointer() *StaminaStaminaModelRecoverDistributionStatistics {
	return &p
}

func CastStaminaStaminaModelRecoverDistributionStatisticses(data []interface{}) []StaminaStaminaModelRecoverDistributionStatistics {
	v := make([]StaminaStaminaModelRecoverDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewStaminaStaminaModelRecoverDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaStaminaModelRecoverDistributionStatisticsesFromDict(data []StaminaStaminaModelRecoverDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaStaminaModelRecoverDistributionSegment struct {
	Min   *int64 `json:"min"`
	Max   *int64 `json:"max"`
	Count *int64 `json:"count"`
}

func NewStaminaStaminaModelRecoverDistributionSegmentFromJson(data string) StaminaStaminaModelRecoverDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaStaminaModelRecoverDistributionSegmentFromDict(dict)
}

func NewStaminaStaminaModelRecoverDistributionSegmentFromDict(data map[string]interface{}) StaminaStaminaModelRecoverDistributionSegment {
	return StaminaStaminaModelRecoverDistributionSegment{
		Min:   core.CastInt64(data["min"]),
		Max:   core.CastInt64(data["max"]),
		Count: core.CastInt64(data["count"]),
	}
}

func (p StaminaStaminaModelRecoverDistributionSegment) ToDict() map[string]interface{} {

	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"min":   min,
		"max":   max,
		"count": count,
	}
}

func (p StaminaStaminaModelRecoverDistributionSegment) Pointer() *StaminaStaminaModelRecoverDistributionSegment {
	return &p
}

func CastStaminaStaminaModelRecoverDistributionSegments(data []interface{}) []StaminaStaminaModelRecoverDistributionSegment {
	v := make([]StaminaStaminaModelRecoverDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewStaminaStaminaModelRecoverDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaStaminaModelRecoverDistributionSegmentsFromDict(data []StaminaStaminaModelRecoverDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaStaminaModelRecoverDistribution struct {
	Statistics   *StaminaStaminaModelRecoverDistributionStatistics `json:"statistics"`
	Distribution []StaminaStaminaModelRecoverDistributionSegment   `json:"distribution"`
}

func NewStaminaStaminaModelRecoverDistributionFromJson(data string) StaminaStaminaModelRecoverDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaStaminaModelRecoverDistributionFromDict(dict)
}

func NewStaminaStaminaModelRecoverDistributionFromDict(data map[string]interface{}) StaminaStaminaModelRecoverDistribution {
	return StaminaStaminaModelRecoverDistribution{
		Statistics:   NewStaminaStaminaModelRecoverDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastStaminaStaminaModelRecoverDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p StaminaStaminaModelRecoverDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastStaminaStaminaModelRecoverDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p StaminaStaminaModelRecoverDistribution) Pointer() *StaminaStaminaModelRecoverDistribution {
	return &p
}

func CastStaminaStaminaModelRecoverDistributions(data []interface{}) []StaminaStaminaModelRecoverDistribution {
	v := make([]StaminaStaminaModelRecoverDistribution, 0)
	for _, d := range data {
		v = append(v, NewStaminaStaminaModelRecoverDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaStaminaModelRecoverDistributionsFromDict(data []StaminaStaminaModelRecoverDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaStaminaModelDistributions struct {
	Consume *StaminaStaminaModelConsumeDistribution `json:"consume"`
	Recover *StaminaStaminaModelRecoverDistribution `json:"recover"`
}

func NewStaminaStaminaModelDistributionsFromJson(data string) StaminaStaminaModelDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaStaminaModelDistributionsFromDict(dict)
}

func NewStaminaStaminaModelDistributionsFromDict(data map[string]interface{}) StaminaStaminaModelDistributions {
	return StaminaStaminaModelDistributions{
		Consume: NewStaminaStaminaModelConsumeDistributionFromDict(core.CastMap(data["consume"])).Pointer(),
		Recover: NewStaminaStaminaModelRecoverDistributionFromDict(core.CastMap(data["recover"])).Pointer(),
	}
}

func (p StaminaStaminaModelDistributions) ToDict() map[string]interface{} {

	var consume map[string]interface{}
	if p.Consume != nil {
		consume = p.Consume.ToDict()
	}
	var recover map[string]interface{}
	if p.Recover != nil {
		recover = p.Recover.ToDict()
	}
	return map[string]interface{}{
		"consume": consume,
		"recover": recover,
	}
}

func (p StaminaStaminaModelDistributions) Pointer() *StaminaStaminaModelDistributions {
	return &p
}

func CastStaminaStaminaModelDistributionses(data []interface{}) []StaminaStaminaModelDistributions {
	v := make([]StaminaStaminaModelDistributions, 0)
	for _, d := range data {
		v = append(v, NewStaminaStaminaModelDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaStaminaModelDistributionsesFromDict(data []StaminaStaminaModelDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaStaminaModel struct {
	StaminaModelId *string                           `json:"staminaModelId"`
	StaminaName    *string                           `json:"staminaName"`
	Statistics     *StaminaStaminaModelStatistics    `json:"statistics"`
	Distributions  *StaminaStaminaModelDistributions `json:"distributions"`
}

func NewStaminaStaminaModelFromJson(data string) StaminaStaminaModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaStaminaModelFromDict(dict)
}

func NewStaminaStaminaModelFromDict(data map[string]interface{}) StaminaStaminaModel {
	return StaminaStaminaModel{
		StaminaModelId: core.CastString(data["staminaModelId"]),
		StaminaName:    core.CastString(data["staminaName"]),
		Statistics:     NewStaminaStaminaModelStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions:  NewStaminaStaminaModelDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
	}
}

func (p StaminaStaminaModel) ToDict() map[string]interface{} {

	var staminaModelId *string
	if p.StaminaModelId != nil {
		staminaModelId = p.StaminaModelId
	}
	var staminaName *string
	if p.StaminaName != nil {
		staminaName = p.StaminaName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	return map[string]interface{}{
		"staminaModelId": staminaModelId,
		"staminaName":    staminaName,
		"statistics":     statistics,
		"distributions":  distributions,
	}
}

func (p StaminaStaminaModel) Pointer() *StaminaStaminaModel {
	return &p
}

func CastStaminaStaminaModels(data []interface{}) []StaminaStaminaModel {
	v := make([]StaminaStaminaModel, 0)
	for _, d := range data {
		v = append(v, NewStaminaStaminaModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaStaminaModelsFromDict(data []StaminaStaminaModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaNamespaceStatistics struct {
	Consume *int64 `json:"consume"`
	Recover *int64 `json:"recover"`
}

func NewStaminaNamespaceStatisticsFromJson(data string) StaminaNamespaceStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaNamespaceStatisticsFromDict(dict)
}

func NewStaminaNamespaceStatisticsFromDict(data map[string]interface{}) StaminaNamespaceStatistics {
	return StaminaNamespaceStatistics{
		Consume: core.CastInt64(data["consume"]),
		Recover: core.CastInt64(data["recover"]),
	}
}

func (p StaminaNamespaceStatistics) ToDict() map[string]interface{} {

	var consume *int64
	if p.Consume != nil {
		consume = p.Consume
	}
	var recover *int64
	if p.Recover != nil {
		recover = p.Recover
	}
	return map[string]interface{}{
		"consume": consume,
		"recover": recover,
	}
}

func (p StaminaNamespaceStatistics) Pointer() *StaminaNamespaceStatistics {
	return &p
}

func CastStaminaNamespaceStatisticses(data []interface{}) []StaminaNamespaceStatistics {
	v := make([]StaminaNamespaceStatistics, 0)
	for _, d := range data {
		v = append(v, NewStaminaNamespaceStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaNamespaceStatisticsesFromDict(data []StaminaNamespaceStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaNamespaceConsumeDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewStaminaNamespaceConsumeDistributionStatisticsFromJson(data string) StaminaNamespaceConsumeDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaNamespaceConsumeDistributionStatisticsFromDict(dict)
}

func NewStaminaNamespaceConsumeDistributionStatisticsFromDict(data map[string]interface{}) StaminaNamespaceConsumeDistributionStatistics {
	return StaminaNamespaceConsumeDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p StaminaNamespaceConsumeDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p StaminaNamespaceConsumeDistributionStatistics) Pointer() *StaminaNamespaceConsumeDistributionStatistics {
	return &p
}

func CastStaminaNamespaceConsumeDistributionStatisticses(data []interface{}) []StaminaNamespaceConsumeDistributionStatistics {
	v := make([]StaminaNamespaceConsumeDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewStaminaNamespaceConsumeDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaNamespaceConsumeDistributionStatisticsesFromDict(data []StaminaNamespaceConsumeDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaNamespaceConsumeDistributionSegment struct {
	StaminaName *string `json:"staminaName"`
	Count       *int64  `json:"count"`
}

func NewStaminaNamespaceConsumeDistributionSegmentFromJson(data string) StaminaNamespaceConsumeDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaNamespaceConsumeDistributionSegmentFromDict(dict)
}

func NewStaminaNamespaceConsumeDistributionSegmentFromDict(data map[string]interface{}) StaminaNamespaceConsumeDistributionSegment {
	return StaminaNamespaceConsumeDistributionSegment{
		StaminaName: core.CastString(data["staminaName"]),
		Count:       core.CastInt64(data["count"]),
	}
}

func (p StaminaNamespaceConsumeDistributionSegment) ToDict() map[string]interface{} {

	var staminaName *string
	if p.StaminaName != nil {
		staminaName = p.StaminaName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"staminaName": staminaName,
		"count":       count,
	}
}

func (p StaminaNamespaceConsumeDistributionSegment) Pointer() *StaminaNamespaceConsumeDistributionSegment {
	return &p
}

func CastStaminaNamespaceConsumeDistributionSegments(data []interface{}) []StaminaNamespaceConsumeDistributionSegment {
	v := make([]StaminaNamespaceConsumeDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewStaminaNamespaceConsumeDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaNamespaceConsumeDistributionSegmentsFromDict(data []StaminaNamespaceConsumeDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaNamespaceConsumeDistribution struct {
	Statistics   *StaminaNamespaceConsumeDistributionStatistics `json:"statistics"`
	Distribution []StaminaNamespaceConsumeDistributionSegment   `json:"distribution"`
}

func NewStaminaNamespaceConsumeDistributionFromJson(data string) StaminaNamespaceConsumeDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaNamespaceConsumeDistributionFromDict(dict)
}

func NewStaminaNamespaceConsumeDistributionFromDict(data map[string]interface{}) StaminaNamespaceConsumeDistribution {
	return StaminaNamespaceConsumeDistribution{
		Statistics:   NewStaminaNamespaceConsumeDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastStaminaNamespaceConsumeDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p StaminaNamespaceConsumeDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastStaminaNamespaceConsumeDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p StaminaNamespaceConsumeDistribution) Pointer() *StaminaNamespaceConsumeDistribution {
	return &p
}

func CastStaminaNamespaceConsumeDistributions(data []interface{}) []StaminaNamespaceConsumeDistribution {
	v := make([]StaminaNamespaceConsumeDistribution, 0)
	for _, d := range data {
		v = append(v, NewStaminaNamespaceConsumeDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaNamespaceConsumeDistributionsFromDict(data []StaminaNamespaceConsumeDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaNamespaceRecoverDistributionStatistics struct {
	Count  *int64   `json:"count"`
	Min    *int64   `json:"min"`
	Max    *int64   `json:"max"`
	Avg    *float32 `json:"avg"`
	Median *int64   `json:"median"`
	Stddev *float32 `json:"stddev"`
}

func NewStaminaNamespaceRecoverDistributionStatisticsFromJson(data string) StaminaNamespaceRecoverDistributionStatistics {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaNamespaceRecoverDistributionStatisticsFromDict(dict)
}

func NewStaminaNamespaceRecoverDistributionStatisticsFromDict(data map[string]interface{}) StaminaNamespaceRecoverDistributionStatistics {
	return StaminaNamespaceRecoverDistributionStatistics{
		Count:  core.CastInt64(data["count"]),
		Min:    core.CastInt64(data["min"]),
		Max:    core.CastInt64(data["max"]),
		Avg:    core.CastFloat32(data["avg"]),
		Median: core.CastInt64(data["median"]),
		Stddev: core.CastFloat32(data["stddev"]),
	}
}

func (p StaminaNamespaceRecoverDistributionStatistics) ToDict() map[string]interface{} {

	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var min *int64
	if p.Min != nil {
		min = p.Min
	}
	var max *int64
	if p.Max != nil {
		max = p.Max
	}
	var avg *float32
	if p.Avg != nil {
		avg = p.Avg
	}
	var median *int64
	if p.Median != nil {
		median = p.Median
	}
	var stddev *float32
	if p.Stddev != nil {
		stddev = p.Stddev
	}
	return map[string]interface{}{
		"count":  count,
		"min":    min,
		"max":    max,
		"avg":    avg,
		"median": median,
		"stddev": stddev,
	}
}

func (p StaminaNamespaceRecoverDistributionStatistics) Pointer() *StaminaNamespaceRecoverDistributionStatistics {
	return &p
}

func CastStaminaNamespaceRecoverDistributionStatisticses(data []interface{}) []StaminaNamespaceRecoverDistributionStatistics {
	v := make([]StaminaNamespaceRecoverDistributionStatistics, 0)
	for _, d := range data {
		v = append(v, NewStaminaNamespaceRecoverDistributionStatisticsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaNamespaceRecoverDistributionStatisticsesFromDict(data []StaminaNamespaceRecoverDistributionStatistics) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaNamespaceRecoverDistributionSegment struct {
	StaminaName *string `json:"staminaName"`
	Count       *int64  `json:"count"`
}

func NewStaminaNamespaceRecoverDistributionSegmentFromJson(data string) StaminaNamespaceRecoverDistributionSegment {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaNamespaceRecoverDistributionSegmentFromDict(dict)
}

func NewStaminaNamespaceRecoverDistributionSegmentFromDict(data map[string]interface{}) StaminaNamespaceRecoverDistributionSegment {
	return StaminaNamespaceRecoverDistributionSegment{
		StaminaName: core.CastString(data["staminaName"]),
		Count:       core.CastInt64(data["count"]),
	}
}

func (p StaminaNamespaceRecoverDistributionSegment) ToDict() map[string]interface{} {

	var staminaName *string
	if p.StaminaName != nil {
		staminaName = p.StaminaName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"staminaName": staminaName,
		"count":       count,
	}
}

func (p StaminaNamespaceRecoverDistributionSegment) Pointer() *StaminaNamespaceRecoverDistributionSegment {
	return &p
}

func CastStaminaNamespaceRecoverDistributionSegments(data []interface{}) []StaminaNamespaceRecoverDistributionSegment {
	v := make([]StaminaNamespaceRecoverDistributionSegment, 0)
	for _, d := range data {
		v = append(v, NewStaminaNamespaceRecoverDistributionSegmentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaNamespaceRecoverDistributionSegmentsFromDict(data []StaminaNamespaceRecoverDistributionSegment) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaNamespaceRecoverDistribution struct {
	Statistics   *StaminaNamespaceRecoverDistributionStatistics `json:"statistics"`
	Distribution []StaminaNamespaceRecoverDistributionSegment   `json:"distribution"`
}

func NewStaminaNamespaceRecoverDistributionFromJson(data string) StaminaNamespaceRecoverDistribution {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaNamespaceRecoverDistributionFromDict(dict)
}

func NewStaminaNamespaceRecoverDistributionFromDict(data map[string]interface{}) StaminaNamespaceRecoverDistribution {
	return StaminaNamespaceRecoverDistribution{
		Statistics:   NewStaminaNamespaceRecoverDistributionStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distribution: CastStaminaNamespaceRecoverDistributionSegments(core.CastArray(data["distribution"])),
	}
}

func (p StaminaNamespaceRecoverDistribution) ToDict() map[string]interface{} {

	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distribution []interface{}
	if p.Distribution != nil {
		distribution = CastStaminaNamespaceRecoverDistributionSegmentsFromDict(
			p.Distribution,
		)
	}
	return map[string]interface{}{
		"statistics":   statistics,
		"distribution": distribution,
	}
}

func (p StaminaNamespaceRecoverDistribution) Pointer() *StaminaNamespaceRecoverDistribution {
	return &p
}

func CastStaminaNamespaceRecoverDistributions(data []interface{}) []StaminaNamespaceRecoverDistribution {
	v := make([]StaminaNamespaceRecoverDistribution, 0)
	for _, d := range data {
		v = append(v, NewStaminaNamespaceRecoverDistributionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaNamespaceRecoverDistributionsFromDict(data []StaminaNamespaceRecoverDistribution) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaNamespaceDistributions struct {
	Consume *StaminaNamespaceConsumeDistribution `json:"consume"`
	Recover *StaminaNamespaceRecoverDistribution `json:"recover"`
}

func NewStaminaNamespaceDistributionsFromJson(data string) StaminaNamespaceDistributions {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaNamespaceDistributionsFromDict(dict)
}

func NewStaminaNamespaceDistributionsFromDict(data map[string]interface{}) StaminaNamespaceDistributions {
	return StaminaNamespaceDistributions{
		Consume: NewStaminaNamespaceConsumeDistributionFromDict(core.CastMap(data["consume"])).Pointer(),
		Recover: NewStaminaNamespaceRecoverDistributionFromDict(core.CastMap(data["recover"])).Pointer(),
	}
}

func (p StaminaNamespaceDistributions) ToDict() map[string]interface{} {

	var consume map[string]interface{}
	if p.Consume != nil {
		consume = p.Consume.ToDict()
	}
	var recover map[string]interface{}
	if p.Recover != nil {
		recover = p.Recover.ToDict()
	}
	return map[string]interface{}{
		"consume": consume,
		"recover": recover,
	}
}

func (p StaminaNamespaceDistributions) Pointer() *StaminaNamespaceDistributions {
	return &p
}

func CastStaminaNamespaceDistributionses(data []interface{}) []StaminaNamespaceDistributions {
	v := make([]StaminaNamespaceDistributions, 0)
	for _, d := range data {
		v = append(v, NewStaminaNamespaceDistributionsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaNamespaceDistributionsesFromDict(data []StaminaNamespaceDistributions) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaNamespace struct {
	NamespaceId   *string                        `json:"namespaceId"`
	Year          *int32                         `json:"year"`
	Month         *int32                         `json:"month"`
	Day           *int32                         `json:"day"`
	NamespaceName *string                        `json:"namespaceName"`
	Statistics    *StaminaNamespaceStatistics    `json:"statistics"`
	Distributions *StaminaNamespaceDistributions `json:"distributions"`
	StaminaModels []StaminaStaminaModel          `json:"staminaModels"`
}

func NewStaminaNamespaceFromJson(data string) StaminaNamespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStaminaNamespaceFromDict(dict)
}

func NewStaminaNamespaceFromDict(data map[string]interface{}) StaminaNamespace {
	return StaminaNamespace{
		NamespaceId:   core.CastString(data["namespaceId"]),
		Year:          core.CastInt32(data["year"]),
		Month:         core.CastInt32(data["month"]),
		Day:           core.CastInt32(data["day"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		Statistics:    NewStaminaNamespaceStatisticsFromDict(core.CastMap(data["statistics"])).Pointer(),
		Distributions: NewStaminaNamespaceDistributionsFromDict(core.CastMap(data["distributions"])).Pointer(),
		StaminaModels: CastStaminaStaminaModels(core.CastArray(data["staminaModels"])),
	}
}

func (p StaminaNamespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var day *int32
	if p.Day != nil {
		day = p.Day
	}
	var namespaceName *string
	if p.NamespaceName != nil {
		namespaceName = p.NamespaceName
	}
	var statistics map[string]interface{}
	if p.Statistics != nil {
		statistics = p.Statistics.ToDict()
	}
	var distributions map[string]interface{}
	if p.Distributions != nil {
		distributions = p.Distributions.ToDict()
	}
	var staminaModels []interface{}
	if p.StaminaModels != nil {
		staminaModels = CastStaminaStaminaModelsFromDict(
			p.StaminaModels,
		)
	}
	return map[string]interface{}{
		"namespaceId":   namespaceId,
		"year":          year,
		"month":         month,
		"day":           day,
		"namespaceName": namespaceName,
		"statistics":    statistics,
		"distributions": distributions,
		"staminaModels": staminaModels,
	}
}

func (p StaminaNamespace) Pointer() *StaminaNamespace {
	return &p
}

func CastStaminaNamespaces(data []interface{}) []StaminaNamespace {
	v := make([]StaminaNamespace, 0)
	for _, d := range data {
		v = append(v, NewStaminaNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaNamespacesFromDict(data []StaminaNamespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
