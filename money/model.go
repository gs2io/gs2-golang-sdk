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
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
    /** ネームスペース */
	NamespaceId *core.String   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ネームスペースの名前 */
	Name *core.String   `json:"name"`
    /** ネームスペースの説明 */
	Description *core.String   `json:"description"`
    /** 消費優先度 */
	Priority *core.String   `json:"priority"`
    /** 無償課金通貨を異なるスロットで共有するか */
	ShareFree *bool   `json:"shareFree"`
    /** 通貨の種類 */
	Currency *core.String   `json:"currency"`
    /** Apple AppStore のバンドルID */
	AppleKey *core.String   `json:"appleKey"`
    /** Google PlayStore の秘密鍵 */
	GoogleKey *core.String   `json:"googleKey"`
    /** UnityEditorが出力する偽のレシートで決済できるようにするか */
	EnableFakeReceipt *bool   `json:"enableFakeReceipt"`
    /** ウォレット新規作成したときに実行するスクリプト */
	CreateWalletScript *ScriptSetting   `json:"createWalletScript"`
    /** ウォレット残高加算したときに実行するスクリプト */
	DepositScript *ScriptSetting   `json:"depositScript"`
    /** ウォレット残高消費したときに実行するスクリプト */
	WithdrawScript *ScriptSetting   `json:"withdrawScript"`
    /** 未使用残高 */
	Balance *float64   `json:"balance"`
    /** ログの出力設定 */
	LogSetting *LogSetting   `json:"logSetting"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Namespace) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["ownerId"] = p.OwnerId
    data["name"] = p.Name
    data["description"] = p.Description
    data["priority"] = p.Priority
    data["shareFree"] = p.ShareFree
    data["currency"] = p.Currency
    data["appleKey"] = p.AppleKey
    data["googleKey"] = p.GoogleKey
    data["enableFakeReceipt"] = p.EnableFakeReceipt
    if p.CreateWalletScript != nil {
        data["createWalletScript"] = *p.CreateWalletScript.ToDict()
    }
    if p.DepositScript != nil {
        data["depositScript"] = *p.DepositScript.ToDict()
    }
    if p.WithdrawScript != nil {
        data["withdrawScript"] = *p.WithdrawScript.ToDict()
    }
    data["balance"] = p.Balance
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Wallet struct {
    /** ウォレット */
	WalletId *core.String   `json:"walletId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** スロット番号 */
	Slot *int32   `json:"slot"`
    /** 有償課金通貨所持量 */
	Paid *int32   `json:"paid"`
    /** 無償課金通貨所持量 */
	Free *int32   `json:"free"`
    /** 詳細 */
	Detail *[]*WalletDetail   `json:"detail"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Wallet) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["walletId"] = p.WalletId
    data["userId"] = p.UserId
    data["slot"] = p.Slot
    data["paid"] = p.Paid
    data["free"] = p.Free
    if p.Detail != nil {
        var _detail []*map[string]interface {}
        for _, item := range *p.Detail {
            _detail = append(_detail, item.ToDict())
        }
        data["detail"] = &_detail
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Receipt struct {
    /** レシート */
	ReceiptId *core.String   `json:"receiptId"`
    /** トランザクションID */
	TransactionId *core.String   `json:"transactionId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 種類 */
	Type *core.String   `json:"type"`
    /** スロット番号 */
	Slot *int32   `json:"slot"`
    /** 単価 */
	Price *float32   `json:"price"`
    /** 有償課金通貨 */
	Paid *int32   `json:"paid"`
    /** 無償課金通貨 */
	Free *int32   `json:"free"`
    /** 総数 */
	Total *int32   `json:"total"`
    /** ストアプラットフォームで販売されているコンテンツID */
	ContentsId *core.String   `json:"contentsId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Receipt) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["receiptId"] = p.ReceiptId
    data["transactionId"] = p.TransactionId
    data["userId"] = p.UserId
    data["type"] = p.Type
    data["slot"] = p.Slot
    data["price"] = p.Price
    data["paid"] = p.Paid
    data["free"] = p.Free
    data["total"] = p.Total
    data["contentsId"] = p.ContentsId
    data["createdAt"] = p.CreatedAt
    return &data
}

type ResponseCache struct {
    /** None */
	Region *core.String   `json:"region"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** レスポンスキャッシュ のGRN */
	ResponseCacheId *core.String   `json:"responseCacheId"`
    /** None */
	RequestHash *core.String   `json:"requestHash"`
    /** APIの応答内容 */
	Result *core.String   `json:"result"`
}

func (p *ResponseCache) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["region"] = p.Region
    data["ownerId"] = p.OwnerId
    data["responseCacheId"] = p.ResponseCacheId
    data["requestHash"] = p.RequestHash
    data["result"] = p.Result
    return &data
}

type WalletDetail struct {
    /** 単価 */
	Price *float32   `json:"price"`
    /** 所持量 */
	Count *int32   `json:"count"`
}

func (p *WalletDetail) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["price"] = p.Price
    data["count"] = p.Count
    return &data
}

type ScriptSetting struct {
    /** 実行前に使用する GS2-Script のスクリプト のGRN */
	TriggerScriptId *core.String   `json:"triggerScriptId"`
    /** 完了通知の通知先 */
	DoneTriggerTargetType *core.String   `json:"doneTriggerTargetType"`
    /** 完了時に使用する GS2-Script のスクリプト のGRN */
	DoneTriggerScriptId *core.String   `json:"doneTriggerScriptId"`
    /** 完了時に使用する GS2-JobQueue のネームスペース のGRN */
	DoneTriggerQueueNamespaceId *core.String   `json:"doneTriggerQueueNamespaceId"`
}

func (p *ScriptSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["triggerScriptId"] = p.TriggerScriptId
    data["doneTriggerTargetType"] = p.DoneTriggerTargetType
    data["doneTriggerScriptId"] = p.DoneTriggerScriptId
    data["doneTriggerQueueNamespaceId"] = p.DoneTriggerQueueNamespaceId
    return &data
}

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *core.String   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}
