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

package exchange

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId          *string             `json:"namespaceId"`
	Name                 *string             `json:"name"`
	Description          *string             `json:"description"`
	EnableDirectExchange *bool               `json:"enableDirectExchange"`
	EnableAwaitExchange  *bool               `json:"enableAwaitExchange"`
	TransactionSetting   *TransactionSetting `json:"transactionSetting"`
	ExchangeScript       *ScriptSetting      `json:"exchangeScript"`
	LogSetting           *LogSetting         `json:"logSetting"`
	CreatedAt            *int64              `json:"createdAt"`
	UpdatedAt            *int64              `json:"updatedAt"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId    *string `json:"keyId"`
	Revision *int64  `json:"revision"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:          core.CastString(data["namespaceId"]),
		Name:                 core.CastString(data["name"]),
		Description:          core.CastString(data["description"]),
		EnableDirectExchange: core.CastBool(data["enableDirectExchange"]),
		EnableAwaitExchange:  core.CastBool(data["enableAwaitExchange"]),
		TransactionSetting:   NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ExchangeScript:       NewScriptSettingFromDict(core.CastMap(data["exchangeScript"])).Pointer(),
		LogSetting:           NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:            core.CastInt64(data["createdAt"]),
		UpdatedAt:            core.CastInt64(data["updatedAt"]),
		QueueNamespaceId:     core.CastString(data["queueNamespaceId"]),
		KeyId:                core.CastString(data["keyId"]),
		Revision:             core.CastInt64(data["revision"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var enableDirectExchange *bool
	if p.EnableDirectExchange != nil {
		enableDirectExchange = p.EnableDirectExchange
	}
	var enableAwaitExchange *bool
	if p.EnableAwaitExchange != nil {
		enableAwaitExchange = p.EnableAwaitExchange
	}
	var transactionSetting map[string]interface{}
	if p.TransactionSetting != nil {
		transactionSetting = p.TransactionSetting.ToDict()
	}
	var exchangeScript map[string]interface{}
	if p.ExchangeScript != nil {
		exchangeScript = p.ExchangeScript.ToDict()
	}
	var logSetting map[string]interface{}
	if p.LogSetting != nil {
		logSetting = p.LogSetting.ToDict()
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId":          namespaceId,
		"name":                 name,
		"description":          description,
		"enableDirectExchange": enableDirectExchange,
		"enableAwaitExchange":  enableAwaitExchange,
		"transactionSetting":   transactionSetting,
		"exchangeScript":       exchangeScript,
		"logSetting":           logSetting,
		"createdAt":            createdAt,
		"updatedAt":            updatedAt,
		"queueNamespaceId":     queueNamespaceId,
		"keyId":                keyId,
		"revision":             revision,
	}
}

func (p Namespace) Pointer() *Namespace {
	return &p
}

func CastNamespaces(data []interface{}) []Namespace {
	v := make([]Namespace, 0)
	for _, d := range data {
		v = append(v, NewNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNamespacesFromDict(data []Namespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RateModel struct {
	RateModelId        *string         `json:"rateModelId"`
	Name               *string         `json:"name"`
	Metadata           *string         `json:"metadata"`
	ConsumeActions     []ConsumeAction `json:"consumeActions"`
	TimingType         *string         `json:"timingType"`
	LockTime           *int32          `json:"lockTime"`
	EnableSkip         *bool           `json:"enableSkip"`
	SkipConsumeActions []ConsumeAction `json:"skipConsumeActions"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
}

func NewRateModelFromJson(data string) RateModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRateModelFromDict(dict)
}

func NewRateModelFromDict(data map[string]interface{}) RateModel {
	return RateModel{
		RateModelId:        core.CastString(data["rateModelId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		ConsumeActions:     CastConsumeActions(core.CastArray(data["consumeActions"])),
		TimingType:         core.CastString(data["timingType"]),
		LockTime:           core.CastInt32(data["lockTime"]),
		EnableSkip:         core.CastBool(data["enableSkip"]),
		SkipConsumeActions: CastConsumeActions(core.CastArray(data["skipConsumeActions"])),
		AcquireActions:     CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p RateModel) ToDict() map[string]interface{} {

	var rateModelId *string
	if p.RateModelId != nil {
		rateModelId = p.RateModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var timingType *string
	if p.TimingType != nil {
		timingType = p.TimingType
	}
	var lockTime *int32
	if p.LockTime != nil {
		lockTime = p.LockTime
	}
	var enableSkip *bool
	if p.EnableSkip != nil {
		enableSkip = p.EnableSkip
	}
	var skipConsumeActions []interface{}
	if p.SkipConsumeActions != nil {
		skipConsumeActions = CastConsumeActionsFromDict(
			p.SkipConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	return map[string]interface{}{
		"rateModelId":        rateModelId,
		"name":               name,
		"metadata":           metadata,
		"consumeActions":     consumeActions,
		"timingType":         timingType,
		"lockTime":           lockTime,
		"enableSkip":         enableSkip,
		"skipConsumeActions": skipConsumeActions,
		"acquireActions":     acquireActions,
	}
}

func (p RateModel) Pointer() *RateModel {
	return &p
}

func CastRateModels(data []interface{}) []RateModel {
	v := make([]RateModel, 0)
	for _, d := range data {
		v = append(v, NewRateModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRateModelsFromDict(data []RateModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RateModelMaster struct {
	RateModelId        *string         `json:"rateModelId"`
	Name               *string         `json:"name"`
	Description        *string         `json:"description"`
	Metadata           *string         `json:"metadata"`
	ConsumeActions     []ConsumeAction `json:"consumeActions"`
	TimingType         *string         `json:"timingType"`
	LockTime           *int32          `json:"lockTime"`
	EnableSkip         *bool           `json:"enableSkip"`
	SkipConsumeActions []ConsumeAction `json:"skipConsumeActions"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
	CreatedAt          *int64          `json:"createdAt"`
	UpdatedAt          *int64          `json:"updatedAt"`
	Revision           *int64          `json:"revision"`
}

func NewRateModelMasterFromJson(data string) RateModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRateModelMasterFromDict(dict)
}

func NewRateModelMasterFromDict(data map[string]interface{}) RateModelMaster {
	return RateModelMaster{
		RateModelId:        core.CastString(data["rateModelId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		ConsumeActions:     CastConsumeActions(core.CastArray(data["consumeActions"])),
		TimingType:         core.CastString(data["timingType"]),
		LockTime:           core.CastInt32(data["lockTime"]),
		EnableSkip:         core.CastBool(data["enableSkip"]),
		SkipConsumeActions: CastConsumeActions(core.CastArray(data["skipConsumeActions"])),
		AcquireActions:     CastAcquireActions(core.CastArray(data["acquireActions"])),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		Revision:           core.CastInt64(data["revision"]),
	}
}

func (p RateModelMaster) ToDict() map[string]interface{} {

	var rateModelId *string
	if p.RateModelId != nil {
		rateModelId = p.RateModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var timingType *string
	if p.TimingType != nil {
		timingType = p.TimingType
	}
	var lockTime *int32
	if p.LockTime != nil {
		lockTime = p.LockTime
	}
	var enableSkip *bool
	if p.EnableSkip != nil {
		enableSkip = p.EnableSkip
	}
	var skipConsumeActions []interface{}
	if p.SkipConsumeActions != nil {
		skipConsumeActions = CastConsumeActionsFromDict(
			p.SkipConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
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
		"rateModelId":        rateModelId,
		"name":               name,
		"description":        description,
		"metadata":           metadata,
		"consumeActions":     consumeActions,
		"timingType":         timingType,
		"lockTime":           lockTime,
		"enableSkip":         enableSkip,
		"skipConsumeActions": skipConsumeActions,
		"acquireActions":     acquireActions,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"revision":           revision,
	}
}

func (p RateModelMaster) Pointer() *RateModelMaster {
	return &p
}

func CastRateModelMasters(data []interface{}) []RateModelMaster {
	v := make([]RateModelMaster, 0)
	for _, d := range data {
		v = append(v, NewRateModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRateModelMastersFromDict(data []RateModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type IncrementalRateModel struct {
	IncrementalRateModelId *string         `json:"incrementalRateModelId"`
	Name                   *string         `json:"name"`
	Metadata               *string         `json:"metadata"`
	ConsumeAction          *ConsumeAction  `json:"consumeAction"`
	CalculateType          *string         `json:"calculateType"`
	BaseValue              *int64          `json:"baseValue"`
	CoefficientValue       *int64          `json:"coefficientValue"`
	CalculateScriptId      *string         `json:"calculateScriptId"`
	ExchangeCountId        *string         `json:"exchangeCountId"`
	MaximumExchangeCount   *int32          `json:"maximumExchangeCount"`
	AcquireActions         []AcquireAction `json:"acquireActions"`
}

func NewIncrementalRateModelFromJson(data string) IncrementalRateModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementalRateModelFromDict(dict)
}

func NewIncrementalRateModelFromDict(data map[string]interface{}) IncrementalRateModel {
	return IncrementalRateModel{
		IncrementalRateModelId: core.CastString(data["incrementalRateModelId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		ConsumeAction:          NewConsumeActionFromDict(core.CastMap(data["consumeAction"])).Pointer(),
		CalculateType:          core.CastString(data["calculateType"]),
		BaseValue:              core.CastInt64(data["baseValue"]),
		CoefficientValue:       core.CastInt64(data["coefficientValue"]),
		CalculateScriptId:      core.CastString(data["calculateScriptId"]),
		ExchangeCountId:        core.CastString(data["exchangeCountId"]),
		MaximumExchangeCount:   core.CastInt32(data["maximumExchangeCount"]),
		AcquireActions:         CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p IncrementalRateModel) ToDict() map[string]interface{} {

	var incrementalRateModelId *string
	if p.IncrementalRateModelId != nil {
		incrementalRateModelId = p.IncrementalRateModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeAction map[string]interface{}
	if p.ConsumeAction != nil {
		consumeAction = p.ConsumeAction.ToDict()
	}
	var calculateType *string
	if p.CalculateType != nil {
		calculateType = p.CalculateType
	}
	var baseValue *int64
	if p.BaseValue != nil {
		baseValue = p.BaseValue
	}
	var coefficientValue *int64
	if p.CoefficientValue != nil {
		coefficientValue = p.CoefficientValue
	}
	var calculateScriptId *string
	if p.CalculateScriptId != nil {
		calculateScriptId = p.CalculateScriptId
	}
	var exchangeCountId *string
	if p.ExchangeCountId != nil {
		exchangeCountId = p.ExchangeCountId
	}
	var maximumExchangeCount *int32
	if p.MaximumExchangeCount != nil {
		maximumExchangeCount = p.MaximumExchangeCount
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	return map[string]interface{}{
		"incrementalRateModelId": incrementalRateModelId,
		"name":                   name,
		"metadata":               metadata,
		"consumeAction":          consumeAction,
		"calculateType":          calculateType,
		"baseValue":              baseValue,
		"coefficientValue":       coefficientValue,
		"calculateScriptId":      calculateScriptId,
		"exchangeCountId":        exchangeCountId,
		"maximumExchangeCount":   maximumExchangeCount,
		"acquireActions":         acquireActions,
	}
}

func (p IncrementalRateModel) Pointer() *IncrementalRateModel {
	return &p
}

func CastIncrementalRateModels(data []interface{}) []IncrementalRateModel {
	v := make([]IncrementalRateModel, 0)
	for _, d := range data {
		v = append(v, NewIncrementalRateModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastIncrementalRateModelsFromDict(data []IncrementalRateModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type IncrementalRateModelMaster struct {
	IncrementalRateModelId *string         `json:"incrementalRateModelId"`
	Name                   *string         `json:"name"`
	Description            *string         `json:"description"`
	Metadata               *string         `json:"metadata"`
	ConsumeAction          *ConsumeAction  `json:"consumeAction"`
	CalculateType          *string         `json:"calculateType"`
	BaseValue              *int64          `json:"baseValue"`
	CoefficientValue       *int64          `json:"coefficientValue"`
	CalculateScriptId      *string         `json:"calculateScriptId"`
	ExchangeCountId        *string         `json:"exchangeCountId"`
	MaximumExchangeCount   *int32          `json:"maximumExchangeCount"`
	AcquireActions         []AcquireAction `json:"acquireActions"`
	CreatedAt              *int64          `json:"createdAt"`
	UpdatedAt              *int64          `json:"updatedAt"`
	Revision               *int64          `json:"revision"`
}

func NewIncrementalRateModelMasterFromJson(data string) IncrementalRateModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementalRateModelMasterFromDict(dict)
}

func NewIncrementalRateModelMasterFromDict(data map[string]interface{}) IncrementalRateModelMaster {
	return IncrementalRateModelMaster{
		IncrementalRateModelId: core.CastString(data["incrementalRateModelId"]),
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		Metadata:               core.CastString(data["metadata"]),
		ConsumeAction:          NewConsumeActionFromDict(core.CastMap(data["consumeAction"])).Pointer(),
		CalculateType:          core.CastString(data["calculateType"]),
		BaseValue:              core.CastInt64(data["baseValue"]),
		CoefficientValue:       core.CastInt64(data["coefficientValue"]),
		CalculateScriptId:      core.CastString(data["calculateScriptId"]),
		ExchangeCountId:        core.CastString(data["exchangeCountId"]),
		MaximumExchangeCount:   core.CastInt32(data["maximumExchangeCount"]),
		AcquireActions:         CastAcquireActions(core.CastArray(data["acquireActions"])),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
		Revision:               core.CastInt64(data["revision"]),
	}
}

func (p IncrementalRateModelMaster) ToDict() map[string]interface{} {

	var incrementalRateModelId *string
	if p.IncrementalRateModelId != nil {
		incrementalRateModelId = p.IncrementalRateModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeAction map[string]interface{}
	if p.ConsumeAction != nil {
		consumeAction = p.ConsumeAction.ToDict()
	}
	var calculateType *string
	if p.CalculateType != nil {
		calculateType = p.CalculateType
	}
	var baseValue *int64
	if p.BaseValue != nil {
		baseValue = p.BaseValue
	}
	var coefficientValue *int64
	if p.CoefficientValue != nil {
		coefficientValue = p.CoefficientValue
	}
	var calculateScriptId *string
	if p.CalculateScriptId != nil {
		calculateScriptId = p.CalculateScriptId
	}
	var exchangeCountId *string
	if p.ExchangeCountId != nil {
		exchangeCountId = p.ExchangeCountId
	}
	var maximumExchangeCount *int32
	if p.MaximumExchangeCount != nil {
		maximumExchangeCount = p.MaximumExchangeCount
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
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
		"incrementalRateModelId": incrementalRateModelId,
		"name":                   name,
		"description":            description,
		"metadata":               metadata,
		"consumeAction":          consumeAction,
		"calculateType":          calculateType,
		"baseValue":              baseValue,
		"coefficientValue":       coefficientValue,
		"calculateScriptId":      calculateScriptId,
		"exchangeCountId":        exchangeCountId,
		"maximumExchangeCount":   maximumExchangeCount,
		"acquireActions":         acquireActions,
		"createdAt":              createdAt,
		"updatedAt":              updatedAt,
		"revision":               revision,
	}
}

func (p IncrementalRateModelMaster) Pointer() *IncrementalRateModelMaster {
	return &p
}

func CastIncrementalRateModelMasters(data []interface{}) []IncrementalRateModelMaster {
	v := make([]IncrementalRateModelMaster, 0)
	for _, d := range data {
		v = append(v, NewIncrementalRateModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastIncrementalRateModelMastersFromDict(data []IncrementalRateModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentRateMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentRateMasterFromJson(data string) CurrentRateMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentRateMasterFromDict(dict)
}

func NewCurrentRateMasterFromDict(data map[string]interface{}) CurrentRateMaster {
	return CurrentRateMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentRateMaster) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var settings *string
	if p.Settings != nil {
		settings = p.Settings
	}
	return map[string]interface{}{
		"namespaceId": namespaceId,
		"settings":    settings,
	}
}

func (p CurrentRateMaster) Pointer() *CurrentRateMaster {
	return &p
}

func CastCurrentRateMasters(data []interface{}) []CurrentRateMaster {
	v := make([]CurrentRateMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentRateMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentRateMastersFromDict(data []CurrentRateMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Await struct {
	AwaitId     *string `json:"awaitId"`
	UserId      *string `json:"userId"`
	RateName    *string `json:"rateName"`
	Name        *string `json:"name"`
	Count       *int32  `json:"count"`
	ExchangedAt *int64  `json:"exchangedAt"`
	Revision    *int64  `json:"revision"`
}

func NewAwaitFromJson(data string) Await {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAwaitFromDict(dict)
}

func NewAwaitFromDict(data map[string]interface{}) Await {
	return Await{
		AwaitId:     core.CastString(data["awaitId"]),
		UserId:      core.CastString(data["userId"]),
		RateName:    core.CastString(data["rateName"]),
		Name:        core.CastString(data["name"]),
		Count:       core.CastInt32(data["count"]),
		ExchangedAt: core.CastInt64(data["exchangedAt"]),
		Revision:    core.CastInt64(data["revision"]),
	}
}

func (p Await) ToDict() map[string]interface{} {

	var awaitId *string
	if p.AwaitId != nil {
		awaitId = p.AwaitId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var rateName *string
	if p.RateName != nil {
		rateName = p.RateName
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var count *int32
	if p.Count != nil {
		count = p.Count
	}
	var exchangedAt *int64
	if p.ExchangedAt != nil {
		exchangedAt = p.ExchangedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"awaitId":     awaitId,
		"userId":      userId,
		"rateName":    rateName,
		"name":        name,
		"count":       count,
		"exchangedAt": exchangedAt,
		"revision":    revision,
	}
}

func (p Await) Pointer() *Await {
	return &p
}

func CastAwaits(data []interface{}) []Await {
	v := make([]Await, 0)
	for _, d := range data {
		v = append(v, NewAwaitFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAwaitsFromDict(data []Await) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcquireAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func NewAcquireActionFromJson(data string) AcquireAction {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionFromDict(dict)
}

func NewAcquireActionFromDict(data map[string]interface{}) AcquireAction {
	return AcquireAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p AcquireAction) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
	}
}

func (p AcquireAction) Pointer() *AcquireAction {
	return &p
}

func CastAcquireActions(data []interface{}) []AcquireAction {
	v := make([]AcquireAction, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionsFromDict(data []AcquireAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ConsumeAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func NewConsumeActionFromJson(data string) ConsumeAction {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeActionFromDict(dict)
}

func NewConsumeActionFromDict(data map[string]interface{}) ConsumeAction {
	return ConsumeAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p ConsumeAction) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
	}
}

func (p ConsumeAction) Pointer() *ConsumeAction {
	return &p
}

func CastConsumeActions(data []interface{}) []ConsumeAction {
	v := make([]ConsumeAction, 0)
	for _, d := range data {
		v = append(v, NewConsumeActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeActionsFromDict(data []ConsumeAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Config struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

func NewConfigFromJson(data string) Config {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConfigFromDict(dict)
}

func NewConfigFromDict(data map[string]interface{}) Config {
	return Config{
		Key:   core.CastString(data["key"]),
		Value: core.CastString(data["value"]),
	}
}

func (p Config) ToDict() map[string]interface{} {

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

func (p Config) Pointer() *Config {
	return &p
}

func CastConfigs(data []interface{}) []Config {
	v := make([]Config, 0)
	for _, d := range data {
		v = append(v, NewConfigFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConfigsFromDict(data []Config) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GitHubCheckoutSetting struct {
	ApiKeyId       *string `json:"apiKeyId"`
	RepositoryName *string `json:"repositoryName"`
	SourcePath     *string `json:"sourcePath"`
	ReferenceType  *string `json:"referenceType"`
	CommitHash     *string `json:"commitHash"`
	BranchName     *string `json:"branchName"`
	TagName        *string `json:"tagName"`
}

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGitHubCheckoutSettingFromDict(dict)
}

func NewGitHubCheckoutSettingFromDict(data map[string]interface{}) GitHubCheckoutSetting {
	return GitHubCheckoutSetting{
		ApiKeyId:       core.CastString(data["apiKeyId"]),
		RepositoryName: core.CastString(data["repositoryName"]),
		SourcePath:     core.CastString(data["sourcePath"]),
		ReferenceType:  core.CastString(data["referenceType"]),
		CommitHash:     core.CastString(data["commitHash"]),
		BranchName:     core.CastString(data["branchName"]),
		TagName:        core.CastString(data["tagName"]),
	}
}

func (p GitHubCheckoutSetting) ToDict() map[string]interface{} {

	var apiKeyId *string
	if p.ApiKeyId != nil {
		apiKeyId = p.ApiKeyId
	}
	var repositoryName *string
	if p.RepositoryName != nil {
		repositoryName = p.RepositoryName
	}
	var sourcePath *string
	if p.SourcePath != nil {
		sourcePath = p.SourcePath
	}
	var referenceType *string
	if p.ReferenceType != nil {
		referenceType = p.ReferenceType
	}
	var commitHash *string
	if p.CommitHash != nil {
		commitHash = p.CommitHash
	}
	var branchName *string
	if p.BranchName != nil {
		branchName = p.BranchName
	}
	var tagName *string
	if p.TagName != nil {
		tagName = p.TagName
	}
	return map[string]interface{}{
		"apiKeyId":       apiKeyId,
		"repositoryName": repositoryName,
		"sourcePath":     sourcePath,
		"referenceType":  referenceType,
		"commitHash":     commitHash,
		"branchName":     branchName,
		"tagName":        tagName,
	}
}

func (p GitHubCheckoutSetting) Pointer() *GitHubCheckoutSetting {
	return &p
}

func CastGitHubCheckoutSettings(data []interface{}) []GitHubCheckoutSetting {
	v := make([]GitHubCheckoutSetting, 0)
	for _, d := range data {
		v = append(v, NewGitHubCheckoutSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGitHubCheckoutSettingsFromDict(data []GitHubCheckoutSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ScriptSetting struct {
	TriggerScriptId             *string `json:"triggerScriptId"`
	DoneTriggerTargetType       *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId         *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func NewScriptSettingFromJson(data string) ScriptSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewScriptSettingFromDict(dict)
}

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
	return ScriptSetting{
		TriggerScriptId:             core.CastString(data["triggerScriptId"]),
		DoneTriggerTargetType:       core.CastString(data["doneTriggerTargetType"]),
		DoneTriggerScriptId:         core.CastString(data["doneTriggerScriptId"]),
		DoneTriggerQueueNamespaceId: core.CastString(data["doneTriggerQueueNamespaceId"]),
	}
}

func (p ScriptSetting) ToDict() map[string]interface{} {

	var triggerScriptId *string
	if p.TriggerScriptId != nil {
		triggerScriptId = p.TriggerScriptId
	}
	var doneTriggerTargetType *string
	if p.DoneTriggerTargetType != nil {
		doneTriggerTargetType = p.DoneTriggerTargetType
	}
	var doneTriggerScriptId *string
	if p.DoneTriggerScriptId != nil {
		doneTriggerScriptId = p.DoneTriggerScriptId
	}
	var doneTriggerQueueNamespaceId *string
	if p.DoneTriggerQueueNamespaceId != nil {
		doneTriggerQueueNamespaceId = p.DoneTriggerQueueNamespaceId
	}
	return map[string]interface{}{
		"triggerScriptId":             triggerScriptId,
		"doneTriggerTargetType":       doneTriggerTargetType,
		"doneTriggerScriptId":         doneTriggerScriptId,
		"doneTriggerQueueNamespaceId": doneTriggerQueueNamespaceId,
	}
}

func (p ScriptSetting) Pointer() *ScriptSetting {
	return &p
}

func CastScriptSettings(data []interface{}) []ScriptSetting {
	v := make([]ScriptSetting, 0)
	for _, d := range data {
		v = append(v, NewScriptSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScriptSettingsFromDict(data []ScriptSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LogSetting struct {
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func NewLogSettingFromJson(data string) LogSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLogSettingFromDict(dict)
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
	return LogSetting{
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {

	var loggingNamespaceId *string
	if p.LoggingNamespaceId != nil {
		loggingNamespaceId = p.LoggingNamespaceId
	}
	return map[string]interface{}{
		"loggingNamespaceId": loggingNamespaceId,
	}
}

func (p LogSetting) Pointer() *LogSetting {
	return &p
}

func CastLogSettings(data []interface{}) []LogSetting {
	v := make([]LogSetting, 0)
	for _, d := range data {
		v = append(v, NewLogSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLogSettingsFromDict(data []LogSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionSetting struct {
	EnableAutoRun          *bool   `json:"enableAutoRun"`
	DistributorNamespaceId *string `json:"distributorNamespaceId"`
	KeyId                  *string `json:"keyId"`
	QueueNamespaceId       *string `json:"queueNamespaceId"`
}

func NewTransactionSettingFromJson(data string) TransactionSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTransactionSettingFromDict(dict)
}

func NewTransactionSettingFromDict(data map[string]interface{}) TransactionSetting {
	return TransactionSetting{
		EnableAutoRun:          core.CastBool(data["enableAutoRun"]),
		DistributorNamespaceId: core.CastString(data["distributorNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {

	var enableAutoRun *bool
	if p.EnableAutoRun != nil {
		enableAutoRun = p.EnableAutoRun
	}
	var distributorNamespaceId *string
	if p.DistributorNamespaceId != nil {
		distributorNamespaceId = p.DistributorNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	return map[string]interface{}{
		"enableAutoRun":          enableAutoRun,
		"distributorNamespaceId": distributorNamespaceId,
		"keyId":                  keyId,
		"queueNamespaceId":       queueNamespaceId,
	}
}

func (p TransactionSetting) Pointer() *TransactionSetting {
	return &p
}

func CastTransactionSettings(data []interface{}) []TransactionSetting {
	v := make([]TransactionSetting, 0)
	for _, d := range data {
		v = append(v, NewTransactionSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionSettingsFromDict(data []TransactionSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
