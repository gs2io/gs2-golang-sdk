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

type GetChartResult struct {
	Items         []Chart `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type GetChartAsyncResult struct {
	result *GetChartResult
	err    error
}

func NewGetChartResultFromJson(data string) GetChartResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetChartResultFromDict(dict)
}

func NewGetChartResultFromDict(data map[string]interface{}) GetChartResult {
	return GetChartResult{
		Items:         CastCharts(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p GetChartResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastChartsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p GetChartResult) Pointer() *GetChartResult {
	return &p
}

type GetDistributionResult struct {
	Items         []Distribution `json:"items"`
	NextPageToken *string        `json:"nextPageToken"`
}

type GetDistributionAsyncResult struct {
	result *GetDistributionResult
	err    error
}

func NewGetDistributionResultFromJson(data string) GetDistributionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDistributionResultFromDict(dict)
}

func NewGetDistributionResultFromDict(data map[string]interface{}) GetDistributionResult {
	return GetDistributionResult{
		Items:         CastDistributions(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p GetDistributionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDistributionsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p GetDistributionResult) Pointer() *GetDistributionResult {
	return &p
}

type GetCumulativeResult struct {
	Item *Cumulative `json:"item"`
}

type GetCumulativeAsyncResult struct {
	result *GetCumulativeResult
	err    error
}

func NewGetCumulativeResultFromJson(data string) GetCumulativeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCumulativeResultFromDict(dict)
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

func NewDescribeBillingActivitiesResultFromJson(data string) DescribeBillingActivitiesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBillingActivitiesResultFromDict(dict)
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

func NewGetBillingActivityResultFromJson(data string) GetBillingActivityResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBillingActivityResultFromDict(dict)
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

type GetGeneralMetricsResult struct {
	Item *GeneralMetrics `json:"item"`
}

type GetGeneralMetricsAsyncResult struct {
	result *GetGeneralMetricsResult
	err    error
}

func NewGetGeneralMetricsResultFromJson(data string) GetGeneralMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGeneralMetricsResultFromDict(dict)
}

func NewGetGeneralMetricsResultFromDict(data map[string]interface{}) GetGeneralMetricsResult {
	return GetGeneralMetricsResult{
		Item: NewGeneralMetricsFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetGeneralMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetGeneralMetricsResult) Pointer() *GetGeneralMetricsResult {
	return &p
}

type DescribeAccountNamespaceMetricsResult struct {
	Items []AccountNamespace `json:"items"`
}

type DescribeAccountNamespaceMetricsAsyncResult struct {
	result *DescribeAccountNamespaceMetricsResult
	err    error
}

func NewDescribeAccountNamespaceMetricsResultFromJson(data string) DescribeAccountNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAccountNamespaceMetricsResultFromDict(dict)
}

func NewDescribeAccountNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeAccountNamespaceMetricsResult {
	return DescribeAccountNamespaceMetricsResult{
		Items: CastAccountNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeAccountNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAccountNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeAccountNamespaceMetricsResult) Pointer() *DescribeAccountNamespaceMetricsResult {
	return &p
}

type GetAccountNamespaceMetricsResult struct {
	Item *AccountNamespace `json:"item"`
}

type GetAccountNamespaceMetricsAsyncResult struct {
	result *GetAccountNamespaceMetricsResult
	err    error
}

func NewGetAccountNamespaceMetricsResultFromJson(data string) GetAccountNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAccountNamespaceMetricsResultFromDict(dict)
}

func NewGetAccountNamespaceMetricsResultFromDict(data map[string]interface{}) GetAccountNamespaceMetricsResult {
	return GetAccountNamespaceMetricsResult{
		Item: NewAccountNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetAccountNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetAccountNamespaceMetricsResult) Pointer() *GetAccountNamespaceMetricsResult {
	return &p
}

type DescribeChatNamespaceMetricsResult struct {
	Items []ChatNamespace `json:"items"`
}

type DescribeChatNamespaceMetricsAsyncResult struct {
	result *DescribeChatNamespaceMetricsResult
	err    error
}

func NewDescribeChatNamespaceMetricsResultFromJson(data string) DescribeChatNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeChatNamespaceMetricsResultFromDict(dict)
}

func NewDescribeChatNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeChatNamespaceMetricsResult {
	return DescribeChatNamespaceMetricsResult{
		Items: CastChatNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeChatNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastChatNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeChatNamespaceMetricsResult) Pointer() *DescribeChatNamespaceMetricsResult {
	return &p
}

type GetChatNamespaceMetricsResult struct {
	Item *ChatNamespace `json:"item"`
}

type GetChatNamespaceMetricsAsyncResult struct {
	result *GetChatNamespaceMetricsResult
	err    error
}

func NewGetChatNamespaceMetricsResultFromJson(data string) GetChatNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetChatNamespaceMetricsResultFromDict(dict)
}

func NewGetChatNamespaceMetricsResultFromDict(data map[string]interface{}) GetChatNamespaceMetricsResult {
	return GetChatNamespaceMetricsResult{
		Item: NewChatNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetChatNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetChatNamespaceMetricsResult) Pointer() *GetChatNamespaceMetricsResult {
	return &p
}

type DescribeDatastoreNamespaceMetricsResult struct {
	Items []DatastoreNamespace `json:"items"`
}

type DescribeDatastoreNamespaceMetricsAsyncResult struct {
	result *DescribeDatastoreNamespaceMetricsResult
	err    error
}

func NewDescribeDatastoreNamespaceMetricsResultFromJson(data string) DescribeDatastoreNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDatastoreNamespaceMetricsResultFromDict(dict)
}

func NewDescribeDatastoreNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeDatastoreNamespaceMetricsResult {
	return DescribeDatastoreNamespaceMetricsResult{
		Items: CastDatastoreNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeDatastoreNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDatastoreNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeDatastoreNamespaceMetricsResult) Pointer() *DescribeDatastoreNamespaceMetricsResult {
	return &p
}

type GetDatastoreNamespaceMetricsResult struct {
	Item *DatastoreNamespace `json:"item"`
}

type GetDatastoreNamespaceMetricsAsyncResult struct {
	result *GetDatastoreNamespaceMetricsResult
	err    error
}

func NewGetDatastoreNamespaceMetricsResultFromJson(data string) GetDatastoreNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDatastoreNamespaceMetricsResultFromDict(dict)
}

func NewGetDatastoreNamespaceMetricsResultFromDict(data map[string]interface{}) GetDatastoreNamespaceMetricsResult {
	return GetDatastoreNamespaceMetricsResult{
		Item: NewDatastoreNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetDatastoreNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetDatastoreNamespaceMetricsResult) Pointer() *GetDatastoreNamespaceMetricsResult {
	return &p
}

type DescribeDictionaryNamespaceMetricsResult struct {
	Items []DictionaryNamespace `json:"items"`
}

type DescribeDictionaryNamespaceMetricsAsyncResult struct {
	result *DescribeDictionaryNamespaceMetricsResult
	err    error
}

func NewDescribeDictionaryNamespaceMetricsResultFromJson(data string) DescribeDictionaryNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDictionaryNamespaceMetricsResultFromDict(dict)
}

func NewDescribeDictionaryNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeDictionaryNamespaceMetricsResult {
	return DescribeDictionaryNamespaceMetricsResult{
		Items: CastDictionaryNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeDictionaryNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDictionaryNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeDictionaryNamespaceMetricsResult) Pointer() *DescribeDictionaryNamespaceMetricsResult {
	return &p
}

type GetDictionaryNamespaceMetricsResult struct {
	Item *DictionaryNamespace `json:"item"`
}

type GetDictionaryNamespaceMetricsAsyncResult struct {
	result *GetDictionaryNamespaceMetricsResult
	err    error
}

func NewGetDictionaryNamespaceMetricsResultFromJson(data string) GetDictionaryNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDictionaryNamespaceMetricsResultFromDict(dict)
}

func NewGetDictionaryNamespaceMetricsResultFromDict(data map[string]interface{}) GetDictionaryNamespaceMetricsResult {
	return GetDictionaryNamespaceMetricsResult{
		Item: NewDictionaryNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetDictionaryNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetDictionaryNamespaceMetricsResult) Pointer() *GetDictionaryNamespaceMetricsResult {
	return &p
}

type DescribeExchangeRateModelMetricsResult struct {
	Items []ExchangeRateModel `json:"items"`
}

type DescribeExchangeRateModelMetricsAsyncResult struct {
	result *DescribeExchangeRateModelMetricsResult
	err    error
}

func NewDescribeExchangeRateModelMetricsResultFromJson(data string) DescribeExchangeRateModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeExchangeRateModelMetricsResultFromDict(dict)
}

func NewDescribeExchangeRateModelMetricsResultFromDict(data map[string]interface{}) DescribeExchangeRateModelMetricsResult {
	return DescribeExchangeRateModelMetricsResult{
		Items: CastExchangeRateModels(core.CastArray(data["items"])),
	}
}

func (p DescribeExchangeRateModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastExchangeRateModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeExchangeRateModelMetricsResult) Pointer() *DescribeExchangeRateModelMetricsResult {
	return &p
}

type GetExchangeRateModelMetricsResult struct {
	Item *ExchangeRateModel `json:"item"`
}

type GetExchangeRateModelMetricsAsyncResult struct {
	result *GetExchangeRateModelMetricsResult
	err    error
}

func NewGetExchangeRateModelMetricsResultFromJson(data string) GetExchangeRateModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetExchangeRateModelMetricsResultFromDict(dict)
}

func NewGetExchangeRateModelMetricsResultFromDict(data map[string]interface{}) GetExchangeRateModelMetricsResult {
	return GetExchangeRateModelMetricsResult{
		Item: NewExchangeRateModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetExchangeRateModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetExchangeRateModelMetricsResult) Pointer() *GetExchangeRateModelMetricsResult {
	return &p
}

type DescribeExchangeNamespaceMetricsResult struct {
	Items []ExchangeNamespace `json:"items"`
}

type DescribeExchangeNamespaceMetricsAsyncResult struct {
	result *DescribeExchangeNamespaceMetricsResult
	err    error
}

func NewDescribeExchangeNamespaceMetricsResultFromJson(data string) DescribeExchangeNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeExchangeNamespaceMetricsResultFromDict(dict)
}

func NewDescribeExchangeNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeExchangeNamespaceMetricsResult {
	return DescribeExchangeNamespaceMetricsResult{
		Items: CastExchangeNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeExchangeNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastExchangeNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeExchangeNamespaceMetricsResult) Pointer() *DescribeExchangeNamespaceMetricsResult {
	return &p
}

type GetExchangeNamespaceMetricsResult struct {
	Item *ExchangeNamespace `json:"item"`
}

type GetExchangeNamespaceMetricsAsyncResult struct {
	result *GetExchangeNamespaceMetricsResult
	err    error
}

func NewGetExchangeNamespaceMetricsResultFromJson(data string) GetExchangeNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetExchangeNamespaceMetricsResultFromDict(dict)
}

func NewGetExchangeNamespaceMetricsResultFromDict(data map[string]interface{}) GetExchangeNamespaceMetricsResult {
	return GetExchangeNamespaceMetricsResult{
		Item: NewExchangeNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetExchangeNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetExchangeNamespaceMetricsResult) Pointer() *GetExchangeNamespaceMetricsResult {
	return &p
}

type DescribeExperienceStatusMetricsResult struct {
	Items []ExperienceStatus `json:"items"`
}

type DescribeExperienceStatusMetricsAsyncResult struct {
	result *DescribeExperienceStatusMetricsResult
	err    error
}

func NewDescribeExperienceStatusMetricsResultFromJson(data string) DescribeExperienceStatusMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeExperienceStatusMetricsResultFromDict(dict)
}

func NewDescribeExperienceStatusMetricsResultFromDict(data map[string]interface{}) DescribeExperienceStatusMetricsResult {
	return DescribeExperienceStatusMetricsResult{
		Items: CastExperienceStatuses(core.CastArray(data["items"])),
	}
}

func (p DescribeExperienceStatusMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastExperienceStatusesFromDict(
			p.Items,
		),
	}
}

func (p DescribeExperienceStatusMetricsResult) Pointer() *DescribeExperienceStatusMetricsResult {
	return &p
}

type DescribeExperienceExperienceModelMetricsResult struct {
	Items []ExperienceExperienceModel `json:"items"`
}

type DescribeExperienceExperienceModelMetricsAsyncResult struct {
	result *DescribeExperienceExperienceModelMetricsResult
	err    error
}

func NewDescribeExperienceExperienceModelMetricsResultFromJson(data string) DescribeExperienceExperienceModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeExperienceExperienceModelMetricsResultFromDict(dict)
}

func NewDescribeExperienceExperienceModelMetricsResultFromDict(data map[string]interface{}) DescribeExperienceExperienceModelMetricsResult {
	return DescribeExperienceExperienceModelMetricsResult{
		Items: CastExperienceExperienceModels(core.CastArray(data["items"])),
	}
}

func (p DescribeExperienceExperienceModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastExperienceExperienceModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeExperienceExperienceModelMetricsResult) Pointer() *DescribeExperienceExperienceModelMetricsResult {
	return &p
}

type GetExperienceExperienceModelMetricsResult struct {
	Item *ExperienceExperienceModel `json:"item"`
}

type GetExperienceExperienceModelMetricsAsyncResult struct {
	result *GetExperienceExperienceModelMetricsResult
	err    error
}

func NewGetExperienceExperienceModelMetricsResultFromJson(data string) GetExperienceExperienceModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetExperienceExperienceModelMetricsResultFromDict(dict)
}

func NewGetExperienceExperienceModelMetricsResultFromDict(data map[string]interface{}) GetExperienceExperienceModelMetricsResult {
	return GetExperienceExperienceModelMetricsResult{
		Item: NewExperienceExperienceModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetExperienceExperienceModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetExperienceExperienceModelMetricsResult) Pointer() *GetExperienceExperienceModelMetricsResult {
	return &p
}

type DescribeExperienceNamespaceMetricsResult struct {
	Items []ExperienceNamespace `json:"items"`
}

type DescribeExperienceNamespaceMetricsAsyncResult struct {
	result *DescribeExperienceNamespaceMetricsResult
	err    error
}

func NewDescribeExperienceNamespaceMetricsResultFromJson(data string) DescribeExperienceNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeExperienceNamespaceMetricsResultFromDict(dict)
}

func NewDescribeExperienceNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeExperienceNamespaceMetricsResult {
	return DescribeExperienceNamespaceMetricsResult{
		Items: CastExperienceNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeExperienceNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastExperienceNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeExperienceNamespaceMetricsResult) Pointer() *DescribeExperienceNamespaceMetricsResult {
	return &p
}

type GetExperienceNamespaceMetricsResult struct {
	Item *ExperienceNamespace `json:"item"`
}

type GetExperienceNamespaceMetricsAsyncResult struct {
	result *GetExperienceNamespaceMetricsResult
	err    error
}

func NewGetExperienceNamespaceMetricsResultFromJson(data string) GetExperienceNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetExperienceNamespaceMetricsResultFromDict(dict)
}

func NewGetExperienceNamespaceMetricsResultFromDict(data map[string]interface{}) GetExperienceNamespaceMetricsResult {
	return GetExperienceNamespaceMetricsResult{
		Item: NewExperienceNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetExperienceNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetExperienceNamespaceMetricsResult) Pointer() *GetExperienceNamespaceMetricsResult {
	return &p
}

type DescribeFormationFormMetricsResult struct {
	Items []FormationForm `json:"items"`
}

type DescribeFormationFormMetricsAsyncResult struct {
	result *DescribeFormationFormMetricsResult
	err    error
}

func NewDescribeFormationFormMetricsResultFromJson(data string) DescribeFormationFormMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFormationFormMetricsResultFromDict(dict)
}

func NewDescribeFormationFormMetricsResultFromDict(data map[string]interface{}) DescribeFormationFormMetricsResult {
	return DescribeFormationFormMetricsResult{
		Items: CastFormationForms(core.CastArray(data["items"])),
	}
}

func (p DescribeFormationFormMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFormationFormsFromDict(
			p.Items,
		),
	}
}

func (p DescribeFormationFormMetricsResult) Pointer() *DescribeFormationFormMetricsResult {
	return &p
}

type DescribeFormationMoldMetricsResult struct {
	Items []FormationMold `json:"items"`
}

type DescribeFormationMoldMetricsAsyncResult struct {
	result *DescribeFormationMoldMetricsResult
	err    error
}

func NewDescribeFormationMoldMetricsResultFromJson(data string) DescribeFormationMoldMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFormationMoldMetricsResultFromDict(dict)
}

func NewDescribeFormationMoldMetricsResultFromDict(data map[string]interface{}) DescribeFormationMoldMetricsResult {
	return DescribeFormationMoldMetricsResult{
		Items: CastFormationMolds(core.CastArray(data["items"])),
	}
}

func (p DescribeFormationMoldMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFormationMoldsFromDict(
			p.Items,
		),
	}
}

func (p DescribeFormationMoldMetricsResult) Pointer() *DescribeFormationMoldMetricsResult {
	return &p
}

type DescribeFormationNamespaceMetricsResult struct {
	Items []FormationNamespace `json:"items"`
}

type DescribeFormationNamespaceMetricsAsyncResult struct {
	result *DescribeFormationNamespaceMetricsResult
	err    error
}

func NewDescribeFormationNamespaceMetricsResultFromJson(data string) DescribeFormationNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFormationNamespaceMetricsResultFromDict(dict)
}

func NewDescribeFormationNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeFormationNamespaceMetricsResult {
	return DescribeFormationNamespaceMetricsResult{
		Items: CastFormationNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeFormationNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFormationNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeFormationNamespaceMetricsResult) Pointer() *DescribeFormationNamespaceMetricsResult {
	return &p
}

type GetFormationNamespaceMetricsResult struct {
	Item *FormationNamespace `json:"item"`
}

type GetFormationNamespaceMetricsAsyncResult struct {
	result *GetFormationNamespaceMetricsResult
	err    error
}

func NewGetFormationNamespaceMetricsResultFromJson(data string) GetFormationNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormationNamespaceMetricsResultFromDict(dict)
}

func NewGetFormationNamespaceMetricsResultFromDict(data map[string]interface{}) GetFormationNamespaceMetricsResult {
	return GetFormationNamespaceMetricsResult{
		Item: NewFormationNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetFormationNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetFormationNamespaceMetricsResult) Pointer() *GetFormationNamespaceMetricsResult {
	return &p
}

type DescribeFriendNamespaceMetricsResult struct {
	Items []FriendNamespace `json:"items"`
}

type DescribeFriendNamespaceMetricsAsyncResult struct {
	result *DescribeFriendNamespaceMetricsResult
	err    error
}

func NewDescribeFriendNamespaceMetricsResultFromJson(data string) DescribeFriendNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFriendNamespaceMetricsResultFromDict(dict)
}

func NewDescribeFriendNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeFriendNamespaceMetricsResult {
	return DescribeFriendNamespaceMetricsResult{
		Items: CastFriendNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeFriendNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFriendNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeFriendNamespaceMetricsResult) Pointer() *DescribeFriendNamespaceMetricsResult {
	return &p
}

type GetFriendNamespaceMetricsResult struct {
	Item *FriendNamespace `json:"item"`
}

type GetFriendNamespaceMetricsAsyncResult struct {
	result *GetFriendNamespaceMetricsResult
	err    error
}

func NewGetFriendNamespaceMetricsResultFromJson(data string) GetFriendNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFriendNamespaceMetricsResultFromDict(dict)
}

func NewGetFriendNamespaceMetricsResultFromDict(data map[string]interface{}) GetFriendNamespaceMetricsResult {
	return GetFriendNamespaceMetricsResult{
		Item: NewFriendNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetFriendNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetFriendNamespaceMetricsResult) Pointer() *GetFriendNamespaceMetricsResult {
	return &p
}

type DescribeInboxNamespaceMetricsResult struct {
	Items []InboxNamespace `json:"items"`
}

type DescribeInboxNamespaceMetricsAsyncResult struct {
	result *DescribeInboxNamespaceMetricsResult
	err    error
}

func NewDescribeInboxNamespaceMetricsResultFromJson(data string) DescribeInboxNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInboxNamespaceMetricsResultFromDict(dict)
}

func NewDescribeInboxNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeInboxNamespaceMetricsResult {
	return DescribeInboxNamespaceMetricsResult{
		Items: CastInboxNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeInboxNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInboxNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeInboxNamespaceMetricsResult) Pointer() *DescribeInboxNamespaceMetricsResult {
	return &p
}

type GetInboxNamespaceMetricsResult struct {
	Item *InboxNamespace `json:"item"`
}

type GetInboxNamespaceMetricsAsyncResult struct {
	result *GetInboxNamespaceMetricsResult
	err    error
}

func NewGetInboxNamespaceMetricsResultFromJson(data string) GetInboxNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInboxNamespaceMetricsResultFromDict(dict)
}

func NewGetInboxNamespaceMetricsResultFromDict(data map[string]interface{}) GetInboxNamespaceMetricsResult {
	return GetInboxNamespaceMetricsResult{
		Item: NewInboxNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetInboxNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetInboxNamespaceMetricsResult) Pointer() *GetInboxNamespaceMetricsResult {
	return &p
}

type DescribeInventoryItemSetMetricsResult struct {
	Items []InventoryItemSet `json:"items"`
}

type DescribeInventoryItemSetMetricsAsyncResult struct {
	result *DescribeInventoryItemSetMetricsResult
	err    error
}

func NewDescribeInventoryItemSetMetricsResultFromJson(data string) DescribeInventoryItemSetMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoryItemSetMetricsResultFromDict(dict)
}

func NewDescribeInventoryItemSetMetricsResultFromDict(data map[string]interface{}) DescribeInventoryItemSetMetricsResult {
	return DescribeInventoryItemSetMetricsResult{
		Items: CastInventoryItemSets(core.CastArray(data["items"])),
	}
}

func (p DescribeInventoryItemSetMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoryItemSetsFromDict(
			p.Items,
		),
	}
}

func (p DescribeInventoryItemSetMetricsResult) Pointer() *DescribeInventoryItemSetMetricsResult {
	return &p
}

type DescribeInventoryInventoryMetricsResult struct {
	Items []InventoryInventory `json:"items"`
}

type DescribeInventoryInventoryMetricsAsyncResult struct {
	result *DescribeInventoryInventoryMetricsResult
	err    error
}

func NewDescribeInventoryInventoryMetricsResultFromJson(data string) DescribeInventoryInventoryMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoryInventoryMetricsResultFromDict(dict)
}

func NewDescribeInventoryInventoryMetricsResultFromDict(data map[string]interface{}) DescribeInventoryInventoryMetricsResult {
	return DescribeInventoryInventoryMetricsResult{
		Items: CastInventoryInventories(core.CastArray(data["items"])),
	}
}

func (p DescribeInventoryInventoryMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoryInventoriesFromDict(
			p.Items,
		),
	}
}

func (p DescribeInventoryInventoryMetricsResult) Pointer() *DescribeInventoryInventoryMetricsResult {
	return &p
}

type DescribeInventoryNamespaceMetricsResult struct {
	Items []InventoryNamespace `json:"items"`
}

type DescribeInventoryNamespaceMetricsAsyncResult struct {
	result *DescribeInventoryNamespaceMetricsResult
	err    error
}

func NewDescribeInventoryNamespaceMetricsResultFromJson(data string) DescribeInventoryNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoryNamespaceMetricsResultFromDict(dict)
}

func NewDescribeInventoryNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeInventoryNamespaceMetricsResult {
	return DescribeInventoryNamespaceMetricsResult{
		Items: CastInventoryNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeInventoryNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoryNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeInventoryNamespaceMetricsResult) Pointer() *DescribeInventoryNamespaceMetricsResult {
	return &p
}

type GetInventoryNamespaceMetricsResult struct {
	Item *InventoryNamespace `json:"item"`
}

type GetInventoryNamespaceMetricsAsyncResult struct {
	result *GetInventoryNamespaceMetricsResult
	err    error
}

func NewGetInventoryNamespaceMetricsResultFromJson(data string) GetInventoryNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInventoryNamespaceMetricsResultFromDict(dict)
}

func NewGetInventoryNamespaceMetricsResultFromDict(data map[string]interface{}) GetInventoryNamespaceMetricsResult {
	return GetInventoryNamespaceMetricsResult{
		Item: NewInventoryNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetInventoryNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetInventoryNamespaceMetricsResult) Pointer() *GetInventoryNamespaceMetricsResult {
	return &p
}

type DescribeKeyNamespaceMetricsResult struct {
	Items []KeyNamespace `json:"items"`
}

type DescribeKeyNamespaceMetricsAsyncResult struct {
	result *DescribeKeyNamespaceMetricsResult
	err    error
}

func NewDescribeKeyNamespaceMetricsResultFromJson(data string) DescribeKeyNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeKeyNamespaceMetricsResultFromDict(dict)
}

func NewDescribeKeyNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeKeyNamespaceMetricsResult {
	return DescribeKeyNamespaceMetricsResult{
		Items: CastKeyNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeKeyNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastKeyNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeKeyNamespaceMetricsResult) Pointer() *DescribeKeyNamespaceMetricsResult {
	return &p
}

type GetKeyNamespaceMetricsResult struct {
	Item *KeyNamespace `json:"item"`
}

type GetKeyNamespaceMetricsAsyncResult struct {
	result *GetKeyNamespaceMetricsResult
	err    error
}

func NewGetKeyNamespaceMetricsResultFromJson(data string) GetKeyNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetKeyNamespaceMetricsResultFromDict(dict)
}

func NewGetKeyNamespaceMetricsResultFromDict(data map[string]interface{}) GetKeyNamespaceMetricsResult {
	return GetKeyNamespaceMetricsResult{
		Item: NewKeyNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetKeyNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetKeyNamespaceMetricsResult) Pointer() *GetKeyNamespaceMetricsResult {
	return &p
}

type DescribeLimitCounterMetricsResult struct {
	Items []LimitCounter `json:"items"`
}

type DescribeLimitCounterMetricsAsyncResult struct {
	result *DescribeLimitCounterMetricsResult
	err    error
}

func NewDescribeLimitCounterMetricsResultFromJson(data string) DescribeLimitCounterMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLimitCounterMetricsResultFromDict(dict)
}

func NewDescribeLimitCounterMetricsResultFromDict(data map[string]interface{}) DescribeLimitCounterMetricsResult {
	return DescribeLimitCounterMetricsResult{
		Items: CastLimitCounters(core.CastArray(data["items"])),
	}
}

func (p DescribeLimitCounterMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLimitCountersFromDict(
			p.Items,
		),
	}
}

func (p DescribeLimitCounterMetricsResult) Pointer() *DescribeLimitCounterMetricsResult {
	return &p
}

type DescribeLimitLimitModelMetricsResult struct {
	Items []LimitLimitModel `json:"items"`
}

type DescribeLimitLimitModelMetricsAsyncResult struct {
	result *DescribeLimitLimitModelMetricsResult
	err    error
}

func NewDescribeLimitLimitModelMetricsResultFromJson(data string) DescribeLimitLimitModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLimitLimitModelMetricsResultFromDict(dict)
}

func NewDescribeLimitLimitModelMetricsResultFromDict(data map[string]interface{}) DescribeLimitLimitModelMetricsResult {
	return DescribeLimitLimitModelMetricsResult{
		Items: CastLimitLimitModels(core.CastArray(data["items"])),
	}
}

func (p DescribeLimitLimitModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLimitLimitModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeLimitLimitModelMetricsResult) Pointer() *DescribeLimitLimitModelMetricsResult {
	return &p
}

type GetLimitLimitModelMetricsResult struct {
	Item *LimitLimitModel `json:"item"`
}

type GetLimitLimitModelMetricsAsyncResult struct {
	result *GetLimitLimitModelMetricsResult
	err    error
}

func NewGetLimitLimitModelMetricsResultFromJson(data string) GetLimitLimitModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLimitLimitModelMetricsResultFromDict(dict)
}

func NewGetLimitLimitModelMetricsResultFromDict(data map[string]interface{}) GetLimitLimitModelMetricsResult {
	return GetLimitLimitModelMetricsResult{
		Item: NewLimitLimitModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetLimitLimitModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetLimitLimitModelMetricsResult) Pointer() *GetLimitLimitModelMetricsResult {
	return &p
}

type DescribeLimitNamespaceMetricsResult struct {
	Items []LimitNamespace `json:"items"`
}

type DescribeLimitNamespaceMetricsAsyncResult struct {
	result *DescribeLimitNamespaceMetricsResult
	err    error
}

func NewDescribeLimitNamespaceMetricsResultFromJson(data string) DescribeLimitNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLimitNamespaceMetricsResultFromDict(dict)
}

func NewDescribeLimitNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeLimitNamespaceMetricsResult {
	return DescribeLimitNamespaceMetricsResult{
		Items: CastLimitNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeLimitNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLimitNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeLimitNamespaceMetricsResult) Pointer() *DescribeLimitNamespaceMetricsResult {
	return &p
}

type GetLimitNamespaceMetricsResult struct {
	Item *LimitNamespace `json:"item"`
}

type GetLimitNamespaceMetricsAsyncResult struct {
	result *GetLimitNamespaceMetricsResult
	err    error
}

func NewGetLimitNamespaceMetricsResultFromJson(data string) GetLimitNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLimitNamespaceMetricsResultFromDict(dict)
}

func NewGetLimitNamespaceMetricsResultFromDict(data map[string]interface{}) GetLimitNamespaceMetricsResult {
	return GetLimitNamespaceMetricsResult{
		Item: NewLimitNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetLimitNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetLimitNamespaceMetricsResult) Pointer() *GetLimitNamespaceMetricsResult {
	return &p
}

type DescribeLotteryLotteryMetricsResult struct {
	Items []LotteryLottery `json:"items"`
}

type DescribeLotteryLotteryMetricsAsyncResult struct {
	result *DescribeLotteryLotteryMetricsResult
	err    error
}

func NewDescribeLotteryLotteryMetricsResultFromJson(data string) DescribeLotteryLotteryMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLotteryLotteryMetricsResultFromDict(dict)
}

func NewDescribeLotteryLotteryMetricsResultFromDict(data map[string]interface{}) DescribeLotteryLotteryMetricsResult {
	return DescribeLotteryLotteryMetricsResult{
		Items: CastLotteryLotteries(core.CastArray(data["items"])),
	}
}

func (p DescribeLotteryLotteryMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLotteryLotteriesFromDict(
			p.Items,
		),
	}
}

func (p DescribeLotteryLotteryMetricsResult) Pointer() *DescribeLotteryLotteryMetricsResult {
	return &p
}

type GetLotteryLotteryMetricsResult struct {
	Item *LotteryLottery `json:"item"`
}

type GetLotteryLotteryMetricsAsyncResult struct {
	result *GetLotteryLotteryMetricsResult
	err    error
}

func NewGetLotteryLotteryMetricsResultFromJson(data string) GetLotteryLotteryMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLotteryLotteryMetricsResultFromDict(dict)
}

func NewGetLotteryLotteryMetricsResultFromDict(data map[string]interface{}) GetLotteryLotteryMetricsResult {
	return GetLotteryLotteryMetricsResult{
		Item: NewLotteryLotteryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetLotteryLotteryMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetLotteryLotteryMetricsResult) Pointer() *GetLotteryLotteryMetricsResult {
	return &p
}

type DescribeLotteryNamespaceMetricsResult struct {
	Items []LotteryNamespace `json:"items"`
}

type DescribeLotteryNamespaceMetricsAsyncResult struct {
	result *DescribeLotteryNamespaceMetricsResult
	err    error
}

func NewDescribeLotteryNamespaceMetricsResultFromJson(data string) DescribeLotteryNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLotteryNamespaceMetricsResultFromDict(dict)
}

func NewDescribeLotteryNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeLotteryNamespaceMetricsResult {
	return DescribeLotteryNamespaceMetricsResult{
		Items: CastLotteryNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeLotteryNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLotteryNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeLotteryNamespaceMetricsResult) Pointer() *DescribeLotteryNamespaceMetricsResult {
	return &p
}

type GetLotteryNamespaceMetricsResult struct {
	Item *LotteryNamespace `json:"item"`
}

type GetLotteryNamespaceMetricsAsyncResult struct {
	result *GetLotteryNamespaceMetricsResult
	err    error
}

func NewGetLotteryNamespaceMetricsResultFromJson(data string) GetLotteryNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLotteryNamespaceMetricsResultFromDict(dict)
}

func NewGetLotteryNamespaceMetricsResultFromDict(data map[string]interface{}) GetLotteryNamespaceMetricsResult {
	return GetLotteryNamespaceMetricsResult{
		Item: NewLotteryNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetLotteryNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetLotteryNamespaceMetricsResult) Pointer() *GetLotteryNamespaceMetricsResult {
	return &p
}

type DescribeMatchmakingNamespaceMetricsResult struct {
	Items []MatchmakingNamespace `json:"items"`
}

type DescribeMatchmakingNamespaceMetricsAsyncResult struct {
	result *DescribeMatchmakingNamespaceMetricsResult
	err    error
}

func NewDescribeMatchmakingNamespaceMetricsResultFromJson(data string) DescribeMatchmakingNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMatchmakingNamespaceMetricsResultFromDict(dict)
}

func NewDescribeMatchmakingNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeMatchmakingNamespaceMetricsResult {
	return DescribeMatchmakingNamespaceMetricsResult{
		Items: CastMatchmakingNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeMatchmakingNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMatchmakingNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeMatchmakingNamespaceMetricsResult) Pointer() *DescribeMatchmakingNamespaceMetricsResult {
	return &p
}

type GetMatchmakingNamespaceMetricsResult struct {
	Item *MatchmakingNamespace `json:"item"`
}

type GetMatchmakingNamespaceMetricsAsyncResult struct {
	result *GetMatchmakingNamespaceMetricsResult
	err    error
}

func NewGetMatchmakingNamespaceMetricsResultFromJson(data string) GetMatchmakingNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMatchmakingNamespaceMetricsResultFromDict(dict)
}

func NewGetMatchmakingNamespaceMetricsResultFromDict(data map[string]interface{}) GetMatchmakingNamespaceMetricsResult {
	return GetMatchmakingNamespaceMetricsResult{
		Item: NewMatchmakingNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMatchmakingNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMatchmakingNamespaceMetricsResult) Pointer() *GetMatchmakingNamespaceMetricsResult {
	return &p
}

type DescribeMissionCounterMetricsResult struct {
	Items []MissionCounter `json:"items"`
}

type DescribeMissionCounterMetricsAsyncResult struct {
	result *DescribeMissionCounterMetricsResult
	err    error
}

func NewDescribeMissionCounterMetricsResultFromJson(data string) DescribeMissionCounterMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionCounterMetricsResultFromDict(dict)
}

func NewDescribeMissionCounterMetricsResultFromDict(data map[string]interface{}) DescribeMissionCounterMetricsResult {
	return DescribeMissionCounterMetricsResult{
		Items: CastMissionCounters(core.CastArray(data["items"])),
	}
}

func (p DescribeMissionCounterMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMissionCountersFromDict(
			p.Items,
		),
	}
}

func (p DescribeMissionCounterMetricsResult) Pointer() *DescribeMissionCounterMetricsResult {
	return &p
}

type DescribeMissionMissionGroupModelMetricsResult struct {
	Items []MissionMissionGroupModel `json:"items"`
}

type DescribeMissionMissionGroupModelMetricsAsyncResult struct {
	result *DescribeMissionMissionGroupModelMetricsResult
	err    error
}

func NewDescribeMissionMissionGroupModelMetricsResultFromJson(data string) DescribeMissionMissionGroupModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionMissionGroupModelMetricsResultFromDict(dict)
}

func NewDescribeMissionMissionGroupModelMetricsResultFromDict(data map[string]interface{}) DescribeMissionMissionGroupModelMetricsResult {
	return DescribeMissionMissionGroupModelMetricsResult{
		Items: CastMissionMissionGroupModels(core.CastArray(data["items"])),
	}
}

func (p DescribeMissionMissionGroupModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMissionMissionGroupModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeMissionMissionGroupModelMetricsResult) Pointer() *DescribeMissionMissionGroupModelMetricsResult {
	return &p
}

type GetMissionMissionGroupModelMetricsResult struct {
	Item *MissionMissionGroupModel `json:"item"`
}

type GetMissionMissionGroupModelMetricsAsyncResult struct {
	result *GetMissionMissionGroupModelMetricsResult
	err    error
}

func NewGetMissionMissionGroupModelMetricsResultFromJson(data string) GetMissionMissionGroupModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionMissionGroupModelMetricsResultFromDict(dict)
}

func NewGetMissionMissionGroupModelMetricsResultFromDict(data map[string]interface{}) GetMissionMissionGroupModelMetricsResult {
	return GetMissionMissionGroupModelMetricsResult{
		Item: NewMissionMissionGroupModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMissionMissionGroupModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMissionMissionGroupModelMetricsResult) Pointer() *GetMissionMissionGroupModelMetricsResult {
	return &p
}

type DescribeMissionNamespaceMetricsResult struct {
	Items []MissionNamespace `json:"items"`
}

type DescribeMissionNamespaceMetricsAsyncResult struct {
	result *DescribeMissionNamespaceMetricsResult
	err    error
}

func NewDescribeMissionNamespaceMetricsResultFromJson(data string) DescribeMissionNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionNamespaceMetricsResultFromDict(dict)
}

func NewDescribeMissionNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeMissionNamespaceMetricsResult {
	return DescribeMissionNamespaceMetricsResult{
		Items: CastMissionNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeMissionNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMissionNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeMissionNamespaceMetricsResult) Pointer() *DescribeMissionNamespaceMetricsResult {
	return &p
}

type GetMissionNamespaceMetricsResult struct {
	Item *MissionNamespace `json:"item"`
}

type GetMissionNamespaceMetricsAsyncResult struct {
	result *GetMissionNamespaceMetricsResult
	err    error
}

func NewGetMissionNamespaceMetricsResultFromJson(data string) GetMissionNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionNamespaceMetricsResultFromDict(dict)
}

func NewGetMissionNamespaceMetricsResultFromDict(data map[string]interface{}) GetMissionNamespaceMetricsResult {
	return GetMissionNamespaceMetricsResult{
		Item: NewMissionNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMissionNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMissionNamespaceMetricsResult) Pointer() *GetMissionNamespaceMetricsResult {
	return &p
}

type DescribeMoneyWalletMetricsResult struct {
	Items []MoneyWallet `json:"items"`
}

type DescribeMoneyWalletMetricsAsyncResult struct {
	result *DescribeMoneyWalletMetricsResult
	err    error
}

func NewDescribeMoneyWalletMetricsResultFromJson(data string) DescribeMoneyWalletMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoneyWalletMetricsResultFromDict(dict)
}

func NewDescribeMoneyWalletMetricsResultFromDict(data map[string]interface{}) DescribeMoneyWalletMetricsResult {
	return DescribeMoneyWalletMetricsResult{
		Items: CastMoneyWallets(core.CastArray(data["items"])),
	}
}

func (p DescribeMoneyWalletMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMoneyWalletsFromDict(
			p.Items,
		),
	}
}

func (p DescribeMoneyWalletMetricsResult) Pointer() *DescribeMoneyWalletMetricsResult {
	return &p
}

type DescribeMoneyReceiptMetricsResult struct {
	Items []MoneyReceipt `json:"items"`
}

type DescribeMoneyReceiptMetricsAsyncResult struct {
	result *DescribeMoneyReceiptMetricsResult
	err    error
}

func NewDescribeMoneyReceiptMetricsResultFromJson(data string) DescribeMoneyReceiptMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoneyReceiptMetricsResultFromDict(dict)
}

func NewDescribeMoneyReceiptMetricsResultFromDict(data map[string]interface{}) DescribeMoneyReceiptMetricsResult {
	return DescribeMoneyReceiptMetricsResult{
		Items: CastMoneyReceipts(core.CastArray(data["items"])),
	}
}

func (p DescribeMoneyReceiptMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMoneyReceiptsFromDict(
			p.Items,
		),
	}
}

func (p DescribeMoneyReceiptMetricsResult) Pointer() *DescribeMoneyReceiptMetricsResult {
	return &p
}

type DescribeMoneyNamespaceMetricsResult struct {
	Items []MoneyNamespace `json:"items"`
}

type DescribeMoneyNamespaceMetricsAsyncResult struct {
	result *DescribeMoneyNamespaceMetricsResult
	err    error
}

func NewDescribeMoneyNamespaceMetricsResultFromJson(data string) DescribeMoneyNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoneyNamespaceMetricsResultFromDict(dict)
}

func NewDescribeMoneyNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeMoneyNamespaceMetricsResult {
	return DescribeMoneyNamespaceMetricsResult{
		Items: CastMoneyNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeMoneyNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMoneyNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeMoneyNamespaceMetricsResult) Pointer() *DescribeMoneyNamespaceMetricsResult {
	return &p
}

type GetMoneyNamespaceMetricsResult struct {
	Item *MoneyNamespace `json:"item"`
}

type GetMoneyNamespaceMetricsAsyncResult struct {
	result *GetMoneyNamespaceMetricsResult
	err    error
}

func NewGetMoneyNamespaceMetricsResultFromJson(data string) GetMoneyNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMoneyNamespaceMetricsResultFromDict(dict)
}

func NewGetMoneyNamespaceMetricsResultFromDict(data map[string]interface{}) GetMoneyNamespaceMetricsResult {
	return GetMoneyNamespaceMetricsResult{
		Item: NewMoneyNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMoneyNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMoneyNamespaceMetricsResult) Pointer() *GetMoneyNamespaceMetricsResult {
	return &p
}

type DescribeQuestQuestModelMetricsResult struct {
	Items []QuestQuestModel `json:"items"`
}

type DescribeQuestQuestModelMetricsAsyncResult struct {
	result *DescribeQuestQuestModelMetricsResult
	err    error
}

func NewDescribeQuestQuestModelMetricsResultFromJson(data string) DescribeQuestQuestModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeQuestQuestModelMetricsResultFromDict(dict)
}

func NewDescribeQuestQuestModelMetricsResultFromDict(data map[string]interface{}) DescribeQuestQuestModelMetricsResult {
	return DescribeQuestQuestModelMetricsResult{
		Items: CastQuestQuestModels(core.CastArray(data["items"])),
	}
}

func (p DescribeQuestQuestModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastQuestQuestModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeQuestQuestModelMetricsResult) Pointer() *DescribeQuestQuestModelMetricsResult {
	return &p
}

type GetQuestQuestModelMetricsResult struct {
	Item *QuestQuestModel `json:"item"`
}

type GetQuestQuestModelMetricsAsyncResult struct {
	result *GetQuestQuestModelMetricsResult
	err    error
}

func NewGetQuestQuestModelMetricsResultFromJson(data string) GetQuestQuestModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetQuestQuestModelMetricsResultFromDict(dict)
}

func NewGetQuestQuestModelMetricsResultFromDict(data map[string]interface{}) GetQuestQuestModelMetricsResult {
	return GetQuestQuestModelMetricsResult{
		Item: NewQuestQuestModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetQuestQuestModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetQuestQuestModelMetricsResult) Pointer() *GetQuestQuestModelMetricsResult {
	return &p
}

type DescribeQuestQuestGroupModelMetricsResult struct {
	Items []QuestQuestGroupModel `json:"items"`
}

type DescribeQuestQuestGroupModelMetricsAsyncResult struct {
	result *DescribeQuestQuestGroupModelMetricsResult
	err    error
}

func NewDescribeQuestQuestGroupModelMetricsResultFromJson(data string) DescribeQuestQuestGroupModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeQuestQuestGroupModelMetricsResultFromDict(dict)
}

func NewDescribeQuestQuestGroupModelMetricsResultFromDict(data map[string]interface{}) DescribeQuestQuestGroupModelMetricsResult {
	return DescribeQuestQuestGroupModelMetricsResult{
		Items: CastQuestQuestGroupModels(core.CastArray(data["items"])),
	}
}

func (p DescribeQuestQuestGroupModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastQuestQuestGroupModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeQuestQuestGroupModelMetricsResult) Pointer() *DescribeQuestQuestGroupModelMetricsResult {
	return &p
}

type GetQuestQuestGroupModelMetricsResult struct {
	Item *QuestQuestGroupModel `json:"item"`
}

type GetQuestQuestGroupModelMetricsAsyncResult struct {
	result *GetQuestQuestGroupModelMetricsResult
	err    error
}

func NewGetQuestQuestGroupModelMetricsResultFromJson(data string) GetQuestQuestGroupModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetQuestQuestGroupModelMetricsResultFromDict(dict)
}

func NewGetQuestQuestGroupModelMetricsResultFromDict(data map[string]interface{}) GetQuestQuestGroupModelMetricsResult {
	return GetQuestQuestGroupModelMetricsResult{
		Item: NewQuestQuestGroupModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetQuestQuestGroupModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetQuestQuestGroupModelMetricsResult) Pointer() *GetQuestQuestGroupModelMetricsResult {
	return &p
}

type DescribeQuestNamespaceMetricsResult struct {
	Items []QuestNamespace `json:"items"`
}

type DescribeQuestNamespaceMetricsAsyncResult struct {
	result *DescribeQuestNamespaceMetricsResult
	err    error
}

func NewDescribeQuestNamespaceMetricsResultFromJson(data string) DescribeQuestNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeQuestNamespaceMetricsResultFromDict(dict)
}

func NewDescribeQuestNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeQuestNamespaceMetricsResult {
	return DescribeQuestNamespaceMetricsResult{
		Items: CastQuestNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeQuestNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastQuestNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeQuestNamespaceMetricsResult) Pointer() *DescribeQuestNamespaceMetricsResult {
	return &p
}

type GetQuestNamespaceMetricsResult struct {
	Item *QuestNamespace `json:"item"`
}

type GetQuestNamespaceMetricsAsyncResult struct {
	result *GetQuestNamespaceMetricsResult
	err    error
}

func NewGetQuestNamespaceMetricsResultFromJson(data string) GetQuestNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetQuestNamespaceMetricsResultFromDict(dict)
}

func NewGetQuestNamespaceMetricsResultFromDict(data map[string]interface{}) GetQuestNamespaceMetricsResult {
	return GetQuestNamespaceMetricsResult{
		Item: NewQuestNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetQuestNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetQuestNamespaceMetricsResult) Pointer() *GetQuestNamespaceMetricsResult {
	return &p
}

type DescribeRankingCategoryModelMetricsResult struct {
	Items []RankingCategoryModel `json:"items"`
}

type DescribeRankingCategoryModelMetricsAsyncResult struct {
	result *DescribeRankingCategoryModelMetricsResult
	err    error
}

func NewDescribeRankingCategoryModelMetricsResultFromJson(data string) DescribeRankingCategoryModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRankingCategoryModelMetricsResultFromDict(dict)
}

func NewDescribeRankingCategoryModelMetricsResultFromDict(data map[string]interface{}) DescribeRankingCategoryModelMetricsResult {
	return DescribeRankingCategoryModelMetricsResult{
		Items: CastRankingCategoryModels(core.CastArray(data["items"])),
	}
}

func (p DescribeRankingCategoryModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRankingCategoryModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeRankingCategoryModelMetricsResult) Pointer() *DescribeRankingCategoryModelMetricsResult {
	return &p
}

type GetRankingCategoryModelMetricsResult struct {
	Item *RankingCategoryModel `json:"item"`
}

type GetRankingCategoryModelMetricsAsyncResult struct {
	result *GetRankingCategoryModelMetricsResult
	err    error
}

func NewGetRankingCategoryModelMetricsResultFromJson(data string) GetRankingCategoryModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRankingCategoryModelMetricsResultFromDict(dict)
}

func NewGetRankingCategoryModelMetricsResultFromDict(data map[string]interface{}) GetRankingCategoryModelMetricsResult {
	return GetRankingCategoryModelMetricsResult{
		Item: NewRankingCategoryModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRankingCategoryModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRankingCategoryModelMetricsResult) Pointer() *GetRankingCategoryModelMetricsResult {
	return &p
}

type DescribeRankingNamespaceMetricsResult struct {
	Items []RankingNamespace `json:"items"`
}

type DescribeRankingNamespaceMetricsAsyncResult struct {
	result *DescribeRankingNamespaceMetricsResult
	err    error
}

func NewDescribeRankingNamespaceMetricsResultFromJson(data string) DescribeRankingNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRankingNamespaceMetricsResultFromDict(dict)
}

func NewDescribeRankingNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeRankingNamespaceMetricsResult {
	return DescribeRankingNamespaceMetricsResult{
		Items: CastRankingNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeRankingNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRankingNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeRankingNamespaceMetricsResult) Pointer() *DescribeRankingNamespaceMetricsResult {
	return &p
}

type GetRankingNamespaceMetricsResult struct {
	Item *RankingNamespace `json:"item"`
}

type GetRankingNamespaceMetricsAsyncResult struct {
	result *GetRankingNamespaceMetricsResult
	err    error
}

func NewGetRankingNamespaceMetricsResultFromJson(data string) GetRankingNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRankingNamespaceMetricsResultFromDict(dict)
}

func NewGetRankingNamespaceMetricsResultFromDict(data map[string]interface{}) GetRankingNamespaceMetricsResult {
	return GetRankingNamespaceMetricsResult{
		Item: NewRankingNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRankingNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRankingNamespaceMetricsResult) Pointer() *GetRankingNamespaceMetricsResult {
	return &p
}

type DescribeShowcaseDisplayItemMetricsResult struct {
	Items []ShowcaseDisplayItem `json:"items"`
}

type DescribeShowcaseDisplayItemMetricsAsyncResult struct {
	result *DescribeShowcaseDisplayItemMetricsResult
	err    error
}

func NewDescribeShowcaseDisplayItemMetricsResultFromJson(data string) DescribeShowcaseDisplayItemMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeShowcaseDisplayItemMetricsResultFromDict(dict)
}

func NewDescribeShowcaseDisplayItemMetricsResultFromDict(data map[string]interface{}) DescribeShowcaseDisplayItemMetricsResult {
	return DescribeShowcaseDisplayItemMetricsResult{
		Items: CastShowcaseDisplayItems(core.CastArray(data["items"])),
	}
}

func (p DescribeShowcaseDisplayItemMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastShowcaseDisplayItemsFromDict(
			p.Items,
		),
	}
}

func (p DescribeShowcaseDisplayItemMetricsResult) Pointer() *DescribeShowcaseDisplayItemMetricsResult {
	return &p
}

type GetShowcaseDisplayItemMetricsResult struct {
	Item *ShowcaseDisplayItem `json:"item"`
}

type GetShowcaseDisplayItemMetricsAsyncResult struct {
	result *GetShowcaseDisplayItemMetricsResult
	err    error
}

func NewGetShowcaseDisplayItemMetricsResultFromJson(data string) GetShowcaseDisplayItemMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetShowcaseDisplayItemMetricsResultFromDict(dict)
}

func NewGetShowcaseDisplayItemMetricsResultFromDict(data map[string]interface{}) GetShowcaseDisplayItemMetricsResult {
	return GetShowcaseDisplayItemMetricsResult{
		Item: NewShowcaseDisplayItemFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetShowcaseDisplayItemMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetShowcaseDisplayItemMetricsResult) Pointer() *GetShowcaseDisplayItemMetricsResult {
	return &p
}

type DescribeShowcaseShowcaseMetricsResult struct {
	Items []ShowcaseShowcase `json:"items"`
}

type DescribeShowcaseShowcaseMetricsAsyncResult struct {
	result *DescribeShowcaseShowcaseMetricsResult
	err    error
}

func NewDescribeShowcaseShowcaseMetricsResultFromJson(data string) DescribeShowcaseShowcaseMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeShowcaseShowcaseMetricsResultFromDict(dict)
}

func NewDescribeShowcaseShowcaseMetricsResultFromDict(data map[string]interface{}) DescribeShowcaseShowcaseMetricsResult {
	return DescribeShowcaseShowcaseMetricsResult{
		Items: CastShowcaseShowcases(core.CastArray(data["items"])),
	}
}

func (p DescribeShowcaseShowcaseMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastShowcaseShowcasesFromDict(
			p.Items,
		),
	}
}

func (p DescribeShowcaseShowcaseMetricsResult) Pointer() *DescribeShowcaseShowcaseMetricsResult {
	return &p
}

type GetShowcaseShowcaseMetricsResult struct {
	Item *ShowcaseShowcase `json:"item"`
}

type GetShowcaseShowcaseMetricsAsyncResult struct {
	result *GetShowcaseShowcaseMetricsResult
	err    error
}

func NewGetShowcaseShowcaseMetricsResultFromJson(data string) GetShowcaseShowcaseMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetShowcaseShowcaseMetricsResultFromDict(dict)
}

func NewGetShowcaseShowcaseMetricsResultFromDict(data map[string]interface{}) GetShowcaseShowcaseMetricsResult {
	return GetShowcaseShowcaseMetricsResult{
		Item: NewShowcaseShowcaseFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetShowcaseShowcaseMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetShowcaseShowcaseMetricsResult) Pointer() *GetShowcaseShowcaseMetricsResult {
	return &p
}

type DescribeShowcaseNamespaceMetricsResult struct {
	Items []ShowcaseNamespace `json:"items"`
}

type DescribeShowcaseNamespaceMetricsAsyncResult struct {
	result *DescribeShowcaseNamespaceMetricsResult
	err    error
}

func NewDescribeShowcaseNamespaceMetricsResultFromJson(data string) DescribeShowcaseNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeShowcaseNamespaceMetricsResultFromDict(dict)
}

func NewDescribeShowcaseNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeShowcaseNamespaceMetricsResult {
	return DescribeShowcaseNamespaceMetricsResult{
		Items: CastShowcaseNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeShowcaseNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastShowcaseNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeShowcaseNamespaceMetricsResult) Pointer() *DescribeShowcaseNamespaceMetricsResult {
	return &p
}

type GetShowcaseNamespaceMetricsResult struct {
	Item *ShowcaseNamespace `json:"item"`
}

type GetShowcaseNamespaceMetricsAsyncResult struct {
	result *GetShowcaseNamespaceMetricsResult
	err    error
}

func NewGetShowcaseNamespaceMetricsResultFromJson(data string) GetShowcaseNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetShowcaseNamespaceMetricsResultFromDict(dict)
}

func NewGetShowcaseNamespaceMetricsResultFromDict(data map[string]interface{}) GetShowcaseNamespaceMetricsResult {
	return GetShowcaseNamespaceMetricsResult{
		Item: NewShowcaseNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetShowcaseNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetShowcaseNamespaceMetricsResult) Pointer() *GetShowcaseNamespaceMetricsResult {
	return &p
}

type DescribeStaminaStaminaModelMetricsResult struct {
	Items []StaminaStaminaModel `json:"items"`
}

type DescribeStaminaStaminaModelMetricsAsyncResult struct {
	result *DescribeStaminaStaminaModelMetricsResult
	err    error
}

func NewDescribeStaminaStaminaModelMetricsResultFromJson(data string) DescribeStaminaStaminaModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminaStaminaModelMetricsResultFromDict(dict)
}

func NewDescribeStaminaStaminaModelMetricsResultFromDict(data map[string]interface{}) DescribeStaminaStaminaModelMetricsResult {
	return DescribeStaminaStaminaModelMetricsResult{
		Items: CastStaminaStaminaModels(core.CastArray(data["items"])),
	}
}

func (p DescribeStaminaStaminaModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminaStaminaModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeStaminaStaminaModelMetricsResult) Pointer() *DescribeStaminaStaminaModelMetricsResult {
	return &p
}

type GetStaminaStaminaModelMetricsResult struct {
	Item *StaminaStaminaModel `json:"item"`
}

type GetStaminaStaminaModelMetricsAsyncResult struct {
	result *GetStaminaStaminaModelMetricsResult
	err    error
}

func NewGetStaminaStaminaModelMetricsResultFromJson(data string) GetStaminaStaminaModelMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaStaminaModelMetricsResultFromDict(dict)
}

func NewGetStaminaStaminaModelMetricsResultFromDict(data map[string]interface{}) GetStaminaStaminaModelMetricsResult {
	return GetStaminaStaminaModelMetricsResult{
		Item: NewStaminaStaminaModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStaminaStaminaModelMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetStaminaStaminaModelMetricsResult) Pointer() *GetStaminaStaminaModelMetricsResult {
	return &p
}

type DescribeStaminaNamespaceMetricsResult struct {
	Items []StaminaNamespace `json:"items"`
}

type DescribeStaminaNamespaceMetricsAsyncResult struct {
	result *DescribeStaminaNamespaceMetricsResult
	err    error
}

func NewDescribeStaminaNamespaceMetricsResultFromJson(data string) DescribeStaminaNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminaNamespaceMetricsResultFromDict(dict)
}

func NewDescribeStaminaNamespaceMetricsResultFromDict(data map[string]interface{}) DescribeStaminaNamespaceMetricsResult {
	return DescribeStaminaNamespaceMetricsResult{
		Items: CastStaminaNamespaces(core.CastArray(data["items"])),
	}
}

func (p DescribeStaminaNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminaNamespacesFromDict(
			p.Items,
		),
	}
}

func (p DescribeStaminaNamespaceMetricsResult) Pointer() *DescribeStaminaNamespaceMetricsResult {
	return &p
}

type GetStaminaNamespaceMetricsResult struct {
	Item *StaminaNamespace `json:"item"`
}

type GetStaminaNamespaceMetricsAsyncResult struct {
	result *GetStaminaNamespaceMetricsResult
	err    error
}

func NewGetStaminaNamespaceMetricsResultFromJson(data string) GetStaminaNamespaceMetricsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaNamespaceMetricsResultFromDict(dict)
}

func NewGetStaminaNamespaceMetricsResultFromDict(data map[string]interface{}) GetStaminaNamespaceMetricsResult {
	return GetStaminaNamespaceMetricsResult{
		Item: NewStaminaNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStaminaNamespaceMetricsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetStaminaNamespaceMetricsResult) Pointer() *GetStaminaNamespaceMetricsResult {
	return &p
}
