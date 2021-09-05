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

package deploy

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Stack struct {
	StackId *string `json:"stackId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Template *string `json:"template"`
	Status *string `json:"status"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewStackFromJson(data string) Stack {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStackFromDict(dict)
}

func NewStackFromDict(data map[string]interface{}) Stack {
    return Stack {
        StackId: core.CastString(data["stackId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Template: core.CastString(data["template"]),
        Status: core.CastString(data["status"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Stack) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stackId": p.StackId,
        "name": p.Name,
        "description": p.Description,
        "template": p.Template,
        "status": p.Status,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
    }
}

func (p Stack) Pointer() *Stack {
    return &p
}

func CastStacks(data []interface{}) []Stack {
	v := make([]Stack, 0)
	for _, d := range data {
		v = append(v, NewStackFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStacksFromDict(data []Stack) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Resource struct {
	ResourceId *string `json:"resourceId"`
	Type *string `json:"type"`
	Name *string `json:"name"`
	Request *string `json:"request"`
	Response *string `json:"response"`
	RollbackContext *string `json:"rollbackContext"`
	RollbackRequest *string `json:"rollbackRequest"`
	RollbackAfter []string `json:"rollbackAfter"`
	OutputFields []OutputField `json:"outputFields"`
	WorkId *string `json:"workId"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewResourceFromJson(data string) Resource {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewResourceFromDict(dict)
}

func NewResourceFromDict(data map[string]interface{}) Resource {
    return Resource {
        ResourceId: core.CastString(data["resourceId"]),
        Type: core.CastString(data["type"]),
        Name: core.CastString(data["name"]),
        Request: core.CastString(data["request"]),
        Response: core.CastString(data["response"]),
        RollbackContext: core.CastString(data["rollbackContext"]),
        RollbackRequest: core.CastString(data["rollbackRequest"]),
        RollbackAfter: core.CastStrings(core.CastArray(data["rollbackAfter"])),
        OutputFields: CastOutputFields(core.CastArray(data["outputFields"])),
        WorkId: core.CastString(data["workId"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
    }
}

func (p Resource) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "resourceId": p.ResourceId,
        "type": p.Type,
        "name": p.Name,
        "request": p.Request,
        "response": p.Response,
        "rollbackContext": p.RollbackContext,
        "rollbackRequest": p.RollbackRequest,
        "rollbackAfter": core.CastStringsFromDict(
        p.RollbackAfter,
    ),
        "outputFields": CastOutputFieldsFromDict(
        p.OutputFields,
    ),
        "workId": p.WorkId,
        "createdAt": p.CreatedAt,
    }
}

func (p Resource) Pointer() *Resource {
    return &p
}

func CastResources(data []interface{}) []Resource {
	v := make([]Resource, 0)
	for _, d := range data {
		v = append(v, NewResourceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastResourcesFromDict(data []Resource) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Event struct {
	EventId *string `json:"eventId"`
	Name *string `json:"name"`
	ResourceName *string `json:"resourceName"`
	Type *string `json:"type"`
	Message *string `json:"message"`
	EventAt *int64 `json:"eventAt"`
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
        ResourceName: core.CastString(data["resourceName"]),
        Type: core.CastString(data["type"]),
        Message: core.CastString(data["message"]),
        EventAt: core.CastInt64(data["eventAt"]),
    }
}

func (p Event) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "eventId": p.EventId,
        "name": p.Name,
        "resourceName": p.ResourceName,
        "type": p.Type,
        "message": p.Message,
        "eventAt": p.EventAt,
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

type Output struct {
	OutputId *string `json:"outputId"`
	Name *string `json:"name"`
	Value *string `json:"value"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewOutputFromJson(data string) Output {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewOutputFromDict(dict)
}

func NewOutputFromDict(data map[string]interface{}) Output {
    return Output {
        OutputId: core.CastString(data["outputId"]),
        Name: core.CastString(data["name"]),
        Value: core.CastString(data["value"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
    }
}

func (p Output) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "outputId": p.OutputId,
        "name": p.Name,
        "value": p.Value,
        "createdAt": p.CreatedAt,
    }
}

func (p Output) Pointer() *Output {
    return &p
}

func CastOutputs(data []interface{}) []Output {
	v := make([]Output, 0)
	for _, d := range data {
		v = append(v, NewOutputFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastOutputsFromDict(data []Output) []interface{} {
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

type OutputField struct {
	Name *string `json:"name"`
	FieldName *string `json:"fieldName"`
}

func NewOutputFieldFromJson(data string) OutputField {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewOutputFieldFromDict(dict)
}

func NewOutputFieldFromDict(data map[string]interface{}) OutputField {
    return OutputField {
        Name: core.CastString(data["name"]),
        FieldName: core.CastString(data["fieldName"]),
    }
}

func (p OutputField) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "fieldName": p.FieldName,
    }
}

func (p OutputField) Pointer() *OutputField {
    return &p
}

func CastOutputFields(data []interface{}) []OutputField {
	v := make([]OutputField, 0)
	for _, d := range data {
		v = append(v, NewOutputFieldFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastOutputFieldsFromDict(data []OutputField) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}