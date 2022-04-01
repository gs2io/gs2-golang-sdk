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
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Priority *string `json:"priority"`
	ShareFree *bool `json:"shareFree"`
	Currency *string `json:"currency"`
	AppleKey *string `json:"appleKey"`
	GoogleKey *string `json:"googleKey"`
	EnableFakeReceipt *bool `json:"enableFakeReceipt"`
	CreateWalletScript *ScriptSetting `json:"createWalletScript"`
	DepositScript *ScriptSetting `json:"depositScript"`
	WithdrawScript *ScriptSetting `json:"withdrawScript"`
	Balance *float64 `json:"balance"`
	LogSetting *LogSetting `json:"logSetting"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
    return Namespace {
        NamespaceId: core.CastString(data["namespaceId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Priority: core.CastString(data["priority"]),
        ShareFree: core.CastBool(data["shareFree"]),
        Currency: core.CastString(data["currency"]),
        AppleKey: core.CastString(data["appleKey"]),
        GoogleKey: core.CastString(data["googleKey"]),
        EnableFakeReceipt: core.CastBool(data["enableFakeReceipt"]),
        CreateWalletScript: NewScriptSettingFromDict(core.CastMap(data["createWalletScript"])).Pointer(),
        DepositScript: NewScriptSettingFromDict(core.CastMap(data["depositScript"])).Pointer(),
        WithdrawScript: NewScriptSettingFromDict(core.CastMap(data["withdrawScript"])).Pointer(),
        Balance: core.CastFloat64(data["balance"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
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
    var priority *string
    if p.Priority != nil {
        priority = p.Priority
    }
    var shareFree *bool
    if p.ShareFree != nil {
        shareFree = p.ShareFree
    }
    var currency *string
    if p.Currency != nil {
        currency = p.Currency
    }
    var appleKey *string
    if p.AppleKey != nil {
        appleKey = p.AppleKey
    }
    var googleKey *string
    if p.GoogleKey != nil {
        googleKey = p.GoogleKey
    }
    var enableFakeReceipt *bool
    if p.EnableFakeReceipt != nil {
        enableFakeReceipt = p.EnableFakeReceipt
    }
    var createWalletScript map[string]interface{}
    if p.CreateWalletScript != nil {
        createWalletScript = p.CreateWalletScript.ToDict()
    }
    var depositScript map[string]interface{}
    if p.DepositScript != nil {
        depositScript = p.DepositScript.ToDict()
    }
    var withdrawScript map[string]interface{}
    if p.WithdrawScript != nil {
        withdrawScript = p.WithdrawScript.ToDict()
    }
    var balance *float64
    if p.Balance != nil {
        balance = p.Balance
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
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "priority": priority,
        "shareFree": shareFree,
        "currency": currency,
        "appleKey": appleKey,
        "googleKey": googleKey,
        "enableFakeReceipt": enableFakeReceipt,
        "createWalletScript": createWalletScript,
        "depositScript": depositScript,
        "withdrawScript": withdrawScript,
        "balance": balance,
        "logSetting": logSetting,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
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
	WalletId *string `json:"walletId"`
	UserId *string `json:"userId"`
	Slot *int32 `json:"slot"`
	Paid *int32 `json:"paid"`
	Free *int32 `json:"free"`
	Detail []WalletDetail `json:"detail"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewWalletFromJson(data string) Wallet {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewWalletFromDict(dict)
}

