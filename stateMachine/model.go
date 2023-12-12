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

package stateMachine

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId                 *string             `json:"namespaceId"`
	Name                        *string             `json:"name"`
	Description                 *string             `json:"description"`
	SupportSpeculativeExecution *string             `json:"supportSpeculativeExecution"`
	TransactionSetting          *TransactionSetting `json:"transactionSetting"`
	StartScript                 *ScriptSetting      `json:"startScript"`
	PassScript                  *ScriptSetting      `json:"passScript"`
	ErrorScript                 *ScriptSetting      `json:"errorScript"`
	LowestStateMachineVersion   *int64              `json:"lowestStateMachineVersion"`
	LogSetting                  *LogSetting         `json:"logSetting"`
	CreatedAt                   *int64              `json:"createdAt"`
	UpdatedAt                   *int64              `json:"updatedAt"`
	Revision                    *int64              `json:"revision"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:                 core.CastString(data["namespaceId"]),
		Name:                        core.CastString(data["name"]),
		Description:                 core.CastString(data["description"]),
		SupportSpeculativeExecution: core.CastString(data["supportSpeculativeExecution"]),
		TransactionSetting:          NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		StartScript:                 NewScriptSettingFromDict(core.CastMap(data["startScript"])).Pointer(),
		PassScript:                  NewScriptSettingFromDict(core.CastMap(data["passScript"])).Pointer(),
		ErrorScript:                 NewScriptSettingFromDict(core.CastMap(data["errorScript"])).Pointer(),
		LowestStateMachineVersion:   core.CastInt64(data["lowestStateMachineVersion"]),
		LogSetting:                  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                   core.CastInt64(data["createdAt"]),
		UpdatedAt:                   core.CastInt64(data["updatedAt"]),
		Revision:                    core.CastInt64(data["revision"]),
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
	var supportSpeculativeExecution *string
	if p.SupportSpeculativeExecution != nil {
		supportSpeculativeExecution = p.SupportSpeculativeExecution
	}
	var transactionSetting map[string]interface{}
	if p.TransactionSetting != nil {
		transactionSetting = p.TransactionSetting.ToDict()
	}
	var startScript map[string]interface{}
	if p.StartScript != nil {
		startScript = p.StartScript.ToDict()
	}
	var passScript map[string]interface{}
	if p.PassScript != nil {
		passScript = p.PassScript.ToDict()
	}
	var errorScript map[string]interface{}
	if p.ErrorScript != nil {
		errorScript = p.ErrorScript.ToDict()
	}
	var lowestStateMachineVersion *int64
	if p.LowestStateMachineVersion != nil {
		lowestStateMachineVersion = p.LowestStateMachineVersion
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId":                 namespaceId,
		"name":                        name,
		"description":                 description,
		"supportSpeculativeExecution": supportSpeculativeExecution,
		"transactionSetting":          transactionSetting,
		"startScript":                 startScript,
		"passScript":                  passScript,
		"errorScript":                 errorScript,
		"lowestStateMachineVersion":   lowestStateMachineVersion,
		"logSetting":                  logSetting,
		"createdAt":                   createdAt,
		"updatedAt":                   updatedAt,
		"revision":                    revision,
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

type StateMachineMaster struct {
	StateMachineId       *string `json:"stateMachineId"`
	MainStateMachineName *string `json:"mainStateMachineName"`
	Payload              *string `json:"payload"`
	Version              *int64  `json:"version"`
	CreatedAt            *int64  `json:"createdAt"`
	UpdatedAt            *int64  `json:"updatedAt"`
	Revision             *int64  `json:"revision"`
}

func NewStateMachineMasterFromJson(data string) StateMachineMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStateMachineMasterFromDict(dict)
}

func NewStateMachineMasterFromDict(data map[string]interface{}) StateMachineMaster {
	return StateMachineMaster{
		StateMachineId:       core.CastString(data["stateMachineId"]),
		MainStateMachineName: core.CastString(data["mainStateMachineName"]),
		Payload:              core.CastString(data["payload"]),
		Version:              core.CastInt64(data["version"]),
		CreatedAt:            core.CastInt64(data["createdAt"]),
		UpdatedAt:            core.CastInt64(data["updatedAt"]),
		Revision:             core.CastInt64(data["revision"]),
	}
}

func (p StateMachineMaster) ToDict() map[string]interface{} {

	var stateMachineId *string
	if p.StateMachineId != nil {
		stateMachineId = p.StateMachineId
	}
	var mainStateMachineName *string
	if p.MainStateMachineName != nil {
		mainStateMachineName = p.MainStateMachineName
	}
	var payload *string
	if p.Payload != nil {
		payload = p.Payload
	}
	var version *int64
	if p.Version != nil {
		version = p.Version
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
		"stateMachineId":       stateMachineId,
		"mainStateMachineName": mainStateMachineName,
		"payload":              payload,
		"version":              version,
		"createdAt":            createdAt,
		"updatedAt":            updatedAt,
		"revision":             revision,
	}
}

func (p StateMachineMaster) Pointer() *StateMachineMaster {
	return &p
}

func CastStateMachineMasters(data []interface{}) []StateMachineMaster {
	v := make([]StateMachineMaster, 0)
	for _, d := range data {
		v = append(v, NewStateMachineMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStateMachineMastersFromDict(data []StateMachineMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Status struct {
	StatusId                   *string       `json:"statusId"`
	UserId                     *string       `json:"userId"`
	Name                       *string       `json:"name"`
	StateMachineVersion        *int64        `json:"stateMachineVersion"`
	EnableSpeculativeExecution *string       `json:"enableSpeculativeExecution"`
	StateMachineDefinition     *string       `json:"stateMachineDefinition"`
	RandomStatus               *RandomStatus `json:"randomStatus"`
	Stacks                     []StackEntry  `json:"stacks"`
	Variables                  []Variable    `json:"variables"`
	Status                     *string       `json:"status"`
	LastError                  *string       `json:"lastError"`
	TransitionCount            *int32        `json:"transitionCount"`
	CreatedAt                  *int64        `json:"createdAt"`
	UpdatedAt                  *int64        `json:"updatedAt"`
}

func NewStatusFromJson(data string) Status {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStatusFromDict(dict)
}

func NewStatusFromDict(data map[string]interface{}) Status {
	return Status{
		StatusId:                   core.CastString(data["statusId"]),
		UserId:                     core.CastString(data["userId"]),
		Name:                       core.CastString(data["name"]),
		StateMachineVersion:        core.CastInt64(data["stateMachineVersion"]),
		EnableSpeculativeExecution: core.CastString(data["enableSpeculativeExecution"]),
		StateMachineDefinition:     core.CastString(data["stateMachineDefinition"]),
		RandomStatus:               NewRandomStatusFromDict(core.CastMap(data["randomStatus"])).Pointer(),
		Stacks:                     CastStackEntries(core.CastArray(data["stacks"])),
		Variables:                  CastVariables(core.CastArray(data["variables"])),
		Status:                     core.CastString(data["status"]),
		LastError:                  core.CastString(data["lastError"]),
		TransitionCount:            core.CastInt32(data["transitionCount"]),
		CreatedAt:                  core.CastInt64(data["createdAt"]),
		UpdatedAt:                  core.CastInt64(data["updatedAt"]),
	}
}

func (p Status) ToDict() map[string]interface{} {

	var statusId *string
	if p.StatusId != nil {
		statusId = p.StatusId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var stateMachineVersion *int64
	if p.StateMachineVersion != nil {
		stateMachineVersion = p.StateMachineVersion
	}
	var enableSpeculativeExecution *string
	if p.EnableSpeculativeExecution != nil {
		enableSpeculativeExecution = p.EnableSpeculativeExecution
	}
	var stateMachineDefinition *string
	if p.StateMachineDefinition != nil {
		stateMachineDefinition = p.StateMachineDefinition
	}
	var randomStatus map[string]interface{}
	if p.RandomStatus != nil {
		randomStatus = p.RandomStatus.ToDict()
	}
	var stacks []interface{}
	if p.Stacks != nil {
		stacks = CastStackEntriesFromDict(
			p.Stacks,
		)
	}
	var variables []interface{}
	if p.Variables != nil {
		variables = CastVariablesFromDict(
			p.Variables,
		)
	}
	var status *string
	if p.Status != nil {
		status = p.Status
	}
	var lastError *string
	if p.LastError != nil {
		lastError = p.LastError
	}
	var transitionCount *int32
	if p.TransitionCount != nil {
		transitionCount = p.TransitionCount
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	return map[string]interface{}{
		"statusId":                   statusId,
		"userId":                     userId,
		"name":                       name,
		"stateMachineVersion":        stateMachineVersion,
		"enableSpeculativeExecution": enableSpeculativeExecution,
		"stateMachineDefinition":     stateMachineDefinition,
		"randomStatus":               randomStatus,
		"stacks":                     stacks,
		"variables":                  variables,
		"status":                     status,
		"lastError":                  lastError,
		"transitionCount":            transitionCount,
		"createdAt":                  createdAt,
		"updatedAt":                  updatedAt,
	}
}

func (p Status) Pointer() *Status {
	return &p
}

func CastStatuses(data []interface{}) []Status {
	v := make([]Status, 0)
	for _, d := range data {
		v = append(v, NewStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStatusesFromDict(data []Status) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StackEntry struct {
	StateMachineName *string `json:"stateMachineName"`
	TaskName         *string `json:"taskName"`
}

func NewStackEntryFromJson(data string) StackEntry {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStackEntryFromDict(dict)
}

func NewStackEntryFromDict(data map[string]interface{}) StackEntry {
	return StackEntry{
		StateMachineName: core.CastString(data["stateMachineName"]),
		TaskName:         core.CastString(data["taskName"]),
	}
}

func (p StackEntry) ToDict() map[string]interface{} {

	var stateMachineName *string
	if p.StateMachineName != nil {
		stateMachineName = p.StateMachineName
	}
	var taskName *string
	if p.TaskName != nil {
		taskName = p.TaskName
	}
	return map[string]interface{}{
		"stateMachineName": stateMachineName,
		"taskName":         taskName,
	}
}

func (p StackEntry) Pointer() *StackEntry {
	return &p
}

func CastStackEntries(data []interface{}) []StackEntry {
	v := make([]StackEntry, 0)
	for _, d := range data {
		v = append(v, NewStackEntryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStackEntriesFromDict(data []StackEntry) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Variable struct {
	StateMachineName *string `json:"stateMachineName"`
	Value            *string `json:"value"`
}

func NewVariableFromJson(data string) Variable {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVariableFromDict(dict)
}

func NewVariableFromDict(data map[string]interface{}) Variable {
	return Variable{
		StateMachineName: core.CastString(data["stateMachineName"]),
		Value:            core.CastString(data["value"]),
	}
}

func (p Variable) ToDict() map[string]interface{} {

	var stateMachineName *string
	if p.StateMachineName != nil {
		stateMachineName = p.StateMachineName
	}
	var value *string
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"stateMachineName": stateMachineName,
		"value":            value,
	}
}

func (p Variable) Pointer() *Variable {
	return &p
}

func CastVariables(data []interface{}) []Variable {
	v := make([]Variable, 0)
	for _, d := range data {
		v = append(v, NewVariableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVariablesFromDict(data []Variable) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionSetting struct {
	EnableAutoRun          *bool   `json:"enableAutoRun"`
	DistributorNamespaceId *string `json:"distributorNamespaceId"`
	// Deprecated: should not be used
	KeyId            *string `json:"keyId"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
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

