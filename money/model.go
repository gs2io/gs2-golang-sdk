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

package money

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId        *string             `json:"namespaceId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	Priority           *string             `json:"priority"`
	ShareFree          *bool               `json:"shareFree"`
	Currency           *string             `json:"currency"`
	AppleKey           *string             `json:"appleKey"`
	GoogleKey          *string             `json:"googleKey"`
	EnableFakeReceipt  *bool               `json:"enableFakeReceipt"`
	CreateWalletScript *ScriptSetting      `json:"createWalletScript"`
	DepositScript      *ScriptSetting      `json:"depositScript"`
	WithdrawScript     *ScriptSetting      `json:"withdrawScript"`
	Balance            *float64            `json:"balance"`
	LogSetting         *LogSetting         `json:"logSetting"`
	CreatedAt          *int64              `json:"createdAt"`
	UpdatedAt          *int64              `json:"updatedAt"`
	Revision           *int64              `json:"revision"`
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
		if v, ok := d["transactionSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionSetting)
		}
		if v, ok := d["priority"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Priority = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Priority = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Priority = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Priority = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Priority = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Priority)
				}
			}
		}
		if v, ok := d["shareFree"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ShareFree)
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
		if v, ok := d["appleKey"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AppleKey = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AppleKey = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AppleKey = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AppleKey = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AppleKey = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AppleKey)
				}
			}
		}
		if v, ok := d["googleKey"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GoogleKey = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GoogleKey = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GoogleKey = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GoogleKey = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GoogleKey = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GoogleKey)
				}
			}
		}
		if v, ok := d["enableFakeReceipt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableFakeReceipt)
		}
		if v, ok := d["createWalletScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreateWalletScript)
		}
		if v, ok := d["depositScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositScript)
		}
		if v, ok := d["withdrawScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WithdrawScript)
		}
		if v, ok := d["balance"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Balance)
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
		TransactionSetting: func() *TransactionSetting {
			v, ok := data["transactionSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer()
		}(),
		Priority: func() *string {
			v, ok := data["priority"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["priority"])
		}(),
		ShareFree: func() *bool {
			v, ok := data["shareFree"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["shareFree"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
		}(),
		AppleKey: func() *string {
			v, ok := data["appleKey"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["appleKey"])
		}(),
		GoogleKey: func() *string {
			v, ok := data["googleKey"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["googleKey"])
		}(),
		EnableFakeReceipt: func() *bool {
			v, ok := data["enableFakeReceipt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableFakeReceipt"])
		}(),
		CreateWalletScript: func() *ScriptSetting {
			v, ok := data["createWalletScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["createWalletScript"])).Pointer()
		}(),
		DepositScript: func() *ScriptSetting {
			v, ok := data["depositScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["depositScript"])).Pointer()
		}(),
		WithdrawScript: func() *ScriptSetting {
			v, ok := data["withdrawScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["withdrawScript"])).Pointer()
		}(),
		Balance: func() *float64 {
			v, ok := data["balance"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat64(data["balance"])
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
	if p.TransactionSetting != nil {
		m["transactionSetting"] = func() map[string]interface{} {
			if p.TransactionSetting == nil {
				return nil
			}
			return p.TransactionSetting.ToDict()
		}()
	}
	if p.Priority != nil {
		m["priority"] = p.Priority
	}
	if p.ShareFree != nil {
		m["shareFree"] = p.ShareFree
	}
	if p.Currency != nil {
		m["currency"] = p.Currency
	}
	if p.AppleKey != nil {
		m["appleKey"] = p.AppleKey
	}
	if p.GoogleKey != nil {
		m["googleKey"] = p.GoogleKey
	}
	if p.EnableFakeReceipt != nil {
		m["enableFakeReceipt"] = p.EnableFakeReceipt
	}
	if p.CreateWalletScript != nil {
		m["createWalletScript"] = func() map[string]interface{} {
			if p.CreateWalletScript == nil {
				return nil
			}
			return p.CreateWalletScript.ToDict()
		}()
	}
	if p.DepositScript != nil {
		m["depositScript"] = func() map[string]interface{} {
			if p.DepositScript == nil {
				return nil
			}
			return p.DepositScript.ToDict()
		}()
	}
	if p.WithdrawScript != nil {
		m["withdrawScript"] = func() map[string]interface{} {
			if p.WithdrawScript == nil {
				return nil
			}
			return p.WithdrawScript.ToDict()
		}()
	}
	if p.Balance != nil {
		m["balance"] = p.Balance
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
	WalletId  *string        `json:"walletId"`
	UserId    *string        `json:"userId"`
	Slot      *int32         `json:"slot"`
	Paid      *int32         `json:"paid"`
	Free      *int32         `json:"free"`
	Detail    []WalletDetail `json:"detail"`
	ShareFree *bool          `json:"shareFree"`
	CreatedAt *int64         `json:"createdAt"`
	UpdatedAt *int64         `json:"updatedAt"`
	Revision  *int64         `json:"revision"`
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
		if v, ok := d["paid"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Paid)
		}
		if v, ok := d["free"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Free)
		}
		if v, ok := d["detail"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Detail)
		}
		if v, ok := d["shareFree"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ShareFree)
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
		Detail: func() []WalletDetail {
			if data["detail"] == nil {
				return nil
			}
			return CastWalletDetails(core.CastArray(data["detail"]))
		}(),
		ShareFree: func() *bool {
			v, ok := data["shareFree"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["shareFree"])
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
	if p.Paid != nil {
		m["paid"] = p.Paid
	}
	if p.Free != nil {
		m["free"] = p.Free
	}
	if p.Detail != nil {
		m["detail"] = CastWalletDetailsFromDict(
			p.Detail,
		)
	}
	if p.ShareFree != nil {
		m["shareFree"] = p.ShareFree
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

type Receipt struct {
	ReceiptId     *string  `json:"receiptId"`
	TransactionId *string  `json:"transactionId"`
	PurchaseToken *string  `json:"purchaseToken"`
	UserId        *string  `json:"userId"`
	Type          *string  `json:"type"`
	Slot          *int32   `json:"slot"`
	Price         *float32 `json:"price"`
	Paid          *int32   `json:"paid"`
	Free          *int32   `json:"free"`
	Total         *int32   `json:"total"`
	ContentsId    *string  `json:"contentsId"`
	CreatedAt     *int64   `json:"createdAt"`
	Revision      *int64   `json:"revision"`
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
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Type = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Type = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Type = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Type)
				}
			}
		}
		if v, ok := d["slot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Slot)
		}
		if v, ok := d["price"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Price)
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
		if v, ok := d["contentsId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentsId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentsId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentsId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentsId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentsId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentsId)
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
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		PurchaseToken: func() *string {
			v, ok := data["purchaseToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["purchaseToken"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		Slot: func() *int32 {
			v, ok := data["slot"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["slot"])
		}(),
		Price: func() *float32 {
			v, ok := data["price"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["price"])
		}(),
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
		ContentsId: func() *string {
			v, ok := data["contentsId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentsId"])
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

func (p Receipt) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ReceiptId != nil {
		m["receiptId"] = p.ReceiptId
	}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.PurchaseToken != nil {
		m["purchaseToken"] = p.PurchaseToken
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.Slot != nil {
		m["slot"] = p.Slot
	}
	if p.Price != nil {
		m["price"] = p.Price
	}
	if p.Paid != nil {
		m["paid"] = p.Paid
	}
	if p.Free != nil {
		m["free"] = p.Free
	}
	if p.Total != nil {
		m["total"] = p.Total
	}
	if p.ContentsId != nil {
		m["contentsId"] = p.ContentsId
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
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

type WalletDetail struct {
	Price *float32 `json:"price"`
	Count *int32   `json:"count"`
}

func (p *WalletDetail) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WalletDetail{}
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
		*p = WalletDetail{}
	} else {
		*p = WalletDetail{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["price"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Price)
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewWalletDetailFromJson(data string) WalletDetail {
	req := WalletDetail{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewWalletDetailFromDict(data map[string]interface{}) WalletDetail {
	return WalletDetail{
		Price: func() *float32 {
			v, ok := data["price"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["price"])
		}(),
		Count: func() *int32 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["count"])
		}(),
	}
}

func (p WalletDetail) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Price != nil {
		m["price"] = p.Price
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	return m
}

func (p WalletDetail) Pointer() *WalletDetail {
	return &p
}

func CastWalletDetails(data []interface{}) []WalletDetail {
	v := make([]WalletDetail, 0)
	for _, d := range data {
		v = append(v, NewWalletDetailFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastWalletDetailsFromDict(data []WalletDetail) []interface{} {
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

type TransactionSetting struct {
	EnableAutoRun                      *bool   `json:"enableAutoRun"`
	EnableAtomicCommit                 *bool   `json:"enableAtomicCommit"`
	TransactionUseDistributor          *bool   `json:"transactionUseDistributor"`
	CommitScriptResultInUseDistributor *bool   `json:"commitScriptResultInUseDistributor"`
	AcquireActionUseJobQueue           *bool   `json:"acquireActionUseJobQueue"`
	DistributorNamespaceId             *string `json:"distributorNamespaceId"`
	// Deprecated: should not be used
	KeyId            *string `json:"keyId"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
}