func NewWalletFromDict(data map[string]interface{}) Wallet {
    return Wallet {
        WalletId: core.CastString(data["walletId"]),
        UserId: core.CastString(data["userId"]),
        Slot: core.CastInt32(data["slot"]),
        Paid: core.CastInt32(data["paid"]),
        Free: core.CastInt32(data["free"]),
        Detail: CastWalletDetails(core.CastArray(data["detail"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
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
    var paid *int32
    if p.Paid != nil {
        paid = p.Paid
    }
    var free *int32
    if p.Free != nil {
        free = p.Free
    }
    var detail []interface{}
    if p.Detail != nil {
        detail = CastWalletDetailsFromDict(
            p.Detail,
        )
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "walletId": walletId,
        "userId": userId,
        "slot": slot,
        "paid": paid,
        "free": free,
        "detail": detail,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
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

type Receipt struct {
	ReceiptId *string `json:"receiptId"`
	TransactionId *string `json:"transactionId"`
	PurchaseToken *string `json:"purchaseToken"`
	UserId *string `json:"userId"`
	Type *string `json:"type"`
	Slot *int32 `json:"slot"`
	Price *float32 `json:"price"`
	Paid *int32 `json:"paid"`
	Free *int32 `json:"free"`
	Total *int32 `json:"total"`
	ContentsId *string `json:"contentsId"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewReceiptFromJson(data string) Receipt {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewReceiptFromDict(dict)
}

func NewReceiptFromDict(data map[string]interface{}) Receipt {
    return Receipt {
        ReceiptId: core.CastString(data["receiptId"]),
        TransactionId: core.CastString(data["transactionId"]),
        PurchaseToken: core.CastString(data["purchaseToken"]),
        UserId: core.CastString(data["userId"]),
        Type: core.CastString(data["type"]),
        Slot: core.CastInt32(data["slot"]),
        Price: core.CastFloat32(data["price"]),
        Paid: core.CastInt32(data["paid"]),
        Free: core.CastInt32(data["free"]),
        Total: core.CastInt32(data["total"]),
        ContentsId: core.CastString(data["contentsId"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
    }
}

func (p Receipt) ToDict() map[string]interface{} {
    
    var receiptId *string
    if p.ReceiptId != nil {
        receiptId = p.ReceiptId
    }
    var transactionId *string
    if p.TransactionId != nil {
        transactionId = p.TransactionId
    }
    var purchaseToken *string
    if p.PurchaseToken != nil {
        purchaseToken = p.PurchaseToken
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var _type *string
    if p.Type != nil {
        _type = p.Type
    }
    var slot *int32
    if p.Slot != nil {
        slot = p.Slot
    }
    var price *float32
    if p.Price != nil {
        price = p.Price
    }
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
    var contentsId *string
    if p.ContentsId != nil {
        contentsId = p.ContentsId
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    return map[string]interface{} {
        "receiptId": receiptId,
        "transactionId": transactionId,
        "purchaseToken": purchaseToken,
        "userId": userId,
        "type": _type,
        "slot": slot,
        "price": price,
        "paid": paid,
        "free": free,
        "total": total,
        "contentsId": contentsId,
        "createdAt": createdAt,
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

type WalletDetail struct {
	Price *float32 `json:"price"`
	Count *int32 `json:"count"`
}

func NewWalletDetailFromJson(data string) WalletDetail {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewWalletDetailFromDict(dict)
}

func NewWalletDetailFromDict(data map[string]interface{}) WalletDetail {
    return WalletDetail {
        Price: core.CastFloat32(data["price"]),
        Count: core.CastInt32(data["count"]),
    }
}

func (p WalletDetail) ToDict() map[string]interface{} {
    
    var price *float32
    if p.Price != nil {
        price = p.Price
    }
    var count *int32
    if p.Count != nil {
        count = p.Count
    }
    return map[string]interface{} {
        "price": price,
        "count": count,
    }
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
	TriggerScriptId *string `json:"triggerScriptId"`
	DoneTriggerTargetType *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func NewScriptSettingFromJson(data string) ScriptSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewScriptSettingFromDict(dict)
}

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
    return ScriptSetting {
        TriggerScriptId: core.CastString(data["triggerScriptId"]),
        DoneTriggerTargetType: core.CastString(data["doneTriggerTargetType"]),
        DoneTriggerScriptId: core.CastString(data["doneTriggerScriptId"]),
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
    return map[string]interface{} {
        "triggerScriptId": triggerScriptId,
        "doneTriggerTargetType": doneTriggerTargetType,
        "doneTriggerScriptId": doneTriggerScriptId,
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

func NewLogSettingFromJson(data string) LogSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLogSettingFromDict(dict)
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
    return LogSetting {
        LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
    }
}

func (p LogSetting) ToDict() map[string]interface{} {
    
    var loggingNamespaceId *string
    if p.LoggingNamespaceId != nil {
        loggingNamespaceId = p.LoggingNamespaceId
    }
    return map[string]interface{} {
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