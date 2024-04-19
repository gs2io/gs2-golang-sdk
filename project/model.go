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
	AccountId   *string `json:"accountId"`
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	FullName    *string `json:"fullName"`
	CompanyName *string `json:"companyName"`
	Status      *string `json:"status"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
}

func NewAccountFromJson(data string) Account {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAccountFromDict(dict)
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

func NewProjectFromJson(data string) Project {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewProjectFromDict(dict)
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

func NewGs2RegionFromJson(data string) Gs2Region {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGs2RegionFromDict(dict)
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

func NewBillingMethodFromJson(data string) BillingMethod {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBillingMethodFromDict(dict)
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

func NewReceiptFromJson(data string) Receipt {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiptFromDict(dict)
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

func NewBillingFromJson(data string) Billing {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBillingFromDict(dict)
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

func NewDumpProgressFromJson(data string) DumpProgress {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpProgressFromDict(dict)
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

func NewCleanProgressFromJson(data string) CleanProgress {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanProgressFromDict(dict)
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

func NewImportProgressFromJson(data string) ImportProgress {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportProgressFromDict(dict)
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

func NewImportErrorLogFromJson(data string) ImportErrorLog {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportErrorLogFromDict(dict)
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
