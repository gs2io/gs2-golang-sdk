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

package mission

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Complete struct {
	CompleteId *string `json:"completeId"`
	UserId *string `json:"userId"`
	MissionGroupName *string `json:"missionGroupName"`
	CompletedMissionTaskNames []string `json:"completedMissionTaskNames"`
	ReceivedMissionTaskNames []string `json:"receivedMissionTaskNames"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewCompleteFromJson(data string) Complete {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCompleteFromDict(dict)
}

func NewCompleteFromDict(data map[string]interface{}) Complete {
    return Complete {
        CompleteId: core.CastString(data["completeId"]),
        UserId: core.CastString(data["userId"]),
        MissionGroupName: core.CastString(data["missionGroupName"]),
        CompletedMissionTaskNames: core.CastStrings(core.CastArray(data["completedMissionTaskNames"])),
        ReceivedMissionTaskNames: core.CastStrings(core.CastArray(data["receivedMissionTaskNames"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Complete) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "completeId": p.CompleteId,
        "userId": p.UserId,
        "missionGroupName": p.MissionGroupName,
        "completedMissionTaskNames": core.CastStringsFromDict(
        p.CompletedMissionTaskNames,
    ),
        "receivedMissionTaskNames": core.CastStringsFromDict(
        p.ReceivedMissionTaskNames,
    ),
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
    }
}

func (p Complete) Pointer() *Complete {
    return &p
}

func CastCompletes(data []interface{}) []Complete {
	v := make([]Complete, 0)
	for _, d := range data {
		v = append(v, NewCompleteFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCompletesFromDict(data []Complete) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type NotificationSetting struct {
	GatewayNamespaceId *string `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool `json:"enableTransferMobileNotification"`
	Sound *string `json:"sound"`
}

func NewNotificationSettingFromJson(data string) NotificationSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNotificationSettingFromDict(dict)
}

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
    return NotificationSetting {
        GatewayNamespaceId: core.CastString(data["gatewayNamespaceId"]),
        EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
        Sound: core.CastString(data["sound"]),
    }
}

func (p NotificationSetting) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "gatewayNamespaceId": p.GatewayNamespaceId,
        "enableTransferMobileNotification": p.EnableTransferMobileNotification,
        "sound": p.Sound,
    }
}

func (p NotificationSetting) Pointer() *NotificationSetting {
    return &p
}

