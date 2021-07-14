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

import "core"

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
	return map[string]interface{}{
		"namespaceId":              p.NamespaceId,
		"name":                     p.Name,
		"description":              p.Description,
		"changePasswordIfTakeOver": p.ChangePasswordIfTakeOver,
		"createAccountScript":      p.CreateAccountScript.ToDict(),
		"authenticationScript":     p.AuthenticationScript.ToDict(),
		"createTakeOverScript":     p.CreateTakeOverScript.ToDict(),
		"doTakeOverScript":         p.DoTakeOverScript.ToDict(),
		"logSetting":               p.LogSetting.ToDict(),
		"createdAt":                p.CreatedAt,
		"updatedAt":                p.UpdatedAt,
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
	return map[string]interface{}{
		"accountId":  p.AccountId,
		"userId":     p.UserId,
		"password":   p.Password,
		"timeOffset": p.TimeOffset,
		"createdAt":  p.CreatedAt,
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
	return map[string]interface{}{
		"takeOverId":     p.TakeOverId,
		"userId":         p.UserId,
		"type":           p.Type,
		"userIdentifier": p.UserIdentifier,
		"password":       p.Password,
		"createdAt":      p.CreatedAt,
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

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
	return ScriptSetting{
		TriggerScriptId:             core.CastString(data["triggerScriptId"]),
		DoneTriggerTargetType:       core.CastString(data["doneTriggerTargetType"]),
		DoneTriggerScriptId:         core.CastString(data["doneTriggerScriptId"]),
		DoneTriggerQueueNamespaceId: core.CastString(data["doneTriggerQueueNamespaceId"]),
	}
}

func (p ScriptSetting) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"triggerScriptId":             p.TriggerScriptId,
		"doneTriggerTargetType":       p.DoneTriggerTargetType,
		"doneTriggerScriptId":         p.DoneTriggerScriptId,
		"doneTriggerQueueNamespaceId": p.DoneTriggerQueueNamespaceId,
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

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
	return LogSetting{
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"loggingNamespaceId": p.LoggingNamespaceId,
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