type Event struct {
	EventType        *string           `json:"eventType"`
	ChangeStateEvent *ChangeStateEvent `json:"changeStateEvent"`
	EmitEvent        *EmitEvent        `json:"emitEvent"`
}

func NewEventFromJson(data string) Event {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEventFromDict(dict)
}

func NewEventFromDict(data map[string]interface{}) Event {
	return Event{
		EventType:        core.CastString(data["eventType"]),
		ChangeStateEvent: NewChangeStateEventFromDict(core.CastMap(data["changeStateEvent"])).Pointer(),
		EmitEvent:        NewEmitEventFromDict(core.CastMap(data["emitEvent"])).Pointer(),
	}
}

func (p Event) ToDict() map[string]interface{} {

	var eventType *string
	if p.EventType != nil {
		eventType = p.EventType
	}
	var changeStateEvent map[string]interface{}
	if p.ChangeStateEvent != nil {
		changeStateEvent = p.ChangeStateEvent.ToDict()
	}
	var emitEvent map[string]interface{}
	if p.EmitEvent != nil {
		emitEvent = p.EmitEvent.ToDict()
	}
	return map[string]interface{}{
		"eventType":        eventType,
		"changeStateEvent": changeStateEvent,
		"emitEvent":        emitEvent,
	}
}

