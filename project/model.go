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

package project

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Account struct {
	AccountId                      *string                         `json:"accountId"`
	Name                           *string                         `json:"name"`
	Email                          *string                         `json:"email"`
	FullName                       *string                         `json:"fullName"`
	CompanyName                    *string                         `json:"companyName"`
	EnableTwoFactorAuthentication  *string                         `json:"enableTwoFactorAuthentication"`
	TwoFactorAuthenticationSetting *TwoFactorAuthenticationSetting `json:"twoFactorAuthenticationSetting"`
	Status                         *string                         `json:"status"`
	CreatedAt                      *int64                          `json:"createdAt"`
	UpdatedAt                      *int64                          `json:"updatedAt"`
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
		if v, ok := d["email"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Email = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Email = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Email = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Email)
				}
			}
		}
		if v, ok := d["fullName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FullName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FullName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FullName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FullName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FullName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FullName)
				}
			}
		}
		if v, ok := d["companyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompanyName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompanyName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompanyName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompanyName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompanyName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompanyName)
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
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
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
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Email: func() *string {
			v, ok := data["email"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["email"])
		}(),
		FullName: func() *string {
			v, ok := data["fullName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fullName"])
		}(),
		CompanyName: func() *string {
			v, ok := data["companyName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["companyName"])
		}(),
		EnableTwoFactorAuthentication: func() *string {
			v, ok := data["enableTwoFactorAuthentication"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableTwoFactorAuthentication"])
		}(),
		TwoFactorAuthenticationSetting: func() *TwoFactorAuthenticationSetting {
			v, ok := data["twoFactorAuthenticationSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewTwoFactorAuthenticationSettingFromDict(core.CastMap(data["twoFactorAuthenticationSetting"])).Pointer()
		}(),
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
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
	}
}

func (p Account) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.AccountId != nil {
		m["accountId"] = p.AccountId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Email != nil {
		m["email"] = p.Email
	}
	if p.FullName != nil {
		m["fullName"] = p.FullName
	}
	if p.CompanyName != nil {
		m["companyName"] = p.CompanyName
	}
	if p.EnableTwoFactorAuthentication != nil {
		m["enableTwoFactorAuthentication"] = p.EnableTwoFactorAuthentication
	}
	if p.TwoFactorAuthenticationSetting != nil {
		m["twoFactorAuthenticationSetting"] = func() map[string]interface{} {
			if p.TwoFactorAuthenticationSetting == nil {
				return nil
			}
			return p.TwoFactorAuthenticationSetting.ToDict()
		}()
	}
	if p.Status != nil {
		m["status"] = p.Status
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
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

type Project struct {
	ProjectId               *string     `json:"projectId"`
	AccountName             *string     `json:"accountName"`
	Name                    *string     `json:"name"`
	Description             *string     `json:"description"`
	Plan                    *string     `json:"plan"`
	Regions                 []Gs2Region `json:"regions"`
	BillingMethodName       *string     `json:"billingMethodName"`
	EnableEventBridge       *string     `json:"enableEventBridge"`
	Currency                *string     `json:"currency"`
	EventBridgeAwsAccountId *string     `json:"eventBridgeAwsAccountId"`
	EventBridgeAwsRegion    *string     `json:"eventBridgeAwsRegion"`
	CreatedAt               *int64      `json:"createdAt"`
	UpdatedAt               *int64      `json:"updatedAt"`
}

func (p *Project) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Project{}
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
		*p = Project{}
	} else {
		*p = Project{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["projectId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectId)
				}
			}
		}
		if v, ok := d["accountName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountName)
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
		if v, ok := d["plan"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Plan = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Plan = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Plan = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Plan = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Plan = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Plan)
				}
			}
		}
		if v, ok := d["regions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Regions)
		}
		if v, ok := d["billingMethodName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BillingMethodName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BillingMethodName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BillingMethodName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BillingMethodName)
				}
			}
		}
		if v, ok := d["enableEventBridge"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableEventBridge = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableEventBridge = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableEventBridge = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableEventBridge = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableEventBridge = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableEventBridge)
				}
			}
		}
		if v, ok := d["currency"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Currency = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Currency = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Currency = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Currency)
				}
			}
		}
		if v, ok := d["eventBridgeAwsAccountId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventBridgeAwsAccountId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventBridgeAwsAccountId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventBridgeAwsAccountId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsAccountId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsAccountId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventBridgeAwsAccountId)
				}
			}
		}
		if v, ok := d["eventBridgeAwsRegion"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventBridgeAwsRegion = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventBridgeAwsRegion = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventBridgeAwsRegion = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsRegion = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsRegion = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventBridgeAwsRegion)
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

