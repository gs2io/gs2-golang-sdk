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
		NamespaceId: func() *string {
			v, ok := data["namespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		ChangePasswordIfTakeOver: func() *bool {
			v, ok := data["changePasswordIfTakeOver"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["changePasswordIfTakeOver"])
		}(),
		DifferentUserIdForLoginAndDataRetention: func() *bool {
			v, ok := data["differentUserIdForLoginAndDataRetention"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["differentUserIdForLoginAndDataRetention"])
		}(),
		CreateAccountScript: func() *ScriptSetting {
			v, ok := data["createAccountScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["createAccountScript"])).Pointer()
		}(),
		AuthenticationScript: func() *ScriptSetting {
			v, ok := data["authenticationScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["authenticationScript"])).Pointer()
		}(),
		CreateTakeOverScript: func() *ScriptSetting {
			v, ok := data["createTakeOverScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["createTakeOverScript"])).Pointer()
		}(),
		DoTakeOverScript: func() *ScriptSetting {
			v, ok := data["doTakeOverScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["doTakeOverScript"])).Pointer()
		}(),
		LogSetting: func() *LogSetting {
			v, ok := data["logSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer()
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.ChangePasswordIfTakeOver != nil {
		m["changePasswordIfTakeOver"] = p.ChangePasswordIfTakeOver
	}
	if p.DifferentUserIdForLoginAndDataRetention != nil {
		m["differentUserIdForLoginAndDataRetention"] = p.DifferentUserIdForLoginAndDataRetention
	}
	if p.CreateAccountScript != nil {
		m["createAccountScript"] = func() map[string]interface{} {
			if p.CreateAccountScript == nil {
				return nil
			}
			return p.CreateAccountScript.ToDict()
		}()
	}
	if p.AuthenticationScript != nil {
		m["authenticationScript"] = func() map[string]interface{} {
			if p.AuthenticationScript == nil {
				return nil
			}
			return p.AuthenticationScript.ToDict()
		}()
	}
	if p.CreateTakeOverScript != nil {
		m["createTakeOverScript"] = func() map[string]interface{} {
			if p.CreateTakeOverScript == nil {
				return nil
			}
			return p.CreateTakeOverScript.ToDict()
		}()
	}
	if p.DoTakeOverScript != nil {
		m["doTakeOverScript"] = func() map[string]interface{} {
			if p.DoTakeOverScript == nil {
				return nil
			}
			return p.DoTakeOverScript.ToDict()
		}()
	}
	if p.LogSetting != nil {
		m["logSetting"] = func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}()
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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
		AccountId: func() *string {
			v, ok := data["accountId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Password: func() *string {
			v, ok := data["password"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["password"])
		}(),
		TimeOffset: func() *int32 {
			v, ok := data["timeOffset"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["timeOffset"])
		}(),
		BanStatuses: func() []BanStatus {
			if data["banStatuses"] == nil {
				return nil
			}
			return CastBanStatuses(core.CastArray(data["banStatuses"]))
		}(),
		Banned: func() *bool {
			v, ok := data["banned"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["banned"])
		}(),
		LastAuthenticatedAt: func() *int64 {
			v, ok := data["lastAuthenticatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["lastAuthenticatedAt"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p Account) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.AccountId != nil {
		m["accountId"] = p.AccountId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Password != nil {
		m["password"] = p.Password
	}
	if p.TimeOffset != nil {
		m["timeOffset"] = p.TimeOffset
	}
	if p.BanStatuses != nil {
		m["banStatuses"] = CastBanStatusesFromDict(
			p.BanStatuses,
		)
	}
	if p.Banned != nil {
		m["banned"] = p.Banned
	}
	if p.LastAuthenticatedAt != nil {
		m["lastAuthenticatedAt"] = p.LastAuthenticatedAt
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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
		TakeOverId: func() *string {
			v, ok := data["takeOverId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["takeOverId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Type: func() *int32 {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["type"])
		}(),
		UserIdentifier: func() *string {
			v, ok := data["userIdentifier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userIdentifier"])
		}(),
		Password: func() *string {
			v, ok := data["password"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["password"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p TakeOver) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TakeOverId != nil {
		m["takeOverId"] = p.TakeOverId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.UserIdentifier != nil {
		m["userIdentifier"] = p.UserIdentifier
	}
	if p.Password != nil {
		m["password"] = p.Password
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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
		PlatformId: func() *string {
			v, ok := data["platformId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["platformId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Type: func() *int32 {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["type"])
		}(),
		UserIdentifier: func() *string {
			v, ok := data["userIdentifier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userIdentifier"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p PlatformId) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PlatformId != nil {
		m["platformId"] = p.PlatformId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.UserIdentifier != nil {
		m["userIdentifier"] = p.UserIdentifier
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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
		DataOwnerId: func() *string {
			v, ok := data["dataOwnerId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["dataOwnerId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p DataOwner) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.DataOwnerId != nil {
		m["dataOwnerId"] = p.DataOwnerId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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

type TakeOverTypeModel struct {
	TakeOverTypeModelId  *string               `json:"takeOverTypeModelId"`
	Type                 *int32                `json:"type"`
	Metadata             *string               `json:"metadata"`
	OpenIdConnectSetting *OpenIdConnectSetting `json:"openIdConnectSetting"`
}

func (p *TakeOverTypeModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TakeOverTypeModel{}
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
		*p = TakeOverTypeModel{}
	} else {
		*p = TakeOverTypeModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["takeOverTypeModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TakeOverTypeModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TakeOverTypeModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TakeOverTypeModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TakeOverTypeModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TakeOverTypeModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TakeOverTypeModelId)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Type)
		}
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["openIdConnectSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.OpenIdConnectSetting)
		}
	}
	return nil
}

func NewTakeOverTypeModelFromJson(data string) TakeOverTypeModel {
	req := TakeOverTypeModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTakeOverTypeModelFromDict(data map[string]interface{}) TakeOverTypeModel {
	return TakeOverTypeModel{
		TakeOverTypeModelId: func() *string {
			v, ok := data["takeOverTypeModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["takeOverTypeModelId"])
		}(),
		Type: func() *int32 {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["type"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		OpenIdConnectSetting: func() *OpenIdConnectSetting {
			v, ok := data["openIdConnectSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewOpenIdConnectSettingFromDict(core.CastMap(data["openIdConnectSetting"])).Pointer()
		}(),
	}
}

func (p TakeOverTypeModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TakeOverTypeModelId != nil {
		m["takeOverTypeModelId"] = p.TakeOverTypeModelId
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.OpenIdConnectSetting != nil {
		m["openIdConnectSetting"] = func() map[string]interface{} {
			if p.OpenIdConnectSetting == nil {
				return nil
			}
			return p.OpenIdConnectSetting.ToDict()
		}()
	}
	return m
}

func (p TakeOverTypeModel) Pointer() *TakeOverTypeModel {
	return &p
}

func CastTakeOverTypeModels(data []interface{}) []TakeOverTypeModel {
	v := make([]TakeOverTypeModel, 0)
	for _, d := range data {
		v = append(v, NewTakeOverTypeModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTakeOverTypeModelsFromDict(data []TakeOverTypeModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TakeOverTypeModelMaster struct {
	TakeOverTypeModelId  *string               `json:"takeOverTypeModelId"`
	Type                 *int32                `json:"type"`
	Description          *string               `json:"description"`
	Metadata             *string               `json:"metadata"`
	OpenIdConnectSetting *OpenIdConnectSetting `json:"openIdConnectSetting"`
	CreatedAt            *int64                `json:"createdAt"`
	UpdatedAt            *int64                `json:"updatedAt"`
	Revision             *int64                `json:"revision"`
}

func (p *TakeOverTypeModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TakeOverTypeModelMaster{}
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
		*p = TakeOverTypeModelMaster{}
	} else {
		*p = TakeOverTypeModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["takeOverTypeModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TakeOverTypeModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TakeOverTypeModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TakeOverTypeModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TakeOverTypeModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TakeOverTypeModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TakeOverTypeModelId)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Type)
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["openIdConnectSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.OpenIdConnectSetting)
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

func NewTakeOverTypeModelMasterFromJson(data string) TakeOverTypeModelMaster {
	req := TakeOverTypeModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTakeOverTypeModelMasterFromDict(data map[string]interface{}) TakeOverTypeModelMaster {
	return TakeOverTypeModelMaster{
		TakeOverTypeModelId: func() *string {
			v, ok := data["takeOverTypeModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["takeOverTypeModelId"])
		}(),
		Type: func() *int32 {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["type"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		OpenIdConnectSetting: func() *OpenIdConnectSetting {
			v, ok := data["openIdConnectSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewOpenIdConnectSettingFromDict(core.CastMap(data["openIdConnectSetting"])).Pointer()
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p TakeOverTypeModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TakeOverTypeModelId != nil {
		m["takeOverTypeModelId"] = p.TakeOverTypeModelId
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.OpenIdConnectSetting != nil {
		m["openIdConnectSetting"] = func() map[string]interface{} {
			if p.OpenIdConnectSetting == nil {
				return nil
			}
			return p.OpenIdConnectSetting.ToDict()
		}()
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p TakeOverTypeModelMaster) Pointer() *TakeOverTypeModelMaster {
	return &p
}

func CastTakeOverTypeModelMasters(data []interface{}) []TakeOverTypeModelMaster {
	v := make([]TakeOverTypeModelMaster, 0)
	for _, d := range data {
		v = append(v, NewTakeOverTypeModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTakeOverTypeModelMastersFromDict(data []TakeOverTypeModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentModelMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentModelMaster{}
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
		*p = CurrentModelMaster{}
	} else {
		*p = CurrentModelMaster{}
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
		if v, ok := d["settings"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Settings = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Settings = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Settings = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Settings)
				}
			}
		}
	}
	return nil
}

func NewCurrentModelMasterFromJson(data string) CurrentModelMaster {
	req := CurrentModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentModelMasterFromDict(data map[string]interface{}) CurrentModelMaster {
	return CurrentModelMaster{
		NamespaceId: func() *string {
			v, ok := data["namespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceId"])
		}(),
		Settings: func() *string {
			v, ok := data["settings"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["settings"])
		}(),
	}
}

func (p CurrentModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
}

func (p CurrentModelMaster) Pointer() *CurrentModelMaster {
	return &p
}

func CastCurrentModelMasters(data []interface{}) []CurrentModelMaster {
	v := make([]CurrentModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentModelMastersFromDict(data []CurrentModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type OpenIdConnectSetting struct {
	ConfigurationPath      *string      `json:"configurationPath"`
	ClientId               *string      `json:"clientId"`
	ClientSecret           *string      `json:"clientSecret"`
	AppleTeamId            *string      `json:"appleTeamId"`
	AppleKeyId             *string      `json:"appleKeyId"`
	ApplePrivateKeyPem     *string      `json:"applePrivateKeyPem"`
	DoneEndpointUrl        *string      `json:"doneEndpointUrl"`
	AdditionalScopeValues  []ScopeValue `json:"additionalScopeValues"`
	AdditionalReturnValues []*string    `json:"additionalReturnValues"`
}

func (p *OpenIdConnectSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = OpenIdConnectSetting{}
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
		*p = OpenIdConnectSetting{}
	} else {
		*p = OpenIdConnectSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["configurationPath"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ConfigurationPath = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ConfigurationPath = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ConfigurationPath = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ConfigurationPath = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ConfigurationPath = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ConfigurationPath)
				}
			}
		}
		if v, ok := d["clientId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClientId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClientId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClientId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClientId)
				}
			}
		}
		if v, ok := d["clientSecret"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClientSecret = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClientSecret = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClientSecret = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClientSecret = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClientSecret = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClientSecret)
				}
			}
		}
		if v, ok := d["appleTeamId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AppleTeamId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AppleTeamId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AppleTeamId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AppleTeamId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AppleTeamId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AppleTeamId)
				}
			}
		}
		if v, ok := d["appleKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AppleKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AppleKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AppleKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AppleKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AppleKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AppleKeyId)
				}
			}
		}
		if v, ok := d["applePrivateKeyPem"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ApplePrivateKeyPem = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ApplePrivateKeyPem = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ApplePrivateKeyPem = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ApplePrivateKeyPem = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ApplePrivateKeyPem = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ApplePrivateKeyPem)
				}
			}
		}
		if v, ok := d["doneEndpointUrl"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneEndpointUrl = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneEndpointUrl = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneEndpointUrl = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneEndpointUrl = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneEndpointUrl = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneEndpointUrl)
				}
			}
		}
		if v, ok := d["additionalScopeValues"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AdditionalScopeValues)
		}
		if v, ok := d["additionalReturnValues"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.AdditionalReturnValues = l
			}
		}
	}
	return nil
}

func NewOpenIdConnectSettingFromJson(data string) OpenIdConnectSetting {
	req := OpenIdConnectSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewOpenIdConnectSettingFromDict(data map[string]interface{}) OpenIdConnectSetting {
	return OpenIdConnectSetting{
		ConfigurationPath: func() *string {
			v, ok := data["configurationPath"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["configurationPath"])
		}(),
		ClientId: func() *string {
			v, ok := data["clientId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clientId"])
		}(),
		ClientSecret: func() *string {
			v, ok := data["clientSecret"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clientSecret"])
		}(),
		AppleTeamId: func() *string {
			v, ok := data["appleTeamId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["appleTeamId"])
		}(),
		AppleKeyId: func() *string {
			v, ok := data["appleKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["appleKeyId"])
		}(),
		ApplePrivateKeyPem: func() *string {
			v, ok := data["applePrivateKeyPem"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["applePrivateKeyPem"])
		}(),
		DoneEndpointUrl: func() *string {
			v, ok := data["doneEndpointUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneEndpointUrl"])
		}(),
		AdditionalScopeValues: func() []ScopeValue {
			if data["additionalScopeValues"] == nil {
				return nil
			}
			return CastScopeValues(core.CastArray(data["additionalScopeValues"]))
		}(),
		AdditionalReturnValues: func() []*string {
			v, ok := data["additionalReturnValues"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
	}
}

func (p OpenIdConnectSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ConfigurationPath != nil {
		m["configurationPath"] = p.ConfigurationPath
	}
	if p.ClientId != nil {
		m["clientId"] = p.ClientId
	}
	if p.ClientSecret != nil {
		m["clientSecret"] = p.ClientSecret
	}
	if p.AppleTeamId != nil {
		m["appleTeamId"] = p.AppleTeamId
	}
	if p.AppleKeyId != nil {
		m["appleKeyId"] = p.AppleKeyId
	}
	if p.ApplePrivateKeyPem != nil {
		m["applePrivateKeyPem"] = p.ApplePrivateKeyPem
	}
	if p.DoneEndpointUrl != nil {
		m["doneEndpointUrl"] = p.DoneEndpointUrl
	}
	if p.AdditionalScopeValues != nil {
		m["additionalScopeValues"] = CastScopeValuesFromDict(
			p.AdditionalScopeValues,
		)
	}
	if p.AdditionalReturnValues != nil {
		m["additionalReturnValues"] = core.CastStringsFromDict(
			p.AdditionalReturnValues,
		)
	}
	return m
}

func (p OpenIdConnectSetting) Pointer() *OpenIdConnectSetting {
	return &p
}

func CastOpenIdConnectSettings(data []interface{}) []OpenIdConnectSetting {
	v := make([]OpenIdConnectSetting, 0)
	for _, d := range data {
		v = append(v, NewOpenIdConnectSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastOpenIdConnectSettingsFromDict(data []OpenIdConnectSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ScopeValue struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

func (p *ScopeValue) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ScopeValue{}
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
		*p = ScopeValue{}
	} else {
		*p = ScopeValue{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["key"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Key = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Key = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Key = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Key)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Value = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Value = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Value = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Value)
				}
			}
		}
	}
	return nil
}

func NewScopeValueFromJson(data string) ScopeValue {
	req := ScopeValue{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewScopeValueFromDict(data map[string]interface{}) ScopeValue {
	return ScopeValue{
		Key: func() *string {
			v, ok := data["key"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["key"])
		}(),
		Value: func() *string {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["value"])
		}(),
	}
}

func (p ScopeValue) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Key != nil {
		m["key"] = p.Key
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
}

func (p ScopeValue) Pointer() *ScopeValue {
	return &p
}

func CastScopeValues(data []interface{}) []ScopeValue {
	v := make([]ScopeValue, 0)
	for _, d := range data {
		v = append(v, NewScopeValueFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScopeValuesFromDict(data []ScopeValue) []interface{} {
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
		Type: func() *int32 {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["type"])
		}(),
		UserIdentifier: func() *string {
			v, ok := data["userIdentifier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userIdentifier"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
	}
}

func (p PlatformUser) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.UserIdentifier != nil {
		m["userIdentifier"] = p.UserIdentifier
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	return m
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
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Reason: func() *string {
			v, ok := data["reason"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["reason"])
		}(),
		ReleaseTimestamp: func() *int64 {
			v, ok := data["releaseTimestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["releaseTimestamp"])
		}(),
	}
}

func (p BanStatus) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Reason != nil {
		m["reason"] = p.Reason
	}
	if p.ReleaseTimestamp != nil {
		m["releaseTimestamp"] = p.ReleaseTimestamp
	}
	return m
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

type GitHubCheckoutSetting struct {
	ApiKeyId       *string `json:"apiKeyId"`
	RepositoryName *string `json:"repositoryName"`
	SourcePath     *string `json:"sourcePath"`
	ReferenceType  *string `json:"referenceType"`
	CommitHash     *string `json:"commitHash"`
	BranchName     *string `json:"branchName"`
	TagName        *string `json:"tagName"`
}

func (p *GitHubCheckoutSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GitHubCheckoutSetting{}
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
		*p = GitHubCheckoutSetting{}
	} else {
		*p = GitHubCheckoutSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["apiKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ApiKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ApiKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ApiKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ApiKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ApiKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ApiKeyId)
				}
			}
		}
		if v, ok := d["repositoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepositoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepositoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepositoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepositoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepositoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepositoryName)
				}
			}
		}
		if v, ok := d["sourcePath"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SourcePath = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SourcePath = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SourcePath = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SourcePath = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SourcePath = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SourcePath)
				}
			}
		}
		if v, ok := d["referenceType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceType)
				}
			}
		}
		if v, ok := d["commitHash"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CommitHash = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CommitHash = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CommitHash = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CommitHash = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CommitHash = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CommitHash)
				}
			}
		}
		if v, ok := d["branchName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BranchName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BranchName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BranchName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BranchName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BranchName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BranchName)
				}
			}
		}
		if v, ok := d["tagName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TagName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TagName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TagName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TagName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TagName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TagName)
				}
			}
		}
	}
	return nil
}

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
	req := GitHubCheckoutSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGitHubCheckoutSettingFromDict(data map[string]interface{}) GitHubCheckoutSetting {
	return GitHubCheckoutSetting{
		ApiKeyId: func() *string {
			v, ok := data["apiKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["apiKeyId"])
		}(),
		RepositoryName: func() *string {
			v, ok := data["repositoryName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repositoryName"])
		}(),
		SourcePath: func() *string {
			v, ok := data["sourcePath"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sourcePath"])
		}(),
		ReferenceType: func() *string {
			v, ok := data["referenceType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["referenceType"])
		}(),
		CommitHash: func() *string {
			v, ok := data["commitHash"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["commitHash"])
		}(),
		BranchName: func() *string {
			v, ok := data["branchName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["branchName"])
		}(),
		TagName: func() *string {
			v, ok := data["tagName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["tagName"])
		}(),
	}
}

func (p GitHubCheckoutSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ApiKeyId != nil {
		m["apiKeyId"] = p.ApiKeyId
	}
	if p.RepositoryName != nil {
		m["repositoryName"] = p.RepositoryName
	}
	if p.SourcePath != nil {
		m["sourcePath"] = p.SourcePath
	}
	if p.ReferenceType != nil {
		m["referenceType"] = p.ReferenceType
	}
	if p.CommitHash != nil {
		m["commitHash"] = p.CommitHash
	}
	if p.BranchName != nil {
		m["branchName"] = p.BranchName
	}
	if p.TagName != nil {
		m["tagName"] = p.TagName
	}
	return m
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
		TriggerScriptId: func() *string {
			v, ok := data["triggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerScriptId"])
		}(),
		DoneTriggerTargetType: func() *string {
			v, ok := data["doneTriggerTargetType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerTargetType"])
		}(),
		DoneTriggerScriptId: func() *string {
			v, ok := data["doneTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerScriptId"])
		}(),
		DoneTriggerQueueNamespaceId: func() *string {
			v, ok := data["doneTriggerQueueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerQueueNamespaceId"])
		}(),
	}
}

func (p ScriptSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TriggerScriptId != nil {
		m["triggerScriptId"] = p.TriggerScriptId
	}
	if p.DoneTriggerTargetType != nil {
		m["doneTriggerTargetType"] = p.DoneTriggerTargetType
	}
	if p.DoneTriggerScriptId != nil {
		m["doneTriggerScriptId"] = p.DoneTriggerScriptId
	}
	if p.DoneTriggerQueueNamespaceId != nil {
		m["doneTriggerQueueNamespaceId"] = p.DoneTriggerQueueNamespaceId
	}
	return m
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
		LoggingNamespaceId: func() *string {
			v, ok := data["loggingNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["loggingNamespaceId"])
		}(),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LoggingNamespaceId != nil {
		m["loggingNamespaceId"] = p.LoggingNamespaceId
	}
	return m
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
