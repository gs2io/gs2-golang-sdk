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

package lottery

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	LotteryTriggerScriptId *string `json:"lotteryTriggerScriptId"`
	ChoicePrizeTableScriptId *string `json:"choicePrizeTableScriptId"`
	LogSetting *LogSetting `json:"logSetting"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
	KeyId *string `json:"keyId"`
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
        TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
        LotteryTriggerScriptId: core.CastString(data["lotteryTriggerScriptId"]),
        ChoicePrizeTableScriptId: core.CastString(data["choicePrizeTableScriptId"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
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
    var transactionSetting map[string]interface{}
    if p.TransactionSetting != nil {
        transactionSetting = p.TransactionSetting.ToDict()
    }
    var lotteryTriggerScriptId *string
    if p.LotteryTriggerScriptId != nil {
        lotteryTriggerScriptId = p.LotteryTriggerScriptId
    }
    var choicePrizeTableScriptId *string
    if p.ChoicePrizeTableScriptId != nil {
        choicePrizeTableScriptId = p.ChoicePrizeTableScriptId
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
    var queueNamespaceId *string
    if p.QueueNamespaceId != nil {
        queueNamespaceId = p.QueueNamespaceId
    }
    var keyId *string
    if p.KeyId != nil {
        keyId = p.KeyId
    }
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "transactionSetting": transactionSetting,
        "lotteryTriggerScriptId": lotteryTriggerScriptId,
        "choicePrizeTableScriptId": choicePrizeTableScriptId,
        "logSetting": logSetting,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "queueNamespaceId": queueNamespaceId,
        "keyId": keyId,
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

type LotteryModelMaster struct {
	LotteryModelId *string `json:"lotteryModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	Mode *string `json:"mode"`
	Method *string `json:"method"`
	PrizeTableName *string `json:"prizeTableName"`
	ChoicePrizeTableScriptId *string `json:"choicePrizeTableScriptId"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewLotteryModelMasterFromJson(data string) LotteryModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLotteryModelMasterFromDict(dict)
}

func NewLotteryModelMasterFromDict(data map[string]interface{}) LotteryModelMaster {
    return LotteryModelMaster {
        LotteryModelId: core.CastString(data["lotteryModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        Mode: core.CastString(data["mode"]),
        Method: core.CastString(data["method"]),
        PrizeTableName: core.CastString(data["prizeTableName"]),
        ChoicePrizeTableScriptId: core.CastString(data["choicePrizeTableScriptId"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p LotteryModelMaster) ToDict() map[string]interface{} {
    
    var lotteryModelId *string
    if p.LotteryModelId != nil {
        lotteryModelId = p.LotteryModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var mode *string
    if p.Mode != nil {
        mode = p.Mode
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var prizeTableName *string
    if p.PrizeTableName != nil {
        prizeTableName = p.PrizeTableName
    }
    var choicePrizeTableScriptId *string
    if p.ChoicePrizeTableScriptId != nil {
        choicePrizeTableScriptId = p.ChoicePrizeTableScriptId
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
        "lotteryModelId": lotteryModelId,
        "name": name,
        "metadata": metadata,
        "description": description,
        "mode": mode,
        "method": method,
        "prizeTableName": prizeTableName,
        "choicePrizeTableScriptId": choicePrizeTableScriptId,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p LotteryModelMaster) Pointer() *LotteryModelMaster {
    return &p
}

func CastLotteryModelMasters(data []interface{}) []LotteryModelMaster {
	v := make([]LotteryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewLotteryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryModelMastersFromDict(data []LotteryModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type PrizeTableMaster struct {
	PrizeTableId *string `json:"prizeTableId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	Prizes []Prize `json:"prizes"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewPrizeTableMasterFromJson(data string) PrizeTableMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrizeTableMasterFromDict(dict)
}