func (p Event) Pointer() *Event {
	return &p
}

func CastEvents(data []interface{}) []Event {
	v := make([]Event, 0)
	for _, d := range data {
		v = append(v, NewEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEventsFromDict(data []Event) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChangeStateEvent struct {
	TaskName  *string `json:"taskName"`
	Hash      *string `json:"hash"`
	Timestamp *int64  `json:"timestamp"`
}

func NewChangeStateEventFromJson(data string) ChangeStateEvent {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChangeStateEventFromDict(dict)
}

func NewChangeStateEventFromDict(data map[string]interface{}) ChangeStateEvent {
	return ChangeStateEvent{
		TaskName:  core.CastString(data["taskName"]),
		Hash:      core.CastString(data["hash"]),
		Timestamp: core.CastInt64(data["timestamp"]),
	}
}

func (p ChangeStateEvent) ToDict() map[string]interface{} {

	var taskName *string
	if p.TaskName != nil {
		taskName = p.TaskName
	}
	var hash *string
	if p.Hash != nil {
		hash = p.Hash
	}
	var timestamp *int64
	if p.Timestamp != nil {
		timestamp = p.Timestamp
	}
	return map[string]interface{}{
		"taskName":  taskName,
		"hash":      hash,
		"timestamp": timestamp,
	}
}

func (p ChangeStateEvent) Pointer() *ChangeStateEvent {
	return &p
}

func CastChangeStateEvents(data []interface{}) []ChangeStateEvent {
	v := make([]ChangeStateEvent, 0)
	for _, d := range data {
		v = append(v, NewChangeStateEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChangeStateEventsFromDict(data []ChangeStateEvent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type EmitEvent struct {
	Event      *string `json:"event"`
	Parameters *string `json:"parameters"`
	Timestamp  *int64  `json:"timestamp"`
}

func NewEmitEventFromJson(data string) EmitEvent {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEmitEventFromDict(dict)
}

func NewEmitEventFromDict(data map[string]interface{}) EmitEvent {
	return EmitEvent{
		Event:      core.CastString(data["event"]),
		Parameters: core.CastString(data["parameters"]),
		Timestamp:  core.CastInt64(data["timestamp"]),
	}
}

func (p EmitEvent) ToDict() map[string]interface{} {

	var event *string
	if p.Event != nil {
		event = p.Event
	}
	var parameters *string
	if p.Parameters != nil {
		parameters = p.Parameters
	}
	var timestamp *int64
	if p.Timestamp != nil {
		timestamp = p.Timestamp
	}
	return map[string]interface{}{
		"event":      event,
		"parameters": parameters,
		"timestamp":  timestamp,
	}
}

func (p EmitEvent) Pointer() *EmitEvent {
	return &p
}

func CastEmitEvents(data []interface{}) []EmitEvent {
	v := make([]EmitEvent, 0)
	for _, d := range data {
		v = append(v, NewEmitEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEmitEventsFromDict(data []EmitEvent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomStatus struct {
	Seed *int64       `json:"seed"`
	Used []RandomUsed `json:"used"`
}

func NewRandomStatusFromJson(data string) RandomStatus {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomStatusFromDict(dict)
}

func NewRandomStatusFromDict(data map[string]interface{}) RandomStatus {
	return RandomStatus{
		Seed: core.CastInt64(data["seed"]),
		Used: CastRandomUseds(core.CastArray(data["used"])),
	}
}

func (p RandomStatus) ToDict() map[string]interface{} {

	var seed *int64
	if p.Seed != nil {
		seed = p.Seed
	}
	var used []interface{}
	if p.Used != nil {
		used = CastRandomUsedsFromDict(
			p.Used,
		)
	}
	return map[string]interface{}{
		"seed": seed,
		"used": used,
	}
}

func (p RandomStatus) Pointer() *RandomStatus {
	return &p
}

func CastRandomStatuses(data []interface{}) []RandomStatus {
	v := make([]RandomStatus, 0)
	for _, d := range data {
		v = append(v, NewRandomStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomStatusesFromDict(data []RandomStatus) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomUsed struct {
	Category *int64 `json:"category"`
	Used     *int64 `json:"used"`
}

func NewRandomUsedFromJson(data string) RandomUsed {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomUsedFromDict(dict)
}

func NewRandomUsedFromDict(data map[string]interface{}) RandomUsed {
	return RandomUsed{
		Category: core.CastInt64(data["category"]),
		Used:     core.CastInt64(data["used"]),
	}
}

func (p RandomUsed) ToDict() map[string]interface{} {

	var category *int64
	if p.Category != nil {
		category = p.Category
	}
	var used *int64
	if p.Used != nil {
		used = p.Used
	}
	return map[string]interface{}{
		"category": category,
		"used":     used,
	}
}

func (p RandomUsed) Pointer() *RandomUsed {
	return &p
}

func CastRandomUseds(data []interface{}) []RandomUsed {
	v := make([]RandomUsed, 0)
	for _, d := range data {
		v = append(v, NewRandomUsedFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomUsedsFromDict(data []RandomUsed) []interface{} {
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
