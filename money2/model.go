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
	DepositBalanceScript  *ScriptSetting   `json:"depositBalanceScript"`
	WithdrawBalanceScript *ScriptSetting   `json:"withdrawBalanceScript"`
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
		if v, ok := d["depositBalanceScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositBalanceScript)
		}
		if v, ok := d["withdrawBalanceScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WithdrawBalanceScript)
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
		CurrencyUsagePriority: func() *string {
			v, ok := data["currencyUsagePriority"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currencyUsagePriority"])
		}(),
		SharedFreeCurrency: func() *bool {
			v, ok := data["sharedFreeCurrency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sharedFreeCurrency"])
		}(),
		PlatformSetting: func() *PlatformSetting {
			v, ok := data["platformSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformSettingFromDict(core.CastMap(data["platformSetting"])).Pointer()
		}(),
		DepositBalanceScript: func() *ScriptSetting {
			v, ok := data["depositBalanceScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["depositBalanceScript"])).Pointer()
		}(),
		WithdrawBalanceScript: func() *ScriptSetting {
			v, ok := data["withdrawBalanceScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["withdrawBalanceScript"])).Pointer()
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
	if p.CurrencyUsagePriority != nil {
		m["currencyUsagePriority"] = p.CurrencyUsagePriority
	}
	if p.SharedFreeCurrency != nil {
		m["sharedFreeCurrency"] = p.SharedFreeCurrency
	}
	if p.PlatformSetting != nil {
		m["platformSetting"] = func() map[string]interface{} {
			if p.PlatformSetting == nil {
				return nil
			}
			return p.PlatformSetting.ToDict()
		}()
	}
	if p.DepositBalanceScript != nil {
		m["depositBalanceScript"] = func() map[string]interface{} {
			if p.DepositBalanceScript == nil {
				return nil
			}
			return p.DepositBalanceScript.ToDict()
		}()
	}
	if p.WithdrawBalanceScript != nil {
		m["withdrawBalanceScript"] = func() map[string]interface{} {
			if p.WithdrawBalanceScript == nil {
				return nil
			}
			return p.WithdrawBalanceScript.ToDict()
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
		WalletId: func() *string {
			v, ok := data["walletId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["walletId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Slot: func() *int32 {
			v, ok := data["slot"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["slot"])
		}(),
		Summary: func() *WalletSummary {
			v, ok := data["summary"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletSummaryFromDict(core.CastMap(data["summary"])).Pointer()
		}(),
		DepositTransactions: func() []DepositTransaction {
			if data["depositTransactions"] == nil {
				return nil
			}
			return CastDepositTransactions(core.CastArray(data["depositTransactions"]))
		}(),
		SharedFreeCurrency: func() *bool {
			v, ok := data["sharedFreeCurrency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sharedFreeCurrency"])
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

func (p Wallet) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.WalletId != nil {
		m["walletId"] = p.WalletId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Slot != nil {
		m["slot"] = p.Slot
	}
	if p.Summary != nil {
		m["summary"] = func() map[string]interface{} {
			if p.Summary == nil {
				return nil
			}
			return p.Summary.ToDict()
		}()
	}
	if p.DepositTransactions != nil {
		m["depositTransactions"] = CastDepositTransactionsFromDict(
			p.DepositTransactions,
		)
	}
	if p.SharedFreeCurrency != nil {
		m["sharedFreeCurrency"] = p.SharedFreeCurrency
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
	RefundEvent        *RefundEvent        `json:"refundEvent"`
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
		if v, ok := d["refundEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RefundEvent)
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
		EventId: func() *string {
			v, ok := data["eventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventId"])
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
		EventType: func() *string {
			v, ok := data["eventType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventType"])
		}(),
		VerifyReceiptEvent: func() *VerifyReceiptEvent {
			v, ok := data["verifyReceiptEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewVerifyReceiptEventFromDict(core.CastMap(data["verifyReceiptEvent"])).Pointer()
		}(),
		DepositEvent: func() *DepositEvent {
			v, ok := data["depositEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewDepositEventFromDict(core.CastMap(data["depositEvent"])).Pointer()
		}(),
		WithdrawEvent: func() *WithdrawEvent {
			v, ok := data["withdrawEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewWithdrawEventFromDict(core.CastMap(data["withdrawEvent"])).Pointer()
		}(),
		RefundEvent: func() *RefundEvent {
			v, ok := data["refundEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewRefundEventFromDict(core.CastMap(data["refundEvent"])).Pointer()
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

func (p Event) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.EventId != nil {
		m["eventId"] = p.EventId
	}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.EventType != nil {
		m["eventType"] = p.EventType
	}
	if p.VerifyReceiptEvent != nil {
		m["verifyReceiptEvent"] = func() map[string]interface{} {
			if p.VerifyReceiptEvent == nil {
				return nil
			}
			return p.VerifyReceiptEvent.ToDict()
		}()
	}
	if p.DepositEvent != nil {
		m["depositEvent"] = func() map[string]interface{} {
			if p.DepositEvent == nil {
				return nil
			}
			return p.DepositEvent.ToDict()
		}()
	}
	if p.WithdrawEvent != nil {
		m["withdrawEvent"] = func() map[string]interface{} {
			if p.WithdrawEvent == nil {
				return nil
			}
			return p.WithdrawEvent.ToDict()
		}()
	}
	if p.RefundEvent != nil {
		m["refundEvent"] = func() map[string]interface{} {
			if p.RefundEvent == nil {
				return nil
			}
			return p.RefundEvent.ToDict()
		}()
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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

type SubscribeTransaction struct {
	SubscribeTransactionId *string `json:"subscribeTransactionId"`
	ContentName            *string `json:"contentName"`
	TransactionId          *string `json:"transactionId"`
	Store                  *string `json:"store"`
	UserId                 *string `json:"userId"`
	StatusDetail           *string `json:"statusDetail"`
	ExpiresAt              *int64  `json:"expiresAt"`
	CreatedAt              *int64  `json:"createdAt"`
	UpdatedAt              *int64  `json:"updatedAt"`
	Revision               *int64  `json:"revision"`
}

func (p *SubscribeTransaction) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubscribeTransaction{}
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
		*p = SubscribeTransaction{}
	} else {
		*p = SubscribeTransaction{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["subscribeTransactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscribeTransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscribeTransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscribeTransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeTransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeTransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscribeTransactionId)
				}
			}
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
		if v, ok := d["store"]; ok && v != nil {
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
		if v, ok := d["statusDetail"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StatusDetail = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StatusDetail = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StatusDetail = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StatusDetail = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StatusDetail = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StatusDetail)
				}
			}
		}
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
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

func NewSubscribeTransactionFromJson(data string) SubscribeTransaction {
	req := SubscribeTransaction{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSubscribeTransactionFromDict(data map[string]interface{}) SubscribeTransaction {
	return SubscribeTransaction{
		SubscribeTransactionId: func() *string {
			v, ok := data["subscribeTransactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["subscribeTransactionId"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		Store: func() *string {
			v, ok := data["store"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["store"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		StatusDetail: func() *string {
			v, ok := data["statusDetail"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["statusDetail"])
		}(),
		ExpiresAt: func() *int64 {
			v, ok := data["expiresAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expiresAt"])
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

func (p SubscribeTransaction) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SubscribeTransactionId != nil {
		m["subscribeTransactionId"] = p.SubscribeTransactionId
	}
	if p.ContentName != nil {
		m["contentName"] = p.ContentName
	}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.Store != nil {
		m["store"] = p.Store
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.StatusDetail != nil {
		m["statusDetail"] = p.StatusDetail
	}
	if p.ExpiresAt != nil {
		m["expiresAt"] = p.ExpiresAt
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

func (p SubscribeTransaction) Pointer() *SubscribeTransaction {
	return &p
}

func CastSubscribeTransactions(data []interface{}) []SubscribeTransaction {
	v := make([]SubscribeTransaction, 0)
	for _, d := range data {
		v = append(v, NewSubscribeTransactionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSubscribeTransactionsFromDict(data []SubscribeTransaction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SubscriptionStatus struct {
	UserId      *string                `json:"userId"`
	ContentName *string                `json:"contentName"`
	Status      *string                `json:"status"`
	ExpiresAt   *int64                 `json:"expiresAt"`
	Detail      []SubscribeTransaction `json:"detail"`
}

func (p *SubscriptionStatus) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubscriptionStatus{}
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
		*p = SubscriptionStatus{}
	} else {
		*p = SubscriptionStatus{}
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
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
		}
		if v, ok := d["detail"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Detail)
		}
	}
	return nil
}

func NewSubscriptionStatusFromJson(data string) SubscriptionStatus {
	req := SubscriptionStatus{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSubscriptionStatusFromDict(data map[string]interface{}) SubscriptionStatus {
	return SubscriptionStatus{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
		ExpiresAt: func() *int64 {
			v, ok := data["expiresAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expiresAt"])
		}(),
		Detail: func() []SubscribeTransaction {
			if data["detail"] == nil {
				return nil
			}
			return CastSubscribeTransactions(core.CastArray(data["detail"]))
		}(),
	}
}

func (p SubscriptionStatus) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.ContentName != nil {
		m["contentName"] = p.ContentName
	}
	if p.Status != nil {
		m["status"] = p.Status
	}
	if p.ExpiresAt != nil {
		m["expiresAt"] = p.ExpiresAt
	}
	if p.Detail != nil {
		m["detail"] = CastSubscribeTransactionsFromDict(
			p.Detail,
		)
	}
	return m
}

func (p SubscriptionStatus) Pointer() *SubscriptionStatus {
	return &p
}

func CastSubscriptionStatuses(data []interface{}) []SubscriptionStatus {
	v := make([]SubscriptionStatus, 0)
	for _, d := range data {
		v = append(v, NewSubscriptionStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSubscriptionStatusesFromDict(data []SubscriptionStatus) []interface{} {
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
		StoreContentModelId: func() *string {
			v, ok := data["storeContentModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["storeContentModelId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		AppleAppStore: func() *AppleAppStoreContent {
			v, ok := data["appleAppStore"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreContentFromDict(core.CastMap(data["appleAppStore"])).Pointer()
		}(),
		GooglePlay: func() *GooglePlayContent {
			v, ok := data["googlePlay"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlayContentFromDict(core.CastMap(data["googlePlay"])).Pointer()
		}(),
	}
}

func (p StoreContentModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StoreContentModelId != nil {
		m["storeContentModelId"] = p.StoreContentModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.AppleAppStore != nil {
		m["appleAppStore"] = func() map[string]interface{} {
			if p.AppleAppStore == nil {
				return nil
			}
			return p.AppleAppStore.ToDict()
		}()
	}
	if p.GooglePlay != nil {
		m["googlePlay"] = func() map[string]interface{} {
			if p.GooglePlay == nil {
				return nil
			}
			return p.GooglePlay.ToDict()
		}()
	}
	return m
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
		StoreContentModelId: func() *string {
			v, ok := data["storeContentModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["storeContentModelId"])
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
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		AppleAppStore: func() *AppleAppStoreContent {
			v, ok := data["appleAppStore"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreContentFromDict(core.CastMap(data["appleAppStore"])).Pointer()
		}(),
		GooglePlay: func() *GooglePlayContent {
			v, ok := data["googlePlay"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlayContentFromDict(core.CastMap(data["googlePlay"])).Pointer()
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

func (p StoreContentModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StoreContentModelId != nil {
		m["storeContentModelId"] = p.StoreContentModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.AppleAppStore != nil {
		m["appleAppStore"] = func() map[string]interface{} {
			if p.AppleAppStore == nil {
				return nil
			}
			return p.AppleAppStore.ToDict()
		}()
	}
	if p.GooglePlay != nil {
		m["googlePlay"] = func() map[string]interface{} {
			if p.GooglePlay == nil {
				return nil
			}
			return p.GooglePlay.ToDict()
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

type StoreSubscriptionContentModel struct {
	StoreSubscriptionContentModelId *string                           `json:"storeSubscriptionContentModelId"`
	Name                            *string                           `json:"name"`
	Metadata                        *string                           `json:"metadata"`
	ScheduleNamespaceId             *string                           `json:"scheduleNamespaceId"`
	TriggerName                     *string                           `json:"triggerName"`
	AppleAppStore                   *AppleAppStoreSubscriptionContent `json:"appleAppStore"`
	GooglePlay                      *GooglePlaySubscriptionContent    `json:"googlePlay"`
}

func (p *StoreSubscriptionContentModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = StoreSubscriptionContentModel{}
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
		*p = StoreSubscriptionContentModel{}
	} else {
		*p = StoreSubscriptionContentModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["storeSubscriptionContentModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StoreSubscriptionContentModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StoreSubscriptionContentModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StoreSubscriptionContentModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StoreSubscriptionContentModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StoreSubscriptionContentModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StoreSubscriptionContentModelId)
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
		if v, ok := d["scheduleNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScheduleNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScheduleNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScheduleNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScheduleNamespaceId)
				}
			}
		}
		if v, ok := d["triggerName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerName)
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

func NewStoreSubscriptionContentModelFromJson(data string) StoreSubscriptionContentModel {
	req := StoreSubscriptionContentModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStoreSubscriptionContentModelFromDict(data map[string]interface{}) StoreSubscriptionContentModel {
	return StoreSubscriptionContentModel{
		StoreSubscriptionContentModelId: func() *string {
			v, ok := data["storeSubscriptionContentModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["storeSubscriptionContentModelId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		ScheduleNamespaceId: func() *string {
			v, ok := data["scheduleNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scheduleNamespaceId"])
		}(),
		TriggerName: func() *string {
			v, ok := data["triggerName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerName"])
		}(),
		AppleAppStore: func() *AppleAppStoreSubscriptionContent {
			v, ok := data["appleAppStore"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreSubscriptionContentFromDict(core.CastMap(data["appleAppStore"])).Pointer()
		}(),
		GooglePlay: func() *GooglePlaySubscriptionContent {
			v, ok := data["googlePlay"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlaySubscriptionContentFromDict(core.CastMap(data["googlePlay"])).Pointer()
		}(),
	}
}

func (p StoreSubscriptionContentModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StoreSubscriptionContentModelId != nil {
		m["storeSubscriptionContentModelId"] = p.StoreSubscriptionContentModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ScheduleNamespaceId != nil {
		m["scheduleNamespaceId"] = p.ScheduleNamespaceId
	}
	if p.TriggerName != nil {
		m["triggerName"] = p.TriggerName
	}
	if p.AppleAppStore != nil {
		m["appleAppStore"] = func() map[string]interface{} {
			if p.AppleAppStore == nil {
				return nil
			}
			return p.AppleAppStore.ToDict()
		}()
	}
	if p.GooglePlay != nil {
		m["googlePlay"] = func() map[string]interface{} {
			if p.GooglePlay == nil {
				return nil
			}
			return p.GooglePlay.ToDict()
		}()
	}
	return m
}

func (p StoreSubscriptionContentModel) Pointer() *StoreSubscriptionContentModel {
	return &p
}

func CastStoreSubscriptionContentModels(data []interface{}) []StoreSubscriptionContentModel {
	v := make([]StoreSubscriptionContentModel, 0)
	for _, d := range data {
		v = append(v, NewStoreSubscriptionContentModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStoreSubscriptionContentModelsFromDict(data []StoreSubscriptionContentModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StoreSubscriptionContentModelMaster struct {
	StoreSubscriptionContentModelId *string                           `json:"storeSubscriptionContentModelId"`
	Name                            *string                           `json:"name"`
	Description                     *string                           `json:"description"`
	Metadata                        *string                           `json:"metadata"`
	ScheduleNamespaceId             *string                           `json:"scheduleNamespaceId"`
	TriggerName                     *string                           `json:"triggerName"`
	AppleAppStore                   *AppleAppStoreSubscriptionContent `json:"appleAppStore"`
	GooglePlay                      *GooglePlaySubscriptionContent    `json:"googlePlay"`
	CreatedAt                       *int64                            `json:"createdAt"`
	UpdatedAt                       *int64                            `json:"updatedAt"`
	Revision                        *int64                            `json:"revision"`
}

func (p *StoreSubscriptionContentModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = StoreSubscriptionContentModelMaster{}
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
		*p = StoreSubscriptionContentModelMaster{}
	} else {
		*p = StoreSubscriptionContentModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["storeSubscriptionContentModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StoreSubscriptionContentModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StoreSubscriptionContentModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StoreSubscriptionContentModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StoreSubscriptionContentModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StoreSubscriptionContentModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StoreSubscriptionContentModelId)
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
		if v, ok := d["scheduleNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScheduleNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScheduleNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScheduleNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScheduleNamespaceId)
				}
			}
		}
		if v, ok := d["triggerName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerName)
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

func NewStoreSubscriptionContentModelMasterFromJson(data string) StoreSubscriptionContentModelMaster {
	req := StoreSubscriptionContentModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStoreSubscriptionContentModelMasterFromDict(data map[string]interface{}) StoreSubscriptionContentModelMaster {
	return StoreSubscriptionContentModelMaster{
		StoreSubscriptionContentModelId: func() *string {
			v, ok := data["storeSubscriptionContentModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["storeSubscriptionContentModelId"])
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
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		ScheduleNamespaceId: func() *string {
			v, ok := data["scheduleNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scheduleNamespaceId"])
		}(),
		TriggerName: func() *string {
			v, ok := data["triggerName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerName"])
		}(),
		AppleAppStore: func() *AppleAppStoreSubscriptionContent {
			v, ok := data["appleAppStore"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreSubscriptionContentFromDict(core.CastMap(data["appleAppStore"])).Pointer()
		}(),
		GooglePlay: func() *GooglePlaySubscriptionContent {
			v, ok := data["googlePlay"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlaySubscriptionContentFromDict(core.CastMap(data["googlePlay"])).Pointer()
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

func (p StoreSubscriptionContentModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StoreSubscriptionContentModelId != nil {
		m["storeSubscriptionContentModelId"] = p.StoreSubscriptionContentModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ScheduleNamespaceId != nil {
		m["scheduleNamespaceId"] = p.ScheduleNamespaceId
	}
	if p.TriggerName != nil {
		m["triggerName"] = p.TriggerName
	}
	if p.AppleAppStore != nil {
		m["appleAppStore"] = func() map[string]interface{} {
			if p.AppleAppStore == nil {
				return nil
			}
			return p.AppleAppStore.ToDict()
		}()
	}
	if p.GooglePlay != nil {
		m["googlePlay"] = func() map[string]interface{} {
			if p.GooglePlay == nil {
				return nil
			}
			return p.GooglePlay.ToDict()
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

func (p StoreSubscriptionContentModelMaster) Pointer() *StoreSubscriptionContentModelMaster {
	return &p
}

func CastStoreSubscriptionContentModelMasters(data []interface{}) []StoreSubscriptionContentModelMaster {
	v := make([]StoreSubscriptionContentModelMaster, 0)
	for _, d := range data {
		v = append(v, NewStoreSubscriptionContentModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStoreSubscriptionContentModelMastersFromDict(data []StoreSubscriptionContentModelMaster) []interface{} {
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
		Store: func() *string {
			v, ok := data["store"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["store"])
		}(),
		TransactionID: func() *string {
			v, ok := data["transactionID"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionID"])
		}(),
		Payload: func() *string {
			v, ok := data["payload"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["payload"])
		}(),
	}
}

func (p Receipt) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Store != nil {
		m["Store"] = p.Store
	}
	if p.TransactionID != nil {
		m["TransactionID"] = p.TransactionID
	}
	if p.Payload != nil {
		m["Payload"] = p.Payload
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
		AppleAppStore: func() *AppleAppStoreSetting {
			v, ok := data["appleAppStore"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreSettingFromDict(core.CastMap(data["appleAppStore"])).Pointer()
		}(),
		GooglePlay: func() *GooglePlaySetting {
			v, ok := data["googlePlay"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlaySettingFromDict(core.CastMap(data["googlePlay"])).Pointer()
		}(),
		Fake: func() *FakeSetting {
			v, ok := data["fake"]
			if !ok || v == nil {
				return nil
			}
			return NewFakeSettingFromDict(core.CastMap(data["fake"])).Pointer()
		}(),
	}
}

func (p PlatformSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.AppleAppStore != nil {
		m["appleAppStore"] = func() map[string]interface{} {
			if p.AppleAppStore == nil {
				return nil
			}
			return p.AppleAppStore.ToDict()
		}()
	}
	if p.GooglePlay != nil {
		m["googlePlay"] = func() map[string]interface{} {
			if p.GooglePlay == nil {
				return nil
			}
			return p.GooglePlay.ToDict()
		}()
	}
	if p.Fake != nil {
		m["fake"] = func() map[string]interface{} {
			if p.Fake == nil {
				return nil
			}
			return p.Fake.ToDict()
		}()
	}
	return m
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
	BundleId      *string `json:"bundleId"`
	IssuerId      *string `json:"issuerId"`
	KeyId         *string `json:"keyId"`
	PrivateKeyPem *string `json:"privateKeyPem"`
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
		if v, ok := d["issuerId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.IssuerId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.IssuerId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.IssuerId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.IssuerId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.IssuerId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.IssuerId)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.KeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.KeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.KeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
		if v, ok := d["privateKeyPem"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrivateKeyPem = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrivateKeyPem = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrivateKeyPem = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrivateKeyPem = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrivateKeyPem = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrivateKeyPem)
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
		BundleId: func() *string {
			v, ok := data["bundleId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["bundleId"])
		}(),
		IssuerId: func() *string {
			v, ok := data["issuerId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["issuerId"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
		PrivateKeyPem: func() *string {
			v, ok := data["privateKeyPem"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["privateKeyPem"])
		}(),
	}
}

func (p AppleAppStoreSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.BundleId != nil {
		m["bundleId"] = p.BundleId
	}
	if p.IssuerId != nil {
		m["issuerId"] = p.IssuerId
	}
	if p.KeyId != nil {
		m["keyId"] = p.KeyId
	}
	if p.PrivateKeyPem != nil {
		m["privateKeyPem"] = p.PrivateKeyPem
	}
	return m
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
	PackageName     *string `json:"packageName"`
	PublicKey       *string `json:"publicKey"`
	CredentialsJSON *string `json:"credentialsJSON"`
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
		if v, ok := d["credentialsJSON"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CredentialsJSON = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CredentialsJSON = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CredentialsJSON = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CredentialsJSON = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CredentialsJSON = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CredentialsJSON)
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
		PackageName: func() *string {
			v, ok := data["packageName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["packageName"])
		}(),
		PublicKey: func() *string {
			v, ok := data["publicKey"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["publicKey"])
		}(),
		CredentialsJSON: func() *string {
			v, ok := data["credentialsJSON"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["credentialsJSON"])
		}(),
	}
}

func (p GooglePlaySetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PackageName != nil {
		m["packageName"] = p.PackageName
	}
	if p.PublicKey != nil {
		m["publicKey"] = p.PublicKey
	}
	if p.CredentialsJSON != nil {
		m["credentialsJSON"] = p.CredentialsJSON
	}
	return m
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
		AcceptFakeReceipt: func() *string {
			v, ok := data["acceptFakeReceipt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["acceptFakeReceipt"])
		}(),
	}
}

func (p FakeSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.AcceptFakeReceipt != nil {
		m["acceptFakeReceipt"] = p.AcceptFakeReceipt
	}
	return m
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
		Paid: func() *int32 {
			v, ok := data["paid"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["paid"])
		}(),
		Free: func() *int32 {
			v, ok := data["free"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["free"])
		}(),
		Total: func() *int32 {
			v, ok := data["total"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["total"])
		}(),
	}
}

func (p WalletSummary) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Paid != nil {
		m["paid"] = p.Paid
	}
	if p.Free != nil {
		m["free"] = p.Free
	}
	if p.Total != nil {
		m["total"] = p.Total
	}
	return m
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
		Price: func() *float32 {
			v, ok := data["price"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["price"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
		}(),
		Count: func() *int32 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["count"])
		}(),
		DepositedAt: func() *int64 {
			v, ok := data["depositedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["depositedAt"])
		}(),
	}
}

func (p DepositTransaction) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Price != nil {
		m["price"] = p.Price
	}
	if p.Currency != nil {
		m["currency"] = p.Currency
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	if p.DepositedAt != nil {
		m["depositedAt"] = p.DepositedAt
	}
	return m
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
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
		Platform: func() *string {
			v, ok := data["platform"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["platform"])
		}(),
		AppleAppStoreVerifyReceiptEvent: func() *AppleAppStoreVerifyReceiptEvent {
			v, ok := data["appleAppStoreVerifyReceiptEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreVerifyReceiptEventFromDict(core.CastMap(data["appleAppStoreVerifyReceiptEvent"])).Pointer()
		}(),
		GooglePlayVerifyReceiptEvent: func() *GooglePlayVerifyReceiptEvent {
			v, ok := data["googlePlayVerifyReceiptEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlayVerifyReceiptEventFromDict(core.CastMap(data["googlePlayVerifyReceiptEvent"])).Pointer()
		}(),
	}
}

func (p VerifyReceiptEvent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ContentName != nil {
		m["contentName"] = p.ContentName
	}
	if p.Platform != nil {
		m["platform"] = p.Platform
	}
	if p.AppleAppStoreVerifyReceiptEvent != nil {
		m["appleAppStoreVerifyReceiptEvent"] = func() map[string]interface{} {
			if p.AppleAppStoreVerifyReceiptEvent == nil {
				return nil
			}
			return p.AppleAppStoreVerifyReceiptEvent.ToDict()
		}()
	}
	if p.GooglePlayVerifyReceiptEvent != nil {
		m["googlePlayVerifyReceiptEvent"] = func() map[string]interface{} {
			if p.GooglePlayVerifyReceiptEvent == nil {
				return nil
			}
			return p.GooglePlayVerifyReceiptEvent.ToDict()
		}()
	}
	return m
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
		Slot: func() *int32 {
			v, ok := data["slot"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["slot"])
		}(),
		DepositTransactions: func() []DepositTransaction {
			if data["depositTransactions"] == nil {
				return nil
			}
			return CastDepositTransactions(core.CastArray(data["depositTransactions"]))
		}(),
		Status: func() *WalletSummary {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletSummaryFromDict(core.CastMap(data["status"])).Pointer()
		}(),
	}
}

func (p DepositEvent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Slot != nil {
		m["slot"] = p.Slot
	}
	if p.DepositTransactions != nil {
		m["depositTransactions"] = CastDepositTransactionsFromDict(
			p.DepositTransactions,
		)
	}
	if p.Status != nil {
		m["status"] = func() map[string]interface{} {
			if p.Status == nil {
				return nil
			}
			return p.Status.ToDict()
		}()
	}
	return m
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
		Slot: func() *int32 {
			v, ok := data["slot"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["slot"])
		}(),
		WithdrawDetails: func() []DepositTransaction {
			if data["withdrawDetails"] == nil {
				return nil
			}
			return CastDepositTransactions(core.CastArray(data["withdrawDetails"]))
		}(),
		Status: func() *WalletSummary {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletSummaryFromDict(core.CastMap(data["status"])).Pointer()
		}(),
	}
}

func (p WithdrawEvent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Slot != nil {
		m["slot"] = p.Slot
	}
	if p.WithdrawDetails != nil {
		m["withdrawDetails"] = CastDepositTransactionsFromDict(
			p.WithdrawDetails,
		)
	}
	if p.Status != nil {
		m["status"] = func() map[string]interface{} {
			if p.Status == nil {
				return nil
			}
			return p.Status.ToDict()
		}()
	}
	return m
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

type RefundEvent struct {
	ContentName              *string                          `json:"contentName"`
	Platform                 *string                          `json:"platform"`
	AppleAppStoreRefundEvent *AppleAppStoreVerifyReceiptEvent `json:"appleAppStoreRefundEvent"`
	GooglePlayRefundEvent    *GooglePlayVerifyReceiptEvent    `json:"googlePlayRefundEvent"`
}

func (p *RefundEvent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RefundEvent{}
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
		*p = RefundEvent{}
	} else {
		*p = RefundEvent{}
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
		if v, ok := d["appleAppStoreRefundEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppleAppStoreRefundEvent)
		}
		if v, ok := d["googlePlayRefundEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GooglePlayRefundEvent)
		}
	}
	return nil
}

func NewRefundEventFromJson(data string) RefundEvent {
	req := RefundEvent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRefundEventFromDict(data map[string]interface{}) RefundEvent {
	return RefundEvent{
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
		Platform: func() *string {
			v, ok := data["platform"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["platform"])
		}(),
		AppleAppStoreRefundEvent: func() *AppleAppStoreVerifyReceiptEvent {
			v, ok := data["appleAppStoreRefundEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreVerifyReceiptEventFromDict(core.CastMap(data["appleAppStoreRefundEvent"])).Pointer()
		}(),
		GooglePlayRefundEvent: func() *GooglePlayVerifyReceiptEvent {
			v, ok := data["googlePlayRefundEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlayVerifyReceiptEventFromDict(core.CastMap(data["googlePlayRefundEvent"])).Pointer()
		}(),
	}
}

func (p RefundEvent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ContentName != nil {
		m["contentName"] = p.ContentName
	}
	if p.Platform != nil {
		m["platform"] = p.Platform
	}
	if p.AppleAppStoreRefundEvent != nil {
		m["appleAppStoreRefundEvent"] = func() map[string]interface{} {
			if p.AppleAppStoreRefundEvent == nil {
				return nil
			}
			return p.AppleAppStoreRefundEvent.ToDict()
		}()
	}
	if p.GooglePlayRefundEvent != nil {
		m["googlePlayRefundEvent"] = func() map[string]interface{} {
			if p.GooglePlayRefundEvent == nil {
				return nil
			}
			return p.GooglePlayRefundEvent.ToDict()
		}()
	}
	return m
}

func (p RefundEvent) Pointer() *RefundEvent {
	return &p
}

func CastRefundEvents(data []interface{}) []RefundEvent {
	v := make([]RefundEvent, 0)
	for _, d := range data {
		v = append(v, NewRefundEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRefundEventsFromDict(data []RefundEvent) []interface{} {
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
		Environment: func() *string {
			v, ok := data["environment"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["environment"])
		}(),
	}
}

func (p AppleAppStoreVerifyReceiptEvent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Environment != nil {
		m["environment"] = p.Environment
	}
	return m
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
		PurchaseToken: func() *string {
			v, ok := data["purchaseToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["purchaseToken"])
		}(),
	}
}

func (p GooglePlayVerifyReceiptEvent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PurchaseToken != nil {
		m["purchaseToken"] = p.PurchaseToken
	}
	return m
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
		ProductId: func() *string {
			v, ok := data["productId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["productId"])
		}(),
	}
}

func (p AppleAppStoreContent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ProductId != nil {
		m["productId"] = p.ProductId
	}
	return m
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
		ProductId: func() *string {
			v, ok := data["productId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["productId"])
		}(),
	}
}

func (p GooglePlayContent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ProductId != nil {
		m["productId"] = p.ProductId
	}
	return m
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

type AppleAppStoreSubscriptionContent struct {
	SubscriptionGroupIdentifier *string `json:"subscriptionGroupIdentifier"`
}

func (p *AppleAppStoreSubscriptionContent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AppleAppStoreSubscriptionContent{}
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
		*p = AppleAppStoreSubscriptionContent{}
	} else {
		*p = AppleAppStoreSubscriptionContent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["subscriptionGroupIdentifier"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscriptionGroupIdentifier = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscriptionGroupIdentifier = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscriptionGroupIdentifier = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscriptionGroupIdentifier = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscriptionGroupIdentifier = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscriptionGroupIdentifier)
				}
			}
		}
	}
	return nil
}

func NewAppleAppStoreSubscriptionContentFromJson(data string) AppleAppStoreSubscriptionContent {
	req := AppleAppStoreSubscriptionContent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAppleAppStoreSubscriptionContentFromDict(data map[string]interface{}) AppleAppStoreSubscriptionContent {
	return AppleAppStoreSubscriptionContent{
		SubscriptionGroupIdentifier: func() *string {
			v, ok := data["subscriptionGroupIdentifier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["subscriptionGroupIdentifier"])
		}(),
	}
}

func (p AppleAppStoreSubscriptionContent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SubscriptionGroupIdentifier != nil {
		m["subscriptionGroupIdentifier"] = p.SubscriptionGroupIdentifier
	}
	return m
}

func (p AppleAppStoreSubscriptionContent) Pointer() *AppleAppStoreSubscriptionContent {
	return &p
}

func CastAppleAppStoreSubscriptionContents(data []interface{}) []AppleAppStoreSubscriptionContent {
	v := make([]AppleAppStoreSubscriptionContent, 0)
	for _, d := range data {
		v = append(v, NewAppleAppStoreSubscriptionContentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAppleAppStoreSubscriptionContentsFromDict(data []AppleAppStoreSubscriptionContent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GooglePlaySubscriptionContent struct {
	ProductId *string `json:"productId"`
}

func (p *GooglePlaySubscriptionContent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GooglePlaySubscriptionContent{}
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
		*p = GooglePlaySubscriptionContent{}
	} else {
		*p = GooglePlaySubscriptionContent{}
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

func NewGooglePlaySubscriptionContentFromJson(data string) GooglePlaySubscriptionContent {
	req := GooglePlaySubscriptionContent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGooglePlaySubscriptionContentFromDict(data map[string]interface{}) GooglePlaySubscriptionContent {
	return GooglePlaySubscriptionContent{
		ProductId: func() *string {
			v, ok := data["productId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["productId"])
		}(),
	}
}

func (p GooglePlaySubscriptionContent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ProductId != nil {
		m["productId"] = p.ProductId
	}
	return m
}

func (p GooglePlaySubscriptionContent) Pointer() *GooglePlaySubscriptionContent {
	return &p
}

func CastGooglePlaySubscriptionContents(data []interface{}) []GooglePlaySubscriptionContent {
	v := make([]GooglePlaySubscriptionContent, 0)
	for _, d := range data {
		v = append(v, NewGooglePlaySubscriptionContentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGooglePlaySubscriptionContentsFromDict(data []GooglePlaySubscriptionContent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GooglePlayRealtimeNotificationMessage struct {
	Data        *string `json:"data"`
	MessageId   *string `json:"messageId"`
	PublishTime *string `json:"publishTime"`
}

func (p *GooglePlayRealtimeNotificationMessage) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GooglePlayRealtimeNotificationMessage{}
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
		*p = GooglePlayRealtimeNotificationMessage{}
	} else {
		*p = GooglePlayRealtimeNotificationMessage{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["data"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Data = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Data = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Data = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Data = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Data = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Data)
				}
			}
		}
		if v, ok := d["messageId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MessageId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MessageId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MessageId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MessageId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MessageId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MessageId)
				}
			}
		}
		if v, ok := d["publishTime"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PublishTime = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PublishTime = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PublishTime = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PublishTime = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PublishTime = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PublishTime)
				}
			}
		}
	}
	return nil
}

func NewGooglePlayRealtimeNotificationMessageFromJson(data string) GooglePlayRealtimeNotificationMessage {
	req := GooglePlayRealtimeNotificationMessage{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGooglePlayRealtimeNotificationMessageFromDict(data map[string]interface{}) GooglePlayRealtimeNotificationMessage {
	return GooglePlayRealtimeNotificationMessage{
		Data: func() *string {
			v, ok := data["data"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["data"])
		}(),
		MessageId: func() *string {
			v, ok := data["messageId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["messageId"])
		}(),
		PublishTime: func() *string {
			v, ok := data["publishTime"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["publishTime"])
		}(),
	}
}

func (p GooglePlayRealtimeNotificationMessage) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Data != nil {
		m["data"] = p.Data
	}
	if p.MessageId != nil {
		m["messageId"] = p.MessageId
	}
	if p.PublishTime != nil {
		m["publishTime"] = p.PublishTime
	}
	return m
}

func (p GooglePlayRealtimeNotificationMessage) Pointer() *GooglePlayRealtimeNotificationMessage {
	return &p
}

func CastGooglePlayRealtimeNotificationMessages(data []interface{}) []GooglePlayRealtimeNotificationMessage {
	v := make([]GooglePlayRealtimeNotificationMessage, 0)
	for _, d := range data {
		v = append(v, NewGooglePlayRealtimeNotificationMessageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGooglePlayRealtimeNotificationMessagesFromDict(data []GooglePlayRealtimeNotificationMessage) []interface{} {
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

type DailyTransactionHistory struct {
	DailyTransactionHistoryId *string  `json:"dailyTransactionHistoryId"`
	Year                      *int32   `json:"year"`
	Month                     *int32   `json:"month"`
	Day                       *int32   `json:"day"`
	Currency                  *string  `json:"currency"`
	DepositAmount             *float32 `json:"depositAmount"`
	WithdrawAmount            *float32 `json:"withdrawAmount"`
	UpdatedAt                 *int64   `json:"updatedAt"`
	Revision                  *int64   `json:"revision"`
}

func (p *DailyTransactionHistory) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DailyTransactionHistory{}
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
		*p = DailyTransactionHistory{}
	} else {
		*p = DailyTransactionHistory{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["dailyTransactionHistoryId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DailyTransactionHistoryId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DailyTransactionHistoryId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DailyTransactionHistoryId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DailyTransactionHistoryId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DailyTransactionHistoryId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DailyTransactionHistoryId)
				}
			}
		}
		if v, ok := d["year"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Year)
		}
		if v, ok := d["month"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Month)
		}
		if v, ok := d["day"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Day)
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
		if v, ok := d["depositAmount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositAmount)
		}
		if v, ok := d["withdrawAmount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WithdrawAmount)
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

func NewDailyTransactionHistoryFromJson(data string) DailyTransactionHistory {
	req := DailyTransactionHistory{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDailyTransactionHistoryFromDict(data map[string]interface{}) DailyTransactionHistory {
	return DailyTransactionHistory{
		DailyTransactionHistoryId: func() *string {
			v, ok := data["dailyTransactionHistoryId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["dailyTransactionHistoryId"])
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
		Day: func() *int32 {
			v, ok := data["day"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["day"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
		}(),
		DepositAmount: func() *float32 {
			v, ok := data["depositAmount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["depositAmount"])
		}(),
		WithdrawAmount: func() *float32 {
			v, ok := data["withdrawAmount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["withdrawAmount"])
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

func (p DailyTransactionHistory) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.DailyTransactionHistoryId != nil {
		m["dailyTransactionHistoryId"] = p.DailyTransactionHistoryId
	}
	if p.Year != nil {
		m["year"] = p.Year
	}
	if p.Month != nil {
		m["month"] = p.Month
	}
	if p.Day != nil {
		m["day"] = p.Day
	}
	if p.Currency != nil {
		m["currency"] = p.Currency
	}
	if p.DepositAmount != nil {
		m["depositAmount"] = p.DepositAmount
	}
	if p.WithdrawAmount != nil {
		m["withdrawAmount"] = p.WithdrawAmount
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p DailyTransactionHistory) Pointer() *DailyTransactionHistory {
	return &p
}

func CastDailyTransactionHistories(data []interface{}) []DailyTransactionHistory {
	v := make([]DailyTransactionHistory, 0)
	for _, d := range data {
		v = append(v, NewDailyTransactionHistoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDailyTransactionHistoriesFromDict(data []DailyTransactionHistory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type UnusedBalance struct {
	UnusedBalanceId *string  `json:"unusedBalanceId"`
	Currency        *string  `json:"currency"`
	Balance         *float32 `json:"balance"`
	UpdatedAt       *int64   `json:"updatedAt"`
	Revision        *int64   `json:"revision"`
}

func (p *UnusedBalance) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UnusedBalance{}
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
		*p = UnusedBalance{}
	} else {
		*p = UnusedBalance{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["unusedBalanceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UnusedBalanceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UnusedBalanceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UnusedBalanceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UnusedBalanceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UnusedBalanceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UnusedBalanceId)
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
		if v, ok := d["balance"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Balance)
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

func NewUnusedBalanceFromJson(data string) UnusedBalance {
	req := UnusedBalance{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewUnusedBalanceFromDict(data map[string]interface{}) UnusedBalance {
	return UnusedBalance{
		UnusedBalanceId: func() *string {
			v, ok := data["unusedBalanceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["unusedBalanceId"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
		}(),
		Balance: func() *float32 {
			v, ok := data["balance"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["balance"])
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

func (p UnusedBalance) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UnusedBalanceId != nil {
		m["unusedBalanceId"] = p.UnusedBalanceId
	}
	if p.Currency != nil {
		m["currency"] = p.Currency
	}
	if p.Balance != nil {
		m["balance"] = p.Balance
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p UnusedBalance) Pointer() *UnusedBalance {
	return &p
}

func CastUnusedBalances(data []interface{}) []UnusedBalance {
	v := make([]UnusedBalance, 0)
	for _, d := range data {
		v = append(v, NewUnusedBalanceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastUnusedBalancesFromDict(data []UnusedBalance) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