func NewPrizeTableMasterFromDict(data map[string]interface{}) PrizeTableMaster {
    return PrizeTableMaster {
        PrizeTableId: core.CastString(data["prizeTableId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        Prizes: CastPrizes(core.CastArray(data["prizes"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p PrizeTableMaster) ToDict() map[string]interface{} {
    
    var prizeTableId *string
    if p.PrizeTableId != nil {
        prizeTableId = p.PrizeTableId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var prizes []interface{}
    if p.Prizes != nil {
        prizes = CastPrizesFromDict(
            p.Prizes,
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
        "prizeTableId": prizeTableId,
        "name": name,
        "metadata": metadata,
        "description": description,
        "prizes": prizes,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p PrizeTableMaster) Pointer() *PrizeTableMaster {
    return &p
}

func CastPrizeTableMasters(data []interface{}) []PrizeTableMaster {
	v := make([]PrizeTableMaster, 0)
	for _, d := range data {
		v = append(v, NewPrizeTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPrizeTableMastersFromDict(data []PrizeTableMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Box struct {
	BoxId *string `json:"boxId"`
	PrizeTableName *string `json:"prizeTableName"`
	Index *int32 `json:"index"`
	UserId *string `json:"userId"`
	DrawnIndexes []int32 `json:"drawnIndexes"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewBoxFromJson(data string) Box {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBoxFromDict(dict)
}

func NewBoxFromDict(data map[string]interface{}) Box {
    return Box {
        BoxId: core.CastString(data["boxId"]),
        PrizeTableName: core.CastString(data["prizeTableName"]),
        Index: core.CastInt32(data["index"]),
        UserId: core.CastString(data["userId"]),
        DrawnIndexes: core.CastInt32s(core.CastArray(data["drawnIndexes"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Box) ToDict() map[string]interface{} {
    
    var boxId *string
    if p.BoxId != nil {
        boxId = p.BoxId
    }
    var prizeTableName *string
    if p.PrizeTableName != nil {
        prizeTableName = p.PrizeTableName
    }
    var index *int32
    if p.Index != nil {
        index = p.Index
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var drawnIndexes []interface{}
    if p.DrawnIndexes != nil {
        drawnIndexes = core.CastInt32sFromDict(
            p.DrawnIndexes,
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
        "boxId": boxId,
        "prizeTableName": prizeTableName,
        "index": index,
        "userId": userId,
        "drawnIndexes": drawnIndexes,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p Box) Pointer() *Box {
    return &p
}

func CastBoxes(data []interface{}) []Box {
	v := make([]Box, 0)
	for _, d := range data {
		v = append(v, NewBoxFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBoxesFromDict(data []Box) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type LotteryModel struct {
	LotteryModelId *string `json:"lotteryModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Mode *string `json:"mode"`
	Method *string `json:"method"`
	PrizeTableName *string `json:"prizeTableName"`
	ChoicePrizeTableScriptId *string `json:"choicePrizeTableScriptId"`
}

func NewLotteryModelFromJson(data string) LotteryModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLotteryModelFromDict(dict)
}

func NewLotteryModelFromDict(data map[string]interface{}) LotteryModel {
    return LotteryModel {
        LotteryModelId: core.CastString(data["lotteryModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Mode: core.CastString(data["mode"]),
        Method: core.CastString(data["method"]),
        PrizeTableName: core.CastString(data["prizeTableName"]),
        ChoicePrizeTableScriptId: core.CastString(data["choicePrizeTableScriptId"]),
    }
}

func (p LotteryModel) ToDict() map[string]interface{} {
    
    var lotteryModelId *string
    if p.LotteryModelId != nil {
        lotteryModelId = p.LotteryModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var mode *string
    if p.Mode != nil {
        mode = p.Mode
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var prizeTableName *string
    if p.PrizeTableName != nil {
        prizeTableName = p.PrizeTableName
    }
    var choicePrizeTableScriptId *string
    if p.ChoicePrizeTableScriptId != nil {
        choicePrizeTableScriptId = p.ChoicePrizeTableScriptId
    }
    return map[string]interface{} {
        "lotteryModelId": lotteryModelId,
        "name": name,
        "metadata": metadata,
        "mode": mode,
        "method": method,
        "prizeTableName": prizeTableName,
        "choicePrizeTableScriptId": choicePrizeTableScriptId,
    }
}

func (p LotteryModel) Pointer() *LotteryModel {
    return &p
}

func CastLotteryModels(data []interface{}) []LotteryModel {
	v := make([]LotteryModel, 0)
	for _, d := range data {
		v = append(v, NewLotteryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryModelsFromDict(data []LotteryModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type PrizeTable struct {
	PrizeTableId *string `json:"prizeTableId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Prizes []Prize `json:"prizes"`
}

func NewPrizeTableFromJson(data string) PrizeTable {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrizeTableFromDict(dict)
}

func NewPrizeTableFromDict(data map[string]interface{}) PrizeTable {
    return PrizeTable {
        PrizeTableId: core.CastString(data["prizeTableId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Prizes: CastPrizes(core.CastArray(data["prizes"])),
    }
}

func (p PrizeTable) ToDict() map[string]interface{} {
    
    var prizeTableId *string
    if p.PrizeTableId != nil {
        prizeTableId = p.PrizeTableId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var prizes []interface{}
    if p.Prizes != nil {
        prizes = CastPrizesFromDict(
            p.Prizes,
        )
    }
    return map[string]interface{} {
        "prizeTableId": prizeTableId,
        "name": name,
        "metadata": metadata,
        "prizes": prizes,
    }
}

func (p PrizeTable) Pointer() *PrizeTable {
    return &p
}

func CastPrizeTables(data []interface{}) []PrizeTable {
	v := make([]PrizeTable, 0)
	for _, d := range data {
		v = append(v, NewPrizeTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPrizeTablesFromDict(data []PrizeTable) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Probability struct {
	Prize *DrawnPrize `json:"prize"`
	Rate *float32 `json:"rate"`
}

func NewProbabilityFromJson(data string) Probability {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewProbabilityFromDict(dict)
}

func NewProbabilityFromDict(data map[string]interface{}) Probability {
    return Probability {
        Prize: NewDrawnPrizeFromDict(core.CastMap(data["prize"])).Pointer(),
        Rate: core.CastFloat32(data["rate"]),
    }
}

func (p Probability) ToDict() map[string]interface{} {
    
    var prize map[string]interface{}
    if p.Prize != nil {
        prize = p.Prize.ToDict()
    }
    var rate *float32
    if p.Rate != nil {
        rate = p.Rate
    }
    return map[string]interface{} {
        "prize": prize,
        "rate": rate,
    }
}

func (p Probability) Pointer() *Probability {
    return &p
}

func CastProbabilities(data []interface{}) []Probability {
	v := make([]Probability, 0)
	for _, d := range data {
		v = append(v, NewProbabilityFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProbabilitiesFromDict(data []Probability) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CurrentLotteryMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings *string `json:"settings"`
}

func NewCurrentLotteryMasterFromJson(data string) CurrentLotteryMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentLotteryMasterFromDict(dict)
}

func NewCurrentLotteryMasterFromDict(data map[string]interface{}) CurrentLotteryMaster {
    return CurrentLotteryMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentLotteryMaster) ToDict() map[string]interface{} {
    
    var namespaceId *string
    if p.NamespaceId != nil {
        namespaceId = p.NamespaceId
    }
    var settings *string
    if p.Settings != nil {
        settings = p.Settings
    }
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "settings": settings,
    }
}

func (p CurrentLotteryMaster) Pointer() *CurrentLotteryMaster {
    return &p
}

func CastCurrentLotteryMasters(data []interface{}) []CurrentLotteryMaster {
	v := make([]CurrentLotteryMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentLotteryMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentLotteryMastersFromDict(data []CurrentLotteryMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Prize struct {
	PrizeId *string `json:"prizeId"`
	Type *string `json:"type"`
	AcquireActions []AcquireAction `json:"acquireActions"`
	PrizeTableName *string `json:"prizeTableName"`
	Weight *int32 `json:"weight"`
}

func NewPrizeFromJson(data string) Prize {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrizeFromDict(dict)
}

func NewPrizeFromDict(data map[string]interface{}) Prize {
    return Prize {
        PrizeId: core.CastString(data["prizeId"]),
        Type: core.CastString(data["type"]),
        AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
        PrizeTableName: core.CastString(data["prizeTableName"]),
        Weight: core.CastInt32(data["weight"]),
    }
}

func (p Prize) ToDict() map[string]interface{} {
    
    var prizeId *string
    if p.PrizeId != nil {
        prizeId = p.PrizeId
    }
    var _type *string
    if p.Type != nil {
        _type = p.Type
    }
    var acquireActions []interface{}
    if p.AcquireActions != nil {
        acquireActions = CastAcquireActionsFromDict(
            p.AcquireActions,
        )
    }
    var prizeTableName *string
    if p.PrizeTableName != nil {
        prizeTableName = p.PrizeTableName
    }
    var weight *int32
    if p.Weight != nil {
        weight = p.Weight
    }
    return map[string]interface{} {
        "prizeId": prizeId,
        "type": _type,
        "acquireActions": acquireActions,
        "prizeTableName": prizeTableName,
        "weight": weight,
    }
}

func (p Prize) Pointer() *Prize {
    return &p
}

func CastPrizes(data []interface{}) []Prize {
	v := make([]Prize, 0)
	for _, d := range data {
		v = append(v, NewPrizeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPrizesFromDict(data []Prize) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type DrawnPrize struct {
	AcquireActions []AcquireAction `json:"acquireActions"`
}

func NewDrawnPrizeFromJson(data string) DrawnPrize {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDrawnPrizeFromDict(dict)
}

func NewDrawnPrizeFromDict(data map[string]interface{}) DrawnPrize {
    return DrawnPrize {
        AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
    }
}

func (p DrawnPrize) ToDict() map[string]interface{} {
    
    var acquireActions []interface{}
    if p.AcquireActions != nil {
        acquireActions = CastAcquireActionsFromDict(
            p.AcquireActions,
        )
    }
    return map[string]interface{} {
        "acquireActions": acquireActions,
    }
}

func (p DrawnPrize) Pointer() *DrawnPrize {
    return &p
}

func CastDrawnPrizes(data []interface{}) []DrawnPrize {
	v := make([]DrawnPrize, 0)
	for _, d := range data {
		v = append(v, NewDrawnPrizeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDrawnPrizesFromDict(data []DrawnPrize) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type BoxItem struct {
	AcquireActions []AcquireAction `json:"acquireActions"`
	Remaining *int32 `json:"remaining"`
	Initial *int32 `json:"initial"`
}

func NewBoxItemFromJson(data string) BoxItem {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBoxItemFromDict(dict)
}

func NewBoxItemFromDict(data map[string]interface{}) BoxItem {
    return BoxItem {
        AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
        Remaining: core.CastInt32(data["remaining"]),
        Initial: core.CastInt32(data["initial"]),
    }
}

func (p BoxItem) ToDict() map[string]interface{} {
    
    var acquireActions []interface{}
    if p.AcquireActions != nil {
        acquireActions = CastAcquireActionsFromDict(
            p.AcquireActions,
        )
    }
    var remaining *int32
    if p.Remaining != nil {
        remaining = p.Remaining
    }
    var initial *int32
    if p.Initial != nil {
        initial = p.Initial
    }
    return map[string]interface{} {
        "acquireActions": acquireActions,
        "remaining": remaining,
        "initial": initial,
    }
}

func (p BoxItem) Pointer() *BoxItem {
    return &p
}

func CastBoxItems(data []interface{}) []BoxItem {
	v := make([]BoxItem, 0)
	for _, d := range data {
		v = append(v, NewBoxItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBoxItemsFromDict(data []BoxItem) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type BoxItems struct {
	BoxId *string `json:"boxId"`
	PrizeTableName *string `json:"prizeTableName"`
	UserId *string `json:"userId"`
	Items []BoxItem `json:"items"`
}

func NewBoxItemsFromJson(data string) BoxItems {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBoxItemsFromDict(dict)
}

func NewBoxItemsFromDict(data map[string]interface{}) BoxItems {
    return BoxItems {
        BoxId: core.CastString(data["boxId"]),
        PrizeTableName: core.CastString(data["prizeTableName"]),
        UserId: core.CastString(data["userId"]),
        Items: CastBoxItems(core.CastArray(data["items"])),
    }
}

func (p BoxItems) ToDict() map[string]interface{} {
    
    var boxId *string
    if p.BoxId != nil {
        boxId = p.BoxId
    }
    var prizeTableName *string
    if p.PrizeTableName != nil {
        prizeTableName = p.PrizeTableName
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var items []interface{}
    if p.Items != nil {
        items = CastBoxItemsFromDict(
            p.Items,
        )
    }
    return map[string]interface{} {
        "boxId": boxId,
        "prizeTableName": prizeTableName,
        "userId": userId,
        "items": items,
    }
}

func (p BoxItems) Pointer() *BoxItems {
    return &p
}

func CastBoxItemses(data []interface{}) []BoxItems {
	v := make([]BoxItems, 0)
	for _, d := range data {
		v = append(v, NewBoxItemsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBoxItemsesFromDict(data []BoxItems) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type AcquireAction struct {
	Action *string `json:"action"`
	Request *string `json:"request"`
}

func NewAcquireActionFromJson(data string) AcquireAction {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcquireActionFromDict(dict)
}

func NewAcquireActionFromDict(data map[string]interface{}) AcquireAction {
    return AcquireAction {
        Action: core.CastString(data["action"]),
        Request: core.CastString(data["request"]),
    }
}

func (p AcquireAction) ToDict() map[string]interface{} {
    
    var action *string
    if p.Action != nil {
        action = p.Action
    }
    var request *string
    if p.Request != nil {
        request = p.Request
    }
    return map[string]interface{} {
        "action": action,
        "request": request,
    }
}

func (p AcquireAction) Pointer() *AcquireAction {
    return &p
}

func CastAcquireActions(data []interface{}) []AcquireAction {
	v := make([]AcquireAction, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionsFromDict(data []AcquireAction) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Config struct {
	Key *string `json:"key"`
	Value *string `json:"value"`
}

func NewConfigFromJson(data string) Config {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConfigFromDict(dict)
}

func NewConfigFromDict(data map[string]interface{}) Config {
    return Config {
        Key: core.CastString(data["key"]),
        Value: core.CastString(data["value"]),
    }
}

func (p Config) ToDict() map[string]interface{} {
    
    var key *string
    if p.Key != nil {
        key = p.Key
    }
    var value *string
    if p.Value != nil {
        value = p.Value
    }
    return map[string]interface{} {
        "key": key,
        "value": value,
    }
}

func (p Config) Pointer() *Config {
    return &p
}

func CastConfigs(data []interface{}) []Config {
	v := make([]Config, 0)
	for _, d := range data {
		v = append(v, NewConfigFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConfigsFromDict(data []Config) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type GitHubCheckoutSetting struct {
	ApiKeyId *string `json:"apiKeyId"`
	RepositoryName *string `json:"repositoryName"`
	SourcePath *string `json:"sourcePath"`
	ReferenceType *string `json:"referenceType"`
	CommitHash *string `json:"commitHash"`
	BranchName *string `json:"branchName"`
	TagName *string `json:"tagName"`
}

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGitHubCheckoutSettingFromDict(dict)
}

func NewGitHubCheckoutSettingFromDict(data map[string]interface{}) GitHubCheckoutSetting {
    return GitHubCheckoutSetting {
        ApiKeyId: core.CastString(data["apiKeyId"]),
        RepositoryName: core.CastString(data["repositoryName"]),
        SourcePath: core.CastString(data["sourcePath"]),
        ReferenceType: core.CastString(data["referenceType"]),
        CommitHash: core.CastString(data["commitHash"]),
        BranchName: core.CastString(data["branchName"]),
        TagName: core.CastString(data["tagName"]),
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
    return map[string]interface{} {
        "apiKeyId": apiKeyId,
        "repositoryName": repositoryName,
        "sourcePath": sourcePath,
        "referenceType": referenceType,
        "commitHash": commitHash,
        "branchName": branchName,
        "tagName": tagName,
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

type TransactionSetting struct {
	EnableAutoRun *bool `json:"enableAutoRun"`
	DistributorNamespaceId *string `json:"distributorNamespaceId"`
	KeyId *string `json:"keyId"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
}

func NewTransactionSettingFromJson(data string) TransactionSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewTransactionSettingFromDict(dict)
}

func NewTransactionSettingFromDict(data map[string]interface{}) TransactionSetting {
    return TransactionSetting {
        EnableAutoRun: core.CastBool(data["enableAutoRun"]),
        DistributorNamespaceId: core.CastString(data["distributorNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
    }
}

func (p TransactionSetting) ToDict() map[string]interface{} {
    
    var enableAutoRun *bool
    if p.EnableAutoRun != nil {
        enableAutoRun = p.EnableAutoRun
    }
    var distributorNamespaceId *string
    if p.DistributorNamespaceId != nil {
        distributorNamespaceId = p.DistributorNamespaceId
    }
    var keyId *string
    if p.KeyId != nil {
        keyId = p.KeyId
    }
    var queueNamespaceId *string
    if p.QueueNamespaceId != nil {
        queueNamespaceId = p.QueueNamespaceId
    }
    return map[string]interface{} {
        "enableAutoRun": enableAutoRun,
        "distributorNamespaceId": distributorNamespaceId,
        "keyId": keyId,
        "queueNamespaceId": queueNamespaceId,
    }
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