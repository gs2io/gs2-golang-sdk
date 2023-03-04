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

package matchmaking

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	EnableRating *bool `json:"enableRating"`
	CreateGatheringTriggerType *string `json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId *string `json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId *string `json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType *string `json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *string `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId *string `json:"completeMatchmakingTriggerScriptId"`
	ChangeRatingScript *ScriptSetting `json:"changeRatingScript"`
	JoinNotification *NotificationSetting `json:"joinNotification"`
	LeaveNotification *NotificationSetting `json:"leaveNotification"`
	CompleteNotification *NotificationSetting `json:"completeNotification"`
	ChangeRatingNotification *NotificationSetting `json:"changeRatingNotification"`
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
        EnableRating: core.CastBool(data["enableRating"]),
        CreateGatheringTriggerType: core.CastString(data["createGatheringTriggerType"]),
        CreateGatheringTriggerRealtimeNamespaceId: core.CastString(data["createGatheringTriggerRealtimeNamespaceId"]),
        CreateGatheringTriggerScriptId: core.CastString(data["createGatheringTriggerScriptId"]),
        CompleteMatchmakingTriggerType: core.CastString(data["completeMatchmakingTriggerType"]),
        CompleteMatchmakingTriggerRealtimeNamespaceId: core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"]),
        CompleteMatchmakingTriggerScriptId: core.CastString(data["completeMatchmakingTriggerScriptId"]),
        ChangeRatingScript: NewScriptSettingFromDict(core.CastMap(data["changeRatingScript"])).Pointer(),
        JoinNotification: NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
        LeaveNotification: NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
        CompleteNotification: NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
        ChangeRatingNotification: NewNotificationSettingFromDict(core.CastMap(data["changeRatingNotification"])).Pointer(),
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
    var enableRating *bool
    if p.EnableRating != nil {
        enableRating = p.EnableRating
    }
    var createGatheringTriggerType *string
    if p.CreateGatheringTriggerType != nil {
        createGatheringTriggerType = p.CreateGatheringTriggerType
    }
    var createGatheringTriggerRealtimeNamespaceId *string
    if p.CreateGatheringTriggerRealtimeNamespaceId != nil {
        createGatheringTriggerRealtimeNamespaceId = p.CreateGatheringTriggerRealtimeNamespaceId
    }
    var createGatheringTriggerScriptId *string
    if p.CreateGatheringTriggerScriptId != nil {
        createGatheringTriggerScriptId = p.CreateGatheringTriggerScriptId
    }
    var completeMatchmakingTriggerType *string
    if p.CompleteMatchmakingTriggerType != nil {
        completeMatchmakingTriggerType = p.CompleteMatchmakingTriggerType
    }
    var completeMatchmakingTriggerRealtimeNamespaceId *string
    if p.CompleteMatchmakingTriggerRealtimeNamespaceId != nil {
        completeMatchmakingTriggerRealtimeNamespaceId = p.CompleteMatchmakingTriggerRealtimeNamespaceId
    }
    var completeMatchmakingTriggerScriptId *string
    if p.CompleteMatchmakingTriggerScriptId != nil {
        completeMatchmakingTriggerScriptId = p.CompleteMatchmakingTriggerScriptId
    }
    var changeRatingScript map[string]interface{}
    if p.ChangeRatingScript != nil {
        changeRatingScript = p.ChangeRatingScript.ToDict()
    }
    var joinNotification map[string]interface{}
    if p.JoinNotification != nil {
        joinNotification = p.JoinNotification.ToDict()
    }
    var leaveNotification map[string]interface{}
    if p.LeaveNotification != nil {
        leaveNotification = p.LeaveNotification.ToDict()
    }
    var completeNotification map[string]interface{}
    if p.CompleteNotification != nil {
        completeNotification = p.CompleteNotification.ToDict()
    }
    var changeRatingNotification map[string]interface{}
    if p.ChangeRatingNotification != nil {
        changeRatingNotification = p.ChangeRatingNotification.ToDict()
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
        "enableRating": enableRating,
        "createGatheringTriggerType": createGatheringTriggerType,
        "createGatheringTriggerRealtimeNamespaceId": createGatheringTriggerRealtimeNamespaceId,
        "createGatheringTriggerScriptId": createGatheringTriggerScriptId,
        "completeMatchmakingTriggerType": completeMatchmakingTriggerType,
        "completeMatchmakingTriggerRealtimeNamespaceId": completeMatchmakingTriggerRealtimeNamespaceId,
        "completeMatchmakingTriggerScriptId": completeMatchmakingTriggerScriptId,
        "changeRatingScript": changeRatingScript,
        "joinNotification": joinNotification,
        "leaveNotification": leaveNotification,
        "completeNotification": completeNotification,
        "changeRatingNotification": changeRatingNotification,
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

type Gathering struct {
	GatheringId *string `json:"gatheringId"`
	Name *string `json:"name"`
	AttributeRanges []AttributeRange `json:"attributeRanges"`
	CapacityOfRoles []CapacityOfRole `json:"capacityOfRoles"`
	AllowUserIds []string `json:"allowUserIds"`
	Metadata *string `json:"metadata"`
	ExpiresAt *int64 `json:"expiresAt"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewGatheringFromJson(data string) Gathering {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGatheringFromDict(dict)
}

