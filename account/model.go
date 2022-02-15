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

package account

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId              *string        `json:"namespaceId"`
	Name                     *string        `json:"name"`
	Description              *string        `json:"description"`
	ChangePasswordIfTakeOver *bool          `json:"changePasswordIfTakeOver"`
	CreateAccountScript      *ScriptSetting `json:"createAccountScript"`
	AuthenticationScript     *ScriptSetting `json:"authenticationScript"`
	CreateTakeOverScript     *ScriptSetting `json:"createTakeOverScript"`
	DoTakeOverScript         *ScriptSetting `json:"doTakeOverScript"`
	LogSetting               *LogSetting    `json:"logSetting"`
	CreatedAt                *int64         `json:"createdAt"`
	UpdatedAt                *int64         `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:              core.CastString(data["namespaceId"]),
		Name:                     core.CastString(data["name"]),
		Description:              core.CastString(data["description"]),
		ChangePasswordIfTakeOver: core.CastBool(data["changePasswordIfTakeOver"]),
		CreateAccountScript:      NewScriptSettingFromDict(core.CastMap(data["createAccountScript"])).Pointer(),
		AuthenticationScript:     NewScriptSettingFromDict(core.CastMap(data["authenticationScript"])).Pointer(),
		CreateTakeOverScript:     NewScriptSettingFromDict(core.CastMap(data["createTakeOverScript"])).Pointer(),
		DoTakeOverScript:         NewScriptSettingFromDict(core.CastMap(data["doTakeOverScript"])).Pointer(),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                core.CastInt64(data["createdAt"]),
		UpdatedAt:                core.CastInt64(data["updatedAt"]),
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
	var changePasswordIfTakeOver *bool
	if p.ChangePasswordIfTakeOver != nil {
		changePasswordIfTakeOver = p.ChangePasswordIfTakeOver
	}
	var createAccountScript map[string]interface{}
	if p.CreateAccountScript != nil {
		createAccountScript = p.CreateAccountScript.ToDict()
	}
	var authenticationScript map[string]interface{}
	if p.AuthenticationScript != nil {
		authenticationScript = p.AuthenticationScript.ToDict()
	}
	var createTakeOverScript map[string]interface{}
	if p.CreateTakeOverScript != nil {
		createTakeOverScript = p.CreateTakeOverScript.ToDict()
	}
	var doTakeOverScript map[string]interface{}
	if p.DoTakeOverScript != nil {
		doTakeOverScript = p.DoTakeOverScript.ToDict()
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
	return map[string]interface{}{
		"namespaceId":              namespaceId,
		"name":                     name,
		"description":              description,
		"changePasswordIfTakeOver": changePasswordIfTakeOver,
		"createAccountScript":      createAccountScript,
		"authenticationScript":     authenticationScript,
		"createTakeOverScript":     createTakeOverScript,
		"doTakeOverScript":         doTakeOverScript,
		"logSetting":               logSetting,
		"createdAt":                createdAt,
		"updatedAt":                updatedAt,
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

type Account struct {
	AccountId  *string `json:"accountId"`
	UserId     *string `json:"userId"`
	Password   *string `json:"password"`
	TimeOffset *int32  `json:"timeOffset"`
	CreatedAt  *int64  `json:"createdAt"`
}

func NewAccountFromJson(data string) Account {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAccountFromDict(dict)
}

func NewAccountFromDict(data map[string]interface{}) Account {
	return Account{
		AccountId:  core.CastString(data["accountId"]),
		UserId:     core.CastString(data["userId"]),
		Password:   core.CastString(data["password"]),
		TimeOffset: core.CastInt32(data["timeOffset"]),
		CreatedAt:  core.CastInt64(data["createdAt"]),
	}
}

func (p Account) ToDict() map[string]interface{} {

	var accountId *string
	if p.AccountId != nil {
		accountId = p.AccountId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var password *string
	if p.Password != nil {
		password = p.Password
	}
	var timeOffset *int32
	if p.TimeOffset != nil {
		timeOffset = p.TimeOffset
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	return map[string]interface{}{
		"accountId":  accountId,
		"userId":     userId,
		"password":   password,
		"timeOffset": timeOffset,
		"createdAt":  createdAt,
	}
}

func (p Account) Pointer() *Account {
	return &p
}

func CastAccounts(data []interface{}) []Account {
	v := make([]Account, 0)
	for _, d := range data {
		v = append(v, NewAccountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccountsFromDict(data []Account) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TakeOver struct {
	TakeOverId     *string `json:"takeOverId"`
	UserId         *string `json:"userId"`
	Type           *int32  `json:"type"`
	UserIdentifier *string `json:"userIdentifier"`
	Password       *string `json:"password"`
	CreatedAt      *int64  `json:"createdAt"`
}

func NewTakeOverFromJson(data string) TakeOver {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTakeOverFromDict(dict)
}

func NewTakeOverFromDict(data map[string]interface{}) TakeOver {
	return TakeOver{
		TakeOverId:     core.CastString(data["takeOverId"]),
		UserId:         core.CastString(data["userId"]),
		Type:           core.CastInt32(data["type"]),
		UserIdentifier: core.CastString(data["userIdentifier"]),
		Password:       core.CastString(data["password"]),
		CreatedAt:      core.CastInt64(data["createdAt"]),
	}
}

func (p TakeOver) ToDict() map[string]interface{} {

	var takeOverId *string
	if p.TakeOverId != nil {
		takeOverId = p.TakeOverId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var _type *int32
	if p.Type != nil {
		_type = p.Type
	}
	var userIdentifier *string
	if p.UserIdentifier != nil {
		userIdentifier = p.UserIdentifier
	}
	var password *string
	if p.Password != nil {
		password = p.Password
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	return map[string]interface{}{
		"takeOverId":     takeOverId,
		"userId":         userId,
		"type":           _type,
		"userIdentifier": userIdentifier,
		"password":       password,
		"createdAt":      createdAt,
	}
}

func (p TakeOver) Pointer() *TakeOver {
	return &p
}

func CastTakeOvers(data []interface{}) []TakeOver {
	v := make([]TakeOver, 0)
	for _, d := range data {
		v = append(v, NewTakeOverFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTakeOversFromDict(data []TakeOver) []interface{} {
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
