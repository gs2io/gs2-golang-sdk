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
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId                             *string        `json:"namespaceId"`
	Name                                    *string        `json:"name"`
	Description                             *string        `json:"description"`
	ChangePasswordIfTakeOver                *bool          `json:"changePasswordIfTakeOver"`
	DifferentUserIdForLoginAndDataRetention *bool          `json:"differentUserIdForLoginAndDataRetention"`
	CreateAccountScript                     *ScriptSetting `json:"createAccountScript"`
	AuthenticationScript                    *ScriptSetting `json:"authenticationScript"`
	CreateTakeOverScript                    *ScriptSetting `json:"createTakeOverScript"`
	DoTakeOverScript                        *ScriptSetting `json:"doTakeOverScript"`
	LogSetting                              *LogSetting    `json:"logSetting"`
	CreatedAt                               *int64         `json:"createdAt"`
	UpdatedAt                               *int64         `json:"updatedAt"`
	Revision                                *int64         `json:"revision"`
}

func (p *Namespace) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Namespace{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Namespace{}
	} else {
		*p = Namespace{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceId)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["changePasswordIfTakeOver"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangePasswordIfTakeOver)
		}
		if v, ok := d["differentUserIdForLoginAndDataRetention"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DifferentUserIdForLoginAndDataRetention)
		}
		if v, ok := d["createAccountScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreateAccountScript)
		}
		if v, ok := d["authenticationScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AuthenticationScript)
		}
		if v, ok := d["createTakeOverScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreateTakeOverScript)
		}
		if v, ok := d["doTakeOverScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DoTakeOverScript)
		}
		if v, ok := d["logSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LogSetting)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewNamespaceFromJson(data string) Namespace {
	req := Namespace{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:                             core.CastString(data["namespaceId"]),
		Name:                                    core.CastString(data["name"]),
		Description:                             core.CastString(data["description"]),
		ChangePasswordIfTakeOver:                core.CastBool(data["changePasswordIfTakeOver"]),
		DifferentUserIdForLoginAndDataRetention: core.CastBool(data["differentUserIdForLoginAndDataRetention"]),
		CreateAccountScript:                     NewScriptSettingFromDict(core.CastMap(data["createAccountScript"])).Pointer(),
		AuthenticationScript:                    NewScriptSettingFromDict(core.CastMap(data["authenticationScript"])).Pointer(),
		CreateTakeOverScript:                    NewScriptSettingFromDict(core.CastMap(data["createTakeOverScript"])).Pointer(),
		DoTakeOverScript:                        NewScriptSettingFromDict(core.CastMap(data["doTakeOverScript"])).Pointer(),
		LogSetting:                              NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                               core.CastInt64(data["createdAt"]),
		UpdatedAt:                               core.CastInt64(data["updatedAt"]),
		Revision:                                core.CastInt64(data["revision"]),
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
	var differentUserIdForLoginAndDataRetention *bool
	if p.DifferentUserIdForLoginAndDataRetention != nil {
		differentUserIdForLoginAndDataRetention = p.DifferentUserIdForLoginAndDataRetention
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId":              namespaceId,
		"name":                     name,
		"description":              description,
		"changePasswordIfTakeOver": changePasswordIfTakeOver,
		"differentUserIdForLoginAndDataRetention": differentUserIdForLoginAndDataRetention,
		"createAccountScript":                     createAccountScript,
		"authenticationScript":                    authenticationScript,
		"createTakeOverScript":                    createTakeOverScript,
		"doTakeOverScript":                        doTakeOverScript,
		"logSetting":                              logSetting,
		"createdAt":                               createdAt,
		"updatedAt":                               updatedAt,
		"revision":                                revision,
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
	AccountId           *string     `json:"accountId"`
	UserId              *string     `json:"userId"`
	Password            *string     `json:"password"`
	TimeOffset          *int32      `json:"timeOffset"`
	BanStatuses         []BanStatus `json:"banStatuses"`
	Banned              *bool       `json:"banned"`
	LastAuthenticatedAt *int64      `json:"lastAuthenticatedAt"`
	CreatedAt           *int64      `json:"createdAt"`
	Revision            *int64      `json:"revision"`
}

func (p *Account) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Account{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Account{}
	} else {
		*p = Account{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountId)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["password"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Password = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Password = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Password = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Password)
				}
			}
		}
		if v, ok := d["timeOffset"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TimeOffset)
		}
		if v, ok := d["banStatuses"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BanStatuses)
		}
		if v, ok := d["banned"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Banned)
		}
		if v, ok := d["lastAuthenticatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LastAuthenticatedAt)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewAccountFromJson(data string) Account {
	req := Account{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAccountFromDict(data map[string]interface{}) Account {
	return Account{
		AccountId:           core.CastString(data["accountId"]),
		UserId:              core.CastString(data["userId"]),
		Password:            core.CastString(data["password"]),
		TimeOffset:          core.CastInt32(data["timeOffset"]),
		BanStatuses:         CastBanStatuses(core.CastArray(data["banStatuses"])),
		Banned:              core.CastBool(data["banned"]),
		LastAuthenticatedAt: core.CastInt64(data["lastAuthenticatedAt"]),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		Revision:            core.CastInt64(data["revision"]),
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
	var banStatuses []interface{}
	if p.BanStatuses != nil {
		banStatuses = CastBanStatusesFromDict(
			p.BanStatuses,
		)
	}
	var banned *bool
	if p.Banned != nil {
		banned = p.Banned
	}
	var lastAuthenticatedAt *int64
	if p.LastAuthenticatedAt != nil {
		lastAuthenticatedAt = p.LastAuthenticatedAt
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"accountId":           accountId,
		"userId":              userId,
		"password":            password,
		"timeOffset":          timeOffset,
		"banStatuses":         banStatuses,
		"banned":              banned,
		"lastAuthenticatedAt": lastAuthenticatedAt,
		"createdAt":           createdAt,
		"revision":            revision,
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
	Revision       *int64  `json:"revision"`
}

func (p *TakeOver) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TakeOver{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = TakeOver{}
	} else {
		*p = TakeOver{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["takeOverId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TakeOverId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TakeOverId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TakeOverId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TakeOverId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TakeOverId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TakeOverId)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Type)
		}
		if v, ok := d["userIdentifier"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserIdentifier = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserIdentifier = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserIdentifier = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserIdentifier = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserIdentifier = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserIdentifier)
				}
			}
		}
		if v, ok := d["password"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Password = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Password = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Password = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Password)
				}
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewTakeOverFromJson(data string) TakeOver {
	req := TakeOver{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTakeOverFromDict(data map[string]interface{}) TakeOver {
	return TakeOver{
		TakeOverId:     core.CastString(data["takeOverId"]),
		UserId:         core.CastString(data["userId"]),
		Type:           core.CastInt32(data["type"]),
		UserIdentifier: core.CastString(data["userIdentifier"]),
		Password:       core.CastString(data["password"]),
		CreatedAt:      core.CastInt64(data["createdAt"]),
		Revision:       core.CastInt64(data["revision"]),
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"takeOverId":     takeOverId,
		"userId":         userId,
		"type":           _type,
		"userIdentifier": userIdentifier,
		"password":       password,
		"createdAt":      createdAt,
		"revision":       revision,
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

type PlatformId struct {
	PlatformId     *string `json:"platformId"`
	UserId         *string `json:"userId"`
	Type           *int32  `json:"type"`
	UserIdentifier *string `json:"userIdentifier"`
	CreatedAt      *int64  `json:"createdAt"`
	Revision       *int64  `json:"revision"`
}

func (p *PlatformId) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PlatformId{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PlatformId{}
	} else {
		*p = PlatformId{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["platformId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PlatformId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PlatformId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PlatformId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PlatformId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PlatformId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PlatformId)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Type)
		}
		if v, ok := d["userIdentifier"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserIdentifier = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserIdentifier = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserIdentifier = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserIdentifier = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserIdentifier = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserIdentifier)
				}
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewPlatformIdFromJson(data string) PlatformId {
	req := PlatformId{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPlatformIdFromDict(data map[string]interface{}) PlatformId {
	return PlatformId{
		PlatformId:     core.CastString(data["platformId"]),
		UserId:         core.CastString(data["userId"]),
		Type:           core.CastInt32(data["type"]),
		UserIdentifier: core.CastString(data["userIdentifier"]),
		CreatedAt:      core.CastInt64(data["createdAt"]),
		Revision:       core.CastInt64(data["revision"]),
	}
}

func (p PlatformId) ToDict() map[string]interface{} {

	var platformId *string
	if p.PlatformId != nil {
		platformId = p.PlatformId
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
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"platformId":     platformId,
		"userId":         userId,
		"type":           _type,
		"userIdentifier": userIdentifier,
		"createdAt":      createdAt,
		"revision":       revision,
	}
}

func (p PlatformId) Pointer() *PlatformId {
	return &p
}

func CastPlatformIds(data []interface{}) []PlatformId {
	v := make([]PlatformId, 0)
	for _, d := range data {
		v = append(v, NewPlatformIdFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPlatformIdsFromDict(data []PlatformId) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DataOwner struct {
	DataOwnerId *string `json:"dataOwnerId"`
	UserId      *string `json:"userId"`
	Name        *string `json:"name"`
	CreatedAt   *int64  `json:"createdAt"`
	Revision    *int64  `json:"revision"`
}

func (p *DataOwner) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DataOwner{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DataOwner{}
	} else {
		*p = DataOwner{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["dataOwnerId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DataOwnerId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DataOwnerId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DataOwnerId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DataOwnerId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DataOwnerId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DataOwnerId)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewDataOwnerFromJson(data string) DataOwner {
	req := DataOwner{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDataOwnerFromDict(data map[string]interface{}) DataOwner {
	return DataOwner{
		DataOwnerId: core.CastString(data["dataOwnerId"]),
		UserId:      core.CastString(data["userId"]),
		Name:        core.CastString(data["name"]),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		Revision:    core.CastInt64(data["revision"]),
	}
}

func (p DataOwner) ToDict() map[string]interface{} {

	var dataOwnerId *string
	if p.DataOwnerId != nil {
		dataOwnerId = p.DataOwnerId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"dataOwnerId": dataOwnerId,
		"userId":      userId,
		"name":        name,
		"createdAt":   createdAt,
		"revision":    revision,
	}
}

func (p DataOwner) Pointer() *DataOwner {
	return &p
}

func CastDataOwners(data []interface{}) []DataOwner {
	v := make([]DataOwner, 0)
	for _, d := range data {
		v = append(v, NewDataOwnerFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDataOwnersFromDict(data []DataOwner) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type PlatformUser struct {
	Type           *int32  `json:"type"`
	UserIdentifier *string `json:"userIdentifier"`
	UserId         *string `json:"userId"`
}

func (p *PlatformUser) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PlatformUser{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PlatformUser{}
	} else {
		*p = PlatformUser{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["type"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Type)
		}
		if v, ok := d["userIdentifier"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserIdentifier = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserIdentifier = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserIdentifier = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserIdentifier = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserIdentifier = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserIdentifier)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
	}
	return nil
}

func NewPlatformUserFromJson(data string) PlatformUser {
	req := PlatformUser{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPlatformUserFromDict(data map[string]interface{}) PlatformUser {
	return PlatformUser{
		Type:           core.CastInt32(data["type"]),
		UserIdentifier: core.CastString(data["userIdentifier"]),
		UserId:         core.CastString(data["userId"]),
	}
}

func (p PlatformUser) ToDict() map[string]interface{} {

	var _type *int32
	if p.Type != nil {
		_type = p.Type
	}
	var userIdentifier *string
	if p.UserIdentifier != nil {
		userIdentifier = p.UserIdentifier
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	return map[string]interface{}{
		"type":           _type,
		"userIdentifier": userIdentifier,
		"userId":         userId,
	}
}

func (p PlatformUser) Pointer() *PlatformUser {
	return &p
}

func CastPlatformUsers(data []interface{}) []PlatformUser {
	v := make([]PlatformUser, 0)
	for _, d := range data {
		v = append(v, NewPlatformUserFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPlatformUsersFromDict(data []PlatformUser) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BanStatus struct {
	Name             *string `json:"name"`
	Reason           *string `json:"reason"`
	ReleaseTimestamp *int64  `json:"releaseTimestamp"`
}

func (p *BanStatus) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BanStatus{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BanStatus{}
	} else {
		*p = BanStatus{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["reason"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Reason = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Reason = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Reason = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Reason = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Reason = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Reason)
				}
			}
		}
		if v, ok := d["releaseTimestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReleaseTimestamp)
		}
	}
	return nil
}

func NewBanStatusFromJson(data string) BanStatus {
	req := BanStatus{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBanStatusFromDict(data map[string]interface{}) BanStatus {
	return BanStatus{
		Name:             core.CastString(data["name"]),
		Reason:           core.CastString(data["reason"]),
		ReleaseTimestamp: core.CastInt64(data["releaseTimestamp"]),
	}
}

func (p BanStatus) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var reason *string
	if p.Reason != nil {
		reason = p.Reason
	}
	var releaseTimestamp *int64
	if p.ReleaseTimestamp != nil {
		releaseTimestamp = p.ReleaseTimestamp
	}
	return map[string]interface{}{
		"name":             name,
		"reason":           reason,
		"releaseTimestamp": releaseTimestamp,
	}
}

func (p BanStatus) Pointer() *BanStatus {
	return &p
}

func CastBanStatuses(data []interface{}) []BanStatus {
	v := make([]BanStatus, 0)
	for _, d := range data {
		v = append(v, NewBanStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBanStatusesFromDict(data []BanStatus) []interface{} {
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

func (p *ScriptSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ScriptSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ScriptSetting{}
	} else {
		*p = ScriptSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["triggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerScriptId)
				}
			}
		}
		if v, ok := d["doneTriggerTargetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerTargetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerTargetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerTargetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerTargetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerTargetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerTargetType)
				}
			}
		}
		if v, ok := d["doneTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerScriptId)
				}
			}
		}
		if v, ok := d["doneTriggerQueueNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerQueueNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerQueueNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerQueueNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerQueueNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerQueueNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerQueueNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewScriptSettingFromJson(data string) ScriptSetting {
	req := ScriptSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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

func (p *LogSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LogSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = LogSetting{}
	} else {
		*p = LogSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["loggingNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LoggingNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LoggingNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LoggingNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LoggingNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LoggingNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LoggingNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewLogSettingFromJson(data string) LogSetting {
	req := LogSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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
