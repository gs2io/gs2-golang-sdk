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

package schedule

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
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

type EventMaster struct {
	EventId *string `json:"eventId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	ScheduleType *string `json:"scheduleType"`
	RepeatType *string `json:"repeatType"`
	AbsoluteBegin *int64 `json:"absoluteBegin"`
	AbsoluteEnd *int64 `json:"absoluteEnd"`
	RepeatBeginDayOfMonth *int32 `json:"repeatBeginDayOfMonth"`
	RepeatEndDayOfMonth *int32 `json:"repeatEndDayOfMonth"`
	RepeatBeginDayOfWeek *string `json:"repeatBeginDayOfWeek"`
	RepeatEndDayOfWeek *string `json:"repeatEndDayOfWeek"`
	RepeatBeginHour *int32 `json:"repeatBeginHour"`
	RepeatEndHour *int32 `json:"repeatEndHour"`
	RelativeTriggerName *string `json:"relativeTriggerName"`
	RelativeDuration *int32 `json:"relativeDuration"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewEventMasterFromJson(data string) EventMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewEventMasterFromDict(dict)
}

func NewEventMasterFromDict(data map[string]interface{}) EventMaster {
    return EventMaster {
        EventId: core.CastString(data["eventId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        ScheduleType: core.CastString(data["scheduleType"]),
        RepeatType: core.CastString(data["repeatType"]),
        AbsoluteBegin: core.CastInt64(data["absoluteBegin"]),
        AbsoluteEnd: core.CastInt64(data["absoluteEnd"]),
        RepeatBeginDayOfMonth: core.CastInt32(data["repeatBeginDayOfMonth"]),
        RepeatEndDayOfMonth: core.CastInt32(data["repeatEndDayOfMonth"]),
        RepeatBeginDayOfWeek: core.CastString(data["repeatBeginDayOfWeek"]),
        RepeatEndDayOfWeek: core.CastString(data["repeatEndDayOfWeek"]),
        RepeatBeginHour: core.CastInt32(data["repeatBeginHour"]),
        RepeatEndHour: core.CastInt32(data["repeatEndHour"]),
        RelativeTriggerName: core.CastString(data["relativeTriggerName"]),
        RelativeDuration: core.CastInt32(data["relativeDuration"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p EventMaster) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "eventId": p.EventId,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "scheduleType": p.ScheduleType,
        "repeatType": p.RepeatType,
        "absoluteBegin": p.AbsoluteBegin,
        "absoluteEnd": p.AbsoluteEnd,
        "repeatBeginDayOfMonth": p.RepeatBeginDayOfMonth,
        "repeatEndDayOfMonth": p.RepeatEndDayOfMonth,
        "repeatBeginDayOfWeek": p.RepeatBeginDayOfWeek,
        "repeatEndDayOfWeek": p.RepeatEndDayOfWeek,
        "repeatBeginHour": p.RepeatBeginHour,
        "repeatEndHour": p.RepeatEndHour,
        "relativeTriggerName": p.RelativeTriggerName,
        "relativeDuration": p.RelativeDuration,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
    }
}

func (p EventMaster) Pointer() *EventMaster {
    return &p
}

func CastEventMasters(data []interface{}) []EventMaster {
	v := make([]EventMaster, 0)
	for _, d := range data {
		v = append(v, NewEventMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEventMastersFromDict(data []EventMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Trigger struct {
	TriggerId *string `json:"triggerId"`
	Name *string `json:"name"`
	UserId *string `json:"userId"`
	CreatedAt *int64 `json:"createdAt"`
	ExpiresAt *int64 `json:"expiresAt"`
}

func NewTriggerFromJson(data string) Trigger {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewTriggerFromDict(dict)
}

func NewTriggerFromDict(data map[string]interface{}) Trigger {
    return Trigger {
        TriggerId: core.CastString(data["triggerId"]),
        Name: core.CastString(data["name"]),
        UserId: core.CastString(data["userId"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        ExpiresAt: core.CastInt64(data["expiresAt"]),
    }
}

func (p Trigger) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "triggerId": p.TriggerId,
        "name": p.Name,
        "userId": p.UserId,
        "createdAt": p.CreatedAt,
        "expiresAt": p.ExpiresAt,
    }
}

func (p Trigger) Pointer() *Trigger {
    return &p
}

func CastTriggers(data []interface{}) []Trigger {
	v := make([]Trigger, 0)
	for _, d := range data {
		v = append(v, NewTriggerFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTriggersFromDict(data []Trigger) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Event struct {
	EventId *string `json:"eventId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	ScheduleType *string `json:"scheduleType"`
	RepeatType *string `json:"repeatType"`
	AbsoluteBegin *int64 `json:"absoluteBegin"`
	AbsoluteEnd *int64 `json:"absoluteEnd"`
	RepeatBeginDayOfMonth *int32 `json:"repeatBeginDayOfMonth"`
	RepeatEndDayOfMonth *int32 `json:"repeatEndDayOfMonth"`
	RepeatBeginDayOfWeek *string `json:"repeatBeginDayOfWeek"`
	RepeatEndDayOfWeek *string `json:"repeatEndDayOfWeek"`
	RepeatBeginHour *int32 `json:"repeatBeginHour"`
	RepeatEndHour *int32 `json:"repeatEndHour"`
	RelativeTriggerName *string `json:"relativeTriggerName"`
	RelativeDuration *int32 `json:"relativeDuration"`
}

func NewEventFromJson(data string) Event {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewEventFromDict(dict)
}

func NewEventFromDict(data map[string]interface{}) Event {
    return Event {
        EventId: core.CastString(data["eventId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        ScheduleType: core.CastString(data["scheduleType"]),
        RepeatType: core.CastString(data["repeatType"]),
        AbsoluteBegin: core.CastInt64(data["absoluteBegin"]),
        AbsoluteEnd: core.CastInt64(data["absoluteEnd"]),
        RepeatBeginDayOfMonth: core.CastInt32(data["repeatBeginDayOfMonth"]),
        RepeatEndDayOfMonth: core.CastInt32(data["repeatEndDayOfMonth"]),
        RepeatBeginDayOfWeek: core.CastString(data["repeatBeginDayOfWeek"]),
        RepeatEndDayOfWeek: core.CastString(data["repeatEndDayOfWeek"]),
        RepeatBeginHour: core.CastInt32(data["repeatBeginHour"]),
        RepeatEndHour: core.CastInt32(data["repeatEndHour"]),
        RelativeTriggerName: core.CastString(data["relativeTriggerName"]),
        RelativeDuration: core.CastInt32(data["relativeDuration"]),
    }
}

func (p Event) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "eventId": p.EventId,
        "name": p.Name,
        "metadata": p.Metadata,
        "scheduleType": p.ScheduleType,
        "repeatType": p.RepeatType,
        "absoluteBegin": p.AbsoluteBegin,
        "absoluteEnd": p.AbsoluteEnd,
        "repeatBeginDayOfMonth": p.RepeatBeginDayOfMonth,
        "repeatEndDayOfMonth": p.RepeatEndDayOfMonth,
        "repeatBeginDayOfWeek": p.RepeatBeginDayOfWeek,
        "repeatEndDayOfWeek": p.RepeatEndDayOfWeek,
        "repeatBeginHour": p.RepeatBeginHour,
        "repeatEndHour": p.RepeatEndHour,
        "relativeTriggerName": p.RelativeTriggerName,
        "relativeDuration": p.RelativeDuration,
    }
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

type CurrentEventMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings *string `json:"settings"`
}

func NewCurrentEventMasterFromJson(data string) CurrentEventMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentEventMasterFromDict(dict)
}

func NewCurrentEventMasterFromDict(data map[string]interface{}) CurrentEventMaster {
    return CurrentEventMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentEventMaster) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceId": p.NamespaceId,
        "settings": p.Settings,
    }
}

func (p CurrentEventMaster) Pointer() *CurrentEventMaster {
    return &p
}

func CastCurrentEventMasters(data []interface{}) []CurrentEventMaster {
	v := make([]CurrentEventMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentEventMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentEventMastersFromDict(data []CurrentEventMaster) []interface{} {
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