func NewProjectFromJson(data string) Project {
	req := Project{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewProjectFromDict(data map[string]interface{}) Project {
	return Project{
		ProjectId: func() *string {
			v, ok := data["projectId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectId"])
		}(),
		AccountName: func() *string {
			v, ok := data["accountName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountName"])
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
		Plan: func() *string {
			v, ok := data["plan"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["plan"])
		}(),
		Regions: func() []Gs2Region {
			if data["regions"] == nil {
				return nil
			}
			return CastGs2Regions(core.CastArray(data["regions"]))
		}(),
		BillingMethodName: func() *string {
			v, ok := data["billingMethodName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["billingMethodName"])
		}(),
		EnableEventBridge: func() *string {
			v, ok := data["enableEventBridge"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableEventBridge"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
		}(),
		EventBridgeAwsAccountId: func() *string {
			v, ok := data["eventBridgeAwsAccountId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventBridgeAwsAccountId"])
		}(),
		EventBridgeAwsRegion: func() *string {
			v, ok := data["eventBridgeAwsRegion"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventBridgeAwsRegion"])
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
	}
}

func (p Project) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ProjectId != nil {
		m["projectId"] = p.ProjectId
	}
	if p.AccountName != nil {
		m["accountName"] = p.AccountName
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.Plan != nil {
		m["plan"] = p.Plan
	}
	if p.Regions != nil {
		m["regions"] = CastGs2RegionsFromDict(
			p.Regions,
		)
	}
	if p.BillingMethodName != nil {
		m["billingMethodName"] = p.BillingMethodName
	}
	if p.EnableEventBridge != nil {
		m["enableEventBridge"] = p.EnableEventBridge
	}
	if p.Currency != nil {
		m["currency"] = p.Currency
	}
	if p.EventBridgeAwsAccountId != nil {
		m["eventBridgeAwsAccountId"] = p.EventBridgeAwsAccountId
	}
	if p.EventBridgeAwsRegion != nil {
		m["eventBridgeAwsRegion"] = p.EventBridgeAwsRegion
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	return m
}

func (p Project) Pointer() *Project {
	return &p
}

func CastProjects(data []interface{}) []Project {
	v := make([]Project, 0)
	for _, d := range data {
		v = append(v, NewProjectFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProjectsFromDict(data []Project) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Gs2Region struct {
	RegionName *string `json:"regionName"`
	Status     *string `json:"status"`
}

func (p *Gs2Region) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Gs2Region{}
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
		*p = Gs2Region{}
	} else {
		*p = Gs2Region{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["regionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RegionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RegionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RegionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RegionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RegionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RegionName)
				}
			}
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

func NewGs2RegionFromJson(data string) Gs2Region {
	req := Gs2Region{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGs2RegionFromDict(data map[string]interface{}) Gs2Region {
	return Gs2Region{
		RegionName: func() *string {
			v, ok := data["regionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["regionName"])
		}(),
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
	}
}

func (p Gs2Region) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RegionName != nil {
		m["regionName"] = p.RegionName
	}
	if p.Status != nil {
		m["status"] = p.Status
	}
	return m
}

func (p Gs2Region) Pointer() *Gs2Region {
	return &p
}

func CastGs2Regions(data []interface{}) []Gs2Region {
	v := make([]Gs2Region, 0)
	for _, d := range data {
		v = append(v, NewGs2RegionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGs2RegionsFromDict(data []Gs2Region) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BillingMethod struct {
	BillingMethodId   *string `json:"billingMethodId"`
	AccountName       *string `json:"accountName"`
	Name              *string `json:"name"`
	Description       *string `json:"description"`
	MethodType        *string `json:"methodType"`
	CardSignatureName *string `json:"cardSignatureName"`
	CardBrand         *string `json:"cardBrand"`
	CardLast4         *string `json:"cardLast4"`
	PartnerId         *string `json:"partnerId"`
	CreatedAt         *int64  `json:"createdAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
}

func (p *BillingMethod) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BillingMethod{}
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
		*p = BillingMethod{}
	} else {
		*p = BillingMethod{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["billingMethodId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BillingMethodId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BillingMethodId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BillingMethodId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BillingMethodId)
				}
			}
		}
		if v, ok := d["accountName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountName)
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
		if v, ok := d["methodType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MethodType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MethodType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MethodType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MethodType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MethodType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MethodType)
				}
			}
		}
		if v, ok := d["cardSignatureName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CardSignatureName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CardSignatureName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CardSignatureName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CardSignatureName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CardSignatureName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CardSignatureName)
				}
			}
		}
		if v, ok := d["cardBrand"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CardBrand = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CardBrand = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CardBrand = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CardBrand = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CardBrand = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CardBrand)
				}
			}
		}
		if v, ok := d["cardLast4"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CardLast4 = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CardLast4 = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CardLast4 = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CardLast4 = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CardLast4 = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CardLast4)
				}
			}
		}
		if v, ok := d["partnerId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PartnerId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PartnerId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PartnerId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PartnerId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PartnerId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PartnerId)
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

func NewBillingMethodFromJson(data string) BillingMethod {
	req := BillingMethod{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBillingMethodFromDict(data map[string]interface{}) BillingMethod {
	return BillingMethod{
		BillingMethodId: func() *string {
			v, ok := data["billingMethodId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["billingMethodId"])
		}(),
		AccountName: func() *string {
			v, ok := data["accountName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountName"])
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
		MethodType: func() *string {
			v, ok := data["methodType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["methodType"])
		}(),
		CardSignatureName: func() *string {
			v, ok := data["cardSignatureName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["cardSignatureName"])
		}(),
		CardBrand: func() *string {
			v, ok := data["cardBrand"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["cardBrand"])
		}(),
		CardLast4: func() *string {
			v, ok := data["cardLast4"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["cardLast4"])
		}(),
		PartnerId: func() *string {
			v, ok := data["partnerId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["partnerId"])
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
	}
}

func (p BillingMethod) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.BillingMethodId != nil {
		m["billingMethodId"] = p.BillingMethodId
	}
	if p.AccountName != nil {
		m["accountName"] = p.AccountName
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.MethodType != nil {
		m["methodType"] = p.MethodType
	}
	if p.CardSignatureName != nil {
		m["cardSignatureName"] = p.CardSignatureName
	}
	if p.CardBrand != nil {
		m["cardBrand"] = p.CardBrand
	}
	if p.CardLast4 != nil {
		m["cardLast4"] = p.CardLast4
	}
	if p.PartnerId != nil {
		m["partnerId"] = p.PartnerId
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	return m
}

func (p BillingMethod) Pointer() *BillingMethod {
	return &p
}

func CastBillingMethods(data []interface{}) []BillingMethod {
	v := make([]BillingMethod, 0)
	for _, d := range data {
		v = append(v, NewBillingMethodFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBillingMethodsFromDict(data []BillingMethod) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Receipt struct {
	ReceiptId   *string `json:"receiptId"`
	AccountName *string `json:"accountName"`
	Name        *string `json:"name"`
	Date        *int64  `json:"date"`
	Amount      *string `json:"amount"`
	PdfUrl      *string `json:"pdfUrl"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
}

func (p *Receipt) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Receipt{}
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
		*p = Receipt{}
	} else {
		*p = Receipt{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["receiptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReceiptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReceiptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReceiptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReceiptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReceiptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReceiptId)
				}
			}
		}
		if v, ok := d["accountName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountName)
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
		if v, ok := d["date"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Date)
		}
		if v, ok := d["amount"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Amount = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Amount = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Amount = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Amount = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Amount = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Amount)
				}
			}
		}
		if v, ok := d["pdfUrl"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PdfUrl = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PdfUrl = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PdfUrl = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PdfUrl = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PdfUrl = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PdfUrl)
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

func NewReceiptFromJson(data string) Receipt {
	req := Receipt{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewReceiptFromDict(data map[string]interface{}) Receipt {
	return Receipt{
		ReceiptId: func() *string {
			v, ok := data["receiptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["receiptId"])
		}(),
		AccountName: func() *string {
			v, ok := data["accountName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountName"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Date: func() *int64 {
			v, ok := data["date"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["date"])
		}(),
		Amount: func() *string {
			v, ok := data["amount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["amount"])
		}(),
		PdfUrl: func() *string {
			v, ok := data["pdfUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pdfUrl"])
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
	}
}

func (p Receipt) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ReceiptId != nil {
		m["receiptId"] = p.ReceiptId
	}
	if p.AccountName != nil {
		m["accountName"] = p.AccountName
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Date != nil {
		m["date"] = p.Date
	}
	if p.Amount != nil {
		m["amount"] = p.Amount
	}
	if p.PdfUrl != nil {
		m["pdfUrl"] = p.PdfUrl
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	return m
}

func (p Receipt) Pointer() *Receipt {
	return &p
}

func CastReceipts(data []interface{}) []Receipt {
	v := make([]Receipt, 0)
	for _, d := range data {
		v = append(v, NewReceiptFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReceiptsFromDict(data []Receipt) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Billing struct {
	BillingId    *string  `json:"billingId"`
	ProjectName  *string  `json:"projectName"`
	Year         *int32   `json:"year"`
	Month        *int32   `json:"month"`
	Region       *string  `json:"region"`
	Service      *string  `json:"service"`
	ActivityType *string  `json:"activityType"`
	Unit         *float64 `json:"unit"`
	UnitName     *string  `json:"unitName"`
	Price        *float64 `json:"price"`
	Currency     *string  `json:"currency"`
	CreatedAt    *int64   `json:"createdAt"`
	UpdatedAt    *int64   `json:"updatedAt"`
}

func (p *Billing) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Billing{}
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
		*p = Billing{}
	} else {
		*p = Billing{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["billingId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BillingId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BillingId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BillingId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BillingId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BillingId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BillingId)
				}
			}
		}
		if v, ok := d["projectName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectName)
				}
			}
		}
		if v, ok := d["year"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Year)
		}
		if v, ok := d["month"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Month)
		}
		if v, ok := d["region"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Region = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Region = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Region = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Region = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Region = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Region)
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["activityType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ActivityType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ActivityType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ActivityType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ActivityType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ActivityType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ActivityType)
				}
			}
		}
		if v, ok := d["unit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Unit)
		}
		if v, ok := d["unitName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UnitName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UnitName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UnitName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UnitName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UnitName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UnitName)
				}
			}
		}
		if v, ok := d["price"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Price)
		}
		if v, ok := d["currency"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Currency = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Currency = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Currency = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Currency)
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

func NewBillingFromJson(data string) Billing {
	req := Billing{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBillingFromDict(data map[string]interface{}) Billing {
	return Billing{
		BillingId: func() *string {
			v, ok := data["billingId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["billingId"])
		}(),
		ProjectName: func() *string {
			v, ok := data["projectName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectName"])
		}(),
		Year: func() *int32 {
			v, ok := data["year"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["year"])
		}(),
		Month: func() *int32 {
			v, ok := data["month"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["month"])
		}(),
		Region: func() *string {
			v, ok := data["region"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["region"])
		}(),
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		ActivityType: func() *string {
			v, ok := data["activityType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["activityType"])
		}(),
		Unit: func() *float64 {
			v, ok := data["unit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat64(data["unit"])
		}(),
		UnitName: func() *string {
			v, ok := data["unitName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["unitName"])
		}(),
		Price: func() *float64 {
			v, ok := data["price"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat64(data["price"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
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
	}
}

func (p Billing) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.BillingId != nil {
		m["billingId"] = p.BillingId
	}
	if p.ProjectName != nil {
		m["projectName"] = p.ProjectName
	}
	if p.Year != nil {
		m["year"] = p.Year
	}
	if p.Month != nil {
		m["month"] = p.Month
	}
	if p.Region != nil {
		m["region"] = p.Region
	}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.ActivityType != nil {
		m["activityType"] = p.ActivityType
	}
	if p.Unit != nil {
		m["unit"] = p.Unit
	}
	if p.UnitName != nil {
		m["unitName"] = p.UnitName
	}
	if p.Price != nil {
		m["price"] = p.Price
	}
	if p.Currency != nil {
		m["currency"] = p.Currency
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	return m
}

func (p Billing) Pointer() *Billing {
	return &p
}

func CastBillings(data []interface{}) []Billing {
	v := make([]Billing, 0)
	for _, d := range data {
		v = append(v, NewBillingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBillingsFromDict(data []Billing) []interface{} {
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
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
	}
}

func (p TwoFactorAuthenticationSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Status != nil {
		m["status"] = p.Status
	}
	return m
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

type DumpProgress struct {
	DumpProgressId    *string `json:"dumpProgressId"`
	TransactionId     *string `json:"transactionId"`
	UserId            *string `json:"userId"`
	Dumped            *int32  `json:"dumped"`
	MicroserviceCount *int32  `json:"microserviceCount"`
	CreatedAt         *int64  `json:"createdAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
	Revision          *int64  `json:"revision"`
}

func (p *DumpProgress) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DumpProgress{}
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
		*p = DumpProgress{}
	} else {
		*p = DumpProgress{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["dumpProgressId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DumpProgressId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DumpProgressId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DumpProgressId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DumpProgressId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DumpProgressId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DumpProgressId)
				}
			}
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
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
		if v, ok := d["dumped"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Dumped)
		}
		if v, ok := d["microserviceCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MicroserviceCount)
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

func NewDumpProgressFromJson(data string) DumpProgress {
	req := DumpProgress{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDumpProgressFromDict(data map[string]interface{}) DumpProgress {
	return DumpProgress{
		DumpProgressId: func() *string {
			v, ok := data["dumpProgressId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["dumpProgressId"])
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Dumped: func() *int32 {
			v, ok := data["dumped"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["dumped"])
		}(),
		MicroserviceCount: func() *int32 {
			v, ok := data["microserviceCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["microserviceCount"])
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

func (p DumpProgress) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.DumpProgressId != nil {
		m["dumpProgressId"] = p.DumpProgressId
	}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Dumped != nil {
		m["dumped"] = p.Dumped
	}
	if p.MicroserviceCount != nil {
		m["microserviceCount"] = p.MicroserviceCount
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

func (p DumpProgress) Pointer() *DumpProgress {
	return &p
}

func CastDumpProgresses(data []interface{}) []DumpProgress {
	v := make([]DumpProgress, 0)
	for _, d := range data {
		v = append(v, NewDumpProgressFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDumpProgressesFromDict(data []DumpProgress) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CleanProgress struct {
	CleanProgressId   *string `json:"cleanProgressId"`
	TransactionId     *string `json:"transactionId"`
	UserId            *string `json:"userId"`
	Cleaned           *int32  `json:"cleaned"`
	MicroserviceCount *int32  `json:"microserviceCount"`
	CreatedAt         *int64  `json:"createdAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
	Revision          *int64  `json:"revision"`
}

func (p *CleanProgress) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CleanProgress{}
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
		*p = CleanProgress{}
	} else {
		*p = CleanProgress{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["cleanProgressId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CleanProgressId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CleanProgressId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CleanProgressId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CleanProgressId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CleanProgressId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CleanProgressId)
				}
			}
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
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
		if v, ok := d["cleaned"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Cleaned)
		}
		if v, ok := d["microserviceCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MicroserviceCount)
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

func NewCleanProgressFromJson(data string) CleanProgress {
	req := CleanProgress{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCleanProgressFromDict(data map[string]interface{}) CleanProgress {
	return CleanProgress{
		CleanProgressId: func() *string {
			v, ok := data["cleanProgressId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["cleanProgressId"])
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Cleaned: func() *int32 {
			v, ok := data["cleaned"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["cleaned"])
		}(),
		MicroserviceCount: func() *int32 {
			v, ok := data["microserviceCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["microserviceCount"])
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

func (p CleanProgress) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.CleanProgressId != nil {
		m["cleanProgressId"] = p.CleanProgressId
	}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Cleaned != nil {
		m["cleaned"] = p.Cleaned
	}
	if p.MicroserviceCount != nil {
		m["microserviceCount"] = p.MicroserviceCount
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

func (p CleanProgress) Pointer() *CleanProgress {
	return &p
}

func CastCleanProgresses(data []interface{}) []CleanProgress {
	v := make([]CleanProgress, 0)
	for _, d := range data {
		v = append(v, NewCleanProgressFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCleanProgressesFromDict(data []CleanProgress) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ImportProgress struct {
	ImportProgressId  *string `json:"importProgressId"`
	TransactionId     *string `json:"transactionId"`
	UserId            *string `json:"userId"`
	Imported          *int32  `json:"imported"`
	MicroserviceCount *int32  `json:"microserviceCount"`
	CreatedAt         *int64  `json:"createdAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
	Revision          *int64  `json:"revision"`
}

func (p *ImportProgress) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ImportProgress{}
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
		*p = ImportProgress{}
	} else {
		*p = ImportProgress{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["importProgressId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ImportProgressId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ImportProgressId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ImportProgressId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ImportProgressId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ImportProgressId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ImportProgressId)
				}
			}
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
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
		if v, ok := d["imported"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Imported)
		}
		if v, ok := d["microserviceCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MicroserviceCount)
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

func NewImportProgressFromJson(data string) ImportProgress {
	req := ImportProgress{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewImportProgressFromDict(data map[string]interface{}) ImportProgress {
	return ImportProgress{
		ImportProgressId: func() *string {
			v, ok := data["importProgressId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["importProgressId"])
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Imported: func() *int32 {
			v, ok := data["imported"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["imported"])
		}(),
		MicroserviceCount: func() *int32 {
			v, ok := data["microserviceCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["microserviceCount"])
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

func (p ImportProgress) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ImportProgressId != nil {
		m["importProgressId"] = p.ImportProgressId
	}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Imported != nil {
		m["imported"] = p.Imported
	}
	if p.MicroserviceCount != nil {
		m["microserviceCount"] = p.MicroserviceCount
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

func (p ImportProgress) Pointer() *ImportProgress {
	return &p
}

func CastImportProgresses(data []interface{}) []ImportProgress {
	v := make([]ImportProgress, 0)
	for _, d := range data {
		v = append(v, NewImportProgressFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastImportProgressesFromDict(data []ImportProgress) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ImportErrorLog struct {
	DumpProgressId   *string `json:"dumpProgressId"`
	Name             *string `json:"name"`
	MicroserviceName *string `json:"microserviceName"`
	Message          *string `json:"message"`
	CreatedAt        *int64  `json:"createdAt"`
	Revision         *int64  `json:"revision"`
}

func (p *ImportErrorLog) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ImportErrorLog{}
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
		*p = ImportErrorLog{}
	} else {
		*p = ImportErrorLog{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["dumpProgressId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DumpProgressId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DumpProgressId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DumpProgressId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DumpProgressId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DumpProgressId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DumpProgressId)
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
		if v, ok := d["microserviceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MicroserviceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MicroserviceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MicroserviceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MicroserviceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MicroserviceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MicroserviceName)
				}
			}
		}
		if v, ok := d["message"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Message = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Message = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Message = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Message = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Message = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Message)
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

func NewImportErrorLogFromJson(data string) ImportErrorLog {
	req := ImportErrorLog{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewImportErrorLogFromDict(data map[string]interface{}) ImportErrorLog {
	return ImportErrorLog{
		DumpProgressId: func() *string {
			v, ok := data["dumpProgressId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["dumpProgressId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		MicroserviceName: func() *string {
			v, ok := data["microserviceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["microserviceName"])
		}(),
		Message: func() *string {
			v, ok := data["message"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["message"])
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

func (p ImportErrorLog) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.DumpProgressId != nil {
		m["dumpProgressId"] = p.DumpProgressId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.MicroserviceName != nil {
		m["microserviceName"] = p.MicroserviceName
	}
	if p.Message != nil {
		m["message"] = p.Message
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p ImportErrorLog) Pointer() *ImportErrorLog {
	return &p
}

func CastImportErrorLogs(data []interface{}) []ImportErrorLog {
	v := make([]ImportErrorLog, 0)
	for _, d := range data {
		v = append(v, NewImportErrorLogFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastImportErrorLogsFromDict(data []ImportErrorLog) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
