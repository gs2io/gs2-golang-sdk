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
	Revision *int64 `json:"revision"`
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
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Stack) ToDict() map[string]interface{} {
    
    var stackId *string
    if p.StackId != nil {
        stackId = p.StackId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var template *string
    if p.Template != nil {
        template = p.Template
    }
    var status *string
    if p.Status != nil {
        status = p.Status
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
        "stackId": stackId,
        "name": name,
        "description": description,
        "template": template,
        "status": status,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
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
	RollbackAfter []*string `json:"rollbackAfter"`
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
    
    var resourceId *string
    if p.ResourceId != nil {
        resourceId = p.ResourceId
    }
    var _type *string
    if p.Type != nil {
        _type = p.Type
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var request *string
    if p.Request != nil {
        request = p.Request
    }
    var response *string
    if p.Response != nil {
        response = p.Response
    }
    var rollbackContext *string
    if p.RollbackContext != nil {
        rollbackContext = p.RollbackContext
    }
    var rollbackRequest *string
    if p.RollbackRequest != nil {
        rollbackRequest = p.RollbackRequest
    }
    var rollbackAfter []interface{}
    if p.RollbackAfter != nil {
        rollbackAfter = core.CastStringsFromDict(
            p.RollbackAfter,
        )
    }
    var outputFields []interface{}
    if p.OutputFields != nil {
        outputFields = CastOutputFieldsFromDict(
            p.OutputFields,
        )
    }
    var workId *string
    if p.WorkId != nil {
        workId = p.WorkId
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    return map[string]interface{} {
        "resourceId": resourceId,
        "type": _type,
        "name": name,
        "request": request,
        "response": response,
        "rollbackContext": rollbackContext,
        "rollbackRequest": rollbackRequest,
        "rollbackAfter": rollbackAfter,
        "outputFields": outputFields,
        "workId": workId,
        "createdAt": createdAt,
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
	Revision *int64 `json:"revision"`
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
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Event) ToDict() map[string]interface{} {
    
    var eventId *string
    if p.EventId != nil {
        eventId = p.EventId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var resourceName *string
    if p.ResourceName != nil {
        resourceName = p.ResourceName
    }
    var _type *string
    if p.Type != nil {
        _type = p.Type
    }
    var message *string
    if p.Message != nil {
        message = p.Message
    }
    var eventAt *int64
    if p.EventAt != nil {
        eventAt = p.EventAt
    }
    var revision *int64
    if p.Revision != nil {
        revision = p.Revision
    }
    return map[string]interface{} {
        "eventId": eventId,
        "name": name,
        "resourceName": resourceName,
        "type": _type,
        "message": message,
        "eventAt": eventAt,
        "revision": revision,
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
    
    var outputId *string
    if p.OutputId != nil {
        outputId = p.OutputId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var value *string
    if p.Value != nil {
        value = p.Value
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    return map[string]interface{} {
        "outputId": outputId,
        "name": name,
        "value": value,
        "createdAt": createdAt,
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
    
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var fieldName *string
    if p.FieldName != nil {
        fieldName = p.FieldName
    }
    return map[string]interface{} {
        "name": name,
        "fieldName": fieldName,
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