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

package quest

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	StartQuestScript *ScriptSetting `json:"startQuestScript"`
	CompleteQuestScript *ScriptSetting `json:"completeQuestScript"`
	FailedQuestScript *ScriptSetting `json:"failedQuestScript"`
	LogSetting *LogSetting `json:"logSetting"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
    // Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
    // Deprecated: should not be used
	KeyId *string `json:"keyId"`
	Revision *int64 `json:"revision"`
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
        StartQuestScript: NewScriptSettingFromDict(core.CastMap(data["startQuestScript"])).Pointer(),
        CompleteQuestScript: NewScriptSettingFromDict(core.CastMap(data["completeQuestScript"])).Pointer(),
        FailedQuestScript: NewScriptSettingFromDict(core.CastMap(data["failedQuestScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
        Revision: core.CastInt64(data["revision"]),
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
    var startQuestScript map[string]interface{}
    if p.StartQuestScript != nil {
        startQuestScript = p.StartQuestScript.ToDict()
    }
    var completeQuestScript map[string]interface{}
    if p.CompleteQuestScript != nil {
        completeQuestScript = p.CompleteQuestScript.ToDict()
    }
    var failedQuestScript map[string]interface{}
    if p.FailedQuestScript != nil {
        failedQuestScript = p.FailedQuestScript.ToDict()
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
    var revision *int64
    if p.Revision != nil {
        revision = p.Revision
    }
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "transactionSetting": transactionSetting,
        "startQuestScript": startQuestScript,
        "completeQuestScript": completeQuestScript,
        "failedQuestScript": failedQuestScript,
        "logSetting": logSetting,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "queueNamespaceId": queueNamespaceId,
        "keyId": keyId,
        "revision": revision,
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

type QuestGroupModelMaster struct {
	QuestGroupModelId *string `json:"questGroupModelId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewQuestGroupModelMasterFromJson(data string) QuestGroupModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewQuestGroupModelMasterFromDict(dict)
}

