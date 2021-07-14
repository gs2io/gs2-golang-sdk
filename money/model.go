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

import "core"

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
    return map[string]interface{} {
        "namespaceId": p.NamespaceId,
        "name": p.Name,
        "description": p.Description,
        "priority": p.Priority,
        "shareFree": p.ShareFree,
        "currency": p.Currency,
        "appleKey": p.AppleKey,
        "googleKey": p.GoogleKey,
        "enableFakeReceipt": p.EnableFakeReceipt,
        "createWalletScript": p.CreateWalletScript.ToDict(),
        "depositScript": p.DepositScript.ToDict(),
        "withdrawScript": p.WithdrawScript.ToDict(),
        "balance": p.Balance,
        "logSetting": p.LogSetting.ToDict(),
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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

type User struct {
	UserId *string `json:"userId"`
}

func NewUserFromDict(data map[string]interface{}) User {
    return User {
        UserId: core.CastString(data["userId"]),
    }
}

func (p User) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userId": p.UserId,
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
    return map[string]interface{} {
        "walletId": p.WalletId,
        "userId": p.UserId,
        "slot": p.Slot,
        "paid": p.Paid,
        "free": p.Free,
        "detail": CastWalletDetailsFromDict(
        p.Detail,
    ),
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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

func NewReceiptFromDict(data map[string]interface{}) Receipt {
    return Receipt {
        ReceiptId: core.CastString(data["receiptId"]),
        TransactionId: core.CastString(data["transactionId"]),
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
    return map[string]interface{} {
        "receiptId": p.ReceiptId,
        "transactionId": p.TransactionId,
        "userId": p.UserId,
        "type": p.Type,
        "slot": p.Slot,
        "price": p.Price,
        "paid": p.Paid,
        "free": p.Free,
        "total": p.Total,
        "contentsId": p.ContentsId,
        "createdAt": p.CreatedAt,
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

type ResponseCache struct {
	Region *string `json:"region"`
	ResponseCacheId *string `json:"responseCacheId"`
	RequestHash *string `json:"requestHash"`
	Result *string `json:"result"`
}

func NewResponseCacheFromDict(data map[string]interface{}) ResponseCache {
    return ResponseCache {
        Region: core.CastString(data["region"]),
        ResponseCacheId: core.CastString(data["responseCacheId"]),
        RequestHash: core.CastString(data["requestHash"]),
        Result: core.CastString(data["result"]),
    }
}

func (p ResponseCache) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "region": p.Region,
        "responseCacheId": p.ResponseCacheId,
        "requestHash": p.RequestHash,
        "result": p.Result,
    }
}

func (p ResponseCache) Pointer() *ResponseCache {
    return &p
}

func CastResponseCaches(data []interface{}) []ResponseCache {
	v := make([]ResponseCache, 0)
	for _, d := range data {
		v = append(v, NewResponseCacheFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastResponseCachesFromDict(data []ResponseCache) []interface{} {
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

func NewWalletDetailFromDict(data map[string]interface{}) WalletDetail {
    return WalletDetail {
        Price: core.CastFloat32(data["price"]),
        Count: core.CastInt32(data["count"]),
    }
}

func (p WalletDetail) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "price": p.Price,
        "count": p.Count,
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

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
    return ScriptSetting {
        TriggerScriptId: core.CastString(data["triggerScriptId"]),
        DoneTriggerTargetType: core.CastString(data["doneTriggerTargetType"]),
        DoneTriggerScriptId: core.CastString(data["doneTriggerScriptId"]),
        DoneTriggerQueueNamespaceId: core.CastString(data["doneTriggerQueueNamespaceId"]),
    }
}

func (p ScriptSetting) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "triggerScriptId": p.TriggerScriptId,
        "doneTriggerTargetType": p.DoneTriggerTargetType,
        "doneTriggerScriptId": p.DoneTriggerScriptId,
        "doneTriggerQueueNamespaceId": p.DoneTriggerQueueNamespaceId,
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

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
    return LogSetting {
        LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
    }
}

func (p LogSetting) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "loggingNamespaceId": p.LoggingNamespaceId,
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