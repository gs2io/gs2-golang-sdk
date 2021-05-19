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

type Account struct {
    /** GS2アカウント */
	AccountId *string   `json:"accountId"`
    /** None */
	OwnerId *string   `json:"ownerId"`
    /** GS2アカウントの名前 */
	Name *string   `json:"name"`
    /** メールアドレス */
	Email *string   `json:"email"`
    /** フルネーム */
	FullName *string   `json:"fullName"`
    /** 会社名 */
	CompanyName *string   `json:"companyName"`
    /** パスワード */
	Password *string   `json:"password"`
    /** ステータス */
	Status *string   `json:"status"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Account) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["accountId"] = p.AccountId
    data["ownerId"] = p.OwnerId
    data["name"] = p.Name
    data["email"] = p.Email
    data["fullName"] = p.FullName
    data["companyName"] = p.CompanyName
    data["password"] = p.Password
    data["status"] = p.Status
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Project struct {
    /** プロジェクト */
	ProjectId *string   `json:"projectId"`
    /** GS2アカウントの名前 */
	AccountName *string   `json:"accountName"`
    /** プロジェクト名 */
	Name *string   `json:"name"`
    /** プロジェクトの説明 */
	Description *string   `json:"description"`
    /** 契約プラン */
	Plan *string   `json:"plan"`
    /** 支払い方法名 */
	BillingMethodName *string   `json:"billingMethodName"`
    /** AWS EventBridge の設定 */
	EnableEventBridge *string   `json:"enableEventBridge"`
    /** 通知に使用するAWSアカウントのID */
	EventBridgeAwsAccountId *string   `json:"eventBridgeAwsAccountId"`
    /** 通知に使用するAWSリージョン */
	EventBridgeAwsRegion *string   `json:"eventBridgeAwsRegion"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Project) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["projectId"] = p.ProjectId
    data["accountName"] = p.AccountName
    data["name"] = p.Name
    data["description"] = p.Description
    data["plan"] = p.Plan
    data["billingMethodName"] = p.BillingMethodName
    data["enableEventBridge"] = p.EnableEventBridge
    data["eventBridgeAwsAccountId"] = p.EventBridgeAwsAccountId
    data["eventBridgeAwsRegion"] = p.EventBridgeAwsRegion
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type BillingMethod struct {
    /** 支払い方法 */
	BillingMethodId *string   `json:"billingMethodId"`
    /** GS2アカウントの名前 */
	AccountName *string   `json:"accountName"`
    /** 名前 */
	Name *string   `json:"name"`
    /** 名前 */
	Description *string   `json:"description"`
    /** 支払い方法 */
	MethodType *string   `json:"methodType"`
    /** クレジットカードカスタマーID */
	CardCustomerId *string   `json:"cardCustomerId"`
    /** カード署名 */
	CardSignatureName *string   `json:"cardSignatureName"`
    /** カードブランド */
	CardBrand *string   `json:"cardBrand"`
    /** 末尾4桁 */
	CardLast4 *string   `json:"cardLast4"`
    /** パートナーID */
	PartnerId *string   `json:"partnerId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *BillingMethod) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["billingMethodId"] = p.BillingMethodId
    data["accountName"] = p.AccountName
    data["name"] = p.Name
    data["description"] = p.Description
    data["methodType"] = p.MethodType
    data["cardCustomerId"] = p.CardCustomerId
    data["cardSignatureName"] = p.CardSignatureName
    data["cardBrand"] = p.CardBrand
    data["cardLast4"] = p.CardLast4
    data["partnerId"] = p.PartnerId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Receipt struct {
    /** 領収書 */
	ReceiptId *string   `json:"receiptId"`
    /** GS2アカウントの名前 */
	AccountName *string   `json:"accountName"`
    /** 請求書名 */
	Name *string   `json:"name"`
    /** 請求月 */
	Date *int64   `json:"date"`
    /** 請求金額 */
	Amount *string   `json:"amount"`
    /** PDF URL */
	PdfUrl *string   `json:"pdfUrl"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Receipt) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["receiptId"] = p.ReceiptId
    data["accountName"] = p.AccountName
    data["name"] = p.Name
    data["date"] = p.Date
    data["amount"] = p.Amount
    data["pdfUrl"] = p.PdfUrl
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Billing struct {
    /** 利用状況 */
	BillingId *string   `json:"billingId"`
    /** プロジェクト名 */
	ProjectName *string   `json:"projectName"`
    /** イベントの発生年 */
	Year *int32   `json:"year"`
    /** イベントの発生月 */
	Month *int32   `json:"month"`
    /** リージョン */
	Region *string   `json:"region"`
    /** サービスの種類 */
	Service *string   `json:"service"`
    /** イベントの種類 */
	ActivityType *string   `json:"activityType"`
    /** 回数 */
	Unit *int64   `json:"unit"`
    /** 単位 */
	UnitName *string   `json:"unitName"`
    /** 料金 */
	Price *int64   `json:"price"`
    /** 通貨 */
	Currency *string   `json:"currency"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Billing) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["billingId"] = p.BillingId
    data["projectName"] = p.ProjectName
    data["year"] = p.Year
    data["month"] = p.Month
    data["region"] = p.Region
    data["service"] = p.Service
    data["activityType"] = p.ActivityType
    data["unit"] = p.Unit
    data["unitName"] = p.UnitName
    data["price"] = p.Price
    data["currency"] = p.Currency
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}
