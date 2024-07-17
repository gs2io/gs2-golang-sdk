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

package identifier

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type User struct {
	UserId      *string `json:"userId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
	Revision    *int64  `json:"revision"`
}

func (p *User) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = User{}
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
		*p = User{}
	} else {
		*p = User{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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

func NewUserFromJson(data string) User {
	req := User{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewUserFromDict(data map[string]interface{}) User {
	return User{
		UserId:      core.CastString(data["userId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
		Revision:    core.CastInt64(data["revision"]),
	}
}

func (p User) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
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
		"userId":      userId,
		"name":        name,
		"description": description,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
	}
}

func (p User) Pointer() *User {
	return &p
}

func CastUsers(data []interface{}) []User {
	v := make([]User, 0)
	for _, d := range data {
		v = append(v, NewUserFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastUsersFromDict(data []User) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SecurityPolicy struct {
	SecurityPolicyId *string `json:"securityPolicyId"`
	Name             *string `json:"name"`
	Description      *string `json:"description"`
	Policy           *string `json:"policy"`
	CreatedAt        *int64  `json:"createdAt"`
	UpdatedAt        *int64  `json:"updatedAt"`
}

func (p *SecurityPolicy) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SecurityPolicy{}
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
		*p = SecurityPolicy{}
	} else {
		*p = SecurityPolicy{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["securityPolicyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SecurityPolicyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SecurityPolicyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SecurityPolicyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SecurityPolicyId)
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
		if v, ok := d["policy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Policy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Policy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Policy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Policy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Policy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Policy)
				}
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
	}
	return nil
}

func NewSecurityPolicyFromJson(data string) SecurityPolicy {
	req := SecurityPolicy{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSecurityPolicyFromDict(data map[string]interface{}) SecurityPolicy {
	return SecurityPolicy{
		SecurityPolicyId: core.CastString(data["securityPolicyId"]),
		Name:             core.CastString(data["name"]),
		Description:      core.CastString(data["description"]),
		Policy:           core.CastString(data["policy"]),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
	}
}

func (p SecurityPolicy) ToDict() map[string]interface{} {

	var securityPolicyId *string
	if p.SecurityPolicyId != nil {
		securityPolicyId = p.SecurityPolicyId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var policy *string
	if p.Policy != nil {
		policy = p.Policy
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
		"securityPolicyId": securityPolicyId,
		"name":             name,
		"description":      description,
		"policy":           policy,
		"createdAt":        createdAt,
		"updatedAt":        updatedAt,
	}
}

func (p SecurityPolicy) Pointer() *SecurityPolicy {
	return &p
}

func CastSecurityPolicies(data []interface{}) []SecurityPolicy {
	v := make([]SecurityPolicy, 0)
	for _, d := range data {
		v = append(v, NewSecurityPolicyFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSecurityPoliciesFromDict(data []SecurityPolicy) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Identifier struct {
	ClientId     *string `json:"clientId"`
	UserName     *string `json:"userName"`
	ClientSecret *string `json:"clientSecret"`
	CreatedAt    *int64  `json:"createdAt"`
	Revision     *int64  `json:"revision"`
}

func (p *Identifier) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Identifier{}
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
		*p = Identifier{}
	} else {
		*p = Identifier{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["userName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserName)
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
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewIdentifierFromJson(data string) Identifier {
	req := Identifier{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewIdentifierFromDict(data map[string]interface{}) Identifier {
	return Identifier{
		ClientId:     core.CastString(data["clientId"]),
		UserName:     core.CastString(data["userName"]),
		ClientSecret: core.CastString(data["clientSecret"]),
		CreatedAt:    core.CastInt64(data["createdAt"]),
		Revision:     core.CastInt64(data["revision"]),
	}
}

func (p Identifier) ToDict() map[string]interface{} {

	var clientId *string
	if p.ClientId != nil {
		clientId = p.ClientId
	}
	var userName *string
	if p.UserName != nil {
		userName = p.UserName
	}
	var clientSecret *string
	if p.ClientSecret != nil {
		clientSecret = p.ClientSecret
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
		"clientId":     clientId,
		"userName":     userName,
		"clientSecret": clientSecret,
		"createdAt":    createdAt,
		"revision":     revision,
	}
}

func (p Identifier) Pointer() *Identifier {
	return &p
}

func CastIdentifiers(data []interface{}) []Identifier {
	v := make([]Identifier, 0)
	for _, d := range data {
		v = append(v, NewIdentifierFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastIdentifiersFromDict(data []Identifier) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Password struct {
	PasswordId                     *string                         `json:"passwordId"`
	UserId                         *string                         `json:"userId"`
	UserName                       *string                         `json:"userName"`
	EnableTwoFactorAuthentication  *string                         `json:"enableTwoFactorAuthentication"`
	TwoFactorAuthenticationSetting *TwoFactorAuthenticationSetting `json:"twoFactorAuthenticationSetting"`
	CreatedAt                      *int64                          `json:"createdAt"`
	Revision                       *int64                          `json:"revision"`
}

func (p *Password) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Password{}
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
		*p = Password{}
	} else {
		*p = Password{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["passwordId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PasswordId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PasswordId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PasswordId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PasswordId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PasswordId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PasswordId)
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
		if v, ok := d["userName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserName)
				}
			}
		}
		if v, ok := d["enableTwoFactorAuthentication"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableTwoFactorAuthentication = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableTwoFactorAuthentication = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableTwoFactorAuthentication = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableTwoFactorAuthentication = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableTwoFactorAuthentication = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableTwoFactorAuthentication)
				}
			}
		}
		if v, ok := d["twoFactorAuthenticationSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TwoFactorAuthenticationSetting)
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

func NewPasswordFromJson(data string) Password {
	req := Password{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPasswordFromDict(data map[string]interface{}) Password {
	return Password{
		PasswordId:                     core.CastString(data["passwordId"]),
		UserId:                         core.CastString(data["userId"]),
		UserName:                       core.CastString(data["userName"]),
		EnableTwoFactorAuthentication:  core.CastString(data["enableTwoFactorAuthentication"]),
		TwoFactorAuthenticationSetting: NewTwoFactorAuthenticationSettingFromDict(core.CastMap(data["twoFactorAuthenticationSetting"])).Pointer(),
		CreatedAt:                      core.CastInt64(data["createdAt"]),
		Revision:                       core.CastInt64(data["revision"]),
	}
}

func (p Password) ToDict() map[string]interface{} {

	var passwordId *string
	if p.PasswordId != nil {
		passwordId = p.PasswordId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var userName *string
	if p.UserName != nil {
		userName = p.UserName
	}
	var enableTwoFactorAuthentication *string
	if p.EnableTwoFactorAuthentication != nil {
		enableTwoFactorAuthentication = p.EnableTwoFactorAuthentication
	}
	var twoFactorAuthenticationSetting map[string]interface{}
	if p.TwoFactorAuthenticationSetting != nil {
		twoFactorAuthenticationSetting = p.TwoFactorAuthenticationSetting.ToDict()
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
		"passwordId":                     passwordId,
		"userId":                         userId,
		"userName":                       userName,
		"enableTwoFactorAuthentication":  enableTwoFactorAuthentication,
		"twoFactorAuthenticationSetting": twoFactorAuthenticationSetting,
		"createdAt":                      createdAt,
		"revision":                       revision,
	}
}

func (p Password) Pointer() *Password {
	return &p
}

func CastPasswords(data []interface{}) []Password {
	v := make([]Password, 0)
	for _, d := range data {
		v = append(v, NewPasswordFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPasswordsFromDict(data []Password) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AttachSecurityPolicy struct {
	UserId            *string   `json:"userId"`
	SecurityPolicyIds []*string `json:"securityPolicyIds"`
	AttachedAt        *int64    `json:"attachedAt"`
	Revision          *int64    `json:"revision"`
}

func (p *AttachSecurityPolicy) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AttachSecurityPolicy{}
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
		*p = AttachSecurityPolicy{}
	} else {
		*p = AttachSecurityPolicy{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["securityPolicyIds"]; ok && v != nil {
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
				p.SecurityPolicyIds = l
			}
		}
		if v, ok := d["attachedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AttachedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewAttachSecurityPolicyFromJson(data string) AttachSecurityPolicy {
	req := AttachSecurityPolicy{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAttachSecurityPolicyFromDict(data map[string]interface{}) AttachSecurityPolicy {
	return AttachSecurityPolicy{
		UserId:            core.CastString(data["userId"]),
		SecurityPolicyIds: core.CastStrings(core.CastArray(data["securityPolicyIds"])),
		AttachedAt:        core.CastInt64(data["attachedAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p AttachSecurityPolicy) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var securityPolicyIds []interface{}
	if p.SecurityPolicyIds != nil {
		securityPolicyIds = core.CastStringsFromDict(
			p.SecurityPolicyIds,
		)
	}
	var attachedAt *int64
	if p.AttachedAt != nil {
		attachedAt = p.AttachedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"userId":            userId,
		"securityPolicyIds": securityPolicyIds,
		"attachedAt":        attachedAt,
		"revision":          revision,
	}
}

func (p AttachSecurityPolicy) Pointer() *AttachSecurityPolicy {
	return &p
}

func CastAttachSecurityPolicies(data []interface{}) []AttachSecurityPolicy {
	v := make([]AttachSecurityPolicy, 0)
	for _, d := range data {
		v = append(v, NewAttachSecurityPolicyFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAttachSecurityPoliciesFromDict(data []AttachSecurityPolicy) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ProjectToken struct {
	Token *string `json:"token"`
}

func (p *ProjectToken) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ProjectToken{}
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
		*p = ProjectToken{}
	} else {
		*p = ProjectToken{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["token"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Token = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Token = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Token = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Token = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Token = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Token)
				}
			}
		}
	}
	return nil
}

func NewProjectTokenFromJson(data string) ProjectToken {
	req := ProjectToken{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewProjectTokenFromDict(data map[string]interface{}) ProjectToken {
	return ProjectToken{
		Token: core.CastString(data["token"]),
	}
}

func (p ProjectToken) ToDict() map[string]interface{} {

	var token *string
	if p.Token != nil {
		token = p.Token
	}
	return map[string]interface{}{
		"token": token,
	}
}

func (p ProjectToken) Pointer() *ProjectToken {
	return &p
}

func CastProjectTokens(data []interface{}) []ProjectToken {
	v := make([]ProjectToken, 0)
	for _, d := range data {
		v = append(v, NewProjectTokenFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProjectTokensFromDict(data []ProjectToken) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TwoFactorAuthenticationSetting struct {
	Status *string `json:"status"`
}

func (p *TwoFactorAuthenticationSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TwoFactorAuthenticationSetting{}
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
		*p = TwoFactorAuthenticationSetting{}
	} else {
		*p = TwoFactorAuthenticationSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["status"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Status = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Status = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Status = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Status = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Status = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Status)
				}
			}
		}
	}
	return nil
}

func NewTwoFactorAuthenticationSettingFromJson(data string) TwoFactorAuthenticationSetting {
	req := TwoFactorAuthenticationSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTwoFactorAuthenticationSettingFromDict(data map[string]interface{}) TwoFactorAuthenticationSetting {
	return TwoFactorAuthenticationSetting{
		Status: core.CastString(data["status"]),
	}
}

func (p TwoFactorAuthenticationSetting) ToDict() map[string]interface{} {

	var status *string
	if p.Status != nil {
		status = p.Status
	}
	return map[string]interface{}{
		"status": status,
	}
}

func (p TwoFactorAuthenticationSetting) Pointer() *TwoFactorAuthenticationSetting {
	return &p
}

func CastTwoFactorAuthenticationSettings(data []interface{}) []TwoFactorAuthenticationSetting {
	v := make([]TwoFactorAuthenticationSetting, 0)
	for _, d := range data {
		v = append(v, NewTwoFactorAuthenticationSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTwoFactorAuthenticationSettingsFromDict(data []TwoFactorAuthenticationSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