func (p *TransactionSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TransactionSetting{}
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
		*p = TransactionSetting{}
	} else {
		*p = TransactionSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["enableAutoRun"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableAutoRun)
		}
		if v, ok := d["enableAtomicCommit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableAtomicCommit)
		}
		if v, ok := d["transactionUseDistributor"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionUseDistributor)
		}
		if v, ok := d["commitScriptResultInUseDistributor"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CommitScriptResultInUseDistributor)
		}
		if v, ok := d["acquireActionUseJobQueue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActionUseJobQueue)
		}
		if v, ok := d["distributorNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DistributorNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DistributorNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DistributorNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DistributorNamespaceId)
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
		if v, ok := d["queueNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QueueNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QueueNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QueueNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QueueNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewTransactionSettingFromJson(data string) TransactionSetting {
	req := TransactionSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTransactionSettingFromDict(data map[string]interface{}) TransactionSetting {
	return TransactionSetting{
		EnableAutoRun: func() *bool {
			v, ok := data["enableAutoRun"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableAutoRun"])
		}(),
		EnableAtomicCommit: func() *bool {
			v, ok := data["enableAtomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableAtomicCommit"])
		}(),
		TransactionUseDistributor: func() *bool {
			v, ok := data["transactionUseDistributor"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["transactionUseDistributor"])
		}(),
		CommitScriptResultInUseDistributor: func() *bool {
			v, ok := data["commitScriptResultInUseDistributor"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["commitScriptResultInUseDistributor"])
		}(),
		AcquireActionUseJobQueue: func() *bool {
			v, ok := data["acquireActionUseJobQueue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["acquireActionUseJobQueue"])
		}(),
		DistributorNamespaceId: func() *string {
			v, ok := data["distributorNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["distributorNamespaceId"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
		QueueNamespaceId: func() *string {
			v, ok := data["queueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["queueNamespaceId"])
		}(),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.EnableAutoRun != nil {
		m["enableAutoRun"] = p.EnableAutoRun
	}
	if p.EnableAtomicCommit != nil {
		m["enableAtomicCommit"] = p.EnableAtomicCommit
	}
	if p.TransactionUseDistributor != nil {
		m["transactionUseDistributor"] = p.TransactionUseDistributor
	}
	if p.CommitScriptResultInUseDistributor != nil {
		m["commitScriptResultInUseDistributor"] = p.CommitScriptResultInUseDistributor
	}
	if p.AcquireActionUseJobQueue != nil {
		m["acquireActionUseJobQueue"] = p.AcquireActionUseJobQueue
	}
	if p.DistributorNamespaceId != nil {
		m["distributorNamespaceId"] = p.DistributorNamespaceId
	}
	if p.KeyId != nil {
		m["keyId"] = p.KeyId
	}
	if p.QueueNamespaceId != nil {
		m["queueNamespaceId"] = p.QueueNamespaceId
	}
	return m
}

func (p TransactionSetting) Pointer() *TransactionSetting {
	return &p
}

func CastTransactionSettings(data []interface{}) []TransactionSetting {
	v := make([]TransactionSetting, 0)
	for _, d := range data {
		v = append(v, NewTransactionSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionSettingsFromDict(data []TransactionSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
