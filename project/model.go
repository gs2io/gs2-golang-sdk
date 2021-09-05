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
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Account struct {
	AccountId *string `json:"accountId"`
	OwnerId *string `json:"ownerId"`
	Name *string `json:"name"`
	Email *string `json:"email"`
	FullName *string `json:"fullName"`
	CompanyName *string `json:"companyName"`
	Status *string `json:"status"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewAccountFromJson(data string) Account {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAccountFromDict(dict)
}

func NewAccountFromDict(data map[string]interface{}) Account {
    return Account {
        AccountId: core.CastString(data["accountId"]),
        OwnerId: core.CastString(data["ownerId"]),
        Name: core.CastString(data["name"]),
        Email: core.CastString(data["email"]),
        FullName: core.CastString(data["fullName"]),
        CompanyName: core.CastString(data["companyName"]),
        Status: core.CastString(data["status"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Account) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountId": p.AccountId,
        "ownerId": p.OwnerId,
        "name": p.Name,
        "email": p.Email,
        "fullName": p.FullName,
        "companyName": p.CompanyName,
        "status": p.Status,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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
	ProjectId *string `json:"projectId"`
	AccountName *string `json:"accountName"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Plan *string `json:"plan"`
	BillingMethodName *string `json:"billingMethodName"`
	EnableEventBridge *string `json:"enableEventBridge"`
	EventBridgeAwsAccountId *string `json:"eventBridgeAwsAccountId"`
	EventBridgeAwsRegion *string `json:"eventBridgeAwsRegion"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewProjectFromJson(data string) Project {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewProjectFromDict(dict)
}

func NewProjectFromDict(data map[string]interface{}) Project {
    return Project {
        ProjectId: core.CastString(data["projectId"]),
        AccountName: core.CastString(data["accountName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Plan: core.CastString(data["plan"]),
        BillingMethodName: core.CastString(data["billingMethodName"]),
        EnableEventBridge: core.CastString(data["enableEventBridge"]),
        EventBridgeAwsAccountId: core.CastString(data["eventBridgeAwsAccountId"]),
        EventBridgeAwsRegion: core.CastString(data["eventBridgeAwsRegion"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Project) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "projectId": p.ProjectId,
        "accountName": p.AccountName,
        "name": p.Name,
        "description": p.Description,
        "plan": p.Plan,
        "billingMethodName": p.BillingMethodName,
        "enableEventBridge": p.EnableEventBridge,
        "eventBridgeAwsAccountId": p.EventBridgeAwsAccountId,
        "eventBridgeAwsRegion": p.EventBridgeAwsRegion,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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

type BillingMethod struct {
	BillingMethodId *string `json:"billingMethodId"`
	AccountName *string `json:"accountName"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	MethodType *string `json:"methodType"`
	CardSignatureName *string `json:"cardSignatureName"`
	CardBrand *string `json:"cardBrand"`
	CardLast4 *string `json:"cardLast4"`
	PartnerId *string `json:"partnerId"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewBillingMethodFromJson(data string) BillingMethod {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBillingMethodFromDict(dict)
}

func NewBillingMethodFromDict(data map[string]interface{}) BillingMethod {
    return BillingMethod {
        BillingMethodId: core.CastString(data["billingMethodId"]),
        AccountName: core.CastString(data["accountName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        MethodType: core.CastString(data["methodType"]),
        CardSignatureName: core.CastString(data["cardSignatureName"]),
        CardBrand: core.CastString(data["cardBrand"]),
        CardLast4: core.CastString(data["cardLast4"]),
        PartnerId: core.CastString(data["partnerId"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p BillingMethod) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "billingMethodId": p.BillingMethodId,
        "accountName": p.AccountName,
        "name": p.Name,
        "description": p.Description,
        "methodType": p.MethodType,
        "cardSignatureName": p.CardSignatureName,
        "cardBrand": p.CardBrand,
        "cardLast4": p.CardLast4,
        "partnerId": p.PartnerId,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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
	ReceiptId *string `json:"receiptId"`
	AccountName *string `json:"accountName"`
	Name *string `json:"name"`
	Date *int64 `json:"date"`
	Amount *string `json:"amount"`
	PdfUrl *string `json:"pdfUrl"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewReceiptFromJson(data string) Receipt {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewReceiptFromDict(dict)
}

func NewReceiptFromDict(data map[string]interface{}) Receipt {
    return Receipt {
        ReceiptId: core.CastString(data["receiptId"]),
        AccountName: core.CastString(data["accountName"]),
        Name: core.CastString(data["name"]),
        Date: core.CastInt64(data["date"]),
        Amount: core.CastString(data["amount"]),
        PdfUrl: core.CastString(data["pdfUrl"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Receipt) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "receiptId": p.ReceiptId,
        "accountName": p.AccountName,
        "name": p.Name,
        "date": p.Date,
        "amount": p.Amount,
        "pdfUrl": p.PdfUrl,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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
	BillingId *string `json:"billingId"`
	ProjectName *string `json:"projectName"`
	Year *int32 `json:"year"`
	Month *int32 `json:"month"`
	Region *string `json:"region"`
	Service *string `json:"service"`
	ActivityType *string `json:"activityType"`
	Unit *int64 `json:"unit"`
	UnitName *string `json:"unitName"`
	Price *int64 `json:"price"`
	Currency *string `json:"currency"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewBillingFromJson(data string) Billing {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBillingFromDict(dict)
}

func NewBillingFromDict(data map[string]interface{}) Billing {
    return Billing {
        BillingId: core.CastString(data["billingId"]),
        ProjectName: core.CastString(data["projectName"]),
        Year: core.CastInt32(data["year"]),
        Month: core.CastInt32(data["month"]),
        Region: core.CastString(data["region"]),
        Service: core.CastString(data["service"]),
        ActivityType: core.CastString(data["activityType"]),
        Unit: core.CastInt64(data["unit"]),
        UnitName: core.CastString(data["unitName"]),
        Price: core.CastInt64(data["price"]),
        Currency: core.CastString(data["currency"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Billing) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "billingId": p.BillingId,
        "projectName": p.ProjectName,
        "year": p.Year,
        "month": p.Month,
        "region": p.Region,
        "service": p.Service,
        "activityType": p.ActivityType,
        "unit": p.Unit,
        "unitName": p.UnitName,
        "price": p.Price,
        "currency": p.Currency,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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