func NewQuestGroupModelMasterFromDict(data map[string]interface{}) QuestGroupModelMaster {
    return QuestGroupModelMaster {
        QuestGroupModelId: core.CastString(data["questGroupModelId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p QuestGroupModelMaster) ToDict() map[string]interface{} {
    
    var questGroupModelId *string
    if p.QuestGroupModelId != nil {
        questGroupModelId = p.QuestGroupModelId
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
    var challengePeriodEventId *string
    if p.ChallengePeriodEventId != nil {
        challengePeriodEventId = p.ChallengePeriodEventId
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
    return map[string]interface{} {
        "questGroupModelId": questGroupModelId,
        "name": name,
        "description": description,
        "metadata": metadata,
        "challengePeriodEventId": challengePeriodEventId,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p QuestGroupModelMaster) Pointer() *QuestGroupModelMaster {
    return &p
}

func CastQuestGroupModelMasters(data []interface{}) []QuestGroupModelMaster {
	v := make([]QuestGroupModelMaster, 0)
	for _, d := range data {
		v = append(v, NewQuestGroupModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestGroupModelMastersFromDict(data []QuestGroupModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type QuestModelMaster struct {
	QuestModelId *string `json:"questModelId"`
	QuestGroupName *string `json:"questGroupName"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	Contents []Contents `json:"contents"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
	FirstCompleteAcquireActions []AcquireAction `json:"firstCompleteAcquireActions"`
	ConsumeActions []ConsumeAction `json:"consumeActions"`
	FailedAcquireActions []AcquireAction `json:"failedAcquireActions"`
	PremiseQuestNames []*string `json:"premiseQuestNames"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewQuestModelMasterFromJson(data string) QuestModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewQuestModelMasterFromDict(dict)
}

func NewQuestModelMasterFromDict(data map[string]interface{}) QuestModelMaster {
    return QuestModelMaster {
        QuestModelId: core.CastString(data["questModelId"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Contents: CastContentses(core.CastArray(data["contents"])),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
        FirstCompleteAcquireActions: CastAcquireActions(core.CastArray(data["firstCompleteAcquireActions"])),
        ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
        FailedAcquireActions: CastAcquireActions(core.CastArray(data["failedAcquireActions"])),
        PremiseQuestNames: core.CastStrings(core.CastArray(data["premiseQuestNames"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p QuestModelMaster) ToDict() map[string]interface{} {
    
    var questModelId *string
    if p.QuestModelId != nil {
        questModelId = p.QuestModelId
    }
    var questGroupName *string
    if p.QuestGroupName != nil {
        questGroupName = p.QuestGroupName
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
    var contents []interface{}
    if p.Contents != nil {
        contents = CastContentsesFromDict(
            p.Contents,
        )
    }
    var challengePeriodEventId *string
    if p.ChallengePeriodEventId != nil {
        challengePeriodEventId = p.ChallengePeriodEventId
    }
    var firstCompleteAcquireActions []interface{}
    if p.FirstCompleteAcquireActions != nil {
        firstCompleteAcquireActions = CastAcquireActionsFromDict(
            p.FirstCompleteAcquireActions,
        )
    }
    var consumeActions []interface{}
    if p.ConsumeActions != nil {
        consumeActions = CastConsumeActionsFromDict(
            p.ConsumeActions,
        )
    }
    var failedAcquireActions []interface{}
    if p.FailedAcquireActions != nil {
        failedAcquireActions = CastAcquireActionsFromDict(
            p.FailedAcquireActions,
        )
    }
    var premiseQuestNames []interface{}
    if p.PremiseQuestNames != nil {
        premiseQuestNames = core.CastStringsFromDict(
            p.PremiseQuestNames,
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
    var revision *int64
    if p.Revision != nil {
        revision = p.Revision
    }
    return map[string]interface{} {
        "questModelId": questModelId,
        "questGroupName": questGroupName,
        "name": name,
        "description": description,
        "metadata": metadata,
        "contents": contents,
        "challengePeriodEventId": challengePeriodEventId,
        "firstCompleteAcquireActions": firstCompleteAcquireActions,
        "consumeActions": consumeActions,
        "failedAcquireActions": failedAcquireActions,
        "premiseQuestNames": premiseQuestNames,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p QuestModelMaster) Pointer() *QuestModelMaster {
    return &p
}

func CastQuestModelMasters(data []interface{}) []QuestModelMaster {
	v := make([]QuestModelMaster, 0)
	for _, d := range data {
		v = append(v, NewQuestModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestModelMastersFromDict(data []QuestModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CurrentQuestMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings *string `json:"settings"`
}

func NewCurrentQuestMasterFromJson(data string) CurrentQuestMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentQuestMasterFromDict(dict)
}

func NewCurrentQuestMasterFromDict(data map[string]interface{}) CurrentQuestMaster {
    return CurrentQuestMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentQuestMaster) ToDict() map[string]interface{} {
    
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

func (p CurrentQuestMaster) Pointer() *CurrentQuestMaster {
    return &p
}

func CastCurrentQuestMasters(data []interface{}) []CurrentQuestMaster {
	v := make([]CurrentQuestMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentQuestMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentQuestMastersFromDict(data []CurrentQuestMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Contents struct {
	Metadata *string `json:"metadata"`
	CompleteAcquireActions []AcquireAction `json:"completeAcquireActions"`
	Weight *int32 `json:"weight"`
}

func NewContentsFromJson(data string) Contents {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewContentsFromDict(dict)
}

func NewContentsFromDict(data map[string]interface{}) Contents {
    return Contents {
        Metadata: core.CastString(data["metadata"]),
        CompleteAcquireActions: CastAcquireActions(core.CastArray(data["completeAcquireActions"])),
        Weight: core.CastInt32(data["weight"]),
    }
}

func (p Contents) ToDict() map[string]interface{} {
    
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var completeAcquireActions []interface{}
    if p.CompleteAcquireActions != nil {
        completeAcquireActions = CastAcquireActionsFromDict(
            p.CompleteAcquireActions,
        )
    }
    var weight *int32
    if p.Weight != nil {
        weight = p.Weight
    }
    return map[string]interface{} {
        "metadata": metadata,
        "completeAcquireActions": completeAcquireActions,
        "weight": weight,
    }
}

func (p Contents) Pointer() *Contents {
    return &p
}

func CastContentses(data []interface{}) []Contents {
	v := make([]Contents, 0)
	for _, d := range data {
		v = append(v, NewContentsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastContentsesFromDict(data []Contents) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Reward struct {
	Action *string `json:"action"`
	Request *string `json:"request"`
	ItemId *string `json:"itemId"`
	Value *int32 `json:"value"`
}

func NewRewardFromJson(data string) Reward {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRewardFromDict(dict)
}

func NewRewardFromDict(data map[string]interface{}) Reward {
    return Reward {
        Action: core.CastString(data["action"]),
        Request: core.CastString(data["request"]),
        ItemId: core.CastString(data["itemId"]),
        Value: core.CastInt32(data["value"]),
    }
}

func (p Reward) ToDict() map[string]interface{} {
    
    var action *string
    if p.Action != nil {
        action = p.Action
    }
    var request *string
    if p.Request != nil {
        request = p.Request
    }
    var itemId *string
    if p.ItemId != nil {
        itemId = p.ItemId
    }
    var value *int32
    if p.Value != nil {
        value = p.Value
    }
    return map[string]interface{} {
        "action": action,
        "request": request,
        "itemId": itemId,
        "value": value,
    }
}

func (p Reward) Pointer() *Reward {
    return &p
}

func CastRewards(data []interface{}) []Reward {
	v := make([]Reward, 0)
	for _, d := range data {
		v = append(v, NewRewardFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRewardsFromDict(data []Reward) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Progress struct {
	ProgressId *string `json:"progressId"`
	UserId *string `json:"userId"`
	TransactionId *string `json:"transactionId"`
	QuestModelId *string `json:"questModelId"`
	RandomSeed *int64 `json:"randomSeed"`
	Rewards []Reward `json:"rewards"`
	Metadata *string `json:"metadata"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewProgressFromJson(data string) Progress {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewProgressFromDict(dict)
}

func NewProgressFromDict(data map[string]interface{}) Progress {
    return Progress {
        ProgressId: core.CastString(data["progressId"]),
        UserId: core.CastString(data["userId"]),
        TransactionId: core.CastString(data["transactionId"]),
        QuestModelId: core.CastString(data["questModelId"]),
        RandomSeed: core.CastInt64(data["randomSeed"]),
        Rewards: CastRewards(core.CastArray(data["rewards"])),
        Metadata: core.CastString(data["metadata"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Progress) ToDict() map[string]interface{} {
    
    var progressId *string
    if p.ProgressId != nil {
        progressId = p.ProgressId
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var transactionId *string
    if p.TransactionId != nil {
        transactionId = p.TransactionId
    }
    var questModelId *string
    if p.QuestModelId != nil {
        questModelId = p.QuestModelId
    }
    var randomSeed *int64
    if p.RandomSeed != nil {
        randomSeed = p.RandomSeed
    }
    var rewards []interface{}
    if p.Rewards != nil {
        rewards = CastRewardsFromDict(
            p.Rewards,
        )
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
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
    return map[string]interface{} {
        "progressId": progressId,
        "userId": userId,
        "transactionId": transactionId,
        "questModelId": questModelId,
        "randomSeed": randomSeed,
        "rewards": rewards,
        "metadata": metadata,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p Progress) Pointer() *Progress {
    return &p
}

func CastProgresses(data []interface{}) []Progress {
	v := make([]Progress, 0)
	for _, d := range data {
		v = append(v, NewProgressFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProgressesFromDict(data []Progress) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CompletedQuestList struct {
	CompletedQuestListId *string `json:"completedQuestListId"`
	UserId *string `json:"userId"`
	QuestGroupName *string `json:"questGroupName"`
	CompleteQuestNames []*string `json:"completeQuestNames"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewCompletedQuestListFromJson(data string) CompletedQuestList {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCompletedQuestListFromDict(dict)
}

func NewCompletedQuestListFromDict(data map[string]interface{}) CompletedQuestList {
    return CompletedQuestList {
        CompletedQuestListId: core.CastString(data["completedQuestListId"]),
        UserId: core.CastString(data["userId"]),
        QuestGroupName: core.CastString(data["questGroupName"]),
        CompleteQuestNames: core.CastStrings(core.CastArray(data["completeQuestNames"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p CompletedQuestList) ToDict() map[string]interface{} {
    
    var completedQuestListId *string
    if p.CompletedQuestListId != nil {
        completedQuestListId = p.CompletedQuestListId
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var questGroupName *string
    if p.QuestGroupName != nil {
        questGroupName = p.QuestGroupName
    }
    var completeQuestNames []interface{}
    if p.CompleteQuestNames != nil {
        completeQuestNames = core.CastStringsFromDict(
            p.CompleteQuestNames,
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
    var revision *int64
    if p.Revision != nil {
        revision = p.Revision
    }
    return map[string]interface{} {
        "completedQuestListId": completedQuestListId,
        "userId": userId,
        "questGroupName": questGroupName,
        "completeQuestNames": completeQuestNames,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p CompletedQuestList) Pointer() *CompletedQuestList {
    return &p
}

func CastCompletedQuestList(data []interface{}) []CompletedQuestList {
	v := make([]CompletedQuestList, 0)
	for _, d := range data {
		v = append(v, NewCompletedQuestListFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCompletedQuestListFromDict(data []CompletedQuestList) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type QuestGroupModel struct {
	QuestGroupModelId *string `json:"questGroupModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Quests []QuestModel `json:"quests"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
}

func NewQuestGroupModelFromJson(data string) QuestGroupModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewQuestGroupModelFromDict(dict)
}

func NewQuestGroupModelFromDict(data map[string]interface{}) QuestGroupModel {
    return QuestGroupModel {
        QuestGroupModelId: core.CastString(data["questGroupModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Quests: CastQuestModels(core.CastArray(data["quests"])),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
    }
}

func (p QuestGroupModel) ToDict() map[string]interface{} {
    
    var questGroupModelId *string
    if p.QuestGroupModelId != nil {
        questGroupModelId = p.QuestGroupModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var quests []interface{}
    if p.Quests != nil {
        quests = CastQuestModelsFromDict(
            p.Quests,
        )
    }
    var challengePeriodEventId *string
    if p.ChallengePeriodEventId != nil {
        challengePeriodEventId = p.ChallengePeriodEventId
    }
    return map[string]interface{} {
        "questGroupModelId": questGroupModelId,
        "name": name,
        "metadata": metadata,
        "quests": quests,
        "challengePeriodEventId": challengePeriodEventId,
    }
}

func (p QuestGroupModel) Pointer() *QuestGroupModel {
    return &p
}

func CastQuestGroupModels(data []interface{}) []QuestGroupModel {
	v := make([]QuestGroupModel, 0)
	for _, d := range data {
		v = append(v, NewQuestGroupModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestGroupModelsFromDict(data []QuestGroupModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type QuestModel struct {
	QuestModelId *string `json:"questModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Contents []Contents `json:"contents"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
	FirstCompleteAcquireActions []AcquireAction `json:"firstCompleteAcquireActions"`
	ConsumeActions []ConsumeAction `json:"consumeActions"`
	FailedAcquireActions []AcquireAction `json:"failedAcquireActions"`
	PremiseQuestNames []*string `json:"premiseQuestNames"`
}

func NewQuestModelFromJson(data string) QuestModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewQuestModelFromDict(dict)
}

func NewQuestModelFromDict(data map[string]interface{}) QuestModel {
    return QuestModel {
        QuestModelId: core.CastString(data["questModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Contents: CastContentses(core.CastArray(data["contents"])),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
        FirstCompleteAcquireActions: CastAcquireActions(core.CastArray(data["firstCompleteAcquireActions"])),
        ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
        FailedAcquireActions: CastAcquireActions(core.CastArray(data["failedAcquireActions"])),
        PremiseQuestNames: core.CastStrings(core.CastArray(data["premiseQuestNames"])),
    }
}

func (p QuestModel) ToDict() map[string]interface{} {
    
    var questModelId *string
    if p.QuestModelId != nil {
        questModelId = p.QuestModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var contents []interface{}
    if p.Contents != nil {
        contents = CastContentsesFromDict(
            p.Contents,
        )
    }
    var challengePeriodEventId *string
    if p.ChallengePeriodEventId != nil {
        challengePeriodEventId = p.ChallengePeriodEventId
    }
    var firstCompleteAcquireActions []interface{}
    if p.FirstCompleteAcquireActions != nil {
        firstCompleteAcquireActions = CastAcquireActionsFromDict(
            p.FirstCompleteAcquireActions,
        )
    }
    var consumeActions []interface{}
    if p.ConsumeActions != nil {
        consumeActions = CastConsumeActionsFromDict(
            p.ConsumeActions,
        )
    }
    var failedAcquireActions []interface{}
    if p.FailedAcquireActions != nil {
        failedAcquireActions = CastAcquireActionsFromDict(
            p.FailedAcquireActions,
        )
    }
    var premiseQuestNames []interface{}
    if p.PremiseQuestNames != nil {
        premiseQuestNames = core.CastStringsFromDict(
            p.PremiseQuestNames,
        )
    }
    return map[string]interface{} {
        "questModelId": questModelId,
        "name": name,
        "metadata": metadata,
        "contents": contents,
        "challengePeriodEventId": challengePeriodEventId,
        "firstCompleteAcquireActions": firstCompleteAcquireActions,
        "consumeActions": consumeActions,
        "failedAcquireActions": failedAcquireActions,
        "premiseQuestNames": premiseQuestNames,
    }
}

func (p QuestModel) Pointer() *QuestModel {
    return &p
}

func CastQuestModels(data []interface{}) []QuestModel {
	v := make([]QuestModel, 0)
	for _, d := range data {
		v = append(v, NewQuestModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestModelsFromDict(data []QuestModel) []interface{} {
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

type ConsumeAction struct {
	Action *string `json:"action"`
	Request *string `json:"request"`
}

func NewConsumeActionFromJson(data string) ConsumeAction {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConsumeActionFromDict(dict)
}

func NewConsumeActionFromDict(data map[string]interface{}) ConsumeAction {
    return ConsumeAction {
        Action: core.CastString(data["action"]),
        Request: core.CastString(data["request"]),
    }
}

func (p ConsumeAction) ToDict() map[string]interface{} {
    
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

func (p ConsumeAction) Pointer() *ConsumeAction {
    return &p
}

func CastConsumeActions(data []interface{}) []ConsumeAction {
	v := make([]ConsumeAction, 0)
	for _, d := range data {
		v = append(v, NewConsumeActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeActionsFromDict(data []ConsumeAction) []interface{} {
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