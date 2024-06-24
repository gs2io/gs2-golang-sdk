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

package money2

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId           *string          `json:"namespaceId"`
	Name                  *string          `json:"name"`
	Description           *string          `json:"description"`
	CurrencyUsagePriority *string          `json:"currencyUsagePriority"`
	SharedFreeCurrency    *bool            `json:"sharedFreeCurrency"`
	PlatformSetting       *PlatformSetting `json:"platformSetting"`
	ChangeBalanceScript   *ScriptSetting   `json:"changeBalanceScript"`
	LogSetting            *LogSetting      `json:"logSetting"`
	CreatedAt             *int64           `json:"createdAt"`
	UpdatedAt             *int64           `json:"updatedAt"`
	Revision              *int64           `json:"revision"`
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
		if v, ok := d["currencyUsagePriority"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CurrencyUsagePriority = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CurrencyUsagePriority = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CurrencyUsagePriority = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CurrencyUsagePriority = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CurrencyUsagePriority = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CurrencyUsagePriority)
				}
			}
		}
		if v, ok := d["sharedFreeCurrency"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SharedFreeCurrency)
		}
		if v, ok := d["platformSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.PlatformSetting)
		}
		if v, ok := d["changeBalanceScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeBalanceScript)
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
		NamespaceId:           core.CastString(data["namespaceId"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		CurrencyUsagePriority: core.CastString(data["currencyUsagePriority"]),
		SharedFreeCurrency:    core.CastBool(data["sharedFreeCurrency"]),
		PlatformSetting:       NewPlatformSettingFromDict(core.CastMap(data["platformSetting"])).Pointer(),
		ChangeBalanceScript:   NewScriptSettingFromDict(core.CastMap(data["changeBalanceScript"])).Pointer(),
		LogSetting:            NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
		Revision:              core.CastInt64(data["revision"]),
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
	var currencyUsagePriority *string
	if p.CurrencyUsagePriority != nil {
		currencyUsagePriority = p.CurrencyUsagePriority
	}
	var sharedFreeCurrency *bool
	if p.SharedFreeCurrency != nil {
		sharedFreeCurrency = p.SharedFreeCurrency
	}
	var platformSetting map[string]interface{}
	if p.PlatformSetting != nil {
		platformSetting = p.PlatformSetting.ToDict()
	}
	var changeBalanceScript map[string]interface{}
	if p.ChangeBalanceScript != nil {
		changeBalanceScript = p.ChangeBalanceScript.ToDict()
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
		"namespaceId":           namespaceId,
		"name":                  name,
		"description":           description,
		"currencyUsagePriority": currencyUsagePriority,
		"sharedFreeCurrency":    sharedFreeCurrency,
		"platformSetting":       platformSetting,
		"changeBalanceScript":   changeBalanceScript,
		"logSetting":            logSetting,
		"createdAt":             createdAt,
		"updatedAt":             updatedAt,
		"revision":              revision,
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

type Wallet struct {
	WalletId            *string              `json:"walletId"`
	UserId              *string              `json:"userId"`
	Slot                *int32               `json:"slot"`
	Summary             *WalletSummary       `json:"summary"`
	DepositTransactions []DepositTransaction `json:"depositTransactions"`
	SharedFreeCurrency  *bool                `json:"sharedFreeCurrency"`
	CreatedAt           *int64               `json:"createdAt"`
	UpdatedAt           *int64               `json:"updatedAt"`
	Revision            *int64               `json:"revision"`
}

func (p *Wallet) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Wallet{}
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
		*p = Wallet{}
	} else {
		*p = Wallet{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["walletId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.WalletId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.WalletId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.WalletId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.WalletId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.WalletId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.WalletId)
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
		if v, ok := d["slot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Slot)
		}
		if v, ok := d["summary"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Summary)
		}
		if v, ok := d["depositTransactions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositTransactions)
		}
		if v, ok := d["sharedFreeCurrency"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SharedFreeCurrency)
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

func NewWalletFromJson(data string) Wallet {
	req := Wallet{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewWalletFromDict(data map[string]interface{}) Wallet {
	return Wallet{
		WalletId:            core.CastString(data["walletId"]),
		UserId:              core.CastString(data["userId"]),
		Slot:                core.CastInt32(data["slot"]),
		Summary:             NewWalletSummaryFromDict(core.CastMap(data["summary"])).Pointer(),
		DepositTransactions: CastDepositTransactions(core.CastArray(data["depositTransactions"])),
		SharedFreeCurrency:  core.CastBool(data["sharedFreeCurrency"]),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
		Revision:            core.CastInt64(data["revision"]),
	}
}

func (p Wallet) ToDict() map[string]interface{} {

	var walletId *string
	if p.WalletId != nil {
		walletId = p.WalletId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var slot *int32
	if p.Slot != nil {
		slot = p.Slot
	}
	var summary map[string]interface{}
	if p.Summary != nil {
		summary = p.Summary.ToDict()
	}
	var depositTransactions []interface{}
	if p.DepositTransactions != nil {
		depositTransactions = CastDepositTransactionsFromDict(
			p.DepositTransactions,
		)
	}
	var sharedFreeCurrency *bool
	if p.SharedFreeCurrency != nil {
		sharedFreeCurrency = p.SharedFreeCurrency
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
		"walletId":            walletId,
		"userId":              userId,
		"slot":                slot,
		"summary":             summary,
		"depositTransactions": depositTransactions,
		"sharedFreeCurrency":  sharedFreeCurrency,
		"createdAt":           createdAt,
		"updatedAt":           updatedAt,
		"revision":            revision,
	}
}

func (p Wallet) Pointer() *Wallet {
	return &p
}

func CastWallets(data []interface{}) []Wallet {
	v := make([]Wallet, 0)
	for _, d := range data {
		v = append(v, NewWalletFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastWalletsFromDict(data []Wallet) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Event struct {
	EventId            *string             `json:"eventId"`
	TransactionId      *string             `json:"transactionId"`
	UserId             *string             `json:"userId"`
	EventType          *string             `json:"eventType"`
	VerifyReceiptEvent *VerifyReceiptEvent `json:"verifyReceiptEvent"`
	DepositEvent       *DepositEvent       `json:"depositEvent"`
	WithdrawEvent      *WithdrawEvent      `json:"withdrawEvent"`
	CreatedAt          *int64              `json:"createdAt"`
	Revision           *int64              `json:"revision"`
}

func (p *Event) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Event{}
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
		*p = Event{}
	} else {
		*p = Event{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["eventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventId)
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
		if v, ok := d["eventType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventType)
				}
			}
		}
		if v, ok := d["verifyReceiptEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.VerifyReceiptEvent)
		}
		if v, ok := d["depositEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositEvent)
		}
		if v, ok := d["withdrawEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WithdrawEvent)
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

func NewEventFromJson(data string) Event {
	req := Event{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewEventFromDict(data map[string]interface{}) Event {
	return Event{
		EventId:            core.CastString(data["eventId"]),
		TransactionId:      core.CastString(data["transactionId"]),
		UserId:             core.CastString(data["userId"]),
		EventType:          core.CastString(data["eventType"]),
		VerifyReceiptEvent: NewVerifyReceiptEventFromDict(core.CastMap(data["verifyReceiptEvent"])).Pointer(),
		DepositEvent:       NewDepositEventFromDict(core.CastMap(data["depositEvent"])).Pointer(),
		WithdrawEvent:      NewWithdrawEventFromDict(core.CastMap(data["withdrawEvent"])).Pointer(),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		Revision:           core.CastInt64(data["revision"]),
	}
}

func (p Event) ToDict() map[string]interface{} {

	var eventId *string
	if p.EventId != nil {
		eventId = p.EventId
	}
	var transactionId *string
	if p.TransactionId != nil {
		transactionId = p.TransactionId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var eventType *string
	if p.EventType != nil {
		eventType = p.EventType
	}
	var verifyReceiptEvent map[string]interface{}
	if p.VerifyReceiptEvent != nil {
		verifyReceiptEvent = p.VerifyReceiptEvent.ToDict()
	}
	var depositEvent map[string]interface{}
	if p.DepositEvent != nil {
		depositEvent = p.DepositEvent.ToDict()
	}
	var withdrawEvent map[string]interface{}
	if p.WithdrawEvent != nil {
		withdrawEvent = p.WithdrawEvent.ToDict()
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
		"eventId":            eventId,
		"transactionId":      transactionId,
		"userId":             userId,
		"eventType":          eventType,
		"verifyReceiptEvent": verifyReceiptEvent,
		"depositEvent":       depositEvent,
		"withdrawEvent":      withdrawEvent,
		"createdAt":          createdAt,
		"revision":           revision,
	}
}

func (p Event) Pointer() *Event {
	return &p
}

func CastEvents(data []interface{}) []Event {
	v := make([]Event, 0)
	for _, d := range data {
		v = append(v, NewEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEventsFromDict(data []Event) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StoreContentModel struct {
	StoreContentModelId *string               `json:"storeContentModelId"`
	Name                *string               `json:"name"`
	Metadata            *string               `json:"metadata"`
	AppleAppStore       *AppleAppStoreContent `json:"appleAppStore"`
	GooglePlay          *GooglePlayContent    `json:"googlePlay"`
}

func (p *StoreContentModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = StoreContentModel{}
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
		*p = StoreContentModel{}
	} else {
		*p = StoreContentModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["storeContentModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StoreContentModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StoreContentModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StoreContentModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StoreContentModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StoreContentModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StoreContentModelId)
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
		if v, ok := d["appleAppStore"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppleAppStore)
		}
		if v, ok := d["googlePlay"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GooglePlay)
		}
	}
	return nil
}

func NewStoreContentModelFromJson(data string) StoreContentModel {
	req := StoreContentModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStoreContentModelFromDict(data map[string]interface{}) StoreContentModel {
	return StoreContentModel{
		StoreContentModelId: core.CastString(data["storeContentModelId"]),
		Name:                core.CastString(data["name"]),
		Metadata:            core.CastString(data["metadata"]),
		AppleAppStore:       NewAppleAppStoreContentFromDict(core.CastMap(data["appleAppStore"])).Pointer(),
		GooglePlay:          NewGooglePlayContentFromDict(core.CastMap(data["googlePlay"])).Pointer(),
	}
}

func (p StoreContentModel) ToDict() map[string]interface{} {

	var storeContentModelId *string
	if p.StoreContentModelId != nil {
		storeContentModelId = p.StoreContentModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var appleAppStore map[string]interface{}
	if p.AppleAppStore != nil {
		appleAppStore = p.AppleAppStore.ToDict()
	}
	var googlePlay map[string]interface{}
	if p.GooglePlay != nil {
		googlePlay = p.GooglePlay.ToDict()
	}
	return map[string]interface{}{
		"storeContentModelId": storeContentModelId,
		"name":                name,
		"metadata":            metadata,
		"appleAppStore":       appleAppStore,
		"googlePlay":          googlePlay,
	}
}

func (p StoreContentModel) Pointer() *StoreContentModel {
	return &p
}

func CastStoreContentModels(data []interface{}) []StoreContentModel {
	v := make([]StoreContentModel, 0)
	for _, d := range data {
		v = append(v, NewStoreContentModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStoreContentModelsFromDict(data []StoreContentModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StoreContentModelMaster struct {
	StoreContentModelId *string               `json:"storeContentModelId"`
	Name                *string               `json:"name"`
	Description         *string               `json:"description"`
	Metadata            *string               `json:"metadata"`
	AppleAppStore       *AppleAppStoreContent `json:"appleAppStore"`
	GooglePlay          *GooglePlayContent    `json:"googlePlay"`
	CreatedAt           *int64                `json:"createdAt"`
	UpdatedAt           *int64                `json:"updatedAt"`
	Revision            *int64                `json:"revision"`
}

func (p *StoreContentModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = StoreContentModelMaster{}
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
		*p = StoreContentModelMaster{}
	} else {
		*p = StoreContentModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["storeContentModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StoreContentModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StoreContentModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StoreContentModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StoreContentModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StoreContentModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StoreContentModelId)
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
		if v, ok := d["appleAppStore"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppleAppStore)
		}
		if v, ok := d["googlePlay"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GooglePlay)
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

func NewStoreContentModelMasterFromJson(data string) StoreContentModelMaster {
	req := StoreContentModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStoreContentModelMasterFromDict(data map[string]interface{}) StoreContentModelMaster {
	return StoreContentModelMaster{
		StoreContentModelId: core.CastString(data["storeContentModelId"]),
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		Metadata:            core.CastString(data["metadata"]),
		AppleAppStore:       NewAppleAppStoreContentFromDict(core.CastMap(data["appleAppStore"])).Pointer(),
		GooglePlay:          NewGooglePlayContentFromDict(core.CastMap(data["googlePlay"])).Pointer(),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
		Revision:            core.CastInt64(data["revision"]),
	}
}

func (p StoreContentModelMaster) ToDict() map[string]interface{} {

	var storeContentModelId *string
	if p.StoreContentModelId != nil {
		storeContentModelId = p.StoreContentModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var appleAppStore map[string]interface{}
	if p.AppleAppStore != nil {
		appleAppStore = p.AppleAppStore.ToDict()
	}
	var googlePlay map[string]interface{}
	if p.GooglePlay != nil {
		googlePlay = p.GooglePlay.ToDict()
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
		"storeContentModelId": storeContentModelId,
		"name":                name,
		"description":         description,
		"metadata":            metadata,
		"appleAppStore":       appleAppStore,
		"googlePlay":          googlePlay,
		"createdAt":           createdAt,
		"updatedAt":           updatedAt,
		"revision":            revision,
	}
}

func (p StoreContentModelMaster) Pointer() *StoreContentModelMaster {
	return &p
}

func CastStoreContentModelMasters(data []interface{}) []StoreContentModelMaster {
	v := make([]StoreContentModelMaster, 0)
	for _, d := range data {
		v = append(v, NewStoreContentModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStoreContentModelMastersFromDict(data []StoreContentModelMaster) []interface{} {
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
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentModelMaster) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var settings *string
	if p.Settings != nil {
		settings = p.Settings
	}
	return map[string]interface{}{
		"namespaceId": namespaceId,
		"settings":    settings,
	}
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

type Receipt struct {
	Store         *string `json:"Store"`
	TransactionID *string `json:"TransactionID"`
	Payload       *string `json:"Payload"`
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
		if v, ok := d["Store"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Store = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Store = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Store = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Store = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Store = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Store)
				}
			}
		}
		if v, ok := d["TransactionID"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionID = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionID = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionID = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionID = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionID = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionID)
				}
			}
		}
		if v, ok := d["Payload"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Payload = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Payload = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Payload = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Payload = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Payload = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Payload)
				}
			}
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
		Store:         core.CastString(data["store"]),
		TransactionID: core.CastString(data["transactionID"]),
		Payload:       core.CastString(data["payload"]),
	}
}

func (p Receipt) ToDict() map[string]interface{} {

	var store *string
	if p.Store != nil {
		store = p.Store
	}
	var transactionID *string
	if p.TransactionID != nil {
		transactionID = p.TransactionID
	}
	var payload *string
	if p.Payload != nil {
		payload = p.Payload
	}
	return map[string]interface{}{
		"Store":         store,
		"TransactionID": transactionID,
		"Payload":       payload,
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

type PlatformSetting struct {
	AppleAppStore *AppleAppStoreSetting `json:"appleAppStore"`
	GooglePlay    *GooglePlaySetting    `json:"googlePlay"`
	Fake          *FakeSetting          `json:"fake"`
}

func (p *PlatformSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PlatformSetting{}
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
		*p = PlatformSetting{}
	} else {
		*p = PlatformSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["appleAppStore"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppleAppStore)
		}
		if v, ok := d["googlePlay"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GooglePlay)
		}
		if v, ok := d["fake"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Fake)
		}
	}
	return nil
}

func NewPlatformSettingFromJson(data string) PlatformSetting {
	req := PlatformSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPlatformSettingFromDict(data map[string]interface{}) PlatformSetting {
	return PlatformSetting{
		AppleAppStore: NewAppleAppStoreSettingFromDict(core.CastMap(data["appleAppStore"])).Pointer(),
		GooglePlay:    NewGooglePlaySettingFromDict(core.CastMap(data["googlePlay"])).Pointer(),
		Fake:          NewFakeSettingFromDict(core.CastMap(data["fake"])).Pointer(),
	}
}

func (p PlatformSetting) ToDict() map[string]interface{} {

	var appleAppStore map[string]interface{}
	if p.AppleAppStore != nil {
		appleAppStore = p.AppleAppStore.ToDict()
	}
	var googlePlay map[string]interface{}
	if p.GooglePlay != nil {
		googlePlay = p.GooglePlay.ToDict()
	}
	var fake map[string]interface{}
	if p.Fake != nil {
		fake = p.Fake.ToDict()
	}
	return map[string]interface{}{
		"appleAppStore": appleAppStore,
		"googlePlay":    googlePlay,
		"fake":          fake,
	}
}

func (p PlatformSetting) Pointer() *PlatformSetting {
	return &p
}

func CastPlatformSettings(data []interface{}) []PlatformSetting {
	v := make([]PlatformSetting, 0)
	for _, d := range data {
		v = append(v, NewPlatformSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPlatformSettingsFromDict(data []PlatformSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AppleAppStoreSetting struct {
	BundleId *string `json:"bundleId"`
}

func (p *AppleAppStoreSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AppleAppStoreSetting{}
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
		*p = AppleAppStoreSetting{}
	} else {
		*p = AppleAppStoreSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["bundleId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BundleId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BundleId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BundleId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BundleId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BundleId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BundleId)
				}
			}
		}
	}
	return nil
}

func NewAppleAppStoreSettingFromJson(data string) AppleAppStoreSetting {
	req := AppleAppStoreSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAppleAppStoreSettingFromDict(data map[string]interface{}) AppleAppStoreSetting {
	return AppleAppStoreSetting{
		BundleId: core.CastString(data["bundleId"]),
	}
}

func (p AppleAppStoreSetting) ToDict() map[string]interface{} {

	var bundleId *string
	if p.BundleId != nil {
		bundleId = p.BundleId
	}
	return map[string]interface{}{
		"bundleId": bundleId,
	}
}

func (p AppleAppStoreSetting) Pointer() *AppleAppStoreSetting {
	return &p
}

func CastAppleAppStoreSettings(data []interface{}) []AppleAppStoreSetting {
	v := make([]AppleAppStoreSetting, 0)
	for _, d := range data {
		v = append(v, NewAppleAppStoreSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAppleAppStoreSettingsFromDict(data []AppleAppStoreSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GooglePlaySetting struct {
	PackageName *string `json:"packageName"`
	PublicKey   *string `json:"publicKey"`
}

func (p *GooglePlaySetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GooglePlaySetting{}
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
		*p = GooglePlaySetting{}
	} else {
		*p = GooglePlaySetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["packageName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PackageName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PackageName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PackageName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PackageName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PackageName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PackageName)
				}
			}
		}
		if v, ok := d["publicKey"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PublicKey = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PublicKey = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PublicKey = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PublicKey = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PublicKey = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PublicKey)
				}
			}
		}
	}
	return nil
}

func NewGooglePlaySettingFromJson(data string) GooglePlaySetting {
	req := GooglePlaySetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGooglePlaySettingFromDict(data map[string]interface{}) GooglePlaySetting {
	return GooglePlaySetting{
		PackageName: core.CastString(data["packageName"]),
		PublicKey:   core.CastString(data["publicKey"]),
	}
}

func (p GooglePlaySetting) ToDict() map[string]interface{} {

	var packageName *string
	if p.PackageName != nil {
		packageName = p.PackageName
	}
	var publicKey *string
	if p.PublicKey != nil {
		publicKey = p.PublicKey
	}
	return map[string]interface{}{
		"packageName": packageName,
		"publicKey":   publicKey,
	}
}

func (p GooglePlaySetting) Pointer() *GooglePlaySetting {
	return &p
}

func CastGooglePlaySettings(data []interface{}) []GooglePlaySetting {
	v := make([]GooglePlaySetting, 0)
	for _, d := range data {
		v = append(v, NewGooglePlaySettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGooglePlaySettingsFromDict(data []GooglePlaySetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FakeSetting struct {
	AcceptFakeReceipt *string `json:"acceptFakeReceipt"`
}

func (p *FakeSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FakeSetting{}
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
		*p = FakeSetting{}
	} else {
		*p = FakeSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["acceptFakeReceipt"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AcceptFakeReceipt = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AcceptFakeReceipt = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AcceptFakeReceipt = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AcceptFakeReceipt = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AcceptFakeReceipt = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AcceptFakeReceipt)
				}
			}
		}
	}
	return nil
}

func NewFakeSettingFromJson(data string) FakeSetting {
	req := FakeSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFakeSettingFromDict(data map[string]interface{}) FakeSetting {
	return FakeSetting{
		AcceptFakeReceipt: core.CastString(data["acceptFakeReceipt"]),
	}
}

func (p FakeSetting) ToDict() map[string]interface{} {

	var acceptFakeReceipt *string
	if p.AcceptFakeReceipt != nil {
		acceptFakeReceipt = p.AcceptFakeReceipt
	}
	return map[string]interface{}{
		"acceptFakeReceipt": acceptFakeReceipt,
	}
}

func (p FakeSetting) Pointer() *FakeSetting {
	return &p
}

func CastFakeSettings(data []interface{}) []FakeSetting {
	v := make([]FakeSetting, 0)
	for _, d := range data {
		v = append(v, NewFakeSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFakeSettingsFromDict(data []FakeSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type WalletSummary struct {
	Paid  *int32 `json:"paid"`
	Free  *int32 `json:"free"`
	Total *int32 `json:"total"`
}

func (p *WalletSummary) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WalletSummary{}
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
		*p = WalletSummary{}
	} else {
		*p = WalletSummary{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["paid"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Paid)
		}
		if v, ok := d["free"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Free)
		}
		if v, ok := d["total"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Total)
		}
	}
	return nil
}

func NewWalletSummaryFromJson(data string) WalletSummary {
	req := WalletSummary{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewWalletSummaryFromDict(data map[string]interface{}) WalletSummary {
	return WalletSummary{
		Paid:  core.CastInt32(data["paid"]),
		Free:  core.CastInt32(data["free"]),
		Total: core.CastInt32(data["total"]),
	}
}

func (p WalletSummary) ToDict() map[string]interface{} {

	var paid *int32
	if p.Paid != nil {
		paid = p.Paid
	}
	var free *int32
	if p.Free != nil {
		free = p.Free
	}
	var total *int32
	if p.Total != nil {
		total = p.Total
	}
	return map[string]interface{}{
		"paid":  paid,
		"free":  free,
		"total": total,
	}
}

func (p WalletSummary) Pointer() *WalletSummary {
	return &p
}

func CastWalletSummaries(data []interface{}) []WalletSummary {
	v := make([]WalletSummary, 0)
	for _, d := range data {
		v = append(v, NewWalletSummaryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastWalletSummariesFromDict(data []WalletSummary) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DepositTransaction struct {
	Price       *float32 `json:"price"`
	Currency    *string  `json:"currency"`
	Count       *int32   `json:"count"`
	DepositedAt *int64   `json:"depositedAt"`
}

func (p *DepositTransaction) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DepositTransaction{}
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
		*p = DepositTransaction{}
	} else {
		*p = DepositTransaction{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
		if v, ok := d["depositedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositedAt)
		}
	}
	return nil
}

func NewDepositTransactionFromJson(data string) DepositTransaction {
	req := DepositTransaction{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDepositTransactionFromDict(data map[string]interface{}) DepositTransaction {
	return DepositTransaction{
		Price:       core.CastFloat32(data["price"]),
		Currency:    core.CastString(data["currency"]),
		Count:       core.CastInt32(data["count"]),
		DepositedAt: core.CastInt64(data["depositedAt"]),
	}
}

func (p DepositTransaction) ToDict() map[string]interface{} {

	var price *float32
	if p.Price != nil {
		price = p.Price
	}
	var currency *string
	if p.Currency != nil {
		currency = p.Currency
	}
	var count *int32
	if p.Count != nil {
		count = p.Count
	}
	var depositedAt *int64
	if p.DepositedAt != nil {
		depositedAt = p.DepositedAt
	}
	return map[string]interface{}{
		"price":       price,
		"currency":    currency,
		"count":       count,
		"depositedAt": depositedAt,
	}
}

func (p DepositTransaction) Pointer() *DepositTransaction {
	return &p
}

func CastDepositTransactions(data []interface{}) []DepositTransaction {
	v := make([]DepositTransaction, 0)
	for _, d := range data {
		v = append(v, NewDepositTransactionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDepositTransactionsFromDict(data []DepositTransaction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type VerifyReceiptEvent struct {
	ContentName                     *string                          `json:"contentName"`
	Platform                        *string                          `json:"platform"`
	AppleAppStoreVerifyReceiptEvent *AppleAppStoreVerifyReceiptEvent `json:"appleAppStoreVerifyReceiptEvent"`
	GooglePlayVerifyReceiptEvent    *GooglePlayVerifyReceiptEvent    `json:"googlePlayVerifyReceiptEvent"`
}

func (p *VerifyReceiptEvent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyReceiptEvent{}
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
		*p = VerifyReceiptEvent{}
	} else {
		*p = VerifyReceiptEvent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
		if v, ok := d["platform"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Platform = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Platform = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Platform = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Platform = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Platform = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Platform)
				}
			}
		}
		if v, ok := d["appleAppStoreVerifyReceiptEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppleAppStoreVerifyReceiptEvent)
		}
		if v, ok := d["googlePlayVerifyReceiptEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GooglePlayVerifyReceiptEvent)
		}
	}
	return nil
}

func NewVerifyReceiptEventFromJson(data string) VerifyReceiptEvent {
	req := VerifyReceiptEvent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVerifyReceiptEventFromDict(data map[string]interface{}) VerifyReceiptEvent {
	return VerifyReceiptEvent{
		ContentName:                     core.CastString(data["contentName"]),
		Platform:                        core.CastString(data["platform"]),
		AppleAppStoreVerifyReceiptEvent: NewAppleAppStoreVerifyReceiptEventFromDict(core.CastMap(data["appleAppStoreVerifyReceiptEvent"])).Pointer(),
		GooglePlayVerifyReceiptEvent:    NewGooglePlayVerifyReceiptEventFromDict(core.CastMap(data["googlePlayVerifyReceiptEvent"])).Pointer(),
	}
}

func (p VerifyReceiptEvent) ToDict() map[string]interface{} {

	var contentName *string
	if p.ContentName != nil {
		contentName = p.ContentName
	}
	var platform *string
	if p.Platform != nil {
		platform = p.Platform
	}
	var appleAppStoreVerifyReceiptEvent map[string]interface{}
	if p.AppleAppStoreVerifyReceiptEvent != nil {
		appleAppStoreVerifyReceiptEvent = p.AppleAppStoreVerifyReceiptEvent.ToDict()
	}
	var googlePlayVerifyReceiptEvent map[string]interface{}
	if p.GooglePlayVerifyReceiptEvent != nil {
		googlePlayVerifyReceiptEvent = p.GooglePlayVerifyReceiptEvent.ToDict()
	}
	return map[string]interface{}{
		"contentName":                     contentName,
		"platform":                        platform,
		"appleAppStoreVerifyReceiptEvent": appleAppStoreVerifyReceiptEvent,
		"googlePlayVerifyReceiptEvent":    googlePlayVerifyReceiptEvent,
	}
}

func (p VerifyReceiptEvent) Pointer() *VerifyReceiptEvent {
	return &p
}

func CastVerifyReceiptEvents(data []interface{}) []VerifyReceiptEvent {
	v := make([]VerifyReceiptEvent, 0)
	for _, d := range data {
		v = append(v, NewVerifyReceiptEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVerifyReceiptEventsFromDict(data []VerifyReceiptEvent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DepositEvent struct {
	Slot                *int32               `json:"slot"`
	DepositTransactions []DepositTransaction `json:"depositTransactions"`
	Status              *WalletSummary       `json:"status"`
}

func (p *DepositEvent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DepositEvent{}
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
		*p = DepositEvent{}
	} else {
		*p = DepositEvent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["slot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Slot)
		}
		if v, ok := d["depositTransactions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositTransactions)
		}
		if v, ok := d["status"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Status)
		}
	}
	return nil
}

func NewDepositEventFromJson(data string) DepositEvent {
	req := DepositEvent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDepositEventFromDict(data map[string]interface{}) DepositEvent {
	return DepositEvent{
		Slot:                core.CastInt32(data["slot"]),
		DepositTransactions: CastDepositTransactions(core.CastArray(data["depositTransactions"])),
		Status:              NewWalletSummaryFromDict(core.CastMap(data["status"])).Pointer(),
	}
}

func (p DepositEvent) ToDict() map[string]interface{} {

	var slot *int32
	if p.Slot != nil {
		slot = p.Slot
	}
	var depositTransactions []interface{}
	if p.DepositTransactions != nil {
		depositTransactions = CastDepositTransactionsFromDict(
			p.DepositTransactions,
		)
	}
	var status map[string]interface{}
	if p.Status != nil {
		status = p.Status.ToDict()
	}
	return map[string]interface{}{
		"slot":                slot,
		"depositTransactions": depositTransactions,
		"status":              status,
	}
}

func (p DepositEvent) Pointer() *DepositEvent {
	return &p
}

func CastDepositEvents(data []interface{}) []DepositEvent {
	v := make([]DepositEvent, 0)
	for _, d := range data {
		v = append(v, NewDepositEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDepositEventsFromDict(data []DepositEvent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type WithdrawEvent struct {
	Slot            *int32               `json:"slot"`
	WithdrawDetails []DepositTransaction `json:"withdrawDetails"`
	Status          *WalletSummary       `json:"status"`
}

func (p *WithdrawEvent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WithdrawEvent{}
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
		*p = WithdrawEvent{}
	} else {
		*p = WithdrawEvent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["slot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Slot)
		}
		if v, ok := d["withdrawDetails"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WithdrawDetails)
		}
		if v, ok := d["status"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Status)
		}
	}
	return nil
}

func NewWithdrawEventFromJson(data string) WithdrawEvent {
	req := WithdrawEvent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewWithdrawEventFromDict(data map[string]interface{}) WithdrawEvent {
	return WithdrawEvent{
		Slot:            core.CastInt32(data["slot"]),
		WithdrawDetails: CastDepositTransactions(core.CastArray(data["withdrawDetails"])),
		Status:          NewWalletSummaryFromDict(core.CastMap(data["status"])).Pointer(),
	}
}

func (p WithdrawEvent) ToDict() map[string]interface{} {

	var slot *int32
	if p.Slot != nil {
		slot = p.Slot
	}
	var withdrawDetails []interface{}
	if p.WithdrawDetails != nil {
		withdrawDetails = CastDepositTransactionsFromDict(
			p.WithdrawDetails,
		)
	}
	var status map[string]interface{}
	if p.Status != nil {
		status = p.Status.ToDict()
	}
	return map[string]interface{}{
		"slot":            slot,
		"withdrawDetails": withdrawDetails,
		"status":          status,
	}
}

func (p WithdrawEvent) Pointer() *WithdrawEvent {
	return &p
}

func CastWithdrawEvents(data []interface{}) []WithdrawEvent {
	v := make([]WithdrawEvent, 0)
	for _, d := range data {
		v = append(v, NewWithdrawEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastWithdrawEventsFromDict(data []WithdrawEvent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AppleAppStoreVerifyReceiptEvent struct {
	Environment *string `json:"environment"`
}

func (p *AppleAppStoreVerifyReceiptEvent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AppleAppStoreVerifyReceiptEvent{}
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
		*p = AppleAppStoreVerifyReceiptEvent{}
	} else {
		*p = AppleAppStoreVerifyReceiptEvent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["environment"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Environment = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Environment = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Environment = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Environment = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Environment = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Environment)
				}
			}
		}
	}
	return nil
}

func NewAppleAppStoreVerifyReceiptEventFromJson(data string) AppleAppStoreVerifyReceiptEvent {
	req := AppleAppStoreVerifyReceiptEvent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAppleAppStoreVerifyReceiptEventFromDict(data map[string]interface{}) AppleAppStoreVerifyReceiptEvent {
	return AppleAppStoreVerifyReceiptEvent{
		Environment: core.CastString(data["environment"]),
	}
}

func (p AppleAppStoreVerifyReceiptEvent) ToDict() map[string]interface{} {

	var environment *string
	if p.Environment != nil {
		environment = p.Environment
	}
	return map[string]interface{}{
		"environment": environment,
	}
}

func (p AppleAppStoreVerifyReceiptEvent) Pointer() *AppleAppStoreVerifyReceiptEvent {
	return &p
}

func CastAppleAppStoreVerifyReceiptEvents(data []interface{}) []AppleAppStoreVerifyReceiptEvent {
	v := make([]AppleAppStoreVerifyReceiptEvent, 0)
	for _, d := range data {
		v = append(v, NewAppleAppStoreVerifyReceiptEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAppleAppStoreVerifyReceiptEventsFromDict(data []AppleAppStoreVerifyReceiptEvent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GooglePlayVerifyReceiptEvent struct {
	PurchaseToken *string `json:"purchaseToken"`
}

func (p *GooglePlayVerifyReceiptEvent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GooglePlayVerifyReceiptEvent{}
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
		*p = GooglePlayVerifyReceiptEvent{}
	} else {
		*p = GooglePlayVerifyReceiptEvent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["purchaseToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PurchaseToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PurchaseToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PurchaseToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PurchaseToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PurchaseToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PurchaseToken)
				}
			}
		}
	}
	return nil
}

func NewGooglePlayVerifyReceiptEventFromJson(data string) GooglePlayVerifyReceiptEvent {
	req := GooglePlayVerifyReceiptEvent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGooglePlayVerifyReceiptEventFromDict(data map[string]interface{}) GooglePlayVerifyReceiptEvent {
	return GooglePlayVerifyReceiptEvent{
		PurchaseToken: core.CastString(data["purchaseToken"]),
	}
}

func (p GooglePlayVerifyReceiptEvent) ToDict() map[string]interface{} {

	var purchaseToken *string
	if p.PurchaseToken != nil {
		purchaseToken = p.PurchaseToken
	}
	return map[string]interface{}{
		"purchaseToken": purchaseToken,
	}
}

func (p GooglePlayVerifyReceiptEvent) Pointer() *GooglePlayVerifyReceiptEvent {
	return &p
}

func CastGooglePlayVerifyReceiptEvents(data []interface{}) []GooglePlayVerifyReceiptEvent {
	v := make([]GooglePlayVerifyReceiptEvent, 0)
	for _, d := range data {
		v = append(v, NewGooglePlayVerifyReceiptEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGooglePlayVerifyReceiptEventsFromDict(data []GooglePlayVerifyReceiptEvent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AppleAppStoreContent struct {
	ProductId *string `json:"productId"`
}

func (p *AppleAppStoreContent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AppleAppStoreContent{}
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
		*p = AppleAppStoreContent{}
	} else {
		*p = AppleAppStoreContent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["productId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProductId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProductId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProductId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProductId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProductId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProductId)
				}
			}
		}
	}
	return nil
}

func NewAppleAppStoreContentFromJson(data string) AppleAppStoreContent {
	req := AppleAppStoreContent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAppleAppStoreContentFromDict(data map[string]interface{}) AppleAppStoreContent {
	return AppleAppStoreContent{
		ProductId: core.CastString(data["productId"]),
	}
}

func (p AppleAppStoreContent) ToDict() map[string]interface{} {

	var productId *string
	if p.ProductId != nil {
		productId = p.ProductId
	}
	return map[string]interface{}{
		"productId": productId,
	}
}

func (p AppleAppStoreContent) Pointer() *AppleAppStoreContent {
	return &p
}

func CastAppleAppStoreContents(data []interface{}) []AppleAppStoreContent {
	v := make([]AppleAppStoreContent, 0)
	for _, d := range data {
		v = append(v, NewAppleAppStoreContentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAppleAppStoreContentsFromDict(data []AppleAppStoreContent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GooglePlayContent struct {
	ProductId *string `json:"productId"`
}

func (p *GooglePlayContent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GooglePlayContent{}
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
		*p = GooglePlayContent{}
	} else {
		*p = GooglePlayContent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["productId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProductId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProductId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProductId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProductId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProductId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProductId)
				}
			}
		}
	}
	return nil
}

func NewGooglePlayContentFromJson(data string) GooglePlayContent {
	req := GooglePlayContent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGooglePlayContentFromDict(data map[string]interface{}) GooglePlayContent {
	return GooglePlayContent{
		ProductId: core.CastString(data["productId"]),
	}
}

func (p GooglePlayContent) ToDict() map[string]interface{} {

	var productId *string
	if p.ProductId != nil {
		productId = p.ProductId
	}
	return map[string]interface{}{
		"productId": productId,
	}
}

func (p GooglePlayContent) Pointer() *GooglePlayContent {
	return &p
}

func CastGooglePlayContents(data []interface{}) []GooglePlayContent {
	v := make([]GooglePlayContent, 0)
	for _, d := range data {
		v = append(v, NewGooglePlayContentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGooglePlayContentsFromDict(data []GooglePlayContent) []interface{} {
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
		ApiKeyId:       core.CastString(data["apiKeyId"]),
		RepositoryName: core.CastString(data["repositoryName"]),
		SourcePath:     core.CastString(data["sourcePath"]),
		ReferenceType:  core.CastString(data["referenceType"]),
		CommitHash:     core.CastString(data["commitHash"]),
		BranchName:     core.CastString(data["branchName"]),
		TagName:        core.CastString(data["tagName"]),
	}
}

func (p GitHubCheckoutSetting) ToDict() map[string]interface{} {

	var apiKeyId *string
	if p.ApiKeyId != nil {
		apiKeyId = p.ApiKeyId
	}
	var repositoryName *string
	if p.RepositoryName != nil {
		repositoryName = p.RepositoryName
	}
	var sourcePath *string
	if p.SourcePath != nil {
		sourcePath = p.SourcePath
	}
	var referenceType *string
	if p.ReferenceType != nil {
		referenceType = p.ReferenceType
	}
	var commitHash *string
	if p.CommitHash != nil {
		commitHash = p.CommitHash
	}
	var branchName *string
	if p.BranchName != nil {
		branchName = p.BranchName
	}
	var tagName *string
	if p.TagName != nil {
		tagName = p.TagName
	}
	return map[string]interface{}{
		"apiKeyId":       apiKeyId,
		"repositoryName": repositoryName,
		"sourcePath":     sourcePath,
		"referenceType":  referenceType,
		"commitHash":     commitHash,
		"branchName":     branchName,
		"tagName":        tagName,
	}
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