func NewGatheringFromDict(data map[string]interface{}) Gathering {
    return Gathering {
        GatheringId: core.CastString(data["gatheringId"]),
        Name: core.CastString(data["name"]),
        AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
        CapacityOfRoles: CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"])),
        AllowUserIds: core.CastStrings(core.CastArray(data["allowUserIds"])),
        Metadata: core.CastString(data["metadata"]),
        ExpiresAt: core.CastInt64(data["expiresAt"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Gathering) ToDict() map[string]interface{} {
    
    var gatheringId *string
    if p.GatheringId != nil {
        gatheringId = p.GatheringId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var attributeRanges []interface{}
    if p.AttributeRanges != nil {
        attributeRanges = CastAttributeRangesFromDict(
            p.AttributeRanges,
        )
    }
    var capacityOfRoles []interface{}
    if p.CapacityOfRoles != nil {
        capacityOfRoles = CastCapacityOfRolesFromDict(
            p.CapacityOfRoles,
        )
    }
    var allowUserIds []interface{}
    if p.AllowUserIds != nil {
        allowUserIds = core.CastStringsFromDict(
            p.AllowUserIds,
        )
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var expiresAt *int64
    if p.ExpiresAt != nil {
        expiresAt = p.ExpiresAt
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
        "gatheringId": gatheringId,
        "name": name,
        "attributeRanges": attributeRanges,
        "capacityOfRoles": capacityOfRoles,
        "allowUserIds": allowUserIds,
        "metadata": metadata,
        "expiresAt": expiresAt,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p Gathering) Pointer() *Gathering {
    return &p
}

func CastGatherings(data []interface{}) []Gathering {
	v := make([]Gathering, 0)
	for _, d := range data {
		v = append(v, NewGatheringFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGatheringsFromDict(data []Gathering) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type RatingModelMaster struct {
	RatingModelId *string `json:"ratingModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	Volatility *int32 `json:"volatility"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewRatingModelMasterFromJson(data string) RatingModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRatingModelMasterFromDict(dict)
}

func NewRatingModelMasterFromDict(data map[string]interface{}) RatingModelMaster {
    return RatingModelMaster {
        RatingModelId: core.CastString(data["ratingModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        Volatility: core.CastInt32(data["volatility"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p RatingModelMaster) ToDict() map[string]interface{} {
    
    var ratingModelId *string
    if p.RatingModelId != nil {
        ratingModelId = p.RatingModelId
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
    var volatility *int32
    if p.Volatility != nil {
        volatility = p.Volatility
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
        "ratingModelId": ratingModelId,
        "name": name,
        "metadata": metadata,
        "description": description,
        "volatility": volatility,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p RatingModelMaster) Pointer() *RatingModelMaster {
    return &p
}

func CastRatingModelMasters(data []interface{}) []RatingModelMaster {
	v := make([]RatingModelMaster, 0)
	for _, d := range data {
		v = append(v, NewRatingModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRatingModelMastersFromDict(data []RatingModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type RatingModel struct {
	RatingModelId *string `json:"ratingModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Volatility *int32 `json:"volatility"`
}

func NewRatingModelFromJson(data string) RatingModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRatingModelFromDict(dict)
}

func NewRatingModelFromDict(data map[string]interface{}) RatingModel {
    return RatingModel {
        RatingModelId: core.CastString(data["ratingModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Volatility: core.CastInt32(data["volatility"]),
    }
}

func (p RatingModel) ToDict() map[string]interface{} {
    
    var ratingModelId *string
    if p.RatingModelId != nil {
        ratingModelId = p.RatingModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var volatility *int32
    if p.Volatility != nil {
        volatility = p.Volatility
    }
    return map[string]interface{} {
        "ratingModelId": ratingModelId,
        "name": name,
        "metadata": metadata,
        "volatility": volatility,
    }
}

func (p RatingModel) Pointer() *RatingModel {
    return &p
}

func CastRatingModels(data []interface{}) []RatingModel {
	v := make([]RatingModel, 0)
	for _, d := range data {
		v = append(v, NewRatingModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRatingModelsFromDict(data []RatingModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CurrentRatingModelMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings *string `json:"settings"`
}

func NewCurrentRatingModelMasterFromJson(data string) CurrentRatingModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentRatingModelMasterFromDict(dict)
}

func NewCurrentRatingModelMasterFromDict(data map[string]interface{}) CurrentRatingModelMaster {
    return CurrentRatingModelMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentRatingModelMaster) ToDict() map[string]interface{} {
    
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

func (p CurrentRatingModelMaster) Pointer() *CurrentRatingModelMaster {
    return &p
}

func CastCurrentRatingModelMasters(data []interface{}) []CurrentRatingModelMaster {
	v := make([]CurrentRatingModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentRatingModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentRatingModelMastersFromDict(data []CurrentRatingModelMaster) []interface{} {
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
    
    var gatewayNamespaceId *string
    if p.GatewayNamespaceId != nil {
        gatewayNamespaceId = p.GatewayNamespaceId
    }
    var enableTransferMobileNotification *bool
    if p.EnableTransferMobileNotification != nil {
        enableTransferMobileNotification = p.EnableTransferMobileNotification
    }
    var sound *string
    if p.Sound != nil {
        sound = p.Sound
    }
    return map[string]interface{} {
        "gatewayNamespaceId": gatewayNamespaceId,
        "enableTransferMobileNotification": enableTransferMobileNotification,
        "sound": sound,
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

type AttributeRange struct {
	Name *string `json:"name"`
	Min *int32 `json:"min"`
	Max *int32 `json:"max"`
}

func NewAttributeRangeFromJson(data string) AttributeRange {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAttributeRangeFromDict(dict)
}

func NewAttributeRangeFromDict(data map[string]interface{}) AttributeRange {
    return AttributeRange {
        Name: core.CastString(data["name"]),
        Min: core.CastInt32(data["min"]),
        Max: core.CastInt32(data["max"]),
    }
}

func (p AttributeRange) ToDict() map[string]interface{} {
    
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var min *int32
    if p.Min != nil {
        min = p.Min
    }
    var max *int32
    if p.Max != nil {
        max = p.Max
    }
    return map[string]interface{} {
        "name": name,
        "min": min,
        "max": max,
    }
}

func (p AttributeRange) Pointer() *AttributeRange {
    return &p
}

func CastAttributeRanges(data []interface{}) []AttributeRange {
	v := make([]AttributeRange, 0)
	for _, d := range data {
		v = append(v, NewAttributeRangeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAttributeRangesFromDict(data []AttributeRange) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CapacityOfRole struct {
	RoleName *string `json:"roleName"`
	RoleAliases []string `json:"roleAliases"`
	Capacity *int32 `json:"capacity"`
	Participants []Player `json:"participants"`
}

func NewCapacityOfRoleFromJson(data string) CapacityOfRole {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCapacityOfRoleFromDict(dict)
}

func NewCapacityOfRoleFromDict(data map[string]interface{}) CapacityOfRole {
    return CapacityOfRole {
        RoleName: core.CastString(data["roleName"]),
        RoleAliases: core.CastStrings(core.CastArray(data["roleAliases"])),
        Capacity: core.CastInt32(data["capacity"]),
        Participants: CastPlayers(core.CastArray(data["participants"])),
    }
}

func (p CapacityOfRole) ToDict() map[string]interface{} {
    
    var roleName *string
    if p.RoleName != nil {
        roleName = p.RoleName
    }
    var roleAliases []interface{}
    if p.RoleAliases != nil {
        roleAliases = core.CastStringsFromDict(
            p.RoleAliases,
        )
    }
    var capacity *int32
    if p.Capacity != nil {
        capacity = p.Capacity
    }
    var participants []interface{}
    if p.Participants != nil {
        participants = CastPlayersFromDict(
            p.Participants,
        )
    }
    return map[string]interface{} {
        "roleName": roleName,
        "roleAliases": roleAliases,
        "capacity": capacity,
        "participants": participants,
    }
}

func (p CapacityOfRole) Pointer() *CapacityOfRole {
    return &p
}

func CastCapacityOfRoles(data []interface{}) []CapacityOfRole {
	v := make([]CapacityOfRole, 0)
	for _, d := range data {
		v = append(v, NewCapacityOfRoleFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCapacityOfRolesFromDict(data []CapacityOfRole) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Attribute struct {
	Name *string `json:"name"`
	Value *int32 `json:"value"`
}

func NewAttributeFromJson(data string) Attribute {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAttributeFromDict(dict)
}

func NewAttributeFromDict(data map[string]interface{}) Attribute {
    return Attribute {
        Name: core.CastString(data["name"]),
        Value: core.CastInt32(data["value"]),
    }
}

func (p Attribute) ToDict() map[string]interface{} {
    
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var value *int32
    if p.Value != nil {
        value = p.Value
    }
    return map[string]interface{} {
        "name": name,
        "value": value,
    }
}

func (p Attribute) Pointer() *Attribute {
    return &p
}

func CastAttributes(data []interface{}) []Attribute {
	v := make([]Attribute, 0)
	for _, d := range data {
		v = append(v, NewAttributeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAttributesFromDict(data []Attribute) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Player struct {
	UserId *string `json:"userId"`
	Attributes []Attribute `json:"attributes"`
	RoleName *string `json:"roleName"`
	DenyUserIds []string `json:"denyUserIds"`
}

func NewPlayerFromJson(data string) Player {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPlayerFromDict(dict)
}

func NewPlayerFromDict(data map[string]interface{}) Player {
    return Player {
        UserId: core.CastString(data["userId"]),
        Attributes: CastAttributes(core.CastArray(data["attributes"])),
        RoleName: core.CastString(data["roleName"]),
        DenyUserIds: core.CastStrings(core.CastArray(data["denyUserIds"])),
    }
}

func (p Player) ToDict() map[string]interface{} {
    
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var attributes []interface{}
    if p.Attributes != nil {
        attributes = CastAttributesFromDict(
            p.Attributes,
        )
    }
    var roleName *string
    if p.RoleName != nil {
        roleName = p.RoleName
    }
    var denyUserIds []interface{}
    if p.DenyUserIds != nil {
        denyUserIds = core.CastStringsFromDict(
            p.DenyUserIds,
        )
    }
    return map[string]interface{} {
        "userId": userId,
        "attributes": attributes,
        "roleName": roleName,
        "denyUserIds": denyUserIds,
    }
}

func (p Player) Pointer() *Player {
    return &p
}

func CastPlayers(data []interface{}) []Player {
	v := make([]Player, 0)
	for _, d := range data {
		v = append(v, NewPlayerFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPlayersFromDict(data []Player) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Rating struct {
	RatingId *string `json:"ratingId"`
	Name *string `json:"name"`
	UserId *string `json:"userId"`
	RateValue *float32 `json:"rateValue"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewRatingFromJson(data string) Rating {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRatingFromDict(dict)
}

func NewRatingFromDict(data map[string]interface{}) Rating {
    return Rating {
        RatingId: core.CastString(data["ratingId"]),
        Name: core.CastString(data["name"]),
        UserId: core.CastString(data["userId"]),
        RateValue: core.CastFloat32(data["rateValue"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Rating) ToDict() map[string]interface{} {
    
    var ratingId *string
    if p.RatingId != nil {
        ratingId = p.RatingId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var rateValue *float32
    if p.RateValue != nil {
        rateValue = p.RateValue
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
        "ratingId": ratingId,
        "name": name,
        "userId": userId,
        "rateValue": rateValue,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p Rating) Pointer() *Rating {
    return &p
}

func CastRatings(data []interface{}) []Rating {
	v := make([]Rating, 0)
	for _, d := range data {
		v = append(v, NewRatingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRatingsFromDict(data []Rating) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type GameResult struct {
	Rank *int32 `json:"rank"`
	UserId *string `json:"userId"`
}

func NewGameResultFromJson(data string) GameResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGameResultFromDict(dict)
}

func NewGameResultFromDict(data map[string]interface{}) GameResult {
    return GameResult {
        Rank: core.CastInt32(data["rank"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p GameResult) ToDict() map[string]interface{} {
    
    var rank *int32
    if p.Rank != nil {
        rank = p.Rank
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    return map[string]interface{} {
        "rank": rank,
        "userId": userId,
    }
}

func (p GameResult) Pointer() *GameResult {
    return &p
}

func CastGameResults(data []interface{}) []GameResult {
	v := make([]GameResult, 0)
	for _, d := range data {
		v = append(v, NewGameResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGameResultsFromDict(data []GameResult) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Ballot struct {
	UserId *string `json:"userId"`
	RatingName *string `json:"ratingName"`
	GatheringName *string `json:"gatheringName"`
	NumberOfPlayer *int32 `json:"numberOfPlayer"`
}

func NewBallotFromJson(data string) Ballot {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBallotFromDict(dict)
}

func NewBallotFromDict(data map[string]interface{}) Ballot {
    return Ballot {
        UserId: core.CastString(data["userId"]),
        RatingName: core.CastString(data["ratingName"]),
        GatheringName: core.CastString(data["gatheringName"]),
        NumberOfPlayer: core.CastInt32(data["numberOfPlayer"]),
    }
}

func (p Ballot) ToDict() map[string]interface{} {
    
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var ratingName *string
    if p.RatingName != nil {
        ratingName = p.RatingName
    }
    var gatheringName *string
    if p.GatheringName != nil {
        gatheringName = p.GatheringName
    }
    var numberOfPlayer *int32
    if p.NumberOfPlayer != nil {
        numberOfPlayer = p.NumberOfPlayer
    }
    return map[string]interface{} {
        "userId": userId,
        "ratingName": ratingName,
        "gatheringName": gatheringName,
        "numberOfPlayer": numberOfPlayer,
    }
}

func (p Ballot) Pointer() *Ballot {
    return &p
}

func CastBallots(data []interface{}) []Ballot {
	v := make([]Ballot, 0)
	for _, d := range data {
		v = append(v, NewBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBallotsFromDict(data []Ballot) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type SignedBallot struct {
	Body *string `json:"body"`
	Signature *string `json:"signature"`
}

func NewSignedBallotFromJson(data string) SignedBallot {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSignedBallotFromDict(dict)
}

func NewSignedBallotFromDict(data map[string]interface{}) SignedBallot {
    return SignedBallot {
        Body: core.CastString(data["body"]),
        Signature: core.CastString(data["signature"]),
    }
}

func (p SignedBallot) ToDict() map[string]interface{} {
    
    var body *string
    if p.Body != nil {
        body = p.Body
    }
    var signature *string
    if p.Signature != nil {
        signature = p.Signature
    }
    return map[string]interface{} {
        "body": body,
        "signature": signature,
    }
}

func (p SignedBallot) Pointer() *SignedBallot {
    return &p
}

func CastSignedBallots(data []interface{}) []SignedBallot {
	v := make([]SignedBallot, 0)
	for _, d := range data {
		v = append(v, NewSignedBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSignedBallotsFromDict(data []SignedBallot) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type WrittenBallot struct {
	Ballot *Ballot `json:"ballot"`
	GameResults []GameResult `json:"gameResults"`
}

func NewWrittenBallotFromJson(data string) WrittenBallot {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewWrittenBallotFromDict(dict)
}

func NewWrittenBallotFromDict(data map[string]interface{}) WrittenBallot {
    return WrittenBallot {
        Ballot: NewBallotFromDict(core.CastMap(data["ballot"])).Pointer(),
        GameResults: CastGameResults(core.CastArray(data["gameResults"])),
    }
}

func (p WrittenBallot) ToDict() map[string]interface{} {
    
    var ballot map[string]interface{}
    if p.Ballot != nil {
        ballot = p.Ballot.ToDict()
    }
    var gameResults []interface{}
    if p.GameResults != nil {
        gameResults = CastGameResultsFromDict(
            p.GameResults,
        )
    }
    return map[string]interface{} {
        "ballot": ballot,
        "gameResults": gameResults,
    }
}

func (p WrittenBallot) Pointer() *WrittenBallot {
    return &p
}

func CastWrittenBallots(data []interface{}) []WrittenBallot {
	v := make([]WrittenBallot, 0)
	for _, d := range data {
		v = append(v, NewWrittenBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastWrittenBallotsFromDict(data []WrittenBallot) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Vote struct {
	VoteId *string `json:"voteId"`
	RatingName *string `json:"ratingName"`
	GatheringName *string `json:"gatheringName"`
	WrittenBallots []WrittenBallot `json:"writtenBallots"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewVoteFromJson(data string) Vote {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVoteFromDict(dict)
}

func NewVoteFromDict(data map[string]interface{}) Vote {
    return Vote {
        VoteId: core.CastString(data["voteId"]),
        RatingName: core.CastString(data["ratingName"]),
        GatheringName: core.CastString(data["gatheringName"]),
        WrittenBallots: CastWrittenBallots(core.CastArray(data["writtenBallots"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Vote) ToDict() map[string]interface{} {
    
    var voteId *string
    if p.VoteId != nil {
        voteId = p.VoteId
    }
    var ratingName *string
    if p.RatingName != nil {
        ratingName = p.RatingName
    }
    var gatheringName *string
    if p.GatheringName != nil {
        gatheringName = p.GatheringName
    }
    var writtenBallots []interface{}
    if p.WrittenBallots != nil {
        writtenBallots = CastWrittenBallotsFromDict(
            p.WrittenBallots,
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
        "voteId": voteId,
        "ratingName": ratingName,
        "gatheringName": gatheringName,
        "writtenBallots": writtenBallots,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p Vote) Pointer() *Vote {
    return &p
}

func CastVotes(data []interface{}) []Vote {
	v := make([]Vote, 0)
	for _, d := range data {
		v = append(v, NewVoteFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVotesFromDict(data []Vote) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type TimeSpan struct {
	Days *int32 `json:"days"`
	Hours *int32 `json:"hours"`
	Minutes *int32 `json:"minutes"`
}

func NewTimeSpanFromJson(data string) TimeSpan {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewTimeSpanFromDict(dict)
}

func NewTimeSpanFromDict(data map[string]interface{}) TimeSpan {
    return TimeSpan {
        Days: core.CastInt32(data["days"]),
        Hours: core.CastInt32(data["hours"]),
        Minutes: core.CastInt32(data["minutes"]),
    }
}

func (p TimeSpan) ToDict() map[string]interface{} {
    
    var days *int32
    if p.Days != nil {
        days = p.Days
    }
    var hours *int32
    if p.Hours != nil {
        hours = p.Hours
    }
    var minutes *int32
    if p.Minutes != nil {
        minutes = p.Minutes
    }
    return map[string]interface{} {
        "days": days,
        "hours": hours,
        "minutes": minutes,
    }
}

func (p TimeSpan) Pointer() *TimeSpan {
    return &p
}

func CastTimeSpans(data []interface{}) []TimeSpan {
	v := make([]TimeSpan, 0)
	for _, d := range data {
		v = append(v, NewTimeSpanFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTimeSpansFromDict(data []TimeSpan) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}