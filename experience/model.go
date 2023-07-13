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

package experience

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	ExperienceCapScriptId *string `json:"experienceCapScriptId"`
	ChangeExperienceScript *ScriptSetting `json:"changeExperienceScript"`
	ChangeRankScript *ScriptSetting `json:"changeRankScript"`
	ChangeRankCapScript *ScriptSetting `json:"changeRankCapScript"`
	OverflowExperienceScript *ScriptSetting `json:"overflowExperienceScript"`
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
        ExperienceCapScriptId: core.CastString(data["experienceCapScriptId"]),
        ChangeExperienceScript: NewScriptSettingFromDict(core.CastMap(data["changeExperienceScript"])).Pointer(),
        ChangeRankScript: NewScriptSettingFromDict(core.CastMap(data["changeRankScript"])).Pointer(),
        ChangeRankCapScript: NewScriptSettingFromDict(core.CastMap(data["changeRankCapScript"])).Pointer(),
        OverflowExperienceScript: NewScriptSettingFromDict(core.CastMap(data["overflowExperienceScript"])).Pointer(),
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
    var experienceCapScriptId *string
    if p.ExperienceCapScriptId != nil {
        experienceCapScriptId = p.ExperienceCapScriptId
    }
    var changeExperienceScript map[string]interface{}
    if p.ChangeExperienceScript != nil {
        changeExperienceScript = p.ChangeExperienceScript.ToDict()
    }
    var changeRankScript map[string]interface{}
    if p.ChangeRankScript != nil {
        changeRankScript = p.ChangeRankScript.ToDict()
    }
    var changeRankCapScript map[string]interface{}
    if p.ChangeRankCapScript != nil {
        changeRankCapScript = p.ChangeRankCapScript.ToDict()
    }
    var overflowExperienceScript map[string]interface{}
    if p.OverflowExperienceScript != nil {
        overflowExperienceScript = p.OverflowExperienceScript.ToDict()
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
        "experienceCapScriptId": experienceCapScriptId,
        "changeExperienceScript": changeExperienceScript,
        "changeRankScript": changeRankScript,
        "changeRankCapScript": changeRankCapScript,
        "overflowExperienceScript": overflowExperienceScript,
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

type ExperienceModelMaster struct {
	ExperienceModelId *string `json:"experienceModelId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	DefaultExperience *int64 `json:"defaultExperience"`
	DefaultRankCap *int64 `json:"defaultRankCap"`
	MaxRankCap *int64 `json:"maxRankCap"`
	RankThresholdName *string `json:"rankThresholdName"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewExperienceModelMasterFromJson(data string) ExperienceModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExperienceModelMasterFromDict(dict)
}

func NewExperienceModelMasterFromDict(data map[string]interface{}) ExperienceModelMaster {
    return ExperienceModelMaster {
        ExperienceModelId: core.CastString(data["experienceModelId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        DefaultExperience: core.CastInt64(data["defaultExperience"]),
        DefaultRankCap: core.CastInt64(data["defaultRankCap"]),
        MaxRankCap: core.CastInt64(data["maxRankCap"]),
        RankThresholdName: core.CastString(data["rankThresholdName"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p ExperienceModelMaster) ToDict() map[string]interface{} {
    
    var experienceModelId *string
    if p.ExperienceModelId != nil {
        experienceModelId = p.ExperienceModelId
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
    var defaultExperience *int64
    if p.DefaultExperience != nil {
        defaultExperience = p.DefaultExperience
    }
    var defaultRankCap *int64
    if p.DefaultRankCap != nil {
        defaultRankCap = p.DefaultRankCap
    }
    var maxRankCap *int64
    if p.MaxRankCap != nil {
        maxRankCap = p.MaxRankCap
    }
    var rankThresholdName *string
    if p.RankThresholdName != nil {
        rankThresholdName = p.RankThresholdName
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
        "experienceModelId": experienceModelId,
        "name": name,
        "description": description,
        "metadata": metadata,
        "defaultExperience": defaultExperience,
        "defaultRankCap": defaultRankCap,
        "maxRankCap": maxRankCap,
        "rankThresholdName": rankThresholdName,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p ExperienceModelMaster) Pointer() *ExperienceModelMaster {
    return &p
}

func CastExperienceModelMasters(data []interface{}) []ExperienceModelMaster {
	v := make([]ExperienceModelMaster, 0)
	for _, d := range data {
		v = append(v, NewExperienceModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceModelMastersFromDict(data []ExperienceModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type ExperienceModel struct {
	ExperienceModelId *string `json:"experienceModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	DefaultExperience *int64 `json:"defaultExperience"`
	DefaultRankCap *int64 `json:"defaultRankCap"`
	MaxRankCap *int64 `json:"maxRankCap"`
	RankThreshold *Threshold `json:"rankThreshold"`
}

func NewExperienceModelFromJson(data string) ExperienceModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExperienceModelFromDict(dict)
}

func NewExperienceModelFromDict(data map[string]interface{}) ExperienceModel {
    return ExperienceModel {
        ExperienceModelId: core.CastString(data["experienceModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        DefaultExperience: core.CastInt64(data["defaultExperience"]),
        DefaultRankCap: core.CastInt64(data["defaultRankCap"]),
        MaxRankCap: core.CastInt64(data["maxRankCap"]),
        RankThreshold: NewThresholdFromDict(core.CastMap(data["rankThreshold"])).Pointer(),
    }
}

func (p ExperienceModel) ToDict() map[string]interface{} {
    
    var experienceModelId *string
    if p.ExperienceModelId != nil {
        experienceModelId = p.ExperienceModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var defaultExperience *int64
    if p.DefaultExperience != nil {
        defaultExperience = p.DefaultExperience
    }
    var defaultRankCap *int64
    if p.DefaultRankCap != nil {
        defaultRankCap = p.DefaultRankCap
    }
    var maxRankCap *int64
    if p.MaxRankCap != nil {
        maxRankCap = p.MaxRankCap
    }
    var rankThreshold map[string]interface{}
    if p.RankThreshold != nil {
        rankThreshold = p.RankThreshold.ToDict()
    }
    return map[string]interface{} {
        "experienceModelId": experienceModelId,
        "name": name,
        "metadata": metadata,
        "defaultExperience": defaultExperience,
        "defaultRankCap": defaultRankCap,
        "maxRankCap": maxRankCap,
        "rankThreshold": rankThreshold,
    }
}

func (p ExperienceModel) Pointer() *ExperienceModel {
    return &p
}

func CastExperienceModels(data []interface{}) []ExperienceModel {
	v := make([]ExperienceModel, 0)
	for _, d := range data {
		v = append(v, NewExperienceModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceModelsFromDict(data []ExperienceModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type ThresholdMaster struct {
	ThresholdId *string `json:"thresholdId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	Values []*int64 `json:"values"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewThresholdMasterFromJson(data string) ThresholdMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewThresholdMasterFromDict(dict)
}

func NewThresholdMasterFromDict(data map[string]interface{}) ThresholdMaster {
    return ThresholdMaster {
        ThresholdId: core.CastString(data["thresholdId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Values: core.CastInt64s(core.CastArray(data["values"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p ThresholdMaster) ToDict() map[string]interface{} {
    
    var thresholdId *string
    if p.ThresholdId != nil {
        thresholdId = p.ThresholdId
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
    var values []interface{}
    if p.Values != nil {
        values = core.CastInt64sFromDict(
            p.Values,
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
        "thresholdId": thresholdId,
        "name": name,
        "description": description,
        "metadata": metadata,
        "values": values,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p ThresholdMaster) Pointer() *ThresholdMaster {
    return &p
}

func CastThresholdMasters(data []interface{}) []ThresholdMaster {
	v := make([]ThresholdMaster, 0)
	for _, d := range data {
		v = append(v, NewThresholdMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastThresholdMastersFromDict(data []ThresholdMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Threshold struct {
	Metadata *string `json:"metadata"`
	Values []*int64 `json:"values"`
}

func NewThresholdFromJson(data string) Threshold {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewThresholdFromDict(dict)
}

func NewThresholdFromDict(data map[string]interface{}) Threshold {
    return Threshold {
        Metadata: core.CastString(data["metadata"]),
        Values: core.CastInt64s(core.CastArray(data["values"])),
    }
}

func (p Threshold) ToDict() map[string]interface{} {
    
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var values []interface{}
    if p.Values != nil {
        values = core.CastInt64sFromDict(
            p.Values,
        )
    }
    return map[string]interface{} {
        "metadata": metadata,
        "values": values,
    }
}

func (p Threshold) Pointer() *Threshold {
    return &p
}

func CastThresholds(data []interface{}) []Threshold {
	v := make([]Threshold, 0)
	for _, d := range data {
		v = append(v, NewThresholdFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastThresholdsFromDict(data []Threshold) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CurrentExperienceMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings *string `json:"settings"`
}

func NewCurrentExperienceMasterFromJson(data string) CurrentExperienceMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentExperienceMasterFromDict(dict)
}

func NewCurrentExperienceMasterFromDict(data map[string]interface{}) CurrentExperienceMaster {
    return CurrentExperienceMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentExperienceMaster) ToDict() map[string]interface{} {
    
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

func (p CurrentExperienceMaster) Pointer() *CurrentExperienceMaster {
    return &p
}

func CastCurrentExperienceMasters(data []interface{}) []CurrentExperienceMaster {
	v := make([]CurrentExperienceMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentExperienceMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentExperienceMastersFromDict(data []CurrentExperienceMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Status struct {
	StatusId *string `json:"statusId"`
	ExperienceName *string `json:"experienceName"`
	UserId *string `json:"userId"`
	PropertyId *string `json:"propertyId"`
	ExperienceValue *int64 `json:"experienceValue"`
	RankValue *int64 `json:"rankValue"`
	RankCapValue *int64 `json:"rankCapValue"`
	NextRankUpExperienceValue *int64 `json:"nextRankUpExperienceValue"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewStatusFromJson(data string) Status {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStatusFromDict(dict)
}

func NewStatusFromDict(data map[string]interface{}) Status {
    return Status {
        StatusId: core.CastString(data["statusId"]),
        ExperienceName: core.CastString(data["experienceName"]),
        UserId: core.CastString(data["userId"]),
        PropertyId: core.CastString(data["propertyId"]),
        ExperienceValue: core.CastInt64(data["experienceValue"]),
        RankValue: core.CastInt64(data["rankValue"]),
        RankCapValue: core.CastInt64(data["rankCapValue"]),
        NextRankUpExperienceValue: core.CastInt64(data["nextRankUpExperienceValue"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Status) ToDict() map[string]interface{} {
    
    var statusId *string
    if p.StatusId != nil {
        statusId = p.StatusId
    }
    var experienceName *string
    if p.ExperienceName != nil {
        experienceName = p.ExperienceName
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var propertyId *string
    if p.PropertyId != nil {
        propertyId = p.PropertyId
    }
    var experienceValue *int64
    if p.ExperienceValue != nil {
        experienceValue = p.ExperienceValue
    }
    var rankValue *int64
    if p.RankValue != nil {
        rankValue = p.RankValue
    }
    var rankCapValue *int64
    if p.RankCapValue != nil {
        rankCapValue = p.RankCapValue
    }
    var nextRankUpExperienceValue *int64
    if p.NextRankUpExperienceValue != nil {
        nextRankUpExperienceValue = p.NextRankUpExperienceValue
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
        "statusId": statusId,
        "experienceName": experienceName,
        "userId": userId,
        "propertyId": propertyId,
        "experienceValue": experienceValue,
        "rankValue": rankValue,
        "rankCapValue": rankCapValue,
        "nextRankUpExperienceValue": nextRankUpExperienceValue,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p Status) Pointer() *Status {
    return &p
}

func CastStatuses(data []interface{}) []Status {
	v := make([]Status, 0)
	for _, d := range data {
		v = append(v, NewStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStatusesFromDict(data []Status) []interface{} {
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