func CastNotificationSettings(data []interface{}) []NotificationSetting {
	v := make([]NotificationSetting, 0)
	for _, d := range data {
		v = append(v, NewNotificationSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNotificationSettingsFromDict(data []NotificationSetting) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CounterModelMaster struct {
	CounterId *string `json:"counterId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	Scopes []CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewCounterModelMasterFromJson(data string) CounterModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCounterModelMasterFromDict(dict)
}

func NewCounterModelMasterFromDict(data map[string]interface{}) CounterModelMaster {
    return CounterModelMaster {
        CounterId: core.CastString(data["counterId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        Scopes: CastCounterScopeModels(core.CastArray(data["scopes"])),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p CounterModelMaster) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "counterId": p.CounterId,
        "name": p.Name,
        "metadata": p.Metadata,
        "description": p.Description,
        "scopes": CastCounterScopeModelsFromDict(
        p.Scopes,
    ),
        "challengePeriodEventId": p.ChallengePeriodEventId,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
    }
}

func (p CounterModelMaster) Pointer() *CounterModelMaster {
    return &p
}

func CastCounterModelMasters(data []interface{}) []CounterModelMaster {
	v := make([]CounterModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCounterModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCounterModelMastersFromDict(data []CounterModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CounterScopeModel struct {
	ResetType *string `json:"resetType"`
	ResetDayOfMonth *int32 `json:"resetDayOfMonth"`
	ResetDayOfWeek *string `json:"resetDayOfWeek"`
	ResetHour *int32 `json:"resetHour"`
}

func NewCounterScopeModelFromJson(data string) CounterScopeModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCounterScopeModelFromDict(dict)
}

func NewCounterScopeModelFromDict(data map[string]interface{}) CounterScopeModel {
    return CounterScopeModel {
        ResetType: core.CastString(data["resetType"]),
        ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
        ResetDayOfWeek: core.CastString(data["resetDayOfWeek"]),
        ResetHour: core.CastInt32(data["resetHour"]),
    }
}

func (p CounterScopeModel) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "resetType": p.ResetType,
        "resetDayOfMonth": p.ResetDayOfMonth,
        "resetDayOfWeek": p.ResetDayOfWeek,
        "resetHour": p.ResetHour,
    }
}

func (p CounterScopeModel) Pointer() *CounterScopeModel {
    return &p
}

func CastCounterScopeModels(data []interface{}) []CounterScopeModel {
	v := make([]CounterScopeModel, 0)
	for _, d := range data {
		v = append(v, NewCounterScopeModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCounterScopeModelsFromDict(data []CounterScopeModel) []interface{} {
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
    return map[string]interface{} {
        "key": p.Key,
        "value": p.Value,
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
    return map[string]interface{} {
        "action": p.Action,
        "request": p.Request,
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

type MissionGroupModelMaster struct {
	MissionGroupId *string `json:"missionGroupId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	ResetType *string `json:"resetType"`
	ResetDayOfMonth *int32 `json:"resetDayOfMonth"`
	ResetDayOfWeek *string `json:"resetDayOfWeek"`
	ResetHour *int32 `json:"resetHour"`
	CompleteNotificationNamespaceId *string `json:"completeNotificationNamespaceId"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewMissionGroupModelMasterFromJson(data string) MissionGroupModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMissionGroupModelMasterFromDict(dict)
}

func NewMissionGroupModelMasterFromDict(data map[string]interface{}) MissionGroupModelMaster {
    return MissionGroupModelMaster {
        MissionGroupId: core.CastString(data["missionGroupId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        ResetType: core.CastString(data["resetType"]),
        ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
        ResetDayOfWeek: core.CastString(data["resetDayOfWeek"]),
        ResetHour: core.CastInt32(data["resetHour"]),
        CompleteNotificationNamespaceId: core.CastString(data["completeNotificationNamespaceId"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p MissionGroupModelMaster) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "missionGroupId": p.MissionGroupId,
        "name": p.Name,
        "metadata": p.Metadata,
        "description": p.Description,
        "resetType": p.ResetType,
        "resetDayOfMonth": p.ResetDayOfMonth,
        "resetDayOfWeek": p.ResetDayOfWeek,
        "resetHour": p.ResetHour,
        "completeNotificationNamespaceId": p.CompleteNotificationNamespaceId,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
    }
}

func (p MissionGroupModelMaster) Pointer() *MissionGroupModelMaster {
    return &p
}

func CastMissionGroupModelMasters(data []interface{}) []MissionGroupModelMaster {
	v := make([]MissionGroupModelMaster, 0)
	for _, d := range data {
		v = append(v, NewMissionGroupModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionGroupModelMastersFromDict(data []MissionGroupModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	MissionCompleteScript *ScriptSetting `json:"missionCompleteScript"`
	CounterIncrementScript *ScriptSetting `json:"counterIncrementScript"`
	ReceiveRewardsScript *ScriptSetting `json:"receiveRewardsScript"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
	KeyId *string `json:"keyId"`
	CompleteNotification *NotificationSetting `json:"completeNotification"`
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
        MissionCompleteScript: NewScriptSettingFromDict(core.CastMap(data["missionCompleteScript"])).Pointer(),
        CounterIncrementScript: NewScriptSettingFromDict(core.CastMap(data["counterIncrementScript"])).Pointer(),
        ReceiveRewardsScript: NewScriptSettingFromDict(core.CastMap(data["receiveRewardsScript"])).Pointer(),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
        CompleteNotification: NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
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
        "missionCompleteScript": p.MissionCompleteScript.ToDict(),
        "counterIncrementScript": p.CounterIncrementScript.ToDict(),
        "receiveRewardsScript": p.ReceiveRewardsScript.ToDict(),
        "queueNamespaceId": p.QueueNamespaceId,
        "keyId": p.KeyId,
        "completeNotification": p.CompleteNotification.ToDict(),
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

type Counter struct {
	CounterId *string `json:"counterId"`
	UserId *string `json:"userId"`
	Name *string `json:"name"`
	Values []ScopedValue `json:"values"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewCounterFromJson(data string) Counter {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCounterFromDict(dict)
}

func NewCounterFromDict(data map[string]interface{}) Counter {
    return Counter {
        CounterId: core.CastString(data["counterId"]),
        UserId: core.CastString(data["userId"]),
        Name: core.CastString(data["name"]),
        Values: CastScopedValues(core.CastArray(data["values"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Counter) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "counterId": p.CounterId,
        "userId": p.UserId,
        "name": p.Name,
        "values": CastScopedValuesFromDict(
        p.Values,
    ),
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
    }
}

func (p Counter) Pointer() *Counter {
    return &p
}

func CastCounters(data []interface{}) []Counter {
	v := make([]Counter, 0)
	for _, d := range data {
		v = append(v, NewCounterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCountersFromDict(data []Counter) []interface{} {
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
    return map[string]interface{} {
        "apiKeyId": p.ApiKeyId,
        "repositoryName": p.RepositoryName,
        "sourcePath": p.SourcePath,
        "referenceType": p.ReferenceType,
        "commitHash": p.CommitHash,
        "branchName": p.BranchName,
        "tagName": p.TagName,
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

type CurrentMissionMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings *string `json:"settings"`
}

func NewCurrentMissionMasterFromJson(data string) CurrentMissionMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentMissionMasterFromDict(dict)
}

func NewCurrentMissionMasterFromDict(data map[string]interface{}) CurrentMissionMaster {
    return CurrentMissionMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentMissionMaster) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceId": p.NamespaceId,
        "settings": p.Settings,
    }
}

func (p CurrentMissionMaster) Pointer() *CurrentMissionMaster {
    return &p
}

func CastCurrentMissionMasters(data []interface{}) []CurrentMissionMaster {
	v := make([]CurrentMissionMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentMissionMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentMissionMastersFromDict(data []CurrentMissionMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CounterModel struct {
	CounterId *string `json:"counterId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Scopes []CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
}

func NewCounterModelFromJson(data string) CounterModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCounterModelFromDict(dict)
}

func NewCounterModelFromDict(data map[string]interface{}) CounterModel {
    return CounterModel {
        CounterId: core.CastString(data["counterId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Scopes: CastCounterScopeModels(core.CastArray(data["scopes"])),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
    }
}

func (p CounterModel) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "counterId": p.CounterId,
        "name": p.Name,
        "metadata": p.Metadata,
        "scopes": CastCounterScopeModelsFromDict(
        p.Scopes,
    ),
        "challengePeriodEventId": p.ChallengePeriodEventId,
    }
}

func (p CounterModel) Pointer() *CounterModel {
    return &p
}

func CastCounterModels(data []interface{}) []CounterModel {
	v := make([]CounterModel, 0)
	for _, d := range data {
		v = append(v, NewCounterModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCounterModelsFromDict(data []CounterModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type MissionGroupModel struct {
	MissionGroupId *string `json:"missionGroupId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Tasks []MissionTaskModel `json:"tasks"`
	ResetType *string `json:"resetType"`
	ResetDayOfMonth *int32 `json:"resetDayOfMonth"`
	ResetDayOfWeek *string `json:"resetDayOfWeek"`
	ResetHour *int32 `json:"resetHour"`
	CompleteNotificationNamespaceId *string `json:"completeNotificationNamespaceId"`
}

func NewMissionGroupModelFromJson(data string) MissionGroupModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMissionGroupModelFromDict(dict)
}

func NewMissionGroupModelFromDict(data map[string]interface{}) MissionGroupModel {
    return MissionGroupModel {
        MissionGroupId: core.CastString(data["missionGroupId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Tasks: CastMissionTaskModels(core.CastArray(data["tasks"])),
        ResetType: core.CastString(data["resetType"]),
        ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
        ResetDayOfWeek: core.CastString(data["resetDayOfWeek"]),
        ResetHour: core.CastInt32(data["resetHour"]),
        CompleteNotificationNamespaceId: core.CastString(data["completeNotificationNamespaceId"]),
    }
}

func (p MissionGroupModel) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "missionGroupId": p.MissionGroupId,
        "name": p.Name,
        "metadata": p.Metadata,
        "tasks": CastMissionTaskModelsFromDict(
        p.Tasks,
    ),
        "resetType": p.ResetType,
        "resetDayOfMonth": p.ResetDayOfMonth,
        "resetDayOfWeek": p.ResetDayOfWeek,
        "resetHour": p.ResetHour,
        "completeNotificationNamespaceId": p.CompleteNotificationNamespaceId,
    }
}

func (p MissionGroupModel) Pointer() *MissionGroupModel {
    return &p
}

func CastMissionGroupModels(data []interface{}) []MissionGroupModel {
	v := make([]MissionGroupModel, 0)
	for _, d := range data {
		v = append(v, NewMissionGroupModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionGroupModelsFromDict(data []MissionGroupModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type MissionTaskModel struct {
	MissionTaskId *string `json:"missionTaskId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	CounterName *string `json:"counterName"`
	TargetValue *int64 `json:"targetValue"`
	CompleteAcquireActions []AcquireAction `json:"completeAcquireActions"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
	PremiseMissionTaskName *string `json:"premiseMissionTaskName"`
}

func NewMissionTaskModelFromJson(data string) MissionTaskModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMissionTaskModelFromDict(dict)
}

func NewMissionTaskModelFromDict(data map[string]interface{}) MissionTaskModel {
    return MissionTaskModel {
        MissionTaskId: core.CastString(data["missionTaskId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        CounterName: core.CastString(data["counterName"]),
        TargetValue: core.CastInt64(data["targetValue"]),
        CompleteAcquireActions: CastAcquireActions(core.CastArray(data["completeAcquireActions"])),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
        PremiseMissionTaskName: core.CastString(data["premiseMissionTaskName"]),
    }
}

func (p MissionTaskModel) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "missionTaskId": p.MissionTaskId,
        "name": p.Name,
        "metadata": p.Metadata,
        "counterName": p.CounterName,
        "targetValue": p.TargetValue,
        "completeAcquireActions": CastAcquireActionsFromDict(
        p.CompleteAcquireActions,
    ),
        "challengePeriodEventId": p.ChallengePeriodEventId,
        "premiseMissionTaskName": p.PremiseMissionTaskName,
    }
}

func (p MissionTaskModel) Pointer() *MissionTaskModel {
    return &p
}

func CastMissionTaskModels(data []interface{}) []MissionTaskModel {
	v := make([]MissionTaskModel, 0)
	for _, d := range data {
		v = append(v, NewMissionTaskModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionTaskModelsFromDict(data []MissionTaskModel) []interface{} {
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

type ScopedValue struct {
	ResetType *string `json:"resetType"`
	Value *int64 `json:"value"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewScopedValueFromJson(data string) ScopedValue {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewScopedValueFromDict(dict)
}

func NewScopedValueFromDict(data map[string]interface{}) ScopedValue {
    return ScopedValue {
        ResetType: core.CastString(data["resetType"]),
        Value: core.CastInt64(data["value"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p ScopedValue) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "resetType": p.ResetType,
        "value": p.Value,
        "updatedAt": p.UpdatedAt,
    }
}

func (p ScopedValue) Pointer() *ScopedValue {
    return &p
}

func CastScopedValues(data []interface{}) []ScopedValue {
	v := make([]ScopedValue, 0)
	for _, d := range data {
		v = append(v, NewScopedValueFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScopedValuesFromDict(data []ScopedValue) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type MissionTaskModelMaster struct {
	MissionTaskId *string `json:"missionTaskId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	CounterName *string `json:"counterName"`
	TargetValue *int64 `json:"targetValue"`
	CompleteAcquireActions []AcquireAction `json:"completeAcquireActions"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
	PremiseMissionTaskName *string `json:"premiseMissionTaskName"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewMissionTaskModelMasterFromJson(data string) MissionTaskModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMissionTaskModelMasterFromDict(dict)
}

func NewMissionTaskModelMasterFromDict(data map[string]interface{}) MissionTaskModelMaster {
    return MissionTaskModelMaster {
        MissionTaskId: core.CastString(data["missionTaskId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        CounterName: core.CastString(data["counterName"]),
        TargetValue: core.CastInt64(data["targetValue"]),
        CompleteAcquireActions: CastAcquireActions(core.CastArray(data["completeAcquireActions"])),
        ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
        PremiseMissionTaskName: core.CastString(data["premiseMissionTaskName"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p MissionTaskModelMaster) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "missionTaskId": p.MissionTaskId,
        "name": p.Name,
        "metadata": p.Metadata,
        "description": p.Description,
        "counterName": p.CounterName,
        "targetValue": p.TargetValue,
        "completeAcquireActions": CastAcquireActionsFromDict(
        p.CompleteAcquireActions,
    ),
        "challengePeriodEventId": p.ChallengePeriodEventId,
        "premiseMissionTaskName": p.PremiseMissionTaskName,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
    }
}

func (p MissionTaskModelMaster) Pointer() *MissionTaskModelMaster {
    return &p
}

func CastMissionTaskModelMasters(data []interface{}) []MissionTaskModelMaster {
	v := make([]MissionTaskModelMaster, 0)
	for _, d := range data {
		v = append(v, NewMissionTaskModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionTaskModelMastersFromDict(data []MissionTaskModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}