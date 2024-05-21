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
	AccountId   *string `json:"accountId"`
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	FullName    *string `json:"fullName"`
	CompanyName *string `json:"companyName"`
	Status      *string `json:"status"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
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
		AccountId:   core.CastString(data["accountId"]),
		Name:        core.CastString(data["name"]),
		Email:       core.CastString(data["email"]),
		FullName:    core.CastString(data["fullName"]),
		CompanyName: core.CastString(data["companyName"]),
		Status:      core.CastString(data["status"]),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
	}
}

func (p Account) ToDict() map[string]interface{} {

	var accountId *string
	if p.AccountId != nil {
		accountId = p.AccountId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var email *string
	if p.Email != nil {
		email = p.Email
	}
	var fullName *string
	if p.FullName != nil {
		fullName = p.FullName
	}
	var companyName *string
	if p.CompanyName != nil {
		companyName = p.CompanyName
	}
	var status *string
	if p.Status != nil {
		status = p.Status
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
		"accountId":   accountId,
		"name":        name,
		"email":       email,
		"fullName":    fullName,
		"companyName": companyName,
		"status":      status,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
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
		ProjectId:               core.CastString(data["projectId"]),
		AccountName:             core.CastString(data["accountName"]),
		Name:                    core.CastString(data["name"]),
		Description:             core.CastString(data["description"]),
		Plan:                    core.CastString(data["plan"]),
		Regions:                 CastGs2Regions(core.CastArray(data["regions"])),
		BillingMethodName:       core.CastString(data["billingMethodName"]),
		EnableEventBridge:       core.CastString(data["enableEventBridge"]),
		Currency:                core.CastString(data["currency"]),
		EventBridgeAwsAccountId: core.CastString(data["eventBridgeAwsAccountId"]),
		EventBridgeAwsRegion:    core.CastString(data["eventBridgeAwsRegion"]),
		CreatedAt:               core.CastInt64(data["createdAt"]),
		UpdatedAt:               core.CastInt64(data["updatedAt"]),
	}
}

func (p Project) ToDict() map[string]interface{} {

	var projectId *string
	if p.ProjectId != nil {
		projectId = p.ProjectId
	}
	var accountName *string
	if p.AccountName != nil {
		accountName = p.AccountName
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var plan *string
	if p.Plan != nil {
		plan = p.Plan
	}
	var regions []interface{}
	if p.Regions != nil {
		regions = CastGs2RegionsFromDict(
			p.Regions,
		)
	}
	var billingMethodName *string
	if p.BillingMethodName != nil {
		billingMethodName = p.BillingMethodName
	}
	var enableEventBridge *string
	if p.EnableEventBridge != nil {
		enableEventBridge = p.EnableEventBridge
	}
	var currency *string
	if p.Currency != nil {
		currency = p.Currency
	}
	var eventBridgeAwsAccountId *string
	if p.EventBridgeAwsAccountId != nil {
		eventBridgeAwsAccountId = p.EventBridgeAwsAccountId
	}
	var eventBridgeAwsRegion *string
	if p.EventBridgeAwsRegion != nil {
		eventBridgeAwsRegion = p.EventBridgeAwsRegion
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
		"projectId":               projectId,
		"accountName":             accountName,
		"name":                    name,
		"description":             description,
		"plan":                    plan,
		"regions":                 regions,
		"billingMethodName":       billingMethodName,
		"enableEventBridge":       enableEventBridge,
		"currency":                currency,
		"eventBridgeAwsAccountId": eventBridgeAwsAccountId,
		"eventBridgeAwsRegion":    eventBridgeAwsRegion,
		"createdAt":               createdAt,
		"updatedAt":               updatedAt,
	}
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
		RegionName: core.CastString(data["regionName"]),
		Status:     core.CastString(data["status"]),
	}
}

func (p Gs2Region) ToDict() map[string]interface{} {

	var regionName *string
	if p.RegionName != nil {
		regionName = p.RegionName
	}
	var status *string
	if p.Status != nil {
		status = p.Status
	}
	return map[string]interface{}{
		"regionName": regionName,
		"status":     status,
	}
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
		BillingMethodId:   core.CastString(data["billingMethodId"]),
		AccountName:       core.CastString(data["accountName"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		MethodType:        core.CastString(data["methodType"]),
		CardSignatureName: core.CastString(data["cardSignatureName"]),
		CardBrand:         core.CastString(data["cardBrand"]),
		CardLast4:         core.CastString(data["cardLast4"]),
		PartnerId:         core.CastString(data["partnerId"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
	}
}

func (p BillingMethod) ToDict() map[string]interface{} {

	var billingMethodId *string
	if p.BillingMethodId != nil {
		billingMethodId = p.BillingMethodId
	}
	var accountName *string
	if p.AccountName != nil {
		accountName = p.AccountName
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var methodType *string
	if p.MethodType != nil {
		methodType = p.MethodType
	}
	var cardSignatureName *string
	if p.CardSignatureName != nil {
		cardSignatureName = p.CardSignatureName
	}
	var cardBrand *string
	if p.CardBrand != nil {
		cardBrand = p.CardBrand
	}
	var cardLast4 *string
	if p.CardLast4 != nil {
		cardLast4 = p.CardLast4
	}
	var partnerId *string
	if p.PartnerId != nil {
		partnerId = p.PartnerId
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
		"billingMethodId":   billingMethodId,
		"accountName":       accountName,
		"name":              name,
		"description":       description,
		"methodType":        methodType,
		"cardSignatureName": cardSignatureName,
		"cardBrand":         cardBrand,
		"cardLast4":         cardLast4,
		"partnerId":         partnerId,
		"createdAt":         createdAt,
		"updatedAt":         updatedAt,
	}
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
		ReceiptId:   core.CastString(data["receiptId"]),
		AccountName: core.CastString(data["accountName"]),
		Name:        core.CastString(data["name"]),
		Date:        core.CastInt64(data["date"]),
		Amount:      core.CastString(data["amount"]),
		PdfUrl:      core.CastString(data["pdfUrl"]),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
	}
}

func (p Receipt) ToDict() map[string]interface{} {

	var receiptId *string
	if p.ReceiptId != nil {
		receiptId = p.ReceiptId
	}
	var accountName *string
	if p.AccountName != nil {
		accountName = p.AccountName
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var date *int64
	if p.Date != nil {
		date = p.Date
	}
	var amount *string
	if p.Amount != nil {
		amount = p.Amount
	}
	var pdfUrl *string
	if p.PdfUrl != nil {
		pdfUrl = p.PdfUrl
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
		"receiptId":   receiptId,
		"accountName": accountName,
		"name":        name,
		"date":        date,
		"amount":      amount,
		"pdfUrl":      pdfUrl,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
	}
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
		BillingId:    core.CastString(data["billingId"]),
		ProjectName:  core.CastString(data["projectName"]),
		Year:         core.CastInt32(data["year"]),
		Month:        core.CastInt32(data["month"]),
		Region:       core.CastString(data["region"]),
		Service:      core.CastString(data["service"]),
		ActivityType: core.CastString(data["activityType"]),
		Unit:         core.CastFloat64(data["unit"]),
		UnitName:     core.CastString(data["unitName"]),
		Price:        core.CastFloat64(data["price"]),
		Currency:     core.CastString(data["currency"]),
		CreatedAt:    core.CastInt64(data["createdAt"]),
		UpdatedAt:    core.CastInt64(data["updatedAt"]),
	}
}

func (p Billing) ToDict() map[string]interface{} {

	var billingId *string
	if p.BillingId != nil {
		billingId = p.BillingId
	}
	var projectName *string
	if p.ProjectName != nil {
		projectName = p.ProjectName
	}
	var year *int32
	if p.Year != nil {
		year = p.Year
	}
	var month *int32
	if p.Month != nil {
		month = p.Month
	}
	var region *string
	if p.Region != nil {
		region = p.Region
	}
	var service *string
	if p.Service != nil {
		service = p.Service
	}
	var activityType *string
	if p.ActivityType != nil {
		activityType = p.ActivityType
	}
	var unit *float64
	if p.Unit != nil {
		unit = p.Unit
	}
	var unitName *string
	if p.UnitName != nil {
		unitName = p.UnitName
	}
	var price *float64
	if p.Price != nil {
		price = p.Price
	}
	var currency *string
	if p.Currency != nil {
		currency = p.Currency
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
		"billingId":    billingId,
		"projectName":  projectName,
		"year":         year,
		"month":        month,
		"region":       region,
		"service":      service,
		"activityType": activityType,
		"unit":         unit,
		"unitName":     unitName,
		"price":        price,
		"currency":     currency,
		"createdAt":    createdAt,
		"updatedAt":    updatedAt,
	}
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
		DumpProgressId:    core.CastString(data["dumpProgressId"]),
		TransactionId:     core.CastString(data["transactionId"]),
		UserId:            core.CastString(data["userId"]),
		Dumped:            core.CastInt32(data["dumped"]),
		MicroserviceCount: core.CastInt32(data["microserviceCount"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p DumpProgress) ToDict() map[string]interface{} {

	var dumpProgressId *string
	if p.DumpProgressId != nil {
		dumpProgressId = p.DumpProgressId
	}
	var transactionId *string
	if p.TransactionId != nil {
		transactionId = p.TransactionId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var dumped *int32
	if p.Dumped != nil {
		dumped = p.Dumped
	}
	var microserviceCount *int32
	if p.MicroserviceCount != nil {
		microserviceCount = p.MicroserviceCount
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
		"dumpProgressId":    dumpProgressId,
		"transactionId":     transactionId,
		"userId":            userId,
		"dumped":            dumped,
		"microserviceCount": microserviceCount,
		"createdAt":         createdAt,
		"updatedAt":         updatedAt,
		"revision":          revision,
	}
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
		CleanProgressId:   core.CastString(data["cleanProgressId"]),
		TransactionId:     core.CastString(data["transactionId"]),
		UserId:            core.CastString(data["userId"]),
		Cleaned:           core.CastInt32(data["cleaned"]),
		MicroserviceCount: core.CastInt32(data["microserviceCount"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p CleanProgress) ToDict() map[string]interface{} {

	var cleanProgressId *string
	if p.CleanProgressId != nil {
		cleanProgressId = p.CleanProgressId
	}
	var transactionId *string
	if p.TransactionId != nil {
		transactionId = p.TransactionId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var cleaned *int32
	if p.Cleaned != nil {
		cleaned = p.Cleaned
	}
	var microserviceCount *int32
	if p.MicroserviceCount != nil {
		microserviceCount = p.MicroserviceCount
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
		"cleanProgressId":   cleanProgressId,
		"transactionId":     transactionId,
		"userId":            userId,
		"cleaned":           cleaned,
		"microserviceCount": microserviceCount,
		"createdAt":         createdAt,
		"updatedAt":         updatedAt,
		"revision":          revision,
	}
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
		ImportProgressId:  core.CastString(data["importProgressId"]),
		TransactionId:     core.CastString(data["transactionId"]),
		UserId:            core.CastString(data["userId"]),
		Imported:          core.CastInt32(data["imported"]),
		MicroserviceCount: core.CastInt32(data["microserviceCount"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p ImportProgress) ToDict() map[string]interface{} {

	var importProgressId *string
	if p.ImportProgressId != nil {
		importProgressId = p.ImportProgressId
	}
	var transactionId *string
	if p.TransactionId != nil {
		transactionId = p.TransactionId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var imported *int32
	if p.Imported != nil {
		imported = p.Imported
	}
	var microserviceCount *int32
	if p.MicroserviceCount != nil {
		microserviceCount = p.MicroserviceCount
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
		"importProgressId":  importProgressId,
		"transactionId":     transactionId,
		"userId":            userId,
		"imported":          imported,
		"microserviceCount": microserviceCount,
		"createdAt":         createdAt,
		"updatedAt":         updatedAt,
		"revision":          revision,
	}
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
		DumpProgressId:   core.CastString(data["dumpProgressId"]),
		Name:             core.CastString(data["name"]),
		MicroserviceName: core.CastString(data["microserviceName"]),
		Message:          core.CastString(data["message"]),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		Revision:         core.CastInt64(data["revision"]),
	}
}

func (p ImportErrorLog) ToDict() map[string]interface{} {

	var dumpProgressId *string
	if p.DumpProgressId != nil {
		dumpProgressId = p.DumpProgressId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var microserviceName *string
	if p.MicroserviceName != nil {
		microserviceName = p.MicroserviceName
	}
	var message *string
	if p.Message != nil {
		message = p.Message
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
		"dumpProgressId":   dumpProgressId,
		"name":             name,
		"microserviceName": microserviceName,
		"message":          message,
		"createdAt":        createdAt,
		"revision":         revision,
	}
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
